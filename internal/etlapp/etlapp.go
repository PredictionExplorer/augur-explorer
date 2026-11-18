// Package etlapp is the shared process skeleton of the ETL binaries
// (cmd/cg-etl, cmd/rw-etl, cmd/cz-etl). Every ETL wires the same
// process-wide dependencies in the same order — typed configuration,
// structured logger, RPC client, store with server-side time bounds, pool
// metrics, the module's handler set, the optional private metrics listener
// and the shared indexing engine — and differs only in the module wiring
// step. Run owns the shared order; the per-binary WireFunc owns the module:
// contract bootstrap, handler registry, watermark adapter and any
// module-specific startup passes (cmd/cg-etl's parameter-drift audit and
// recovery hooks).
package etlapp

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/PredictionExplorer/augur-explorer/internal/config"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/version"
)

// Server-side database time bounds (D22 defense in depth). ETL statements
// are small single-row inserts and short reads; one that runs a minute is a
// fault, and PostgreSQL aborts it even if the client-side cancellation never
// arrives. The idle-in-transaction bound caps how long a per-block ingestion
// transaction (ADR-0010) could pin locks if the process stalled between
// statements.
const (
	dbStatementTimeout = time.Minute
	dbIdleInTxTimeout  = 5 * time.Minute
)

// Deps carries the process-wide dependencies Run has built by the time it
// calls the module's WireFunc.
type Deps struct {
	// Config is the loaded typed ETL configuration.
	Config *config.ETL
	// Logger is the process logger ("component=db" derivatives included).
	Logger *slog.Logger
	// Store is the open PostgreSQL store with the D22 session bounds set.
	Store *store.Store
	// Client is the dialed RPC client.
	Client *ethclient.Client
}

// Module is what a WireFunc returns: the pieces the shared engine needs,
// plus an optional post-engine hook.
type Module struct {
	// Progress persists the module's processing watermark.
	Progress indexer.Progress
	// Registry dispatches the module's event handlers.
	Registry *indexer.Registry
	// Contracts is the FilterLogs address set.
	Contracts []ethcommon.Address
	// PostEngine, when non-nil, runs after the engine is built and before
	// it starts: cmd/cg-etl uses it for the implementation/prizes-wallet
	// recovery passes and their live hooks, which need the engine and the
	// process function.
	PostEngine func(ctx context.Context, engine *indexer.Engine, process indexer.ProcessFunc) error
}

// WireFunc builds one module over the shared dependencies. Errors are
// returned verbatim, so wire implementations wrap them with the
// module-specific context ("contract address bootstrap failed: …").
type WireFunc func(ctx context.Context, deps Deps) (Module, error)

// Run wires every shared dependency and drives the indexing engine until
// ctx is cancelled (returning nil) or a fatal error occurs. Environment
// access goes through getenv and structured logs go to logOut; the
// Prometheus registerer/gatherer pair is injected so tests can use isolated
// registries (the default registry rejects duplicate registration across
// runs).
func Run(
	ctx context.Context,
	getenv func(string) string,
	logOut io.Writer,
	reg prometheus.Registerer,
	gatherer prometheus.Gatherer,
	wire WireFunc,
) error {
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

	module, err := wire(ctx, Deps{Config: cfg, Logger: logger, Store: dbStore, Client: eclient})
	if err != nil {
		return err
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

	process := indexer.LogProcessor(dbStore, module.Registry)
	engine, err := indexer.New(indexer.Config{
		Store:     dbStore,
		Client:    eclient,
		Progress:  module.Progress,
		Process:   process,
		Contracts: module.Contracts,
		Logger:    logger,
		Metrics:   metrics,
		TopicName: module.Registry.TopicName,
	})
	// The error branch below is unreachable through Run(): indexer.New only
	// rejects a nil Store/Client, both built above. Kept as a guard against
	// future constructor validation.
	if err != nil {
		return fmt.Errorf("can't build indexer engine: %w", err)
	}

	if module.PostEngine != nil {
		if err := module.PostEngine(ctx, engine, process); err != nil {
			return err
		}
	}

	if err := engine.Run(ctx); err != nil {
		logger.Error("Event processing loop terminated", "err", err)
		return fmt.Errorf("event processing loop terminated: %w", err)
	}
	return nil
}
