package cosmicgame

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

const prizeClaimsSelect = `SELECT
			p.evtlog_id,
			p.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,
			p.time_stamp,
			p.winner_aid,
			wa.addr,
			p.timeout,
			p.amount,
			p.amount/1e18 amount_eth,
			p.round_num,
			p.token_id,
			m.seed,
			s.total_bids,
			s.total_nft_donated,
			s.num_erc20_donations,
			s.total_raffle_eth_deposits,
			s.total_raffle_eth_deposits/1e18 eth_deposits,
			s.total_raffle_nfts,
			d.donation_amount,
			d.donation_amount/1e18 AS amount_eth,
			d.charity_addr,
			dp.deposit_amount,
			dp.deposit_amount/1e18,
			dp.amount_per_token,
			dp.amount_per_token/1e18,
			dp.deposit_id,
			dp.num_staked_nfts,
			p.cst_amount,
			p.cst_amount/1e18,
			end_a.addr,
			endu.erc721_token_id,
			endu.erc20_amount,
			endu.erc20_amount/1e18,
			cw_a.addr,
			cw.eth_amount,
			cw.eth_amount/1e18,
			cw.cst_amount,
			cw.cst_amount/1e18,
			cw.nft_id
		FROM cg_prize_claim p
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
			LEFT JOIN cg_mint_event m ON m.token_id=p.token_id
			LEFT JOIN cg_round_stats s ON p.round_num=s.round_num
			LEFT JOIN cg_endurance_prize endu ON endu.round_num=p.round_num
			LEFT JOIN address end_a ON endu.winner_aid=end_a.address_id
			LEFT JOIN cg_chrono_warrior_prize cw ON cw.round_num=p.round_num
			LEFT JOIN address cw_a ON cw.winner_aid=cw_a.address_id
			LEFT JOIN cg_staking_eth_deposit dp ON dp.round_num=p.round_num
			LEFT JOIN (
				SELECT round_num, SUM(amount) as donation_amount, STRING_AGG(DISTINCT cha.addr, ', ') as charity_addr
					FROM cg_donation_received d
					LEFT JOIN address cha ON d.contract_aid=cha.address_id
					WHERE round_num >= 0
					GROUP BY round_num
			) d ON p.round_num = d.round_num
	`

