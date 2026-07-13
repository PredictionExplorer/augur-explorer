//go:build integration

package archive_test

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/PredictionExplorer/augur-explorer/internal/ops/archive"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testchain"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
	"github.com/PredictionExplorer/augur-explorer/internal/testfixtures"
)

func TestArchiveExportVerifyAndIdempotentResume(t *testing.T) {
	db := testdb.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}

	exporter := &archive.SQLExporter{
		Source:      db.SQL,
		Destination: db.SQL,
		BatchSize:   7,
	}
	if _, err := exporter.ExportProject(ctx, archive.ProjectRandomWalk); err != nil {
		t.Fatalf("first export: %v", err)
	}
	before := archiveCounts(t, ctx, db)

	report, err := (&archive.SQLVerifier{DB: db.SQL}).VerifyProject(
		ctx,
		archive.ProjectRandomWalk,
		archive.VerifyOptions{},
	)
	if err != nil {
		t.Fatalf("verify exported archive: %v", err)
	}
	if !report.Passed {
		t.Fatalf("exported archive did not verify: %+v", report.Stats)
	}

	if _, err := exporter.ExportProject(ctx, archive.ProjectRandomWalk); err != nil {
		t.Fatalf("resume export: %v", err)
	}
	after := archiveCounts(t, ctx, db)
	if before != after {
		t.Fatalf("idempotent resume changed archive counts: before=%+v after=%+v", before, after)
	}
}

type counts struct {
	events int64
	txs    int64
	blocks int64
}

func archiveCounts(t *testing.T, ctx context.Context, db *testdb.DB) counts {
	t.Helper()
	var result counts
	for query, destination := range map[string]*int64{
		"SELECT COUNT(*) FROM arch_evtlog": &result.events,
		"SELECT COUNT(*) FROM arch_tx":     &result.txs,
		"SELECT COUNT(*) FROM arch_block":  &result.blocks,
	} {
		if err := db.SQL.QueryRowContext(ctx, query).Scan(destination); err != nil {
			t.Fatalf("counting archive rows: %v", err)
		}
	}
	return result
}

