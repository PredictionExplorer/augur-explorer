package testutil

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
)

func TestScriptedPgxServesRowsAndRecordsViolations(t *testing.T) {
	t.Parallel()
	script := NewScriptedPgx(PgxOp{
		Kind:     "query",
		Contains: "FROM widgets",
		Parts:    []string{"ORDER BY id", "not-in-the-query"},
		ArgCount: new(2),
		Rows:     [][]any{{int64(1), "a"}, {int64(2), "b"}},
	})
	rows, err := script.Query(context.Background(), "SELECT id, name FROM widgets ORDER BY id", 1)
	if err != nil {
		t.Fatalf("Query() error = %v", err)
	}
	var got []string
	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			t.Fatalf("Scan() error = %v", err)
		}
		got = append(got, name)
		if values, err := rows.Values(); err != nil || len(values) != 2 {
			t.Fatalf("Values() = %v, %v", values, err)
		}
	}
	if rows.Err() != nil || len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("rows = %v, err = %v", got, rows.Err())
	}
	rows.Close()
	if rows.Next() {
		t.Fatal("Next() after Close() served a row")
	}

	// The Parts and ArgCount mismatches above are recorded for AssertDone
	// (same-package access keeps this test from having to fail itself).
	script.mu.Lock()
	violations := len(script.violations)
	script.mu.Unlock()
	if violations != 2 {
		t.Fatalf("recorded violations = %d, want the Parts and ArgCount mismatches", violations)
	}

	// The pgx.Rows surface the production code never touches stays inert.
	if rows.CommandTag().String() != "" || rows.FieldDescriptions() != nil ||
		rows.RawValues() != nil || rows.Conn() != nil {
		t.Fatal("stub pgx.Rows surface returned data")
	}
}

func TestScriptedPgxOperationMatching(t *testing.T) {
	t.Parallel()
	queryError := func(t *testing.T, script *ScriptedPgx, query, want string) {
		t.Helper()
		rows, err := script.Query(context.Background(), query)
		if err == nil {
			rows.Close()
			t.Fatalf("query %q unexpectedly succeeded", query)
		}
		if !strings.Contains(err.Error(), want) {
			t.Fatalf("error = %v, want containing %q", err, want)
		}
	}
	t.Run("unexpected operation", func(t *testing.T) {
		t.Parallel()
		queryError(t, NewScriptedPgx(), "SELECT 1", "unexpected query")
	})
	t.Run("kind mismatch", func(t *testing.T) {
		t.Parallel()
		queryError(t, NewScriptedPgx(PgxOp{Kind: "exec", Contains: "UPDATE"}), "SELECT 1", "want exec")
	})
	t.Run("contains mismatch", func(t *testing.T) {
		t.Parallel()
		queryError(
			t,
			NewScriptedPgx(PgxOp{Kind: "query", Contains: "FROM other_table"}),
			"SELECT 1 FROM widgets",
			"does not contain",
		)
	})
	t.Run("on call hook", func(t *testing.T) {
		t.Parallel()
		called := false
		script := NewScriptedPgx(PgxOp{Kind: "exec", OnCall: func() { called = true }, Affected: 3})
		tag, err := script.Exec(context.Background(), "DELETE FROM widgets")
		if err != nil || !called {
			t.Fatalf("exec called=%v error=%v", called, err)
		}
		if tag.RowsAffected() != 3 {
			t.Fatalf("rows affected = %d", tag.RowsAffected())
		}
	})
	t.Run("exec error and unexpected exec", func(t *testing.T) {
		t.Parallel()
		sentinel := errors.New("exec failed")
		script := NewScriptedPgx(PgxOp{Kind: "exec", Err: sentinel})
		if _, err := script.Exec(context.Background(), "DELETE"); !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
		if _, err := script.Exec(context.Background(), "DELETE"); err == nil ||
			!strings.Contains(err.Error(), "unexpected exec") {
			t.Fatalf("error = %v", err)
		}
	})
}

