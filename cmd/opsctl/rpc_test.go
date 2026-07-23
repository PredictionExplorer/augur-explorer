package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer/logscan"
)

func TestDialBoundedEthClientRejectsMalformedURL(t *testing.T) {
	t.Parallel()
	if _, err := dialBoundedEthClient(context.Background(), "://not-a-url"); err == nil {
		t.Fatal("malformed RPC URL was accepted")
	}
}

func TestDialBoundedEthClientConnects(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))
	}))
	t.Cleanup(srv.Close)

	client, err := dialBoundedEthClient(context.Background(), srv.URL)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(client.Close)
	if _, err := client.ChainID(context.Background()); err != nil {
		t.Fatalf("ChainID through the bounded transport: %v", err)
	}
}

// TestOpsctlRPCTimeoutExceedsLogscanFetchBound pins the layering invariant:
// the transport-level bound must sit above logscan's per-fetch deadline, or
// it would silently truncate the scanner's own budget.
func TestOpsctlRPCTimeoutExceedsLogscanFetchBound(t *testing.T) {
	t.Parallel()
	if rpcRequestTimeout <= logscan.DefaultFetchTimeout {
		t.Fatalf("rpcRequestTimeout %v must exceed logscan.DefaultFetchTimeout %v",
			rpcRequestTimeout, logscan.DefaultFetchTimeout)
	}
}
