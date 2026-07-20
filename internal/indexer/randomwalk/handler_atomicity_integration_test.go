//go:build integration

package randomwalk

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

// stageAtomicityScenario records three production-shaped blocks: a mint, a
// sell offer, then ItemBought followed by Transfer in one transaction. The
// final block is deliberately ordered so a failing Transfer handler must
// roll back the already-applied ItemBought handler and its profit triggers.
func stageAtomicityScenario(t *testing.T, firstBlock int64) {
	t.Helper()
	const tokenID = int64(77)
	stageTx(t, firstBlock, addr(fxRandomWalkAddr), 0, []*types.Log{
		mintLog(tokenID, fxCarol)(t),
	})
	stageTx(t, firstBlock+1, addr(fxMarketplaceAddr), 0, []*types.Log{
		offerLog(1, tokenID, fxCarol, fxZero, 3)(t),
	})
	stageTx(t, firstBlock+2, addr(fxMarketplaceAddr), 0, []*types.Log{
		buildLog(t, fxMarketABI, "ItemBought", addr(fxMarketplaceAddr),
			[]any{bigInt(1), addr(fxCarol), addr(fxDave)}, nil),
		transferLog(fxCarol, fxDave, tokenID)(t),
	})
}

// TestRealHandlersRollbackAndRetryConverges proves the ADR-0010 guarantee
// through actual RandomWalk handlers and triggers. The mint and offer blocks
// commit; ItemBought writes, then Transfer fails because rw_transfer is
// hidden. The entire sale block must vanish, and its retry must converge to
// exactly the state of a clean three-block run.
func TestRealHandlersRollbackAndRetryConverges(t *testing.T) {
	resetDB(t)
	const firstBlock = int64(4000)
	const targetBlock = firstBlock + 2
	testChain.Reorg(firstBlock) // make -count=N runs reuse-safe
	stageAtomicityScenario(t, firstBlock)

	if err := runHarnessRange(t, firstBlock, firstBlock+1, 2); err != nil {
		t.Fatalf("clean prerequisite blocks: %v", err)
	}
	expectedPrior := snapshot(t)

	resetDB(t)
	restore := faultHarnessTable(t, "rw_transfer")
	defer restore()
	err := runHarnessRange(t, firstBlock, targetBlock, 3)
	if err == nil || !strings.Contains(err.Error(), "stage process") {
		restore()
		t.Fatalf("faulted run error = %v, want process-stage failure", err)
	}
	restore()

	requireNoDiff(t, expectedPrior, snapshot(t), "real-handler rollback")
	if _, err := dbStore.LookupAddressID(context.Background(), fxDave); err == nil {
		t.Fatal("rolled-back buyer address remained visible")
	} else if !errors.Is(err, store.ErrNotFound) {
		t.Fatalf("rolled-back buyer lookup = %v, want not found", err)
	}

	if err := runHarnessRange(t, targetBlock, targetBlock, 1); err != nil {
		t.Fatalf("retry failed block: %v", err)
	}
	retried := snapshot(t)

	resetDB(t)
	if err := runHarnessRange(t, firstBlock, targetBlock, 3); err != nil {
		t.Fatalf("clean control run: %v", err)
	}
	requireNoDiff(t, retried, snapshot(t), "retry versus clean control")
}

// TestRealHandlerWritesAndWatermarkBecomeVisibleTogether holds a mint
// transaction after the real handler writes but before the progress update.
// A second connection must see the pre-block state until the transaction is
// released, then observe the mint, trigger aggregates, layer-1 rows and
// watermark together.
func TestRealHandlerWritesAndWatermarkBecomeVisibleTogether(t *testing.T) {
	resetDB(t)
	const blockNum = int64(4100)
	testChain.Reorg(blockNum) // make -count=N runs reuse-safe
	stageTx(t, blockNum, addr(fxRandomWalkAddr), 0, []*types.Log{
		mintLog(88, fxAlice)(t),
	})
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
	process := indexer.ProcessFunc(func(ctx context.Context, evtID int64) error {
		if err := testProcess(ctx, evtID); err != nil {
			return err
		}
		handlersDone <- struct{}{}
		<-release
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
		t.Fatal("handler did not reach the pre-commit barrier")
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
	status, err := rwRepo.ProcessingStatus(context.Background())
	if err != nil {
		t.Fatalf("reading committed progress: %v", err)
	}
	if status.LastBlockNum != blockNum {
		t.Fatalf("committed progress = %d, want %d", status.LastBlockNum, blockNum)
	}
}
