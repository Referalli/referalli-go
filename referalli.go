// Package referalli is the official Go SDK for the Referalli API (standard library only).
//
//	client := referalli.New(os.Getenv("REFERALLI_API_KEY"))
//	link, err := client.PostLinks(ctx, referalli.PostLinksBody{ProgramID: "prg_123", PartnerID: "ptr_123", DestinationURL: "https://example.com"})
//	for partner, err := range client.GetPartnersEach(ctx, &referalli.GetPartnersParams{Status: referalli.Ptr("active")}) {
//		…
//	}
//
// Every operation has a typed method named after its operationId. Writes
// carry an Idempotency-Key (generated unless WithIdempotencyKey is passed)
// that is reused across retries. Idempotent requests are retried on 429, 5xx
// and network errors with exponential backoff and jitter, honoring
// Retry-After. Requests are pinned to APIVersion.
package referalli

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iter"
	"math"
	mrand "math/rand/v2"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// DefaultBaseURL is the hosted Referalli API.
const DefaultBaseURL = "https://referalli.com"

// DefaultTimeout bounds each attempt of a request.
const DefaultTimeout = 30 * time.Second

const maxRetryDelay = 30 * time.Second

// Client talks to the Referalli API with an rk_… API key. Keep it on your server.
type Client struct {
	APIKey         string
	BaseURL        string
	HTTP           *http.Client
	Timeout        time.Duration
	MaxRetries     int
	RetryBaseDelay time.Duration
	APIVersion     string
	// Mode, when "test" or "live", is asserted on every request: the API
	// refuses a request whose key is in the other mode.
	Mode    string
	Headers http.Header
}

// Option configures a Client.
type Option func(*Client)

// WithBaseURL points the client at another host (no /v1 suffix).
func WithBaseURL(u string) Option { return func(c *Client) { c.BaseURL = u } }

// WithHTTPClient sets the HTTP client (transport, proxies). Its own Timeout still applies.
func WithHTTPClient(h *http.Client) Option { return func(c *Client) { c.HTTP = h } }

// WithTimeout bounds each attempt of a request (default 30s).
func WithTimeout(d time.Duration) Option { return func(c *Client) { c.Timeout = d } }

// WithMaxRetries sets how many times idempotent requests are retried (default 3).
func WithMaxRetries(n int) Option { return func(c *Client) { c.MaxRetries = n } }

// WithRetryBaseDelay sets the first backoff delay (default 500ms), doubled per retry.
func WithRetryBaseDelay(d time.Duration) Option { return func(c *Client) { c.RetryBaseDelay = d } }

// WithAPIVersion pins another API version.
func WithAPIVersion(v string) Option { return func(c *Client) { c.APIVersion = v } }

// WithMode asserts the key's mode ("test" or "live") on every request.
func WithMode(mode string) Option { return func(c *Client) { c.Mode = mode } }

// New builds a client for the hosted API; pass WithBaseURL for another host.
func New(apiKey string, opts ...Option) *Client {
	c := &Client{
		APIKey:         apiKey,
		BaseURL:        DefaultBaseURL,
		HTTP:           &http.Client{},
		Timeout:        DefaultTimeout,
		MaxRetries:     3,
		RetryBaseDelay: 500 * time.Millisecond,
		APIVersion:     APIVersion,
		Headers:        http.Header{},
	}
	for _, o := range opts {
		o(c)
	}
	if c.BaseURL == "" {
		c.BaseURL = DefaultBaseURL
	}
	return c
}

// WithOptions returns a copy of the client with some settings changed.
func (c *Client) WithOptions(opts ...Option) *Client {
	cp := *c
	cp.Headers = c.Headers.Clone()
	for _, o := range opts {
		o(&cp)
	}
	return &cp
}

// RequestOption changes one call.
type RequestOption func(*requestOptions)

type requestOptions struct {
	idempotencyKey string
	headers        http.Header
	timeout        time.Duration
}

// WithIdempotencyKey sends this Idempotency-Key instead of a generated one.
func WithIdempotencyKey(key string) RequestOption {
	return func(o *requestOptions) { o.idempotencyKey = key }
}

// WithRequestTimeout bounds each attempt of this call.
func WithRequestTimeout(d time.Duration) RequestOption {
	return func(o *requestOptions) { o.timeout = d }
}

// WithHeader adds a header to one call.
func WithHeader(name, value string) RequestOption {
	return func(o *requestOptions) {
		if o.headers == nil {
			o.headers = http.Header{}
		}
		o.headers.Set(name, value)
	}
}

