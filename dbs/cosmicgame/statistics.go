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
		&stats.TotalNFTDonated,
		&stats.NumBidsCST,
		&stats.TotalCSTConsumed,
		&stats.TotalCSTConsumedEth,
		&stats.TotalMktRewards,
		&stats.TotalMktRewardsEth,
		&stats.NumMktRewards,
	)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
		os.Exit(1)
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
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
		os.Exit(1)
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
	if null_winners.Valid { stats.NumUniqueWinners = uint64(null_winners.Int64) }
	if null_sum_wei.Valid { stats.TotalPrizesPaidAmountWei = null_sum_wei.String }
	if null_sum_eth.Valid { stats.TotalPrizesPaidAmountEth = null_sum_eth.Float64 }
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var null_stakers sql.NullInt64
	query = "SELECT "+
				"COUNT(*) AS total "+
				"FROM cg_staker " +
				"WHERE num_stake_actions > 0"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(&null_stakers)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_stakers.Valid { stats.NumUniqueStakers = uint64(null_stakers.Int64) }

	var null_donated_nfts sql.NullInt64
	query = "SELECT "+
				"SUM(num_donated) as total FROM cg_nft_stats"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_donated_nfts,
	)
	stats.NumDonatedNFTs=uint64(null_donated_nfts.Int64)
	query = "SELECT count(*) AS total FROM cg_mint_event "+
			"WHERE LENGTH(token_name) > 0 "
	var null_named_tokens sql.NullInt64
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_named_tokens.Int64,
	)
	stats.TotalNamedTokens = null_named_tokens.Int64

	query = "SELECT count(winner_aid) AS total FROM cg_raffle_Winner_stats "+
			"WHERE amount_sum > 0 "
	var null_num_users_missing_withdrawal sql.NullInt64
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_num_users_missing_withdrawal.Int64,
	)
	stats.NumWinnersWithPendingRaffleWithdrawal = null_num_users_missing_withdrawal.Int64

	stats.DonatedTokenDistribution = sw.Get_donated_token_distribution();
	stats.StakeStatisticsCST = sw.Get_stake_statistics()
	return stats
}
func (sw *SQLStorageWrapper) Get_stake_statistics() p.CGStakeStatsCST {

	var stats p.CGStakeStatsCST
	var query string
	query = "SELECT "+
				"total_tokens_staked, "+
				"total_reward_amount,"+
				"total_reward_amount/1e18,"+
				"total_unclaimed_reward,"+
				"total_unclaimed_reward/1e18,"+
				"total_num_stakers, "+
				"num_deposits "+
			"FROM cg_stake_stats LIMIT 1"

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
				"total_raffle_nfts "+
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
				"w.winner_aid,"+
				"a.addr,"+
				"w.prizes_count,"+
				"w.max_win_amount,"+
				"w.max_win_amount/1e18 max_win_eth, "+
				"w.prizes_sum/1e18 prizes_sum_eth "+
			"FROM "+sw.S.SchemaName()+".cg_winner w "+
				"LEFT JOIN address a ON w.winner_aid=a.address_id " +
			"WHERE w.prizes_count > 0 " +
			"ORDER BY prizes_count DESC "
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
func (sw *SQLStorageWrapper) Get_unique_stakers() []p.CGUniqueStaker {

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
				"s.unclaimed_reward/1e18 "+
			"FROM "+sw.S.SchemaName()+".cg_staker s "+
				"LEFT JOIN address a ON s.staker_aid=a.address_id " +
			"WHERE num_stake_actions > 0 "+
			"ORDER BY total_reward DESC "
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGUniqueStaker ,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGUniqueStaker
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
func (sw *SQLStorageWrapper) Get_user_global_winnings(winner_aid int64) p.CGClaimInfo {

	var output p.CGClaimInfo
	var query string
	query = "SELECT " +
				"s.amount_sum,"+ 
				"s.amount_sum/1e18, " +
				"w.unclaimed_nfts  " +
			"FROM cg_raffle_winner_stats s " +
				"LEFT JOIN cg_winner w ON s.winner_aid=w.winner_aid "+
			"WHERE s.winner_aid = $1"


	row := sw.S.Db().QueryRow(query,winner_aid)
	var err error
	var null_wei sql.NullString
	var null_eth sql.NullFloat64
	var null_nfts sql.NullInt64

	err=row.Scan(&null_wei,&null_eth,&null_nfts);
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	if null_eth.Valid {
		output.ETHRaffleToClaim = null_eth.Float64
	}
	if null_wei.Valid {
		output.ETHRaffleToClaimWei = null_wei.String
	}
	if null_nfts.Valid {
		output.NumDonatedNFTToClaim = null_nfts.Int64
	}

	var null_staking_rewards sql.NullFloat64
	query = "SELECT unclaimed_reward/1e18 FROM cg_staker WHERE staker_aid=$1"
	row = sw.S.Db().QueryRow(query,winner_aid)
	fmt.Printf("query=  %v \n",query);
	fmt.Printf("user = %v\n",winner_aid);
	err=row.Scan(&null_staking_rewards);
	if err != nil {
		if err == sql.ErrNoRows {
			return output;
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	if null_staking_rewards.Valid {
		output.UnclaimedStakingReward = null_staking_rewards.Float64
	}
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
