package main

import (
	"time"
	"bytes"
	"encoding/hex"
	"math/big"
	"os"
	"errors"
	//"context"
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func roll_back_blocks(diverging_block *types.Header) error {
	// Finds the block from which the fork started
	var err error
	var bhash common.Hash
	//ctx := context.Background()
	bhash,diverging_block,_, err = get_full_block(diverging_block.Number.Int64())
		//eclient.HeaderByNumber(ctx, diverging_block.Number)
	if err != nil {
		return errors.New(fmt.Sprintf("During chainsplit an error getting HeaderByHash happened: %v\n",err))
	}
	starting_block_num := diverging_block.Number.Int64()
	block_num:=starting_block_num
	block_hash:=bhash.String()
	Info.Printf("roll_back_blocks(): block_num = %v , block_hash %v",block_num,block_hash)
	Info.Printf("\t\t\tparent_hash %v\n",diverging_block.ParentHash.String())
	for {
		my_block_num,err := storage.Get_block_num_by_hash(block_hash)
		Info.Printf("Chainsplit fix: diverging hash %v, my_block_num=%v err=%v\n",block_hash,my_block_num,err)
		if err == nil {
			total_blocks := block_num - my_block_num
			if total_blocks < 0 { total_blocks = -total_blocks }
			if total_blocks > MAX_BLOCKS_CHAIN_SPLIT {
				Info.Printf(
					"Chainsplit fix: Chain split is longer than reasonal length, aborting. " +
					"(starting_block_num=%v, cur_block_num=%v",
					starting_block_num,block_num,
				)
				return errors.New("Chain split max size overflow")
			}
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
				"Chainsplit occurred at block %v and was fixed at block %v",starting_block_num,my_block_num,
			))
		} else {
			Info.Printf(
				"Chainsplit fix: block %v doesn't fit, block_hash=%v not found in my DB. Trying more...\n",
				block_num,block_hash,
			)
		}
		total_blocks := starting_block_num - block_num
		if total_blocks < 0 { total_blocks = -total_blocks }//just an extra safety against any bug before
		if total_blocks > MAX_BLOCKS_CHAIN_SPLIT {
			Info.Printf(
				"Chainsplit fix: Chain split is longer than reasonal length, aborting. " +
				"(starting_block_num=%v, cur_block_num=%v",
				starting_block_num,block_num,
			)
			return errors.New("Chain split max size overflow")
		}
		// keep trying by following parent hash
		bhash,diverging_block,_, err = get_full_block(diverging_block.Number.Int64()-1)
		//diverging_block, err = eclient.HeaderByHash(ctx, diverging_block.ParentHash)
		if err != nil {
			return errors.New(fmt.Sprintf("During chainsplit an error getting BlockByNumber happened: %v\n",err))
		}
		block_num = diverging_block.Number.Int64()
		block_hash = bhash.String()
		Info.Printf("Current block has been set to number %v , hash = %v\n",block_num,diverging_block.Hash().String())
	}
	return errors.New("Chainsplit fix: Undefined behaviour")
}
func process_transactions(bnum int64,transactions []*AugurTx,receipt_calls []*receiptCallResult,block_receipts types.Receipts) error {
	//	if receipt_calls is not nil then the old slow getTrasnactionReceipt call is used
	//	if block_receipts is not nil then we are using new getBlockReceipts RPC call
	for tnum,agtx := range transactions {
		if agtx.From == "0x0000000000000000000000000000000000000000" {
			continue // this is Polygon's automatic transaction
		}
		var rcpt *types.Receipt
		if receipt_calls != nil {
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
					agtx.TxHash,bnum,receipt_calls[tnum].err,
				)
				Error.Printf(
					"Failed to get Tx Receipt for %v, block num=%v. Aborting block processing: %v\n",
					agtx.TxHash,bnum,receipt_calls[tnum].err,
				)
				return receipt_calls[tnum].err
			}
			rcpt = receipt_calls[tnum].receipt
		} else {
			// receipts were fetched using eth_getBlockReceipts, we only need to reference the receipt
			rcpt = block_receipts[tnum]
		}
		//Info.Printf("\ttx: %v of %v : %v at blockNum=%v\n",tnum,len(transactions),agtx.TxHash,bnum)
		//Info.Printf("\t from=%v\n",agtx.From)
		//Info.Printf("\t to=%v for $%v (%v bytes data)\n",
		//				agtx.To,agtx.Value,len(agtx.Input))
		if rcpt.Status == types.ReceiptStatusFailed {
			//Info.Printf("Tx (index %v) %v . Status: Failed. Skipping this transaciton.\n",tnum,agtx.TxHash)
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
		transaction_hash := common.HexToHash(agtx.TxHash)
		if !bytes.Equal(rcpt.TxHash.Bytes(),transaction_hash.Bytes()) { // can be removed later
			Error.Printf("Receipt's hash doesn't match Tx hash, aborting (tx_hash=%v)",agtx.TxHash)
			os.Exit(1)
		}
		agtx.TxId = 0
		if agtx.CtrctCreate == true {
			agtx.To = rcpt.ContractAddress.String()
		}
		agtx.GasUsed = int64(rcpt.GasUsed)
		agtx.TxIndex = int32(tnum)
		agtx.NumLogs = int32(len(rcpt.Logs))
		logs_to_insert := prepare_event_log_batch(agtx,rcpt.Logs)
		if len(logs_to_insert) > 0 {
			storage.Insert_all_tx_event_logs(logs_to_insert)
		}
	}
	return nil
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
	//Info.Printf("block %v hash = %v, num_tx=%v\n",bnum,block_hash_str,num_transactions)
	if bnum!=header.Number.Int64() {
		Info.Printf("Retrieved block number %v but Block object contains another number (%v)",bnum,header.Number.Int64())
		Error.Printf("Retrieved block number %v but Block object contains another number (%v)",bnum,header.Number.Int64())
		return errors.New("Block object inconsistency")
	}
	storage.Block_delete_with_everything(big_bnum.Int64())
	var receipt_calls []*receiptCallResult = nil
	var block_receipts types.Receipts = nil
	if USE_BLOCK_RECEIPTS_RPC_CALL {
		block_receipts,err = get_block_receipts(block_hash)
		if err != nil {
			Error.Printf("Error getting receipts of the block: %v\n",err)
			return err
		}

	} else {
		receipt_calls = make([]*receiptCallResult,num_transactions,num_transactions)
		for i,tx := range transactions {
			hash := common.HexToHash(tx.TxHash)
			go get_receipt_async(i,hash,&receipt_calls)
		}
	}
	err = storage.Insert_block(block_hash_str,header,num_transactions,no_chainsplit_check)
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
	process_transactions(bnum,transactions,receipt_calls,block_receipts)
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
