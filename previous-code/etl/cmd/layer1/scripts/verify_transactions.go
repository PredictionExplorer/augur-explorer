// Checks that all transacitons of the block are present in the database, notifies integrity failure
package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"time"
	"encoding/json"
	"context"
	"math/big"
	"database/sql"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	NUM_BLOCKS_TO_LAG	int64 = 256	// to avoid false positives due to chain split
)
type rpcBlock struct {
	Hash         common.Hash      `json:"hash"`
	Transactions []rpcTransaction `json:"transactions"`
	UncleHashes  []common.Hash    `json:"uncles"`
}
type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}
type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
	Hash		*common.Hash	`json:"hash,omitempty"`
}
var (
	storage *SQLStorage

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")

	eclient *ethclient.Client
	rpcclient *rpc.Client

	market_order_id int64 = 0

	Info	*log.Logger

)
func process_block(bnum int64) {

	var raw json.RawMessage
	ctx := context.Background()
	err := rpcclient.CallContext(ctx, &raw,"eth_getBlockByNumber", hexutil.EncodeBig(big.NewInt(int64(bnum))),true)
	if err != nil {
		Info.Printf("Error getting block %v, exiting\n",bnum)
		os.Exit(1)
	}
	var head *types.Header
	var body rpcBlock
	err = json.Unmarshal(raw, &body);
	if err != nil {
		Info.Printf("Error unmarshalling transactions of the block (num=%v): %v, exiting.\n",bnum,err)
		os.Exit(1)
	}
	err = json.Unmarshal(raw,&head)
	if err!= nil {
		Info.Printf("Error unmarshalling hash of the block (num=%v): %v, exiting.\n",bnum,err)
		os.Exit(1)
	}
	for _, tx := range body.Transactions {
		receipt,err := eclient.TransactionReceipt(context.Background(),*tx.Hash)
		if err != nil {
			Info.Printf("Error getting receipt for tx %v of block %v: %v, exiting.",tx.Hash.String(),bnum,err)
			os.Exit(1)
		}
		Info.Printf("block %v, tx %v\n",bnum,tx.Hash.String())
		if receipt.Status == types.ReceiptStatusFailed {
			continue
		}
		logs := receipt.Logs
		has_events := false
		for i:=0; i<len(logs); i++ {
			l := logs[i]
			if len(l.Topics) > 0 {
				has_events = true
			}
		}
		if !has_events {
			continue
		}
		tx_id,err := storage.Get_tx_id_by_hash(tx.Hash.String())
		if (err != nil) && (err != sql.ErrNoRows) {
			Info.Printf("Error querying tx (hash=%v, block=%v): %v, exiting.\n",tx.Hash.String(),bnum,err)
			os.Exit(1)
		}
		if tx_id == 0 {
			Info.Printf("MISSING TX: %v block %v\n",tx.Hash.String(),bnum)
		}
	}
}
func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v [starting_block_num]\n",os.Args[0])
		fmt.Printf("\tScans blocks and verifies the DB has all transactions it should have\n")
		os.Exit(1)
	}

	starting_block,err := strconv.ParseInt(os.Args[1],10,64)
	if err != nil {
		fmt.Printf("Error parsing starting block num: %v\n",err)
		os.Exit(1)
	}

	RPC_URL = os.Getenv("RPC_URL")
	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	storage = Connect_to_storage(&market_order_id,Info)

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("Connected to ETH node: %v\n",RPC_URL)
	eclient = ethclient.NewClient(rpcclient)

	for {
		ctx := context.Background()
		latestBlock, err := eclient.BlockByNumber(ctx, nil)
		if err != nil {
			Info.Printf("BlockByNumber() failed on block %v, exiting.\n",starting_block)
			os.Exit(1)
		}
		latest_bnum := latestBlock.Number().Int64()
		if starting_block < (latest_bnum-NUM_BLOCKS_TO_LAG) {
			for ;starting_block < (latest_bnum-NUM_BLOCKS_TO_LAG);starting_block=starting_block+1 {
				process_block(starting_block)
				mod := starting_block % 1000
				if mod == 0 {
					Info.Printf("Checking errors in blocks above number %v ...\n",starting_block)
				}
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
