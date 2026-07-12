package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ListRoundPrizes implements
// GET /api/v2/cosmicgame/rounds/{round}/prizes.
func (s *Server) ListRoundPrizes(
	ctx context.Context,
	request ListRoundPrizesRequestObject,
) (ListRoundPrizesResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d/prizes", request.Round)
	if request.Round < 0 {
		return listRoundPrizesBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Round must be zero or greater.",
			instance,
		)), nil
	}

	limit, validLimit := resolvePageLimit(request.Params.Limit)
	if !validLimit {
		return listRoundPrizesBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	var after *cgstore.PrizePageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodePrizeCursor(*request.Params.Cursor, request.Round)
		if err != nil {
			return listRoundPrizesBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another round.",
				instance,
			)), nil
		}
		after = &cgstore.PrizePageCursor{
			PrizeType:   cursor.PrizeType,
			WinnerIndex: cursor.WinnerIndex,
		}
	}

	exists, err := s.prizes.CompletedRoundExists(ctx, request.Round)
	if err != nil {
		s.logInternal(ctx, "check completed round for prize list", err, "round", request.Round)
		return listRoundPrizesInternal(internalProblem(instance)), nil
	}
	if !exists {
		return listRoundPrizesNotFound(roundNotFoundProblem(instance)), nil
	}

	records, hasMore, err := s.prizes.AllPrizesForRoundPage(ctx, request.Round, after, limit)
	if err != nil {
		s.logInternal(ctx, "list round prizes", err, "round", request.Round)
		return listRoundPrizesInternal(internalProblem(instance)), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate round prize page cardinality", err, "round", request.Round)
		return listRoundPrizesInternal(internalProblem(instance)), nil
	}

	data := make([]RoundPrize, 0, len(records))
	previous := after
	for i := range records {
		current := cgstore.PrizePageCursor{
			PrizeType:   records[i].RecordType,
			WinnerIndex: records[i].WinnerIndex,
		}
		if records[i].RoundNum != request.Round ||
			(previous != nil && !prizeCursorFollows(current, *previous)) {
			err := errors.New("repository returned an out-of-scope or unordered prize page")
			s.logInternal(ctx, "validate round prize page", err,
				"round", request.Round,
				"prize_type", current.PrizeType,
				"winner_index", current.WinnerIndex)
			return listRoundPrizesInternal(internalProblem(instance)), nil
		}

		prize, err := mapRoundPrize(records[i])
		if err != nil {
			s.logInternal(ctx, "map round prize", err,
				"round", request.Round,
				"prize_type", current.PrizeType,
				"winner_index", current.WinnerIndex)
			return listRoundPrizesInternal(internalProblem(instance)), nil
		}
		data = append(data, prize)
		previous = &current
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if previous == nil || len(records) == 0 {
			err := errors.New("repository reported another prize page without a cursor row")
			s.logInternal(ctx, "list round prizes", err, "round", request.Round)
			return listRoundPrizesInternal(internalProblem(instance)), nil
		}
		next, err := encodePrizeCursor(prizeCursor{
			Version:     prizeCursorVersion,
			Round:       request.Round,
			PrizeType:   previous.PrizeType,
			WinnerIndex: previous.WinnerIndex,
		})
		if err != nil {
			s.logInternal(ctx, "encode round-prize cursor", err, "round", request.Round)
			return listRoundPrizesInternal(internalProblem(instance)), nil
		}
		meta.NextCursor = &next
	}

	return ListRoundPrizes200JSONResponse{
		RoundPrizePageJSONResponse: RoundPrizePageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func prizeCursorFollows(current, previous cgstore.PrizePageCursor) bool {
	return current.PrizeType > previous.PrizeType ||
		(current.PrizeType == previous.PrizeType && current.WinnerIndex > previous.WinnerIndex)
}

func listRoundPrizesBadRequest(problem Problem) ListRoundPrizesResponseObject {
	return ListRoundPrizes400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundPrizesNotFound(problem Problem) ListRoundPrizesResponseObject {
	return ListRoundPrizes404ApplicationProblemPlusJSONResponse{
		NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundPrizesInternal(problem Problem) ListRoundPrizesResponseObject {
	return ListRoundPrizes500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}
