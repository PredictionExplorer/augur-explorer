package archive

import (
	"context"
	"database/sql/driver"
	"errors"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type coverageRepository struct {
	contracts    Contracts
	contractsErr error
	start        uint64
	startErr     error
	archived     []int64
	archivedErr  error
	writer       NodeFillWriter
	prepareErr   error
}

func (r *coverageRepository) ProjectContracts(context.Context, string) (Contracts, error) {
	return r.contracts, r.contractsErr
}

func (r *coverageRepository) ResolveStartBlock(context.Context, Contracts, uint64) (uint64, error) {
	return r.start, r.startErr
}

func (r *coverageRepository) ArchivedBlockNumbers(
	context.Context,
	Contracts,
	uint64,
	uint64,
) ([]int64, error) {
	return r.archived, r.archivedErr
}

func (r *coverageRepository) PrepareWriter(context.Context) (NodeFillWriter, error) {
	return r.writer, r.prepareErr
}

type coverageWriter struct {
	eventExistsFn       func(context.Context, string, int) (bool, error)
	insertEventFn       func(context.Context, EventLog) (bool, error)
	transactionExistsFn func(context.Context, string) (bool, error)
	insertTransactionFn func(context.Context, Transaction) (bool, error)
	blockExistsFn       func(context.Context, int64) (bool, error)
	insertBlockFn       func(context.Context, Block) (bool, error)
	closeErr            error
	closeCalls          int
	insertedTransaction Transaction
	insertedBlock       Block
	forceBlockCleanup   bool
}

func (w *coverageWriter) EventLogExists(ctx context.Context, hash string, index int) (bool, error) {
	if w.eventExistsFn != nil {
		return w.eventExistsFn(ctx, hash, index)
	}
	return false, nil
}

func (w *coverageWriter) InsertEventLog(ctx context.Context, event EventLog) (bool, error) {
	if w.insertEventFn != nil {
		return w.insertEventFn(ctx, event)
	}
	return true, nil
}

func (w *coverageWriter) TransactionExists(ctx context.Context, hash string) (bool, error) {
	if w.transactionExistsFn != nil {
		return w.transactionExistsFn(ctx, hash)
	}
	return true, nil
}

func (w *coverageWriter) InsertTransaction(ctx context.Context, tx Transaction) (bool, error) {
	w.insertedTransaction = tx
	if w.insertTransactionFn != nil {
		return w.insertTransactionFn(ctx, tx)
	}
	return true, nil
}

func (w *coverageWriter) BlockExists(ctx context.Context, number int64, _ string) (bool, error) {
	if w.blockExistsFn != nil {
		return w.blockExistsFn(ctx, number)
	}
	return true, nil
}

func (w *coverageWriter) InsertBlock(
	ctx context.Context,
	block Block,
	_ []string,
	forceCleanup bool,
) (bool, error) {
	w.insertedBlock = block
	w.forceBlockCleanup = forceCleanup
	if w.insertBlockFn != nil {
		return w.insertBlockFn(ctx, block)
	}
	return true, nil
}

func (w *coverageWriter) Close() error {
	w.closeCalls++
	return w.closeErr
}

type filterResponse struct {
	logs []types.Log
	err  error
}

type coverageClient struct {
	filterResponses []filterResponse
	filterCalls     int
	tx              *types.Transaction
	pending         bool
	txErr           error
	receipt         *types.Receipt
	receiptErr      error
	block           *types.Block
	blockErr        error
	nilBlock        bool
}

func (c *coverageClient) FilterLogs(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
	c.filterCalls++
	if len(c.filterResponses) == 0 {
		return nil, nil
	}
	response := c.filterResponses[0]
	c.filterResponses = c.filterResponses[1:]
	return response.logs, response.err
}

func (c *coverageClient) TransactionByHash(context.Context, common.Hash) (*types.Transaction, bool, error) {
	return c.tx, c.pending, c.txErr
}

func (c *coverageClient) TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error) {
	return c.receipt, c.receiptErr
}

func (c *coverageClient) BlockByNumber(context.Context, *big.Int) (*types.Block, error) {
	if c.block == nil && c.blockErr == nil && !c.nilBlock {
		return types.NewBlockWithHeader(&types.Header{
			Number:     big.NewInt(10),
			Time:       123,
			ParentHash: common.HexToHash("0x01"),
		}), nil
	}
	return c.block, c.blockErr
}

