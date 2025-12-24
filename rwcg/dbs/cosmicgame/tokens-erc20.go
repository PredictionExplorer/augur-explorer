package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"

	p "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/cosmicgame"
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
	query = "SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_chrono_warrior WHERE cst_amount > 0"
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
	query = "SELECT COALESCE(SUM(cst_amount), 0), COALESCE(SUM(cst_amount)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_chrono_warrior WHERE winner_aid=$1 AND cst_amount > 0"
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
