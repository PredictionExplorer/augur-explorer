package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

func (s *Server) ListCosmicGameBidders(
	ctx context.Context,
	request ListCosmicGameBiddersRequestObject,
) (ListCosmicGameBiddersResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/participants/bidders"
	limit, after, problem := participantPageInput(
		request.Params.Cursor, request.Params.Limit, cgstore.ParticipantBidders, instance,
	)
	if problem != nil {
		return ListCosmicGameBidders400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	records, hasMore, err := s.participants.BidderParticipantsPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list bidder participants", err)
		return ListCosmicGameBidders500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	data, meta, err := buildParticipantPage(
		records, hasMore, cgstore.ParticipantBidders, after, limit,
		func(record cgstore.BidderParticipantRecord) cgstore.ParticipantPageCursor {
			return cgstore.ParticipantPageCursor{
				Kind:      cgstore.ParticipantBidders,
				SortValue: strconv.FormatInt(record.BidCount, 10),
				AddressID: record.BidderAid,
			}
		},
		mapBidderParticipant,
	)
	if err != nil {
		s.logInternal(ctx, "build bidder participant page", err)
		return ListCosmicGameBidders500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	return ListCosmicGameBidders200JSONResponse{
		BidderParticipantPageJSONResponse: BidderParticipantPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func (s *Server) ListCosmicGameWinners(
	ctx context.Context,
	request ListCosmicGameWinnersRequestObject,
) (ListCosmicGameWinnersResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/participants/winners"
	limit, after, problem := participantPageInput(
		request.Params.Cursor, request.Params.Limit, cgstore.ParticipantWinners, instance,
	)
	if problem != nil {
		return ListCosmicGameWinners400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	records, hasMore, err := s.participants.WinnerParticipantsPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list winner participants", err)
		return ListCosmicGameWinners500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	data, meta, err := buildParticipantPage(
		records, hasMore, cgstore.ParticipantWinners, after, limit,
		func(record cgstore.WinnerParticipantRecord) cgstore.ParticipantPageCursor {
			return cgstore.ParticipantPageCursor{
				Kind:      cgstore.ParticipantWinners,
				SortValue: strconv.FormatInt(record.PrizeCount, 10),
				AddressID: record.WinnerAid,
			}
		},
		mapWinnerParticipant,
	)
	if err != nil {
		s.logInternal(ctx, "build winner participant page", err)
		return ListCosmicGameWinners500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	return ListCosmicGameWinners200JSONResponse{
		WinnerParticipantPageJSONResponse: WinnerParticipantPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func (s *Server) ListCosmicGameDonors(
	ctx context.Context,
	request ListCosmicGameDonorsRequestObject,
) (ListCosmicGameDonorsResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/participants/donors"
	limit, after, problem := participantPageInput(
		request.Params.Cursor, request.Params.Limit, cgstore.ParticipantDonors, instance,
	)
	if problem != nil {
		return ListCosmicGameDonors400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	records, hasMore, err := s.participants.DonorParticipantsPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list donor participants", err)
		return ListCosmicGameDonors500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	data, meta, err := buildParticipantPage(
		records, hasMore, cgstore.ParticipantDonors, after, limit,
		func(record cgstore.DonorParticipantRecord) cgstore.ParticipantPageCursor {
			return cgstore.ParticipantPageCursor{
				Kind:      cgstore.ParticipantDonors,
				SortValue: record.TotalDonatedWei,
				AddressID: record.DonorAid,
			}
		},
		mapDonorParticipant,
	)
	if err != nil {
		s.logInternal(ctx, "build donor participant page", err)
		return ListCosmicGameDonors500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	return ListCosmicGameDonors200JSONResponse{
		DonorParticipantPageJSONResponse: DonorParticipantPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func (s *Server) ListCosmicGameCstStakers(
	ctx context.Context,
	request ListCosmicGameCstStakersRequestObject,
) (ListCosmicGameCstStakersResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/participants/stakers/cst"
	limit, after, problem := participantPageInput(
		request.Params.Cursor, request.Params.Limit, cgstore.ParticipantCSTStakers, instance,
	)
	if problem != nil {
		return ListCosmicGameCstStakers400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	records, hasMore, err := s.participants.CSTStakerParticipantsPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list CST staker participants", err)
		return ListCosmicGameCstStakers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	data, meta, err := buildParticipantPage(
		records, hasMore, cgstore.ParticipantCSTStakers, after, limit,
		func(record cgstore.CSTStakerParticipantRecord) cgstore.ParticipantPageCursor {
			return cgstore.ParticipantPageCursor{
				Kind:      cgstore.ParticipantCSTStakers,
				SortValue: record.TotalRewardWei,
				AddressID: record.StakerAid,
			}
		},
		mapCSTStakerParticipant,
	)
	if err != nil {
		s.logInternal(ctx, "build CST staker participant page", err)
		return ListCosmicGameCstStakers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	return ListCosmicGameCstStakers200JSONResponse{
		CstStakerParticipantPageJSONResponse: CstStakerParticipantPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func (s *Server) ListCosmicGameRandomWalkStakers(
	ctx context.Context,
	request ListCosmicGameRandomWalkStakersRequestObject,
) (ListCosmicGameRandomWalkStakersResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/participants/stakers/random-walk"
	limit, after, problem := participantPageInput(
		request.Params.Cursor, request.Params.Limit, cgstore.ParticipantRandomWalkStakers, instance,
	)
	if problem != nil {
		return ListCosmicGameRandomWalkStakers400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	records, hasMore, err := s.participants.RandomWalkStakerParticipantsPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list RandomWalk staker participants", err)
		return ListCosmicGameRandomWalkStakers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	data, meta, err := buildParticipantPage(
		records, hasMore, cgstore.ParticipantRandomWalkStakers, after, limit,
		func(record cgstore.RandomWalkStakerParticipantRecord) cgstore.ParticipantPageCursor {
			return cgstore.ParticipantPageCursor{
				Kind:      cgstore.ParticipantRandomWalkStakers,
				SortValue: strconv.FormatInt(record.StakedTokenCount, 10),
				AddressID: record.StakerAid,
			}
		},
		mapRandomWalkStakerParticipant,
	)
	if err != nil {
		s.logInternal(ctx, "build RandomWalk staker participant page", err)
		return ListCosmicGameRandomWalkStakers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	return ListCosmicGameRandomWalkStakers200JSONResponse{
		RandomWalkStakerParticipantPageJSONResponse: RandomWalkStakerParticipantPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func (s *Server) ListCosmicGameDualStakers(
	ctx context.Context,
	request ListCosmicGameDualStakersRequestObject,
) (ListCosmicGameDualStakersResponseObject, error) {
	const instance = "/api/v2/cosmicgame/statistics/participants/stakers/both"
	limit, after, problem := participantPageInput(
		request.Params.Cursor, request.Params.Limit, cgstore.ParticipantDualStakers, instance,
	)
	if problem != nil {
		return ListCosmicGameDualStakers400ApplicationProblemPlusJSONResponse{
			BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(*problem),
		}, nil
	}
	records, hasMore, err := s.participants.DualStakerParticipantsPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list dual-staker participants", err)
		return ListCosmicGameDualStakers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	data, meta, err := buildParticipantPage(
		records, hasMore, cgstore.ParticipantDualStakers, after, limit,
		func(record cgstore.DualStakerParticipantRecord) cgstore.ParticipantPageCursor {
			return cgstore.ParticipantPageCursor{
				Kind:      cgstore.ParticipantDualStakers,
				SortValue: strconv.FormatInt(record.TotalStakedTokenCount, 10),
				AddressID: record.StakerAid,
			}
		},
		mapDualStakerParticipant,
	)
	if err != nil {
		s.logInternal(ctx, "build dual-staker participant page", err)
		return ListCosmicGameDualStakers500ApplicationProblemPlusJSONResponse{
			InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(internalProblem(instance)),
		}, nil
	}
	return ListCosmicGameDualStakers200JSONResponse{
		DualStakerParticipantPageJSONResponse: DualStakerParticipantPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func participantPageInput(
	encoded *string,
	requestedLimit *int,
	kind cgstore.ParticipantKind,
	instance string,
) (int, *cgstore.ParticipantPageCursor, *Problem) {
	limit, valid := resolvePageLimit(requestedLimit)
	if !valid {
		problem := newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)
		return 0, nil, &problem
	}
	if encoded == nil {
		return limit, nil, nil
	}
	cursor, err := decodeParticipantCursor(*encoded, kind)
	if err != nil {
		problem := newProblem(
			http.StatusBadRequest,
			"invalid-cursor",
			"Invalid cursor",
			"The cursor is malformed, unsupported, or belongs to another participant directory.",
			instance,
		)
		return 0, nil, &problem
	}
	return limit, &cgstore.ParticipantPageCursor{
		Kind:      cursor.Kind,
		SortValue: cursor.SortValue,
		AddressID: cursor.AddressID,
	}, nil
}

func buildParticipantPage[Record, Model any](
	records []Record,
	hasMore bool,
	kind cgstore.ParticipantKind,
	after *cgstore.ParticipantPageCursor,
	limit int,
	cursorFor func(Record) cgstore.ParticipantPageCursor,
	mapRecord func(Record) (Model, error),
) ([]Model, PageMeta, error) {
	if len(records) > limit {
		return nil, PageMeta{}, errors.New("repository returned more participant rows than requested")
	}
	data := make([]Model, 0, len(records))
	previous := after
	for i := range records {
		current := cursorFor(records[i])
		if current.Kind != kind ||
			!validParticipantCursor(participantCursor{
				Version:   participantCursorVersion,
				Kind:      current.Kind,
				SortValue: current.SortValue,
				AddressID: current.AddressID,
			}, kind) ||
			(previous != nil && !participantCursorFollows(current, *previous)) {
			return nil, PageMeta{}, errors.New("repository returned an unordered participant page")
		}
		mapped, err := mapRecord(records[i])
		if err != nil {
			return nil, PageMeta{}, err
		}
		data = append(data, mapped)
		previous = &current
	}
	meta := PageMeta{Limit: limit}
	if hasMore {
		if len(records) == 0 || previous == nil {
			return nil, PageMeta{}, errors.New("participant page hasMore without a cursor row")
		}
		next, err := encodeParticipantCursor(participantCursor{
			Version:   participantCursorVersion,
			Kind:      previous.Kind,
			SortValue: previous.SortValue,
			AddressID: previous.AddressID,
		})
		if err != nil {
			return nil, PageMeta{}, fmt.Errorf("encode participant cursor: %w", err)
		}
		meta.NextCursor = &next
	}
	return data, meta, nil
}

func participantCursorFollows(
	current cgstore.ParticipantPageCursor,
	previous cgstore.ParticipantPageCursor,
) bool {
	if current.Kind != previous.Kind {
		return false
	}
	comparison, err := compareDecimal(current.SortValue, previous.SortValue)
	if err != nil || comparison > 0 {
		return false
	}
	if comparison < 0 {
		return true
	}
	return current.AddressID > previous.AddressID
}
