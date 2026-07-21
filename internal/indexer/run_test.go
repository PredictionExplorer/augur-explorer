// Unit tests (no Docker) for the polling loop's control flow: startup
// validation, shutdown, the caught-up wait, retry/backoff on chain failures,
// the circuit breaker and adaptive batch sizing. Paths that persist rows run
// in run_integration_test.go against a real database.
package indexer

import (
	"context"
	"errors"
	"log/slog"
	"math"
	"math/big"
	"strings"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// fakeClient scripts the Client surface per test.
type fakeClient struct {
	mu              sync.Mutex
	blockNumberFn   func() (uint64, error)
	filterLogsFn    func(from, to uint64) ([]types.Log, error)
	headerFn        func(number *big.Int) (*types.Header, error)
	fetchedRanges   [][2]uint64
	blockNumberCall int
}

func (f *fakeClient) BlockNumber(ctx context.Context) (uint64, error) {
	f.mu.Lock()
	f.blockNumberCall++
	fn := f.blockNumberFn
	f.mu.Unlock()
	if fn == nil {
		return 0, errors.New("fakeClient: BlockNumber not scripted")
	}
	return fn()
}

func (f *fakeClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	from, to := q.FromBlock.Uint64(), q.ToBlock.Uint64()
	f.mu.Lock()
	f.fetchedRanges = append(f.fetchedRanges, [2]uint64{from, to})
	fn := f.filterLogsFn
	f.mu.Unlock()
	if fn == nil {
		return nil, errors.New("fakeClient: FilterLogs not scripted")
	}
	return fn(from, to)
}

func (f *fakeClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	f.mu.Lock()
	fn := f.headerFn
	f.mu.Unlock()
	if fn == nil {
		return nil, errors.New("fakeClient: HeaderByNumber not scripted")
	}
	return fn(number)
}

func (f *fakeClient) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	return nil, false, errors.New("fakeClient: TransactionByHash not scripted")
}

func (f *fakeClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, errors.New("fakeClient: TransactionReceipt not scripted")
}

func (f *fakeClient) ranges() [][2]uint64 {
	f.mu.Lock()
	defer f.mu.Unlock()
	out := make([][2]uint64, len(f.fetchedRanges))
	copy(out, f.fetchedRanges)
	return out
}

func (f *fakeClient) blockCalls() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.blockNumberCall
}

// fakeProgress is an in-memory watermark.
type fakeProgress struct {
	mu     sync.Mutex
	last   int64
	writes []int64
}

func (p *fakeProgress) LastBlock(ctx context.Context) (int64, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.last, nil
}

func (p *fakeProgress) SetLastBlock(ctx context.Context, block int64) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.last = block
	p.writes = append(p.writes, block)
	return nil
}

func (p *fakeProgress) writesCopy() []int64 {
	p.mu.Lock()
	defer p.mu.Unlock()
	out := make([]int64, len(p.writes))
	copy(out, p.writes)
	return out
}

// unitEngine builds an Engine over fakes; the store handle is structurally
// present but the scripted paths never touch the database.
func unitEngine(t *testing.T, cfg Config) *Engine {
	t.Helper()
	if cfg.Store == nil {
		cfg.Store = store.NewFromPool(nil)
	}
	if cfg.Logger == nil {
		cfg.Logger = slog.New(slog.DiscardHandler)
	}
	if cfg.Retry.MinDelay == 0 {
		cfg.Retry.MinDelay = time.Millisecond
	}
	if cfg.Retry.MaxDelay == 0 {
		cfg.Retry.MaxDelay = 5 * time.Millisecond
	}
	if cfg.CaughtUpDelay == 0 {
		cfg.CaughtUpDelay = time.Millisecond
	}
	e, err := New(cfg)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return e
}

