package cosmicgame

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// erc20DonationColumns is the shared SELECT list of the ERC-20 donation
// queries that join the round's prize claim (alias tok = cg_erc20_donation,
// t = transaction, da = donor address, tokaddr = token contract address,
// p = cg_prize_claim, wa = winner address).
const erc20DonationColumns = `
			tok.id,
			tok.evtlog_id,
			tok.block_num,
			tok.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM tok.time_stamp)::BIGINT,
			tok.time_stamp,
			tok.round_num,
			tok.donor_aid,
			da.addr,
			tokaddr.address_id,
			tokaddr.addr,
			tok.amount,
			tok.amount/1e18,
			p.winner_aid,
			wa.addr`

// scanERC20DonationWithWinner scans erc20DonationColumns; the winner columns
// come from a join on cg_prize_claim and are NULL until the round is
// claimed.
func scanERC20DonationWithWinner(rows pgx.Rows, rec *p.CGERC20Donation) error {
	var nullWinnerAid sql.NullInt64
	var nullWinnerAddr sql.NullString
	err := rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.RoundNum,
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.TokenAid,
		&rec.TokenAddr,
		&rec.Amount,
		&rec.AmountEth,
		&nullWinnerAid,
		&nullWinnerAddr,
	)
	if err != nil {
		return err
	}
	if nullWinnerAid.Valid {
		rec.WinnerAid = nullWinnerAid.Int64
	}
	if nullWinnerAddr.Valid {
		rec.WinnerAddr = nullWinnerAddr.String
	}
	return nil
}

// ERC20DonationsByRoundDetailed returns one row per ERC-20 donation of a
// claimed round, with the main-prize winner attached. Rounds without a prize
// claim yield nothing (INNER JOIN); use ERC20DonationsByRoundAll to include
// donations of unclaimed rounds.
func (r *Repo) ERC20DonationsByRoundDetailed(ctx context.Context, roundNum int64) ([]p.CGERC20Donation, error) {
	query := `SELECT ` + erc20DonationColumns + `
		FROM cg_erc20_donation tok
			INNER JOIN cg_prize_claim p ON p.round_num=tok.round_num
			LEFT JOIN transaction t ON t.id=tok.tx_id
			LEFT JOIN address da ON tok.donor_aid=da.address_id
			LEFT JOIN address tokaddr ON tok.token_aid=tokaddr.address_id
			LEFT JOIN address wa ON wa.address_id = p.winner_aid
		WHERE tok.round_num= $1
		ORDER BY tok.id DESC`
	return queryList(ctx, r, "erc20 donations by round detailed", 256, query, scanERC20DonationWithWinner, roundNum)
}

// ERC20DonationsByRoundAll returns every ERC-20 donation of a round, with
// the main-prize winner fields populated once the round is claimed (LEFT
// JOIN) — unlike ERC20DonationsByRoundDetailed, rows appear before the claim.
func (r *Repo) ERC20DonationsByRoundAll(ctx context.Context, roundNum int64) ([]p.CGERC20Donation, error) {
	query := `SELECT DISTINCT ON (tok.id) ` + erc20DonationColumns + `
		FROM cg_erc20_donation tok
			LEFT JOIN cg_prize_claim p ON p.round_num=tok.round_num
			LEFT JOIN transaction t ON t.id=tok.tx_id
			LEFT JOIN address da ON tok.donor_aid=da.address_id
			LEFT JOIN address tokaddr ON tok.token_aid=tokaddr.address_id
			LEFT JOIN address wa ON wa.address_id = p.winner_aid
		WHERE tok.round_num= $1
		ORDER BY tok.id DESC, p.id DESC NULLS LAST`
	return queryList(ctx, r, "erc20 donations by round all", 256, query, scanERC20DonationWithWinner, roundNum)
}

