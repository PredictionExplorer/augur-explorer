package v2

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// #nosec G101 -- route instances and problem details, not credentials.
const (
	globalTokensInstance         = "/api/v2/cosmicgame/cosmic-signature-tokens"
	csHoldersInstance            = "/api/v2/cosmicgame/cosmic-signature-tokens/holders"
	cosmicTokenHoldersInstance   = "/api/v2/cosmicgame/cosmic-token/holders"
	cosmicTokenStatsInstance     = "/api/v2/cosmicgame/cosmic-token/statistics"
	supplyByBidInstance          = "/api/v2/cosmicgame/cosmic-token/supply-history/by-bid"
	supplyDailyInstance          = "/api/v2/cosmicgame/cosmic-token/supply-history/daily"
	globalMarketingInstance      = "/api/v2/cosmicgame/marketing-rewards"
	contradictoryTokenFilterMsg  = "The named and name filters are mutually exclusive."
	invalidTokenCursorProblemMsg = "The cursor is malformed, unsupported, or belongs to another token or resource."
)

// ListCosmicGameCosmicSignatureTokens implements
// GET /api/v2/cosmicgame/cosmic-signature-tokens.
func (s *Server) ListCosmicGameCosmicSignatureTokens(
	ctx context.Context,
	request ListCosmicGameCosmicSignatureTokensRequestObject,
) (ListCosmicGameCosmicSignatureTokensResponseObject, error) {
	badRequest := func(problem Problem) ListCosmicGameCosmicSignatureTokensResponseObject {
		return ListCosmicGameCosmicSignatureTokens400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListCosmicGameCosmicSignatureTokensResponseObject {
		return ListCosmicGameCosmicSignatureTokens500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(globalTokensInstance),
			),
		}
	}

	scope := globalTokenFilterScope{}
	if request.Params.Named != nil && *request.Params.Named {
		scope.Named = true
	}
	if request.Params.Name != nil {
		scope.Name = *request.Params.Name
	}
	if scope.Named && scope.Name != "" {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			contradictoryTokenFilterMsg,
			globalTokensInstance,
		)), nil
	}
	if len(scope.Name) > maxTokenNameSearchLength || (request.Params.Name != nil && scope.Name == "") {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			fmt.Sprintf("The name filter must be 1 through %d characters.", maxTokenNameSearchLength),
			globalTokensInstance,
		)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			globalTokensInstance,
		)), nil
	}

	var after *cgstore.GlobalTokenPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeGlobalTokenCursor(*request.Params.Cursor, scope)
		if err != nil {
			return badRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another filter or resource.",
				globalTokensInstance,
			)), nil
		}
		after = &cgstore.GlobalTokenPageCursor{TokenID: cursor.TokenID}
	}

	filter := cgstore.GlobalTokenFilter{NamedOnly: scope.Named, NameContains: scope.Name}
	records, hasMore, err := s.globalDirectories.CosmicSignatureTokensGlobalPage(ctx, filter, after, limit)
	if err != nil {
		s.logInternal(ctx, "list global cosmic signature tokens", err)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate global token page cardinality", err)
		return internal(), nil
	}

	data := make([]CosmicSignatureToken, 0, len(records))
	previousTokenID := int64(0)
	hasPrevious := false
	if after != nil {
		previousTokenID = after.TokenID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if record.TokenID < 0 || (hasPrevious && record.TokenID >= previousTokenID) {
			s.logInternal(ctx, "validate global token page",
				errors.New("repository returned an unordered global token row"),
				"token_id", record.TokenID)
			return internal(), nil
		}
		token, err := mapGlobalToken(record)
		if err != nil {
			s.logInternal(ctx, "map global token", err, "token_id", record.TokenID)
			return internal(), nil
		}
		data = append(data, token)
		previousTokenID = record.TokenID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list global cosmic signature tokens",
				errors.New("repository reported another page without a cursor row"))
			return internal(), nil
		}
		next, err := encodeGlobalTokenCursor(globalTokenCursor{
			Version: globalTokenCursorVersion,
			Filter:  scope,
			TokenID: previousTokenID,
		})
		if err != nil {
			s.logInternal(ctx, "encode global token cursor", err)
			return internal(), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameCosmicSignatureTokens200JSONResponse{
		CosmicGameCosmicSignatureTokenPageJSONResponse: CosmicGameCosmicSignatureTokenPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetCosmicGameCosmicSignatureToken implements
// GET /api/v2/cosmicgame/cosmic-signature-tokens/{nftTokenId}.
func (s *Server) GetCosmicGameCosmicSignatureToken(
	ctx context.Context,
	request GetCosmicGameCosmicSignatureTokenRequestObject,
) (GetCosmicGameCosmicSignatureTokenResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/cosmic-signature-tokens/%d", request.NftTokenId)
	if request.NftTokenId < 0 {
		return GetCosmicGameCosmicSignatureToken400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				negativeTokenIDProblem(instance),
			),
		}, nil
	}
	record, err := s.globalDirectories.CosmicSignatureTokenDetailV2(ctx, request.NftTokenId)
	if errors.Is(err, store.ErrNotFound) {
		return GetCosmicGameCosmicSignatureToken404ApplicationProblemPlusJSONResponse{
			NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(
				nftTokenNotFoundProblem(instance),
			),
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "get cosmic signature token detail", err,
			"token_id", request.NftTokenId)
		return getGlobalTokenInternal(instance), nil
	}
	if record.TokenID != request.NftTokenId {
		s.logInternal(ctx, "validate cosmic signature token detail",
			errors.New("repository returned another token"),
			"token_id", request.NftTokenId)
		return getGlobalTokenInternal(instance), nil
	}
	detail, err := mapGlobalTokenDetail(record)
	if err != nil {
		s.logInternal(ctx, "map cosmic signature token detail", err,
			"token_id", request.NftTokenId)
		return getGlobalTokenInternal(instance), nil
	}
	return GetCosmicGameCosmicSignatureToken200JSONResponse{
		CosmicGameCosmicSignatureTokenDetailJSONResponse: CosmicGameCosmicSignatureTokenDetailJSONResponse(detail),
	}, nil
}

