//go:build integration

// Integration tests for the context-scoped transaction seam (tx.go): commit
// and rollback atomicity, cache-overlay flushing, isolation from the pool,
// the wrong-Store guard, nested joins, panic safety and the in-transaction
// address-create conflict that the pre-ON CONFLICT code could not survive
// (SQLSTATE 25P02).
package store_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

// addressRowExists reads the address table through the raw pool, bypassing
// every cache, so tests can observe exactly what has been committed.
func addressRowExists(t *testing.T, db *testdb.DB, addr string) bool {
	t.Helper()
	var n int
	if err := db.Pool.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM address WHERE addr=$1", addr).Scan(&n); err != nil {
		t.Fatalf("counting address rows: %v", err)
	}
	return n > 0
}

func TestInTxCommitsAtomicallyAndFlushesCache(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	st := store.NewFromPool(db.Pool)

	const addr = "0x00000000000000000000000000000000000000C1"
	var aid int64
	err := st.InTx(ctx, func(txCtx context.Context) error {
		var err error
		aid, err = st.LookupOrCreateAddress(txCtx, addr, 11, 0)
		if err != nil {
			return err
		}
		if err := st.SetLastBlockNum(txCtx, 11); err != nil {
			return err
		}
		// The uncommitted write is invisible to the pool: the insert really
		// runs inside the transaction, not autocommitted on a pool conn.
		if addressRowExists(t, db, addr) {
			t.Error("address row visible outside the transaction before commit")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("InTx: %v", err)
	}

	if !addressRowExists(t, db, addr) {
		t.Fatal("address row missing after commit")
	}
	last, err := st.LastBlockNum(ctx)
	if err != nil || last != 11 {
		t.Errorf("LastBlockNum after commit = (%d, %v), want (11, nil)", last, err)
	}

	// Flush-on-commit proof: delete the row behind the cache's back — the id
	// must still resolve from the shared LRU without touching the table.
	if _, err := db.Pool.Exec(ctx, "DELETE FROM address WHERE addr=$1", addr); err != nil {
		t.Fatalf("deleting address row: %v", err)
	}
	cached, err := st.LookupAddressID(ctx, addr)
	if err != nil || cached != aid {
		t.Errorf("LookupAddressID after delete = (%d, %v), want cached (%d, nil)", cached, err, aid)
	}
}

func TestInTxRollbackDiscardsWritesAndCache(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	st := store.NewFromPool(db.Pool)

	const addr = "0x00000000000000000000000000000000000000C2"
	sentinel := errors.New("handler failed")
	err := st.InTx(ctx, func(txCtx context.Context) error {
		if _, err := st.LookupOrCreateAddress(txCtx, addr, 12, 0); err != nil {
			return err
		}
		if err := st.SetLastBlockNum(txCtx, 999); err != nil {
			return err
		}
		return sentinel
	})
	if !errors.Is(err, sentinel) {
		t.Fatalf("InTx = %v, want the callback's own error", err)
	}

	if addressRowExists(t, db, addr) {
		t.Error("address row survived the rollback")
	}
	last, err := st.LastBlockNum(ctx)
	if err != nil || last == 999 {
		t.Errorf("LastBlockNum after rollback = (%d, %v), want the watermark untouched", last, err)
	}
	// The rolled-back id must not have poisoned the shared cache: a fresh
	// lookup must miss (cache and table agree the address does not exist).
	if _, err := st.LookupAddressID(ctx, addr); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("LookupAddressID after rollback = %v, want ErrNotFound (clean cache)", err)
	}
}

func TestInTxWrongStoreUsesItsOwnPool(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	stA := store.NewFromPool(db.Pool)
	stB := store.NewFromPool(db.Pool)

	const addr = "0x00000000000000000000000000000000000000C3"
	err := stA.InTx(ctx, func(txCtx context.Context) error {
		if _, err := stA.LookupOrCreateAddress(txCtx, addr, 13, 0); err != nil {
			return err
		}
		// Store B receives A's transaction context but must resolve through
		// its own pool: A's uncommitted row is invisible to it.
		if _, err := stB.LookupAddressID(txCtx, addr); !errors.Is(err, store.ErrNotFound) {
			t.Errorf("foreign Store saw another Store's uncommitted row (err=%v)", err)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("InTx: %v", err)
	}
}

func TestInTxNestedJoinsTheOuterTransaction(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	st := store.NewFromPool(db.Pool)

	const addr = "0x00000000000000000000000000000000000000C4"
	sentinel := errors.New("inner failure")
	err := st.InTx(ctx, func(outerCtx context.Context) error {
		if _, err := st.LookupOrCreateAddress(outerCtx, addr, 14, 0); err != nil {
			return err
		}
		return st.InTx(outerCtx, func(innerCtx context.Context) error {
			// The nested call joined the same transaction: the outer write
			// is visible through it.
			if _, err := st.LookupAddressID(innerCtx, addr); err != nil {
				t.Errorf("nested transaction cannot see the outer write: %v", err)
			}
			return sentinel
		})
	})
	if !errors.Is(err, sentinel) {
		t.Fatalf("InTx = %v, want the inner error", err)
	}
	if addressRowExists(t, db, addr) {
		t.Error("outer write survived: the nested join must share one transaction")
	}
}

func TestInTxPanicRollsBack(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	st := store.NewFromPool(db.Pool)

	const addr = "0x00000000000000000000000000000000000000C5"
	func() {
		defer func() {
			if recover() == nil {
				t.Error("the callback's panic must propagate")
			}
		}()
		_ = st.InTx(ctx, func(txCtx context.Context) error {
			if _, err := st.LookupOrCreateAddress(txCtx, addr, 15, 0); err != nil {
				return err
			}
			panic("handler exploded")
		})
	}()

	if addressRowExists(t, db, addr) {
		t.Error("address row survived the panic rollback")
	}
	if _, err := st.LookupAddressID(ctx, addr); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("LookupAddressID after panic = %v, want ErrNotFound", err)
	}
}

// TestInTxRepoMethodsJoinTheTransaction proves the domain repos resolve the
// context-carried transaction: a watermark write through the RandomWalk
// repo inside a failed InTx leaves the status row untouched.
func TestInTxRepoMethodsJoinTheTransaction(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	st := store.NewFromPool(db.Pool)
	repo := rwstore.NewRepo(st)

	before, err := repo.ProcessingStatus(ctx) // self-seeds the singleton row
	if err != nil {
		t.Fatalf("ProcessingStatus: %v", err)
	}

	sentinel := errors.New("later failure")
	err = st.InTx(ctx, func(txCtx context.Context) error {
		status := before
		status.LastBlockNum = before.LastBlockNum + 777
		if err := repo.UpdateProcessingStatus(txCtx, &status); err != nil {
			return err
		}
		// Inside the transaction the repo reads its own uncommitted write.
		mid, err := repo.ProcessingStatus(txCtx)
		if err != nil {
			return err
		}
		if mid.LastBlockNum != status.LastBlockNum {
			t.Errorf("repo read inside tx = %d, want %d", mid.LastBlockNum, status.LastBlockNum)
		}
		return sentinel
	})
	if !errors.Is(err, sentinel) {
		t.Fatalf("InTx = %v, want the sentinel", err)
	}

	after, err := repo.ProcessingStatus(ctx)
	if err != nil {
		t.Fatalf("ProcessingStatus after rollback: %v", err)
	}
	if after.LastBlockNum != before.LastBlockNum {
		t.Errorf("watermark after rollback = %d, want %d (repo write must join the tx)",
			after.LastBlockNum, before.LastBlockNum)
	}
}

// TestInTxBeginFailure pins the begin-transaction error arm: a closed pool
// cannot open a transaction and the callback must never run.
func TestInTxBeginFailure(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()

	cfg := db.Pool.Config().Copy()
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		t.Fatalf("building second pool: %v", err)
	}
	st := store.NewFromPool(pool)
	st.Close()

	ran := false
	err = st.InTx(ctx, func(context.Context) error { ran = true; return nil })
	if err == nil || !strings.Contains(err.Error(), "begin transaction") {
		t.Fatalf("InTx on closed pool = %v, want wrapped begin failure", err)
	}
	if ran {
		t.Error("callback ran despite the failed begin")
	}
}

// TestInTxCommitFailureDiscardsWrites pins the commit error arm: a context
// canceled between the callback and the commit fails the commit, the error
// is wrapped, and the transaction's writes are gone.
func TestInTxCommitFailureDiscardsWrites(t *testing.T) {
	db := testdb.New(t)
	st := store.NewFromPool(db.Pool)

	const addr = "0x00000000000000000000000000000000000000C8"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := st.InTx(ctx, func(txCtx context.Context) error {
		if _, err := st.LookupOrCreateAddress(txCtx, addr, 18, 0); err != nil {
			return err
		}
		cancel() // the commit must fail
		return nil
	})
	if err == nil || !strings.Contains(err.Error(), "commit transaction") {
		t.Fatalf("InTx with canceled commit = %v, want wrapped commit failure", err)
	}

	if addressRowExists(t, db, addr) {
		t.Error("address row survived the failed commit")
	}
	if _, err := st.LookupAddressID(context.Background(), addr); !errors.Is(err, store.ErrNotFound) {
		t.Errorf("LookupAddressID after failed commit = %v, want ErrNotFound (clean cache)", err)
	}
}

// TestProcessingStatusSeedInsertFailurePropagates drives the lazy
// watermark-row seeding of both domain repos into its insert-failure arm:
// on a fresh database (no status row) a read-only session turns the
// seeding INSERT into an error, which must surface wrapped instead of
// being swallowed as a zero watermark.
func TestProcessingStatusSeedInsertFailurePropagates(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()

	cfg := db.Pool.Config().Copy()
	cfg.ConnConfig.RuntimeParams["default_transaction_read_only"] = "on"
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		t.Fatalf("building read-only pool: %v", err)
	}
	st := store.NewFromPool(pool)
	defer st.Close()

	if _, err := cgstore.NewRepo(st).ProcessingStatus(ctx); err == nil ||
		!strings.Contains(err.Error(), "insert default row") {
		t.Errorf("CosmicGame ProcessingStatus on read-only session = %v, want wrapped seed-insert failure", err)
	}
	if _, err := rwstore.NewRepo(st).ProcessingStatus(ctx); err == nil ||
		!strings.Contains(err.Error(), "insert default row") {
		t.Errorf("RandomWalk ProcessingStatus on read-only session = %v, want wrapped seed-insert failure", err)
	}
}

