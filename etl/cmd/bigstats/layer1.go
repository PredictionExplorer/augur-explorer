// Augur ETL: Converts Augur Data to SQL database

package main

import (
	"os"
	"os/signal"
	"syscall"
	"io/ioutil"
	"strings"
	"time"
	"strconv"
	"fmt"
	"context"
	"log"
	"math/big"
	"flag"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	//. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (

	DEFAULT_STATISTICS_DURATION	int64 = 24*60*60 // in seconds
	DEFAULT_WAIT_TIME = 2000	// 2 seconds
	DEFAULT_DB_LOG				= "db.log"
	//DEFAULT_LOG_DIR				= "ae_logs"
	MAX_APPROVAL_BASE10 string = "115792089237316195423570985008687907853269984665640564039457584007913129639935"
	NUM_AUGUR_CONTRACTS int = 35
	//USE_BLOCK_RECEIPTS_RPC_CALL bool = false // flag for using patch in ./geth-patch/README.txt
)
var (
	storage *SQLStorage

	//RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")

	eclient *ethclient.Client
	rpcclient *rpc.Client

	market_order_id int64 = 0
	owner_fld_offset int64 = int64(OWNER_FIELD_OFFSET)	// offset to AugurContract::owner field obtained with eth_getStorage()

	Error   *log.Logger
	Info	*log.Logger

	//DISCONTINUED ErrChainSplit error = errors.New("Chainsplit detected")
	split_simulated bool = false

	max_approval *big.Int = big.NewInt(0)
	use_block_receipts_call	*bool = nil

)
type rpcBlockHash struct {
	Hash		string
}
func read_block_numbers(fname string) []int64 {
	data,err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("Can't open file %v containing comma-separated block numbers to be processed\n")
		os.Exit(1)
	}
	blocks_str := string(data)
	numbers := strings.Split(blocks_str,",")
	output := make([]int64,0,512)
	for i:=0 ; i<len(numbers); i++ {
		trimmed:=strings.ReplaceAll(numbers[i],"\n","")
		bnum,err:=strconv.Atoi(trimmed)
		if err!=nil {
			fmt.Printf("Can't convert block %v to number: %v . Aborting\n",numbers[i],err)
			os.Exit(1)
		}
		output = append(output,int64(bnum))
	}
	return output
}
func multi_threaded_loop_routine(retval *int64,wg_ptr *sync.WaitGroup,num_threads,tid int64,stop_chan chan bool, bnum_low,bnum_high int64,no_rollback_blocks bool) {
	// tid - thread id
	// every go routine will attend a sharded block number
	// shards are specified by 'block_num' on the commandline

	// Notes: multithreaded routine is not calculating statistics, it is meant to 
	//			accelerate initial load, until the point you reach the last block in the chain, and
	//			then you switch to single threaded routine

	last_block_num := int64(0)
	bnum := bnum_low
	Info.Printf("Thread %v , processing blocks from %v to %v\n",tid,bnum,bnum_high)

	for ; bnum<bnum_high; bnum=bnum+num_threads{
		select {
			case exit_flag := <-stop_chan:
				if exit_flag {
					Info.Printf("Thread %v is exiting (set last block to %v)",tid,last_block_num)
					*retval=last_block_num
					wg_ptr.Done()
					return
				}
			default:
		}
		Info.Printf("Thread %v : processing block %v\n",tid,bnum)
		// note: we do not update last_block table in multithreaded run, we do it in main routine
		update_last_block := false
		no_chainsplit_check := true
		err := process_block(bnum,update_last_block,no_chainsplit_check,no_rollback_blocks)
		if err!=nil {
			Error.Printf("Block processing error: %v. Aborting\n",err)
			Error.Printf("Update last_block manually (irregular exit)\n")
			os.Exit(1)
		} else {
			last_block_num = bnum
		}
	}// for block_num
	// if the thread has ended, we just lock for 'read' operation and wait for main thread
	//		to notify us to exit
	*retval = last_block_num
	exit_flag := <-stop_chan
	_ = exit_flag
	wg_ptr.Done()
}
func single_threaded_loop_routine(exit_chan chan bool, no_rollback_blocks bool) {
  main_loop:

	  latestBlock, err := eclient.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("oops:", err)
	}
	bnum_high := latestBlock.Number().Int64()

	bnum,exists := storage.Bigstats_get_last_block_num()
	if !exists {
		bnum = 0
	} else {
		bnum = bnum + 1
	}
	bnum_high = latestBlock.Number().Int64()
	Info.Printf("Latest block=%v, bnum=%v\n",latestBlock.Number().Int64(),bnum)
	if bnum_high < bnum {
		Info.Printf("Database has more blocks than the blockchain, aborting. Sleeping to wait\n")
		time.Sleep(10 * time.Second)
		goto main_loop
	}
	for ; bnum<bnum_high; bnum++ {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting")
					os.Exit(0)
				}
			default:
		}
		err := process_block(bnum,true,false,no_rollback_blocks)
		if err!=nil {
			// this is probably happening due to RPC unavailability, so we use a delay
			time.Sleep(1 * time.Second)
			Error.Printf("Block processing error: %v\n",err)
			break
		}
		mod := bnum % 1000
		if mod == 0 {
			manage_stat_periods(storage,"ethprice",DEFAULT_STATISTICS_DURATION)
		}
	}// for block_num
	time.Sleep(DEFAULT_WAIT_TIME * time.Millisecond)
	goto main_loop // infinite loop without loss of one indentation level
}
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
	num_threads := flag.Int64("num_threads",1,"Number of parallel threads for block processing")
	close_periods := flag.Bool("closeperiods",false,"Close periods and exit")
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
	use_block_receipts_call	= block_rcpts

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
	db_log_file:=fmt.Sprintf("%v/bigstats_%v_%v",log_dir,*schema_name,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/bigstats_%v_info.log",log_dir,*schema_name)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/bigstats_%v_error.log",log_dir,*schema_name)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	max_approval.SetString(MAX_APPROVAL_BASE10,10)

	Info.Printf("Selected schema name: %v\n",*schema_name)
	Info.Printf("Use our custom ethclient.GetBlockReceipts() call: %v\n",*use_block_receipts_call)
	rpcclient, err=rpc.DialContext(context.Background(), *rpc_url)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",*rpc_url)
	eclient = ethclient.NewClient(rpcclient)

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")
	storage.Db_set_schema_name(*schema_name)

	ctx := context.Background()
	stored_chain_id := storage.Bigstats_get_stored_chain_id()
	network_chain_id,err :=eclient.NetworkID(ctx)
	if err != nil {
		Fatalf("Can't get Network ID: %v\n",err)
	}
	if stored_chain_id != network_chain_id.Int64() {
		if stored_chain_id == 0 {
			// not initialized yet
			storage.Bigstats_set_chain_id(network_chain_id.Int64())
		} else {
			Fatalf(
				"Network chain_id = %v , my chain_id = %v. Mismatch, exiting",
				network_chain_id.Int64(),stored_chain_id,
			)
		}
	}
	if *close_periods {
		Info.Printf("Closing periods only\n")
		manage_stat_periods(storage,"ethprice",DEFAULT_STATISTICS_DURATION)
		Info.Printf("Done closing periods, exiting\n")
		os.Exit(0)
	}
	if *block_num > 0 {
		Info.Printf("Processing single block %v\n",*block_num)
		process_block(*block_num,false,true,true)
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

	bnum,exists := storage.Bigstats_get_last_block_num()
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
	no_rollback_blocks := false
	if *num_threads > 1 {
		no_rollback_blocks = true // with multiple threds we can't recover from chainsplit
	}
	if *num_threads < 1 {
		fmt.Printf("Error: num_threads variable must be greater than 1\n")
		os.Exit(1)
	}
	if *num_threads == 1 {
		single_threaded_loop_routine(exit_chan,no_rollback_blocks)
	} else {
		latestBlock, err := eclient.BlockByNumber(context.Background(), nil)
		if err != nil {
			log.Fatal("oops:", err)
		}
		bnum_high := latestBlock.Number().Int64()
		stop_chans := make([]chan bool,*num_threads)
		return_values := make([]int64,*num_threads)
		for i:=int64(0);i<*num_threads;i++ {
			// we need only 1 but we use 2 to have extra space for safety (to avoid deadlocks)
			stop_chans[i]=make(chan bool,2)
		}
		// send child threads
		var wg sync.WaitGroup
		for tid:=int64(0);tid<*num_threads;tid++ {
			wg.Add(1)
			bnum_low := bnum + tid
			go multi_threaded_loop_routine(
				&return_values[tid],
				&wg,
				*num_threads,
				tid,
				stop_chans[tid],
				bnum_low,
				bnum_high,
				no_rollback_blocks,
			)
		}
		// this is the main routine (controller)
		for {
			time.Sleep(5*time.Second)	// reasonable delay
			select {
				case exit_flag := <-exit_chan:
					if exit_flag {
						Info.Printf("Sending 'exit' messages to all %v threads",*num_threads)
						for i:=int64(0);i<*num_threads;i++ {
							stop_chans[i] <- true
						}
					}
					wg.Wait()
					Info.Printf("All threads exited\n")
					Info.Printf("Updating last block\n")
					const min_block_const int64 = 1000000000000
					var min_block int64 = min_block_const
					for i:=int64(0);i<*num_threads;i++ {
						Info.Printf("retval for %v is %v\n",i,return_values[i])
						if min_block > return_values[i] {
							min_block = return_values[i]
							Info.Printf("set min_block to %v\n",min_block)
						}
					}
					Info.Printf("loop exited, min_block=%v\n",min_block)
					if min_block != min_block_const {
						Info.Printf("Set last block number to %v\n",min_block)
						storage.Bigstats_set_last_block_num(min_block)
					}
					Info.Printf("Exiting main thread\nBye\n")
					os.Exit(1)
				default:
			}
		}
	}
}
