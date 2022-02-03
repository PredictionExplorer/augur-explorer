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

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (

	DEFAULT_STATISTICS_DURATION	int64 = 15*60 // in seconds
	DEFAULT_WAIT_TIME = 2000	// 2 seconds
	DEFAULT_DB_LOG				= "db.log"
	//DEFAULT_LOG_DIR				= "ae_logs"
	MAX_APPROVAL_BASE10 string = "115792089237316195423570985008687907853269984665640564039457584007913129639935"
	NUM_AUGUR_CONTRACTS int = 35
	USE_BLOCK_RECEIPTS_RPC_CALL bool = false // flag for using patch in ./geth-patch/README.txt
)
var (
	storage *SQLStorage

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")

	eclient *ethclient.Client
	rpcclient *rpc.Client

	market_order_id int64 = 0
	owner_fld_offset int64 = int64(OWNER_FIELD_OFFSET)	// offset to AugurContract::owner field obtained with eth_getStorage()

	Error   *log.Logger
	Info	*log.Logger

	//DISCONTINUED ErrChainSplit error = errors.New("Chainsplit detected")
	split_simulated bool = false

	max_approval *big.Int = big.NewInt(0)

)
type rpcBlockHash struct {
	Hash		string
}
func read_block_numbers(fname string)  []int64 {
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
func main() {

	schema_name := flag.String("schema", "", "Schema name")
	flag.Parse()
	if len(*schema_name) < 3 {
		fmt.Printf("Schema name must be larger than 2 characters\n")
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

	if len(RPC_URL) == 0 {
		Fatalf("Configuration error: RPC URL of Ethereum node is not set."+
				" Please set AUGUR_ETH_NODE_RPC environment variable")
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
	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Init_log(db_log_file)
	storage.Log_msg("Log initialized\n")
	storage.Set_schema_name(*schema_name)

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
  main_loop:
	latestBlock, err = eclient.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("oops:", err)
	}

	bnum,exists = storage.Bigstats_get_last_block_num()
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
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		err := process_block(bnum,true,false)
		if err!=nil {
			// this is probably happening due to RPC unavailability, so we use a delay
			time.Sleep(1 * time.Second)
			Error.Printf("Block processing error: %v\n",err)
			break
		}
		manage_stat_periods(storage,DEFAULT_STATISTICS_DURATION)
	}// for block_num
	time.Sleep(DEFAULT_WAIT_TIME * time.Millisecond)
	goto main_loop // infinite loop without loss of one indentation level
}
