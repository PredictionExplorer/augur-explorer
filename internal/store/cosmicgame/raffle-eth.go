package cosmicgame

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// prizeDepositColumns is the shared SELECT list of the cg_prize_deposit
// queries (alias p = cg_prize_deposit, t = transaction, wa = winner address).
// The record_type column differs per query and is appended by each caller.
const prizeDepositColumns = `
			p.id,
			p.evtlog_id,
			p.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,
			p.time_stamp,
			p.winner_aid,
			wa.addr,
			p.winner_index,
			p.round_num,
			p.amount/1e18 amount_eth,
			p.claimed, `

func scanPrizeDeposit(rows pgx.Rows, rec *p.CGPrizeDepositRec) error {
	return rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.WinnerAid,
		&rec.WinnerAddr,
		&rec.WinnerIndex,
		&rec.RoundNum,
		&rec.Amount,
		&rec.Claimed,
		&rec.RecordType,
	)
}

// UnclaimedPrizeEthDeposits returns winnerAid's ETH deposits in the prizes
// wallet that have not been withdrawn yet, newest first. Chrono warrior
// deposits are tagged record_type 7, plain raffle deposits 10.
func (r *Repo) UnclaimedPrizeEthDeposits(ctx context.Context, winnerAid int64, offset, limit int) ([]p.CGPrizeDepositRec, error) {
	query := `SELECT
			rd.id,
			rd.evtlog_id,
			rd.block_num,
			rd.tx_id,
			t.tx_hash,
			EXTRACT(EPOCH FROM rd.time_stamp)::BIGINT AS tstmp,
			rd.time_stamp AS date_time,
			wa.addr,
			rd.winner_aid,
			rd.winner_index,
			rd.round_num,
			rd.amount/1e18 AS amount_eth,
			rd.claimed,
			EXTRACT(EPOCH FROM rw.time_stamp)::BIGINT AS tstmp,
			rw.time_stamp,
			CASE WHEN cw.round_num IS NOT NULL THEN 7 ELSE 10 END AS record_type
		FROM cg_prize_deposit rd
			LEFT JOIN cg_prize_withdrawal rw ON rw.evtlog_id=rd.withdrawal_id
			LEFT JOIN transaction t ON t.id=rd.tx_id
			LEFT JOIN address wa ON rd.winner_aid = wa.address_id
			LEFT JOIN cg_chrono_warrior_prize cw ON (rd.round_num = cw.round_num AND rd.winner_index = cw.winner_index)
		WHERE rd.winner_aid=$1 AND rd.claimed='F'
		ORDER BY rd.id DESC
		OFFSET $2 LIMIT $3`
	scan := func(rows pgx.Rows, rec *p.CGPrizeDepositRec) error {
		var claimTs sql.NullInt64
		err := rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.WinnerAddr,
			&rec.WinnerAid,
			&rec.WinnerIndex,
			&rec.RoundNum,
			&rec.Amount,
			&rec.Claimed,
			&claimTs,
			store.NullTimeText(&rec.ClaimDateTime),
			&rec.RecordType,
		)
		if err != nil {
			return err
		}
		if claimTs.Valid {
			rec.ClaimTimeStamp = claimTs.Int64
		}
		return nil
	}
	return queryList(ctx, r, "unclaimed prize eth deposits", 32, query, scan, winnerAid, offset, limit)
}

// PrizeEthDeposits returns every ETH deposit made into the prizes wallet,
// newest first (chrono warrior deposits tagged record_type 7, raffle 10).
// limit 0 means no effective limit.
func (r *Repo) PrizeEthDeposits(ctx context.Context, offset, limit int) ([]p.CGPrizeDepositRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := `SELECT ` + prizeDepositColumns + `
			CASE WHEN cw.round_num IS NOT NULL THEN 7 ELSE 10 END AS record_type
		FROM cg_prize_deposit p
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
			LEFT JOIN cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index)
		ORDER BY p.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "prize eth deposits", 256, query, scanPrizeDeposit, offset, limit)
}

// RaffleEthDeposits returns the plain raffle ETH deposits (excluding chrono
// warrior deposits), newest first. limit 0 means no effective limit.
func (r *Repo) RaffleEthDeposits(ctx context.Context, offset, limit int) ([]p.CGPrizeDepositRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := `SELECT ` + prizeDepositColumns + `
			10 AS record_type
		FROM cg_prize_deposit p
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
			LEFT JOIN cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index)
		WHERE cw.round_num IS NULL
		ORDER BY p.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "raffle eth deposits", 256, query, scanPrizeDeposit, offset, limit)
}