func scanPrizeClaimRow(rows pgx.Rows, rec *p.CGRoundRec) error {
	var nullSeed sql.NullString
	var nullDepAmount, nullDepAmountPerTok sql.NullString
	var nullDepAmountEth, nullDepAmountPerTokenEth sql.NullFloat64
	var nullDepDepositNum, nullNumStakedNfts sql.NullInt64
	var nullCharityAmount, nullCharityAddr sql.NullString
	var nullCharityAmountEth sql.NullFloat64
	var nullMainCstAmount sql.NullString
	var nullMainCstEthFloat sql.NullFloat64
	var nullEnduranceAddr sql.NullString
	var nullEnduranceTid sql.NullInt64
	var nullEnduranceErc20Amount sql.NullString
	var nullEnduranceCstEth sql.NullFloat64
	var nullChronoAddr sql.NullString
	var nullChronoEthAmount, nullChronoCstAmount sql.NullString
	var nullChronoEthEth, nullChronoCstEth sql.NullFloat64
	var nullChronoNftID sql.NullInt64
	// Scan order must match prizeClaimsSelect exactly.
	err := rows.Scan(
		&rec.ClaimPrizeTx.Tx.EvtLogId,
		&rec.ClaimPrizeTx.Tx.BlockNum,
		&rec.ClaimPrizeTx.Tx.TxId,
		&rec.ClaimPrizeTx.Tx.TxHash,
		&rec.ClaimPrizeTx.Tx.TimeStamp,
		store.TimeText(&rec.ClaimPrizeTx.Tx.DateTime),
		&rec.MainPrize.WinnerAid,
		&rec.MainPrize.WinnerAddr,
		&rec.MainPrize.TimeoutTs,
		&rec.MainPrize.EthAmount,
		&rec.MainPrize.EthAmountEth,
		&rec.RoundNum,
		&rec.MainPrize.NftTokenId,
		&nullSeed,
		&rec.RoundStats.TotalBids,
		&rec.RoundStats.TotalDonatedNFTs,
		&rec.RoundStats.NumERC20Donations,
		&rec.RoundStats.TotalRaffleEthDeposits,
		&rec.RoundStats.TotalRaffleEthDepositsEth,
		&rec.RoundStats.TotalRaffleNFTs,
		&nullCharityAmount,
		&nullCharityAmountEth,
		&nullCharityAddr,
		&nullDepAmount,
		&nullDepAmountEth,
		&nullDepAmountPerTok,
		&nullDepAmountPerTokenEth,
		&nullDepDepositNum,
		&nullNumStakedNfts,
		&nullMainCstAmount,
		&nullMainCstEthFloat,
		&nullEnduranceAddr,
		&nullEnduranceTid,
		&nullEnduranceErc20Amount,
		&nullEnduranceCstEth,
		&nullChronoAddr,
		&nullChronoEthAmount,
		&nullChronoEthEth,
		&nullChronoCstAmount,
		&nullChronoCstEth,
		&nullChronoNftID,
	)
	if err != nil {
		return err
	}
	if nullSeed.Valid {
		rec.MainPrize.Seed = nullSeed.String
	} else {
		rec.MainPrize.Seed = "???"
	}
	if nullMainCstAmount.Valid {
		rec.MainPrize.CstAmount = nullMainCstAmount.String
		if nullMainCstEthFloat.Valid {
			rec.MainPrize.CstAmountEth = nullMainCstEthFloat.Float64
		}
	}
	if nullEnduranceAddr.Valid {
		rec.EnduranceChampion.WinnerAddr = nullEnduranceAddr.String
	}
	if nullEnduranceTid.Valid {
		rec.EnduranceChampion.NftTokenId = nullEnduranceTid.Int64
	}
	if nullEnduranceErc20Amount.Valid {
		rec.EnduranceChampion.CstAmount = nullEnduranceErc20Amount.String
		if nullEnduranceCstEth.Valid {
			rec.EnduranceChampion.CstAmountEth = nullEnduranceCstEth.Float64
		}
	}
	if nullChronoAddr.Valid {
		rec.ChronoWarrior.WinnerAddr = nullChronoAddr.String
	}
	if nullChronoEthAmount.Valid {
		rec.ChronoWarrior.EthAmount = nullChronoEthAmount.String
		if nullChronoEthEth.Valid {
			rec.ChronoWarrior.EthAmountEth = nullChronoEthEth.Float64
		}
	}
	if nullChronoCstAmount.Valid {
		rec.ChronoWarrior.CstAmount = nullChronoCstAmount.String
		if nullChronoCstEth.Valid {
			rec.ChronoWarrior.CstAmountEth = nullChronoCstEth.Float64
		}
	}
	if nullChronoNftID.Valid {
		rec.ChronoWarrior.NftTokenId = nullChronoNftID.Int64
	}
	if nullCharityAmount.Valid {
		rec.CharityDeposit.CharityAmount = nullCharityAmount.String
	}
	if nullCharityAmountEth.Valid {
		rec.CharityDeposit.CharityAmountETH = nullCharityAmountEth.Float64
	}
	if nullCharityAddr.Valid {
		rec.CharityDeposit.CharityAddress = nullCharityAddr.String
	}
	applyStakingDeposit(&rec.StakingDeposit, nullDepAmount, nullDepAmountEth, nullDepAmountPerTok, nullDepAmountPerTokenEth, nullDepDepositNum, nullNumStakedNfts)
	return nil
}

// PrizeClaims returns the main prize claims of all rounds (newest first,
// offset/limit paginated; limit 0 means unbounded) with the per-round stats,
// charity, staking, endurance and chrono warrior columns the round list needs.
func (r *Repo) PrizeClaims(ctx context.Context, offset, limit int) ([]p.CGRoundRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := prizeClaimsSelect + " ORDER BY p.id DESC OFFSET $1 LIMIT $2"
	return queryList(ctx, r, "prize claims", 256, query, scanPrizeClaimRow, offset, limit)
}

// RoundPageCursor identifies the last completed round returned by
// PrizeClaimsPage.
type RoundPageCursor struct {
	RoundNum   int64
	EventLogID int64
}

