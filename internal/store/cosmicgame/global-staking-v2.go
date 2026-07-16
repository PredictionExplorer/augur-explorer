package cosmicgame

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// GlobalStakingActionPageCursor identifies the last immutable event returned
// by a newest-first global staking action page.
type GlobalStakingActionPageCursor struct {
	EventLogID int64
}

func (c *GlobalStakingActionPageCursor) valid() bool {
	return c == nil || c.EventLogID >= 1
}

// GlobalStakingActionRecord is one stake or unstake event in a global staking
// ledger. RewardWei and RewardPerTokenWei are populated only for Cosmic
// Signature unstake events.
type GlobalStakingActionRecord struct {
	Tx                cgmodel.Transaction
	Kind              UserStakingActionKind
	StakerAid         int64
	StakerAddress     string
	ActionID          int64
	TokenID           int64
	RoundNum          int64
	TotalStakedNfts   int64
	RewardWei         string
	RewardPerTokenWei string
}

// globalStakingActionsPageSQL builds the bounded two-branch merge used by
// both global action ledgers. All table names and reward expressions are
// compile-time literals supplied by the two public methods below.
func globalStakingActionsPageSQL(
	stakeTable, unstakeTable, rewardExpr, rewardPerTokenExpr string,
	after bool,
) string {
	filter := ""
	limitPlaceholder := "$1"
	if after {
		filter = "WHERE e.evtlog_id < $1"
		limitPlaceholder = "$2"
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
			staker_address,
			action_id,
			token_id,
			round_num,
			num_staked_nfts,
			reward_wei,
			reward_per_token_wei
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
				a.addr AS staker_address,
				e.action_id,
				e.token_id,
				e.round_num,
				e.num_staked_nfts,
				NULL::TEXT AS reward_wei,
				NULL::TEXT AS reward_per_token_wei
			FROM %s e
				INNER JOIN transaction t ON t.id=e.tx_id
				INNER JOIN address a ON a.address_id=e.staker_aid
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
				a.addr AS staker_address,
				e.action_id,
				e.token_id,
				e.round_num,
				e.num_staked_nfts,
				%s AS reward_wei,
				%s AS reward_per_token_wei
			FROM %s e
				INNER JOIN transaction t ON t.id=e.tx_id
				INNER JOIN address a ON a.address_id=e.staker_aid
			%s
			ORDER BY e.evtlog_id DESC
			LIMIT %s)
		) actions
		ORDER BY evtlog_id DESC
		LIMIT %s`,
		stakeTable, filter, limitPlaceholder,
		rewardExpr, rewardPerTokenExpr, unstakeTable, filter, limitPlaceholder,
		limitPlaceholder)
}

func scanGlobalStakingAction(rows pgx.Rows, record *GlobalStakingActionRecord) error {
	var reward, rewardPerToken sql.NullString
	if err := rows.Scan(
		&record.Kind,
		&record.Tx.EvtLogId,
		&record.Tx.BlockNum,
		&record.Tx.TxId,
		&record.Tx.TxHash,
		&record.Tx.TimeStamp,
		store.TimeText(&record.Tx.DateTime),
		&record.StakerAid,
		&record.StakerAddress,
		&record.ActionID,
		&record.TokenID,
		&record.RoundNum,
		&record.TotalStakedNfts,
		&reward,
		&rewardPerToken,
	); err != nil {
		return err
	}
	if reward.Valid {
		record.RewardWei = reward.String
	}
	if rewardPerToken.Valid {
		record.RewardPerTokenWei = rewardPerToken.String
	}
	return nil
}

func (r *Repo) globalStakingActionsPage(
	ctx context.Context,
	op, stakeTable, unstakeTable, rewardExpr, rewardPerTokenExpr string,
	after *GlobalStakingActionPageCursor,
	limit int,
) (records []GlobalStakingActionRecord, hasMore bool, err error) {
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	args := []any{limit + 1}
	if after != nil {
		args = []any{after.EventLogID, limit + 1}
	}
	records, err = queryList(
		ctx,
		r,
		op,
		limit+1,
		globalStakingActionsPageSQL(
			stakeTable,
			unstakeTable,
			rewardExpr,
			rewardPerTokenExpr,
			after != nil,
		),
		scanGlobalStakingAction,
		args...,
	)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// GlobalCstStakingActionsPage returns at most limit Cosmic Signature stake
// and unstake events, newest first by immutable event-log ID.
func (r *Repo) GlobalCstStakingActionsPage(
	ctx context.Context,
	after *GlobalStakingActionPageCursor,
	limit int,
) ([]GlobalStakingActionRecord, bool, error) {
	return r.globalStakingActionsPage(
		ctx,
		"global cst staking actions page",
		"cg_nft_staked_cst",
		"cg_nft_unstaked_cst",
		"e.reward::TEXT",
		"e.reward_per_tok::TEXT",
		after,
		limit,
	)
}

// GlobalRwalkStakingActionsPage returns at most limit RandomWalk stake and
// unstake events, newest first by immutable event-log ID.
func (r *Repo) GlobalRwalkStakingActionsPage(
	ctx context.Context,
	after *GlobalStakingActionPageCursor,
	limit int,
) ([]GlobalStakingActionRecord, bool, error) {
	return r.globalStakingActionsPage(
		ctx,
		"global rwalk staking actions page",
		"cg_nft_staked_rwalk",
		"cg_nft_unstaked_rwalk",
		"NULL::TEXT",
		"NULL::TEXT",
		after,
		limit,
	)
}

// GlobalStakedTokenPageCursor identifies the last token returned by an
// ascending global live-membership page.
type GlobalStakedTokenPageCursor struct {
	TokenID int64
}

func (c *GlobalStakedTokenPageCursor) valid() bool {
	return c == nil || c.TokenID >= 0
}

// GlobalStakedCstTokenRecord is one currently staked Cosmic Signature token,
// its staker, locking action and mint provenance.
type GlobalStakedCstTokenRecord struct {
	StakeTx       cgmodel.Transaction
	StakerAid     int64
	StakerAddress string
	ActionID      int64
	TokenID       int64
	MintRound     int64
	Seed          string
	TokenName     string
}

// GlobalStakedCstTokensPage returns at most limit currently staked Cosmic
// Signature tokens in ascending token order.
func (r *Repo) GlobalStakedCstTokensPage(
	ctx context.Context,
	after *GlobalStakedTokenPageCursor,
	limit int,
) (records []GlobalStakedCstTokenRecord, hasMore bool, err error) {
	const op = "global staked cst tokens page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
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
			sa.addr,
			a.action_id,
			st.token_id,
			m.round_num,
			m.seed,
			m.token_name
		FROM cg_staked_token_cst st
			INNER JOIN cg_nft_staked_cst a ON a.action_id=st.stake_action_id
			INNER JOIN cg_mint_event m ON m.token_id=st.token_id
			INNER JOIN address sa ON sa.address_id=st.staker_aid
			INNER JOIN transaction t ON t.id=a.tx_id`
	args := []any{}
	if after != nil {
		query += " WHERE st.token_id > $1"
		args = append(args, after.TokenID)
	}
	query += fmt.Sprintf(`
		ORDER BY st.token_id
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)
	scan := func(rows pgx.Rows, record *GlobalStakedCstTokenRecord) error {
		return rows.Scan(
			&record.StakeTx.EvtLogId,
			&record.StakeTx.BlockNum,
			&record.StakeTx.TxId,
			&record.StakeTx.TxHash,
			&record.StakeTx.TimeStamp,
			store.TimeText(&record.StakeTx.DateTime),
			&record.StakerAid,
			&record.StakerAddress,
			&record.ActionID,
			&record.TokenID,
			&record.MintRound,
			&record.Seed,
			&record.TokenName,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// GlobalStakedRwalkTokenRecord is one currently staked RandomWalk token, its
// staker and locking action.
type GlobalStakedRwalkTokenRecord struct {
	StakeTx       cgmodel.Transaction
	StakerAid     int64
	StakerAddress string
	ActionID      int64
	TokenID       int64
}

// GlobalStakedRwalkTokensPage returns at most limit currently staked
// RandomWalk tokens in ascending token order.
func (r *Repo) GlobalStakedRwalkTokensPage(
	ctx context.Context,
	after *GlobalStakedTokenPageCursor,
	limit int,
) (records []GlobalStakedRwalkTokenRecord, hasMore bool, err error) {
	const op = "global staked rwalk tokens page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
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
			sa.addr,
			a.action_id,
			st.token_id
		FROM cg_staked_token_rwalk st
			INNER JOIN cg_nft_staked_rwalk a ON a.action_id=st.stake_action_id
			INNER JOIN address sa ON sa.address_id=st.staker_aid
			INNER JOIN transaction t ON t.id=a.tx_id`
	args := []any{}
	if after != nil {
		query += " WHERE st.token_id > $1"
		args = append(args, after.TokenID)
	}
	query += fmt.Sprintf(`
		ORDER BY st.token_id
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)
	scan := func(rows pgx.Rows, record *GlobalStakedRwalkTokenRecord) error {
		return rows.Scan(
			&record.StakeTx.EvtLogId,
			&record.StakeTx.BlockNum,
			&record.StakeTx.TxId,
			&record.StakeTx.TxHash,
			&record.StakeTx.TimeStamp,
			store.TimeText(&record.StakeTx.DateTime),
			&record.StakerAid,
			&record.StakerAddress,
			&record.ActionID,
			&record.TokenID,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// GlobalStakingDepositPageCursor identifies the last staking deposit returned
// by a newest-first global deposit page.
type GlobalStakingDepositPageCursor struct {
	DepositID int64
}

func (c *GlobalStakingDepositPageCursor) valid() bool {
	return c == nil || c.DepositID >= 0
}

// GlobalStakingDepositRecord is one staking-wallet ETH deposit with exact
// collected, pending and integer-division remainder amounts.
type GlobalStakingDepositRecord struct {
	Tx                 cgmodel.Transaction
	DepositID          int64
	RoundNum           int64
	TotalDepositWei    string
	TotalStakedNfts    int64
	AmountPerTokenWei  string
	CollectedWei       string
	PendingWei         string
	RemainderWei       string
	RewardCount        int64
	PendingRewardCount int64
}

// GlobalStakingDepositsPage returns at most limit staking-wallet ETH deposits,
// newest first by contract-assigned deposit ID.
func (r *Repo) GlobalStakingDepositsPage(
	ctx context.Context,
	after *GlobalStakingDepositPageCursor,
	limit int,
) (records []GlobalStakingDepositRecord, hasMore bool, err error) {
	const op = "global staking deposits page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	filter := ""
	args := []any{}
	if after != nil {
		filter = "WHERE d.deposit_id < $1"
		args = append(args, after.DepositID)
	}
	query := fmt.Sprintf(`WITH deposits AS (
			SELECT
				d.evtlog_id,
				d.block_num,
				t.id AS tx_id,
				t.tx_hash,
				EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS ts,
				d.time_stamp,
				d.deposit_id,
				d.round_num,
				d.deposit_amount,
				d.accumulated_nfts,
				d.amount_per_token,
				d.modulo
			FROM cg_staking_eth_deposit d
				INNER JOIN transaction t ON t.id=d.tx_id
			%s
			ORDER BY d.deposit_id DESC
			LIMIT $%d
		), rewards AS (
			SELECT
				r.deposit_id,
				COALESCE(SUM(r.reward) FILTER (WHERE r.collected), 0) AS collected_wei,
				COALESCE(SUM(r.reward) FILTER (WHERE NOT r.collected), 0) AS pending_wei,
				COUNT(*) AS reward_count,
				COUNT(*) FILTER (WHERE NOT r.collected) AS pending_count
			FROM cg_st_reward r
				INNER JOIN deposits d ON d.deposit_id=r.deposit_id
			GROUP BY r.deposit_id
		)
		SELECT
			d.evtlog_id,
			d.block_num,
			d.tx_id,
			d.tx_hash,
			d.ts,
			d.time_stamp,
			d.deposit_id,
			d.round_num,
			d.deposit_amount::TEXT,
			d.accumulated_nfts,
			d.amount_per_token::TEXT,
			COALESCE(r.collected_wei, 0)::TEXT,
			COALESCE(r.pending_wei, 0)::TEXT,
			d.modulo::TEXT,
			COALESCE(r.reward_count, 0),
			COALESCE(r.pending_count, 0)
		FROM deposits d
			LEFT JOIN rewards r ON r.deposit_id=d.deposit_id
		ORDER BY d.deposit_id DESC`,
		filter,
		len(args)+1,
	)
	args = append(args, limit+1)
	scan := func(rows pgx.Rows, record *GlobalStakingDepositRecord) error {
		return rows.Scan(
			&record.Tx.EvtLogId,
			&record.Tx.BlockNum,
			&record.Tx.TxId,
			&record.Tx.TxHash,
			&record.Tx.TimeStamp,
			store.TimeText(&record.Tx.DateTime),
			&record.DepositID,
			&record.RoundNum,
			&record.TotalDepositWei,
			&record.TotalStakedNfts,
			&record.AmountPerTokenWei,
			&record.CollectedWei,
			&record.PendingWei,
			&record.RemainderWei,
			&record.RewardCount,
			&record.PendingRewardCount,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// RoundStakingRewardPageCursor identifies the last (deposit, staker) pair
// returned by a round-scoped reward page.
type RoundStakingRewardPageCursor struct {
	DepositID int64
	StakerAid int64
}

func (c *RoundStakingRewardPageCursor) valid() bool {
	return c == nil || (c.DepositID >= 0 && c.StakerAid >= 1)
}

// RoundStakingRewardRecord is one staker's exact share of one staking-wallet
// ETH deposit.
type RoundStakingRewardRecord struct {
	DepositID      int64
	RoundNum       int64
	StakerAid      int64
	StakerAddress  string
	StakedNftCount int64
	RewardWei      string
	CollectedWei   string
	PendingWei     string
}

// RoundStakingRewardsPage returns at most limit staker allocations in one
// round, newest deposit first and then by stable address ID.
func (r *Repo) RoundStakingRewardsPage(
	ctx context.Context,
	roundNum int64,
	after *RoundStakingRewardPageCursor,
	limit int,
) (records []RoundStakingRewardRecord, hasMore bool, err error) {
	const op = "round staking rewards page"
	if roundNum < 0 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid round or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	query := `SELECT
			sd.deposit_id,
			d.round_num,
			sd.staker_aid,
			a.addr,
			sd.tokens_staked,
			sd.amount_deposited::TEXT,
			(sd.amount_deposited - sd.amount_to_claim)::TEXT,
			sd.amount_to_claim::TEXT
		FROM cg_staker_deposit sd
			INNER JOIN cg_staking_eth_deposit d ON d.deposit_id=sd.deposit_id
			INNER JOIN address a ON a.address_id=sd.staker_aid
		WHERE d.round_num=$1`
	args := []any{roundNum}
	if after != nil {
		query += ` AND (
			sd.deposit_id < $2 OR
			(sd.deposit_id = $2 AND sd.staker_aid > $3)
		)`
		args = append(args, after.DepositID, after.StakerAid)
	}
	query += fmt.Sprintf(`
		ORDER BY sd.deposit_id DESC, sd.staker_aid
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)
	scan := func(rows pgx.Rows, record *RoundStakingRewardRecord) error {
		return rows.Scan(
			&record.DepositID,
			&record.RoundNum,
			&record.StakerAid,
			&record.StakerAddress,
			&record.StakedNftCount,
			&record.RewardWei,
			&record.CollectedWei,
			&record.PendingWei,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// GlobalStakerRafflePageCursor identifies the last immutable event returned
// by a newest-first global staker-raffle page.
type GlobalStakerRafflePageCursor struct {
	EventLogID int64
}

func (c *GlobalStakerRafflePageCursor) valid() bool {
	return c == nil || c.EventLogID >= 1
}

// GlobalStakerRaffleNftWinsPage returns at most limit Cosmic Signature NFT
// wins for the selected staking pool, newest first by event-log ID.
func (r *Repo) GlobalStakerRaffleNftWinsPage(
	ctx context.Context,
	isRwalk bool,
	after *GlobalStakerRafflePageCursor,
	limit int,
) (records []cgmodel.CGRaffleNFTWinnerRec, hasMore bool, err error) {
	const op = "global staker raffle nft wins page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	query := `SELECT
			w.id,
			w.evtlog_id,
			w.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM w.time_stamp)::BIGINT,
			w.time_stamp,
			w.token_id,
			w.cst_amount,
			w.cst_amount/1e18,
			w.winner_idx,
			w.round_num,
			w.winner_aid,
			wa.addr
		FROM cg_raffle_nft_prize w
			INNER JOIN transaction t ON t.id=w.tx_id
			INNER JOIN address wa ON wa.address_id=w.winner_aid
		WHERE w.is_staker=TRUE AND w.is_rwalk=$1`
	args := []any{isRwalk}
	if after != nil {
		query += " AND w.evtlog_id < $2"
		args = append(args, after.EventLogID)
	}
	query += fmt.Sprintf(`
		ORDER BY w.evtlog_id DESC
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)
	scan := func(rows pgx.Rows, record *cgmodel.CGRaffleNFTWinnerRec) error {
		if err := scanStakingMint(rows, record); err != nil {
			return err
		}
		record.IsRWalk = isRwalk
		record.IsStaker = true
		return nil
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}