// TestLogBlockNum pins the chain-boundary conversion every pipeline entry
// point applies to log block numbers: values through math.MaxInt64 convert
// exactly, anything larger is rejected as corrupt node data.
func TestLogBlockNum(t *testing.T) {
	for _, tc := range []struct {
		in   uint64
		want int64
	}{
		{0, 0},
		{455_767_500, 455_767_500},
		{math.MaxInt64, math.MaxInt64},
	} {
		got, err := logBlockNum(&types.Log{BlockNumber: tc.in})
		if err != nil || got != tc.want {
			t.Errorf("logBlockNum(%d) = %d, %v; want %d", tc.in, got, err, tc.want)
		}
	}
	if _, err := logBlockNum(&types.Log{BlockNumber: math.MaxInt64 + 1}); err == nil {
		t.Error("logBlockNum(MaxInt64+1): want error, got nil")
	}
}

// TestPipelineRejectsOverflowingBlockNumber drives a log whose block number
// exceeds int64 through processBatch, InsertEventLog and the backfill scan:
// all must fail loudly before touching the database (the engine store here
// has no pool).
func TestPipelineRejectsOverflowingBlockNumber(t *testing.T) {
	overflowing := types.Log{BlockNumber: math.MaxUint64}
	client := &fakeClient{
		filterLogsFn: func(_, _ uint64) ([]types.Log, error) {
			return []types.Log{overflowing}, nil
		},
	}
	e := unitEngine(t, Config{Client: client})

	last, stage, err := e.processBatch(context.Background(), []types.Log{overflowing})
	if err == nil || !strings.Contains(err.Error(), "overflows int64") {
		t.Fatalf("processBatch error = %v, want block-number overflow", err)
	}
	if last != 0 || stage != "block" {
		t.Errorf("processBatch = (last %d, stage %q), want (0, block)", last, stage)
	}

	if _, err := e.InsertEventLog(context.Background(), overflowing, 1); err == nil ||
		!strings.Contains(err.Error(), "overflows int64") {
		t.Errorf("InsertEventLog error = %v, want block-number overflow", err)
	}

	stats, err := e.BackfillContractEvtLogs(context.Background(), []common.Address{{0x01}}, 1, 1, 10)
	if err == nil || !strings.Contains(err.Error(), "overflows int64") {
		t.Errorf("BackfillContractEvtLogs error = %v, want block-number overflow", err)
	}
	if stats != (BackfillStats{}) {
		t.Errorf("backfill reported non-durable work before rejecting: %+v", stats)
	}
	if stats, err := e.backfillBlock(context.Background(), nil); err != nil || stats != (BackfillStats{}) {
		t.Errorf("empty backfill block = %+v, %v", stats, err)
	}
	valid := types.Log{BlockNumber: 1}
	if stats, err := e.backfillBlock(context.Background(), []types.Log{valid, overflowing}); err == nil ||
		!strings.Contains(err.Error(), "overflows int64") ||
		stats != (BackfillStats{}) {
		t.Errorf("later overflow backfill = %+v, %v", stats, err)
	}
}

// failingProgress accepts the watermark read but fails every write, driving
// the run loop's watermark-failure arm.
type failingProgress struct {
	fakeProgress

	writeErr error
}

func (p *failingProgress) SetLastBlock(context.Context, int64) error { return p.writeErr }

// TestRunWatermarkWriteFailureTripsBreaker pins the watermark stage: a
// healthy batch whose progress write keeps failing must retry and then trip
// the circuit breaker with the watermark stage in the error.
func TestRunWatermarkWriteFailureTripsBreaker(t *testing.T) {
	writeErr := errors.New("watermark write refused")
	client := &fakeClient{
		blockNumberFn: func() (uint64, error) { return 12, nil },
		filterLogsFn:  func(_, _ uint64) ([]types.Log, error) { return nil, nil },
	}
	e := unitEngine(t, Config{
		Client:    client,
		Progress:  &failingProgress{fakeProgress: fakeProgress{last: 10}, writeErr: writeErr},
		Process:   func(context.Context, int64) error { return nil },
		Contracts: []common.Address{{0x01}},
		Retry:     RetryConfig{MaxConsecutiveFailures: 2, MinDelay: time.Millisecond, MaxDelay: 2 * time.Millisecond},
	})
	err := runWithTimeout(t, e, context.Background())
	if err == nil || !errors.Is(err, writeErr) {
		t.Fatalf("Run = %v, want the watermark write failure", err)
	}
	if !strings.Contains(err.Error(), "watermark") {
		t.Errorf("breaker error missing watermark stage: %v", err)
	}
}