// ERC20DonationsByRoundSummarized aggregates a claimed round's ERC-20
// donations per token contract: total donated, total claimed so far and the
// remaining difference, plus the round winner entitled to claim.
func (r *Repo) ERC20DonationsByRoundSummarized(ctx context.Context, roundNum int64) ([]p.CGSummarizedERC20Donation, error) {
	query := `WITH claim AS (
			SELECT SUM(amount) total,round_num,token_aid,winner_aid
			FROM cg_donated_tok_claimed GROUP BY round_num,token_aid,winner_aid
		)
		SELECT
			p.id,
			p.evtlog_id,
			p.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,
			p.time_stamp,
			dt20.round_num,
			tokaddr.address_id,
			tokaddr.addr,
			dt20.total_amount,
			dt20.total_amount/1e18,
			COALESCE(claim.total,0),
			COALESCE(claim.total,0)/1e18,
			dt20.total_amount-COALESCE(claim.total,0),
			(dt20.total_amount-COALESCE(claim.total,0))/1e18,
			p.winner_aid,
			wa.addr,
			dt20.claimed
		FROM cg_erc20_donation_stats dt20
			INNER JOIN cg_prize_claim p ON p.round_num=dt20.round_num
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address tokaddr ON dt20.token_aid=tokaddr.address_id
			LEFT JOIN claim ON (claim.token_aid=dt20.token_aid AND dt20.round_num=claim.round_num)
			LEFT JOIN address wa ON wa.address_id = p.winner_aid
		WHERE p.round_num= $1
		ORDER BY dt20.token_aid DESC`
	scan := func(rows pgx.Rows, rec *p.CGSummarizedERC20Donation) error {
		var nullWinnerAid sql.NullInt64
		var nullWinnerAddr sql.NullString
		err := rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.RoundNum,
			&rec.TokenAid,
			&rec.TokenAddr,
			&rec.AmountDonated,
			&rec.AmountDonatedEth,
			&rec.AmountClaimed,
			&rec.AmountClaimedEth,
			&rec.DonateClaimDiff,
			&rec.DonateClaimDiffEth,
			&nullWinnerAid,
			&nullWinnerAddr,
			&rec.Claimed,
		)
		if err != nil {
			return err
		}
		if nullWinnerAid.Valid {
			rec.WinnerAid = nullWinnerAid.Int64
		}
		if nullWinnerAddr.Valid {
			rec.WinnerAddr = nullWinnerAddr.String
		}
		return nil
	}
	return queryList(ctx, r, "erc20 donations by round summarized", 256, query, scan, roundNum)
}

// ERC20Donations returns the ERC-20 donations of claimed rounds, newest
// first, with the round winner attached (INNER JOIN on the prize claim).
func (r *Repo) ERC20Donations(ctx context.Context, offset, limit int) ([]p.CGERC20Donation, error) {
	query := `SELECT ` + erc20DonationColumns + `
		FROM cg_erc20_donation tok
			INNER JOIN cg_prize_claim p ON p.round_num=tok.round_num
			LEFT JOIN transaction t ON t.id=tok.tx_id
			LEFT JOIN address da ON tok.donor_aid=da.address_id
			LEFT JOIN address tokaddr ON tok.token_aid=tokaddr.address_id
			LEFT JOIN address wa ON wa.address_id = p.winner_aid
		ORDER BY tok.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "erc20 donations", 256, query, scanERC20DonationWithWinner, offset, limit)
}

// ERC20DonationInfo returns one ERC-20 donation by record id, or
// store.ErrNotFound when the id does not exist. The returned record does not
// carry amounts or winner fields (the legacy query never selected them).
func (r *Repo) ERC20DonationInfo(ctx context.Context, id int64) (p.CGERC20Donation, error) {
	query := `SELECT
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			d.round_num,
			d.donor_aid,
			da.addr,
			toka.address_id,
			toka.addr
		FROM cg_erc20_donation d
			LEFT JOIN transaction t ON t.id=tx_id
			LEFT JOIN address da ON d.donor_aid=da.address_id
			LEFT JOIN address toka ON d.token_aid=toka.address_id
		WHERE d.id=$1`
	var rec p.CGERC20Donation
	rec.RecordId = id
	err := r.pool().QueryRow(ctx, query, id).Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.RoundNum,
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.TokenAid,
		&rec.TokenAddr,
	)
	if err != nil {
		return p.CGERC20Donation{RecordId: id}, store.WrapError("erc20 donation info", err)
	}
	return rec, nil
}

// ERC20DonationsByUser returns every ERC-20 donation made by one donor,
// newest first (no winner columns — donations are keyed by donor here).
func (r *Repo) ERC20DonationsByUser(ctx context.Context, donorAid int64) ([]p.CGERC20Donation, error) {
	query := `SELECT
			tok.id,
			tok.evtlog_id,
			tok.block_num,
			tok.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM tok.time_stamp)::BIGINT,
			tok.time_stamp,
			tok.round_num,
			tok.donor_aid,
			da.addr,
			tokaddr.address_id,
			tokaddr.addr,
			tok.amount,
			tok.amount/1e18
		FROM cg_erc20_donation tok
			LEFT JOIN transaction t ON t.id=tok.tx_id
			LEFT JOIN address da ON tok.donor_aid=da.address_id
			LEFT JOIN address tokaddr ON tok.token_aid=tokaddr.address_id
		WHERE tok.donor_aid=$1
		ORDER BY tok.id DESC`
	scan := func(rows pgx.Rows, rec *p.CGERC20Donation) error {
		return rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.RoundNum,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.TokenAid,
			&rec.TokenAddr,
			&rec.Amount,
			&rec.AmountEth,
		)
	}
	return queryList(ctx, r, "erc20 donations by user", 256, query, scan, donorAid)
}

// erc20ClaimSelectSQL is the shared SELECT of the donated-token claim
// queries (alias c = cg_donated_tok_claimed); the donor address comes from
// joining the donation of the same round and token and is NULL when that
// donation is missing.
const erc20ClaimSelectSQL = `SELECT
			c.id,
			c.evtlog_id,
			c.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM c.time_stamp)::BIGINT,
			c.time_stamp,
			c.round_num,
			c.idx,
			tokaddr.address_id,
			tokaddr.addr,
			c.amount,
			c.amount/1e18,
			c.winner_aid,
			wa.addr,
			da.addr
		FROM cg_donated_tok_claimed c
			LEFT JOIN transaction t ON t.id=c.tx_id
			LEFT JOIN address tokaddr ON c.token_aid=tokaddr.address_id
			LEFT JOIN address wa ON c.winner_aid=wa.address_id
			LEFT JOIN cg_erc20_donation d ON (d.round_num=c.round_num AND d.token_aid=c.token_aid)
			LEFT JOIN address da ON d.donor_aid=da.address_id`

func scanERC20Claim(rows pgx.Rows, rec *p.CGERC20ClaimRec) error {
	var nullDonorAddr sql.NullString
	err := rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.RoundNum,
		&rec.Index,
		&rec.TokenAid,
		&rec.TokenAddr,
		&rec.Amount,
		&rec.AmountEth,
		&rec.WinnerAid,
		&rec.WinnerAddr,
		&nullDonorAddr,
	)
	if err != nil {
		return err
	}
	if nullDonorAddr.Valid {
		rec.DonorAddr = nullDonorAddr.String
	}
	return nil
}

// ERC20DonationClaims returns every claim of donated ERC-20 tokens, newest
// first.
func (r *Repo) ERC20DonationClaims(ctx context.Context, offset, limit int) ([]p.CGERC20ClaimRec, error) {
	query := erc20ClaimSelectSQL + `
		ORDER BY c.id DESC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "erc20 donation claims", 256, query, scanERC20Claim, offset, limit)
}

