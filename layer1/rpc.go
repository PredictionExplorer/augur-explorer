package layer1

import (
	"fmt"
	"errors"
	//"reflect"
	"encoding/json"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common/hexutil"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
type rpcBlockHash struct {
	Hash		string
}
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
func get_full_block(etl *ETL_Layer1,bnum int64) (common.Hash,*types.Header,[]*AugurTx,error) {

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
		etl.Error.Printf("Error unmarshalling transactions of the block: %v\n",err)
		return common.Hash{},head, make([]*AugurTx,0,0),err
	}
	err = json.Unmarshal(raw,&head)
	if err!= nil {
		etl.Error.Printf("Error unmarshalling hash of the block: %v\n",err)
		return body.Hash,head,make([]*AugurTx,0,0),err
	}
	txs := make([]*AugurTx, len(body.Transactions))
	for i, tx := range body.Transactions {
		agtx := new(AugurTx)
		agtx.BlockNum = bnum
		agtx.BlockHash = tx.txExtraInfo.BlockHash.String()
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
	}

	return body.Hash,head,txs,nil
}
func get_block_hash(etl *ETL_Layer1,block_num int64) (string,error) {

	// this function is needed because Parity doesn't return the correct block hash over RPC, the hash
	// it returns is re-calculated while some fileds of the types.Header object are unset, giving wrong hash
	ctx := context.Background()
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getBlockByNumber", hexutil.EncodeBig(big.NewInt(int64(block_num))),false)
	var blockHash rpcBlockHash
	err = json.Unmarshal(raw,&blockHash)
	if err!= nil {
		etl.Error.Printf("Error unmarshalling hash of the block: %v\n",err)
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
func get_receipt_async_custom_rpc(etl *ETL_Layer1,idx int,tx_hash common.Hash,receipt_results *[]*receiptCallResult) {
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
			etl.Error.Printf("Error unmarshalling receipt object : %v\n",err)
			result.err=err
		}
		var rcpt_extra rcptExtraInfo
		err = json.Unmarshal(raw, &rcpt_extra)
		if err != nil {
			etl.Error.Printf("Error unmarshalling receipt extra data : %v\n",err)
			result.err=err
		}
		cum_gas,err := hexutil.DecodeUint64(rcpt_extra.CumulativeGasUsed)
		if err != nil {
			etl.Error.Printf("Error parsing CumulativeGas %v: %v\n",rcpt_extra.CumulativeGasUsed,err)
			result.err=err
		}
		extra.CumulativeGasUsed = cum_gas
		egasp,err := hexutil.DecodeBig(rcpt_extra.EffectiveGasPrice)
		if err != nil {
			etl.Error.Printf("Error parsing EffectiveGasPrice %v : %v\n",rcpt_extra.EffectiveGasPrice,err)
			result.err=err
		}
		extra.EffectiveGasPrice = egasp
		result.receipt = rcpt
		result.extra = extra
	}
	(*receipt_results)[idx]=result
}
func get_block_receipts_v1(etl *ETL_Layer1,block_hash common.Hash) (types.Receipts,error) {

	// this is version 1, doesn't have effectiveTxGasPrice field
	ctx := context.Background()
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getBlockReceipts", block_hash)
	var receipts types.Receipts
	err = json.Unmarshal(raw,&receipts)
	if err!= nil {
		etl.Error.Printf("Error unmarshalling receipts after eth-getBlockReceipts: %v\n",err)
		return nil,err
	} else {
		return receipts,nil
	}
}
type PackedReceipt map[string]interface{}
type ReceiptsPackage struct {
	Receipts		[]PackedReceipt
}
func get_block_receipts_v2(etl *ETL_Layer1,block_hash common.Hash) ([]types.Receipt,[]ReceiptExtraInfo ,error) {
	// this is version 2, matches version 2 of the patch to Geth, which sends full transaction info
	ctx := context.Background()
	var raw json.RawMessage
	err := rpcclient.CallContext(ctx, &raw,"eth_getBlockReceipts", block_hash)
	var packed_receipts ReceiptsPackage
	err = json.Unmarshal(raw,&packed_receipts)
	if err!= nil {
		etl.Error.Printf("Error unmarshalling receipts after eth-getBlockReceipts on block %v : %v\n",block_hash.String(),err)
		etl.Error.Printf("Raw JSON msg:\n%v\n",string(raw))
		return nil,nil,err
	}
	output_receipts := make([]types.Receipt,0,256)
	output_extra := make([]ReceiptExtraInfo,0,256)
	for i:=0; i<len(packed_receipts.Receipts);i++ {
		var r types.Receipt
		receipt := packed_receipts.Receipts[i]
		tmp_val,err := hexutil.DecodeUint64(receipt["type"].(string))
		if err != nil {
			return nil,nil,errors.New(fmt.Sprintf("Error: receipt (i=%v) type (%v): %v",i,receipt["type"],err))
		}
		r.Type = uint8(tmp_val)
		tmp_val,err = hexutil.DecodeUint64(receipt["status"].(string))
		if err != nil {
			return nil,nil,errors.New(fmt.Sprintf("Error: receipt (i=%v) status (%v): %v",i,receipt["status"],err))
		}
		r.Status = tmp_val
		if receipt["logs"] != nil {
			logs_in := receipt["logs"].([]interface{})
			var logs_out []*types.Log
			if len(logs_in) >0 {
				logs_out = make([]*types.Log,0,32)
			}
			for j:=0;j<len(logs_in);j++ {
				lin := logs_in[j].(map[string]interface{})
				lout := &types.Log{}

				lout.Address = common.HexToAddress(lin["address"].(string))
				tmp_val, err := hexutil.DecodeUint64(lin["blockNumber"].(string))
				if err != nil {
					return nil,nil,errors.New(
						fmt.Sprintf(
							"Error: receipt (i=%v) log (j=%v) blockNumber (%v): %v",
							i,j,lin["blockNumber"].(string),err,
						),
					)
				}
				lout.BlockNumber = tmp_val

				lout.TxHash = common.HexToHash(lin["transactionHash"].(string))
				tmp_val, err = hexutil.DecodeUint64(lin["transactionIndex"].(string))
				if err != nil {
					return nil,nil,errors.New(
						fmt.Sprintf(
							"Error: receipt (i=%v) log (j=%v) transactionIndex (%v): %v",
							i,j,lin["transactionIndex"].(string),err,
						),
					)
				}
				lout.TxIndex = uint(tmp_val)

				lout.BlockHash = common.HexToHash(lin["blockHash"].(string))

				tmp_val, err = hexutil.DecodeUint64(lin["logIndex"].(string))
				if err != nil {
					return nil,nil,errors.New(
						fmt.Sprintf(
							"Error: receipt (i=%v) log (j=%v) logIndex (%v): %v",
							i,j,lin["logIndex"].(string),err,
						),
					)
				}
				lout.Index = uint(tmp_val)

				lout.Removed = lin["removed"].(bool)

				tmp_data,err := hexutil.Decode(lin["data"].(string))
				if err != nil {
					return nil,nil,errors.New(
						fmt.Sprintf(
							"Error: receipt (i=%v) log (j=%v) data (%v): %v",
							i,j,lin["data"].(string),err,
						),
					)
				}
				lout.Data = tmp_data
				topics_in := lin["topics"].([]interface{})
				if len(topics_in) >0 {
					lout.Topics = make([]common.Hash,0,3)
					for k:=0;k<len(topics_in);k++ {
						t := common.HexToHash(topics_in[k].(string))
						lout.Topics = append(lout.Topics,t)
					}
				}
				logs_out = append(logs_out,lout)
			}
			if len(logs_out) > 0 {
				r.Logs = logs_out
			}
		}
		r.TxHash = common.HexToHash(receipt["transactionHash"].(string))
		if receipt["contractAddress"] != nil {
			r.ContractAddress = common.HexToAddress(receipt["contractAddress"].(string))
		}
		tmp_val,err = hexutil.DecodeUint64(receipt["gasUsed"].(string))
		if err != nil {
			return nil,nil,errors.New(fmt.Sprintf("Error: receipt (i=%v) gasUsed(%v): %v",i,receipt["gasUsed"].(string),err))
		}
		r.GasUsed = tmp_val
		r.BlockHash = common.HexToHash(receipt["blockHash"].(string))
		tmp_val,err = hexutil.DecodeUint64(receipt["blockNumber"].(string))
		if err != nil {
			return nil,nil,errors.New(fmt.Sprintf("Error: receipt (i=%v) blockNumber(%v): %v",i,receipt["blockNumber"].(string),err))
		}
		r.BlockNumber = big.NewInt(0).SetUint64(tmp_val)
		tmp_val,err = hexutil.DecodeUint64(receipt["transactionIndex"].(string))
		if err != nil {
			return nil,nil,errors.New(fmt.Sprintf("Error: receipt (i=%v) transactionIndex(%v): %v",i,receipt["transactionIndex"].(string),err))
		}
		r.TransactionIndex = uint(tmp_val)

		var extra_fields ReceiptExtraInfo
		tmp_val,err = hexutil.DecodeUint64(receipt["effectiveGasPrice"].(string))
		if err != nil {
			return nil,nil,errors.New(fmt.Sprintf("Error: receipt (i=%v) effectiveGasPrice(%v): %v",i,receipt["effectiveGasPrice"],err))
		}
		extra_fields.EffectiveGasPrice = big.NewInt(0).SetUint64(tmp_val)
		tmp_val, err = hexutil.DecodeUint64(receipt["cumulativeGasUsed"].(string))
		if err != nil {
			return nil,nil,errors.New(fmt.Sprintf("Error: receipt (i=%v) cumulativeGasUsed(%v): %v",i,receipt["cumulativeGasUsed"],err))
		}
		extra_fields.CumulativeGasUsed = tmp_val
		output_receipts = append(output_receipts,r)
		output_extra = append(output_extra,extra_fields)
	}
	return output_receipts,output_extra,nil
}