func TestScriptedPgxQueryRowSemantics(t *testing.T) {
	t.Parallel()
	sentinel := errors.New("query failed")
	t.Run("no rows yields pgx.ErrNoRows", func(t *testing.T) {
		t.Parallel()
		script := NewScriptedPgx(PgxOp{Kind: "query"})
		var one int
		if err := script.QueryRow(context.Background(), "SELECT 1").Scan(&one); !errors.Is(err, pgx.ErrNoRows) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("operation error defers to Scan", func(t *testing.T) {
		t.Parallel()
		script := NewScriptedPgx(PgxOp{Kind: "query", Err: sentinel})
		var one int
		if err := script.QueryRow(context.Background(), "SELECT 1").Scan(&one); !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("iteration error defers to Scan", func(t *testing.T) {
		t.Parallel()
		script := NewScriptedPgx(PgxOp{Kind: "query", RowsErr: sentinel})
		var one int
		if err := script.QueryRow(context.Background(), "SELECT 1").Scan(&one); !errors.Is(err, sentinel) {
			t.Fatalf("error = %v", err)
		}
	})
	t.Run("script mismatch defers to Scan", func(t *testing.T) {
		t.Parallel()
		script := NewScriptedPgx()
		var one int
		if err := script.QueryRow(context.Background(), "SELECT 1").Scan(&one); err == nil {
			t.Fatal("unscripted QueryRow did not error")
		}
	})
	t.Run("first row wins", func(t *testing.T) {
		t.Parallel()
		script := NewScriptedPgx(PgxOp{Kind: "query", Rows: [][]any{{int64(7)}, {int64(9)}}})
		var got int64
		if err := script.QueryRow(context.Background(), "SELECT 1").Scan(&got); err != nil || got != 7 {
			t.Fatalf("value/error = %d/%v", got, err)
		}
	})
}

func TestScriptedPgxIterationErrors(t *testing.T) {
	t.Parallel()
	sentinel := errors.New("iteration failed")
	t.Run("after all rows by default", func(t *testing.T) {
		t.Parallel()
		script := NewScriptedPgx(PgxOp{Kind: "query", Rows: [][]any{{int64(1)}}, RowsErr: sentinel})
		rows, err := script.Query(context.Background(), "SELECT 1")
		if err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("first row was not served")
		}
		if rows.Next() || !errors.Is(rows.Err(), sentinel) {
			t.Fatalf("iteration error = %v", rows.Err())
		}
		var one int64
		if err := rows.Scan(&one); !errors.Is(err, sentinel) {
			t.Fatalf("Scan after failed iteration = %v", err)
		}
	})
	t.Run("at a scripted index", func(t *testing.T) {
		t.Parallel()
		script := NewScriptedPgx(PgxOp{
			Kind:      "query",
			Rows:      [][]any{{int64(1)}, {int64(2)}},
			RowsErr:   sentinel,
			RowsErrAt: 1,
		})
		rows, err := script.Query(context.Background(), "SELECT 1")
		if err != nil {
			t.Fatal(err)
		}
		if !rows.Next() {
			t.Fatal("first row was not served")
		}
		if rows.Next() || !errors.Is(rows.Err(), sentinel) {
			t.Fatalf("iteration error = %v", rows.Err())
		}
	})
}

func TestScriptedPgxScanConversions(t *testing.T) {
	t.Parallel()
	when := time.Unix(1_700_000_000, 0)
	script := NewScriptedPgx(PgxOp{Kind: "query", Rows: [][]any{{
		int64(7), int64(8), "text", []byte{0x01}, true, when, "null-string", nil,
	}}})
	rows, err := script.Query(context.Background(), "SELECT everything")
	if err != nil || !rows.Next() {
		t.Fatalf("rows setup failed: %v", err)
	}
	var (
		i64      int64
		i        int
		s        string
		b        []byte
		flag     bool
		ts       time.Time
		scanner  sql.NullString
		nullable sql.NullInt64
	)
	if err := rows.Scan(&i64, &i, &s, &b, &flag, &ts, &scanner, &nullable); err != nil {
		t.Fatalf("Scan() error = %v", err)
	}
	if i64 != 7 || i != 8 || s != "text" || len(b) != 1 || !flag || !ts.Equal(when) ||
		scanner.String != "null-string" || nullable.Valid {
		t.Fatalf("scanned = %v %v %v %v %v %v %#v %#v", i64, i, s, b, flag, ts, scanner, nullable)
	}

	for name, run := range map[string]func() error{
		"arity mismatch": func() error {
			return scanValues([]any{int64(1)}, []any{})
		},
		"null into plain destination": func() error {
			var dst string
			return scanValues([]any{nil}, []any{&dst})
		},
		"type mismatch": func() error {
			var dst int64
			return scanValues([]any{"nope"}, []any{&dst})
		},
		"int mismatch": func() error {
			var dst int
			return scanValues([]any{"nope"}, []any{&dst})
		},
		"string mismatch": func() error {
			var dst string
			return scanValues([]any{int64(1)}, []any{&dst})
		},
		"bytes mismatch": func() error {
			var dst []byte
			return scanValues([]any{"nope"}, []any{&dst})
		},
		"bool mismatch": func() error {
			var dst bool
			return scanValues([]any{"nope"}, []any{&dst})
		},
		"time mismatch": func() error {
			var dst time.Time
			return scanValues([]any{"nope"}, []any{&dst})
		},
		"unsupported destination": func() error {
			var dst float64
			return scanValues([]any{"nope"}, []any{&dst})
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if err := run(); err == nil {
				t.Fatal("conversion unexpectedly succeeded")
			}
		})
	}

	// A scan failure poisons the rows like pgx: Err reports it and further
	// scans keep failing.
	failing := NewScriptedPgx(PgxOp{Kind: "query", Rows: [][]any{{"nope"}}})
	badRows, err := failing.Query(context.Background(), "SELECT 1")
	if err != nil || !badRows.Next() {
		t.Fatalf("rows setup failed: %v", err)
	}
	var dst int64
	if err := badRows.Scan(&dst); err == nil || badRows.Err() == nil {
		t.Fatalf("scan failure not recorded: %v / %v", err, badRows.Err())
	}
}

func TestScriptedPgxPoolSurface(t *testing.T) {
	t.Parallel()
	script := NewScriptedPgx()
	if err := script.Ping(context.Background()); err != nil {
		t.Fatalf("Ping() error = %v", err)
	}
	if script.Closed() {
		t.Fatal("handle reported closed before Close")
	}
	script.Close()
	if !script.Closed() {
		t.Fatal("Close() was not recorded")
	}
	if err := script.Ping(context.Background()); err == nil {
		t.Fatal("Ping() after Close() succeeded")
	}
	script.AssertDone(t)
}
