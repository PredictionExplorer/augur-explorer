package v2

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/cosmicgame/contractstate"
)

const (
	contractAddressesInstance     = "/api/v2/cosmicgame/contracts/addresses"
	contractConfigurationInstance = "/api/v2/cosmicgame/contracts/configuration"
	contractBalancesInstance      = "/api/v2/cosmicgame/contracts/balances"
	currentBidPricesInstance      = "/api/v2/cosmicgame/rounds/current/bid-prices"
	currentSpecialWinnersInstance = "/api/v2/cosmicgame/rounds/current/special-winners"
	contractCacheRetryAfter       = int(contractstate.DefaultVariablesInterval / time.Second)
	contractConstantsRetryAfter   = int(contractstate.DefaultConstantsInterval / time.Second)
	specialWinnersRetryAfter      = int(contractstate.DefaultSpecialWinnersInterval / time.Second)
)

// GetCosmicGameContractAddresses implements
// GET /api/v2/cosmicgame/contracts/addresses.
func (s *Server) GetCosmicGameContractAddresses(
	ctx context.Context,
	_ GetCosmicGameContractAddressesRequestObject,
) (GetCosmicGameContractAddressesResponseObject, error) {
	record, err := s.contractAddresses.ContractAddrs(ctx)
	if err != nil {
		s.logInternal(ctx, "get contract address registry", err)
		return getContractAddressesInternal(), nil
	}
	result, err := mapContractAddressRegistry(record)
	if err != nil {
		s.logInternal(ctx, "map contract address registry", err)
		return getContractAddressesInternal(), nil
	}
	return GetCosmicGameContractAddresses200JSONResponse{
		CosmicGameContractAddressesJSONResponse: CosmicGameContractAddressesJSONResponse(result),
	}, nil
}

// GetCosmicGameContractConfiguration implements
// GET /api/v2/cosmicgame/contracts/configuration.
func (s *Server) GetCosmicGameContractConfiguration(
	ctx context.Context,
	_ GetCosmicGameContractConfigurationRequestObject,
) (GetCosmicGameContractConfigurationResponseObject, error) {
	snapshot := s.contractState.Snapshot()
	result, err := mapContractConfiguration(snapshot)
	if err != nil {
		if errors.Is(err, errCachedLiveUnavailable) {
			retryAfter := contractCacheRetryAfter
			if !snapshot.ConstantsReady {
				retryAfter = contractConstantsRetryAfter
			}
			return getContractConfigurationUnavailable(retryAfter), nil
		}
		s.logInternal(ctx, "map cached contract configuration", err)
		return getContractConfigurationInternal(), nil
	}
	return GetCosmicGameContractConfiguration200JSONResponse{
		CosmicGameContractConfigurationJSONResponse: CosmicGameContractConfigurationJSONResponse(result),
	}, nil
}

// GetCosmicGameContractBalances implements
// GET /api/v2/cosmicgame/contracts/balances.
func (s *Server) GetCosmicGameContractBalances(
	ctx context.Context,
	_ GetCosmicGameContractBalancesRequestObject,
) (GetCosmicGameContractBalancesResponseObject, error) {
	result, err := mapContractBalances(s.contractState.Snapshot())
	if err != nil {
		if errors.Is(err, errCachedLiveUnavailable) {
			return getContractBalancesUnavailable(), nil
		}
		s.logInternal(ctx, "map cached contract balances", err)
		return getContractBalancesInternal(), nil
	}
	return GetCosmicGameContractBalances200JSONResponse{
		CosmicGameContractBalancesJSONResponse: CosmicGameContractBalancesJSONResponse(result),
	}, nil
}

// GetCosmicGameCurrentBidPrices implements
// GET /api/v2/cosmicgame/rounds/current/bid-prices.
func (s *Server) GetCosmicGameCurrentBidPrices(
	ctx context.Context,
	_ GetCosmicGameCurrentBidPricesRequestObject,
) (GetCosmicGameCurrentBidPricesResponseObject, error) {
	result, err := mapCurrentBidPrices(s.contractState.Snapshot())
	if err != nil {
		if errors.Is(err, errCachedLiveUnavailable) {
			return getCurrentBidPricesUnavailable(), nil
		}
		s.logInternal(ctx, "map cached current bid prices", err)
		return getCurrentBidPricesInternal(), nil
	}
	return GetCosmicGameCurrentBidPrices200JSONResponse{
		CosmicGameCurrentBidPricesJSONResponse: CosmicGameCurrentBidPricesJSONResponse(result),
	}, nil
}

