package testutil

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// PgxOp is one scripted database operation served by ScriptedPgx, in order.
// Kind is "query" (Query and QueryRow both consume it) or "exec".
type PgxOp struct {
	Kind     string   // "query" or "exec"
	Contains string   // substring the whitespace-normalized SQL must contain
	Parts    []string // additional required substrings
	ArgCount *int     // exact bound-argument count when non-nil

	// Query results. RowsErrAt is the row index at which iteration stops
	// with RowsErr; -1 (the default set by NewScriptedPgx for a zero value
	// of 0 with RowsErr set... see normalize) means after all rows.
	Rows      [][]any
	RowsErr   error
	RowsErrAt int

	// Exec result.
	Affected int64

	// Err fails the operation itself (Query/Exec error; QueryRow defers it
	// to Scan like pgx does).
	Err error

	// OnCall runs when the operation is consumed (e.g. cancel a context).
	OnCall func()
}

// ScriptedPgx serves scripted results through the pgx query surface the ops
// packages consume (Exec/Query/QueryRow) plus the pool-shaped Ping/Close,
// so it satisfies both the per-package Querier interfaces and command
// wiring seams. It is safe for concurrent use.
type ScriptedPgx struct {
	mu         sync.Mutex
	ops        []PgxOp
	violations []string
	closed     atomic.Bool
}

// NewScriptedPgx builds a scripted querier serving ops in order.
func NewScriptedPgx(ops ...PgxOp) *ScriptedPgx {
	return &ScriptedPgx{ops: append([]PgxOp(nil), ops...)}
}

// Ping mimics pgxpool.Pool.Ping against the scripted handle.
func (s *ScriptedPgx) Ping(context.Context) error {
	if s.closed.Load() {
		return errors.New("scripted querier is closed")
	}
	return nil
}

// Close marks the handle closed like pgxpool.Pool.Close.
func (s *ScriptedPgx) Close() { s.closed.Store(true) }

// Closed reports whether Close was called.
func (s *ScriptedPgx) Closed() bool { return s.closed.Load() }

// AssertDone fails the test when scripted operations remain unconsumed or
// any recorded expectation was violated.
func (s *ScriptedPgx) AssertDone(tb testing.TB) {
	tb.Helper()
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.ops) != 0 {
		tb.Errorf("script has %d unconsumed operations", len(s.ops))
	}
	for _, violation := range s.violations {
		tb.Error(violation)
	}
}

func (s *ScriptedPgx) next(kind, query string, args []any) (PgxOp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	normalized := CompactSQL(query)
	if len(s.ops) == 0 {
		return PgxOp{}, fmt.Errorf("unexpected %s: %s", kind, normalized)
	}
	op := s.ops[0]
	s.ops = s.ops[1:]
	if op.Kind != kind {
		return PgxOp{}, fmt.Errorf("got %s %q, want %s containing %q", kind, normalized, op.Kind, op.Contains)
	}
	if op.Contains != "" && !strings.Contains(normalized, CompactSQL(op.Contains)) {
		return PgxOp{}, fmt.Errorf("%s %q does not contain %q", kind, normalized, CompactSQL(op.Contains))
	}
	for _, part := range op.Parts {
		if !strings.Contains(normalized, CompactSQL(part)) {
			s.violations = append(s.violations, fmt.Sprintf("%s %q does not contain %q", kind, normalized, part))
		}
	}
	if op.ArgCount != nil && len(args) != *op.ArgCount {
		s.violations = append(s.violations, fmt.Sprintf("%s %q has %d args, want %d", kind, normalized, len(args), *op.ArgCount))
	}
	if op.OnCall != nil {
		op.OnCall()
	}
	return op, nil
}

// CompactSQL normalizes whitespace for substring assertions on SQL text.
func CompactSQL(query string) string {
	return strings.Join(strings.Fields(query), " ")
}

// Exec implements the pgx Exec surface.
func (s *ScriptedPgx) Exec(_ context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	op, err := s.next("exec", query, args)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	if op.Err != nil {
		return pgconn.CommandTag{}, op.Err
	}
	return pgconn.NewCommandTag(fmt.Sprintf("SCRIPT %d", op.Affected)), nil
}

