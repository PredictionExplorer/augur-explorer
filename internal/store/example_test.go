package store_test

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ExampleNew shows the production wiring: one pgx pool per process, shared
// by the base Store and the domain repositories. There is no Output
// directive because the example needs a reachable PostgreSQL server; it is
// compiled but not executed (the integration suites in
// internal/store/{cosmicgame,randomwalk} run this exact wiring against a
// testcontainers database).
func ExampleNew() {
	ctx := context.Background()

	// ConfigFromEnv reads DATABASE_URL (which wins) or the PGSQL_*
	// variables; services build the same Config through internal/config.
	cfg := store.ConfigFromEnv()
	cfg.Logger = slog.Default().With("component", "db")

	st, err := store.New(ctx, cfg)
	if err != nil {
		fmt.Println(store.ConnectHint(err))
		return
	}
	defer st.Close()

	// Domain queries live on per-domain repositories sharing the pool (D3).
	repo := cgstore.NewRepo(st)
	status, err := repo.ProcessingStatus(ctx)
	if err != nil {
		fmt.Println("read watermark:", err)
		return
	}
	fmt.Println("indexed up to block", status.LastBlockNum)
}
