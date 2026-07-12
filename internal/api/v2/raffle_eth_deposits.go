package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ListRoundRaffleEthDeposits implements
// GET /api/v2/cosmicgame/rounds/{round}/raffle-eth-deposits.
func (s *Server) ListRoundRaffleEthDeposits(
	ctx context.Context,
	request ListRoundRaffleEthDepositsRequestObject,
) (ListRoundRaffleEthDepositsResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d/raffle-eth-deposits", request.Round)
	if request.Round < 0 {
		return listRoundRaffleEthDepositsBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Round must be zero or greater.",
			instance,
		)), nil
	}
	limit, validLimit := resolvePageLimit(request.Params.Limit)
	if !validLimit {
		return listRoundRaffleEthDepositsBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	var after *cgstore.RaffleEthDepositPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeRaffleEthDepositCursor(*request.Params.Cursor, request.Round)
		if err != nil {
			return listRoundRaffleEthDepositsBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another round.",
				instance,
			)), nil
		}
		after = &cgstore.RaffleEthDepositPageCursor{
			WinnerIndex: cursor.WinnerIndex,
			EventLogID:  cursor.EventLogID,
		}
	}

	exists, err := s.raffles.CompletedRoundExists(ctx, request.Round)
	if err != nil {
		s.logInternal(ctx, "check completed round for raffle ETH deposits", err, "round", request.Round)
		return listRoundRaffleEthDepositsInternal(internalProblem(instance)), nil
	}
	if !exists {
		return listRoundRaffleEthDepositsNotFound(roundNotFoundProblem(instance)), nil
	}

	records, hasMore, err := s.raffles.RaffleEthDepositsByRoundPage(ctx, request.Round, after, limit)
	if err != nil {
		s.logInternal(ctx, "list round raffle ETH deposits", err, "round", request.Round)
		return listRoundRaffleEthDepositsInternal(internalProblem(instance)), nil
	}
	if err := validatePageCardinality(len(records), limit); err != nil {
		s.logInternal(ctx, "validate raffle ETH deposit page cardinality", err, "round", request.Round)
		return listRoundRaffleEthDepositsInternal(internalProblem(instance)), nil
	}

	data := make([]RoundRaffleEthDeposit, 0, len(records))
	previous := after
	for i := range records {
		current := cgstore.RaffleEthDepositPageCursor{
			WinnerIndex: records[i].WinnerIndex,
			EventLogID:  records[i].Tx.EvtLogId,
		}
		if records[i].RoundNum != request.Round ||
			(previous != nil && !raffleEthDepositCursorFollows(current, *previous)) {
			err := errors.New("repository returned an out-of-scope or unordered raffle ETH deposit page")
			s.logInternal(ctx, "validate raffle ETH deposit page", err,
				"round", request.Round,
				"winner_index", current.WinnerIndex,
				"event_log_id", current.EventLogID)
			return listRoundRaffleEthDepositsInternal(internalProblem(instance)), nil
		}
		deposit, err := mapRoundRaffleEthDeposit(records[i])
		if err != nil {
			s.logInternal(ctx, "map raffle ETH deposit", err,
				"round", request.Round,
				"winner_index", current.WinnerIndex,
				"event_log_id", current.EventLogID)
			return listRoundRaffleEthDepositsInternal(internalProblem(instance)), nil
		}
		data = append(data, deposit)
		previous = &current
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if previous == nil || len(records) == 0 {
			err := errors.New("repository reported another raffle ETH deposit page without a cursor row")
			s.logInternal(ctx, "list round raffle ETH deposits", err, "round", request.Round)
			return listRoundRaffleEthDepositsInternal(internalProblem(instance)), nil
		}
		next, err := encodeRaffleEthDepositCursor(raffleEthDepositCursor{
			Version:     raffleEthDepositCursorVersion,
			Round:       request.Round,
			WinnerIndex: previous.WinnerIndex,
			EventLogID:  previous.EventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode raffle ETH deposit cursor", err, "round", request.Round)
			return listRoundRaffleEthDepositsInternal(internalProblem(instance)), nil
		}
		meta.NextCursor = &next
	}

	return ListRoundRaffleEthDeposits200JSONResponse{
		RoundRaffleEthDepositPageJSONResponse: RoundRaffleEthDepositPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func raffleEthDepositCursorFollows(
	current,
	previous cgstore.RaffleEthDepositPageCursor,
) bool {
	return current.WinnerIndex > previous.WinnerIndex ||
		(current.WinnerIndex == previous.WinnerIndex && current.EventLogID > previous.EventLogID)
}

func listRoundRaffleEthDepositsBadRequest(problem Problem) ListRoundRaffleEthDepositsResponseObject {
	return ListRoundRaffleEthDeposits400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundRaffleEthDepositsNotFound(problem Problem) ListRoundRaffleEthDepositsResponseObject {
	return ListRoundRaffleEthDeposits404ApplicationProblemPlusJSONResponse{
		NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundRaffleEthDepositsInternal(problem Problem) ListRoundRaffleEthDepositsResponseObject {
	return ListRoundRaffleEthDeposits500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}
