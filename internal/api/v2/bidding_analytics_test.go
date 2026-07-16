package v2

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"testing"
	"time"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
)

func TestBiddingActivityReturnsBoundedSpikes(t *testing.T) {
	t.Parallel()
	records := make([]cgmodel.CGBidFrequencyBucket, 0, 11)
	for i := 0; i < 10; i++ {
		count := int64(0)
		if i == 5 {
			count = 20
		}
		records = append(records, cgmodel.CGBidFrequencyBucket{
			BucketTs:      1000 + int64(i*10),
			NumBids:       count,
			UniqueBidders: count / 2,
		})
	}
	// The legacy generate_series query includes a bucket at an exactly
	// divisible exclusive boundary. V2 deliberately omits it.
	records = append(records, cgmodel.CGBidFrequencyBucket{BucketTs: 1100})
	server := newBiddingAnalyticsTestServer(t, fakeBiddingAnalyticsReader{
		frequency: func(_ context.Context, from, to, interval int) ([]cgmodel.CGBidFrequencyBucket, error) {
			if from != 1000 || to != 1100 || interval != 10 {
				t.Fatalf("repository args = %d,%d,%d", from, to, interval)
			}
			return records, nil
		},
	})
	server.now = func() time.Time { return time.Unix(1080, 0) }

	response := serve(t, server,
		"/api/v2/cosmicgame/statistics/bidding/activity?from=1000&to=1100&intervalSeconds=10")
	if response.Code != http.StatusOK {
		t.Fatalf("response = %d %s", response.Code, response.Body.String())
	}
	var activity BiddingActivity
	decodeResponse(t, response, &activity)
	if len(activity.Buckets) != 10 || len(activity.Spikes) != 1 {
		t.Fatalf("activity = %+v", activity)
	}
	if activity.Buckets[len(activity.Buckets)-1].BucketTimestamp != 1090 {
		t.Fatalf("terminal bucket was not removed: %+v", activity.Buckets)
	}
	if activity.Spikes[0].StartTimestamp != 1050 ||
		activity.Spikes[0].PeakBidCount != 20 ||
		activity.RecentSpikeIndex == nil || *activity.RecentSpikeIndex != 0 {
		t.Fatalf("spikes = %+v recent=%v", activity.Spikes, activity.RecentSpikeIndex)
	}
}

