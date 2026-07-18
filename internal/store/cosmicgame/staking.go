package cosmicgame

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// stakeActionQueryCST returns the unified stake/unstake info query for the
// CST tables (with reward columns). The table names are compile-time
// constants at the single call site; the function exists so the SQL shape is
// unit-testable without a database.
func stakeActionQueryCST(stakeTable, unstakeTable string) string {
	return "SELECT " +
		"st.id," +
		"st.evtlog_id," +
		"st.block_num," +
		"ts.id," +
		"ts.tx_hash," +
		"EXTRACT(EPOCH FROM st.time_stamp)::BIGINT," +
		"st.time_stamp," +
		"st.action_id," +
		"st.token_id," +
		"st.round_num," +
		"st.num_staked_nfts," +
		"st.staker_aid," +
		"sa.addr," +
		"u.id," +
		"u.evtlog_id," +
		"u.block_num," +
		"tu.id," +
		"tu.tx_hash," +
		"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT," +
		"u.time_stamp," +
		"u.action_id," +
		"u.token_id, " +
		"u.round_num," +
		"u.num_staked_nfts, " +
		"u.reward," +
		"u.reward/1e18," +
		"u.reward_per_tok," +
		"u.reward_per_tok/1e18," +
		"u.staker_aid, " +
		"ua.addr " +
		"FROM " + stakeTable + " st " +
		"LEFT JOIN transaction ts ON ts.id=st.tx_id " +
		"LEFT JOIN address sa ON st.staker_aid=sa.address_id " +
		"LEFT JOIN " + unstakeTable + " u ON st.action_id=u.action_id " +
		"LEFT JOIN transaction tu ON tu.id=u.tx_id " +
		"LEFT JOIN address ua ON u.staker_aid=ua.address_id " +
		"WHERE st.action_id=$1"
}

// stakeActionQueryRWalk is stakeActionQueryCST for the RandomWalk tables,
// whose unstake rows carry no reward columns.
func stakeActionQueryRWalk(stakeTable, unstakeTable string) string {
	return "SELECT " +
		"st.id," +
		"st.evtlog_id," +
		"st.block_num," +
		"ts.id," +
		"ts.tx_hash," +
		"EXTRACT(EPOCH FROM st.time_stamp)::BIGINT," +
		"st.time_stamp," +
		"st.action_id," +
		"st.token_id," +
		"st.round_num," +
		"st.num_staked_nfts," +
		"st.staker_aid," +
		"sa.addr," +
		"u.id," +
		"u.evtlog_id," +
		"u.block_num," +
		"tu.id," +
		"tu.tx_hash," +
		"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT," +
		"u.time_stamp," +
		"u.action_id," +
		"u.token_id, " +
		"u.round_num," +
		"u.num_staked_nfts, " +
		"u.staker_aid, " +
		"ua.addr " +
		"FROM " + stakeTable + " st " +
		"LEFT JOIN transaction ts ON ts.id=st.tx_id " +
		"LEFT JOIN address sa ON st.staker_aid=sa.address_id " +
		"LEFT JOIN " + unstakeTable + " u ON st.action_id=u.action_id " +
		"LEFT JOIN transaction tu ON tu.id=u.tx_id " +
		"LEFT JOIN address ua ON u.staker_aid=ua.address_id " +
		"WHERE st.action_id=$1"
}

// StakeActionCstInfo returns the combined stake/unstake record of one CST
// stake action (unstake fields zero-valued while the token is still staked),
// or store.ErrNotFound for an unknown action id.
func (r *Repo) StakeActionCstInfo(ctx context.Context, actionID int64) (cgmodel.CGStakeUnstakeCombined, error) {
	const op = "stake action cst info"
	query := stakeActionQueryCST("cg_nft_staked_cst", "cg_nft_unstaked_cst")

	var rec cgmodel.CGStakeUnstakeCombined
	var nullRecordID, nullEvtlogID, nullTxID, nullUnstakeTs, nullActionID sql.NullInt64
	var nullBlockNum, nullTokenID, nullRoundNum, nullNumStakedNFTs, nullStakerAid sql.NullInt64
	var nullTxHash, nullStakerAddr, nullReward, nullRewardPerTok sql.NullString
	var nullRewardEth, nullRewardPerTokEth sql.NullFloat64

	err := r.q(ctx).QueryRow(ctx, query, actionID).Scan(
		&rec.Stake.RecordId, &rec.Stake.Tx.EvtLogId, &rec.Stake.Tx.BlockNum, &rec.Stake.Tx.TxId, &rec.Stake.Tx.TxHash,
		&rec.Stake.Tx.TimeStamp, store.TimeText(&rec.Stake.Tx.DateTime), &rec.Stake.ActionId, &rec.Stake.TokenId,
		&rec.Stake.RoundNum, &rec.Stake.NumStakedNFTs, &rec.Stake.StakerAid, &rec.Stake.StakerAddr,
		&nullRecordID, &nullEvtlogID, &nullBlockNum, &nullTxID, &nullTxHash,
		&nullUnstakeTs, store.NullTimeText(&rec.Unstake.Tx.DateTime), &nullActionID, &nullTokenID,
		&nullRoundNum, &nullNumStakedNFTs, &nullReward, &nullRewardEth,
		&nullRewardPerTok, &nullRewardPerTokEth, &nullStakerAid, &nullStakerAddr,
	)
	if err != nil {
		return cgmodel.CGStakeUnstakeCombined{}, store.WrapError(op, err)
	}

	if nullRecordID.Valid {
		rec.Unstake.RecordId = nullRecordID.Int64
	}
	if nullEvtlogID.Valid {
		rec.Unstake.Tx.EvtLogId = nullEvtlogID.Int64
	}
	if nullBlockNum.Valid {
		rec.Unstake.Tx.BlockNum = nullBlockNum.Int64
	}
	if nullTxID.Valid {
		rec.Unstake.Tx.TxId = nullTxID.Int64
	}
	if nullTxHash.Valid {
		rec.Unstake.Tx.TxHash = nullTxHash.String
	}
	if nullUnstakeTs.Valid {
		rec.Unstake.Tx.TimeStamp = nullUnstakeTs.Int64
	}
	if nullActionID.Valid {
		rec.Unstake.ActionId = nullActionID.Int64
	}
	if nullTokenID.Valid {
		rec.Unstake.TokenId = nullTokenID.Int64
	}
	if nullRoundNum.Valid {
		rec.Unstake.RoundNum = nullRoundNum.Int64
	}
	if nullNumStakedNFTs.Valid {
		rec.Unstake.NumStakedNFTs = nullNumStakedNFTs.Int64
	}
	if nullReward.Valid {
		rec.Unstake.RewardAmount = nullReward.String
	}
	if nullRewardEth.Valid {
		rec.Unstake.RewardAmountEth = nullRewardEth.Float64
	}
	if nullRewardPerTok.Valid {
		rec.Unstake.RewardPerToken = nullRewardPerTok.String
	}
	if nullRewardPerTokEth.Valid {
		rec.Unstake.RewardPerTokenEth = nullRewardPerTokEth.Float64
	}
	if nullStakerAid.Valid {
		rec.Unstake.StakerAid = nullStakerAid.Int64
	}
	if nullStakerAddr.Valid {
		rec.Unstake.StakerAddr = nullStakerAddr.String
	}
	return rec, nil
}

