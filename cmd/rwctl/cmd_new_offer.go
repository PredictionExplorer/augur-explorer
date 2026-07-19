package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
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
				return errors.New("invalid operation: must be BUY or SELL")
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

			s, err := newTxSession(cmd, verbose)
			if err != nil {
				return err
			}
			market, err := rwcontracts.NewRWMarket(marketAddr, s.Net.Client)
			if err != nil {
				return fmt.Errorf("failed to instantiate RWMarket contract: %w", err)
			}

			s.Out.Section("OFFER INFO")
			s.Out.KeyValue("Type", method)
			s.Out.KeyValue("Market", marketAddr.String())
			s.Out.KeyValue("NFT", nftAddr.String())
			s.Out.KeyValue("Token ID", tokenID)
			s.Out.KeyValueEth("Price", amount)

			gasLimit := ethtx.GasLimitHighComplexity
			value := big.NewInt(0)
			if method == "BUY" {
				value.Set(amount)
			}
			txopts := s.TransactOpts(value, gasLimit)
			s.Out.TxSubmitting("Make"+method+"Offer", value, gasLimit, ethtx.AdjustGasPrice(s.Net.GasPrice))

			var tx *types.Transaction
			if method == "BUY" {
				tx, err = market.MakeBuyOffer(txopts, nftAddr, big.NewInt(tokenID))
			} else {
				tx, err = market.MakeSellOffer(txopts, nftAddr, big.NewInt(tokenID), amount)
			}
			return s.FinishTx(cmd.Context(), tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newNewOfferCmd()) }
