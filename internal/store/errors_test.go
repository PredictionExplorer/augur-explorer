package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func TestWrapErrorNil(t *testing.T) {
	if err := WrapError("op", nil); err != nil {
		t.Errorf("WrapError(op, nil) = %v, want nil", err)
	}
}

func TestWrapErrorNotFound(t *testing.T) {
	for _, src := range []error{pgx.ErrNoRows, sql.ErrNoRows, fmt.Errorf("scan: %w", pgx.ErrNoRows)} {
		err := WrapError("prize info round=5", src)
		if !errors.Is(err, ErrNotFound) {
			t.Errorf("WrapError(%v) = %v, want ErrNotFound in chain", src, err)
		}
		if !strings.Contains(err.Error(), "prize info round=5") {
			t.Errorf("error %q missing operation context", err)
		}
	}
}

func TestWrapErrorConflict(t *testing.T) {
	pgErr := &pgconn.PgError{Code: uniqueViolation, ConstraintName: "cg_banned_bids_pkey"}
	err := WrapError("insert banned bid", pgErr)
	if !errors.Is(err, ErrConflict) {
		t.Errorf("WrapError(unique violation) = %v, want ErrConflict in chain", err)
	}
	// The driver error must stay reachable for callers that need details.
	var got *pgconn.PgError
	if !errors.As(err, &got) || got.ConstraintName != "cg_banned_bids_pkey" {
		t.Errorf("WrapError lost the pgconn.PgError: %v", err)
	}
}

func TestWrapErrorOther(t *testing.T) {
	src := errors.New("connection reset")
	err := WrapError("list bids", src)
	if errors.Is(err, ErrNotFound) || errors.Is(err, ErrConflict) {
		t.Errorf("WrapError mapped a generic error to a sentinel: %v", err)
	}
	if !errors.Is(err, src) {
		t.Errorf("WrapError broke the error chain: %v", err)
	}
	// Non-unique-violation Postgres errors are not conflicts.
	fk := &pgconn.PgError{Code: "23503"} // foreign_key_violation
	if errors.Is(WrapError("op", fk), ErrConflict) {
		t.Error("foreign key violation mapped to ErrConflict")
	}
}