// StakeActionRwalkInfo is StakeActionCstInfo for RandomWalk stake actions
// (whose unstake rows carry no rewards), or store.ErrNotFound.
func (r *Repo) StakeActionRwalkInfo(ctx context.Context, actionID int64) (cgmodel.CGStakeUnstakeCombined, error) {
	const op = "stake action rwalk info"
	query := stakeActionQueryRWalk("cg_nft_staked_rwalk", "cg_nft_unstaked_rwalk")

	var rec cgmodel.CGStakeUnstakeCombined
	var nullRecordID, nullEvtlogID, nullTxID, nullUnstakeTs, nullActionID sql.NullInt64
	var nullBlockNum, nullTokenID, nullRoundNum, nullNumStakedNFTs, nullStakerAid sql.NullInt64
	var nullTxHash, nullStakerAddr sql.NullString

	err := r.q(ctx).QueryRow(ctx, query, actionID).Scan(
		&rec.Stake.RecordId, &rec.Stake.Tx.EvtLogId, &rec.Stake.Tx.BlockNum, &rec.Stake.Tx.TxId, &rec.Stake.Tx.TxHash,
		&rec.Stake.Tx.TimeStamp, store.TimeText(&rec.Stake.Tx.DateTime), &rec.Stake.ActionId, &rec.Stake.TokenId,
		&rec.Stake.RoundNum, &rec.Stake.NumStakedNFTs, &rec.Stake.StakerAid, &rec.Stake.StakerAddr,
		&nullRecordID, &nullEvtlogID, &nullBlockNum, &nullTxID, &nullTxHash,
		&nullUnstakeTs, store.NullTimeText(&rec.Unstake.Tx.DateTime), &nullActionID, &nullTokenID,
		&nullRoundNum, &nullNumStakedNFTs, &nullStakerAid, &nullStakerAddr,
	)
	if err != nil {
		return cgmodel.CGStakeUnstakeCombined{}, store.WrapError(op, err)
	}
	if nullRecordID.Valid {
		rec.Unstake.RecordId = nullRecordID.Int64
	}
	if nullEvtlogID.Valid {
		rec.Unstake.Tx.EvtLogId = nullEvtlogID.Int64
	}
	if nullBlockNum.Valid {
		rec.Unstake.Tx.BlockNum = nullBlockNum.Int64
	}
	if nullTxID.Valid {
		rec.Unstake.Tx.TxId = nullTxID.Int64
	}
	if nullTxHash.Valid {
		rec.Unstake.Tx.TxHash = nullTxHash.String
	}
	if nullUnstakeTs.Valid {
		rec.Unstake.Tx.TimeStamp = nullUnstakeTs.Int64
	}
	if nullActionID.Valid {
		rec.Unstake.ActionId = nullActionID.Int64
	}
	if nullTokenID.Valid {
		rec.Unstake.TokenId = nullTokenID.Int64
	}
	if nullRoundNum.Valid {
		rec.Unstake.RoundNum = nullRoundNum.Int64
	}
	if nullNumStakedNFTs.Valid {
		rec.Unstake.NumStakedNFTs = nullNumStakedNFTs.Int64
	}
	if nullStakerAid.Valid {
		rec.Unstake.StakerAid = nullStakerAid.Int64
	}
	if nullStakerAddr.Valid {
		rec.Unstake.StakerAddr = nullStakerAddr.String
	}
	// RWalk unstake rows have no reward columns; the reward fields stay empty.
	return rec, nil
}

// StakingRewardsToBeClaimed returns, per staking deposit, the reward amounts
// userAid can still collect, newest deposit first.
func (r *Repo) StakingRewardsToBeClaimed(ctx context.Context, userAid int64) ([]cgmodel.CGRewardToClaim, error) {
	query := "WITH rwd AS (" +
		"SELECT " +
		"COUNT(token_id) AS num_toks_to_collect," +
		"SUM(reward) AS pending_reward," +
		"SUM(reward)/1e18 AS pending_reward_eth," +
		"deposit_id " +
		"FROM cg_st_reward " +
		"WHERE staker_aid=$1 AND collected='f' " +
		"GROUP BY deposit_id " +
		") " +
		"SELECT " +
		"d.id," +
		"d.evtlog_id," +
		"d.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT," +
		"d.time_stamp," +
		"EXTRACT(EPOCH FROM d.deposit_time)::BIGINT," +
		"d.deposit_time," +
		"d.deposit_id," +
		"d.accumulated_nfts," +
		"d.deposit_amount," +
		"d.deposit_amount/1e18," +
		"sd.tokens_staked," +
		"sd.amount_deposited," +
		"sd.amount_deposited/1e18," +
		"sd.amount_deposited-COALESCE(rwd.pending_reward,0)," +
		"(sd.amount_deposited-COALESCE(rwd.pending_reward,0))/1e18," +
		"rwd.pending_reward," +
		"rwd.pending_reward_eth, " +
		"rwd.num_toks_to_collect," +
		"d.deposit_amount/d.accumulated_nfts," +
		"(d.deposit_amount/d.accumulated_nfts)/1e18 " +
		"FROM cg_staker_deposit sd " +
		"INNER JOIN cg_staking_eth_deposit d ON sd.deposit_id=d.deposit_id " +
		"INNER JOIN transaction tx ON tx.id=d.tx_id " +
		"INNER JOIN rwd ON rwd.deposit_id=sd.deposit_id " +
		"WHERE (sd.staker_aid = $1) " +
		"ORDER BY d.id DESC "
	scan := func(rows pgx.Rows, rec *cgmodel.CGRewardToClaim) error {
		var nullPendingReward sql.NullString
		var nullPendingRewardEth sql.NullFloat64
		var nullPendingNumToks sql.NullInt64
		err := rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.DepositTimeStamp,
			store.TimeText(&rec.DepositDate),
			&rec.DepositId,
			&rec.NumStakedNFTs,
			&rec.DepositAmount,
			&rec.DepositAmountEth,
			&rec.YourTokensStaked,
			&rec.YourRewardAmount,
			&rec.YourRewardAmountEth,
			&rec.YourCollectedAmount,
			&rec.YourCollectedAmountEth,
			&nullPendingReward,
			&nullPendingRewardEth,
			&nullPendingNumToks,
			&rec.AmountPerToken,
			&rec.AmountPerTokenEth,
		)
		if err != nil {
			return err
		}
		if nullPendingReward.Valid {
			rec.PendingToClaim = nullPendingReward.String
		}
		if nullPendingRewardEth.Valid {
			rec.PendingToClaimEth = nullPendingRewardEth.Float64
		}
		if nullPendingNumToks.Valid {
			rec.NumUnclaimedTokens = nullPendingNumToks.Int64
		}
		return nil
	}
	return queryList(ctx, r, "staking rewards to be claimed", 16, query, scan, userAid)
}

