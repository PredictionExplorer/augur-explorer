// The CosmicGame ETL: indexes every CosmicGame-family contract event into
// PostgreSQL. main wires the process-wide dependencies (loggers, RPC client,
// store, the handler set of internal/indexer/cosmicgame), runs the startup
// contract-parameter sync and hands control to the shared indexing engine
// (internal/indexer).
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	cgindexer "github.com/PredictionExplorer/augur-explorer/internal/indexer/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const (
	defaultDBLog = "db.log"

	// defaultLogDir is the legacy operational log directory under $HOME.
	defaultLogDir = "ae_logs"
)

// openAppendLog opens (creating if needed) one of the ETL's append-only log
// files.
func openAppendLog(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666) // #nosec G302 G304 G703 -- operational log under $HOME/ae_logs, world-readable by design
	if err != nil {
		return nil, fmt.Errorf("can't open log file: %w", err)
	}
	return f, nil
}

// cgProgress adapts the cg_proc_status row to the engine's watermark
// interface, preserving last_evt_id across writes.
type cgProgress struct {
	repo *cgstore.Repo
}

func (p cgProgress) LastBlock(ctx context.Context) (int64, error) {
	status, err := p.repo.ProcessingStatus(ctx)
	if err != nil {
		return 0, err
	}
	return status.LastBlockNum, nil
}

func (p cgProgress) SetLastBlock(ctx context.Context, block int64) error {
	status, err := p.repo.ProcessingStatus(ctx)
	if err != nil {
		return err
	}
	status.LastBlockNum = block
	return p.repo.UpdateProcessingStatus(ctx, &status)
}

func main() {
	// Graceful shutdown: on SIGINT/SIGTERM/SIGHUP finish the current event
	// batch, write status, and exit 0 cleanly. The engine checks ctx between
	// batches and during waits.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	if err := run(ctx, os.Getenv, prometheus.DefaultRegisterer, prometheus.DefaultGatherer); err != nil {
		fmt.Fprintf(os.Stderr, "cg-etl: %v\n", err)
		os.Exit(1)
	}
}

// run wires every dependency and drives the indexing engine until ctx is
// cancelled (returning nil) or a fatal error occurs. The Prometheus
// registerer/gatherer pair is injected so tests can use isolated registries
// (the default registry rejects duplicate registration across runs).
func run(ctx context.Context, getenv func(string) string, reg prometheus.Registerer, gatherer prometheus.Gatherer) error {
	logDir := fmt.Sprintf("%v/%v", getenv("HOME"), defaultLogDir)
	_ = os.MkdirAll(logDir, os.ModePerm) // #nosec G301 G703 -- legacy log dir under $HOME; openAppendLog fails loudly if unusable

	infoFile, err := openAppendLog(fmt.Sprintf("%v/cosmicgame_info.log", logDir))
	if err != nil {
		return err
	}
	errFile, err := openAppendLog(fmt.Sprintf("%v/cosmicgame_error.log", logDir))
	if err != nil {
		return err
	}
	// The engine, the handlers and the startup sync log structured records
	// into the legacy two-file layout: everything to the info log, errors
	// duplicated to the error log.
	logger := slog.New(indexer.NewDualLogHandler(infoFile, errFile))

	go func() {
		<-ctx.Done()
		logger.Info("Got signal, will exit after current batch is processed." +
			" To interrupt abruptly send SIGKILL (9) to the kernel.")
	}()

	rpcURL := getenv("RPC_URL")
	rpcclient, err := rpc.DialContext(ctx, rpcURL)
	if err != nil {
		return fmt.Errorf("dialing RPC node: %w", err)
	}
	logger.Info("Connected to ETH node", "rpc_url", rpcURL)
	eclient := ethclient.NewClient(rpcclient)

	// Database log output (failed and slow queries) goes through the pgx
	// slog tracer into the file the legacy Init_log wrote to.
	dbLogHandle, err := openAppendLog(fmt.Sprintf("%v/cosmicgame_%v", logDir, defaultDBLog))
	if err != nil {
		return err
	}
	cfg := store.ConfigFromEnv()
	cfg.Logger = slog.New(slog.NewTextHandler(dbLogHandle, nil))
	dbStore, err := store.New(ctx, cfg)
	if err != nil {
		return fmt.Errorf("can't connect to PostgreSQL database: %w\n%s", err, store.ConnectHint(err))
	}
	defer dbStore.Close()
	cgRepo := cgstore.NewRepo(dbStore)

	// Register the contract addresses (fresh-database bootstrap) and build
	// the event-handler set over them.
	contracts, cgAddrs, err := cgindexer.BootstrapContracts(ctx, cgRepo, dbStore)
	if err != nil {
		return fmt.Errorf("contract address bootstrap failed: %w", err)
	}
	logger.Info("All contract addresses registered in address table")

	handlers, err := cgindexer.New(cgindexer.Config{
		Repo:      cgRepo,
		Store:     dbStore,
		Caller:    eclient,
		Contracts: contracts,
		Logger:    logger,
	})
	if err != nil {
		return fmt.Errorf("can't build event handlers: %w", err)
	}

	if err := cgindexer.SyncContractParams(ctx, cgRepo, dbStore, eclient, cgAddrs.CosmicGameAddr, cgAddrs.PrizesWalletAddr, logger); err != nil {
		logger.Error("Contract param chain sync failed", "err", err)
		return fmt.Errorf("contract param chain sync failed: %w", err)
	}

	// Private metrics/pprof listener, enabled by METRICS_ADDR (never expose
	// it publicly; use a different port per process on shared hosts).
	metrics := indexer.NewMetrics(reg)
	if addr := strings.TrimSpace(getenv("METRICS_ADDR")); addr != "" {
		srv, _, err := indexer.StartMetricsServer(ctx, addr, gatherer, logger)
		if err != nil {
			return fmt.Errorf("can't start metrics server on %v: %w", addr, err)
		}
		defer func() { _ = srv.Close() }()
	}

	registry := handlers.Registry()
	engine, err := indexer.New(indexer.Config{
		Store:     dbStore,
		Client:    eclient,
		Progress:  cgProgress{repo: cgRepo},
		Process:   indexer.LogProcessor(dbStore, registry),
		Contracts: contracts.All(),
		Logger:    logger,
		Metrics:   metrics,
		TopicName: registry.TopicName,
	})
	if err != nil {
		return fmt.Errorf("can't build indexer engine: %w", err)
	}

	if err := engine.Run(ctx); err != nil {
		logger.Error("Event processing loop terminated", "err", err)
		return fmt.Errorf("event processing loop terminated: %w", err)
	}
	return nil
}
