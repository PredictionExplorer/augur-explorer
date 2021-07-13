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
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	DEFAULT_DB_LOG				= "db.log"

	CONDITION_PREPARATION = "ab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177"
	CONDITION_RESOLUTION = "b44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894"
	PAYOUT_REDEMPTION = "2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d"
	POSITION_SPLIT = "2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298"
	POSITION_MERGE = "f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca"
	URI = "6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b"

)
var (
	evt_condition_preparation,_ = hex.DecodeString(CONDITION_PREPARATION)
	evt_condition_resolution,_ = hex.DecodeString(CONDITION_RESOLUTION)
	evt_payout_redemption,_ = hex.DecodeString(PAYOUT_REDEMPTION)
	evt_position_split,_ = hex.DecodeString(POSITION_SPLIT)
	evt_position_merge,_ = hex.DecodeString(POSITION_MERGE)
	evt_uri,_ = hex.DecodeString(URI)

	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	Error   *log.Logger
	Info	*log.Logger
	market_order_id int64 = 0
	inspected_events []InspectedEvent

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
func process_conditional_token_events(exit_chan chan bool) {

	var max_batch_size int64 = 1024*100
	for {
		status := storage.Get_polymarkets_processing_status()
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		Info.Printf("scanning event range from %v to %v\n",status.LastIdProcessed,status.LastIdProcessed+max_batch_size)
		id_upper_limit := status.LastIdProcessed + max_batch_size
		last_evt_id,err := storage.Get_last_evtlog_id()
		if err != nil {
			Error.Printf("Error: %v. Possibly 'evt_log' table is empty, aborting",err)
			os.Exit(1)
		}
		if  id_upper_limit > last_evt_id {
			id_upper_limit = last_evt_id
		}
		events := get_event_ids(status.LastIdProcessed,id_upper_limit)
		for _,evt_id := range events {
			err := process_polymarket_event(evt_id)
			if err != nil {
				Error.Printf("Pausing event processing loop for 5 sec due to error")
				time.Sleep(5 * time.Second)
				break
			}
			status.LastIdProcessed=evt_id
			storage.Update_polymarkets_process_status(&status)
		}
		if len(events) == 0 {
			last_evt_log_id_on_chain,err := storage.Get_last_evtlog_id()
			if err != nil {
				Info.Printf("Error getting last event log id: %v\n",err)
				os.Exit(1)
			}
			if last_evt_log_id_on_chain > id_upper_limit {
				// only advance upper range if events within the range have filled id value space
				status.LastIdProcessed = id_upper_limit
				storage.Update_polymarkets_process_status(&status)
			}
			time.Sleep(1 * time.Second) // sleep only if there is no data
		}
	}
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/polymarkets_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/polymarkets_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/polymarkets_error.log",log_dir)
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

	abi_parsed := strings.NewReader(ConditionalTokenABI)
	aa_abi,err = abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse Polymarkets ABI: %v\n",err)
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

	inspected_events = build_list_of_inspected_events()

	process_conditional_token_events(exit_chan)
}