// Query implements the pgx Query surface.
func (s *ScriptedPgx) Query(_ context.Context, query string, args ...any) (pgx.Rows, error) {
	op, err := s.next("query", query, args)
	if err != nil {
		return nil, err
	}
	if op.Err != nil {
		return nil, op.Err
	}
	errAt := op.RowsErrAt
	if op.RowsErr != nil && errAt == 0 && len(op.Rows) > 0 {
		// Default: the iteration error surfaces after all scripted rows.
		errAt = len(op.Rows)
	}
	return &scriptedRows{rows: op.Rows, iterErr: op.RowsErr, iterErrAt: errAt}, nil
}

// QueryRow implements the pgx QueryRow surface: operation errors and scan
// results are deferred to Row.Scan, and zero scripted rows yield
// pgx.ErrNoRows exactly like a live query.
func (s *ScriptedPgx) QueryRow(_ context.Context, query string, args ...any) pgx.Row {
	op, err := s.next("query", query, args)
	if err != nil {
		return scriptedRow{err: err}
	}
	if op.Err != nil {
		return scriptedRow{err: op.Err}
	}
	if op.RowsErr != nil {
		return scriptedRow{err: op.RowsErr}
	}
	if len(op.Rows) == 0 {
		return scriptedRow{err: pgx.ErrNoRows}
	}
	return scriptedRow{values: op.Rows[0]}
}

type scriptedRow struct {
	values []any
	err    error
}

func (r scriptedRow) Scan(dest ...any) error {
	if r.err != nil {
		return r.err
	}
	return scanValues(r.values, dest)
}

// scriptedRows implements pgx.Rows over scripted values.
type scriptedRows struct {
	rows      [][]any
	index     int
	current   []any
	iterErr   error
	iterErrAt int
	err       error
	closed    bool
}

func (r *scriptedRows) Next() bool {
	if r.closed || r.err != nil {
		return false
	}
	if r.iterErr != nil && r.index == r.iterErrAt {
		r.err = r.iterErr
		return false
	}
	if r.index >= len(r.rows) {
		return false
	}
	r.current = r.rows[r.index]
	r.index++
	return true
}

func (r *scriptedRows) Scan(dest ...any) error {
	if r.err != nil {
		return r.err
	}
	if err := scanValues(r.current, dest); err != nil {
		r.err = err
		return err
	}
	return nil
}

func (r *scriptedRows) Err() error { return r.err }
func (r *scriptedRows) Close()     { r.closed = true }
func (r *scriptedRows) Conn() *pgx.Conn {
	return nil
}
func (r *scriptedRows) CommandTag() pgconn.CommandTag                { return pgconn.CommandTag{} }
func (r *scriptedRows) FieldDescriptions() []pgconn.FieldDescription { return nil }
func (r *scriptedRows) RawValues() [][]byte                          { return nil }

func (r *scriptedRows) Values() ([]any, error) {
	return append([]any(nil), r.current...), nil
}

func scanValues(values []any, dest []any) error {
	if len(values) != len(dest) {
		return fmt.Errorf("scan: %d destination arguments for %d values", len(dest), len(values))
	}
	for i, value := range values {
		if err := assignScanValue(dest[i], value); err != nil {
			return fmt.Errorf("scan column %d: %w", i, err)
		}
	}
	return nil
}

// assignScanValue mirrors the narrow set of conversions the converted ops
// code relies on (plain values, sql.Scanner destinations); anything else is
// a scan error, like a live driver's type mismatch.
func assignScanValue(dst, src any) error {
	if scanner, ok := dst.(sql.Scanner); ok {
		return scanner.Scan(src)
	}
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}
	switch d := dst.(type) {
	case *int64:
		v, ok := src.(int64)
		if !ok {
			return scanMismatch(dst, src)
		}
		*d = v
	case *int:
		v, ok := src.(int64)
		if !ok {
			return scanMismatch(dst, src)
		}
		*d = int(v)
	case *string:
		v, ok := src.(string)
		if !ok {
			return scanMismatch(dst, src)
		}
		*d = v
	case *[]byte:
		v, ok := src.([]byte)
		if !ok {
			return scanMismatch(dst, src)
		}
		*d = append([]byte(nil), v...)
	case *bool:
		v, ok := src.(bool)
		if !ok {
			return scanMismatch(dst, src)
		}
		*d = v
	case *time.Time:
		v, ok := src.(time.Time)
		if !ok {
			return scanMismatch(dst, src)
		}
		*d = v
	default:
		return fmt.Errorf("unsupported scan destination %T", dst)
	}
	return nil
}

func scanMismatch(dst, src any) error {
	return fmt.Errorf("cannot scan %T into %T", src, dst)
}
