package cosmicgame

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// CosmicGameStatistics returns the global dashboard aggregate: the
// cg_glob_stats row plus unique-participant counts, prize/donation sums,
// donated-token distribution and both staking summaries. Sub-aggregates
// whose tables are empty keep their zero values.
func (r *Repo) CosmicGameStatistics(ctx context.Context) (cgmodel.CGStatistics, error) {
	const op = "cosmic game statistics"
	var stats cgmodel.CGStatistics

	query := "SELECT " +
		"num_vol_donations, " +
		"vol_donations_total/1e18 as voluntary_donations_sum," +
		"num_cg_donations," +
		"cg_donations_total/1e18," +
		"direct_donations/1e18," +
		"num_direct_donations," +
		"num_withdrawals," +
		"sum_withdrawals/1e18," +
		"num_bids," +
		"cur_num_bids," +
		"num_wins, " +
		"num_rwalk_used, " +
		"num_mints, " +
		"total_raffle_eth_deposits/1e18, " +
		"total_raffle_eth_withdrawn/1e18, " +
		"total_chrono_warrior_eth_deposits/1e18, " +
		"total_cst_given_in_prizes/1e18, " +
		"total_nft_donated," +
		"num_bids_cst," +
		"total_cst_consumed," +
		"total_cst_consumed/1e18, " +
		"total_mkt_rewards," +
		"total_mkt_rewards/1e18," +
		"num_mkt_rewards " +
		"FROM cg_glob_stats LIMIT 1"
	err := r.q(ctx).QueryRow(ctx, query).Scan(
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
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": glob stats", err)
	}

	var nullBidders sql.NullInt64
	err = r.q(ctx).QueryRow(ctx,
		"SELECT COUNT(*) AS total FROM cg_bidder WHERE num_bids > 0",
	).Scan(&nullBidders)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": unique bidders", err)
	}
	if nullBidders.Valid {
		stats.NumUniqueBidders = uint64(nullBidders.Int64) // #nosec G115 -- COUNT(*) is non-negative
	}

	var nullWinners sql.NullInt64
	var nullSumWei sql.NullString
	var nullSumEth sql.NullFloat64
	var nullTotalPrizeAwards sql.NullInt64
	err = r.q(ctx).QueryRow(ctx,
		"SELECT "+
			"COUNT(*) AS total,"+
			"SUM(prizes_sum) AS sum_wei,"+
			"SUM(prizes_sum)/1e18 AS sum_eth,"+
			"COALESCE(SUM(prizes_count),0) AS total_prize_awards "+
			"FROM cg_winner "+
			"WHERE prizes_count > 0",
	).Scan(
		&nullWinners,
		&nullSumWei,
		&nullSumEth,
		&nullTotalPrizeAwards,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": unique winners", err)
	}
	if nullWinners.Valid {
		stats.NumUniqueWinners = uint64(nullWinners.Int64) // #nosec G115 -- COUNT(*) is non-negative
	}
	if nullSumWei.Valid {
		stats.TotalPrizesPaidAmountWei = nullSumWei.String
	}
	if nullSumEth.Valid {
		stats.TotalPrizesPaidAmountEth = nullSumEth.Float64
	}
	if nullTotalPrizeAwards.Valid {
		stats.TotalPrizeAwards = uint64(nullTotalPrizeAwards.Int64) // #nosec G115 -- COALESCE(SUM(count),0) is non-negative
	}

	var nullDonors sql.NullInt64
	err = r.q(ctx).QueryRow(ctx,
		"SELECT "+
			"COUNT(*) AS total,"+
			"SUM(total_eth_donated) AS sum_wei,"+
			"SUM(total_eth_donated)/1e18 AS sum_eth "+
			"FROM cg_donor "+
			"WHERE total_eth_donated > 0",
	).Scan(
		&nullDonors,
		&nullSumWei,
		&nullSumEth,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": unique donors", err)
	}
	if nullDonors.Valid {
		stats.NumUniqueDonors = nullDonors.Int64
	}
	if nullSumWei.Valid {
		stats.TotalEthDonatedAmount = nullSumWei.String
	}
	if nullSumEth.Valid {
		stats.TotalEthDonatedAmountEth = nullSumEth.Float64
	}

	var nullStakers sql.NullInt64
	err = r.q(ctx).QueryRow(ctx,
		"SELECT COUNT(*) AS total FROM cg_staker_cst WHERE num_stake_actions > 0",
	).Scan(&nullStakers)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": unique stakers cst", err)
	}
	if nullStakers.Valid {
		stats.NumUniqueStakersCST = uint64(nullStakers.Int64) // #nosec G115 -- COUNT(*) is non-negative
	}
	err = r.q(ctx).QueryRow(ctx,
		"SELECT COUNT(*) AS total FROM cg_staker_rwalk WHERE num_stake_actions > 0",
	).Scan(&nullStakers)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": unique stakers rwalk", err)
	}
	if nullStakers.Valid {
		stats.NumUniqueStakersRWalk = uint64(nullStakers.Int64) // #nosec G115 -- COUNT(*) is non-negative
	}
	err = r.q(ctx).QueryRow(ctx,
		"SELECT "+
			"COUNT(*) all_tokens_num "+
			"FROM address a "+
			"LEFT JOIN cg_staker_cst c ON a.address_id = c.staker_aid "+
			"LEFT JOIN cg_staker_rwalk r ON a.address_id = r.staker_aid "+
			"WHERE "+
			"(COALESCE(c.total_tokens_staked,0) >0) AND (COALESCE(r.total_tokens_staked,0) > 0) ",
	).Scan(&nullStakers)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": unique stakers both", err)
	}
	if nullStakers.Valid {
		stats.NumUniqueStakersBoth = uint64(nullStakers.Int64) // #nosec G115 -- COUNT(*) is non-negative
	}

	var nullDonatedNfts sql.NullInt64
	err = r.q(ctx).QueryRow(ctx,
		"SELECT SUM(num_donated) as total FROM cg_nft_stats",
	).Scan(&nullDonatedNfts)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": donated nfts", err)
	}
	stats.NumDonatedNFTs = uint64(nullDonatedNfts.Int64) // #nosec G115 -- SUM of non-negative counters

	var namedTokens int64
	err = r.q(ctx).QueryRow(ctx,
		"SELECT count(*) AS total FROM cg_mint_event WHERE LENGTH(token_name) > 0 ",
	).Scan(&namedTokens)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": named tokens", err)
	}
	stats.TotalNamedTokens = namedTokens

	var numUsersMissingWithdrawal int64
	err = r.q(ctx).QueryRow(ctx,
		"SELECT count(winner_aid) AS total FROM cg_raffle_winner_stats WHERE amount_sum > 0 ",
	).Scan(&numUsersMissingWithdrawal)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": pending raffle withdrawals", err)
	}
	stats.NumWinnersWithPendingRaffleWithdrawal = numUsersMissingWithdrawal

	var nullCgPrizeRows sql.NullInt64
	err = r.q(ctx).QueryRow(ctx, "SELECT COUNT(*) AS total FROM cg_prize").Scan(&nullCgPrizeRows)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGStatistics{}, store.WrapError(op+": cg_prize rows", err)
	}
	if nullCgPrizeRows.Valid {
		stats.CgPrizeRowCount = uint64(nullCgPrizeRows.Int64) // #nosec G115 -- COUNT(*) is non-negative
	}

	stats.DonatedTokenDistribution, err = r.DonatedTokenDistribution(ctx)
	if err != nil {
		return cgmodel.CGStatistics{}, err
	}
	stats.StakeStatisticsCST, err = r.StakeStatisticsCst(ctx)
	if err != nil {
		return cgmodel.CGStatistics{}, err
	}
	stats.StakeStatisticsRWalk, err = r.StakeStatisticsRwalk(ctx)
	if err != nil {
		return cgmodel.CGStatistics{}, err
	}
	return stats, nil
}

