package txcollector

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/fs"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer/logscan"
	"github.com/PredictionExplorer/augur-explorer/internal/toolutil"
)

type txResponse struct {
	tx      *types.Transaction
	pending bool
	err     error
}

type receiptResponse struct {
	receipt *types.Receipt
	err     error
}

type fakeCollectorClient struct {
	filter func(context.Context, ethereum.FilterQuery) ([]types.Log, error)
	tx     func(context.Context, common.Hash) (*types.Transaction, bool, error)
	rcpt   func(context.Context, common.Hash) (*types.Receipt, error)

	mu           sync.Mutex
	txOrder      []common.Hash
	receiptCalls []common.Hash
}

func (f *fakeCollectorClient) FilterLogs(
	ctx context.Context,
	query ethereum.FilterQuery,
) ([]types.Log, error) {
	if f.filter == nil {
		return nil, nil
	}
	return f.filter(ctx, query)
}

func (f *fakeCollectorClient) TransactionByHash(
	ctx context.Context,
	hash common.Hash,
) (*types.Transaction, bool, error) {
	f.mu.Lock()
	f.txOrder = append(f.txOrder, hash)
	f.mu.Unlock()
	if f.tx == nil {
		return nil, false, errors.New("unexpected TransactionByHash")
	}
	return f.tx(ctx, hash)
}

func (f *fakeCollectorClient) TransactionReceipt(
	ctx context.Context,
	hash common.Hash,
) (*types.Receipt, error) {
	f.mu.Lock()
	f.receiptCalls = append(f.receiptCalls, hash)
	f.mu.Unlock()
	if f.rcpt == nil {
		return nil, errors.New("unexpected TransactionReceipt")
	}
	return f.rcpt(ctx, hash)
}

func (f *fakeCollectorClient) transactionOrder() []common.Hash {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]common.Hash(nil), f.txOrder...)
}

type testLogger struct {
	mu    sync.Mutex
	lines []string
}

func (l *testLogger) Printf(format string, args ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lines = append(l.lines, fmt.Sprintf(format, args...))
}

func (l *testLogger) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return strings.Join(l.lines, "\n")
}

type hookLogger struct {
	hook func(string)
}

func (l hookLogger) Printf(format string, args ...any) {
	if l.hook != nil {
		l.hook(fmt.Sprintf(format, args...))
	}
}

func testRunOptions(from, to uint64) RunOptions {
	return RunOptions{
		FromBlock:    from,
		ToBlock:      to,
		InitialBatch: to - from + 1,
		MinBatch:     1,
		RetryDelay:   time.Millisecond,
	}
}

func fixtureTx(nonce uint64) *types.Transaction {
	return types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: big.NewInt(1),
		Gas:      21_000,
		To:       ptr(common.HexToAddress("0x1000000000000000000000000000000000000001")),
		Value:    big.NewInt(0),
	})
}

func ptr[T any](value T) *T {
	return &value
}

func fixtureReceipt(tx *types.Transaction, logs ...*types.Log) *types.Receipt {
	return &types.Receipt{
		Type:              tx.Type(),
		Status:            types.ReceiptStatusSuccessful,
		CumulativeGasUsed: 21_000,
		Logs:              logs,
		TxHash:            tx.Hash(),
	}
}

func TestRunHappyPathDeduplicatesAndUsesDeterministicOrder(t *testing.T) {
	dir := t.TempDir()
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")
	txAtBlockTwo := fixtureTx(2)
	txAtBlockOne := fixtureTx(1)
	transactions := map[common.Hash]*types.Transaction{
		txAtBlockTwo.Hash(): txAtBlockTwo,
		txAtBlockOne.Hash(): txAtBlockOne,
	}
	receipts := map[common.Hash]*types.Receipt{
		txAtBlockTwo.Hash(): fixtureReceipt(txAtBlockTwo),
		txAtBlockOne.Hash(): fixtureReceipt(txAtBlockOne),
	}
	client := &fakeCollectorClient{
		filter: func(_ context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
			if got := query.Addresses; !reflect.DeepEqual(got, []common.Address{contract}) {
				t.Fatalf("query addresses = %v", got)
			}
			return []types.Log{
				{Address: contract, BlockNumber: 2, TxHash: txAtBlockTwo.Hash()},
				{Address: contract, BlockNumber: 1, TxHash: txAtBlockOne.Hash()},
				{Address: contract, BlockNumber: 2, TxHash: txAtBlockTwo.Hash(), Index: 1},
			}, nil
		},
		tx: func(_ context.Context, hash common.Hash) (*types.Transaction, bool, error) {
			return transactions[hash], false, nil
		},
		rcpt: func(_ context.Context, hash common.Hash) (*types.Receipt, error) {
			return receipts[hash], nil
		},
	}

	stats, err := Run(context.Background(), Config{
		Client:    client,
		OutputDir: dir,
		Contracts: []common.Address{contract},
	}, testRunOptions(1, 3))
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	wantStats := RunStats{
		BlocksScanned:  3,
		LogsSeen:       3,
		TxUnique:       2,
		TxWritten:      2,
		ReceiptWritten: 2,
	}
	if stats != wantStats {
		t.Fatalf("stats = %+v, want %+v", stats, wantStats)
	}
	if got, want := client.transactionOrder(), []common.Hash{txAtBlockOne.Hash(), txAtBlockTwo.Hash()}; !reflect.DeepEqual(got, want) {
		t.Fatalf("transaction fetch order = %v, want %v", got, want)
	}

	for block, tx := range map[uint64]*types.Transaction{1: txAtBlockOne, 2: txAtBlockTwo} {
		txData, err := os.ReadFile(toolutil.TxRLPPath(dir, block, tx.Hash().Hex()))
		if err != nil {
			t.Fatalf("read tx backup: %v", err)
		}
		var decodedTx types.Transaction
		if err := rlp.DecodeBytes(txData, &decodedTx); err != nil {
			t.Fatalf("decode tx backup: %v", err)
		}
		if decodedTx.Hash() != tx.Hash() {
			t.Fatalf("decoded hash = %s, want %s", decodedTx.Hash(), tx.Hash())
		}
		receiptData, err := os.ReadFile(toolutil.ReceiptRLPPath(dir, block, tx.Hash().Hex()))
		if err != nil {
			t.Fatalf("read receipt backup: %v", err)
		}
		if _, legacy, err := toolutil.TryDecodeReceiptRLP(receiptData); err != nil || legacy {
			t.Fatalf("decode receipt backup: legacy=%v err=%v", legacy, err)
		}
	}
}