// runWithTimeout runs Run on a goroutine and fails the test if it does not
// return within the deadline.
func runWithTimeout(t *testing.T, e *Engine, ctx context.Context) error {
	t.Helper()
	done := make(chan error, 1)
	go func() { done <- e.Run(ctx) }()
	select {
	case err := <-done:
		return err
	case <-time.After(30 * time.Second):
		t.Fatal("Run did not return within 30s")
		return nil
	}
}

func TestNewValidatesRequiredDeps(t *testing.T) {
	if _, err := New(Config{Client: &fakeClient{}}); err == nil || !strings.Contains(err.Error(), "Store") {
		t.Errorf("New without Store: err = %v, want mention of Store", err)
	}
	if _, err := New(Config{Store: store.NewFromPool(nil)}); err == nil || !strings.Contains(err.Error(), "Client") {
		t.Errorf("New without Client: err = %v, want mention of Client", err)
	}
}

func TestRunRequiresLoopDeps(t *testing.T) {
	e := unitEngine(t, Config{Client: &fakeClient{}})
	if err := e.Run(context.Background()); err == nil || !strings.Contains(err.Error(), "Progress") {
		t.Errorf("Run without Progress/Process: err = %v, want validation error", err)
	}

	e = unitEngine(t, Config{
		Client:   &fakeClient{},
		Progress: &fakeProgress{},
		Process:  func(context.Context, int64) error { return nil },
	})
	if err := e.Run(context.Background()); err == nil || !strings.Contains(err.Error(), "Contracts") {
		t.Errorf("Run without Contracts: err = %v, want validation error", err)
	}
}

func TestRunExitsOnCanceledContext(t *testing.T) {
	e := unitEngine(t, Config{
		Client:    &fakeClient{},
		Progress:  &fakeProgress{last: 10},
		Process:   func(context.Context, int64) error { return nil },
		Contracts: []common.Address{{0x01}},
	})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := runWithTimeout(t, e, ctx); err != nil {
		t.Fatalf("Run on canceled context = %v, want nil (clean shutdown)", err)
	}
}

func TestRunCaughtUpWaitsAndHonorsCancel(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		client := &fakeClient{
			blockNumberFn: func() (uint64, error) { return 10, nil }, // head == watermark
		}
		e := unitEngine(t, Config{
			Client:        client,
			Progress:      &fakeProgress{last: 10},
			Process:       func(context.Context, int64) error { return nil },
			Contracts:     []common.Address{{0x01}},
			CaughtUpDelay: time.Hour, // cancellation must interrupt the wait
		})
		ctx, cancel := context.WithCancel(t.Context())
		done := make(chan error, 1)
		go func() { done <- e.Run(ctx) }()

		synctest.Wait() // Run is blocked in the one-hour caught-up wait.
		if got := client.blockCalls(); got != 1 {
			t.Fatalf("BlockNumber calls before cancel = %d, want 1", got)
		}
		cancel()
		synctest.Wait()
		if err := <-done; err != nil {
			t.Fatalf("Run = %v, want nil", err)
		}
		if got := len(client.ranges()); got != 0 {
			t.Errorf("caught-up loop fetched %d ranges, want 0", got)
		}
	})
}

