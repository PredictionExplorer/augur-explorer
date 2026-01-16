// Gets all transfers of the token , for verification purposes
package main
import (
	"os"
	"log"
	"fmt"
	"strconv"
	"math/big"
	"context"
	"encoding/hex"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	//. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	MAX_BLOCKS	int64 = 1024*1
)
var (

	Info    *log.Logger
	RPC_URL string

	rpcclient *rpc.Client
	eclient *ethclient.Client

	storage *SQLStorage
	contract_addr	common.Address
)
func process_log(log *types.Log) {

	if len(log.Topics) == 0 {
		return
	}
	sigcode := hex.EncodeToString(log.Topics[0][0:4])
	evt_name := storage.Get_abi_event_name_by_signature(sigcode)
	fmt.Printf(
		"Block %v: Tx %v Signature %v (%v)\n",
		log.BlockNumber,log.TxHash.String(),sigcode,evt_name,
	)
}
func main() {


	if len(os.Args) < 3 {
		fmt.Printf("Usage: %v [contract_addr] [from_block]\n",os.Args[0])
		os.Exit(1)
	}
	contract_addr = common.HexToAddress(os.Args[1])
	block_num_start,err := strconv.ParseInt(os.Args[2],10,64)

	RPC_URL = os.Getenv("RPC_URL")

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	storage = Connect_to_storage(Info)
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
	latestBlock, err := eclient.BlockByNumber(ctx, nil)
	if err != nil {
		fmt.Printf("Error geting latest block: %v\n",err)
		os.Exit(1)
	}
	latest_bnum := latestBlock.Number().Int64()

	from_block := int64(block_num_start)
	to_block := int64(0)
	filter := ethereum.FilterQuery{}
	addresses := make([]common.Address,0,1)
	addresses = append(addresses,contract_addr)
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
