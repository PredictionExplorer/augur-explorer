package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
)

// newNewOfferCmd builds the new-offer subcommand (legacy new_offer script).
func newNewOfferCmd() *cobra.Command {
	var verbose bool
	c := &cobra.Command{
		Use:   "new-offer [BUY|SELL] [market_addr] [nft_addr] [token_id] [price_wei]",
		Short: "Create a buy or sell offer on the RandomWalk marketplace",
		Long:  "Creates a BUY or SELL offer on the RandomWalk marketplace.\n\n" + txEnvHelp,
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			method := args[0]
			if method != "BUY" && method != "SELL" {
				return fmt.Errorf("invalid operation: must be BUY or SELL")
			}
			marketAddr := common.HexToAddress(args[1])
			nftAddr := common.HexToAddress(args[2])
			tokenID, err := parseInt64("token_id", args[3])
			if err != nil {
				return err
			}
			amount, err := parseBigInt("price_wei", args[4])
			if err != nil {
				return err
			}

			s, err := newTxSession(verbose)
			if err != nil {
				return err
			}
			market, err := rwcontracts.NewRWMarket(marketAddr, s.net.client)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWMarket contract: %w", err)
			}

			s.out.section("OFFER INFO")
			s.out.keyValue("Type", method)
			s.out.keyValue("Market", marketAddr.String())
			s.out.keyValue("NFT", nftAddr.String())
			s.out.keyValue("Token ID", tokenID)
			s.out.keyValueEth("Price", amount)

			gasLimit := gasLimitHighComplexity
			value := big.NewInt(0)
			if method == "BUY" {
				value.Set(amount)
			}
			txopts := transactOpts(s.net, s.acc, value, gasLimit)
			s.out.txSubmitting("Make"+method+"Offer", value, gasLimit, adjustGasPrice(s.net.gasPrice))

			var tx *types.Transaction
			if method == "BUY" {
				tx, err = market.MakeBuyOffer(txopts, nftAddr, big.NewInt(tokenID))
			} else {
				tx, err = market.MakeSellOffer(txopts, nftAddr, big.NewInt(tokenID), amount)
			}
			return s.finishTx(tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newNewOfferCmd()) }