// Ptr returns a pointer to v, for optional fields and params.
func Ptr[T any](v T) *T { return &v }

// Page is one page of a list: pass NextCursor as StartingAfter for the next.
type Page[T any] struct {
	Object     string  `json:"object"`
	Data       []T     `json:"data"`
	HasMore    bool    `json:"has_more"`
	NextCursor *string `json:"next_cursor"`
}

// APIError is an error the API answered with. Code is stable; DocsURL explains it.
// Each documented code also has its own type (NotFoundError, …): match with errors.As.
type APIError struct {
	Status     int
	Code       string
	Message    string
	RequestID  string
	DocsURL    string
	Details    json.RawMessage
	RetryAfter *time.Duration
	Header     http.Header
}

func (e *APIError) Error() string { return fmt.Sprintf("%s (%d): %s", e.Code, e.Status, e.Message) }

// ConnectionError means the API could not be reached (network error or timeout), after retries.
type ConnectionError struct{ Err error }

func (e *ConnectionError) Error() string {
	return "referalli: could not reach the API: " + e.Err.Error()
}
func (e *ConnectionError) Unwrap() error { return e.Err }

var retryStatuses = map[int]bool{429: true, 500: true, 502: true, 503: true, 504: true}

// Methods that are safe to retry without an idempotency key.
var idempotentMethods = map[string]bool{http.MethodGet: true, http.MethodHead: true, http.MethodPut: true, http.MethodDelete: true, http.MethodOptions: true}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any, out any, kind string, opts []RequestOption) error {
	ro := requestOptions{}
	for _, o := range opts {
		o(&ro)
	}
	u := c.BaseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	var payload []byte
	if body != nil {
		var err error
		if payload, err = json.Marshal(body); err != nil {
			return err
		}
	}
	key := ro.idempotencyKey
	if key == "" && (method == http.MethodPost || method == http.MethodPatch) {
		key = newUUID()
	}
	// Only idempotent requests are retried: safe methods, or writes that carry an idempotency key.
	retryable := idempotentMethods[method] || key != ""
	timeout := c.Timeout
	if ro.timeout > 0 {
		timeout = ro.timeout
	}
	for attempt := 0; ; attempt++ {
		res, data, err := c.attempt(ctx, method, u, payload, key, ro.headers, timeout)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			if retryable && attempt < c.MaxRetries {
				if err := sleep(ctx, c.delay(attempt, nil)); err != nil {
					return err
				}
				continue
			}
			return &ConnectionError{Err: err}
		}
		if res.StatusCode >= 200 && res.StatusCode < 300 {
			switch kind {
			case "text":
				*(out.(*string)) = string(data)
				return nil
			case "empty":
				return nil
			}
			if out == nil || len(data) == 0 {
				return nil
			}
			return json.Unmarshal(data, out)
		}
		retryAfter := parseRetryAfter(res.Header.Get("retry-after"))
		if retryStatuses[res.StatusCode] && retryable && attempt < c.MaxRetries {
			if err := sleep(ctx, c.delay(attempt, retryAfter)); err != nil {
				return err
			}
			continue
		}
		return apiError(res, data, retryAfter)
	}
}

// attempt sends one request and reads its whole body within the timeout.
func (c *Client) attempt(ctx context.Context, method, u string, payload []byte, key string, extra http.Header, timeout time.Duration) (*http.Response, []byte, error) {
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	var reader io.Reader
	if payload != nil {
		reader = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(ctx, method, u, reader)
	if err != nil {
		return nil, nil, err
	}
	for k, v := range c.Headers {
		req.Header[k] = v
	}
	req.Header.Set("accept", "application/json")
	// Wire protocol: the API reads these header names.
	req.Header.Set("referalli-version", c.APIVersion)
	if c.Mode != "" {
		req.Header.Set("referalli-mode", c.Mode)
	}
	req.Header.Set("user-agent", "referalli-go/"+Version)
	if c.APIKey != "" {
		req.Header.Set("authorization", "Bearer "+c.APIKey)
	}
	if payload != nil {
		req.Header.Set("content-type", "application/json")
	}
	if key != "" {
		req.Header.Set("idempotency-key", key)
	}
	for k, v := range extra {
		req.Header[k] = v
	}
	httpClient := c.HTTP
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	return res, data, nil
}

// delay is exponential backoff with jitter; Retry-After wins when the server sends it.
func (c *Client) delay(attempt int, retryAfter *time.Duration) time.Duration {
	if retryAfter != nil {
		return min(max(*retryAfter, 0), maxRetryDelay)
	}
	d := float64(c.RetryBaseDelay) * math.Pow(2, float64(attempt))
	d = math.Min(d, float64(maxRetryDelay))
	return time.Duration(d * (0.5 + mrand.Float64()*0.5))
}

func sleep(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}

func parseRetryAfter(v string) *time.Duration {
	if v == "" {
		return nil
	}
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		d := time.Duration(s * float64(time.Second))
		return &d
	}
	if at, err := http.ParseTime(v); err == nil {
		d := max(time.Until(at), 0)
		return &d
	}
	return nil
}

