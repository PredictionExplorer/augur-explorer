//go:build integration

package main

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

func getenvFrom(env map[string]string) func(string) string {
	return func(key string) string { return env[key] }
}

func TestRunRequiresDSN(t *testing.T) {
	var out bytes.Buffer
	err := run(context.Background(), getenvFrom(nil), &out)
	if err == nil || !strings.Contains(err.Error(), "CZCHECK_DSN is required") {
		t.Fatalf("expected missing-DSN error, got %v", err)
	}
}

func TestRunMigrateFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var out bytes.Buffer
	err := run(ctx, getenvFrom(map[string]string{
		"CZCHECK_DSN": "postgres://nobody:nothing@127.0.0.1:1/czcheck_missing?connect_timeout=1",
	}), &out)
	if err == nil || !strings.Contains(err.Error(), "migrate:") {
		t.Fatalf("expected migrate error, got %v", err)
	}
}

func TestRunMigratesAndChecksInserts(t *testing.T) {
	db := testdb.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Without CZCHECK_INSERTS the run stops after the (idempotent) migration.
	var out bytes.Buffer
	if err := run(ctx, getenvFrom(map[string]string{
		"CZCHECK_DSN": db.ConnString,
	}), &out); err != nil {
		t.Fatalf("run without inserts: %v", err)
	}
	if got := out.String(); !strings.Contains(got, "all migrations applied") ||
		strings.Contains(got, "insert shapes verified") {
		t.Fatalf("unexpected output without inserts:\n%s", got)
	}

	// With CZCHECK_INSERTS=1 the rolled-back fixture exercise runs too.
	out.Reset()
	if err := run(ctx, getenvFrom(map[string]string{
		"CZCHECK_DSN":     db.ConnString,
		"CZCHECK_INSERTS": "1",
	}), &out); err != nil {
		t.Fatalf("run with inserts: %v", err)
	}
	if got := out.String(); !strings.Contains(got, "all migrations applied") ||
		!strings.Contains(got, "insert shapes verified (uint256 max survives NUMERIC[])") {
		t.Fatalf("unexpected output with inserts:\n%s", got)
	}

	// The insert exercise rolls everything back: the database stays clean.
	var leftover int
	if err := db.Pool.QueryRow(ctx,
		"SELECT COUNT(*) FROM cz_round_funded").Scan(&leftover); err != nil {
		t.Fatalf("counting cz_round_funded: %v", err)
	}
	if leftover != 0 {
		t.Fatalf("insert check leaked %d cz_round_funded rows", leftover)
	}
}

func TestRunReportsInsertCheckFailure(t *testing.T) {
	db := testdb.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// Hide a table the insert exercise needs; the migration itself stays
	// green because goose only consults its version table.
	if _, err := db.Pool.Exec(ctx,
		"ALTER TABLE cz_round_funded RENAME TO cz_round_funded_hidden"); err != nil {
		t.Fatalf("hiding cz_round_funded: %v", err)
	}
	var out bytes.Buffer
	err := run(ctx, getenvFrom(map[string]string{
		"CZCHECK_DSN":     db.ConnString,
		"CZCHECK_INSERTS": "1",
	}), &out)
	if err == nil || !strings.Contains(err.Error(), "insert check:") {
		t.Fatalf("expected insert check error, got %v", err)
	}
}
