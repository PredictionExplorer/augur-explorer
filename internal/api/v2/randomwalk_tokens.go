package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// #nosec G101 -- a route instance and problem detail, not credentials.
const (
	randomWalkTokensInstance      = "/api/v2/randomwalk/tokens"
	invalidRandomWalkCursorDetail = "The cursor is malformed, unsupported, or belongs to another resource, filter, sort or wallet."
)

func randomWalkTokenNotFoundProblem(instance string) Problem {
	return newProblem(
		http.StatusNotFound,
		"token-not-found",
		"Token not found",
		"The RandomWalk collection has not minted a token with that ID.",
		instance,
	)
}

func invalidRandomWalkCursorProblem(instance string) Problem {
	return newProblem(
		http.StatusBadRequest,
		"invalid-cursor",
		"Invalid cursor",
		invalidRandomWalkCursorDetail,
		instance,
	)
}

func negativeRandomWalkTokenProblem(instance string) Problem {
	return newProblem(
		http.StatusBadRequest,
		"invalid-parameter",
		"Invalid parameter",
		"The token ID must be zero or greater.",
		instance,
	)
}

// randomWalkTokenScope validates the directory's filter and sort inputs and
// carries them between the request, the cursor and the repository call.
type randomWalkTokenScope struct {
	filter randomWalkTokenFilterScope
	sort   RandomWalkTokenSort
}

func resolveRandomWalkTokenScope(params ListRandomWalkTokensParams) (randomWalkTokenScope, error) {
	scope := randomWalkTokenScope{sort: TokenId}
	if params.Sort != nil {
		scope.sort = *params.Sort
	}
	if !scope.sort.Valid() {
		return randomWalkTokenScope{}, errors.New("sort must be tokenId or mostTraded")
	}
	if params.Named != nil && *params.Named {
		scope.filter.Named = true
	}
	if params.Name != nil {
		scope.filter.Name = *params.Name
	}
	if scope.filter.Named && scope.filter.Name != "" {
		return randomWalkTokenScope{}, errors.New("the named and name filters are mutually exclusive")
	}
	if len(scope.filter.Name) > maxTokenNameSearchLength ||
		(params.Name != nil && scope.filter.Name == "") {
		return randomWalkTokenScope{}, fmt.Errorf(
			"the name filter must be 1 through %d characters", maxTokenNameSearchLength)
	}
	if params.MintedFrom != nil {
		from := *params.MintedFrom
		if from < 0 || from > analyticsMaxTimestamp {
			return randomWalkTokenScope{}, errors.New("mintedFrom is outside the supported range")
		}
		scope.filter.MintedFrom = &from
	}
	if params.MintedUntil != nil {
		until := *params.MintedUntil
		if until < 1 || until > analyticsMaxTimestamp {
			return randomWalkTokenScope{}, errors.New("mintedUntil is outside the supported range")
		}
		scope.filter.MintedUntil = &until
	}
	if scope.filter.MintedFrom != nil && scope.filter.MintedUntil != nil &&
		*scope.filter.MintedUntil <= *scope.filter.MintedFrom {
		return randomWalkTokenScope{}, errors.New("mintedUntil must be greater than mintedFrom")
	}
	return scope, nil
}

// randomWalkTokenPageOrdered reports whether row (tradeCount, tokenID)
// strictly follows the previous position under the scope's sort.
func randomWalkTokenPageOrdered(
	sort RandomWalkTokenSort,
	previousTradeCount, previousTokenID int64,
	hasPrevious bool,
	record rwstore.TokenRecord,
) bool {
	if record.TokenID < 0 || record.TradeCount < 0 {
		return false
	}
	if !hasPrevious {
		return true
	}
	if sort == MostTraded {
		return record.TradeCount < previousTradeCount ||
			(record.TradeCount == previousTradeCount && record.TokenID > previousTokenID)
	}
	return record.TokenID > previousTokenID
}

