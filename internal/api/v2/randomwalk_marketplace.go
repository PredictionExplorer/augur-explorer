package v2

import (
	"context"
	"errors"
	"math/big"
	"net/http"

	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

const (
	randomWalkOffersInstance       = "/api/v2/randomwalk/marketplace/offers"
	randomWalkOfferHistoryInstance = "/api/v2/randomwalk/marketplace/offer-history"
	randomWalkTradesInstance       = "/api/v2/randomwalk/marketplace/trades"
	randomWalkFloorInstance        = "/api/v2/randomwalk/marketplace/floor-price"
)

// randomWalkOfferBookOrdered reports whether an order-book row strictly
// follows the previous position under the requested sort. Price ranks
// compare exact integers; ties break on ascending event-log ID.
func randomWalkOfferBookOrdered(
	sort RandomWalkOfferSort,
	previousPrice *big.Int,
	previousEventLogID int64,
	record rwstore.OfferRecord,
	recordPrice *big.Int,
) bool {
	if record.ListTx.EvtLogID < 1 {
		return false
	}
	switch sort {
	case Newest:
		return record.ListTx.EvtLogID < previousEventLogID
	case Oldest:
		return record.ListTx.EvtLogID > previousEventLogID
	case PriceAsc:
		comparison := recordPrice.Cmp(previousPrice)
		return comparison > 0 ||
			(comparison == 0 && record.ListTx.EvtLogID > previousEventLogID)
	case PriceDesc:
		comparison := recordPrice.Cmp(previousPrice)
		return comparison < 0 ||
			(comparison == 0 && record.ListTx.EvtLogID > previousEventLogID)
	default:
		return false
	}
}

// ListRandomWalkMarketplaceOffers implements
// GET /api/v2/randomwalk/marketplace/offers.
func (s *Server) ListRandomWalkMarketplaceOffers(
	ctx context.Context,
	request ListRandomWalkMarketplaceOffersRequestObject,
) (ListRandomWalkMarketplaceOffersResponseObject, error) {
	badRequest := func(problem Problem) ListRandomWalkMarketplaceOffersResponseObject {
		return ListRandomWalkMarketplaceOffers400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListRandomWalkMarketplaceOffersResponseObject {
		return ListRandomWalkMarketplaceOffers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(randomWalkOffersInstance),
			),
		}
	}

	sort := Newest
	if request.Params.Sort != nil {
		sort = *request.Params.Sort
	}
	if !sort.Valid() {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Sort must be newest, oldest, priceAsc or priceDesc.",
			randomWalkOffersInstance,
		)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			randomWalkOffersInstance,
		)), nil
	}

	var after *rwstore.OfferPageCursor
	var previousPrice *big.Int
	previousEventLogID := int64(0)
	hasPrevious := false
	if request.Params.Cursor != nil {
		cursor, err := decodeRandomWalkOfferBookCursor(*request.Params.Cursor, sort)
		if err != nil {
			return badRequest(invalidRandomWalkCursorProblem(randomWalkOffersInstance)), nil
		}
		after = &rwstore.OfferPageCursor{
			EventLogID: cursor.EventLogID,
			PriceWei:   cursor.PriceWei,
		}
		previousEventLogID = cursor.EventLogID
		hasPrevious = true
		if cursor.PriceWei != "" {
			price, ok := new(big.Int).SetString(cursor.PriceWei, 10)
			if !ok {
				return badRequest(invalidRandomWalkCursorProblem(randomWalkOffersInstance)), nil
			}
			previousPrice = price
		}
	}

	storeSort := map[RandomWalkOfferSort]rwstore.OfferSort{
		Newest:    rwstore.OfferSortNewest,
		Oldest:    rwstore.OfferSortOldest,
		PriceAsc:  rwstore.OfferSortPriceAsc,
		PriceDesc: rwstore.OfferSortPriceDesc,
	}[sort]
	records, hasMore, err := s.randomWalk.ActiveOffersPage(ctx, storeSort, after, limit)
	if err != nil {
		s.logInternal(ctx, "list random walk offers", err, "sort", sort)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate random walk offer page cardinality", err)
		return internal(), nil
	}

	priceSorted := sort == PriceAsc || sort == PriceDesc
	data := make([]RandomWalkMarketplaceOffer, 0, len(records))
	for i := range records {
		record := records[i]
		offer, err := mapRandomWalkMarketplaceOffer(record)
		if err != nil {
			s.logInternal(ctx, "map random walk offer", err,
				"event_log_id", record.ListTx.EvtLogID)
			return internal(), nil
		}
		recordPrice, ok := new(big.Int).SetString(offer.PriceWei, 10)
		if !ok {
			s.logInternal(ctx, "validate random walk offer page",
				errors.New("mapped offer price is not an integer"),
				"event_log_id", record.ListTx.EvtLogID)
			return internal(), nil
		}
		if hasPrevious && !randomWalkOfferBookOrdered(
			sort, previousPrice, previousEventLogID, record, recordPrice) {
			s.logInternal(ctx, "validate random walk offer page",
				errors.New("repository returned an unordered offer"),
				"event_log_id", record.ListTx.EvtLogID)
			return internal(), nil
		}
		data = append(data, offer)
		previousPrice = recordPrice
		previousEventLogID = record.ListTx.EvtLogID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list random walk offers",
				errors.New("repository reported another page without a cursor row"))
			return internal(), nil
		}
		cursor := randomWalkOfferBookCursor{
			Version:    randomWalkCursorVersion,
			Resource:   randomWalkResourceOfferBook,
			Sort:       sort,
			EventLogID: previousEventLogID,
		}
		if priceSorted {
			cursor.PriceWei = previousPrice.String()
		}
		next, err := encodeRandomWalkOfferBookCursor(cursor)
		if err != nil {
			s.logInternal(ctx, "encode random walk offer cursor", err)
			return internal(), nil
		}
		meta.NextCursor = &next
	}

	return ListRandomWalkMarketplaceOffers200JSONResponse{
		RandomWalkMarketplaceOfferPageJSONResponse: RandomWalkMarketplaceOfferPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// listRandomWalkLedgerPage is the shared shape of the newest-first
// RandomWalk event ledgers.
func listRandomWalkLedgerPage[R any, T any](
	ctx context.Context,
	s *Server,
	instance string,
	requestedCursor *Cursor,
	requestedLimit *Limit,
	resource randomWalkResource,
	address string,
	fetch func(context.Context, *rwstore.EventPageCursor, int) ([]R, bool, error),
	eventLogID func(R) int64,
	mapRecord func(R) (T, error),
) (data []T, meta PageMeta, problem *Problem) {
	fail := func(p Problem) ([]T, PageMeta, *Problem) {
		return nil, PageMeta{}, &p
	}
	limit, valid := resolvePageLimit(requestedLimit)
	if !valid {
		return fail(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		))
	}
	var after *rwstore.EventPageCursor
	if requestedCursor != nil {
		cursor, err := decodeRandomWalkLedgerCursor(*requestedCursor, resource, address)
		if err != nil {
			return fail(invalidRandomWalkCursorProblem(instance))
		}
		after = &rwstore.EventPageCursor{EventLogID: cursor.EventLogID}
	}

	records, hasMore, err := fetch(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list random walk ledger", err, "resource", resource)
		return fail(internalProblem(instance))
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate random walk ledger cardinality", err, "resource", resource)
		return fail(internalProblem(instance))
	}

	data = make([]T, 0, len(records))
	previousEventLogID := int64(0)
	hasPrevious := false
	if after != nil {
		previousEventLogID = after.EventLogID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		id := eventLogID(record)
		if id < 1 || (hasPrevious && id >= previousEventLogID) {
			s.logInternal(ctx, "validate random walk ledger page",
				errors.New("repository returned an unordered event"),
				"resource", resource,
				"event_log_id", id)
			return fail(internalProblem(instance))
		}
		mapped, err := mapRecord(record)
		if err != nil {
			s.logInternal(ctx, "map random walk ledger row", err,
				"resource", resource,
				"event_log_id", id)
			return fail(internalProblem(instance))
		}
		data = append(data, mapped)
		previousEventLogID = id
		hasPrevious = true
	}

	meta = PageMeta{Limit: limit}
	if hasMore {
		if !hasPrevious {
			s.logInternal(ctx, "list random walk ledger",
				errors.New("repository reported another page without a cursor row"),
				"resource", resource)
			return fail(internalProblem(instance))
		}
		next, err := encodeRandomWalkLedgerCursor(randomWalkLedgerCursor{
			Version:    randomWalkCursorVersion,
			Resource:   resource,
			Address:    address,
			EventLogID: previousEventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode random walk ledger cursor", err, "resource", resource)
			return fail(internalProblem(instance))
		}
		meta.NextCursor = &next
	}
	return data, meta, nil
}

// ListRandomWalkMarketplaceOfferHistory implements
// GET /api/v2/randomwalk/marketplace/offer-history.
func (s *Server) ListRandomWalkMarketplaceOfferHistory(
	ctx context.Context,
	request ListRandomWalkMarketplaceOfferHistoryRequestObject,
) (ListRandomWalkMarketplaceOfferHistoryResponseObject, error) {
	data, meta, problem := listRandomWalkLedgerPage(
		ctx,
		s,
		randomWalkOfferHistoryInstance,
		request.Params.Cursor,
		request.Params.Limit,
		randomWalkResourceOfferHistory,
		"",
		s.randomWalk.OfferHistoryPage,
		func(record rwstore.OfferHistoryRecord) int64 { return record.ListTx.EvtLogID },
		mapRandomWalkOfferHistoryEntry,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListRandomWalkMarketplaceOfferHistory400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListRandomWalkMarketplaceOfferHistory500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListRandomWalkMarketplaceOfferHistory200JSONResponse{
		RandomWalkOfferHistoryPageJSONResponse: RandomWalkOfferHistoryPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListRandomWalkMarketplaceTrades implements
// GET /api/v2/randomwalk/marketplace/trades.
func (s *Server) ListRandomWalkMarketplaceTrades(
	ctx context.Context,
	request ListRandomWalkMarketplaceTradesRequestObject,
) (ListRandomWalkMarketplaceTradesResponseObject, error) {
	data, meta, problem := listRandomWalkLedgerPage(
		ctx,
		s,
		randomWalkTradesInstance,
		request.Params.Cursor,
		request.Params.Limit,
		randomWalkResourceTrades,
		"",
		s.randomWalk.TradesPage,
		func(record rwstore.TradeRecord) int64 { return record.Tx.EvtLogID },
		mapRandomWalkTrade,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListRandomWalkMarketplaceTrades400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListRandomWalkMarketplaceTrades500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListRandomWalkMarketplaceTrades200JSONResponse{
		RandomWalkTradePageJSONResponse: RandomWalkTradePageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetRandomWalkMarketplaceFloorPrice implements
// GET /api/v2/randomwalk/marketplace/floor-price.
func (s *Server) GetRandomWalkMarketplaceFloorPrice(
	ctx context.Context,
	_ GetRandomWalkMarketplaceFloorPriceRequestObject,
) (GetRandomWalkMarketplaceFloorPriceResponseObject, error) {
	internal := func() GetRandomWalkMarketplaceFloorPriceResponseObject {
		return GetRandomWalkMarketplaceFloorPrice500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(randomWalkFloorInstance),
			),
		}
	}
	record, err := s.randomWalk.FloorPriceV2(ctx)
	if err != nil {
		s.logInternal(ctx, "get random walk floor price", err)
		return internal(), nil
	}
	floor, err := mapRandomWalkFloorPrice(record)
	if err != nil {
		s.logInternal(ctx, "map random walk floor price", err)
		return internal(), nil
	}
	return GetRandomWalkMarketplaceFloorPrice200JSONResponse{
		RandomWalkFloorPriceJSONResponse: RandomWalkFloorPriceJSONResponse(floor),
	}, nil
}
