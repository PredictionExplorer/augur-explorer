// Augur ETL: Converts Augur Data to SQL database
// Notes:
//		Arbitrum starting block: 217636
//		MainNet starting block: 13000000
package layer1

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
	. "github.com/PredictionExplorer/augur-explorer/layer1"
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

	eclient *ethclient.Client
	rpcclient *rpc.Client

	Error   *log.Logger
	Info	*log.Logger

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
	num_threads := flag.Int64("num_threads",1,"Number of parallel threads for block processing")
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

	Info.Printf("Selected schema name: %v\n",*schema_name)
	Info.Printf("Use our custom ethclient.GetBlockReceipts() call: %v\n",*use_block_receipts_call)
	rpcclient, err=rpc.DialContext(context.Background(), *rpc_url)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",*rpc_url)
	eclient = ethclient.NewClient(rpcclient)

	storage = Connect_to_storage(Info)
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
