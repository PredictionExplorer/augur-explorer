// Read queries for the ChaosZero JSON API (internal/api/chaoszero): per-table
// lists with address joins, by-user filters, and the combined round
// summary/detail and statistics views. CST amounts are scaled to whole CST
// in SQL (value/1e18), matching the RandomWalk API convention.
package chaoszero

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	czmodel "github.com/PredictionExplorer/augur-explorer/internal/model/chaoszero"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// queryList runs a SELECT and scans every row with scanRow, wrapping any
// failure with op context. The result is always a non-nil slice so an empty
// result marshals as [].
func queryList[T any](ctx context.Context, r *Repo, op string, capHint int, query string, scanRow func(pgx.Rows, *T) error, args ...any) ([]T, error) {
	return store.QueryList(ctx, r.q(ctx), op, capHint, query, scanRow, args...)
}

// eventRowColumns are the shared leading columns of every per-event list
// query; scanEventRow consumes them in the same order. The tx hash is a
// primary-key subquery so every event query gets it without repeating a
// JOIN; lists are offset/limit bounded so the per-row lookup stays cheap.
const eventRowColumns = `t.id,COALESCE(t.evtlog_id,0),t.block_num,t.tx_id,
	COALESCE((SELECT tx_hash FROM transaction WHERE id=t.tx_id),''),
	EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,t.time_stamp`

func scanEventRow(rows pgx.Rows, rec *czmodel.EventRow, rest ...any) error {
	dest := append([]any{
		&rec.RecordId, &rec.EvtLogId, &rec.BlockNum, &rec.TxId, &rec.TxHash,
		&rec.TimeStamp, store.TimeText(&rec.DateTime),
	}, rest...)
	return rows.Scan(dest...)
}

// =============================================================================
// COMBINED VIEWS
// =============================================================================

// Statistics aggregates the dashboard totals over every cz_ table.
func (r *Repo) Statistics(ctx context.Context) (czmodel.ApiStatistics, error) {
	const op = "chaoszero statistics"
	var s czmodel.ApiStatistics
	query := `SELECT
		(SELECT COUNT(*) FROM cz_round_initialized),
		(SELECT COALESCE(MAX(round_id),0) FROM cz_round_initialized),
		(SELECT COUNT(*) FROM cz_round_funded),
		(SELECT COALESCE(SUM((donated+transfers_used+carry_used)/1e18),0)::FLOAT8 FROM cz_round_funded),
		(SELECT COUNT(*) FROM cz_trade),
		(SELECT COALESCE(SUM(cash/1e18),0)::FLOAT8 FROM cz_trade),
		(SELECT COUNT(DISTINCT user_aid) FROM cz_trade),
		(SELECT COUNT(*) FROM cz_position_split),
		(SELECT COUNT(*) FROM cz_positions_merged),
		(SELECT COUNT(*) FROM cz_round_settled),
		(SELECT COALESCE(SUM(carried/1e18),0)::FLOAT8 FROM cz_round_settled),
		(SELECT COUNT(*) FROM cz_claim),
		(SELECT COALESCE(SUM(total/1e18),0)::FLOAT8 FROM cz_claim),
		(SELECT COUNT(*) FROM cz_paid),
		(SELECT COALESCE(SUM(amount/1e18),0)::FLOAT8 FROM cz_paid),
		(SELECT COUNT(*) FROM cz_position_transfer),
		(SELECT COUNT(*) FROM cz_approval_for_all),
		(SELECT COUNT(*) FROM cz_answers_observed)`
	err := r.q(ctx).QueryRow(ctx, query).Scan(
		&s.NumRounds, &s.LatestRound, &s.NumFundings, &s.FundedTotal,
		&s.NumTrades, &s.Volume, &s.NumTraders,
		&s.NumSplits, &s.NumMerges,
		&s.NumSettled, &s.CarriedTotal,
		&s.NumClaims, &s.ClaimedTotal,
		&s.NumPayouts, &s.PaidTotal,
		&s.NumTransfers, &s.NumApprovals, &s.NumAnswerEvents,
	)
	if err != nil {
		return s, store.WrapError(op, err)
	}
	return s, nil
}

