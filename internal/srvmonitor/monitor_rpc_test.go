package srvmonitor

import (
	"context"
	"math"
	"strings"
	"sync"
	"testing"

	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
)

// rpcTestSetup builds an RPC monitor over two fake chains: an official node
// and a lagging follower, both on chain id 1. The blockWait hook advances
// the chains so the head moves between the two reads of a check.
func TestRPCMonitorAdvancingAndLag(t *testing.T) {
	t.Parallel()
	official := testchain.New(t)
	official.EnsureBlock(100)
	follower := testchain.New(t)
	follower.EnsureBlock(95)

	nodes := []RPCConfig{
		{Name: "official-main", URL: official.URL(), ChainID: "1", IsOfficial: true},
		{Name: "follower", URL: follower.URL(), ChainID: "1"},
	}
	shared := NewSharedRPCState()
	m := NewRPCMonitor(nodes, map[string]string{"mainnet": "official-main"}, shared, testIntervals())
	// Both nodes are checked concurrently and share this hook. The barrier
	// makes sure every node finished its first read before either chain
	// advances — otherwise one goroutine could advance the official chain
	// before the other's first read, which then observes a zero block
	// difference (a real flake this test had under full-suite load).
	var firstReadsDone sync.WaitGroup
	firstReadsDone.Add(len(nodes))
	m.blockWait = func(context.Context) {
		firstReadsDone.Done()
		firstReadsDone.Wait()
		official.EnsureBlock(110)
		follower.EnsureBlock(97)
	}

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unexpected errors: %v", msgs)
	}
	officialSt, followerSt := m.statuses[0], m.statuses[1]
	if !officialSt.Alive || officialSt.LastBlockNum != 110 {
		t.Fatalf("official status = %+v", officialSt)
	}
	if !followerSt.Alive || followerSt.LastBlockNum != 97 {
		t.Fatalf("follower status = %+v", followerSt)
	}
	if followerSt.OfficialLagDiff != 13 {
		t.Fatalf("lag = %d, want 13", followerSt.OfficialLagDiff)
	}
	if got := shared.Official("1"); got != 110 {
		t.Fatalf("shared official block = %d, want 110", got)
	}

	// Official row shows N/A for lag; follower shows the number.
	if row := disp.Row(1); !strings.Contains(row, "Alive") || !strings.Contains(row, "N/A") || !strings.Contains(row, "official-main") {
		t.Fatalf("official row = %q", row)
	}
	if row := disp.Row(2); !strings.Contains(row, "13") || !strings.Contains(row, "follower") {
		t.Fatalf("follower row = %q", row)
	}
	if header := disp.Row(0); !strings.Contains(header, "RPC Nodes") {
		t.Fatalf("header = %q", header)
	}
}

func TestRPCMonitorStalledChain(t *testing.T) {
	t.Parallel()
	chain := testchain.New(t)
	chain.EnsureBlock(50)

	nodes := []RPCConfig{{Name: "n", URL: chain.URL(), ChainID: "1"}}
	// The default blockWait (a 1ms sleep from testIntervals) is kept: the
	// chain does not advance during it.
	m := NewRPCMonitor(nodes, nil, NewSharedRPCState(), testIntervals())

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	st := m.statuses[0]
	if st.Alive {
		t.Fatal("stalled chain must not be alive")
	}
	if want := "Block difference is zero (last block = 50)"; st.ErrStr != want {
		t.Fatalf("ErrStr = %q, want %q", st.ErrStr, want)
	}
	if msgs := drain(errCh); len(msgs) != 1 || msgs[0] != st.ErrStr {
		t.Fatalf("errors = %v", msgs)
	}
	if row := disp.Row(1); !strings.Contains(row, "DOWN") {
		t.Fatalf("row = %q", row)
	}
}

