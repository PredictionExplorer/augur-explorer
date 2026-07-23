package autobid

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestDialWithRequestTimeoutBoundsHungEndpoint proves the default dial's
// injected HTTP client bounds every JSON-RPC exchange: a black-holed
// endpoint fails the call at the client timeout with no context deadline
// involved — market reads, bid submissions and receipt polls all ride this
// transport.
func TestDialWithRequestTimeoutBoundsHungEndpoint(t *testing.T) {
	t.Parallel()
	hung := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		// Drain the body so the server notices the aborted client and
		// cancels the request context (otherwise Close would wait forever).
		_, _ = io.Copy(io.Discard, r.Body)
		<-r.Context().Done()
	}))
	t.Cleanup(hung.Close)

	client, err := dialWithRequestTimeout(100*time.Millisecond)(context.Background(), hung.URL)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(client.Close)

	start := time.Now()
	var result string
	// Deliberately context.Background(): the bound must come from the
	// transport, not a caller deadline.
	err = client.CallContext(context.Background(), &result, "eth_chainId")
	if err == nil {
		t.Fatal("call against a hung endpoint returned nil")
	}
	if elapsed := time.Since(start); elapsed > 5*time.Second {
		t.Fatalf("call took %v, want ~100ms (the transport bound)", elapsed)
	}
}

// TestDefaultDialAndTimeoutConstants pins the production wiring: a nil
// Config.Dial selects the bounded dial and the bound is the documented 30s.
func TestDefaultDialAndTimeoutConstants(t *testing.T) {
	t.Parallel()
	if rpcRequestTimeout != 30*time.Second {
		t.Fatalf("rpcRequestTimeout = %v, want the documented 30s", rpcRequestTimeout)
	}
	if dialBoundedRPC == nil {
		t.Fatal("dialBoundedRPC must be initialized (the nil-Dial default)")
	}
}
