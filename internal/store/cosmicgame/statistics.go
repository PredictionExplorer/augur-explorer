package cosmicgame

import (
	"os"
	"fmt"
	"time"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
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
			"FROM "+sw.S.SchemaName()+".cg_glob_stats LIMIT 1"

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
			"FROM "+sw.S.SchemaName()+".cg_bidder " +
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
	var null_total_prize_awards sql.NullInt64
	query = "SELECT "+
				"COUNT(*) AS total,"+
				"SUM(prizes_sum) AS sum_wei,"+
				"SUM(prizes_sum)/1e18 AS sum_eth,"+
				"COALESCE(SUM(prizes_count),0) AS total_prize_awards "+
				"FROM "+sw.S.SchemaName()+".cg_winner " +
				"WHERE prizes_count > 0"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_winners,
		&null_sum_wei,
		&null_sum_eth,
		&null_total_prize_awards,
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
	if null_total_prize_awards.Valid { stats.TotalPrizeAwards = uint64(null_total_prize_awards.Int64) }

	var null_donors sql.NullInt64
	query = "SELECT "+
				"COUNT(*) AS total,"+
				"SUM(total_eth_donated) AS sum_wei,"+
				"SUM(total_eth_donated)/1e18 AS sum_eth "+
				"FROM "+sw.S.SchemaName()+".cg_donor " +
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
				"FROM "+sw.S.SchemaName()+".cg_staker_cst " +
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
				"FROM "+sw.S.SchemaName()+".cg_staker_rwalk " +
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
	query = "SELECT count(*) AS total FROM "+sw.S.SchemaName()+".cg_mint_event "+
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

	query = "SELECT count(winner_aid) AS total FROM "+sw.S.SchemaName()+".cg_raffle_winner_stats "+
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

	var null_cg_prize_rows sql.NullInt64
	query = "SELECT COUNT(*) AS total FROM " + sw.S.SchemaName() + ".cg_prize"
	row = sw.S.Db().QueryRow(query)
	err = row.Scan(&null_cg_prize_rows)
	if err != nil {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v", err, query))
			os.Exit(1)
		}
	}
	if null_cg_prize_rows.Valid {
		stats.CgPrizeRowCount = uint64(null_cg_prize_rows.Int64)
	}

	stats.DonatedTokenDistribution = sw.donatedTokenDistributionLegacy()
	stats.StakeStatisticsCST = sw.Get_stake_statistics_cst()
	stats.StakeStatisticsRWalk = sw.Get_stake_statistics_rwalk()
	return stats
}