// StakingRewardsCollected returns, per staking deposit, the rewards userAid
// already collected, newest deposit first.
func (r *Repo) StakingRewardsCollected(ctx context.Context, userAid int64, offset, limit int) ([]cgmodel.CGCollectedReward, error) {
	query := "WITH rwd AS (" +
		"SELECT " +
		"COUNT(token_id) AS num_toks_collected," +
		"SUM(reward) AS collected_reward," +
		"SUM(reward)/1e18 AS collected_reward_eth," +
		"deposit_id " +
		"FROM cg_st_reward " +
		"WHERE staker_aid=$1 AND collected='T' " +
		"GROUP BY deposit_id " +
		") " +
		"SELECT " +
		"d.id," +
		"d.evtlog_id," +
		"d.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT," +
		"d.time_stamp," +
		"EXTRACT(EPOCH FROM d.deposit_time)::BIGINT," +
		"d.deposit_time," +
		"d.deposit_id," +
		"d.accumulated_nfts," +
		"d.deposit_amount," +
		"d.deposit_amount/1e18," +
		"sd.tokens_staked," +
		"sd.amount_to_claim," +
		"sd.amount_to_claim/1e18," +
		"rwd.num_toks_collected," +
		"d.deposit_amount/accumulated_nfts," +
		"(d.deposit_amount/accumulated_nfts)/1e18, " +
		"modulo," +
		"modulo/1e+18, " +
		"d.round_num, " +
		"rwd.collected_reward," +
		"rwd.collected_reward_eth " +
		"FROM cg_staker_deposit sd " +
		"INNER JOIN cg_staking_eth_deposit d ON sd.deposit_id=d.deposit_id " +
		"INNER JOIN transaction tx ON tx.id=d.tx_id " +
		"INNER JOIN rwd ON rwd.deposit_id=sd.deposit_id " +
		"WHERE sd.staker_aid=$1 " +
		"ORDER BY d.id DESC " +
		"OFFSET $2 LIMIT $3"
	scan := func(rows pgx.Rows, rec *cgmodel.CGCollectedReward) error {
		err := rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.DepositTimeStamp,
			store.TimeText(&rec.DepositDate),
			&rec.DepositId,
			&rec.NumStakedNFTs,
			&rec.TotalDepositAmount,
			&rec.TotalDepositAmountEth,
			&rec.YourTokensStaked,
			&rec.YourAmountToClaim,
			&rec.YourAmountToClaimEth,
			&rec.NumTokensCollected,
			&rec.DepositAmountPerToken,
			&rec.DepositAmountPerTokenEth,
			&rec.Modulo,
			&rec.ModuloF64,
			&rec.RoundNum,
			&rec.YourCollectedAmount,
			&rec.YourCollectedAmountEth,
		)
		if err != nil {
			return err
		}
		if rec.YourAmountToClaimEth == rec.YourCollectedAmountEth {
			rec.FullyClaimed = true
		}
		return nil
	}
	return queryList(ctx, r, "staking rewards collected", 16, query, scan, userAid, offset, limit)
}

// StakedTokensCstGlobal returns every currently staked Cosmic Signature
// token with its mint provenance and the stake action that locked it.
func (r *Repo) StakedTokensCstGlobal(ctx context.Context) ([]cgmodel.CGStakedTokenCSTRec, error) {
	query := "SELECT " +
		"m.id," +
		"m.evtlog_id," +
		"m.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT," +
		"m.time_stamp," +
		"m.owner_aid," +
		"wa.addr," +
		"m.cur_owner_aid," +
		"oa.addr," +
		"m.seed, " +
		"m.token_id," +
		"m.round_num," +
		"p.round_num, " +
		"m.token_name, " +
		"a.evtlog_id," +
		"a.block_num," +
		"a.action_id," +
		"EXTRACT(EPOCH FROM a.time_stamp)::BIGINT," +
		"a.time_stamp," +
		"sa.addr," +
		"sa.address_id " +
		"FROM cg_staked_token_cst st " +
		"LEFT JOIN cg_mint_event m ON st.token_id=m.token_id " +
		"LEFT JOIN transaction t ON t.id=m.tx_id " +
		"LEFT JOIN address wa ON m.owner_aid=wa.address_id " +
		"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id " +
		"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id " +
		"LEFT JOIN cg_nft_staked_cst a ON a.action_id=st.stake_action_id " +
		"LEFT JOIN address sa ON a.staker_aid = sa.address_id " +
		"ORDER BY m.token_id"
	scan := func(rows pgx.Rows, rec *cgmodel.CGStakedTokenCSTRec) error {
		var nullPrizeNum sql.NullInt64
		err := rows.Scan(
			&rec.TokenInfo.RecordId,
			&rec.TokenInfo.Tx.EvtLogId,
			&rec.TokenInfo.Tx.BlockNum,
			&rec.TokenInfo.Tx.TxId,
			&rec.TokenInfo.Tx.TxHash,
			&rec.TokenInfo.Tx.TimeStamp,
			store.TimeText(&rec.TokenInfo.Tx.DateTime),
			&rec.TokenInfo.WinnerAid,
			&rec.TokenInfo.WinnerAddr,
			&rec.TokenInfo.CurOwnerAid,
			&rec.TokenInfo.CurOwnerAddr,
			&rec.TokenInfo.Seed,
			&rec.TokenInfo.TokenId,
			&rec.TokenInfo.RoundNum,
			&nullPrizeNum,
			&rec.TokenInfo.TokenName,
			&rec.StakeEvtLogId,
			&rec.StakeBlockNum,
			&rec.StakeActionId,
			&rec.StakeTimeStamp,
			store.TimeText(&rec.StakeDateTime),
			&rec.UserAddr,
			&rec.UserAid,
		)
		if err != nil {
			return err
		}
		if nullPrizeNum.Valid {
			rec.TokenInfo.RecordType = 3
		} else {
			rec.TokenInfo.RecordType = 1
		}
		return nil
	}
	return queryList(ctx, r, "staked tokens cst global", 16, query, scan)
}

