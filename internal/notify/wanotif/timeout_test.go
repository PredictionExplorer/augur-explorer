package wanotif

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestDefaultClientIsBounded pins the D22 fallback: a Whatsapp constructed
// without an HTTPClient sends through the package's bounded client — never
// the timeout-less http.DefaultClient that could hang a notification
// forever.
func TestDefaultClientIsBounded(t *testing.T) {
	t.Parallel()
	if defaultHTTPClient.Timeout != defaultRequestTimeout || defaultRequestTimeout != 30*time.Second {
		t.Fatalf("fallback client timeout = %v (const %v), want the documented 30s",
			defaultHTTPClient.Timeout, defaultRequestTimeout)
	}

	// The fallback client must actually carry a nil-HTTPClient send.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"messages":[{"id":"wamid.Z"}]}`))
	}))
	t.Cleanup(srv.Close)
	wa := NewWhatsapp("tok", "1")
	wa.BaseURL = srv.URL
	if wa.HTTPClient != nil {
		t.Fatal("test premise broken: HTTPClient must start nil")
	}
	if _, err := wa.SendText("155", "ping"); err != nil {
		t.Fatalf("SendText through the fallback client: %v", err)
	}
}

// TestBoundedClientCutsHungGraphAPI proves the timeout semantics the
// fallback relies on: a black-holed Graph API fails the send at the client
// timeout without any caller cancellation.
func TestBoundedClientCutsHungGraphAPI(t *testing.T) {
	t.Parallel()
	hung := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		// Drain the body so the server notices the aborted client and
		// cancels the request context (otherwise Close would wait forever).
		_, _ = io.Copy(io.Discard, r.Body)
		<-r.Context().Done()
	}))
	t.Cleanup(hung.Close)

	wa := NewWhatsapp("tok", "1")
	wa.BaseURL = hung.URL
	wa.HTTPClient = &http.Client{Timeout: 100 * time.Millisecond}

	start := time.Now()
	if _, err := wa.SendText("155", "ping"); err == nil {
		t.Fatal("SendText against a hung Graph API returned nil")
	}
	if elapsed := time.Since(start); elapsed > 5*time.Second {
		t.Fatalf("SendText took %v, want ~100ms (the client bound)", elapsed)
	}
}
