package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// GetCosmicGameUser implements GET /api/v2/cosmicgame/users/{address}.
func (s *Server) GetCosmicGameUser(
	ctx context.Context,
	request GetCosmicGameUserRequestObject,
) (GetCosmicGameUserResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s", request.Address)
	address, _, valid := userAddressInput(request.Address)
	if !valid {
		return GetCosmicGameUser400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				invalidUserAddressProblem(instance),
			),
		}, nil
	}

	userAid, err := s.users.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return GetCosmicGameUser200JSONResponse{
			CosmicGameUserProfileJSONResponse: CosmicGameUserProfileJSONResponse(
				zeroUserProfile(address),
			),
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address", err, "address", address)
		return GetCosmicGameUser500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}, nil
	}

	record, err := s.users.UserProfile(ctx, userAid)
	if err != nil {
		s.logInternal(ctx, "get user profile", err, "address", address)
		return GetCosmicGameUser500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}, nil
	}
	profile, err := mapUserProfile(record)
	if err != nil || !strings.EqualFold(profile.Address, address) {
		if err == nil {
			err = errors.New("repository returned a profile for another address")
		}
		s.logInternal(ctx, "map user profile", err, "address", address)
		return GetCosmicGameUser500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}, nil
	}
	return GetCosmicGameUser200JSONResponse{
		CosmicGameUserProfileJSONResponse: CosmicGameUserProfileJSONResponse(profile),
	}, nil
}

// ListCosmicGameUserBids implements
// GET /api/v2/cosmicgame/users/{address}/bids.
func (s *Server) ListCosmicGameUserBids(
	ctx context.Context,
	request ListCosmicGameUserBidsRequestObject,
) (ListCosmicGameUserBidsResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/bids", request.Address)
	address, scope, valid := userAddressInput(request.Address)
	if !valid {
		return userBidsBadRequest(invalidUserAddressProblem(instance)), nil
	}

	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return userBidsBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	var after *cgstore.UserBidPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeUserBidCursor(*request.Params.Cursor, scope)
		if err != nil {
			return userBidsBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another wallet.",
				instance,
			)), nil
		}
		after = &cgstore.UserBidPageCursor{EventLogID: cursor.EventLogID}
	}

	userAid, err := s.users.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return ListCosmicGameUserBids200JSONResponse{
			CosmicGameUserBidPageJSONResponse: CosmicGameUserBidPageJSONResponse{
				Data: []Bid{},
				Meta: PageMeta{Limit: limit},
			},
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for bids", err, "address", address)
		return userBidsInternal(instance), nil
	}

	records, hasMore, err := s.users.BidsByUserPage(ctx, userAid, after, limit)
	if err != nil {
		s.logInternal(ctx, "list user bids", err, "address", address)
		return userBidsInternal(instance), nil
	}
	if len(records) > limit {
		s.logInternal(ctx, "validate user bid page",
			errors.New("repository returned more user bids than requested"),
			"address", address)
		return userBidsInternal(instance), nil
	}

	data := make([]Bid, 0, len(records))
	var previousEventLogID int64
	if after != nil {
		previousEventLogID = after.EventLogID
	}
	for i := range records {
		eventLogID := records[i].Tx.EvtLogId
		if records[i].BidderAid != userAid ||
			!strings.EqualFold(records[i].BidderAddr, address) ||
			eventLogID < 1 ||
			(previousEventLogID > 0 && eventLogID >= previousEventLogID) {
			s.logInternal(ctx, "validate user bid page",
				errors.New("repository returned an out-of-scope or unordered user bid"),
				"address", address,
				"event_log_id", eventLogID)
			return userBidsInternal(instance), nil
		}
		bid, err := mapBid(records[i])
		if err != nil {
			s.logInternal(ctx, "map user bid", err,
				"address", address,
				"event_log_id", eventLogID)
			return userBidsInternal(instance), nil
		}
		data = append(data, bid)
		previousEventLogID = eventLogID
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 {
			s.logInternal(ctx, "list user bids",
				errors.New("repository reported another page without a cursor row"),
				"address", address)
			return userBidsInternal(instance), nil
		}
		next, err := encodeUserBidCursor(userBidCursor{
			Version:    userBidCursorVersion,
			Address:    scope,
			EventLogID: records[len(records)-1].Tx.EvtLogId,
		})
		if err != nil {
			s.logInternal(ctx, "encode user bid cursor", err, "address", address)
			return userBidsInternal(instance), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameUserBids200JSONResponse{
		CosmicGameUserBidPageJSONResponse: CosmicGameUserBidPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func userAddressInput(raw string) (checksum string, scope string, valid bool) {
	if !strings.HasPrefix(raw, "0x") || len(raw) != 42 || !ethcommon.IsHexAddress(raw) {
		return "", "", false
	}
	checksum = ethcommon.HexToAddress(raw).Hex()
	scope, valid = normalizedUserAddress(checksum)
	return checksum, scope, valid
}

func invalidUserAddressProblem(instance string) Problem {
	return newProblem(
		http.StatusBadRequest,
		"invalid-parameter",
		"Invalid parameter",
		"Address must be a 20-byte 0x-prefixed hexadecimal Ethereum address.",
		instance,
	)
}

func userBidsBadRequest(problem Problem) ListCosmicGameUserBids400ApplicationProblemPlusJSONResponse {
	return ListCosmicGameUserBids400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func userBidsInternal(instance string) ListCosmicGameUserBids500ApplicationProblemPlusJSONResponse {
	return ListCosmicGameUserBids500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(instance),
		),
	}
}
