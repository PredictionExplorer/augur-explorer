//go:build integration

// Benchmark for the per-block ingestion transaction (ADR-0010), run against
// a real PostgreSQL container and the fake chain. Baselines live in
// docs/benchmarks.md; re-run with:
//
//	go test -tags=integration ./internal/indexer/ -bench BenchmarkIngestBlock -benchmem -count=6 -run '^$' -timeout 15m
//
// The transactional sub-benchmark measures processBatch end to end — BEGIN,
// the steady-state pipeline (block/transaction reads hit their SELECTs, the
// event log replays via delete-then-insert), the watermark write and COMMIT.
// The autocommit sub-benchmark runs the identical pipeline without the
// transaction wrapper, isolating the cost the atomicity guarantee adds.
package indexer

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

func BenchmarkIngestBlock(b *testing.B) {
	ctx := context.Background()
	db, stopDB, err := testdb.Start(ctx)
	if err != nil {
		if errors.Is(err, testdb.ErrContainerUnavailable) {
			b.Skipf("skipping: %v", err)
		}
		b.Fatalf("starting test database: %v", err)
	}
	b.Cleanup(stopDB)

	chain, stopChain := testchain.Start()
	b.Cleanup(stopChain)

	rpcClient, err := rpc.DialContext(ctx, chain.URL())
	if err != nil {
		b.Fatalf("dialing fake chain: %v", err)
	}
	ethClient := ethclient.NewClient(rpcClient)
	b.Cleanup(ethClient.Close)

	// One block, three logs across two transactions — the shape of a busy
	// CosmicGame block. The processor is a no-op: the benchmark isolates
	// the storage pipeline, not the handlers.
	const blockNum = int64(100)
	contract := ethcommon.HexToAddress("0x2000000000000000000000000000000000000002")
	topic := ethcommon.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	var logs []types.Log
	for txIdx, logCount := range []int{2, 1} {
		tx := chain.AddTx(blockNum, contract, nil)
		attached := make([]*types.Log, 0, logCount)
		for i := range logCount {
			attached = append(attached, &types.Log{
				Address:     contract,
				Topics:      []ethcommon.Hash{topic},
				BlockNumber: uint64(blockNum),
				BlockHash:   chain.BlockHash(blockNum),
				TxHash:      tx.Hash(),
				Index:       uint(2*txIdx + i), // #nosec G115 -- small positive constants
			})
		}
		chain.AttachLogs(tx.Hash(), attached)
		for _, l := range attached {
			logs = append(logs, *l)
		}
	}

	st := store.NewFromPool(db.Pool)
	progress := &fakeProgress{last: blockNum - 1}
	engine, err := New(Config{
		Store:    st,
		Client:   ethClient,
		Progress: progress,
		Process:  func(context.Context, int64) error { return nil },
		Logger:   slog.New(slog.DiscardHandler),
	})
	if err != nil {
		b.Fatalf("New: %v", err)
	}

	// Warm the pipeline once: the transactions are fetched from the fake
	// chain and inserted; every timed iteration then replays the block in
	// steady state (SELECT hits plus the evt_log delete-then-insert).
	if _, stage, err := engine.processBatch(ctx, logs); err != nil {
		b.Fatalf("warmup processBatch (stage %s): %v", stage, err)
	}

	b.Run("tx_block_3_logs", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			if _, stage, err := engine.processBatch(ctx, logs); err != nil {
				b.Fatalf("processBatch (stage %s): %v", stage, err)
			}
		}
	})

	b.Run("autocommit_3_logs", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			for i := range logs {
				log := logs[i]
				if _, err := engine.EnsureBlockExists(ctx, blockNum, log.BlockHash.Hex()); err != nil {
					b.Fatalf("EnsureBlockExists: %v", err)
				}
				txID, _, err := engine.EnsureTransactionExists(ctx, log.TxHash, blockNum)
				if err != nil {
					b.Fatalf("EnsureTransactionExists: %v", err)
				}
				if _, err := engine.InsertEventLog(ctx, log, txID); err != nil {
					b.Fatalf("InsertEventLog: %v", err)
				}
			}
			if err := progress.SetLastBlock(ctx, blockNum); err != nil {
				b.Fatalf("SetLastBlock: %v", err)
			}
		}
	})
}