func TestNodeFillDryRunInsertIdempotencyAndRPCError(t *testing.T) {
	db := testdb.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := testfixtures.Apply(ctx, db.SQL); err != nil {
		t.Fatalf("applying fixtures: %v", err)
	}

	chain := testchain.New(t)
	contract := common.HexToAddress("0x8000000000000000000000000000000000000008")
	topic := common.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	addChainLog := func(block int64) common.Hash {
		tx := chain.AddTx(block, contract, []byte{0x01, 0x02, 0x03, 0x04})
		chain.AttachLogs(tx.Hash(), []*types.Log{{
			Address:     contract,
			Topics:      []common.Hash{topic},
			Data:        []byte{byte(block)},
			BlockNumber: uint64(block),
			BlockHash:   chain.BlockHash(block),
			TxHash:      tx.Hash(),
			Index:       0,
		}})
		return tx.Hash()
	}
	dryHash := addChainLog(900)
	insertHash := addChainLog(901)

	client, err := ethclient.DialContext(ctx, chain.URL())
	if err != nil {
		t.Fatalf("dialing test chain: %v", err)
	}
	defer client.Close()
	filler := &archive.NodeFiller{
		Repository:   &archive.SQLNodeFillRepository{DB: db.SQL},
		AddressStore: store.NewFromPool(db.Pool),
		Client:       client,
		Sleep: func(context.Context, time.Duration) error {
			return nil
		},
	}

	dryStats, err := filler.RunProject(ctx, archive.ProjectRandomWalk, archive.NodeFillOptions{
		FromBlock: 900,
		EndBlock:  900,
		BatchSize: 1,
		DryRun:    true,
	})
	if err != nil {
		t.Fatalf("dry run: %v", err)
	}
	if dryStats.LogsInserted != 1 || archiveEventExists(t, ctx, db, dryHash.Hex()) {
		t.Fatalf("dry-run stats/row = %+v / exists=%v", dryStats, archiveEventExists(t, ctx, db, dryHash.Hex()))
	}

	chain.FailNextRPC("eth_getLogs", "transient filter failure")
	insertStats, err := filler.RunProject(ctx, archive.ProjectRandomWalk, archive.NodeFillOptions{
		FromBlock: 901,
		EndBlock:  901,
		BatchSize: 1,
	})
	if err != nil {
		t.Fatalf("insert run: %v", err)
	}
	if insertStats.FilterRetries != 1 ||
		insertStats.RPCErrors != 0 ||
		insertStats.LogsInserted != 1 ||
		insertStats.TxInserted != 1 ||
		insertStats.BlockInserted != 1 {
		t.Fatalf("insert stats = %+v", insertStats)
	}
	if !archiveEventExists(t, ctx, db, insertHash.Hex()) {
		t.Fatal("node-filled event was not persisted")
	}

	// Simulate a partial prior run: arch_evtlog committed, then the dependent
	// transaction/block writes failed. A retry must repair those rows even
	// though the event log itself is already present.
	if _, err := db.SQL.ExecContext(ctx, `DELETE FROM arch_tx WHERE tx_hash = $1`, insertHash.Hex()); err != nil {
		t.Fatalf("deleting archived transaction: %v", err)
	}
	if _, err := db.SQL.ExecContext(ctx, `DELETE FROM arch_block WHERE block_num = $1`, int64(901)); err != nil {
		t.Fatalf("deleting archived block: %v", err)
	}
	repairStats, err := filler.RunProject(ctx, archive.ProjectRandomWalk, archive.NodeFillOptions{
		FromBlock: 901,
		EndBlock:  901,
		BatchSize: 1,
	})
	if err != nil {
		t.Fatalf("repair run: %v", err)
	}
	if repairStats.LogsSkipped != 1 ||
		repairStats.TxInserted != 1 ||
		repairStats.BlockInserted != 1 {
		t.Fatalf("repair stats = %+v", repairStats)
	}

	resumeStats, err := filler.RunProject(ctx, archive.ProjectRandomWalk, archive.NodeFillOptions{
		FromBlock: 901,
		EndBlock:  901,
		BatchSize: 1,
	})
	if err != nil {
		t.Fatalf("idempotent run: %v", err)
	}
	if resumeStats.LogsSkipped != 0 ||
		resumeStats.LogsInserted != 1 ||
		resumeStats.TxInserted != 1 ||
		resumeStats.BlockInserted != 0 ||
		resumeStats.BlockSkipped != 1 {
		t.Fatalf("idempotent stats = %+v", resumeStats)
	}

	oldBlockHash := chain.BlockHash(901).Hex()
	chain.Reorg(901)
	replacementHash := addChainLog(901)
	replacementBlockHash := chain.BlockHash(901).Hex()
	reorgStats, err := filler.RunProject(ctx, archive.ProjectRandomWalk, archive.NodeFillOptions{
		FromBlock: 901,
		EndBlock:  901,
		BatchSize: 1,
	})
	if err != nil {
		t.Fatalf("reorg convergence run: %v", err)
	}
	if reorgStats.LogsInserted != 1 || reorgStats.TxInserted != 1 || reorgStats.BlockInserted != 1 {
		t.Fatalf("reorg stats = %+v", reorgStats)
	}
	if !archiveEventExists(t, ctx, db, replacementHash.Hex()) {
		t.Fatal("replacement-fork event was not archived")
	}
	var (
		blockCount       int
		txCount          int
		eventCount       int
		currentBlockHash string
	)
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM arch_block WHERE block_num = $1`,
		int64(901),
	).Scan(&blockCount); err != nil {
		t.Fatalf("counting canonical archive blocks: %v", err)
	}
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM arch_tx WHERE tx_hash = $1`,
		insertHash.Hex(),
	).Scan(&txCount); err != nil {
		t.Fatalf("counting canonical archive transaction: %v", err)
	}
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM arch_evtlog WHERE block_num = $1`,
		int64(901),
	).Scan(&eventCount); err != nil {
		t.Fatalf("counting canonical archive events: %v", err)
	}
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT btrim(block_hash) FROM arch_block WHERE block_num = $1`,
		int64(901),
	).Scan(&currentBlockHash); err != nil {
		t.Fatalf("loading canonical archive block: %v", err)
	}
	if blockCount != 1 || txCount != 1 || eventCount != 1 ||
		currentBlockHash != replacementBlockHash || currentBlockHash == oldBlockHash {
		t.Fatalf(
			"post-reorg block=%d tx=%d event=%d hash=%s old=%s replacement=%s",
			blockCount,
			txCount,
			eventCount,
			currentBlockHash,
			oldBlockHash,
			replacementBlockHash,
		)
	}

	emptyReplacementTx := addChainLog(902)
	if _, err := filler.RunProject(ctx, archive.ProjectRandomWalk, archive.NodeFillOptions{
		FromBlock: 902,
		EndBlock:  902,
		BatchSize: 1,
	}); err != nil {
		t.Fatalf("initial block-902 fill: %v", err)
	}
	chain.Reorg(902) // canonical replacement deliberately emits no project log
	emptyReplacementHash := chain.BlockHash(902).Hex()
	emptyStats, err := filler.RunProject(ctx, archive.ProjectRandomWalk, archive.NodeFillOptions{
		FromBlock: 902,
		EndBlock:  902,
		BatchSize: 1,
	})
	if err != nil {
		t.Fatalf("empty replacement cleanup: %v", err)
	}
	var emptyEvents, emptyTransactions, emptyBlocks int
	var archivedEmptyHash string
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM arch_evtlog WHERE block_num = $1`,
		int64(902),
	).Scan(&emptyEvents); err != nil {
		t.Fatal(err)
	}
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM arch_tx WHERE tx_hash = $1`,
		emptyReplacementTx.Hex(),
	).Scan(&emptyTransactions); err != nil {
		t.Fatal(err)
	}
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT COUNT(*), COALESCE(MIN(btrim(block_hash)), '') FROM arch_block WHERE block_num = $1`,
		int64(902),
	).Scan(&emptyBlocks, &archivedEmptyHash); err != nil {
		t.Fatal(err)
	}
	if emptyStats.LogsFromNode != 0 || emptyEvents != 0 || emptyTransactions != 0 ||
		emptyBlocks != 1 || archivedEmptyHash != emptyReplacementHash {
		t.Fatalf(
			"empty replacement stats=%+v events=%d tx=%d blocks=%d hash=%s want=%s",
			emptyStats,
			emptyEvents,
			emptyTransactions,
			emptyBlocks,
			archivedEmptyHash,
			emptyReplacementHash,
		)
	}
}

func archiveEventExists(t *testing.T, ctx context.Context, db *testdb.DB, txHash string) bool {
	t.Helper()
	var exists bool
	if err := db.SQL.QueryRowContext(ctx,
		`SELECT EXISTS (SELECT 1 FROM arch_evtlog WHERE tx_hash = $1)`,
		txHash,
	).Scan(&exists); err != nil {
		t.Fatalf("checking archive event: %v", err)
	}
	return exists
}