// StakedTokensRwalkGlobal returns every currently staked RandomWalk token
// with its stake action.
func (r *Repo) StakedTokensRwalkGlobal(ctx context.Context) ([]cgmodel.CGStakedTokenRWalkRec, error) {
	query := "SELECT " +
		"a.evtlog_id," +
		"a.block_num," +
		"a.action_id," +
		"EXTRACT(EPOCH FROM a.time_stamp)::BIGINT," +
		"a.time_Stamp," +
		"sa.addr," +
		"sa.address_id, " +
		"st.token_id " +
		"FROM cg_staked_token_rwalk st " +
		"LEFT JOIN cg_mint_event m ON st.token_id=m.token_id " +
		"LEFT JOIN transaction t ON t.id=m.tx_id " +
		"LEFT JOIN address wa ON m.owner_aid=wa.address_id " +
		"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id " +
		"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id " +
		"LEFT JOIN cg_nft_staked_rwalk a ON a.action_id=st.stake_action_id " +
		"LEFT JOIN address sa ON a.staker_aid = sa.address_id " +
		"ORDER BY m.token_id"
	scan := func(rows pgx.Rows, rec *cgmodel.CGStakedTokenRWalkRec) error {
		return rows.Scan(
			&rec.StakeEvtLogId,
			&rec.StakeBlockNum,
			&rec.StakeActionId,
			&rec.StakeTimeStamp,
			store.TimeText(&rec.StakeDateTime),
			&rec.UserAddr,
			&rec.UserAid,
			&rec.StakedTokenId,
		)
	}
	return queryList(ctx, r, "staked tokens rwalk global", 16, query, scan)
}

// ActionIDsForDepositWithClaimInfo returns the stake actions of userAid that
// participate in one staking deposit, with per-action claim status.
func (r *Repo) ActionIDsForDepositWithClaimInfo(ctx context.Context, depositID, userAid int64) ([]cgmodel.CGActionIdsForDepositWithClaimInfo, error) {
	query := "SELECT " +
		"d.id," +
		"a.action_id, " +
		"a.token_id, " +
		"d.deposit_id, " +
		"d.block_num, " +
		"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT," +
		"d.time_stamp," +
		"tx.id," +
		"tx.tx_hash," +
		"r.reward," +
		"r.reward/1e18, " +
		"r.collected " +
		"FROM cg_nft_staked_cst a " +
		"JOIN cg_st_reward r ON (a.action_id=r.action_id) AND (r.deposit_id=$1) AND (r.staker_aid=a.staker_aid) AND (r.staker_aid=$2) " +
		"JOIN cg_staking_eth_deposit d ON r.deposit_id=d.deposit_id " +
		"LEFT JOIN cg_nft_unstaked_cst u ON a.action_id=u.action_id " +
		"LEFT JOIN transaction tx ON tx.id=d.tx_id " +
		"WHERE " +
		"(a.staker_aid = $2) AND (" +
		"(" +
		"(a.action_id < $1) AND (u.action_counter IS NULL)" +
		") OR " +
		"(" +
		"(a.action_id<$1 AND " +
		"(u.action_counter IS NOT NULL) AND " +
		"($1<=u.action_counter) " +
		")" +
		")" +
		") " +
		"ORDER BY d.evtlog_id DESC "
	scan := func(rows pgx.Rows, rec *cgmodel.CGActionIdsForDepositWithClaimInfo) error {
		var nullDepositID sql.NullInt64
		var nullRwdBlockNum, nullRwdTimestamp, nullRwdTxID sql.NullInt64
		var nullRwdTxHash, nullReward sql.NullString
		var nullRewardEth sql.NullFloat64
		err := rows.Scan(
			&rec.RecordId,
			&rec.StakeActionId,
			&rec.TokenId,
			&nullDepositID,
			&nullRwdBlockNum,
			&nullRwdTimestamp,
			store.NullTimeText(&rec.ClaimDateTime),
			&nullRwdTxID,
			&nullRwdTxHash,
			&nullReward,
			&nullRewardEth,
			&rec.Claimed,
		)
		if err != nil {
			return err
		}
		rec.DepositId = depositID
		rec.UserAid = userAid
		if nullRwdBlockNum.Valid {
			rec.ClaimBlockNum = nullRwdBlockNum.Int64
		}
		if nullRwdTimestamp.Valid {
			rec.ClaimTimeStamp = nullRwdTimestamp.Int64
		}
		if nullRwdTxID.Valid {
			rec.ClaimTxId = nullRwdTxID.Int64
		}
		if nullRwdTxHash.Valid {
			rec.ClaimTxHash = nullRwdTxHash.String
		}
		if nullReward.Valid {
			rec.ClaimRewardAmount = nullReward.String
		}
		if nullRewardEth.Valid {
			rec.ClaimRewardAmountEth = nullRewardEth.Float64
		}
		return nil
	}
	return queryList(ctx, r, "action ids for deposit with claim info", 16, query, scan, depositID, userAid)
}

// GlobalStakingRewards returns every staking ETH deposit with collected vs
// pending totals across all stakers, newest first.
func (r *Repo) GlobalStakingRewards(ctx context.Context) ([]cgmodel.CGStakingRewardGlobal, error) {
	query := "WITH rwd AS (" +
		"SELECT " +
		"SUM(CASE WHEN collected='T' THEN reward ELSE 0 END) AS collected_reward," +
		"SUM(CASE WHEN collected='T' THEN reward/1e18 ELSE 0 END) AS collected_reward_eth," +
		"COUNT(token_id) AS count_total," +
		"SUM(CASE WHEN collected='F' THEN 1 ELSE 0 END) AS count_not_collected, " +
		"round_num " +
		"FROM cg_st_reward " +
		"GROUP BY round_num " +
		") " +
		"SELECT " +
		"d.id," +
		"d.evtlog_id," +
		"d.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT," +
		"d.time_stamp," +
		"EXTRACT(EPOCH FROM d.deposit_time)::BIGINT," +
		"d.deposit_time," +
		"d.accumulated_nfts," +
		"d.deposit_amount," +
		"d.deposit_amount/1e18," +
		"d.round_num, " +
		"COALESCE(rwd.collected_reward,0)," +
		"COALESCE(rwd.collected_reward_eth,0), " +
		"COALESCE((d.deposit_amount-rwd.collected_reward),0)," +
		"COALESCE((d.deposit_amount-rwd.collected_reward)/1e18,0)," +
		"COALESCE(rwd.count_total,0)," +
		"COALESCE(rwd.count_not_collected,0) " +
		"FROM cg_staking_eth_deposit d " +
		"INNER JOIN transaction tx ON tx.id=d.tx_id " +
		"LEFT JOIN rwd ON (rwd.round_num=d.round_num) " +
		"ORDER BY d.id DESC "
	scan := func(rows pgx.Rows, rec *cgmodel.CGStakingRewardGlobal) error {
		var countTotal, countNotCollected int64
		err := rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.DepositTimeStamp,
			store.TimeText(&rec.DepositDate),
			&rec.NumStakedNFTs,
			&rec.TotalDepositAmount,
			&rec.TotalDepositAmountEth,
			&rec.RoundNum,
			&rec.AlreadyCollected,
			&rec.AlreadyCollectedEth,
			&rec.PendingToCollect,
			&rec.PendingToCollectEth,
			&countTotal,
			&countNotCollected,
		)
		if err != nil {
			return err
		}
		rec.FullyClaimed = countNotCollected == 0
		return nil
	}
	return queryList(ctx, r, "global staking rewards", 32, query, scan)
}

