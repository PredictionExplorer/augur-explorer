package cosmicgame

import (
	"math"
	"strconv"
	"strings"
	"testing"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
)

func TestParseOptionalIntQuery(t *testing.T) {
	cases := []struct {
		in         string
		defaultVal int
		want       int
	}{
		{"", 5, 5},
		{"7", 5, 7},
		{"-3", 5, -3},
		{"0", 5, 0},
		{"+7", 5, 7},
		{"abc", 5, 5},
		{"7.5", 5, 5},
		{" 7", 5, 5},                     // Atoi rejects surrounding whitespace
		{"999999999999999999999", 5, 5},  // out of int range
		{"-999999999999999999999", 5, 5}, // out of int range
	}
	for _, tc := range cases {
		if got := ParseOptionalIntQuery(tc.in, tc.defaultVal); got != tc.want {
			t.Errorf("ParseOptionalIntQuery(%q, %d) = %d, want %d", tc.in, tc.defaultVal, got, tc.want)
		}
	}
}

func FuzzParseOptionalIntQuery(f *testing.F) {
	f.Add("", 0)
	f.Add("42", 7)
	f.Add("-1", 7)
	f.Add("junk", 7)
	f.Add("999999999999999999999", 7)
	f.Fuzz(func(t *testing.T, s string, defaultVal int) {
		got := ParseOptionalIntQuery(s, defaultVal)
		want := defaultVal
		if v, err := strconv.Atoi(s); err == nil {
			want = v
		}
		if got != want {
			t.Fatalf("ParseOptionalIntQuery(%q, %d) = %d, want %d", s, defaultVal, got, want)
		}
	})
}

// TestBidFrequencySQLBranchSelection pins which sampling intervals get the
// UTC epoch-aligned bucket query (hour and day windows) versus the
// initTs-anchored one, and that each branch binds the parameters it
// documents: the anchored query takes the interval as a third bound
// parameter, the epoch-aligned one interpolates it as an integer literal.
func TestBidFrequencySQLBranchSelection(t *testing.T) {
	cases := []struct {
		intervalSecs     int
		wantEpochAligned bool
	}{
		{3600, true},
		{86400, true},
		{1, false},
		{900, false},
		{3599, false},
		{3601, false},
		{86401, false},
		{999999, false},
	}
	for _, tc := range cases {
		query, epochAligned := bidFrequencySQL(tc.intervalSecs)
		if epochAligned != tc.wantEpochAligned {
			t.Errorf("bidFrequencySQL(%d): epochAligned = %v, want %v",
				tc.intervalSecs, epochAligned, tc.wantEpochAligned)
		}
		if !strings.Contains(query, "$1") || !strings.Contains(query, "$2") {
			t.Errorf("bidFrequencySQL(%d): query must bind $1/$2 range params", tc.intervalSecs)
		}
		if epochAligned {
			if strings.Contains(query, "$3") {
				t.Errorf("bidFrequencySQL(%d): epoch-aligned query must not reference $3", tc.intervalSecs)
			}
		} else {
			if !strings.Contains(query, "($3 || ' seconds')::interval") {
				t.Errorf("bidFrequencySQL(%d): anchored query must bind the interval as $3", tc.intervalSecs)
			}
		}
		if !strings.Contains(query, "cg_round_stats") {
			t.Errorf("bidFrequencySQL(%d): first-hour-after-round-start exclusion missing", tc.intervalSecs)
		}
	}
}

// TestBidFrequencySQLInterpolationIsNumeric proves the only value ever
// interpolated into the epoch-aligned query text is the decimal rendering of
// the interval: strip the two known literals and the remaining text must be
// identical for any two intervals (i.e. nothing else varies with input).
func TestBidFrequencySQLInterpolationIsNumeric(t *testing.T) {
	a, _ := bidFrequencySQL(3600)
	b, _ := bidFrequencySQL(86400)
	excl := excludeFirstHourAfterRoundStartSQL() // shared literal, contains "3600" itself
	normalized := func(q, interval string) string {
		q = strings.ReplaceAll(q, excl, "<EXCL>")
		return strings.ReplaceAll(q, interval, "<N>")
	}
	if normalized(a, "3600") != normalized(b, "86400") {
		t.Error("epoch-aligned queries differ beyond the interval literal")
	}
}

func TestBidFrequencyBoundedSQLFiltersBeforeBucketing(t *testing.T) {
	t.Parallel()
	for _, interval := range []int{900, 3600} {
		query, _ := bidFrequencyBoundedSQL(interval)
		for _, required := range []string{
			"bucketed AS",
			"DATE_BIN(",
			"b.time_stamp >= TO_TIMESTAMP($1)",
			"b.time_stamp < TO_TIMESTAMP($2)",
			"$2::bigint - 1",
		} {
			if !strings.Contains(query, required) {
				t.Errorf("bidFrequencyBoundedSQL(%d) missing %q", interval, required)
			}
		}
		if strings.Contains(query, "b.time_stamp >= p.start_ts") {
			t.Errorf("bidFrequencyBoundedSQL(%d) retained the bucket×bid range join", interval)
		}
	}
}

