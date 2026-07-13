package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/PredictionExplorer/augur-explorer/internal/config"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// serverDeps bundles the process-wide dependencies run assembles at boot:
// the shared store pool and the process logger.
type serverDeps struct {
	store  *store.Store
	logger *slog.Logger
}

// newServerDeps connects the shared store pool (query tracing routed through
// the process logger with a component attribute) and returns the bundle.
// Errors are fatal to the caller: the API server cannot run without its
// database.
func newServerDeps(ctx context.Context, cfg *config.APIServer, logger *slog.Logger) (*serverDeps, error) {
	storeCfg := cfg.DB.StoreConfig()
	storeCfg.Logger = logger.With("component", "db")
	st, err := store.New(ctx, storeCfg)
	if err != nil {
		return nil, fmt.Errorf("can't connect to PostgreSQL database: %w\n%s", err, store.ConnectHint(err))
	}
	return &serverDeps{store: st, logger: logger}, nil
}