func TestRunSortsSameBlockTransactionsByHash(t *testing.T) {
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")
	first := fixtureTx(3)
	second := fixtureTx(4)
	transactions := map[common.Hash]*types.Transaction{
		first.Hash():  first,
		second.Hash(): second,
	}
	client := &fakeCollectorClient{
		filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
			return []types.Log{
				{Address: contract, BlockNumber: 5, TxHash: second.Hash()},
				{Address: contract, BlockNumber: 5, TxHash: first.Hash()},
			}, nil
		},
		tx: func(_ context.Context, hash common.Hash) (*types.Transaction, bool, error) {
			return transactions[hash], false, nil
		},
		rcpt: func(_ context.Context, hash common.Hash) (*types.Receipt, error) {
			return fixtureReceipt(transactions[hash]), nil
		},
	}

	_, err := Run(context.Background(), Config{
		Client: client, OutputDir: t.TempDir(), Contracts: []common.Address{contract},
	}, testRunOptions(5, 5))
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	want := []common.Hash{first.Hash(), second.Hash()}
	sort.Slice(want, func(i, j int) bool {
		return bytes.Compare(want[i][:], want[j][:]) < 0
	})
	if got := client.transactionOrder(); !reflect.DeepEqual(got, want) {
		t.Fatalf("transaction order = %v, want %v", got, want)
	}
}

