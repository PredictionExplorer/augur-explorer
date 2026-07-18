package cosmicgame

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// ethDutchAuctionDurationNumerator matches the on-chain ETH Dutch auction duration formula:
// durationSeconds = roundActivationSpanSeconds * numerator / ethDutchAuctionDurationDivisor.
const ethDutchAuctionDurationNumerator int64 = 11614

// initialDurationRawNumerator matches GetInitialDurationUntilMainPrize raw numerator:
// bumpSeconds = mainPrizeTimeIncrementMicroSeconds * numerator / 1000 / initialDurationUntilMainPrizeDivisor.
const initialDurationRawNumerator int64 = 1071

// ResolveAdminEventValues fills the ResolvedValue of each admin event with a
// human-readable rendering of its parameter change (durations, percentages,
// derived auction values). Events whose type has no resolver, or whose
// context rows are missing, keep an empty ResolvedValue — only real database
// failures return an error.
func (r *Repo) ResolveAdminEventValues(ctx context.Context, events []cgmodel.CGAdminEvent) error {
	for i := range events {
		value, err := r.resolveAdminEventValue(ctx, &events[i])
		if err != nil {
			return err
		}
		events[i].ResolvedValue = value
	}
	return nil
}

func (r *Repo) resolveAdminEventValue(ctx context.Context, rec *cgmodel.CGAdminEvent) (string, error) {
	if rec == nil || rec.IntegerValue == 0 && rec.FloatValue == 0 && rec.StringValue == "" {
		return "", nil
	}
	switch rec.RecordType {
	case 7, 19, 35:
		return formatDurationSeconds(rec.IntegerValue), nil
	case 18:
		return r.resolveTimeIncreaseChanged(ctx, rec)
	case 20:
		return formatPercentFromDivisor(rec.IntegerValue), nil
	case 21:
		return formatMicrosecondsAsSeconds(rec.IntegerValue), nil
	case 22:
		return r.resolveInitialSecondsUntilPrizeChanged(ctx, rec)
	case 25:
		return formatDurationSeconds(rec.IntegerValue), nil
	case 36:
		return r.resolveEthDutchAuctionDurationDivisor(ctx, rec)
	case 37:
		return r.resolveEthDutchAuctionEndingBidPriceDivisor(ctx, rec)
	case 39:
		return r.resolveCstDutchAuctionDurationChangeDivisor(ctx, rec)
	default:
		return "", nil
	}
}

func (r *Repo) resolveTimeIncreaseChanged(ctx context.Context, rec *cgmodel.CGAdminEvent) (string, error) {
	if rec.IntegerValue <= 0 {
		return "", nil
	}
	baseMicros, ok, err := r.latestInt64ParamAtEvtlog(ctx, "cg_adm_prize_microsec", "new_microseconds", rec.EvtLogId, true)
	if err != nil {
		return "", err
	}
	if !ok || baseMicros <= 0 {
		return fmt.Sprintf("×%.2f per bid", float64(rec.IntegerValue)/100.0), nil
	}
	afterMicros := baseMicros * rec.IntegerValue / 100
	return fmt.Sprintf("%s after next bid (×%.2f)", formatMicrosecondsAsSeconds(afterMicros), float64(rec.IntegerValue)/100.0), nil
}

func (r *Repo) resolveInitialSecondsUntilPrizeChanged(ctx context.Context, rec *cgmodel.CGAdminEvent) (string, error) {
	if rec.IntegerValue <= 0 {
		return "", nil
	}
	pct := formatPercentFromDivisor(rec.IntegerValue)
	baseMicros, ok, err := r.latestInt64ParamAtEvtlog(ctx, "cg_adm_prize_microsec", "new_microseconds", rec.EvtLogId, true)
	if err != nil {
		return "", err
	}
	if !ok || baseMicros <= 0 {
		return pct, nil
	}
	raw := baseMicros * initialDurationRawNumerator / 1000
	secs := raw / rec.IntegerValue
	return fmt.Sprintf("%s (%s)", formatDurationSeconds(secs), pct), nil
}

func (r *Repo) resolveEthDutchAuctionDurationDivisor(ctx context.Context, rec *cgmodel.CGAdminEvent) (string, error) {
	if rec.IntegerValue <= 0 {
		return "", nil
	}
	span, ok, err := r.latestActivationSpanSeconds(ctx, rec.EvtLogId)
	if err != nil {
		return "", err
	}
	if !ok || span <= 0 {
		return "", nil
	}
	secs := span * ethDutchAuctionDurationNumerator / rec.IntegerValue
	return formatDurationSeconds(secs), nil
}

func (r *Repo) resolveEthDutchAuctionEndingBidPriceDivisor(ctx context.Context, rec *cgmodel.CGAdminEvent) (string, error) {
	if rec.IntegerValue <= 0 {
		return "", nil
	}
	beginWei, ok, err := r.latestEthBidPriceWeiBeforeEvtlog(ctx, rec.EvtLogId)
	if err != nil {
		return "", err
	}
	if !ok || beginWei <= 0 {
		return "", nil
	}
	endWei := beginWei / rec.IntegerValue
	return formatEthFromWei(endWei), nil
}

func (r *Repo) resolveCstDutchAuctionDurationChangeDivisor(ctx context.Context, rec *cgmodel.CGAdminEvent) (string, error) {
	if rec.IntegerValue <= 0 {
		return "", nil
	}
	duration, ok, err := r.latestInt64ParamAtEvtlog(ctx, "cg_adm_cst_auclen", "new_len", rec.EvtLogId, true)
	if err != nil {
		return "", err
	}
	if !ok || duration <= 0 {
		return "", nil
	}
	delta := duration / rec.IntegerValue
	if delta <= 0 {
		return "", nil
	}
	return fmt.Sprintf("%s change per bid", formatDurationSeconds(delta)), nil
}

