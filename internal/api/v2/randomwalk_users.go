package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	rwstore "github.com/PredictionExplorer/augur-explorer/internal/store/randomwalk"
)

// GetRandomWalkUser implements GET /api/v2/randomwalk/users/{address}.
func (s *Server) GetRandomWalkUser(
	ctx context.Context,
	request GetRandomWalkUserRequestObject,
) (GetRandomWalkUserResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/randomwalk/users/%s", request.Address)
	internal := func() GetRandomWalkUserResponseObject {
		return GetRandomWalkUser500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}
	}
	address, _, valid := userAddressInput(request.Address)
	if !valid {
		return GetRandomWalkUser400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				invalidUserAddressProblem(instance),
			),
		}, nil
	}

	userAid, err := s.randomWalk.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return GetRandomWalkUser200JSONResponse{
			RandomWalkUserProfileJSONResponse: RandomWalkUserProfileJSONResponse(
				zeroRandomWalkUserProfile(address),
			),
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve random walk user address", err, "address", address)
		return internal(), nil
	}

	record, err := s.randomWalk.UserProfileV2(ctx, userAid)
	if err != nil {
		s.logInternal(ctx, "get random walk user profile", err, "address", address)
		return internal(), nil
	}
	profile, err := mapRandomWalkUserProfile(record)
	if err != nil || !strings.EqualFold(profile.Address, address) {
		if err == nil {
			err = errors.New("repository returned a profile for another address")
		}
		s.logInternal(ctx, "map random walk user profile", err, "address", address)
		return internal(), nil
	}
	return GetRandomWalkUser200JSONResponse{
		RandomWalkUserProfileJSONResponse: RandomWalkUserProfileJSONResponse(profile),
	}, nil
}

