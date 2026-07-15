package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ListCosmicGameUserPrizes implements
// GET /api/v2/cosmicgame/users/{address}/prizes.
func (s *Server) ListCosmicGameUserPrizes(
	ctx context.Context,
	request ListCosmicGameUserPrizesRequestObject,
) (ListCosmicGameUserPrizesResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/prizes", request.Address)
	address, scope, valid := userAddressInput(request.Address)
	if !valid {
		return userPrizesBadRequest(invalidUserAddressProblem(instance)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return userPrizesBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	var after *cgstore.UserPrizePageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeUserPrizeCursor(*request.Params.Cursor, scope)
		if err != nil {
			return userPrizesBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another wallet.",
				instance,
			)), nil
		}
		after = &cgstore.UserPrizePageCursor{
			Round:       cursor.Round,
			PrizeType:   cursor.PrizeType,
			WinnerIndex: cursor.WinnerIndex,
		}
	}

	userAid, err := s.userHistories.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return ListCosmicGameUserPrizes200JSONResponse{
			CosmicGameUserPrizePageJSONResponse: CosmicGameUserPrizePageJSONResponse{
				Data: []RoundPrize{},
				Meta: PageMeta{Limit: limit},
			},
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for prizes", err, "address", address)
		return userPrizesInternal(instance), nil
	}

	records, hasMore, err := s.userHistories.UserPrizesPage(ctx, userAid, after, limit)
	if err != nil {
		s.logInternal(ctx, "list user prizes", err, "address", address)
		return userPrizesInternal(instance), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate user prize page cardinality", err, "address", address)
		return userPrizesInternal(instance), nil
	}

	data := make([]RoundPrize, 0, len(records))
	previous := after
	for i := range records {
		current := cgstore.UserPrizePageCursor{
			Round:       records[i].RoundNum,
			PrizeType:   records[i].RecordType,
			WinnerIndex: records[i].WinnerIndex,
		}
		if records[i].WinnerAid != userAid ||
			!strings.EqualFold(records[i].WinnerAddr, address) ||
			(previous != nil && !userPrizeCursorFollows(current, *previous)) {
			s.logInternal(ctx, "validate user prize page",
				errors.New("repository returned an out-of-scope or unordered user prize"),
				"address", address,
				"round", current.Round,
				"prize_type", current.PrizeType,
				"winner_index", current.WinnerIndex)
			return userPrizesInternal(instance), nil
		}
		prize, err := mapRoundPrize(records[i])
		if err != nil {
			s.logInternal(ctx, "map user prize", err,
				"address", address,
				"round", current.Round,
				"prize_type", current.PrizeType,
				"winner_index", current.WinnerIndex)
			return userPrizesInternal(instance), nil
		}
		data = append(data, prize)
		previous = &current
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if previous == nil || len(records) == 0 {
			s.logInternal(ctx, "list user prizes",
				errors.New("repository reported another prize page without a cursor row"),
				"address", address)
			return userPrizesInternal(instance), nil
		}
		next, err := encodeUserPrizeCursor(userPrizeCursor{
			Version:     userPrizeCursorVersion,
			Address:     scope,
			Round:       previous.Round,
			PrizeType:   previous.PrizeType,
			WinnerIndex: previous.WinnerIndex,
		})
		if err != nil {
			s.logInternal(ctx, "encode user prize cursor", err, "address", address)
			return userPrizesInternal(instance), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameUserPrizes200JSONResponse{
		CosmicGameUserPrizePageJSONResponse: CosmicGameUserPrizePageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// userPrizeCursorFollows reports whether current comes strictly after
// previous in the (round DESC, ptype ASC, winner_index ASC) prize order.
func userPrizeCursorFollows(current, previous cgstore.UserPrizePageCursor) bool {
	if current.Round != previous.Round {
		return current.Round < previous.Round
	}
	return current.PrizeType > previous.PrizeType ||
		(current.PrizeType == previous.PrizeType && current.WinnerIndex > previous.WinnerIndex)
}

// ListCosmicGameUserRaffleEthDeposits implements
// GET /api/v2/cosmicgame/users/{address}/raffle-eth-deposits.
func (s *Server) ListCosmicGameUserRaffleEthDeposits(
	ctx context.Context,
	request ListCosmicGameUserRaffleEthDepositsRequestObject,
) (ListCosmicGameUserRaffleEthDepositsResponseObject, error) {
	var claimed *bool
	if request.Params.Claimed != nil {
		value := *request.Params.Claimed
		claimed = &value
	}
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"raffle-eth-deposits",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceRaffleEthDeposits,
		func(ctx context.Context, userAid int64, after *cgstore.UserEventPageCursor, limit int) ([]cgstore.UserRaffleEthDepositRecord, bool, error) {
			return s.userHistories.UserRaffleEthDepositsPage(ctx, userAid, claimed, after, limit)
		},
		func(record cgstore.UserRaffleEthDepositRecord, userAid int64, address string) (bool, int64) {
			inScope := record.WinnerAid == userAid &&
				strings.EqualFold(record.WinnerAddr, address) &&
				(claimed == nil || record.Claimed == *claimed)
			return inScope, record.Tx.EvtLogId
		},
		mapUserRaffleEthDeposit,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserRaffleEthDeposits400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserRaffleEthDeposits500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserRaffleEthDeposits200JSONResponse{
		CosmicGameUserRaffleEthDepositPageJSONResponse: CosmicGameUserRaffleEthDepositPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserRaffleNftWins implements
// GET /api/v2/cosmicgame/users/{address}/raffle-nft-wins.
func (s *Server) ListCosmicGameUserRaffleNftWins(
	ctx context.Context,
	request ListCosmicGameUserRaffleNftWinsRequestObject,
) (ListCosmicGameUserRaffleNftWinsResponseObject, error) {
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"raffle-nft-wins",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceRaffleNftWins,
		s.userHistories.UserRaffleNftWinsPage,
		func(record cgstore.UserRaffleNftWinRecord, userAid int64, address string) (bool, int64) {
			inScope := record.WinnerAid == userAid &&
				strings.EqualFold(record.WinnerAddr, address)
			return inScope, record.Tx.EvtLogId
		},
		mapUserRaffleNftWin,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserRaffleNftWins400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserRaffleNftWins500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserRaffleNftWins200JSONResponse{
		CosmicGameUserRaffleNftWinPageJSONResponse: CosmicGameUserRaffleNftWinPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserEthDonations implements
// GET /api/v2/cosmicgame/users/{address}/eth-donations.
func (s *Server) ListCosmicGameUserEthDonations(
	ctx context.Context,
	request ListCosmicGameUserEthDonationsRequestObject,
) (ListCosmicGameUserEthDonationsResponseObject, error) {
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"eth-donations",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceEthDonations,
		s.userHistories.EthDonationsByUserPage,
		func(record cgstore.RoundEthDonationRecord, _ int64, address string) (bool, int64) {
			return strings.EqualFold(record.DonorAddr, address), record.Tx.EvtLogId
		},
		mapRoundEthDonation,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserEthDonations400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserEthDonations500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserEthDonations200JSONResponse{
		CosmicGameUserEthDonationPageJSONResponse: CosmicGameUserEthDonationPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserErc20Donations implements
// GET /api/v2/cosmicgame/users/{address}/erc20-donations.
func (s *Server) ListCosmicGameUserErc20Donations(
	ctx context.Context,
	request ListCosmicGameUserErc20DonationsRequestObject,
) (ListCosmicGameUserErc20DonationsResponseObject, error) {
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"erc20-donations",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceErc20Donations,
		s.userHistories.ERC20DonationsByUserPage,
		func(record cgstore.RoundERC20DonationRecord, _ int64, address string) (bool, int64) {
			return strings.EqualFold(record.DonorAddr, address), record.Tx.EvtLogId
		},
		mapRoundERC20Donation,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserErc20Donations400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserErc20Donations500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserErc20Donations200JSONResponse{
		CosmicGameUserErc20DonationPageJSONResponse: CosmicGameUserErc20DonationPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserNftDonations implements
// GET /api/v2/cosmicgame/users/{address}/nft-donations.
func (s *Server) ListCosmicGameUserNftDonations(
	ctx context.Context,
	request ListCosmicGameUserNftDonationsRequestObject,
) (ListCosmicGameUserNftDonationsResponseObject, error) {
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"nft-donations",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceNftDonations,
		s.userHistories.NFTDonationsByUserPage,
		func(record cgstore.RoundNFTDonationRecord, _ int64, address string) (bool, int64) {
			return strings.EqualFold(record.DonorAddr, address), record.Tx.EvtLogId
		},
		mapRoundNFTDonation,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserNftDonations400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserNftDonations500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserNftDonations200JSONResponse{
		CosmicGameUserNftDonationPageJSONResponse: CosmicGameUserNftDonationPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserDonatedNfts implements
// GET /api/v2/cosmicgame/users/{address}/donated-nfts.
func (s *Server) ListCosmicGameUserDonatedNfts(
	ctx context.Context,
	request ListCosmicGameUserDonatedNftsRequestObject,
) (ListCosmicGameUserDonatedNftsResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/donated-nfts", request.Address)
	var claimed *bool
	if request.Params.Status != nil {
		status := *request.Params.Status
		if !status.Valid() {
			return ListCosmicGameUserDonatedNfts400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(newProblem(
					http.StatusBadRequest,
					"invalid-parameter",
					"Invalid parameter",
					"Status must be claimed or unclaimed.",
					instance,
				)),
			}, nil
		}
		value := status == Claimed
		claimed = &value
	}
	data, meta, problem := listUserEventPage(
		ctx,
		s,
		request.Address,
		"donated-nfts",
		request.Params.Cursor,
		request.Params.Limit,
		userEventResourceDonatedNfts,
		func(ctx context.Context, userAid int64, after *cgstore.UserEventPageCursor, limit int) ([]cgstore.UserDonatedNftRecord, bool, error) {
			return s.userHistories.UserDonatedNftsPage(ctx, userAid, claimed, after, limit)
		},
		func(record cgstore.UserDonatedNftRecord, userAid int64, _ string) (bool, int64) {
			claimedByUser := record.Claim != nil && record.Claim.ClaimerAid == userAid
			inScope := (record.RoundWinnerAid == userAid || claimedByUser) &&
				(claimed == nil || record.Claimed == *claimed)
			return inScope, record.Tx.EvtLogId
		},
		mapUserDonatedNft,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return ListCosmicGameUserDonatedNfts400ApplicationProblemPlusJSONResponse{
				BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
			}, nil
		}
		return ListCosmicGameUserDonatedNfts500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	return ListCosmicGameUserDonatedNfts200JSONResponse{
		CosmicGameUserDonatedNftPageJSONResponse: CosmicGameUserDonatedNftPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListCosmicGameUserDonatedErc20 implements
// GET /api/v2/cosmicgame/users/{address}/donated-erc20.
func (s *Server) ListCosmicGameUserDonatedErc20(
	ctx context.Context,
	request ListCosmicGameUserDonatedErc20RequestObject,
) (ListCosmicGameUserDonatedErc20ResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/users/%s/donated-erc20", request.Address)
	address, scope, valid := userAddressInput(request.Address)
	if !valid {
		return userDonatedErc20BadRequest(invalidUserAddressProblem(instance)), nil
	}
	limit, valid := resolvePageLimit(request.Params.Limit)
	if !valid {
		return userDonatedErc20BadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	var after *cgstore.UserDonatedErc20PageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeUserDonatedErc20Cursor(*request.Params.Cursor, scope)
		if err != nil {
			return userDonatedErc20BadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another wallet.",
				instance,
			)), nil
		}
		after = &cgstore.UserDonatedErc20PageCursor{
			Round:    cursor.Round,
			TokenAid: cursor.TokenID,
		}
	}

	userAid, err := s.userHistories.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return ListCosmicGameUserDonatedErc20200JSONResponse{
			CosmicGameUserDonatedErc20PageJSONResponse: CosmicGameUserDonatedErc20PageJSONResponse{
				Data: []UserDonatedErc20{},
				Meta: PageMeta{Limit: limit},
			},
		}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for donated erc20", err, "address", address)
		return userDonatedErc20Internal(instance), nil
	}

	records, hasMore, err := s.userHistories.UserDonatedErc20Page(ctx, userAid, after, limit)
	if err != nil {
		s.logInternal(ctx, "list user donated erc20", err, "address", address)
		return userDonatedErc20Internal(instance), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate user donated erc20 page cardinality", err, "address", address)
		return userDonatedErc20Internal(instance), nil
	}

	data := make([]UserDonatedErc20, 0, len(records))
	previous := after
	for i := range records {
		current := cgstore.UserDonatedErc20PageCursor{
			Round:    records[i].RoundNum,
			TokenAid: records[i].TokenAid,
		}
		if current.TokenAid < 1 ||
			(previous != nil && !userDonatedErc20CursorFollows(current, *previous)) {
			s.logInternal(ctx, "validate user donated erc20 page",
				errors.New("repository returned an unordered donated ERC-20 page"),
				"address", address,
				"round", current.Round)
			return userDonatedErc20Internal(instance), nil
		}
		summary, err := mapUserDonatedErc20(records[i])
		if err != nil {
			s.logInternal(ctx, "map user donated erc20", err,
				"address", address,
				"round", current.Round)
			return userDonatedErc20Internal(instance), nil
		}
		data = append(data, summary)
		previous = &current
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if previous == nil || len(records) == 0 {
			s.logInternal(ctx, "list user donated erc20",
				errors.New("repository reported another page without a cursor row"),
				"address", address)
			return userDonatedErc20Internal(instance), nil
		}
		next, err := encodeUserDonatedErc20Cursor(userDonatedErc20Cursor{
			Version: userDonatedErc20CursorVersion,
			Address: scope,
			Round:   previous.Round,
			TokenID: previous.TokenAid,
		})
		if err != nil {
			s.logInternal(ctx, "encode user donated erc20 cursor", err, "address", address)
			return userDonatedErc20Internal(instance), nil
		}
		meta.NextCursor = &next
	}

	return ListCosmicGameUserDonatedErc20200JSONResponse{
		CosmicGameUserDonatedErc20PageJSONResponse: CosmicGameUserDonatedErc20PageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// userDonatedErc20CursorFollows reports whether current comes strictly
// after previous in the (round DESC, token ASC) summary order.
func userDonatedErc20CursorFollows(current, previous cgstore.UserDonatedErc20PageCursor) bool {
	if current.Round != previous.Round {
		return current.Round < previous.Round
	}
	return current.TokenAid > previous.TokenAid
}

type userEventPageFetch[T any] func(
	context.Context,
	int64,
	*cgstore.UserEventPageCursor,
	int,
) ([]T, bool, error)

// listUserEventPage centralizes the event-keyed user history flow: address
// and limit validation, wallet-and-resource scoped cursor decoding, the
// unindexed-wallet empty page, repository cardinality/scope/order
// enforcement and continuation-cursor encoding.
func listUserEventPage[StoreRecord, APIRecord any](
	ctx context.Context,
	s *Server,
	rawAddress string,
	pathSegment string,
	cursor *Cursor,
	requestedLimit *Limit,
	resource userEventResource,
	fetch userEventPageFetch[StoreRecord],
	identity func(record StoreRecord, userAid int64, address string) (inScope bool, eventLogID int64),
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

	var after *cgstore.UserEventPageCursor
	if cursor != nil {
		decoded, err := decodeUserEventCursor(*cursor, scope, resource)
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
		after = &cgstore.UserEventPageCursor{EventLogID: decoded.EventLogID}
	}

	userAid, err := s.userHistories.UserAddressID(ctx, address)
	if errors.Is(err, store.ErrNotFound) {
		return []APIRecord{}, PageMeta{Limit: limit}, nil
	}
	if err != nil {
		s.logInternal(ctx, "resolve user address for history", err,
			"resource", string(resource),
			"address", address)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}

	records, hasMore, err := fetch(ctx, userAid, after, limit)
	if err != nil {
		s.logInternal(ctx, "list user history", err,
			"resource", string(resource),
			"address", address)
		problem := internalProblem(instance)
		return nil, PageMeta{}, &problem
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate user history page cardinality", err,
			"resource", string(resource),
			"address", address)
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
		inScope, eventLogID := identity(records[i], userAid, address)
		if !inScope || eventLogID < 1 ||
			(hasPrevious && eventLogID >= previousEventLogID) {
			s.logInternal(ctx, "validate user history page",
				errors.New("repository returned an out-of-scope or unordered user history row"),
				"resource", string(resource),
				"address", address,
				"event_log_id", eventLogID)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		mapped, err := mapRecord(records[i])
		if err != nil {
			s.logInternal(ctx, "map user history row", err,
				"resource", string(resource),
				"address", address,
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
			s.logInternal(ctx, "list user history",
				errors.New("repository reported another page without a cursor row"),
				"resource", string(resource),
				"address", address)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		next, err := encodeUserEventCursor(userEventCursor{
			Version:    userEventCursorVersion,
			Address:    scope,
			Resource:   resource,
			EventLogID: previousEventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode user history cursor", err,
				"resource", string(resource),
				"address", address)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		meta.NextCursor = &next
	}
	return data, meta, nil
}

func userPrizesBadRequest(problem Problem) ListCosmicGameUserPrizesResponseObject {
	return ListCosmicGameUserPrizes400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func userPrizesInternal(instance string) ListCosmicGameUserPrizesResponseObject {
	return ListCosmicGameUserPrizes500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(instance),
		),
	}
}

func userDonatedErc20BadRequest(problem Problem) ListCosmicGameUserDonatedErc20ResponseObject {
	return ListCosmicGameUserDonatedErc20400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func userDonatedErc20Internal(instance string) ListCosmicGameUserDonatedErc20ResponseObject {
	return ListCosmicGameUserDonatedErc20500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(instance),
		),
	}
}