// ERC20DonationClaimsByUser returns one winner's claims of donated ERC-20
// tokens, newest first.
func (r *Repo) ERC20DonationClaimsByUser(ctx context.Context, winnerAid int64) ([]p.CGERC20ClaimRec, error) {
	query := erc20ClaimSelectSQL + `
		WHERE c.winner_aid=$1
		ORDER BY c.id DESC`
	return queryList(ctx, r, "erc20 donation claims by user", 256, query, scanERC20Claim, winnerAid)
}

// ERC20DonationClaimsByRound returns the claims of donated ERC-20 tokens in
// one round, newest first.
func (r *Repo) ERC20DonationClaimsByRound(ctx context.Context, roundNum int64) ([]p.CGERC20ClaimRec, error) {
	query := erc20ClaimSelectSQL + `
		WHERE c.round_num=$1
		ORDER BY c.id DESC`
	return queryList(ctx, r, "erc20 donation claims by round", 256, query, scanERC20Claim, roundNum)
}

// RoundERC20DonationRecord is the exact-base-unit event projection used by
// the v2 round donation collection. It intentionally does not join the
// eventual prize winner or aggregate claims onto an immutable donation event.
type RoundERC20DonationRecord struct {
	Tx              p.Transaction
	RoundNum        int64
	DonorAddr       string
	TokenAddr       string
	AmountBaseUnits string
}

const roundERC20DonationsSelect = `SELECT
			tok.evtlog_id,
			tok.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM tok.time_stamp)::BIGINT,
			tok.time_stamp,
			tok.round_num,
			da.addr,
			token.addr,
			tok.amount::TEXT
		FROM cg_erc20_donation tok
			LEFT JOIN transaction t ON t.id=tok.tx_id
			LEFT JOIN address da ON da.address_id=tok.donor_aid
			LEFT JOIN address token ON token.address_id=tok.token_aid`

func scanRoundERC20Donation(rows pgx.Rows, rec *RoundERC20DonationRecord) error {
	return rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.RoundNum,
		&rec.DonorAddr,
		&rec.TokenAddr,
		&rec.AmountBaseUnits,
	)
}

// ERC20DonationsByRoundPage returns at most limit ERC-20 donation events
// before the supplied newest-first event cursor.
func (r *Repo) ERC20DonationsByRoundPage(
	ctx context.Context,
	roundNum int64,
	after *DonationPageCursor,
	limit int,
) (records []RoundERC20DonationRecord, hasMore bool, err error) {
	const op = "erc20 donations by round page"
	if roundNum < 0 {
		return nil, false, fmt.Errorf("%s: round must be non-negative", op)
	}
	if limit <= 0 || limit > maxDonationPageLimit {
		return nil, false, fmt.Errorf("%s: limit must be between 1 and %d", op, maxDonationPageLimit)
	}

	query := roundERC20DonationsSelect + `
		WHERE tok.round_num=$1
		ORDER BY tok.evtlog_id DESC
		LIMIT $2`
	args := []any{roundNum, limit + 1}
	if after != nil {
		if after.EventLogID < 1 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		query = roundERC20DonationsSelect + `
			WHERE tok.round_num=$1 AND tok.evtlog_id < $2
			ORDER BY tok.evtlog_id DESC
			LIMIT $3`
		args = []any{roundNum, after.EventLogID, limit + 1}
	}

	records, err = queryList(ctx, r, op, limit+1, query, scanRoundERC20Donation, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		records = records[:limit]
		hasMore = true
	}
	return records, hasMore, nil
}
