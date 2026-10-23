package referalli

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"
)

// TestVerifyWebhookVectors checks VerifyWebhook against the vectors every
// Referalli SDK shares (REFERALLI_WEBHOOK_VECTORS overrides the path).
func TestVerifyWebhookVectors(t *testing.T) {
	path := os.Getenv("REFERALLI_WEBHOOK_VECTORS")
	if path == "" {
		path = "testdata/vectors.json"
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var file struct {
		Cases []struct {
			Name    string            `json:"name"`
			Body    string            `json:"body"`
			Headers map[string]string `json:"headers"`
			Secrets []string          `json:"secrets"`
			Now     int64             `json:"now"`
			Valid   bool              `json:"valid"`
		} `json:"cases"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	for _, c := range file.Cases {
		h := http.Header{}
		for k, v := range c.Headers {
			h.Set(k, v)
		}
		if got := VerifyWebhook([]byte(c.Body), h, c.Secrets, time.Unix(c.Now, 0)); got != c.Valid {
			t.Errorf("%s: got %v, want %v", c.Name, got, c.Valid)
		}
	}
}
