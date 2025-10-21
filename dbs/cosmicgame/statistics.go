package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_cosmic_game_statistics() p.CGStatistics {

	var stats p.CGStatistics
	var query string
	query = "SELECT "+
				"num_vol_donations, "+
				"vol_donations_total/1e18 as voluntary_donations_sum,"+
				"num_cg_donations,"+
				"cg_donations_total/1e18,"+
				"direct_donations/1e18,"+
				"num_direct_donations,"+
				"num_withdrawals,"+
				"sum_withdrawals/1e18,"+
				"num_bids," +
				"cur_num_bids,"+
				"num_wins, "+
				"num_rwalk_used, "+
				"num_mints, "+
			"total_raffle_eth_deposits/1e18, "+
			"total_raffle_eth_withdrawn/1e18, "+
			"total_chrono_warrior_eth_deposits/1e18, "+
			"total_cst_given_in_prizes/1e18, "+
			"total_nft_donated,"+
				"num_bids_cst,"+
				"total_cst_consumed,"+
				"total_cst_consumed/1e18, "+
				"total_mkt_rewards,"+
				"total_mkt_rewards/1e18,"+
				"num_mkt_rewards "+
			"FROM cg_glob_stats LIMIT 1"

	row := sw.S.Db().QueryRow(query)
	var err error
	err=row.Scan(
		&stats.NumVoluntaryDonations,
		&stats.SumVoluntaryDonationsEth,
		&stats.NumCosmicGameDonations,
		&stats.SumCosmicGameDonationsEth,
		&stats.DirectDonationsEth,
		&stats.NumDirectDonations,
		&stats.NumWithdrawals,
		&stats.SumWithdrawals,
		&stats.TotalBids,
		&stats.CurNumBids,
		&stats.TotalPrizes,
		&stats.NumRwalkTokensUsed,
		&stats.NumCSTokenMints,
		&stats.TotalRaffleEthDeposits,
		&stats.TotalRaffleEthWithdrawn,
		&stats.TotalChronoWarriorEthDeposits,
		&stats.TotalCSTGivenInPrizes,
		&stats.TotalNFTDonated,
		&stats.NumBidsCST,
		&stats.TotalCSTConsumed,
		&stats.TotalCSTConsumedEth,
		&stats.TotalMktRewards,
		&stats.TotalMktRewardsEth,
		&stats.NumMktRewards,
	)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	var null_bidders sql.NullInt64
	query = "SELECT "+
				"COUNT(*) AS total "+
			"FROM cg_bidder " +
			"WHERE num_bids > 0 "
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(&null_bidders)
	if null_bidders.Valid  { stats.NumUniqueBidders = uint64(null_bidders.Int64) }
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	var null_winners sql.NullInt64
	var null_sum_wei sql.NullString
	var null_sum_eth sql.NullFloat64
	query = "SELECT "+
				"COUNT(*) AS total,"+
				"SUM(prizes_sum) AS sum_wei,"+
				"SUM(prizes_sum)/1e18 AS sum_eth "+
				"FROM cg_winner " +
				"WHERE prizes_count > 0"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_winners,
		&null_sum_wei,
		&null_sum_eth,
	)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	if null_winners.Valid { stats.NumUniqueWinners = uint64(null_winners.Int64) }
	if null_sum_wei.Valid { stats.TotalPrizesPaidAmountWei = null_sum_wei.String }
	if null_sum_eth.Valid { stats.TotalPrizesPaidAmountEth = null_sum_eth.Float64 }

	var null_donors sql.NullInt64
	query = "SELECT "+
				"COUNT(*) AS total,"+
				"SUM(total_eth_donated) AS sum_wei,"+
				"SUM(total_eth_donated)/1e18 AS sum_eth "+
				"FROM cg_donor " +
				"WHERE total_eth_donated > 0"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_donors,
		&null_sum_wei,
		&null_sum_eth,
	)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	if null_donors.Valid { stats.NumUniqueDonors = int64(null_donors.Int64) }
	if null_sum_wei.Valid { stats.TotalEthDonatedAmount = null_sum_wei.String }
	if null_sum_eth.Valid { stats.TotalEthDonatedAmountEth = null_sum_eth.Float64 }


	query = "SELECT "+
				"COUNT(*) AS total "+
				"FROM cg_staker_cst " +
				"WHERE num_stake_actions > 0"
	row = sw.S.Db().QueryRow(query)
	var null_stakers sql.NullInt64
	err=row.Scan(&null_stakers)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	if null_stakers.Valid { stats.NumUniqueStakersCST = uint64(null_stakers.Int64) }
	query = "SELECT "+
				"COUNT(*) AS total "+
				"FROM cg_staker_rwalk " +
				"WHERE num_stake_actions > 0"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(&null_stakers)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	if null_stakers.Valid { stats.NumUniqueStakersRWalk = uint64(null_stakers.Int64) }
	query = "SELECT "+
				"COUNT(*) all_tokens_num "+
			"FROM address a "+
				"LEFT JOIN cg_staker_cst c ON a.address_id = c.staker_aid "+
				"LEFT JOIN cg_staker_rwalk r ON a.address_id = r.staker_aid "+
			"WHERE "+
				"(COALESCE(c.total_tokens_staked,0) >0) AND (COALESCE(r.total_tokens_staked,0) > 0) "
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(&null_stakers)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	if null_stakers.Valid { stats.NumUniqueStakersBoth= uint64(null_stakers.Int64) }

	var null_donated_nfts sql.NullInt64
	query = "SELECT "+
				"SUM(num_donated) as total FROM cg_nft_stats"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_donated_nfts,
	)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	stats.NumDonatedNFTs=uint64(null_donated_nfts.Int64)
	query = "SELECT count(*) AS total FROM cg_mint_event "+
			"WHERE LENGTH(token_name) > 0 "
	var null_named_tokens sql.NullInt64
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_named_tokens.Int64,
	)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	stats.TotalNamedTokens = null_named_tokens.Int64

	query = "SELECT count(winner_aid) AS total FROM cg_raffle_Winner_stats "+
			"WHERE amount_sum > 0 "
	var null_num_users_missing_withdrawal sql.NullInt64
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_num_users_missing_withdrawal.Int64,
	)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
			os.Exit(1)
		}
	}
	stats.NumWinnersWithPendingRaffleWithdrawal = null_num_users_missing_withdrawal.Int64

	stats.DonatedTokenDistribution = sw.Get_donated_token_distribution();
	stats.StakeStatisticsCST = sw.Get_stake_statistics_cst()
	stats.StakeStatisticsRWalk = sw.Get_stake_statistics_rwalk()
	return stats
}
func (sw *SQLStorageWrapper) Get_stake_statistics_cst() p.CGStakeStatsCST {

	var stats p.CGStakeStatsCST
	var query string
	query = "SELECT "+
				"total_tokens_staked, "+
				"total_reward_amount,"+
				"total_reward_amount/1e18,"+
				"total_unclaimed_reward,"+
				"total_unclaimed_reward/1e18,"+
				"total_num_stakers, "+
				"num_deposits, "+
				"total_nft_mints "+
			"FROM cg_stake_stats_cst LIMIT 1"

	row := sw.S.Db().QueryRow(query)
	var err error
	err=row.Scan(
		&stats.TotalTokensStaked,
		&stats.TotalReward,
		&stats.TotalRewardEth,
		&stats.UnclaimedReward,
		&stats.UnclaimedRewardEth,
		&stats.NumActiveStakers,
		&stats.NumDeposits,
		&stats.TotalTokensMinted,
	)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_stake_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return stats
}
func (sw *SQLStorageWrapper) Get_stake_statistics_rwalk() p.CGStakeStatsRWalk {

	var stats p.CGStakeStatsRWalk
	var query string
	query = "SELECT "+
				"total_tokens_staked, "+
				"total_num_stakers, "+
				"total_nft_mints "+
			"FROM cg_stake_stats_rwalk LIMIT 1"

	row := sw.S.Db().QueryRow(query)
	var err error
	err=row.Scan(
		&stats.TotalTokensStaked,
		&stats.NumActiveStakers,
		&stats.TotalTokensMinted,
	)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_stake_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return stats
}
func (sw *SQLStorageWrapper) Get_cosmic_game_round_statistics(round_num int64) p.CGRoundStats {

	var stats p.CGRoundStats
	var query string
	query = "SELECT "+
				"round_num, "+
				"total_bids,"+
				"total_nft_donated," +
				"total_raffle_eth_deposits,"+
				"total_raffle_eth_deposits/1e18,"+
				"total_raffle_nfts, "+
				"donations_round_count,"+
				"donations_round_total,"+
				"donations_round_total/1e18 "+
			"FROM cg_round_stats WHERE round_num=$1"

	row := sw.S.Db().QueryRow(query,round_num)
	var err error
	err=row.Scan(
		&stats.RoundNum,
		&stats.TotalBids,
		&stats.TotalDonatedNFTs,
		&stats.TotalRaffleEthDeposits,
		&stats.TotalRaffleEthDepositsEth,
		&stats.TotalRaffleNFTs,
		&stats.TotalDonatedCount,
		&stats.TotalDonatedAmount,
		&stats.TotalDonatedAmountEth,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return stats
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_round_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return stats
}
func (sw *SQLStorageWrapper) Get_unique_bidders() []p.CGUniqueBidder {

	var query string
	query = "SELECT "+
				"b.bidder_aid,"+
				"a.addr,"+
				"b.num_bids,"+
				"b.max_bid,"+
				"b.max_bid/1e18 max_bid_eth "+
			"FROM "+sw.S.SchemaName()+".cg_bidder b "+
				"LEFT JOIN address a ON b.bidder_aid=a.address_id " +
			"ORDER BY num_bids DESC "
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGUniqueBidder,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGUniqueBidder
		err=rows.Scan(
			&rec.BidderAid,
			&rec.BidderAddr,
			&rec.NumBids,
			&rec.MaxBidAmount,
			&rec.MaxBidAmountEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_unique_winners() []p.CGUniqueWinner {

	var query string
	query = "SELECT "+
				"winner_aid,"+
				"winner_addr,"+
				"COUNT(*) AS prizes_count,"+
				"0 AS max_win_amount,"+
				"0 AS max_win_eth,"+
				"0 AS prizes_sum_eth "+
			"FROM ("+
				"SELECT DISTINCT "+
					"COALESCE(pc.winner_aid, rew.winner_aid, rnw.winner_aid, ew.winner_aid, lw.winner_aid, cw.winner_aid) AS winner_aid,"+
					"COALESCE(wa_pc.addr, wa_rew.addr, wa_rnw.addr, wa_ew.addr, wa_lw.addr, wa_cw.addr) AS winner_addr,"+
					"p.round_num,"+
					"p.ptype "+
				"FROM "+sw.S.SchemaName()+".cg_prize p "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_prize_claim pc ON (p.round_num = pc.round_num AND p.ptype IN (0,1,2)) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_pc ON pc.winner_aid = wa_pc.address_id "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_eth_winner rew ON (p.round_num = rew.round_num AND p.winner_index = rew.winner_idx AND p.ptype = 10) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_rew ON rew.winner_aid = wa_rew.address_id "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_nft_winner rnw ON (p.round_num = rnw.round_num AND p.winner_index = rnw.winner_idx AND p.ptype IN (11,12,13,14)) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_rnw ON rnw.winner_aid = wa_rnw.address_id "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_endurance_winner ew ON (p.round_num = ew.round_num AND p.winner_index = ew.winner_idx AND p.ptype IN (5,6)) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_ew ON ew.winner_aid = wa_ew.address_id "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_lastcst_winner lw ON (p.round_num = lw.round_num AND p.winner_index = lw.winner_idx AND p.ptype IN (3,4)) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_lw ON lw.winner_aid = wa_lw.address_id "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_chrono_warrior cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index AND p.ptype IN (7,8,9)) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_cw ON cw.winner_aid = wa_cw.address_id "+
				"WHERE p.ptype != 15"+
			") unique_prizes "+
			"GROUP BY winner_aid, winner_addr "+
			"ORDER BY prizes_count DESC"
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGUniqueWinner,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGUniqueWinner
		err=rows.Scan(
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.PrizesCount,
			&rec.MaxWinAmount,
			&rec.MaxWinAmountEth,
			&rec.PrizesSum,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_unique_stakers_cst() []p.CGUniqueStakerCST {

	var query string
	query = "SELECT "+
				"s.staker_aid,"+
				"a.addr,"+
				"s.total_tokens_staked,"+
				"s.num_stake_actions,"+
				"s.num_unstake_actions,"+
				"s.total_reward,"+
				"s.total_reward/1e18, "+
				"s.unclaimed_reward,"+
				"s.unclaimed_reward/1e18, "+
				"s.num_tokens_minted "+
			"FROM "+sw.S.SchemaName()+".cg_staker_cst s "+
				"LEFT JOIN address a ON s.staker_aid=a.address_id " +
			"WHERE num_stake_actions> 0 "+
			"ORDER BY total_reward DESC "
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGUniqueStakerCST ,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGUniqueStakerCST
		err=rows.Scan(
			&rec.StakerAid,
			&rec.StakerAddr,
			&rec.TotalTokensStaked,
			&rec.NumStakeActions,
			&rec.NumUnstakeActions,
			&rec.TotalReward,
			&rec.TotalRewardEth,
			&rec.UnclaimedReward,
			&rec.UnclaimedRewardEth,
			&rec.TotalTokensMinted,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_unique_stakers_rwalk() []p.CGUniqueStakerRWalk {

	var query string
	query = "SELECT "+
				"s.staker_aid,"+
				"a.addr,"+
				"s.total_tokens_staked,"+
				"s.num_stake_actions,"+
				"s.num_unstake_actions,"+
				"s.num_tokens_minted "+
			"FROM "+sw.S.SchemaName()+".cg_staker_rwalk s "+
				"LEFT JOIN address a ON s.staker_aid=a.address_id " +
			"WHERE num_stake_actions > 0 "+
			"ORDER BY total_tokens_staked DESC "
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGUniqueStakerRWalk ,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGUniqueStakerRWalk
		err=rows.Scan(
			&rec.StakerAid,
			&rec.StakerAddr,
			&rec.TotalTokensStaked,
			&rec.NumStakeActions,
			&rec.NumUnstakeActions,
			&rec.TotalTokensMinted,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_unique_stakers_both() []p.CGUniqueStakersBoth {

	var query string
	query = "SELECT "+
				"a.address_id,"+
				"a.addr,"+

				"COALESCE(c.total_tokens_staked,0) cst_total_tokens_staked,"+
				"COALESCE(c.num_stake_actions,0) cst_num_stake_actions,"+
				"COALESCE(c.num_unstake_actions,0) cst_num_unstake_actions,"+
				"COALESCE(c.total_reward,0) cst_total_reward,"+
				"COALESCE(c.total_reward/1e18,0) cst_total_reward_eth,"+
				"COALESCE(c.unclaimed_reward,0) cst_unclaimed_reward,"+
				"COALESCE(c.unclaimed_reward/1e18,0) cst_unclaimed_reward_eth, "+
				"COALESCE(c.num_tokens_minted,0) cst_num_tokens_minted, "+

				"COALESCE(r.total_tokens_staked,0) rw_total_tokens_staked,"+
				"COALESCE(r.num_stake_actions,0) rw_num_stake_actions,"+
				"COALESCE(r.num_unstake_actions,0) rw_num_unstake_actions,"+
				"COALESCE(r.num_tokens_minted,0) rw_num_tokens_minted,"+

				"(COALESCE(c.total_tokens_staked,0) + COALESCE(r.total_tokens_staked,0)) all_tokens_num "+

			"FROM address a "+
				"LEFT JOIN cg_staker_cst c ON a.address_id = c.staker_aid "+
				"LEFT JOIN cg_staker_rwalk r ON a.address_id = r.staker_aid "+
			"WHERE "+
				"(COALESCE(c.total_tokens_staked,0)>0) AND (COALESCE(r.total_tokens_staked,0) > 0) "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGUniqueStakersBoth ,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGUniqueStakersBoth
		err=rows.Scan(
			&rec.StakerAid,
			&rec.StakerAddr,
			&rec.CSTStats.TotalTokensStaked,
			&rec.CSTStats.NumStakeActions,
			&rec.CSTStats.NumUnstakeActions,
			&rec.CSTStats.TotalReward,
			&rec.CSTStats.TotalRewardEth,
			&rec.CSTStats.UnclaimedReward,
			&rec.CSTStats.UnclaimedRewardEth,
			&rec.CSTStats.TotalTokensMinted,
			&rec.RWalkStats.TotalTokensStaked,
			&rec.RWalkStats.NumStakeActions,
			&rec.RWalkStats.NumUnstakeActions,
			&rec.RWalkStats.TotalTokensMinted,
			&rec.TotalStakedTokensBoth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_unique_donors() []p.CGUniqueDonor {

	var query string
	query = "SELECT "+
				"d.donor_aid,"+
				"a.addr,"+
				"d.count_donations,"+
				"d.total_eth_donated,"+
				"d.total_eth_donated/1e18 total_eth_donated_eth "+
			"FROM "+sw.S.SchemaName()+".cg_donor d "+
				"LEFT JOIN address a ON d.donor_aid=a.address_id " +
			"WHERE d.count_donations > 0 " +
			"ORDER BY total_eth_donated DESC "
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGUniqueDonor,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGUniqueDonor
		err=rows.Scan(
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.CountDonations,
			&rec.TotalDonated,
			&rec.TotalDonatedEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_NFT_donation_stats() []p.CGNFTDonationStats {

	var query string
	query = "SELECT "+
				"s.contract_aid,"+
				"a.addr,"+
				"s.num_donated "+
			"FROM "+sw.S.SchemaName()+".cg_nft_stats s " +
				"LEFT JOIN address a ON s.contract_aid=a.address_id "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGNFTDonationStats,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGNFTDonationStats
		err=rows.Scan(
			&rec.TokenAddressId,
			&rec.TokenAddress,
			&rec.NumDonations,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_record_counters() p.CGRecordCounters {

	var output p.CGRecordCounters
	var null_total_bids,null_total_prizes,null_total_tok_donations sql.NullInt64
	var query string
	query = "SELECT count(*) AS total FROM "+sw.S.SchemaName()+".cg_bid"
	res := sw.S.Db().QueryRow(query)
	err := res.Scan(&null_total_bids)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	} else { output.TotalBids = null_total_bids.Int64 }

	query = "SELECT count(*) AS total FROM "+sw.S.SchemaName()+".cg_prize_claim"
	res = sw.S.Db().QueryRow(query)
	err = res.Scan(&null_total_prizes)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	} else { output.TotalPrizes= null_total_prizes.Int64 }

	query = "SELECT count(*) AS total FROM "+sw.S.SchemaName()+".cg_nft_donation"
	res = sw.S.Db().QueryRow(query)
	err = res.Scan(&null_total_tok_donations)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	} else { output.TotalDonatedNFTs = null_total_tok_donations.Int64 }

	return output
}
func (sw *SQLStorageWrapper) Get_num_prize_claims() int64 {

	var query string
	query = "SELECT num_wins FROM cg_glob_stats"
	row := sw.S.Db().QueryRow(query)
	var err error
	var null_num_claims sql.NullInt64
	err=row.Scan(&null_num_claims);
	if err != nil {
		if err == sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	return null_num_claims.Int64
}
