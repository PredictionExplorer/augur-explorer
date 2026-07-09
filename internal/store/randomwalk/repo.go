// Package randomwalk provides the RandomWalk domain's database access: a
// Repo of context-first, error-returning, pgx-native query methods over the
// shared store pool (ETL writes, marketplace/token API reads, notification
// watermarks and the beauty-contest ranking storage).
package randomwalk

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Repo executes RandomWalk queries on the shared store pool. It is the
// Phase 1 replacement for SQLStorageWrapper: methods take a context as their
// first parameter, return errors instead of exiting, use idiomatic names and
// run natively on pgx.
type Repo struct {
	store *store.Store
}

// NewRepo returns a Repo backed by st's connection pool.
func NewRepo(st *store.Store) *Repo {
	return &Repo{store: st}
}

// pool is a shorthand used by the query methods.
func (r *Repo) pool() *pgxpool.Pool { return r.store.Pool() }

// addrID resolves addr to its address_id through the shared Store's
// lookup-or-create cache; the insert methods use it for every foreign-key
// address column.
func (r *Repo) addrID(ctx context.Context, addr string, blockNum, txID int64) (int64, error) {
	return r.store.LookupOrCreateAddress(ctx, addr, blockNum, txID)
}

// queryList runs a SELECT and scans every row with scanRow, wrapping any
// failure with op context. The result is always a non-nil slice (capacity
// capHint) so an empty result marshals as [] — the shape every legacy caller
// and golden file relies on.
func queryList[T any](ctx context.Context, r *Repo, op string, capHint int, query string, scanRow func(pgx.Rows, *T) error, args ...any) ([]T, error) {
	return store.QueryList(ctx, r.pool(), op, capHint, query, scanRow, args...)
}