type coverageAddressStore struct {
	calls int
	errAt int
	err   error
}

type coverageCloser struct {
	err error
}

func (c coverageCloser) Close() error { return c.err }

func (s *coverageAddressStore) LookupOrCreateAddress(context.Context, string, int64, int64) (int64, error) {
	s.calls++
	if s.err != nil && s.calls == s.errAt {
		return 0, s.err
	}
	return int64(100 + s.calls), nil
}

func validCoverageFiller(writer NodeFillWriter, client NodeClient) *NodeFiller {
	return &NodeFiller{
		Repository: &coverageRepository{
			contracts: Contracts{Addresses: []string{"0x1000000000000000000000000000000000000001"}},
			start:     10,
			writer:    writer,
		},
		AddressStore: &coverageAddressStore{},
		Client:       client,
	}
}

func coverageLog() types.Log {
	return types.Log{
		Address:     common.HexToAddress("0x1000000000000000000000000000000000000001"),
		Topics:      []common.Hash{common.HexToHash("0x01")},
		BlockNumber: 10,
		TxHash:      common.HexToHash("0x02"),
		Index:       1,
	}
}

func TestNodeFillerDependencyAndSetupFailures(t *testing.T) {
	options := NodeFillOptions{FromBlock: 10, EndBlock: 10, BatchSize: 1}
	validWriter := &coverageWriter{}
	validClient := &coverageClient{}
	tests := []struct {
		name   string
		filler *NodeFiller
		ctx    context.Context
		want   string
	}{
		{name: "repository", filler: &NodeFiller{}, ctx: context.Background(), want: "repository"},
		{
			name:   "client",
			filler: &NodeFiller{Repository: &coverageRepository{}},
			ctx:    context.Background(),
			want:   "client",
		},
		{
			name: "address store",
			filler: &NodeFiller{
				Repository: &coverageRepository{},
				Client:     validClient,
			},
			ctx:  context.Background(),
			want: "address store",
		},
		{
			name: "batch",
			filler: &NodeFiller{
				Repository:   &coverageRepository{},
				Client:       validClient,
				AddressStore: &coverageAddressStore{},
			},
			ctx:  context.Background(),
			want: "batch size",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			runOptions := options
			if test.name == "batch" {
				runOptions.BatchSize = 0
			}
			_, err := test.filler.RunProject(test.ctx, ProjectRandomWalk, runOptions)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("error = %v, want containing %q", err, test.want)
			}
		})
	}

	t.Run("canceled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err := validCoverageFiller(validWriter, validClient).RunProject(ctx, ProjectRandomWalk, options)
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("contracts", func(t *testing.T) {
		sentinel := errors.New("contracts")
		filler := validCoverageFiller(validWriter, validClient)
		filler.Repository = &coverageRepository{contractsErr: sentinel}
		_, err := filler.RunProject(context.Background(), ProjectRandomWalk, options)
		if !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("start", func(t *testing.T) {
		sentinel := errors.New("start")
		filler := validCoverageFiller(validWriter, validClient)
		filler.Repository = &coverageRepository{
			contracts: Contracts{},
			startErr:  sentinel,
		}
		_, err := filler.RunProject(context.Background(), ProjectRandomWalk, options)
		if !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("empty range", func(t *testing.T) {
		logger := &recordingLogger{}
		filler := validCoverageFiller(validWriter, validClient)
		filler.Logger = logger
		filler.Repository.(*coverageRepository).start = 11
		stats, err := filler.RunProject(context.Background(), ProjectRandomWalk, options)
		if err != nil || stats != (FillStats{}) || !logger.contains("nothing to scan") {
			t.Fatalf("stats/error/log = %+v / %v / %v", stats, err, logger.lines)
		}
	})
	t.Run("prepare writer", func(t *testing.T) {
		sentinel := errors.New("prepare")
		filler := validCoverageFiller(validWriter, validClient)
		filler.Repository.(*coverageRepository).prepareErr = sentinel
		_, err := filler.RunProject(context.Background(), ProjectRandomWalk, options)
		if !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("archived block lookup", func(t *testing.T) {
		sentinel := errors.New("archived blocks")
		filler := validCoverageFiller(validWriter, validClient)
		filler.Repository.(*coverageRepository).archivedErr = sentinel
		_, err := filler.RunProject(context.Background(), ProjectRandomWalk, options)
		if !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
}

func TestNodeFillerCleansArchivedBlockWithoutReplacementLogs(t *testing.T) {
	writer := &coverageWriter{
		blockExistsFn: func(context.Context, int64) (bool, error) { return true, nil },
	}
	client := &coverageClient{filterResponses: []filterResponse{{}}}
	filler := validCoverageFiller(writer, client)
	filler.Repository.(*coverageRepository).archived = []int64{10}
	stats, err := filler.RunProject(context.Background(), ProjectRandomWalk, NodeFillOptions{
		FromBlock: 10,
		EndBlock:  10,
		BatchSize: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if stats.LogsFromNode != 0 || stats.BlockSkipped != 1 || !writer.forceBlockCleanup {
		t.Fatalf("stats=%+v forced=%v", stats, writer.forceBlockCleanup)
	}
}

func TestNodeFillerAdaptiveRetryAndLogging(t *testing.T) {
	sentinel := errors.New("filter failed")
	writer := &coverageWriter{}
	client := &coverageClient{filterResponses: []filterResponse{
		{err: sentinel},
		{},
	}}
	logger := &recordingLogger{}
	filler := validCoverageFiller(writer, client)
	filler.Repository.(*coverageRepository).start = 0
	filler.Logger = logger
	stats, err := filler.RunProject(context.Background(), ProjectRandomWalk, NodeFillOptions{
		EndBlock:  10,
		BatchSize: 4_000,
	})
	if err != nil {
		t.Fatalf("RunProject() error = %v", err)
	}
	if stats.FilterRetries != 1 || stats.RPCErrors != 0 ||
		stats.BlocksScanned != 11 || client.filterCalls != 2 {
		t.Fatalf("stats/calls = %+v / %d", stats, client.filterCalls)
	}
	if !logger.contains("FilterLogs error") || !logger.contains("Reducing batch") {
		t.Fatalf("retry logs = %v", logger.lines)
	}
}

func TestNodeFillerPerRowBranches(t *testing.T) {
	sentinel := errors.New("row failed")
	tests := []struct {
		name        string
		writer      *coverageWriter
		client      *coverageClient
		encode      func(*types.Log) ([]byte, error)
		wantDB      int64
		wantRPC     int64
		wantLogIns  int64
		wantLogSkip int64
		wantTxSkip  int64
		wantBlkSkip int64
		wantErr     error
		dryRun      bool
	}{
		{
			name: "exists context error",
			writer: &coverageWriter{eventExistsFn: func(context.Context, string, int) (bool, error) {
				return false, context.Canceled
			}},
			client:  &coverageClient{},
			wantErr: context.Canceled,
		},
		{
			name:        "encoding error",
			writer:      &coverageWriter{},
			client:      &coverageClient{},
			encode:      func(*types.Log) ([]byte, error) { return nil, sentinel },
			wantDB:      1,
			wantBlkSkip: 1,
		},
		{
			name: "insert error",
			writer: &coverageWriter{insertEventFn: func(context.Context, EventLog) (bool, error) {
				return false, sentinel
			}},
			client:      &coverageClient{},
			wantDB:      1,
			wantBlkSkip: 1,
		},
		{
			name: "insert context error",
			writer: &coverageWriter{insertEventFn: func(context.Context, EventLog) (bool, error) {
				return false, context.DeadlineExceeded
			}},
			client:  &coverageClient{},
			wantErr: context.DeadlineExceeded,
		},
		{
			name: "conflict repairs dependencies",
			writer: &coverageWriter{
				insertEventFn: func(context.Context, EventLog) (bool, error) { return false, nil },
			},
			client:      &coverageClient{},
			wantLogSkip: 1,
			wantTxSkip:  1,
			wantBlkSkip: 1,
		},
		{
			name: "existing log dry run does not repair dependencies",
			writer: &coverageWriter{eventExistsFn: func(context.Context, string, int) (bool, error) {
				return true, nil
			}},
			client:      &coverageClient{},
			wantLogSkip: 1,
			dryRun:      true,
		},
		{
			name: "transaction db error",
			writer: &coverageWriter{
				transactionExistsFn: func(context.Context, string) (bool, error) { return false, sentinel },
			},
			client:      &coverageClient{},
			wantLogIns:  1,
			wantDB:      1,
			wantBlkSkip: 1,
		},
		{
			name: "transaction context error",
			writer: &coverageWriter{
				transactionExistsFn: func(context.Context, string) (bool, error) {
					return false, context.Canceled
				},
			},
			client:     &coverageClient{},
			wantLogIns: 1,
			wantErr:    context.Canceled,
		},
		{
			name: "transaction rpc error",
			writer: &coverageWriter{
				transactionExistsFn: func(context.Context, string) (bool, error) { return false, nil },
			},
			client:      &coverageClient{txErr: sentinel},
			wantLogIns:  1,
			wantRPC:     1,
			wantBlkSkip: 1,
		},
		{
			name: "block db error",
			writer: &coverageWriter{
				blockExistsFn: func(context.Context, int64) (bool, error) { return false, sentinel },
			},
			client: &coverageClient{},
			wantDB: 1,
		},
		{
			name: "block context error",
			writer: &coverageWriter{
				blockExistsFn: func(context.Context, int64) (bool, error) {
					return false, context.DeadlineExceeded
				},
			},
			client:     &coverageClient{},
			wantLogIns: 1,
			wantErr:    context.DeadlineExceeded,
		},
		{
			name: "block rpc error",
			writer: &coverageWriter{
				blockExistsFn: func(context.Context, int64) (bool, error) { return false, nil },
			},
			client:  &coverageClient{blockErr: sentinel},
			wantRPC: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.client.filterResponses = []filterResponse{{logs: []types.Log{coverageLog()}}}
			filler := validCoverageFiller(test.writer, test.client)
			filler.EncodeLog = test.encode
			stats, err := filler.RunProject(context.Background(), ProjectRandomWalk, NodeFillOptions{
				FromBlock: 10,
				EndBlock:  10,
				BatchSize: 1,
				DryRun:    test.dryRun,
			})
			if test.wantErr != nil {
				if !errors.Is(err, test.wantErr) {
					t.Fatalf("error = %v, want %v", err, test.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("RunProject() error = %v", err)
			}
			if stats.DBErrors != test.wantDB ||
				stats.RPCErrors != test.wantRPC ||
				stats.LogsInserted != test.wantLogIns ||
				stats.LogsSkipped != test.wantLogSkip ||
				stats.TxSkipped != test.wantTxSkip ||
				stats.BlockSkipped != test.wantBlkSkip {
				t.Fatalf("stats = %+v", stats)
			}
		})
	}
}

func TestNodeFillerJoinsScanAndCloseErrors(t *testing.T) {
	scanErr := errors.New("stop retrying")
	closeErr := errors.New("close writer")
	writer := &coverageWriter{closeErr: closeErr}
	client := &coverageClient{filterResponses: []filterResponse{{err: errors.New("rpc")}}}
	filler := validCoverageFiller(writer, client)
	filler.Sleep = func(context.Context, time.Duration) error { return scanErr }
	_, err := filler.RunProject(context.Background(), ProjectRandomWalk, NodeFillOptions{
		FromBlock: 10,
		EndBlock:  10,
		BatchSize: 1,
	})
	if !errors.Is(err, scanErr) || !errors.Is(err, closeErr) {
		t.Fatalf("joined error = %v", err)
	}
}

func TestNodeFillerPanicStillClosesWriter(t *testing.T) {
	writer := &coverageWriter{eventExistsFn: func(context.Context, string, int) (bool, error) {
		panic("writer panic")
	}}
	client := &coverageClient{filterResponses: []filterResponse{{logs: []types.Log{coverageLog()}}}}
	filler := validCoverageFiller(writer, client)
	func() {
		defer func() {
			if recovered := recover(); recovered == nil {
				t.Fatal("expected panic")
			}
		}()
		_, _ = filler.RunProject(context.Background(), ProjectRandomWalk, NodeFillOptions{
			FromBlock: 10,
			EndBlock:  10,
			BatchSize: 1,
		})
	}()
	if writer.closeCalls != 1 {
		t.Fatalf("writer close calls = %d", writer.closeCalls)
	}
}

func signTransaction(t *testing.T, tx *types.Transaction, signer types.Signer) *types.Transaction {
	t.Helper()
	key, err := crypto.HexToECDSA("2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6")
	if err != nil {
		t.Fatal(err)
	}
	signed, err := types.SignTx(tx, signer, key)
	if err != nil {
		t.Fatalf("signing transaction: %v", err)
	}
	return signed
}

func legacyTransaction(t *testing.T, to *common.Address, data []byte, homestead bool) *types.Transaction {
	t.Helper()
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    1,
		GasPrice: big.NewInt(10),
		Gas:      100_000,
		To:       to,
		Value:    big.NewInt(5),
		Data:     data,
	})
	if homestead {
		return signTransaction(t, tx, types.HomesteadSigner{})
	}
	return signTransaction(t, tx, types.LatestSignerForChainID(big.NewInt(1337)))
}

func dynamicTransaction(t *testing.T, to *common.Address) *types.Transaction {
	t.Helper()
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(1337),
		Nonce:     2,
		GasTipCap: big.NewInt(2),
		GasFeeCap: big.NewInt(20),
		Gas:       100_000,
		To:        to,
		Value:     big.NewInt(6),
		Data:      []byte{1, 2, 3, 4},
	})
	return signTransaction(t, tx, types.LatestSignerForChainID(big.NewInt(1337)))
}

func successfulTxClient(tx *types.Transaction) *coverageClient {
	return &coverageClient{
		tx: tx,
		receipt: &types.Receipt{
			GasUsed:          21_000,
			TransactionIndex: 3,
			Logs:             []*types.Log{{}},
		},
	}
}

func TestEnsureTransactionBranches(t *testing.T) {
	sentinel := errors.New("transaction branch")
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	baseTx := legacyTransaction(t, &to, []byte{1, 2, 3, 4}, false)
	tests := []struct {
		name      string
		writer    *coverageWriter
		client    *coverageClient
		addresses *coverageAddressStore
		wantIns   int64
		wantSkip  int64
		wantKind  errorKind
		wantErr   bool
	}{
		{
			name: "exists query error",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return false, sentinel
			}},
			client:    successfulTxClient(baseTx),
			addresses: &coverageAddressStore{},
			wantKind:  dbError,
			wantErr:   true,
		},
		{
			name: "already exists",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return true, nil
			}},
			client:    successfulTxClient(baseTx),
			addresses: &coverageAddressStore{},
			wantSkip:  1,
			wantKind:  dbError,
		},
		{
			name: "rpc transaction error",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return false, nil
			}},
			client:    &coverageClient{txErr: sentinel},
			addresses: &coverageAddressStore{},
			wantKind:  rpcError,
			wantErr:   true,
		},
		{
			name: "nil transaction",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return false, nil
			}},
			client:    &coverageClient{},
			addresses: &coverageAddressStore{},
			wantKind:  rpcError,
			wantErr:   true,
		},
		{
			name: "pending transaction",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return false, nil
			}},
			client:    &coverageClient{tx: baseTx, pending: true},
			addresses: &coverageAddressStore{},
			wantKind:  rpcError,
			wantErr:   true,
		},
		{
			name: "receipt error",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return false, nil
			}},
			client:    &coverageClient{tx: baseTx, receiptErr: sentinel},
			addresses: &coverageAddressStore{},
			wantKind:  rpcError,
			wantErr:   true,
		},
		{
			name: "nil receipt",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return false, nil
			}},
			client:    &coverageClient{tx: baseTx},
			addresses: &coverageAddressStore{},
			wantKind:  rpcError,
			wantErr:   true,
		},
		{
			name: "sender recovery error",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return false, nil
			}},
			client: successfulTxClient(types.NewTx(&types.LegacyTx{
				GasPrice: big.NewInt(1),
				Gas:      21_000,
				To:       &to,
				Value:    big.NewInt(0),
			})),
			addresses: &coverageAddressStore{},
			wantKind:  rpcError,
			wantErr:   true,
		},
		{
			name: "from address error",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return false, nil
			}},
			client:    successfulTxClient(baseTx),
			addresses: &coverageAddressStore{errAt: 1, err: sentinel},
			wantKind:  dbError,
			wantErr:   true,
		},
		{
			name: "to address error",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return false, nil
			}},
			client:    successfulTxClient(baseTx),
			addresses: &coverageAddressStore{errAt: 2, err: sentinel},
			wantKind:  dbError,
			wantErr:   true,
		},
		{
			name: "insert error",
			writer: &coverageWriter{
				transactionExistsFn: func(context.Context, string) (bool, error) { return false, nil },
				insertTransactionFn: func(context.Context, Transaction) (bool, error) {
					return false, sentinel
				},
			},
			client:    successfulTxClient(baseTx),
			addresses: &coverageAddressStore{},
			wantKind:  dbError,
			wantErr:   true,
		},
		{
			name: "insert conflict",
			writer: &coverageWriter{
				transactionExistsFn: func(context.Context, string) (bool, error) { return false, nil },
				insertTransactionFn: func(context.Context, Transaction) (bool, error) {
					return false, nil
				},
			},
			client:    successfulTxClient(baseTx),
			addresses: &coverageAddressStore{},
			wantSkip:  1,
			wantKind:  dbError,
		},
		{
			name: "inserted",
			writer: &coverageWriter{transactionExistsFn: func(context.Context, string) (bool, error) {
				return false, nil
			}},
			client:    successfulTxClient(baseTx),
			addresses: &coverageAddressStore{},
			wantIns:   1,
			wantKind:  dbError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filler := &NodeFiller{Client: test.client, AddressStore: test.addresses}
			inserted, skipped, kind, err := filler.ensureTransaction(
				context.Background(),
				test.writer,
				baseTx.Hash().Hex(),
				10,
			)
			if (err != nil) != test.wantErr ||
				inserted != test.wantIns ||
				skipped != test.wantSkip ||
				kind != test.wantKind {
				t.Fatalf("inserted/skipped/kind/error = %d/%d/%d/%v", inserted, skipped, kind, err)
			}
		})
	}
}

