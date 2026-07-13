package logscan

import (
	"context"
	"errors"
	"math"
	"math/big"
	"reflect"
	"strings"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type filtererFunc func(context.Context, ethereum.FilterQuery) ([]types.Log, error)

func (f filtererFunc) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return f(ctx, query)
}

func validOptions() Options {
	return Options{
		FromBlock:    10,
		ToBlock:      20,
		InitialBatch: 4,
		MinBatch:     2,
		RetryDelay:   time.Second,
	}
}

func TestScanInclusiveRangesAndStats(t *testing.T) {
	var ranges [][2]uint64
	client := filtererFunc(func(_ context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
		ranges = append(ranges, [2]uint64{query.FromBlock.Uint64(), query.ToBlock.Uint64()})
		return nil, nil
	})

	stats, err := Scan(context.Background(), client, validOptions(), func(context.Context, types.Log) error {
		return nil
	})
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	wantRanges := [][2]uint64{{10, 13}, {14, 17}, {18, 20}}
	if !reflect.DeepEqual(ranges, wantRanges) {
		t.Fatalf("ranges = %v, want %v", ranges, wantRanges)
	}
	wantStats := Stats{BlocksScanned: 11, RangesScanned: 3, FetchAttempts: 3}
	if stats != wantStats {
		t.Fatalf("stats = %+v, want %+v", stats, wantStats)
	}
}

func TestScanPreservesQueryAddressTopicsAndInput(t *testing.T) {
	addresses := []common.Address{
		common.HexToAddress("0x1000000000000000000000000000000000000001"),
		common.HexToAddress("0x2000000000000000000000000000000000000002"),
	}
	topics := [][]common.Hash{
		{common.HexToHash("0x01"), common.HexToHash("0x02")},
		nil,
		{common.HexToHash("0x03")},
	}
	originalFrom := big.NewInt(1)
	originalTo := big.NewInt(2)
	opts := validOptions()
	opts.Query = ethereum.FilterQuery{
		FromBlock: originalFrom,
		ToBlock:   originalTo,
		Addresses: addresses,
		Topics:    topics,
	}

	client := filtererFunc(func(_ context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
		if !reflect.DeepEqual(query.Addresses, addresses) {
			t.Errorf("addresses = %v, want %v", query.Addresses, addresses)
		}
		if !reflect.DeepEqual(query.Topics, topics) {
			t.Errorf("topics = %v, want %v", query.Topics, topics)
		}
		query.Addresses[0] = common.Address{}
		query.Topics[0][0] = common.Hash{}
		return nil, nil
	})
	_, err := Scan(context.Background(), client, opts, func(context.Context, types.Log) error { return nil })
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	if originalFrom.Int64() != 1 || originalTo.Int64() != 2 {
		t.Fatalf("input query range was mutated: %s..%s", originalFrom, originalTo)
	}
	if !reflect.DeepEqual(opts.Query.Addresses, addresses) || !reflect.DeepEqual(opts.Query.Topics, topics) {
		t.Fatal("input query filters were mutated")
	}
}

func TestScanRetriesSameRangeAndHalvesToMinimum(t *testing.T) {
	fetchErr := errors.New("query too wide")
	var progress []Progress
	var sleeps []time.Duration
	calls := 0
	client := filtererFunc(func(_ context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
		calls++
		if calls <= 3 {
			return nil, fetchErr
		}
		return nil, nil
	})
	opts := Options{
		FromBlock:    0,
		ToBlock:      9,
		InitialBatch: 10,
		MinBatch:     3,
		RetryDelay:   7 * time.Second,
		OnProgress: func(_ context.Context, p Progress) error {
			progress = append(progress, p)
			return nil
		},
		Sleep: func(_ context.Context, delay time.Duration) error {
			sleeps = append(sleeps, delay)
			return nil
		},
	}

	stats, err := Scan(context.Background(), client, opts, func(context.Context, types.Log) error { return nil })
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	wantRanges := [][3]uint64{
		{0, 9, 10},
		{0, 4, 5},
		{0, 2, 3},
		{0, 2, 3},
		{3, 5, 3},
		{6, 8, 3},
		{9, 9, 3},
	}
	if len(progress) != len(wantRanges) {
		t.Fatalf("progress calls = %d, want %d: %+v", len(progress), len(wantRanges), progress)
	}
	for i, want := range wantRanges {
		got := progress[i]
		if got.FromBlock != want[0] || got.ToBlock != want[1] || got.BatchSize != want[2] {
			t.Errorf("progress[%d] = %+v, want range/batch %v", i, got, want)
		}
	}
	if !reflect.DeepEqual(sleeps, []time.Duration{7 * time.Second}) {
		t.Fatalf("sleeps = %v, want one minimum-batch delay", sleeps)
	}
	wantStats := Stats{
		BlocksScanned: 10,
		RangesScanned: 4,
		FetchAttempts: 7,
		FilterErrors:  3,
	}
	if stats != wantStats {
		t.Fatalf("stats = %+v, want %+v", stats, wantStats)
	}
}