// roundSummaryQuery aggregates the combined per-round view; %s receives an
// optional WHERE clause on ri.round_id.
const roundSummaryQuery = `SELECT
		ri.round_id,
		EXTRACT(EPOCH FROM ri.time_stamp)::BIGINT,
		ri.time_stamp,
		COALESCE(f.num_fundings,0), COALESCE(f.funded_total,0)::FLOAT8,
		COALESCE(tr.num_trades,0), COALESCE(tr.volume,0)::FLOAT8, COALESCE(tr.num_traders,0),
		(s.round_id IS NOT NULL), COALESCE(s.values_mask,0), COALESCE(s.carried,0)::FLOAT8,
		COALESCE(e.escrowed_total,0)::FLOAT8,
		COALESCE(c.num_claims,0), COALESCE(c.claimed_total,0)::FLOAT8
	FROM cz_round_initialized ri
	LEFT JOIN (
		SELECT round_id, COUNT(*) num_fundings,
			SUM((donated+transfers_used+carry_used)/1e18) funded_total
		FROM cz_round_funded GROUP BY round_id
	) f ON f.round_id = ri.round_id
	LEFT JOIN (
		SELECT round_id, COUNT(*) num_trades, SUM(cash/1e18) volume,
			COUNT(DISTINCT user_aid) num_traders
		FROM cz_trade GROUP BY round_id
	) tr ON tr.round_id = ri.round_id
	LEFT JOIN (
		SELECT round_id, values_mask, carried/1e18 carried
		FROM cz_round_settled
	) s ON s.round_id = ri.round_id
	LEFT JOIN (
		SELECT round_id,
			(SELECT SUM(x)/1e18 FROM UNNEST(amounts) AS x) escrowed_total
		FROM cz_cycle_escrowed
	) e ON e.round_id = ri.round_id
	LEFT JOIN (
		SELECT round_id, COUNT(*) num_claims, SUM(total/1e18) claimed_total
		FROM cz_claim GROUP BY round_id
	) c ON c.round_id = ri.round_id
	%s
	ORDER BY ri.round_id DESC`

func scanRoundSummary(rows pgx.Rows, rec *czmodel.ApiRoundSummary) error {
	return rows.Scan(
		&rec.RoundId, &rec.FirstSeenTs, store.TimeText(&rec.FirstSeenDate),
		&rec.NumFundings, &rec.FundedTotal,
		&rec.NumTrades, &rec.Volume, &rec.NumTraders,
		&rec.Settled, &rec.ValuesMask, &rec.Carried,
		&rec.EscrowedTotal,
		&rec.NumClaims, &rec.ClaimedTotal,
	)
}

// RoundsList returns the combined per-round summaries, newest round first.
func (r *Repo) RoundsList(ctx context.Context, offset, limit int) ([]czmodel.ApiRoundSummary, error) {
	query := fmt.Sprintf(roundSummaryQuery, "") + " OFFSET $1 LIMIT $2"
	return queryList(ctx, r, "chaoszero rounds list", 16, query, scanRoundSummary, offset, limit)
}

// RoundDetail returns the combined single-round view: the summary row plus
// the round's fundings, settlement, escrow, answer observations, claims and
// vault payouts. A round the indexer never saw yields store.ErrNotFound.
func (r *Repo) RoundDetail(ctx context.Context, roundID int64) (czmodel.ApiRoundDetail, error) {
	var out czmodel.ApiRoundDetail
	summaries, err := queryList(ctx, r, "chaoszero round detail", 1,
		fmt.Sprintf(roundSummaryQuery, "WHERE ri.round_id = $1"), scanRoundSummary, roundID)
	if err != nil {
		return out, err
	}
	if len(summaries) == 0 {
		return out, store.ErrNotFound
	}
	out.Summary = summaries[0]
	if out.Fundings, err = r.FundingByRound(ctx, roundID); err != nil {
		return out, err
	}
	settlements, err := r.settlements(ctx, "WHERE t.round_id = $1", 1, roundID)
	if err != nil {
		return out, err
	}
	if len(settlements) > 0 {
		out.Settlement = &settlements[0]
	}
	escrows, err := r.escrows(ctx, "WHERE t.round_id = $1", 1, roundID)
	if err != nil {
		return out, err
	}
	if len(escrows) > 0 {
		out.Escrow = &escrows[0]
	}
	if out.Answers, err = r.answers(ctx, "WHERE t.round_id = $1", 4, roundID); err != nil {
		return out, err
	}
	if out.Claims, err = r.claims(ctx, "WHERE t.round_id = $1 ORDER BY t.id", 8, roundID); err != nil {
		return out, err
	}
	if out.Payouts, err = r.payouts(ctx, "WHERE t.round_id = $1 ORDER BY t.id", 8, roundID); err != nil {
		return out, err
	}
	return out, nil
}

