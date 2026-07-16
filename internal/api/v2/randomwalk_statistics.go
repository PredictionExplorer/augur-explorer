package v2

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

const (
	randomWalkStatisticsInstance   = "/api/v2/randomwalk/statistics"
	randomWalkVolumeInstance       = "/api/v2/randomwalk/statistics/trading-volume"
	randomWalkListingFloorInstance = "/api/v2/randomwalk/statistics/listing-floor-history"
	randomWalkMintReportInstance   = "/api/v2/randomwalk/statistics/mint-report"
	randomWalkWithdrawalsInstance  = "/api/v2/randomwalk/withdrawals"
	randomWalkContractsInstance    = "/api/v2/randomwalk/contracts/addresses"
)

// GetRandomWalkStatistics implements GET /api/v2/randomwalk/statistics.
func (s *Server) GetRandomWalkStatistics(
	ctx context.Context,
	_ GetRandomWalkStatisticsRequestObject,
) (GetRandomWalkStatisticsResponseObject, error) {
	internal := func() GetRandomWalkStatisticsResponseObject {
		return GetRandomWalkStatistics500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(randomWalkStatisticsInstance),
			),
		}
	}
	record, err := s.randomWalk.StatisticsV2(ctx)
	if err != nil {
		s.logInternal(ctx, "get random walk statistics", err)
		return internal(), nil
	}
	statistics, err := mapRandomWalkStatistics(record)
	if err != nil {
		s.logInternal(ctx, "map random walk statistics", err)
		return internal(), nil
	}
	return GetRandomWalkStatistics200JSONResponse{
		RandomWalkStatisticsJSONResponse: RandomWalkStatisticsJSONResponse(statistics),
	}, nil
}

// GetRandomWalkTradingVolume implements
// GET /api/v2/randomwalk/statistics/trading-volume.
func (s *Server) GetRandomWalkTradingVolume(
	ctx context.Context,
	request GetRandomWalkTradingVolumeRequestObject,
) (GetRandomWalkTradingVolumeResponseObject, error) {
	badRequest := func(problem Problem) GetRandomWalkTradingVolumeResponseObject {
		return GetRandomWalkTradingVolume400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() GetRandomWalkTradingVolumeResponseObject {
		return GetRandomWalkTradingVolume500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(randomWalkVolumeInstance),
			),
		}
	}
	window, err := resolveRandomWalkSeriesWindow(
		request.Params.From, request.Params.To, request.Params.IntervalSeconds)
	if err != nil {
		return badRequest(analyticsParameterProblem(randomWalkVolumeInstance, err)), nil
	}
	queryContext, cancel := context.WithTimeout(ctx, analyticsQueryTimeout)
	defer cancel()
	baseVolume, records, err := s.randomWalk.TradingVolumeSeries(
		queryContext, window.storeFrom, window.storeTo, window.storeInterval)
	if err != nil {
		s.logInternal(ctx, "get random walk trading volume", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return internal(), nil
	}
	result, err := mapRandomWalkTradingVolume(window, baseVolume, records)
	if err != nil {
		s.logInternal(ctx, "map random walk trading volume", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return internal(), nil
	}
	return GetRandomWalkTradingVolume200JSONResponse{
		RandomWalkTradingVolumeJSONResponse: RandomWalkTradingVolumeJSONResponse(result),
	}, nil
}

// GetRandomWalkListingFloorHistory implements
// GET /api/v2/randomwalk/statistics/listing-floor-history.
func (s *Server) GetRandomWalkListingFloorHistory(
	ctx context.Context,
	request GetRandomWalkListingFloorHistoryRequestObject,
) (GetRandomWalkListingFloorHistoryResponseObject, error) {
	badRequest := func(problem Problem) GetRandomWalkListingFloorHistoryResponseObject {
		return GetRandomWalkListingFloorHistory400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() GetRandomWalkListingFloorHistoryResponseObject {
		return GetRandomWalkListingFloorHistory500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(randomWalkListingFloorInstance),
			),
		}
	}
	window, err := resolveRandomWalkSeriesWindow(
		request.Params.From, request.Params.To, request.Params.IntervalSeconds)
	if err != nil {
		return badRequest(analyticsParameterProblem(randomWalkListingFloorInstance, err)), nil
	}
	queryContext, cancel := context.WithTimeout(ctx, analyticsQueryTimeout)
	defer cancel()
	records, err := s.randomWalk.ListingFloorSeries(
		queryContext, window.storeFrom, window.storeTo, window.storeInterval)
	if err != nil {
		s.logInternal(ctx, "get random walk listing floor history", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return internal(), nil
	}
	result, err := mapRandomWalkListingFloorHistory(window, records)
	if err != nil {
		s.logInternal(ctx, "map random walk listing floor history", err,
			"from", window.from, "to", window.to, "interval_seconds", window.intervalSeconds)
		return internal(), nil
	}
	return GetRandomWalkListingFloorHistory200JSONResponse{
		RandomWalkListingFloorHistoryJSONResponse: RandomWalkListingFloorHistoryJSONResponse(result),
	}, nil
}

// GetRandomWalkMintReport implements
// GET /api/v2/randomwalk/statistics/mint-report.
func (s *Server) GetRandomWalkMintReport(
	ctx context.Context,
	_ GetRandomWalkMintReportRequestObject,
) (GetRandomWalkMintReportResponseObject, error) {
	internal := func() GetRandomWalkMintReportResponseObject {
		return GetRandomWalkMintReport500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(randomWalkMintReportInstance),
			),
		}
	}
	records, err := s.randomWalk.MintReportV2(ctx)
	if err != nil {
		s.logInternal(ctx, "get random walk mint report", err)
		return internal(), nil
	}
	result, err := mapRandomWalkMintReport(records)
	if err != nil {
		s.logInternal(ctx, "map random walk mint report", err)
		return internal(), nil
	}
	return GetRandomWalkMintReport200JSONResponse{
		RandomWalkMintReportJSONResponse: RandomWalkMintReportJSONResponse(result),
	}, nil
}