func TestBiddingAnalyticsV2TieBreakersAreIsolated(t *testing.T) {
	t.Parallel()
	if got := topBiddersOrderBy(false); got != "ORDER BY b.num_bids DESC " {
		t.Fatalf("legacy top-bidder order = %q", got)
	}
	if got := topBiddersOrderBy(true); got != "ORDER BY b.num_bids DESC, b.bidder_aid " {
		t.Fatalf("stable top-bidder order = %q", got)
	}
	if got := activePeriodsOrderBy(false); got != "ORDER BY MIN(time_stamp)" {
		t.Fatalf("legacy active-period order = %q", got)
	}
	if got := activePeriodsOrderBy(true); got !=
		"ORDER BY MIN(time_stamp), bidder_aid, MAX(time_stamp)" {
		t.Fatalf("stable active-period order = %q", got)
	}
}

func TestDetectBidSpikes(t *testing.T) {
	t.Parallel()
	buckets := make([]p.CGBidFrequencyBucket, 10)
	for i := range buckets {
		buckets[i].BucketTs = int64(i * 10)
	}
	buckets[4].NumBids = 20
	buckets[5].NumBids = 20

	spikes := DetectBidSpikes(buckets, 10)
	if len(spikes) != 1 {
		t.Fatalf("spikes = %+v", spikes)
	}
	got := spikes[0]
	if got.Index != 0 || got.StartTs != 40 || got.EndTs != 60 ||
		got.PeakTs != 50 || got.PeakNumBids != 20 ||
		got.TotalBids != 40 || got.BucketCount != 2 {
		t.Fatalf("spike = %+v", got)
	}
}

func TestDetectBidSpikesKeepsSeparateRunsOrdered(t *testing.T) {
	t.Parallel()
	buckets := make([]p.CGBidFrequencyBucket, 10)
	for i := range buckets {
		buckets[i].BucketTs = int64(i * 60)
	}
	buckets[2].NumBids = 20
	buckets[7].NumBids = 20

	spikes := DetectBidSpikes(buckets, 60)
	if len(spikes) != 2 ||
		spikes[0].Index != 0 || spikes[0].StartTs != 120 ||
		spikes[1].Index != 1 || spikes[1].StartTs != 420 {
		t.Fatalf("spikes = %+v", spikes)
	}
}

func TestDetectBidSpikesRejectsInvalidInput(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		buckets  []p.CGBidFrequencyBucket
		interval int
	}{
		"empty":         {interval: 1},
		"zero interval": {buckets: []p.CGBidFrequencyBucket{{BucketTs: 1, NumBids: 10}}},
		"negative count": {
			buckets:  []p.CGBidFrequencyBucket{{BucketTs: 1, NumBids: -1}},
			interval: 1,
		},
		"unordered": {
			buckets: []p.CGBidFrequencyBucket{
				{BucketTs: 2, NumBids: 10},
				{BucketTs: 1, NumBids: 10},
			},
			interval: 1,
		},
		"duplicate timestamp": {
			buckets: []p.CGBidFrequencyBucket{
				{BucketTs: 1, NumBids: 10},
				{BucketTs: 1, NumBids: 10},
			},
			interval: 1,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := DetectBidSpikes(test.buckets, test.interval); got != nil {
				t.Fatalf("DetectBidSpikes = %+v, want nil", got)
			}
		})
	}
}

func TestDetectBidSpikesSupportsPreEpochWindows(t *testing.T) {
	t.Parallel()
	buckets := make([]p.CGBidFrequencyBucket, 10)
	for i := range buckets {
		buckets[i].BucketTs = int64(-100 + i*10)
	}
	buckets[5].NumBids = 20
	spikes := DetectBidSpikes(buckets, 10)
	if len(spikes) != 1 || spikes[0].StartTs != -50 || spikes[0].EndTs != -40 {
		t.Fatalf("pre-epoch spikes = %+v", spikes)
	}
}

func TestDetectBidSpikesSaturatesOverflow(t *testing.T) {
	t.Parallel()
	buckets := make([]p.CGBidFrequencyBucket, 10)
	for i := range buckets {
		buckets[i].BucketTs = int64(i)
	}
	buckets[9].BucketTs = math.MaxInt64 - 5
	buckets[9].NumBids = math.MaxInt64

	spikes := DetectBidSpikes(buckets, 10)
	if len(spikes) != 1 || spikes[0].EndTs != math.MaxInt64 ||
		spikes[0].TotalBids != math.MaxInt64 {
		t.Fatalf("spikes = %+v", spikes)
	}
}

func FuzzDetectBidSpikes(f *testing.F) {
	f.Add([]byte{0, 0, 20, 20, 0}, 60)
	f.Add([]byte{}, 1)
	f.Add([]byte{255}, 0)
	f.Fuzz(func(t *testing.T, counts []byte, interval int) {
		if len(counts) > 512 {
			counts = counts[:512]
		}
		buckets := make([]p.CGBidFrequencyBucket, len(counts))
		for i := range counts {
			buckets[i] = p.CGBidFrequencyBucket{
				BucketTs: int64(i * 1000),
				NumBids:  int64(counts[i]),
			}
		}
		spikes := DetectBidSpikes(buckets, interval)
		if interval <= 0 {
			if spikes != nil {
				t.Fatalf("non-positive interval returned %+v", spikes)
			}
			return
		}
		for i := range spikes {
			if spikes[i].Index != i || spikes[i].StartTs >= spikes[i].EndTs ||
				spikes[i].PeakTs < spikes[i].StartTs ||
				spikes[i].PeakTs >= spikes[i].EndTs ||
				spikes[i].PeakNumBids > spikes[i].TotalBids ||
				spikes[i].BucketCount < 1 ||
				(i > 0 && spikes[i].StartTs <= spikes[i-1].StartTs) {
				t.Fatalf("invalid spike %d: %+v", i, spikes[i])
			}
		}
	})
}
