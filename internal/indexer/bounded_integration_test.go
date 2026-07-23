//go:build integration

// Bounded-time integration proof (D22 + ADR-0010): a chain RPC call that
// black-holes in the middle of a block's ingestion transaction fails at the
// engine's per-call deadline, the whole block rolls back with zero residue,
// and the retry converges once the endpoint recovers — against a real
// migrated PostgreSQL.
package indexer

import (
	"context"
	"errors"
	"log/slog"
	"sync/atomic"
	"testing"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

// blackholeTxClient serves everything from the fake chain except
// TransactionByHash, which hangs until healed.
type blackholeTxClient struct {
	Client

	healed atomic.Bool
}

func (c *blackholeTxClient) TransactionByHash(ctx context.Context, hash ethcommon.Hash) (*types.Transaction, bool, error) {
	if c.healed.Load() {
		return c.Client.TransactionByHash(ctx, hash)
	}
	<-ctx.Done()
	return nil, false, ctx.Err()
}

func TestProcessBlockHungRPCRollsBackAndRetryConverges(t *testing.T) {
	db := testdb.New(t)
	chain := testchain.New(t)
	client := &blackholeTxClient{Client: testchainClient(t, chain)}
	st := store.NewFromPool(db.Pool)
	progress := &fakeProgress{}

	engine, err := New(Config{
		Store:          st,
		Client:         client,
		Progress:       progress,
		Process:        func(context.Context, int64) error { return nil },
		Contracts:      []ethcommon.Address{{0x01}},
		Logger:         slog.New(slog.DiscardHandler),
		RPCCallTimeout: 300 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	const blockNum = int64(120)
	chain.EnsureBlock(blockNum)
	lg := types.Log{
		Address:     ethcommon.Address{0xAA},
		Topics:      []ethcommon.Hash{{0x01}},
		Data:        []byte{0x02},
		BlockNumber: uint64(blockNum),
		BlockHash:   chain.BlockHash(blockNum),
		TxHash:      ethcommon.Hash{0xBB}, // unknown to the DB: forces TransactionByHash
		Index:       0,
	}

	// The ingestion path runs on a shutdown-immune context (run.go's
	// WithoutCancel); the per-call bound must still cut the hang short.
	dbCtx := context.WithoutCancel(context.Background())
	start := time.Now()
	stage, err := engine.processBlock(dbCtx, blockNum, []types.Log{lg})
	elapsed := time.Since(start)
	if err == nil {
		t.Fatal("processBlock with a black-holed TransactionByHash must fail")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v, want the per-call DeadlineExceeded cause", err)
	}
	if stage != "transaction" {
		t.Fatalf("failed stage = %q, want transaction", stage)
	}
	if elapsed > 5*time.Second {
		t.Fatalf("processBlock took %v; the 300ms per-call bound must cut the hang short", elapsed)
	}

	// Zero residue: the whole block transaction rolled back.
	if _, err := st.BlockHash(context.Background(), blockNum); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("block row after rollback: err = %v, want ErrNotFound", err)
	}
	if _, err := st.TransactionIDByHash(context.Background(), lg.TxHash.Hex()); !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("transaction row after rollback: err = %v, want ErrNotFound", err)
	}
	if progress.last != 0 {
		t.Fatalf("watermark after rollback = %d, want untouched 0", progress.last)
	}

	// Endpoint recovers: the same block replays to a committed state (the
	// unknown transaction degrades to the documented minimal record).
	client.healed.Store(true)
	if stage, err := engine.processBlock(dbCtx, blockNum, []types.Log{lg}); err != nil {
		t.Fatalf("retry after heal failed at stage %q: %v", stage, err)
	}
	if got := blockHash(t, st, blockNum); got != lg.BlockHash.Hex() {
		t.Fatalf("committed block hash = %s, want %s", got, lg.BlockHash.Hex())
	}
	if _, err := st.TransactionIDByHash(context.Background(), lg.TxHash.Hex()); err != nil {
		t.Fatalf("committed transaction lookup: %v", err)
	}
	if progress.last != blockNum {
		t.Fatalf("watermark after retry = %d, want %d", progress.last, blockNum)
	}
}

// testchainClient dials the fake chain and returns the ethclient surface.
func testchainClient(t *testing.T, chain *testchain.Chain) Client {
	t.Helper()
	rpcClient, err := rpc.DialContext(context.Background(), chain.URL())
	if err != nil {
		t.Fatalf("dialing fake chain: %v", err)
	}
	ethClient := ethclient.NewClient(rpcClient)
	t.Cleanup(ethClient.Close)
	return ethClient
}

func blockHash(t *testing.T, st *store.Store, blockNum int64) string {
	t.Helper()
	hash, err := st.BlockHash(context.Background(), blockNum)
	if err != nil {
		t.Fatalf("BlockHash(%d): %v", blockNum, err)
	}
	return hash
}
