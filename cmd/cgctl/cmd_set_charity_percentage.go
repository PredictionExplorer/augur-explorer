package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

// charityPercentageSpec drives the set-charity-percentage subcommand.
var charityPercentageSpec = gameSetterSpec{
	use:       "set-charity-percentage <cosmicgame-addr> <percentage>",
	short:     "Set charityEthDonationAmountPercentage (owner only)",
	long:      "Set charityEthDonationAmountPercentage, the percentage of funds donated to charity each round.",
	section:   "CHARITY DONATION PERCENTAGE CONFIG",
	action:    "SetCharityEthDonationAmountPercentage",
	valueName: "percentage",
	percent:   true,
	read: func(g *cgcontracts.CosmicSignatureGame, o *bind.CallOpts) (*big.Int, error) {
		return g.CharityEthDonationAmountPercentage(o)
	},
	write: func(g *cgcontracts.CosmicSignatureGame, o *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
		return g.SetCharityEthDonationAmountPercentage(o, v)
	},
}

func init() { register(newGameSetterCmd(charityPercentageSpec)) }