// StakingCstRewardsByRound returns each staker's share of the staking ETH
// deposit made in one round.
func (r *Repo) StakingCstRewardsByRound(ctx context.Context, roundNum int64) ([]cgmodel.CGEthDepositAsReward, error) {
	query := "SELECT " +
		"d.id," +
		"d.evtlog_id," +
		"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS tstmp, " +
		"d.time_stamp AS date_time, " +
		"d.block_num," +
		"d.tx_id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM d.deposit_time)::BIGINT, " +
		"d.deposit_time, " +
		"d.accumulated_nfts," +
		"d.deposit_amount," +
		"d.deposit_amount/1e18 AS deposit_amount_eth," +
		"d.amount_per_token," +
		"d.amount_per_token/1e18 AS amount_per_token_eth, " +
		"sd.staker_aid, " +
		"sa.addr," +
		"sd.tokens_staked," +
		"sd.amount_deposited," +
		"sd.amount_deposited/1e18 AS amount_deposited_eth," +
		"(sd.amount_deposited - sd.amount_to_claim)," +
		"(sd.amount_deposited - sd.amount_to_claim)/1e18 AS amount_collected_eth," +
		"sd.amount_to_claim, " +
		"sd.amount_to_claim/1e18 AS amount_to_claim_eth " +
		"FROM cg_staker_deposit sd " +
		"LEFT JOIN cg_staking_eth_deposit d ON sd.deposit_id=d.deposit_id " +
		"LEFT JOIN address sa ON sd.staker_aid = sa.address_id " +
		"LEFT JOIN transaction t ON t.id=d.tx_id " +
		"WHERE d.round_num=$1 "
	scan := func(rows pgx.Rows, rec *cgmodel.CGEthDepositAsReward) error {
		return rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.DepositTimeStamp,
			store.TimeText(&rec.DepositDate),
			&rec.NumStakedNFTsTotal,
			&rec.Amount,
			&rec.AmountEth,
			&rec.AmountPerToken,
			&rec.AmountPerTokenEth,
			&rec.StakerAid,
			&rec.StakerAddr,
			&rec.StakerNumStakedNFTs,
			&rec.StakerAmount,
			&rec.StakerAmountEth,
			&rec.AmountCollected,
			&rec.AmountCollectedEth,
			&rec.AmountPendingToClaim,
			&rec.AmountPendingToClaimEth,
		)
	}
	return queryList(ctx, r, "staking cst rewards by round", 32, query, scan, roundNum)
}

// GlobalStakingCstHistory returns the interleaved CST stake/unstake event
// history (newest first) with a running staked-NFT accumulator.
func (r *Repo) GlobalStakingCstHistory(ctx context.Context, offset, limit int) ([]cgmodel.CGStakingCSTHistoryRec, error) {
	query := "(" +
		"SELECT " +
		"0 AS action_type," +
		"s.id," +
		"s.evtlog_id," +
		"s.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT," +
		"s.time_stamp," +
		"-1 AS usts," +
		"TO_TIMESTAMP(0) AS unstake_time," +
		"s.action_id," +
		"s.token_id," +
		"s.round_num," +
		"s.num_staked_nfts, " +
		"s.staker_aid, " +
		"sa.addr staker_addr " +
		"FROM cg_nft_staked_cst s " +
		"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
		"LEFT JOIN address sa ON s.staker_aid=sa.address_id " +
		") UNION ALL (" +
		"SELECT " +
		"1 AS action_type," +
		"u.id," +
		"u.evtlog_id," +
		"u.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT AS usts," +
		"u.time_stamp," +
		"0 AS usts," +
		"TO_TIMESTAMP(0) AS unstake_time," +
		"u.action_id," +
		"u.token_id," +
		"u.round_num," +
		"u.num_staked_nfts, " +
		"u.staker_aid," +
		"ua.addr staker_addr " +
		"FROM cg_nft_unstaked_cst u " +
		"LEFT JOIN transaction tx ON tx.id=u.tx_id " +
		"LEFT JOIN address ua ON u.staker_aid=ua.address_id " +
		") ORDER BY evtlog_id DESC " +
		"OFFSET $1 LIMIT $2 "
	var accumStakedNFTs int64
	scan := func(rows pgx.Rows, rec *cgmodel.CGStakingCSTHistoryRec) error {
		err := rows.Scan(
			&rec.ActionType,
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.UnstakeTimeStamp,
			store.TimeText(&rec.UnstakeDate),
			&rec.ActionId,
			&rec.TokenId,
			&rec.RoundNum,
			&rec.NumStakedNFTs,
			&rec.StakerAid,
			&rec.StakerAddr,
		)
		if err != nil {
			return err
		}
		accumStakedNFTs += rec.NumStakedNFTs
		rec.AccumNumStakedNFTs = accumStakedNFTs
		return nil
	}
	return queryList(ctx, r, "global staking cst history", 16, query, scan, offset, limit)
}

// GlobalStakingRwalkHistory returns the interleaved RandomWalk stake/unstake
// event history (newest first) with a running staked-NFT accumulator and the
// last indexed block timestamp attached to each row.
func (r *Repo) GlobalStakingRwalkHistory(ctx context.Context, offset, limit int) ([]cgmodel.CGStakingRWalkHistoryRec, error) {
	lastTs, err := r.lastBlockTimestamp(ctx)
	if err != nil {
		return nil, err
	}
	query := "(" +
		"SELECT " +
		"0 AS action_type," +
		"s.id," +
		"s.evtlog_id," +
		"s.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT," +
		"s.time_stamp," +
		"-1 AS usts," +
		"TO_TIMESTAMP(0) AS unstake_time," +
		"s.action_id," +
		"s.token_id," +
		"s.round_num," +
		"s.num_staked_nfts, " +
		"s.staker_aid, " +
		"sa.addr staker_addr " +
		"FROM cg_nft_staked_rwalk s " +
		"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
		"LEFT JOIN address sa ON s.staker_aid=sa.address_id " +
		"ORDER BY s.id DESC " +
		"OFFSET $1 LIMIT $2 " +
		") UNION ALL (" +
		"SELECT " +
		"1 AS action_type," +
		"u.id," +
		"u.evtlog_id," +
		"u.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT AS usts," +
		"u.time_stamp," +
		"0 AS usts," +
		"TO_TIMESTAMP(0) AS unnstake_time," +
		"u.action_id," +
		"u.token_id," +
		"u.round_num," +
		"u.num_staked_nfts, " +
		"u.staker_aid," +
		"ua.addr staker_addr " +
		"FROM cg_nft_unstaked_rwalk u " +
		"LEFT JOIN transaction tx ON tx.id=u.tx_id " +
		"LEFT JOIN address ua ON u.staker_aid=ua.address_id " +
		"LEFT JOIN cg_nft_staked_rwalk s ON u.action_id=s.action_id " +
		"ORDER BY u.id DESC " +
		"OFFSET $1 LIMIT $2 " +
		") order by evtlog_id DESC"
	var accumStakedNFTs int64
	scan := func(rows pgx.Rows, rec *cgmodel.CGStakingRWalkHistoryRec) error {
		err := rows.Scan(
			&rec.ActionType,
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.UnstakeTimeStamp,
			store.TimeText(&rec.UnstakeDate),
			&rec.ActionId,
			&rec.TokenId,
			&rec.RoundNum,
			&rec.NumStakedNFTs,
			&rec.StakerAid,
			&rec.StakerAddr,
		)
		if err != nil {
			return err
		}
		accumStakedNFTs += rec.NumStakedNFTs
		rec.AccumNumStakedNFTs = accumStakedNFTs
		rec.LastBlockTS = lastTs
		rec.UnstakeExpirationDiff = -1
		return nil
	}
	return queryList(ctx, r, "global staking rwalk history", 16, query, scan, offset, limit)
}