func TestRunExistingBackupFiles(t *testing.T) {
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")
	tx := fixtureTx(10)
	logEntry := types.Log{Address: contract, BlockNumber: 7, TxHash: tx.Hash()}

	t.Run("complete pair is skipped without RPC fetch", func(t *testing.T) {
		dir := t.TempDir()
		writeBackupFixture(t, dir, 7, tx, fixtureReceipt(tx))
		client := &fakeCollectorClient{
			filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
				return []types.Log{logEntry}, nil
			},
		}

		stats, err := Run(context.Background(), Config{
			Client: client, OutputDir: dir, Contracts: []common.Address{contract},
		}, testRunOptions(7, 7))
		if err != nil {
			t.Fatalf("Run: %v", err)
		}
		if stats.TxSkippedExists != 1 || stats.TxWritten != 0 || stats.ReceiptWritten != 0 {
			t.Fatalf("stats = %+v", stats)
		}
		if len(client.transactionOrder()) != 0 {
			t.Fatal("existing complete pair triggered an RPC fetch")
		}
	})

	t.Run("partial pair writes only missing receipt", func(t *testing.T) {
		dir := t.TempDir()
		txPath := toolutil.TxRLPPath(dir, 7, tx.Hash().Hex())
		writeTxFixture(t, dir, 7, tx.Hash().Hex(), tx)
		// #nosec G304 -- txPath is deterministic under t.TempDir.
		existing, err := os.ReadFile(txPath)
		if err != nil {
			t.Fatal(err)
		}
		client := &fakeCollectorClient{
			filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
				return []types.Log{logEntry}, nil
			},
			tx: func(context.Context, common.Hash) (*types.Transaction, bool, error) {
				return tx, false, nil
			},
			rcpt: func(context.Context, common.Hash) (*types.Receipt, error) {
				return fixtureReceipt(tx), nil
			},
		}

		stats, err := Run(context.Background(), Config{
			Client: client, OutputDir: dir, Contracts: []common.Address{contract},
		}, testRunOptions(7, 7))
		if err != nil {
			t.Fatalf("Run: %v", err)
		}
		if stats.TxWritten != 0 || stats.ReceiptWritten != 1 || stats.TxSkippedExists != 0 {
			t.Fatalf("stats = %+v", stats)
		}
		// #nosec G304 -- txPath is a deterministic path under t.TempDir.
		if got, err := os.ReadFile(txPath); err != nil || !bytes.Equal(got, existing) {
			t.Fatalf("existing tx changed: %q err=%v", got, err)
		}
	})

	t.Run("corrupt pair is detected and replaced", func(t *testing.T) {
		dir := t.TempDir()
		txPath := toolutil.TxRLPPath(dir, 7, tx.Hash().Hex())
		receiptPath := toolutil.ReceiptRLPPath(dir, 7, tx.Hash().Hex())
		writeTxFixture(t, dir, 7, tx.Hash().Hex(), fixtureTx(11))
		if err := writeFileAtomic(receiptPath, []byte("corrupt receipt")); err != nil {
			t.Fatal(err)
		}
		client := &fakeCollectorClient{
			filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
				return []types.Log{logEntry}, nil
			},
			tx: func(context.Context, common.Hash) (*types.Transaction, bool, error) {
				return tx, false, nil
			},
			rcpt: func(context.Context, common.Hash) (*types.Receipt, error) {
				return fixtureReceipt(tx), nil
			},
		}

		stats, err := Run(context.Background(), Config{
			Client: client, OutputDir: dir, Contracts: []common.Address{contract},
		}, testRunOptions(7, 7))
		if err != nil {
			t.Fatalf("Run: %v", err)
		}
		if stats.InvalidBackups != 2 || stats.TxWritten != 1 || stats.ReceiptWritten != 1 ||
			stats.TxSkippedExists != 0 {
			t.Fatalf("stats = %+v", stats)
		}
		if err := validateTransactionBackup(txPath, tx.Hash()); err != nil {
			t.Fatalf("replacement transaction: %v", err)
		}
		if err := validateReceiptBackup(receiptPath); err != nil {
			t.Fatalf("replacement receipt: %v", err)
		}
	})

	t.Run("unavailable corrupt pair is a blocking backup error", func(t *testing.T) {
		dir := t.TempDir()
		txPath := toolutil.TxRLPPath(dir, 7, tx.Hash().Hex())
		receiptPath := toolutil.ReceiptRLPPath(dir, 7, tx.Hash().Hex())
		if err := writeFileAtomic(txPath, []byte("corrupt transaction")); err != nil {
			t.Fatal(err)
		}
		if err := writeFileAtomic(receiptPath, []byte("corrupt receipt")); err != nil {
			t.Fatal(err)
		}
		client := &fakeCollectorClient{
			filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
				return []types.Log{logEntry}, nil
			},
			tx: func(context.Context, common.Hash) (*types.Transaction, bool, error) {
				return nil, false, errors.New("not found")
			},
		}

		stats, err := Run(context.Background(), Config{
			Client: client, OutputDir: dir, Contracts: []common.Address{contract},
		}, testRunOptions(7, 7))
		if err == nil || !strings.Contains(err.Error(), "repair invalid backup") {
			t.Fatalf("Run error = %v", err)
		}
		if stats.InvalidBackups != 2 || stats.BackupErrors != 1 || stats.TxMissingNode != 1 {
			t.Fatalf("stats = %+v", stats)
		}
	})
}

func TestRunTransactionWarningsDoNotAbort(t *testing.T) {
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")
	tx := fixtureTx(20)
	cases := []struct {
		name         string
		txResponse   txResponse
		rcptResponse receiptResponse
		wantMissing  uint64
		wantErrors   uint64
	}{
		{
			name:        "transaction missing",
			txResponse:  txResponse{err: errors.New("not found")},
			wantMissing: 1,
		},
		{
			name:       "transaction pending",
			txResponse: txResponse{tx: tx, pending: true},
			wantErrors: 1,
		},
		{
			name:       "transaction RPC error",
			txResponse: txResponse{err: errors.New("connection reset")},
			wantErrors: 1,
		},
		{
			name:       "empty transaction response",
			txResponse: txResponse{},
			wantErrors: 1,
		},
		{
			name:         "receipt missing",
			txResponse:   txResponse{tx: tx},
			rcptResponse: receiptResponse{err: errors.New("header not found")},
			wantMissing:  1,
		},
		{
			name:         "receipt RPC error",
			txResponse:   txResponse{tx: tx},
			rcptResponse: receiptResponse{err: errors.New("rate limited")},
			wantErrors:   1,
		},
		{
			name:       "empty receipt response",
			txResponse: txResponse{tx: tx},
			wantErrors: 1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			logger := &testLogger{}
			client := &fakeCollectorClient{
				filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
					return []types.Log{{Address: contract, BlockNumber: 9, TxHash: tx.Hash()}}, nil
				},
				tx: func(context.Context, common.Hash) (*types.Transaction, bool, error) {
					r := tt.txResponse
					return r.tx, r.pending, r.err
				},
				rcpt: func(context.Context, common.Hash) (*types.Receipt, error) {
					return tt.rcptResponse.receipt, tt.rcptResponse.err
				},
			}

			stats, err := Run(context.Background(), Config{
				Client: client, OutputDir: dir, Contracts: []common.Address{contract}, Logger: logger,
			}, testRunOptions(9, 9))
			if err != nil {
				t.Fatalf("Run: %v", err)
			}
			if stats.TxMissingNode != tt.wantMissing || stats.TxFetchErrors != tt.wantErrors {
				t.Fatalf("stats = %+v, want missing=%d errors=%d", stats, tt.wantMissing, tt.wantErrors)
			}
			if logger.String() == "" {
				t.Fatal("warning/progress logger received no output")
			}
		})
	}
}

