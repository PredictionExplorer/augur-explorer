package cosmicgame

import (
	"context"
	"math"
	"sort"
	"strconv"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

const roundOpenExcludeSecs = 3600

// excludeFirstHourAfterRoundStartSQL filters out bids placed in the first
// hour after each cycle's round_start_time (FirstBidPlacedInRound), so the
// frequency series measures organic activity rather than round-open rushes.
func excludeFirstHourAfterRoundStartSQL() string {
	return " AND NOT EXISTS (" +
		"SELECT 1 FROM cg_round_stats rs " +
		"WHERE rs.round_num = b.round_num " +
		"AND rs.round_start_time IS NOT NULL " +
		"AND b.time_stamp >= rs.round_start_time " +
		"AND b.time_stamp < rs.round_start_time + INTERVAL '" + strconv.Itoa(roundOpenExcludeSecs) + " seconds'" +
		")"
}

// bidFrequencySQL returns the bucket query for one sampling interval and
// reports whether the buckets are epoch-aligned. Hour and day windows align
// to UTC epoch boundaries so counts match regardless of the lookback window;
// every other interval anchors the series at initTs. The interval value is
// an integer rendered with strconv.Itoa, so the interpolation cannot carry
// SQL. Epoch-aligned queries bind (initTs, finTs); anchored queries bind
// (initTs, finTs, interval-as-text) — the third parameter feeds a text
// concatenation ("$3 || ' seconds'") and must be a string because pgx infers
// the placeholder as text.
func bidFrequencySQL(intervalSecs int) (query string, epochAligned bool) {
	intervalStr := strconv.Itoa(intervalSecs)
	excl := excludeFirstHourAfterRoundStartSQL()

	if intervalSecs == 3600 || intervalSecs == 86400 {
		return "WITH periods AS (" +
			"SELECT generate_series AS start_ts, " +
			"generate_series + (" + intervalStr + " || ' seconds')::interval AS end_ts " +
			"FROM generate_series(" +
			"TO_TIMESTAMP((FLOOR($1::float / " + intervalStr + ") * " + intervalStr + ")::bigint), " +
			"TO_TIMESTAMP((FLOOR($2::float / " + intervalStr + ") * " + intervalStr + ")::bigint), " +
			"('" + intervalStr + " seconds')::interval" +
			") AS generate_series" +
			") " +
			"SELECT " +
			"FLOOR(EXTRACT(EPOCH FROM p.start_ts))::BIGINT AS bucket_ts, " +
			"COALESCE(COUNT(b.id), 0)::BIGINT AS num_bids, " +
			"COALESCE(COUNT(DISTINCT b.bidder_aid), 0)::BIGINT AS unique_bidders " +
			"FROM periods p " +
			"LEFT JOIN cg_bid b ON " +
			"FLOOR(EXTRACT(EPOCH FROM b.time_stamp)::bigint / " + intervalStr + ") * " + intervalStr +
			" = FLOOR(EXTRACT(EPOCH FROM p.start_ts)::bigint) " +
			"AND b.time_stamp >= TO_TIMESTAMP($1) AND b.time_stamp < TO_TIMESTAMP($2)" +
			excl + " " +
			"GROUP BY p.start_ts " +
			"ORDER BY p.start_ts", true
	}
	return "WITH periods AS (" +
		"SELECT generate_series AS start_ts, " +
		"generate_series + ($3 || ' seconds')::interval AS end_ts " +
		"FROM generate_series(" +
		"TO_TIMESTAMP($1), " +
		"TO_TIMESTAMP($2), " +
		"($3 || ' seconds')::interval" +
		") AS generate_series" +
		") " +
		"SELECT " +
		"FLOOR(EXTRACT(EPOCH FROM p.start_ts))::BIGINT AS bucket_ts, " +
		"COALESCE(COUNT(b.id), 0)::BIGINT AS num_bids, " +
		"COALESCE(COUNT(DISTINCT b.bidder_aid), 0)::BIGINT AS unique_bidders " +
		"FROM periods p " +
		"LEFT JOIN cg_bid b ON " +
		"b.time_stamp >= p.start_ts AND b.time_stamp < p.end_ts" +
		excl + " " +
		"GROUP BY p.start_ts " +
		"ORDER BY p.start_ts", false
}

// BidFrequencyByPeriod returns the bid-count series over [initTs, finTs) in
// buckets of intervalSecs seconds. Empty buckets are zero-filled so the
// caller can render a continuous series; bids in the first hour after a
// round opens are excluded (see excludeFirstHourAfterRoundStartSQL). A
// non-positive interval means one bucket spanning the whole range (falling
// back to an hour if the range itself is empty).
func (r *Repo) BidFrequencyByPeriod(ctx context.Context, initTs, finTs, intervalSecs int) ([]p.CGBidFrequencyBucket, error) {
	if intervalSecs <= 0 {
		intervalSecs = finTs - initTs
		if intervalSecs <= 0 {
			intervalSecs = 3600
		}
	}

	query, epochAligned := bidFrequencySQL(intervalSecs)
	args := []any{initTs, finTs}
	if !epochAligned {
		args = append(args, strconv.Itoa(intervalSecs))
	}
	scan := func(rows pgx.Rows, rec *p.CGBidFrequencyBucket) error {
		return rows.Scan(&rec.BucketTs, &rec.NumBids, &rec.UniqueBidders)
	}
	return queryList(ctx, r, "bid frequency by period", 256, query, scan, args...)
}

// BidTypeRatioByPeriod returns the bid-type composition per sampling window
// over [initTs, finTs). Each window reports raw counts per bid_type plus
// those counts normalized to a windowed 100% (per-interval mix, not
// cumulative). Windows with no bids report 0 counts and 0% for every type,
// so the caller can render a continuous, gap-free series.
// bid_type: 0=ETH, 1=RandomWalk, 2=CST.
func (r *Repo) BidTypeRatioByPeriod(ctx context.Context, initTs, finTs, intervalSecs int) ([]p.CGBidTypeRatioBucket, error) {
	if intervalSecs <= 0 {
		intervalSecs = finTs - initTs
		if intervalSecs <= 0 {
			intervalSecs = 86400
		}
	}

	intervalStr := strconv.Itoa(intervalSecs)
	query := "WITH periods AS (" +
		"SELECT generate_series AS start_ts, " +
		"generate_series + ('" + intervalStr + " seconds')::interval AS end_ts " +
		"FROM generate_series(" +
		"TO_TIMESTAMP($1), " +
		"TO_TIMESTAMP($2), " +
		"('" + intervalStr + " seconds')::interval" +
		") AS generate_series" +
		") " +
		"SELECT " +
		"FLOOR(EXTRACT(EPOCH FROM p.start_ts))::BIGINT AS bucket_ts, " +
		"COALESCE(COUNT(b.id) FILTER (WHERE b.bid_type = 0), 0)::BIGINT AS eth_bids, " +
		"COALESCE(COUNT(b.id) FILTER (WHERE b.bid_type = 1), 0)::BIGINT AS rwalk_bids, " +
		"COALESCE(COUNT(b.id) FILTER (WHERE b.bid_type = 2), 0)::BIGINT AS cst_bids, " +
		"COALESCE(COUNT(b.id), 0)::BIGINT AS total_bids " +
		"FROM periods p " +
		"LEFT JOIN cg_bid b ON " +
		"b.time_stamp >= p.start_ts AND b.time_stamp < p.end_ts " +
		"GROUP BY p.start_ts " +
		"ORDER BY p.start_ts"

	scan := func(rows pgx.Rows, rec *p.CGBidTypeRatioBucket) error {
		if err := rows.Scan(&rec.BucketTs, &rec.EthBids, &rec.RwalkBids, &rec.CstBids, &rec.TotalBids); err != nil {
			return err
		}
		if rec.TotalBids > 0 {
			total := float64(rec.TotalBids)
			rec.EthPct = math.Round(float64(rec.EthBids)/total*10000) / 100
			rec.RwalkPct = math.Round(float64(rec.RwalkBids)/total*10000) / 100
			rec.CstPct = math.Round(float64(rec.CstBids)/total*10000) / 100
		}
		return nil
	}
	return queryList(ctx, r, "bid type ratio by period", 256, query, scan, initTs, finTs)
}

// TopBidders returns the most active bidders by lifetime bid count,
// descending. A non-positive limit defaults to 20.
func (r *Repo) TopBidders(ctx context.Context, limit int) ([]p.CGTopBidderInfo, error) {
	if limit <= 0 {
		limit = 20
	}
	query := "SELECT b.bidder_aid, a.addr, b.num_bids " +
		"FROM cg_bidder b " +
		"LEFT JOIN address a ON b.bidder_aid = a.address_id " +
		"WHERE b.num_bids > 0 " +
		"ORDER BY b.num_bids DESC " +
		"LIMIT $1"
	scan := func(rows pgx.Rows, rec *p.CGTopBidderInfo) error {
		return rows.Scan(&rec.BidderAid, &rec.BidderAddr, &rec.NumBids)
	}
	return queryList(ctx, r, "top bidders", limit, query, scan, limit)
}

// TopBidderActivePeriods segments the bids of the topN bidders inside
// [initTs, finTs) into activity sessions: a gap longer than gapHours starts
// a new session, and sessions with fewer than minBids bids are dropped.
// Non-positive knobs default to topN=20, gapHours=6, minBids=2. When there
// are no bidders at all the period list is nil.
func (r *Repo) TopBidderActivePeriods(ctx context.Context, topN, initTs, finTs, gapHours, minBids int) ([]p.CGTopBidderInfo, []p.CGBidderActivePeriod, error) {
	if topN <= 0 {
		topN = 20
	}
	if gapHours <= 0 {
		gapHours = 6
	}
	if minBids <= 0 {
		minBids = 2
	}

	topBidders, err := r.TopBidders(ctx, topN)
	if err != nil {
		return nil, nil, err
	}
	if len(topBidders) == 0 {
		return topBidders, nil, nil
	}

	bidderAids := make([]int64, 0, len(topBidders))
	for _, b := range topBidders {
		bidderAids = append(bidderAids, b.BidderAid)
	}

	query := "WITH top_bidders AS (" +
		"SELECT UNNEST($1::bigint[]) AS bidder_aid" +
		"), ordered AS (" +
		"SELECT b.bidder_aid, a.addr, b.time_stamp, " +
		"b.time_stamp - LAG(b.time_stamp) OVER (PARTITION BY b.bidder_aid ORDER BY b.time_stamp) AS gap " +
		"FROM cg_bid b " +
		"JOIN address a ON a.address_id = b.bidder_aid " +
		"WHERE b.bidder_aid IN (SELECT bidder_aid FROM top_bidders) " +
		"AND b.time_stamp >= TO_TIMESTAMP($2) " +
		"AND b.time_stamp < TO_TIMESTAMP($3)" +
		"), sessions AS (" +
		"SELECT bidder_aid, addr, time_stamp, " +
		"SUM(CASE WHEN gap IS NULL OR gap > ($4 || ' hours')::interval THEN 1 ELSE 0 END) " +
		"OVER (PARTITION BY bidder_aid ORDER BY time_stamp) AS session_id " +
		"FROM ordered" +
		") " +
		"SELECT bidder_aid, addr, session_id, " +
		"FLOOR(EXTRACT(EPOCH FROM MIN(time_stamp)))::BIGINT, " +
		"FLOOR(EXTRACT(EPOCH FROM MAX(time_stamp)))::BIGINT, " +
		"COUNT(*)::BIGINT " +
		"FROM sessions " +
		"GROUP BY bidder_aid, addr, session_id " +
		"HAVING COUNT(*) >= $5 " +
		"ORDER BY MIN(time_stamp)"

	scan := func(rows pgx.Rows, rec *p.CGBidderActivePeriod) error {
		var sessionID int64
		if err := rows.Scan(
			&rec.BidderAid,
			&rec.BidderAddr,
			&sessionID,
			&rec.PeriodStart,
			&rec.PeriodEnd,
			&rec.NumBids,
		); err != nil {
			return err
		}
		rec.DurationSecs = rec.PeriodEnd - rec.PeriodStart
		if rec.DurationSecs < 0 {
			rec.DurationSecs = 0
		}
		return nil
	}
	// bidderAids is a native []int64, encoded by pgx as bigint[]. gapHours
	// feeds a text concatenation ("$4 || ' hours'") and must be a string for
	// the same reason as in bidFrequencySQL.
	periods, err := queryList(ctx, r, "top bidder active periods", 128, query, scan,
		bidderAids, initTs, finTs, strconv.Itoa(gapHours), minBids)
	if err != nil {
		return nil, nil, err
	}
	return topBidders, periods, nil
}

// BidTimeBounds returns the epoch timestamps of the first and last bid ever
// placed (0, 0 when there are no bids).
func (r *Repo) BidTimeBounds(ctx context.Context) (minTs, maxTs int64, err error) {
	query := "SELECT " +
		"COALESCE(FLOOR(EXTRACT(EPOCH FROM MIN(time_stamp))), 0)::BIGINT, " +
		"COALESCE(FLOOR(EXTRACT(EPOCH FROM MAX(time_stamp))), 0)::BIGINT " +
		"FROM cg_bid"
	if err := r.pool().QueryRow(ctx, query).Scan(&minTs, &maxTs); err != nil {
		return 0, 0, store.WrapError("bid time bounds", err)
	}
	return minTs, maxTs, nil
}

// DetectBidSpikes finds merged runs of buckets whose bid count exceeds a dynamic threshold.
func DetectBidSpikes(buckets []p.CGBidFrequencyBucket, intervalSecs int) []p.CGBidSpike {
	if len(buckets) == 0 {
		return nil
	}

	counts := make([]int64, len(buckets))
	var sum int64
	for i, b := range buckets {
		counts[i] = b.NumBids
		sum += b.NumBids
	}
	mean := float64(sum) / float64(len(buckets))

	var variance float64
	for _, c := range counts {
		d := float64(c) - mean
		variance += d * d
	}
	stddev := math.Sqrt(variance / float64(len(counts)))

	threshold := mean + 2*stddev
	if threshold < 5 {
		threshold = 5
	}
	if threshold < mean*2 && mean >= 2 {
		threshold = mean * 2
	}

	type spikeRun struct {
		startIdx int
		endIdx   int
	}
	var runs []spikeRun
	inRun := false
	var runStart int

	for i, c := range counts {
		if float64(c) >= threshold && c >= 3 {
			if !inRun {
				inRun = true
				runStart = i
			}
			continue
		}
		if inRun {
			runs = append(runs, spikeRun{startIdx: runStart, endIdx: i - 1})
			inRun = false
		}
	}
	if inRun {
		runs = append(runs, spikeRun{startIdx: runStart, endIdx: len(counts) - 1})
	}

	spikes := make([]p.CGBidSpike, 0, len(runs))
	for idx, run := range runs {
		spike := p.CGBidSpike{Index: idx}
		spike.StartTs = buckets[run.startIdx].BucketTs
		lastBucket := buckets[run.endIdx]
		spike.EndTs = lastBucket.BucketTs + int64(intervalSecs)

		var peakIdx int
		var peakBids int64
		for i := run.startIdx; i <= run.endIdx; i++ {
			spike.TotalBids += buckets[i].NumBids
			if buckets[i].NumBids >= peakBids {
				peakBids = buckets[i].NumBids
				peakIdx = i
			}
		}
		spike.PeakTs = buckets[peakIdx].BucketTs
		spike.PeakNumBids = peakBids
		spike.BucketCount = int64(run.endIdx - run.startIdx + 1)
		spikes = append(spikes, spike)
	}

	sort.Slice(spikes, func(i, j int) bool {
		return spikes[i].StartTs < spikes[j].StartTs
	})
	for i := range spikes {
		spikes[i].Index = i
	}
	return spikes
}

// ParseOptionalIntQuery parses an optional integer query/path parameter,
// returning defaultVal when the value is empty or not an integer.
func ParseOptionalIntQuery(cVal string, defaultVal int) int {
	if cVal == "" {
		return defaultVal
	}
	v, err := strconv.Atoi(cVal)
	if err != nil {
		return defaultVal
	}
	return v
}
