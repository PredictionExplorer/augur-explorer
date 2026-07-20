//go:build integration

package cosmicgame

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testutil"
)

// stageStoryTransaction materializes one scripted transaction on the fake
// chain without ingesting it.
func stageStoryTransaction(t *testing.T, tx storyTx) {
	t.Helper()
	logs := make([]*types.Log, 0, len(tx.logs))
	for _, fixture := range tx.logs {
		logs = append(logs, fixture.build(t))
	}
	stageTx(t, tx.block, addr(tx.to), 0, logs)
}

// faultHarnessTable temporarily renames one table and resets pooled
// connections so pgx cannot reuse a cached statement description.
func faultHarnessTable(t *testing.T, table string) func() {
	t.Helper()
	backup := table + "_atomicity_backup"
	tableID := pgx.Identifier{table}.Sanitize()
	backupID := pgx.Identifier{backup}.Sanitize()
	rename := func(from, to string) {
		t.Helper()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if _, err := testDB.Pool.Exec(ctx, "ALTER TABLE "+from+" RENAME TO "+to); err != nil {
			t.Fatalf("renaming %s to %s: %v", from, to, err)
		}
		testDB.Pool.Reset()
	}
	rename(tableID, backupID)
	restored := false
	return func() {
		if restored {
			return
		}
		restored = true
		rename(backupID, tableID)
	}
}

// TestRealHandlersRollbackAndRetryConverges proves the ADR-0010 guarantee
// through the actual CosmicGame handlers and plpgsql triggers. Block 5000
// commits, block 5001 performs two real writes and then fails in InsertBid;
// the failed block must vanish while the earlier block remains byte-identical
// to a clean one-block ingest. Restoring cg_bid and retrying must converge to
// the exact state of a clean two-block run.
func TestRealHandlersRollbackAndRetryConverges(t *testing.T) {
	resetDB(t)
	const firstBlock = int64(5000)
	testChain.Reorg(firstBlock) // make -count=N runs reuse-safe
	story := scriptedRound(firstBlock)
	prior := story[0]
	failing := story[1]
	stageStoryTransaction(t, prior)
	stageStoryTransaction(t, failing)

	if err := runHarnessRange(t, firstBlock, firstBlock, 1); err != nil {
		t.Fatalf("clean prior block: %v", err)
	}
	expectedPrior := snapshot(t)

	resetDB(t)
	restore := faultHarnessTable(t, "cg_bid")
	defer restore()
	err := runHarnessRange(t, firstBlock, failing.block, 2)
	if err == nil || !strings.Contains(err.Error(), "stage process") {
		restore()
		t.Fatalf("faulted run error = %v, want process-stage failure", err)
	}
	restore()

	requireNoDiff(t, expectedPrior, snapshot(t), "real-handler rollback")
	if _, err := dbStore.LookupAddressID(context.Background(), fxAlice); err == nil {
		t.Fatal("rolled-back bidder address remained visible")
	} else if !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("rolled-back bidder lookup = %v, want not found", err)
	}

	if err := runHarnessRange(t, failing.block, failing.block, 1); err != nil {
		t.Fatalf("retry failed block: %v", err)
	}
	retried := snapshot(t)

	resetDB(t)
	if err := runHarnessRange(t, firstBlock, failing.block, 2); err != nil {
		t.Fatalf("clean control run: %v", err)
	}
	requireNoDiff(t, retried, snapshot(t), "retry versus clean control")
}

// TestRealHandlerWritesAndWatermarkBecomeVisibleTogether holds the
// transaction after every real handler has written but before the progress
// update. A second connection must still observe the pre-block snapshot;
// after release, handler rows, trigger effects, layer-1 rows and watermark
// become visible together.
func TestRealHandlerWritesAndWatermarkBecomeVisibleTogether(t *testing.T) {
	resetDB(t)
	const blockNum = int64(5100)
	testChain.Reorg(blockNum) // make -count=N runs reuse-safe
	tx := scriptedRound(blockNum)[0]
	stageStoryTransaction(t, tx)
	progress := primeHarnessProgress(t, blockNum)
	baseline := snapshot(t)

	handlersDone := make(chan struct{}, 1)
	release := make(chan struct{})
	released := false
	defer func() {
		if !released {
			close(release)
		}
	}()
	processed := 0
	process := indexer.ProcessFunc(func(ctx context.Context, evtID int64) error {
		if err := testProcess(ctx, evtID); err != nil {
			return err
		}
		processed++
		if processed == len(tx.logs) {
			handlersDone <- struct{}{}
			<-release
		}
		return nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	progress.onSet = func(block int64) {
		if block == blockNum {
			cancel()
		}
	}
	engine := newHarnessRunEngine(t, progress, process, 1)
	done := make(chan error, 1)
	go func() {
		done <- engine.Run(ctx)
	}()

	select {
	case <-handlersDone:
	case <-ctx.Done():
		t.Fatal("handlers did not reach the pre-commit barrier")
	}
	requireNoDiff(t, baseline, snapshot(t), "uncommitted real-handler writes")
	released = true
	close(release)

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("Engine.Run: %v", err)
		}
	case <-time.After(30 * time.Second):
		t.Fatal("Engine.Run did not return after releasing commit")
	}

	committed := snapshot(t)
	diff, err := testutil.DiffSnapshots(baseline, committed)
	if err != nil {
		t.Fatalf("diffing committed state: %v", err)
	}
	if strings.TrimSpace(string(diff)) == "{}" {
		t.Fatal("commit published no state")
	}
	status, err := cgRepo.ProcessingStatus(context.Background())
	if err != nil {
		t.Fatalf("reading committed progress: %v", err)
	}
	if status.LastBlockNum != blockNum {
		t.Fatalf("committed progress = %d, want %d", status.LastBlockNum, blockNum)
	}
}
