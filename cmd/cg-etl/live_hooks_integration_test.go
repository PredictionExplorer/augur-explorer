//go:build integration

package main

import (
	"context"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/prometheus/client_golang/prometheus"

	cgc "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

// fixture addresses from the seeded dataset plus two fresh contracts the
// live events introduce.
const (
	liveGameAddr      = "0x2000000000000000000000000000000000000002"
	liveNewImplAddr   = "0x7700000000000000000000000000000000000077"
	liveNewWalletAddr = "0x8800000000000000000000000000000000000088"
)

// TestRunLiveUpgradeAndWalletSwitchBackfills drives the two live hooks
// through the running production wiring: an Upgraded event and a
// PrizesWalletAddressChanged event land in a new block; the hooks must add
// the new contracts to the filter set and complete the after-commit
// backfills (pinned by their completion records on the log stream).
func TestRunLiveUpgradeAndWalletSwitchBackfills(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}

	chain := testchain.New(t)
	chain.EnsureBlock(150)
	setEnv(t, db.ConnString, chain.URL())

	var logBuf syncBuffer
	runCtx, cancel := context.WithCancel(ctx)
	done := make(chan error, 1)
	reg := prometheus.NewRegistry()
	go func() { done <- run(runCtx, os.Getenv, &logBuf, reg, reg) }()

	waitForCondition(t, done, cancel, &logBuf, "watermark at chain tip", func() bool {
		var lastBlock int64
		_ = db.Pool.QueryRow(ctx, "SELECT last_block_num FROM cg_proc_status").Scan(&lastBlock)
		return lastBlock == 150
	})

	// One transaction at block 151 carries both announcements, exactly the
	// shape of a production upgrade transaction.
	gameABI, err := abi.JSON(strings.NewReader(cgc.CosmicSignatureGameABI))
	if err != nil {
		t.Fatal(err)
	}
	game := ethcommon.HexToAddress(liveGameAddr)
	tx := chain.AddTx(151, game, nil)
	chain.AttachLogs(tx.Hash(), []*types.Log{
		{
			Address: game,
			Topics: []ethcommon.Hash{
				gameABI.Events["Upgraded"].ID,
				ethcommon.HexToHash(liveNewImplAddr),
			},
			BlockNumber: 151,
			BlockHash:   chain.BlockHash(151),
			TxHash:      tx.Hash(),
			Index:       0,
		},
		{
			Address: game,
			Topics: []ethcommon.Hash{
				gameABI.Events["PrizesWalletAddressChanged"].ID,
				ethcommon.HexToHash(liveNewWalletAddr),
			},
			BlockNumber: 151,
			BlockHash:   chain.BlockHash(151),
			TxHash:      tx.Hash(),
			Index:       1,
		},
	})

	// The hooks fire inside block 151's transaction; the after-commit
	// backfills wait for the watermark and log one completion record each.
	waitForCondition(t, done, cancel, &logBuf, "implementation backfill completion", func() bool {
		return strings.Contains(logBuf.String(), "implementation events backfilled")
	})
	waitForCondition(t, done, cancel, &logBuf, "prizes wallet backfill completion", func() bool {
		return strings.Contains(logBuf.String(), "prizes wallet events backfilled")
	})

	cancel()
	if err := awaitRun(done); err != nil {
		t.Fatalf("run returned %v, want nil on graceful shutdown", err)
	}
	logs := logBuf.String()
	for _, banned := range []string{"backfill failed", "backfill skipped"} {
		if strings.Contains(logs, banned) {
			t.Errorf("log stream reports %q:\n%s", banned, logs)
		}
	}
}

// TestRunRecoveryPassFailures seeds one unfinished announcement, fails the
// startup recovery pass's log fetch, and asserts run reports the failed
// stage instead of indexing with silent gaps. Each subtest gets its own
// database: a seeded Upgraded row would otherwise leak its recovery attempt
// into the other subtest's RPC failure budget.
func TestRunRecoveryPassFailures(t *testing.T) {
	ctx := context.Background()
	seededDB := func(t *testing.T, seed string) *testdb.DB {
		t.Helper()
		db := testdb.New(t)
		if err := testfixtures.Apply(ctx, db.SQL); err != nil {
			t.Fatalf("applying fixtures: %v", err)
		}
		if _, err := db.Pool.Exec(ctx, seed); err != nil {
			t.Fatalf("seeding announcement: %v", err)
		}
		return db
	}
	expectRunError := func(t *testing.T, db *testdb.DB, want string) {
		t.Helper()
		chain := testchain.New(t)
		chain.EnsureBlock(150)
		setEnv(t, db.ConnString, chain.URL())
		chain.FailNextRPC("eth_getLogs", "log fetch refused")
		err := run(ctx, os.Getenv, io.Discard, prometheus.NewRegistry(), prometheus.NewRegistry())
		if err == nil || !strings.Contains(err.Error(), want) {
			t.Fatalf("err = %v, want %q", err, want)
		}
	}

	t.Run("implementation recovery", func(t *testing.T) {
		// An Upgraded row whose implementation has no Initialized row: the
		// recovery pass must backfill it and surface the fetch failure.
		db := seededDB(t, `
			WITH impl AS (
				INSERT INTO address (block_num, tx_id, addr)
				VALUES (140, 0, '0x9900000000000000000000000000000000000099')
				RETURNING address_id
			)
			INSERT INTO cg_adm_upgraded (block_num, tx_id, time_stamp, contract_aid, implementation_aid)
			SELECT 140, 0, NOW(), address_id, address_id FROM impl`)
		expectRunError(t, db, "implementation event recovery failed")
	})

	t.Run("prizes wallet watermark read", func(t *testing.T) {
		// The recovery pass starts by reading the committed watermark; with
		// the status table unavailable the read's failure must abort the
		// boot.
		db := seededDB(t, `ALTER TABLE cg_proc_status RENAME TO cg_proc_status_hidden`)
		expectRunError(t, db, "prizes wallet event recovery: read watermark")
	})

	t.Run("prizes wallet recovery", func(t *testing.T) {
		// Announcement at block 0 so it sits at or below the boot watermark
		// (announcements past the watermark are deferred to the live hook).
		db := seededDB(t, `
			WITH wallet AS (
				INSERT INTO address (block_num, tx_id, addr)
				VALUES (0, 0, '0xaa000000000000000000000000000000000000aa')
				RETURNING address_id
			)
			INSERT INTO cg_adm_prizes_wallet_addr (block_num, tx_id, time_stamp, contract_aid, new_wallet_aid)
			SELECT 0, 0, NOW(), address_id, address_id FROM wallet`)
		expectRunError(t, db, "prizes wallet event recovery failed")
	})
}
