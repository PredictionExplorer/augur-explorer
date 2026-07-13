package cstscan

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer/logscan"
)

type fakeClient struct {
	head      uint64
	headErr   error
	headCalls int
	filter    func(context.Context, ethereum.FilterQuery) ([]types.Log, error)
}

func (f *fakeClient) BlockNumber(context.Context) (uint64, error) {
	f.headCalls++
	return f.head, f.headErr
}

func (f *fakeClient) FilterLogs(
	ctx context.Context,
	query ethereum.FilterQuery,
) ([]types.Log, error) {
	if f.filter == nil {
		return nil, nil
	}
	return f.filter(ctx, query)
}

type captureLogger struct {
	mu    sync.Mutex
	lines []string
}

func (l *captureLogger) Printf(format string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lines = append(l.lines, fmt.Sprintf(format, args...))
}

func (l *captureLogger) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return strings.Join(l.lines, "\n")
}

func uint256(value int64) []byte {
	return common.LeftPadBytes(big.NewInt(value).Bytes(), common.HashLength)
}

func testOptions(from, to uint64) Options {
	return Options{
		FromBlock:    from,
		ToBlock:      to,
		InitialBatch: 2,
		MinBatch:     1,
		RetryDelay:   time.Millisecond,
	}
}

func TestDecodeAuctionLength(t *testing.T) {
	value, err := DecodeAuctionLength(uint256(123456))
	if err != nil {
		t.Fatalf("DecodeAuctionLength: %v", err)
	}
	if value.String() != "123456" {
		t.Fatalf("value = %s, want 123456", value)
	}
	for _, data := range [][]byte{nil, make([]byte, 31), make([]byte, 33)} {
		if _, err := DecodeAuctionLength(data); err == nil {
			t.Fatalf("DecodeAuctionLength accepted %d bytes", len(data))
		}
	}
}

func TestScanOutputAndLatestBlock(t *testing.T) {
	contract := common.HexToAddress("0x4000000000000000000000000000000000000004")
	topic := common.HexToHash("0xabc")
	firstHash := common.HexToHash("0x01")
	secondHash := common.HexToHash("0x02")
	var ranges [][2]uint64
	client := &fakeClient{
		head: 12,
		filter: func(_ context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
			if !reflect.DeepEqual(query.Addresses, []common.Address{contract}) {
				t.Fatalf("addresses = %v", query.Addresses)
			}
			if !reflect.DeepEqual(query.Topics, [][]common.Hash{{topic}}) {
				t.Fatalf("topics = %v", query.Topics)
			}
			from, to := query.FromBlock.Uint64(), query.ToBlock.Uint64()
			ranges = append(ranges, [2]uint64{from, to})
			switch from {
			case 10:
				return []types.Log{
					{BlockNumber: 10, TxHash: firstHash, Index: 3, Data: uint256(42)},
					{BlockNumber: 11, Removed: true},
				}, nil
			case 12:
				return []types.Log{
					{BlockNumber: 12, TxHash: secondHash, Index: 4, Data: uint256(99)},
				}, nil
			default:
				return nil, nil
			}
		},
	}
	var output bytes.Buffer
	logger := &captureLogger{}

	stats, err := Scan(context.Background(), Config{
		Client:   client,
		Contract: contract,
		Topic0:   topic,
		Output:   &output,
		Logger:   logger,
	}, testOptions(10, 0))
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	wantOutput := "block_num\ttx_hash\tlog_index\tnew_len\n" +
		"10\t" + firstHash.Hex() + "\t3\t42\n" +
		"12\t" + secondHash.Hex() + "\t4\t99\n"
	if output.String() != wantOutput {
		t.Fatalf("output:\n%s\nwant:\n%s", output.String(), wantOutput)
	}
	if !reflect.DeepEqual(ranges, [][2]uint64{{10, 11}, {12, 12}}) {
		t.Fatalf("ranges = %v", ranges)
	}
	wantStats := Stats{
		FromBlock:     10,
		ToBlock:       12,
		BlocksScanned: 3,
		Events:        2,
		RemovedLogs:   1,
	}
	if stats != wantStats {
		t.Fatalf("stats = %+v, want %+v", stats, wantStats)
	}
	if client.headCalls != 1 {
		t.Fatalf("BlockNumber calls = %d, want 1", client.headCalls)
	}
	if !strings.Contains(logger.String(), "Done. on_chain_events=2 blocks=10..12") {
		t.Fatalf("logger output = %q", logger.String())
	}
}

