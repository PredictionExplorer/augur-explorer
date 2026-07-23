package indexer

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"testing"
	"testing/synctest"
	"time"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// hangingClient blocks every call until its context dies — the black-holed
// JSON-RPC endpoint. calls counts attempts across all methods.
type hangingClient struct {
	calls int
}

func (h *hangingClient) hang(ctx context.Context) error {
	h.calls++
	<-ctx.Done()
	return ctx.Err()
}

func (h *hangingClient) BlockNumber(ctx context.Context) (uint64, error) {
	return 0, h.hang(ctx)
}

func (h *hangingClient) HeaderByNumber(ctx context.Context, _ *big.Int) (*types.Header, error) {
	return nil, h.hang(ctx)
}

func (h *hangingClient) FilterLogs(ctx context.Context, _ ethereum.FilterQuery) ([]types.Log, error) {
	return nil, h.hang(ctx)
}

func (h *hangingClient) TransactionByHash(ctx context.Context, _ ethcommon.Hash) (*types.Transaction, bool, error) {
	return nil, false, h.hang(ctx)
}

func (h *hangingClient) TransactionReceipt(ctx context.Context, _ ethcommon.Hash) (*types.Receipt, error) {
	return nil, h.hang(ctx)
}

// TestBoundedClientBoundsEveryMethod proves each of the five Client methods
// returns context.DeadlineExceeded after exactly the configured bound when
// the endpoint black-holes — the calls can no longer hang forever.
func TestBoundedClientBoundsEveryMethod(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const bound = 42 * time.Second
		inner := &hangingClient{}
		c := newBoundedClient(inner, bound)

		calls := map[string]func(context.Context) error{
			"BlockNumber": func(ctx context.Context) error { _, err := c.BlockNumber(ctx); return err },
			"HeaderByNumber": func(ctx context.Context) error {
				_, err := c.HeaderByNumber(ctx, big.NewInt(1))
				return err
			},
			"FilterLogs": func(ctx context.Context) error {
				_, err := c.FilterLogs(ctx, ethereum.FilterQuery{})
				return err
			},
			"TransactionByHash": func(ctx context.Context) error {
				_, _, err := c.TransactionByHash(ctx, ethcommon.Hash{0x01})
				return err
			},
			"TransactionReceipt": func(ctx context.Context) error {
				_, err := c.TransactionReceipt(ctx, ethcommon.Hash{0x01})
				return err
			},
		}
		for name, call := range calls {
			start := time.Now()
			err := call(t.Context())
			if !errors.Is(err, context.DeadlineExceeded) {
				t.Errorf("%s on a black-holed endpoint: err = %v, want DeadlineExceeded", name, err)
			}
			if elapsed := time.Since(start); elapsed != bound {
				t.Errorf("%s returned after %v (fake time), want exactly the %v bound", name, elapsed, bound)
			}
		}
		if inner.calls != len(calls) {
			t.Fatalf("inner client saw %d calls, want %d", inner.calls, len(calls))
		}
	})
}

// TestBoundedClientKeepsEarlierCallerDeadline pins that the wrapper never
// extends a caller's tighter budget: WithTimeout picks the earlier deadline.
func TestBoundedClientKeepsEarlierCallerDeadline(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := newBoundedClient(&hangingClient{}, time.Hour)
		ctx, cancel := context.WithTimeout(t.Context(), 10*time.Millisecond)
		defer cancel()

		start := time.Now()
		_, err := c.BlockNumber(ctx)
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("err = %v, want DeadlineExceeded", err)
		}
		if elapsed := time.Since(start); elapsed != 10*time.Millisecond {
			t.Fatalf("returned after %v, want the caller's tighter 10ms deadline", elapsed)
		}
	})
}

// TestBoundedClientBoundsShutdownImmuneContexts pins the load-bearing
// property for ADR-0010's finish-the-batch semantics: even a
// context.WithoutCancel descendant — immune to SIGTERM — gets a fresh
// per-call deadline, so "finish the batch" can no longer mean "forever".
func TestBoundedClientBoundsShutdownImmuneContexts(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		const bound = 30 * time.Second
		c := newBoundedClient(&hangingClient{}, bound)

		parent, cancel := context.WithCancel(t.Context())
		cancel() // shutdown already signaled
		dbCtx := context.WithoutCancel(parent)

		start := time.Now()
		_, _, err := c.TransactionByHash(dbCtx, ethcommon.Hash{0x01})
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("err = %v, want DeadlineExceeded", err)
		}
		if elapsed := time.Since(start); elapsed != bound {
			t.Fatalf("returned after %v, want the %v bound despite WithoutCancel", elapsed, bound)
		}
	})
}

