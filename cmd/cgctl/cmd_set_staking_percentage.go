package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

func init() {
	register(newGameSetterCmd(gameSetterSpec{
		use:       "set-staking-percentage <cosmicgame-addr> <percentage>",
		short:     "Set cosmicSignatureNftStakingTotalEthRewardAmountPercentage (owner only)",
		long:      "Set cosmicSignatureNftStakingTotalEthRewardAmountPercentage, the percentage of funds used for staking rewards.",
		section:   "STAKING REWARD PERCENTAGE CONFIG",
		action:    "SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage",
		valueName: "percentage",
		percent:   true,
		read: func(g *cgcontracts.CosmicSignatureGame, o *bind.CallOpts) (*big.Int, error) {
			return g.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(o)
		},
		write: func(g *cgcontracts.CosmicSignatureGame, o *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
			return g.SetCosmicSignatureNftStakingTotalEthRewardAmountPercentage(o, v)
		},
	}))
}
