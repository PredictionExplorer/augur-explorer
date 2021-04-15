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
	NEW_HATCHERY= "08afdadd49d632c11dbde177a7ab47701b5adaac8f633beedd892c8da8d4393f"
	TURBO_CREATED = "2c4d919a4805caed2e2fdd9bb8a122413c2a643b61e08b957445484bbbfd8f4f"
	COMPLETE_SETS_MINTED = "51b2bca5bb2f65b2670950591ce7b54cfc4d99b2db85abfea36b8b92d10ac380"
	COMPLETE_SETS_BURNED = "2df8f390c89a8c8e8b89875f61085269c64b16b81e7745b844ba42a40a3dde27"
	CLAIM = "7bb2b3c10797baccb6f8c4791f1edd6ca2f0d028ee0eda64b01a9a57e3a653f7"

	NUM_AUGUR_CONTRACTS int = 35
)
var (
	evt_pool_created,_ = hex.DecodeString(POOL_CREATED)
	//event PoolCreated(
	//	address _pool,
	//	address indexed _hatchery,
	//	uint256 indexed _turboId,
	//	address indexed _creator
	//);

	evt_new_hatchery,_ = hex.DecodeString(NEW_HATCHERY)
	//event NewHatchery(
	//	address id,
	//	address indexed collateral,
	//	address shareTokenFactory,
	//	address feePot
	//);

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

	evt_complete_sets_minted,_ = hex.DecodeString(COMPLETE_SETS_MINTED)
	//event CompleteSetsMinted(
	//	uint256 turboId,
	//	uint256 amount,
	//	address target
	//);
	evt_complete_sets_burned,_ = hex.DecodeString(COMPLETE_SETS_BURNED)
	//event CompleteSetsBurned(
	//	uint256 turboId,
	//	uint256 amount,
	//	address target
	//);

	evt_claim,_ = hex.DecodeString(CLAIM)
	//event Claim(uint256 turboId);

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

	var max_batch_size int64 = 1024*100
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
		if er != nil {
			Error.Printf("Error: %v. Possibly 'evt_log' table is empty, aborting",err)
			os.Exit(1)
		}
		if  id_upper_limit > last_evt_id {
			_upper_limit = last_evt_id
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