// StakeStatisticsCst returns the global CST staking summary row.
func (r *Repo) StakeStatisticsCst(ctx context.Context) (cgmodel.CGStakeStatsCST, error) {
	var stats cgmodel.CGStakeStatsCST
	query := "SELECT " +
		"total_tokens_staked, " +
		"total_reward_amount," +
		"total_reward_amount/1e18," +
		"total_unclaimed_reward," +
		"total_unclaimed_reward/1e18," +
		"total_num_stakers, " +
		"num_deposits " +
		"FROM cg_stake_stats_cst LIMIT 1"
	err := r.q(ctx).QueryRow(ctx, query).Scan(
		&stats.TotalTokensStaked,
		&stats.TotalReward,
		&stats.TotalRewardEth,
		&stats.UnclaimedReward,
		&stats.UnclaimedRewardEth,
		&stats.NumActiveStakers,
		&stats.NumDeposits,
	)
	if err != nil {
		return cgmodel.CGStakeStatsCST{}, store.WrapError("stake statistics cst", err)
	}
	return stats, nil
}

// StakeStatisticsRwalk returns the global RandomWalk staking summary row.
func (r *Repo) StakeStatisticsRwalk(ctx context.Context) (cgmodel.CGStakeStatsRWalk, error) {
	var stats cgmodel.CGStakeStatsRWalk
	query := "SELECT " +
		"total_tokens_staked, " +
		"total_num_stakers, " +
		"total_nft_mints " +
		"FROM cg_stake_stats_rwalk LIMIT 1"
	err := r.q(ctx).QueryRow(ctx, query).Scan(
		&stats.TotalTokensStaked,
		&stats.NumActiveStakers,
		&stats.TotalTokensMinted,
	)
	if err != nil {
		return cgmodel.CGStakeStatsRWalk{}, store.WrapError("stake statistics rwalk", err)
	}
	return stats, nil
}

// CosmicGameRoundStatistics returns the per-round aggregate row. A round
// without a stats row yet (the open round) comes back with only RoundNum and
// the activation time derived from admin events.
func (r *Repo) CosmicGameRoundStatistics(ctx context.Context, roundNum int64) (cgmodel.CGRoundStats, error) {
	const op = "cosmic game round statistics"
	var stats cgmodel.CGRoundStats
	query := "SELECT " +
		"round_num, " +
		"total_bids," +
		"total_nft_donated," +
		"total_raffle_eth_deposits," +
		"total_raffle_eth_deposits/1e18," +
		"total_raffle_nfts, " +
		"donations_round_count," +
		"donations_round_total," +
		"donations_round_total/1e18," +
		"param_window_start_time::text," +
		"EXTRACT(EPOCH FROM activation_time)::BIGINT," +
		"param_window_duration_seconds," +
		"round_start_time::text," +
		"round_end_time::text," +
		"round_duration_seconds, " +
		"total_cst_in_bids/1e18, " +
		"total_eth_in_bids/1e18, " +
		"endurance_champion_duration, " +
		"chrono_warrior_duration " +
		"FROM cg_round_stats WHERE round_num=$1"

	var nullParamWindowStart, nullRoundStart, nullRoundEnd sql.NullString
	var nullActivationTime sql.NullInt64
	var nullParamWindowDuration, nullRoundDuration sql.NullInt64
	err := r.q(ctx).QueryRow(ctx, query, roundNum).Scan(
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
		&stats.EnduranceChampionDuration,
		&stats.ChronoWarriorDuration,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			stats.RoundNum = roundNum
			stats.ActivationTime, err = r.activationTimeFromEvents(ctx, roundNum)
			if err != nil {
				return cgmodel.CGRoundStats{}, err
			}
			return stats, nil
		}
		return cgmodel.CGRoundStats{}, store.WrapError(op, err)
	}
	if nullParamWindowStart.Valid {
		stats.ParamWindowStartTime = nullParamWindowStart.String
	}
	if nullActivationTime.Valid {
		stats.ActivationTime = nullActivationTime.Int64
	} else {
		stats.ActivationTime, err = r.activationTimeFromEvents(ctx, roundNum)
		if err != nil {
			return cgmodel.CGRoundStats{}, err
		}
	}
	if nullParamWindowDuration.Valid {
		stats.ParamWindowDurationSeconds = nullParamWindowDuration.Int64
	}
	if nullRoundStart.Valid {
		stats.RoundStartTime = nullRoundStart.String
	}
	if nullRoundEnd.Valid {
		stats.RoundEndTime = nullRoundEnd.String
	}
	if nullRoundDuration.Valid {
		stats.RoundDurationSeconds = nullRoundDuration.Int64
	}
	return stats, nil
}

