// Context-scoped transaction support: InTx opens one pgx transaction and
// carries it through the context, so every Store method and every domain
// repo method (they all resolve their querier via Store.Querier) joins the
// transaction without signature changes. This is the seam the indexer uses
// to commit a block's writes and its processing watermark atomically
// (docs/adr/0010-transactional-ingestion.md).

package store

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Querier is the query surface shared by the connection pool and an open
// transaction; every Store and repo query runs through one. Both
// *pgxpool.Pool and pgx.Tx satisfy it.
type Querier interface {
	// Exec executes a statement that returns no rows.
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	// Query executes a rows-returning statement.
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	// QueryRow executes a statement expected to return at most one row.
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// txState carries an open transaction through a context: the owning Store
// (so a foreign Store's methods never run on another Store's transaction),
// the transaction itself, and the overlay of address-cache entries created
// inside it (flushed into the shared LRU only on commit).
type txState struct {
	owner   *Store
	tx      pgx.Tx
	overlay map[string]int64
}

// txKey is the context key carrying a *txState.
type txKey struct{}

// txState returns the transaction state ctx carries when — and only when —
// it was begun by this Store's InTx; any other transaction state is
// ignored, so multi-Store processes and test harnesses cannot leak a
// transaction across pools.
func (s *Store) txState(ctx context.Context) *txState {
	if st, ok := ctx.Value(txKey{}).(*txState); ok && st.owner == s {
		return st
	}
	return nil
}

// Querier resolves the query surface for ctx: the transaction the context
// carries when it was begun by this Store's InTx, otherwise the shared
// pool. Domain repos in subpackages call it for every query, which is what
// makes all of them transaction-aware without signature changes.
func (s *Store) Querier(ctx context.Context) Querier {
	if st := s.txState(ctx); st != nil {
		return st.tx
	}
	return s.pool
}

// q is the package-internal shorthand for Querier.
func (s *Store) q(ctx context.Context) Querier { return s.Querier(ctx) }

// InTx runs fn inside one database transaction. The context passed to fn
// carries the transaction: every Store method and every domain repo method
// backed by this Store joins it automatically, and address-cache entries
// created inside the transaction stay in a per-transaction overlay until
// commit (a rollback must not poison the shared cache with ids whose rows
// vanished).
//
// A non-nil error from fn rolls the transaction back and is returned
// unchanged; Begin/Commit failures are wrapped. When ctx already carries a
// transaction begun by this Store, fn joins it directly (no savepoint): the
// enclosing InTx still owns commit and rollback.
//
// The transaction is single-flight state: fn must not share its context
// across goroutines (pgx.Tx itself is not safe for concurrent use).
func (s *Store) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	if s.txState(ctx) != nil {
		return fn(ctx)
	}
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return WrapError("begin transaction", err)
	}
	st := &txState{owner: s, tx: tx}
	committed := false
	defer func() {
		if !committed {
			// Error or panic: discard the transaction. Rollback on an
			// already-failed connection just closes it; the pool replenishes.
			_ = tx.Rollback(ctx)
		}
	}()
	if err := fn(context.WithValue(ctx, txKey{}, st)); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return WrapError("commit transaction", err)
	}
	committed = true
	s.addrCache.putAll(st.overlay)
	return nil
}
