package main

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/ethtx"
)

// newAcceptOfferCmd builds the accept-offer subcommand (legacy accept_offer script).
func newAcceptOfferCmd() *cobra.Command {
	var verbose bool
	c := &cobra.Command{
		Use:   "accept-offer [market_addr] [offer_id]",
		Short: "Accept a buy or sell offer on the RandomWalk marketplace",
		Long:  "Accepts a buy or sell offer (purchases token on the marketplace).\n\n" + txEnvHelp,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			marketAddr := common.HexToAddress(args[0])
			offerID, err := parseInt64("offer_id", args[1])
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

			offer, err := market.Offers(callOpts(), big.NewInt(offerID))
			if err != nil {
				return fmt.Errorf("error calling offers(offer_id=%v): %w", offerID, err)
			}
			var zeroAddr common.Address
			isSellOffer := bytes.Equal(zeroAddr.Bytes(), offer.Buyer.Bytes())
			amount := new(big.Int).Set(offer.Price)

			offerType := "SELL"
			if !isSellOffer {
				offerType = "BUY"
			}
			s.Out.Section("ACCEPT OFFER INFO")
			s.Out.KeyValue("Offer ID", offerID)
			s.Out.KeyValue("Token ID", offer.TokenId.String())
			s.Out.KeyValue("Type", offerType)
			s.Out.KeyValueEth("Price", amount)

			gasLimit := ethtx.GasLimitHighComplexity
			value := big.NewInt(0)
			if isSellOffer {
				value.Set(amount)
			}
			txopts := s.TransactOpts(value, gasLimit)
			s.Out.TxSubmitting("Accept"+offerType+"Offer", value, gasLimit, ethtx.AdjustGasPrice(s.Net.GasPrice))

			var tx *types.Transaction
			if isSellOffer {
				tx, err = market.AcceptSellOffer(txopts, big.NewInt(offerID))
			} else {
				tx, err = market.AcceptBuyOffer(txopts, big.NewInt(offerID))
			}
			return s.FinishTx(cmd.Context(), tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newAcceptOfferCmd()) }
