package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const (
	globalCstActionsInstance      = "/api/v2/cosmicgame/staking/cst/actions"
	globalCstStakedInstance       = "/api/v2/cosmicgame/staking/cst/staked-tokens"
	globalStakingDepositsInstance = "/api/v2/cosmicgame/staking/cst/deposits"
	globalRwalkActionsInstance    = "/api/v2/cosmicgame/staking/random-walk/actions"
	globalRwalkStakedInstance     = "/api/v2/cosmicgame/staking/random-walk/staked-tokens"
	globalStakerRaffleInstance    = "/api/v2/cosmicgame/staking/raffle-nft-wins"
)

type globalStakingActionFetch func(
	context.Context,
	*cgstore.GlobalStakingActionPageCursor,
	int,
) ([]cgstore.GlobalStakingActionRecord, bool, error)

func listGlobalStakingActions[T any](
	ctx context.Context,
	s *Server,
	instance string,
	cursor *Cursor,
	requestedLimit *Limit,
	resource globalStakingEventResource,
	fetch globalStakingActionFetch,
	mapRecord func(cgstore.GlobalStakingActionRecord) (T, error),
) ([]T, PageMeta, *Problem) {
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
	var after *cgstore.GlobalStakingActionPageCursor
	if cursor != nil {
		decoded, err := decodeGlobalStakingEventCursor(*cursor, resource)
		if err != nil {
			problem := invalidGlobalStakingCursorProblem(instance)
			return nil, PageMeta{}, &problem
		}
		after = &cgstore.GlobalStakingActionPageCursor{EventLogID: decoded.EventLogID}
	}
	records, hasMore, err := fetch(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list global staking actions", err, "resource", resource)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate global staking action cardinality", err, "resource", resource)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}
	data := make([]T, 0, len(records))
	previousEventLogID := int64(0)
	hasPrevious := false
	if after != nil {
		previousEventLogID = after.EventLogID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if record.Tx.EvtLogId < 1 ||
			(hasPrevious && record.Tx.EvtLogId >= previousEventLogID) {
			s.logInternal(ctx, "validate global staking action page",
				errors.New("repository returned an unordered staking action"),
				"resource", resource,
				"event_log_id", record.Tx.EvtLogId)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		mapped, err := mapRecord(record)
		if err != nil {
			s.logInternal(ctx, "map global staking action", err,
				"resource", resource,
				"event_log_id", record.Tx.EvtLogId)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		data = append(data, mapped)
		previousEventLogID = record.Tx.EvtLogId
		hasPrevious = true
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list global staking actions",
				errors.New("repository reported another page without a cursor row"),
				"resource", resource)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		next, err := encodeGlobalStakingEventCursor(globalStakingEventCursor{
			Version:    globalStakingEventCursorVersion,
			Resource:   resource,
			EventLogID: previousEventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode global staking action cursor", err, "resource", resource)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		meta.NextCursor = &next
	}
	return data, meta, nil
}

// ListCosmicGameGlobalCstStakingActions implements
// GET /api/v2/cosmicgame/staking/cst/actions.
func (s *Server) ListCosmicGameGlobalCstStakingActions(
	ctx context.Context,
	request ListCosmicGameGlobalCstStakingActionsRequestObject,
) (ListCosmicGameGlobalCstStakingActionsResponseObject, error) {
	data, meta, problem := listGlobalStakingActions(
		ctx,
		s,
		globalCstActionsInstance,
		request.Params.Cursor,
		request.Params.Limit,
		globalStakingEventCstActions,
		s.globalStaking.GlobalCstStakingActionsPage,
		mapGlobalCstStakingAction,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameGlobalCstStakingActions400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameGlobalCstStakingActions500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameGlobalCstStakingActions200JSONResponse{
		CosmicGameGlobalCstStakingActionPageJSONResponse: CosmicGameGlobalCstStakingActionPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameGlobalRandomWalkStakingActions implements
// GET /api/v2/cosmicgame/staking/random-walk/actions.
func (s *Server) ListCosmicGameGlobalRandomWalkStakingActions(
	ctx context.Context,
	request ListCosmicGameGlobalRandomWalkStakingActionsRequestObject,
) (ListCosmicGameGlobalRandomWalkStakingActionsResponseObject, error) {
	data, meta, problem := listGlobalStakingActions(
		ctx,
		s,
		globalRwalkActionsInstance,
		request.Params.Cursor,
		request.Params.Limit,
		globalStakingEventRwalkActions,
		s.globalStaking.GlobalRwalkStakingActionsPage,
		mapGlobalRandomWalkStakingAction,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameGlobalRandomWalkStakingActions400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameGlobalRandomWalkStakingActions500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameGlobalRandomWalkStakingActions200JSONResponse{
		CosmicGameGlobalRandomWalkStakingActionPageJSONResponse: CosmicGameGlobalRandomWalkStakingActionPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetCosmicGameGlobalCstStakingAction implements
// GET /api/v2/cosmicgame/staking/cst/actions/{actionId}.
func (s *Server) GetCosmicGameGlobalCstStakingAction(
	ctx context.Context,
	request GetCosmicGameGlobalCstStakingActionRequestObject,
) (GetCosmicGameGlobalCstStakingActionResponseObject, error) {
	instance := fmt.Sprintf("%s/%d", globalCstActionsInstance, request.ActionId)
	if request.ActionId < 0 {
		return GetCosmicGameGlobalCstStakingAction400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				negativeStakingActionIDProblem(instance),
			),
		}, nil
	}
	record, err := s.globalStaking.StakeActionCstInfo(ctx, request.ActionId)
	if errors.Is(err, store.ErrNotFound) {
		return GetCosmicGameGlobalCstStakingAction404ApplicationProblemPlusJSONResponse{
			NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(
				stakingActionNotFoundProblem(instance),
			),
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "get global CST staking action", err, "action_id", request.ActionId)
		return GetCosmicGameGlobalCstStakingAction500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}, nil
	}
	action, err := mapGlobalCstStakingActionDetail(record)
	if err != nil {
		s.logInternal(ctx, "map global CST staking action", err, "action_id", request.ActionId)
		return GetCosmicGameGlobalCstStakingAction500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}, nil
	}
	return GetCosmicGameGlobalCstStakingAction200JSONResponse{
		CosmicGameGlobalCstStakingActionJSONResponse: CosmicGameGlobalCstStakingActionJSONResponse(action),
	}, nil
}

// GetCosmicGameGlobalRandomWalkStakingAction implements
// GET /api/v2/cosmicgame/staking/random-walk/actions/{actionId}.
func (s *Server) GetCosmicGameGlobalRandomWalkStakingAction(
	ctx context.Context,
	request GetCosmicGameGlobalRandomWalkStakingActionRequestObject,
) (GetCosmicGameGlobalRandomWalkStakingActionResponseObject, error) {
	instance := fmt.Sprintf("%s/%d", globalRwalkActionsInstance, request.ActionId)
	if request.ActionId < 0 {
		return GetCosmicGameGlobalRandomWalkStakingAction400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(
				negativeStakingActionIDProblem(instance),
			),
		}, nil
	}
	record, err := s.globalStaking.StakeActionRwalkInfo(ctx, request.ActionId)
	if errors.Is(err, store.ErrNotFound) {
		return GetCosmicGameGlobalRandomWalkStakingAction404ApplicationProblemPlusJSONResponse{
			NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(
				stakingActionNotFoundProblem(instance),
			),
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "get global RandomWalk staking action", err, "action_id", request.ActionId)
		return GetCosmicGameGlobalRandomWalkStakingAction500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}, nil
	}
	action, err := mapGlobalRandomWalkStakingActionDetail(record)
	if err != nil {
		s.logInternal(ctx, "map global RandomWalk staking action", err, "action_id", request.ActionId)
		return GetCosmicGameGlobalRandomWalkStakingAction500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}, nil
	}
	return GetCosmicGameGlobalRandomWalkStakingAction200JSONResponse{
		CosmicGameGlobalRandomWalkStakingActionJSONResponse: CosmicGameGlobalRandomWalkStakingActionJSONResponse(action),
	}, nil
}

type globalStakedTokenFetch[T any] func(
	context.Context,
	*cgstore.GlobalStakedTokenPageCursor,
	int,
) ([]T, bool, error)

func listGlobalStakedTokens[StoreRecord, APIRecord any](
	ctx context.Context,
	s *Server,
	instance string,
	cursor *Cursor,
	requestedLimit *Limit,
	resource globalStakedTokenResource,
	fetch globalStakedTokenFetch[StoreRecord],
	identity func(StoreRecord) int64,
	mapRecord func(StoreRecord) (APIRecord, error),
) ([]APIRecord, PageMeta, *Problem) {
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
	var after *cgstore.GlobalStakedTokenPageCursor
	if cursor != nil {
		decoded, err := decodeGlobalStakedTokenCursor(*cursor, resource)
		if err != nil {
			problem := invalidGlobalStakingCursorProblem(instance)
			return nil, PageMeta{}, &problem
		}
		after = &cgstore.GlobalStakedTokenPageCursor{TokenID: decoded.TokenID}
	}
	records, hasMore, err := fetch(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list globally staked tokens", err, "resource", resource)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate globally staked token cardinality", err, "resource", resource)
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
		tokenID := identity(records[i])
		if tokenID < 0 || (hasPrevious && tokenID <= previousTokenID) {
			s.logInternal(ctx, "validate globally staked token page",
				errors.New("repository returned an unordered staked token"),
				"resource", resource,
				"token_id", tokenID)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		mapped, err := mapRecord(records[i])
		if err != nil {
			s.logInternal(ctx, "map globally staked token", err,
				"resource", resource,
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
			s.logInternal(ctx, "list globally staked tokens",
				errors.New("repository reported another page without a cursor row"),
				"resource", resource)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		next, err := encodeGlobalStakedTokenCursor(globalStakedTokenCursor{
			Version:  globalStakedTokenCursorVersion,
			Resource: resource,
			TokenID:  previousTokenID,
		})
		if err != nil {
			s.logInternal(ctx, "encode globally staked token cursor", err, "resource", resource)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		meta.NextCursor = &next
	}
	return data, meta, nil
}

// ListCosmicGameGlobalStakedCstTokens implements
// GET /api/v2/cosmicgame/staking/cst/staked-tokens.
func (s *Server) ListCosmicGameGlobalStakedCstTokens(
	ctx context.Context,
	request ListCosmicGameGlobalStakedCstTokensRequestObject,
) (ListCosmicGameGlobalStakedCstTokensResponseObject, error) {
	data, meta, problem := listGlobalStakedTokens(
		ctx,
		s,
		globalCstStakedInstance,
		request.Params.Cursor,
		request.Params.Limit,
		globalStakedTokenCst,
		s.globalStaking.GlobalStakedCstTokensPage,
		func(record cgstore.GlobalStakedCstTokenRecord) int64 { return record.TokenID },
		mapGlobalStakedCstToken,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameGlobalStakedCstTokens400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameGlobalStakedCstTokens500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameGlobalStakedCstTokens200JSONResponse{
		CosmicGameGlobalStakedCstTokenPageJSONResponse: CosmicGameGlobalStakedCstTokenPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameGlobalStakedRandomWalkTokens implements
// GET /api/v2/cosmicgame/staking/random-walk/staked-tokens.
func (s *Server) ListCosmicGameGlobalStakedRandomWalkTokens(
	ctx context.Context,
	request ListCosmicGameGlobalStakedRandomWalkTokensRequestObject,
) (ListCosmicGameGlobalStakedRandomWalkTokensResponseObject, error) {
	data, meta, problem := listGlobalStakedTokens(
		ctx,
		s,
		globalRwalkStakedInstance,
		request.Params.Cursor,
		request.Params.Limit,
		globalStakedTokenRwalk,
		s.globalStaking.GlobalStakedRwalkTokensPage,
		func(record cgstore.GlobalStakedRwalkTokenRecord) int64 { return record.TokenID },
		mapGlobalStakedRandomWalkToken,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameGlobalStakedRandomWalkTokens400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameGlobalStakedRandomWalkTokens500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameGlobalStakedRandomWalkTokens200JSONResponse{
		CosmicGameGlobalStakedRandomWalkTokenPageJSONResponse: CosmicGameGlobalStakedRandomWalkTokenPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameGlobalStakingDeposits implements
// GET /api/v2/cosmicgame/staking/cst/deposits.
func (s *Server) ListCosmicGameGlobalStakingDeposits(
	ctx context.Context,
	request ListCosmicGameGlobalStakingDepositsRequestObject,
) (ListCosmicGameGlobalStakingDepositsResponseObject, error) {
	badRequest := func(problem Problem) ListCosmicGameGlobalStakingDepositsResponseObject {
		return ListCosmicGameGlobalStakingDeposits400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListCosmicGameGlobalStakingDepositsResponseObject {
		return ListCosmicGameGlobalStakingDeposits500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(globalStakingDepositsInstance),
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
			globalStakingDepositsInstance,
		)), nil
	}
	var after *cgstore.GlobalStakingDepositPageCursor
	if request.Params.Cursor != nil {
		decoded, err := decodeGlobalStakingDepositCursor(*request.Params.Cursor)
		if err != nil {
			return badRequest(invalidGlobalStakingCursorProblem(globalStakingDepositsInstance)), nil
		}
		after = &cgstore.GlobalStakingDepositPageCursor{DepositID: decoded.DepositID}
	}
	records, hasMore, err := s.globalStaking.GlobalStakingDepositsPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list global staking deposits", err)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate global staking deposit cardinality", err)
		return internal(), nil
	}
	data := make([]GlobalStakingDeposit, 0, len(records))
	previousDepositID := int64(0)
	hasPrevious := false
	if after != nil {
		previousDepositID = after.DepositID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if record.DepositID < 0 || (hasPrevious && record.DepositID >= previousDepositID) {
			s.logInternal(ctx, "validate global staking deposit page",
				errors.New("repository returned an unordered staking deposit"),
				"deposit_id", record.DepositID)
			return internal(), nil
		}
		deposit, err := mapGlobalStakingDeposit(record)
		if err != nil {
			s.logInternal(ctx, "map global staking deposit", err, "deposit_id", record.DepositID)
			return internal(), nil
		}
		data = append(data, deposit)
		previousDepositID = record.DepositID
		hasPrevious = true
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list global staking deposits",
				errors.New("repository reported another page without a cursor row"))
			return internal(), nil
		}
		next, err := encodeGlobalStakingDepositCursor(globalStakingDepositCursor{
			Version:   globalStakingDepositCursorVersion,
			DepositID: previousDepositID,
		})
		if err != nil {
			s.logInternal(ctx, "encode global staking deposit cursor", err)
			return internal(), nil
		}
		meta.NextCursor = &next
	}
	return ListCosmicGameGlobalStakingDeposits200JSONResponse{
		CosmicGameGlobalStakingDepositPageJSONResponse: CosmicGameGlobalStakingDepositPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameRoundStakingRewards implements
// GET /api/v2/cosmicgame/rounds/{round}/staking-rewards.
func (s *Server) ListCosmicGameRoundStakingRewards(
	ctx context.Context,
	request ListCosmicGameRoundStakingRewardsRequestObject,
) (ListCosmicGameRoundStakingRewardsResponseObject, error) {
	roundNum := request.Round
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d/staking-rewards", roundNum)
	badRequest := func(problem Problem) ListCosmicGameRoundStakingRewardsResponseObject {
		return ListCosmicGameRoundStakingRewards400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListCosmicGameRoundStakingRewardsResponseObject {
		return ListCosmicGameRoundStakingRewards500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(instance),
			),
		}
	}
	if roundNum < 0 {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Round must be a non-negative integer.",
			instance,
		)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}
	var after *cgstore.RoundStakingRewardPageCursor
	if request.Params.Cursor != nil {
		decoded, err := decodeRoundStakingRewardCursor(*request.Params.Cursor, roundNum)
		if err != nil {
			return badRequest(invalidGlobalStakingCursorProblem(instance)), nil
		}
		after = &cgstore.RoundStakingRewardPageCursor{
			DepositID: decoded.DepositID,
			StakerAid: decoded.StakerAid,
		}
	}
	exists, err := s.globalStaking.CompletedRoundExists(ctx, roundNum)
	if err != nil {
		s.logInternal(ctx, "check round for staking rewards", err, "round", roundNum)
		return internal(), nil
	}
	if !exists {
		return ListCosmicGameRoundStakingRewards404ApplicationProblemPlusJSONResponse{
			NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(
				roundNotFoundProblem(instance),
			),
		}, nil
	}
	records, hasMore, err := s.globalStaking.RoundStakingRewardsPage(ctx, roundNum, after, limit)
	if err != nil {
		s.logInternal(ctx, "list round staking rewards", err, "round", roundNum)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate round staking reward cardinality", err, "round", roundNum)
		return internal(), nil
	}
	data := make([]RoundStakingReward, 0, len(records))
	previousDepositID := int64(0)
	previousStakerAid := int64(0)
	hasPrevious := false
	if after != nil {
		previousDepositID = after.DepositID
		previousStakerAid = after.StakerAid
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		ordered := !hasPrevious ||
			record.DepositID < previousDepositID ||
			(record.DepositID == previousDepositID && record.StakerAid > previousStakerAid)
		if record.RoundNum != roundNum || record.DepositID < 0 || record.StakerAid < 1 || !ordered {
			s.logInternal(ctx, "validate round staking reward page",
				errors.New("repository returned an out-of-scope or unordered staking reward"),
				"round", roundNum,
				"deposit_id", record.DepositID,
				"staker_aid", record.StakerAid)
			return internal(), nil
		}
		reward, err := mapRoundStakingReward(record)
		if err != nil {
			s.logInternal(ctx, "map round staking reward", err,
				"round", roundNum,
				"deposit_id", record.DepositID,
				"staker_aid", record.StakerAid)
			return internal(), nil
		}
		data = append(data, reward)
		previousDepositID = record.DepositID
		previousStakerAid = record.StakerAid
		hasPrevious = true
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list round staking rewards",
				errors.New("repository reported another page without a cursor row"),
				"round", roundNum)
			return internal(), nil
		}
		next, err := encodeRoundStakingRewardCursor(roundStakingRewardCursor{
			Version:   roundStakingRewardCursorVersion,
			Round:     roundNum,
			DepositID: previousDepositID,
			StakerAid: previousStakerAid,
		})
		if err != nil {
			s.logInternal(ctx, "encode round staking reward cursor", err, "round", roundNum)
			return internal(), nil
		}
		meta.NextCursor = &next
	}
	return ListCosmicGameRoundStakingRewards200JSONResponse{
		CosmicGameRoundStakingRewardPageJSONResponse: CosmicGameRoundStakingRewardPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameGlobalStakerRaffleNftWins implements
// GET /api/v2/cosmicgame/staking/raffle-nft-wins.
func (s *Server) ListCosmicGameGlobalStakerRaffleNftWins(
	ctx context.Context,
	request ListCosmicGameGlobalStakerRaffleNftWinsRequestObject,
) (ListCosmicGameGlobalStakerRaffleNftWinsResponseObject, error) {
	badRequest := func(problem Problem) ListCosmicGameGlobalStakerRaffleNftWinsResponseObject {
		return ListCosmicGameGlobalStakerRaffleNftWins400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
		}
	}
	internal := func() ListCosmicGameGlobalStakerRaffleNftWinsResponseObject {
		return ListCosmicGameGlobalStakerRaffleNftWins500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
				internalProblem(globalStakerRaffleInstance),
			),
		}
	}
	var isRwalk bool
	var resource globalStakingEventResource
	switch request.Params.Pool {
	case StakerRafflePoolCst:
		resource = globalStakingEventCstRaffle
	case StakerRafflePoolRandomWalk:
		isRwalk = true
		resource = globalStakingEventRwalkRaffle
	default:
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"The pool must be cst or randomWalk.",
			globalStakerRaffleInstance,
		)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return badRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			globalStakerRaffleInstance,
		)), nil
	}
	var after *cgstore.GlobalStakerRafflePageCursor
	if request.Params.Cursor != nil {
		decoded, err := decodeGlobalStakingEventCursor(*request.Params.Cursor, resource)
		if err != nil {
			return badRequest(invalidGlobalStakingCursorProblem(globalStakerRaffleInstance)), nil
		}
		after = &cgstore.GlobalStakerRafflePageCursor{EventLogID: decoded.EventLogID}
	}
	records, hasMore, err := s.globalStaking.GlobalStakerRaffleNftWinsPage(ctx, isRwalk, after, limit)
	if err != nil {
		s.logInternal(ctx, "list global staker raffle NFT wins", err, "pool", request.Params.Pool)
		return internal(), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate global staker raffle cardinality", err, "pool", request.Params.Pool)
		return internal(), nil
	}
	data := make([]RoundRaffleNftWinner, 0, len(records))
	previousEventLogID := int64(0)
	hasPrevious := false
	if after != nil {
		previousEventLogID = after.EventLogID
		hasPrevious = true
	}
	for i := range records {
		record := records[i]
		if !record.IsStaker || record.IsRWalk != isRwalk || record.Tx.EvtLogId < 1 ||
			(hasPrevious && record.Tx.EvtLogId >= previousEventLogID) {
			s.logInternal(ctx, "validate global staker raffle page",
				errors.New("repository returned an out-of-scope or unordered raffle win"),
				"pool", request.Params.Pool,
				"event_log_id", record.Tx.EvtLogId)
			return internal(), nil
		}
		win, err := mapRoundRaffleNftWinner(record, true)
		if err != nil {
			s.logInternal(ctx, "map global staker raffle NFT win", err,
				"pool", request.Params.Pool,
				"event_log_id", record.Tx.EvtLogId)
			return internal(), nil
		}
		data = append(data, win)
		previousEventLogID = record.Tx.EvtLogId
		hasPrevious = true
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || !hasPrevious {
			s.logInternal(ctx, "list global staker raffle NFT wins",
				errors.New("repository reported another page without a cursor row"),
				"pool", request.Params.Pool)
			return internal(), nil
		}
		next, err := encodeGlobalStakingEventCursor(globalStakingEventCursor{
			Version:    globalStakingEventCursorVersion,
			Resource:   resource,
			EventLogID: previousEventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode global staker raffle cursor", err, "pool", request.Params.Pool)
			return internal(), nil
		}
		meta.NextCursor = &next
	}
	return ListCosmicGameGlobalStakerRaffleNftWins200JSONResponse{
		CosmicGameGlobalStakerRaffleNftWinPageJSONResponse: CosmicGameGlobalStakerRaffleNftWinPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func invalidGlobalStakingCursorProblem(instance string) Problem {
	return newProblem(
		http.StatusBadRequest,
		"invalid-cursor",
		"Invalid cursor",
		"The cursor is malformed, unsupported, or belongs to another staking resource.",
		instance,
	)
}

func negativeStakingActionIDProblem(instance string) Problem {
	return newProblem(
		http.StatusBadRequest,
		"invalid-parameter",
		"Invalid parameter",
		"Action ID must be a non-negative integer.",
		instance,
	)
}

func stakingActionNotFoundProblem(instance string) Problem {
	return newProblem(
		http.StatusNotFound,
		"staking-action-not-found",
		"Staking action not found",
		"No staking action exists with that ID.",
		instance,
	)
}