// donatedTokenDistributionLegacy is the pre-conversion body of what is now
// Repo.DonatedTokenDistribution, kept private for Get_cosmic_game_statistics
// until this file converts to the Repo in its own right.
func (sw *SQLStorageWrapper) donatedTokenDistributionLegacy() []p.CGDonatedTokenDistrRec {
	query := "SELECT " +
		"ca.addr," +
		"ns.num_donated " +
		"FROM " + sw.S.SchemaName() + ".cg_nft_stats ns " +
		"LEFT JOIN address ca ON ns.contract_aid=ca.address_id " +
		"ORDER BY ns.num_donated DESC "

	rows, err := sw.S.Db().Query(query)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	records := make([]p.CGDonatedTokenDistrRec, 0, 16)
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var rec p.CGDonatedTokenDistrRec
		err = rows.Scan(
			&rec.ContractAddr,
			&rec.NumDonatedTokens,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		records = append(records, rec)
	}
	if err := rows.Err(); err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	return records
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
				"num_deposits "+
			"FROM "+sw.S.SchemaName()+".cg_stake_stats_cst LIMIT 1"

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
func (sw *SQLStorageWrapper) Get_stake_statistics_rwalk() p.CGStakeStatsRWalk {

	var stats p.CGStakeStatsRWalk
	var query string
	query = "SELECT "+
				"total_tokens_staked, "+
				"total_num_stakers, "+
				"total_nft_mints "+
			"FROM "+sw.S.SchemaName()+".cg_stake_stats_rwalk LIMIT 1"

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
				"donations_round_total/1e18,"+
				"param_window_start_time::text,"+
				"EXTRACT(EPOCH FROM activation_time)::BIGINT,"+
				"param_window_duration_seconds,"+
				"round_start_time::text,"+
				"round_end_time::text,"+
				"round_duration_seconds, "+
				"total_cst_in_bids/1e18, "+
				"total_eth_in_bids/1e18 "+
			"FROM "+sw.S.SchemaName()+".cg_round_stats WHERE round_num=$1"

	row := sw.S.Db().QueryRow(query,round_num)
	var nullParamWindowStart, nullRoundStart, nullRoundEnd sql.NullString
	var nullActivationTime sql.NullInt64
	var nullParamWindowDuration, nullRoundDuration sql.NullInt64
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
		&nullParamWindowStart,
		&nullActivationTime,
		&nullParamWindowDuration,
		&nullRoundStart,
		&nullRoundEnd,
		&nullRoundDuration,
		&stats.TotalCstInBidsEth,
		&stats.TotalEthInBidsEth,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			stats.RoundNum = round_num
			stats.ActivationTime = sw.get_activation_time_from_events(round_num)
			return stats
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_round_statistics(): %v, q=%v", err, query))
		os.Exit(1)
	}
	if nullParamWindowStart.Valid { stats.ParamWindowStartTime = nullParamWindowStart.String }
	if nullActivationTime.Valid {
		stats.ActivationTime = nullActivationTime.Int64
	} else {
		stats.ActivationTime = sw.get_activation_time_from_events(round_num)
	}
	if nullParamWindowDuration.Valid { stats.ParamWindowDurationSeconds = nullParamWindowDuration.Int64 }
	if nullRoundStart.Valid { stats.RoundStartTime = nullRoundStart.String }
	if nullRoundEnd.Valid { stats.RoundEndTime = nullRoundEnd.String }
	if nullRoundDuration.Valid { stats.RoundDurationSeconds = nullRoundDuration.Int64 }
	return stats
}