// lastBlockTimestamp returns the timestamp of the last indexed block, or 0
// when the chain watermark has no matching block row yet.
func (r *Repo) lastBlockTimestamp(ctx context.Context) (int64, error) {
	query := "SELECT FLOOR(EXTRACT(EPOCH FROM block.ts))::BIGINT AS ts " +
		"FROM block,last_block WHERE last_block.block_num=block.block_num"
	var ts int64
	err := r.q(ctx).QueryRow(ctx, query).Scan(&ts)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, store.WrapError("last block timestamp", err)
	}
	return ts, nil
}

// StakingRwalkMintsGlobal returns the Cosmic Signature tokens minted to
// RandomWalk-staker raffle winners, newest first.
func (r *Repo) StakingRwalkMintsGlobal(ctx context.Context, offset, limit int) ([]cgmodel.CGRaffleNFTWinnerRec, error) {
	query := stakingMintsQuery("(is_rwalk=TRUE) AND (is_staker=TRUE)")
	scan := func(rows pgx.Rows, rec *cgmodel.CGRaffleNFTWinnerRec) error {
		if err := scanStakingMint(rows, rec); err != nil {
			return err
		}
		rec.IsRWalk = true
		rec.IsStaker = true
		return nil
	}
	return queryList(ctx, r, "staking rwalk mints global", 16, query, scan, offset, limit)
}

// StakingCstMintsGlobal returns the Cosmic Signature tokens minted to
// CST-staker raffle winners, newest first.
func (r *Repo) StakingCstMintsGlobal(ctx context.Context, offset, limit int) ([]cgmodel.CGRaffleNFTWinnerRec, error) {
	query := stakingMintsQuery("(is_rwalk=FALSE) AND (is_staker=TRUE)")
	scan := func(rows pgx.Rows, rec *cgmodel.CGRaffleNFTWinnerRec) error {
		if err := scanStakingMint(rows, rec); err != nil {
			return err
		}
		rec.IsRWalk = false
		rec.IsStaker = true
		return nil
	}
	return queryList(ctx, r, "staking cst mints global", 16, query, scan, offset, limit)
}

// stakingMintsQuery is the shared SELECT of the two global staking-mint
// queries; whereClause is a compile-time literal from this file.
func stakingMintsQuery(whereClause string) string {
	return "SELECT " +
		"w.id," +
		"w.evtlog_id," +
		"w.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM w.time_stamp)::BIGINT," +
		"w.time_stamp," +
		"w.token_id," +
		"w.cst_amount," +
		"w.cst_amount/1e18 cst_amount_eth," +
		"w.winner_idx," +
		"w.round_num," +
		"w.winner_aid," +
		"wa.addr " +
		"FROM cg_raffle_nft_prize w " +
		"LEFT JOIN transaction t ON t.id=w.tx_id " +
		"LEFT JOIN address wa ON w.winner_aid=wa.address_id " +
		"WHERE " + whereClause + " " +
		"ORDER BY w.evtlog_id DESC " +
		"OFFSET $1 LIMIT $2"
}

func scanStakingMint(rows pgx.Rows, rec *cgmodel.CGRaffleNFTWinnerRec) error {
	return rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.TokenId,
		&rec.CstAmount,
		&rec.CstAmountEth,
		&rec.WinnerIndex,
		&rec.RoundNum,
		&rec.WinnerAid,
		&rec.WinnerAddr,
	)
}

