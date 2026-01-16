package main

import (
	"os"
	"time"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func update_non_augur_flag(exit_chan chan bool) {


	last_block_chain,last_block_processed,err := storage.Get_status_not_augur_block_num()
	if err != nil {
		Info.Printf("No records in status table, aborting update non-augur flag process\n")
		return
	}
	from_block := last_block_processed + 1
	/*diff := last_block_chain - from_block
	if diff > 10000 {
		diff = 10000
	}*/
	to_block := last_block_chain - 1
	//to_block := from_block + diff
	Info.Printf("status: from_block=%v, to_block=%v\n",from_block,to_block)
	records := storage.Get_shares_minted_burned_in_block_range("aa_shares_minted",from_block,to_block)
	for i:=0; i<len(records) ; i++ {
		rec := records[i]
		Info.Printf("status: sharesburned rec = %+v\n",rec)
		if (rec.SharesSwappedId == 0) && (rec.LiquidityId==0) && (rec.BalancerId==0) {
			Info.Printf("Detected SharesMinted event_id=%v as non-augur\n",rec.RecordId)
			storage.Insert_not_augur_mark(rec.RecordId,RecTypeMint)
		}
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
	}
	Info.Printf("status: finished SharesMinted\n")
	records = storage.Get_shares_minted_burned_in_block_range("aa_shares_burned",from_block,to_block)
	for i:=0; i<len(records) ; i++ {
		rec := records[i]
		Info.Printf("status: sharesminted rec = %+v\n",rec)
		if (rec.SharesSwappedId == 0) && (rec.LiquidityId==0) && (rec.BalancerId==0)  {
			Info.Printf("Detected SharesBurned event_id=%v as non-augur\n",rec.RecordId)
			storage.Insert_not_augur_mark(rec.RecordId,RecTypeBurn)
		}
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
	}
	Info.Printf("status: finished SharesBurned\n")
	records_swap := storage.Get_balancer_swaps_for_augur_markets(from_block,to_block)
	for i:=0; i<len(records_swap); i++ {
		rec := records_swap[i]
		Info.Printf("status: balancer rec = %+v\n",rec)
		if (rec.SharesSwappedId == 0) && (rec.LiquidityId==0) {
			Info.Printf("Detected Balancer swap event_id=%v as non-augur\n",rec.RecordId)
			storage.Insert_not_augur_mark(rec.RecordId,RecTypeBalancer)
		}
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
	}
	Info.Printf("status: finished Balancer swaps\n")
	Info.Printf("status: updating last block to %v\n",to_block)
	storage.Update_status_not_augur_block_num(to_block)
}
func update_non_augur_flag_manager(exit_chan chan bool) {

	for {
		update_non_augur_flag(exit_chan)
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		time.Sleep(10 * time.Second)
	}
}
func update_non_augur_flag_for_erc20_transfers(exit_chan chan bool) {


	last_block_chain,last_block_processed,err := storage.Get_status_erc20_transf_not_augur_block_num()
	if err != nil {
		Info.Printf("No records in status table for erc20 transfers, aborting\n")
		return
	}
	last_erc20_evtlog_id := storage.Get_last_erc20_evt_id()

	Info.Printf("last_block_chain=%v, last_block_processed=%v\n",last_block_chain,last_block_processed)
	proc_status := storage.Get_arbitrum_augur_processing_status()
	if proc_status.LastEvtId > last_erc20_evtlog_id {
		Info.Printf("ERC20 token database is older than my local DB (check ERC20 daemon is running)\n")
		return // erc20 table has older data than non-augur tables
	}
	from_block := last_block_processed + 1
	/*diff := last_block_chain - from_block
	if diff > 10000 {
		diff = 10000
	}*/
	to_block := last_block_chain - 1
	//to_block := from_block + diff
	Info.Printf("status: from_block=%v, to_block=%v\n",from_block,to_block)
	records := storage.Get_erc20transfers_for_augur_markets(from_block,to_block)
	for i:=0; i<len(records) ; i++ {
		rec := records[i]
		Info.Printf("status: rec = %+v\n",rec)
		if (rec.SharesSwappedId == 0) && (rec.LiquidityId==0) && (rec.SharesBurnedId ==0) &&
			(rec.SharesMintedId==0 ) && (rec.BalancerId==0) && (rec.WinningsClaimedId ==0) &&
			(rec.BalExitId == 0 ) {
			Info.Printf("Detected ERC@0 event_id=%v as non-augur\n",rec.RecordId)
			storage.Insert_not_augur_mark(rec.RecordId,RecTypeERC20)

		}
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
	}
	Info.Printf("status: updating last block to %v\n",to_block)
	storage.Update_status_erc20_not_augur_block_num(to_block)
}
func update_erc20_non_augur_flag_manager(exit_chan chan bool) {

	for {
		update_non_augur_flag_for_erc20_transfers(exit_chan)
		select {
			case exit_flag := <-exit_chan:
				if exit_flag {
					Info.Println("Exiting by user request.")
					os.Exit(0)
				}
			default:
		}
		time.Sleep(10 * time.Second)
	}
}
