package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
	register(newGameSetterCmd(gameSetterSpec{
		use:       "set-main-prize-percentage <cosmicgame-addr> <percentage>",
		short:     "Set mainEthPrizeAmountPercentage (owner only)",
		long:      "Set mainEthPrizeAmountPercentage, the percentage of the contract balance paid as the main prize.",
		section:   "MAIN PRIZE PERCENTAGE CONFIG",
		action:    "SetMainEthPrizeAmountPercentage",
		valueName: "percentage",
		percent:   true,
		read: func(g *cgcontracts.CosmicSignatureGame, o *bind.CallOpts) (*big.Int, error) {
			return g.MainEthPrizeAmountPercentage(o)
		},
		write: func(g *cgcontracts.CosmicSignatureGame, o *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
			return g.SetMainEthPrizeAmountPercentage(o, v)
		},
	}))
}
