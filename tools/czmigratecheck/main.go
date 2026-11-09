// One-off migration check: applies every goose migration in db/migrations to
// the database named by CZCHECK_DSN. Used to validate new migrations against
// a scratch database without docker.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/PredictionExplorer/augur-explorer/internal/testdb"
)

func main() {
	dsn := os.Getenv("CZCHECK_DSN")
	if dsn == "" {
		fmt.Fprintln(os.Stderr, "CZCHECK_DSN is required")
		os.Exit(1)
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "open:", err)
		os.Exit(1)
	}
	defer db.Close()
	if err := testdb.Migrate(context.Background(), db); err != nil {
		fmt.Fprintln(os.Stderr, "migrate:", err)
		os.Exit(1)
	}
	fmt.Println("all migrations applied")
	maybeInsertCheck(context.Background(), dsn)
}
