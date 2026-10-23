package referalli

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
	"time"
)

func testServer(t *testing.T, handler func(n int, w http.ResponseWriter, r *http.Request)) (*Client, *[]*http.Request) {
	t.Helper()
	var mu sync.Mutex
	var seen []*http.Request
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		seen = append(seen, r.Clone(context.Background()))
		n := len(seen)
		mu.Unlock()
		handler(n, w, r)
	}))
	t.Cleanup(srv.Close)
	return New("rk_test_1", WithBaseURL(srv.URL), WithRetryBaseDelay(time.Millisecond), WithMode("test")), &seen
}

func TestPostRetriesWithTheSameKeyAndSendsHeaders(t *testing.T) {
	client, seen := testServer(t, func(n int, w http.ResponseWriter, _ *http.Request) {
		if n == 1 {
			w.Header().Set("retry-after", "0")
			w.WriteHeader(429)
			_, _ = w.Write([]byte(`{"error":{"code":"rate_limited","message":"slow"}}`))
			return
		}
		_, _ = w.Write([]byte(`{"id":"ptr_1","email":"a@example.com","displayName":"","externalRef":null,"status":"active","tags":[],"notes":null,"createdAt":"2026-10-21T12:00:00Z"}`))
	})
	p, err := client.PostPartners(context.Background(), PostPartnersBody{Email: "a@example.com"})
	if err != nil || p.ID != "ptr_1" {
		t.Fatalf("got %v, %v", p, err)
	}
	if len(*seen) != 2 {
		t.Fatalf("want 2 requests, got %d", len(*seen))
	}
	first, second := (*seen)[0].Header, (*seen)[1].Header
	if first.Get("idempotency-key") == "" || first.Get("idempotency-key") != second.Get("idempotency-key") {
		t.Fatalf("idempotency key not reused: %q vs %q", first.Get("idempotency-key"), second.Get("idempotency-key"))
	}
	if first.Get("referalli-mode") != "test" || first.Get("user-agent") != "referalli-go/"+Version || first.Get("referalli-version") != APIVersion {
		t.Fatalf("headers: %v", first)
	}
}

func TestTypedErrorsAndTimeouts(t *testing.T) {
	client, seen := testServer(t, func(_ int, w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/partners/slow" {
			time.Sleep(200 * time.Millisecond)
		}
		w.WriteHeader(404)
		_, _ = w.Write([]byte(`{"error":{"code":"not_found","message":"nope","requestId":"req_1"}}`))
	})
	_, err := client.GetPartnersID(context.Background(), "ptr_x")
	var nf *NotFoundError
	if !errors.As(err, &nf) || nf.RequestID != "req_1" || len(*seen) != 1 {
		t.Fatalf("want NotFoundError after 1 request, got %v (%d requests)", err, len(*seen))
	}
	_, err = client.WithOptions(WithTimeout(20*time.Millisecond), WithMaxRetries(0)).GetPartnersID(context.Background(), "slow")
	var ce *ConnectionError
	if !errors.As(err, &ce) {
		t.Fatalf("want ConnectionError on timeout, got %v", err)
	}
}

func TestEachFollowsCursor(t *testing.T) {
	client, seen := testServer(t, func(n int, w http.ResponseWriter, _ *http.Request) {
		if n == 1 {
			_, _ = w.Write([]byte(`{"object":"list","data":[{"id":"ptr_1"}],"has_more":true,"next_cursor":"ptr_1"}`))
			return
		}
		_, _ = w.Write([]byte(`{"object":"list","data":[{"id":"ptr_2"}],"has_more":false,"next_cursor":null}`))
	})
	var ids []string
	for p, err := range client.GetPartnersEach(context.Background(), &GetPartnersParams{Limit: Ptr[int64](1)}) {
		if err != nil {
			t.Fatal(err)
		}
		ids = append(ids, p.ID)
	}
	if len(ids) != 2 || ids[1] != "ptr_2" || (*seen)[1].URL.Query().Get("starting_after") != "ptr_1" {
		t.Fatalf("ids %v", ids)
	}
}

func TestConstructEvent(t *testing.T) {
	secret := "whsec_" + base64.StdEncoding.EncodeToString([]byte("0123456789abcdef0123456789abcdef"))
	body := []byte(`{"id":"evt_1","topic":"referral.created","orgId":"org_1","livemode":false,"createdAt":"2026-10-21T12:00:00Z","data":{"partnerId":"ptr_1"}}`)
	sign := func(ts int64) http.Header {
		mac := hmac.New(sha256.New, webhookKey(secret))
		mac.Write([]byte("msg_1." + strconv.FormatInt(ts, 10) + "."))
		mac.Write(body)
		h := http.Header{}
		h.Set("webhook-id", "msg_1")
		h.Set("webhook-timestamp", strconv.FormatInt(ts, 10))
		h.Set("webhook-signature", "v1,"+base64.StdEncoding.EncodeToString(mac.Sum(nil)))
		return h
	}
	event, err := ConstructEvent(body, sign(time.Now().Unix()), []string{secret}, VerifyOptions{})
	if err != nil || event.Topic != TopicReferralCreated || event.OrgID != "org_1" {
		t.Fatalf("got %+v, %v", event, err)
	}
	if _, err := event.AsReferralCreated(); err != nil {
		t.Fatal(err)
	}
	var wrong *WrongTopicError
	if _, err := event.AsCommissionEarned(); !errors.As(err, &wrong) {
		t.Fatalf("want WrongTopicError, got %v", err)
	}
	var bad *WebhookVerificationError
	if _, err := ConstructEvent(body, sign(time.Now().Unix()-301), []string{secret}, VerifyOptions{}); !errors.As(err, &bad) {
		t.Fatalf("stale delivery accepted: %v", err)
	}
	if _, err := ConstructEvent(append(body, ' '), sign(time.Now().Unix()), []string{secret}, VerifyOptions{}); !errors.As(err, &bad) {
		t.Fatalf("tampered delivery accepted: %v", err)
	}
}