func TestScanSkipsRemovedLogs(t *testing.T) {
	logs := []types.Log{
		{BlockNumber: 10, Index: 1},
		{BlockNumber: 10, Index: 2, Removed: true},
		{BlockNumber: 11, Index: 3},
	}
	var handled []uint
	client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
		return logs, nil
	})
	opts := validOptions()
	opts.ToBlock = opts.FromBlock

	stats, err := Scan(context.Background(), client, opts, func(_ context.Context, log types.Log) error {
		handled = append(handled, log.Index)
		return nil
	})
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	if !reflect.DeepEqual(handled, []uint{1, 3}) {
		t.Fatalf("handled indices = %v, want [1 3]", handled)
	}
	if stats.LogsSeen != 2 || stats.RemovedLogs != 1 {
		t.Fatalf("stats = %+v, want two seen and one removed", stats)
	}
}

func TestScanCallbackErrorStopsRange(t *testing.T) {
	boom := errors.New("write failed")
	client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
		return []types.Log{{BlockNumber: 10, Index: 4}, {BlockNumber: 10, Index: 5}}, nil
	})
	opts := validOptions()
	opts.ToBlock = opts.FromBlock
	calls := 0

	stats, err := Scan(context.Background(), client, opts, func(context.Context, types.Log) error {
		calls++
		return boom
	})
	if !errors.Is(err, boom) {
		t.Fatalf("error = %v, want callback error", err)
	}
	if !strings.Contains(err.Error(), "block=10 index=4") {
		t.Fatalf("error = %q, want deterministic log identity", err)
	}
	if calls != 1 {
		t.Fatalf("callback calls = %d, want 1", calls)
	}
	if stats.LogsSeen != 1 || stats.RangesScanned != 0 || stats.BlocksScanned != 0 {
		t.Fatalf("partial stats = %+v", stats)
	}
}

func TestScanCancellationDuringFetch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	entered := make(chan struct{})
	client := filtererFunc(func(ctx context.Context, _ ethereum.FilterQuery) ([]types.Log, error) {
		close(entered)
		<-ctx.Done()
		return nil, ctx.Err()
	})
	go func() {
		<-entered
		cancel()
	}()

	stats, err := Scan(ctx, client, validOptions(), func(context.Context, types.Log) error { return nil })
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context canceled", err)
	}
	if stats.FetchAttempts != 1 || stats.FilterErrors != 1 {
		t.Fatalf("stats = %+v, want one canceled fetch", stats)
	}
}

func TestScanHonorsCancellationWhenFetcherReturnsSuccess(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
		cancel()
		return nil, nil
	})

	stats, err := Scan(ctx, client, validOptions(), func(context.Context, types.Log) error { return nil })
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context canceled", err)
	}
	if stats.FetchAttempts != 1 || stats.RangesScanned != 0 {
		t.Fatalf("stats = %+v, want canceled fetch without completed range", stats)
	}
}

