//go:build integration

package cosmicgame

import (
	"context"
	"testing"
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
	golden(t, "top_bidder_active_periods", func() any {
		bidders, periods, err := r.TopBidderActivePeriods(context.Background(), 3, fixtureStartTs, fixtureEndTs, 1, 1)
		if err != nil {
			t.Fatalf("TopBidderActivePeriods: %v", err)
		}
		return map[string]any{"topBidders": bidders, "activePeriods": periods}
	})
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