// activationTimeFromEvents returns the activation time (Unix seconds) for
// roundNum from cg_adm_acttime when that round is the one the latest event
// applies to (same logic as the trigger: 0 when no claims, else
// last_claimed+1); 0 when no event matches.
func (r *Repo) activationTimeFromEvents(ctx context.Context, roundNum int64) (int64, error) {
	query := "SELECT r.new_atime FROM cg_adm_acttime r " +
		"WHERE (SELECT COALESCE(MAX(p.round_num), -1) + 1 FROM cg_prize_claim p) = $1 " +
		"ORDER BY r.id DESC LIMIT 1"
	var t int64
	err := r.q(ctx).QueryRow(ctx, query, roundNum).Scan(&t)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, store.WrapError("activation time from events", err)
	}
	return t, nil
}

// UniqueBidders returns every address that ever bid, with bid counts and
// maximum bid, most active first.
func (r *Repo) UniqueBidders(ctx context.Context) ([]cgmodel.CGUniqueBidder, error) {
	query := "SELECT " +
		"b.bidder_aid," +
		"a.addr," +
		"b.num_bids," +
		"b.max_bid," +
		"b.max_bid/1e18 max_bid_eth " +
		"FROM cg_bidder b " +
		"LEFT JOIN address a ON b.bidder_aid=a.address_id " +
		"ORDER BY num_bids DESC "
	scan := func(rows pgx.Rows, rec *cgmodel.CGUniqueBidder) error {
		return rows.Scan(
			&rec.BidderAid,
			&rec.BidderAddr,
			&rec.NumBids,
			&rec.MaxBidAmount,
			&rec.MaxBidAmountEth,
		)
	}
	return queryList(ctx, r, "unique bidders", 32, query, scan)
}