// ListRandomWalkTokens implements GET /api/v2/randomwalk/tokens.
func (s *Server) ListRandomWalkTokens(
	ctx context.Context,
	request ListRandomWalkTokensRequestObject,
) (ListRandomWalkTokensResponseObject, error) {
	badRequest := func(problem Problem) ListRandomWalkTokensResponseObject {
		return ListRandomWalkTokens400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListRandomWalkTokensResponseObject {
		return ListRandomWalkTokens500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(randomWalkTokensInstance),
			),
		}
	}

	scope, err := resolveRandomWalkTokenScope(request.Params)
	if err != nil {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			err.Error()+".",
			randomWalkTokensInstance,
		)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			randomWalkTokensInstance,
		)), nil
	}

	var after *rwstore.TokenPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeRandomWalkTokenCursor(*request.Params.Cursor, scope.filter, scope.sort)
		if err != nil {
			return badRequest(invalidRandomWalkCursorProblem(randomWalkTokensInstance)), nil
		}
		after = &rwstore.TokenPageCursor{
			TokenID:    cursor.TokenID,
			TradeCount: cursor.TradeCount,
		}
	}

	storeSort := rwstore.TokenSortByID
	if scope.sort == MostTraded {
		storeSort = rwstore.TokenSortByTrades
	}
	filter := rwstore.TokenFilter{
		NamedOnly:    scope.filter.Named,
		NameContains: scope.filter.Name,
		MintedFrom:   scope.filter.MintedFrom,
		MintedUntil:  scope.filter.MintedUntil,
	}
	records, hasMore, err := s.randomWalk.TokensPage(ctx, filter, storeSort, after, limit)
	if err != nil {
		s.logInternal(ctx, "list random walk tokens", err, "sort", scope.sort)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate random walk token page cardinality", err)
		return internal(), nil
	}

	data := make([]RandomWalkToken, 0, len(records))
	previousTradeCount, previousTokenID := int64(0), int64(0)
	hasPrevious := false
	if after != nil {
		previousTradeCount, previousTokenID = after.TradeCount, after.TokenID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if !randomWalkTokenPageOrdered(
			scope.sort, previousTradeCount, previousTokenID, hasPrevious, record) {
			s.logInternal(ctx, "validate random walk token page",
				errors.New("repository returned an unordered token row"),
				"token_id", record.TokenID)
			return internal(), nil
		}
		token, err := mapRandomWalkToken(record)
		if err != nil {
			s.logInternal(ctx, "map random walk token", err, "token_id", record.TokenID)
			return internal(), nil
		}
		data = append(data, token)
		previousTradeCount, previousTokenID = record.TradeCount, record.TokenID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if !hasPrevious {
			s.logInternal(ctx, "list random walk tokens",
				errors.New("repository reported another page without a cursor row"))
			return internal(), nil
		}
		cursor := randomWalkTokenCursor{
			Version:  randomWalkCursorVersion,
			Resource: randomWalkResourceTokens,
			Filter:   scope.filter,
			Sort:     scope.sort,
			TokenID:  previousTokenID,
		}
		if scope.sort == MostTraded {
			cursor.TradeCount = previousTradeCount
		}
		next, err := encodeRandomWalkTokenCursor(cursor)
		if err != nil {
			s.logInternal(ctx, "encode random walk token cursor", err)
			return internal(), nil
		}
		meta.NextCursor = &next
	}

	return ListRandomWalkTokens200JSONResponse{
		RandomWalkTokenPageJSONResponse: RandomWalkTokenPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetRandomWalkToken implements GET /api/v2/randomwalk/tokens/{tokenId}.
func (s *Server) GetRandomWalkToken(
	ctx context.Context,
	request GetRandomWalkTokenRequestObject,
) (GetRandomWalkTokenResponseObject, error) {
	instance := fmt.Sprintf("%s/%d", randomWalkTokensInstance, request.TokenId)
	internal := func() GetRandomWalkTokenResponseObject {
		return GetRandomWalkToken500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}
	}
	if request.TokenId < 0 {
		return GetRandomWalkToken400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				negativeRandomWalkTokenProblem(instance),
			),
		}, nil
	}
	record, err := s.randomWalk.TokenDetailV2(ctx, request.TokenId)
	if errors.Is(err, store.ErrNotFound) {
		return GetRandomWalkToken404ApplicationProblemPlusJSONResponse{
			NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(
				randomWalkTokenNotFoundProblem(instance),
			),
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "get random walk token", err, "token_id", request.TokenId)
		return internal(), nil
	}
	if record.TokenID != request.TokenId {
		s.logInternal(ctx, "validate random walk token detail",
			errors.New("repository returned another token"),
			"token_id", request.TokenId)
		return internal(), nil
	}
	detail, err := mapRandomWalkTokenDetail(record)
	if err != nil {
		s.logInternal(ctx, "map random walk token detail", err, "token_id", request.TokenId)
		return internal(), nil
	}
	return GetRandomWalkToken200JSONResponse{
		RandomWalkTokenDetailJSONResponse: RandomWalkTokenDetailJSONResponse(detail),
	}, nil
}