func TestRunCancellationDuringTransactionFetch(t *testing.T) {
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")
	tx := fixtureTx(30)
	ctx, cancel := context.WithCancel(context.Background())
	entered := make(chan struct{})
	client := &fakeCollectorClient{
		filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
			return []types.Log{{Address: contract, BlockNumber: 4, TxHash: tx.Hash()}}, nil
		},
		tx: func(ctx context.Context, _ common.Hash) (*types.Transaction, bool, error) {
			close(entered)
			<-ctx.Done()
			return nil, false, ctx.Err()
		},
	}
	go func() {
		<-entered
		cancel()
	}()

	stats, err := Run(ctx, Config{
		Client: client, OutputDir: t.TempDir(), Contracts: []common.Address{contract},
	}, testRunOptions(4, 4))
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context canceled", err)
	}
	if stats.TxUnique != 1 || stats.TxFetchErrors != 0 || stats.TxMissingNode != 0 {
		t.Fatalf("stats = %+v", stats)
	}
}

func TestRunCancellationAndSetupErrors(t *testing.T) {
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")

	t.Run("pre-canceled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		stats, err := Run(ctx, Config{
			Client: &fakeCollectorClient{}, OutputDir: t.TempDir(), Contracts: []common.Address{contract},
		}, testRunOptions(1, 1))
		if !errors.Is(err, context.Canceled) || stats != (RunStats{}) {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("output directory creation", func(t *testing.T) {
		root := t.TempDir()
		parentFile := filepath.Join(root, "not-a-directory")
		if err := os.WriteFile(parentFile, []byte("x"), 0o600); err != nil {
			t.Fatal(err)
		}
		stats, err := Run(context.Background(), Config{
			Client:    &fakeCollectorClient{},
			OutputDir: filepath.Join(parentFile, "backup"),
			Contracts: []common.Address{contract},
		}, testRunOptions(1, 1))
		if err == nil || !strings.Contains(err.Error(), "create output directory") {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("progress callback", func(t *testing.T) {
		boom := errors.New("progress failed")
		opts := testRunOptions(1, 1)
		opts.OnProgress = func(context.Context, logscan.Progress) error {
			return boom
		}
		stats, err := Run(context.Background(), Config{
			Client: &fakeCollectorClient{}, OutputDir: t.TempDir(), Contracts: []common.Address{contract},
		}, opts)
		if !errors.Is(err, boom) || stats.BlocksScanned != 0 {
			t.Fatalf("stats=%+v err=%v", stats, err)
		}
	})

	t.Run("before transaction loop", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		tx := fixtureTx(31)
		client := &fakeCollectorClient{
			filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
				return []types.Log{{BlockNumber: 1, TxHash: tx.Hash()}}, nil
			},
		}
		logger := hookLogger{hook: func(line string) {
			if strings.HasPrefix(line, "Collected ") {
				cancel()
			}
		}}
		stats, err := Run(ctx, Config{
			Client: client, OutputDir: t.TempDir(), Contracts: []common.Address{contract}, Logger: logger,
		}, testRunOptions(1, 1))
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
		if stats.TxUnique != 1 || len(client.transactionOrder()) != 0 {
			t.Fatalf("stats=%+v txOrder=%v", stats, client.transactionOrder())
		}
	})
}

func TestRunReportsFilterRetriesFromLogscan(t *testing.T) {
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")
	calls := 0
	var sleeps []time.Duration
	client := &fakeCollectorClient{
		filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
			calls++
			if calls <= 2 {
				return nil, errors.New("too many results")
			}
			return nil, nil
		},
	}
	opts := RunOptions{
		FromBlock:    0,
		ToBlock:      1,
		InitialBatch: 2,
		MinBatch:     1,
		RetryDelay:   5 * time.Second,
		Sleep: func(context.Context, time.Duration) error {
			sleeps = append(sleeps, 5*time.Second)
			return nil
		},
	}
	stats, err := Run(context.Background(), Config{
		Client: client, OutputDir: t.TempDir(), Contracts: []common.Address{contract},
	}, opts)
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if stats.FilterLogErrors != 2 || stats.BlocksScanned != 2 {
		t.Fatalf("stats = %+v", stats)
	}
	if len(sleeps) != 1 {
		t.Fatalf("sleep calls = %d, want 1", len(sleeps))
	}
}