// TestNewWrapsClientWithRPCCallTimeout pins that every Engine gets the
// bounded client: the default when Config.RPCCallTimeout is zero, the
// configured value otherwise.
func TestNewWrapsClientWithRPCCallTimeout(t *testing.T) {
	e := unitEngine(t, Config{Client: &fakeClient{}})
	bounded, ok := e.client.(boundedClient)
	if !ok {
		t.Fatalf("Engine client is %T, want the boundedClient decorator", e.client)
	}
	if bounded.timeout != DefaultRPCCallTimeout {
		t.Fatalf("default RPC call timeout = %v, want %v", bounded.timeout, DefaultRPCCallTimeout)
	}

	e = unitEngine(t, Config{Client: &fakeClient{}, RPCCallTimeout: 3 * time.Second})
	if got := e.client.(boundedClient).timeout; got != 3*time.Second {
		t.Fatalf("configured RPC call timeout = %v, want 3s", got)
	}
}

// TestRunHungFetchFailsBatchesAndTripsBreaker proves the end-to-end
// degradation story: a black-holed eth_getLogs endpoint now produces
// ordinary timed-out fetch failures — counted by the breaker until Run
// returns and the supervisor restarts the process — where it used to hang
// the ETL forever without a single failure recorded.
func TestRunHungFetchFailsBatchesAndTripsBreaker(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		client := &hangingFetchClient{head: 100}
		e := unitEngine(t, Config{
			Client:         client,
			Progress:       &fakeProgress{last: 10},
			Process:        func(context.Context, int64) error { return nil },
			Contracts:      []ethcommon.Address{{0x01}},
			RPCCallTimeout: time.Second,
			Retry:          RetryConfig{MaxConsecutiveFailures: 3, MinDelay: time.Millisecond, MaxDelay: 2 * time.Millisecond},
		})

		start := time.Now()
		err := e.Run(t.Context())
		if err == nil || !strings.Contains(err.Error(), "giving up after 3 consecutive batch failures (stage fetch)") {
			t.Fatalf("Run = %v, want the fetch-stage circuit breaker", err)
		}
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("breaker error should carry the per-call deadline cause, got %v", err)
		}
		// Three bounded one-second hangs plus two sub-3ms backoffs; without
		// the bound this test would deadlock in the bubble.
		if elapsed := time.Since(start); elapsed < 3*time.Second || elapsed > 4*time.Second {
			t.Fatalf("Run returned after %v, want ~3s of bounded fetch hangs", elapsed)
		}
		if client.fetches != 3 {
			t.Fatalf("FilterLogs attempts = %d, want 3", client.fetches)
		}
	})
}

// hangingFetchClient answers the head probe instantly and black-holes every
// FilterLogs call.
type hangingFetchClient struct {
	head    uint64
	fetches int
}

func (h *hangingFetchClient) BlockNumber(context.Context) (uint64, error) { return h.head, nil }

func (h *hangingFetchClient) HeaderByNumber(ctx context.Context, _ *big.Int) (*types.Header, error) {
	<-ctx.Done()
	return nil, ctx.Err()
}

func (h *hangingFetchClient) FilterLogs(ctx context.Context, _ ethereum.FilterQuery) ([]types.Log, error) {
	h.fetches++
	<-ctx.Done()
	return nil, ctx.Err()
}

func (h *hangingFetchClient) TransactionByHash(ctx context.Context, _ ethcommon.Hash) (*types.Transaction, bool, error) {
	<-ctx.Done()
	return nil, false, ctx.Err()
}

func (h *hangingFetchClient) TransactionReceipt(ctx context.Context, _ ethcommon.Hash) (*types.Receipt, error) {
	<-ctx.Done()
	return nil, ctx.Err()
}

// TestFetchLogsPackageFunctionStaysUnbounded pins the split of concerns: the
// package-level FetchLogs helper adds no deadline of its own (opsctl callers
// bound their calls via logscan or their dialed client), while engine-owned
// calls are bounded by construction (TestNewWrapsClientWithRPCCallTimeout).
func TestFetchLogsPackageFunctionStaysUnbounded(t *testing.T) {
	probe := &deadlineProbeClient{}
	if _, err := FetchLogs(context.Background(), probe, 1, 2, []ethcommon.Address{{0x01}}); err != nil {
		t.Fatalf("FetchLogs: %v", err)
	}
	if probe.sawDeadline {
		t.Fatal("package-level FetchLogs must not add its own deadline")
	}
}

type deadlineProbeClient struct {
	fakeClient

	sawDeadline bool
}

func (p *deadlineProbeClient) FilterLogs(ctx context.Context, _ ethereum.FilterQuery) ([]types.Log, error) {
	_, p.sawDeadline = ctx.Deadline()
	return nil, nil
}