// listRandomWalkTokenScopedPage is the shared shape of the two token-scoped
// event collections: validate inputs, decode the token-scoped cursor, gate
// on mint existence, fetch one page and validate its descending order.
func listRandomWalkTokenScopedPage[R any, T any](
	ctx context.Context,
	s *Server,
	instance string,
	tokenID int64,
	requestedCursor *Cursor,
	requestedLimit *Limit,
	resource randomWalkResource,
	fetch func(context.Context, int64, *rwstore.EventPageCursor, int) ([]R, bool, error),
	eventLogID func(R) int64,
	mapRecord func(R) (T, error),
) (data []T, meta PageMeta, problem *Problem) {
	fail := func(p Problem) ([]T, PageMeta, *Problem) {
		return nil, PageMeta{}, &p
	}
	if tokenID < 0 {
		return fail(negativeRandomWalkTokenProblem(instance))
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
		cursor, err := decodeRandomWalkTokenEventCursor(*requestedCursor, tokenID, resource)
		if err != nil {
			return fail(invalidRandomWalkCursorProblem(instance))
		}
		after = &rwstore.EventPageCursor{EventLogID: cursor.EventLogID}
	}

	exists, err := s.randomWalk.CollectionTokenExists(ctx, tokenID)
	if err != nil {
		s.logInternal(ctx, "check random walk token existence", err, "token_id", tokenID)
		return fail(internalProblem(instance))
	}
	if !exists {
		return fail(randomWalkTokenNotFoundProblem(instance))
	}

	records, hasMore, err := fetch(ctx, tokenID, after, limit)
	if err != nil {
		s.logInternal(ctx, "list random walk token events", err,
			"token_id", tokenID, "resource", resource)
		return fail(internalProblem(instance))
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate random walk token event cardinality", err,
			"token_id", tokenID, "resource", resource)
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
			s.logInternal(ctx, "validate random walk token event page",
				errors.New("repository returned an unordered event"),
				"token_id", tokenID,
				"resource", resource,
				"event_log_id", id)
			return fail(internalProblem(instance))
		}
		mapped, err := mapRecord(record)
		if err != nil {
			s.logInternal(ctx, "map random walk token event", err,
				"token_id", tokenID,
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
			s.logInternal(ctx, "list random walk token events",
				errors.New("repository reported another page without a cursor row"),
				"token_id", tokenID, "resource", resource)
			return fail(internalProblem(instance))
		}
		next, err := encodeRandomWalkTokenEventCursor(randomWalkTokenEventCursor{
			Version:    randomWalkCursorVersion,
			Resource:   resource,
			TokenID:    tokenID,
			EventLogID: previousEventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode random walk token event cursor", err,
				"token_id", tokenID, "resource", resource)
			return fail(internalProblem(instance))
		}
		meta.NextCursor = &next
	}
	return data, meta, nil
}

// ListRandomWalkTokenNameHistory implements
// GET /api/v2/randomwalk/tokens/{tokenId}/name-history.
func (s *Server) ListRandomWalkTokenNameHistory(
	ctx context.Context,
	request ListRandomWalkTokenNameHistoryRequestObject,
) (ListRandomWalkTokenNameHistoryResponseObject, error) {
	instance := fmt.Sprintf("%s/%d/name-history", randomWalkTokensInstance, request.TokenId)
	data, meta, problem := listRandomWalkTokenScopedPage(
		ctx,
		s,
		instance,
		request.TokenId,
		request.Params.Cursor,
		request.Params.Limit,
		randomWalkResourceNameHistory,
		s.randomWalk.TokenNameChangesPageV2,
		func(record rwstore.TokenNameChangeRecord) int64 { return record.Tx.EvtLogID },
		mapRandomWalkTokenNameChange,
	)
	if problem != nil {
		switch problem.Status {
		case http.StatusBadRequest:
			return ListRandomWalkTokenNameHistory400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		case http.StatusNotFound:
			return ListRandomWalkTokenNameHistory404ApplicationProblemPlusJSONResponse{
				NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(*problem),
			}, nil
		default:
			return ListRandomWalkTokenNameHistory500ApplicationProblemPlusJSONResponse{
				InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
	}
	return ListRandomWalkTokenNameHistory200JSONResponse{
		RandomWalkTokenNameChangePageJSONResponse: RandomWalkTokenNameChangePageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListRandomWalkTokenEvents implements
// GET /api/v2/randomwalk/tokens/{tokenId}/events.
func (s *Server) ListRandomWalkTokenEvents(
	ctx context.Context,
	request ListRandomWalkTokenEventsRequestObject,
) (ListRandomWalkTokenEventsResponseObject, error) {
	instance := fmt.Sprintf("%s/%d/events", randomWalkTokensInstance, request.TokenId)
	data, meta, problem := listRandomWalkTokenScopedPage(
		ctx,
		s,
		instance,
		request.TokenId,
		request.Params.Cursor,
		request.Params.Limit,
		randomWalkResourceTokenEvents,
		s.randomWalk.TokenEventsPage,
		func(record rwstore.TokenEventRecord) int64 { return record.Tx.EvtLogID },
		mapRandomWalkTokenEvent,
	)
	if problem != nil {
		switch problem.Status {
		case http.StatusBadRequest:
			return ListRandomWalkTokenEvents400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		case http.StatusNotFound:
			return ListRandomWalkTokenEvents404ApplicationProblemPlusJSONResponse{
				NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(*problem),
			}, nil
		default:
			return ListRandomWalkTokenEvents500ApplicationProblemPlusJSONResponse{
				InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
	}
	return ListRandomWalkTokenEvents200JSONResponse{
		RandomWalkTokenEventPageJSONResponse: RandomWalkTokenEventPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}
