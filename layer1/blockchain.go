package layer1

import (
	"time"
	"bytes"
	"math/big"
	"os"
	"errors"
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
var (
	addrs sync.Map
)
func lookup_or_insert_addr(addr common.Address) int64 {

	var aid int64
	var exists bool
	var result interface{}
	result,exists = addrs.Load(addr)
	if exists {
		aid = result.(int64)
	} else {
		var err error
		aid,is_contract,err = storage.Layer1_lookup_address_id(addr.String())
		if err != nil {
			aid = storage.Layer1_insert_address(addr.String(),is_contract)
			addrs.Store(addr,aid)
		}
	}
	return aid
}
func roll_back_blocks(diverging_block *types.Header) error {
	// Finds the block from which the fork started
	var err error
	var bhash common.Hash
	bhash,diverging_block,_, err = get_full_block(diverging_block.Number.Int64())
	if err != nil {
		return errors.New(fmt.Sprintf("During chainsplit an error getting HeaderByHash happened: %v\n",err))
	}
	starting_block_num := diverging_block.Number.Int64()
	block_num:=starting_block_num
	block_hash:=bhash.String()
	Info.Printf("roll_back_blocks(): block_num = %v , block_hash %v",block_num,block_hash)
	Info.Printf("\t\t\tparent_hash %v\n",diverging_block.ParentHash.String())
	for {
		my_block_num,err := storage.Layer1_get_block_num_by_hash(block_hash)
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
			storage.Layer1_chainsplit_delete_blocks(my_block_num)
			storage.Layer1_set_last_block_num(my_block_num)
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
		if err != nil {
			return errors.New(fmt.Sprintf("During chainsplit an error getting BlockByNumber happened: %v\n",err))
		}
		block_num = diverging_block.Number.Int64()
		block_hash = bhash.String()
		Info.Printf("Current block has been set to number %v , hash = %v\n",block_num,diverging_block.Hash().String())
	}
	return errors.New("Chainsplit fix: Undefined behaviour")
}
func add_address_stat_entry(addr_stats_log []AddrStatsLog,block_num,tx_index,aid int64) []AddrStatsLog {

	var entry AddrStatsLog
	entry.BlockNum = block_num
	entry.TxIndex = tx_index
	entry.Aid = aid
	addr_stats_log = append(addr_stats_log,entry)
	return addr_stats_log
}
func process_transactions(bnum int64,transactions []*AugurTx,receipt_calls []*receiptCallResult,block_receipts []types.Receipt,extra_info []ReceiptExtraInfo) (*big.Int,*big.Int,error) {
	//	if receipt_calls is not nil then the old slow getTrasnactionReceipt call is used
	//	if block_receipts is not nil then we are using new getBlockReceipts RPC call

	total_eth := big.NewInt(0)
	total_fees := big.NewInt(0)
	for tnum,agtx := range transactions {
		if agtx.From == "0x0000000000000000000000000000000000000000" {
			continue // this is Polygon's automatic transaction
		}
		agtx.TxIndex = int32(tnum)
		tmp_log_slice := make([]AddrStatsLog,0,2)
		from_addr := common.HexToAddress(agtx.From)
		from_aid,_:= lookup_or_insert_addr(from_addr)
		tmp_log_slice = add_address_stat_entry(tmp_log_slice,agtx.BlockNum,int64(agtx.TxIndex),from_aid)
		to_addr := common.HexToAddress(agtx.To)
		to_aid,_:= lookup_or_insert_addr(to_addr)
		tmp_log_slice = add_address_stat_entry(tmp_log_slice,agtx.BlockNum,int64(agtx.TxIndex),to_aid)

		storage.Bigstats_insert_all_addr_stat_logs(tmp_log_slice)
		var rcpt *types.Receipt = nil
		var rcpt_extra *ReceiptExtraInfo = nil
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
				return total_eth,total_fees,receipt_calls[tnum].err
			}
			rcpt = receipt_calls[tnum].receipt
			rcpt_extra = receipt_calls[tnum].extra
		} else {
			// receipts were fetched using eth_getBlockReceipts, we only need to reference the receipt
			rcpt = &block_receipts[tnum]
			rcpt_extra = &extra_info[tnum]
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
			return total_eth,total_fees,errors.New("Block changed during processing")
		}
		gas_price := big.NewInt(0)
		gas_price.SetString(agtx.GasPrice,10)
		var tx_short TxShort
		tx_short.BlockNum = bnum
		tx_short.TxIndex = int64(agtx.TxIndex)
		if rcpt_extra == nil {
			Info.Printf("Receipt Extra info struct is nil for tx %v\n",agtx.TxHash)
			Error.Printf("Receipt Extra info struct is nil for tx %v\n",agtx.TxHash)
			tx_short.TxFee = "0"
		} else {
			tx_fee := big.NewInt(int64(rcpt.GasUsed))
			//Info.Printf("tnum=%v: Multiplying gas used %v by gas price %v\n",tnum,tx_fee.String(),rcpt_extra.EffectiveGasPrice.String())
			tx_fee.Mul(tx_fee,rcpt_extra.EffectiveGasPrice)
			tx_short.TxFee = tx_fee.String()
			total_fees.Add(total_fees,tx_fee)
		}
		storage.Bigstats_insert_transaction(&tx_short)	// at this point we are sure Tx is without error
		transaction_hash := common.HexToHash(agtx.TxHash)
		if !bytes.Equal(rcpt.TxHash.Bytes(),transaction_hash.Bytes()) { // can be removed later
			Error.Printf("Receipt's hash doesn't match Tx hash, aborting (tx_hash=%v)\n",agtx.TxHash)
			os.Exit(1)
		}
		agtx.TxId = 0
		if agtx.CtrctCreate == true {
			agtx.To = rcpt.ContractAddress.String()
		}
		big_value := big.NewInt(0)
		big_value.SetString(agtx.Value,10)
		total_eth.Add(total_eth,big_value)
		agtx.GasUsed = int64(rcpt.GasUsed)
		agtx.NumLogs = int32(len(rcpt.Logs))
		logs_to_insert := extract_addresses_from_event_logs(agtx,rcpt.Logs)
		if len(logs_to_insert) > 0 {
			storage.Bigstats_insert_all_addr_stat_logs(logs_to_insert)
		}
	}
	return total_eth,total_fees,nil
}
func process_block(bnum int64,update_last_block bool,no_chainsplit_check bool,norollback bool) error {

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
		Info.Printf("Retrieved block number %v but Block object contains another number (%v)\n",bnum,header.Number.Int64())
		Error.Printf("Retrieved block number %v but Block object contains another number (%v)\n",bnum,header.Number.Int64())
		return errors.New("Block object inconsistency")
	}
	storage.Bigstats_block_delete_with_everything(big_bnum.Int64())
	var receipt_calls []*receiptCallResult = nil
	var block_receipts []types.Receipt = nil
	var extra_fields []ReceiptExtraInfo
	if *use_block_receipts_call {
		block_receipts,extra_fields,err = get_block_receipts_v2(block_hash)
		if err != nil {
			Error.Printf("Error getting receipts of the block: %v\n",err)
			return err
		}

	} else {
		receipt_calls = make([]*receiptCallResult,num_transactions,num_transactions)
		for i,tx := range transactions {
			hash := common.HexToHash(tx.TxHash)
			go get_receipt_async_custom_rpc(i,hash,&receipt_calls)
		}
	}
	err = storage.Bigstats_insert_block(block_hash_str,header,num_transactions,no_chainsplit_check)
	if err != nil {
		if !norollback {
			err = roll_back_blocks(header)
		}
		return err
	}
	if num_transactions == 0 {
		if update_last_block {
			storage.Bigstats_set_last_block_num(bnum)
		}
		return nil
	}
	err := process_transactions(bnum,transactions,receipt_calls,block_receipts,extra_fields)
	if update_last_block {
		storage.Bigstats_set_last_block_num(bnum)
	}
	return nil
}
func extract_addresses_from_event_logs(agtx *AugurTx,logs []*types.Log) []AddrStatsLog {

	output := make([]AddrStatsLog,0,64)
	num_logs := len(logs)
	for i:=0 ; i < num_logs ; i++ {
		log := logs[i]
		var astats AddrStatsLog
		var is_contract bool
		astats.BlockNum = agtx.BlockNum
		astats.TxIndex = int64(agtx.TxIndex)
		astats.Aid,is_contract = lookup_or_insert_addr(log.Address)
		astats.Aid = storage.Bigstats_lookup_or_create_address(log.Address.String(),is_contract)
		output = append(output,astats)
	}

	return output
}
