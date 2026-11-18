// The ChaosZero ETL: indexes the IndependentOutcomeMarket and
// OutcomeSettlementVault contract events into PostgreSQL. The shared process
// skeleton (typed configuration, process logger, RPC client, store, metrics,
// the indexing engine) lives in internal/etlapp; this binary contributes
// only the ChaosZero module wiring: contract bootstrap, the handler set of
// internal/indexer/chaoszero and the cz_proc_status watermark adapter.
//
// The ChaosZero contracts are immutable singletons (no proxy, no owner), so
// unlike cmd/cg-etl there is no contract-parameter drift audit and no
// implementation discovery: this binary is the simpler cmd/rw-etl shape.
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
	czindexer "github.com/PredictionExplorer/augur-explorer/internal/indexer/chaoszero"
	czstore "github.com/PredictionExplorer/augur-explorer/internal/store/chaoszero"
	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

// czProgress adapts the cz_proc_status row to the engine's watermark
// interface, preserving last_evt_id across writes.
type czProgress struct {
	repo *czstore.Repo
}

func (p czProgress) LastBlock(ctx context.Context) (int64, error) {
	status, err := p.repo.ProcessingStatus(ctx)
	if err != nil {
		return 0, err
	}
	return status.LastBlockNum, nil
}

func (p czProgress) SetLastBlock(ctx context.Context, block int64) error {
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
		fmt.Fprintf(os.Stderr, "cz-etl: %v\n", err)
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

// run drives the shared ETL skeleton with the ChaosZero module wiring; the
// signature is the integration tests' seam.
func run(ctx context.Context, getenv func(string) string, logOut io.Writer, reg prometheus.Registerer, gatherer prometheus.Gatherer) error {
	return etlapp.Run(ctx, getenv, logOut, reg, gatherer, wireChaosZero)
}

// wireChaosZero registers the contract addresses (fresh-database bootstrap)
// and builds the event-handler set over them.
func wireChaosZero(ctx context.Context, deps etlapp.Deps) (etlapp.Module, error) {
	czRepo := czstore.NewRepo(deps.Store)

	contracts, addrs, err := czindexer.BootstrapContracts(ctx, czRepo, deps.Store)
	if err != nil {
		return etlapp.Module{}, fmt.Errorf("contract address bootstrap failed: %w", err)
	}
	deps.Logger.Info("Contract addresses registered in address table",
		"market", contracts.Market.String(), "vault", contracts.Vault.String(), "source", addrs.Source)

	handlers, err := czindexer.New(czindexer.Config{
		Repo:      czRepo,
		Contracts: contracts,
		Logger:    deps.Logger,
	})
	// The error branch below is unreachable through run(): czindexer.New only
	// rejects a nil Repo, and this wiring always supplies one. Kept as a
	// guard against future constructor validation.
	if err != nil {
		return etlapp.Module{}, fmt.Errorf("can't build event handlers: %w", err)
	}

	return etlapp.Module{
		Progress:  czProgress{repo: czRepo},
		Registry:  handlers.Registry(),
		Contracts: contracts.All(),
	}, nil
}
