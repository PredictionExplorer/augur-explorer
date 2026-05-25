package cosmicgame

import (
	"database/sql"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/rwcg/primitives/cosmicgame"
)

const roundOpenExcludeSecs = 3600

// Excludes bids in the first hour after each cycle's round_start_time (FirstBidPlacedInRound).
func (sw *SQLStorageWrapper) excludeFirstHourAfterRoundStartSQL() string {
	schema := sw.S.SchemaName()
	return " AND NOT EXISTS (" +
		"SELECT 1 FROM " + schema + ".cg_round_stats rs " +
		"WHERE rs.round_num = b.round_num " +
		"AND rs.round_start_time IS NOT NULL " +
		"AND b.time_stamp >= rs.round_start_time " +
		"AND b.time_stamp < rs.round_start_time + INTERVAL '" + strconv.Itoa(roundOpenExcludeSecs) + " seconds'" +
		")"
}

func (sw *SQLStorageWrapper) Get_bid_frequency_by_period(initTs, finTs, intervalSecs int) []p.CGBidFrequencyBucket {
	if intervalSecs <= 0 {
		intervalSecs = finTs - initTs
		if intervalSecs <= 0 {
			intervalSecs = 3600
		}
	}

	intervalStr := strconv.Itoa(intervalSecs)
	schema := sw.S.SchemaName()
	excl := sw.excludeFirstHourAfterRoundStartSQL()

	var query string
	if intervalSecs == 3600 || intervalSecs == 86400 {
		// UTC epoch-aligned buckets so counts match regardless of init_ts lookback window.
		query = "WITH periods AS (" +
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
			"LEFT JOIN " + schema + ".cg_bid b ON " +
			"FLOOR(EXTRACT(EPOCH FROM b.time_stamp)::bigint / " + intervalStr + ") * " + intervalStr +
			" = FLOOR(EXTRACT(EPOCH FROM p.start_ts)::bigint) " +
			"AND b.time_stamp >= TO_TIMESTAMP($1) AND b.time_stamp < TO_TIMESTAMP($2)" +
			excl + " " +
			"GROUP BY p.start_ts " +
			"ORDER BY p.start_ts"
	} else {
		query = "WITH periods AS (" +
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
			"LEFT JOIN " + schema + ".cg_bid b ON " +
			"b.time_stamp >= p.start_ts AND b.time_stamp < p.end_ts" +
			excl + " " +
			"GROUP BY p.start_ts " +
			"ORDER BY p.start_ts"
	}

	var rows *sql.Rows
	var err error
	if intervalSecs == 3600 || intervalSecs == 86400 {
		rows, err = sw.S.Db().Query(query, initTs, finTs)
	} else {
		rows, err = sw.S.Db().Query(query, initTs, finTs, intervalSecs)
	}
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	defer rows.Close()

	records := make([]p.CGBidFrequencyBucket, 0, 256)
	for rows.Next() {
		var rec p.CGBidFrequencyBucket
		if err := rows.Scan(&rec.BucketTs, &rec.NumBids, &rec.UniqueBidders); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		records = append(records, rec)
	}
	return records
}

func (sw *SQLStorageWrapper) Get_top_bidders(limit int) []p.CGTopBidderInfo {
	if limit <= 0 {
		limit = 20
	}
	query := "SELECT b.bidder_aid, a.addr, b.num_bids " +
		"FROM " + sw.S.SchemaName() + ".cg_bidder b " +
		"LEFT JOIN address a ON b.bidder_aid = a.address_id " +
		"WHERE b.num_bids > 0 " +
		"ORDER BY b.num_bids DESC " +
		"LIMIT $1"

	rows, err := sw.S.Db().Query(query, limit)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	defer rows.Close()

	records := make([]p.CGTopBidderInfo, 0, limit)
	for rows.Next() {
		var rec p.CGTopBidderInfo
		if err := rows.Scan(&rec.BidderAid, &rec.BidderAddr, &rec.NumBids); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		records = append(records, rec)
	}
	return records
}

func (sw *SQLStorageWrapper) Get_top_bidder_active_periods(topN, initTs, finTs, gapHours, minBids int) ([]p.CGTopBidderInfo, []p.CGBidderActivePeriod) {
	if topN <= 0 {
		topN = 20
	}
	if gapHours <= 0 {
		gapHours = 6
	}
	if minBids <= 0 {
		minBids = 2
	}

	topBidders := sw.Get_top_bidders(topN)
	if len(topBidders) == 0 {
		return topBidders, nil
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
		"FROM " + sw.S.SchemaName() + ".cg_bid b " +
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

	rows, err := sw.S.Db().Query(query, pq.Array(bidderAids), initTs, finTs, gapHours, minBids)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	defer rows.Close()

	periods := make([]p.CGBidderActivePeriod, 0, 128)
	for rows.Next() {
		var rec p.CGBidderActivePeriod
		var sessionID int64
		if err := rows.Scan(
			&rec.BidderAid,
			&rec.BidderAddr,
			&sessionID,
			&rec.PeriodStart,
			&rec.PeriodEnd,
			&rec.NumBids,
		); err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		rec.DurationSecs = rec.PeriodEnd - rec.PeriodStart
		if rec.DurationSecs < 0 {
			rec.DurationSecs = 0
		}
		periods = append(periods, rec)
	}
	return topBidders, periods
}

func (sw *SQLStorageWrapper) Get_bid_time_bounds() (minTs, maxTs int64) {
	query := "SELECT " +
		"COALESCE(FLOOR(EXTRACT(EPOCH FROM MIN(time_stamp))), 0)::BIGINT, " +
		"COALESCE(FLOOR(EXTRACT(EPOCH FROM MAX(time_stamp))), 0)::BIGINT " +
		"FROM " + sw.S.SchemaName() + ".cg_bid"
	row := sw.S.Db().QueryRow(query)
	if err := row.Scan(&minTs, &maxTs); err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	return minTs, maxTs
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
