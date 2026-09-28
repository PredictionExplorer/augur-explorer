package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

// totalNumBids returns the bid count of roundNum, preferring the consolidated
// v3.1 roundStats() getter (which replaced getTotalNumBids) and falling back
// to the pre-consolidation getter still exposed by older deployments.
func totalNumBids(
	copts *bind.CallOpts,
	game *cgcontracts.CosmicSignatureGame,
	gameV3 *cgcontracts.CosmicSignatureGameV3,
	roundNum *big.Int,
) (*big.Int, error) {
	if gameV3 != nil {
		if stats, err := gameV3.RoundStats(copts, roundNum); err == nil && stats.NumBids != nil {
			return stats.NumBids, nil
		}
	}
	return game.GetTotalNumBids(copts, roundNum)
}
