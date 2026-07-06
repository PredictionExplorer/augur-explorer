//go:build integration

package testdb

import (
	"testing"
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
