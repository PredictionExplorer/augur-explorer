// The CosmicGame ETL: indexes every CosmicGame-family contract event into
// PostgreSQL. main wires the process-wide dependencies (typed configuration,
// process logger, RPC client, store, the handler set of
// internal/indexer/cosmicgame), audits contract-parameter drift, recovers
// event-less V3 champion durations and hands control to the shared indexing
// engine (internal/indexer).
package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/PredictionExplorer/augur-explorer/internal/config"
	"github.com/PredictionExplorer/augur-explorer/internal/ethcall"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	cgindexer "github.com/PredictionExplorer/augur-explorer/internal/indexer/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

// Server-side database time bounds (D22 defense in depth). ETL statements
// are small single-row inserts and short reads; one that runs a minute is a
// fault, and PostgreSQL aborts it even if the client-side cancellation never
// arrives. The idle-in-transaction bound caps how long a per-block ingestion
// transaction (ADR-0010) could pin locks if the process stalled between
// statements — for example on a chain RPC call that the engine's per-call
// deadline somehow failed to cut short.
const (
	dbStatementTimeout = time.Minute
	dbIdleInTxTimeout  = 5 * time.Minute
)

// cgProgress adapts the cg_proc_status row to the engine's watermark
// interface, preserving last_evt_id across writes.
// implementationCommitWait bounds how long the implementation backfill waits
// for the block carrying an Upgraded event to commit before giving up (the
// startup recovery pass catches anything left behind).
const implementationCommitWait = 10 * time.Minute

// backfillImplementationAfterCommit waits until the engine has committed
// blockNum (the Upgraded event's block; the hook fires inside that block's
// transaction) and then backfills the new implementation's constructor
// Initialized event. Failures are logged, not fatal: the startup recovery
// pass (RecoverImplementationEvents) repeats the attempt on the next run.
func backfillImplementationAfterCommit(
	ctx context.Context,
	logger *slog.Logger,
	progress cgProgress,
	handlers *cgindexer.Handlers,
	engine *indexer.Engine,
	process indexer.ProcessFunc,
	impl ethcommon.Address,
	blockNum int64,
) {
	deadline := time.Now().Add(implementationCommitWait)
	for {
		last, err := progress.LastBlock(ctx)
		if err == nil && last >= blockNum {
			break
		}
		if ctx.Err() != nil {
			return
		}
		if time.Now().After(deadline) {
			logger.Warn("implementation backfill skipped: Upgraded block not committed in time",
				"implementation", impl.Hex(), "block", blockNum)
			return
		}
		time.Sleep(time.Second)
	}
	if blockNum < 0 {
		return
	}
	if _, err := handlers.BackfillImplementationEvents(ctx, engine, process, impl, uint64(blockNum)); err != nil {
		logger.Error("implementation backfill failed; the startup recovery pass will retry",
			"implementation", impl.Hex(), "block", blockNum, "err", err)
	}
}

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

// osExit is stubbed by tests that drive main through its failure arm.
var osExit = os.Exit

func main() {
	if version.HandleFlag(os.Args[1:], os.Stdout) {
		return
	}
	if err := runMain(); err != nil {
		fmt.Fprintf(os.Stderr, "cg-etl: %v\n", err)
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
	storeCfg.StatementTimeout = dbStatementTimeout
	storeCfg.IdleInTxSessionTimeout = dbIdleInTxTimeout
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

	if _, err := cgindexer.CheckContractParamsDrift(
		ctx,
		cgRepo,
		eclient,
		cgAddrs.CosmicGameAddr,
		cgAddrs.PrizesWalletAddr,
		logger,
	); err != nil {
		logger.Error("Contract parameter drift audit failed", "err", err)
		return fmt.Errorf("contract parameter drift audit failed: %w", err)
	}

	recoveryCtx, cancelRecovery := context.WithTimeout(ctx, ethcall.DefaultTimeout)
	header, headerErr := eclient.HeaderByNumber(recoveryCtx, nil)
	cancelRecovery()
	if headerErr != nil {
		logger.Warn("Champion duration recovery skipped: latest header unavailable", "err", headerErr)
	} else if header.Time > math.MaxInt64 {
		return fmt.Errorf("champion duration recovery: header timestamp %d overflows int64", header.Time)
	} else if err := handlers.RecoverChampionDurations(
		ctx,
		header.Number.Int64(),
		int64(header.Time), // #nosec G115 -- bounded above
	); err != nil {
		return fmt.Errorf("champion duration recovery failed: %w", err)
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
	process := indexer.LogProcessor(dbStore, registry)
	progress := cgProgress{repo: cgRepo}
	engine, err := indexer.New(indexer.Config{
		Store:     dbStore,
		Client:    eclient,
		Progress:  progress,
		Process:   process,
		Contracts: contracts.All(),
		Logger:    logger,
		Metrics:   metrics,
		TopicName: registry.TopicName,
	})
	if err != nil {
		return fmt.Errorf("can't build indexer engine: %w", err)
	}

	// Game implementations: constructor events of implementations the ETL
	// was not watching when they were deployed are backfilled now, and a
	// new implementation reported by Upgraded joins the FilterLogs set at
	// once, with its constructor event backfilled once the block commits.
	if err := handlers.RecoverImplementationEvents(ctx, engine, process); err != nil {
		return fmt.Errorf("implementation event recovery failed: %w", err)
	}
	handlers.SetOnNewImplementation(func(_ context.Context, impl ethcommon.Address, blockNum int64) {
		engine.AddContracts(impl)
		go backfillImplementationAfterCommit(ctx, logger, progress, handlers, engine, process, impl, blockNum)
	})

	// A new prizes wallet (PrizesWalletAddressChanged) joins the FilterLogs
	// set at once. No backfill is needed: a prizes wallet only emits events
	// when the game calls into it, which cannot precede the announcement.
	handlers.SetOnNewPrizesWallet(func(_ context.Context, wallet ethcommon.Address, _ int64) {
		engine.AddContracts(wallet)
	})

	if err := engine.Run(ctx); err != nil {
		logger.Error("Event processing loop terminated", "err", err)
		return fmt.Errorf("event processing loop terminated: %w", err)
	}
	return nil
}
