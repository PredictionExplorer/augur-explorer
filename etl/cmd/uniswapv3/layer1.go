
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
	"github.com/ethereum/go-ethereum/common"

	//. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
	//. "github.com/PredictionExplorer/augur-explorer/primitives/uniswapv3"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/dbs/uniswapv3"
	. "github.com/PredictionExplorer/augur-explorer/layer1"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	INITIALIZE					= "98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95"
	MINT						= "7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde"
	COLLECT						= "70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0"
	BURN						= "0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c"
	SWAP						= "c42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"
	FLASH						= "bdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633"
	INCREASE_OBSERV_CARDIN_NEXT = "ac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a"
	SET_FEE_PROTOCOL			= "973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133"
	COLLECT_PROTOCOL			= "596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151"
	OWNER_CHANGED				= "b532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c"
	POOL_CREATED				= "783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118"
	FEE_AMOUNT_ENABLED			= "c66a3fdf07232cdd185febcc6579d408c241b47ae2f9907d84be655141eeaecc"
	ERC20_TRANSFER				= "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	COLLECT_PERIFERY			= "40d0efd1a53d60ecbf40971b9daf7dc90178c3aadc7aab1765632738fa8b8f01"
	INCREASE_LIQUIDITY			= "3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f"
	DECREASE_LIQUIDITY			= "26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4"

	DEFAULT_STATISTICS_DURATION	int64 = 24*60*60 // in seconds
	DEFAULT_WAIT_TIME = 2000	// 2 seconds
	DEFAULT_DB_LOG				= "db.log"
	MAX_APPROVAL_BASE10 string = "115792089237316195423570985008687907853269984665640564039457584007913129639935"
	NUM_AUGUR_CONTRACTS int = 35
)

var (
	evt_initialize,_ = hex.DecodeString(INITIALIZE)
	evt_mint,_ = hex.DecodeString(MINT)
	evt_collect,_ = hex.DecodeString(COLLECT)
	evt_burn,_ = hex.DecodeString(BURN)
	evt_swap,_ = hex.DecodeString(SWAP)
	evt_flash,_ = hex.DecodeString(FLASH)
	evt_inc_obs_cardin_next,_ = hex.DecodeString(INCREASE_OBSERV_CARDIN_NEXT)
	evt_set_fee_protocol,_ = hex.DecodeString(SET_FEE_PROTOCOL)
	evt_collect_protocol,_ = hex.DecodeString(COLLECT_PROTOCOL)
	evt_owner_changed,_ = hex.DecodeString(OWNER_CHANGED)
	evt_pool_created,_ = hex.DecodeString(POOL_CREATED)
	evt_fee_amount_enabled,_ = hex.DecodeString(FEE_AMOUNT_ENABLED)
	evt_erc20_transfer,_ = hex.DecodeString(ERC20_TRANSFER)
	evt_collect_perifery,_ = hex.DecodeString(COLLECT_PERIFERY)
	evt_increase_liquidity,_ = hex.DecodeString(INCREASE_LIQUIDITY)
	evt_decrease_liquidity,_ = hex.DecodeString(DECREASE_LIQUIDITY)

	eclient *ethclient.Client
	rpcclient *rpc.Client

	Error   *log.Logger
	Info	*log.Logger

	manager	ETL_Manager
	layer1 ETL_Layer1
	factory_abi *abi.ABI
	pool_abi *abi.ABI
	nfpm_abi *abi.ABI
	erc20_abi *abi.ABI

	caddrs		UniswapV3Addrs

	processor	ETL_Processor
	storagew            SQLStorageWrapper
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
	layer1.AppName = "uniswapv3"
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
	storagew.S = Connect_to_storage_with_schema(Info,*schema_name)

	storagew.S.Init_log(db_log_file)
	storagew.S.Log_msg("Log initialized\n")
	spath := storagew.S.Get_search_path()
	Info.Printf("search path : %v\n",spath)
	layer1.Storage = storagew.S

	ctx := context.Background()
	stored_chain_id := storagew.S.Layer1_get_stored_chain_id()
	network_chain_id,err :=eclient.NetworkID(ctx)
	if err != nil {
		Fatalf("Can't get Network ID: %v\n",err)
	}
	if stored_chain_id != network_chain_id.Int64() {
		if stored_chain_id == 0 {
			// not initialized yet
			storagew.S.Layer1_set_chain_id(network_chain_id.Int64())
		} else {
			Fatalf(
				"Network chain_id = %v , my chain_id = %v. Mismatch, exiting",
				network_chain_id.Int64(),stored_chain_id,
			)
		}
	}
	tmp_caddrs := storagew.Uniswap_get_contract_addrs()
	caddrs.FactoryAddr = common.HexToAddress(tmp_caddrs.FactoryAddr)
	caddrs.NFTPosMgrAddr = common.HexToAddress(tmp_caddrs.NFTPosMgrAddr)

	if *block_num > 0 {
		Info.Printf("Processing single block %v\n",*block_num)
		layer1.SingleBlockNum = *block_num
		layer1.UpdateLastBlock = false
		layer1.NoChainSplitCheck = true
		layer1.NoRollbackBlocks = true
		Process_single_block(&layer1,eclient,rpcclient)
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

	abi_parsed1 := strings.NewReader(UniswapV3FactoryABI)
	abi1,err := abi.JSON(abi_parsed1)
	if err!= nil {
		Info.Printf("Can't parse UniswapV2Factory ABI: %v\n",err)
		os.Exit(1)
	}
	factory_abi = &abi1

	abi_parsed2 := strings.NewReader(UniswapV3PoolABI)
	abi2,err := abi.JSON(abi_parsed2)
	if err!= nil {
		Info.Printf("Can't parse Uniswap Pool ABI: %v\n",err)
		os.Exit(1)
	}
	pool_abi = &abi2

	abi_parsed3 := strings.NewReader(ERC20ABI)
	abi3,err := abi.JSON(abi_parsed3)
	if err != nil {
		Info.Printf("Can't parse ERC20 token ABI")
		os.Exit(1)
	}
	erc20_abi = &abi3

	abi_parsed4 := strings.NewReader(NonfungiblePositionManagerABI)
	abi4,err := abi.JSON(abi_parsed4)
	if err != nil {
		Info.Printf("Can't parse NonFunbiglePositionManager contract ABI")
		os.Exit(1)
	}
	nfpm_abi = &abi4

	bnum,exists := storagew.S.Layer1_get_last_block_num()
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
	Init(&layer1,eclient,rpcclient)
	if *num_threads == 1 {
		Main_event_loop_single_thread(&layer1,exit_chan)
	} else {
		Main_event_loop_multithreaded(&layer1,exit_chan)
	}
}
