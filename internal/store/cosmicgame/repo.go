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
	rows, err := r.pool().Query(ctx, query, args...)
	if err != nil {
		return nil, store.WrapError(op, err)
	}
	defer rows.Close()
	records := make([]T, 0, capHint)
	for rows.Next() {
		var rec T
		if err := scanRow(rows, &rec); err != nil {
			return nil, store.WrapError(op, err)
		}
		records = append(records, rec)
	}
	if err := rows.Err(); err != nil {
		return nil, store.WrapError(op, err)
	}
	return records, nil
}
