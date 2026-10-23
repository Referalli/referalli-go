package referalli

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// WebhookTolerance is how old (or how far ahead) a signed delivery may be.
const WebhookTolerance = 300 * time.Second

// WebhookEvent is a webhook delivery. Data holds the topic's payload: decode
// it with the event's As<Topic> method (AsReferralCreated, …).
type WebhookEvent struct {
	ID        string          `json:"id"`
	Topic     string          `json:"topic"`
	OrgID     string          `json:"orgId"`
	Livemode  bool            `json:"livemode"`
	CreatedAt string          `json:"createdAt"`
	Data      json.RawMessage `json:"data"`
}

// WebhookVerificationError means a delivery failed signature or timestamp
// verification, or isn't a JSON event.
type WebhookVerificationError struct{ Reason string }

func (e *WebhookVerificationError) Error() string { return "referalli: webhook rejected: " + e.Reason }

// WrongTopicError is returned by an As<Topic> method called on an event of another topic.
type WrongTopicError struct{ Want, Got string }

func (e *WrongTopicError) Error() string {
	return "referalli: webhook event is " + strconv.Quote(e.Got) + ", not " + strconv.Quote(e.Want)
}

// VerifyOptions tune webhook verification. The zero value means now and WebhookTolerance.
type VerifyOptions struct {
	Now       time.Time
	Tolerance time.Duration
}

func webhookKey(secret string) []byte {
	if strings.HasPrefix(secret, "whsec_") {
		if b, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(secret, "whsec_")); err == nil {
			return b
		}
	}
	return []byte(secret)
}

func webhookProblem(rawBody []byte, headers http.Header, secrets []string, opts VerifyOptions) string {
	id := headers.Get("webhook-id")
	ts := headers.Get("webhook-timestamp")
	sig := headers.Get("webhook-signature")
	if id == "" || ts == "" || sig == "" {
		return "missing webhook-id, webhook-timestamp or webhook-signature header"
	}
	for _, c := range ts {
		if c < '0' || c > '9' {
			return "webhook-timestamp is not a unix timestamp"
		}
	}
	sec, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return "webhook-timestamp is not a unix timestamp"
	}
	now, tolerance := opts.Now, opts.Tolerance
	if now.IsZero() {
		now = time.Now()
	}
	if tolerance == 0 {
		tolerance = WebhookTolerance
	}
	if d := now.Sub(time.Unix(sec, 0)); d > tolerance || d < -tolerance {
		return "webhook-timestamp is outside the tolerance (replayed or clock skew)"
	}
	var provided []string
	for _, p := range strings.Split(sig, " ") {
		if strings.HasPrefix(p, "v1,") {
			provided = append(provided, p)
		}
	}
	for _, secret := range secrets {
		mac := hmac.New(sha256.New, webhookKey(secret))
		mac.Write([]byte(id + "." + ts + "."))
		mac.Write(rawBody)
		expected := "v1," + base64.StdEncoding.EncodeToString(mac.Sum(nil))
		for _, p := range provided {
			if hmac.Equal([]byte(p), []byte(expected)) {
				return ""
			}
		}
	}
	return "no signature matches the secret"
}

// VerifyWebhook reports whether rawBody (exactly as received) was signed by
// Referalli with one of secrets within WebhookTolerance of now (Standard
// Webhooks; the signature is compared in constant time). Pass two secrets
// while you rotate one.
func VerifyWebhook(rawBody []byte, headers http.Header, secrets []string, now time.Time) bool {
	return webhookProblem(rawBody, headers, secrets, VerifyOptions{Now: now}) == ""
}

// ConstructEvent verifies a delivery and decodes its event. It returns a
// *WebhookVerificationError when the signature, timestamp or body is bad.
//
//	event, err := referalli.ConstructEvent(body, r.Header, []string{secret}, referalli.VerifyOptions{})
//	if event.Topic == referalli.TopicCommissionEarned { data, _ := event.AsCommissionEarned() … }
func ConstructEvent(rawBody []byte, headers http.Header, secrets []string, opts VerifyOptions) (*WebhookEvent, error) {
	if problem := webhookProblem(rawBody, headers, secrets, opts); problem != "" {
		return nil, &WebhookVerificationError{Reason: problem}
	}
	var event WebhookEvent
	if err := json.Unmarshal(rawBody, &event); err != nil {
		return nil, &WebhookVerificationError{Reason: "the body is not JSON: " + err.Error()}
	}
	if event.Topic == "" {
		return nil, &WebhookVerificationError{Reason: "the body is not a webhook event"}
	}
	return &event, nil
}
