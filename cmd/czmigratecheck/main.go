// Command czmigratecheck applies every goose migration in db/migrations to
// the database named by CZCHECK_DSN. It validates new migrations against a
// scratch database without docker; with CZCHECK_INSERTS=1 it also exercises
// the ChaosZero insert parameter shapes (see insertcheck.go).
package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

func main() {
	if err := run(context.Background(), os.Getenv, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context, getenv func(string) string, out io.Writer) error {
	dsn := getenv("CZCHECK_DSN")
	if dsn == "" {
		return errors.New("CZCHECK_DSN is required")
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	if err := testdb.Migrate(ctx, db); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	fmt.Fprintln(out, "all migrations applied")
	if getenv("CZCHECK_INSERTS") != "1" {
		return nil
	}
	if err := insertCheck(ctx, dsn, out); err != nil {
		return fmt.Errorf("insert check: %w", err)
	}
	return nil
}