// TestLookupOrCreateAddressConflictInsideTx is the regression test for the
// in-transaction create race: a plain INSERT losing the unique race raises
// 23505, which aborts the whole enclosing transaction (every later
// statement fails with 25P02). With ON CONFLICT DO NOTHING the loser keeps
// its transaction healthy, adopts the winner's id and continues.
//
// Orchestration: a raw transaction inserts the address and stays open, the
// InTx side blocks on the speculative-insert lock, and the raw transaction
// commits only once pg_stat_activity shows the waiter — so the conflict
// branch is hit deterministically.
func TestLookupOrCreateAddressConflictInsideTx(t *testing.T) {
	db := testdb.New(t)
	ctx := context.Background()
	st := store.NewFromPool(db.Pool)

	const addr = "0x00000000000000000000000000000000000000C6"
	const followUp = "0x00000000000000000000000000000000000000C7"

	winner, err := db.Pool.Begin(ctx)
	if err != nil {
		t.Fatalf("beginning the winning transaction: %v", err)
	}
	defer func() { _ = winner.Rollback(ctx) }()
	var winnerID int64
	if err := winner.QueryRow(ctx,
		"INSERT INTO address(addr,block_num,tx_id) VALUES($1,16,0) RETURNING address_id",
		addr).Scan(&winnerID); err != nil {
		t.Fatalf("winning insert: %v", err)
	}

	type result struct {
		aid int64
		err error
	}
	loserDone := make(chan result, 1)
	go func() {
		var aid int64
		err := st.InTx(ctx, func(txCtx context.Context) error {
			var err error
			// SELECT misses (the winner is uncommitted), INSERT blocks on
			// the winner's speculative lock, resolves to a conflict once it
			// commits, and the re-read adopts the winner's id.
			aid, err = st.LookupOrCreateAddress(txCtx, addr, 17, 0)
			if err != nil {
				return err
			}
			// The transaction must still be healthy — the pre-fix code dies
			// here with "current transaction is aborted" (25P02).
			if _, err := st.LookupOrCreateAddress(txCtx, followUp, 17, 0); err != nil {
				return fmt.Errorf("transaction poisoned after conflict: %w", err)
			}
			return nil
		})
		loserDone <- result{aid: aid, err: err}
	}()

	// Commit the winner only once the loser is provably waiting on the lock.
	deadline := time.After(15 * time.Second)
	for waiting := false; !waiting; {
		select {
		case r := <-loserDone:
			t.Fatalf("loser finished before the winner committed: (%d, %v)", r.aid, r.err)
		case <-deadline:
			t.Fatal("loser never blocked on the conflicting insert")
		case <-time.After(10 * time.Millisecond):
		}
		if err := db.Pool.QueryRow(ctx, `SELECT EXISTS (
				SELECT 1 FROM pg_stat_activity
				WHERE wait_event_type='Lock' AND query LIKE 'INSERT INTO address%'
			)`).Scan(&waiting); err != nil {
			t.Fatalf("polling pg_stat_activity: %v", err)
		}
	}
	if err := winner.Commit(ctx); err != nil {
		t.Fatalf("committing the winner: %v", err)
	}

	r := <-loserDone
	if r.err != nil {
		t.Fatalf("InTx after conflict = %v, want nil (healthy transaction)", r.err)
	}
	if r.aid != winnerID {
		t.Errorf("conflict resolution returned id %d, want the winner's %d", r.aid, winnerID)
	}
	if !addressRowExists(t, db, followUp) {
		t.Error("follow-up insert missing: the loser's transaction did not stay usable")
	}
}