// latestInt64ParamAtEvtlog returns the most recent value of an admin
// parameter column at (or, when inclusive, up to and including) evtlogID.
// ok is false when no usable row exists. The identifiers pass the same guard
// as the contract-params queries; call sites use compile-time literals.
func (r *Repo) latestInt64ParamAtEvtlog(ctx context.Context, table, column string, evtlogID int64, inclusive bool) (int64, bool, error) {
	if err := checkAdminIdent("table", table); err != nil {
		return 0, false, err
	}
	if err := checkAdminIdent("column", column); err != nil {
		return 0, false, err
	}
	cmp := "<"
	if inclusive {
		cmp = "<="
	}
	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE evtlog_id %s $1 ORDER BY evtlog_id DESC, id DESC LIMIT 1",
		column, table, cmp,
	)
	var val sql.NullString
	err := r.q(ctx).QueryRow(ctx, query, evtlogID).Scan(&val)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, store.WrapError("latest param "+table+"."+column, err)
	}
	if !val.Valid || val.String == "" {
		return 0, false, nil
	}
	bi, ok := new(big.Int).SetString(val.String, 10)
	if !ok {
		return 0, false, nil
	}
	return bi.Int64(), true, nil
}

// latestActivationSpanSeconds derives the span between the two most recent
// activation-time settings at evtlogID (the round length used by the ETH
// Dutch auction resolvers). ok is false with fewer than two usable rows.
func (r *Repo) latestActivationSpanSeconds(ctx context.Context, evtlogID int64) (int64, bool, error) {
	const op = "latest activation span seconds"
	query := "SELECT new_atime FROM cg_adm_acttime WHERE evtlog_id <= $1 ORDER BY evtlog_id DESC, id DESC LIMIT 2"
	rows, err := r.q(ctx).Query(ctx, query, evtlogID)
	if err != nil {
		return 0, false, store.WrapError(op, err)
	}
	defer rows.Close()
	var times []int64
	for rows.Next() {
		var raw sql.NullString
		if err := rows.Scan(&raw); err != nil {
			return 0, false, store.WrapError(op, err)
		}
		if !raw.Valid {
			continue
		}
		bi, ok := new(big.Int).SetString(raw.String, 10)
		if !ok {
			continue
		}
		times = append(times, bi.Int64())
	}
	if err := rows.Err(); err != nil {
		return 0, false, store.WrapError(op, err)
	}
	if len(times) < 2 {
		return 0, false, nil
	}
	span := times[0] - times[1]
	if span <= 0 {
		return 0, false, nil
	}
	return span, true, nil
}

// latestEthBidPriceWeiBeforeEvtlog returns the last positive ETH bid price
// preceding evtlogID; ok is false when no such bid exists.
func (r *Repo) latestEthBidPriceWeiBeforeEvtlog(ctx context.Context, evtlogID int64) (int64, bool, error) {
	query := "SELECT eth_price FROM cg_bid WHERE evtlog_id < $1 AND eth_price > 0 ORDER BY evtlog_id DESC LIMIT 1"
	var val sql.NullString
	err := r.q(ctx).QueryRow(ctx, query, evtlogID).Scan(&val)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, store.WrapError("latest eth bid price before evtlog", err)
	}
	if !val.Valid || val.String == "" {
		return 0, false, nil
	}
	bi, ok := new(big.Int).SetString(val.String, 10)
	if !ok {
		return 0, false, nil
	}
	return bi.Int64(), true, nil
}

func formatPercentFromDivisor(divisor int64) string {
	if divisor <= 0 {
		return ""
	}
	pct := 100.0 / float64(divisor)
	return fmt.Sprintf("%.4f%%", pct)
}

func formatMicrosecondsAsSeconds(micros int64) string {
	if micros <= 0 {
		return "0 sec"
	}
	whole := micros / 1_000_000
	frac := micros % 1_000_000
	if frac == 0 {
		return fmt.Sprintf("%d sec", whole)
	}
	return fmt.Sprintf("%.6f sec", float64(micros)/1_000_000.0)
}

func formatDurationSeconds(secs int64) string {
	if secs <= 0 {
		return "0 sec"
	}
	if secs < 120 {
		return fmt.Sprintf("%d sec", secs)
	}
	hours := secs / 3600
	mins := (secs % 3600) / 60
	rem := secs % 60
	if hours > 0 {
		if rem == 0 && mins == 0 {
			return fmt.Sprintf("%dh", hours)
		}
		if rem == 0 {
			return fmt.Sprintf("%dh %dm", hours, mins)
		}
		return fmt.Sprintf("%dh %dm %ds", hours, mins, rem)
	}
	if rem == 0 {
		return fmt.Sprintf("%dm", mins)
	}
	return fmt.Sprintf("%dm %ds", mins, rem)
}

func formatEthFromWei(wei int64) string {
	if wei <= 0 {
		return "0 ETH"
	}
	f := new(big.Float).SetInt(big.NewInt(wei))
	f.Quo(f, big.NewFloat(1e18))
	val, _ := f.Float64()
	if math.Abs(val) >= 0.0001 {
		return fmt.Sprintf("%.8f ETH", val)
	}
	return fmt.Sprintf("%.12f ETH", val)
}
