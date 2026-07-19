// Package testfixtures owns the shared seed dataset used by every
// integration suite that needs a populated database (API parity, store read
// suite, statistics benchmarks). The SQL files are embedded so suites in any
// package can apply them without depending on the process working directory.
//
// The dataset models 3 complete CosmicGame rounds plus 1 open round, 5
// bidders, every prize type, donations of every kind, staking activity, and
// the RandomWalk marketplace with Elo ranking state. Aggregate tables
// (cg_glob_stats, cg_bidder, rw_stats, ...) are not seeded directly: the
// inserts fire the production plpgsql triggers, which compute them.
//
// Suites that need extra rows (e.g. notification state only the store suite
// reads) layer suite-local extension seeds on top via ApplyFS; the shared
// dataset itself must not change without regenerating every dependent golden.
package testfixtures

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"slices"
)

//go:embed seed/*.sql
var seedFS embed.FS

// Apply loads the shared fixture dataset, executing every embedded seed file
// in lexical order inside the given database.
func Apply(ctx context.Context, db *sql.DB) error {
	return ApplyFS(ctx, db, seedFS, "seed")
}

// ApplyFS executes every *.sql file directly under dir of fsys in lexical
// order. Integration suites use it to apply their own extension seeds with
// the same semantics as the shared dataset.
func ApplyFS(ctx context.Context, db *sql.DB, fsys fs.FS, dir string) error {
	entries, err := fs.Glob(fsys, dir+"/*.sql")
	if err != nil {
		return fmt.Errorf("globbing seed files under %s: %w", dir, err)
	}
	if len(entries) == 0 {
		return fmt.Errorf("no seed files found under %s", dir)
	}
	slices.Sort(entries)
	for _, path := range entries {
		contents, err := fs.ReadFile(fsys, path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}
		if _, err := db.ExecContext(ctx, string(contents)); err != nil {
			return fmt.Errorf("applying %s: %w", path, err)
		}
	}
	return nil
}
