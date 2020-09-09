package main

import (
	"os"
	"os/signal"
	"syscall"
//	"io/ioutil"
//	"strings"
//	"time"
//	"strconv"
	"fmt"
//	"context"
	"log"
//	"math/big"
//	"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/ethclient"
//	"github.com/ethereum/go-ethereum/accounts/abi"
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
)
func process_tokens(exit_chan chan bool,caddrs *ContractAddresses) {
	conf := storage.Get_token_etl_process_config()

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
		dai_events := storage.Get_token_transfers_batch(ERC20_TRANSFER,contract_aid,conf.LastIdDAI)
		for _,evt := range dai_events {
			Info.Printf("event = %+v\n",evt)
			eth_log := new(types.Log)
			_ = eth_log
		}
		Info.Printf("processed %v events\n",len(dai_events))
		os.Exit(1)
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
	process_tokens(exit_chan,&caddrs_obj)
}
