// Gets all transfers of the token , for verification purposes
package main
import (
	"os"
	"log"
	"fmt"
	"math/big"
	"context"
	"encoding/hex"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
const (
	TRANSFER_EVT=   "c42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"
	MAX_BLOCKS	int64 = 1024*1
	USDC
)
var (

	Info    *log.Logger
	RPC_URL string
	evt_transfer,_   = hex.DecodeString(TRANSFER_EVT)
	rwalk_addr common.Address

	rpcclient *rpc.Client
	eclient *ethclient.Client
)
func process_log(log *types.Log) {

	var evt IUniswapV3PoolSwap
	evt.Sender = common.BytesToAddress(log.Topics[1][12:])
	evt.Recipient = common.BytesToAddress(log.Topics[2][12:])

}
func main() {


	RPC_URL = os.Getenv("RPC_URL")

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)
	db_log_file:=fmt.Sprintf("%v/ethprice_%v",log_dir,DEFAULT_DB_LOG)

	fname:=fmt.Sprintf("%v/ethprice_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ltime|log.Lshortfile)

	fname=fmt.Sprintf("%v/randomwalk_error.log",log_dir)
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Error = log.New(logfile,"ERROR: ",log.Ltime|log.Lshortfile)

	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}
	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err!= nil {
		fmt.Printf("Error connecting to node: %v\n",err)
		os.Exit(1)
	}
	ctx := context.Background()
	starting_block := storage.Ethprice_get_last_processed_block()

	from_block := int64(starting_block)
	to_block := int64(0)
	filter := ethereum.FilterQuery{}
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_transfer)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	addresses := make([]common.Address,0,1)
	addresses = append(addresses,rwalk_addr)
	filter.Addresses = addresses

	for ;from_block<=latest_bnum;from_block=from_block+MAX_BLOCKS {
		to_block = from_block + MAX_BLOCKS
		filter.FromBlock = big.NewInt(from_block)
		filter.ToBlock = big.NewInt(to_block)
		logs,err := eclient.FilterLogs(context.Background(),filter)
		if err != nil {
			fmt.Printf("Error querying events: %v\n",err)
			os.Exit(1)
		}
		for _,log := range logs {
			if log.Removed {
				continue
			} else {
				process_log(&log)
			}
		}
	}
}