// PrizeClaimsPage returns at most limit completed rounds after the supplied
// descending keyset cursor. A nil cursor starts at the newest round.
func (r *Repo) PrizeClaimsPage(ctx context.Context, after *RoundPageCursor, limit int) (records []p.CGRoundRec, hasMore bool, err error) {
	const op = "prize claims page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: limit must be positive", op)
	}

	query := prizeClaimsSelect + " ORDER BY p.round_num DESC, p.evtlog_id DESC LIMIT $1"
	args := []any{limit + 1}
	if after != nil {
		if after.RoundNum < 0 || after.EventLogID < 1 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		query = prizeClaimsSelect +
			" WHERE (p.round_num, p.evtlog_id) < ($1, $2)" +
			" ORDER BY p.round_num DESC, p.evtlog_id DESC LIMIT $3"
		args = []any{after.RoundNum, after.EventLogID, limit + 1}
	}

	records, err = queryList(ctx, r, op, limit+1, query, scanPrizeClaimRow, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		records = records[:limit]
		hasMore = true
	}
	return records, hasMore, nil
}

// applyStakingDeposit copies the nullable staking-deposit columns into the
// record, defaulting the deposit id to -1 when the round has no deposit.
func applyStakingDeposit(dst *p.CGStakingDeposit, amount sql.NullString, amountEth sql.NullFloat64, perTok sql.NullString, perTokEth sql.NullFloat64, depositNum, numStaked sql.NullInt64) {
	if amount.Valid {
		dst.StakingDepositAmount = amount.String
	}
	if amountEth.Valid {
		dst.StakingDepositAmountEth = amountEth.Float64
	}
	if perTok.Valid {
		dst.StakingPerToken = perTok.String
	}
	if perTokEth.Valid {
		dst.StakingPerTokenEth = perTokEth.Float64
	}
	if depositNum.Valid {
		dst.StakingDepositId = depositNum.Int64
	} else {
		dst.StakingDepositId = -1
	}
	if numStaked.Valid {
		dst.StakingNumStakedTokens = numStaked.Int64
	}
}

