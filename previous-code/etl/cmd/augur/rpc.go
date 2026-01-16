package main

import (
	//"fmt"
	"errors"
	"encoding/json"
	//"encoding/hex"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}
func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}
type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
	Hash		*common.Hash	`json:"hash,omitempty"`
}
type rpcBlock struct {
	Hash         common.Hash      `json:"hash"`
	Transactions []rpcTransaction `json:"transactions"`
	UncleHashes  []common.Hash    `json:"uncles"`
}
type rpcBlockHash struct {
	Hash		string
}
type receiptCallResult struct {
	receipt		*types.Receipt
	err			error
	idx			int
}
/*
// senderFromServer is a types.Signer that remembers the sender address returned by the RPC
// server. It is stored in the transaction's sender address cache to avoid an additional
// request in TransactionSender.
type senderFromServer struct {
	addr      common.Address
	blockhash common.Hash
}
func setSenderFromServer(tx *types.Transaction, addr common.Address, block common.Hash) {
	// Use types.Sender for side-effect to store our signer into the cache.
	types.Sender(&senderFromServer{addr, block}, tx)
}*/
func get_full_block(bnum int64) (common.Hash,*types.Header,[]*AugurTx,error) {

	var head *types.Header
	ctx := context.Background()
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getBlockByNumber", hexutil.EncodeBig(big.NewInt(int64(bnum))),true)
	if err != nil {
		return common.Hash{},head,make([]*AugurTx,0,0),err
	}
	var body rpcBlock
	err = json.Unmarshal(raw, &body);
	if err != nil {
		Error.Printf("Error unmarshalling transactions of the block: %v\n",err)
		return common.Hash{},head, make([]*AugurTx,0,0),err
	}
	err = json.Unmarshal(raw,&head)
	if err!= nil {
		Error.Printf("Error unmarshalling hash of the block: %v\n",err)
		return body.Hash,head,make([]*AugurTx,0,0),err
	}
	txs := make([]*AugurTx, len(body.Transactions))
	for i, tx := range body.Transactions {
		/*if tx.From != nil {
			setSenderFromServer(tx.tx, *tx.From, body.Hash)
		}*/
		agtx := new(AugurTx)
		agtx.BlockNum = bnum
		agtx.TxHash = tx.txExtraInfo.Hash.String()
		agtx.From = tx.txExtraInfo.From.String()
		if tx.tx.To() != nil {
			agtx.To  = tx.tx.To().String()
		} else {
			agtx.CtrctCreate = true
			agtx.To = "0x0000000000000000000000000000000000000000"
		}
		agtx.Value = tx.tx.Value().String()
		agtx.Input = tx.tx.Data()
		agtx.GasPrice = tx.tx.GasPrice().String()
		txs[i]=agtx
		//txs[i] = tx.tx
	}

	return body.Hash,head,txs,nil
}
func rpc_get_transaction(hash common.Hash) (*AugurTx,error) {

	ctx := context.Background()
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getTransactionByHash", hash)
	if err != nil {
		return nil,err
	}
	var tx *rpcTransaction
	err = json.Unmarshal(raw, &tx);
	if err != nil {
		Error.Printf("Error unmarshalling transactions of the block: %v\n",err)
		return nil,err
	}
	if tx == nil {
		return nil,errors.New("augur rpc: transaction not found")
	}
	_,r,_ := tx.tx.RawSignatureValues()
	if r ==nil {
		return nil,errors.New("transaction without signature")
	}
	agtx := new(AugurTx)
	big_block_num,err := hexutil.DecodeBig(*tx.txExtraInfo.BlockNumber)
	if err != nil {
		Error.Printf("cant decode hex encoded block number from RPC")
		return nil,err
	}
	agtx.BlockNum = big_block_num.Int64()
	agtx.TxHash = tx.txExtraInfo.Hash.String()
	agtx.From = tx.txExtraInfo.From.String()
	if tx.tx.To() != nil {
		agtx.To  = tx.tx.To().String()
	} else {
		agtx.CtrctCreate = true
		agtx.To = "0x0000000000000000000000000000000000000000"
	}
	agtx.Value = tx.tx.Value().String()
	agtx.Input = tx.tx.Data()
	agtx.GasPrice = tx.tx.GasPrice().String()

	return agtx,nil
}
func get_block_hash(block_num int64) (string,error) {

	// this function is needed because Parity doesn't return the correct block hash over RPC, the hash
	// it returns is re-calculated while some fileds of the types.Header object are unset, giving wrong hash
	ctx := context.Background()
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getBlockByNumber", hexutil.EncodeBig(big.NewInt(int64(block_num))),false)
	var blockHash rpcBlockHash
	err = json.Unmarshal(raw,&blockHash)
	if err!= nil {
		Error.Printf("Error unmarshalling hash of the block: %v\n",err)
		return "",err
	} else {
		return blockHash.Hash,nil
	}
}
func get_receipt_async(idx int,tx_hash common.Hash,receipt_results *[]*receiptCallResult) {
	// this func is launched as go-routine
	ctx := context.Background()
	result := new(receiptCallResult)
	result.receipt,result.err = eclient.TransactionReceipt(ctx,tx_hash)
	result.idx = idx
	(*receipt_results)[idx]=result
}