func TestEnsureTransactionSpecialTypes(t *testing.T) {
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	tests := []struct {
		name string
		tx   *types.Transaction
	}{
		{name: "homestead signer fallback", tx: legacyTransaction(t, &to, nil, true)},
		{name: "dynamic fee", tx: dynamicTransaction(t, &to)},
		{name: "contract creation", tx: legacyTransaction(t, nil, nil, false)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := &coverageWriter{
				transactionExistsFn: func(context.Context, string) (bool, error) { return false, nil },
			}
			filler := &NodeFiller{
				Client:       successfulTxClient(test.tx),
				AddressStore: &coverageAddressStore{},
			}
			inserted, _, _, err := filler.ensureTransaction(
				context.Background(),
				writer,
				test.tx.Hash().Hex(),
				10,
			)
			if err != nil || inserted != 1 {
				t.Fatalf("insert/error = %d/%v", inserted, err)
			}
			if test.name == "dynamic fee" && writer.insertedTransaction.GasPrice != "20" {
				t.Errorf("dynamic gas price = %s", writer.insertedTransaction.GasPrice)
			}
			if test.name == "contract creation" && (!writer.insertedTransaction.ContractCreate || writer.insertedTransaction.ToAddressID != 0) {
				t.Errorf("contract-create row = %+v", writer.insertedTransaction)
			}
		})
	}
}

