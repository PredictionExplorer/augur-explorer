package cosmicgame

import (
	"context"
	"math"
	"sort"
	"strconv"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

const (
	roundOpenExcludeSecs = 3600

	// MaxBiddingActivePeriods bounds the non-paginated v2 analytics response.
	MaxBiddingActivePeriods = 2000
)

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

// bidFrequencyBoundedSQL is the v2 query shape. It first filters and assigns
// each bid to one DATE_BIN bucket, then joins those aggregates to the bounded
// zero-fill series. This prevents a range join from comparing every bid with
// every generated bucket and enforces the exclusive upper timestamp.
func bidFrequencyBoundedSQL(intervalSecs int) (query string, epochAligned bool) {
	intervalStr := strconv.Itoa(intervalSecs)
	excl := excludeFirstHourAfterRoundStartSQL()
	if intervalSecs == 3600 || intervalSecs == 86400 {
		return "WITH periods AS (" +
			"SELECT generate_series AS start_ts " +
			"FROM generate_series(" +
			"TO_TIMESTAMP((FLOOR($1::float / " + intervalStr + ") * " + intervalStr + ")::bigint), " +
			"TO_TIMESTAMP((FLOOR(($2::bigint - 1)::float / " + intervalStr + ") * " + intervalStr + ")::bigint), " +
			"('" + intervalStr + " seconds')::interval" +
			") AS generate_series" +
			"), bucketed AS (" +
			"SELECT DATE_BIN(('" + intervalStr + " seconds')::interval, b.time_stamp, TO_TIMESTAMP(0)) AS start_ts, " +
			"COUNT(*)::BIGINT AS num_bids, COUNT(DISTINCT b.bidder_aid)::BIGINT AS unique_bidders " +
			"FROM cg_bid b " +
			"WHERE b.time_stamp >= TO_TIMESTAMP($1) AND b.time_stamp < TO_TIMESTAMP($2)" +
			excl + " " +
			"GROUP BY 1" +
			") " +
			"SELECT FLOOR(EXTRACT(EPOCH FROM p.start_ts))::BIGINT, " +
			"COALESCE(b.num_bids, 0)::BIGINT, COALESCE(b.unique_bidders, 0)::BIGINT " +
			"FROM periods p LEFT JOIN bucketed b ON b.start_ts = p.start_ts " +
			"ORDER BY p.start_ts", true
	}
	return "WITH periods AS (" +
		"SELECT generate_series AS start_ts " +
		"FROM generate_series(" +
		"TO_TIMESTAMP($1), TO_TIMESTAMP($2::bigint - 1), ($3 || ' seconds')::interval" +
		") AS generate_series" +
		"), bucketed AS (" +
		"SELECT DATE_BIN(($3 || ' seconds')::interval, b.time_stamp, TO_TIMESTAMP($1)) AS start_ts, " +
		"COUNT(*)::BIGINT AS num_bids, COUNT(DISTINCT b.bidder_aid)::BIGINT AS unique_bidders " +
		"FROM cg_bid b " +
		"WHERE b.time_stamp >= TO_TIMESTAMP($1) AND b.time_stamp < TO_TIMESTAMP($2)" +
		excl + " " +
		"GROUP BY 1" +
		") " +
		"SELECT FLOOR(EXTRACT(EPOCH FROM p.start_ts))::BIGINT, " +
		"COALESCE(b.num_bids, 0)::BIGINT, COALESCE(b.unique_bidders, 0)::BIGINT " +
		"FROM periods p LEFT JOIN bucketed b ON b.start_ts = p.start_ts " +
		"ORDER BY p.start_ts", false
}

// BidFrequencyByPeriod returns the bid-count series over [initTs, finTs) in
// buckets of intervalSecs seconds. Empty buckets are zero-filled so the
// caller can render a continuous series; bids in the first hour after a
// round opens are excluded (see excludeFirstHourAfterRoundStartSQL). A
// non-positive interval means one bucket spanning the whole range (falling
// back to an hour if the range itself is empty).
func (r *Repo) BidFrequencyByPeriod(ctx context.Context, initTs, finTs, intervalSecs int) ([]cgmodel.CGBidFrequencyBucket, error) {
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
	scan := func(rows pgx.Rows, rec *cgmodel.CGBidFrequencyBucket) error {
		return rows.Scan(&rec.BucketTs, &rec.NumBids, &rec.UniqueBidders)
	}
	return queryList(ctx, r, "bid frequency by period", 256, query, scan, args...)
}

// BidFrequencyByPeriodBounded is the v2 frequency projection. It preserves
// BidFrequencyByPeriod's bucket alignment and round-open exclusion while
// guaranteeing that bids at or after finTs cannot enter a partial final
// bucket.
func (r *Repo) BidFrequencyByPeriodBounded(ctx context.Context, initTs, finTs, intervalSecs int) ([]cgmodel.CGBidFrequencyBucket, error) {
	if intervalSecs <= 0 {
		intervalSecs = finTs - initTs
		if intervalSecs <= 0 {
			intervalSecs = 3600
		}
	}
	query, epochAligned := bidFrequencyBoundedSQL(intervalSecs)
	args := []any{initTs, finTs}
	if !epochAligned {
		args = append(args, strconv.Itoa(intervalSecs))
	}
	scan := func(rows pgx.Rows, rec *cgmodel.CGBidFrequencyBucket) error {
		return rows.Scan(&rec.BucketTs, &rec.NumBids, &rec.UniqueBidders)
	}
	return queryList(ctx, r, "bounded bid frequency by period", 256, query, scan, args...)
}

// BidTypeRatioByPeriod returns the bid-type composition per sampling window
// over [initTs, finTs). Each window reports raw counts per bid_type plus
// those counts normalized to a windowed 100% (per-interval mix, not
// cumulative). Windows with no bids report 0 counts and 0% for every type,
// so the caller can render a continuous, gap-free series.
// bid_type: 0=ETH, 1=RandomWalk, 2=CST.
func (r *Repo) BidTypeRatioByPeriod(ctx context.Context, initTs, finTs, intervalSecs int) ([]cgmodel.CGBidTypeRatioBucket, error) {
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

	return queryList(ctx, r, "bid type ratio by period", 256, query, scanBidTypeRatioBucket, initTs, finTs)
}

// BidTypeRatioByPeriodBounded is the v2 bid-type projection. It assigns each
// filtered bid to one bucket before zero-filling the series and strictly
// excludes bids at or after finTs from the partial final bucket.
func (r *Repo) BidTypeRatioByPeriodBounded(ctx context.Context, initTs, finTs, intervalSecs int) ([]cgmodel.CGBidTypeRatioBucket, error) {
	if intervalSecs <= 0 {
		intervalSecs = finTs - initTs
		if intervalSecs <= 0 {
			intervalSecs = 86400
		}
	}
	intervalStr := strconv.Itoa(intervalSecs)
	query := "WITH periods AS (" +
		"SELECT generate_series AS start_ts " +
		"FROM generate_series(" +
		"TO_TIMESTAMP($1), TO_TIMESTAMP($2::bigint - 1), ('" + intervalStr + " seconds')::interval" +
		") AS generate_series" +
		"), bucketed AS (" +
		"SELECT DATE_BIN(('" + intervalStr + " seconds')::interval, b.time_stamp, TO_TIMESTAMP($1)) AS start_ts, " +
		"COUNT(*) FILTER (WHERE b.bid_type = 0)::BIGINT AS eth_bids, " +
		"COUNT(*) FILTER (WHERE b.bid_type = 1)::BIGINT AS rwalk_bids, " +
		"COUNT(*) FILTER (WHERE b.bid_type = 2)::BIGINT AS cst_bids, " +
		"COUNT(*)::BIGINT AS total_bids " +
		"FROM cg_bid b " +
		"WHERE b.time_stamp >= TO_TIMESTAMP($1) AND b.time_stamp < TO_TIMESTAMP($2) " +
		"GROUP BY 1" +
		") " +
		"SELECT FLOOR(EXTRACT(EPOCH FROM p.start_ts))::BIGINT, " +
		"COALESCE(b.eth_bids, 0)::BIGINT, COALESCE(b.rwalk_bids, 0)::BIGINT, " +
		"COALESCE(b.cst_bids, 0)::BIGINT, COALESCE(b.total_bids, 0)::BIGINT " +
		"FROM periods p LEFT JOIN bucketed b ON b.start_ts = p.start_ts " +
		"ORDER BY p.start_ts"
	return queryList(
		ctx,
		r,
		"bounded bid type ratio by period",
		256,
		query,
		scanBidTypeRatioBucket,
		initTs,
		finTs,
	)
}

func scanBidTypeRatioBucket(rows pgx.Rows, rec *cgmodel.CGBidTypeRatioBucket) error {
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

// TopBidders returns the most active bidders by lifetime bid count,
// descending. A non-positive limit defaults to 20.
func (r *Repo) TopBidders(ctx context.Context, limit int) ([]cgmodel.CGTopBidderInfo, error) {
	return r.topBidders(ctx, limit, false)
}

func (r *Repo) topBidders(ctx context.Context, limit int, stableTies bool) ([]cgmodel.CGTopBidderInfo, error) {
	if limit <= 0 {
		limit = 20
	}
	query := "SELECT b.bidder_aid, a.addr, b.num_bids " +
		"FROM cg_bidder b " +
		"LEFT JOIN address a ON b.bidder_aid = a.address_id " +
		"WHERE b.num_bids > 0 " +
		topBiddersOrderBy(stableTies) +
		"LIMIT $1"
	scan := func(rows pgx.Rows, rec *cgmodel.CGTopBidderInfo) error {
		return rows.Scan(&rec.BidderAid, &rec.BidderAddr, &rec.NumBids)
	}
	return queryList(ctx, r, "top bidders", limit, query, scan, limit)
}

func topBiddersOrderBy(stableTies bool) string {
	if stableTies {
		return "ORDER BY b.num_bids DESC, b.bidder_aid "
	}
	return "ORDER BY b.num_bids DESC "
}

// TopBidderActivePeriods segments the bids of the topN bidders inside
// [initTs, finTs) into activity sessions: a gap longer than gapHours starts
// a new session, and sessions with fewer than minBids bids are dropped.
// Non-positive knobs default to topN=20, gapHours=6, minBids=2. When there
// are no bidders at all the period list is nil.
func (r *Repo) TopBidderActivePeriods(ctx context.Context, topN, initTs, finTs, gapHours, minBids int) ([]cgmodel.CGTopBidderInfo, []cgmodel.CGBidderActivePeriod, error) {
	bidders, periods, _, err := r.topBidderActivePeriods(
		ctx, topN, initTs, finTs, gapHours, minBids, false, 0,
	)
	return bidders, periods, err
}

// TopBidderActivePeriodsBounded is the deterministic v2 projection of
// TopBidderActivePeriods. Equal lifetime bid counts use the internal address
// ID as a final tie-breaker, and equal period starts have stable secondary
// keys. At most MaxBiddingActivePeriods rows are returned; hasMore asks the
// caller to reject an over-broad request instead of serializing an unbounded
// response. Internal IDs remain repository-only.
func (r *Repo) TopBidderActivePeriodsBounded(ctx context.Context, topN, initTs, finTs, gapHours, minBids int) ([]cgmodel.CGTopBidderInfo, []cgmodel.CGBidderActivePeriod, bool, error) {
	return r.topBidderActivePeriods(
		ctx,
		topN,
		initTs,
		finTs,
		gapHours,
		minBids,
		true,
		MaxBiddingActivePeriods,
	)
}

func (r *Repo) topBidderActivePeriods(
	ctx context.Context,
	topN, initTs, finTs, gapHours, minBids int,
	stableTies bool,
	periodLimit int,
) ([]cgmodel.CGTopBidderInfo, []cgmodel.CGBidderActivePeriod, bool, error) {
	if topN <= 0 {
		topN = 20
	}
	if gapHours <= 0 {
		gapHours = 6
	}
	if minBids <= 0 {
		minBids = 2
	}

	topBidders, err := r.topBidders(ctx, topN, stableTies)
	if err != nil {
		return nil, nil, false, err
	}
	if len(topBidders) == 0 {
		return topBidders, nil, false, nil
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
		activePeriodsOrderBy(stableTies)
	args := []any{bidderAids, initTs, finTs, strconv.Itoa(gapHours), minBids}
	if periodLimit > 0 {
		query += " LIMIT $6"
		args = append(args, periodLimit+1)
	}

	scan := func(rows pgx.Rows, rec *cgmodel.CGBidderActivePeriod) error {
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
	periods, err := queryList(ctx, r, "top bidder active periods", 128, query, scan, args...)
	if err != nil {
		return nil, nil, false, err
	}
	hasMore := periodLimit > 0 && len(periods) > periodLimit
	if hasMore {
		periods = periods[:periodLimit]
	}
	return topBidders, periods, hasMore, nil
}

func activePeriodsOrderBy(stableTies bool) string {
	if stableTies {
		return "ORDER BY MIN(time_stamp), bidder_aid, MAX(time_stamp)"
	}
	return "ORDER BY MIN(time_stamp)"
}

// BidTimeBounds returns the epoch timestamps of the first and last bid ever
// placed (0, 0 when there are no bids).
func (r *Repo) BidTimeBounds(ctx context.Context) (minTs, maxTs int64, err error) {
	query := "SELECT " +
		"COALESCE(FLOOR(EXTRACT(EPOCH FROM MIN(time_stamp))), 0)::BIGINT, " +
		"COALESCE(FLOOR(EXTRACT(EPOCH FROM MAX(time_stamp))), 0)::BIGINT " +
		"FROM cg_bid"
	if err := r.q(ctx).QueryRow(ctx, query).Scan(&minTs, &maxTs); err != nil {
		return 0, 0, store.WrapError("bid time bounds", err)
	}
	return minTs, maxTs, nil
}

// DetectBidSpikes finds merged runs of ordered buckets whose non-negative bid
// count exceeds a dynamic threshold. Invalid input returns no spikes.
func DetectBidSpikes(buckets []cgmodel.CGBidFrequencyBucket, intervalSecs int) []cgmodel.CGBidSpike {
	if len(buckets) == 0 || intervalSecs <= 0 {
		return nil
	}

	counts := make([]int64, len(buckets))
	var sum float64
	var previousTimestamp int64
	for i, b := range buckets {
		if b.NumBids < 0 {
			return nil
		}
		if i > 0 && b.BucketTs <= previousTimestamp {
			return nil
		}
		counts[i] = b.NumBids
		sum += float64(b.NumBids)
		previousTimestamp = b.BucketTs
	}
	mean := sum / float64(len(buckets))

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

	spikes := make([]cgmodel.CGBidSpike, 0, len(runs))
	for idx, run := range runs {
		spike := cgmodel.CGBidSpike{Index: idx}
		spike.StartTs = buckets[run.startIdx].BucketTs
		lastBucket := buckets[run.endIdx]
		interval := int64(intervalSecs)
		if lastBucket.BucketTs > math.MaxInt64-interval {
			spike.EndTs = math.MaxInt64
		} else {
			spike.EndTs = lastBucket.BucketTs + interval
		}

		var peakIdx int
		var peakBids int64
		for i := run.startIdx; i <= run.endIdx; i++ {
			if buckets[i].NumBids > math.MaxInt64-spike.TotalBids {
				spike.TotalBids = math.MaxInt64
			} else {
				spike.TotalBids += buckets[i].NumBids
			}
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