// =============================================================================
// PER-TABLE LISTS
// =============================================================================

const tradeSelect = `SELECT ` + eventRowColumns + `,
		t.round_id, ua.addr, t.question, t.yes, t.buy,
		(t.cash/1e18)::FLOAT8, (t.units/1e18)::FLOAT8
	FROM cz_trade t JOIN address ua ON t.user_aid = ua.address_id `

func scanTrade(rows pgx.Rows, rec *czmodel.ApiTrade) error {
	return scanEventRow(rows, &rec.EventRow,
		&rec.RoundId, &rec.UserAddr, &rec.Question, &rec.Yes, &rec.Buy,
		&rec.Cash, &rec.Units)
}

// TradesList returns trades newest-first.
func (r *Repo) TradesList(ctx context.Context, offset, limit int) ([]czmodel.ApiTrade, error) {
	query := tradeSelect + "ORDER BY t.id DESC OFFSET $1 LIMIT $2"
	return queryList(ctx, r, "chaoszero trades list", 32, query, scanTrade, offset, limit)
}

// TradesByRound returns one round's trades in event order.
func (r *Repo) TradesByRound(ctx context.Context, roundID int64, offset, limit int) ([]czmodel.ApiTrade, error) {
	query := tradeSelect + "WHERE t.round_id = $1 ORDER BY t.id OFFSET $2 LIMIT $3"
	return queryList(ctx, r, "chaoszero trades by round", 32, query, scanTrade, roundID, offset, limit)
}

// TradesByUser returns one user's trades newest-first.
func (r *Repo) TradesByUser(ctx context.Context, userAid int64, offset, limit int) ([]czmodel.ApiTrade, error) {
	query := tradeSelect + "WHERE t.user_aid = $1 ORDER BY t.id DESC OFFSET $2 LIMIT $3"
	return queryList(ctx, r, "chaoszero trades by user", 32, query, scanTrade, userAid, offset, limit)
}

const fundingSelect = `SELECT ` + eventRowColumns + `,
		t.round_id, ca.addr,
		(t.donated/1e18)::FLOAT8, (t.transfers_used/1e18)::FLOAT8, (t.carry_used/1e18)::FLOAT8,
		` + `ARRAY(SELECT (x/1e18)::FLOAT8 FROM UNNEST(t.allocations) AS x)` + `
	FROM cz_round_funded t JOIN address ca ON t.caller_aid = ca.address_id `

func scanFunding(rows pgx.Rows, rec *czmodel.ApiRoundFunded) error {
	return scanEventRow(rows, &rec.EventRow,
		&rec.RoundId, &rec.CallerAddr,
		&rec.Donated, &rec.TransfersUsed, &rec.CarryUsed, &rec.Allocations)
}

// FundingList returns round fundings newest-first.
func (r *Repo) FundingList(ctx context.Context, offset, limit int) ([]czmodel.ApiRoundFunded, error) {
	query := fundingSelect + "ORDER BY t.id DESC OFFSET $1 LIMIT $2"
	return queryList(ctx, r, "chaoszero funding list", 16, query, scanFunding, offset, limit)
}

// FundingByRound returns one round's fundings in event order.
func (r *Repo) FundingByRound(ctx context.Context, roundID int64) ([]czmodel.ApiRoundFunded, error) {
	query := fundingSelect + "WHERE t.round_id = $1 ORDER BY t.id"
	return queryList(ctx, r, "chaoszero funding by round", 4, query, scanFunding, roundID)
}

// SplitsList returns position splits newest-first.
func (r *Repo) SplitsList(ctx context.Context, offset, limit int) ([]czmodel.ApiPositionSplit, error) {
	query := `SELECT ` + eventRowColumns + `,
			t.round_id, ua.addr, t.question, (t.amount/1e18)::FLOAT8
		FROM cz_position_split t JOIN address ua ON t.user_aid = ua.address_id
		ORDER BY t.id DESC OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "chaoszero splits list", 16, query,
		func(rows pgx.Rows, rec *czmodel.ApiPositionSplit) error {
			return scanEventRow(rows, &rec.EventRow,
				&rec.RoundId, &rec.UserAddr, &rec.Question, &rec.Amount)
		}, offset, limit)
}

// MergesList returns position merges newest-first.
func (r *Repo) MergesList(ctx context.Context, offset, limit int) ([]czmodel.ApiPositionsMerged, error) {
	query := `SELECT ` + eventRowColumns + `,
			t.round_id, ua.addr,
			ARRAY(SELECT (x/1e18)::FLOAT8 FROM UNNEST(t.amounts) AS x),
			(t.collateral_out/1e18)::FLOAT8
		FROM cz_positions_merged t JOIN address ua ON t.user_aid = ua.address_id
		ORDER BY t.id DESC OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "chaoszero merges list", 16, query,
		func(rows pgx.Rows, rec *czmodel.ApiPositionsMerged) error {
			return scanEventRow(rows, &rec.EventRow,
				&rec.RoundId, &rec.UserAddr, &rec.Amounts, &rec.CollateralOut)
		}, offset, limit)
}