func TestScanDatabaseCrossCheck(t *testing.T) {
	contract := common.HexToAddress("0x4000000000000000000000000000000000000004")
	topic := common.HexToHash("0xabc")
	inDBHash := common.HexToHash("0x11")
	missingHash := common.HexToHash("0x22")
	keySourceCalled := false
	keySource := KeySourceFunc(func(ctx context.Context) (map[EventKey]struct{}, error) {
		if ctx == nil {
			t.Fatal("key source received nil context")
		}
		keySourceCalled = true
		return map[EventKey]struct{}{
			{TxHash: inDBHash, LogIndex: 1}: {},
		}, nil
	})
	client := &fakeClient{
		filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
			return []types.Log{
				{BlockNumber: 20, TxHash: inDBHash, Index: 1, Data: uint256(5)},
				{BlockNumber: 20, TxHash: missingHash, Index: 2, Data: uint256(6)},
			}, nil
		},
	}
	var output bytes.Buffer
	logger := &captureLogger{}

	stats, err := Scan(context.Background(), Config{
		Client:    client,
		Contract:  contract,
		Topic0:    topic,
		KeySource: keySource,
		Output:    &output,
		Logger:    logger,
	}, testOptions(20, 20))
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	if !keySourceCalled {
		t.Fatal("key source was not called")
	}
	wantOutput := "block_num\ttx_hash\tlog_index\tnew_len\tin_db\n" +
		"20\t" + inDBHash.Hex() + "\t1\t5\tyes\n" +
		"20\t" + missingHash.Hex() + "\t2\t6\tNO\n"
	if output.String() != wantOutput {
		t.Fatalf("output:\n%s\nwant:\n%s", output.String(), wantOutput)
	}
	if stats.Events != 2 || stats.InDB != 1 || stats.MissingFromDB != 1 {
		t.Fatalf("stats = %+v", stats)
	}
	if !strings.Contains(logger.String(), "MISSING_FROM_DB=1") {
		t.Fatalf("logger output = %q", logger.String())
	}
}

func TestScanNilKeyMapStillEnablesCrossCheck(t *testing.T) {
	hash := common.HexToHash("0x33")
	client := &fakeClient{
		filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
			return []types.Log{{BlockNumber: 1, TxHash: hash, Index: 2, Data: uint256(7)}}, nil
		},
	}
	var output bytes.Buffer
	_, err := Scan(context.Background(), Config{
		Client: client,
		KeySource: KeySourceFunc(func(context.Context) (map[EventKey]struct{}, error) {
			return nil, nil
		}),
		Output: &output,
	}, testOptions(1, 1))
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	if !strings.Contains(output.String(), "in_db\n") || !strings.Contains(output.String(), "\tNO\n") {
		t.Fatalf("output = %q, want enabled empty DB cross-check", output.String())
	}
}

func TestScanCancellationDuringFilterLogs(t *testing.T) {
	contract := common.HexToAddress("0x4000000000000000000000000000000000000004")
	ctx, cancel := context.WithCancel(context.Background())
	entered := make(chan struct{})
	client := &fakeClient{
		filter: func(ctx context.Context, _ ethereum.FilterQuery) ([]types.Log, error) {
			close(entered)
			<-ctx.Done()
			return nil, ctx.Err()
		},
	}
	go func() {
		<-entered
		cancel()
	}()

	stats, err := Scan(ctx, Config{
		Client:   client,
		Contract: contract,
		Topic0:   common.HexToHash("0xabc"),
		Output:   io.Discard,
	}, testOptions(30, 30))
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context canceled", err)
	}
	if stats.FilterErrors != 1 || stats.BlocksScanned != 0 {
		t.Fatalf("stats = %+v", stats)
	}
}

