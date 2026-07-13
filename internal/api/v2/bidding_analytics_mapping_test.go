package v2

import (
	"math"
	"testing"
	"time"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

func TestMapBidFrequencyBuckets(t *testing.T) {
	t.Parallel()
	records := []cgmodel.CGBidFrequencyBucket{
		{BucketTs: 0, NumBids: 2, UniqueBidders: 1},
		{BucketTs: 10, NumBids: 0, UniqueBidders: 0},
		{BucketTs: 20},
	}
	mapped, normalized, err := mapBidFrequencyBuckets(records, 20)
	if err != nil {
		t.Fatalf("mapBidFrequencyBuckets: %v", err)
	}
	if len(mapped) != 2 || len(normalized) != 2 ||
		mapped[0].BidCount != 2 || mapped[1].BucketTimestamp != 10 {
		t.Fatalf("mapped=%+v normalized=%+v", mapped, normalized)
	}
}

func TestMapBidFrequencyBucketsRejectsInvalidRows(t *testing.T) {
	t.Parallel()
	tests := map[string][]cgmodel.CGBidFrequencyBucket{
		"unordered": {
			{BucketTs: 10},
			{BucketTs: 5},
		},
		"duplicate": {
			{BucketTs: 5},
			{BucketTs: 5},
		},
		"after boundary": {
			{BucketTs: 20},
			{BucketTs: 21},
		},
		"too many unique": {
			{BucketTs: 5, NumBids: 1, UniqueBidders: 2},
		},
		"negative": {
			{BucketTs: 5, NumBids: -1},
		},
	}
	for name, records := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if _, _, err := mapBidFrequencyBuckets(records, 20); err == nil {
				t.Fatalf("records accepted: %+v", records)
			}
		})
	}
}

func TestMapBidSpikesClipsPartialFinalBucket(t *testing.T) {
	t.Parallel()
	mapped, err := mapBidSpikes([]cgmodel.CGBidSpike{{
		Index:       0,
		StartTs:     90,
		EndTs:       110,
		PeakTs:      90,
		PeakNumBids: 10,
		TotalBids:   10,
		BucketCount: 1,
	}}, 100)
	if err != nil {
		t.Fatalf("mapBidSpikes: %v", err)
	}
	if len(mapped) != 1 || mapped[0].EndTimestamp != 100 {
		t.Fatalf("spikes = %+v", mapped)
	}
}

func TestRecentBidSpikeIndex(t *testing.T) {
	t.Parallel()
	now := time.Unix(10_000_000, 0)
	cutoff := now.Unix() - recentBidSpikeWindowSeconds
	spikes := []BidSpike{
		{Index: 0, StartTimestamp: cutoff - 1},
		{Index: 1, StartTimestamp: cutoff},
		{Index: 2, StartTimestamp: now.Unix()},
		{Index: 3, StartTimestamp: now.Unix() + 1},
	}
	index := recentBidSpikeIndex(spikes, now)
	if index == nil || *index != 2 {
		t.Fatalf("recent index = %v", index)
	}
	if got := recentBidSpikeIndex(spikes[:1], now); got != nil {
		t.Fatalf("old spike marked recent: %v", got)
	}
}

func TestPercentageString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		part  int64
		total int64
		want  string
	}{
		{0, 0, "0"},
		{0, 3, "0"},
		{1, 4, "25"},
		{1, 3, "33.33"},
		{2, 3, "66.67"},
		{math.MaxInt64, math.MaxInt64, "100"},
	}
	for _, test := range tests {
		got, err := percentageString(test.part, test.total)
		if err != nil {
			t.Fatalf("percentageString(%d,%d): %v", test.part, test.total, err)
		}
		if got != test.want {
			t.Fatalf("percentageString(%d,%d) = %q, want %q",
				test.part, test.total, got, test.want)
		}
	}
	for _, invalid := range [][2]int64{{-1, 1}, {1, -1}, {2, 1}, {1, 0}} {
		if _, err := percentageString(invalid[0], invalid[1]); err == nil {
			t.Fatalf("percentageString(%d,%d) accepted invalid counts", invalid[0], invalid[1])
		}
	}
}

func TestMapBidTypeRatioBucketsRejectsInconsistentTotal(t *testing.T) {
	t.Parallel()
	_, err := mapBidTypeRatioBuckets([]cgmodel.CGBidTypeRatioBucket{{
		BucketTs:  1,
		EthBids:   1,
		RwalkBids: 1,
		CstBids:   1,
		TotalBids: 2,
	}}, 2)
	if err == nil {
		t.Fatal("inconsistent ratio total accepted")
	}
}

func TestMapTopBidderActivePeriodsCanonicalizesAndSorts(t *testing.T) {
	t.Parallel()
	address1 := "0x00000000000000000000000000000000000000aa"
	address2 := "0x00000000000000000000000000000000000000bb"
	bidders, periods, err := mapTopBidderActivePeriods(
		[]cgmodel.CGTopBidderInfo{
			{BidderAid: 2, BidderAddr: address2, NumBids: 2},
			{BidderAid: 1, BidderAddr: address1, NumBids: 2},
		},
		[]cgmodel.CGBidderActivePeriod{
			{BidderAid: 2, BidderAddr: address2, PeriodStart: 20, PeriodEnd: 21, NumBids: 2, DurationSecs: 1},
			{BidderAid: 1, BidderAddr: address1, PeriodStart: 10, PeriodEnd: 12, NumBids: 2, DurationSecs: 2},
		},
		0,
		100,
		2,
		2,
	)
	if err != nil {
		t.Fatalf("mapTopBidderActivePeriods: %v", err)
	}
	if len(bidders) != 2 || bidders[0].BidderAddress != "0x00000000000000000000000000000000000000AA" ||
		len(periods) != 2 || periods[0].StartTimestamp != 10 {
		t.Fatalf("bidders=%+v periods=%+v", bidders, periods)
	}
}