// RoundInfo returns the lean detail of one completed round without loading
// the legacy nested prize collections. A round with no prize claim yet yields
// store.ErrNotFound.
func (r *Repo) RoundInfo(ctx context.Context, roundNum int64) (p.CGRoundRec, error) {
	const op = "prize info"
	var rec p.CGRoundRec
	query := `SELECT
			p.evtlog_id,
			p.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,
			p.time_stamp,
			p.winner_aid,
			wa.addr,
			p.timeout,
			p.amount,
			p.amount/1e18 amount_eth,
			p.cst_amount,
			p.cst_amount/1e18 cst_amount_eth,
			p.round_num,
			p.token_id,
			m.seed,
			s.total_bids,
			s.total_nft_donated,
			s.num_erc20_donations,
			s.total_raffle_eth_deposits,
			s.total_raffle_eth_deposits/1e18,
			s.total_raffle_nfts,
			d.donation_amount,
			d.donation_amount/1e+18,
			d.charity_addr,
			dp.deposit_amount,
			dp.deposit_amount/1e18,
			dp.amount_per_token,
			dp.amount_per_token/1e18,
			dp.deposit_id,
			dp.num_staked_nfts,
			endu.erc721_token_id,
			end_a.addr,
			top.erc721_token_id,
			top_a.addr,
			w_a.addr,
			endu.erc20_amount,
			endu.erc20_amount/1e18,
			top.erc20_amount,
			top.erc20_amount/1e18,
			w.eth_amount,
			w.eth_amount/1e18,
			w.cst_amount,
			w.cst_amount/1e18,
			w.nft_id,
			s.donations_round_count,
			s.donations_round_total,
			s.donations_round_total/1e18,
			s.param_window_start_time,
			EXTRACT(EPOCH FROM s.activation_time)::BIGINT,
			s.param_window_duration_seconds,
			s.round_start_time,
			s.round_end_time,
			s.round_duration_seconds
		FROM cg_prize_claim p
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
			LEFT JOIN cg_mint_event m ON m.token_id=p.token_id
			LEFT JOIN cg_staking_eth_deposit dp ON dp.round_num=p.round_num
			LEFT JOIN cg_round_stats s ON s.round_num=p.round_num
			LEFT JOIN cg_winner ws ON p.winner_aid=ws.winner_aid
			LEFT JOIN cg_endurance_prize endu ON endu.round_num=p.round_num
			LEFT JOIN cg_lastcst_prize top ON top.round_num=p.round_num
			LEFT JOIN cg_chrono_warrior_prize w ON w.round_num = p.round_num
			LEFT JOIN address end_a ON endu.winner_aid=end_a.address_id
			LEFT JOIN address top_a ON top.winner_aid=top_a.address_id
			LEFT JOIN address w_a ON w.winner_aid=w_a.address_id
			LEFT JOIN (
				SELECT round_num, SUM(amount) as donation_amount, STRING_AGG(DISTINCT cha.addr, ', ') as charity_addr
					FROM cg_donation_received d
					LEFT JOIN address cha ON d.contract_aid=cha.address_id
					WHERE round_num >= 0
					GROUP BY round_num
			) d ON p.round_num = d.round_num
		WHERE p.round_num=$1`

	row := r.pool().QueryRow(ctx, query, roundNum)
	var nullSeed sql.NullString
	var nullDepAmount, nullDepAmountPerTok sql.NullString
	var nullDepAmountEth, nullDepAmountPerTokenEth sql.NullFloat64
	var nullDepDepositNum, nullNumStakedNfts sql.NullInt64
	var nullMainCstAmount sql.NullString
	var nullMainCstAmountEth sql.NullFloat64
	var nullCharityAmount, nullCharityAddr sql.NullString
	var nullCharityAmountEth sql.NullFloat64
	var nullEnduranceTid, nullLastCstTid, nullWarriorNftID sql.NullInt64
	var nullEnduranceAddr, nullLastCstAddr, nullWarriorAddr sql.NullString
	var nullEnduranceErc20Amount, nullLastCstErc20Amount, nullWarriorEthAmount, nullWarriorCstAmount sql.NullString
	var nullEnduranceErc20AmountFloat, nullLastCstErc20AmountFloat, nullWarriorEthAmountFloat, nullWarriorCstAmountFloat sql.NullFloat64
	// Round timing fields (nullable)
	var nullParamWindowStart sql.NullString
	var nullActivationTime sql.NullInt64
	var nullParamWindowDuration sql.NullInt64
	var nullRoundStartTime sql.NullString
	var nullRoundEndTime sql.NullString
	var nullRoundDuration sql.NullInt64
	err := row.Scan(
		&rec.ClaimPrizeTx.Tx.EvtLogId,
		&rec.ClaimPrizeTx.Tx.BlockNum,
		&rec.ClaimPrizeTx.Tx.TxId,
		&rec.ClaimPrizeTx.Tx.TxHash,
		&rec.ClaimPrizeTx.Tx.TimeStamp,
		store.TimeText(&rec.ClaimPrizeTx.Tx.DateTime),
		&rec.MainPrize.WinnerAid,
		&rec.MainPrize.WinnerAddr,
		&rec.MainPrize.TimeoutTs,
		&rec.MainPrize.EthAmount,
		&rec.MainPrize.EthAmountEth,
		&nullMainCstAmount,
		&nullMainCstAmountEth,
		&rec.RoundNum,
		&rec.MainPrize.NftTokenId,
		&nullSeed,
		&rec.RoundStats.TotalBids,
		&rec.RoundStats.TotalDonatedNFTs,
		&rec.RoundStats.NumERC20Donations,
		&rec.RoundStats.TotalRaffleEthDeposits,
		&rec.RoundStats.TotalRaffleEthDepositsEth,
		&rec.RoundStats.TotalRaffleNFTs,
		&nullCharityAmount,
		&nullCharityAmountEth,
		&nullCharityAddr,
		&nullDepAmount,
		&nullDepAmountEth,
		&nullDepAmountPerTok,
		&nullDepAmountPerTokenEth,
		&nullDepDepositNum,
		&nullNumStakedNfts,
		&nullEnduranceTid,
		&nullEnduranceAddr,
		&nullLastCstTid,
		&nullLastCstAddr,
		&nullWarriorAddr,
		&nullEnduranceErc20Amount,
		&nullEnduranceErc20AmountFloat,
		&nullLastCstErc20Amount,
		&nullLastCstErc20AmountFloat,
		&nullWarriorEthAmount,
		&nullWarriorEthAmountFloat,
		&nullWarriorCstAmount,
		&nullWarriorCstAmountFloat,
		&nullWarriorNftID,
		&rec.RoundStats.TotalDonatedCount,
		&rec.RoundStats.TotalDonatedAmount,
		&rec.RoundStats.TotalDonatedAmountEth,
		&nullParamWindowStart,
		&nullActivationTime,
		&nullParamWindowDuration,
		&nullRoundStartTime,
		&nullRoundEndTime,
		&nullRoundDuration,
	)
	if err != nil {
		return rec, store.WrapError(op, err)
	}
	if nullSeed.Valid {
		rec.MainPrize.Seed = nullSeed.String
	} else {
		rec.MainPrize.Seed = "???"
	}
	if nullMainCstAmount.Valid {
		rec.MainPrize.CstAmount = nullMainCstAmount.String
		rec.MainPrize.CstAmountEth = nullMainCstAmountEth.Float64
	}
	if nullCharityAmount.Valid {
		rec.CharityDeposit.CharityAmount = nullCharityAmount.String
	}
	if nullCharityAmountEth.Valid {
		rec.CharityDeposit.CharityAmountETH = nullCharityAmountEth.Float64
	}
	if nullCharityAddr.Valid {
		rec.CharityDeposit.CharityAddress = nullCharityAddr.String
	}

	applyStakingDeposit(&rec.StakingDeposit, nullDepAmount, nullDepAmountEth, nullDepAmountPerTok, nullDepAmountPerTokenEth, nullDepDepositNum, nullNumStakedNfts)
	if nullEnduranceTid.Valid {
		rec.EnduranceChampion.WinnerAddr = nullEnduranceAddr.String
		rec.EnduranceChampion.NftTokenId = nullEnduranceTid.Int64
	}
	if nullLastCstTid.Valid {
		rec.LastCstBidder.WinnerAddr = nullLastCstAddr.String
		rec.LastCstBidder.NftTokenId = nullLastCstTid.Int64
	}
	if nullEnduranceErc20Amount.Valid {
		rec.EnduranceChampion.CstAmount = nullEnduranceErc20Amount.String
		rec.EnduranceChampion.CstAmountEth = nullEnduranceErc20AmountFloat.Float64
	}
	if nullLastCstErc20Amount.Valid {
		rec.LastCstBidder.CstAmount = nullLastCstErc20Amount.String
		rec.LastCstBidder.CstAmountEth = nullLastCstErc20AmountFloat.Float64
	}
	if nullWarriorEthAmount.Valid {
		rec.ChronoWarrior.EthAmount = nullWarriorEthAmount.String
		rec.ChronoWarrior.EthAmountEth = nullWarriorEthAmountFloat.Float64
	}
	if nullWarriorCstAmount.Valid {
		rec.ChronoWarrior.CstAmount = nullWarriorCstAmount.String
		rec.ChronoWarrior.CstAmountEth = nullWarriorCstAmountFloat.Float64
	}
	if nullWarriorNftID.Valid {
		rec.ChronoWarrior.NftTokenId = nullWarriorNftID.Int64
	}
	if nullWarriorAddr.Valid {
		rec.ChronoWarrior.WinnerAddr = nullWarriorAddr.String
	}

	// Round timing fields
	if nullParamWindowStart.Valid {
		rec.RoundStats.ParamWindowStartTime = nullParamWindowStart.String
	}
	if nullActivationTime.Valid {
		rec.RoundStats.ActivationTime = nullActivationTime.Int64
	}
	if nullParamWindowDuration.Valid {
		rec.RoundStats.ParamWindowDurationSeconds = nullParamWindowDuration.Int64
	}
	if nullRoundStartTime.Valid {
		rec.RoundStats.RoundStartTime = nullRoundStartTime.String
	}
	if nullRoundEndTime.Valid {
		rec.RoundStats.RoundEndTime = nullRoundEndTime.String
	}
	if nullRoundDuration.Valid {
		rec.RoundStats.RoundDurationSeconds = nullRoundDuration.Int64
	}

	return rec, nil
}

