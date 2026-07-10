package v2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// ListRoundRaffleNftWinners implements
// GET /api/v2/cosmicgame/rounds/{round}/raffle-nft-winners.
func (s *Server) ListRoundRaffleNftWinners(
	ctx context.Context,
	request ListRoundRaffleNftWinnersRequestObject,
) (ListRoundRaffleNftWinnersResponseObject, error) {
	instance := fmt.Sprintf("/api/v2/cosmicgame/rounds/%d/raffle-nft-winners", request.Round)
	if request.Round < 0 {
		return listRoundRaffleNftWinnersBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Round must be zero or greater.",
			instance,
		)), nil
	}
	pool := request.Params.Pool
	if !pool.Valid() {
		return listRoundRaffleNftWinnersBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			"Pool must be bidder or randomWalkStaker.",
			instance,
		)), nil
	}
	limit, validLimit := resolvePageLimit(request.Params.Limit)
	if !validLimit {
		return listRoundRaffleNftWinnersBadRequest(newProblem(
			http.StatusBadRequest,
			"invalid-parameter",
			"Invalid parameter",
			pageLimitProblemDetail(),
			instance,
		)), nil
	}

	var after *cgstore.RaffleNFTWinnerPageCursor
	if request.Params.Cursor != nil {
		cursor, err := decodeRaffleNftWinnerCursor(*request.Params.Cursor, request.Round, pool)
		if err != nil {
			return listRoundRaffleNftWinnersBadRequest(newProblem(
				http.StatusBadRequest,
				"invalid-cursor",
				"Invalid cursor",
				"The cursor is malformed, unsupported, or belongs to another round or pool.",
				instance,
			)), nil
		}
		after = &cgstore.RaffleNFTWinnerPageCursor{
			WinnerIndex: cursor.WinnerIndex,
			EventLogID:  cursor.EventLogID,
		}
	}

	exists, err := s.raffles.CompletedRoundExists(ctx, request.Round)
	if err != nil {
		s.logInternal(ctx, "check completed round for raffle NFT winners", err,
			"round", request.Round,
			"pool", pool)
		return listRoundRaffleNftWinnersInternal(internalProblem(instance)), nil
	}
	if !exists {
		return listRoundRaffleNftWinnersNotFound(roundNotFoundProblem(instance)), nil
	}

	isStaker := pool == RandomWalkStaker
	records, hasMore, err := s.raffles.RaffleNFTWinnersByRoundPage(
		ctx,
		request.Round,
		isStaker,
		after,
		limit,
	)
	if err != nil {
		s.logInternal(ctx, "list round raffle NFT winners", err,
			"round", request.Round,
			"pool", pool)
		return listRoundRaffleNftWinnersInternal(internalProblem(instance)), nil
	}

	data := make([]RoundRaffleNftWinner, 0, len(records))
	previous := after
	for i := range records {
		current := cgstore.RaffleNFTWinnerPageCursor{
			WinnerIndex: records[i].WinnerIndex,
			EventLogID:  records[i].Tx.EvtLogId,
		}
		if records[i].RoundNum != request.Round ||
			records[i].IsStaker != isStaker ||
			(previous != nil && !raffleNftWinnerCursorFollows(current, *previous)) {
			err := errors.New("repository returned an out-of-scope or unordered raffle NFT winner page")
			s.logInternal(ctx, "validate raffle NFT winner page", err,
				"round", request.Round,
				"pool", pool,
				"winner_index", current.WinnerIndex,
				"event_log_id", current.EventLogID)
			return listRoundRaffleNftWinnersInternal(internalProblem(instance)), nil
		}
		winner, err := mapRoundRaffleNftWinner(records[i], isStaker)
		if err != nil {
			s.logInternal(ctx, "map raffle NFT winner", err,
				"round", request.Round,
				"pool", pool,
				"winner_index", current.WinnerIndex,
				"event_log_id", current.EventLogID)
			return listRoundRaffleNftWinnersInternal(internalProblem(instance)), nil
		}
		data = append(data, winner)
		previous = &current
	}

	meta := PageMeta{Limit: limit}
	if hasMore {
		if previous == nil || len(records) == 0 {
			err := errors.New("repository reported another raffle NFT winner page without a cursor row")
			s.logInternal(ctx, "list round raffle NFT winners", err,
				"round", request.Round,
				"pool", pool)
			return listRoundRaffleNftWinnersInternal(internalProblem(instance)), nil
		}
		next, err := encodeRaffleNftWinnerCursor(raffleNftWinnerCursor{
			Version:     raffleNftWinnerCursorVersion,
			Round:       request.Round,
			Pool:        pool,
			WinnerIndex: previous.WinnerIndex,
			EventLogID:  previous.EventLogID,
		})
		if err != nil {
			s.logInternal(ctx, "encode raffle NFT winner cursor", err,
				"round", request.Round,
				"pool", pool)
			return listRoundRaffleNftWinnersInternal(internalProblem(instance)), nil
		}
		meta.NextCursor = &next
	}

	return ListRoundRaffleNftWinners200JSONResponse{
		RoundRaffleNftWinnerPageJSONResponse: RoundRaffleNftWinnerPageJSONResponse{
			Data: data,
			Meta: meta,
		},
	}, nil
}

func raffleNftWinnerCursorFollows(
	current,
	previous cgstore.RaffleNFTWinnerPageCursor,
) bool {
	return current.WinnerIndex > previous.WinnerIndex ||
		(current.WinnerIndex == previous.WinnerIndex && current.EventLogID > previous.EventLogID)
}

func listRoundRaffleNftWinnersBadRequest(problem Problem) ListRoundRaffleNftWinnersResponseObject {
	return ListRoundRaffleNftWinners400ApplicationProblemPlusJSONResponse{
		BadRequestApplicationProblemPlusJSONResponse: BadRequestApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundRaffleNftWinnersNotFound(problem Problem) ListRoundRaffleNftWinnersResponseObject {
	return ListRoundRaffleNftWinners404ApplicationProblemPlusJSONResponse{
		NotFoundApplicationProblemPlusJSONResponse: NotFoundApplicationProblemPlusJSONResponse(problem),
	}
}

func listRoundRaffleNftWinnersInternal(problem Problem) ListRoundRaffleNftWinnersResponseObject {
	return ListRoundRaffleNftWinners500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(problem),
	}
}
