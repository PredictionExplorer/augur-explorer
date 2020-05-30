package main

import (
	"sort"
)
var (
	storage *SQLStorage

	fill_order_id int64 = 0			// during event processing, holds id of record in mktord from Fill evt
	market_order_id int64 = 0
)
func update_profit_ranks(records []RankStats) {

	num_recs := len(records)
	for i:=0 ; i < num_recs  ; i++ {
		record := &records[i]
		rank_value := float64(i/100) + 1
		storage.update_top_profit_rank(record.EoaAid,rank_value)
	}
}
func update_trade_ranks(records []RankStats) {

	num_recs := len(records)
	for i:=0 ; i < num_recs  ; i++ {
		record := &records[i]
		rank_value := float64(i/100) + 1
		storage.update_top_total_trades_rank(record.EoaAid,rank_value)
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

	storage = connect_to_storage()
	records := storage.get_ranking_data_for_all_users()
	sort.SliceStable(records,func(i,j int) bool {
		return records[i].TotalTrades < records[i].TotalTrades
	})
	update_trade_ranks(records)
	sort.SliceStable(records,func(i,j int) bool {
		return records[i].ProfitLoss < records[i].ProfitLoss
	})
	update_profit_ranks(records)
}