// PrizeInfo preserves the v1 round detail by composing RoundInfo with the
// nested raffle, staking, deposit and prize collections.
func (r *Repo) PrizeInfo(ctx context.Context, roundNum int64) (p.CGRoundRec, error) {
	const op = "prize info"
	rec, err := r.RoundInfo(ctx, roundNum)
	if err != nil {
		return rec, err
	}
	if rec.RaffleNFTWinners, err = r.RaffleNFTWinnersByRound(ctx, roundNum, false); err != nil {
		return rec, store.WrapError(op, err)
	}
	if rec.StakingNFTWinners, err = r.RaffleNFTWinnersByRound(ctx, roundNum, true); err != nil {
		return rec, store.WrapError(op, err)
	}
	if rec.RaffleETHDeposits, err = r.PrizeDepositsByRound(ctx, roundNum); err != nil {
		return rec, store.WrapError(op, err)
	}
	if rec.AllPrizes, err = r.AllPrizesForRound(ctx, roundNum); err != nil {
		return rec, store.WrapError(op, err)
	}
	return rec, nil
}

// AllPrizesForRound returns every prize row of one round from the cg_prize
// registry with the type-specific amount/token columns resolved, ordered by
// prize type then winner index.
func (r *Repo) AllPrizesForRound(ctx context.Context, roundNum int64) ([]p.CGPrizeHistory, error) {
	query := `SELECT
			p.ptype AS record_type,
			COALESCE(pc.evtlog_id, rew.evtlog_id, rnw_bidder.evtlog_id, rnw_rwalk.evtlog_id, ew.evtlog_id, lw.evtlog_id, cw.evtlog_id, ed.evtlog_id) AS evtlog_id,
			COALESCE(EXTRACT(EPOCH FROM pc.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rnw_bidder.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rnw_rwalk.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM lw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM cw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ed.time_stamp)::BIGINT) AS tstmp,
			COALESCE(pc.time_stamp, rew.time_stamp, rnw_bidder.time_stamp, rnw_rwalk.time_stamp, ew.time_stamp, lw.time_stamp, cw.time_stamp, ed.time_stamp) AS date_time,
			COALESCE(pc.block_num, rew.block_num, rnw_bidder.block_num, rnw_rwalk.block_num, ew.block_num, lw.block_num, cw.block_num, ed.block_num) AS block_num,
			COALESCE(tc.id, trew.id, trnw_bidder.id, trnw_rwalk.id, tew.id, tlw.id, tcw.id, ted.id) AS tx_id,
			COALESCE(tc.tx_hash, trew.tx_hash, trnw_bidder.tx_hash, trnw_rwalk.tx_hash, tew.tx_hash, tlw.tx_hash, tcw.tx_hash, ted.tx_hash) AS tx_hash,
			p.round_num,
			CASE
				WHEN p.ptype = 0 THEN pc.amount
				WHEN p.ptype = 1 THEN pc.cst_amount
				WHEN p.ptype = 4 THEN lw.erc20_amount
				WHEN p.ptype = 6 THEN ew.erc20_amount
				WHEN p.ptype = 7 THEN cw.eth_amount
				WHEN p.ptype = 8 THEN cw.cst_amount
				WHEN p.ptype = 10 THEN rew.amount
				WHEN p.ptype = 11 THEN rnw_bidder.cst_amount
				WHEN p.ptype = 13 THEN rnw_rwalk.cst_amount
				WHEN p.ptype = 15 THEN ed.deposit_amount
				ELSE '0'
			END AS amount,
			CASE
				WHEN p.ptype = 0 THEN pc.amount/1e18
				WHEN p.ptype = 1 THEN pc.cst_amount/1e18
				WHEN p.ptype = 4 THEN lw.erc20_amount/1e18
				WHEN p.ptype = 6 THEN ew.erc20_amount/1e18
				WHEN p.ptype = 7 THEN cw.eth_amount/1e18
				WHEN p.ptype = 8 THEN cw.cst_amount/1e18
				WHEN p.ptype = 10 THEN rew.amount/1e18
				WHEN p.ptype = 11 THEN rnw_bidder.cst_amount/1e18
				WHEN p.ptype = 13 THEN rnw_rwalk.cst_amount/1e18
				WHEN p.ptype = 15 THEN ed.deposit_amount/1e18
				ELSE 0
			END AS amount_eth,
			'' AS token_addr,
			CASE
				WHEN p.ptype = 2 THEN pc.token_id
				WHEN p.ptype = 3 THEN lw.erc721_token_id
				WHEN p.ptype = 5 THEN ew.erc721_token_id
				WHEN p.ptype = 9 THEN cw.nft_id
				WHEN p.ptype = 12 THEN rnw_bidder.token_id
				WHEN p.ptype = 14 THEN rnw_rwalk.token_id
				ELSE -1
			END AS token_id,
			'' AS token_uri,
			p.winner_index,
			TRUE AS claimed,
			CASE WHEN p.ptype = 15 THEN '(All CS NFT Stakers)' ELSE COALESCE(wa_pc.addr, wa_rew.addr, wa_rnw_bidder.addr, wa_rnw_rwalk.addr, wa_ew.addr, wa_lw.addr, wa_cw.addr, '') END AS winner_addr,
			COALESCE(pc.winner_aid, rew.winner_aid, rnw_bidder.winner_aid, rnw_rwalk.winner_aid, ew.winner_aid, lw.winner_aid, cw.winner_aid, 0) AS winner_aid
		FROM cg_prize p
			LEFT JOIN cg_prize_claim pc ON (p.round_num = pc.round_num AND p.ptype IN (0,1,2))
			LEFT JOIN transaction tc ON tc.id = pc.tx_id
			LEFT JOIN address wa_pc ON pc.winner_aid = wa_pc.address_id
			LEFT JOIN cg_lastcst_prize lw ON (p.round_num = lw.round_num AND p.ptype IN (3,4))
			LEFT JOIN transaction tlw ON tlw.id = lw.tx_id
			LEFT JOIN address wa_lw ON lw.winner_aid = wa_lw.address_id
			LEFT JOIN cg_endurance_prize ew ON (p.round_num = ew.round_num AND p.ptype IN (5,6))
			LEFT JOIN transaction tew ON tew.id = ew.tx_id
			LEFT JOIN address wa_ew ON ew.winner_aid = wa_ew.address_id
			LEFT JOIN cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index AND p.ptype IN (7,8,9))
			LEFT JOIN transaction tcw ON tcw.id = cw.tx_id
			LEFT JOIN address wa_cw ON cw.winner_aid = wa_cw.address_id
			LEFT JOIN cg_raffle_eth_prize rew ON (p.round_num = rew.round_num AND p.winner_index = rew.winner_idx AND p.ptype = 10)
			LEFT JOIN transaction trew ON trew.id = rew.tx_id
			LEFT JOIN address wa_rew ON rew.winner_aid = wa_rew.address_id
			LEFT JOIN cg_raffle_nft_prize rnw_bidder ON (p.round_num = rnw_bidder.round_num AND p.winner_index = rnw_bidder.winner_idx AND p.ptype IN (11,12) AND rnw_bidder.is_rwalk = false)
			LEFT JOIN transaction trnw_bidder ON trnw_bidder.id = rnw_bidder.tx_id
			LEFT JOIN address wa_rnw_bidder ON rnw_bidder.winner_aid = wa_rnw_bidder.address_id
			LEFT JOIN cg_raffle_nft_prize rnw_rwalk ON (p.round_num = rnw_rwalk.round_num AND p.winner_index = rnw_rwalk.winner_idx AND p.ptype IN (13,14) AND rnw_rwalk.is_rwalk = true)
			LEFT JOIN transaction trnw_rwalk ON trnw_rwalk.id = rnw_rwalk.tx_id
			LEFT JOIN address wa_rnw_rwalk ON rnw_rwalk.winner_aid = wa_rnw_rwalk.address_id
			LEFT JOIN cg_staking_eth_deposit ed ON (p.round_num = ed.round_num AND p.ptype = 15)
			LEFT JOIN transaction ted ON ted.id = ed.tx_id
		WHERE p.round_num = $1
			ORDER BY p.ptype, p.winner_index`
	scan := func(rows pgx.Rows, rec *p.CGPrizeHistory) error {
		return rows.Scan(
			&rec.RecordType,
			&rec.Tx.EvtLogId,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.RoundNum,
			&rec.Amount,
			&rec.AmountEth,
			&rec.TokenAddress,
			&rec.TokenId,
			&rec.TokenURI,
			&rec.WinnerIndex,
			&rec.Claimed,
			&rec.WinnerAddr,
			&rec.WinnerAid,
		)
	}
	return queryList(ctx, r, "all prizes for round", 64, query, scan, roundNum)
}
