//go:build integration

package cosmicgame

import (
	"context"
	"reflect"
	"testing"
	"time"
)

// Fixture bid activity spans 1767225600..1767233200 (2026-01-01, 100s blocks).
const (
	fixtureStartTs = 1767225600
	fixtureEndTs   = 1767234000
)

func TestBidFrequencyByPeriod(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "bid_frequency_by_period", func() any {
		buckets, err := r.BidFrequencyByPeriod(ctx, fixtureStartTs, fixtureEndTs, 900)
		if err != nil {
			t.Fatalf("BidFrequencyByPeriod: %v", err)
		}
		return buckets
	})
	// Epoch-aligned branch (3600/86400 uses a different query shape).
	golden(t, "bid_frequency_by_period_hourly", func() any {
		buckets, err := r.BidFrequencyByPeriod(ctx, fixtureStartTs, fixtureEndTs, 3600)
		if err != nil {
			t.Fatalf("BidFrequencyByPeriod(3600): %v", err)
		}
		return buckets
	})
	// Buckets are zero-filled by generate_series: a range with no bids still
	// yields one bucket per interval, all with zero counts.
	buckets, err := r.BidFrequencyByPeriod(ctx, 100, 200, 60)
	if err != nil {
		t.Fatalf("BidFrequencyByPeriod(1970): %v", err)
	}
	for _, bucket := range buckets {
		if bucket.NumBids != 0 || bucket.UniqueBidders != 0 {
			t.Errorf("expected empty bucket in 1970, got %+v", bucket)
		}
	}
}

func TestBidTypeRatioByPeriod(t *testing.T) {
	r := repo(t)
	golden(t, "bid_type_ratio_by_period", func() any {
		buckets, err := r.BidTypeRatioByPeriod(context.Background(), fixtureStartTs, fixtureEndTs, 900)
		if err != nil {
			t.Fatalf("BidTypeRatioByPeriod: %v", err)
		}
		return buckets
	})
}

func TestBoundedBiddingAnalyticsExcludePartialBucketTail(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	var indexExists bool
	if err := r.pool().QueryRow(
		ctx,
		`SELECT EXISTS (
			SELECT 1 FROM pg_indexes
			WHERE schemaname = 'public' AND indexname = 'cg_bid_analytics_time_idx'
		)`,
	).Scan(&indexExists); err != nil {
		t.Fatalf("checking analytics timestamp index: %v", err)
	}
	if !indexExists {
		t.Fatal("cg_bid_analytics_time_idx is missing")
	}
	var originalRoundStart time.Time
	if err := r.pool().QueryRow(
		ctx,
		"SELECT round_start_time FROM cg_round_stats WHERE round_num = 0",
	).Scan(&originalRoundStart); err != nil {
		t.Fatalf("reading round start: %v", err)
	}
	if _, err := r.pool().Exec(
		ctx,
		"UPDATE cg_round_stats SET round_start_time = TO_TIMESTAMP($1) WHERE round_num = 0",
		1767222000,
	); err != nil {
		t.Fatalf("moving round start: %v", err)
	}
	t.Cleanup(func() {
		if _, err := r.pool().Exec(
			context.Background(),
			"UPDATE cg_round_stats SET round_start_time = $1 WHERE round_num = 0",
			originalRoundStart,
		); err != nil {
			t.Errorf("restoring round start: %v", err)
		}
	})

	const partialEnd = 1767225750
	legacyFrequency, err := r.BidFrequencyByPeriod(ctx, fixtureStartTs, partialEnd, 600)
	if err != nil {
		t.Fatalf("legacy frequency: %v", err)
	}
	boundedFrequency, err := r.BidFrequencyByPeriodBounded(ctx, fixtureStartTs, partialEnd, 600)
	if err != nil {
		t.Fatalf("bounded frequency: %v", err)
	}
	if len(legacyFrequency) != 1 || legacyFrequency[0].NumBids != 4 ||
		len(boundedFrequency) != 1 || boundedFrequency[0].NumBids != 1 {
		t.Fatalf("frequency legacy=%+v bounded=%+v", legacyFrequency, boundedFrequency)
	}

	legacyRatio, err := r.BidTypeRatioByPeriod(ctx, fixtureStartTs, partialEnd, 600)
	if err != nil {
		t.Fatalf("legacy type ratio: %v", err)
	}
	boundedRatio, err := r.BidTypeRatioByPeriodBounded(ctx, fixtureStartTs, partialEnd, 600)
	if err != nil {
		t.Fatalf("bounded type ratio: %v", err)
	}
	if len(legacyRatio) != 1 || legacyRatio[0].TotalBids != 4 ||
		len(boundedRatio) != 1 || boundedRatio[0].TotalBids != 1 {
		t.Fatalf("type ratio legacy=%+v bounded=%+v", legacyRatio, boundedRatio)
	}

	const post2038Start = 2_200_000_000
	for _, interval := range []int{60, 3600} {
		buckets, err := r.BidFrequencyByPeriodBounded(
			ctx,
			post2038Start,
			post2038Start+100,
			interval,
		)
		if err != nil {
			t.Fatalf("post-2038 bounded frequency (%d): %v", interval, err)
		}
		if len(buckets) == 0 {
			t.Fatalf("post-2038 bounded frequency (%d) returned no zero-fill buckets", interval)
		}
	}
	post2038Ratio, err := r.BidTypeRatioByPeriodBounded(
		ctx,
		post2038Start,
		post2038Start+100,
		60,
	)
	if err != nil {
		t.Fatalf("post-2038 bounded type ratio: %v", err)
	}
	if len(post2038Ratio) == 0 {
		t.Fatal("post-2038 bounded type ratio returned no zero-fill buckets")
	}
}

