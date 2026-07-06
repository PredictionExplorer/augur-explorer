package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_cosmic_token_holders() []p.CGCosmicTokenHolderRec {

	var query string
	query = "SELECT "+
				"o.owner_aid,"+
				"oa.addr,"+
				"o.cur_balance,"+
				"o.cur_balance/1e18 " +
			"FROM "+sw.S.SchemaName()+".cg_costok_owner o "+
				"LEFT JOIN address oa ON o.owner_aid=oa.address_id "+
			"ORDER BY o.cur_balance DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicTokenHolderRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicTokenHolderRec
		err=rows.Scan(
			&rec.OwnerAid,
			&rec.OwnerAddr,
			&rec.Balance,
			&rec.BalanceFloat,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_cosmic_token_statistics() p.CGCosmicTokenStats {

	var stats p.CGCosmicTokenStats
	
	// Total supply and holder count
	var query string
	query = "SELECT "+
				"COUNT(*) as holder_count, "+
				"COALESCE(SUM(cur_balance), 0) as total_supply, "+
				"COALESCE(SUM(cur_balance)/1e18, 0) as total_supply_eth "+
			"FROM "+sw.S.SchemaName()+".cg_costok_owner "+
			"WHERE cur_balance > 0"
	
	row := sw.S.Db().QueryRow(query)
	err := row.Scan(
		&stats.TotalHolders,
		&stats.TotalSupply,
		&stats.TotalSupplyEth,
	)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error in Get_cosmic_token_statistics: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Bidding rewards (CST ERC20 given for bidding)
	query = "SELECT COALESCE(SUM(cst_reward), 0), COALESCE(SUM(cst_reward)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_bid WHERE cst_reward > 0"
	row = sw.S.Db().QueryRow(query)
	err = row.Scan(&stats.EarnedFromBidding, &stats.EarnedFromBiddingEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Note: Staking rewards removed - those are ETH rewards for staking ERC721 NFTs, not CST ERC20
	
	// Marketing rewards (CST ERC20)
	query = "SELECT COALESCE(SUM(amount), 0), COALESCE(SUM(amount)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_mkt_reward"
	row = sw.S.Db().QueryRow(query)
	err = row.Scan(&stats.DistributedToMarketers, &stats.DistributedToMarketersEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Main prizes
	query = "SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_prize_claim WHERE cst_amount > 0"
	row = sw.S.Db().QueryRow(query)
	err = row.Scan(&stats.GivenAsMainPrizes, &stats.GivenAsMainPrizesEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Raffle prizes
	query = "SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_raffle_nft_prize WHERE cst_amount > 0"
	row = sw.S.Db().QueryRow(query)
	err = row.Scan(&stats.GivenAsRafflePrizes, &stats.GivenAsRafflePrizesEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Chrono warrior prizes
	query = "SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_chrono_warrior_prize WHERE cst_amount > 0"
	row = sw.S.Db().QueryRow(query)
	err = row.Scan(&stats.GivenAsChronoWarriorPrizes, &stats.GivenAsChronoWarriorPrizesEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Consumed in bids
	query = "SELECT COALESCE(total_cst_consumed, 0), COALESCE(total_cst_consumed/1e18, 0) FROM "+sw.S.SchemaName()+".cg_glob_stats LIMIT 1"
	row = sw.S.Db().QueryRow(query)
	err = row.Scan(&stats.ConsumedInBids, &stats.ConsumedInBidsEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Transfer counts by type
	query = "SELECT "+
				"COUNT(CASE WHEN otype = 1 THEN 1 END) as mints, "+
				"COUNT(CASE WHEN otype = 2 THEN 1 END) as burns, "+
				"COUNT(CASE WHEN otype = 0 THEN 1 END) as transfers "+
			"FROM "+sw.S.SchemaName()+".cg_erc20_transfer"
	row = sw.S.Db().QueryRow(query)
	err = row.Scan(&stats.TotalMints, &stats.TotalBurns, &stats.TotalTransfers)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Top 10 holders
	query = "SELECT "+
				"o.owner_aid, "+
				"a.addr, "+
				"o.cur_balance, "+
				"o.cur_balance/1e18 "+
			"FROM "+sw.S.SchemaName()+".cg_costok_owner o "+
				"LEFT JOIN address a ON o.owner_aid = a.address_id "+
			"WHERE o.cur_balance > 0 "+
			"ORDER BY o.cur_balance DESC "+
			"LIMIT 10"
	
	rows, err := sw.S.Db().Query(query)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	
	stats.TopHolders = make([]p.CGCosmicTokenHolderRec, 0, 10)
	for rows.Next() {
		var holder p.CGCosmicTokenHolderRec
		err = rows.Scan(
			&holder.OwnerAid,
			&holder.OwnerAddr,
			&holder.Balance,
			&holder.BalanceFloat,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		// Calculate percent of supply
		if stats.TotalSupplyEth > 0 {
			holder.PercentOfSupply = (holder.BalanceFloat / stats.TotalSupplyEth) * 100.0
		}
		stats.TopHolders = append(stats.TopHolders, holder)
	}
	
	return stats
}
func (sw *SQLStorageWrapper) Get_user_cosmic_token_summary(user_aid int64) p.CGUserCosmicTokenSummary {

	var summary p.CGUserCosmicTokenSummary
	summary.UserAid = user_aid
	
	// Current balance
	var query string
	query = "SELECT COALESCE(cur_balance, 0), COALESCE(cur_balance/1e18, 0) FROM "+sw.S.SchemaName()+".cg_costok_owner WHERE owner_aid=$1"
	row := sw.S.Db().QueryRow(query, user_aid)
	err := row.Scan(&summary.CurrentBalance, &summary.CurrentBalanceEth)
	if err != nil {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		summary.CurrentBalance = "0"
		summary.CurrentBalanceEth = 0
	}
	
	// Earned from bidding rewards (CST ERC20)
	query = "SELECT COALESCE(SUM(cst_reward), 0), COALESCE(SUM(cst_reward)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_bid WHERE bidder_aid=$1 AND cst_reward > 0"
	row = sw.S.Db().QueryRow(query, user_aid)
	err = row.Scan(&summary.EarnedFromBidding, &summary.EarnedFromBiddingEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Earned from main prizes (CST ERC20)
	query = "SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_prize_claim WHERE winner_aid=$1 AND cst_amount > 0"
	row = sw.S.Db().QueryRow(query, user_aid)
	err = row.Scan(&summary.EarnedFromMainPrizes, &summary.EarnedFromMainPrizesEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Earned from raffle prizes
	query = "SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_raffle_nft_prize WHERE winner_aid=$1 AND cst_amount > 0"
	row = sw.S.Db().QueryRow(query, user_aid)
	err = row.Scan(&summary.EarnedFromRafflePrizes, &summary.EarnedFromRafflePrizesEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Earned from chrono warrior
	query = "SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_chrono_warrior_prize WHERE winner_aid=$1 AND cst_amount > 0"
	row = sw.S.Db().QueryRow(query, user_aid)
	err = row.Scan(&summary.EarnedFromChronoWarrior, &summary.EarnedFromChronoWarriorEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Earned from marketing
	query = "SELECT COALESCE(SUM(amount), 0), COALESCE(SUM(amount)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_mkt_reward WHERE marketer_aid=$1"
	row = sw.S.Db().QueryRow(query, user_aid)
	err = row.Scan(&summary.EarnedFromMarketing, &summary.EarnedFromMarketingEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Consumed in bids (CST spent for bidding)
	query = "SELECT COALESCE(SUM(cst_price), 0), COALESCE(SUM(cst_price)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_bid WHERE bidder_aid=$1 AND cst_price > 0"
	row = sw.S.Db().QueryRow(query, user_aid)
	err = row.Scan(&summary.ConsumedInBids, &summary.ConsumedInBidsEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Calculate total CST earned
	summary.TotalEarnedEth = summary.EarnedFromBiddingEth + 
							summary.EarnedFromMainPrizesEth + 
							summary.EarnedFromRafflePrizesEth + 
							summary.EarnedFromChronoWarriorEth + 
							summary.EarnedFromMarketingEth
	
	// Calculate net flow
	summary.NetCSTFlowEth = summary.TotalEarnedEth - summary.ConsumedInBidsEth
	
	// Activity counts
	query = "SELECT "+
				"COUNT(CASE WHEN otype = 1 THEN 1 END) as mints, "+
				"COUNT(CASE WHEN otype = 2 THEN 1 END) as burns, "+
				"COUNT(*) as total_transfers "+
			"FROM "+sw.S.SchemaName()+".cg_erc20_transfer "+
			"WHERE from_aid=$1 OR to_aid=$1"
	row = sw.S.Db().QueryRow(query, user_aid)
	err = row.Scan(&summary.NumMints, &summary.NumBurns, &summary.NumTransfers)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Note: Staking info removed - that's about ERC721 NFT staking (not ERC20 CST)
	// Staking rewards are ETH, not CST tokens
	
	return summary
}

// Get_cosmic_token_total_supply_history_by_bid returns one row per bid with net CST supply change
// (cst_reward mint minus cst_price burn on CST bids) and running totals computed in SQL.
func (sw *SQLStorageWrapper) Get_cosmic_token_total_supply_history_by_bid() []p.CGTotalSupplyHistoryRec {

	query := "SELECT " +
		"b.evtlog_id, b.bid_type, COALESCE(ba.addr, ''), b.block_num, COALESCE(t.id, 0), COALESCE(t.tx_hash, ''), " +
		"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT, b.time_stamp, " +
		"GREATEST(COALESCE(b.cst_reward, 0), 0)::text, " +
		"GREATEST(COALESCE(b.cst_reward, 0), 0)/1e18, " +
		"(CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)::text, " +
		"(CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)/1e18, " +
		"(GREATEST(COALESCE(b.cst_reward, 0), 0) - CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)::text, " +
		"(GREATEST(COALESCE(b.cst_reward, 0), 0) - CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)/1e18, " +
		"SUM(GREATEST(COALESCE(b.cst_reward, 0), 0) - CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END) " +
		"OVER (ORDER BY b.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)::text, " +
		"SUM((GREATEST(COALESCE(b.cst_reward, 0), 0) - CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)/1e18) " +
		"OVER (ORDER BY b.id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) " +
		"FROM " + sw.S.SchemaName() + ".cg_bid b " +
		"LEFT JOIN " + sw.S.SchemaName() + ".address ba ON b.bidder_aid = ba.address_id " +
		"LEFT JOIN " + sw.S.SchemaName() + ".transaction t ON t.id = b.tx_id " +
		"ORDER BY b.id"

	rows, err := sw.S.Db().Query(query)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	records := make([]p.CGTotalSupplyHistoryRec, 0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGTotalSupplyHistoryRec
		err = rows.Scan(
			&rec.BidInfoId,
			&rec.BidType,
			&rec.BidderAddr,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
			&rec.MintAmount,
			&rec.MintAmountEth,
			&rec.BurnAmount,
			&rec.BurnAmountEth,
			&rec.Amount,
			&rec.AmountEth,
			&rec.TotalSupply,
			&rec.TotalSupplyEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		rec.Tx.EvtLogId = rec.BidInfoId
		records = append(records, rec)
	}
	return records
}

// Get_cosmic_token_total_supply_history_by_date returns daily aggregates of CST supply change
// between fromDate and toDate (inclusive), with running totals over all history up to each day.
func (sw *SQLStorageWrapper) Get_cosmic_token_total_supply_history_by_date(fromDate, toDate string) []p.CGTotalSupplyHistoryByDateRec {

	query := "WITH daily AS (" +
		"SELECT " +
		"DATE(b.time_stamp) AS bid_date, " +
		"COUNT(*)::bigint AS num_bids, " +
		"SUM(GREATEST(COALESCE(b.cst_reward, 0), 0)) AS mint_amt, " +
		"SUM(CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END) AS burn_amt, " +
		"SUM(GREATEST(COALESCE(b.cst_reward, 0), 0) - CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END) AS net_amt " +
		"FROM " + sw.S.SchemaName() + ".cg_bid b " +
		"GROUP BY DATE(b.time_stamp) " +
		"), with_totals AS (" +
		"SELECT " +
		"d.bid_date, " +
		"d.num_bids, " +
		"d.mint_amt, d.burn_amt, d.net_amt, " +
		"SUM(d.net_amt) OVER (ORDER BY d.bid_date ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS total_supply " +
		"FROM daily d " +
		") " +
		"SELECT " +
		"TO_CHAR(w.bid_date, 'YYYYMMDD'), " +
		"EXTRACT(EPOCH FROM w.bid_date)::BIGINT, " +
		"w.bid_date::text, " +
		"w.num_bids, " +
		"w.mint_amt::text, w.mint_amt/1e18, " +
		"w.burn_amt::text, w.burn_amt/1e18, " +
		"w.net_amt::text, w.net_amt/1e18, " +
		"w.total_supply::text, w.total_supply/1e18 " +
		"FROM with_totals w " +
		"WHERE w.bid_date >= TO_DATE($1, 'YYYYMMDD') AND w.bid_date <= TO_DATE($2, 'YYYYMMDD') " +
		"ORDER BY w.bid_date"

	rows, err := sw.S.Db().Query(query, fromDate, toDate)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	records := make([]p.CGTotalSupplyHistoryByDateRec, 0, 64)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGTotalSupplyHistoryByDateRec
		err = rows.Scan(
			&rec.Date,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.NumBids,
			&rec.MintAmount,
			&rec.MintAmountEth,
			&rec.BurnAmount,
			&rec.BurnAmountEth,
			&rec.Amount,
			&rec.AmountEth,
			&rec.TotalSupply,
			&rec.TotalSupplyEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		records = append(records, rec)
	}
	return records
}