func TestStoreTransactionPairErrorPaths(t *testing.T) {
	tx := fixtureTx(40)
	ref := txRef{hash: tx.Hash(), blockNum: 8}

	t.Run("transaction stat", func(t *testing.T) {
		dir := t.TempDir()
		blockDir := filepath.Join(dir, "8")
		if err := os.Mkdir(blockDir, 0o750); err != nil {
			t.Fatal(err)
		}
		txPath := toolutil.TxRLPPath(dir, ref.blockNum, ref.hash.Hex())
		if err := os.Symlink(filepath.Base(txPath), txPath); err != nil {
			t.Fatal(err)
		}
		err := storeTransactionPair(context.Background(), Config{
			Client: &fakeCollectorClient{}, OutputDir: dir,
		}, ref, &RunStats{})
		if err == nil || !strings.Contains(err.Error(), "stat tx backup") {
			t.Fatalf("error = %v", err)
		}
	})

	t.Run("receipt stat", func(t *testing.T) {
		dir := t.TempDir()
		txPath := toolutil.TxRLPPath(dir, ref.blockNum, ref.hash.Hex())
		receiptPath := toolutil.ReceiptRLPPath(dir, ref.blockNum, ref.hash.Hex())
		if err := os.MkdirAll(filepath.Dir(txPath), 0o750); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(txPath, []byte("existing"), 0o600); err != nil {
			t.Fatal(err)
		}
		if err := os.Symlink(filepath.Base(receiptPath), receiptPath); err != nil {
			t.Fatal(err)
		}
		err := storeTransactionPair(context.Background(), Config{
			Client: &fakeCollectorClient{}, OutputDir: dir,
		}, ref, &RunStats{})
		if err == nil || !strings.Contains(err.Error(), "stat receipt backup") {
			t.Fatalf("error = %v", err)
		}
	})

	t.Run("receipt cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		client := &fakeCollectorClient{
			tx: func(context.Context, common.Hash) (*types.Transaction, bool, error) {
				return tx, false, nil
			},
			rcpt: func(context.Context, common.Hash) (*types.Receipt, error) {
				cancel()
				return nil, context.Canceled
			},
		}
		err := storeTransactionPair(ctx, Config{
			Client: client, OutputDir: t.TempDir(),
		}, ref, &RunStats{})
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v, want context canceled", err)
		}
	})

	t.Run("transaction encoding", func(t *testing.T) {
		client := &fakeCollectorClient{
			tx: func(context.Context, common.Hash) (*types.Transaction, bool, error) {
				return new(types.Transaction), false, nil
			},
			rcpt: func(context.Context, common.Hash) (*types.Receipt, error) {
				return fixtureReceipt(tx), nil
			},
		}
		err := storeTransactionPair(context.Background(), Config{
			Client: client, OutputDir: t.TempDir(),
		}, ref, &RunStats{})
		if err == nil || !strings.Contains(err.Error(), "encode transaction RLP") {
			t.Fatalf("error = %v", err)
		}
	})

	t.Run("receipt encoding", func(t *testing.T) {
		boom := errors.New("receipt encoding failed")
		client := &fakeCollectorClient{
			tx: func(context.Context, common.Hash) (*types.Transaction, bool, error) {
				return tx, false, nil
			},
			rcpt: func(context.Context, common.Hash) (*types.Receipt, error) {
				return fixtureReceipt(tx), nil
			},
		}
		encoders := defaultPairEncoders
		encoders.receipt = func(*types.Receipt) ([]byte, error) {
			return nil, boom
		}
		err := storeTransactionPairWithEncoders(context.Background(), Config{
			Client: client, OutputDir: t.TempDir(),
		}, ref, &RunStats{}, encoders)
		if !errors.Is(err, boom) || !strings.Contains(err.Error(), "encode receipt RLP") {
			t.Fatalf("error = %v", err)
		}
	})

	t.Run("receipt write", func(t *testing.T) {
		dir := t.TempDir()
		txPath := toolutil.TxRLPPath(dir, ref.blockNum, ref.hash.Hex())
		receiptPath := toolutil.ReceiptRLPPath(dir, ref.blockNum, ref.hash.Hex())
		if err := os.MkdirAll(filepath.Dir(txPath), 0o750); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(txPath, []byte("existing"), 0o600); err != nil {
			t.Fatal(err)
		}
		client := &fakeCollectorClient{
			tx: func(context.Context, common.Hash) (*types.Transaction, bool, error) {
				return tx, false, nil
			},
			rcpt: func(context.Context, common.Hash) (*types.Receipt, error) {
				if err := os.Mkdir(receiptPath, 0o750); err != nil {
					t.Fatalf("creating receipt destination directory: %v", err)
				}
				return fixtureReceipt(tx), nil
			},
		}
		err := storeTransactionPair(context.Background(), Config{
			Client: client, OutputDir: dir,
		}, ref, &RunStats{})
		if err == nil || !strings.Contains(err.Error(), "write receipt backup") {
			t.Fatalf("error = %v", err)
		}
		assertNoAtomicTemps(t, filepath.Dir(receiptPath))
	})
}

