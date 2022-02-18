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
type ReceiptExtraInfo struct {
	EffectiveGasPrice			*big.Int
	CumulativeGasUsed			uint64
}
type rcptExtraInfo struct {
	EffectiveGasPrice string	`json: effectiveGasPrice,omitempty`
	CumulativeGasUsed string	`json: "cumulativeGasUsed,omitempty"`
}
type receiptCallResult struct {
	receipt		*types.Receipt
	extra		*ReceiptExtraInfo
	err			error
	idx			int
}
var (
	storage *SQLStorage

	RPC_URL string

	eclient *ethclient.Client
	rpcclient *rpc.Client

	market_order_id int64 = 0

	Info	*log.Logger

)
func get_receipt_custom_rpc(tx_hash common.Hash) (*receiptCallResult,error) {
	ctx := context.Background()
	result := new(receiptCallResult)
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getTransactionReceipt", tx_hash)
	if err != nil {
		result.err = err
	} else {
		extra := new(ReceiptExtraInfo)
		rcpt := new(types.Receipt)
		err = json.Unmarshal(raw, &rcpt);
		if err != nil {
			Info.Printf("Error unmarshalling receipt object : %v\n",err)
			return nil,err
		}
		var rcpt_extra rcptExtraInfo
		err = json.Unmarshal(raw, &rcpt_extra)
		if err != nil {
			Info.Printf("Error unmarshalling receipt extra data : %v\n",err)
			return nil,err
		}
		cum_gas,err := hexutil.DecodeUint64(rcpt_extra.CumulativeGasUsed)
		if err != nil {
			Info.Printf("Error parsing CumulativeGas %v: %v\n",rcpt_extra.CumulativeGasUsed,err)
			return nil,err
		}
		extra.CumulativeGasUsed = cum_gas
		egasp,err := hexutil.DecodeBig(rcpt_extra.EffectiveGasPrice)
		if err != nil {
			Info.Printf("Error parsing EffectiveGasPrice %v : %v\n",rcpt_extra.EffectiveGasPrice,err)
			return nil,err
		}
		extra.EffectiveGasPrice = egasp
		result.receipt = rcpt
		result.extra = extra
	}
	return result,err
}
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
	for tx_index, tx := range body.Transactions {
		receipt_result,err := get_receipt_custom_rpc(*tx.Hash)
		if err != nil {
			Info.Printf("Error getting receipt for tx %v of block %v: %v, exiting.",tx.Hash.String(),bnum,err)
			os.Exit(1)
		}
		//Info.Printf("block %v, tx %v\n",bnum,tx.Hash.String())
		if receipt_result.receipt.Status == types.ReceiptStatusFailed {
			continue
		}
		db_tx_fee,err := storage.Bigstats_get_tx_fee_by_index(bnum,int64(tx_index))
		if err != nil {
			if err != sql.ErrNoRows {
				Info.Printf("Error querying tx (block=%v,index=%v): %v, exiting.\n",bnum,tx_index,err)
				os.Exit(1)
			} else {
				Info.Printf("Reached the end of the data in the database, exiting\n")
				os.Exit(0)
			}
		}
		tx_fee := big.NewInt(0).SetUint64(receipt_result.receipt.GasUsed)
		tx_fee.Mul(tx_fee,receipt_result.extra.EffectiveGasPrice)
		if tx_fee.String() != db_tx_fee {
			Info.Printf(
				"MISMATCH: b %v, i %v ( %v ): fee should be %v but the DB has %v\n",
				bnum,tx_index,tx.Hash.String(),tx_fee.String(),db_tx_fee,
			)
		}
	}
}
func main() {

	if len(os.Args) < 4 {
		fmt.Printf("Usage: %v  [schema] [rpc_url] [starting_block_num]\n",os.Args[0])
		fmt.Printf("\tScans blocks and verifies the DB has all transactions it should have and checks the value of the transaction fee\n")
		os.Exit(1)
	}

	starting_block,err := strconv.ParseInt(os.Args[3],10,64)
	if err != nil {
		fmt.Printf("Error parsing starting block num: %v\n",err)
		os.Exit(1)
	}

	RPC_URL = os.Args[2]
	schema_name := os.Args[1]

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	storage = Connect_to_storage(&market_order_id,Info)
	storage.Db_set_schema_name(schema_name)

	rpcclient, err=rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	Info.Printf("RPC: %v\n",RPC_URL)
	Info.Printf("Schema name: %v\n",schema_name)
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