// UniqueWinners returns every address that won any prize, with per-winner
// aggregate statistics, most prizes first.
func (r *Repo) UniqueWinners(ctx context.Context) ([]cgmodel.CGUniqueWinner, error) {
	query := "WITH prize_winners AS (" +
		"SELECT " +
		"p.round_num," +
		"p.winner_index," +
		"p.ptype," +
		"COALESCE(pc.winner_aid, rew.winner_aid, rnw.winner_aid, ew.winner_aid, lw.winner_aid, cw.winner_aid) AS winner_aid," +
		"COALESCE(wa_pc.addr, wa_rew.addr, wa_rnw.addr, wa_ew.addr, wa_lw.addr, wa_cw.addr) AS winner_addr " +
		"FROM cg_prize p " +
		"LEFT JOIN cg_prize_claim pc ON (p.round_num = pc.round_num AND p.ptype IN (0,1,2)) " +
		"LEFT JOIN address wa_pc ON pc.winner_aid = wa_pc.address_id " +
		"LEFT JOIN cg_raffle_eth_prize rew ON (p.round_num = rew.round_num AND p.winner_index = rew.winner_idx AND p.ptype = 10) " +
		"LEFT JOIN address wa_rew ON rew.winner_aid = wa_rew.address_id " +
		"LEFT JOIN cg_raffle_nft_prize rnw ON (p.round_num = rnw.round_num AND p.winner_index = rnw.winner_idx AND p.ptype IN (11,12,13,14) AND ((p.ptype IN (11,12) AND rnw.is_rwalk=false) OR (p.ptype IN (13,14) AND rnw.is_rwalk=true))) " +
		"LEFT JOIN address wa_rnw ON rnw.winner_aid = wa_rnw.address_id " +
		"LEFT JOIN cg_endurance_prize ew ON (p.round_num = ew.round_num AND p.ptype IN (5,6)) " +
		"LEFT JOIN address wa_ew ON ew.winner_aid = wa_ew.address_id " +
		"LEFT JOIN cg_lastcst_prize lw ON (p.round_num = lw.round_num AND p.ptype IN (3,4)) " +
		"LEFT JOIN address wa_lw ON lw.winner_aid = wa_lw.address_id " +
		"LEFT JOIN cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index AND p.ptype IN (7,8,9)) " +
		"LEFT JOIN address wa_cw ON cw.winner_aid = wa_cw.address_id " +
		"WHERE p.ptype != 15" +
		"), " +
		"bidder_spending AS (" +
		"SELECT bidder_aid, SUM(CASE WHEN eth_price > 0 THEN eth_price ELSE 0 END) AS total_spent " +
		"FROM cg_bid " +
		"GROUP BY bidder_aid" +
		") " +
		"SELECT " +
		"pw.winner_aid," +
		"pw.winner_addr," +
		"COUNT(*) AS prizes_count," +
		"COALESCE(w.max_win_amount,0) AS max_win_amount," +
		"COALESCE(w.max_win_amount,0)/1e18 AS max_win_eth," +
		"COALESCE(w.prizes_sum,0)/1e18 AS prizes_sum_eth," +
		"COALESCE(w.max_win_amount,0)," +
		"COALESCE(w.max_win_amount,0)/1e18," +
		"COALESCE(w.prizes_count,0)," +
		"COALESCE(w.prizes_sum,0)," +
		"COALESCE(w.prizes_sum,0)/1e18," +
		"COALESCE(w.tokens_count,0)," +
		"COALESCE(w.erc20_count,0)," +
		"COALESCE(w.erc721_count,0)," +
		"COALESCE(w.unclaimed_nfts,0)," +
		"COALESCE(bs.total_spent,0)," +
		"COALESCE(bs.total_spent,0)/1e18 " +
		"FROM prize_winners pw " +
		"LEFT JOIN cg_winner w ON pw.winner_aid=w.winner_aid " +
		"LEFT JOIN bidder_spending bs ON pw.winner_aid=bs.bidder_aid " +
		"WHERE pw.winner_aid IS NOT NULL " +
		"GROUP BY pw.winner_aid, pw.winner_addr, w.max_win_amount, w.prizes_count, w.prizes_sum, w.tokens_count, w.erc20_count, w.erc721_count, w.unclaimed_nfts, bs.total_spent " +
		"ORDER BY prizes_count DESC"
	scan := func(rows pgx.Rows, rec *cgmodel.CGUniqueWinner) error {
		return rows.Scan(
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
	}
	return queryList(ctx, r, "unique winners", 32, query, scan)
}

// roiLeaderboardOrderClause maps a caller-supplied sort key onto one of the
// whitelisted ORDER BY clauses for RoiLeaderboard. Every unrecognized key
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

// RoiLeaderboard returns per-player bidding profitability (ETH-only ROI,
// Tier 1), joining maintained spend (cg_bidder) with winnings (cg_winner)
// plus on-demand rounds-participated / rounds-won (for win rate). sortBy is
// whitelisted; minBids filters out one-lucky-bid noise; offset/limit
// paginate.
func (r *Repo) RoiLeaderboard(ctx context.Context, minBids int, sortBy string, offset, limit int) ([]cgmodel.CGRoiLeaderboardEntry, error) {
	order := roiLeaderboardOrderClause(sortBy)
	query := "WITH rounds_part AS (" +
		"SELECT bidder_aid, COUNT(DISTINCT round_num) AS rounds_participated " +
		"FROM cg_bid GROUP BY bidder_aid" +
		"), rounds_won AS (" +
		"SELECT aid, COUNT(DISTINCT round_num) AS rounds_won FROM (" +
		"SELECT winner_aid AS aid, round_num FROM cg_prize_claim " +
		"UNION SELECT winner_aid, round_num FROM cg_raffle_eth_prize " +
		"UNION SELECT winner_aid, round_num FROM cg_raffle_nft_prize " +
		"UNION SELECT winner_aid, round_num FROM cg_endurance_prize " +
		"UNION SELECT winner_aid, round_num FROM cg_lastcst_prize " +
		"UNION SELECT winner_aid, round_num FROM cg_chrono_warrior_prize" +
		") u GROUP BY aid" +
		") " +
		"SELECT " +
		"b.bidder_aid," +
		"a.addr," +
		"b.num_bids," +
		"COALESCE(rp.rounds_participated,0) AS rounds_participated," +
		"COALESCE(rw.rounds_won,0) AS rounds_won," +
		"CASE WHEN COALESCE(rp.rounds_participated,0) > 0 " +
		"THEN COALESCE(rw.rounds_won,0)::numeric / rp.rounds_participated ELSE 0 END AS win_rate," +
		"b.total_eth_spent," +
		"b.total_eth_spent/1e18," +
		"b.total_cst_spent," +
		"b.total_cst_spent/1e18," +
		"COALESCE(w.prizes_sum,0)," +
		"COALESCE(w.prizes_sum,0)/1e18," +
		"COALESCE(w.prizes_count,0)," +
		"COALESCE(w.erc20_count,0)," +
		"COALESCE(w.erc721_count,0) AS nft_prizes_count," +
		"(COALESCE(w.prizes_sum,0) - b.total_eth_spent)/1e18 AS net_pl_eth," +
		"CASE WHEN b.total_eth_spent > 0 " +
		"THEN (COALESCE(w.prizes_sum,0) - b.total_eth_spent)/b.total_eth_spent ELSE 0 END AS roi " +
		"FROM cg_bidder b " +
		"LEFT JOIN address a ON b.bidder_aid=a.address_id " +
		"LEFT JOIN cg_winner w ON b.bidder_aid=w.winner_aid " +
		"LEFT JOIN rounds_part rp ON b.bidder_aid=rp.bidder_aid " +
		"LEFT JOIN rounds_won rw ON b.bidder_aid=rw.aid " +
		"WHERE b.num_bids >= $1 " +
		"ORDER BY " + order + " " +
		"OFFSET $2 LIMIT $3"
	scan := func(rows pgx.Rows, rec *cgmodel.CGRoiLeaderboardEntry) error {
		return rows.Scan(
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
	}
	return queryList(ctx, r, "roi leaderboard", 64, query, scan, minBids, offset, limit)
}

// ClaimsByRound returns, per finalized cycle that awarded claimable assets
// (secondary ETH prizes, donated NFTs, donated ERC-20s held in
// PrizesWallet), how many were awarded vs still unclaimed, the claim-window
// expiry, the average time recipients took to claim, and the list of
// still-unclaimed items for a drill-down. Directly-paid assets (main-prize
// ETH, minted CST/NFT) are not claimable and are intentionally excluded.
func (r *Repo) ClaimsByRound(ctx context.Context) ([]cgmodel.CGRoundClaimSummary, error) {
	const op = "claims by round"
	summaryQ := "SELECT " +
		"pc.round_num," +
		"pc.timeout," +
		"EXTRACT(EPOCH FROM pc.time_stamp)::bigint," +
		"COALESCE(eth.awarded,0), COALESCE(eth.unclaimed,0), COALESCE(eth.unclaimed_amt,0)/1e18," +
		"COALESCE(nft.awarded,0), COALESCE(nft.unclaimed,0)," +
		"COALESCE(erc.awarded,0), COALESCE(erc.unclaimed,0)," +
		"COALESCE(cp.avg_secs,0)::bigint " +
		"FROM cg_prize_claim pc " +
		"LEFT JOIN (SELECT round_num, COUNT(*) awarded, COUNT(*) FILTER (WHERE NOT claimed) unclaimed, " +
		"SUM(amount) FILTER (WHERE NOT claimed) unclaimed_amt FROM cg_prize_deposit GROUP BY round_num) eth " +
		"ON eth.round_num=pc.round_num " +
		"LEFT JOIN (SELECT d.round_num, COUNT(*) awarded, COUNT(*) FILTER (WHERE c.round_num IS NULL) unclaimed " +
		"FROM cg_nft_donation d " +
		"LEFT JOIN cg_donated_nft_claimed c ON c.round_num=d.round_num AND c.idx=d.idx " +
		"GROUP BY d.round_num) nft ON nft.round_num=pc.round_num " +
		"LEFT JOIN (SELECT round_num, COUNT(*) awarded, COUNT(*) FILTER (WHERE NOT claimed) unclaimed " +
		"FROM cg_erc20_donation_stats GROUP BY round_num) erc ON erc.round_num=pc.round_num " +
		"LEFT JOIN (SELECT rn round_num, AVG(secs) avg_secs FROM (" +
		"SELECT w.round_num rn, EXTRACT(EPOCH FROM (w.time_stamp - pcw.time_stamp)) secs " +
		"FROM cg_prize_withdrawal w JOIN cg_prize_claim pcw ON pcw.round_num=w.round_num " +
		"UNION ALL SELECT c.round_num, EXTRACT(EPOCH FROM (c.time_stamp - pcn.time_stamp)) " +
		"FROM cg_donated_nft_claimed c JOIN cg_prize_claim pcn ON pcn.round_num=c.round_num " +
		"UNION ALL SELECT t.round_num, EXTRACT(EPOCH FROM (t.time_stamp - pct.time_stamp)) " +
		"FROM cg_donated_tok_claimed t JOIN cg_prize_claim pct ON pct.round_num=t.round_num " +
		") x GROUP BY rn) cp ON cp.round_num=pc.round_num " +
		"WHERE (COALESCE(eth.awarded,0)+COALESCE(nft.awarded,0)+COALESCE(erc.awarded,0)) > 0 " +
		"ORDER BY pc.round_num DESC"

	now := time.Now().Unix()
	scan := func(rows pgx.Rows, rec *cgmodel.CGRoundClaimSummary) error {
		err := rows.Scan(
			&rec.RoundNum,
			&rec.ClaimWindowTimeout,
			&rec.AwardedTs,
			&rec.EthAwarded, &rec.EthUnclaimed, &rec.EthUnclaimedEth,
			&rec.NftAwarded, &rec.NftUnclaimed,
			&rec.Erc20Awarded, &rec.Erc20Unclaimed,
			&rec.AvgClaimPeriodSecs,
		)
		if err != nil {
			return err
		}
		rec.TotalAwarded = rec.EthAwarded + rec.NftAwarded + rec.Erc20Awarded
		rec.TotalUnclaimed = rec.EthUnclaimed + rec.NftUnclaimed + rec.Erc20Unclaimed
		rec.Expired = now >= rec.ClaimWindowTimeout
		rec.UnclaimedItems = make([]cgmodel.CGClaimUnclaimedItem, 0)
		return nil
	}
	records, err := queryList(ctx, r, op, 32, summaryQ, scan)
	if err != nil {
		return nil, err
	}
	byRound := make(map[int64]*cgmodel.CGRoundClaimSummary, len(records))
	for i := range records {
		byRound[records[i].RoundNum] = &records[i]
	}
	appendItem := func(round int64, item cgmodel.CGClaimUnclaimedItem) {
		if s, ok := byRound[round]; ok {
			s.UnclaimedItems = append(s.UnclaimedItems, item)
		}
	}

	// Unclaimed secondary ETH prizes.
	ethQ := "SELECT d.round_num, a.addr, d.amount/1e18 " +
		"FROM cg_prize_deposit d JOIN address a ON a.address_id=d.winner_aid " +
		"WHERE NOT d.claimed " +
		"ORDER BY d.round_num DESC, d.winner_index, d.id"
	if err := r.scanUnclaimedItems(ctx, ethQ, "ETH", appendItem); err != nil {
		return nil, err
	}

	// Unclaimed donated NFTs (claimable by the cycle's main-prize recipient).
	nftQ := "SELECT d.round_num, w.addr, ta.addr, d.token_id " +
		"FROM cg_nft_donation d " +
		"LEFT JOIN cg_donated_nft_claimed c ON c.round_num=d.round_num AND c.idx=d.idx " +
		"JOIN address ta ON ta.address_id=d.token_aid " +
		"LEFT JOIN cg_prize_claim pc ON pc.round_num=d.round_num " +
		"LEFT JOIN address w ON w.address_id=pc.winner_aid " +
		"WHERE c.round_num IS NULL " +
		"ORDER BY d.round_num DESC, d.idx, d.id"
	if err := r.scanUnclaimedNFTItems(ctx, nftQ, appendItem); err != nil {
		return nil, err
	}

	// Unclaimed donated ERC-20 tokens (per cycle + token).
	ercQ := "SELECT s.round_num, w.addr, ta.addr, s.total_amount/1e18 " +
		"FROM cg_erc20_donation_stats s " +
		"JOIN address ta ON ta.address_id=s.token_aid " +
		"LEFT JOIN cg_prize_claim pc ON pc.round_num=s.round_num " +
		"LEFT JOIN address w ON w.address_id=pc.winner_aid " +
		"WHERE NOT s.claimed " +
		"ORDER BY s.round_num DESC, s.token_aid"
	if err := r.scanUnclaimedERC20Items(ctx, ercQ, appendItem); err != nil {
		return nil, err
	}
	return records, nil
}

func (r *Repo) scanUnclaimedItems(ctx context.Context, query, assetType string, appendItem func(int64, cgmodel.CGClaimUnclaimedItem)) error {
	const op = "claims by round: unclaimed items"
	rows, err := r.q(ctx).Query(ctx, query)
	if err != nil {
		return store.WrapError(op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var round int64
		var addr string
		var amount float64
		if err := rows.Scan(&round, &addr, &amount); err != nil {
			return store.WrapError(op, err)
		}
		appendItem(round, cgmodel.CGClaimUnclaimedItem{AssetType: assetType, RecipientAddr: addr, AmountEth: amount, TokenId: -1})
	}
	return store.WrapError(op, rows.Err())
}

func (r *Repo) scanUnclaimedNFTItems(ctx context.Context, query string, appendItem func(int64, cgmodel.CGClaimUnclaimedItem)) error {
	const op = "claims by round: unclaimed nft items"
	rows, err := r.q(ctx).Query(ctx, query)
	if err != nil {
		return store.WrapError(op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var round int64
		var recipient, tokenAddr sql.NullString
		var tokenID int64
		if err := rows.Scan(&round, &recipient, &tokenAddr, &tokenID); err != nil {
			return store.WrapError(op, err)
		}
		appendItem(round, cgmodel.CGClaimUnclaimedItem{AssetType: "ERC721", RecipientAddr: recipient.String, TokenAddr: tokenAddr.String, TokenId: tokenID})
	}
	return store.WrapError(op, rows.Err())
}

func (r *Repo) scanUnclaimedERC20Items(ctx context.Context, query string, appendItem func(int64, cgmodel.CGClaimUnclaimedItem)) error {
	const op = "claims by round: unclaimed erc20 items"
	rows, err := r.q(ctx).Query(ctx, query)
	if err != nil {
		return store.WrapError(op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var round int64
		var recipient, tokenAddr sql.NullString
		var amount float64
		if err := rows.Scan(&round, &recipient, &tokenAddr, &amount); err != nil {
			return store.WrapError(op, err)
		}
		appendItem(round, cgmodel.CGClaimUnclaimedItem{AssetType: "ERC20", RecipientAddr: recipient.String, TokenAddr: tokenAddr.String, AmountEth: amount, TokenId: -1})
	}
	return store.WrapError(op, rows.Err())
}

// ClaimDetailByRound returns, for a single cycle, the claim transactions
// (each recipient's withdrawal of a claimable asset, with the time it took
// after the cycle finalized and the tx hash) and the tokens attached during
// that cycle.
func (r *Repo) ClaimDetailByRound(ctx context.Context, round int64) (cgmodel.CGRoundClaimDetail, error) {
	detail := cgmodel.CGRoundClaimDetail{
		RoundNum:          round,
		ClaimTransactions: make([]cgmodel.CGClaimTxn, 0),
		AttachedTokens:    make([]cgmodel.CGAttachedToken, 0),
	}

	// ---- Claim transactions: secondary ETH allocations ----
	ethQ := "SELECT ben.addr, win.addr, w.amount/1e18, " +
		"EXTRACT(EPOCH FROM (w.time_stamp - pc.time_stamp))::bigint, " +
		"EXTRACT(EPOCH FROM w.time_stamp)::bigint, t.tx_hash " +
		"FROM cg_prize_withdrawal w " +
		"JOIN cg_prize_claim pc ON pc.round_num=w.round_num " +
		"JOIN address ben ON ben.address_id=w.beneficiary_aid " +
		"JOIN address win ON win.address_id=w.winner_aid " +
		"LEFT JOIN transaction t ON t.id=w.tx_id " +
		"WHERE w.round_num=$1"
	ethScan := func(rows pgx.Rows, c *cgmodel.CGClaimTxn) error {
		var txh sql.NullString
		c.AssetType = "ETH"
		c.TokenId = -1
		if err := rows.Scan(&c.BeneficiaryAddr, &c.RecipientAddr, &c.AmountEth, &c.ClaimedAfterSecs, &c.ClaimTs, &txh); err != nil {
			return err
		}
		c.TxHash = txh.String
		return nil
	}
	ethClaims, err := queryList(ctx, r, "claim detail: eth claims", 8, ethQ, ethScan, round)
	if err != nil {
		return cgmodel.CGRoundClaimDetail{}, err
	}
	detail.ClaimTransactions = append(detail.ClaimTransactions, ethClaims...)

	// ---- Claim transactions: attached NFTs ----
	nftQ := "SELECT w.addr, ta.addr, dc.token_id, " +
		"EXTRACT(EPOCH FROM (dc.time_stamp - pc.time_stamp))::bigint, " +
		"EXTRACT(EPOCH FROM dc.time_stamp)::bigint, t.tx_hash " +
		"FROM cg_donated_nft_claimed dc " +
		"JOIN cg_prize_claim pc ON pc.round_num=dc.round_num " +
		"JOIN address w ON w.address_id=dc.winner_aid " +
		"JOIN address ta ON ta.address_id=dc.token_aid " +
		"LEFT JOIN transaction t ON t.id=dc.tx_id " +
		"WHERE dc.round_num=$1"
	nftScan := func(rows pgx.Rows, c *cgmodel.CGClaimTxn) error {
		var txh sql.NullString
		c.AssetType = "ERC721"
		if err := rows.Scan(&c.RecipientAddr, &c.TokenAddr, &c.TokenId, &c.ClaimedAfterSecs, &c.ClaimTs, &txh); err != nil {
			return err
		}
		c.BeneficiaryAddr = c.RecipientAddr
		c.TxHash = txh.String
		return nil
	}
	nftClaims, err := queryList(ctx, r, "claim detail: nft claims", 8, nftQ, nftScan, round)
	if err != nil {
		return cgmodel.CGRoundClaimDetail{}, err
	}
	detail.ClaimTransactions = append(detail.ClaimTransactions, nftClaims...)

	// ---- Claim transactions: attached ERC-20s ----
	ercQ := "SELECT w.addr, ta.addr, dc.amount/1e18, " +
		"EXTRACT(EPOCH FROM (dc.time_stamp - pc.time_stamp))::bigint, " +
		"EXTRACT(EPOCH FROM dc.time_stamp)::bigint, t.tx_hash " +
		"FROM cg_donated_tok_claimed dc " +
		"JOIN cg_prize_claim pc ON pc.round_num=dc.round_num " +
		"JOIN address w ON w.address_id=dc.winner_aid " +
		"JOIN address ta ON ta.address_id=dc.token_aid " +
		"LEFT JOIN transaction t ON t.id=dc.tx_id " +
		"WHERE dc.round_num=$1"
	ercScan := func(rows pgx.Rows, c *cgmodel.CGClaimTxn) error {
		var txh sql.NullString
		c.AssetType = "ERC20"
		c.TokenId = -1
		if err := rows.Scan(&c.RecipientAddr, &c.TokenAddr, &c.AmountEth, &c.ClaimedAfterSecs, &c.ClaimTs, &txh); err != nil {
			return err
		}
		c.BeneficiaryAddr = c.RecipientAddr
		c.TxHash = txh.String
		return nil
	}
	ercClaims, err := queryList(ctx, r, "claim detail: erc20 claims", 8, ercQ, ercScan, round)
	if err != nil {
		return cgmodel.CGRoundClaimDetail{}, err
	}
	detail.ClaimTransactions = append(detail.ClaimTransactions, ercClaims...)

	// ---- Attached NFTs this cycle ----
	anQ := "SELECT d.addr, ta.addr, dn.token_id, EXTRACT(EPOCH FROM dn.time_stamp)::bigint, t.tx_hash " +
		"FROM cg_nft_donation dn " +
		"JOIN address d ON d.address_id=dn.donor_aid " +
		"JOIN address ta ON ta.address_id=dn.token_aid " +
		"LEFT JOIN transaction t ON t.id=dn.tx_id " +
		"WHERE dn.round_num=$1 ORDER BY dn.idx"
	anScan := func(rows pgx.Rows, a *cgmodel.CGAttachedToken) error {
		var txh sql.NullString
		a.AssetType = "ERC721"
		if err := rows.Scan(&a.ContributorAddr, &a.TokenAddr, &a.TokenId, &a.Ts, &txh); err != nil {
			return err
		}
		a.TxHash = txh.String
		return nil
	}
	attachedNFTs, err := queryList(ctx, r, "claim detail: attached nfts", 8, anQ, anScan, round)
	if err != nil {
		return cgmodel.CGRoundClaimDetail{}, err
	}
	detail.AttachedTokens = append(detail.AttachedTokens, attachedNFTs...)

	// ---- Attached ERC-20s this cycle ----
	aeQ := "SELECT d.addr, ta.addr, e.amount/1e18, EXTRACT(EPOCH FROM e.time_stamp)::bigint, t.tx_hash " +
		"FROM cg_erc20_donation e " +
		"JOIN address d ON d.address_id=e.donor_aid " +
		"JOIN address ta ON ta.address_id=e.token_aid " +
		"LEFT JOIN transaction t ON t.id=e.tx_id " +
		"WHERE e.round_num=$1"
	aeScan := func(rows pgx.Rows, a *cgmodel.CGAttachedToken) error {
		var txh sql.NullString
		a.AssetType = "ERC20"
		a.TokenId = -1
		if err := rows.Scan(&a.ContributorAddr, &a.TokenAddr, &a.AmountEth, &a.Ts, &txh); err != nil {
			return err
		}
		a.TxHash = txh.String
		return nil
	}
	attachedERC20s, err := queryList(ctx, r, "claim detail: attached erc20s", 8, aeQ, aeScan, round)
	if err != nil {
		return cgmodel.CGRoundClaimDetail{}, err
	}
	detail.AttachedTokens = append(detail.AttachedTokens, attachedERC20s...)

	return detail, nil
}

// UniqueStakersCst returns every address that ever staked a Cosmic
// Signature token, with reward totals, highest reward first.
func (r *Repo) UniqueStakersCst(ctx context.Context) ([]cgmodel.CGUniqueStakerCST, error) {
	query := "SELECT " +
		"s.staker_aid," +
		"a.addr," +
		"s.total_tokens_staked," +
		"s.num_stake_actions," +
		"s.num_unstake_actions," +
		"s.total_reward," +
		"s.total_reward/1e18, " +
		"s.unclaimed_reward," +
		"s.unclaimed_reward/1e18 " +
		"FROM cg_staker_cst s " +
		"LEFT JOIN address a ON s.staker_aid=a.address_id " +
		"WHERE num_stake_actions> 0 " +
		"ORDER BY total_reward DESC "
	scan := func(rows pgx.Rows, rec *cgmodel.CGUniqueStakerCST) error {
		return rows.Scan(
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
	}
	return queryList(ctx, r, "unique stakers cst", 32, query, scan)
}

// UniqueStakersRwalk returns every address that ever staked a RandomWalk
// token, most tokens staked first.
func (r *Repo) UniqueStakersRwalk(ctx context.Context) ([]cgmodel.CGUniqueStakerRWalk, error) {
	query := "SELECT " +
		"s.staker_aid," +
		"a.addr," +
		"s.total_tokens_staked," +
		"s.num_stake_actions," +
		"s.num_unstake_actions," +
		"s.num_tokens_minted " +
		"FROM cg_staker_rwalk s " +
		"LEFT JOIN address a ON s.staker_aid=a.address_id " +
		"WHERE num_stake_actions > 0 " +
		"ORDER BY total_tokens_staked DESC "
	scan := func(rows pgx.Rows, rec *cgmodel.CGUniqueStakerRWalk) error {
		return rows.Scan(
			&rec.StakerAid,
			&rec.StakerAddr,
			&rec.TotalTokensStaked,
			&rec.NumStakeActions,
			&rec.NumUnstakeActions,
			&rec.TotalTokensMinted,
		)
	}
	return queryList(ctx, r, "unique stakers rwalk", 32, query, scan)
}

// UniqueStakersBoth returns the addresses that currently hold stakes in
// both token families, with per-family statistics.
func (r *Repo) UniqueStakersBoth(ctx context.Context) ([]cgmodel.CGUniqueStakersBoth, error) {
	query := "SELECT " +
		"a.address_id," +
		"a.addr," +

		"COALESCE(c.total_tokens_staked,0) cst_total_tokens_staked," +
		"COALESCE(c.num_stake_actions,0) cst_num_stake_actions," +
		"COALESCE(c.num_unstake_actions,0) cst_num_unstake_actions," +
		"COALESCE(c.total_reward,0) cst_total_reward," +
		"COALESCE(c.total_reward/1e18,0) cst_total_reward_eth," +
		"COALESCE(c.unclaimed_reward,0) cst_unclaimed_reward," +
		"COALESCE(c.unclaimed_reward/1e18,0) cst_unclaimed_reward_eth, " +

		"COALESCE(r.total_tokens_staked,0) rw_total_tokens_staked," +
		"COALESCE(r.num_stake_actions,0) rw_num_stake_actions," +
		"COALESCE(r.num_unstake_actions,0) rw_num_unstake_actions," +
		"COALESCE(r.num_tokens_minted,0) rw_num_tokens_minted," +

		"(COALESCE(c.total_tokens_staked,0) + COALESCE(r.total_tokens_staked,0)) all_tokens_num " +

		"FROM address a " +
		"LEFT JOIN cg_staker_cst c ON a.address_id = c.staker_aid " +
		"LEFT JOIN cg_staker_rwalk r ON a.address_id = r.staker_aid " +
		"WHERE " +
		"(COALESCE(c.total_tokens_staked,0)>0) AND (COALESCE(r.total_tokens_staked,0) > 0) "
	scan := func(rows pgx.Rows, rec *cgmodel.CGUniqueStakersBoth) error {
		return rows.Scan(
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
	}
	return queryList(ctx, r, "unique stakers both", 32, query, scan)
}

// UniqueDonors returns every address that donated ETH, largest total first.
func (r *Repo) UniqueDonors(ctx context.Context) ([]cgmodel.CGUniqueDonor, error) {
	query := "SELECT " +
		"d.donor_aid," +
		"a.addr," +
		"d.count_donations," +
		"d.total_eth_donated," +
		"d.total_eth_donated/1e18 total_eth_donated_eth " +
		"FROM cg_donor d " +
		"LEFT JOIN address a ON d.donor_aid=a.address_id " +
		"WHERE d.count_donations > 0 " +
		"ORDER BY total_eth_donated DESC "
	scan := func(rows pgx.Rows, rec *cgmodel.CGUniqueDonor) error {
		return rows.Scan(
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.CountDonations,
			&rec.TotalDonated,
			&rec.TotalDonatedEth,
		)
	}
	return queryList(ctx, r, "unique donors", 32, query, scan)
}

// NFTDonationStats returns per-contract donated-NFT counts.
func (r *Repo) NFTDonationStats(ctx context.Context) ([]cgmodel.CGNFTDonationStats, error) {
	query := "SELECT " +
		"s.contract_aid," +
		"a.addr," +
		"s.num_donated " +
		"FROM cg_nft_stats s " +
		"LEFT JOIN address a ON s.contract_aid=a.address_id "
	scan := func(rows pgx.Rows, rec *cgmodel.CGNFTDonationStats) error {
		return rows.Scan(
			&rec.TokenAddressId,
			&rec.TokenAddress,
			&rec.NumDonations,
		)
	}
	return queryList(ctx, r, "nft donation stats", 256, query, scan)
}

// RecordCounters returns the raw row counts of the bid, prize-claim and
// NFT-donation tables.
func (r *Repo) RecordCounters(ctx context.Context) (cgmodel.CGRecordCounters, error) {
	const op = "record counters"
	var output cgmodel.CGRecordCounters
	err := r.q(ctx).QueryRow(ctx, "SELECT count(*) AS total FROM cg_bid").Scan(&output.TotalBids)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGRecordCounters{}, store.WrapError(op+": bids", err)
	}
	err = r.q(ctx).QueryRow(ctx, "SELECT count(*) AS total FROM cg_prize_claim").Scan(&output.TotalPrizes)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGRecordCounters{}, store.WrapError(op+": prizes", err)
	}
	err = r.q(ctx).QueryRow(ctx, "SELECT count(*) AS total FROM cg_nft_donation").Scan(&output.TotalDonatedNFTs)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return cgmodel.CGRecordCounters{}, store.WrapError(op+": nft donations", err)
	}
	return output, nil
}
