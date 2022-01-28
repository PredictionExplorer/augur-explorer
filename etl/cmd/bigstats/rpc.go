package main

import (
	//"fmt"
	//"errors"
	"encoding/json"
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
type receiptCallResult struct {
	receipt		*types.Receipt
	extra		*ReceiptExtraInfo
	err			error
	idx			int
}
type rcptExtraInfo struct {
	EffectiveGasPrice string	`json: effectiveGasPrice,omitempty`
	CumulativeGasUsed string	`json: "cumulativeGasUsed,omitempty"`
}
type ReceiptExtraInfo struct {
	EffectiveGasPrice			*big.Int
	CumulativeGasUsed			uint64
}
/*
// senderFromServer is a types.Signer that remembers the sender address returned by the RPCng.
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
		Info.Printf("block %v, index %v, hash = %v\n",bnum,i,tx.tx.Hash().String())
		Info.Printf("gasPrice=%v\n",tx.tx.GasPrice().String())
		Info.Printf("gasTipCap=%v\n",tx.tx.GasTipCap().String())
		Info.Printf("gasFeeCap=%v\n",tx.tx.GasFeeCap().String())
		Info.Printf("Value=%v\n",tx.tx.Value().String())
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
func get_receipt_async_custom_rpc(idx int,tx_hash common.Hash,receipt_results *[]*receiptCallResult) {
	ctx := context.Background()
	result := new(receiptCallResult)
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getTransactionReceipt", tx_hash)
	result.idx = idx
	if err != nil {
		result.err = err
	} else {
		extra := new(ReceiptExtraInfo)
		rcpt := new(types.Receipt)
		err = json.Unmarshal(raw, &rcpt);
		if err != nil {
			Error.Printf("Error unmarshalling receipt object : %v\n",err)
		}
		var rcpt_extra rcptExtraInfo
		err = json.Unmarshal(raw, rcpt_extra)
		if err != nil {
			Error.Printf("Error unmarshalling receipt extra data : %v\n",err)
			result.err=err
		}
		cum_gas,err := hexutil.DecodeUint64(rcpt_extra.CumulativeGasUsed)
		if err != nil {
			Error.Printf("Error parsing CumulativeGas %v: %v\n",rcpt_extra.CumulativeGasUsed,err)
		}
		extra.CumulativeGasUsed = cum_gas
		egasp,err := hexutil.DecodeBig(rcpt_extra.EffectiveGasPrice)
		if err != nil {
			Error.Printf("Error parsing EffectiveGasPrice %v : %v\n",rcpt_extra.EffectiveGasPrice,err)
			result.err=err
		}
		extra.EffectiveGasPrice = egasp
		Info.Printf("CumulativeGas=%v, EffectiveGasPrice=%v\n",extra.CumulativeGasUsed,extra.EffectiveGasPrice)
		result.receipt = rcpt
		result.extra = extra
	}
	(*receipt_results)[idx]=result
}
func get_block_receipts(block_hash common.Hash) (types.Receipts,error) {

	ctx := context.Background()
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getBlockReceipts", block_hash)
	var receipts types.Receipts
	err = json.Unmarshal(raw,&receipts)
	if err!= nil {
		Error.Printf("Error unmarshalling receipts after eth-getBlockReceipts: %v\n",err)
		return nil,err
	} else {
		return receipts,nil
	}
}
