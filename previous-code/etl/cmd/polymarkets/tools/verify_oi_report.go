package main

import (
	"os"
	"log"
	"sort"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
const (
	USDC_ADDR				string = "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"
)
var (
	storage *SQLStorage

	Info    *log.Logger

	condition_id			string
	market_id				int64
	contract_aid			int64

	usdc_adr				common.HexToAddress(USDC_ADDR)
)
type TokTransfs struct {
	SingleTransfers		[]Pol_ERC_Transfer
	BatchTransfers		[]Pol_ERC_Transfer
	SingleId			int64
	BatchId				int64
	SingleIdx			int64
	BatchIdx			int64
}
func get_next_record(tt *TokTransfers) (int64,Pol_ERC_Transfer) {

	next_index_single := tt.SingleIdx + 1
	next_index_batch := tt.BatchIdx + 1
	var next_id_single,next_id_batch int64
	var next_evtlog_single,next_evtlog_batch int64
	if len(tt.SingleTransfers) < next_index_single {
		next_evtlog_single = tt.SingleTransfers[next_index_single].EvtLogId
	}
	if len(tt.BatchTransfers) < next_index_batch {
		next_evtlog_batch = tt.BatchTransfers[next_index_batch].EvtLogId
	}
	if (next_evtlog_single > 0) && (next_evtlog_batch > 0) {
		if next_evtlog_single > next_evtlog_batch {
			tt.BatchIdx = next_index_batch
			return next_index_batch,tt.BatchTransfers[next_index_batch]
		} else {
			tt.SingleIdx = next_index_single
			return next_index_single,tt.BatchSingle[next_index_single]
		}
	} else {
		if next_evtlog_single > 0 {
			tt.SingleIdx = next_index_single
			return next_index_single,tt.BatchSingle[next_index_single]
		}
		if next_evtlog_batch > 0 {
			tt.BatchIdx = next_index_batch
			return next_index_batch,tt.BatchTransfers[next_index_batch]
		}
	}
}
func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v [market_id]\n")
		os.Exit(1)
	}
	var err error
	market_id,err = strconv.ParseInt(os.Args[1],10,64)
	if err != nil {
		fmt.Printf("Error parsing market id: %v\n",err)
		os.Exit(1)
	}

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(Info)

	condition_id = storage.Get_condition_id(market_id)
	contract_aid = storage.Get_fpmm_contract_aid(market_id)

	tokens := storage.Get_tokens_by_condition_id(condition_id)

	usdc_aid := storage.Nonfatal_lookup_address_id(

	data := make([]TokTransfs,0,8)
	for i:=0;i<len(tt_ids); i++ {
		token_id := tt_ids[i]
		var tt TokTransfs
		tt.SingleTransfers = storage.Get_token_erc1155_single_transfers(token_id)
		tt.BatchTransfers = storage.Get_token_erc1155_batch_transfers(token_id)
		data = append(data,tt)
	}
	var cur_single_id int64
	var cur_batch_id int64
	for i:=0;i<len(data);i++ {
		rec := get_next_record(data)
		fmt.Printf("%v   %v\n",rec.TxId,rec.TimeStamp,Amount)
		usdc_tr := storage.Get_poly_usdc_transfers(rec.TxId)
		for j:=0; j<len(usdc_tr);j++ {
			usdc_rec := usdc_tr[j]
			fmt.Printf("\tUSDC : $ %v\n",usdc_rec.Amount)
		}
	}

}