func TestScanCancellationDuringRetrySleep(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	entered := make(chan struct{})
	client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
		return nil, errors.New("node unavailable")
	})
	opts := validOptions()
	opts.InitialBatch = opts.MinBatch
	opts.Sleep = func(ctx context.Context, _ time.Duration) error {
		close(entered)
		<-ctx.Done()
		return ctx.Err()
	}
	go func() {
		<-entered
		cancel()
	}()

	stats, err := Scan(ctx, client, opts, func(context.Context, types.Log) error { return nil })
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context canceled", err)
	}
	if stats.FetchAttempts != 1 || stats.FilterErrors != 1 {
		t.Fatalf("stats = %+v, want one failed fetch", stats)
	}
}

func TestScanDefaultSleepRetriesAfterDelay(t *testing.T) {
	calls := 0
	client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
		calls++
		if calls == 1 {
			return nil, errors.New("temporary failure")
		}
		return nil, nil
	})
	opts := validOptions()
	opts.ToBlock = opts.FromBlock
	opts.MinBatch = opts.InitialBatch
	opts.RetryDelay = time.Millisecond

	stats, err := Scan(context.Background(), client, opts, func(context.Context, types.Log) error { return nil })
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	if calls != 2 || stats.FilterErrors != 1 || stats.FetchAttempts != 2 {
		t.Fatalf("calls=%d stats=%+v", calls, stats)
	}
}

func TestScanDefaultSleepHonorsCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	fetched := make(chan struct{})
	client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
		close(fetched)
		return nil, errors.New("temporary failure")
	})
	opts := validOptions()
	opts.ToBlock = opts.FromBlock
	opts.MinBatch = opts.InitialBatch
	opts.RetryDelay = time.Hour
	go func() {
		<-fetched
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	stats, err := Scan(ctx, client, opts, func(context.Context, types.Log) error { return nil })
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context canceled", err)
	}
	if stats.FilterErrors != 1 || stats.FetchAttempts != 1 {
		t.Fatalf("stats = %+v", stats)
	}
}

func TestScanCancellationCheckpoints(t *testing.T) {
	t.Run("before first fetch", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
			t.Fatal("FilterLogs called for canceled context")
			return nil, nil
		})

		stats, err := Scan(ctx, client, validOptions(), func(context.Context, types.Log) error { return nil })
		if !errors.Is(err, context.Canceled) || stats != (Stats{}) {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("after progress", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		opts := validOptions()
		opts.OnProgress = func(context.Context, Progress) error {
			cancel()
			return nil
		}
		client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
			t.Fatal("FilterLogs called after progress canceled context")
			return nil, nil
		})

		stats, err := Scan(ctx, client, opts, func(context.Context, types.Log) error { return nil })
		if !errors.Is(err, context.Canceled) || stats != (Stats{}) {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("between logs", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
			return []types.Log{{BlockNumber: 10}, {BlockNumber: 10, Index: 1}}, nil
		})
		opts := validOptions()
		opts.ToBlock = opts.FromBlock
		handled := 0

		stats, err := Scan(ctx, client, opts, func(context.Context, types.Log) error {
			handled++
			cancel()
			return nil
		})
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
		if handled != 1 || stats.LogsSeen != 1 || stats.RangesScanned != 0 {
			t.Fatalf("handled=%d stats=%+v", handled, stats)
		}
	})
}

func TestScanProgressError(t *testing.T) {
	boom := errors.New("output closed")
	opts := validOptions()
	opts.OnProgress = func(ctx context.Context, p Progress) error {
		if ctx == nil || p.FromBlock != opts.FromBlock {
			t.Fatalf("progress = %+v, context=%v", p, ctx)
		}
		return boom
	}
	client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
		t.Fatal("FilterLogs called after progress error")
		return nil, nil
	})

	stats, err := Scan(context.Background(), client, opts, func(context.Context, types.Log) error { return nil })
	if !errors.Is(err, boom) {
		t.Fatalf("error = %v, want progress error", err)
	}
	if stats != (Stats{}) {
		t.Fatalf("stats = %+v, want zero", stats)
	}
}

