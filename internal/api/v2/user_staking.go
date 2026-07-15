package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ListCosmicGameUserCstStakingActions implements
// GET /api/v2/cosmicgame/users/{address}/staking/cst/actions.
func (s *Server) ListCosmicGameUserCstStakingActions(
	ctx context.Context,
	request ListCosmicGameUserCstStakingActionsRequestObject,
) (ListCosmicGameUserCstStakingActionsResponseObject, error) {
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"staking/cst/actions",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceCstStakingActions,
		s.userStaking.UserAddressID,
		s.userStaking.UserCstStakingActionsPage,
		func(record cgstore.UserStakingActionRecord, userAid int64, _ string) (bool, int64) {
			return record.StakerAid == userAid, record.Tx.EvtLogId
		},
		mapUserCstStakingAction,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserCstStakingActions400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserCstStakingActions500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserCstStakingActions200JSONResponse{
		CosmicGameUserCstStakingActionPageJSONResponse: CosmicGameUserCstStakingActionPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserRandomWalkStakingActions implements
// GET /api/v2/cosmicgame/users/{address}/staking/random-walk/actions.
func (s *Server) ListCosmicGameUserRandomWalkStakingActions(
	ctx context.Context,
	request ListCosmicGameUserRandomWalkStakingActionsRequestObject,
) (ListCosmicGameUserRandomWalkStakingActionsResponseObject, error) {
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"staking/random-walk/actions",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceRwStakingActions,
		s.userStaking.UserAddressID,
		s.userStaking.UserRwalkStakingActionsPage,
		func(record cgstore.UserStakingActionRecord, userAid int64, _ string) (bool, int64) {
			return record.StakerAid == userAid, record.Tx.EvtLogId
		},
		mapUserRandomWalkStakingAction,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserRandomWalkStakingActions400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserRandomWalkStakingActions500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserRandomWalkStakingActions200JSONResponse{
		CosmicGameUserRandomWalkStakingActionPageJSONResponse: CosmicGameUserRandomWalkStakingActionPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

type userStakingTokenPageFetch[T any] func(
	context.Context,
	int64,
	*cgstore.UserStakingTokenPageCursor,
	int,
) ([]T, bool, error)

// listUserStakingTokenPage centralizes the ascending token-keyed staking
// flow shared by the two staked-token collections and the per-token reward
// totals: address and limit validation, wallet-and-resource scoped cursor
// decoding, the unindexed-wallet empty page, repository
// cardinality/scope/order enforcement and continuation-cursor encoding.
func listUserStakingTokenPage[StoreRecord, APIRecord any](
	ctx context.Context,
	s *Server,
	rawAddress string,
	pathSegment string,
	cursor *Cursor,
	requestedLimit *Limit,
	resource userStakingTokenResource,
	fetch userStakingTokenPageFetch[StoreRecord],
	identity func(record StoreRecord, userAid int64) (inScope bool, tokenID int64),
	mapRecord func(StoreRecord) (APIRecord, error),
) ([]APIRecord, PageMeta, *Problem) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/%s", rawAddress, pathSegment)
	address, scope, valid := userAddressInput(rawAddress)
	if !valid {
		problem := invalidUserAddressProblem(instance)
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

	var after *cgstore.UserStakingTokenPageCursor
	if cursor != nil {
		decoded, err := decodeUserStakingTokenCursor(*cursor, scope, resource)
		if err != nil {
			problem := newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another wallet or resource.",
				instance,
			)
			return nil, PageMeta{}, &problem
		}
		after = &cgstore.UserStakingTokenPageCursor{TokenID: decoded.TokenID}
	}

	userAid, err := s.userStaking.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return []APIRecord{}, PageMeta{Limit: limit}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for staking tokens", err,
			"resource", string(resource),
			"address", address)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}

	records, hasMore, err := fetch(ctx, userAid, after, limit)
	if err != nil {
		s.logInternal(ctx, "list user staking tokens", err,
			"resource", string(resource),
			"address", address)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate user staking token page cardinality", err,
			"resource", string(resource),
			"address", address)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}

	data := make([]APIRecord, 0, len(records))
	previousTokenID := int64(0)
	hasPrevious := false
	if after != nil {
		previousTokenID = after.TokenID
		hasPrevious = true
	}
	for i := range records {
		inScope, tokenID := identity(records[i], userAid)
		if !inScope || tokenID < 0 ||
			(hasPrevious && tokenID <= previousTokenID) {
			s.logInternal(ctx, "validate user staking token page",
				errors.New("repository returned an out-of-scope or unordered staking token row"),
				"resource", string(resource),
				"address", address,
				"token_id", tokenID)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		mapped, err := mapRecord(records[i])
		if err != nil {
			s.logInternal(ctx, "map user staking token row", err,
				"resource", string(resource),
				"address", address,
				"token_id", tokenID)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		data = append(data, mapped)
		previousTokenID = tokenID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list user staking tokens",
				errors.New("repository reported another page without a cursor row"),
				"resource", string(resource),
				"address", address)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		next, err := encodeUserStakingTokenCursor(userStakingTokenCursor{
			Version:  userStakingTokenCursorVersion,
			Address:  scope,
			Resource: resource,
			TokenID:  previousTokenID,
		})
		if err != nil {
			s.logInternal(ctx, "encode user staking token cursor", err,
				"resource", string(resource),
				"address", address)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		meta.NextCursor = &next
	}
	return data, meta, nil
}

// ListCosmicGameUserStakedCstTokens implements
// GET /api/v2/cosmicgame/users/{address}/staking/cst/staked-tokens.
func (s *Server) ListCosmicGameUserStakedCstTokens(
	ctx context.Context,
	request ListCosmicGameUserStakedCstTokensRequestObject,
) (ListCosmicGameUserStakedCstTokensResponseObject, error) {
	data, meta, problem := listUserStakingTokenPage(
		ctx,
		s,
		request.Address,
		"staking/cst/staked-tokens",
		request.Params.Cursor,
		request.Params.Limit,
		userStakingTokenResourceCstStakedTokens,
		s.userStaking.UserStakedCstTokensPage,
		func(record cgstore.UserStakedCstTokenRecord, userAid int64) (bool, int64) {
			return record.StakerAid == userAid, record.TokenID
		},
		mapUserStakedCstToken,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserStakedCstTokens400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserStakedCstTokens500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserStakedCstTokens200JSONResponse{
		CosmicGameUserStakedCstTokenPageJSONResponse: CosmicGameUserStakedCstTokenPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserStakedRandomWalkTokens implements
// GET /api/v2/cosmicgame/users/{address}/staking/random-walk/staked-tokens.
func (s *Server) ListCosmicGameUserStakedRandomWalkTokens(
	ctx context.Context,
	request ListCosmicGameUserStakedRandomWalkTokensRequestObject,
) (ListCosmicGameUserStakedRandomWalkTokensResponseObject, error) {
	data, meta, problem := listUserStakingTokenPage(
		ctx,
		s,
		request.Address,
		"staking/random-walk/staked-tokens",
		request.Params.Cursor,
		request.Params.Limit,
		userStakingTokenResourceRwStakedTokens,
		s.userStaking.UserStakedRwalkTokensPage,
		func(record cgstore.UserStakedRwalkTokenRecord, userAid int64) (bool, int64) {
			return record.StakerAid == userAid, record.TokenID
		},
		mapUserStakedRandomWalkToken,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserStakedRandomWalkTokens400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserStakedRandomWalkTokens500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserStakedRandomWalkTokens200JSONResponse{
		CosmicGameUserStakedRandomWalkTokenPageJSONResponse: CosmicGameUserStakedRandomWalkTokenPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserStakingTokenRewards implements
// GET /api/v2/cosmicgame/users/{address}/staking/cst/token-rewards.
func (s *Server) ListCosmicGameUserStakingTokenRewards(
	ctx context.Context,
	request ListCosmicGameUserStakingTokenRewardsRequestObject,
) (ListCosmicGameUserStakingTokenRewardsResponseObject, error) {
	data, meta, problem := listUserStakingTokenPage(
		ctx,
		s,
		request.Address,
		"staking/cst/token-rewards",
		request.Params.Cursor,
		request.Params.Limit,
		userStakingTokenResourceCstTokenRewards,
		s.userStaking.UserStakingTokenRewardsPage,
		func(record cgstore.UserStakingTokenRewardRecord, _ int64) (bool, int64) {
			// The aggregate rows carry no staker column; the repository
			// filter is the scope and ordering stays enforced.
			return true, record.TokenID
		},
		mapUserStakingTokenReward,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserStakingTokenRewards400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserStakingTokenRewards500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserStakingTokenRewards200JSONResponse{
		CosmicGameUserStakingTokenRewardPageJSONResponse: CosmicGameUserStakingTokenRewardPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserStakingDeposits implements
// GET /api/v2/cosmicgame/users/{address}/staking/cst/deposits.
func (s *Server) ListCosmicGameUserStakingDeposits(
	ctx context.Context,
	request ListCosmicGameUserStakingDepositsRequestObject,
) (ListCosmicGameUserStakingDepositsResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/staking/cst/deposits", request.Address)
	address, scope, valid := userAddressInput(request.Address)
	if !valid {
		return userStakingDepositsBadRequest(invalidUserAddressProblem(instance)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return userStakingDepositsBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}
	var claimed *bool
	if request.Params.Claimed != nil {
		value := *request.Params.Claimed
		claimed = &value
	}

	var after *cgstore.UserStakingDepositPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeUserStakingDepositCursor(*request.Params.Cursor, scope)
		if err != nil {
			return userStakingDepositsBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another wallet.",
				instance,
			)), nil
		}
		after = &cgstore.UserStakingDepositPageCursor{DepositID: cursor.DepositID}
	}

	userAid, err := s.userStaking.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return ListCosmicGameUserStakingDeposits200JSONResponse{
			CosmicGameUserStakingDepositPageJSONResponse: CosmicGameUserStakingDepositPageJSONResponse{
				Data: []UserStakingDeposit{},
				Meta: PageMeta{Limit: limit},
			},
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for staking deposits", err, "address", address)
		return userStakingDepositsInternal(instance), nil
	}

	records, hasMore, err := s.userStaking.UserStakingDepositsPage(ctx, userAid, claimed, after, limit)
	if err != nil {
		s.logInternal(ctx, "list user staking deposits", err, "address", address)
		return userStakingDepositsInternal(instance), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate user staking deposit page cardinality", err, "address", address)
		return userStakingDepositsInternal(instance), nil
	}

	data := make([]UserStakingDeposit, 0, len(records))
	previousDepositID := int64(0)
	hasPrevious := false
	if after != nil {
		previousDepositID = after.DepositID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		fullyClaimed := record.PendingNftCount == 0
		inScope := record.StakerAid == userAid &&
			(claimed == nil || fullyClaimed == *claimed)
		if !inScope || record.DepositID < 0 ||
			(hasPrevious && record.DepositID >= previousDepositID) {
			s.logInternal(ctx, "validate user staking deposit page",
				errors.New("repository returned an out-of-scope or unordered staking deposit"),
				"address", address,
				"deposit_id", record.DepositID)
			return userStakingDepositsInternal(instance), nil
		}
		deposit, err := mapUserStakingDeposit(record)
		if err != nil {
			s.logInternal(ctx, "map user staking deposit", err,
				"address", address,
				"deposit_id", record.DepositID)
			return userStakingDepositsInternal(instance), nil
		}
		data = append(data, deposit)
		previousDepositID = record.DepositID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list user staking deposits",
				errors.New("repository reported another page without a cursor row"),
				"address", address)
			return userStakingDepositsInternal(instance), nil
		}
		next, err := encodeUserStakingDepositCursor(userStakingDepositCursor{
			Version:   userStakingDepositCursorVersion,
			Address:   scope,
			DepositID: previousDepositID,
		})
		if err != nil {
			s.logInternal(ctx, "encode user staking deposit cursor", err, "address", address)
			return userStakingDepositsInternal(instance), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameUserStakingDeposits200JSONResponse{
		CosmicGameUserStakingDepositPageJSONResponse: CosmicGameUserStakingDepositPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserStakingDepositRewards implements GET
// /api/v2/cosmicgame/users/{address}/staking/cst/deposits/{depositId}/rewards.
func (s *Server) ListCosmicGameUserStakingDepositRewards(
	ctx context.Context,
	request ListCosmicGameUserStakingDepositRewardsRequestObject,
) (ListCosmicGameUserStakingDepositRewardsResponseObject, error) {
	depositID := request.DepositId
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/staking/cst/deposits/%d/rewards",
		request.Address, depositID)
	address, scope, valid := userAddressInput(request.Address)
	if !valid {
		return userStakingDepositRewardsBadRequest(invalidUserAddressProblem(instance)), nil
	}
	if depositID < 0 {
		return userStakingDepositRewardsBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Deposit ID must be a non-negative integer.",
			instance,
		)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return userStakingDepositRewardsBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	var after *cgstore.UserStakingRewardPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeUserStakingDepositRewardCursor(*request.Params.Cursor, scope, depositID)
		if err != nil {
			return userStakingDepositRewardsBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another wallet or deposit.",
				instance,
			)), nil
		}
		after = &cgstore.UserStakingRewardPageCursor{ActionID: cursor.ActionID}
	}

	exists, err := s.userStaking.StakingDepositExists(ctx, depositID)
	if err != nil {
		s.logInternal(ctx, "check staking deposit existence", err,
			"address", address,
			"deposit_id", depositID)
		return userStakingDepositRewardsInternal(instance), nil
	}
	if !exists {
		return ListCosmicGameUserStakingDepositRewards404ApplicationProblemPlusJSONResponse{
			NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(
				stakingDepositNotFoundProblem(instance),
			),
		}, nil
	}

	userAid, err := s.userStaking.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return ListCosmicGameUserStakingDepositRewards200JSONResponse{
			CosmicGameUserStakingDepositRewardPageJSONResponse: CosmicGameUserStakingDepositRewardPageJSONResponse{
				Data: []UserStakingDepositReward{},
				Meta: PageMeta{Limit: limit},
			},
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for staking deposit rewards", err, "address", address)
		return userStakingDepositRewardsInternal(instance), nil
	}

	records, hasMore, err := s.userStaking.UserStakingDepositRewardsPage(ctx, userAid, depositID, after, limit)
	if err != nil {
		s.logInternal(ctx, "list user staking deposit rewards", err,
			"address", address,
			"deposit_id", depositID)
		return userStakingDepositRewardsInternal(instance), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate user staking deposit reward page cardinality", err,
			"address", address,
			"deposit_id", depositID)
		return userStakingDepositRewardsInternal(instance), nil
	}

	data := make([]UserStakingDepositReward, 0, len(records))
	previousActionID := int64(0)
	hasPrevious := false
	if after != nil {
		previousActionID = after.ActionID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if record.StakerAid != userAid || record.ActionID < 0 ||
			(hasPrevious && record.ActionID <= previousActionID) {
			s.logInternal(ctx, "validate user staking deposit reward page",
				errors.New("repository returned an out-of-scope or unordered staking reward"),
				"address", address,
				"deposit_id", depositID,
				"action_id", record.ActionID)
			return userStakingDepositRewardsInternal(instance), nil
		}
		reward, err := mapUserStakingDepositReward(record)
		if err != nil {
			s.logInternal(ctx, "map user staking deposit reward", err,
				"address", address,
				"deposit_id", depositID,
				"action_id", record.ActionID)
			return userStakingDepositRewardsInternal(instance), nil
		}
		data = append(data, reward)
		previousActionID = record.ActionID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list user staking deposit rewards",
				errors.New("repository reported another page without a cursor row"),
				"address", address,
				"deposit_id", depositID)
			return userStakingDepositRewardsInternal(instance), nil
		}
		next, err := encodeUserStakingDepositRewardCursor(userStakingDepositRewardCursor{
			Version:   userStakingDepositRewardCursorVersion,
			Address:   scope,
			DepositID: depositID,
			ActionID:  previousActionID,
		})
		if err != nil {
			s.logInternal(ctx, "encode user staking deposit reward cursor", err,
				"address", address,
				"deposit_id", depositID)
			return userStakingDepositRewardsInternal(instance), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameUserStakingDepositRewards200JSONResponse{
		CosmicGameUserStakingDepositRewardPageJSONResponse: CosmicGameUserStakingDepositRewardPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserStakingTokenRewardDeposits implements GET
// /api/v2/cosmicgame/users/{address}/staking/cst/token-rewards/{nftTokenId}/deposits.
func (s *Server) ListCosmicGameUserStakingTokenRewardDeposits(
	ctx context.Context,
	request ListCosmicGameUserStakingTokenRewardDepositsRequestObject,
) (ListCosmicGameUserStakingTokenRewardDepositsResponseObject, error) {
	tokenID := request.NftTokenId
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/staking/cst/token-rewards/%d/deposits",
		request.Address, tokenID)
	address, scope, valid := userAddressInput(request.Address)
	if !valid {
		return userStakingTokenDepositsBadRequest(invalidUserAddressProblem(instance)), nil
	}
	if tokenID < 0 {
		return userStakingTokenDepositsBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Token ID must be a non-negative integer.",
			instance,
		)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return userStakingTokenDepositsBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	var after *cgstore.UserStakingTokenDepositPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeUserStakingTokenDepositCursor(*request.Params.Cursor, scope, tokenID)
		if err != nil {
			return userStakingTokenDepositsBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another wallet or token.",
				instance,
			)), nil
		}
		after = &cgstore.UserStakingTokenDepositPageCursor{DepositID: cursor.DepositID}
	}

	exists, err := s.userStaking.CosmicSignatureTokenExists(ctx, tokenID)
	if err != nil {
		s.logInternal(ctx, "check cosmic signature token existence", err,
			"address", address,
			"token_id", tokenID)
		return userStakingTokenDepositsInternal(instance), nil
	}
	if !exists {
		return ListCosmicGameUserStakingTokenRewardDeposits404ApplicationProblemPlusJSONResponse{
			NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(
				nftTokenNotFoundProblem(instance),
			),
		}, nil
	}

	userAid, err := s.userStaking.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return ListCosmicGameUserStakingTokenRewardDeposits200JSONResponse{
			CosmicGameUserStakingTokenRewardDepositPageJSONResponse: CosmicGameUserStakingTokenRewardDepositPageJSONResponse{
				Data: []UserStakingTokenRewardDeposit{},
				Meta: PageMeta{Limit: limit},
			},
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for staking token deposits", err, "address", address)
		return userStakingTokenDepositsInternal(instance), nil
	}

	records, hasMore, err := s.userStaking.UserStakingTokenRewardDepositsPage(ctx, userAid, tokenID, after, limit)
	if err != nil {
		s.logInternal(ctx, "list user staking token reward deposits", err,
			"address", address,
			"token_id", tokenID)
		return userStakingTokenDepositsInternal(instance), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate user staking token deposit page cardinality", err,
			"address", address,
			"token_id", tokenID)
		return userStakingTokenDepositsInternal(instance), nil
	}

	data := make([]UserStakingTokenRewardDeposit, 0, len(records))
	previousDepositID := int64(0)
	hasPrevious := false
	if after != nil {
		previousDepositID = after.DepositID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if record.StakerAid != userAid || record.DepositID < 0 ||
			(hasPrevious && record.DepositID <= previousDepositID) {
			s.logInternal(ctx, "validate user staking token deposit page",
				errors.New("repository returned an out-of-scope or unordered staking token deposit"),
				"address", address,
				"token_id", tokenID,
				"deposit_id", record.DepositID)
			return userStakingTokenDepositsInternal(instance), nil
		}
		deposit, err := mapUserStakingTokenRewardDeposit(record)
		if err != nil {
			s.logInternal(ctx, "map user staking token reward deposit", err,
				"address", address,
				"token_id", tokenID,
				"deposit_id", record.DepositID)
			return userStakingTokenDepositsInternal(instance), nil
		}
		data = append(data, deposit)
		previousDepositID = record.DepositID
		hasPrevious = true
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list user staking token reward deposits",
				errors.New("repository reported another page without a cursor row"),
				"address", address,
				"token_id", tokenID)
			return userStakingTokenDepositsInternal(instance), nil
		}
		next, err := encodeUserStakingTokenDepositCursor(userStakingTokenDepositCursor{
			Version:   userStakingTokenDepositCursorVersion,
			Address:   scope,
			TokenID:   tokenID,
			DepositID: previousDepositID,
		})
		if err != nil {
			s.logInternal(ctx, "encode user staking token deposit cursor", err,
				"address", address,
				"token_id", tokenID)
			return userStakingTokenDepositsInternal(instance), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameUserStakingTokenRewardDeposits200JSONResponse{
		CosmicGameUserStakingTokenRewardDepositPageJSONResponse: CosmicGameUserStakingTokenRewardDepositPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func stakingDepositNotFoundProblem(instance string) Problem {
	return newProblem(
		http.StatusNotFound,
		"staking-deposit-not-found",
		"Staking deposit not found",
		"No staking deposit exists with that ID.",
		instance,
	)
}

func nftTokenNotFoundProblem(instance string) Problem {
	return newProblem(
		http.StatusNotFound,
		"nft-token-not-found",
		"Token not found",
		"No Cosmic Signature NFT exists with that token ID.",
		instance,
	)
}

func userStakingDepositsBadRequest(problem Problem) ListCosmicGameUserStakingDepositsResponseObject {
	return ListCosmicGameUserStakingDeposits400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func userStakingDepositsInternal(instance string) ListCosmicGameUserStakingDepositsResponseObject {
	return ListCosmicGameUserStakingDeposits500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(instance),
		),
	}
}

func userStakingDepositRewardsBadRequest(problem Problem) ListCosmicGameUserStakingDepositRewardsResponseObject {
	return ListCosmicGameUserStakingDepositRewards400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func userStakingDepositRewardsInternal(instance string) ListCosmicGameUserStakingDepositRewardsResponseObject {
	return ListCosmicGameUserStakingDepositRewards500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(instance),
		),
	}
}

func userStakingTokenDepositsBadRequest(problem Problem) ListCosmicGameUserStakingTokenRewardDepositsResponseObject {
	return ListCosmicGameUserStakingTokenRewardDeposits400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func userStakingTokenDepositsInternal(instance string) ListCosmicGameUserStakingTokenRewardDepositsResponseObject {
	return ListCosmicGameUserStakingTokenRewardDeposits500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(instance),
		),
	}
}
