//go:build integration

// Write-error propagation suite: every dispatched event handler must return
// DB write failures to the polling loop instead of terminating the process
// (the legacy layer called os.Exit(1) inside the store on these paths).
//
// Mechanism: each fixture is ingested normally, then re-processed through a
// second Handlers set built over a pool whose connections run with
// default_transaction_read_only=on. Reads succeed on that pool (the
// existence guards and address lookups), the handler's INSERT fails with
// SQLSTATE 25006, and the error must surface from the processor. Fixtures
// whose events store nothing (wrong address, unknown token/offer) must keep
// replaying cleanly — proving their handlers really perform no writes.
package randomwalk

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwdb "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// fixturesWithoutWrites lists the fixtures whose only event never reaches a
// store write (address-guard and existence-guard negative cases).
var fixturesWithoutWrites = map[string]bool{
	"token_name_unknown_token_skipped":     true,
	"transfer_wrong_address_skipped":       true,
	"offer_canceled_unknown_offer_skipped": true,
}

func TestWriteErrorPropagation(t *testing.T) {
	requireHarness(t)
	ctx := context.Background()

	roCfg, err := pgxpool.ParseConfig(testDB.ConnString)
	if err != nil {
		t.Fatalf("parsing test connection string: %v", err)
	}
	roCfg.ConnConfig.RuntimeParams["default_transaction_read_only"] = "on"
	roCfg.ConnConfig.RuntimeParams["timezone"] = "UTC"
	roCfg.ConnConfig.RuntimeParams["search_path"] = "public"
	roPool, err := pgxpool.NewWithConfig(ctx, roCfg)
	if err != nil {
		t.Fatalf("creating read-only pool: %v", err)
	}
	roStore := store.NewFromPool(roPool)
	defer roStore.Close()
	roRepo := rwdb.NewRepo(roStore)

	for _, fx := range eventFixtures() {
		t.Run(fx.name, func(t *testing.T) {
			resetDB(t)

			var evtIDs []int64
			logIndexByBlock := make(map[int64]uint)
			for _, ftx := range fx.txs {
				block := fx.block + ftx.blockOffset
				logs := make([]*types.Log, 0, len(ftx.logs))
				for _, fl := range ftx.logs {
					logs = append(logs, fl.build(t))
				}
				start := logIndexByBlock[block]
				ids := ingestTx(t, block, addr(ftx.to), start, logs)
				logIndexByBlock[block] = start + uint(len(logs))
				evtIDs = append(evtIDs, ids...)
			}

			// Replay through a handler set whose writes go to the read-only
			// pool (resetDB rebuilt testContracts, so ids are current).
			roHandlers := newTestHandlers(t, roRepo, testContracts)
			roProcess := indexer.LogProcessor(roStore, roHandlers.Registry())

			for _, id := range evtIDs {
				err := roProcess(ctx, id)
				if fixturesWithoutWrites[fx.name] {
					if err != nil {
						t.Errorf("event %d of no-write fixture errored on a read-only pool: %v", id, err)
					}
					continue
				}
				if err == nil {
					t.Errorf("event %d replayed on a read-only pool without error: its handler swallows write failures", id)
				}
			}
		})
	}
}