// GetCosmicGameCurrentSpecialWinners implements
// GET /api/v2/cosmicgame/rounds/current/special-winners.
func (s *Server) GetCosmicGameCurrentSpecialWinners(
	ctx context.Context,
	_ GetCosmicGameCurrentSpecialWinnersRequestObject,
) (GetCosmicGameCurrentSpecialWinnersResponseObject, error) {
	result, err := mapCurrentSpecialWinners(s.contractState.Snapshot())
	if err != nil {
		if errors.Is(err, errCachedLiveUnavailable) {
			return getCurrentSpecialWinnersUnavailable(), nil
		}
		s.logInternal(ctx, "map cached current special winners", err)
		return getCurrentSpecialWinnersInternal(), nil
	}
	return GetCosmicGameCurrentSpecialWinners200JSONResponse{
		CosmicGameCurrentSpecialWinnersJSONResponse: CosmicGameCurrentSpecialWinnersJSONResponse(result),
	}, nil
}

func cachedLiveProblem(instance, resource string) Problem {
	return newProblem(
		http.StatusServiceUnavailable,
		"live-state-unavailable",
		"Live state unavailable",
		"The cached live state for "+resource+" is temporarily unavailable.",
		instance,
	)
}

func getContractAddressesInternal() GetCosmicGameContractAddressesResponseObject {
	return GetCosmicGameContractAddresses500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(contractAddressesInstance),
		),
	}
}

func getContractConfigurationUnavailable(retryAfter int) GetCosmicGameContractConfigurationResponseObject {
	return GetCosmicGameContractConfiguration503ApplicationProblemPlusJSONResponse{
		ServiceUnavailableApplicationProblemPlusJSONResponse: ServiceUnavailableApplicationProblemPlusJSONResponse{
			Body:    cachedLiveProblem(contractConfigurationInstance, "contract configuration"),
			Headers: ServiceUnavailableResponseHeaders{RetryAfter: retryAfter},
		},
	}
}

func getContractConfigurationInternal() GetCosmicGameContractConfigurationResponseObject {
	return GetCosmicGameContractConfiguration500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(contractConfigurationInstance),
		),
	}
}

func getContractBalancesUnavailable() GetCosmicGameContractBalancesResponseObject {
	return GetCosmicGameContractBalances503ApplicationProblemPlusJSONResponse{
		ServiceUnavailableApplicationProblemPlusJSONResponse: ServiceUnavailableApplicationProblemPlusJSONResponse{
			Body:    cachedLiveProblem(contractBalancesInstance, "contract balances"),
			Headers: ServiceUnavailableResponseHeaders{RetryAfter: contractCacheRetryAfter},
		},
	}
}

func getContractBalancesInternal() GetCosmicGameContractBalancesResponseObject {
	return GetCosmicGameContractBalances500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(contractBalancesInstance),
		),
	}
}

func getCurrentBidPricesUnavailable() GetCosmicGameCurrentBidPricesResponseObject {
	return GetCosmicGameCurrentBidPrices503ApplicationProblemPlusJSONResponse{
		ServiceUnavailableApplicationProblemPlusJSONResponse: ServiceUnavailableApplicationProblemPlusJSONResponse{
			Body:    cachedLiveProblem(currentBidPricesInstance, "current bid prices"),
			Headers: ServiceUnavailableResponseHeaders{RetryAfter: contractCacheRetryAfter},
		},
	}
}

func getCurrentBidPricesInternal() GetCosmicGameCurrentBidPricesResponseObject {
	return GetCosmicGameCurrentBidPrices500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(currentBidPricesInstance),
		),
	}
}

func getCurrentSpecialWinnersUnavailable() GetCosmicGameCurrentSpecialWinnersResponseObject {
	return GetCosmicGameCurrentSpecialWinners503ApplicationProblemPlusJSONResponse{
		ServiceUnavailableApplicationProblemPlusJSONResponse: ServiceUnavailableApplicationProblemPlusJSONResponse{
			Body:    cachedLiveProblem(currentSpecialWinnersInstance, "special-winner standings"),
			Headers: ServiceUnavailableResponseHeaders{RetryAfter: specialWinnersRetryAfter},
		},
	}
}

func getCurrentSpecialWinnersInternal() GetCosmicGameCurrentSpecialWinnersResponseObject {
	return GetCosmicGameCurrentSpecialWinners500ApplicationProblemPlusJSONResponse{
		InternalErrorApplicationProblemPlusJSONResponse: InternalErrorApplicationProblemPlusJSONResponse(
			internalProblem(currentSpecialWinnersInstance),
		),
	}
}