// ListRandomWalkWithdrawals implements GET /api/v2/randomwalk/withdrawals.
func (s *Server) ListRandomWalkWithdrawals(
	ctx context.Context,
	request ListRandomWalkWithdrawalsRequestObject,
) (ListRandomWalkWithdrawalsResponseObject, error) {
	data, meta, problem := listRandomWalkLedgerPage(
		ctx,
		s,
		randomWalkWithdrawalsInstance,
		request.Params.Cursor,
		request.Params.Limit,
		randomWalkResourceWithdrawals,
		"",
		s.randomWalk.WithdrawalsPage,
		func(record rwstore.WithdrawalRecord) int64 { return record.Tx.EvtLogID },
		mapRandomWalkWithdrawal,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListRandomWalkWithdrawals400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListRandomWalkWithdrawals500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListRandomWalkWithdrawals200JSONResponse{
		RandomWalkWithdrawalPageJSONResponse: RandomWalkWithdrawalPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetRandomWalkContractAddresses implements
// GET /api/v2/randomwalk/contracts/addresses.
func (s *Server) GetRandomWalkContractAddresses(
	ctx context.Context,
	_ GetRandomWalkContractAddressesRequestObject,
) (GetRandomWalkContractAddressesResponseObject, error) {
	internal := func() GetRandomWalkContractAddressesResponseObject {
		return GetRandomWalkContractAddresses500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(randomWalkContractsInstance),
			),
		}
	}
	addrs, err := s.randomWalk.ContractAddrs(ctx)
	if err != nil {
		s.logInternal(ctx, "get random walk contract addresses", err)
		return internal(), nil
	}
	nft, err := canonicalNonZeroAddress("RandomWalk NFT contract", addrs.RandomWalk)
	if err != nil {
		s.logInternal(ctx, "map random walk contract addresses", err)
		return internal(), nil
	}
	marketplace, err := canonicalNonZeroAddress("RandomWalk marketplace contract", addrs.MarketPlace)
	if err != nil {
		s.logInternal(ctx, "map random walk contract addresses", err)
		return internal(), nil
	}
	return GetRandomWalkContractAddresses200JSONResponse{
		RandomWalkContractAddressesJSONResponse: RandomWalkContractAddressesJSONResponse{
			MarketplaceAddress: marketplace,
			NftAddress:         nft,
		},
	}, nil
}

// resolveRandomWalkSeriesWindow validates a from/to/intervalSeconds chart
// window: the shared analytics range bounds plus the 2,000-bucket cap of
// the from-anchored series.
func resolveRandomWalkSeriesWindow(
	from, to int64,
	requestedInterval *int64,
) (analyticsWindow, error) {
	timeRange, err := resolveAnalyticsRange(from, to)
	if err != nil {
		return analyticsWindow{}, err
	}
	interval := analyticsDefaultIntervalSeconds
	if requestedInterval != nil {
		interval = *requestedInterval
	}
	if interval < 1 {
		return analyticsWindow{}, errors.New("intervalSeconds must be greater than zero")
	}
	if interval > analyticsMaxWindowSeconds {
		return analyticsWindow{}, fmt.Errorf(
			"intervalSeconds must not exceed %d", analyticsMaxWindowSeconds)
	}
	storeInterval, ok := checkedInt(interval)
	if !ok {
		return analyticsWindow{}, errors.New("intervalSeconds is outside the supported integer range")
	}
	if bucketCount := anchoredBucketCount(from, to, interval); bucketCount > analyticsMaxBuckets {
		return analyticsWindow{}, fmt.Errorf(
			"the requested window produces %d buckets; the maximum is %d",
			bucketCount, analyticsMaxBuckets)
	}
	return analyticsWindow{
		analyticsRange:  timeRange,
		intervalSeconds: interval,
		storeInterval:   storeInterval,
	}, nil
}

func mapRandomWalkTradingVolume(
	window analyticsWindow,
	baseVolumeWei string,
	records []rwstore.VolumeBucketRecord,
) (RandomWalkTradingVolume, error) {
	base, err := requiredAmount(baseVolumeWei)
	if err != nil {
		return RandomWalkTradingVolume{}, fmt.Errorf("base volume: %w", err)
	}
	running, _ := new(big.Int).SetString(base, 10)
	buckets := make([]RandomWalkVolumeBucket, 0, len(records))
	previousStart := int64(-1)
	for i := range records {
		record := records[i]
		if record.BucketStart < window.from || record.BucketStart >= window.to ||
			record.BucketStart <= previousStart {
			return RandomWalkTradingVolume{}, errors.New("bucket outside or out of window order")
		}
		if record.TradeCount < 0 {
			return RandomWalkTradingVolume{}, errors.New("negative bucket trade count")
		}
		volume, err := requiredAmount(record.VolumeWei)
		if err != nil {
			return RandomWalkTradingVolume{}, fmt.Errorf("bucket volume: %w", err)
		}
		if record.TradeCount == 0 && volume != "0" {
			return RandomWalkTradingVolume{}, errors.New("bucket volume without trades")
		}
		volumeInt, _ := new(big.Int).SetString(volume, 10)
		running = new(big.Int).Add(running, volumeInt)
		buckets = append(buckets, RandomWalkVolumeBucket{
			BucketStart:         record.BucketStart,
			CumulativeVolumeWei: running.String(),
			TradeCount:          record.TradeCount,
			VolumeWei:           volume,
		})
		previousStart = record.BucketStart
	}
	return RandomWalkTradingVolume{
		BaseVolumeWei:   base,
		Buckets:         buckets,
		From:            window.from,
		IntervalSeconds: window.intervalSeconds,
		To:              window.to,
	}, nil
}

func mapRandomWalkListingFloorHistory(
	window analyticsWindow,
	records []rwstore.FloorPointRecord,
) (RandomWalkListingFloorHistory, error) {
	points := make([]RandomWalkListingFloorPoint, 0, len(records))
	previousStart := int64(-1)
	for i := range records {
		record := records[i]
		if record.BucketStart < window.from || record.BucketStart >= window.to ||
			record.BucketStart <= previousStart {
			return RandomWalkListingFloorHistory{}, errors.New("floor point outside or out of window order")
		}
		floor, err := requiredAmount(record.FloorWei)
		if err != nil {
			return RandomWalkListingFloorHistory{}, fmt.Errorf("floor point price: %w", err)
		}
		points = append(points, RandomWalkListingFloorPoint{
			BucketStart:   record.BucketStart,
			FloorPriceWei: floor,
		})
		previousStart = record.BucketStart
	}
	return RandomWalkListingFloorHistory{
		From:            window.from,
		IntervalSeconds: window.intervalSeconds,
		Points:          points,
		To:              window.to,
	}, nil
}

func mapRandomWalkMintReport(records []rwstore.MonthlyMintRecord) (RandomWalkMintReport, error) {
	months := make([]RandomWalkMonthlyMints, 0, len(records))
	running := big.NewInt(0)
	previousYear, previousMonth := int64(0), int64(0)
	for i := range records {
		record := records[i]
		if record.Month < 1 || record.Month > 12 || record.Year < 2021 {
			return RandomWalkMintReport{}, errors.New("invalid mint report month")
		}
		if record.Year < previousYear ||
			(record.Year == previousYear && record.Month <= previousMonth) {
			return RandomWalkMintReport{}, errors.New("mint report months out of order")
		}
		if record.MintCount < 1 {
			return RandomWalkMintReport{}, errors.New("mint report month without mints")
		}
		minted, err := requiredAmount(record.MintedWei)
		if err != nil {
			return RandomWalkMintReport{}, fmt.Errorf("mint report amount: %w", err)
		}
		mintedInt, _ := new(big.Int).SetString(minted, 10)
		running = new(big.Int).Add(running, mintedInt)
		months = append(months, RandomWalkMonthlyMints{
			CumulativeMintedWei: running.String(),
			MintCount:           record.MintCount,
			MintedWei:           minted,
			Month:               record.Month,
			Year:                record.Year,
		})
		previousYear, previousMonth = record.Year, record.Month
	}
	return RandomWalkMintReport{Months: months}, nil
}