func getGlobalTokenInternal(instance string) GetCosmicGameCosmicSignatureTokenResponseObject {
	return GetCosmicGameCosmicSignatureToken500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(instance),
		),
	}
}

// negativeTokenIDProblem rejects negative path token IDs; the OpenAPI
// minimum cannot be enforced by the stdlib binder.
func negativeTokenIDProblem(instance string) Problem {
	return newProblem(
		http.StatusBadRequest,
		"invalid-parameter",
		"Invalid parameter",
		"Token ID must be a non-negative integer.",
		instance,
	)
}

// tokenEventPageFetch loads one token-scoped event page from the store.
type tokenEventPageFetch[Record any] func(
	ctx context.Context,
	tokenID int64,
	after *cgstore.TokenEventPageCursor,
	limit int,
) ([]Record, bool, error)

// listTokenEventPage drives one token-scoped, newest-first event resource:
// it validates the limit and cursor, gates on token existence, enforces
// repository scope/order/cardinality and encodes the continuation cursor.
func listTokenEventPage[StoreRecord, APIRecord any](
	ctx context.Context,
	s *Server,
	tokenID int64,
	pathSegment string,
	cursor *Cursor,
	requestedLimit *Limit,
	resource tokenEventResource,
	fetch tokenEventPageFetch[StoreRecord],
	identity func(record StoreRecord) (tokenID int64, eventLogID int64),
	mapRecord func(StoreRecord) (APIRecord, error),
) ([]APIRecord, PageMeta, *Problem) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/cosmic-signature-tokens/%d/%s", tokenID, pathSegment)
	if tokenID < 0 {
		problem := negativeTokenIDProblem(instance)
		return nil, PageMeta{}, &problem
	}
	limit, valid := resolvePageLimit(requestedLimit)
	if !valid {
		problem := newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)
		return nil, PageMeta{}, &problem
	}

	var after *cgstore.TokenEventPageCursor
	if cursor != nil {
		decoded, err := decodeTokenEventCursor(*cursor, tokenID, resource)
		if err != nil {
			problem := newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				invalidTokenCursorProblemMsg,
				instance,
			)
			return nil, PageMeta{}, &problem
		}
		after = &cgstore.TokenEventPageCursor{EventLogID: decoded.EventLogID}
	}

	exists, err := s.globalDirectories.CosmicSignatureTokenExists(ctx, tokenID)
	if err != nil {
		s.logInternal(ctx, "check token existence for token events", err,
			"resource", string(resource),
			"token_id", tokenID)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}
	if !exists {
		problem := nftTokenNotFoundProblem(instance)
		return nil, PageMeta{}, &problem
	}

	records, hasMore, err := fetch(ctx, tokenID, after, limit)
	if err != nil {
		s.logInternal(ctx, "list token events", err,
			"resource", string(resource),
			"token_id", tokenID)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate token event page cardinality", err,
			"resource", string(resource),
			"token_id", tokenID)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}

	data := make([]APIRecord, 0, len(records))
	previousEventLogID := int64(0)
	hasPrevious := false
	if after != nil {
		previousEventLogID = after.EventLogID
		hasPrevious = true
	}
	for i := range records {
		recordTokenID, eventLogID := identity(records[i])
		if recordTokenID != tokenID || eventLogID < 1 ||
			(hasPrevious && eventLogID >= previousEventLogID) {
			s.logInternal(ctx, "validate token event page",
				errors.New("repository returned an out-of-scope or unordered token event row"),
				"resource", string(resource),
				"token_id", tokenID,
				"event_log_id", eventLogID)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		mapped, err := mapRecord(records[i])
		if err != nil {
			s.logInternal(ctx, "map token event row", err,
				"resource", string(resource),
				"token_id", tokenID,
				"event_log_id", eventLogID)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		data = append(data, mapped)
		previousEventLogID = eventLogID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list token events",
				errors.New("repository reported another page without a cursor row"),
				"resource", string(resource),
				"token_id", tokenID)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		next, err := encodeTokenEventCursor(tokenEventCursor{
			Version:    tokenEventCursorVersion,
			Resource:   resource,
			TokenID:    tokenID,
			EventLogID: previousEventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode token event cursor", err,
				"resource", string(resource),
				"token_id", tokenID)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		meta.NextCursor = &next
	}
	return data, meta, nil
}

// ListCosmicGameCosmicSignatureTokenNameHistory implements
// GET /api/v2/cosmicgame/cosmic-signature-tokens/{nftTokenId}/name-history.
func (s *Server) ListCosmicGameCosmicSignatureTokenNameHistory(
	ctx context.Context,
	request ListCosmicGameCosmicSignatureTokenNameHistoryRequestObject,
) (ListCosmicGameCosmicSignatureTokenNameHistoryResponseObject, error) {
	data, meta, problem := listTokenEventPage(
		ctx,
		s,
		request.NftTokenId,
		"name-history",
		request.Params.Cursor,
		request.Params.Limit,
		tokenEventResourceNameHistory,
		s.globalDirectories.TokenNameHistoryPage,
		func(record cgstore.TokenNameChangeRecord) (int64, int64) {
			return record.TokenID, record.Tx.EvtLogId
		},
		mapTokenNameChange,
	)
	if problem != nil {
		switch problem.Status {
		case http.StatusBadRequest:
			return ListCosmicGameCosmicSignatureTokenNameHistory400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		case http.StatusNotFound:
			return ListCosmicGameCosmicSignatureTokenNameHistory404ApplicationProblemPlusJSONResponse{
				NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(*problem),
			}, nil
		default:
			return ListCosmicGameCosmicSignatureTokenNameHistory500ApplicationProblemPlusJSONResponse{
				InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
	}
	return ListCosmicGameCosmicSignatureTokenNameHistory200JSONResponse{
		CosmicGameTokenNameChangePageJSONResponse: CosmicGameTokenNameChangePageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameCosmicSignatureTokenTransfers implements
// GET /api/v2/cosmicgame/cosmic-signature-tokens/{nftTokenId}/transfers.
func (s *Server) ListCosmicGameCosmicSignatureTokenTransfers(
	ctx context.Context,
	request ListCosmicGameCosmicSignatureTokenTransfersRequestObject,
) (ListCosmicGameCosmicSignatureTokenTransfersResponseObject, error) {
	data, meta, problem := listTokenEventPage(
		ctx,
		s,
		request.NftTokenId,
		"transfers",
		request.Params.Cursor,
		request.Params.Limit,
		tokenEventResourceTransfers,
		s.globalDirectories.TokenTransfersPage,
		func(record cgstore.TokenTransferRecord) (int64, int64) {
			return record.TokenID, record.Tx.EvtLogId
		},
		mapTokenTransfer,
	)
	if problem != nil {
		switch problem.Status {
		case http.StatusBadRequest:
			return ListCosmicGameCosmicSignatureTokenTransfers400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		case http.StatusNotFound:
			return ListCosmicGameCosmicSignatureTokenTransfers404ApplicationProblemPlusJSONResponse{
				NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(*problem),
			}, nil
		default:
			return ListCosmicGameCosmicSignatureTokenTransfers500ApplicationProblemPlusJSONResponse{
				InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
	}
	return ListCosmicGameCosmicSignatureTokenTransfers200JSONResponse{
		CosmicGameCosmicSignatureTokenTransferPageJSONResponse: CosmicGameCosmicSignatureTokenTransferPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameCosmicSignatureHolders implements
// GET /api/v2/cosmicgame/cosmic-signature-tokens/holders.
func (s *Server) ListCosmicGameCosmicSignatureHolders(
	ctx context.Context,
	request ListCosmicGameCosmicSignatureHoldersRequestObject,
) (ListCosmicGameCosmicSignatureHoldersResponseObject, error) {
	limit, after, problem := participantPageInput(
		request.Params.Cursor, request.Params.Limit, cgstore.ParticipantCsTokenHolders, csHoldersInstance,
	)
	if problem != nil {
		return ListCosmicGameCosmicSignatureHolders400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	records, hasMore, err := s.globalDirectories.CosmicSignatureHoldersPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list cosmic signature holders", err)
		return ListCosmicGameCosmicSignatureHolders500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(csHoldersInstance),
			),
		}, nil
	}
	data, meta, err := buildParticipantPage(
		records, hasMore, cgstore.ParticipantCsTokenHolders, after, limit,
		func(record cgstore.CosmicSignatureHolderRecord) cgstore.ParticipantPageCursor {
			return cgstore.ParticipantPageCursor{
				Kind:      cgstore.ParticipantCsTokenHolders,
				SortValue: strconv.FormatInt(record.TokenCount, 10),
				AddressID: record.OwnerAid,
			}
		},
		mapCosmicSignatureHolder,
	)
	if err != nil {
		s.logInternal(ctx, "build cosmic signature holder page", err)
		return ListCosmicGameCosmicSignatureHolders500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(csHoldersInstance),
			),
		}, nil
	}
	return ListCosmicGameCosmicSignatureHolders200JSONResponse{
		CosmicGameCosmicSignatureHolderPageJSONResponse: CosmicGameCosmicSignatureHolderPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameCosmicTokenHolders implements
// GET /api/v2/cosmicgame/cosmic-token/holders.
func (s *Server) ListCosmicGameCosmicTokenHolders(
	ctx context.Context,
	request ListCosmicGameCosmicTokenHoldersRequestObject,
) (ListCosmicGameCosmicTokenHoldersResponseObject, error) {
	limit, after, problem := participantPageInput(
		request.Params.Cursor, request.Params.Limit, cgstore.ParticipantCosmicTokenHolders, cosmicTokenHoldersInstance,
	)
	if problem != nil {
		return ListCosmicGameCosmicTokenHolders400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	records, hasMore, err := s.globalDirectories.CosmicTokenHoldersPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list cosmic token holders", err)
		return ListCosmicGameCosmicTokenHolders500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(cosmicTokenHoldersInstance),
			),
		}, nil
	}
	data, meta, err := buildParticipantPage(
		records, hasMore, cgstore.ParticipantCosmicTokenHolders, after, limit,
		func(record cgstore.CosmicTokenHolderRecord) cgstore.ParticipantPageCursor {
			return cgstore.ParticipantPageCursor{
				Kind:      cgstore.ParticipantCosmicTokenHolders,
				SortValue: record.BalanceWei,
				AddressID: record.OwnerAid,
			}
		},
		mapCosmicTokenHolder,
	)
	if err != nil {
		s.logInternal(ctx, "build cosmic token holder page", err)
		return ListCosmicGameCosmicTokenHolders500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(cosmicTokenHoldersInstance),
			),
		}, nil
	}
	return ListCosmicGameCosmicTokenHolders200JSONResponse{
		CosmicGameCosmicTokenHolderPageJSONResponse: CosmicGameCosmicTokenHolderPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetCosmicGameCosmicTokenStatistics implements
// GET /api/v2/cosmicgame/cosmic-token/statistics.
func (s *Server) GetCosmicGameCosmicTokenStatistics(
	ctx context.Context,
	_ GetCosmicGameCosmicTokenStatisticsRequestObject,
) (GetCosmicGameCosmicTokenStatisticsResponseObject, error) {
	record, err := s.globalDirectories.CosmicTokenStatisticsV2(ctx)
	if err != nil {
		s.logInternal(ctx, "get cosmic token statistics", err)
		return getCosmicTokenStatisticsInternal(), nil
	}
	statistics, err := mapCosmicTokenStatistics(record)
	if err != nil {
		s.logInternal(ctx, "map cosmic token statistics", err)
		return getCosmicTokenStatisticsInternal(), nil
	}
	return GetCosmicGameCosmicTokenStatistics200JSONResponse{
		CosmicGameCosmicTokenStatisticsJSONResponse: CosmicGameCosmicTokenStatisticsJSONResponse(statistics),
	}, nil
}

func getCosmicTokenStatisticsInternal() GetCosmicGameCosmicTokenStatisticsResponseObject {
	return GetCosmicGameCosmicTokenStatistics500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(cosmicTokenStatsInstance),
		),
	}
}

// ListCosmicGameCosmicTokenSupplyByBid implements
// GET /api/v2/cosmicgame/cosmic-token/supply-history/by-bid.
func (s *Server) ListCosmicGameCosmicTokenSupplyByBid(
	ctx context.Context,
	request ListCosmicGameCosmicTokenSupplyByBidRequestObject,
) (ListCosmicGameCosmicTokenSupplyByBidResponseObject, error) {
	badRequest := func(problem Problem) ListCosmicGameCosmicTokenSupplyByBidResponseObject {
		return ListCosmicGameCosmicTokenSupplyByBid400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListCosmicGameCosmicTokenSupplyByBidResponseObject {
		return ListCosmicGameCosmicTokenSupplyByBid500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(supplyByBidInstance),
			),
		}
	}

	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			supplyByBidInstance,
		)), nil
	}
	var after *cgstore.SupplyChangePageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeSupplyChangeCursor(*request.Params.Cursor)
		if err != nil {
			return badRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another resource.",
				supplyByBidInstance,
			)), nil
		}
		after = &cgstore.SupplyChangePageCursor{EventLogID: cursor.EventLogID}
	}

	records, hasMore, err := s.globalDirectories.CosmicTokenSupplyByBidPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list cosmic token supply by bid", err)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate supply-by-bid page cardinality", err)
		return internal(), nil
	}

	data := make([]CosmicTokenSupplyChange, 0, len(records))
	previousEventLogID := int64(0)
	hasPrevious := false
	if after != nil {
		previousEventLogID = after.EventLogID
		hasPrevious = true
	}
	var previousTotal *big.Int
	for i := range records {
		record := records[i]
		if record.Tx.EvtLogId < 1 || (hasPrevious && record.Tx.EvtLogId <= previousEventLogID) {
			s.logInternal(ctx, "validate supply-by-bid page",
				errors.New("repository returned an unordered supply-change row"),
				"event_log_id", record.Tx.EvtLogId)
			return internal(), nil
		}
		change, err := mapSupplyChange(record)
		if err != nil {
			s.logInternal(ctx, "map supply change", err, "event_log_id", record.Tx.EvtLogId)
			return internal(), nil
		}
		total, _ := new(big.Int).SetString(change.TotalSupplyWei, 10)
		if previousTotal != nil {
			net, _ := new(big.Int).SetString(change.NetWei, 10)
			if new(big.Int).Add(previousTotal, net).Cmp(total) != 0 {
				s.logInternal(ctx, "validate supply-by-bid page",
					errors.New("running supply diverges from the previous row"),
					"event_log_id", record.Tx.EvtLogId)
				return internal(), nil
			}
		}
		data = append(data, change)
		previousEventLogID = record.Tx.EvtLogId
		hasPrevious = true
		previousTotal = total
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list cosmic token supply by bid",
				errors.New("repository reported another page without a cursor row"))
			return internal(), nil
		}
		next, err := encodeSupplyChangeCursor(supplyChangeCursor{
			Version:    supplyChangeCursorVersion,
			EventLogID: previousEventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode supply change cursor", err)
			return internal(), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameCosmicTokenSupplyByBid200JSONResponse{
		CosmicGameCosmicTokenSupplyByBidPageJSONResponse: CosmicGameCosmicTokenSupplyByBidPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameCosmicTokenSupplyDaily implements
// GET /api/v2/cosmicgame/cosmic-token/supply-history/daily.
func (s *Server) ListCosmicGameCosmicTokenSupplyDaily(
	ctx context.Context,
	request ListCosmicGameCosmicTokenSupplyDailyRequestObject,
) (ListCosmicGameCosmicTokenSupplyDailyResponseObject, error) {
	badRequest := func(detail string) ListCosmicGameCosmicTokenSupplyDailyResponseObject {
		return ListCosmicGameCosmicTokenSupplyDaily400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				newProblem(
					http.StatusBadRequest,
					"invalid-parameter",
					"Invalid parameter",
					detail,
					supplyDailyInstance,
				),
			),
		}
	}
	internal := func() ListCosmicGameCosmicTokenSupplyDailyResponseObject {
		return ListCosmicGameCosmicTokenSupplyDaily500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(supplyDailyInstance),
			),
		}
	}

	from := request.Params.From.Time
	to := request.Params.To.Time
	if !from.Before(to) {
		return badRequest("The window is empty: from must be before to."), nil
	}
	if to.Sub(from) > cgstore.MaxSupplyDailyWindowDays*24*time.Hour {
		return badRequest(fmt.Sprintf(
			"The window may span at most %d days.", cgstore.MaxSupplyDailyWindowDays)), nil
	}

	records, err := s.globalDirectories.CosmicTokenSupplyDaily(ctx, from, to)
	if err != nil {
		s.logInternal(ctx, "list cosmic token supply daily", err)
		return internal(), nil
	}
	if len(records) > cgstore.MaxSupplyDailyWindowDays {
		s.logInternal(ctx, "validate daily supply rows",
			errors.New("repository returned more daily rows than the window allows"))
		return internal(), nil
	}

	data := make([]CosmicTokenDailySupply, 0, len(records))
	previousDate := ""
	for i := range records {
		record := records[i]
		if record.Date <= previousDate {
			s.logInternal(ctx, "validate daily supply rows",
				errors.New("repository returned unordered daily rows"),
				"date", record.Date)
			return internal(), nil
		}
		row, err := mapDailySupply(record)
		if err != nil {
			s.logInternal(ctx, "map daily supply row", err, "date", record.Date)
			return internal(), nil
		}
		day := row.Date.Time
		if day.Before(from) || !day.Before(to) {
			s.logInternal(ctx, "validate daily supply rows",
				errors.New("repository returned a day outside the window"),
				"date", record.Date)
			return internal(), nil
		}
		data = append(data, row)
		previousDate = record.Date
	}

	return ListCosmicGameCosmicTokenSupplyDaily200JSONResponse{
		CosmicGameCosmicTokenSupplyDailyJSONResponse: CosmicGameCosmicTokenSupplyDailyJSONResponse{
			Data: data,
		},
	}, nil
}

