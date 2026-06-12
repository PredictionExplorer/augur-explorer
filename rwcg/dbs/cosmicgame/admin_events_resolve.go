package cosmicgame

import (
	"database/sql"
	"fmt"
	"math"
	"math/big"

	p "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/cosmicgame"
)

// ethDutchAuctionDurationNumerator matches the on-chain ETH Dutch auction duration formula:
// durationSeconds = roundActivationSpanSeconds * numerator / ethDutchAuctionDurationDivisor.
const ethDutchAuctionDurationNumerator int64 = 11614

// initialDurationRawNumerator matches GetInitialDurationUntilMainPrize raw numerator:
// bumpSeconds = mainPrizeTimeIncrementMicroSeconds * numerator / 1000 / initialDurationUntilMainPrizeDivisor.
const initialDurationRawNumerator int64 = 1071

func (sw *SQLStorageWrapper) Resolve_admin_event_values(events []p.CGAdminEvent) {
	for i := range events {
		events[i].ResolvedValue = sw.resolveAdminEventValueSQL(&events[i])
	}
}

func (sw *SQLStorageWrapper) resolveAdminEventValueSQL(rec *p.CGAdminEvent) string {
	if rec == nil || rec.IntegerValue == 0 && rec.FloatValue == 0 && rec.StringValue == "" {
		return ""
	}
	switch rec.RecordType {
	case 7, 19, 35:
		return formatDurationSeconds(rec.IntegerValue)
	case 18:
		return sw.resolveTimeIncreaseChanged(rec)
	case 20:
		return formatPercentFromDivisor(rec.IntegerValue)
	case 21:
		return formatMicrosecondsAsSeconds(rec.IntegerValue)
	case 22:
		return sw.resolveInitialSecondsUntilPrizeChanged(rec)
	case 25:
		return formatDurationSeconds(rec.IntegerValue)
	case 36:
		return sw.resolveEthDutchAuctionDurationDivisor(rec)
	case 37:
		return sw.resolveEthDutchAuctionEndingBidPriceDivisor(rec)
	case 39:
		return sw.resolveCstDutchAuctionDurationChangeDivisor(rec)
	default:
		return ""
	}
}

func (sw *SQLStorageWrapper) resolveTimeIncreaseChanged(rec *p.CGAdminEvent) string {
	if rec.IntegerValue <= 0 {
		return ""
	}
	baseMicros, ok := sw.getLatestInt64ParamAtEvtlog("cg_adm_prize_microsec", "new_microseconds", rec.EvtLogId, true)
	if !ok || baseMicros <= 0 {
		return fmt.Sprintf("×%.2f per bid", float64(rec.IntegerValue)/100.0)
	}
	afterMicros := baseMicros * rec.IntegerValue / 100
	return fmt.Sprintf("%s after next bid (×%.2f)", formatMicrosecondsAsSeconds(afterMicros), float64(rec.IntegerValue)/100.0)
}

func (sw *SQLStorageWrapper) resolveInitialSecondsUntilPrizeChanged(rec *p.CGAdminEvent) string {
	if rec.IntegerValue <= 0 {
		return ""
	}
	pct := formatPercentFromDivisor(rec.IntegerValue)
	baseMicros, ok := sw.getLatestInt64ParamAtEvtlog("cg_adm_prize_microsec", "new_microseconds", rec.EvtLogId, true)
	if !ok || baseMicros <= 0 {
		return pct
	}
	raw := baseMicros * initialDurationRawNumerator / 1000
	secs := raw / rec.IntegerValue
	return fmt.Sprintf("%s (%s)", formatDurationSeconds(secs), pct)
}

func (sw *SQLStorageWrapper) resolveEthDutchAuctionDurationDivisor(rec *p.CGAdminEvent) string {
	if rec.IntegerValue <= 0 {
		return ""
	}
	span, ok := sw.getLatestActivationSpanSeconds(rec.EvtLogId)
	if !ok || span <= 0 {
		return ""
	}
	secs := span * ethDutchAuctionDurationNumerator / rec.IntegerValue
	return formatDurationSeconds(secs)
}

