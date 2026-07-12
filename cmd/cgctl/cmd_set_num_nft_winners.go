package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	cgcontracts "github.com/PredictionExplorer/augur-explorer/contracts/cosmicgame"
)

// numNftWinnersSpec drives the set-num-nft-winners subcommand.
var numNftWinnersSpec = gameSetterSpec{
	use:       "set-num-nft-winners <cosmicgame-addr> <num-winners>",
	short:     "Set numRaffleCosmicSignatureNftsForBidders (owner only)",
	long:      "Set numRaffleCosmicSignatureNftsForBidders, the number of raffle NFT winners per round.",
	section:   "RAFFLE NFT WINNERS CONFIG",
	action:    "SetNumRaffleCosmicSignatureNftsForBidders",
	valueName: "number",
	read: func(g *cgcontracts.CosmicSignatureGame, o *bind.CallOpts) (*big.Int, error) {
		return g.NumRaffleCosmicSignatureNftsForBidders(o)
	},
	write: func(g *cgcontracts.CosmicSignatureGame, o *bind.TransactOpts, v *big.Int) (*types.Transaction, error) {
		return g.SetNumRaffleCosmicSignatureNftsForBidders(o, v)
	},
}

func init() { register(newGameSetterCmd(numNftWinnersSpec)) }
