package main


import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"context"
	"sort"
	"log"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	DEFAULT_DB_LOG				= "db.log"

	NEW_OFFER =		"55076e90b6b34a2569ffb2e1e34ee0da92d30ca423f0d6cfb317d252ade9a56a"
	//NEW_OFFER =		"8b4d06c200b17b9c1150172953ceb6fa3e7ace7623f6f933707badfa52c354cf"
	ITEM_BOUGHT=	"7f7e375fbeaef0d6ebfc53af15b7aeed1d704e3424f34ef67e914b1f68f8c8ef"
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

	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	Error   *log.Logger
	Info	*log.Logger
	market_order_id int64 = 0
	inspected_events []InspectedEvent

	eclient *ethclient.Client
	rpcclient *rpc.Client

	marketplace_abi *abi.ABI
	randomwalk_abi *abi.ABI

	rw_contracts RW_ContractAddresses
	market_addr common.Address
	rwalk_addr common.Address
)
func get_event_ids(from_evt_id,to_evt_id int64) []int64 {
	output := make([]int64 ,0,1024)
	for _,e := range inspected_events {
		var event_list []int64
		if hex.EncodeToString(evt_transfer[:4]) == e.Signature {
			// Transfer event is filtered by contract address for speed
			event_list = storage.Get_evtlogs_by_signature_in_range(
				e.Signature,fmt.Sprintf("%v",e.ContractAid),from_evt_id,to_evt_id,
			)
		} else {
			// all other events are fetched without filtering by Contract address
			event_list = storage.Get_evtlogs_by_signature_only_in_range(
				e.Signature,from_evt_id,to_evt_id,
			)
		}
		output = append(output,event_list...)
	}
	sort.Slice(output, func(i, j int) bool { return output[i] < output[j] })
	num_elts:=Remove_duplicates_int64(output)
	return output[0:num_elts]
}
func process_events(exit_chan chan bool) {

	var max_batch_size int64 = 1024*100
	for {
		status := storage.Get_randomwalk_processing_status()
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
			err := process_single_event(evt_id)
			if err != nil {
				Error.Printf("Pausing event processing loop for 5 sec due to error")
				time.Sleep(5 * time.Second)
				break
			}
			status.LastIdProcessed=evt_id
			storage.Update_randomwalk_process_status(&status)
		}
		if len(events) == 0 {
			status.LastIdProcessed = id_upper_limit
			storage.Update_randomwalk_process_status(&status)
			time.Sleep(1 * time.Second) // sleep only if there is no data
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

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

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

	rw_contracts = storage.Get_randomwalk_contract_addresses()
	rwalk_addr = common.HexToAddress(rw_contracts.RandomWalk)
	market_addr = common.HexToAddress(rw_contracts.MarketPlace)
	Info.Printf("RandomWalk contract %v\n",rwalk_addr.String())
	Info.Printf("MarketPlace contract %v\n",market_addr.String())

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()


	inspected_events = build_list_of_inspected_events_layer1(rw_contracts.RandomWalkAid)
	process_events(exit_chan)
}
