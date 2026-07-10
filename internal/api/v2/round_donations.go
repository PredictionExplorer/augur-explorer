package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ListRoundEthDonations implements
// GET /api/v2/cosmicgame/rounds/{round}/eth-donations.
func (s *Server) ListRoundEthDonations(
	ctx context.Context,
	request ListRoundEthDonationsRequestObject,
) (ListRoundEthDonationsResponseObject, error) {
	data, meta, problem := listRoundDonationPage(
		ctx,
		s,
		request.Round,
		request.Params.Cursor,
		request.Params.Limit,
		roundDonationResourceETH,
		"eth-donations",
		s.donations.EthDonationsByRoundPage,
		func(record cgstore.RoundEthDonationRecord) (int64, int64) {
			return record.RoundNum, record.Tx.EvtLogId
		},
		mapRoundEthDonation,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return listRoundEthDonationsBadRequest(*problem), nil
		}
		return listRoundEthDonationsInternal(*problem), nil
	}
	return ListRoundEthDonations200JSONResponse{
		RoundEthDonationPageJSONResponse: RoundEthDonationPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListRoundErc20Donations implements
// GET /api/v2/cosmicgame/rounds/{round}/erc20-donations.
func (s *Server) ListRoundErc20Donations(
	ctx context.Context,
	request ListRoundErc20DonationsRequestObject,
) (ListRoundErc20DonationsResponseObject, error) {
	data, meta, problem := listRoundDonationPage(
		ctx,
		s,
		request.Round,
		request.Params.Cursor,
		request.Params.Limit,
		roundDonationResourceERC20,
		"erc20-donations",
		s.donations.ERC20DonationsByRoundPage,
		func(record cgstore.RoundERC20DonationRecord) (int64, int64) {
			return record.RoundNum, record.Tx.EvtLogId
		},
		mapRoundERC20Donation,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return listRoundErc20DonationsBadRequest(*problem), nil
		}
		return listRoundErc20DonationsInternal(*problem), nil
	}
	return ListRoundErc20Donations200JSONResponse{
		RoundErc20DonationPageJSONResponse: RoundErc20DonationPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// ListRoundNftDonations implements
// GET /api/v2/cosmicgame/rounds/{round}/nft-donations.
func (s *Server) ListRoundNftDonations(
	ctx context.Context,
	request ListRoundNftDonationsRequestObject,
) (ListRoundNftDonationsResponseObject, error) {
	data, meta, problem := listRoundDonationPage(
		ctx,
		s,
		request.Round,
		request.Params.Cursor,
		request.Params.Limit,
		roundDonationResourceNFT,
		"nft-donations",
		s.donations.NFTDonationsByRoundPage,
		func(record cgstore.RoundNFTDonationRecord) (int64, int64) {
			return record.RoundNum, record.Tx.EvtLogId
		},
		mapRoundNFTDonation,
	)
	if problem != nil {
		if problem.Status == http.StatusBadRequest {
			return listRoundNftDonationsBadRequest(*problem), nil
		}
		return listRoundNftDonationsInternal(*problem), nil
	}
	return ListRoundNftDonations200JSONResponse{
		RoundNftDonationPageJSONResponse: RoundNftDonationPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

type roundDonationPageFetch[T any] func(
	context.Context,
	int64,
	*cgstore.DonationPageCursor,
	int,
) ([]T, bool, error)

func listRoundDonationPage[StoreRecord, APIRecord any](
	ctx context.Context,
	s *Server,
	round int64,
	cursor *Cursor,
	requestedLimit *Limit,
	resource roundDonationResource,
	pathSegment string,
	fetch roundDonationPageFetch[StoreRecord],
	identity func(StoreRecord) (round int64, eventLogID int64),
	mapRecord func(StoreRecord) (APIRecord, error),
) ([]APIRecord, PageMeta, *Problem) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d/%s", round, pathSegment)
	if round < 0 {
		problem := newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Round must be zero or greater.",
			instance,
		)
		return nil, PageMeta{}, &problem
	}
	limit, validLimit := resolvePageLimit(requestedLimit)
	if !validLimit {
		problem := newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)
		return nil, PageMeta{}, &problem
	}

	var after *cgstore.DonationPageCursor
	if cursor != nil {
		decoded, err := decodeRoundDonationCursor(*cursor, round, resource)
		if err != nil {
			problem := newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another round or donation resource.",
				instance,
			)
			return nil, PageMeta{}, &problem
		}
		after = &cgstore.DonationPageCursor{EventLogID: decoded.EventLogID}
	}

	records, hasMore, err := fetch(ctx, round, after, limit)
	if err != nil {
		s.logInternal(ctx, "list round donations", err,
			"resource", string(resource),
			"round", round)
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
		recordRound, eventLogID := identity(records[i])
		if recordRound != round || eventLogID < 1 ||
			(hasPrevious && eventLogID >= previousEventLogID) {
			err := errors.New("repository returned an out-of-scope or unordered donation page")
			s.logInternal(ctx, "validate round donation page", err,
				"resource", string(resource),
				"round", round,
				"record_round", recordRound,
				"event_log_id", eventLogID)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		mapped, err := mapRecord(records[i])
		if err != nil {
			s.logInternal(ctx, "map round donation", err,
				"resource", string(resource),
				"round", round,
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
			err := errors.New("repository reported another donation page without a cursor row")
			s.logInternal(ctx, "list round donations", err,
				"resource", string(resource),
				"round", round)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		next, err := encodeRoundDonationCursor(roundDonationCursor{
			Version:    roundDonationCursorVersion,
			Round:      round,
			Resource:   resource,
			EventLogID: previousEventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode round donation cursor", err,
				"resource", string(resource),
				"round", round)
			problem := internalProblem(instance)
			return nil, PageMeta{}, &problem
		}
		meta.NextCursor = &next
	}
	return data, meta, nil
}

func listRoundEthDonationsBadRequest(problem Problem) ListRoundEthDonationsResponseObject {
	return ListRoundEthDonations400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundEthDonationsInternal(problem Problem) ListRoundEthDonationsResponseObject {
	return ListRoundEthDonations500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundErc20DonationsBadRequest(problem Problem) ListRoundErc20DonationsResponseObject {
	return ListRoundErc20Donations400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundErc20DonationsInternal(problem Problem) ListRoundErc20DonationsResponseObject {
	return ListRoundErc20Donations500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundNftDonationsBadRequest(problem Problem) ListRoundNftDonationsResponseObject {
	return ListRoundNftDonations400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundNftDonationsInternal(problem Problem) ListRoundNftDonationsResponseObject {
	return ListRoundNftDonations500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}