// ListCosmicGameMarketingRewards implements
// GET /api/v2/cosmicgame/marketing-rewards.
func (s *Server) ListCosmicGameMarketingRewards(
	ctx context.Context,
	request ListCosmicGameMarketingRewardsRequestObject,
) (ListCosmicGameMarketingRewardsResponseObject, error) {
	badRequest := func(problem Problem) ListCosmicGameMarketingRewardsResponseObject {
		return ListCosmicGameMarketingRewards400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListCosmicGameMarketingRewardsResponseObject {
		return ListCosmicGameMarketingRewards500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(globalMarketingInstance),
			),
		}
	}

	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			globalMarketingInstance,
		)), nil
	}
	var after *cgstore.UserEventPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeGlobalMarketingCursor(*request.Params.Cursor)
		if err != nil {
			return badRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another resource.",
				globalMarketingInstance,
			)), nil
		}
		after = &cgstore.UserEventPageCursor{EventLogID: cursor.EventLogID}
	}

	records, hasMore, err := s.globalDirectories.MarketingRewardsGlobalPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list global marketing rewards", err)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate global marketing page cardinality", err)
		return internal(), nil
	}

	data := make([]MarketingReward, 0, len(records))
	previousEventLogID := int64(0)
	hasPrevious := false
	if after != nil {
		previousEventLogID = after.EventLogID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if record.Tx.EvtLogId < 1 || (hasPrevious && record.Tx.EvtLogId >= previousEventLogID) {
			s.logInternal(ctx, "validate global marketing page",
				errors.New("repository returned an unordered marketing reward row"),
				"event_log_id", record.Tx.EvtLogId)
			return internal(), nil
		}
		reward, err := mapMarketingReward(record)
		if err != nil {
			s.logInternal(ctx, "map global marketing reward", err,
				"event_log_id", record.Tx.EvtLogId)
			return internal(), nil
		}
		data = append(data, reward)
		previousEventLogID = record.Tx.EvtLogId
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list global marketing rewards",
				errors.New("repository reported another page without a cursor row"))
			return internal(), nil
		}
		next, err := encodeGlobalMarketingCursor(globalMarketingCursor{
			Version:    globalMarketingCursorVersion,
			EventLogID: previousEventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode global marketing cursor", err)
			return internal(), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameMarketingRewards200JSONResponse{
		CosmicGameMarketingRewardPageJSONResponse: CosmicGameMarketingRewardPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}
