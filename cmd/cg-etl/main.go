// The CosmicGame ETL: indexes every CosmicGame-family contract event into
// PostgreSQL. main wires the process-wide dependencies (loggers, RPC client,
// store, ABIs, contract addresses), runs the startup contract-parameter sync
// and hands control to the shared indexing engine (internal/indexer).
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

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/indexer"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	. "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const DEFAULT_DB_LOG = "db.log"

var (
	eclient *ethclient.Client

	cosmic_game_abi          *abi.ABI
	cosmic_game_v2_abi       *abi.ABI
	cosmic_signature_abi     *abi.ABI
	charity_wallet_abi       *abi.ABI
	prizes_wallet_abi        *abi.ABI
	staking_wallet_cst_abi   *abi.ABI
	staking_wallet_rwalk_abi *abi.ABI
	marketing_wallet_abi     *abi.ABI
	erc20_abi                *abi.ABI
	erc721_abi               *abi.ABI
	erc1967_abi              *abi.ABI

	cosmic_game_addr          ethcommon.Address
	cosmic_signature_addr     ethcommon.Address
	cosmic_token_addr         ethcommon.Address
	cosmic_dao_addr           ethcommon.Address
	charity_wallet_addr       ethcommon.Address
	prizes_wallet_addr        ethcommon.Address
	staking_wallet_cst_addr   ethcommon.Address
	staking_wallet_rwalk_addr ethcommon.Address
	marketing_wallet_addr     ethcommon.Address
	implementation_addr       ethcommon.Address
	cosmic_tok_aid            int64

	cg_contracts CosmicGameContractAddrs
	// dbStore owns the process-wide connection pool; cgRepo runs every
	// CosmicGame query on it.
	dbStore *store.Store
	cgRepo  *Repo
	RPC_URL = os.Getenv("RPC_URL")
	Error   *log.Logger
	Info    *log.Logger
)

func get_abi(abi_str string) *abi.ABI {
	abi_parsed := strings.NewReader(abi_str)
	abi_obj, err := abi.JSON(abi_parsed)
	if err != nil {
		Info.Printf("Can't parse ABI: %v\n", err)
		os.Exit(1)
	}
	return &abi_obj
}

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