func apiError(res *http.Response, data []byte, retryAfter *time.Duration) error {
	var envelope struct {
		Error *struct {
			Code      string          `json:"code"`
			Message   string          `json:"message"`
			RequestID string          `json:"requestId"`
			DocsURL   string          `json:"docs_url"`
			Details   json.RawMessage `json:"details"`
		} `json:"error"`
	}
	base := &APIError{Status: res.StatusCode, Code: "http_error", Message: string(data), RetryAfter: retryAfter, Header: res.Header}
	if json.Unmarshal(data, &envelope) == nil && envelope.Error != nil {
		base.Code, base.Message, base.RequestID, base.DocsURL, base.Details = envelope.Error.Code, envelope.Error.Message, envelope.Error.RequestID, envelope.Error.DocsURL, envelope.Error.Details
	}
	if base.RequestID == "" {
		base.RequestID = res.Header.Get("x-request-id")
	}
	return typedError(base)
}

// AsAPIError returns the APIError inside err, if any.
func AsAPIError(err error) (*APIError, bool) {
	var e *APIError
	ok := errors.As(err, &e)
	return e, ok
}

func paginate[T any](fetch func(cursor *string) (*Page[T], error)) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var cursor *string
		for {
			page, err := fetch(cursor)
			if err != nil {
				var zero T
				yield(zero, err)
				return
			}
			for _, item := range page.Data {
				if !yield(item, nil) {
					return
				}
			}
			if !page.HasMore || page.NextCursor == nil || *page.NextCursor == "" {
				return
			}
			cursor = page.NextCursor
		}
	}
}

func setQuery[T any](v url.Values, name string, value *T) {
	if value == nil {
		return
	}
	switch x := any(*value).(type) {
	case string:
		v.Set(name, x)
	case bool:
		v.Set(name, strconv.FormatBool(x))
	case int64:
		v.Set(name, strconv.FormatInt(x, 10))
	case float64:
		v.Set(name, strconv.FormatFloat(x, 'f', -1, 64))
	default:
		v.Set(name, fmt.Sprint(x))
	}
}

// nilable turns a nil optional body into an untyped nil (no request body).
func nilable[T any](body *T) any {
	if body == nil {
		return nil
	}
	return body
}

func decodeInto(raw json.RawMessage, out any) error {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	return json.Unmarshal(raw, out)
}

// decodeQuery fills a Params struct from wire-named values; numbers given as strings are accepted.
func decodeQuery(q map[string]any, out any) error {
	if len(q) == 0 {
		return nil
	}
	raw, err := json.Marshal(q)
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, out)
}

func newUUID() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// ListAll yields every item of a list operation (by operation id), following
// next_cursor. Items are raw JSON; the typed <Operation>Each methods decode them.
func (c *Client) ListAll(ctx context.Context, operationID string, path map[string]string, query map[string]any, opts ...RequestOption) iter.Seq2[json.RawMessage, error] {
	return func(yield func(json.RawMessage, error) bool) {
		q := map[string]any{}
		for k, v := range query {
			q[k] = v
		}
		for {
			res, err := c.Call(ctx, operationID, path, q, nil, opts...)
			if err != nil {
				yield(nil, err)
				return
			}
			raw, err := json.Marshal(res)
			if err != nil {
				yield(nil, err)
				return
			}
			var page Page[json.RawMessage]
			if err := json.Unmarshal(raw, &page); err != nil {
				yield(nil, err)
				return
			}
			for _, item := range page.Data {
				if !yield(item, nil) {
					return
				}
			}
			if !page.HasMore || page.NextCursor == nil || *page.NextCursor == "" {
				return
			}
			q["starting_after"] = *page.NextCursor
		}
	}
}