func TestRunReturnsTransactionWriteFailureAfterAttempt(t *testing.T) {
	dir := t.TempDir()
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")
	tx := fixtureTx(41)
	txPath := toolutil.TxRLPPath(dir, 9, tx.Hash().Hex())
	logger := &testLogger{}
	client := &fakeCollectorClient{
		filter: func(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
			return []types.Log{{Address: contract, BlockNumber: 9, TxHash: tx.Hash()}}, nil
		},
		tx: func(context.Context, common.Hash) (*types.Transaction, bool, error) {
			return tx, false, nil
		},
		rcpt: func(context.Context, common.Hash) (*types.Receipt, error) {
			if err := os.MkdirAll(txPath, 0o750); err != nil {
				t.Fatalf("creating transaction destination directory: %v", err)
			}
			return fixtureReceipt(tx), nil
		},
	}

	stats, err := Run(context.Background(), Config{
		Client: client, OutputDir: dir, Contracts: []common.Address{contract}, Logger: logger,
	}, testRunOptions(9, 9))
	if err == nil || !strings.Contains(err.Error(), "1 transaction backup(s) failed") {
		t.Fatalf("Run error = %v", err)
	}
	if stats.TxWritten != 0 || stats.BackupErrors != 1 ||
		!strings.Contains(logger.String(), "write transaction backup") {
		t.Fatalf("stats=%+v logger=%q", stats, logger.String())
	}
	assertNoAtomicTemps(t, filepath.Dir(txPath))
}

func TestMissingOnNodeClassifierVariants(t *testing.T) {
	tests := []struct {
		message string
		want    bool
	}{
		{message: "", want: false},
		{message: "not found", want: true},
		{message: "MISSING historical transaction", want: true},
		{message: "unknown transaction", want: true},
		{message: "transaction indexing is in progress", want: true},
		{message: "header not found", want: true},
		{message: "connection reset", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.message, func(t *testing.T) {
			var err error
			if tt.message != "" {
				err = errors.New(tt.message)
			}
			if got := isMissingOnNodeError(err); got != tt.want {
				t.Fatalf("isMissingOnNodeError(%v) = %v, want %v", err, got, tt.want)
			}
		})
	}
}

