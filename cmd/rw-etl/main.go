// The RandomWalk ETL: indexes RandomWalk NFT and marketplace contract events
// into PostgreSQL. main wires the process-wide dependencies (loggers, RPC
// client, store, the handler set of internal/indexer/randomwalk) and hands
// control to the shared indexing engine (internal/indexer).
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	rwindexer "github.com/PredictionExplorer/augur-explorer/internal/indexer/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/primitives"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

const defaultDBLog = "db.log"

// openAppendLog opens (creating if needed) one of the ETL's append-only log
// files.
func openAppendLog(path string) *os.File {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // #nosec G302 G304 G703 -- operational log under $HOME/ae_logs, world-readable by design
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't start: %v\n", err)
		os.Exit(1)
	}
	return f
}

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

func main() {
	logDir := fmt.Sprintf("%v/%v", os.Getenv("HOME"), primitives.DEFAULT_LOG_DIR)
	_ = os.MkdirAll(logDir, os.ModePerm) // #nosec G301 G703 -- legacy log dir under $HOME; openAppendLog fails loudly if unusable

	infoFile := openAppendLog(fmt.Sprintf("%v/randomwalk_info.log", logDir))
	errFile := openAppendLog(fmt.Sprintf("%v/randomwalk_error.log", logDir))
	// The engine and the handlers log structured records into the legacy
	// two-file layout: everything to the info log, errors duplicated to the
	// error log.
	logger := slog.New(indexer.NewDualLogHandler(infoFile, errFile))

	// Graceful shutdown: on SIGINT/SIGTERM/SIGHUP finish the current event
	// batch, write status, and exit 0 cleanly. The engine checks ctx between
	// batches and during waits.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()
	go func() {
		<-ctx.Done()
		logger.Info("Got signal, will exit after current batch is processed." +
			" To interrupt abruptly send SIGKILL (9) to the kernel.")
	}()

	rpcURL := os.Getenv("RPC_URL")
	rpcclient, err := rpc.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Connected to ETH node", "rpc_url", rpcURL)
	eclient := ethclient.NewClient(rpcclient)

	// Database log output (failed and slow queries) goes through the pgx
	// slog tracer into the file the legacy Init_log wrote to.
	dbLogHandle := openAppendLog(fmt.Sprintf("%v/randomwalk_%v", logDir, defaultDBLog))
	cfg := store.ConfigFromEnv()
	cfg.Logger = slog.New(slog.NewTextHandler(dbLogHandle, nil))
	dbStore, err := store.New(ctx, cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't connect to PostgreSQL database.\nConnection error: %v\n%s", err, store.ConnectHint(err))
		os.Exit(1)
	}
	rwRepo := rwstore.NewRepo(dbStore)

	// Register the contract addresses (fresh-database bootstrap) and build
	// the event-handler set over them.
	contracts, _, err := rwindexer.BootstrapContracts(ctx, rwRepo, dbStore)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Contract address bootstrap failed: %v\n", err)
		os.Exit(1)
	}
	logger.Info("Contract addresses registered in address table",
		"randomwalk", contracts.RandomWalk.String(), "marketplace", contracts.Market.String())

	handlers, err := rwindexer.New(rwindexer.Config{
		Repo:      rwRepo,
		Contracts: contracts,
		Logger:    logger,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build event handlers: %v\n", err)
		os.Exit(1)
	}

	// Private metrics/pprof listener, enabled by METRICS_ADDR (never expose
	// it publicly; use a different port per process on shared hosts).
	metrics := indexer.NewMetrics(prometheus.DefaultRegisterer)
	if addr := strings.TrimSpace(os.Getenv("METRICS_ADDR")); addr != "" {
		srv, _, err := indexer.StartMetricsServer(addr, prometheus.DefaultGatherer, logger)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't start metrics server on %v: %v\n", addr, err)
			os.Exit(1)
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
		fmt.Fprintf(os.Stderr, "Can't build indexer engine: %v\n", err)
		os.Exit(1)
	}

	if err := engine.Run(ctx); err != nil {
		logger.Error("Event processing loop terminated", "err", err)
		os.Exit(1)
	}
}