func (r *Repo) settlements(ctx context.Context, clause string, capHint int, args ...any) ([]czmodel.ApiRoundSettled, error) {
	query := `SELECT ` + eventRowColumns + `,
			t.round_id, t.values_mask, t.answers,
			ARRAY(SELECT (x/1e18)::FLOAT8 FROM UNNEST(t.escrowed) AS x),
			(t.carried/1e18)::FLOAT8
		FROM cz_round_settled t ` + clause
	return queryList(ctx, r, "chaoszero settlements", capHint, query,
		func(rows pgx.Rows, rec *czmodel.ApiRoundSettled) error {
			return scanEventRow(rows, &rec.EventRow,
				&rec.RoundId, &rec.ValuesMask, &rec.Answers, &rec.Escrowed, &rec.Carried)
		}, args...)
}

// SettlementsList returns round settlements newest-first.
func (r *Repo) SettlementsList(ctx context.Context, offset, limit int) ([]czmodel.ApiRoundSettled, error) {
	return r.settlements(ctx, "ORDER BY t.id DESC OFFSET $1 LIMIT $2", 16, offset, limit)
}

func (r *Repo) claims(ctx context.Context, clause string, capHint int, args ...any) ([]czmodel.ApiClaim, error) {
	query := `SELECT ` + eventRowColumns + `,
			t.round_id, ua.addr,
			ARRAY(SELECT (x/1e18)::FLOAT8 FROM UNNEST(t.amounts) AS x),
			(t.total/1e18)::FLOAT8
		FROM cz_claim t JOIN address ua ON t.user_aid = ua.address_id ` + clause
	return queryList(ctx, r, "chaoszero claims", capHint, query,
		func(rows pgx.Rows, rec *czmodel.ApiClaim) error {
			return scanEventRow(rows, &rec.EventRow,
				&rec.RoundId, &rec.UserAddr, &rec.Amounts, &rec.Total)
		}, args...)
}

// ClaimsList returns claims newest-first.
func (r *Repo) ClaimsList(ctx context.Context, offset, limit int) ([]czmodel.ApiClaim, error) {
	return r.claims(ctx, "ORDER BY t.id DESC OFFSET $1 LIMIT $2", 16, offset, limit)
}

// ClaimsByUser returns one user's claims newest-first.
func (r *Repo) ClaimsByUser(ctx context.Context, userAid int64, offset, limit int) ([]czmodel.ApiClaim, error) {
	return r.claims(ctx, "WHERE t.user_aid = $1 ORDER BY t.id DESC OFFSET $2 LIMIT $3", 16, userAid, offset, limit)
}

func (r *Repo) answers(ctx context.Context, clause string, capHint int, args ...any) ([]czmodel.ApiAnswersObserved, error) {
	query := `SELECT ` + eventRowColumns + `,
			t.round_id, t.known_mask, t.values_mask
		FROM cz_answers_observed t ` + clause
	return queryList(ctx, r, "chaoszero answers", capHint, query,
		func(rows pgx.Rows, rec *czmodel.ApiAnswersObserved) error {
			return scanEventRow(rows, &rec.EventRow,
				&rec.RoundId, &rec.KnownMask, &rec.ValuesMask)
		}, args...)
}

// AnswersList returns answer observations newest-first.
func (r *Repo) AnswersList(ctx context.Context, offset, limit int) ([]czmodel.ApiAnswersObserved, error) {
	return r.answers(ctx, "ORDER BY t.id DESC OFFSET $1 LIMIT $2", 16, offset, limit)
}

