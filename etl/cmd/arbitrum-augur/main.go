package main


import (
	"os"
	"os/signal"
	"syscall"
	"sort"
	"time"
	"fmt"
	"strings"
	"context"
	"log"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DEFAULT_DB_LOG				= "db.log"

	POOL_CREATED = "0537d3e5d88810bbfb16874b28bc0f95856d7bb24c8f29511fe463c5b1d27c6b"
	SPORTS_MARKET_CREATED = "afad6545e5200f9fdf4db34dfca61a9e7d72518593fd5155d11fd69c20e1555f"
	PRICE_MARKET_CREATED = "28c8de42a10b7bcc4a65ea3618bc8ada3e24cd7394886eae1b4f5f7440477080"
	TRUSTED_MARKET_CREATED = "a1bb41461c32765a0cc838c35ce6b8e28985bb6a069dfe2af0873796438670d4"
	SHARES_MINTED = "d81c0442e10068a9818f3aa093c9ccb804584690df572d7df3da2d892a6973f2"
	SHARES_BURNED = "b6fdb729b2ed801daf629f0ab713e4a7a73619505790f6f27fd92d6f2c9688d7"
	MARKET_RESOLVED = "c68d106ea6e4bec784925cfd91767212c71ced92adbac107dc364435321113f6"
	WINNINGS_CLAIMED = "2bdd7a5109193ce6270ec3b4afcf4ccd4a06c27742ba11f660498cb41433bb00"

	ERC20_TRANSFER = "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	//																		(used by FeePot)

)
var (
	evt_pool_created,_ = hex.DecodeString(POOL_CREATED)
	// event PoolCreated(
	//	address pool,
	//	address indexed marketFactory,
	//	uint256 indexed marketId,
	//	address indexed creator
	//);
							// DISCONTINUED
							//event PoolCreated(
							//	address _pool,
							//	address indexed _hatchery,
							//	uint256 indexed _turboId,
							//	address indexed _creator
							//);

/*DISCONTINUED							
	evt_new_hatchery,_ = hex.DecodeString(NEW_HATCHERY)
	//event NewHatchery(
	//	address id,
	//	address indexed collateral,
	//	address shareTokenFactory,
	//	address feePot
	//);
*/
/* DISCONTINUED
	evt_turbo_created,_ = hex.DecodeString(TURBO_CREATED)
	//    event TurboCreated(
	//        uint256 id,
	//	        uint256 creatorFee,
	//        string[] outcomeSymbols,
	//        bytes32[] outcomeNames,
	//        uint256 numTicks,
	//        IArbiter arbiter,
	//        bytes arbiterConfiguration,
	//        uint256 indexed index
	//    );
*/
	// Price Market
	evt_price_market_created,_ = hex.DecodeString(PRICE_MARKET_CREATED)
	//event MarketCreated(
	//	uint256 id, address creator,
	//	uint256 endTime,
	//	uint256 spotPrice
	//);

	// Sports  Market
	evt_sports_market_created,_ = hex.DecodeString(SPORTS_MARKET_CREATED)
	//event MarketCreated(
	//  uint256 id,
	//  address creator,
	//	uint256 endTime,
	//	MarketType marketType,
	//	uint256 indexed eventId,
	//	uint256 homeTeamId,
	//	uint256 awayTeamId,
	//	uint256 estimatedStarTime,
	//	int256 score
	//);

	// Trusted Market
	evt_trusted_market_created,_ = hex.DecodeString(TRUSTED_MARKET_CREATED)
	//event MarketCreated(
	//	uint256 id,
	//	address creator,
	//	uint256 _endTime,
	//	string description,
	//	string[] outcomes
	//);

	// MarketResolved
	evt_market_resolved,_= hex.DecodeString(MARKET_RESOLVED)
	//event MarketResolved(
	//	uint256 id,
	//	address winner
	//);


	/* DISCONTINUED
	//event CompleteSetsMinted(
	//	uint256 turboId,
	//	uint256 amount,
	//	address target
	//);
	*/

	evt_shares_minted,_ = hex.DecodeString(SHARES_MINTED)
	//event SharesMinted(
	//	uint256 id,
	//	uint256 amount,
	//	address receiver
	//);

	/*DISCONTINUED
	//event CompleteSetsBurned(
	//	uint256 turboId,
	//	uint256 amount,
	//	address target
	//);
	*/
	evt_shares_burned,_ = hex.DecodeString(SHARES_BURNED)
	//event SharesBurned(
	//	uint256 id,
	//	uint256 amount,
	//	address receiver
	//);


	// DISCONTINUED //event Claim(uint256 turboId);
	evt_winnings_claimed,_ = hex.DecodeString(WINNINGS_CLAIMED)
	//event WinningsClaimed(
	//	uint256 id,
	//	uint256 amount,
	//	address receiver
	//);

	evt_erc20_transfer,_ = hex.DecodeString(ERC20_TRANSFER)

	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	Error   *log.Logger
	Info	*log.Logger
	market_order_id int64 = 0
	inspected_events []InspectedEvent

	caddrs *AA_ContractAddrs
	augur_abi *abi.ABI
	aa_abi abi.ABI

	ctrct_wallet_registry *AugurWalletRegistry

	eclient *ethclient.Client
	rpcclient *rpc.Client

)
func get_event_ids(from_evt_id,to_evt_id int64) []int64 {
	output := make([]int64 ,0,1024)
	for _,e := range inspected_events {
		event_list := storage.Get_evtlogs_by_signature_only_in_range(
			e.Signature,from_evt_id,to_evt_id,
		)
		/*Info.Printf("selected events for signature %v:\n",e.Signature)
		for _,evt_id := range event_list {
			Info.Printf("\tEvtId:\t%9v\n",evt_id)
		}*/
		output = append(output,event_list...)
	}
	sort.Slice(output, func(i, j int) bool { return output[i] < output[j] })
	num_elts:=Remove_duplicates_int64(output)
	return output[0:num_elts]
}
func process_arbitrum_augur_events(exit_chan chan bool) {

	var max_batch_size int64 = 1024*5
	for {
		status := storage.Get_arbitrum_augur_processing_status()
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		Info.Printf("scanning event range from %v to %v\n",status.LastEvtId,status.LastEvtId+max_batch_size)
		id_upper_limit := status.LastEvtId + max_batch_size
		last_evt_id,err := storage.Get_last_evtlog_id()
		if err != nil {
			Error.Printf("Error: %v. Possibly 'evt_log' table is empty, aborting",err)
			os.Exit(1)
		}
		if  id_upper_limit > last_evt_id {
			id_upper_limit = last_evt_id
		}
		events := get_event_ids(status.LastEvtId,id_upper_limit)
		for _,evt_id := range events {
			err := process_arbitrum_augur_event(evt_id)
			if err != nil {
				Error.Printf("Pausing event processing loop for 5 sec due to error")
				time.Sleep(5 * time.Second)
				break
			}
			status.LastEvtId=evt_id
			storage.Update_arbitrum_augur_process_status(&status)
		}
		if len(events) == 0 {
			last_evt_log_id_on_chain,err := storage.Get_last_evtlog_id()
			if err != nil {
				Info.Printf("Error getting last event log id: %v\n",err)
				os.Exit(1)
			}
			if last_evt_log_id_on_chain > id_upper_limit {
				// only advance upper range if events within the range have filled id value space
				status.LastEvtId = id_upper_limit
				storage.Update_arbitrum_augur_process_status(&status)
			}
			time.Sleep(1 * time.Second) // sleep only if there is no data
		}
	}
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/arbitrum_augur_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/arbitrum_augur_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/arbitrum_augur_error.log",log_dir)
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

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

	abi_parsed := strings.NewReader(Arbitrum_ABI)
	aa_abi,err = abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse Arbitrum Augur ABI: %v\n",err)
		os.Exit(1)
	}

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()

	caddrs_obj,err := storage.Get_arbitrum_augur_contract_addresses()
	if err!=nil {
		Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs = &caddrs_obj

	inspected_events = build_list_of_inspected_events()
	process_arbitrum_augur_events(exit_chan)
}
