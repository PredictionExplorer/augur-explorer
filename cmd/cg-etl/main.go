package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives"
	. "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	. "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const (
	DEFAULT_DB_LOG = "db.log"
	IMGGEN_PATH    = "v2/etl/cmd/cosmicgame/imggen_monitor/imggen_exec" // relative to $HOME
)

var (
	eclient   *ethclient.Client
	rpcclient *rpc.Client

	cosmic_game_abi          *abi.ABI
	cosmic_game_v2_abi       *abi.ABI
	cosmic_signature_abi     *abi.ABI
	cosmic_token_abi         *abi.ABI
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
	cosmic_sig_aid            int64
	cosmic_tok_aid            int64

	cg_contracts CosmicGameContractAddrs
	storagew     SQLStorageWrapper
	// cgRepo carries the store queries already converted to the context-first
	// Repo; storagew keeps the rest until Phase 1 completes.
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
func main() {

	log_dir := fmt.Sprintf("%v/%v", os.Getenv("HOME"), DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file := fmt.Sprintf("%v/cosmicgame_%v", log_dir, DEFAULT_DB_LOG)

	fname := fmt.Sprintf("%v/cosmicgame_info.log", log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't start: %v\n", err)
		os.Exit(1)
	}
	Info = log.New(logfile, "INFO: ", log.Ltime|log.Lshortfile)

	fname = fmt.Sprintf("%v/cosmicgame_error.log", log_dir)
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

	st, err := store.New(context.Background(), store.ConfigFromEnv())
	if err != nil {
		Info.Printf("failed to connect to storage: %v", err)
		fmt.Fprintf(os.Stderr, "Can't connect to PostgreSQL database.\nConnection error: %v\n%s", err, store.ConnectHint(err))
		os.Exit(1)
	}
	cgRepo = NewRepo(st)
	storagew.S = store.NewSQLStorageFromDB(st.DB(), Info)
	storagew.S.Db_set_schema_name("public")
	if err := storagew.S.Init_log(db_log_file); err != nil {
		fmt.Fprintf(os.Stderr, "Can't initialize DB log: %v\n", err)
		os.Exit(1)
	}
	storagew.S.Log_msg("Log initialized\n")

	cosmic_game_abi = get_abi(CosmicSignatureGameABI)
	cosmic_game_v2_abi = get_abi(CosmicSignatureGameV2ABI)
	cosmic_signature_abi = get_abi(CosmicSignatureNftABI)
	cosmic_token_abi = get_abi(CosmicSignatureTokenABI)
	charity_wallet_abi = get_abi(CharityWalletABI)
	prizes_wallet_abi = get_abi(PrizesWalletABI)
	staking_wallet_cst_abi = get_abi(IStakingWalletCosmicSignatureNftABI)
	staking_wallet_rwalk_abi = get_abi(IStakingWalletRandomWalkNftABI)
	marketing_wallet_abi = get_abi(MarketingWalletABI)
	erc20_abi = get_abi(ERC20ABI)
	erc721_abi = get_abi(ERC721ABI)
	erc1967_abi = get_abi(IERC1967ABI)

	cg_contracts, err = cgRepo.ContractAddrs(context.Background())
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
		if _, err := storagew.S.Lookup_or_create_address(contract_addr, 0, 0); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to register contract address %v: %v\n", contract_addr, err)
			os.Exit(1)
		}
	}
	Info.Printf("All contract addresses registered in address table\n")

	// Now lookup the address IDs (they are guaranteed to exist now)
	cosmic_sig_aid, err = storagew.S.Nonfatal_lookup_address_id(cg_contracts.CosmicSignatureAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Lookup of CosmicSignatureAddr failed: %v", err)
		os.Exit(1)
	}
	cosmic_tok_aid, err = storagew.S.Nonfatal_lookup_address_id(cg_contracts.CosmicTokenAddr)
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

	if err := syncContractParamsFromChain(context.Background(), cgRepo, &storagew, eclient, cg_contracts.CosmicGameAddr, cg_contracts.PrizesWalletAddr, Info, Error); err != nil {
		Error.Printf("Contract param chain sync failed: %v", err)
		os.Exit(1)
	}

	// Graceful shutdown: on SIGINT/SIGTERM/SIGHUP finish the current event batch,
	// write status, and exit 0 cleanly. The polling loop checks ctx between batches.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()
	go func() {
		<-ctx.Done()
		Info.Printf("Got signal, will exit after current batch is processed." +
			" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
	}()

	// cosmic_dao_addr is in getContractAddresses(); DAO Governor events are stored in
	// evt_log only (no cg_dao_* layer-2 tables). process_single_event ignores unknown topics.
	if err := process_events_filterlog(ctx); err != nil {
		Error.Printf("Event processing loop terminated: %v\n", err)
		Info.Printf("Event processing loop terminated: %v\n", err)
		os.Exit(1)
	}
}