func TestRunCancellationDuringBackoffSleep(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		headErr := errors.New("temporary RPC outage")
		client := &fakeClient{
			blockNumberFn: func() (uint64, error) { return 0, headErr },
		}
		e := unitEngine(t, Config{
			Client:    client,
			Progress:  &fakeProgress{last: 10},
			Process:   func(context.Context, int64) error { return nil },
			Contracts: []common.Address{{0x01}},
			Retry: RetryConfig{
				MaxConsecutiveFailures: 10,
				MinDelay:               time.Hour,
				MaxDelay:               time.Hour,
			},
		})
		ctx, cancel := context.WithCancel(t.Context())
		done := make(chan error, 1)
		go func() { done <- e.Run(ctx) }()

		synctest.Wait() // Run is blocked in its first retry backoff.
		if got := client.blockCalls(); got != 1 {
			t.Fatalf("BlockNumber calls before cancel = %d, want 1", got)
		}
		cancel()
		synctest.Wait()
		if err := <-done; err != nil {
			t.Fatalf("Run after backoff cancellation = %v, want nil", err)
		}
		if got := client.blockCalls(); got != 1 {
			t.Fatalf("BlockNumber calls after cancel = %d, want 1", got)
		}
	})
}

func TestEngineSleepBoundaryAndExpiry(t *testing.T) {
	e := &Engine{}
	if !e.sleep(context.Background(), 0) {
		t.Fatal("zero delay on active context = false, want true")
	}
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	if e.sleep(cancelled, 0) {
		t.Fatal("zero delay on cancelled context = true, want false")
	}

	synctest.Test(t, func(t *testing.T) {
		start := time.Now()
		if !e.sleep(t.Context(), time.Hour) {
			t.Fatal("timer expiry = false, want true")
		}
		if elapsed := time.Since(start); elapsed != time.Hour {
			t.Fatalf("fake elapsed time = %v, want 1h", elapsed)
		}
	})
}

func TestRunCircuitBreakerTripsAfterConsecutiveHeadFailures(t *testing.T) {
	headErr := errors.New("connection refused")
	client := &fakeClient{
		blockNumberFn: func() (uint64, error) { return 0, headErr },
	}
	e := unitEngine(t, Config{
		Client:    client,
		Progress:  &fakeProgress{last: 10},
		Process:   func(context.Context, int64) error { return nil },
		Contracts: []common.Address{{0x01}},
		Retry:     RetryConfig{MaxConsecutiveFailures: 3, MinDelay: time.Millisecond, MaxDelay: 2 * time.Millisecond},
	})
	err := runWithTimeout(t, e, context.Background())
	if err == nil {
		t.Fatal("Run = nil, want circuit-breaker error")
	}
	if !errors.Is(err, headErr) {
		t.Errorf("breaker error does not wrap the cause: %v", err)
	}
	if !strings.Contains(err.Error(), "3 consecutive") || !strings.Contains(err.Error(), "chain_head") {
		t.Errorf("breaker error missing count/stage: %v", err)
	}
	if client.blockNumberCall != 3 {
		t.Errorf("BlockNumber called %d times, want exactly 3", client.blockNumberCall)
	}
}

// TestRunRejectsOverflowingChainHead pins the chain-boundary guard on the
// head read: a head beyond int64 (corrupt node data) is a batch failure that
// trips the breaker instead of wrapping the watermark negative.
func TestRunRejectsOverflowingChainHead(t *testing.T) {
	client := &fakeClient{
		blockNumberFn: func() (uint64, error) { return math.MaxUint64, nil },
	}
	e := unitEngine(t, Config{
		Client:    client,
		Progress:  &fakeProgress{last: 10},
		Process:   func(context.Context, int64) error { return nil },
		Contracts: []common.Address{{0x01}},
		Retry:     RetryConfig{MaxConsecutiveFailures: 2, MinDelay: time.Millisecond, MaxDelay: 2 * time.Millisecond},
	})
	err := runWithTimeout(t, e, context.Background())
	if err == nil || !strings.Contains(err.Error(), "overflows int64") {
		t.Fatalf("Run = %v, want chain-head overflow failure", err)
	}
	if !strings.Contains(err.Error(), "chain_head") {
		t.Errorf("breaker error missing stage: %v", err)
	}
	if got := len(client.ranges()); got != 0 {
		t.Errorf("overflowing head still fetched %d ranges, want 0", got)
	}
}