func TestTopBidders(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "top_bidders", func() any {
		bidders, err := r.TopBidders(ctx, 10)
		if err != nil {
			t.Fatalf("TopBidders(10): %v", err)
		}
		return bidders
	})
	golden(t, "top_bidders_limit_2", func() any {
		bidders, err := r.TopBidders(ctx, 2)
		if err != nil {
			t.Fatalf("TopBidders(2): %v", err)
		}
		return bidders
	})
}

func TestTopBidderActivePeriods(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	golden(t, "top_bidder_active_periods", func() any {
		bidders, periods, err := r.TopBidderActivePeriods(ctx, 3, fixtureStartTs, fixtureEndTs, 1, 1)
		if err != nil {
			t.Fatalf("TopBidderActivePeriods: %v", err)
		}
		return map[string]any{"topBidders": bidders, "activePeriods": periods}
	})
	legacyBidders, legacyPeriods, err := r.TopBidderActivePeriods(
		ctx, 3, fixtureStartTs, fixtureEndTs, 1, 1,
	)
	if err != nil {
		t.Fatalf("TopBidderActivePeriods: %v", err)
	}
	stableBidders, stablePeriods, hasMore, err := r.TopBidderActivePeriodsBounded(
		ctx, 3, fixtureStartTs, fixtureEndTs, 1, 1,
	)
	if err != nil {
		t.Fatalf("TopBidderActivePeriodsBounded: %v", err)
	}
	if hasMore ||
		!reflect.DeepEqual(stableBidders, legacyBidders) ||
		!reflect.DeepEqual(stablePeriods, legacyPeriods) {
		t.Fatalf("bounded projection changed untied fixture: bidders=%+v periods=%+v more=%v",
			stableBidders, stablePeriods, hasMore)
	}
	_, limitedPeriods, hasMore, err := r.topBidderActivePeriods(
		ctx,
		3,
		fixtureStartTs,
		fixtureEndTs,
		1,
		1,
		true,
		1,
	)
	if err != nil {
		t.Fatalf("limited topBidderActivePeriods: %v", err)
	}
	if !hasMore || len(limitedPeriods) != 1 {
		t.Fatalf("period cap: periods=%+v more=%v", limitedPeriods, hasMore)
	}
}

func TestBidTimeBounds(t *testing.T) {
	r := repo(t)
	golden(t, "bid_time_bounds", func() any {
		minTs, maxTs, err := r.BidTimeBounds(context.Background())
		if err != nil {
			t.Fatalf("BidTimeBounds: %v", err)
		}
		return map[string]int64{"minTs": minTs, "maxTs": maxTs}
	})
}
