package cosmicgame

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// UserStakingActionKind labels one row of an interleaved stake/unstake
// history.
type UserStakingActionKind string

// The two staking action kinds emitted by the action history queries.
const (
	UserStakingActionStake   UserStakingActionKind = "stake"
	UserStakingActionUnstake UserStakingActionKind = "unstake"
)

// UserStakingActionRecord is one stake or unstake event of one wallet.
// RewardWei is populated exactly on Cosmic Signature unstake rows (the
// unstake transaction collects that token's accumulated rewards); it stays
// empty on stake rows and on every RandomWalk row.
type UserStakingActionRecord struct {
	Tx              cgmodel.Transaction
	Kind            UserStakingActionKind
	StakerAid       int64
	ActionID        int64
	TokenID         int64
	TotalStakedNfts int64
	RewardWei       string
}

// userStakingActionsPageSQL merges one wallet's stake and unstake events
// newest first. Both sides are bounded before the merge so the outer sort
// never sees more than two pages. The table names and the unstake reward
// expression are compile-time literals from this file; only the cursor
// placeholders vary per call.
func userStakingActionsPageSQL(stakeTable, unstakeTable, unstakeRewardExpr string, after bool) string {
	filter := "WHERE e.staker_aid = $1"
	limitPlaceholder := "$2"
	if after {
		filter += " AND e.evtlog_id < $2"
		limitPlaceholder = "$3"
	}
	return fmt.Sprintf(`SELECT
			action_kind,
			evtlog_id,
			block_num,
			tx_id,
			tx_hash,
			ts,
			date_time,
			staker_aid,
			action_id,
			token_id,
			num_staked_nfts,
			reward_wei
		FROM (
			(SELECT
				'stake'::TEXT AS action_kind,
				e.evtlog_id,
				e.block_num,
				t.id AS tx_id,
				t.tx_hash,
				EXTRACT(EPOCH FROM e.time_stamp)::BIGINT AS ts,
				e.time_stamp AS date_time,
				e.staker_aid,
				e.action_id,
				e.token_id,
				e.num_staked_nfts,
				NULL::TEXT AS reward_wei
			FROM %s e
				LEFT JOIN transaction t ON t.id=e.tx_id
			%s
			ORDER BY e.evtlog_id DESC
			LIMIT %s)
			UNION ALL
			(SELECT
				'unstake'::TEXT AS action_kind,
				e.evtlog_id,
				e.block_num,
				t.id AS tx_id,
				t.tx_hash,
				EXTRACT(EPOCH FROM e.time_stamp)::BIGINT AS ts,
				e.time_stamp AS date_time,
				e.staker_aid,
				e.action_id,
				e.token_id,
				e.num_staked_nfts,
				%s AS reward_wei
			FROM %s e
				LEFT JOIN transaction t ON t.id=e.tx_id
			%s
			ORDER BY e.evtlog_id DESC
			LIMIT %s)
		) actions
		ORDER BY evtlog_id DESC
		LIMIT %s`,
		stakeTable, filter, limitPlaceholder,
		unstakeRewardExpr, unstakeTable, filter, limitPlaceholder,
		limitPlaceholder)
}

func scanUserStakingAction(rows pgx.Rows, rec *UserStakingActionRecord) error {
	var reward sql.NullString
	if err := rows.Scan(
		&rec.Kind,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.StakerAid,
		&rec.ActionID,
		&rec.TokenID,
		&rec.TotalStakedNfts,
		&reward,
	); err != nil {
		return err
	}
	if reward.Valid {
		rec.RewardWei = reward.String
	}
	return nil
}

