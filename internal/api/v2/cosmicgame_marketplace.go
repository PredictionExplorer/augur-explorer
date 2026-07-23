package v2

import (
	"context"
	"errors"
	"math/big"
	"net/http"

	marketstore "github.com/PredictionExplorer/augur-explorer/internal/store/marketplace"
)

const (
	cosmicMarketplaceOffersInstance       = "/api/v2/cosmicgame/marketplace/offers"
	cosmicMarketplaceOfferHistoryInstance = "/api/v2/cosmicgame/marketplace/offer-history"
	cosmicMarketplaceTradesInstance       = "/api/v2/cosmicgame/marketplace/trades"
	cosmicMarketplaceFloorInstance        = "/api/v2/cosmicgame/marketplace/floor-price"
)

func cosmicMarketplaceOfferOrdered(
	sort RandomWalkOfferSort,
	previousPrice *big.Int,
	previousEventLogID int64,
	record marketstore.OfferRecord,
	recordPrice *big.Int,
) bool {
	if record.ListTx.EventLogID < 1 {
		return false
	}
	switch sort {
	case Newest:
		return record.ListTx.EventLogID < previousEventLogID
	case Oldest:
		return record.ListTx.EventLogID > previousEventLogID
	case PriceAsc:
		comparison := recordPrice.Cmp(previousPrice)
		return comparison > 0 ||
			(comparison == 0 && record.ListTx.EventLogID > previousEventLogID)
	case PriceDesc:
		comparison := recordPrice.Cmp(previousPrice)
		return comparison < 0 ||
			(comparison == 0 && record.ListTx.EventLogID > previousEventLogID)
	default:
		return false
	}
}