// StakingCstUserDepositRewards returns userAid's CST staking rewards grouped
// per deposit, each with the stake/unstake actions that earned them and
// claimable/claimed accumulators.
func (r *Repo) StakingCstUserDepositRewards(ctx context.Context, userAid int64) ([]cgmodel.CGCombinedDepositRewardRec, error) {
	const op = "staking cst user deposit rewards"
	// Sorted ASC because the per-deposit accumulators are computed in Go
	// between consecutive rows.
	query := "SELECT " +
		"sa_id,sa_evtlog_id,sa_block_num,sa_tx_id,sa_tx_hash,sa_time_stamp,sa_date_time, " +
		"sa_action_id,sa_token_id,sa_num_staked_nfts," +
		"ua_id,ua_evtlog_id,ua_block_num,ua_tx_id,ua_tx_hash,ua_time_stamp,ua_date_time, " +
		"ua_action_id,ua_token_id,ua_num_staked_nfts,ua_reward,ua_reward_eth," +
		"d.id,d.evtlog_id,d.block_num,tx.id,tx.tx_hash,EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,d.time_stamp," +
		"d.deposit_id,d.round_num,d.num_staked_nfts,d.deposit_amount,d.deposit_amount/1e18,amount_per_token/1e18," +
		"str.reward," +
		"str.reward/1e18," +
		"str.collected " +
		"FROM cg_st_reward str " +
		"INNER JOIN cg_staking_eth_deposit d ON str.deposit_id=d.deposit_id " +
		"INNER JOIN transaction tx ON tx.id=d.tx_id " +
		"INNER JOIN LATERAL (" +
		"SELECT " +
		"sa.id sa_id,sa.evtlog_id sa_evtlog_id,sa.block_num sa_block_num,satx.id sa_tx_id,satx.tx_hash sa_tx_hash,EXTRACT(EPOCH FROM sa.time_stamp)::BIGINT sa_time_stamp,sa.time_stamp sa_date_time, " +
		"sa.action_id sa_action_id,sa.token_id sa_token_id,sa.num_staked_nfts sa_num_staked_nfts," +
		"ua.id ua_id,ua.evtlog_id ua_evtlog_id,ua.block_num ua_block_num,uatx.id ua_tx_id,uatx.tx_hash ua_tx_hash,EXTRACT(EPOCH FROM ua.time_stamp)::BIGINT ua_time_stamp,ua.time_stamp ua_date_time, " +
		"ua.action_id ua_action_id,ua.token_id ua_token_id,ua.num_staked_nfts ua_num_staked_nfts,ua.reward ua_reward,ua.reward/1e18 ua_reward_eth " +
		"FROM cg_nft_staked_cst sa " +
		"LEFT JOIN cg_nft_unstaked_cst ua ON ua.action_id=sa.action_id " +
		"LEFT JOIN transaction satx ON satx.id=sa.tx_id " +
		"LEFT JOIN transaction uatx ON uatx.id=ua.tx_id " +
		") a ON a.sa_action_id=str.action_id " +
		"WHERE str.staker_aid=$1 " +
		"ORDER BY d.id ASC,sa_action_id DESC "

	rows, err := r.q(ctx).Query(ctx, query, userAid)
	if err != nil {
		return nil, store.WrapError(op, err)
	}
	defer rows.Close()

	var rec cgmodel.CGCombinedDepositRewardRec
	curDepositID := int64(-1)
	var yourTokensStaked, numTokensCollected int64
	var yourClaimableAmount, yourClaimedAmount float64
	fullyClaimed := true
	records := make([]cgmodel.CGCombinedDepositRewardRec, 0, 16)
	actions := make([]cgmodel.CGNftStakeUnstakeCombined, 0, 16)

	flush := func() {
		rec.Actions = actions
		rec.YourTokensStaked = yourTokensStaked
		rec.YourClaimableAmountEth = yourClaimableAmount
		rec.FullyClaimed = fullyClaimed
		rec.NumTokensCollected = numTokensCollected
		rec.ClaimedAmountEth = yourClaimedAmount
		records = append(records, rec)
	}

	for rows.Next() {
		var recRow cgmodel.CGNftStakeUnstakeCombined
		var nullRecordID, nullActionID, nullEvtlogID, nullBlockNum, nullTxID, nullTokenID, nullNumStakedNFTs, nullTimeStamp sql.NullInt64
		var nullTxHash, nullReward sql.NullString
		var nullRewardEth sql.NullFloat64
		var recordID, evtlogID, blockNum, txID, depositID, timeStamp, depRound, numStakedNFTs int64
		var txHash, dateTime, depositAmount string
		var depAmountEth, amountPerTokenEth float64
		err := rows.Scan(
			&recRow.Stake.RecordId, &recRow.Stake.Tx.EvtLogId, &recRow.Stake.Tx.BlockNum, &recRow.Stake.Tx.TxId, &recRow.Stake.Tx.TxHash, &recRow.Stake.Tx.TimeStamp, store.TimeText(&recRow.Stake.Tx.DateTime),
			&recRow.Stake.ActionId, &recRow.Stake.TokenId, &recRow.Stake.NumStakedNFTs,
			&nullRecordID, &nullEvtlogID, &nullBlockNum, &nullTxID, &nullTxHash, &nullTimeStamp, store.NullTimeText(&recRow.Unstake.Tx.DateTime),
			&nullActionID, &nullTokenID, &nullNumStakedNFTs, &nullReward, &nullRewardEth,
			&recordID, &evtlogID, &blockNum, &txID, &txHash, &timeStamp, store.TimeText(&dateTime),
			&depositID, &depRound, &numStakedNFTs, &depositAmount, &depAmountEth, &amountPerTokenEth,
			&recRow.Reward, &recRow.RewardEth,
			&recRow.Claimed,
		)
		if err != nil {
			return nil, store.WrapError(op, err)
		}
		if nullRecordID.Valid {
			recRow.Unstake.RecordId = nullRecordID.Int64
			recRow.Unstake.StakerAid = userAid
		}
		if nullEvtlogID.Valid {
			recRow.Unstake.Tx.EvtLogId = nullEvtlogID.Int64
		}
		if nullBlockNum.Valid {
			recRow.Unstake.Tx.BlockNum = nullBlockNum.Int64
		}
		if nullTxID.Valid {
			recRow.Unstake.Tx.TxId = nullTxID.Int64
		}
		if nullTxHash.Valid {
			recRow.Unstake.Tx.TxHash = nullTxHash.String
		}
		if nullTimeStamp.Valid {
			recRow.Unstake.Tx.TimeStamp = nullTimeStamp.Int64
		}
		if nullActionID.Valid {
			recRow.Unstake.ActionId = nullActionID.Int64
		}
		if nullTokenID.Valid {
			recRow.Unstake.TokenId = nullTokenID.Int64
		}
		if nullNumStakedNFTs.Valid {
			recRow.Unstake.NumStakedNFTs = nullNumStakedNFTs.Int64
		}
		if nullReward.Valid {
			recRow.Unstake.RewardAmount = nullReward.String
		}
		if nullRewardEth.Valid {
			recRow.Unstake.RewardAmountEth = nullRewardEth.Float64
		}
		if depositID != curDepositID {
			if curDepositID != -1 {
				flush()
				fullyClaimed = true
				yourTokensStaked = 0
				yourClaimableAmount = 0
				yourClaimedAmount = 0
				numTokensCollected = 0
				actions = make([]cgmodel.CGNftStakeUnstakeCombined, 0, 16)
			}
			rec.RecordId = recordID
			rec.Tx.EvtLogId = evtlogID
			rec.Tx.BlockNum = blockNum
			rec.Tx.TxId = txID
			rec.Tx.TxHash = txHash
			rec.Tx.TimeStamp = timeStamp
			rec.Tx.DateTime = dateTime
			rec.DepositId = depositID
			rec.DepositRoundNum = depRound
			rec.NumStakedNFTs = numStakedNFTs
			rec.DepositAmount = depositAmount
			rec.DepositAmountEth = depAmountEth
			rec.AmountPerTokenEth = amountPerTokenEth
			curDepositID = depositID
		}
		yourTokensStaked++
		if !recRow.Claimed {
			yourClaimableAmount += recRow.RewardEth
			fullyClaimed = false
		} else {
			numTokensCollected++
			yourClaimedAmount += recRow.RewardEth
		}
		actions = append(actions, recRow)
	}
	if err := rows.Err(); err != nil {
		return nil, store.WrapError(op, err)
	}
	if yourTokensStaked > 0 {
		flush()
	}
	return records, nil
}