func (r *Repo) userStakingActionsPage(
	ctx context.Context,
	op string,
	stakeTable, unstakeTable, unstakeRewardExpr string,
	userAid int64,
	after *UserEventPageCursor,
	limit int,
) (records []UserStakingActionRecord, hasMore bool, err error) {
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	args := []any{userAid, limit + 1}
	if after != nil {
		args = []any{userAid, after.EventLogID, limit + 1}
	}
	records, err = queryList(
		ctx,
		r,
		op,
		limit+1,
		userStakingActionsPageSQL(stakeTable, unstakeTable, unstakeRewardExpr, after != nil),
		scanUserStakingAction,
		args...,
	)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserCstStakingActionsPage returns at most limit Cosmic Signature stake and
// unstake events of userAid, newest first by immutable event-log ID.
func (r *Repo) UserCstStakingActionsPage(
	ctx context.Context,
	userAid int64,
	after *UserEventPageCursor,
	limit int,
) ([]UserStakingActionRecord, bool, error) {
	return r.userStakingActionsPage(
		ctx,
		"user cst staking actions page",
		"cg_nft_staked_cst", "cg_nft_unstaked_cst", "e.reward::TEXT",
		userAid, after, limit,
	)
}

// UserRwalkStakingActionsPage returns at most limit RandomWalk stake and
// unstake events of userAid, newest first. RandomWalk unstake rows carry no
// reward columns, so RewardWei stays empty.
func (r *Repo) UserRwalkStakingActionsPage(
	ctx context.Context,
	userAid int64,
	after *UserEventPageCursor,
	limit int,
) ([]UserStakingActionRecord, bool, error) {
	return r.userStakingActionsPage(
		ctx,
		"user rwalk staking actions page",
		"cg_nft_staked_rwalk", "cg_nft_unstaked_rwalk", "NULL::TEXT",
		userAid, after, limit,
	)
}

// UserStakingTokenPageCursor identifies the last token returned by an
// ascending token-keyed staking page.
type UserStakingTokenPageCursor struct {
	TokenID int64
}

func (c *UserStakingTokenPageCursor) valid() bool {
	return c == nil || c.TokenID >= 0
}

// UserStakedCstTokenRecord is one currently staked Cosmic Signature token
// with the stake action that locked it and the token's mint provenance.
type UserStakedCstTokenRecord struct {
	StakeTx   cgmodel.Transaction
	StakerAid int64
	ActionID  int64
	TokenID   int64
	MintRound int64
	Seed      string
	TokenName string
}

// UserStakedCstTokensPage returns at most limit currently staked Cosmic
// Signature tokens of userAid in ascending token order. The collection is
// live membership: rows disappear when tokens unstake.
func (r *Repo) UserStakedCstTokensPage(
	ctx context.Context,
	userAid int64,
	after *UserStakingTokenPageCursor,
	limit int,
) (records []UserStakedCstTokenRecord, hasMore bool, err error) {
	const op = "user staked cst tokens page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
			a.evtlog_id,
			a.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM a.time_stamp)::BIGINT,
			a.time_stamp,
			st.staker_aid,
			a.action_id,
			st.token_id,
			m.round_num,
			m.seed,
			m.token_name
		FROM cg_staked_token_cst st
			INNER JOIN cg_nft_staked_cst a ON a.action_id=st.stake_action_id
			INNER JOIN cg_mint_event m ON m.token_id=st.token_id
			LEFT JOIN transaction t ON t.id=a.tx_id
		WHERE st.staker_aid=$1`
	args := []any{userAid}
	if after != nil {
		query += fmt.Sprintf(" AND st.token_id > $%d", len(args)+1)
		args = append(args, after.TokenID)
	}
	query += fmt.Sprintf(`
		ORDER BY st.token_id
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	scan := func(rows pgx.Rows, rec *UserStakedCstTokenRecord) error {
		return rows.Scan(
			&rec.StakeTx.EvtLogId,
			&rec.StakeTx.BlockNum,
			&rec.StakeTx.TxId,
			&rec.StakeTx.TxHash,
			&rec.StakeTx.TimeStamp,
			store.TimeText(&rec.StakeTx.DateTime),
			&rec.StakerAid,
			&rec.ActionID,
			&rec.TokenID,
			&rec.MintRound,
			&rec.Seed,
			&rec.TokenName,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserStakedRwalkTokenRecord is one currently staked RandomWalk token with
// the stake action that locked it.
type UserStakedRwalkTokenRecord struct {
	StakeTx   cgmodel.Transaction
	StakerAid int64
	ActionID  int64
	TokenID   int64
}

// UserStakedRwalkTokensPage returns at most limit currently staked
// RandomWalk tokens of userAid in ascending token order.
func (r *Repo) UserStakedRwalkTokensPage(
	ctx context.Context,
	userAid int64,
	after *UserStakingTokenPageCursor,
	limit int,
) (records []UserStakedRwalkTokenRecord, hasMore bool, err error) {
	const op = "user staked rwalk tokens page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
			a.evtlog_id,
			a.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM a.time_stamp)::BIGINT,
			a.time_stamp,
			st.staker_aid,
			a.action_id,
			st.token_id
		FROM cg_staked_token_rwalk st
			INNER JOIN cg_nft_staked_rwalk a ON a.action_id=st.stake_action_id
			LEFT JOIN transaction t ON t.id=a.tx_id
		WHERE st.staker_aid=$1`
	args := []any{userAid}
	if after != nil {
		query += fmt.Sprintf(" AND st.token_id > $%d", len(args)+1)
		args = append(args, after.TokenID)
	}
	query += fmt.Sprintf(`
		ORDER BY st.token_id
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	scan := func(rows pgx.Rows, rec *UserStakedRwalkTokenRecord) error {
		return rows.Scan(
			&rec.StakeTx.EvtLogId,
			&rec.StakeTx.BlockNum,
			&rec.StakeTx.TxId,
			&rec.StakeTx.TxHash,
			&rec.StakeTx.TimeStamp,
			store.TimeText(&rec.StakeTx.DateTime),
			&rec.StakerAid,
			&rec.ActionID,
			&rec.TokenID,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserStakingDepositPageCursor identifies the last deposit returned by the
// newest-first deposit ledger.
type UserStakingDepositPageCursor struct {
	DepositID int64
}

func (c *UserStakingDepositPageCursor) valid() bool {
	return c == nil || c.DepositID >= 0
}

// UserStakingDepositRecord is one wallet's share of one staking-wallet ETH
// deposit with exact accumulators from the smallest reward units. The
// AmountDepositedWei/PendingWei pair comes from the cg_staker_deposit
// accumulator, the Claimed*/Pending* sums and counts from cg_st_reward; the
// API mapper cross-checks both sources.
type UserStakingDepositRecord struct {
	Tx                 cgmodel.Transaction
	StakerAid          int64
	DepositID          int64
	RoundNum           int64
	TotalDepositWei    string
	TotalStakedNfts    int64
	AmountPerTokenWei  string
	StakedNftCount     int64
	AmountDepositedWei string
	AmountToClaimWei   string
	ClaimedRewardWei   string
	PendingRewardWei   string
	ClaimedNftCount    int64
	PendingNftCount    int64
}

// userStakingDepositsSelect aggregates the wallet's smallest reward units
// per deposit next to the cg_staker_deposit accumulator row.
const userStakingDepositsSelect = `WITH rewards AS (
			SELECT
				deposit_id,
				COUNT(*) FILTER (WHERE collected) AS claimed_count,
				COUNT(*) FILTER (WHERE NOT collected) AS pending_count,
				COALESCE(SUM(reward) FILTER (WHERE collected), 0) AS claimed_sum,
				COALESCE(SUM(reward) FILTER (WHERE NOT collected), 0) AS pending_sum
			FROM cg_st_reward
			WHERE staker_aid=$1
			GROUP BY deposit_id
		)
		SELECT
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			sd.staker_aid,
			d.deposit_id,
			d.round_num,
			d.deposit_amount::TEXT,
			d.accumulated_nfts,
			d.amount_per_token::TEXT,
			sd.tokens_staked,
			sd.amount_deposited::TEXT,
			sd.amount_to_claim::TEXT,
			COALESCE(r.claimed_sum, 0)::TEXT,
			COALESCE(r.pending_sum, 0)::TEXT,
			COALESCE(r.claimed_count, 0),
			COALESCE(r.pending_count, 0)
		FROM cg_staker_deposit sd
			INNER JOIN cg_staking_eth_deposit d ON d.deposit_id=sd.deposit_id
			INNER JOIN transaction t ON t.id=d.tx_id
			LEFT JOIN rewards r ON r.deposit_id=sd.deposit_id`

func scanUserStakingDeposit(rows pgx.Rows, rec *UserStakingDepositRecord) error {
	return rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.StakerAid,
		&rec.DepositID,
		&rec.RoundNum,
		&rec.TotalDepositWei,
		&rec.TotalStakedNfts,
		&rec.AmountPerTokenWei,
		&rec.StakedNftCount,
		&rec.AmountDepositedWei,
		&rec.AmountToClaimWei,
		&rec.ClaimedRewardWei,
		&rec.PendingRewardWei,
		&rec.ClaimedNftCount,
		&rec.PendingNftCount,
	)
}

// UserStakingDepositsPage returns at most limit staking deposits userAid
// has a share in, newest deposit first. claimed narrows the ledger to fully
// claimed (true) or partially pending (false) deposits; nil returns both.
func (r *Repo) UserStakingDepositsPage(
	ctx context.Context,
	userAid int64,
	claimed *bool,
	after *UserStakingDepositPageCursor,
	limit int,
) (records []UserStakingDepositRecord, hasMore bool, err error) {
	const op = "user staking deposits page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	filter := "WHERE sd.staker_aid=$1"
	args := []any{userAid}
	if claimed != nil {
		if *claimed {
			filter += " AND COALESCE(r.pending_count, 0) = 0"
		} else {
			filter += " AND COALESCE(r.pending_count, 0) > 0"
		}
	}
	if after != nil {
		filter += fmt.Sprintf(" AND sd.deposit_id < $%d", len(args)+1)
		args = append(args, after.DepositID)
	}
	query := fmt.Sprintf(`%s
		%s
		ORDER BY sd.deposit_id DESC
		LIMIT $%d`, userStakingDepositsSelect, filter, len(args)+1)
	args = append(args, limit+1)

	records, err = queryList(ctx, r, op, limit+1, query, scanUserStakingDeposit, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// StakingDepositExists reports whether a staking-wallet ETH deposit with the
// contract-assigned identifier exists.
func (r *Repo) StakingDepositExists(ctx context.Context, depositID int64) (bool, error) {
	const op = "staking deposit exists"
	var exists bool
	err := r.pool().QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM cg_staking_eth_deposit WHERE deposit_id=$1)",
		depositID,
	).Scan(&exists)
	if err != nil {
		return false, store.WrapError(op, err)
	}
	return exists, nil
}

// UserStakingRewardPageCursor identifies the last stake action returned by
// the ascending per-deposit reward page.
type UserStakingRewardPageCursor struct {
	ActionID int64
}

func (c *UserStakingRewardPageCursor) valid() bool {
	return c == nil || c.ActionID >= 0
}

// UserStakingDepositRewardRecord is the smallest reward unit: one wallet's
// reward for one stake action inside one deposit.
type UserStakingDepositRewardRecord struct {
	StakerAid int64
	ActionID  int64
	TokenID   int64
	RewardWei string
	Claimed   bool
}

// UserStakingDepositRewardsPage returns at most limit reward units of
// userAid inside one deposit, in ascending stake action order.
func (r *Repo) UserStakingDepositRewardsPage(
	ctx context.Context,
	userAid int64,
	depositID int64,
	after *UserStakingRewardPageCursor,
	limit int,
) (records []UserStakingDepositRewardRecord, hasMore bool, err error) {
	const op = "user staking deposit rewards page"
	if userAid < 1 || depositID < 0 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id, deposit id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
			r.staker_aid,
			r.action_id,
			r.token_id,
			r.reward::TEXT,
			r.collected
		FROM cg_st_reward r
		WHERE r.staker_aid=$1 AND r.deposit_id=$2`
	args := []any{userAid, depositID}
	if after != nil {
		query += fmt.Sprintf(" AND r.action_id > $%d", len(args)+1)
		args = append(args, after.ActionID)
	}
	query += fmt.Sprintf(`
		ORDER BY r.action_id
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	scan := func(rows pgx.Rows, rec *UserStakingDepositRewardRecord) error {
		return rows.Scan(
			&rec.StakerAid,
			&rec.ActionID,
			&rec.TokenID,
			&rec.RewardWei,
			&rec.Claimed,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserStakingTokenRewardRecord is one wallet's exact staking reward totals
// for one Cosmic Signature token.
type UserStakingTokenRewardRecord struct {
	TokenID      int64
	TotalWei     string
	CollectedWei string
	PendingWei   string
}

// UserStakingTokenRewardsPage returns at most limit per-token reward totals
// of userAid in ascending token order.
func (r *Repo) UserStakingTokenRewardsPage(
	ctx context.Context,
	userAid int64,
	after *UserStakingTokenPageCursor,
	limit int,
) (records []UserStakingTokenRewardRecord, hasMore bool, err error) {
	const op = "user staking token rewards page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
			token_id,
			COALESCE(SUM(reward), 0)::TEXT,
			COALESCE(SUM(reward) FILTER (WHERE collected), 0)::TEXT,
			COALESCE(SUM(reward) FILTER (WHERE NOT collected), 0)::TEXT
		FROM cg_st_reward
		WHERE staker_aid=$1`
	args := []any{userAid}
	if after != nil {
		query += fmt.Sprintf(" AND token_id > $%d", len(args)+1)
		args = append(args, after.TokenID)
	}
	query += fmt.Sprintf(`
		GROUP BY token_id
		ORDER BY token_id
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	scan := func(rows pgx.Rows, rec *UserStakingTokenRewardRecord) error {
		return rows.Scan(
			&rec.TokenID,
			&rec.TotalWei,
			&rec.CollectedWei,
			&rec.PendingWei,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// CosmicSignatureTokenExists reports whether a Cosmic Signature token with
// the given ID has been minted.
func (r *Repo) CosmicSignatureTokenExists(ctx context.Context, tokenID int64) (bool, error) {
	const op = "cosmic signature token exists"
	var exists bool
	err := r.pool().QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM cg_mint_event WHERE token_id=$1)",
		tokenID,
	).Scan(&exists)
	if err != nil {
		return false, store.WrapError(op, err)
	}
	return exists, nil
}

// UserStakingTokenDepositPageCursor identifies the last deposit returned by
// the ascending per-token reward breakdown.
type UserStakingTokenDepositPageCursor struct {
	DepositID int64
}

func (c *UserStakingTokenDepositPageCursor) valid() bool {
	return c == nil || c.DepositID >= 0
}

// UserStakingTokenRewardDepositRecord is one wallet's reward from one
// staking deposit that one staked token participated in. The transaction
// identity is the deposit event's.
type UserStakingTokenRewardDepositRecord struct {
	Tx        cgmodel.Transaction
	StakerAid int64
	DepositID int64
	RoundNum  int64
	RewardWei string
	Claimed   bool
}

// UserStakingTokenRewardDepositsPage returns at most limit per-deposit
// rewards one token earned for userAid, in ascending deposit order.
func (r *Repo) UserStakingTokenRewardDepositsPage(
	ctx context.Context,
	userAid int64,
	tokenID int64,
	after *UserStakingTokenDepositPageCursor,
	limit int,
) (records []UserStakingTokenRewardDepositRecord, hasMore bool, err error) {
	const op = "user staking token reward deposits page"
	if userAid < 1 || tokenID < 0 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id, token id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			r.staker_aid,
			r.deposit_id,
			d.round_num,
			r.reward::TEXT,
			r.collected
		FROM cg_st_reward r
			INNER JOIN cg_staking_eth_deposit d ON d.deposit_id=r.deposit_id
			INNER JOIN transaction t ON t.id=d.tx_id
		WHERE r.staker_aid=$1 AND r.token_id=$2`
	args := []any{userAid, tokenID}
	if after != nil {
		query += fmt.Sprintf(" AND r.deposit_id > $%d", len(args)+1)
		args = append(args, after.DepositID)
	}
	query += fmt.Sprintf(`
		ORDER BY r.deposit_id
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	scan := func(rows pgx.Rows, rec *UserStakingTokenRewardDepositRecord) error {
		return rows.Scan(
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.StakerAid,
			&rec.DepositID,
			&rec.RoundNum,
			&rec.RewardWei,
			&rec.Claimed,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}