// ListCosmicGameMarketplaceOffers implements
// GET /api/v2/cosmicgame/marketplace/offers.
func (s *Server) ListCosmicGameMarketplaceOffers(
	ctx context.Context,
	request ListCosmicGameMarketplaceOffersRequestObject,
) (ListCosmicGameMarketplaceOffersResponseObject, error) {
	badRequest := func(problem Problem) ListCosmicGameMarketplaceOffersResponseObject {
		return ListCosmicGameMarketplaceOffers400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListCosmicGameMarketplaceOffersResponseObject {
		return ListCosmicGameMarketplaceOffers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(cosmicMarketplaceOffersInstance),
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
			cosmicMarketplaceOffersInstance,
		)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			cosmicMarketplaceOffersInstance,
		)), nil
	}

	var after *marketstore.OfferPageCursor
	var previousPrice *big.Int
	previousEventLogID := int64(0)
	hasPrevious := false
	if request.Params.Cursor != nil {
		cursor, err := decodeCosmicMarketplaceOfferCursor(*request.Params.Cursor, sort)
		if err != nil {
			return badRequest(
				invalidCosmicMarketplaceCursorProblem(cosmicMarketplaceOffersInstance),
			), nil
		}
		after = &marketstore.OfferPageCursor{
			EventLogID: cursor.EventLogID,
			PriceWei:   cursor.PriceWei,
		}
		previousEventLogID = cursor.EventLogID
		hasPrevious = true
		if cursor.PriceWei != "" {
			price, ok := new(big.Int).SetString(cursor.PriceWei, 10)
			if !ok {
				return badRequest(
					invalidCosmicMarketplaceCursorProblem(cosmicMarketplaceOffersInstance),
				), nil
			}
			previousPrice = price
		}
	}

	storeSort := map[RandomWalkOfferSort]marketstore.OfferSort{
		Newest:    marketstore.OfferSortNewest,
		Oldest:    marketstore.OfferSortOldest,
		PriceAsc:  marketstore.OfferSortPriceAsc,
		PriceDesc: marketstore.OfferSortPriceDesc,
	}[sort]
	records, hasMore, err := s.cosmicMarketplace.CosmicSignatureMarketplaceOffersPage(
		ctx,
		storeSort,
		after,
		limit,
	)
	if err != nil {
		s.logInternal(ctx, "list Cosmic Signature marketplace offers", err, "sort", sort)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate Cosmic Signature marketplace offer cardinality", err)
		return internal(), nil
	}

	priceSorted := sort == PriceAsc || sort == PriceDesc
	data := make([]CosmicSignatureMarketplaceOffer, 0, len(records))
	for i := range records {
		record := records[i]
		offer, err := mapCosmicSignatureMarketplaceOffer(record)
		if err != nil {
			s.logInternal(ctx, "map Cosmic Signature marketplace offer", err,
				"event_log_id", record.ListTx.EventLogID)
			return internal(), nil
		}
		recordPrice, ok := new(big.Int).SetString(offer.PriceWei, 10)
		if !ok {
			s.logInternal(ctx, "validate Cosmic Signature marketplace offer page",
				errors.New("mapped offer price is not an integer"),
				"event_log_id", record.ListTx.EventLogID)
			return internal(), nil
		}
		if hasPrevious && !cosmicMarketplaceOfferOrdered(
			sort,
			previousPrice,
			previousEventLogID,
			record,
			recordPrice,
		) {
			s.logInternal(ctx, "validate Cosmic Signature marketplace offer order",
				errors.New("repository returned an unordered offer"),
				"event_log_id", record.ListTx.EventLogID)
			return internal(), nil
		}
		data = append(data, offer)
		previousPrice = recordPrice
		previousEventLogID = record.ListTx.EventLogID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list Cosmic Signature marketplace offers",
				errors.New("repository reported another page without a cursor row"))
			return internal(), nil
		}
		cursor := cosmicMarketplaceOfferCursor{
			Version:    cosmicMarketplaceCursorVersion,
			Resource:   cosmicMarketplaceResourceOffers,
			Collection: cosmicSignatureCollectionScope,
			Sort:       sort,
			EventLogID: previousEventLogID,
		}
		if priceSorted {
			cursor.PriceWei = previousPrice.String()
		}
		next, err := encodeCosmicMarketplaceOfferCursor(cursor)
		if err != nil {
			s.logInternal(ctx, "encode Cosmic Signature marketplace offer cursor", err)
			return internal(), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameMarketplaceOffers200JSONResponse{
		CosmicSignatureMarketplaceOfferPageJSONResponse: CosmicSignatureMarketplaceOfferPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func listCosmicMarketplaceLedgerPage[R any, T any](
	ctx context.Context,
	s *Server,
	instance string,
	requestedCursor *Cursor,
	requestedLimit *Limit,
	resource cosmicMarketplaceResource,
	fetch func(context.Context, *marketstore.EventPageCursor, int) ([]R, bool, error),
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
	var after *marketstore.EventPageCursor
	if requestedCursor != nil {
		cursor, err := decodeCosmicMarketplaceLedgerCursor(*requestedCursor, resource)
		if err != nil {
			return fail(invalidCosmicMarketplaceCursorProblem(instance))
		}
		after = &marketstore.EventPageCursor{EventLogID: cursor.EventLogID}
	}

	records, hasMore, err := fetch(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list Cosmic Signature marketplace ledger", err, "resource", resource)
		return fail(internalProblem(instance))
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate Cosmic Signature marketplace ledger cardinality",
			err, "resource", resource)
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
			s.logInternal(ctx, "validate Cosmic Signature marketplace ledger order",
				errors.New("repository returned an unordered event"),
				"resource", resource,
				"event_log_id", id)
			return fail(internalProblem(instance))
		}
		mapped, err := mapRecord(record)
		if err != nil {
			s.logInternal(ctx, "map Cosmic Signature marketplace ledger row", err,
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
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list Cosmic Signature marketplace ledger",
				errors.New("repository reported another page without a cursor row"),
				"resource", resource)
			return fail(internalProblem(instance))
		}
		next, err := encodeCosmicMarketplaceLedgerCursor(cosmicMarketplaceLedgerCursor{
			Version:    cosmicMarketplaceCursorVersion,
			Resource:   resource,
			Collection: cosmicSignatureCollectionScope,
			EventLogID: previousEventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode Cosmic Signature marketplace ledger cursor",
				err, "resource", resource)
			return fail(internalProblem(instance))
		}
		meta.NextCursor = &next
	}
	return data, meta, nil
}

