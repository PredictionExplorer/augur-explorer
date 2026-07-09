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

	. "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	rwp "github.com/PredictionExplorer/augur-explorer/internal/primitives/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwdb "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

const (
	DEFAULT_DB_LOG = "db.log"
)

var (
	// dbStore owns the process-wide connection pool; rwRepo runs every
	// RandomWalk query on it.
	dbStore *store.Store
	rwRepo  *rwdb.Repo
	RPC_URL = os.Getenv("RPC_URL")
	Error   *log.Logger
	Info    *log.Logger

	eclient   *ethclient.Client
	rpcclient *rpc.Client

	marketplace_abi *abi.ABI
	randomwalk_abi  *abi.ABI

	rw_contracts rwp.ContractAddresses
	market_addr  ethcommon.Address
	rwalk_addr   ethcommon.Address
)

func main() {

	log_dir := fmt.Sprintf("%v/%v", os.Getenv("HOME"), DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file := fmt.Sprintf("%v/randomwalk_%v", log_dir, DEFAULT_DB_LOG)

	fname := fmt.Sprintf("%v/randomwalk_info.log", log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't start: %v\n", err)
		os.Exit(1)
	}
	Info = log.New(logfile, "INFO: ", log.Ltime|log.Lshortfile)

	fname = fmt.Sprintf("%v/randomwalk_error.log", log_dir)
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't start: %v\n", err)
		os.Exit(1)
	}
	Error = log.New(logfile, "ERROR: ", log.Ltime|log.Lshortfile)

	rpcclient, err = rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n", RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	// Database log output (failed and slow queries) goes through the pgx
	// slog tracer into the file the legacy Init_log wrote to.
	dbLogHandle, err := os.OpenFile(db_log_file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't initialize DB log: %v\n", err)
		os.Exit(1)
	}
	cfg := store.ConfigFromEnv()
	cfg.Logger = slog.New(slog.NewTextHandler(dbLogHandle, nil))
	dbStore, err = store.New(context.Background(), cfg)
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
	rawMarketplace, rawRandomwalk, err := rwRepo.RawContractAddrs(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read rw_contracts: %v\n", err)
		os.Exit(1)
	}

	// Insert contract addresses into address table (for fresh database)
	for _, contract_addr := range []string{rawMarketplace, rawRandomwalk} {
		if _, err := dbStore.LookupOrCreateAddress(context.Background(), contract_addr, 0, 0); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to register contract address %v: %v\n", contract_addr, err)
			os.Exit(1)
		}
	}
	Info.Printf("Contract addresses registered in address table\n")

	// Now we can safely call the function that joins with address table
	rw_contracts, err = rwRepo.ContractAddrs(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve contract addresses: %v\n", err)
		os.Exit(1)
	}
	rwalk_addr = ethcommon.HexToAddress(rw_contracts.RandomWalk)
	market_addr = ethcommon.HexToAddress(rw_contracts.MarketPlace)
	Info.Printf("RandomWalk contract %v\n", rwalk_addr.String())
	Info.Printf("MarketPlace contract %v\n", market_addr.String())

	// Graceful shutdown: on SIGINT/SIGTERM/SIGHUP finish the current event batch,
	// write status, and exit 0 cleanly. The polling loop checks ctx between batches.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()
	go func() {
		<-ctx.Done()
		Info.Printf("Got signal, will exit after current batch is processed." +
			" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
	}()

	// Use new FilterLogs-based event processing
	if err := process_events_filterlog(ctx); err != nil {
		Error.Printf("Event processing loop terminated: %v\n", err)
		Info.Printf("Event processing loop terminated: %v\n", err)
		os.Exit(1)
	}
}