func TestTransactionSenderProtectedFallbackFailure(t *testing.T) {
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	malformed := types.NewTx(&types.LegacyTx{
		GasPrice: big.NewInt(1),
		Gas:      21_000,
		To:       &to,
		Value:    big.NewInt(0),
		V:        big.NewInt(37),
		R:        big.NewInt(0),
		S:        big.NewInt(0),
	})
	if _, err := transactionSender(malformed); err == nil {
		t.Fatal("malformed protected signature unexpectedly recovered a sender")
	}
}

func TestEnsureBlockBranches(t *testing.T) {
	sentinel := errors.New("block branch")
	block := types.NewBlockWithHeader(&types.Header{
		Number:     big.NewInt(10),
		Time:       123,
		ParentHash: common.HexToHash("0x01"),
	})
	tests := []struct {
		name     string
		writer   *coverageWriter
		client   *coverageClient
		wantIns  int64
		wantSkip int64
		wantKind errorKind
		wantErr  bool
	}{
		{
			name: "exists query error",
			writer: &coverageWriter{blockExistsFn: func(context.Context, int64) (bool, error) {
				return false, sentinel
			}},
			client:   &coverageClient{block: block},
			wantKind: dbError,
			wantErr:  true,
		},
		{
			name: "already exists",
			writer: &coverageWriter{blockExistsFn: func(context.Context, int64) (bool, error) {
				return true, nil
			}},
			client:   &coverageClient{block: block},
			wantSkip: 1,
			wantKind: dbError,
		},
		{
			name: "rpc error",
			writer: &coverageWriter{blockExistsFn: func(context.Context, int64) (bool, error) {
				return false, nil
			}},
			client:   &coverageClient{blockErr: sentinel},
			wantKind: rpcError,
			wantErr:  true,
		},
		{
			name: "nil block",
			writer: &coverageWriter{blockExistsFn: func(context.Context, int64) (bool, error) {
				return false, nil
			}},
			client:   &coverageClient{nilBlock: true},
			wantKind: rpcError,
			wantErr:  true,
		},
		{
			name: "insert error",
			writer: &coverageWriter{
				blockExistsFn: func(context.Context, int64) (bool, error) { return false, nil },
				insertBlockFn: func(context.Context, Block) (bool, error) { return false, sentinel },
			},
			client:   &coverageClient{block: block},
			wantKind: dbError,
			wantErr:  true,
		},
		{
			name: "insert conflict",
			writer: &coverageWriter{
				blockExistsFn: func(context.Context, int64) (bool, error) { return false, nil },
				insertBlockFn: func(context.Context, Block) (bool, error) { return false, nil },
			},
			client:   &coverageClient{block: block},
			wantSkip: 1,
			wantKind: dbError,
		},
		{
			name: "inserted",
			writer: &coverageWriter{blockExistsFn: func(context.Context, int64) (bool, error) {
				return false, nil
			}},
			client:   &coverageClient{block: block},
			wantIns:  1,
			wantKind: dbError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filler := &NodeFiller{Client: test.client}
			inserted, skipped, kind, err := filler.ensureBlock(
				context.Background(),
				test.writer,
				10,
				[]string{"0x1000000000000000000000000000000000000001"},
			)
			if (err != nil) != test.wantErr ||
				inserted != test.wantIns ||
				skipped != test.wantSkip ||
				kind != test.wantKind {
				t.Fatalf("inserted/skipped/kind/error = %d/%d/%d/%v", inserted, skipped, kind, err)
			}
		})
	}
}

