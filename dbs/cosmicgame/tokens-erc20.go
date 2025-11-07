package cosmicgame

import (
	"os"
	"fmt"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
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
	
	// Bidding rewards
	query = "SELECT COALESCE(SUM(cst_reward), 0), COALESCE(SUM(cst_reward)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_bid WHERE cst_reward > 0"
	row = sw.S.Db().QueryRow(query)
	err = row.Scan(&stats.EarnedFromBidding, &stats.EarnedFromBiddingEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Staking rewards
	query = "SELECT COALESCE(SUM(reward), 0), COALESCE(SUM(reward)/1e18, 0) FROM "+sw.S.SchemaName()+".cg_st_reward"
	row = sw.S.Db().QueryRow(query)
	err = row.Scan(&stats.EarnedFromStakingNFTs, &stats.EarnedFromStakingNFTsEth)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	
	// Marketing rewards
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
