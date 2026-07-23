// The RandomWalk ETL: indexes RandomWalk NFT and marketplace contract events
// into PostgreSQL. main wires the process-wide dependencies (typed
// configuration, process logger, RPC client, store, the handler set of
// internal/indexer/randomwalk) and hands control to the shared indexing
// engine (internal/indexer).
package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/PredictionExplorer/augur-explorer/internal/config"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	rwindexer "github.com/PredictionExplorer/augur-explorer/internal/indexer/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
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

// run wires every dependency and drives the indexing engine until ctx is
// cancelled (returning nil) or a fatal error occurs. Environment access goes
// through getenv and structured logs go to logOut; the Prometheus
// registerer/gatherer pair is injected so tests can use isolated registries
// (the default registry rejects duplicate registration across runs).
func run(ctx context.Context, getenv func(string) string, logOut io.Writer, reg prometheus.Registerer, gatherer prometheus.Gatherer) error {
	cfg, err := config.LoadETL(getenv)
	if err != nil {
		return err
	}
	// One structured logger on stdout; journald owns persistence (§8.3 —
	// the legacy $HOME/ae_logs dual-file layout is gone).
	logger := cfg.Log.NewLogger(logOut)
	logger.LogAttrs(ctx, slog.LevelInfo, "build info", version.LogAttrs()...)
	logger.LogAttrs(ctx, slog.LevelInfo, "effective configuration", config.Attrs(cfg)...)

	go func() {
		<-ctx.Done()
		logger.Info("Got signal, will exit after current batch is processed." +
			" To interrupt abruptly send SIGKILL (9) to the kernel.")
	}()

	rpcclient, err := rpc.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return fmt.Errorf("dialing RPC node: %w", err)
	}
	logger.Info("Connected to ETH node", "rpc_url", config.RedactURL(cfg.RPCURL))
	eclient := ethclient.NewClient(rpcclient)

	// Database log output (failed and slow queries) goes through the pgx
	// slog tracer onto the same stream, tagged component=db.
	storeCfg := cfg.DB.StoreConfig()
	storeCfg.Logger = logger.With("component", "db")
	dbStore, err := store.New(ctx, storeCfg)
	if err != nil {
		return fmt.Errorf("can't connect to PostgreSQL database: %w\n%s", err, store.ConnectHint(err))
	}
	defer dbStore.Close()
	poolCollector := store.NewPoolCollector(dbStore.Pool())
	if err := reg.Register(poolCollector); err != nil {
		return fmt.Errorf("registering db pool metrics: %w", err)
	}
	defer reg.Unregister(poolCollector)
	rwRepo := rwstore.NewRepo(dbStore)

	// Register the contract addresses (fresh-database bootstrap) and build
	// the event-handler set over them.
	contracts, _, err := rwindexer.BootstrapContracts(ctx, rwRepo, dbStore)
	if err != nil {
		return fmt.Errorf("contract address bootstrap failed: %w", err)
	}
	logger.Info("Contract addresses registered in address table",
		"randomwalk", contracts.RandomWalk.String(), "marketplace", contracts.Market.String())

	handlers, err := rwindexer.New(rwindexer.Config{
		Repo:      rwRepo,
		Contracts: contracts,
		Logger:    logger,
	})
	if err != nil {
		return fmt.Errorf("can't build event handlers: %w", err)
	}

	// Private metrics/pprof listener, enabled by METRICS_ADDR (never expose
	// it publicly; use a different port per process on shared hosts).
	metrics := indexer.NewMetrics(reg)
	if addr := strings.TrimSpace(cfg.MetricsAddr); addr != "" {
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
		Progress:  rwProgress{repo: rwRepo},
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