func TestRunRecoversAfterTransientFailures(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		// Two failures, one healthy caught-up poll, then two more failures.
		// The healthy poll must reset the breaker; otherwise call 4 would be
		// treated as the third consecutive failure and Run would exit.
		calls := 0
		ctx, cancel := context.WithCancel(t.Context())
		client := &fakeClient{}
		client.blockNumberFn = func() (uint64, error) {
			calls++
			switch calls {
			case 1, 2, 4, 5:
				return 0, errors.New("transient blip")
			default:
				if calls >= 6 {
					cancel()
				}
				return 10, nil
			}
		}
		e := unitEngine(t, Config{
			Client:        client,
			Progress:      &fakeProgress{last: 10},
			Process:       func(context.Context, int64) error { return nil },
			Contracts:     []common.Address{{0x01}},
			Retry:         RetryConfig{MaxConsecutiveFailures: 3, MinDelay: time.Millisecond, MaxDelay: 2 * time.Millisecond},
			CaughtUpDelay: time.Millisecond,
		})
		done := make(chan error, 1)
		go func() { done <- e.Run(ctx) }()

		time.Sleep(20 * time.Millisecond) // fake time; enough for every scripted wait
		synctest.Wait()
		if err := <-done; err != nil {
			t.Fatalf("Run = %v, want nil after separated failure streaks", err)
		}
		if calls != 6 {
			t.Fatalf("BlockNumber calls = %d, want 6", calls)
		}
	})
}

func TestRunFetchErrorShrinksBatchAndEmptySuccessGrowsIt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var mu sync.Mutex
	step := 0
	client := &fakeClient{
		blockNumberFn: func() (uint64, error) { return 100_000, nil },
	}
	client.filterLogsFn = func(from, to uint64) ([]types.Log, error) {
		mu.Lock()
		defer mu.Unlock()
		step++
		switch step {
		case 1, 2:
			return nil, errors.New("query returned more than 10000 results")
		case 3, 4:
			return nil, nil // empty success: batch doubles
		default:
			cancel()
			return nil, ctx.Err()
		}
	}
	progress := &fakeProgress{last: 100}
	e := unitEngine(t, Config{
		Client:    client,
		Progress:  progress,
		Process:   func(context.Context, int64) error { return nil },
		Contracts: []common.Address{{0x01}},
		Batch:     BatchConfig{Initial: 400, Min: 100, Max: 1600},
		Retry:     RetryConfig{MaxConsecutiveFailures: 10, MinDelay: time.Millisecond, MaxDelay: 2 * time.Millisecond},
	})
	if err := runWithTimeout(t, e, ctx); err != nil {
		t.Fatalf("Run = %v, want nil", err)
	}

	ranges := client.ranges()
	if len(ranges) < 4 {
		t.Fatalf("expected at least 4 fetches, got %v", ranges)
	}
	// Fetch 1: initial 400 -> [101,500]. Errors halve: 200, then success at
	// 100... wait: two errors halve 400->200->100; fetch 2 uses 200.
	wantWidths := []uint64{400, 200, 100, 200}
	for i, want := range wantWidths {
		if got := ranges[i][1] - ranges[i][0] + 1; got != want {
			t.Errorf("fetch %d range %v: width = %d, want %d", i+1, ranges[i], got, want)
		}
	}
	// Empty successes acknowledge the scanned range.
	writes := progress.writesCopy()
	if len(writes) < 2 || writes[0] != 200 || writes[1] != 400 {
		t.Errorf("watermark writes = %v, want [200 400 ...] (acknowledged empty ranges)", writes)
	}
}

// errProgress scripts watermark failures.
type errProgress struct {
	readErr  error
	writeErr error
	last     int64
}

