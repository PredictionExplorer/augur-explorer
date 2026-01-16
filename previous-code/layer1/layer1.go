package layer1

import (
	"os"
	"io/ioutil"
	"strings"
	"time"
	"strconv"
	"fmt"
	"context"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
const (

	WAIT_TIME = 2000	// 2 seconds
)
var (

	eclient *ethclient.Client
	rpcclient *rpc.Client


)
func read_block_numbers(fname string) []int64 {
	data,err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("Can't open file %v containing comma-separated block numbers to be processed\n")
		os.Exit(1)
	}
	blocks_str := string(data)
	numbers := strings.Split(blocks_str,",")
	output := make([]int64,0, 512)
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
func multi_threaded_loop_routine(etl *ETL_Layer1,retval *int64,wg_ptr *sync.WaitGroup,num_threads,tid int64,stop_chan chan bool, bnum_low,bnum_high int64) {
	// tid - thread id
	// every go routine will attend a sharded block number
	// shards are specified by 'block_num' on the commandline

	// Notes: multithreaded routine is not calculating statistics, it is meant to 
	//			accelerate initial load, until the point you reach the last block in the chain, and
	//			then you switch to single threaded routine

	last_block_num := int64(0)
	bnum := bnum_low
	etl.Info.Printf("Thread %v , processing blocks from %v to %v\n",tid,bnum,bnum_high)

	for ; bnum<bnum_high; bnum=bnum+num_threads{
		select {
			case exit_flag := <-stop_chan:
				if exit_flag {
					etl.Info.Printf("Thread %v is exiting (set last block to %v)",tid,last_block_num)
					*retval=last_block_num
					wg_ptr.Done()
					return
				}
			default:
		}
		etl.Info.Printf("Thread %v : processing block %v\n",tid,bnum)
		// note: we do not update last_block table in multithreaded run, we do it in main routine
		update_last_block := false
		no_chainsplit_check := true
		err := process_block(etl,bnum,update_last_block,no_chainsplit_check,etl.NoRollbackBlocks)
		if err!=nil {
			etl.Error.Printf(
				"Multithreaded block processing error at block $v : %v. Aborting\n",
				bnum,err,
			)
			etl.Error.Printf("Update last_block manually (irregular exit)\n")
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
func single_threaded_loop_routine(etl *ETL_Layer1,exit_chan chan bool) {
  main_loop:

	select {
		case exit_flag := <-exit_chan:
			if exit_flag {
				etl.Info.Println("Exiting")
				os.Exit(0)
			}
		default:
	}
	  latestBlock, err := eclient.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("oops:", err)
	}
	bnum_high := latestBlock.Number().Int64()

	bnum,exists := etl.Storage.Layer1_get_last_block_num()
	cur_bnum := bnum
	if !exists {
		bnum = 0
		etl.Info.Printf("DB is empty, starting from block 0\n")
	} else {
		bnum = bnum + 1
	}
	if (bnum > etl.EndingBlock) && (etl.EndingBlock > 0) {
		etl.Info.Printf("Reached specified ending block (num=%v), finishing process\n",etl.EndingBlock)
		os.Exit(0)
	}
	bnum_high = latestBlock.Number().Int64()
	if bnum_high < cur_bnum {
		etl.Info.Printf("Database has more blocks than the blockchain, aborting. Sleeping to wait\n")
		time.Sleep(10 * time.Second)
		goto main_loop
	}
	for ; bnum<=bnum_high; bnum++ {
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					etl.Info.Println("Exiting")
					os.Exit(0)
				}
			default:
		}
		err := process_block(etl,bnum,true,false,etl.NoRollbackBlocks)
		if err!=nil {
			// this is probably happening due to RPC unavailability, so we use a delay
			time.Sleep(1 * time.Second)
			etl.Error.Printf("Block processing error: %v\n",err)
			break
		}
	}// for block_num
	time.Sleep(WAIT_TIME * time.Millisecond)
	goto main_loop // infinite loop without loss of one indentation level
}
func Process_single_block(etl *ETL_Layer1,ec *ethclient.Client,	rc *rpc.Client) {

	Validate_params(etl)
	Init(etl,ec,rc)
	etl.Info.Printf("Processing single block %v\n")
	process_block(etl,etl.SingleBlockNum,etl.UpdateLastBlock,etl.NoChainSplitCheck,etl.NoRollbackBlocks)
}
func Main_event_loop_single_thread(etl *ETL_Layer1,exit_chan chan bool) {

	etl.Info.Printf("Thread: single\n")
	latestBlock, err := eclient.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("oops:", err)
	}

	bnum,exists := etl.Storage.Layer1_get_last_block_num()
	cur_bnum := bnum
	if !exists {
		bnum = 0
	} else {
		bnum = bnum + 1
	}
	var bnum_high int64 = latestBlock.Number().Int64()
	if bnum_high < cur_bnum {
		etl.Info.Printf("Database has more blocks than the blockchain, aborting. Fix last_block table.\n")
		os.Exit(1)
	}

	single_threaded_loop_routine(etl,exit_chan)

}
func Validate_params(etl *ETL_Layer1) {

	if len(etl.SchemaName) < 3 {
		fmt.Printf("Schema name must be larger than 2 characters\n")
		os.Exit(1)
	}
	if len(etl.RPC_Url) < 1 {
		fmt.Printf("RPC URL name must be non-empty\n")
		os.Exit(1)
	}
	if etl.NumThreads < 1 {
		fmt.Printf("Error: num_threads variable must be greater than 1\n")
		os.Exit(1)
	}
}
func Init(etl *ETL_Layer1,ec *ethclient.Client,	rc *rpc.Client) {

	Validate_params(etl)

	var err error
	rpcclient = rc
	eclient = ec

	ctx := context.Background()
	stored_chain_id := etl.Storage.Layer1_get_stored_chain_id()
	network_chain_id,err :=eclient.NetworkID(ctx)
	if err != nil {
		Fatalf("Can't get Network ID: %v\n",err)
	}
	if stored_chain_id != network_chain_id.Int64() {
		if stored_chain_id == 0 {
			// not initialized yet
			etl.Storage.Layer1_set_chain_id(network_chain_id.Int64())
		} else {
			Fatalf(
				"Network chain_id = %v , my chain_id = %v. Mismatch, exiting",
				network_chain_id.Int64(),stored_chain_id,
			)
		}
	}

	etl.NoRollbackBlocks = false
	if etl.NumThreads > 1 {
		etl.NoRollbackBlocks = true // with multiple threds we can't recover from chainsplit
	}
	if etl.SingleBlockNum > 0 {
		etl.NoRollbackBlocks = true
	}
}
func Main_event_loop_multithreaded(etl *ETL_Layer1,exit_chan chan bool) {

	etl.Info.Printf("Thread: multithreaded, num_threads=%v\n",etl.NumThreads)
	latestBlock, err := eclient.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("oops:", err)
	}
	bnum,exists := etl.Storage.Layer1_get_last_block_num()
	cur_bnum := bnum
	if !exists {
		bnum = 0
	} else {
		bnum = bnum + 1
	}
	var bnum_high int64 = latestBlock.Number().Int64()
	if bnum_high < cur_bnum {
		etl.Info.Printf("Database has more blocks than the blockchain, aborting. Fix last_block table.\n")
		os.Exit(1)
	}
	stop_chans := make([]chan bool,etl.NumThreads)
	return_values := make([]int64,etl.NumThreads)
	for i:=int64(0);i<etl.NumThreads;i++ {
		// we need only 1 but we use 2 to have extra space for safety (to avoid deadlocks)
		stop_chans[i]=make(chan bool,2)
	}
	// send child threads
	var wg sync.WaitGroup
	for tid:=int64(0);tid<etl.NumThreads;tid++ {
		wg.Add(1)
		bnum_low := bnum + tid
		go multi_threaded_loop_routine(
			etl,
			&return_values[tid],
			&wg,
			etl.NumThreads,
			tid,
			stop_chans[tid],
			bnum_low,
			bnum_high,
		)
	}
	// this is the main routine (controller)
	for {
		time.Sleep(5*time.Second)	// reasonable delay
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					etl.Info.Printf("Sending 'exit' messages to all %v threads",etl.NumThreads)
					for i:=int64(0);i<etl.NumThreads;i++ {
						stop_chans[i] <- true
					}
				}
				wg.Wait()
				etl.Info.Printf("All threads exited\n")
				etl.Info.Printf("Updating last block\n")
				const min_block_const int64 = 1000000000000
				var min_block int64 = min_block_const
				for i:=int64(0);i<etl.NumThreads;i++ {
					etl.Info.Printf("retval for %v is %v\n",i,return_values[i])
					if min_block > return_values[i] {
						min_block = return_values[i]
						etl.Info.Printf("set min_block to %v\n",min_block)
					}
				}
				etl.Info.Printf("loop exited, min_block=%v\n",min_block)
				if min_block != min_block_const {
					etl.Info.Printf("Set last block number to %v\n",min_block)
					etl.Storage.Layer1_set_last_block_num(min_block)
				}
				etl.Info.Printf("Exiting main thread\nBye\n")
				os.Exit(1)
			default:
		}
	}
}
