package main

import (
	"os"
	"log"
	"fmt"
	"sort"
	"time"
	"context"
	"strings"
	"encoding/hex"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/primitives/biddingwar"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/dbs/biddingwar"
)
const (
	DEFAULT_DB_LOG			= "db.log"

	PRIZE_CLAIM_EVENT		= "27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce"
	BID_EVENT				= "22664806da305c71fc17312c8fafba247a01ea6a4ee8a2e97f06cd21d979eefe"
	DONATION_EVENT			= "8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1"
	DONATION_RECEIVED_EVENT	= "46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368"
	DONATION_SENT_EVENT		= "44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32"
	CHARITY_UPDATED			= "a0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe"
	TOKEN_NAME_EVENT		= "8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12"
	MINT_EVENT				= "af162acd8d98cd428a10ad5028b47e3d3e50b4089880be8bf474aa921fed6b2e"
)
var (
	eclient 				*ethclient.Client
	rpcclient 				*rpc.Client

	evt_prize_claim_event,_ = hex.DecodeString(PRIZE_CLAIM_EVENT)
	evt_bid_event,_			= hex.DecodeString(BID_EVENT)
	evt_donation_event,_	= hex.DecodeString(DONATION_EVENT)
	evt_donation_received_event,_=hex.DecodeString(DONATION_RECEIVED_EVENT)
	evt_donation_sent_event,_= hex.DecodeString(DONATION_SENT_EVENT)
	evt_charity_updated,_	= hex.DecodeString(CHARITY_UPDATED)
	evt_token_name_event,_	= hex.DecodeString(TOKEN_NAME_EVENT)
	evt_mint_event,_		= hex.DecodeString(MINT_EVENT)

	inspected_events []InspectedEvent

	biddingwar_abi			*abi.ABI
	cosmic_signature_abi	*abi.ABI
	cosmic_token_abi		*abi.ABI
	charity_wallet_abi		*abi.ABI

	biddingwar_addr			common.Address
	cosmic_signature_addr	common.Address
	cosmic_token_addr		common.Address
	charity_wallet_addr		common.Address

	bw_contracts			BiddingWarContractAddrs
	storagew				SQLStorageWrapper
	RPC_URL					 = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	Error					*log.Logger
	Info					*log.Logger
)

func get_event_ids(from_evt_id,to_evt_id int64) []int64 {
	output := make([]int64 ,0, 1024)
	for _,e := range inspected_events {
		var event_list []int64
		event_list = storagew.S.Get_evtlogs_by_signature_only_in_range(
				e.Signature,from_evt_id,to_evt_id,
		)
		output = append(output,event_list...)
	}
	sort.Slice(output, func(i, j int) bool { return output[i] < output[j] })
	num_elts:=Remove_duplicates_int64(output)
	return output[0:num_elts]
}
func process_events(exit_chan chan bool) {

	var max_batch_size int64 = 1024*200
	for {
		status := storagew.Get_biddingwar_processing_status()
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		Info.Printf(
			"scanning event range from %v to %v\n",
			status.LastEvtIdProcessed,status.LastEvtIdProcessed+max_batch_size,
		)
		id_upper_limit := status.LastEvtIdProcessed + max_batch_size
		last_evt_id,err := storagew.S.Get_last_evtlog_id()
		if err != nil {
			Error.Printf("Error: %v. Possibly 'evt_log' table is empty, aborting",err)
			os.Exit(1)
		}
		if  id_upper_limit > last_evt_id {
			id_upper_limit = last_evt_id
		}
		events := get_event_ids(status.LastEvtIdProcessed,id_upper_limit)
		for _,evt_id := range events {
			err := process_single_event(evt_id)
			if err != nil {
				Error.Printf("Pausing event processing loop for 5 sec due to error")
				time.Sleep(5 * time.Second)
				break
			}
			status.LastEvtIdProcessed=evt_id
			storagew.Update_biddingwar_process_status(&status)
		}
		if len(events) == 0 {
			status.LastEvtIdProcessed = id_upper_limit
			storagew.Update_biddingwar_process_status(&status)
			time.Sleep(1 * time.Second) // sleep only if there is no data
		}
	}
}
func get_abi(abi_str string) *abi.ABI {
	abi_parsed := strings.NewReader(abi_str)
	abi_obj,err := abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse ABI: %v\n",err)
		os.Exit(1)
	}
	return &abi_obj
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/biddingwar_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/biddingwar_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/biddingwar_error.log",log_dir)
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

	storagew.S = Connect_to_storage(Info)
	storagew.S.Init_log(db_log_file)
	storagew.S.Log_msg("Log initialized\n")

	biddingwar_abi = get_abi(BiddingWarABI)
	cosmic_signature_abi = get_abi(CosmicSignatureABI)
	cosmic_token_abi = get_abi(CosmicSignatureTokenABI)
	charity_wallet_abi = get_abi(CharityWalletABI);

	bw_contracts = storagew.Get_biddingwar_contract_addrs()
	biddingwar_addr = common.HexToAddress(bw_contracts.BiddingWarAddr)
	cosmic_signature_addr = common.HexToAddress(bw_contracts.CosmicSignatureAddr)
	cosmic_token_addr = common.HexToAddress(bw_contracts.CosmicSignatureTokenAddr)
	charity_wallet_addr = common.HexToAddress(bw_contracts.CharityWalletAddr)

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()


	inspected_events = build_list_of_inspected_events_layer1()
	process_events(exit_chan)
}
