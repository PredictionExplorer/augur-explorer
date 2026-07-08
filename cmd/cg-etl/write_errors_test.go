//go:build integration

// Write-error propagation suite: every dispatched event handler must return
// DB write failures to the polling loop instead of terminating the process
// (the legacy layer called os.Exit(1) inside the store on these paths).
//
// Mechanism: each fixture is ingested normally, then re-processed with the
// package Repo swapped onto a second pool whose connections run with
// default_transaction_read_only=on. Reads succeed on that pool, the
// handler's first write (the delete-before-insert) fails with SQLSTATE
// 25006, and the error must surface from process_single_event. Fixtures
// whose events store nothing (wrong address, unknown topic, no topics) must
// keep replaying cleanly — proving their handlers really perform no writes.
package main

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// fixturesWithoutWrites lists the fixtures whose events never reach a store
// write (dispatch/address-guard negative cases).
var fixturesWithoutWrites = map[string]bool{
	"bid_wrong_address_skipped": true,
	"unknown_topic_noop":        true,
	"no_topics_noop":            true,
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
	roRepo := cgstore.NewRepo(roStore)

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

			// Swap the package Repo onto the read-only pool for the replay;
			// restore before the next subtest regardless of outcome.
			rwRepo := cgRepo
			cgRepo = roRepo
			defer func() { cgRepo = rwRepo }()

			for _, id := range evtIDs {
				err := process_single_event(ctx, id)
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