func (p *errProgress) LastBlock(ctx context.Context) (int64, error) {
	if p.readErr != nil {
		return 0, p.readErr
	}
	return p.last, nil
}

func (p *errProgress) SetLastBlock(ctx context.Context, block int64) error {
	return p.writeErr
}

func TestRunStartupWatermarkFailureTripsBreaker(t *testing.T) {
	// The startup watermark read is a database failure: retried with
	// backoff, fatal only when persistent (the circuit breaker trips).
	e := unitEngine(t, Config{
		Client:    &fakeClient{},
		Progress:  &errProgress{readErr: errors.New("status table on fire")},
		Process:   func(context.Context, int64) error { return nil },
		Contracts: []common.Address{{0x01}},
		Retry:     RetryConfig{MaxConsecutiveFailures: 2, MinDelay: time.Millisecond, MaxDelay: time.Millisecond},
	})
	err := runWithTimeout(t, e, context.Background())
	if err == nil || !strings.Contains(err.Error(), "reading processing status") {
		t.Fatalf("Run = %v, want the wrapped watermark read failure", err)
	}
	if !strings.Contains(err.Error(), "2 consecutive batch failures") {
		t.Errorf("Run = %v, want the circuit-breaker wrap", err)
	}
}

func TestSetLastBlockPersistFailure(t *testing.T) {
	e := unitEngine(t, Config{
		Client:    &fakeClient{},
		Progress:  &errProgress{writeErr: errors.New("disk full")},
		Process:   func(context.Context, int64) error { return nil },
		Contracts: []common.Address{{0x01}},
	})
	if err := e.setLastBlock(context.Background(), 7); err == nil ||
		!strings.Contains(err.Error(), "updating processing status") {
		t.Fatalf("setLastBlock = %v, want the wrapped persist failure", err)
	}
}

func TestEventTypeLabel(t *testing.T) {
	known := common.HexToHash("0xaaaa")
	e := unitEngine(t, Config{
		Client:    &fakeClient{},
		Progress:  &fakeProgress{},
		Process:   func(context.Context, int64) error { return nil },
		Contracts: []common.Address{{0x01}},
		TopicName: func(topic common.Hash) string {
			if topic == known {
				return "BidPlaced"
			}
			return ""
		},
	})
	cases := []struct {
		name   string
		topics []common.Hash
		want   string
	}{
		{"no topics", nil, "none"},
		{"known topic", []common.Hash{known}, "BidPlaced"},
		{"unknown topic", []common.Hash{common.HexToHash("0xbbbb")}, "other"},
	}
	for _, tc := range cases {
		if got := e.eventTypeLabel(&types.Log{Topics: tc.topics}); got != tc.want {
			t.Errorf("%s: eventTypeLabel = %q, want %q", tc.name, got, tc.want)
		}
	}

	// Without a TopicName resolver every topic is "other".
	e2 := unitEngine(t, Config{
		Client:    &fakeClient{},
		Progress:  &fakeProgress{},
		Process:   func(context.Context, int64) error { return nil },
		Contracts: []common.Address{{0x01}},
	})
	if got := e2.eventTypeLabel(&types.Log{Topics: []common.Hash{known}}); got != "other" {
		t.Errorf("eventTypeLabel without resolver = %q, want other", got)
	}
}

func TestRunWatermarkZeroFallsBackToStoreWatermark(t *testing.T) {
	// A fresh status row (LastBlock 0) must consult the store's block
	// watermark; with a nil pool the call panics, which proves (and pins)
	// that the fallback is exercised — the integration suite covers the
	// value semantics.
	e := unitEngine(t, Config{
		Client:    &fakeClient{},
		Progress:  &fakeProgress{last: 0},
		Process:   func(context.Context, int64) error { return nil },
		Contracts: []common.Address{{0x01}},
	})
	defer func() {
		if recover() == nil {
			t.Fatal("expected the store fallback path to be reached (panic on nil pool)")
		}
	}()
	_, _ = e.lastProcessedBlock(context.Background())
}
