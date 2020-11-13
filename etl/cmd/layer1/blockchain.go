package main

import (
	"time"
	"encoding/hex"
	"math/big"
	"os"
	"errors"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func roll_back_blocks(diverging_block *types.Header) error {
	// Finds the block from which the fork started
	ctx := context.Background()
	block_num:=diverging_block.Number.Int64()
	starting_block_num:=block_num
	for {
		big_block_num := big.NewInt(block_num)
		block, err := eclient.BlockByNumber(ctx,big_block_num)
		if err != nil {
			return err
		}
		if block == nil {
			e:=errors.New(fmt.Sprintf("ETH client api returned NULL block object (bnum=%v)",block_num))
			return e
		}
		block_hash:=block.Hash().String()
		my_block_num,err := storage.Get_block_num_by_hash(block_hash)
		Info.Printf("Chainsplit fix: hash %v, my_block_num=%v err=%v\n",block_hash,my_block_num,err)
		if err == nil {
			if my_block_num == block.Number().Int64() {
				Info.Printf(
					"Chainsplit fix: deleting blocks higher than %v ; good block hash = %v\n",
					my_block_num,block_hash,
				)
				storage.Chainsplit_delete_blocks(my_block_num)
				storage.Set_last_block_num(my_block_num)
				var chain_reorg_event ChainReorg
				chain_reorg_event.BlockNum = my_block_num
				chain_reorg_event.Hash = block_hash
				storage.Insert_chain_reorg_event(&chain_reorg_event)
				return errors.New(fmt.Sprintf(
					"Chainsplit occurred at block %v and was fixedx at block %v",block_num,my_block_num,
				))
			}
		} else {
			Info.Printf(
				"Chainsplit fix: block %v donesn't fit, block_hash=%v not found in my DB.\n",
				block_num,block_hash,
			)
		}
		block_num--
		if (starting_block_num - block_num) > MAX_BLOCKS_CHAIN_SPLIT {
			Info.Printf(
				"Chainsplit fix: Chain split is longer than reasonal length, aborting. " +
				"(starting_block_num=%v, cur_block_num=%v",
				starting_block_num,block_num,
			)
			return errors.New("Chain split max size overflow")
		}
	}
	return errors.New("Chainsplit fix: Undefined behaviour")
}
func process_block(bnum int64,update_last_block bool,no_chainsplit_check bool) error {

	block_hash_str,err:=get_block_hash(bnum)
	if err!=nil {
		return err
	}
	big_bnum:=big.NewInt(int64(bnum))
	block_hash,header,transactions,err := get_full_block(bnum)
	if err!=nil {
		Info.Printf("Can't decode Block object received on RPC: %v. Aborting.\n",err)
		return err
	}
	num_transactions := len(transactions)
	Info.Printf("block %v hash = %v, num_tx=%v\n",bnum,block_hash_str,num_transactions)
	if bnum!=header.Number.Int64() {
		Info.Printf("Retrieved block number %v but Block object contains another number (%v)",bnum,header.Number.Int64())
		Error.Printf("Retrieved block number %v but Block object contains another number (%v)",bnum,header.Number.Int64())
		return errors.New("Block object inconsistency")
	}
	storage.Block_delete_with_everything(big_bnum.Int64())
	receipt_calls := make([]*receiptCallResult,num_transactions,num_transactions)
	for i,tx := range transactions {
		hash := common.HexToHash(tx.TxHash)
		go get_receipt_async(i,hash,&receipt_calls)
	}
	err = storage.Insert_block(block_hash_str,header,no_chainsplit_check)
	if err != nil {
		err = roll_back_blocks(header)
		return err
	}
	if num_transactions == 0 {
		if update_last_block {
			storage.Set_last_block_num(bnum)
		}
		return nil
	}
	for tnum,agtx := range transactions {
		// wait for receipt to arrive
		for {
			if receipt_calls[tnum] != nil {
				break	// receipt arrived from the net, stop waiting
			}
			time.Sleep(1 * time.Millisecond)
		}
		if receipt_calls[tnum].err != nil {
			Info.Printf(
				"Failed to get Tx Receipt for %v, block num=%v. Aborting block processing: %v\n",
				agtx.TxHash,bnum,err,
			)
			Error.Printf(
				"Failed to get Tx Receipt for %v, block num=%v. Aborting block processing: %v\n",
				agtx.TxHash,bnum,err,
			)
			return receipt_calls[tnum].err
		}
		rcpt := receipt_calls[tnum].receipt
		Info.Printf("\ttx: %v of %v : %v at blockNum=%v\n",tnum,num_transactions,agtx.TxHash,bnum)
		Info.Printf("\t from=%v\n",agtx.From)
		Info.Printf("\t to=%v for $%v (%v bytes data)\n",
						agtx.To,agtx.Value,len(agtx.Input))
		if rcpt.Status == types.ReceiptStatusFailed {
			Info.Printf("\t Status: Failed. Skipping this transaciton.\n")
			continue	// transaction failed (i.e. Out of Gas, etc)
		}
		if rcpt.BlockNumber.Int64() != bnum {
			Error.Printf(
				"Transaction's receipt doesn't match current block number. (block possibly changed)" +
				" cur_block_num=%v, receipt.block_num=%v\n",
				bnum,rcpt.BlockNumber.Int64(),
			)
			return errors.New("Block changed during processing")
		}
		agtx.TxId = 0
		if agtx.CtrctCreate == true {
			agtx.To = rcpt.ContractAddress.String()
		}
		agtx.GasUsed = int64(rcpt.GasUsed)
		agtx.TxIndex = int32(tnum)
		logs_to_insert := prepare_event_log_batch(agtx,rcpt.Logs)
		if len(logs_to_insert) > 0 {
			storage.Insert_all_tx_event_logs(logs_to_insert)
		}
	}
	Info.Printf("block_proc: %v %v ; %v transactions\n",bnum,block_hash.String(),num_transactions)
	if update_last_block {
		storage.Set_last_block_num(bnum)
	}
	return nil
}
func prepare_event_log_batch(agtx *AugurTx,logs []*types.Log) []EthereumEventLog {

	var err error
	output := make([]EthereumEventLog,0,64)
	num_logs := len(logs)
	for i:=0 ; i < num_logs ; i++ {
		log := logs[i]
		if len(log.Topics) > 0 {
			tx_insert_if_needed(agtx)
			var eel EthereumEventLog
			eel.BlockNum = agtx.BlockNum
			eel.TxId = agtx.TxId
			eel.ContractAid = storage.Lookup_or_create_address(log.Address.String(),agtx.BlockNum,agtx.TxId)
			eel.Topic0_Sig = hex.EncodeToString(log.Topics[0][0:4])
			eel.RlpLog, err = rlp.EncodeToBytes(log)
			if err != nil {
				Info.Printf("Couldn't RLP-encode log : %v\n",err)
				os.Exit(1)
			}
			output = append(output,eel)
		}
	}

	return output
}
func process_tx_event_log(agtx *AugurTx,log *types.Log) {
	var err error
	var eel EthereumEventLog
	eel.BlockNum = agtx.BlockNum
	eel.TxId = agtx.TxId
	eel.ContractAid = storage.Lookup_or_create_address(log.Address.String(),agtx.BlockNum,agtx.TxId)
	eel.Topic0_Sig = hex.EncodeToString(log.Topics[0][0:4])
	eel.RlpLog, err = rlp.EncodeToBytes(log)
	if err != nil {
		Info.Printf("Couldn't RLP-encode log : %v\n",err)
		os.Exit(1)
	}
	storage.Insert_tx_event_log(&eel)
}
func tx_insert_if_needed(agtx *AugurTx) {
	if agtx.TxId == 0 {
		agtx.TxId=storage.Insert_transaction(agtx)
	}
}
