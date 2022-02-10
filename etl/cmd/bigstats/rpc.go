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
		err = json.Unmarshal(raw, &rcpt_extra)
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
		result.receipt = rcpt
		result.extra = extra
	}
	(*receipt_results)[idx]=result
}
func get_block_receipts_v1(block_hash common.Hash) (types.Receipts,error) {

	// this is version 1, doesn't have effectiveTxGasPrice field
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
type PackedReceipt map[string]interface{}
type ReceiptsPackage struct {
	Receipts		[]PackedReceipt
}
func get_block_receipts_v2(block_hash common.Hash) ([]types.Receipt,[]ReceiptExtraInfo ,error) {
	// this is version 2, matches version 2 of the patch to Geth, which sends full transaction info
	ctx := context.Background()
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getBlockReceipts", block_hash)
	var packed_receipts ReceiptsPackage
	err = json.Unmarshal(raw,&packed_receipts)
	if err!= nil {
		Error.Printf("Error unmarshalling receipts after eth-getBlockReceipts: %v\n",err)
		return nil,nil,err
	}
	output_receipts := make([]types.Receipt,0,256)
	output_extra := make([]ReceiptExtraInfo,0,256)
	for i:=0; i<len(packed_receipts.Receipts);i++ {
		var r types.Receipt
		receipt := packed_receipts.Receipts[i]
		r.Type = receipt["type"].(uint8)
		r.Status = receipt["status"].(uint64)
		if receipt["logs"].([][]*types.Log) != nil {
			r.Logs = receipt["logs"].([]*types.Log)
		}
		r.TxHash.SetBytes(receipt["transactionHash"].(common.Hash).Bytes())
		if receipt["contractAddress"] != nil {
			r.ContractAddress.SetBytes(receipt["contractAddress"].(common.Address).Bytes())
		}
		Info.Printf("Receipt for tx %v (%v)\n",i,r.TxHash.String())
		r.GasUsed = uint64(receipt["gasUsed"].(uint64))
		r.BlockHash.SetBytes(receipt["blockHash"].(common.Hash).Bytes())
		var extra_fields ReceiptExtraInfo
		extra_fields.EffectiveGasPrice = big.NewInt(0).SetUint64(receipt["effectiveGasPrice"].(uint64))
		extra_fields.CumulativeGasUsed = receipt["cumulativeGasUsed"].(uint64)
		output_receipts = append(output_receipts,r)
		output_extra = append(output_extra,extra_fields)
	}
	return output_receipts,output_extra,nil
}
