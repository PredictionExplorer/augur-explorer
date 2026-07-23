//go:build integration

package store_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

// newTimeoutStore builds a Store with tight server-side bounds against the
// test container, the way the services set their production values.
func newTimeoutStore(t *testing.T, statement, idleInTx time.Duration) *store.Store {
	t.Helper()
	db := testdb.New(t)
	cfg := configFromConnString(t, db.ConnString)
	cfg.StatementTimeout = statement
	cfg.IdleInTxSessionTimeout = idleInTx
	st, err := store.New(context.Background(), cfg)
	if err != nil {
		t.Fatalf("store.New: %v", err)
	}
	t.Cleanup(st.Close)
	return st
}

// TestStatementTimeoutGUCsApplied proves the configured bounds reach every
// pool connection as the real PostgreSQL session GUCs.
func TestStatementTimeoutGUCsApplied(t *testing.T) {
	st := newTimeoutStore(t, 30*time.Second, 5*time.Minute)
	ctx := context.Background()

	var statement, idleInTx string
	if err := st.Pool().QueryRow(ctx, "SHOW statement_timeout").Scan(&statement); err != nil {
		t.Fatalf("SHOW statement_timeout: %v", err)
	}
	if err := st.Pool().QueryRow(ctx, "SHOW idle_in_transaction_session_timeout").Scan(&idleInTx); err != nil {
		t.Fatalf("SHOW idle_in_transaction_session_timeout: %v", err)
	}
	if statement != "30s" {
		t.Errorf("statement_timeout = %q, want 30s", statement)
	}
	if idleInTx != "5min" {
		t.Errorf("idle_in_transaction_session_timeout = %q, want 5min", idleInTx)
	}
}

// TestStatementTimeoutAbortsRunawayQuery proves the server-side backstop: a
// statement outliving the bound dies with SQLSTATE 57014 even though the
// client imposed no context deadline, and the pool stays healthy afterwards.
func TestStatementTimeoutAbortsRunawayQuery(t *testing.T) {
	st := newTimeoutStore(t, 200*time.Millisecond, 0)
	ctx := context.Background()

	start := time.Now()
	_, err := st.Pool().Exec(ctx, "SELECT pg_sleep(5)")
	elapsed := time.Since(start)
	if err == nil {
		t.Fatal("pg_sleep(5) survived a 200ms statement_timeout")
	}
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) || pgErr.Code != "57014" {
		t.Fatalf("err = %v, want PostgreSQL 57014 (query_canceled)", err)
	}
	if elapsed > 3*time.Second {
		t.Errorf("abort took %v, want ~200ms", elapsed)
	}

	var one int
	if err := st.Pool().QueryRow(ctx, "SELECT 1").Scan(&one); err != nil || one != 1 {
		t.Fatalf("pool unhealthy after statement abort: (%d, %v)", one, err)
	}
}

// TestContextDeadlineCancelsQueryPromptly proves the client-side half the
// API request deadline relies on: pgx cancels the running statement when the
// context deadline fires, the caller gets context.DeadlineExceeded fast, and
// the connection returns to the pool usable.
func TestContextDeadlineCancelsQueryPromptly(t *testing.T) {
	st := newTimeoutStore(t, 0, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err := st.Pool().Exec(ctx, "SELECT pg_sleep(10)")
	elapsed := time.Since(start)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v, want DeadlineExceeded", err)
	}
	if elapsed > 3*time.Second {
		t.Errorf("cancellation took %v, want ~150ms (the request deadline depends on this being prompt)", elapsed)
	}

	var one int
	if err := st.Pool().QueryRow(context.Background(), "SELECT 1").Scan(&one); err != nil || one != 1 {
		t.Fatalf("pool unhealthy after context cancellation: (%d, %v)", one, err)
	}
}

// TestIdleInTransactionTimeoutReleasesStalledTx proves the ADR-0010 backstop:
// a transaction stalled between statements (the shape of a wedged non-DB
// call inside InTx) is terminated by the server, InTx surfaces the failure,
// and the pool recovers.
func TestIdleInTransactionTimeoutReleasesStalledTx(t *testing.T) {
	st := newTimeoutStore(t, 0, 300*time.Millisecond)
	ctx := context.Background()

	err := st.InTx(ctx, func(txCtx context.Context) error {
		if _, err := st.Querier(txCtx).Exec(txCtx, "SELECT 1"); err != nil {
			return err
		}
		time.Sleep(900 * time.Millisecond) // the stalled non-DB call
		_, err := st.Querier(txCtx).Exec(txCtx, "SELECT 1")
		return err
	})
	if err == nil {
		t.Fatal("a transaction idle past the bound must fail, not commit")
	}

	var one int
	if err := st.Pool().QueryRow(ctx, "SELECT 1").Scan(&one); err != nil || one != 1 {
		t.Fatalf("pool unhealthy after idle-in-transaction termination: (%d, %v)", one, err)
	}
}
