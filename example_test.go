package referalli_test

// The README's samples, compiled by `go vet` and `go test` (they have no
// Output comment, so they are not run).

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	referalli "github.com/referalli/referalli-go"
)

func Example_quickstart() {
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
	if err != nil {
		log.Fatal(err)
	}

	// 3. Track a conversion: a purchase your server has validated
	_, err = client.PostMobilePurchases(ctx, referalli.PostMobilePurchasesBody{
		EventID:     "order_1001",
		CustomerRef: "user_42",
		AmountMinor: 4900,
		Currency:    referalli.Ptr("usd"),
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Example_webhook() {
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
}

func Example_configuration() {
	client := referalli.New("rk_live_...",
		referalli.WithBaseURL("https://referalli.com"), // the default
		referalli.WithTimeout(30*time.Second),          // per attempt
		referalli.WithMaxRetries(3),                    // retries for idempotent requests
		referalli.WithMode("live"),                     // optional: assert the key's mode
	)
	fast := client.WithOptions(referalli.WithTimeout(5*time.Second), referalli.WithMaxRetries(0))
	_ = fast
}

func Example_errors() {
	ctx := context.Background()
	client := referalli.New("rk_live_...")
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
}

func Example_pagination() {
	ctx := context.Background()
	client := referalli.New("rk_live_...")
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
}

func Example_anyOperation() {
	ctx := context.Background()
	client := referalli.New("rk_live_...")
	partner, err := client.Call(ctx, "getPartnersId", map[string]string{"id": "ptr_123"}, nil, nil)
	fmt.Println(partner, err)
	for item, err := range client.ListAll(ctx, "getPartners", nil, map[string]any{"status": "active"}) {
		if err != nil {
			break
		}
		fmt.Println(string(item))
	}
}
