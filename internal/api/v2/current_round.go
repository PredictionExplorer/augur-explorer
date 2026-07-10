package v2

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
)

const (
	currentRoundInstance          = "/api/v2/cosmicgame/rounds/current"
	currentRoundRetryAfterSeconds = int(contractstate.DefaultVariablesInterval / time.Second)
)

// GetCurrentRound implements GET /api/v2/cosmicgame/rounds/current.
func (s *Server) GetCurrentRound(
	ctx context.Context,
	_ GetCurrentRoundRequestObject,
) (GetCurrentRoundResponseObject, error) {
	live, err := normalizeCurrentRoundSnapshot(s.contractState.Snapshot())
	if err != nil {
		if errors.Is(err, errCurrentRoundUnavailable) {
			return getCurrentRoundUnavailable(), nil
		}
		s.logInternal(ctx, "validate current-round live state", err)
		return getCurrentRoundInternal(), nil
	}

	stats, err := s.currentRounds.CosmicGameRoundStatistics(ctx, live.round)
	if err != nil {
		s.logInternal(ctx, "get current-round statistics", err, "round", live.round)
		return getCurrentRoundInternal(), nil
	}
	bidCount, err := s.currentRounds.BidCountForRound(ctx, live.round)
	if err != nil {
		s.logInternal(ctx, "get current-round bid count", err, "round", live.round)
		return getCurrentRoundInternal(), nil
	}

	current, err := mapCurrentRound(live, stats, bidCount)
	if err != nil {
		if errors.Is(err, errCurrentRoundUnavailable) {
			return getCurrentRoundUnavailable(), nil
		}
		s.logInternal(ctx, "map current round", err, "round", live.round)
		return getCurrentRoundInternal(), nil
	}

	return GetCurrentRound200JSONResponse{
		CosmicGameCurrentRoundJSONResponse: CosmicGameCurrentRoundJSONResponse(current),
	}, nil
}

func getCurrentRoundUnavailable() GetCurrentRoundResponseObject {
	return GetCurrentRound503ApplicationProblemPlusJSONResponse{
		ServiceUnavailableApplicationProblemPlusJSONResponse: ServiceUnavailableApplicationProblemPlusJSONResponse{
			Body: newProblem(
				http.StatusServiceUnavailable,
				"live-state-unavailable",
				"Live state unavailable",
				"The cached live round state is temporarily unavailable.",
				currentRoundInstance,
			),
			Headers: ServiceUnavailableResponseHeaders{
				RetryAfter: currentRoundRetryAfterSeconds,
			},
		},
	}
}

func getCurrentRoundInternal() GetCurrentRoundResponseObject {
	return GetCurrentRound500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(currentRoundInstance),
		),
	}
}
