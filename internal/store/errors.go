package store

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// ErrNotFound reports that the requested row does not exist. Query methods
// return it (wrapped with operation context) instead of exposing pgx.ErrNoRows;
// test with errors.Is(err, store.ErrNotFound).
var ErrNotFound = errors.New("not found")

// ErrConflict reports a uniqueness violation: the row being inserted already
// exists. Test with errors.Is(err, store.ErrConflict).
var ErrConflict = errors.New("already exists")

// uniqueViolation is the PostgreSQL error code for unique_violation.
const uniqueViolation = "23505"

// WrapError attaches operation context to a query error and maps driver-level
// conditions onto the package sentinels: no-rows results (pgx or database/sql)
// become ErrNotFound, unique violations become ErrConflict (with the driver
// error preserved in the chain). A nil err returns nil, so callers can wrap
// unconditionally.
func WrapError(op string, err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, pgx.ErrNoRows) || errors.Is(err, sql.ErrNoRows):
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	case isUniqueViolation(err):
		return fmt.Errorf("%s: %w: %w", op, ErrConflict, err)
	default:
		return fmt.Errorf("%s: %w", op, err)
	}
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == uniqueViolation
}