func TestScanUint64MaximumDoesNotOverflow(t *testing.T) {
	var ranges [][2]uint64
	client := filtererFunc(func(_ context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
		ranges = append(ranges, [2]uint64{query.FromBlock.Uint64(), query.ToBlock.Uint64()})
		return nil, nil
	})
	opts := Options{
		FromBlock:    math.MaxUint64 - 1,
		ToBlock:      math.MaxUint64,
		InitialBatch: 100,
		MinBatch:     1,
		RetryDelay:   time.Second,
	}

	stats, err := Scan(context.Background(), client, opts, func(context.Context, types.Log) error { return nil })
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	want := [][2]uint64{{math.MaxUint64 - 1, math.MaxUint64}}
	if !reflect.DeepEqual(ranges, want) {
		t.Fatalf("ranges = %v, want %v", ranges, want)
	}
	if stats.BlocksScanned != 2 || stats.RangesScanned != 1 {
		t.Fatalf("stats = %+v, want two blocks", stats)
	}
}

func TestScanFullUint64DomainSaturatesBlockStats(t *testing.T) {
	var ranges [][2]uint64
	client := filtererFunc(func(_ context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
		ranges = append(ranges, [2]uint64{query.FromBlock.Uint64(), query.ToBlock.Uint64()})
		return nil, nil
	})
	opts := Options{
		FromBlock:    0,
		ToBlock:      math.MaxUint64,
		InitialBatch: math.MaxUint64,
		MinBatch:     1,
		RetryDelay:   time.Second,
	}

	stats, err := Scan(context.Background(), client, opts, func(context.Context, types.Log) error { return nil })
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	want := [][2]uint64{{0, math.MaxUint64 - 1}, {math.MaxUint64, math.MaxUint64}}
	if !reflect.DeepEqual(ranges, want) {
		t.Fatalf("ranges = %v, want %v", ranges, want)
	}
	if stats.BlocksScanned != math.MaxUint64 || stats.RangesScanned != 2 {
		t.Fatalf("stats = %+v, want saturated blocks and two ranges", stats)
	}
}

func TestAddBlocksSaturatingFullDomain(t *testing.T) {
	if got := addBlocksSaturating(0, 0, math.MaxUint64); got != math.MaxUint64 {
		t.Fatalf("addBlocksSaturating = %d, want saturation", got)
	}
}

func TestScanValidatesOptions(t *testing.T) {
	client := filtererFunc(func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
		return nil, nil
	})
	handle := func(context.Context, types.Log) error { return nil }
	blockHash := common.HexToHash("0x01")

	tests := []struct {
		name   string
		client Filterer
		opts   Options
		handle HandleFunc
	}{
		{name: "nil client", opts: validOptions(), handle: handle},
		{name: "nil callback", client: client, opts: validOptions()},
		{name: "reversed range", client: client, opts: func() Options {
			o := validOptions()
			o.FromBlock, o.ToBlock = 21, 20
			return o
		}(), handle: handle},
		{name: "zero initial batch", client: client, opts: func() Options {
			o := validOptions()
			o.InitialBatch = 0
			return o
		}(), handle: handle},
		{name: "zero minimum batch", client: client, opts: func() Options {
			o := validOptions()
			o.MinBatch = 0
			return o
		}(), handle: handle},
		{name: "initial below minimum", client: client, opts: func() Options {
			o := validOptions()
			o.InitialBatch, o.MinBatch = 2, 3
			return o
		}(), handle: handle},
		{name: "zero retry delay", client: client, opts: func() Options {
			o := validOptions()
			o.RetryDelay = 0
			return o
		}(), handle: handle},
		{name: "negative retry delay", client: client, opts: func() Options {
			o := validOptions()
			o.RetryDelay = -time.Second
			return o
		}(), handle: handle},
		{name: "block hash query", client: client, opts: func() Options {
			o := validOptions()
			o.Query.BlockHash = &blockHash
			return o
		}(), handle: handle},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := Scan(context.Background(), tt.client, tt.opts, tt.handle)
			if err == nil {
				t.Fatal("Scan succeeded, want validation error")
			}
			if stats != (Stats{}) {
				t.Fatalf("stats = %+v, want zero", stats)
			}
		})
	}
}
