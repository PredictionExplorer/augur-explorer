// The RandomWalk ETL: indexes RandomWalk NFT and marketplace contract events
// into PostgreSQL. main wires the process-wide dependencies (loggers, RPC
// client, store, ABIs, contract addresses) and hands control to the shared
// indexing engine (internal/indexer).
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

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"

	. "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwdb "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

const DEFAULT_DB_LOG = "db.log"

var (
	// dbStore owns the process-wide connection pool; rwRepo runs every
	// RandomWalk query on it.
	dbStore *store.Store
	rwRepo  *rwdb.Repo
	RPC_URL = os.Getenv("RPC_URL")
	Error   *log.Logger
	Info    *log.Logger

	eclient *ethclient.Client

	marketplace_abi *abi.ABI
	randomwalk_abi  *abi.ABI

	rw_contracts rwp.ContractAddresses
	market_addr  ethcommon.Address
	rwalk_addr   ethcommon.Address
)

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
	repo *rwdb.Repo
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

	log_dir := fmt.Sprintf("%v/%v", os.Getenv("HOME"), DEFAULT_LOG_DIR)
	_ = os.MkdirAll(log_dir, os.ModePerm) // #nosec G301 G703 -- legacy log dir under $HOME; openAppendLog fails loudly if unusable

	infoFile := openAppendLog(fmt.Sprintf("%v/randomwalk_info.log", log_dir))
	errFile := openAppendLog(fmt.Sprintf("%v/randomwalk_error.log", log_dir))
	Info = log.New(infoFile, "INFO: ", log.Ltime|log.Lshortfile)
	Error = log.New(errFile, "ERROR: ", log.Ltime|log.Lshortfile)
	// The engine logs structured records into the same files: everything to
	// the info log, errors duplicated to the error log.
	logger := slog.New(indexer.NewDualLogHandler(infoFile, errFile))

	// Graceful shutdown: on SIGINT/SIGTERM/SIGHUP finish the current event
	// batch, write status, and exit 0 cleanly. The engine checks ctx between
	// batches and during waits.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()
	go func() {
		<-ctx.Done()
		Info.Printf("Got signal, will exit after current batch is processed." +
			" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
	}()

	rpcclient, err := rpc.DialContext(ctx, RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n", RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	// Database log output (failed and slow queries) goes through the pgx
	// slog tracer into the file the legacy Init_log wrote to.
	dbLogHandle := openAppendLog(fmt.Sprintf("%v/randomwalk_%v", log_dir, DEFAULT_DB_LOG))
	cfg := store.ConfigFromEnv()
	cfg.Logger = slog.New(slog.NewTextHandler(dbLogHandle, nil))
	dbStore, err = store.New(ctx, cfg)
	if err != nil {
		Info.Printf("failed to connect to storage: %v", err)
		fmt.Fprintf(os.Stderr, "Can't connect to PostgreSQL database.\nConnection error: %v\n%s", err, store.ConnectHint(err))
		os.Exit(1)
	}
	rwRepo = rwdb.NewRepo(dbStore)

	abi_parsed1 := strings.NewReader(RWMarketABI)
	ab1, err := abi.JSON(abi_parsed1)
	if err != nil {
		Info.Printf("Can't parse Marketplace ABI: %v\n", err)
		os.Exit(1)
	}
	marketplace_abi = &ab1
	abi_parsed2 := strings.NewReader(RWalkABI)
	ab2, err := abi.JSON(abi_parsed2)
	if err != nil {
		Info.Printf("Can't parse RandomWalk ABI: %v\n", err)
		os.Exit(1)
	}
	randomwalk_abi = &ab2

	// First, read raw contract addresses and insert into address table
	rawMarketplace, rawRandomwalk, err := rwRepo.RawContractAddrs(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read rw_contracts: %v\n", err)
		os.Exit(1)
	}

	// Insert contract addresses into address table (for fresh database)
	for _, contract_addr := range []string{rawMarketplace, rawRandomwalk} {
		if _, err := dbStore.LookupOrCreateAddress(ctx, contract_addr, 0, 0); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to register contract address %v: %v\n", contract_addr, err)
			os.Exit(1)
		}
	}
	Info.Printf("Contract addresses registered in address table\n")

	// Now we can safely call the function that joins with address table
	rw_contracts, err = rwRepo.ContractAddrs(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve contract addresses: %v\n", err)
		os.Exit(1)
	}
	rwalk_addr = ethcommon.HexToAddress(rw_contracts.RandomWalk)
	market_addr = ethcommon.HexToAddress(rw_contracts.MarketPlace)
	Info.Printf("RandomWalk contract %v\n", rwalk_addr.String())
	Info.Printf("MarketPlace contract %v\n", market_addr.String())

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

	engine, err := indexer.New(indexer.Config{
		Store:     dbStore,
		Client:    eclient,
		Progress:  rwProgress{repo: rwRepo},
		Process:   process_single_event,
		Contracts: []ethcommon.Address{rwalk_addr, market_addr},
		Logger:    logger,
		Metrics:   metrics,
		TopicName: eventTopicName,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build indexer engine: %v\n", err)
		os.Exit(1)
	}

	if err := engine.Run(ctx); err != nil {
		Error.Printf("Event processing loop terminated: %v\n", err)
		Info.Printf("Event processing loop terminated: %v\n", err)
		os.Exit(1)
	}
}
