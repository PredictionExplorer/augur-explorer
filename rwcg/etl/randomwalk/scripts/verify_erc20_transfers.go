// Verifies that all transfers have corresponding entry in `rw_transfer` table
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

	"github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
	rwdb "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/randomwalk"
	rwp "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/randomwalk"
)
const (
	TRANSFER_EVT=   "ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	MAX_BLOCKS	int64 = 1024*1
)
var (

	Info    *log.Logger
	RPC_URL string
	evt_transfer,_   = hex.DecodeString(TRANSFER_EVT)
	rwalk_addr common.Address

	storagew *rwdb.SQLStorageWrapper
	caddrs   *rwp.ContractAddresses

	rpcclient *rpc.Client
	eclient   *ethclient.Client
)
func process_log(log *types.Log) {
	var evt rwp.TransferEntry
	evt.From = common.BytesToAddress(log.Topics[1][12:]).String()
	evt.To = common.BytesToAddress(log.Topics[2][12:]).String()
	evt.TokenId = log.Topics[3].Big().Int64()
	transfers := storagew.Get_rw_token_transfers_by_tx_hash(log.TxHash.String())
	found := false
	for i:=0; i<len(transfers); i++ {
		trf_item := &transfers[i]
		if evt.From == trf_item.From && evt.To == trf_item.To && evt.TokenId==trf_item.TokenId {
			found = true
		}
	}
	if !found {
		fmt.Printf(
			"Block %v: Transfer %v -> %v (tok %v) tx %v, transfer missing\n",
			log.BlockNumber,evt.From,evt.To,evt.TokenId,log.TxHash.String(),
		)
	}
}
func main() {


	RPC_URL = os.Getenv("RPC_URL")

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	base := dbs.Connect_to_storage(Info)
	storagew = &rwdb.SQLStorageWrapper{S: base}

	caddrs_obj := storagew.Get_randomwalk_contract_addresses()
	caddrs = &caddrs_obj

	rwalk_addr = common.HexToAddress(caddrs.RandomWalk)

	eclient, err := ethclient.Dial(RPC_URL)
	if err != nil {
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
	ctx := context.Background()
	latestBlock, err := eclient.BlockByNumber(ctx, nil)
	if err != nil {
		fmt.Printf("Error geting latest block: %v\n",err)
		os.Exit(1)
	}
	latest_bnum := latestBlock.Number().Int64()

	from_block := int64(0)
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
