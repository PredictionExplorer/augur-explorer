// Generic list-query helper shared by the domain repos (cosmicgame,
// randomwalk). Lives in the base package so each repo's thin wrapper can
// delegate instead of duplicating the scan loop.

package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// QueryList runs a SELECT and scans every row with scanRow, wrapping any
// failure with op context via WrapError. The result is always a non-nil
// slice (capacity capHint) so an empty result marshals as JSON [] — the
// shape every legacy caller and golden file relies on.
func QueryList[T any](ctx context.Context, q Querier, op string, capHint int, query string, scanRow func(pgx.Rows, *T) error, args ...any) ([]T, error) {
	rows, err := q.Query(ctx, query, args...)
	if err != nil {
		return nil, WrapError(op, err)
	}
	defer rows.Close()
	records := make([]T, 0, capHint)
	for rows.Next() {
		var rec T
		if err := scanRow(rows, &rec); err != nil {
			return nil, WrapError(op, err)
		}
		records = append(records, rec)
	}
	if err := rows.Err(); err != nil {
		return nil, WrapError(op, err)
	}
	return records, nil
}