func (sw *SQLStorageWrapper) resolveEthDutchAuctionEndingBidPriceDivisor(rec *p.CGAdminEvent) string {
	if rec.IntegerValue <= 0 {
		return ""
	}
	beginWei, ok := sw.getLatestEthBidPriceWeiBeforeEvtlog(rec.EvtLogId)
	if !ok || beginWei <= 0 {
		return ""
	}
	endWei := beginWei / rec.IntegerValue
	return formatEthFromWei(endWei)
}

func (sw *SQLStorageWrapper) resolveCstDutchAuctionDurationChangeDivisor(rec *p.CGAdminEvent) string {
	if rec.IntegerValue <= 0 {
		return ""
	}
	duration, ok := sw.getLatestInt64ParamAtEvtlog("cg_adm_cst_auclen", "new_len", rec.EvtLogId, true)
	if !ok || duration <= 0 {
		return ""
	}
	delta := duration / rec.IntegerValue
	if delta <= 0 {
		return ""
	}
	return fmt.Sprintf("%s change per bid", formatDurationSeconds(delta))
}

func (sw *SQLStorageWrapper) getLatestInt64ParamAtEvtlog(table, column string, evtlogID int64, inclusive bool) (int64, bool) {
	op := "<"
	if inclusive {
		op = "<="
	}
	query := fmt.Sprintf(
		"SELECT %s FROM %s.%s WHERE evtlog_id %s $1 ORDER BY evtlog_id DESC, id DESC LIMIT 1",
		column, sw.S.SchemaName(), table, op,
	)
	var val sql.NullString
	err := sw.S.Db().QueryRow(query, evtlogID).Scan(&val)
	if err == sql.ErrNoRows || err != nil || !val.Valid || val.String == "" {
		return 0, false
	}
	bi, ok := new(big.Int).SetString(val.String, 10)
	if !ok {
		return 0, false
	}
	return bi.Int64(), true
}

func (sw *SQLStorageWrapper) getLatestActivationSpanSeconds(evtlogID int64) (int64, bool) {
	query := fmt.Sprintf(
		"SELECT new_atime FROM %s.cg_adm_acttime WHERE evtlog_id <= $1 ORDER BY evtlog_id DESC, id DESC LIMIT 2",
		sw.S.SchemaName(),
	)
	rows, err := sw.S.Db().Query(query, evtlogID)
	if err != nil {
		return 0, false
	}
	defer rows.Close()
	var times []int64
	for rows.Next() {
		var raw sql.NullString
		if err := rows.Scan(&raw); err != nil || !raw.Valid {
			continue
		}
		bi, ok := new(big.Int).SetString(raw.String, 10)
		if !ok {
			continue
		}
		times = append(times, bi.Int64())
	}
	if len(times) < 2 {
		return 0, false
	}
	span := times[0] - times[1]
	if span <= 0 {
		return 0, false
	}
	return span, true
}

func (sw *SQLStorageWrapper) getLatestEthBidPriceWeiBeforeEvtlog(evtlogID int64) (int64, bool) {
	query := fmt.Sprintf(
		"SELECT eth_price FROM %s.cg_bid WHERE evtlog_id < $1 AND eth_price > 0 ORDER BY evtlog_id DESC LIMIT 1",
		sw.S.SchemaName(),
	)
	var val sql.NullString
	err := sw.S.Db().QueryRow(query, evtlogID).Scan(&val)
	if err == sql.ErrNoRows || err != nil || !val.Valid || val.String == "" {
		return 0, false
	}
	bi, ok := new(big.Int).SetString(val.String, 10)
	if !ok {
		return 0, false
	}
	return bi.Int64(), true
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