func TestFileExistsErrors(t *testing.T) {
	parent := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(parent, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := fileExists(filepath.Join(parent, "child")); err == nil {
		t.Fatal("fileExists succeeded below a regular file")
	}
}

func TestEnsureDurableDirectoryCreatesNestedRootIdempotently(t *testing.T) {
	path := filepath.Join(t.TempDir(), "first", "second")
	if err := ensureDurableDirectory(path, 0o750); err != nil {
		t.Fatalf("first creation: %v", err)
	}
	if err := ensureDurableDirectory(path, 0o750); err != nil {
		t.Fatalf("idempotent creation: %v", err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if !info.IsDir() || info.Mode().Perm() != 0o750 {
		t.Fatalf("directory mode = %v", info.Mode())
	}
}

func assertNoAtomicTemps(t *testing.T, dir string) {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		if strings.Contains(entry.Name(), ".tmp-") {
			t.Fatalf("temporary file remains: %s", entry.Name())
		}
	}
}

type fakeAtomicFile struct {
	name       string
	writeErr   error
	chmodErr   error
	syncErr    error
	closeErr   error
	closeCalls int
	syncCalls  int
	mode       fs.FileMode
}

func (f *fakeAtomicFile) Name() string {
	return f.name
}

func (f *fakeAtomicFile) Write(data []byte) (int, error) {
	if f.writeErr != nil {
		return 0, f.writeErr
	}
	return len(data), nil
}

func (f *fakeAtomicFile) Chmod(mode fs.FileMode) error {
	f.mode = mode
	return f.chmodErr
}

func (f *fakeAtomicFile) Sync() error {
	f.syncCalls++
	return f.syncErr
}

func (f *fakeAtomicFile) Close() error {
	f.closeCalls++
	return f.closeErr
}

func TestWriteFileAtomicInjectedFailures(t *testing.T) {
	boom := errors.New("injected failure")
	path := filepath.Join("backup", "blob.rlp")

	t.Run("mkdir", func(t *testing.T) {
		ops := successfulAtomicOps(t, &fakeAtomicFile{name: "temp"})
		ops.mkdirAll = func(string, fs.FileMode) error { return boom }
		if err := writeFileAtomicWithOps(path, []byte("data"), ops); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want mkdir error", err)
		}
	})

	t.Run("create temp", func(t *testing.T) {
		ops := successfulAtomicOps(t, &fakeAtomicFile{name: "temp"})
		ops.createTemp = func(string, string) (atomicFile, error) { return nil, boom }
		if err := writeFileAtomicWithOps(path, []byte("data"), ops); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want create error", err)
		}
	})

	t.Run("parent directory sync", func(t *testing.T) {
		ops := successfulAtomicOps(t, &fakeAtomicFile{name: "temp"})
		ops.syncParent = func(string) error { return boom }
		if err := writeFileAtomicWithOps(path, []byte("data"), ops); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want parent sync error", err)
		}
	})

	for _, test := range []struct {
		name           string
		configure      func(*fakeAtomicFile)
		wantCloseCalls int
	}{
		{
			name: "write",
			configure: func(file *fakeAtomicFile) {
				file.writeErr = boom
			},
			wantCloseCalls: 1,
		},
		{
			name: "chmod",
			configure: func(file *fakeAtomicFile) {
				file.chmodErr = boom
			},
			wantCloseCalls: 1,
		},
		{
			name: "close",
			configure: func(file *fakeAtomicFile) {
				file.closeErr = boom
			},
			wantCloseCalls: 2,
		},
		{
			name: "sync",
			configure: func(file *fakeAtomicFile) {
				file.syncErr = boom
			},
			wantCloseCalls: 1,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			file := &fakeAtomicFile{name: "temp"}
			test.configure(file)
			removed := false
			ops := successfulAtomicOps(t, file)
			ops.remove = func(path string) error {
				if path != file.name {
					t.Fatalf("removed path = %q, want %q", path, file.name)
				}
				removed = true
				return nil
			}
			if err := writeFileAtomicWithOps(path, []byte("data"), ops); !errors.Is(err, boom) {
				t.Fatalf("error = %v, want injected error", err)
			}
			if file.closeCalls != test.wantCloseCalls || !removed {
				t.Fatalf("closeCalls=%d removed=%v", file.closeCalls, removed)
			}
		})
	}

	t.Run("rename cleanup", func(t *testing.T) {
		file := &fakeAtomicFile{name: "temp"}
		removed := false
		ops := successfulAtomicOps(t, file)
		ops.rename = func(string, string) error { return boom }
		ops.remove = func(string) error {
			removed = true
			return nil
		}
		if err := writeFileAtomicWithOps(path, []byte("data"), ops); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want rename error", err)
		}
		if file.closeCalls != 1 || !removed {
			t.Fatalf("closeCalls=%d removed=%v", file.closeCalls, removed)
		}
	})

	t.Run("directory sync failure is returned after rename", func(t *testing.T) {
		file := &fakeAtomicFile{name: "temp"}
		ops := successfulAtomicOps(t, file)
		ops.syncDir = func(string) error { return boom }
		if err := writeFileAtomicWithOps(path, []byte("data"), ops); !errors.Is(err, boom) {
			t.Fatalf("error = %v, want directory sync error", err)
		}
	})

	t.Run("success uses secure modes", func(t *testing.T) {
		file := &fakeAtomicFile{name: "temp"}
		ops := successfulAtomicOps(t, file)
		if err := writeFileAtomicWithOps(path, []byte("data"), ops); err != nil {
			t.Fatalf("writeFileAtomicWithOps: %v", err)
		}
		if file.mode != 0o640 || file.syncCalls != 1 || file.closeCalls != 1 {
			t.Fatalf("mode=%o syncCalls=%d closeCalls=%d", file.mode, file.syncCalls, file.closeCalls)
		}
	})
}

func successfulAtomicOps(t *testing.T, file atomicFile) atomicWriteOps {
	t.Helper()
	return atomicWriteOps{
		mkdirAll: func(_ string, mode fs.FileMode) error {
			if mode != 0o750 {
				t.Fatalf("directory mode = %o, want 750", mode)
			}
			return nil
		},
		createTemp: func(_, pattern string) (atomicFile, error) {
			if pattern != ".blob.rlp.tmp-*" {
				t.Fatalf("temp pattern = %q", pattern)
			}
			return file, nil
		},
		remove: func(string) error { return nil },
		rename: func(oldPath, newPath string) error {
			if oldPath != file.Name() || newPath != filepath.Join("backup", "blob.rlp") {
				t.Fatalf("rename %q -> %q", oldPath, newPath)
			}
			return nil
		},
		syncParent: func(path string) error {
			if path != "." {
				t.Fatalf("synced parent directory = %q, want .", path)
			}
			return nil
		},
		syncDir: func(path string) error {
			if path != "backup" {
				t.Fatalf("synced directory = %q, want backup", path)
			}
			return nil
		},
	}
}

