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

// newCancelOfferCmd builds the cancel-offer subcommand (legacy cancel_offer script).
func newCancelOfferCmd() *cobra.Command {
	var verbose bool
	c := &cobra.Command{
		Use:   "cancel-offer [market_addr] [offer_id]",
		Short: "Cancel an offer on the RandomWalk marketplace",
		Long:  "Cancels an existing buy or sell offer.\n\n" + txEnvHelp,
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

			offerType := "SELL"
			if !isSellOffer {
				offerType = "BUY"
			}
			s.out.section("CANCEL OFFER INFO")
			s.out.keyValue("Offer ID", offerID)
			s.out.keyValue("Token ID", offer.TokenId.String())
			s.out.keyValue("Type", offerType)

			gasLimit := gasLimitHighComplexity
			txopts := transactOpts(s.net, s.acc, big.NewInt(0), gasLimit)
			s.out.txSubmitting("CancelOffer", big.NewInt(0), gasLimit, adjustGasPrice(s.net.gasPrice))

			var tx *types.Transaction
			if isSellOffer {
				tx, err = market.CancelSellOffer(txopts, big.NewInt(offerID))
			} else {
				tx, err = market.CancelBuyOffer(txopts, big.NewInt(offerID))
			}
			return s.finishTx(tx, err)
		},
	}
	addInfoFlag(c, &verbose)
	return c
}

func init() { register(newCancelOfferCmd()) }
