package cosmicgame

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Repo executes CosmicGame queries on the shared store pool. It is the
// Phase 1 replacement for SQLStorageWrapper: methods take a context as their
// first parameter, return errors instead of exiting, use idiomatic names and
// run natively on pgx. Files are converted one at a time; anything not yet
// converted remains a SQLStorageWrapper method.
type Repo struct {
	store *store.Store
}

// NewRepo returns a Repo backed by st's connection pool.
func NewRepo(st *store.Store) *Repo {
	return &Repo{store: st}
}

// q resolves the querier for ctx — the transaction the context carries when
// one was begun by this repo's Store (store.InTx), otherwise the shared
// pool. Every query method runs through it, which is what lets the indexer
// commit a block's domain writes atomically with its watermark.
func (r *Repo) q(ctx context.Context) store.Querier { return r.store.Querier(ctx) }

// pool exposes the raw connection pool for the integration tests' fixture
// mutation and cleanup SQL, which must run outside any transaction.
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
	return store.QueryList(ctx, r.q(ctx), op, capHint, query, scanRow, args...)
}