// ListRandomWalkUserTokens implements
// GET /api/v2/randomwalk/users/{address}/tokens.
func (s *Server) ListRandomWalkUserTokens(
	ctx context.Context,
	request ListRandomWalkUserTokensRequestObject,
) (ListRandomWalkUserTokensResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/randomwalk/users/%s/tokens", request.Address)
	badRequest := func(problem Problem) ListRandomWalkUserTokensResponseObject {
		return ListRandomWalkUserTokens400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListRandomWalkUserTokensResponseObject {
		return ListRandomWalkUserTokens500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}
	}
	address, scope, valid := userAddressInput(request.Address)
	if !valid {
		return badRequest(invalidUserAddressProblem(instance)), nil
	}
	limit, validLimit := resolvePageLimit(request.Params.Limit)
	if !validLimit {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}
	var after *rwstore.TokenPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeRandomWalkUserTokenCursor(*request.Params.Cursor, scope)
		if err != nil {
			return badRequest(invalidRandomWalkCursorProblem(instance)), nil
		}
		after = &rwstore.TokenPageCursor{TokenID: cursor.TokenID}
	}

	userAid, err := s.randomWalk.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return ListRandomWalkUserTokens200JSONResponse{
			RandomWalkOwnedTokenPageJSONResponse: RandomWalkOwnedTokenPageJSONResponse{
				Data: []RandomWalkOwnedToken{},
				Meta: PageMeta{Limit: limit},
			},
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve random walk user for tokens", err, "address", address)
		return internal(), nil
	}

	records, hasMore, err := s.randomWalk.UserTokensPage(ctx, userAid, after, limit)
	if err != nil {
		s.logInternal(ctx, "list random walk user tokens", err, "address", address)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate random walk user token cardinality", err, "address", address)
		return internal(), nil
	}

	data := make([]RandomWalkOwnedToken, 0, len(records))
	previousTokenID := int64(0)
	hasPrevious := false
	if after != nil {
		previousTokenID = after.TokenID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if record.TokenID < 0 || (hasPrevious && record.TokenID <= previousTokenID) {
			s.logInternal(ctx, "validate random walk user token page",
				errors.New("repository returned an unordered owned token"),
				"address", address,
				"token_id", record.TokenID)
			return internal(), nil
		}
		token, err := mapRandomWalkOwnedToken(record)
		if err != nil {
			s.logInternal(ctx, "map random walk owned token", err,
				"address", address,
				"token_id", record.TokenID)
			return internal(), nil
		}
		data = append(data, token)
		previousTokenID = record.TokenID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if !hasPrevious {
			s.logInternal(ctx, "list random walk user tokens",
				errors.New("repository reported another page without a cursor row"),
				"address", address)
			return internal(), nil
		}
		next, err := encodeRandomWalkUserTokenCursor(randomWalkUserTokenCursor{
			Version:  randomWalkCursorVersion,
			Resource: randomWalkResourceUserTokens,
			Address:  scope,
			TokenID:  previousTokenID,
		})
		if err != nil {
			s.logInternal(ctx, "encode random walk user token cursor", err, "address", address)
			return internal(), nil
		}
		meta.NextCursor = &next
	}

	return ListRandomWalkUserTokens200JSONResponse{
		RandomWalkOwnedTokenPageJSONResponse: RandomWalkOwnedTokenPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListRandomWalkUserOffers implements
// GET /api/v2/randomwalk/users/{address}/offers.
func (s *Server) ListRandomWalkUserOffers(
	ctx context.Context,
	request ListRandomWalkUserOffersRequestObject,
) (ListRandomWalkUserOffersResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/randomwalk/users/%s/offers", request.Address)
	badRequest := func(problem Problem) ListRandomWalkUserOffersResponseObject {
		return ListRandomWalkUserOffers400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListRandomWalkUserOffersResponseObject {
		return ListRandomWalkUserOffers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}
	}
	address, scope, valid := userAddressInput(request.Address)
	if !valid {
		return badRequest(invalidUserAddressProblem(instance)), nil
	}

	userAid, err := s.randomWalk.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		limit, validLimit := resolvePageLimit(request.Params.Limit)
		if !validLimit {
			return badRequest(newProblem(
				http.StatusBadRequest,
				"invalid-parameter",
				"Invalid parameter",
				pageLimitProblemDetail(),
				instance,
			)), nil
		}
		if request.Params.Cursor != nil {
			if _, err := decodeRandomWalkLedgerCursor(
				*request.Params.Cursor, randomWalkResourceUserOffers, scope); err != nil {
				return badRequest(invalidRandomWalkCursorProblem(instance)), nil
			}
		}
		return ListRandomWalkUserOffers200JSONResponse{
			RandomWalkOfferHistoryPageJSONResponse: RandomWalkOfferHistoryPageJSONResponse{
				Data: []RandomWalkOfferHistoryEntry{},
				Meta: PageMeta{Limit: limit},
			},
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve random walk user for offers", err, "address", address)
		return internal(), nil
	}

	data, meta, problem := listRandomWalkLedgerPage(
		ctx,
		s,
		instance,
		request.Params.Cursor,
		request.Params.Limit,
		randomWalkResourceUserOffers,
		scope,
		func(ctx context.Context, after *rwstore.EventPageCursor, limit int) ([]rwstore.OfferHistoryRecord, bool, error) {
			return s.randomWalk.UserOffersPage(ctx, userAid, after, limit)
		},
		func(record rwstore.OfferHistoryRecord) int64 { return record.ListTx.EvtLogID },
		func(record rwstore.OfferHistoryRecord) (RandomWalkOfferHistoryEntry, error) {
			entry, err := mapRandomWalkOfferHistoryEntry(record)
			if err != nil {
				return RandomWalkOfferHistoryEntry{}, err
			}
			if !randomWalkOfferInvolvesWallet(record, address) {
				return RandomWalkOfferHistoryEntry{}, errors.New(
					"repository returned an offer outside the wallet scope")
			}
			return entry, nil
		},
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return badRequest(*problem), nil
		}
		return internal(), nil
	}
	return ListRandomWalkUserOffers200JSONResponse{
		RandomWalkOfferHistoryPageJSONResponse: RandomWalkOfferHistoryPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// randomWalkOfferInvolvesWallet reports whether the wallet appears on the
// maker or the recorded purchase side of the offer.
func randomWalkOfferInvolvesWallet(record rwstore.OfferHistoryRecord, address string) bool {
	if strings.EqualFold(record.MakerAddr, address) {
		return true
	}
	if record.Purchase != nil {
		return strings.EqualFold(record.Purchase.BuyerAddr, address) ||
			strings.EqualFold(record.Purchase.SellerAddr, address)
	}
	return false
}
