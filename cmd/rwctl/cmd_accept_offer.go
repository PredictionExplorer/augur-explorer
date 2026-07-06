package main

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	rwcontracts "github.com/PredictionExplorer/augur-explorer/contracts/randomwalk"
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

			s, err := newTxSession(verbose)
			if err != nil {
				return err
			}
			market, err := rwcontracts.NewRWMarket(marketAddr, s.net.client)
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
			s.out.section("ACCEPT OFFER INFO")
			s.out.keyValue("Offer ID", offerID)
			s.out.keyValue("Token ID", offer.TokenId.String())
			s.out.keyValue("Type", offerType)
			s.out.keyValueEth("Price", amount)

			gasLimit := gasLimitHighComplexity
			value := big.NewInt(0)
			if isSellOffer {
				value.Set(amount)
			}
			txopts := transactOpts(s.net, s.acc, value, gasLimit)
			s.out.txSubmitting("Accept"+offerType+"Offer", value, gasLimit, adjustGasPrice(s.net.gasPrice))

			var tx *types.Transaction
			if isSellOffer {
				tx, err = market.AcceptSellOffer(txopts, big.NewInt(offerID))
			} else {
				tx, err = market.AcceptBuyOffer(txopts, big.NewInt(offerID))
			}
			return s.finishTx(tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newAcceptOfferCmd()) }