func TestBiddingFrequencyUsesDocumentedDefault(t *testing.T) {
	t.Parallel()
	server := newBiddingAnalyticsTestServer(t, fakeBiddingAnalyticsReader{
		frequency: func(ctx context.Context, from, to, interval int) ([]cgmodel.CGBidFrequencyBucket, error) {
			if from != 1000 || to != 2000 || interval != int(analyticsDefaultIntervalSeconds) {
				t.Fatalf("repository args = %d,%d,%d", from, to, interval)
			}
			deadline, ok := ctx.Deadline()
			if remaining := time.Until(deadline); !ok || remaining <= 0 || remaining > analyticsQueryTimeout {
				t.Fatalf("query deadline = %v, ok=%v remaining=%v", deadline, ok, remaining)
			}
			return nil, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/statistics/bidding/frequency?from=1000&to=2000")
	if response.Code != http.StatusOK {
		t.Fatalf("response = %d %s", response.Code, response.Body.String())
	}
	var frequency BiddingFrequency
	decodeResponse(t, response, &frequency)
	if frequency.IntervalSeconds != analyticsDefaultIntervalSeconds ||
		frequency.Buckets == nil || len(frequency.Buckets) != 0 {
		t.Fatalf("frequency = %+v", frequency)
	}
}

func TestBiddingTypeRatioUsesDecimalPercentages(t *testing.T) {
	t.Parallel()
	server := newBiddingAnalyticsTestServer(t, fakeBiddingAnalyticsReader{
		ratio: func(_ context.Context, from, to, interval int) ([]cgmodel.CGBidTypeRatioBucket, error) {
			if from != 1000 || to != 1100 || interval != 100 {
				t.Fatalf("repository args = %d,%d,%d", from, to, interval)
			}
			return []cgmodel.CGBidTypeRatioBucket{
				{BucketTs: 1000, EthBids: 1, RwalkBids: 1, CstBids: 1, TotalBids: 3},
				{BucketTs: 1100},
			}, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/statistics/bidding/type-ratio?from=1000&to=1100&intervalSeconds=100")
	if response.Code != http.StatusOK {
		t.Fatalf("response = %d %s", response.Code, response.Body.String())
	}
	var ratio BiddingTypeRatio
	decodeResponse(t, response, &ratio)
	if len(ratio.Buckets) != 1 ||
		ratio.Buckets[0].EthPercentage != "33.33" ||
		ratio.Buckets[0].RandomWalkPercentage != "33.33" ||
		ratio.Buckets[0].CstPercentage != "33.33" {
		t.Fatalf("ratio = %+v", ratio)
	}
}

func TestBiddingTopActivePeriodsMapsAndSorts(t *testing.T) {
	t.Parallel()
	address1 := "0x0000000000000000000000000000000000000001"
	address2 := "0x0000000000000000000000000000000000000002"
	server := newBiddingAnalyticsTestServer(t, fakeBiddingAnalyticsReader{
		periods: func(
			_ context.Context,
			top, from, to, gapHours, minBids int,
		) ([]cgmodel.CGTopBidderInfo, []cgmodel.CGBidderActivePeriod, bool, error) {
			if top != 2 || from != 1000 || to != 2000 || gapHours != 4 || minBids != 2 {
				t.Fatalf("repository args = %d,%d,%d,%d,%d", top, from, to, gapHours, minBids)
			}
			return []cgmodel.CGTopBidderInfo{
					{BidderAid: 2, BidderAddr: address2, NumBids: 3},
					{BidderAid: 1, BidderAddr: address1, NumBids: 5},
				}, []cgmodel.CGBidderActivePeriod{
					{BidderAid: 2, BidderAddr: address2, PeriodStart: 1200, PeriodEnd: 1220, NumBids: 2, DurationSecs: 20},
					{BidderAid: 1, BidderAddr: address1, PeriodStart: 1100, PeriodEnd: 1130, NumBids: 3, DurationSecs: 30},
				}, false, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/statistics/bidding/top-active-periods?from=1000&to=2000&top=2&gapHours=4&minBids=2")
	if response.Code != http.StatusOK {
		t.Fatalf("response = %d %s", response.Code, response.Body.String())
	}
	var result BiddingTopActivePeriods
	decodeResponse(t, response, &result)
	if len(result.TopBidders) != 2 || result.TopBidders[0].LifetimeBidCount != 5 ||
		len(result.ActivePeriods) != 2 || result.ActivePeriods[0].StartTimestamp != 1100 {
		t.Fatalf("top active periods = %+v", result)
	}
}

func TestBiddingTopActivePeriodsRejectsOversizedResult(t *testing.T) {
	t.Parallel()
	server := newBiddingAnalyticsTestServer(t, fakeBiddingAnalyticsReader{
		periods: func(
			context.Context,
			int,
			int,
			int,
			int,
			int,
		) ([]cgmodel.CGTopBidderInfo, []cgmodel.CGBidderActivePeriod, bool, error) {
			return nil, nil, true, nil
		},
	})
	response := serve(t, server,
		"/api/v2/cosmicgame/statistics/bidding/top-active-periods?from=1000&to=2000")
	assertProblem(t, response, http.StatusBadRequest)
	if !strings.Contains(response.Body.String(), "result-too-large") {
		t.Fatalf("problem = %s", response.Body.String())
	}
}

func TestBiddingTimeBounds(t *testing.T) {
	t.Parallel()
	server := newBiddingAnalyticsTestServer(t, fakeBiddingAnalyticsReader{
		bounds: func(context.Context) (int64, int64, error) { return 100, 200, nil },
	})
	response := serve(t, server, "/api/v2/cosmicgame/statistics/bidding/time-bounds")
	if response.Code != http.StatusOK {
		t.Fatalf("response = %d %s", response.Code, response.Body.String())
	}
	var bounds BiddingTimeBounds
	decodeResponse(t, response, &bounds)
	if bounds.MinTimestamp != 100 || bounds.MaxTimestamp != 200 {
		t.Fatalf("bounds = %+v", bounds)
	}
}

func TestBiddingAnalyticsRejectsInvalidParameters(t *testing.T) {
	t.Parallel()
	tooWide := analyticsMaxWindowSeconds + 1
	tests := map[string]string{
		"missing required from": "/api/v2/cosmicgame/statistics/bidding/frequency?to=2",
		"negative from":         "/api/v2/cosmicgame/statistics/bidding/frequency?from=-1&to=2",
		"inverted range":        "/api/v2/cosmicgame/statistics/bidding/frequency?from=2&to=2",
		"window too wide":       "/api/v2/cosmicgame/statistics/bidding/frequency?from=0&to=" + int64String(tooWide),
		"timestamp too large":   "/api/v2/cosmicgame/statistics/bidding/frequency?from=" + int64String(analyticsMaxTimestamp) + "&to=" + int64String(analyticsMaxTimestamp+1),
		"zero interval":         "/api/v2/cosmicgame/statistics/bidding/frequency?from=0&to=2&intervalSeconds=0",
		"interval too large":    "/api/v2/cosmicgame/statistics/bidding/frequency?from=0&to=2&intervalSeconds=" + int64String(analyticsMaxWindowSeconds+1),
		"too many buckets":      "/api/v2/cosmicgame/statistics/bidding/frequency?from=0&to=2001&intervalSeconds=1",
		"malformed interval":    "/api/v2/cosmicgame/statistics/bidding/frequency?from=0&to=2&intervalSeconds=bad",
		"zero top":              "/api/v2/cosmicgame/statistics/bidding/top-active-periods?from=0&to=2&top=0",
		"top too large":         "/api/v2/cosmicgame/statistics/bidding/top-active-periods?from=0&to=2&top=101",
		"zero gap":              "/api/v2/cosmicgame/statistics/bidding/top-active-periods?from=0&to=2&gapHours=0",
		"zero minimum":          "/api/v2/cosmicgame/statistics/bidding/top-active-periods?from=0&to=2&minBids=0",
	}
	for name, path := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			response := serve(t, newBiddingAnalyticsTestServer(t, fakeBiddingAnalyticsReader{}), path)
			assertProblem(t, response, http.StatusBadRequest)
		})
	}
}

func TestBiddingAnalyticsHidesRepositoryErrors(t *testing.T) {
	t.Parallel()
	secret := errors.New("password=private")
	server := newBiddingAnalyticsTestServer(t, fakeBiddingAnalyticsReader{
		frequency: func(context.Context, int, int, int) ([]cgmodel.CGBidFrequencyBucket, error) {
			return nil, secret
		},
		ratio: func(context.Context, int, int, int) ([]cgmodel.CGBidTypeRatioBucket, error) {
			return nil, secret
		},
		periods: func(context.Context, int, int, int, int, int) ([]cgmodel.CGTopBidderInfo, []cgmodel.CGBidderActivePeriod, bool, error) {
			return nil, nil, false, secret
		},
		bounds: func(context.Context) (int64, int64, error) { return 0, 0, secret },
	})
	for _, path := range []string{
		"/api/v2/cosmicgame/statistics/bidding/activity?from=0&to=2",
		"/api/v2/cosmicgame/statistics/bidding/frequency?from=0&to=2",
		"/api/v2/cosmicgame/statistics/bidding/type-ratio?from=0&to=2",
		"/api/v2/cosmicgame/statistics/bidding/top-active-periods?from=0&to=2",
		"/api/v2/cosmicgame/statistics/bidding/time-bounds",
	} {
		response := serve(t, server, path)
		assertProblem(t, response, http.StatusInternalServerError)
		if strings.Contains(response.Body.String(), "private") {
			t.Fatalf("internal error leaked for %s: %s", path, response.Body.String())
		}
	}
}

func TestBiddingAnalyticsRejectsInconsistentRepositoryOutput(t *testing.T) {
	t.Parallel()
	address := "0x0000000000000000000000000000000000000001"
	tests := map[string]struct {
		reader fakeBiddingAnalyticsReader
		path   string
	}{
		"frequency counts": {
			reader: fakeBiddingAnalyticsReader{frequency: func(context.Context, int, int, int) ([]cgmodel.CGBidFrequencyBucket, error) {
				return []cgmodel.CGBidFrequencyBucket{{BucketTs: 0, NumBids: 1, UniqueBidders: 2}}, nil
			}},
			path: "/api/v2/cosmicgame/statistics/bidding/frequency?from=0&to=2&intervalSeconds=1",
		},
		"type ratio sum": {
			reader: fakeBiddingAnalyticsReader{ratio: func(context.Context, int, int, int) ([]cgmodel.CGBidTypeRatioBucket, error) {
				return []cgmodel.CGBidTypeRatioBucket{{BucketTs: 0, EthBids: 1, TotalBids: 2}}, nil
			}},
			path: "/api/v2/cosmicgame/statistics/bidding/type-ratio?from=0&to=2&intervalSeconds=1",
		},
		"period outsider": {
			reader: fakeBiddingAnalyticsReader{periods: func(context.Context, int, int, int, int, int) ([]cgmodel.CGTopBidderInfo, []cgmodel.CGBidderActivePeriod, bool, error) {
				return []cgmodel.CGTopBidderInfo{{BidderAid: 1, BidderAddr: address, NumBids: 1}},
					[]cgmodel.CGBidderActivePeriod{{BidderAid: 2, BidderAddr: address, PeriodStart: 0, PeriodEnd: 1, NumBids: 2, DurationSecs: 1}}, false, nil
			}},
			path: "/api/v2/cosmicgame/statistics/bidding/top-active-periods?from=0&to=2",
		},
		"bounds": {
			reader: fakeBiddingAnalyticsReader{bounds: func(context.Context) (int64, int64, error) {
				return 2, 1, nil
			}},
			path: "/api/v2/cosmicgame/statistics/bidding/time-bounds",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assertProblem(t, serve(t, newBiddingAnalyticsTestServer(t, test.reader), test.path),
				http.StatusInternalServerError)
		})
	}
}

func TestAnalyticsBucketCounts(t *testing.T) {
	t.Parallel()
	if got := anchoredBucketCount(0, 2000, 1); got != 2000 {
		t.Fatalf("anchored exact count = %d", got)
	}
	if got := epochAlignedBucketCount(3599, 3601, 3600); got != 2 {
		t.Fatalf("epoch-aligned split count = %d", got)
	}
	if _, err := resolveAnalyticsWindow(0, 2000, int64Pointer(1), 1, false); err != nil {
		t.Fatalf("maximum bucket count rejected: %v", err)
	}
	if _, err := resolveAnalyticsWindow(0, 2001, int64Pointer(1), 1, false); err == nil {
		t.Fatal("over-limit bucket count accepted")
	}
}

func FuzzResolveAnalyticsWindow(f *testing.F) {
	f.Add(int64(0), int64(1), int64(1), false)
	f.Add(int64(3599), int64(3601), int64(3600), true)
	f.Add(int64(-1), int64(2), int64(1), false)
	f.Add(int64(0), analyticsMaxWindowSeconds+1, int64(1), true)
	f.Fuzz(func(t *testing.T, from, to, interval int64, epochAligned bool) {
		window, err := resolveAnalyticsWindow(from, to, &interval, 1, epochAligned)
		if err != nil {
			return
		}
		if window.from < 0 || window.to <= window.from ||
			window.to > analyticsMaxTimestamp ||
			window.to-window.from > analyticsMaxWindowSeconds ||
			window.intervalSeconds < 1 ||
			window.intervalSeconds > analyticsMaxWindowSeconds {
			t.Fatalf("accepted invalid window: %+v", window)
		}
		count := anchoredBucketCount(window.from, window.to, window.intervalSeconds)
		if epochAligned && (window.intervalSeconds == 3600 || window.intervalSeconds == 86400) {
			count = epochAlignedBucketCount(window.from, window.to, window.intervalSeconds)
		}
		if count > analyticsMaxBuckets {
			t.Fatalf("accepted %d buckets", count)
		}
	})
}

func newBiddingAnalyticsTestServer(t *testing.T, analytics biddingAnalyticsReader) *Server {
	t.Helper()
	server, err := newServer(
		nil,
		fakeBidReader{},
		fakeRoundReader{},
		fakeCurrentRoundReader{},
		fakeRoundPrizeReader{},
		fakeRoundRaffleReader{},
		fakeRoundDonationReader{},
		fakeStatisticsReader{},
		analytics,
		fakeContractAddressReader{},
		fakeParticipantReader{},
		fakeUserReader{},
		fakeUserHistoryReader{},
		fakeUserStakingReader{},
		fakeUserActivityReader{},
		fakeGlobalDirectoryReader{},
		fakeGlobalStakingReader{},
		fakeRandomWalkReader{},
		fakeContractState{},
		slog.New(slog.NewTextHandler(io.Discard, nil)),
	)
	if err != nil {
		t.Fatalf("newServer: %v", err)
	}
	return server
}

func int64Pointer(value int64) *int64 {
	return &value
}

func int64String(value int64) string {
	return fmt.Sprintf("%d", value)
}
