package srvmonitor

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestRPCMonitorProbeBoundsHungNode proves a black-holed RPC endpoint —
// reachable TCP, no HTTP answer — shows as DOWN within the probe bound
// instead of parking the probe goroutine forever. The other monitors carry
// the same 10-second bound; this was the one unbounded probe.
func TestRPCMonitorProbeBoundsHungNode(t *testing.T) {
	t.Parallel()
	hung := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		// Drain the body so the server notices the aborted client and
		// cancels the request context (otherwise Close would wait forever).
		_, _ = io.Copy(io.Discard, r.Body)
		<-r.Context().Done()
	}))
	t.Cleanup(hung.Close)

	nodes := []RPCConfig{{Name: "hung", URL: hung.URL, ChainID: "1"}}
	m := NewRPCMonitor(nodes, nil, NewSharedRPCState(), testIntervals())
	m.blockWait = func(context.Context) {}
	m.probeTimeout = 100 * time.Millisecond

	start := time.Now()
	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)
	elapsed := time.Since(start)

	st := m.statuses[0]
	if st.Alive || st.ErrStr == "" {
		t.Fatalf("status = %+v, want DOWN with a timeout error", st)
	}
	if elapsed > 5*time.Second {
		t.Fatalf("check took %v against a hung node, want ~100ms (the probe bound)", elapsed)
	}
	if msgs := drain(errCh); len(msgs) != 1 {
		t.Fatalf("errors = %v, want the one probe failure", msgs)
	}
}

// TestRPCMonitorDefaultProbeTimeout pins the production bound to the shared
// 10-second probe convention.
func TestRPCMonitorDefaultProbeTimeout(t *testing.T) {
	t.Parallel()
	m := NewRPCMonitor(nil, nil, NewSharedRPCState(), testIntervals())
	if m.probeTimeout != rpcProbeTimeout || rpcProbeTimeout != 10*time.Second {
		t.Fatalf("probeTimeout = %v (const %v), want the documented 10s", m.probeTimeout, rpcProbeTimeout)
	}
}
