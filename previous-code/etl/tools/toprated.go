package main

import (
	"os"
	"log"
	"sort"
	//"fmt"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
var (
	storage *SQLStorage

	fill_order_id int64 = 0			// during event processing, holds id of record in mktord from Fill evt
	market_order_id int64 = 0

	Info    *log.Logger
)
func update_profit_ranks(records []RankStats) {

	num_recs := len(records)
	for i:=0 ; i < num_recs  ; i++ {
		record := &records[i]
		rank_value := (float64(i)/float64(num_recs))*100.0 + 1.0
//		fmt.Printf("rank for Profit of %v is %v pl=%v\n",record.EoaAid,rank_value,record.ProfitLoss)
		storage.Update_top_profit_rank(record.Aid,rank_value,record.ProfitLoss)
	}
}
func update_trade_ranks(records []RankStats) {

	num_recs := len(records)
	for i:=0 ; i < num_recs  ; i++ {
		record := &records[i]
		rank_value := (float64(i)/float64(num_recs))*100.0 + 1.0
//		fmt.Printf("rank for TotalTrades of %v is %v trades=%v\n",record.EoaAid,rank_value,record.TotalTrades)
		storage.Update_top_total_trades_rank(record.Aid,rank_value,record.TotalTrades)
	}
}
func update_volume_ranks(records []RankStats) {

	num_recs := len(records)
	for i:=0 ; i < num_recs  ; i++ {
		record := &records[i]
		rank_value := (float64(i)/float64(num_recs))*100.0 + 1.0
//		fmt.Printf("rank for Volume of %v is %v volume=%v\n",record.EoaAid,rank_value,record.VolumeTraded)
		storage.Update_top_volume_rank(record.Aid,rank_value,record.VolumeTraded)
	}
}
func main() {
	// Design Notes:
	//		We don't want to load production environment with a heavy update query as statistics is
	//		not a critical task and is executed once a day
	//		All updates are done using primary key, the sorting is done in memory space of the
	//			Client program
	//		This process is single threaded and light-weight but will take a long time to finish
	//		ToDo: possibly insert a sleep() with millisecond timeout between updates to avoid
	//				server overload

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)
	records := storage.Get_ranking_data_for_all_users()

	sort.SliceStable(records,func(i,j int) bool {
		return records[j].TotalTrades < records[i].TotalTrades
	})
	update_trade_ranks(records)

	sort.SliceStable(records,func(i,j int) bool {
		return records[j].ProfitLoss < records[i].ProfitLoss
	})
	update_profit_ranks(records)

	sort.SliceStable(records,func(i,j int) bool {
		return records[j].VolumeTraded < records[i].VolumeTraded
	})
	update_volume_ranks(records)
}
