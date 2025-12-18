package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
	rwdb "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/randomwalk"
	etlcommon "github.com/PredictionExplorer/augur-explorer/rwcg/etl/common"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/primitives"
)
const (
	DEFAULT_DB_LOG				= "db.log"

	NEW_OFFER =		"55076e90b6b34a2569ffb2e1e34ee0da92d30ca423f0d6cfb317d252ade9a56a"
	//NEW_OFFER =		"8b4d06c200b17b9c1150172953ceb6fa3e7ace7623f6f933707badfa52c354cf"
	//ITEM_BOUGHT=	"7f7e375fbeaef0d6ebfc53af15b7aeed1d704e3424f34ef67e914b1f68f8c8ef"
	ITEM_BOUGHT=	"caacc56f18ca259dc5175dae29eb0ca81407703a4819958c6885acbb7d4f3af3"
	OFFER_CANCELED=	"0ff09947dd7d2583091e8cbfb427fecacb697bf895187b243fd0072c0ee9b951"
	WITHDRAWAL_EVT=	"a11b556ace4b11a5cae8675a293b51e8cde3a06387d34010861789dfd9e9abc7"
	TOKEN_NAME_EVT=	"8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12"
	MINT_EVENT =	"ad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec"
	TRANSFER_EVT=	"ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)
var (
	evt_new_offer ,_ = hex.DecodeString(NEW_OFFER)
	evt_item_bought,_ = hex.DecodeString(ITEM_BOUGHT)
	evt_offer_canceled,_ = hex.DecodeString(OFFER_CANCELED)
	evt_withdrawal,_ = hex.DecodeString(WITHDRAWAL_EVT)
	evt_token_name,_ = hex.DecodeString(TOKEN_NAME_EVT)
	evt_transfer,_	 = hex.DecodeString(TRANSFER_EVT)
	evt_mint_event,_ = hex.DecodeString(MINT_EVENT)

	storage *dbs.SQLStorage
	storagew rwdb.SQLStorageWrapper
	RPC_URL = os.Getenv("RPC_URL")
	Error   *log.Logger
	Info	*log.Logger
	inspected_events []InspectedEvent

	eclient *ethclient.Client
	rpcclient *rpc.Client

	marketplace_abi *abi.ABI
	randomwalk_abi *abi.ABI

	rw_contracts RW_ContractAddresses
	market_addr ethcommon.Address
	rwalk_addr ethcommon.Address
)

// getContractAddresses returns all contract addresses for FilterLogs
func getContractAddresses() []ethcommon.Address {
	return []ethcommon.Address{
		rwalk_addr,
		market_addr,
	}
}

// process_events_filterlog uses FilterLogs to get events directly from blockchain
func process_events_filterlog(exit_chan chan bool) {
	// Create ETL context for common operations
	ctx := &etlcommon.ETLContext{
		Storage:   storage,
		EthClient: eclient,
		RpcClient: rpcclient,
		Info:      Info,
		Error:     Error,
	}

	// Adaptive batch sizing: start large, reduce if we get events
	var batchSize uint64 = 100000      // Start with 100k blocks
	var minBatchSize uint64 = 1000     // Minimum when processing events
	var maxBatchSize uint64 = 1000000  // Maximum when scanning empty ranges
	contracts := getContractAddresses()
	
	// Debug: log the addresses being used for FilterLogs
	Info.Printf("FilterLogs will query these contract addresses:\n")
	for i, addr := range contracts {
		Info.Printf("  [%d] %s\n", i, addr.Hex())
	}

	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}

		// Get last processed block from status
		status := storagew.Get_randomwalk_processing_status()
		lastProcessedBlock := status.LastBlockNum
		if lastProcessedBlock == 0 {
			// If no blocks processed yet, start from the block where contracts were deployed
			lastProcessedBlock, _ = storage.Get_last_block_num()
		}

		// Get current block from chain
		currentBlock, err := etlcommon.GetCurrentBlockNumber(eclient)
		if err != nil {
			Error.Printf("Failed to get current block number: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Calculate block range to process
		fromBlock := uint64(lastProcessedBlock + 1)
		toBlock := fromBlock + batchSize - 1
		if toBlock > currentBlock {
			toBlock = currentBlock
		}

		if fromBlock > currentBlock {
			// Already caught up, wait for new blocks
			time.Sleep(2 * time.Second)
			batchSize = minBatchSize // Reset to small batch for real-time
			continue
		}

		Info.Printf("Fetching events from block %d to %d (batch size: %d)\n", fromBlock, toBlock, batchSize)

		// Fetch events using FilterLogs
		logs, err := etlcommon.FetchEvents(eclient, fromBlock, toBlock, contracts)
		if err != nil {
			Error.Printf("FetchEvents failed: %v", err)
			// Reduce batch size on error (might be too large)
			batchSize = batchSize / 2
			if batchSize < minBatchSize {
				batchSize = minBatchSize
			}
			time.Sleep(5 * time.Second)
			continue
		}

		Info.Printf("Received %d events\n", len(logs))

		// Process each event
		var processingFailed bool
		var lastSuccessfulBlock uint64
		for _, log := range logs {
			// Ensure block exists with correct hash (chain split detection)
			_, err := etlcommon.EnsureBlockExists(ctx, int64(log.BlockNumber), log.BlockHash.Hex())
			if err != nil {
				Error.Printf("EnsureBlockExists failed for block %d: %v", log.BlockNumber, err)
				processingFailed = true
				time.Sleep(5 * time.Second)
				break
			}

			// Ensure transaction exists
			txId, _, err := etlcommon.EnsureTransactionExists(ctx, log.TxHash, int64(log.BlockNumber))
			if err != nil {
				Error.Printf("EnsureTransactionExists failed for tx %s: %v", log.TxHash.Hex(), err)
				processingFailed = true
				time.Sleep(5 * time.Second)
				break
			}

			// Insert event log
			evtId, err := etlcommon.InsertEventLog(ctx, log, txId)
			if err != nil {
				Error.Printf("InsertEventLog failed: %v", err)
				processingFailed = true
				time.Sleep(5 * time.Second)
				break
			}

			// Process the event using existing logic
			err = process_single_event(evtId)
			if err != nil {
				Error.Printf("process_single_event failed for evt %d: %v", evtId, err)
				// Continue processing other events
			}
			
			// Track last successfully processed block
			lastSuccessfulBlock = log.BlockNumber
		}

		// Only update status if processing succeeded
		if !processingFailed {
			status.LastBlockNum = int64(toBlock)
			storagew.Update_randomwalk_process_status(&status)
		} else if lastSuccessfulBlock > 0 {
			// Update to last successfully processed block
			status.LastBlockNum = int64(lastSuccessfulBlock)
			storagew.Update_randomwalk_process_status(&status)
		}
		// If processingFailed and lastSuccessfulBlock==0, don't update - will retry same batch

		// Adaptive batch sizing
		if len(logs) == 0 {
			// No events - increase batch size for faster scanning
			batchSize = batchSize * 2
			if batchSize > maxBatchSize {
				batchSize = maxBatchSize
			}
		} else {
			// Found events - use smaller batch for granularity
			batchSize = minBatchSize
		}
	}
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/randomwalk_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/randomwalk_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/randomwalk_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	storage = dbs.Connect_to_storage(Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")
	storagew.S = storage

	abi_parsed1 := strings.NewReader(RWMarketABI)
	ab1,err := abi.JSON(abi_parsed1)
	if err!= nil {
		Info.Printf("Can't parse Marketplace ABI: %v\n",err)
		os.Exit(1)
	}
	marketplace_abi = &ab1
	abi_parsed2 := strings.NewReader(RWalkABI)
	ab2,err := abi.JSON(abi_parsed2)
	if err != nil {
		Info.Printf("Can't parse RandomWalk ABI: %v\n",err)
		os.Exit(1)
	}
	randomwalk_abi = &ab2

	// First, read raw contract addresses and insert into address table
	var rawMarketplace, rawRandomwalk string
	query := "SELECT marketplace_addr, randomwalk_addr FROM rw_contracts LIMIT 1"
	err = storage.Db().QueryRow(query).Scan(&rawMarketplace, &rawRandomwalk)
	if err != nil {
		fmt.Printf("Failed to read rw_contracts: %v\n", err)
		os.Exit(1)
	}
	
	// Insert contract addresses into address table (for fresh database)
	storage.Lookup_or_create_address(rawMarketplace, 0, 0)
	storage.Lookup_or_create_address(rawRandomwalk, 0, 0)
	Info.Printf("Contract addresses registered in address table\n")
	
	// Now we can safely call the function that joins with address table
	rw_contracts = storagew.Get_randomwalk_contract_addresses()
	rwalk_addr = ethcommon.HexToAddress(rw_contracts.RandomWalk)
	market_addr = ethcommon.HexToAddress(rw_contracts.MarketPlace)
	Info.Printf("RandomWalk contract %v\n",rwalk_addr.String())
	Info.Printf("MarketPlace contract %v\n",market_addr.String())

	c := make(chan os.Signal, 1)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-c
		Info.Printf("Got signal, will exit after current batch is processed." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()

	// Use new FilterLogs-based event processing
	process_events_filterlog(exit_chan)
}
