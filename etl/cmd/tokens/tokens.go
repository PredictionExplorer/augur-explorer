package main

import (
	"os"
	"os/signal"
	"syscall"
	"bytes"
//	"io/ioutil"
//	"strings"
	"time"
//	"strconv"
	"fmt"
//	"context"
	"log"
//	"math/big"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
//	"github.com/ethereum/go-ethereum/rpc"

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
	Error   *log.Logger
	Info	*log.Logger
	market_order_id int64 = 0
	cash_abi *abi.ABI
	all_contracts map[string]interface{}
	caddrs *ContractAddresses
)
func proc_erc20_transfer(log *types.Log,agtx *AugurTx,evtlog_id int64) {
	var mevt ETransfer
	if len(log.Topics)!=3 {
		Info.Printf("ERC20 transfer event is not compliant log.Topics!=3. Tx hash=%v\n",log.TxHash.String())
		return
	}
	mevt.From= common.BytesToAddress(log.Topics[1][12:])
	mevt.To= common.BytesToAddress(log.Topics[2][12:])
	start := time.Now()
	err := cash_abi.Unpack(&mevt,"Transfer",log.Data)
	duration := time.Since(start)
	Info.Printf("BENCH cash_abi.Unpack() took %v micrsec\n",duration.Microseconds())
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
		if bytes.Equal(caddrs.Dai.Bytes(), log.Address.Bytes()) {	// this is DAI contract
			storage.Delete_DAI_transfer_by_evtlog_id(evtlog_id)
			storage.Process_DAI_token_transfer(&mevt,caddrs,agtx,evtlog_id)
		}
		
		if bytes.Equal(caddrs.Reputation.Bytes(), log.Address.Bytes()) {	// this is DAI contract
			storage.Process_REP_token_transfer(&mevt,agtx,evtlog_id)
		}
		
	}
}
func process_erc20_tokens(contract_aid int64) {

	for {
		status := storage.Get_dai_process_status()
		start1 := time.Now()
	//		dai_events := storage.Get_token_transfers_batch(ERC20_TRANSFER,contract_aid,conf.LastIdDAI)
		dai_events := storage.Get_evt_logs_by_signature(ERC20_TRANSFER,contract_aid,status.LastBlock,256)
		duration1 := time.Since(start1)
		Info.Printf("BENCH Get_token_transfers_batch() took %v milliseconds\n",duration1.Milliseconds())
		for _,evt := range dai_events {
			Info.Printf("event = %+v\n",evt)
			start2 := time.Now()
			evtlog := storage.Get_event_log(evt.EvtId)
			var log types.Log
			rlp.DecodeBytes(evtlog.RlpLog,&log)
			duration2 := time.Since(start2)
			Info.Printf("BENCH Get_evnt_log() took %v microsec\n",duration2.Microseconds())
			log.Address.SetBytes(caddrs.Dai.Bytes())
			agtx := storage.Get_augur_transaction(evt.TxId)
			proc_erc20_transfer(&log,agtx,evt.EvtId)
			status.LastBlock=agtx.BlockNum
		}
		storage.Update_dai_process_status(&status)
		Info.Printf("processed %v events\n",len(dai_events))
		if len(dai_events) == 0 {
			return
		}
	}
}
func process_tokens(exit_chan chan bool,caddrs *ContractAddresses) {
	for {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		contract_aid,err := storage.Nonfatal_lookup_address_id(caddrs.Dai.String())
		if err != nil {
			Error.Printf("DAI contract is not in the 'address' table")
			os.Exit(1)
		}
		process_erc20_tokens(contract_aid)
		//contract_aid,err = storage.Nonfatal_lookup_address_id(caddrs.Reputation.String())
		contract_aid= storage.Lookup_or_create_address(caddrs.Reputation.String(),0,0)
		if err != nil {
			Error.Printf("REPv2 contract is not in the 'address' table")
			os.Exit(1)
		}
		process_erc20_tokens(contract_aid)
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

	process_tokens(exit_chan,&caddrs_obj)
}
