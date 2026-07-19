package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const (
	analyticsActivityIntervalSeconds = int64(3600)
	analyticsDefaultIntervalSeconds  = int64(86400)
	analyticsMaxWindowSeconds        = int64(5 * 366 * 24 * 60 * 60)
	analyticsMaxTimestamp            = int64(253402300799)
	analyticsMaxBuckets              = int64(2000)
	analyticsDefaultTop              = 20
	analyticsMaxTop                  = 100
	analyticsDefaultGapHours         = 6
	analyticsMaxGapHours             = 8760
	analyticsDefaultMinBids          = 2
	analyticsMaxMinBids              = 1_000_000
	analyticsMaxActivePeriods        = cgstore.MaxBiddingActivePeriods
	recentBidSpikeWindowSeconds      = int64(30 * 24 * 60 * 60)
	analyticsQueryTimeout            = 5 * time.Second
)

type analyticsRange struct {
	from      int64
	to        int64
	storeFrom int
	storeTo   int
}

type analyticsWindow struct {
	analyticsRange

	intervalSeconds int64
	storeInterval   int
}

// GetCosmicGameBiddingActivity implements
// GET /api/v2/cosmicgame/statistics/bidding/activity.
func (s *Server) GetCosmicGameBiddingActivity(
	ctx context.Context,
	request GetCosmicGameBiddingActivityRequestObject,
) (GetCosmicGameBiddingActivityResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/bidding/activity"
	window, err := resolveAnalyticsWindow(
		request.Params.From,
		request.Params.To,
		request.Params.IntervalSeconds,
		analyticsActivityIntervalSeconds,
		true,
	)
	if err != nil {
		return getBiddingActivityBadRequest(analyticsParameterProblem(instance, err)), nil
	}
	queryContext, cancel := context.WithTimeout(ctx, analyticsQueryTimeout)
	defer cancel()
	records, err := s.analytics.BidFrequencyByPeriodBounded(
		queryContext, window.storeFrom, window.storeTo, window.storeInterval,
	)
	if err != nil {
		s.logInternal(ctx, "get bidding activity", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return getBiddingActivityInternal(internalProblem(instance)), nil
	}
	buckets, normalized, err := mapBidFrequencyBuckets(records, window.to)
	if err != nil {
		s.logInternal(ctx, "map bidding activity buckets", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return getBiddingActivityInternal(internalProblem(instance)), nil
	}
	spikes, err := mapBidSpikes(
		cgstore.DetectBidSpikes(normalized, window.storeInterval),
		window.to,
	)
	if err != nil {
		s.logInternal(ctx, "map bidding activity spikes", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return getBiddingActivityInternal(internalProblem(instance)), nil
	}
	result := BiddingActivity{
		Buckets:             buckets,
		From:                window.from,
		IntervalSeconds:     window.intervalSeconds,
		RecentSpikeIndex:    recentBidSpikeIndex(spikes, s.now()),
		RecentWindowSeconds: recentBidSpikeWindowSeconds,
		Spikes:              spikes,
		To:                  window.to,
	}
	return GetCosmicGameBiddingActivity200JSONResponse{
		CosmicGameBiddingActivityJSONResponse: CosmicGameBiddingActivityJSONResponse(result),
	}, nil
}

// GetCosmicGameBiddingFrequency implements
// GET /api/v2/cosmicgame/statistics/bidding/frequency.
func (s *Server) GetCosmicGameBiddingFrequency(
	ctx context.Context,
	request GetCosmicGameBiddingFrequencyRequestObject,
) (GetCosmicGameBiddingFrequencyResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/bidding/frequency"
	window, err := resolveAnalyticsWindow(
		request.Params.From,
		request.Params.To,
		request.Params.IntervalSeconds,
		analyticsDefaultIntervalSeconds,
		true,
	)
	if err != nil {
		return getBiddingFrequencyBadRequest(analyticsParameterProblem(instance, err)), nil
	}
	queryContext, cancel := context.WithTimeout(ctx, analyticsQueryTimeout)
	defer cancel()
	records, err := s.analytics.BidFrequencyByPeriodBounded(
		queryContext, window.storeFrom, window.storeTo, window.storeInterval,
	)
	if err != nil {
		s.logInternal(ctx, "get bidding frequency", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return getBiddingFrequencyInternal(internalProblem(instance)), nil
	}
	buckets, _, err := mapBidFrequencyBuckets(records, window.to)
	if err != nil {
		s.logInternal(ctx, "map bidding frequency", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return getBiddingFrequencyInternal(internalProblem(instance)), nil
	}
	result := BiddingFrequency{
		Buckets:         buckets,
		From:            window.from,
		IntervalSeconds: window.intervalSeconds,
		To:              window.to,
	}
	return GetCosmicGameBiddingFrequency200JSONResponse{
		CosmicGameBiddingFrequencyJSONResponse: CosmicGameBiddingFrequencyJSONResponse(result),
	}, nil
}

// GetCosmicGameBiddingTypeRatio implements
// GET /api/v2/cosmicgame/statistics/bidding/type-ratio.
func (s *Server) GetCosmicGameBiddingTypeRatio(
	ctx context.Context,
	request GetCosmicGameBiddingTypeRatioRequestObject,
) (GetCosmicGameBiddingTypeRatioResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/bidding/type-ratio"
	window, err := resolveAnalyticsWindow(
		request.Params.From,
		request.Params.To,
		request.Params.IntervalSeconds,
		analyticsDefaultIntervalSeconds,
		false,
	)
	if err != nil {
		return getBiddingTypeRatioBadRequest(analyticsParameterProblem(instance, err)), nil
	}
	queryContext, cancel := context.WithTimeout(ctx, analyticsQueryTimeout)
	defer cancel()
	records, err := s.analytics.BidTypeRatioByPeriodBounded(
		queryContext, window.storeFrom, window.storeTo, window.storeInterval,
	)
	if err != nil {
		s.logInternal(ctx, "get bidding type ratio", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return getBiddingTypeRatioInternal(internalProblem(instance)), nil
	}
	buckets, err := mapBidTypeRatioBuckets(records, window.to)
	if err != nil {
		s.logInternal(ctx, "map bidding type ratio", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return getBiddingTypeRatioInternal(internalProblem(instance)), nil
	}
	result := BiddingTypeRatio{
		Buckets:         buckets,
		From:            window.from,
		IntervalSeconds: window.intervalSeconds,
		To:              window.to,
	}
	return GetCosmicGameBiddingTypeRatio200JSONResponse{
		CosmicGameBiddingTypeRatioJSONResponse: CosmicGameBiddingTypeRatioJSONResponse(result),
	}, nil
}

// GetCosmicGameBiddingTopActivePeriods implements
// GET /api/v2/cosmicgame/statistics/bidding/top-active-periods.
func (s *Server) GetCosmicGameBiddingTopActivePeriods(
	ctx context.Context,
	request GetCosmicGameBiddingTopActivePeriodsRequestObject,
) (GetCosmicGameBiddingTopActivePeriodsResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/bidding/top-active-periods"
	timeRange, err := resolveAnalyticsRange(request.Params.From, request.Params.To)
	if err != nil {
		return getBiddingTopPeriodsBadRequest(analyticsParameterProblem(instance, err)), nil
	}
	top, err := boundedAnalyticsInt(request.Params.Top, analyticsDefaultTop, 1, analyticsMaxTop, "top")
	if err != nil {
		return getBiddingTopPeriodsBadRequest(analyticsParameterProblem(instance, err)), nil
	}
	gapHours, err := boundedAnalyticsInt(
		request.Params.GapHours, analyticsDefaultGapHours, 1, analyticsMaxGapHours, "gapHours",
	)
	if err != nil {
		return getBiddingTopPeriodsBadRequest(analyticsParameterProblem(instance, err)), nil
	}
	minBids, err := boundedAnalyticsInt(
		request.Params.MinBids, analyticsDefaultMinBids, 1, analyticsMaxMinBids, "minBids",
	)
	if err != nil {
		return getBiddingTopPeriodsBadRequest(analyticsParameterProblem(instance, err)), nil
	}
	queryContext, cancel := context.WithTimeout(ctx, analyticsQueryTimeout)
	defer cancel()
	bidderRecords, periodRecords, hasMore, err := s.analytics.TopBidderActivePeriodsBounded(
		queryContext, top, timeRange.storeFrom, timeRange.storeTo, gapHours, minBids,
	)
	if err != nil {
		s.logInternal(ctx, "get top bidder active periods", err,
			"from", timeRange.from, "to", timeRange.to,
			"top", top, "gap_hours", gapHours, "min_bids", minBids)
		return getBiddingTopPeriodsInternal(internalProblem(instance)), nil
	}
	if hasMore {
		return getBiddingTopPeriodsBadRequest(newProblem(
			http.StatusBadRequest,
			"result-too-large",
			"Result too large",
			fmt.Sprintf(
				"The requested filters produce more than %d active periods; narrow the window or increase gapHours or minBids.",
				analyticsMaxActivePeriods,
			),
			instance,
		)), nil
	}
	bidders, periods, err := mapTopBidderActivePeriods(
		bidderRecords, periodRecords, timeRange.from, timeRange.to, top, minBids,
	)
	if err != nil {
		s.logInternal(ctx, "map top bidder active periods", err,
			"from", timeRange.from, "to", timeRange.to,
			"top", top, "gap_hours", gapHours, "min_bids", minBids)
		return getBiddingTopPeriodsInternal(internalProblem(instance)), nil
	}
	result := BiddingTopActivePeriods{
		ActivePeriods: periods,
		From:          timeRange.from,
		GapHours:      gapHours,
		MinBids:       minBids,
		To:            timeRange.to,
		Top:           top,
		TopBidders:    bidders,
	}
	return GetCosmicGameBiddingTopActivePeriods200JSONResponse{
		CosmicGameBiddingTopActivePeriodsJSONResponse: CosmicGameBiddingTopActivePeriodsJSONResponse(result),
	}, nil
}

// GetCosmicGameBiddingTimeBounds implements
// GET /api/v2/cosmicgame/statistics/bidding/time-bounds.
func (s *Server) GetCosmicGameBiddingTimeBounds(
	ctx context.Context,
	_ GetCosmicGameBiddingTimeBoundsRequestObject,
) (GetCosmicGameBiddingTimeBoundsResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/bidding/time-bounds"
	queryContext, cancel := context.WithTimeout(ctx, analyticsQueryTimeout)
	defer cancel()
	minTimestamp, maxTimestamp, err := s.analytics.BidTimeBounds(queryContext)
	if err != nil {
		s.logInternal(ctx, "get bidding time bounds", err)
		return getBiddingTimeBoundsInternal(internalProblem(instance)), nil
	}
	if minTimestamp < 0 || maxTimestamp < 0 || minTimestamp > maxTimestamp {
		s.logInternal(ctx, "map bidding time bounds", errors.New("repository returned invalid bid time bounds"),
			"min_timestamp", minTimestamp, "max_timestamp", maxTimestamp)
		return getBiddingTimeBoundsInternal(internalProblem(instance)), nil
	}
	return GetCosmicGameBiddingTimeBounds200JSONResponse{
		CosmicGameBiddingTimeBoundsJSONResponse: CosmicGameBiddingTimeBoundsJSONResponse{
			MaxTimestamp: maxTimestamp,
			MinTimestamp: minTimestamp,
		},
	}, nil
}

func resolveAnalyticsWindow(
	from int64,
	to int64,
	requestedInterval *int64,
	defaultInterval int64,
	epochAligned bool,
) (analyticsWindow, error) {
	timeRange, err := resolveAnalyticsRange(from, to)
	if err != nil {
		return analyticsWindow{}, err
	}
	interval := defaultInterval
	if requestedInterval != nil {
		interval = *requestedInterval
	}
	if interval < 1 {
		return analyticsWindow{}, errors.New("intervalSeconds must be greater than zero")
	}
	if interval > analyticsMaxWindowSeconds {
		return analyticsWindow{}, fmt.Errorf(
			"intervalSeconds must not exceed %d",
			analyticsMaxWindowSeconds,
		)
	}
	storeInterval, ok := checkedInt(interval)
	if !ok {
		return analyticsWindow{}, errors.New("intervalSeconds is outside the supported integer range")
	}
	bucketCount := anchoredBucketCount(from, to, interval)
	if epochAligned && (interval == 3600 || interval == 86400) {
		bucketCount = epochAlignedBucketCount(from, to, interval)
	}
	if bucketCount > analyticsMaxBuckets {
		return analyticsWindow{}, fmt.Errorf(
			"the requested window produces %d buckets; the maximum is %d",
			bucketCount,
			analyticsMaxBuckets,
		)
	}
	return analyticsWindow{
		analyticsRange:  timeRange,
		intervalSeconds: interval,
		storeInterval:   storeInterval,
	}, nil
}

func resolveAnalyticsRange(from, to int64) (analyticsRange, error) {
	if from < 0 {
		return analyticsRange{}, errors.New("from must be zero or greater")
	}
	if from > analyticsMaxTimestamp || to > analyticsMaxTimestamp {
		return analyticsRange{}, fmt.Errorf(
			"from and to must not exceed %d",
			analyticsMaxTimestamp,
		)
	}
	if to <= from {
		return analyticsRange{}, errors.New("to must be greater than from")
	}
	if to-from > analyticsMaxWindowSeconds {
		return analyticsRange{}, fmt.Errorf(
			"the analytics window must not exceed %d seconds",
			analyticsMaxWindowSeconds,
		)
	}
	storeFrom, fromOK := checkedInt(from)
	storeTo, toOK := checkedInt(to)
	if !fromOK || !toOK {
		return analyticsRange{}, errors.New("from or to is outside the supported integer range")
	}
	return analyticsRange{from: from, to: to, storeFrom: storeFrom, storeTo: storeTo}, nil
}

func anchoredBucketCount(from, to, interval int64) int64 {
	return 1 + (to-from-1)/interval
}

func epochAlignedBucketCount(from, to, interval int64) int64 {
	return (to-1)/interval - from/interval + 1
}

func checkedInt(value int64) (int, bool) {
	converted := int(value)
	return converted, int64(converted) == value
}

func boundedAnalyticsInt(value *int, defaultValue, minimum, maximum int, name string) (int, error) {
	resolved := defaultValue
	if value != nil {
		resolved = *value
	}
	if resolved < minimum || resolved > maximum {
		return 0, fmt.Errorf("%s must be between %d and %d", name, minimum, maximum)
	}
	return resolved, nil
}

func analyticsParameterProblem(instance string, err error) Problem {
	return newProblem(
		http.StatusBadRequest,
		"invalid-parameter",
		"Invalid parameter",
		err.Error()+".",
		instance,
	)
}

func getBiddingActivityBadRequest(problem Problem) GetCosmicGameBiddingActivityResponseObject {
	return GetCosmicGameBiddingActivity400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func getBiddingActivityInternal(problem Problem) GetCosmicGameBiddingActivityResponseObject {
	return GetCosmicGameBiddingActivity500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func getBiddingFrequencyBadRequest(problem Problem) GetCosmicGameBiddingFrequencyResponseObject {
	return GetCosmicGameBiddingFrequency400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func getBiddingFrequencyInternal(problem Problem) GetCosmicGameBiddingFrequencyResponseObject {
	return GetCosmicGameBiddingFrequency500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func getBiddingTypeRatioBadRequest(problem Problem) GetCosmicGameBiddingTypeRatioResponseObject {
	return GetCosmicGameBiddingTypeRatio400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func getBiddingTypeRatioInternal(problem Problem) GetCosmicGameBiddingTypeRatioResponseObject {
	return GetCosmicGameBiddingTypeRatio500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func getBiddingTopPeriodsBadRequest(problem Problem) GetCosmicGameBiddingTopActivePeriodsResponseObject {
	return GetCosmicGameBiddingTopActivePeriods400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func getBiddingTopPeriodsInternal(problem Problem) GetCosmicGameBiddingTopActivePeriodsResponseObject {
	return GetCosmicGameBiddingTopActivePeriods500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func getBiddingTimeBoundsInternal(problem Problem) GetCosmicGameBiddingTimeBoundsResponseObject {
	return GetCosmicGameBiddingTimeBounds500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}
