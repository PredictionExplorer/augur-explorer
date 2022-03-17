// MainNet Balancer v2 starting block: 12272146

package main

import (
	"os"
	"os/signal"
	"syscall"
	//"time"
	"encoding/hex"
	"strings"
	//"strconv"
	"fmt"
	"context"
	"log"
	"flag"
	//"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"

	//. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/layer1"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	POOL_CREATED				= "83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc"
	POOL_BALANCE_CHANGED		= "e5ce249087ce04f05a957192435400fd97868dba0e6a4b4c049abf8af80dae78"
	POOL_REGISTERED				= "3c13bc30b8e878c53fd2a36b679409c073afd75950be43d8858768e956fbc20e"
	EXTERNAL_BALANCE_TRANSFER	= "540a1a3f28340caec336c81d8d7b3df139ee5cdc1839a4f283d7ebb7eaae2d5c"
	INTERNAL_BALANCE_CHANGED	= "18e1ea4139e68413d7d08aa752e71568e36b2c5bf940893314c2c5b01eaa0c42"
	TOKEN_DEREGISTERED			= "7dcdc6d02ef40c7c1a7046a011b058bd7f988fa14e20a66344f9d4e60657d610"
	TOKEN_REGISTERED			= "f5847d3f2197b16cdcd2098ec95d0905cd1abdaf415f07bb7cef2bba8ac5dec4"
	SWAP						= "2170c741c41531aec20e7c107c24eecfdd15e69c9bb0a8dd37b1840b9e0b207b"

	DEFAULT_STATISTICS_DURATION	int64 = 24*60*60 // in seconds
	DEFAULT_WAIT_TIME = 2000	// 2 seconds
	DEFAULT_DB_LOG				= "db.log"
	//DEFAULT_LOG_DIR				= "ae_logs"
	MAX_APPROVAL_BASE10 string = "115792089237316195423570985008687907853269984665640564039457584007913129639935"
	NUM_AUGUR_CONTRACTS int = 35
	//USE_BLOCK_RECEIPTS_RPC_CALL bool = false // flag for using patch in ./geth-patch/README.txt
)
var (
	evt_pool_created,_ = hex.DecodeString(POOL_CREATED)
	evt_pool_balance_changed,_ = hex.DecodeString(POOL_BALANCE_CHANGED)
	evt_pool_registered,_ = hex.DecodeString(POOL_REGISTERED)
	evt_external_balance_transfer,_ = hex.DecodeString(EXTERNAL_BALANCE_TRANSFER)
	evt_internal_balance_changed,_ = hex.DecodeString(INTERNAL_BALANCE_CHANGED)
	evt_token_deregistered,_ = hex.DecodeString(TOKEN_DEREGISTERED)
	evt_token_registered,_ = hex.DecodeString(TOKEN_REGISTERED)
	evt_swap,_ = hex.DecodeString(SWAP)

	storage *SQLStorage

	eclient *ethclient.Client
	rpcclient *rpc.Client

	Error   *log.Logger
	Info	*log.Logger

	manager	ETL_Manager
	layer1 ETL_Layer1
	pool_factory_abi *abi.ABI
	vault_abi *abi.ABI

	processor	ETL_Processor
)
func main() {

	usage_str := fmt.Sprintf("usage: %v --schema [schema_name] --rpc [rpc_url] --blockrcpts [true|false]\n",os.Args[0])
	if len(os.Args)<4 {
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	schema_name := flag.String("schema", "", "Schema name")
	rpc_url := flag.String("rpc","","RPC URL")
	block_rcpts := flag.Bool("blockrcpts",false,"Use block receipts rpc call")
	block_num := flag.Int64("bnum",0,"Single block number to process")
	num_threads := flag.Int64("numthreads",1,"Number of parallel threads for block processing")
	flag.Parse()
	fmt.Printf("num_threads=%v\n",*num_threads)
	if len(*schema_name) < 3 {
		fmt.Printf("Schema name must be larger than 2 characters\n")
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	if len(*rpc_url) < 1 {
		fmt.Printf("RPC URL name must be non-empty\n")
		fmt.Printf("%v",usage_str)
		os.Exit(1)
	}
	var rLimit syscall.Rlimit
	rLimit.Max = 999999
	rLimit.Cur = 999999
	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Setting Rlimit ", err)
		os.Exit(1)
	}
	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	//db_log_file:=fmt.Sprintf("%v/%v_%v_%v",log_dir,etl.AppName,etl.SchemaName,DEFAULT_DB_LOG)

	layer1.UseBlockReceiptsCall = *block_rcpts
	layer1.SchemaName = *schema_name
	layer1.RPC_Url = *rpc_url
	layer1.AppName = "balancerv2"
	processor.ETL = &layer1
	layer1.Manager = &processor


	fname:=fmt.Sprintf("%v/%v_%v_info.log",log_dir,layer1.AppName,layer1.SchemaName)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/%v_%v_error.log",log_dir,layer1.AppName,layer1.SchemaName)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)
	layer1.Info = Info
	layer1.Error = Error

	Info.Printf("Selected schema name: %v\n",*schema_name)
	Info.Printf("Use our custom ethclient.GetBlockReceipts() call: %v\n",layer1.UseBlockReceiptsCall)
	rpcclient, err=rpc.DialContext(context.Background(), *rpc_url)
	if err != nil {
		log.Fatal(err)
	}
	layer1.Info.Printf("Connected to ETH node: %v\n",*rpc_url)
	eclient = ethclient.NewClient(rpcclient)

	db_log_file:=fmt.Sprintf("%v/%v_%v_%v",log_dir,layer1.AppName,layer1.SchemaName,DEFAULT_DB_LOG)
	storage = Connect_to_storage(Info)

	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")
	storage.Db_set_schema_name(*schema_name)
	layer1.Storage = storage

	ctx := context.Background()
	stored_chain_id := storage.Layer1_get_stored_chain_id()
	network_chain_id,err :=eclient.NetworkID(ctx)
	if err != nil {
		Fatalf("Can't get Network ID: %v\n",err)
	}
	if stored_chain_id != network_chain_id.Int64() {
		if stored_chain_id == 0 {
			// not initialized yet
			storage.Layer1_set_chain_id(network_chain_id.Int64())
		} else {
			Fatalf(
				"Network chain_id = %v , my chain_id = %v. Mismatch, exiting",
				network_chain_id.Int64(),stored_chain_id,
			)
		}
	}
	if *block_num > 0 {
		Info.Printf("Processing single block %v\n",*block_num)
		layer1.SingleBlockNum = *block_num
		layer1.UpdateLastBlock = false
		layer1.NoChainSplitCheck = true
		layer1.NoRollbackBlocks = true
		Process_single_block(&layer1)
		os.Exit(0)
	}
	//fmt.Printf("Forced exit");	os.Exit(1);

	c := make(chan os.Signal)
	exit_chan := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		Info.Printf("Got SIGINT signal, will exit after block processing is over." +
					" To interrupt abruptly send SIGKILL (9) to the kernel.\n")
		exit_chan <- true
	}()

	latestBlock, err := eclient.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("oops:", err)
	}

	abi_parsed1 := strings.NewReader(BalancerV2WeightedPoolFactoryABI)
	abi1,err := abi.JSON(abi_parsed1)
	if err!= nil {
		Info.Printf("Can't parse PoolFactory ABI: %v\n",err)
		os.Exit(1)
	}
	pool_factory_abi = &abi1

	abi_parsed2 := strings.NewReader(BalancerV2VaultABI)
	abi2,err := abi.JSON(abi_parsed2)
	if err!= nil {
		Info.Printf("Can't parse Vault ABI: %v\n",err)
		os.Exit(1)
	}
	vault_abi = &abi2

	bnum,exists := storage.Layer1_get_last_block_num()
	if !exists {
		bnum = 0
	} else {
		bnum = bnum + 1
	}
	var bnum_high int64 = latestBlock.Number().Int64()
	if bnum_high < bnum {
		Info.Printf("Database has more blocks than the blockchain, aborting. Fix last_block table.\n")
		os.Exit(1)
	}
	layer1.NumThreads = *num_threads
	Info.Printf("Num threads = %v\n",layer1.NumThreads)
	Init(&layer1)
	if *num_threads == 1 {
		Main_event_loop_single_thread(&layer1,exit_chan)
	} else {
		Main_event_loop_multithreaded(&layer1,exit_chan)
	}
}
