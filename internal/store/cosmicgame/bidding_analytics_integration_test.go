//go:build integration

package cosmicgame

import "testing"

// Fixture bid activity spans 1767225600..1767233200 (2026-01-01, 100s blocks).
const (
	fixtureStartTs = 1767225600
	fixtureEndTs   = 1767234000
)

func TestGetBidFrequencyByPeriod(t *testing.T) {
	sw := store(t)
	golden(t, "bid_frequency_by_period", func() any {
		return sw.Get_bid_frequency_by_period(fixtureStartTs, fixtureEndTs, 900)
	})
	// Buckets are zero-filled by generate_series: a range with no bids still
	// yields one bucket per interval, all with zero counts.
	for _, bucket := range sw.Get_bid_frequency_by_period(100, 200, 60) {
		if bucket.NumBids != 0 || bucket.UniqueBidders != 0 {
			t.Errorf("expected empty bucket in 1970, got %+v", bucket)
		}
	}
}

func TestGetBidTypeRatioByPeriod(t *testing.T) {
	sw := store(t)
	golden(t, "bid_type_ratio_by_period", func() any {
		return sw.Get_bid_type_ratio_by_period(fixtureStartTs, fixtureEndTs, 900)
	})
}

func TestGetTopBidders(t *testing.T) {
	sw := store(t)
	golden(t, "top_bidders", func() any {
		return sw.Get_top_bidders(10)
	})
	golden(t, "top_bidders_limit_2", func() any {
		return sw.Get_top_bidders(2)
	})
}

func TestGetTopBidderActivePeriods(t *testing.T) {
	sw := store(t)
	golden(t, "top_bidder_active_periods", func() any {
		bidders, periods := sw.Get_top_bidder_active_periods(3, fixtureStartTs, fixtureEndTs, 1, 1)
		return map[string]any{"topBidders": bidders, "activePeriods": periods}
	})
}

func TestGetBidTimeBounds(t *testing.T) {
	sw := store(t)
	golden(t, "bid_time_bounds", func() any {
		minTs, maxTs := sw.Get_bid_time_bounds()
		return map[string]int64{"minTs": minTs, "maxTs": maxTs}
	})
}