func (r *Repo) escrows(ctx context.Context, clause string, capHint int, args ...any) ([]czmodel.ApiCycleEscrowed, error) {
	query := `SELECT ` + eventRowColumns + `,
			t.round_id,
			ARRAY(SELECT (x/1e18)::FLOAT8 FROM UNNEST(t.amounts) AS x)
		FROM cz_cycle_escrowed t ` + clause
	return queryList(ctx, r, "chaoszero escrows", capHint, query,
		func(rows pgx.Rows, rec *czmodel.ApiCycleEscrowed) error {
			return scanEventRow(rows, &rec.EventRow, &rec.RoundId, &rec.Amounts)
		}, args...)
}

// EscrowsList returns vault escrows newest-first.
func (r *Repo) EscrowsList(ctx context.Context, offset, limit int) ([]czmodel.ApiCycleEscrowed, error) {
	return r.escrows(ctx, "ORDER BY t.id DESC OFFSET $1 LIMIT $2", 16, offset, limit)
}

func (r *Repo) payouts(ctx context.Context, clause string, capHint int, args ...any) ([]czmodel.ApiPaid, error) {
	query := `SELECT ` + eventRowColumns + `,
			t.round_id, t.question, ra.addr, (t.amount/1e18)::FLOAT8
		FROM cz_paid t JOIN address ra ON t.recipient_aid = ra.address_id ` + clause
	return queryList(ctx, r, "chaoszero payouts", capHint, query,
		func(rows pgx.Rows, rec *czmodel.ApiPaid) error {
			return scanEventRow(rows, &rec.EventRow,
				&rec.RoundId, &rec.Question, &rec.RecipientAddr, &rec.Amount)
		}, args...)
}

// PayoutsList returns vault payouts newest-first.
func (r *Repo) PayoutsList(ctx context.Context, offset, limit int) ([]czmodel.ApiPaid, error) {
	return r.payouts(ctx, "ORDER BY t.id DESC OFFSET $1 LIMIT $2", 16, offset, limit)
}

const transferSelect = `SELECT ` + eventRowColumns + `,
		t.batch_index, oa.addr, fa.addr, ta.addr,
		t.position_id::TEXT, t.round_id, t.question, t.yes, (t.amount/1e18)::FLOAT8
	FROM cz_position_transfer t
		JOIN address oa ON t.operator_aid = oa.address_id
		JOIN address fa ON t.from_aid = fa.address_id
		JOIN address ta ON t.to_aid = ta.address_id `

func scanTransfer(rows pgx.Rows, rec *czmodel.ApiPositionTransfer) error {
	return scanEventRow(rows, &rec.EventRow,
		&rec.BatchIndex, &rec.OperatorAddr, &rec.FromAddr, &rec.ToAddr,
		&rec.PositionId, &rec.RoundId, &rec.Question, &rec.Yes, &rec.Amount)
}

// TransfersList returns position transfers newest-first.
func (r *Repo) TransfersList(ctx context.Context, offset, limit int) ([]czmodel.ApiPositionTransfer, error) {
	query := transferSelect + "ORDER BY t.id DESC OFFSET $1 LIMIT $2"
	return queryList(ctx, r, "chaoszero transfers list", 32, query, scanTransfer, offset, limit)
}

// TransfersByUser returns transfers where the user is sender or receiver,
// newest-first.
func (r *Repo) TransfersByUser(ctx context.Context, userAid int64, offset, limit int) ([]czmodel.ApiPositionTransfer, error) {
	query := transferSelect + "WHERE t.from_aid = $1 OR t.to_aid = $1 ORDER BY t.id DESC OFFSET $2 LIMIT $3"
	return queryList(ctx, r, "chaoszero transfers by user", 32, query, scanTransfer, userAid, offset, limit)
}

// ApprovalsList returns operator approvals newest-first.
func (r *Repo) ApprovalsList(ctx context.Context, offset, limit int) ([]czmodel.ApiApprovalForAll, error) {
	query := `SELECT ` + eventRowColumns + `,
			oa.addr, pa.addr, t.approved
		FROM cz_approval_for_all t
			JOIN address oa ON t.owner_aid = oa.address_id
			JOIN address pa ON t.operator_aid = pa.address_id
		ORDER BY t.id DESC OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "chaoszero approvals list", 16, query,
		func(rows pgx.Rows, rec *czmodel.ApiApprovalForAll) error {
			return scanEventRow(rows, &rec.EventRow,
				&rec.OwnerAddr, &rec.OperatorAddr, &rec.Approved)
		}, offset, limit)
}
