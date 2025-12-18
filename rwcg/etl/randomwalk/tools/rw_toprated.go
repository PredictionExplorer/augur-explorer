package main

import (
	"os"
	"log"
	"sort"

	"github.com/PredictionExplorer/augur-explorer/rwcg/dbs"
	rwdb "github.com/PredictionExplorer/augur-explorer/rwcg/dbs/randomwalk"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/primitives"
)
var (
	storagew *rwdb.SQLStorageWrapper
	Info     *log.Logger
)
func update_profit_ranks(records []RankStats) {

	num_recs := len(records)
	for i:=0 ; i < num_recs  ; i++ {
		record := &records[i]
		rank_value := (float64(i)/float64(num_recs))*100.0 + 1.0
		storagew.Update_randomwalk_top_profit_rank(record.Aid,rank_value,record.ProfitLoss)
	}
}
func update_trade_ranks(records []RankStats) {

	num_recs := len(records)
	for i:=0 ; i < num_recs  ; i++ {
		record := &records[i]
		rank_value := (float64(i)/float64(num_recs))*100.0 + 1.0
		storagew.Update_randomwalk_top_total_trades_rank(record.Aid,rank_value,record.TotalTrades)
	}
}
func update_volume_ranks(records []RankStats) {

	num_recs := len(records)
	for i:=0 ; i < num_recs  ; i++ {
		record := &records[i]
		rank_value := (float64(i)/float64(num_recs))*100.0 + 1.0
		storagew.Update_randomwalk_top_volume_rank(record.Aid,rank_value,record.VolumeTraded)
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
	storage := dbs.Connect_to_storage(Info)
	storagew = &rwdb.SQLStorageWrapper{S: storage}
	records := storagew.Get_randomwalk_ranking_data_for_all_users()

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