func TestScanErrors(t *testing.T) {
	contract := common.HexToAddress("0x4000000000000000000000000000000000000004")
	topic := common.HexToHash("0xabc")

	t.Run("head lookup", func(t *testing.T) {
		boom := errors.New("node down")
		stats, err := Scan(context.Background(), Config{
			Client: &fakeClient{headErr: boom}, Contract: contract, Topic0: topic, Output: io.Discard,
		}, testOptions(1, 0))
		if !errors.Is(err, boom) || stats.BlocksScanned != 0 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("key source", func(t *testing.T) {
		boom := errors.New("database down")
		stats, err := Scan(context.Background(), Config{
			Client:   &fakeClient{},
			Contract: contract,
			Topic0:   topic,
			KeySource: KeySourceFunc(func(context.Context) (map[EventKey]struct{}, error) {
				return nil, boom
			}),
			Output: io.Discard,
		}, testOptions(1, 1))
		if !errors.Is(err, boom) || stats.BlocksScanned != 0 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("malformed event data", func(t *testing.T) {
		client := &fakeClient{
			filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
				return []types.Log{{BlockNumber: 1, Data: []byte{1}}}, nil
			},
		}
		stats, err := Scan(context.Background(), Config{
			Client: client, Contract: contract, Topic0: topic, Output: io.Discard,
		}, testOptions(1, 1))
		if err == nil || !strings.Contains(err.Error(), "want 32") {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
		if stats.Events != 0 || stats.BlocksScanned != 0 {
			t.Fatalf("partial stats = %+v", stats)
		}
	})

	t.Run("output failure", func(t *testing.T) {
		client := &fakeClient{}
		stats, err := Scan(context.Background(), Config{
			Client: client, Contract: contract, Topic0: topic, Output: failingWriter{},
		}, testOptions(1, 1))
		if err == nil || !strings.Contains(err.Error(), "write header") {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("event output failure", func(t *testing.T) {
		client := &fakeClient{
			filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
				return []types.Log{{BlockNumber: 1, Data: uint256(1)}}, nil
			},
		}
		writer := &failNthWriter{failAt: 2}
		stats, err := Scan(context.Background(), Config{
			Client: client, Contract: contract, Topic0: topic, Output: writer,
		}, testOptions(1, 1))
		if err == nil || !strings.Contains(err.Error(), "write event") {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
		if stats.Events != 1 || stats.BlocksScanned != 0 {
			t.Fatalf("partial stats = %+v", stats)
		}
	})
}

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) {
	return 0, errors.New("writer closed")
}

type failNthWriter struct {
	writes int
	failAt int
}

func (w *failNthWriter) Write(data []byte) (int, error) {
	w.writes++
	if w.writes == w.failAt {
		return 0, errors.New("writer closed")
	}
	return len(data), nil
}

func TestScanCancellationCheckpoints(t *testing.T) {
	t.Run("before head lookup", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		client := &fakeClient{}
		stats, err := Scan(ctx, Config{Client: client, Output: io.Discard}, testOptions(1, 0))
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
		if stats.BlocksScanned != 0 || client.headCalls != 0 {
			t.Fatalf("stats=%+v headCalls=%d", stats, client.headCalls)
		}
	})

	t.Run("after key source", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		keySource := KeySourceFunc(func(context.Context) (map[EventKey]struct{}, error) {
			cancel()
			return map[EventKey]struct{}{}, nil
		})
		client := &fakeClient{}
		stats, err := Scan(ctx, Config{
			Client: client, KeySource: keySource, Output: io.Discard,
		}, testOptions(1, 1))
		if !errors.Is(err, context.Canceled) || stats.BlocksScanned != 0 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})
}

func TestScanReportsAdaptiveBatchProgress(t *testing.T) {
	calls := 0
	progressCalls := 0
	client := &fakeClient{
		filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
			calls++
			if calls == 1 {
				return nil, errors.New("range too large")
			}
			return nil, nil
		},
	}
	logger := &captureLogger{}
	opts := testOptions(1, 1)
	opts.OnProgress = func(context.Context, logscan.Progress) error {
		progressCalls++
		return nil
	}

	stats, err := Scan(context.Background(), Config{
		Client: client, Output: io.Discard, Logger: logger,
	}, opts)
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	if stats.FilterErrors != 1 || calls != 2 || progressCalls != 2 {
		t.Fatalf("calls=%d progress=%d stats=%+v", calls, progressCalls, stats)
	}
	if !strings.Contains(logger.String(), "Reducing batch to 1 blocks") {
		t.Fatalf("logger output = %q", logger.String())
	}
}

func TestScanValidatesConfigAndOptions(t *testing.T) {
	client := &fakeClient{}
	validCfg := Config{Client: client, Output: io.Discard}
	validOpts := testOptions(1, 1)

	tests := []struct {
		name string
		cfg  Config
		opts Options
	}{
		{name: "nil client", cfg: Config{Output: io.Discard}, opts: validOpts},
		{name: "nil output", cfg: Config{Client: client}, opts: validOpts},
		{name: "zero initial batch", cfg: validCfg, opts: Options{
			FromBlock: 1, ToBlock: 1, MinBatch: 1, RetryDelay: time.Second,
		}},
		{name: "zero minimum batch", cfg: validCfg, opts: Options{
			FromBlock: 1, ToBlock: 1, InitialBatch: 1, RetryDelay: time.Second,
		}},
		{name: "initial below minimum", cfg: validCfg, opts: Options{
			FromBlock: 1, ToBlock: 1, InitialBatch: 1, MinBatch: 2, RetryDelay: time.Second,
		}},
		{name: "zero delay", cfg: validCfg, opts: Options{
			FromBlock: 1, ToBlock: 1, InitialBatch: 1, MinBatch: 1,
		}},
		{name: "reversed explicit range", cfg: validCfg, opts: testOptions(2, 1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := Scan(context.Background(), tt.cfg, tt.opts)
			if err == nil {
				t.Fatal("Scan succeeded, want validation error")
			}
			if stats.BlocksScanned != 0 || stats.Events != 0 {
				t.Fatalf("stats = %+v", stats)
			}
		})
	}
}