func TestRPCMonitorFailures(t *testing.T) {
	t.Parallel()

	t.Run("unreachable node", func(t *testing.T) {
		t.Parallel()
		nodes := []RPCConfig{{Name: "gone", URL: "http://127.0.0.1:1", ChainID: "1"}}
		m := NewRPCMonitor(nodes, nil, NewSharedRPCState(), testIntervals())
		m.blockWait = func(context.Context) {}

		errCh := make(chan string, 10)
		m.check(context.Background(), newFakeDisplay(), errCh)

		st := m.statuses[0]
		if st.Alive || st.ErrStr == "" {
			t.Fatalf("status = %+v, want unreachable error", st)
		}
		if msgs := drain(errCh); len(msgs) != 1 {
			t.Fatalf("errors = %v", msgs)
		}
	})

	t.Run("malformed url fails at dial", func(t *testing.T) {
		t.Parallel()
		nodes := []RPCConfig{{Name: "bad", URL: "://not-a-url", ChainID: "1"}}
		m := NewRPCMonitor(nodes, nil, NewSharedRPCState(), testIntervals())
		m.blockWait = func(context.Context) {}

		errCh := make(chan string, 10)
		m.check(context.Background(), newFakeDisplay(), errCh)

		st := m.statuses[0]
		if st.Alive || st.ErrStr == "" {
			t.Fatalf("status = %+v, want dial error", st)
		}
		if msgs := drain(errCh); len(msgs) != 1 {
			t.Fatalf("errors = %v", msgs)
		}
	})

	t.Run("first read fails", func(t *testing.T) {
		t.Parallel()
		chain := testchain.New(t)
		chain.EnsureBlock(10)
		chain.FailNextRPC("eth_getBlockByNumber", "read failed")

		nodes := []RPCConfig{{Name: "n", URL: chain.URL(), ChainID: "1"}}
		m := NewRPCMonitor(nodes, nil, NewSharedRPCState(), testIntervals())
		m.blockWait = func(context.Context) {}

		errCh := make(chan string, 10)
		m.check(context.Background(), newFakeDisplay(), errCh)

		st := m.statuses[0]
		if st.Alive || !strings.Contains(st.ErrStr, "read failed") {
			t.Fatalf("status = %+v", st)
		}
	})

	t.Run("second read fails keeps old block", func(t *testing.T) {
		t.Parallel()
		chain := testchain.New(t)
		chain.EnsureBlock(10)

		nodes := []RPCConfig{{Name: "n", URL: chain.URL(), ChainID: "1"}}
		m := NewRPCMonitor(nodes, nil, NewSharedRPCState(), testIntervals())
		m.blockWait = func(context.Context) { chain.EnsureBlock(11) }

		// A successful cycle records block 11.
		m.check(context.Background(), newFakeDisplay(), make(chan string, 10))
		if m.statuses[0].LastBlockNum != 11 {
			t.Fatalf("warm-up block = %d", m.statuses[0].LastBlockNum)
		}

		// Next cycle: the second read fails; LastBlockNum stays at 11.
		m.blockWait = func(context.Context) {
			chain.FailNextRPC("eth_getBlockByNumber", "flaky")
		}
		errCh := make(chan string, 10)
		m.check(context.Background(), newFakeDisplay(), errCh)
		st := m.statuses[0]
		if st.Alive || !strings.Contains(st.ErrStr, "flaky") || st.LastBlockNum != 11 {
			t.Fatalf("status = %+v, want stale block 11 with error", st)
		}
	})
}

func TestRPCMonitorEmptyURL(t *testing.T) {
	t.Parallel()
	nodes := []RPCConfig{{Name: "unset", URL: "", ChainID: "1"}}
	m := NewRPCMonitor(nodes, nil, NewSharedRPCState(), testIntervals())
	m.blockWait = func(context.Context) {}

	disp := newFakeDisplay()
	errCh := make(chan string, 10)
	m.check(context.Background(), disp, errCh)

	if msgs := drain(errCh); len(msgs) != 0 {
		t.Fatalf("unset URL must not raise errors, got %v", msgs)
	}
	if row := disp.Row(1); !strings.Contains(row, "*** not set ***") {
		t.Fatalf("row = %q", row)
	}
}

func TestRPCMonitorOfficialFailureKeepsLagSentinel(t *testing.T) {
	t.Parallel()
	follower := testchain.New(t)
	follower.EnsureBlock(95)

	nodes := []RPCConfig{
		{Name: "official-main", URL: "http://127.0.0.1:1", ChainID: "1", IsOfficial: true},
		{Name: "follower", URL: follower.URL(), ChainID: "1"},
	}
	shared := NewSharedRPCState()
	m := NewRPCMonitor(nodes, map[string]string{"mainnet": "official-main"}, shared, testIntervals())
	m.blockWait = func(context.Context) { follower.EnsureBlock(97) }

	errCh := make(chan string, 10)
	m.check(context.Background(), newFakeDisplay(), errCh)

	// Official never reported a block: shared state stays 0 and the
	// follower's lag stays at the sentinel.
	if got := shared.Official("1"); got != 0 {
		t.Fatalf("shared official = %d, want 0", got)
	}
	if got := m.statuses[1].OfficialLagDiff; got != math.MaxInt64 {
		t.Fatalf("lag = %d, want MaxInt64 sentinel", got)
	}
}

func TestRPCMonitorStartLoopStopsOnCancel(t *testing.T) {
	t.Parallel()
	chain := testchain.New(t)
	chain.EnsureBlock(5)

	nodes := []RPCConfig{{Name: "n", URL: chain.URL(), ChainID: "1"}}
	m := NewRPCMonitor(nodes, nil, NewSharedRPCState(), testIntervals())
	cycles := make(chan struct{}, 100)
	m.blockWait = func(context.Context) {
		select {
		case cycles <- struct{}{}:
		default:
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		m.Start(ctx, newFakeDisplay(), make(chan string, 100))
		close(done)
	}()

	<-cycles
	<-cycles
	cancel()
	waitFor(t, "loop exit", func() bool {
		select {
		case <-done:
			return true
		default:
			return false
		}
	})
}

func TestSharedRPCStateRoundTrip(t *testing.T) {
	t.Parallel()
	s := NewSharedRPCState()
	for i, chainID := range []string{"1", "11155111", "42161", "421614"} {
		s.UpdateOfficial(chainID, int64(100+i))
	}
	for i, chainID := range []string{"1", "11155111", "42161", "421614"} {
		if got := s.Official(chainID); got != int64(100+i) {
			t.Fatalf("Official(%s) = %d, want %d", chainID, got, 100+i)
		}
	}
	// Unknown chain ids are ignored on write and read as zero.
	s.UpdateOfficial("999", 5)
	if got := s.Official("999"); got != 0 {
		t.Fatalf("unknown chain id = %d, want 0", got)
	}
}
