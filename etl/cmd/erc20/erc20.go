package main

import (
	"os"
	"os/signal"
	"syscall"
	"bytes"
	"time"
	"fmt"
	"log"
	"strings"
	"math/big"
	"context"
	"encoding/hex"
	//"encoding/json"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	DEFAULT_DB_LOG				= "db.log"
	ERC20_TRANSFER = "ddf252ad"
)
var (
	storage *SQLStorage
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
	Error   *log.Logger
	Info	*log.Logger
	market_order_id int64 = 0
	cash_abi *abi.ABI
	af_abi abi.ABI
	all_contracts map[string]interface{}
	caddrs *ContractAddresses
)
func proc_erc20_transfer(log *types.Log,block_num,tx_id,evtlog_id int64) {
	var mevt ETransfer
	if len(log.Topics)!=3 {
		Info.Printf("ERC20 transfer event is not compliant log.Topics!=3. Tx hash=%v\n",log.TxHash.String())
		return
	}
	mevt.From= common.BytesToAddress(log.Topics[1][12:])
	mevt.To= common.BytesToAddress(log.Topics[2][12:])
	start := time.Now()
	err := cash_abi.Unpack(&mevt,"Transfer",log.Data)
	if err != nil {
		Error.Printf("signature=%v\n",log.Topics[0].String())
		Error.Printf("address=%v\n",log.Address.String())
		Error.Printf("tx hash = %v\n",log.TxHash.String())
		Error.Printf("log.Data=%v, data len=%v\n",hex.EncodeToString(log.Data[:]),len(log.Data[:]))
		Error.Printf("Event ERC20_Transfer, decode error: %v",err)
	} else {
		Info.Printf("ERC20_Transfer event, contract %v (block=%v) :\n",
									log.Address.String(),log.BlockNumber)
		mevt.Dump(Info)
		storage.Process_ERC_token_transfer(&mevt,block_num,tx_id,evtlog_id)
	}
}
func process_erc20_tokens(exit_chan chan bool,contract_aids string) {

	var max_batch_size int64 = 256
	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}

		status := storage.Get_tok_process_status()
		id_upper_limit := status.LastEvtId + max_batch_size
		last_evt_id,err := storage.Get_last_evtlog_id()
		if err != nil {
			Error.Printf("Error: %v. Possibly 'evt_log' table is empty, aborting",err)
			os.Exit(1)
		}
		if  id_upper_limit > last_evt_id {
			id_upper_limit = last_evt_id
		}

		tok_events := storage.Get_evt_log_ids_by_signature_in_range(
			ERC20_TRANSFER,
			contract_aids,
			status.LastEvtId,
			id_upper_limit,
		)
		for _,evt := range tok_events {
			evtlog := storage.Get_event_log(evt.EvtId)
			var log types.Log
			rlp.DecodeBytes(evtlog.RlpLog,&log)
			log.Address.SetBytes(caddrs.Dai.Bytes())
			agtx := storage.Get_augur_transaction(evt.TxId)
			proc_erc20_transfer(&log,agtx,evt.EvtId)
			status.LastEvtId = evt.EvtId
			storage.Update_tok_process_status(&status)
		}
		Info.Printf("processed %v events\n",len(tok_events))
		if len(tok_events) == 0 {
			last_evt_log_id_on_chain,err := storage.Get_last_evtlog_id()
			if err != nil {
				Info.Printf("Error getting last event log id: %v\n",err)
				os.Exit(1)
			}
			if last_evt_log_id_on_chain > id_upper_limit {
				// only advance upper range if events within the range have filled id value space
				status.LastEvtId = id_upper_limit
				storage.Update_ens_proc_status(&status)
			}
			time.Sleep(1 * time.Second) // sleep only if there is no data
			return
		}
	}
}
func process_tokens(exit_chan chan bool,caddrs *ContractAddresses) {
	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.\n")
					os.Exit(0)
				}
			default:
		}
		dai_contract_aid:= storage.Lookup_or_create_address(caddrs.Dai.String(),0,0)
		rep_contract_aid := storage.Lookup_or_create_address(caddrs.Reputation.String(),0,0)
		_=dai_contract_aid
		_=rep_contract_aid
	//	process_erc20_tokens(fmt.Sprintf("%v,%v",dai_contract_aid,rep_contract_aid))
		process_afoundry_wrapper_created_events()
		process_erc20wrapped_sharetokens()
		process_ethusd_price_events(exit_chan)
		time.Sleep(1 * time.Second)
	}
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/tokens_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/tokens_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/tokens_error.log",log_dir)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()

	caddrs_obj,err := storage.Get_contract_addresses()
	if err!=nil {
		Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}
	caddrs = &caddrs_obj

	all_contracts = Load_all_artifacts("./abis/augur-artifacts-abi.json")

	cash_abi = Abi_from_artifacts(&all_contracts,"Cash")

	abi_parsed := strings.NewReader(AugurFoundryABI)
	af_abi,err = abi.JSON(abi_parsed)
	if err!= nil {
		Info.Printf("Can't parse Augur Foundry ABI: %v\n",err)
		os.Exit(1)
	}

	process_tokens(exit_chan,&caddrs_obj)
}
