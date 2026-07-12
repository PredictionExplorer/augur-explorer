package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

// numRaffleWinnersSpec drives the set-num-raffle-winners subcommand.
var numRaffleWinnersSpec = gameSetterSpec{
	use:       "set-num-raffle-winners <cosmicgame-addr> <num-winners>",
	short:     "Set numRaffleEthPrizesForBidders (owner only)",
	long:      "Set numRaffleEthPrizesForBidders, the number of ETH raffle winners per round.",
	section:   "RAFFLE ETH WINNERS CONFIG",
	action:    "SetNumRaffleEthPrizesForBidders",
	valueName: "number",
	read: func(g *cgcontracts.CosmicSignatureGame, o *bind.CallOpts) (*big.Int, error) {
		return g.NumRaffleEthPrizesForBidders(o)
	},
	write: func(g *cgcontracts.CosmicSignatureGame, o *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
		return g.SetNumRaffleEthPrizesForBidders(o, v)
	},
}

func init() { register(newGameSetterCmd(numRaffleWinnersSpec)) }
