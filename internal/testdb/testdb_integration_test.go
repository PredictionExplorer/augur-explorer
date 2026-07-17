//go:build integration

package testdb

import (
	"context"
	"os"
	"testing"

	"github.com/pressly/goose/v3"
)

// TestMigrationsApply verifies that a fresh container comes up and every
// goose migration applies cleanly — this guards the whole schema.
func TestMigrationsApply(t *testing.T) {
	db := New(t)

	// Spot-check one table per schema group.
	for _, table := range []string{"block", "cg_bid", "rw_mint_evt"} {
		var exists bool
		err := db.SQL.QueryRow(
			`SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)`, table,
		).Scan(&exists)
		if err != nil {
			t.Fatalf("checking table %s: %v", table, err)
		}
		if !exists {
			t.Errorf("expected table %q to exist after migrations", table)
		}
	}
}

func TestBidModerationMigrationDeduplicatesAndEnforcesUniqueness(t *testing.T) {
	db := New(t)
	ctx := context.Background()
	provider, err := goose.NewProvider(goose.DialectPostgres, db.SQL, os.DirFS(migrationsDir()))
	if err != nil {
		t.Fatalf("creating migration provider: %v", err)
	}
	if _, err := provider.DownTo(ctx, 24); err != nil {
		t.Fatalf("rolling back bid-moderation migration: %v", err)
	}

	if _, err := db.SQL.ExecContext(ctx, `INSERT INTO cg_banned_bids(bid_id,user_addr,created_at) VALUES
		(42,'older',100),
		(42,'newer',200)`); err != nil {
		t.Fatalf("seeding duplicate bid bans: %v", err)
	}
	if _, err := provider.UpTo(ctx, 25); err != nil {
		t.Fatalf("reapplying bid-moderation migration: %v", err)
	}

	var count int
	var address string
	if err := db.SQL.QueryRowContext(ctx, `SELECT COUNT(*), MAX(user_addr)
		FROM cg_banned_bids WHERE bid_id=42`).Scan(&count, &address); err != nil {
		t.Fatalf("reading deduplicated bid ban: %v", err)
	}
	if count != 1 || address != "newer" {
		t.Fatalf("deduplicated rows = %d, address=%q; want newest row only", count, address)
	}
	if _, err := db.SQL.ExecContext(ctx,
		`INSERT INTO cg_banned_bids(bid_id,user_addr,created_at) VALUES (42,'duplicate',300)`); err == nil {
		t.Fatal("unique bid-ban index accepted a duplicate bid id")
	}
}