// ListCosmicGameMarketplaceOfferHistory implements
// GET /api/v2/cosmicgame/marketplace/offer-history.
func (s *Server) ListCosmicGameMarketplaceOfferHistory(
	ctx context.Context,
	request ListCosmicGameMarketplaceOfferHistoryRequestObject,
) (ListCosmicGameMarketplaceOfferHistoryResponseObject, error) {
	data, meta, problem := listCosmicMarketplaceLedgerPage(
		ctx,
		s,
		cosmicMarketplaceOfferHistoryInstance,
		request.Params.Cursor,
		request.Params.Limit,
		cosmicMarketplaceResourceOfferHistory,
		s.cosmicMarketplace.CosmicSignatureMarketplaceOfferHistoryPage,
		func(record marketstore.OfferHistoryRecord) int64 {
			return record.ListTx.EventLogID
		},
		mapCosmicSignatureMarketplaceOfferHistory,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameMarketplaceOfferHistory400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameMarketplaceOfferHistory500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameMarketplaceOfferHistory200JSONResponse{
		CosmicSignatureMarketplaceOfferHistoryPageJSONResponse: CosmicSignatureMarketplaceOfferHistoryPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameMarketplaceTrades implements
// GET /api/v2/cosmicgame/marketplace/trades.
func (s *Server) ListCosmicGameMarketplaceTrades(
	ctx context.Context,
	request ListCosmicGameMarketplaceTradesRequestObject,
) (ListCosmicGameMarketplaceTradesResponseObject, error) {
	data, meta, problem := listCosmicMarketplaceLedgerPage(
		ctx,
		s,
		cosmicMarketplaceTradesInstance,
		request.Params.Cursor,
		request.Params.Limit,
		cosmicMarketplaceResourceTrades,
		s.cosmicMarketplace.CosmicSignatureMarketplaceTradesPage,
		func(record marketstore.TradeRecord) int64 { return record.Tx.EventLogID },
		mapCosmicSignatureMarketplaceTrade,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameMarketplaceTrades400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameMarketplaceTrades500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameMarketplaceTrades200JSONResponse{
		CosmicSignatureMarketplaceTradePageJSONResponse: CosmicSignatureMarketplaceTradePageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetCosmicGameMarketplaceFloorPrice implements
// GET /api/v2/cosmicgame/marketplace/floor-price.
func (s *Server) GetCosmicGameMarketplaceFloorPrice(
	ctx context.Context,
	_ GetCosmicGameMarketplaceFloorPriceRequestObject,
) (GetCosmicGameMarketplaceFloorPriceResponseObject, error) {
	internal := func() GetCosmicGameMarketplaceFloorPriceResponseObject {
		return GetCosmicGameMarketplaceFloorPrice500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(cosmicMarketplaceFloorInstance),
			),
		}
	}
	record, err := s.cosmicMarketplace.CosmicSignatureMarketplaceFloorPrice(ctx)
	if err != nil {
		s.logInternal(ctx, "get Cosmic Signature marketplace floor price", err)
		return internal(), nil
	}
	floor, err := mapCosmicSignatureMarketplaceFloorPrice(record)
	if err != nil {
		s.logInternal(ctx, "map Cosmic Signature marketplace floor price", err)
		return internal(), nil
	}
	return GetCosmicGameMarketplaceFloorPrice200JSONResponse{
		CosmicSignatureMarketplaceFloorPriceJSONResponse: CosmicSignatureMarketplaceFloorPriceJSONResponse(floor),
	}, nil
}
