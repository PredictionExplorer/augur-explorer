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
func (sw *SQLStorageWrapper) Get_bids(offset,limit int) []p.BwBidRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"b.evtlog_id,"+
				"b.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT,"+
				"b.time_stamp,"+
				"b.bidder_aid,"+
				"ba.addr,"+
				"b.bid_price,"+
				"b.bid_price/1e18 bid_price_eth, " +
				"b.erc20_amount,"+
				"b.erc20_amount/1e18 erc20_amount_eth "+
			"FROM "+sw.S.SchemaName()+".bw_bid b "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address ba ON b.bidder_aid=ba.address_id "+
			"ORDER BY id OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.BwBidRec,0, limit)
	defer rows.Close()
	for rows.Next() {
		var rec p.BwBidRec
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BidderAid,
			&rec.BidderAddr,
			&rec.BidPrice,
			&rec.BidPriceEth,
			&rec.ERC20_Amount,
			&rec.ERC20_AmountEth,
		)
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_prize_claims(offset,limit int) []p.BwPrizeRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"p.evtlog_id,"+
				"p.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,"+
				"p.time_stamp,"+
				"p.winner_aid,"+
				"wa.addr,"+
				"p.amount, "+
				"p.amount/1e18 amount_eth, " +
				"p.prize_num,"+
				"m.seed "+
			"FROM "+sw.S.SchemaName()+".bw_prize_claim p "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
				"LEFT JOIN bw_mint_event m ON m.token_id=p.prize_num "+
			"ORDER BY id OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var null_seed sql.NullString
	records := make([]p.BwPrizeRec,0, limit)
	defer rows.Close()
	for rows.Next() {
		var rec p.BwPrizeRec
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.Amount,
			&rec.AmountEth,
			&rec.PrizeNum,
			&null_seed,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_seed.Valid { rec.Seed = null_seed.String } else {rec.Seed = "???"}
		rec.TokenId = rec.PrizeNum 
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_bid_info(evtlog_id int64) (bool,p.BwBidRec) {

	var rec p.BwBidRec
	var query string
	query = "SELECT " +
				"b.evtlog_id,"+
				"b.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT,"+
				"b.time_stamp,"+
				"b.bidder_aid,"+
				"ba.addr,"+
				"b.bid_price,"+
				"b.bid_price/1e18 bid_price_eth, " +
				"b.erc20_amount,"+
				"b.erc20_amount/1e18 erc20_amount_eth "+
			"FROM "+sw.S.SchemaName()+".bw_bid b "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address ba ON b.bidder_aid=ba.address_id "+
			"WHERE b.evtlog_id=$1"

	row := sw.S.Db().QueryRow(query,evtlog_id)
	var err error
	err=row.Scan(
		&rec.EvtLogId,
		&rec.BlockNum,
		&rec.TxId,
		&rec.TxHash,
		&rec.TimeStamp,
		&rec.DateTime,
		&rec.BidderAid,
		&rec.BidderAddr,
		&rec.BidPrice,
		&rec.BidPriceEth,
		&rec.ERC20_Amount,
		&rec.ERC20_AmountEth,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	return true,rec
}
func (sw *SQLStorageWrapper) Get_prize_info(prize_num int64) (bool,p.BwPrizeRec) {

	var rec p.BwPrizeRec
	var query string
	query = "SELECT "+
				"p.evtlog_id,"+
				"p.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,"+
				"p.time_stamp,"+
				"p.winner_aid,"+
				"wa.addr,"+
				"p.amount, "+
				"p.amount/1e18 amount_eth, " +
				"p.prize_num,"+
				"m.seed "+
			"FROM "+sw.S.SchemaName()+".bw_prize_claim p "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
				"LEFT JOIN bw_mint_event m ON m.token_id=p.prize_num "+
			"WHERE p.prize_num=$1"

	row := sw.S.Db().QueryRow(query,prize_num)
	var null_seed sql.NullString
	err := row.Scan(
		&rec.EvtLogId,
		&rec.BlockNum,
		&rec.TxId,
		&rec.TxHash,
		&rec.TimeStamp,
		&rec.DateTime,
		&rec.WinnerAid,
		&rec.WinnerAddr,
		&rec.Amount,
		&rec.AmountEth,
		&rec.PrizeNum,
		&null_seed,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_seed.Valid { rec.Seed = null_seed.String } else {rec.Seed = "???"}
	rec.TokenId = rec.PrizeNum 
	return true,rec
}
