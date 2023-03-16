package biddingwar

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/biddingwar"
	//. "github.com/PredictionExplorer/augur-explorer/dbs"
)
func (sw *SQLStorageWrapper) Get_biddingwar_statistics() p.BWStatistics {

	var stats p.BWStatistics
	var query string
	query = "SELECT "+
				"num_vol_donations, "+
				"vol_donations_total/1e18 as voluntary_donations_sum,"+
				"num_bids," +
				"cur_num_bids,"+
				"num_wins, "+
				"num_rwalk_used "+
			"FROM bw_glob_stats LIMIT 1"

	row := sw.S.Db().QueryRow(query)
	var err error
	err=row.Scan(
		&stats.NumVoluntaryDonations,
		&stats.SumVoluntaryDonationsEth,
		&stats.TotalBids,
		&stats.CurNumBids,
		&stats.TotalPrizes,
		&stats.NumRwalkTokensUsed,
	)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_biddingwar_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var null_bidders sql.NullInt64
	query = "SELECT "+
				"COUNT(*) AS total "+
			"FROM bw_bidder"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(&null_bidders)
	if null_bidders.Valid  { stats.NumUniqueBidders = uint64(null_bidders.Int64) }
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_biddingwar_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var null_winners sql.NullInt64
	var null_sum_wei sql.NullString
	var null_sum_eth sql.NullFloat64
	query = "SELECT "+
				"COUNT(*) AS total,"+
				"SUM(prizes_sum) AS sum_wei,"+
				"SUM(prizes_sum)/1e18 AS sum_eth "+
				"FROM bw_winner"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_winners,
		&null_sum_wei,
		&null_sum_eth,
	)
	if null_winners.Valid { stats.NumUniqueWinners = uint64(null_winners.Int64) }
	if null_sum_wei.Valid { stats.TotalPrizesPaidAmountWei = null_sum_wei.String }
	if null_sum_eth.Valid { stats.TotalPrizesPaidAmountEth = null_sum_eth.Float64 }
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_biddingwar_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	
	return stats
}