// ChronoWarriorEthDeposits returns the chrono warrior ETH deposits, newest
// first. limit 0 means no effective limit.
func (r *Repo) ChronoWarriorEthDeposits(ctx context.Context, offset, limit int) ([]p.CGPrizeDepositRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := `SELECT ` + prizeDepositColumns + `
			7 AS record_type
		FROM cg_prize_deposit p
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
			INNER JOIN cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index)
		ORDER BY p.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "chrono warrior eth deposits", 256, query, scanPrizeDeposit, offset, limit)
}

// EthDepositsByUser returns every prize-wallet ETH deposit of one winner,
// newest first (chrono warrior tagged record_type 7, raffle 10).
func (r *Repo) EthDepositsByUser(ctx context.Context, winnerAid int64) ([]p.CGPrizeDepositRec, error) {
	query := `SELECT ` + prizeDepositColumns + `
			CASE WHEN cw.round_num IS NOT NULL THEN 7 ELSE 10 END AS record_type
		FROM cg_prize_deposit p
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
			LEFT JOIN cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index)
		WHERE p.winner_aid=$1
		ORDER BY p.id DESC`
	return queryList(ctx, r, "eth deposits by user", 256, query, scanPrizeDeposit, winnerAid)
}

// RaffleEthDepositsByUser returns one winner's plain raffle ETH deposits
// (excluding chrono warrior deposits), newest first.
func (r *Repo) RaffleEthDepositsByUser(ctx context.Context, winnerAid int64) ([]p.CGPrizeDepositRec, error) {
	query := `SELECT ` + prizeDepositColumns + `
			10 AS record_type
		FROM cg_prize_deposit p
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
			LEFT JOIN cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index)
		WHERE p.winner_aid=$1 AND cw.round_num IS NULL
		ORDER BY p.id DESC`
	return queryList(ctx, r, "raffle eth deposits by user", 256, query, scanPrizeDeposit, winnerAid)
}

// ChronoWarriorEthDepositsByUser returns one winner's chrono warrior ETH
// deposits, newest first.
func (r *Repo) ChronoWarriorEthDepositsByUser(ctx context.Context, winnerAid int64) ([]p.CGPrizeDepositRec, error) {
	query := `SELECT ` + prizeDepositColumns + `
			7 AS record_type
		FROM cg_prize_deposit p
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
			INNER JOIN cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index)
		WHERE p.winner_aid=$1
		ORDER BY p.id DESC`
	return queryList(ctx, r, "chrono warrior eth deposits by user", 256, query, scanPrizeDeposit, winnerAid)
}

// PrizeDepositsByRound returns the bidder-raffle ETH deposits of one round,
// ordered by winner index. The cg_prize ptype=10 join deliberately excludes
// chrono-warrior wallet deposits. Repo.PrizeInfo composes it.
func (r *Repo) PrizeDepositsByRound(ctx context.Context, roundNum int64) ([]p.CGPrizeDepositRec, error) {
	query := `SELECT
			p.id,
			p.evtlog_id,
			p.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,
			p.time_stamp,
			p.winner_aid,
			wa.addr,
			p.winner_index,
			p.round_num,
			p.amount/1e18 amount_eth,
			p.claimed,
			CASE WHEN cw.round_num IS NOT NULL THEN 7 ELSE 10 END AS record_type
		FROM cg_prize_deposit p
			INNER JOIN cg_prize pr ON (pr.round_num = p.round_num AND pr.winner_index = p.winner_index AND pr.ptype = 10)
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
			LEFT JOIN cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index)
		WHERE p.round_num = $1
		ORDER BY p.winner_index`
	return queryList(ctx, r, "prize deposits by round", 32, query, scanPrizeDeposit, roundNum)
}

// RaffleEthDepositRecord is the exact-wei store projection used by the v2
// round raffle-deposit resource. The legacy CGPrizeDepositRec retains its
// float amount for frozen v1 responses.
type RaffleEthDepositRecord struct {
	Tx           p.Transaction
	RoundNum     int64
	WinnerIndex  int64
	WinnerAddr   string
	EthAmountWei string
	Claimed      bool
}

// RaffleEthDepositPageCursor identifies the last deposit returned by
// RaffleEthDepositsByRoundPage.
type RaffleEthDepositPageCursor struct {
	WinnerIndex int64
	EventLogID  int64
}

const raffleEthDepositsByRoundSelect = `SELECT
			p.evtlog_id,
			p.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,
			p.time_stamp,
			wa.addr,
			p.winner_index,
			p.round_num,
			p.amount::text,
			p.claimed
		FROM cg_prize_deposit p
			INNER JOIN cg_prize pr ON (
				pr.round_num = p.round_num
				AND pr.winner_index = p.winner_index
				AND pr.ptype = 10
			)
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id`

func scanRaffleEthDepositRecord(rows pgx.Rows, rec *RaffleEthDepositRecord) error {
	return rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.WinnerAddr,
		&rec.WinnerIndex,
		&rec.RoundNum,
		&rec.EthAmountWei,
		&rec.Claimed,
	)
}

// RaffleEthDepositsByRoundPage returns at most limit bidder-raffle ETH
// wallet deposits after the supplied ascending keyset cursor. A nil cursor
// starts at the first winner index.
func (r *Repo) RaffleEthDepositsByRoundPage(
	ctx context.Context,
	roundNum int64,
	after *RaffleEthDepositPageCursor,
	limit int,
) (records []RaffleEthDepositRecord, hasMore bool, err error) {
	const op = "raffle eth deposits by round page"
	if roundNum < 0 {
		return nil, false, fmt.Errorf("%s: round must be non-negative", op)
	}
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: limit must be positive", op)
	}

	query := raffleEthDepositsByRoundSelect + `
		WHERE p.round_num=$1
		ORDER BY p.winner_index, p.evtlog_id
		LIMIT $2`
	args := []any{roundNum, limit + 1}
	if after != nil {
		if after.WinnerIndex < 0 || after.EventLogID < 1 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		query = raffleEthDepositsByRoundSelect + `
			WHERE p.round_num=$1
				AND (p.winner_index, p.evtlog_id) > ($2, $3)
			ORDER BY p.winner_index, p.evtlog_id
			LIMIT $4`
		args = []any{roundNum, after.WinnerIndex, after.EventLogID, limit + 1}
	}

	records, err = queryList(ctx, r, op, limit+1, query, scanRaffleEthDepositRecord, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		records = records[:limit]
		hasMore = true
	}
	return records, hasMore, nil
}
