// The RandomWalk ETL: indexes RandomWalk NFT and marketplace contract events
// into PostgreSQL. The shared process skeleton (typed configuration, process
// logger, RPC client, store, metrics, the indexing engine) lives in
// internal/etlapp; this binary contributes only the RandomWalk module
// wiring: contract bootstrap, the handler set of internal/indexer/randomwalk
// and the rw_proc_status watermark adapter.
package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/PredictionExplorer/augur-explorer/internal/etlapp"
	rwindexer "github.com/PredictionExplorer/augur-explorer/internal/indexer/randomwalk"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

// rwProgress adapts the rw_proc_status row to the engine's watermark
// interface, preserving last_evt_id across writes.
type rwProgress struct {
	repo *rwstore.Repo
}

func (p rwProgress) LastBlock(ctx context.Context) (int64, error) {
	status, err := p.repo.ProcessingStatus(ctx)
	if err != nil {
		return 0, err
	}
	return status.LastBlockNum, nil
}

func (p rwProgress) SetLastBlock(ctx context.Context, block int64) error {
	status, err := p.repo.ProcessingStatus(ctx)
	if err != nil {
		return err
	}
	status.LastBlockNum = block
	return p.repo.UpdateProcessingStatus(ctx, &status)
}

// osExit is stubbed by tests that drive main through its failure arm.
var osExit = os.Exit

func main() {
	if version.HandleFlag(os.Args[1:], os.Stdout) {
		return
	}
	if err := runMain(); err != nil {
		fmt.Fprintf(os.Stderr, "rw-etl: %v\n", err)
		osExit(1)
	}
}

// runMain owns the signal-scoped context so its deferred cleanup always runs
// before main decides the exit code (os.Exit skips deferred calls).
func runMain() error {
	// Graceful shutdown: on SIGINT/SIGTERM/SIGHUP finish the current event
	// batch, write status, and exit 0 cleanly. The engine checks ctx between
	// batches and during waits.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()
	return run(ctx, os.Getenv, os.Stdout, prometheus.DefaultRegisterer, prometheus.DefaultGatherer)
}

// run drives the shared ETL skeleton with the RandomWalk module wiring; the
// signature is the integration tests' seam.
func run(ctx context.Context, getenv func(string) string, logOut io.Writer, reg prometheus.Registerer, gatherer prometheus.Gatherer) error {
	return etlapp.Run(ctx, getenv, logOut, reg, gatherer, wireRandomWalk)
}

// wireRandomWalk registers the contract addresses (fresh-database bootstrap)
// and builds the event-handler set over them.
func wireRandomWalk(ctx context.Context, deps etlapp.Deps) (etlapp.Module, error) {
	rwRepo := rwstore.NewRepo(deps.Store)

	contracts, _, err := rwindexer.BootstrapContracts(ctx, rwRepo, deps.Store)
	if err != nil {
		return etlapp.Module{}, fmt.Errorf("contract address bootstrap failed: %w", err)
	}
	deps.Logger.Info("Contract addresses registered in address table",
		"randomwalk", contracts.RandomWalk.String(), "marketplace", contracts.Market.String())

	handlers, err := rwindexer.New(rwindexer.Config{
		Repo:      rwRepo,
		Contracts: contracts,
		Logger:    deps.Logger,
	})
	if err != nil {
		return etlapp.Module{}, fmt.Errorf("can't build event handlers: %w", err)
	}

	return etlapp.Module{
		Progress:  rwProgress{repo: rwRepo},
		Registry:  handlers.Registry(),
		Contracts: contracts.All(),
	}, nil
}