func TestEnsureBlockRefreshesProjectRowsForExistingCanonicalBlock(t *testing.T) {
	block := types.NewBlockWithHeader(&types.Header{
		Number:     big.NewInt(10),
		Time:       123,
		ParentHash: common.HexToHash("0x01"),
	})
	writer := &coverageWriter{
		blockExistsFn: func(context.Context, int64) (bool, error) { return true, nil },
	}
	filler := &NodeFiller{
		Client: &coverageClient{block: block},
	}
	inserted, skipped, _, err := filler.ensureBlock(
		context.Background(),
		writer,
		10,
		[]string{"0x1000000000000000000000000000000000000001"},
	)
	if err != nil || inserted != 0 || skipped != 1 || !writer.forceBlockCleanup {
		t.Fatalf(
			"inserted=%d skipped=%d forced=%v error=%v",
			inserted,
			skipped,
			writer.forceBlockCleanup,
			err,
		)
	}
}

func TestNodeFillSQLRepositoryAndSchemaBranches(t *testing.T) {
	sentinel := errors.New("sql repository")
	contracts := Contracts{AddressIDs: []int64{8}, Addresses: []string{"0x08"}}
	t.Run("auto start", func(t *testing.T) {
		repository := &SQLNodeFillRepository{DB: openScriptDB(t,
			queryOp("MIN(block_num) FROM address", []string{"min"}, []driver.Value{int64(20)}),
			queryOp("MIN(block_num) FROM evt_log", []string{"min"}, []driver.Value{int64(10)}),
		)}
		start, err := repository.ResolveStartBlock(context.Background(), contracts, 0)
		if err != nil || start != 10 {
			t.Fatalf("start/error = %d/%v", start, err)
		}
	})
	t.Run("address start query", func(t *testing.T) {
		repository := &SQLNodeFillRepository{DB: openScriptDB(t,
			queryErrorOp("MIN(block_num) FROM address", sentinel),
		)}
		if _, err := repository.ResolveStartBlock(context.Background(), contracts, 0); !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("event start query", func(t *testing.T) {
		repository := &SQLNodeFillRepository{DB: openScriptDB(t,
			queryOp("MIN(block_num) FROM address", []string{"min"}, []driver.Value{int64(20)}),
			queryErrorOp("MIN(block_num) FROM evt_log", sentinel),
		)}
		if _, err := repository.ResolveStartBlock(context.Background(), contracts, 0); !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})

	t.Run("schema query", func(t *testing.T) {
		if err := RequireArchLogIndex(context.Background(), openScriptDB(t,
			queryErrorOp("information_schema.columns", sentinel),
		)); !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("schema missing", func(t *testing.T) {
		err := RequireArchLogIndex(context.Background(), openScriptDB(t,
			queryOp("information_schema.columns", []string{"count"}, []driver.Value{int64(0)}),
		))
		if err == nil || !strings.Contains(err.Error(), "log_index missing") {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("schema present", func(t *testing.T) {
		err := RequireArchLogIndex(context.Background(), openScriptDB(t,
			queryOp("information_schema.columns", []string{"count"}, []driver.Value{int64(1)}),
		))
		if err != nil {
			t.Fatalf("error = %v", err)
		}
	})
}

func TestPrepareWriterEveryFailure(t *testing.T) {
	sentinel := errors.New("prepare failed")
	matches := []string{
		"INSERT INTO arch_evtlog",
		"SELECT 1 FROM arch_evtlog",
		"INSERT INTO arch_tx",
		"SELECT 1 FROM arch_tx",
		"DELETE FROM arch_tx at",
		"DELETE FROM arch_evtlog",
		"DELETE FROM arch_block",
		"INSERT INTO arch_block",
		"SELECT 1 FROM arch_block",
	}
	for failAt := range matches {
		t.Run(matches[failAt], func(t *testing.T) {
			ops := make([]scriptOp, 0, failAt+1)
			for index := 0; index <= failAt; index++ {
				if index == failAt {
					ops = append(ops, prepareErrorOp(matches[index], sentinel))
				} else {
					ops = append(ops, prepareOp(matches[index]))
				}
			}
			_, err := (&SQLNodeFillRepository{DB: openScriptDB(t, ops...)}).PrepareWriter(context.Background())
			if !errors.Is(err, sentinel) {
				t.Fatalf("error = %v", err)
			}
		})
	}
}

func TestPreparedHelpersAndCloseBranches(t *testing.T) {
	sentinel := errors.New("statement failed")
	t.Run("statement query error", func(t *testing.T) {
		statement := prepareStatement(t,
			prepareOp("INSERT test statement"),
			queryErrorOp("INSERT test statement", sentinel),
		)
		if _, err := statementExists(context.Background(), statement); !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("row result error", func(t *testing.T) {
		if _, err := rowInserted(nil, sentinel); !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("rows affected error", func(t *testing.T) {
		if _, err := rowInserted(scriptResult{affectedErr: sentinel}, nil); !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("zero rows", func(t *testing.T) {
		inserted, err := rowInserted(scriptResult{}, nil)
		if err != nil || inserted {
			t.Fatalf("inserted/error = %v/%v", inserted, err)
		}
	})
	t.Run("nil statements", func(t *testing.T) {
		if err := (&sqlNodeFillWriter{}).Close(); err != nil {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("first close error", func(t *testing.T) {
		err := closeStatements([]statementCloser{
			coverageCloser{err: sentinel},
			coverageCloser{err: errors.New("second")},
		})
		if !errors.Is(err, sentinel) {
			t.Fatalf("close error = %v", err)
		}
	})
}

func TestContextErrorAndStatsLogging(t *testing.T) {
	sentinel := errors.New("ordinary")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := contextError(ctx, sentinel); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled context error = %v", err)
	}
	if err := contextError(context.Background(), context.Canceled); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled argument error = %v", err)
	}
	if err := contextError(context.Background(), context.DeadlineExceeded); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("deadline argument error = %v", err)
	}
	if err := contextError(context.Background(), sentinel); err != nil {
		t.Fatalf("ordinary error classified as context error: %v", err)
	}

	var stats FillStats
	stats.countError(rpcError)
	stats.countError(dbError)
	if stats.RPCErrors != 1 || stats.DBErrors != 1 {
		t.Fatalf("classified stats = %+v", stats)
	}

	LogFillStats(nil, "nil", stats)
	logger := &recordingLogger{}
	LogFillStats(logger, "clean", FillStats{BlocksScanned: 1})
	if logger.contains("errors —") {
		t.Fatalf("clean stats logged errors: %v", logger.lines)
	}
	stats.FilterRetries = 2
	LogFillStats(logger, "errors", stats)
	if !logger.contains("recovered FilterLogs retries: 2") ||
		!logger.contains("errors — rpc: 1, db: 1") {
		t.Fatalf("stats logs = %v", logger.lines)
	}
}

func TestNodeFillFiltererWithoutLogger(t *testing.T) {
	sentinel := errors.New("filter")
	filterer := nodeFillFilterer{client: &coverageClient{
		filterResponses: []filterResponse{{err: sentinel}},
	}}
	_, err := filterer.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(1),
		ToBlock:   big.NewInt(2),
	})
	if !errors.Is(err, sentinel) {
		t.Fatalf("error = %v", err)
	}
}
