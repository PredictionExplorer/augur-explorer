package main

import (
	"time"
//	"bytes"
	"encoding/hex"
	"math/big"
//	"bufio"
//	"context"
	"os"
	"errors"
//	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/0xProject/0x-mesh/zeroex"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
/*
func process_block_orig(bnum int64,update_last_block bool,no_chainsplit_check bool) error {

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
		if err == nil {
			return nil
		}
		Error.Printf("Unable to recover from chainsplit: %v. Aborting",err)
		os.Exit(1)
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
		dump_tx_input_if_known(agtx.Input)
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
		sequencer := new(EventSequencer)
		num_logs := len(rcpt.Logs)
		var agtx_type int = AgTxType_Unclassified
		// Step 1: First detect what kind of Augur Transaction we are dealing with
		for i:=0 ; i<num_logs ; i++ {
			if len(rcpt.Logs[i].Topics) > 0 {
				Info.Printf(
					"\t\tlog %v\t for contract %v (%v of %v items)\n",
					hex.EncodeToString(rcpt.Logs[i].Topics[0][0:4]),
					rcpt.Logs[i].Address.String(),(i+1),len(rcpt.Logs))
				if 0 == bytes.Compare(rcpt.Logs[i].Topics[0].Bytes(),evt_market_finalized) {
					if is_warp_sync_event(rcpt.Logs[i]) {
						// WarpSync market emits 2 events MarketFFinalized and MarketCreated
						// MarketFinalized doesn't have ProfitLoss events, so we can process it
						// just using inverse order (i.e. considering it as non-MarketFinalized)
					} else {
						agtx_type = AgTxType_MarketFinalized
					}
				}
				if 0 == bytes.Compare(rcpt.Logs[i].Topics[0].Bytes(),evt_market_order) {
					agtx_type = AgTxType_MarketOrder
				}
			}
			sequencer.append_event(rcpt.Logs[i])
		}
		// Step 1.1 If a Wallet contract has been created, register EOA-Wallet link
		wallet_created,wallet_addr,possible_eoa_addr := was_wallet_created(caddrs,rcpt.Logs)
		if wallet_created {
			tx_insert_if_needed(agtx)
			var from_addr *string = &agtx.From
			var possible_eoa_str string
			if possible_eoa_addr != nil {
				possible_eoa_str = possible_eoa_addr.String()
				from_addr = &possible_eoa_str
			}
			storage.Register_eoa_and_wallet(*from_addr,wallet_addr.String(),agtx.BlockNum,agtx.TxId)
		}
		// Step 1.2 If transaction contains executeWalletTransaction, store it in the DB
		if (agtx.To == caddrs.WalletReg.String()) || (agtx.To == caddrs.WalletReg2.String()) {
			exec_wtx := contains_execute_wallet_transaction_call(agtx.Input)
			if exec_wtx != nil {
				tx_insert_if_needed(agtx)
				// Note: Tests are pending for transactions going through GSN (EOA address must be extracted
				//			from the tx.Data(), in the first 20 bytes
				eoa_aid := storage.Lookup_or_create_address(agtx.From,agtx.BlockNum,agtx.TxId)
				wallet_aid,err := storage.Lookup_wallet_aid(eoa_aid)
				if err != nil {
					Info.Printf(
						"executeWalletTransaction(): wallet_aid=0 for eoa=%v (id=%v)\n",
						agtx.From,eoa_aid,
					)
				}
				exec_wtx.Dump(Info)
				storage.Insert_execute_wallet_tx(eoa_aid,wallet_aid,agtx,exec_wtx)
			}
		}
		// Step 2: Knowing what kind of Augur Transaction, we are sorting events in an order
		//			that is convinient for us to process the event series
		var ordered_list []*types.Log
		switch agtx_type {
			case AgTxType_Unclassified:
				ordered_list = sequencer.get_ordered_event_list()
			case AgTxType_MarketFinalized:
				// logs with Market finalized event need to have special order
				ordered_list = sequencer.get_events_for_market_finalized_case()
			case AgTxType_MarketOrder:
				ordered_list = sequencer.get_events_for_market_order_case()
			default:
				Info.Printf("Undefined behaviour in detecting Augur Transaction type")
				os.Exit(1)
		}
		num_logs = len(ordered_list)
		pl_entries := make([]int64,0,2);// profit loss entries
		// before processing events we need to reset these global vars as they accumulate some data
		market_order_id = 0
		initial_amount = nil
		//
		// Step 3: Execute events using ordered list prepared in previous step
		for i:=0 ; i < num_logs ; i++ {
			if len(ordered_list[i].Topics) > 0 {
				Info.Printf(
					"\t\tchecking log with sig %v\t for contract %v\n",
					hex.EncodeToString(ordered_list[i].Topics[0][0:4]),
					ordered_list[i].Address.String())
				id := process_event(header,agtx,&ordered_list,i)
				if 0 == bytes.Compare(ordered_list[i].Topics[0].Bytes(),evt_profit_loss_changed) {
					pl_entries = append(pl_entries,id)
				}
			}
		}
	}
	Info.Printf("block_proc: %v %v ; %v transactions\n",bnum,block_hash.String(),num_transactions)
	if update_last_block {
		storage.Set_last_block_num(bnum)
	}
	return nil
}
*/
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
		if err == nil {
			return nil
		}
		Error.Printf("Unable to recover from chainsplit: %v. Aborting",err)
		os.Exit(1)
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
		dump_tx_input_if_known(agtx.Input)
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
		num_logs := len(rcpt.Logs)
		//var agtx_type int = AgTxType_Unclassified
		// Step 1.2 If transaction contains executeWalletTransaction, store it in the DB
		for i:=0 ; i < num_logs ; i++ {
			log := rcpt.Logs[i]
			if len(log.Topics) > 0 {
				tx_insert_if_needed(agtx)
				process_tx_event_log(agtx,log)
			}
		}
	}
	Info.Printf("block_proc: %v %v ; %v transactions\n",bnum,block_hash.String(),num_transactions)
	if update_last_block {
		storage.Set_last_block_num(bnum)
	}
	return nil
}
func process_tx_event_log(agtx *AugurTx,log *types.Log) {
	var err error
	var eel EthereumEventLog
	eel.BlockNum = agtx.BlockNum
	eel.TxId = agtx.TxId
	eel.ContractAddress = log.Address.String()
	eel.Topic0_Sig = hex.EncodeToString(log.Topics[0][0:4])
	eel.RlpLog, err = rlp.EncodeToBytes(log)
	if err != nil {
		Info.Printf("Couldn't RLP-encode log : %v\n",err)
		os.Exit(1)
	}
	event_id,contract_aid := storage.Insert_tx_event_log(&eel)

	for pos,topic := range log.Topics {
		var eet EthereumEventTopic
		eet.BlockNum = eel.BlockNum
		eet.TxId = eel.TxId
		eet.EventLogId = event_id
		eet.ContractAid = contract_aid
		eet.Pos = pos
		eet.Value = hex.EncodeToString(topic.Bytes())
		storage.Insert_event_log_topic(&eet)
	}
}
func tx_insert_if_needed(agtx *AugurTx) {
	if agtx.TxId == 0 {
		agtx.TxId=storage.Insert_transaction(agtx)
	}
}
