// Gets all transfers of the token , for verification purposes
package main
import (
	"os"
	"log"
	"fmt"
	"time"
	"math/big"
	"context"
	"encoding/hex"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	SCANNED_EVT=   "ad2bc79f659de022c64ef55c71f16d0cf125452ed5fc5757b2edc331f58565ec"
	MAX_BLOCKS	int64 = 1024*1024
	START_BLOCK int64 = 2000000
)
var (

	Info    *log.Logger
	RPC_URL string
	evt_scanned,_   = hex.DecodeString(SCANNED_EVT)
	rwalk_addr common.Address

	rpcclient *rpc.Client
	eclient *ethclient.Client
	storage *SQLStorage
)
func process_log(log *types.Log) {
	
	//fmt.Printf("processing log %+v\n",log)
	var evt RW_MintEvent
	evt.TokenId = log.Topics[1].Big().Int64()
	evt.Owner = common.BytesToAddress(log.Topics[2][12:]).String()
//	fmt.Printf(
//		"Block %v: RWMint token id %v  tx %v\n",
//		log.BlockNumber,evt.TokenId,log.TxHash.String(),
//	)
  again:
	exists,err := storage.Check_rwalk_token_exists(evt.TokenId)
	if err != nil {
		fmt.Printf("Error accessing database: %v\n",err)
		time.Sleep(1 * time.Second)
		goto again
	}
	if !exists {
		fmt.Printf("Token %v DOES NOT exist in the DB\n",evt.TokenId)
	}
//	fmt.Printf("Token %v\n",evt.TokenId)
}
func main() {


	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v [randomwalk_addr]\n",os.Args[0])
		os.Exit(1)
	}
	rwalk_addr = common.HexToAddress(os.Args[1])

	RPC_URL = os.Getenv("RPC_URL")

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}
	/*rwalk_addr := common.HexToAddress(caddrs.RandomWalk)
	rwalk_ctrct,err := NewRWalk(rwalk_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate RWalk contract: %v\n",err)
		os.Exit(1)
	}
	var copts bind.CallOpts
	*/
	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err!= nil {
		fmt.Printf("Error connecting to node: %v\n",err)
		os.Exit(1)
	}
	storage = Connect_to_storage(Info)
	ctx := context.Background()
	latestBlock, err := eclient.HeaderByNumber(ctx, nil)
	if err != nil {
		fmt.Printf("Error getting latest block: %v\n",err)
		os.Exit(1)
	}
	latest_bnum := latestBlock.Number.Int64()

	from_block := int64(START_BLOCK)
	to_block := int64(0)
	filter := ethereum.FilterQuery{}
	topics := make([]common.Hash,0,1)
	signature := common.BytesToHash(evt_scanned)
	topics = append(topics,signature)
	filter.Topics= append(filter.Topics,topics)
	addresses := make([]common.Address,0,1)
	addresses = append(addresses,rwalk_addr)
	filter.Addresses = addresses

	for ;from_block<=latest_bnum;from_block=from_block+MAX_BLOCKS {
		to_block = from_block + MAX_BLOCKS
		filter.FromBlock = big.NewInt(from_block)
		filter.ToBlock = big.NewInt(to_block)
		fmt.Printf("From %v , to %v\n",filter.FromBlock.Int64(),filter.ToBlock.Int64())
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