// get_activation_time_from_events returns activation_time (Unix seconds) for the given round from cg_adm_acttime when
// that round is the one the latest event applies to (same logic as trigger: 0 when no claims, else last_claimed+1).
func (sw *SQLStorageWrapper) get_activation_time_from_events(round_num int64) int64 {
	q := "SELECT r.new_atime FROM " + sw.S.SchemaName() + ".cg_adm_acttime r " +
		"WHERE (SELECT COALESCE(MAX(p.round_num), -1) + 1 FROM " + sw.S.SchemaName() + ".cg_prize_claim p) = $1 " +
		"ORDER BY r.id DESC LIMIT 1"
	var t int64
	err := sw.S.Db().QueryRow(q, round_num).Scan(&t)
	if err != nil {
		return 0
	}
	return t
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
	query = "WITH prize_winners AS ("+
				"SELECT "+
					"p.round_num,"+
					"p.winner_index,"+
					"p.ptype,"+
					"COALESCE(pc.winner_aid, rew.winner_aid, rnw.winner_aid, ew.winner_aid, lw.winner_aid, cw.winner_aid) AS winner_aid,"+
					"COALESCE(wa_pc.addr, wa_rew.addr, wa_rnw.addr, wa_ew.addr, wa_lw.addr, wa_cw.addr) AS winner_addr "+
				"FROM "+sw.S.SchemaName()+".cg_prize p "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_prize_claim pc ON (p.round_num = pc.round_num AND p.ptype IN (0,1,2)) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_pc ON pc.winner_aid = wa_pc.address_id "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_eth_prize rew ON (p.round_num = rew.round_num AND p.winner_index = rew.winner_idx AND p.ptype = 10) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_rew ON rew.winner_aid = wa_rew.address_id "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_nft_prize rnw ON (p.round_num = rnw.round_num AND p.winner_index = rnw.winner_idx AND p.ptype IN (11,12,13,14) AND ((p.ptype IN (11,12) AND rnw.is_rwalk=false) OR (p.ptype IN (13,14) AND rnw.is_rwalk=true))) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_rnw ON rnw.winner_aid = wa_rnw.address_id "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_endurance_prize ew ON (p.round_num = ew.round_num AND p.ptype IN (5,6)) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_ew ON ew.winner_aid = wa_ew.address_id "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_lastcst_prize lw ON (p.round_num = lw.round_num AND p.ptype IN (3,4)) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_lw ON lw.winner_aid = wa_lw.address_id "+
					"LEFT JOIN "+sw.S.SchemaName()+".cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index AND p.ptype IN (7,8,9)) "+
					"LEFT JOIN "+sw.S.SchemaName()+".address wa_cw ON cw.winner_aid = wa_cw.address_id "+
				"WHERE p.ptype != 15"+
			"), "+
			"bidder_spending AS ("+
				"SELECT bidder_aid, SUM(CASE WHEN eth_price > 0 THEN eth_price ELSE 0 END) AS total_spent "+
				"FROM "+sw.S.SchemaName()+".cg_bid "+
				"GROUP BY bidder_aid"+
			") "+
			"SELECT "+
				"pw.winner_aid,"+
				"pw.winner_addr,"+
				"COUNT(*) AS prizes_count,"+
				"COALESCE(w.max_win_amount,0) AS max_win_amount,"+
				"COALESCE(w.max_win_amount,0)/1e18 AS max_win_eth,"+
				"COALESCE(w.prizes_sum,0)/1e18 AS prizes_sum_eth,"+
				"COALESCE(w.max_win_amount,0),"+
				"COALESCE(w.max_win_amount,0)/1e18,"+
				"COALESCE(w.prizes_count,0),"+
				"COALESCE(w.prizes_sum,0),"+
				"COALESCE(w.prizes_sum,0)/1e18,"+
				"COALESCE(w.tokens_count,0),"+
				"COALESCE(w.erc20_count,0),"+
				"COALESCE(w.erc721_count,0),"+
				"COALESCE(w.unclaimed_nfts,0),"+
				"COALESCE(bs.total_spent,0),"+
				"COALESCE(bs.total_spent,0)/1e18 "+
			"FROM prize_winners pw "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_winner w ON pw.winner_aid=w.winner_aid "+
			"LEFT JOIN bidder_spending bs ON pw.winner_aid=bs.bidder_aid "+
			"WHERE pw.winner_aid IS NOT NULL "+
			"GROUP BY pw.winner_aid, pw.winner_addr, w.max_win_amount, w.prizes_count, w.prizes_sum, w.tokens_count, w.erc20_count, w.erc721_count, w.unclaimed_nfts, bs.total_spent "+
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
			&rec.WinnerStats.MaxWinAmount,
			&rec.WinnerStats.MaxWinAmountEth,
			&rec.WinnerStats.PrizesCount,
			&rec.WinnerStats.PrizesSum,
			&rec.WinnerStats.PrizesSumEth,
			&rec.WinnerStats.TokensCount,
			&rec.WinnerStats.ERC20Count,
			&rec.WinnerStats.ERC721Count,
			&rec.WinnerStats.UnclaimedNfts,
			&rec.WinnerStats.TotalSpent,
			&rec.WinnerStats.TotalSpentEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
// roiLeaderboardOrderClause maps a caller-supplied sort key onto one of the
// whitelisted ORDER BY clauses for Get_roi_leaderboard. Every unrecognized key
// falls back to net profit, so request input can never reach the SQL text.
func roiLeaderboardOrderClause(sortBy string) string {
	switch sortBy {
	case "roi":
		return "roi DESC"
	case "winrate":
		return "win_rate DESC, rounds_participated DESC"
	case "spent":
		return "b.total_eth_spent DESC"
	case "nfts":
		return "nft_prizes_count DESC"
	case "bids":
		return "b.num_bids DESC"
	default:
		return "net_pl_eth DESC"
	}
}

// Get_roi_leaderboard returns per-player bidding profitability (ETH-only ROI, Tier 1),
// joining maintained spend (cg_bidder) with winnings (cg_winner) plus on-demand
// rounds-participated / rounds-won (for win rate). sort_by is whitelisted; min_bids
// filters out one-lucky-bid noise; offset/limit paginate.
func (sw *SQLStorageWrapper) Get_roi_leaderboard(min_bids int, sort_by string, offset int, limit int) []p.CGRoiLeaderboardEntry {

	order := roiLeaderboardOrderClause(sort_by)
	schema := sw.S.SchemaName()
	query := "WITH rounds_part AS ("+
				"SELECT bidder_aid, COUNT(DISTINCT round_num) AS rounds_participated "+
				"FROM "+schema+".cg_bid GROUP BY bidder_aid"+
			"), rounds_won AS ("+
				"SELECT aid, COUNT(DISTINCT round_num) AS rounds_won FROM ("+
					"SELECT winner_aid AS aid, round_num FROM "+schema+".cg_prize_claim "+
					"UNION SELECT winner_aid, round_num FROM "+schema+".cg_raffle_eth_prize "+
					"UNION SELECT winner_aid, round_num FROM "+schema+".cg_raffle_nft_prize "+
					"UNION SELECT winner_aid, round_num FROM "+schema+".cg_endurance_prize "+
					"UNION SELECT winner_aid, round_num FROM "+schema+".cg_lastcst_prize "+
					"UNION SELECT winner_aid, round_num FROM "+schema+".cg_chrono_warrior_prize"+
				") u GROUP BY aid"+
			") "+
			"SELECT "+
				"b.bidder_aid,"+
				"a.addr,"+
				"b.num_bids,"+
				"COALESCE(rp.rounds_participated,0) AS rounds_participated,"+
				"COALESCE(rw.rounds_won,0) AS rounds_won,"+
				"CASE WHEN COALESCE(rp.rounds_participated,0) > 0 "+
					"THEN COALESCE(rw.rounds_won,0)::numeric / rp.rounds_participated ELSE 0 END AS win_rate,"+
				"b.total_eth_spent,"+
				"b.total_eth_spent/1e18,"+
				"b.total_cst_spent,"+
				"b.total_cst_spent/1e18,"+
				"COALESCE(w.prizes_sum,0),"+
				"COALESCE(w.prizes_sum,0)/1e18,"+
				"COALESCE(w.prizes_count,0),"+
				"COALESCE(w.erc20_count,0),"+
				"COALESCE(w.erc721_count,0) AS nft_prizes_count,"+
				"(COALESCE(w.prizes_sum,0) - b.total_eth_spent)/1e18 AS net_pl_eth,"+
				"CASE WHEN b.total_eth_spent > 0 "+
					"THEN (COALESCE(w.prizes_sum,0) - b.total_eth_spent)/b.total_eth_spent ELSE 0 END AS roi "+
			"FROM "+schema+".cg_bidder b "+
				"LEFT JOIN address a ON b.bidder_aid=a.address_id "+
				"LEFT JOIN "+schema+".cg_winner w ON b.bidder_aid=w.winner_aid "+
				"LEFT JOIN rounds_part rp ON b.bidder_aid=rp.bidder_aid "+
				"LEFT JOIN rounds_won rw ON b.bidder_aid=rw.aid "+
			"WHERE b.num_bids >= $1 "+
			"ORDER BY "+order+" "+
			"OFFSET $2 LIMIT $3"
	rows,err := sw.S.Db().Query(query,min_bids,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRoiLeaderboardEntry,0, 64)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRoiLeaderboardEntry
		err=rows.Scan(
			&rec.BidderAid,
			&rec.BidderAddr,
			&rec.NumBids,
			&rec.RoundsParticipated,
			&rec.RoundsWon,
			&rec.WinRate,
			&rec.TotalEthSpent,
			&rec.TotalEthSpentEth,
			&rec.TotalCstSpent,
			&rec.TotalCstSpentEth,
			&rec.EthWon,
			&rec.EthWonEth,
			&rec.PrizesCount,
			&rec.CstPrizesCount,
			&rec.NftPrizesCount,
			&rec.NetPlEth,
			&rec.Roi,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
// Get_claims_by_round returns, per finalized cycle that awarded claimable assets
// (secondary ETH prizes, donated NFTs, donated ERC-20s held in PrizesWallet), how
// many were awarded vs still unclaimed, the claim-window expiry, the average time
// recipients took to claim, and the list of still-unclaimed items for a drill-down.
// Directly-paid assets (main-prize ETH, minted CST/NFT) are not claimable and are
// intentionally excluded.
func (sw *SQLStorageWrapper) Get_claims_by_round() []p.CGRoundClaimSummary {

	schema := sw.S.SchemaName()
	summaryQ := "SELECT "+
			"pc.round_num,"+
			"pc.timeout,"+
			"EXTRACT(EPOCH FROM pc.time_stamp)::bigint,"+
			"COALESCE(eth.awarded,0), COALESCE(eth.unclaimed,0), COALESCE(eth.unclaimed_amt,0)/1e18,"+
			"COALESCE(nft.awarded,0), COALESCE(nft.unclaimed,0),"+
			"COALESCE(erc.awarded,0), COALESCE(erc.unclaimed,0),"+
			"COALESCE(cp.avg_secs,0)::bigint "+
		"FROM "+schema+".cg_prize_claim pc "+
		"LEFT JOIN (SELECT round_num, COUNT(*) awarded, COUNT(*) FILTER (WHERE NOT claimed) unclaimed, "+
			"SUM(amount) FILTER (WHERE NOT claimed) unclaimed_amt FROM "+schema+".cg_prize_deposit GROUP BY round_num) eth "+
			"ON eth.round_num=pc.round_num "+
		"LEFT JOIN (SELECT d.round_num, COUNT(*) awarded, COUNT(*) FILTER (WHERE c.round_num IS NULL) unclaimed "+
			"FROM "+schema+".cg_nft_donation d "+
			"LEFT JOIN "+schema+".cg_donated_nft_claimed c ON c.round_num=d.round_num AND c.idx=d.idx "+
			"GROUP BY d.round_num) nft ON nft.round_num=pc.round_num "+
		"LEFT JOIN (SELECT round_num, COUNT(*) awarded, COUNT(*) FILTER (WHERE NOT claimed) unclaimed "+
			"FROM "+schema+".cg_erc20_donation_stats GROUP BY round_num) erc ON erc.round_num=pc.round_num "+
		"LEFT JOIN (SELECT rn round_num, AVG(secs) avg_secs FROM ("+
				"SELECT w.round_num rn, EXTRACT(EPOCH FROM (w.time_stamp - pcw.time_stamp)) secs "+
					"FROM "+schema+".cg_prize_withdrawal w JOIN "+schema+".cg_prize_claim pcw ON pcw.round_num=w.round_num "+
				"UNION ALL SELECT c.round_num, EXTRACT(EPOCH FROM (c.time_stamp - pcn.time_stamp)) "+
					"FROM "+schema+".cg_donated_nft_claimed c JOIN "+schema+".cg_prize_claim pcn ON pcn.round_num=c.round_num "+
				"UNION ALL SELECT t.round_num, EXTRACT(EPOCH FROM (t.time_stamp - pct.time_stamp)) "+
					"FROM "+schema+".cg_donated_tok_claimed t JOIN "+schema+".cg_prize_claim pct ON pct.round_num=t.round_num "+
			") x GROUP BY rn) cp ON cp.round_num=pc.round_num "+
		"WHERE (COALESCE(eth.awarded,0)+COALESCE(nft.awarded,0)+COALESCE(erc.awarded,0)) > 0 "+
		"ORDER BY pc.round_num DESC"

	rows,err := sw.S.Db().Query(summaryQ)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,summaryQ))
		os.Exit(1)
	}
	now := time.Now().Unix()
	records := make([]p.CGRoundClaimSummary,0, 32)
	byRound := make(map[int64]*p.CGRoundClaimSummary)
	for rows.Next() {
		var r p.CGRoundClaimSummary
		err=rows.Scan(
			&r.RoundNum,
			&r.ClaimWindowTimeout,
			&r.AwardedTs,
			&r.EthAwarded, &r.EthUnclaimed, &r.EthUnclaimedEth,
			&r.NftAwarded, &r.NftUnclaimed,
			&r.Erc20Awarded, &r.Erc20Unclaimed,
			&r.AvgClaimPeriodSecs,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,summaryQ))
			os.Exit(1)
		}
		r.TotalAwarded = r.EthAwarded + r.NftAwarded + r.Erc20Awarded
		r.TotalUnclaimed = r.EthUnclaimed + r.NftUnclaimed + r.Erc20Unclaimed
		r.Expired = now >= r.ClaimWindowTimeout
		r.UnclaimedItems = make([]p.CGClaimUnclaimedItem,0)
		records = append(records,r)
	}
	rows.Close()
	for i := range records {
		byRound[records[i].RoundNum] = &records[i]
	}

	appendItem := func(round int64, item p.CGClaimUnclaimedItem) {
		if s, ok := byRound[round]; ok {
			s.UnclaimedItems = append(s.UnclaimedItems, item)
		}
	}

	// Unclaimed secondary ETH prizes.
	ethQ := "SELECT d.round_num, a.addr, d.amount/1e18 "+
		"FROM "+schema+".cg_prize_deposit d JOIN address a ON a.address_id=d.winner_aid "+
		"WHERE NOT d.claimed"
	sw.scan_unclaimed_items(ethQ, "ETH", appendItem)

	// Unclaimed donated NFTs (claimable by the cycle's main-prize recipient).
	nftQ := "SELECT d.round_num, w.addr, ta.addr, d.token_id "+
		"FROM "+schema+".cg_nft_donation d "+
		"LEFT JOIN "+schema+".cg_donated_nft_claimed c ON c.round_num=d.round_num AND c.idx=d.idx "+
		"JOIN address ta ON ta.address_id=d.token_aid "+
		"LEFT JOIN "+schema+".cg_prize_claim pc ON pc.round_num=d.round_num "+
		"LEFT JOIN address w ON w.address_id=pc.winner_aid "+
		"WHERE c.round_num IS NULL"
	sw.scan_unclaimed_nft_items(nftQ, appendItem)

	// Unclaimed donated ERC-20 tokens (per cycle + token).
	ercQ := "SELECT s.round_num, w.addr, ta.addr, s.total_amount/1e18 "+
		"FROM "+schema+".cg_erc20_donation_stats s "+
		"JOIN address ta ON ta.address_id=s.token_aid "+
		"LEFT JOIN "+schema+".cg_prize_claim pc ON pc.round_num=s.round_num "+
		"LEFT JOIN address w ON w.address_id=pc.winner_aid "+
		"WHERE NOT s.claimed"
	sw.scan_unclaimed_erc20_items(ercQ, appendItem)

	return records
}
func (sw *SQLStorageWrapper) scan_unclaimed_items(query string, assetType string, appendItem func(int64, p.CGClaimUnclaimedItem)) {
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var round int64
		var addr string
		var amount float64
		if err=rows.Scan(&round,&addr,&amount); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		appendItem(round, p.CGClaimUnclaimedItem{AssetType: assetType, RecipientAddr: addr, AmountEth: amount, TokenId: -1})
	}
}
func (sw *SQLStorageWrapper) scan_unclaimed_nft_items(query string, appendItem func(int64, p.CGClaimUnclaimedItem)) {
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var round int64
		var recipient, tokenAddr sql.NullString
		var tokenId int64
		if err=rows.Scan(&round,&recipient,&tokenAddr,&tokenId); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		appendItem(round, p.CGClaimUnclaimedItem{AssetType: "ERC721", RecipientAddr: recipient.String, TokenAddr: tokenAddr.String, TokenId: tokenId})
	}
}
func (sw *SQLStorageWrapper) scan_unclaimed_erc20_items(query string, appendItem func(int64, p.CGClaimUnclaimedItem)) {
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var round int64
		var recipient, tokenAddr sql.NullString
		var amount float64
		if err=rows.Scan(&round,&recipient,&tokenAddr,&amount); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		appendItem(round, p.CGClaimUnclaimedItem{AssetType: "ERC20", RecipientAddr: recipient.String, TokenAddr: tokenAddr.String, AmountEth: amount, TokenId: -1})
	}
}
// Get_claim_detail_by_round returns, for a single cycle, the claim transactions
// (each recipient's withdrawal of a claimable asset, with the time it took after the
// cycle finalized and the tx hash) and the tokens attached during that cycle.
func (sw *SQLStorageWrapper) Get_claim_detail_by_round(round int64) p.CGRoundClaimDetail {

	schema := sw.S.SchemaName()
	detail := p.CGRoundClaimDetail{
		RoundNum:          round,
		ClaimTransactions: make([]p.CGClaimTxn, 0),
		AttachedTokens:    make([]p.CGAttachedToken, 0),
	}

	// ---- Claim transactions: secondary ETH allocations ----
	ethQ := "SELECT ben.addr, win.addr, w.amount/1e18, "+
			"EXTRACT(EPOCH FROM (w.time_stamp - pc.time_stamp))::bigint, "+
			"EXTRACT(EPOCH FROM w.time_stamp)::bigint, t.tx_hash "+
		"FROM "+schema+".cg_prize_withdrawal w "+
		"JOIN "+schema+".cg_prize_claim pc ON pc.round_num=w.round_num "+
		"JOIN address ben ON ben.address_id=w.beneficiary_aid "+
		"JOIN address win ON win.address_id=w.winner_aid "+
		"LEFT JOIN "+schema+".transaction t ON t.id=w.tx_id "+
		"WHERE w.round_num=$1"
	ethRows,err := sw.S.Db().Query(ethQ, round)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,ethQ))
		os.Exit(1)
	}
	for ethRows.Next() {
		var c p.CGClaimTxn
		var txh sql.NullString
		c.AssetType = "ETH"; c.TokenId = -1
		if err=ethRows.Scan(&c.BeneficiaryAddr,&c.RecipientAddr,&c.AmountEth,&c.ClaimedAfterSecs,&c.ClaimTs,&txh); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,ethQ)); os.Exit(1)
		}
		c.TxHash = txh.String
		detail.ClaimTransactions = append(detail.ClaimTransactions, c)
	}
	ethRows.Close()

	// ---- Claim transactions: attached NFTs ----
	nftQ := "SELECT w.addr, ta.addr, dc.token_id, "+
			"EXTRACT(EPOCH FROM (dc.time_stamp - pc.time_stamp))::bigint, "+
			"EXTRACT(EPOCH FROM dc.time_stamp)::bigint, t.tx_hash "+
		"FROM "+schema+".cg_donated_nft_claimed dc "+
		"JOIN "+schema+".cg_prize_claim pc ON pc.round_num=dc.round_num "+
		"JOIN address w ON w.address_id=dc.winner_aid "+
		"JOIN address ta ON ta.address_id=dc.token_aid "+
		"LEFT JOIN "+schema+".transaction t ON t.id=dc.tx_id "+
		"WHERE dc.round_num=$1"
	nftRows,err := sw.S.Db().Query(nftQ, round)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,nftQ)); os.Exit(1)
	}
	for nftRows.Next() {
		var c p.CGClaimTxn
		var txh sql.NullString
		c.AssetType = "ERC721"
		if err=nftRows.Scan(&c.RecipientAddr,&c.TokenAddr,&c.TokenId,&c.ClaimedAfterSecs,&c.ClaimTs,&txh); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,nftQ)); os.Exit(1)
		}
		c.BeneficiaryAddr = c.RecipientAddr; c.TxHash = txh.String
		detail.ClaimTransactions = append(detail.ClaimTransactions, c)
	}
	nftRows.Close()

	// ---- Claim transactions: attached ERC-20s ----
	ercQ := "SELECT w.addr, ta.addr, dc.amount/1e18, "+
			"EXTRACT(EPOCH FROM (dc.time_stamp - pc.time_stamp))::bigint, "+
			"EXTRACT(EPOCH FROM dc.time_stamp)::bigint, t.tx_hash "+
		"FROM "+schema+".cg_donated_tok_claimed dc "+
		"JOIN "+schema+".cg_prize_claim pc ON pc.round_num=dc.round_num "+
		"JOIN address w ON w.address_id=dc.winner_aid "+
		"JOIN address ta ON ta.address_id=dc.token_aid "+
		"LEFT JOIN "+schema+".transaction t ON t.id=dc.tx_id "+
		"WHERE dc.round_num=$1"
	ercRows,err := sw.S.Db().Query(ercQ, round)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,ercQ)); os.Exit(1)
	}
	for ercRows.Next() {
		var c p.CGClaimTxn
		var txh sql.NullString
		c.AssetType = "ERC20"; c.TokenId = -1
		if err=ercRows.Scan(&c.RecipientAddr,&c.TokenAddr,&c.AmountEth,&c.ClaimedAfterSecs,&c.ClaimTs,&txh); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,ercQ)); os.Exit(1)
		}
		c.BeneficiaryAddr = c.RecipientAddr; c.TxHash = txh.String
		detail.ClaimTransactions = append(detail.ClaimTransactions, c)
	}
	ercRows.Close()

	// ---- Attached NFTs this cycle ----
	anQ := "SELECT d.addr, ta.addr, dn.token_id, EXTRACT(EPOCH FROM dn.time_stamp)::bigint, t.tx_hash "+
		"FROM "+schema+".cg_nft_donation dn "+
		"JOIN address d ON d.address_id=dn.donor_aid "+
		"JOIN address ta ON ta.address_id=dn.token_aid "+
		"LEFT JOIN "+schema+".transaction t ON t.id=dn.tx_id "+
		"WHERE dn.round_num=$1 ORDER BY dn.idx"
	anRows,err := sw.S.Db().Query(anQ, round)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,anQ)); os.Exit(1)
	}
	for anRows.Next() {
		var a p.CGAttachedToken
		var txh sql.NullString
		a.AssetType = "ERC721"
		if err=anRows.Scan(&a.ContributorAddr,&a.TokenAddr,&a.TokenId,&a.Ts,&txh); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,anQ)); os.Exit(1)
		}
		a.TxHash = txh.String
		detail.AttachedTokens = append(detail.AttachedTokens, a)
	}
	anRows.Close()

	// ---- Attached ERC-20s this cycle ----
	aeQ := "SELECT d.addr, ta.addr, e.amount/1e18, EXTRACT(EPOCH FROM e.time_stamp)::bigint, t.tx_hash "+
		"FROM "+schema+".cg_erc20_donation e "+
		"JOIN address d ON d.address_id=e.donor_aid "+
		"JOIN address ta ON ta.address_id=e.token_aid "+
		"LEFT JOIN "+schema+".transaction t ON t.id=e.tx_id "+
		"WHERE e.round_num=$1"
	aeRows,err := sw.S.Db().Query(aeQ, round)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,aeQ)); os.Exit(1)
	}
	for aeRows.Next() {
		var a p.CGAttachedToken
		var txh sql.NullString
		a.AssetType = "ERC20"; a.TokenId = -1
		if err=aeRows.Scan(&a.ContributorAddr,&a.TokenAddr,&a.AmountEth,&a.Ts,&txh); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,aeQ)); os.Exit(1)
		}
		a.TxHash = txh.String
		detail.AttachedTokens = append(detail.AttachedTokens, a)
	}
	aeRows.Close()

	return detail
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
				"s.unclaimed_reward/1e18 "+
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
	query = "SELECT num_wins FROM "+sw.S.SchemaName()+".cg_glob_stats"
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
