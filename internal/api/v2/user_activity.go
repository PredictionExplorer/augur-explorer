package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ListCosmicGameUserCosmicSignatureTokens implements
// GET /api/v2/cosmicgame/users/{address}/cosmic-signature-tokens.
func (s *Server) ListCosmicGameUserCosmicSignatureTokens(
	ctx context.Context,
	request ListCosmicGameUserCosmicSignatureTokensRequestObject,
) (ListCosmicGameUserCosmicSignatureTokensResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/cosmic-signature-tokens", request.Address)
	address, scope, valid := userAddressInput(request.Address)
	if !valid {
		return userOwnedTokensBadRequest(invalidUserAddressProblem(instance)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return userOwnedTokensBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	var after *cgstore.UserTokenPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeUserOwnedTokenCursor(*request.Params.Cursor, scope)
		if err != nil {
			return userOwnedTokensBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another wallet or resource.",
				instance,
			)), nil
		}
		after = &cgstore.UserTokenPageCursor{TokenID: cursor.TokenID}
	}

	userAid, err := s.userActivity.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return ListCosmicGameUserCosmicSignatureTokens200JSONResponse{
			CosmicGameUserCosmicSignatureTokenPageJSONResponse: CosmicGameUserCosmicSignatureTokenPageJSONResponse{
				Data: []UserCosmicSignatureToken{},
				Meta: PageMeta{Limit: limit},
			},
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for owned tokens", err, "address", address)
		return userOwnedTokensInternal(instance), nil
	}

	records, hasMore, err := s.userActivity.UserCosmicSignatureTokensPage(ctx, userAid, after, limit)
	if err != nil {
		s.logInternal(ctx, "list user owned tokens", err, "address", address)
		return userOwnedTokensInternal(instance), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate user owned token page cardinality", err, "address", address)
		return userOwnedTokensInternal(instance), nil
	}

	data := make([]UserCosmicSignatureToken, 0, len(records))
	previousTokenID := int64(0)
	hasPrevious := false
	if after != nil {
		previousTokenID = after.TokenID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if record.OwnerAid != userAid || record.TokenID < 0 ||
			(hasPrevious && record.TokenID <= previousTokenID) {
			s.logInternal(ctx, "validate user owned token page",
				errors.New("repository returned an out-of-scope or unordered owned token"),
				"address", address,
				"token_id", record.TokenID)
			return userOwnedTokensInternal(instance), nil
		}
		token, err := mapUserOwnedToken(record)
		if err != nil {
			s.logInternal(ctx, "map user owned token", err,
				"address", address,
				"token_id", record.TokenID)
			return userOwnedTokensInternal(instance), nil
		}
		data = append(data, token)
		previousTokenID = record.TokenID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list user owned tokens",
				errors.New("repository reported another page without a cursor row"),
				"address", address)
			return userOwnedTokensInternal(instance), nil
		}
		next, err := encodeUserOwnedTokenCursor(userOwnedTokenCursor{
			Version: userOwnedTokenCursorVersion,
			Address: scope,
			TokenID: previousTokenID,
		})
		if err != nil {
			s.logInternal(ctx, "encode user owned token cursor", err, "address", address)
			return userOwnedTokensInternal(instance), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameUserCosmicSignatureTokens200JSONResponse{
		CosmicGameUserCosmicSignatureTokenPageJSONResponse: CosmicGameUserCosmicSignatureTokenPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// transferPageIdentity builds the listUserEventPage identity check shared
// by the two transfer ledgers: the row must involve the wallet and its
// precomputed direction must agree with the counterparty columns.
func transferPageIdentity(
	fromAid, toAid int64,
	direction cgstore.UserTransferDirection,
	userAid int64,
	eventLogID int64,
) (bool, int64) {
	involved := fromAid == userAid || toAid == userAid
	var expected cgstore.UserTransferDirection
	switch {
	case fromAid == userAid && toAid == userAid:
		expected = cgstore.UserTransferSelf
	case fromAid == userAid:
		expected = cgstore.UserTransferOut
	default:
		expected = cgstore.UserTransferIn
	}
	return involved && direction == expected, eventLogID
}

// ListCosmicGameUserCosmicSignatureTransfers implements
// GET /api/v2/cosmicgame/users/{address}/cosmic-signature-transfers.
func (s *Server) ListCosmicGameUserCosmicSignatureTransfers(
	ctx context.Context,
	request ListCosmicGameUserCosmicSignatureTransfersRequestObject,
) (ListCosmicGameUserCosmicSignatureTransfersResponseObject, error) {
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"cosmic-signature-transfers",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceCsTransfers,
		s.userActivity.UserAddressID,
		s.userActivity.UserCosmicSignatureTransfersPage,
		func(record cgstore.UserCosmicSignatureTransferRecord, userAid int64, _ string) (bool, int64) {
			return transferPageIdentity(
				record.FromAid, record.ToAid, record.Direction, userAid, record.Tx.EvtLogId)
		},
		mapUserCosmicSignatureTransfer,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserCosmicSignatureTransfers400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserCosmicSignatureTransfers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserCosmicSignatureTransfers200JSONResponse{
		CosmicGameUserCosmicSignatureTransferPageJSONResponse: CosmicGameUserCosmicSignatureTransferPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserCosmicTokenTransfers implements
// GET /api/v2/cosmicgame/users/{address}/cosmic-token-transfers.
func (s *Server) ListCosmicGameUserCosmicTokenTransfers(
	ctx context.Context,
	request ListCosmicGameUserCosmicTokenTransfersRequestObject,
) (ListCosmicGameUserCosmicTokenTransfersResponseObject, error) {
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"cosmic-token-transfers",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceCtTransfers,
		s.userActivity.UserAddressID,
		s.userActivity.UserCosmicTokenTransfersPage,
		func(record cgstore.UserCosmicTokenTransferRecord, userAid int64, _ string) (bool, int64) {
			return transferPageIdentity(
				record.FromAid, record.ToAid, record.Direction, userAid, record.Tx.EvtLogId)
		},
		mapUserCosmicTokenTransfer,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserCosmicTokenTransfers400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserCosmicTokenTransfers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserCosmicTokenTransfers200JSONResponse{
		CosmicGameUserCosmicTokenTransferPageJSONResponse: CosmicGameUserCosmicTokenTransferPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserMarketingRewards implements
// GET /api/v2/cosmicgame/users/{address}/marketing-rewards.
func (s *Server) ListCosmicGameUserMarketingRewards(
	ctx context.Context,
	request ListCosmicGameUserMarketingRewardsRequestObject,
) (ListCosmicGameUserMarketingRewardsResponseObject, error) {
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"marketing-rewards",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceMarketingRewards,
		s.userActivity.UserAddressID,
		s.userActivity.UserMarketingRewardsPage,
		func(record cgstore.UserMarketingRewardRecord, userAid int64, _ string) (bool, int64) {
			return record.MarketerAid == userAid, record.Tx.EvtLogId
		},
		mapUserMarketingReward,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserMarketingRewards400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserMarketingRewards500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserMarketingRewards200JSONResponse{
		CosmicGameUserMarketingRewardPageJSONResponse: CosmicGameUserMarketingRewardPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetCosmicGameUserCosmicTokenSummary implements
// GET /api/v2/cosmicgame/users/{address}/cosmic-token-summary.
func (s *Server) GetCosmicGameUserCosmicTokenSummary(
	ctx context.Context,
	request GetCosmicGameUserCosmicTokenSummaryRequestObject,
) (GetCosmicGameUserCosmicTokenSummaryResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/cosmic-token-summary", request.Address)
	address, _, valid := userAddressInput(request.Address)
	if !valid {
		return GetCosmicGameUserCosmicTokenSummary400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				invalidUserAddressProblem(instance),
			),
		}, nil
	}

	userAid, err := s.userActivity.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return GetCosmicGameUserCosmicTokenSummary200JSONResponse{
			CosmicGameUserCosmicTokenSummaryJSONResponse: CosmicGameUserCosmicTokenSummaryJSONResponse(
				zeroUserCosmicTokenSummary(address),
			),
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for token summary", err, "address", address)
		return userCosmicTokenSummaryInternal(instance), nil
	}

	record, err := s.userActivity.UserCosmicTokenSummaryV2(ctx, userAid)
	if err != nil {
		s.logInternal(ctx, "get user cosmic token summary", err, "address", address)
		return userCosmicTokenSummaryInternal(instance), nil
	}
	summary, err := mapUserCosmicTokenSummary(address, record)
	if err != nil {
		s.logInternal(ctx, "map user cosmic token summary", err, "address", address)
		return userCosmicTokenSummaryInternal(instance), nil
	}
	return GetCosmicGameUserCosmicTokenSummary200JSONResponse{
		CosmicGameUserCosmicTokenSummaryJSONResponse: CosmicGameUserCosmicTokenSummaryJSONResponse(summary),
	}, nil
}

// GetCosmicGameUserPendingWinnings implements
// GET /api/v2/cosmicgame/users/{address}/pending-winnings.
func (s *Server) GetCosmicGameUserPendingWinnings(
	ctx context.Context,
	request GetCosmicGameUserPendingWinningsRequestObject,
) (GetCosmicGameUserPendingWinningsResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/pending-winnings", request.Address)
	address, _, valid := userAddressInput(request.Address)
	if !valid {
		return GetCosmicGameUserPendingWinnings400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				invalidUserAddressProblem(instance),
			),
		}, nil
	}

	userAid, err := s.userActivity.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return GetCosmicGameUserPendingWinnings200JSONResponse{
			CosmicGameUserPendingWinningsJSONResponse: CosmicGameUserPendingWinningsJSONResponse(
				zeroUserPendingWinnings(address),
			),
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for pending winnings", err, "address", address)
		return userPendingWinningsInternal(instance), nil
	}

	record, err := s.userActivity.UserPendingWinnings(ctx, userAid)
	if err != nil {
		s.logInternal(ctx, "get user pending winnings", err, "address", address)
		return userPendingWinningsInternal(instance), nil
	}
	winnings, err := mapUserPendingWinnings(address, record)
	if err != nil {
		s.logInternal(ctx, "map user pending winnings", err, "address", address)
		return userPendingWinningsInternal(instance), nil
	}
	return GetCosmicGameUserPendingWinnings200JSONResponse{
		CosmicGameUserPendingWinningsJSONResponse: CosmicGameUserPendingWinningsJSONResponse(winnings),
	}, nil
}

func userOwnedTokensBadRequest(problem Problem) ListCosmicGameUserCosmicSignatureTokensResponseObject {
	return ListCosmicGameUserCosmicSignatureTokens400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func userOwnedTokensInternal(instance string) ListCosmicGameUserCosmicSignatureTokensResponseObject {
	return ListCosmicGameUserCosmicSignatureTokens500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(instance),
		),
	}
}

func userCosmicTokenSummaryInternal(instance string) GetCosmicGameUserCosmicTokenSummaryResponseObject {
	return GetCosmicGameUserCosmicTokenSummary500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(instance),
		),
	}
}

func userPendingWinningsInternal(instance string) GetCosmicGameUserPendingWinningsResponseObject {
	return GetCosmicGameUserPendingWinnings500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(instance),
		),
	}
}