// StakingCstUserTokenRewards returns userAid's collected vs pending CST
// staking rewards summed per token.
func (r *Repo) StakingCstUserTokenRewards(ctx context.Context, userAid int64) ([]cgmodel.CGStakingCstRewardPerTokenRec, error) {
	query := "WITH rwd AS (" +
		"SELECT " +
		"token_id, " +
		"SUM(" +
		"CASE " +
		"WHEN collected='T' THEN reward " +
		"ELSE 0 " +
		"END " +
		")/1e18 AS reward_collected, " +
		"SUM(" +
		"CASE " +
		"WHEN collected='F' THEN reward " +
		"ELSE 0 " +
		"END" +
		")/1e18 AS reward_to_collect " +
		"FROM cg_st_reward " +
		"WHERE staker_aid=$1 " +
		"GROUP BY token_id " +
		") " +
		"SELECT " +
		"rwd.token_id," +
		"rwd.reward_collected," +
		"rwd.reward_to_collect " +
		"FROM rwd " +
		"ORDER BY token_id "
	scan := func(rows pgx.Rows, rec *cgmodel.CGStakingCstRewardPerTokenRec) error {
		err := rows.Scan(
			&rec.TokenId,
			&rec.RewardCollectedEth,
			&rec.RewardToCollectEth,
		)
		if err != nil {
			return err
		}
		rec.UserAid = userAid
		return nil
	}
	return queryList(ctx, r, "staking cst user token rewards", 16, query, scan, userAid)
}

// StakingCstUserTokenRewardDetails returns the per-deposit reward rows of
// one token staked by userAid, with the stake/unstake actions attached.
func (r *Repo) StakingCstUserTokenRewardDetails(ctx context.Context, userAid, tokenID int64) ([]cgmodel.CGNftStakeUnstakeCombined, error) {
	query := "SELECT " +
		"rwd.reward," +
		"rwd.reward/1e18, " +
		"rwd.collected, " +
		"rwd.round_num," +
		"rwd.deposit_id," +
		// stake action fields
		"sa_id,sa_evtlog_id,sa_block_num,sa_tx_id,sa_tx_hash,sa_time_stamp,sa_date_time, " +
		"sa_action_id,sa_num_staked_nfts," +
		// unstake action fields
		"ua_id,ua_evtlog_id,ua_block_num,ua_tx_id,ua_tx_hash,ua_time_stamp,ua_date_time, " +
		"ua_action_id,ua_num_staked_nfts,ua_reward,ua_reward_eth, " +
		"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,d.time_stamp " +
		"FROM cg_st_reward rwd " +
		"INNER JOIN cg_staking_eth_deposit d ON rwd.deposit_id=d.deposit_id " +
		"INNER JOIN LATERAL (" +
		"SELECT " +
		"sa.id sa_id,sa.evtlog_id sa_evtlog_id,sa.block_num sa_block_num,satx.id sa_tx_id,satx.tx_hash sa_tx_hash,EXTRACT(EPOCH FROM sa.time_stamp)::BIGINT sa_time_stamp,sa.time_stamp sa_date_time, " +
		"sa.action_id sa_action_id,sa.token_id sa_token_id,sa.num_staked_nfts sa_num_staked_nfts," +
		"ua.id ua_id,ua.evtlog_id ua_evtlog_id,ua.block_num ua_block_num,uatx.id ua_tx_id,uatx.tx_hash ua_tx_hash,EXTRACT(EPOCH FROM ua.time_stamp)::BIGINT ua_time_stamp,ua.time_stamp ua_date_time, " +
		"ua.action_id ua_action_id,ua.token_id ua_token_id,ua.num_staked_nfts ua_num_staked_nfts,ua.reward ua_reward,ua.reward/1e18 ua_reward_eth " +
		"FROM cg_nft_staked_cst sa " +
		"LEFT JOIN cg_nft_unstaked_cst ua ON ua.action_id=sa.action_id " +
		"LEFT JOIN transaction satx ON satx.id=sa.tx_id " +
		"LEFT JOIN transaction uatx ON uatx.id=ua.tx_id " +
		") a ON a.sa_action_id=rwd.action_id " +
		"WHERE rwd.staker_aid=$1 AND rwd.token_id=$2 " +
		"ORDER BY rwd.deposit_id"
	scan := func(rows pgx.Rows, rec *cgmodel.CGNftStakeUnstakeCombined) error {
		var nullRecID, nullEvtlogID, nullBlockNum, nullTxID, nullTimestamp, nullActionID, nullStakedNFTs sql.NullInt64
		var nullTxHash, nullReward sql.NullString
		var nullRewardEth sql.NullFloat64
		err := rows.Scan(
			&rec.Reward,
			&rec.RewardEth,
			&rec.Claimed,
			&rec.RoundNum,
			&rec.DepositId,
			// stake action fields
			&rec.Stake.RecordId,
			&rec.Stake.Tx.EvtLogId,
			&rec.Stake.Tx.BlockNum,
			&rec.Stake.Tx.TxId,
			&rec.Stake.Tx.TxHash,
			&rec.Stake.Tx.TimeStamp,
			store.TimeText(&rec.Stake.Tx.DateTime),
			&rec.Stake.ActionId,
			&rec.Stake.NumStakedNFTs,
			// unstake action fields
			&nullRecID,
			&nullEvtlogID,
			&nullBlockNum,
			&nullTxID,
			&nullTxHash,
			&nullTimestamp,
			store.NullTimeText(&rec.Unstake.Tx.DateTime),
			&nullActionID,
			&nullStakedNFTs,
			&nullReward,
			&nullRewardEth,
			&rec.DepositTimeStamp,
			store.TimeText(&rec.DepositDateTime),
		)
		if err != nil {
			return err
		}
		if nullRecID.Valid {
			rec.Unstake.RecordId = nullRecID.Int64
		}
		if nullEvtlogID.Valid {
			rec.Unstake.Tx.EvtLogId = nullEvtlogID.Int64
		}
		if nullBlockNum.Valid {
			rec.Unstake.Tx.BlockNum = nullBlockNum.Int64
		}
		if nullTxID.Valid {
			rec.Unstake.Tx.TxId = nullTxID.Int64
		}
		if nullTxHash.Valid {
			rec.Unstake.Tx.TxHash = nullTxHash.String
		}
		if nullTimestamp.Valid {
			rec.Unstake.Tx.TimeStamp = nullTimestamp.Int64
		}
		if nullActionID.Valid {
			rec.Unstake.ActionId = nullActionID.Int64
		}
		if nullStakedNFTs.Valid {
			rec.Unstake.NumStakedNFTs = nullStakedNFTs.Int64
		}
		if nullReward.Valid {
			rec.Unstake.RewardAmount = nullReward.String
		}
		if nullRewardEth.Valid {
			rec.Unstake.RewardAmountEth = nullRewardEth.Float64
		}
		rec.Unstake.StakerAid = userAid
		rec.Stake.StakerAid = userAid
		return nil
	}
	return queryList(ctx, r, "staking cst user token reward details", 16, query, scan, userAid, tokenID)
}