func TestWriteFileAtomicModeUniquenessAndCleanup(t *testing.T) {
	t.Run("writes mode and leaves no temp", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "nested", "blob.rlp")
		if err := writeFileAtomic(path, []byte("payload")); err != nil {
			t.Fatalf("writeFileAtomic: %v", err)
		}
		info, err := os.Stat(path)
		if err != nil {
			t.Fatal(err)
		}
		if got := info.Mode().Perm(); got != 0o640 {
			t.Fatalf("mode = %o, want 640", got)
		}
		dirInfo, err := os.Stat(filepath.Dir(path))
		if err != nil {
			t.Fatal(err)
		}
		if got := dirInfo.Mode().Perm(); got != 0o750 {
			t.Fatalf("directory mode = %o, want 750", got)
		}
		entries, err := os.ReadDir(filepath.Dir(path))
		if err != nil {
			t.Fatal(err)
		}
		if len(entries) != 1 || entries[0].Name() != "blob.rlp" {
			t.Fatalf("directory entries = %v", entries)
		}
	})

	t.Run("concurrent writers use unique temp files", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "blob.rlp")
		payloads := [][]byte{[]byte("first"), []byte("second")}
		errorsCh := make(chan error, len(payloads))
		for _, payload := range payloads {
			go func() {
				errorsCh <- writeFileAtomic(path, payload)
			}()
		}
		for range payloads {
			if err := <-errorsCh; err != nil {
				t.Fatalf("writeFileAtomic: %v", err)
			}
		}
		// #nosec G304 -- path is a deterministic path under t.TempDir.
		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(got, payloads[0]) && !bytes.Equal(got, payloads[1]) {
			t.Fatalf("final payload = %q", got)
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}
		if len(entries) != 1 || entries[0].Name() != "blob.rlp" {
			t.Fatalf("temp files remain: %v", entries)
		}
	})

	t.Run("directory creation failure", func(t *testing.T) {
		dir := t.TempDir()
		parentFile := filepath.Join(dir, "not-a-directory")
		if err := os.WriteFile(parentFile, []byte("x"), 0o600); err != nil {
			t.Fatal(err)
		}
		if err := writeFileAtomic(filepath.Join(parentFile, "blob.rlp"), []byte("payload")); err == nil {
			t.Fatal("writeFileAtomic succeeded below a regular file")
		}
		assertNoAtomicTemps(t, dir)
	})

	t.Run("rename failure removes temp file", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "destination")
		if err := os.Mkdir(path, 0o750); err != nil {
			t.Fatal(err)
		}
		if err := writeFileAtomic(path, []byte("payload")); err == nil {
			t.Fatal("writeFileAtomic succeeded over a directory")
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			t.Fatal(err)
		}
		if len(entries) != 1 || entries[0].Name() != "destination" || !entries[0].IsDir() {
			t.Fatalf("cleanup left entries: %v", entries)
		}
	})
}

func TestRunValidatesConfigurationAndOptions(t *testing.T) {
	client := &fakeCollectorClient{}
	contract := common.HexToAddress("0x2000000000000000000000000000000000000002")
	validCfg := Config{Client: client, OutputDir: t.TempDir(), Contracts: []common.Address{contract}}
	validOpts := testRunOptions(1, 2)

	cases := []struct {
		name string
		cfg  Config
		opts RunOptions
	}{
		{name: "nil client", cfg: Config{OutputDir: validCfg.OutputDir, Contracts: validCfg.Contracts}, opts: validOpts},
		{name: "empty output", cfg: Config{Client: client, Contracts: validCfg.Contracts}, opts: validOpts},
		{name: "empty contracts", cfg: Config{Client: client, OutputDir: validCfg.OutputDir}, opts: validOpts},
		{name: "reversed range", cfg: validCfg, opts: RunOptions{
			FromBlock: 3, ToBlock: 2, InitialBatch: 1, MinBatch: 1, RetryDelay: time.Second,
		}},
		{name: "zero batch", cfg: validCfg, opts: RunOptions{
			FromBlock: 1, ToBlock: 2, MinBatch: 1, RetryDelay: time.Second,
		}},
		{name: "zero min", cfg: validCfg, opts: RunOptions{
			FromBlock: 1, ToBlock: 2, InitialBatch: 1, RetryDelay: time.Second,
		}},
		{name: "batch below min", cfg: validCfg, opts: RunOptions{
			FromBlock: 1, ToBlock: 2, InitialBatch: 1, MinBatch: 2, RetryDelay: time.Second,
		}},
		{name: "zero delay", cfg: validCfg, opts: RunOptions{
			FromBlock: 1, ToBlock: 2, InitialBatch: 1, MinBatch: 1,
		}},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			stats, err := Run(context.Background(), tt.cfg, tt.opts)
			if err == nil {
				t.Fatal("Run succeeded, want validation error")
			}
			if stats != (RunStats{}) {
				t.Fatalf("stats = %+v, want zero", stats)
			}
		})
	}
}
