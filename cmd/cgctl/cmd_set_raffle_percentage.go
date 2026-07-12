package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

// rafflePercentageSpec drives the set-raffle-percentage subcommand.
var rafflePercentageSpec = gameSetterSpec{
	use:       "set-raffle-percentage <cosmicgame-addr> <percentage>",
	short:     "Set raffleTotalEthPrizeAmountForBiddersPercentage (owner only)",
	long:      "Set raffleTotalEthPrizeAmountForBiddersPercentage, the percentage of funds used for raffle ETH prizes.",
	section:   "RAFFLE ETH PRIZE PERCENTAGE CONFIG",
	action:    "SetRaffleTotalEthPrizeAmountForBiddersPercentage",
	valueName: "percentage",
	percent:   true,
	read: func(g *cgcontracts.CosmicSignatureGame, o *bind.CallOpts) (*big.Int, error) {
		return g.RaffleTotalEthPrizeAmountForBiddersPercentage(o)
	},
	write: func(g *cgcontracts.CosmicSignatureGame, o *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
		return g.SetRaffleTotalEthPrizeAmountForBiddersPercentage(o, v)
	},
}

func init() { register(newGameSetterCmd(rafflePercentageSpec)) }
