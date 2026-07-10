package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ListRounds implements GET /api/v2/cosmicgame/rounds.
func (s *Server) ListRounds(ctx context.Context, request ListRoundsRequestObject) (ListRoundsResponseObject, error) {
	const instance = "/api/v2/cosmicgame/rounds"

	limit, validLimit := resolvePageLimit(request.Params.Limit)
	if !validLimit {
		return listRoundsBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	var after *cgstore.RoundPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeRoundCursor(*request.Params.Cursor)
		if err != nil {
			return listRoundsBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed or uses an unsupported version.",
				instance,
			)), nil
		}
		after = &cgstore.RoundPageCursor{
			RoundNum:   cursor.RoundNum,
			EventLogID: cursor.EventLogID,
		}
	}

	records, hasMore, err := s.rounds.PrizeClaimsPage(ctx, after, limit)
	if err != nil {
		s.logInternal(ctx, "list completed rounds", err)
		return listRoundsInternal(internalProblem(instance)), nil
	}

	data := make([]CosmicGameRoundSummary, 0, len(records))
	previous := after
	for i := range records {
		roundNum, err := roundNumber(records[i].RoundNum)
		if err != nil {
			s.logInternal(ctx, "validate completed round page", err)
			return listRoundsInternal(internalProblem(instance)), nil
		}
		current := cgstore.RoundPageCursor{
			RoundNum:   roundNum,
			EventLogID: records[i].ClaimPrizeTx.Tx.EvtLogId,
		}
		if previous != nil && !roundCursorPrecedes(current, *previous) {
			err := errors.New("repository returned an unordered completed-round page")
			s.logInternal(ctx, "validate completed round page", err,
				"round", roundNum,
				"event_log_id", current.EventLogID)
			return listRoundsInternal(internalProblem(instance)), nil
		}

		summary, err := mapRoundSummary(records[i])
		if err != nil {
			s.logInternal(ctx, "map completed round summary", err,
				"round", roundNum,
				"event_log_id", current.EventLogID)
			return listRoundsInternal(internalProblem(instance)), nil
		}
		data = append(data, summary)
		previous = &current
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if previous == nil || len(records) == 0 {
			err := errors.New("repository reported another completed-round page without a cursor row")
			s.logInternal(ctx, "list completed rounds", err)
			return listRoundsInternal(internalProblem(instance)), nil
		}
		next, err := encodeRoundCursor(roundCursor{
			Version:    roundCursorVersion,
			RoundNum:   previous.RoundNum,
			EventLogID: previous.EventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode completed-round cursor", err)
			return listRoundsInternal(internalProblem(instance)), nil
		}
		meta.NextCursor = &next
	}

	return ListRounds200JSONResponse{
		RoundPageJSONResponse: RoundPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

// GetRound implements GET /api/v2/cosmicgame/rounds/{round}.
func (s *Server) GetRound(ctx context.Context, request GetRoundRequestObject) (GetRoundResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d", request.Round)
	if request.Round < 0 {
		return getRoundBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Round must be zero or greater.",
			instance,
		)), nil
	}

	record, err := s.rounds.RoundInfo(ctx, request.Round)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return GetRound404ApplicationProblemPlusJSONResponse{
				NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(newProblem(
					http.StatusNotFound,
					"round-not-found",
					"Round not found",
					"No completed round exists with that number.",
					instance,
				)),
			}, nil
		}
		s.logInternal(ctx, "get completed round", err, "round", request.Round)
		return getRoundInternal(internalProblem(instance)), nil
	}

	roundNum, err := roundNumber(record.RoundNum)
	if err != nil || roundNum != request.Round {
		if err == nil {
			err = errors.New("repository returned a round outside the requested identity")
		}
		s.logInternal(ctx, "validate completed round", err,
			"requested_round", request.Round,
			"returned_round", record.RoundNum)
		return getRoundInternal(internalProblem(instance)), nil
	}

	mapped, err := mapRound(record)
	if err != nil {
		s.logInternal(ctx, "map completed round", err,
			"round", request.Round,
			"event_log_id", record.ClaimPrizeTx.Tx.EvtLogId)
		return getRoundInternal(internalProblem(instance)), nil
	}
	return GetRound200JSONResponse{
		CosmicGameRoundJSONResponse: CosmicGameRoundJSONResponse(mapped),
	}, nil
}

func roundCursorPrecedes(current, previous cgstore.RoundPageCursor) bool {
	return current.RoundNum < previous.RoundNum ||
		(current.RoundNum == previous.RoundNum && current.EventLogID < previous.EventLogID)
}

func listRoundsBadRequest(problem Problem) ListRoundsResponseObject {
	return ListRounds400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundsInternal(problem Problem) ListRoundsResponseObject {
	return ListRounds500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}

func getRoundBadRequest(problem Problem) GetRoundResponseObject {
	return GetRound400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func getRoundInternal(problem Problem) GetRoundResponseObject {
	return GetRound500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}
