# Referalli Go SDK

The official Go client for the [Referalli](https://referalli.com) API: attribution-first affiliate and referral programs.

- One typed method per API operation, with structs for every model and a `<Operation>Params` struct for each operation's query.
- Retries with backoff, idempotency keys, timeouts, `iter.Seq2` auto-paging and one error type per error code.
- Webhook signature verification and typed webhook payloads.
- Standard library only. Go 1.23+.

Docs: <https://referalli.com/docs>. API reference: <https://pkg.go.dev/github.com/referalli/referalli-go>.

## Install

```bash
go get github.com/referalli/referalli-go
```

## Quickstart

Keep the `rk_` API key on your server. (These samples are compiled as examples in `example_test.go`.)

```go
import referalli "github.com/referalli/referalli-go"

ctx := context.Background()
client := referalli.New(os.Getenv("REFERALLI_API_KEY"))

// 1. Create a referral link for a partner
link, err := client.PostLinks(ctx, referalli.PostLinksBody{
	ProgramID:      "prg_123",
	PartnerID:      "ptr_123",
	DestinationURL: "https://example.com/pricing",
})
if err != nil {
	log.Fatal(err)
}
if link.ShortURL != nil {
	fmt.Println(*link.ShortURL)
}

// 2. Attribute a signup to the partner who referred it
_, err = client.PostReferralsAssociate(ctx, referalli.PostReferralsAssociateBody{
	ProgramID:   "prg_123",
	CustomerRef: "user_42",
	AnonID:      referalli.Ptr("anon_from_cookie"),
})

// 3. Track a conversion: a purchase your server has validated
_, err = client.PostMobilePurchases(ctx, referalli.PostMobilePurchasesBody{
	EventID:     "order_1001",
	CustomerRef: "user_42",
	AmountMinor: 4900,
	Currency:    referalli.Ptr("usd"),
})
```

Verify a webhook delivery (`net/http` shown; pass the raw body exactly as received):

```go
secret := os.Getenv("REFERALLI_WEBHOOK_SECRET")
http.HandleFunc("/webhooks/referalli", func(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "read error", http.StatusBadRequest)
		return
	}
	event, err := referalli.ConstructEvent(body, r.Header, []string{secret}, referalli.VerifyOptions{})
	if err != nil {
		http.Error(w, "bad signature", http.StatusBadRequest)
		return
	}
	if event.Topic == referalli.TopicCommissionEarned {
		data, _ := event.AsCommissionEarned()
		log.Printf("commission %v", data)
	}
	w.WriteHeader(http.StatusNoContent)
})
```

## Configuration

```go
client := referalli.New("rk_live_...",
	referalli.WithBaseURL("https://referalli.com"), // the default
	referalli.WithTimeout(30*time.Second),          // per attempt
	referalli.WithMaxRetries(3),                    // retries for idempotent requests
	referalli.WithMode("live"),                     // optional: assert the key's mode
)
fast := client.WithOptions(referalli.WithTimeout(5*time.Second), referalli.WithMaxRetries(0))
```

Per call: `referalli.WithIdempotencyKey(key)`, `referalli.WithRequestTimeout(d)` and `referalli.WithHeader(name, value)`. Cancel a call with its `context.Context`. Requests are pinned to `referalli.APIVersion`; `WithAPIVersion` pins another. `WithHTTPClient` supplies your own `*http.Client` (transport, proxies).

### Test mode

Use a test-mode key (`rk_test_...`) to work against your organization's test data. `WithMode("test")` makes every request assert it: the API rejects a request whose key is in the other mode with `*referalli.ModeMismatchError`. Webhook events carry `Livemode`.

## Errors

Every API error code has its own type (`*ValidationError`, `*NotFoundError`, `*RateLimitedError`, ...) wrapping `*APIError` (`Code`, `Status`, `Message`, `RequestID`, `DocsURL`, `Details`, `RetryAfter`); match with `errors.As`. Network failures and timeouts return `*ConnectionError` after the retries run out.

```go
_, err := client.GetPartnersID(ctx, "ptr_missing")
var notFound *referalli.NotFoundError
var limited *referalli.RateLimitedError
switch {
case errors.As(err, &notFound):
	fmt.Println(notFound.Code, notFound.Status, notFound.RequestID, notFound.DocsURL)
case errors.As(err, &limited):
	fmt.Println("retry after", limited.RetryAfter)
case err != nil:
	if apiErr, ok := referalli.AsAPIError(err); ok {
		fmt.Println(apiErr.Code, apiErr.Message)
	}
}
```

Retries: GET, PUT and DELETE requests, and writes carrying an `Idempotency-Key`, are retried on 429, 5xx and network errors with exponential backoff and jitter, honoring `Retry-After`. POST and PATCH requests always carry a key (generated unless you pass `WithIdempotencyKey`), reused across retries, so a retried write is applied once.

## Pagination

List methods return one `*Page[T]`; their `Each` twin returns an `iter.Seq2[T, error]` that follows the cursor.

```go
page, err := client.GetPartners(ctx, &referalli.GetPartnersParams{Status: referalli.Ptr("active"), Limit: referalli.Ptr[int64](50)})
if err != nil {
	log.Fatal(err)
}
_, _ = client.GetPartners(ctx, &referalli.GetPartnersParams{Status: referalli.Ptr("active"), StartingAfter: page.NextCursor})

for referral, err := range client.GetReferralsEach(ctx, nil) {
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(referral.ID)
}
```

## Webhooks

`referalli.ConstructEvent(body, headers, secrets, opts)` checks the Standard Webhooks signature (`webhook-id`, `webhook-timestamp`, `webhook-signature`) in constant time, rejects deliveries more than 5 minutes old or ahead (`VerifyOptions.Tolerance`), and returns a `*WebhookEvent`. Decode its payload with the topic's method (`AsReferralCreated`, `AsCommissionEarned`, ...; the topics are the `Topic...` constants and `referalli.Topics`). Pass two secrets while you rotate one. `referalli.VerifyWebhook(body, headers, secrets, time.Now())` returns a bool instead.

## Any operation by id

```go
partner, err := client.Call(ctx, "getPartnersId", map[string]string{"id": "ptr_123"}, nil, nil)
for item, err := range client.ListAll(ctx, "getPartners", nil, map[string]any{"status": "active"}) {
	if err != nil {
		break
	}
	fmt.Println(string(item))
}
```

## Versioning

This module follows semantic versioning; see [CHANGELOG.md](./CHANGELOG.md).

## License

MIT. See [LICENSE](./LICENSE).