// cgProgress adapts the cg_proc_status row to the engine's watermark
// interface, preserving last_evt_id across writes.
type cgProgress struct {
	repo *Repo
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

// contractAddresses returns every contract the FilterLogs subscription
// watches. cosmic_dao_addr is included although DAO Governor events are
// stored in evt_log only (no cg_dao_* layer-2 tables); the dispatcher
// ignores unknown topics.
func contractAddresses() []ethcommon.Address {
	return []ethcommon.Address{
		cosmic_game_addr,
		cosmic_signature_addr,
		cosmic_token_addr,
		cosmic_dao_addr,
		charity_wallet_addr,
		prizes_wallet_addr,
		staking_wallet_cst_addr,
		staking_wallet_rwalk_addr,
		marketing_wallet_addr,
		implementation_addr,
	}
}

func main() {

	log_dir := fmt.Sprintf("%v/%v", os.Getenv("HOME"), DEFAULT_LOG_DIR)
	_ = os.MkdirAll(log_dir, os.ModePerm) // #nosec G301 G703 -- legacy log dir under $HOME; openAppendLog fails loudly if unusable

	infoFile := openAppendLog(fmt.Sprintf("%v/cosmicgame_info.log", log_dir))
	errFile := openAppendLog(fmt.Sprintf("%v/cosmicgame_error.log", log_dir))
	Info = log.New(infoFile, "INFO: ", log.Ltime|log.Lshortfile)
	Error = log.New(errFile, "ERROR: ", log.Ltime|log.Lshortfile)
	// The engine and the startup sync log structured records into the same
	// files: everything to the info log, errors duplicated to the error log.
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
	dbLogHandle := openAppendLog(fmt.Sprintf("%v/cosmicgame_%v", log_dir, DEFAULT_DB_LOG))
	cfg := store.ConfigFromEnv()
	cfg.Logger = slog.New(slog.NewTextHandler(dbLogHandle, nil))
	dbStore, err = store.New(ctx, cfg)
	if err != nil {
		Info.Printf("failed to connect to storage: %v", err)
		fmt.Fprintf(os.Stderr, "Can't connect to PostgreSQL database.\nConnection error: %v\n%s", err, store.ConnectHint(err))
		os.Exit(1)
	}
	cgRepo = NewRepo(dbStore)

	cosmic_game_abi = get_abi(CosmicSignatureGameABI)
	cosmic_game_v2_abi = get_abi(CosmicSignatureGameV2ABI)
	cosmic_signature_abi = get_abi(CosmicSignatureNftABI)
	charity_wallet_abi = get_abi(CharityWalletABI)
	prizes_wallet_abi = get_abi(PrizesWalletABI)
	staking_wallet_cst_abi = get_abi(IStakingWalletCosmicSignatureNftABI)
	staking_wallet_rwalk_abi = get_abi(IStakingWalletRandomWalkNftABI)
	marketing_wallet_abi = get_abi(MarketingWalletABI)
	erc20_abi = get_abi(ERC20ABI)
	erc721_abi = get_abi(ERC721ABI)
	erc1967_abi = get_abi(IERC1967ABI)

	cg_contracts, err = cgRepo.ContractAddrs(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read contract addresses: %v\n", err)
		os.Exit(1)
	}

	// Insert all contract addresses into address table (for fresh database)
	for _, contract_addr := range []string{
		cg_contracts.CosmicGameAddr,
		cg_contracts.CosmicSignatureAddr,
		cg_contracts.CosmicTokenAddr,
		cg_contracts.CosmicDaoAddr,
		cg_contracts.CharityWalletAddr,
		cg_contracts.PrizesWalletAddr,
		cg_contracts.RandomWalkAddr,
		cg_contracts.StakingWalletCSTAddr,
		cg_contracts.StakingWalletRWalkAddr,
		cg_contracts.MarketingWalletAddr,
		cg_contracts.ImplementationAddr,
	} {
		if _, err := dbStore.LookupOrCreateAddress(ctx, contract_addr, 0, 0); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to register contract address %v: %v\n", contract_addr, err)
			os.Exit(1)
		}
	}
	Info.Printf("All contract addresses registered in address table\n")

	// Now lookup the address IDs (they are guaranteed to exist now)
	cosmic_tok_aid, err = dbStore.LookupAddressID(ctx, cg_contracts.CosmicTokenAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Lookup of CosmicTokenAddr failed: %v", err)
		os.Exit(1)
	}
	cosmic_game_addr = ethcommon.HexToAddress(cg_contracts.CosmicGameAddr)
	cosmic_signature_addr = ethcommon.HexToAddress(cg_contracts.CosmicSignatureAddr)
	cosmic_token_addr = ethcommon.HexToAddress(cg_contracts.CosmicTokenAddr)
	cosmic_dao_addr = ethcommon.HexToAddress(cg_contracts.CosmicDaoAddr)
	charity_wallet_addr = ethcommon.HexToAddress(cg_contracts.CharityWalletAddr)
	prizes_wallet_addr = ethcommon.HexToAddress(cg_contracts.PrizesWalletAddr)
	staking_wallet_cst_addr = ethcommon.HexToAddress(cg_contracts.StakingWalletCSTAddr)
	staking_wallet_rwalk_addr = ethcommon.HexToAddress(cg_contracts.StakingWalletRWalkAddr)
	marketing_wallet_addr = ethcommon.HexToAddress(cg_contracts.MarketingWalletAddr)
	implementation_addr = ethcommon.HexToAddress(cg_contracts.ImplementationAddr)

	if err := syncContractParamsFromChain(ctx, cgRepo, dbStore, eclient, cg_contracts.CosmicGameAddr, cg_contracts.PrizesWalletAddr, logger); err != nil {
		Error.Printf("Contract param chain sync failed: %v", err)
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

	engine, err := indexer.New(indexer.Config{
		Store:     dbStore,
		Client:    eclient,
		Progress:  cgProgress{repo: cgRepo},
		Process:   process_single_event,
		Contracts: contractAddresses(),
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
