// Cancels offer
package main

import (
	"bytes"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/randomwalk"
)

func main() {
	cutils.ParseInfoFlag()
	cutils.GasPriceMultiplier = big.NewFloat(2.0)
	if len(os.Args) != 3 {
		cutils.PrintUsage(os.Args[0],
			"[market_addr] [offer_id]",
			"Cancels an existing buy or sell offer.",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)", "PKEY_HEX": "64-char hex private key, no 0x prefix (required)"},
		)
		os.Exit(1)
	}

	marketAddr := common.HexToAddress(os.Args[1])
	offerID, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		cutils.Fatal("Invalid offer_id: %v", err)
	}
	if err != nil {
		cutils.Fatal("Invalid offer_id: %v", err)
	}

	net, err := cutils.ConnectToRPC()
	if err != nil {
		cutils.Fatal("Network connection failed: %v", err)
	}
	cutils.PrintNetworkInfo(net)

	acc, err := cutils.PrepareAccount(net, cutils.MustGetPkeyHex())
	if err != nil {
		cutils.Fatal("Account setup failed: %v", err)
	}
	cutils.PrintAccountInfo(acc)

	marketCtrct, err := NewRWMarket(marketAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate RWMarket contract: %v", err)
	}

	var zeroAddr common.Address
	offer, err := marketCtrct.Offers(cutils.CreateCallOpts(), big.NewInt(offerID))
	if err != nil {
		cutils.Fatal("Error calling offers(offer_id=%v): %v", offerID, err)
	}
	isSellOffer := bytes.Equal(zeroAddr.Bytes(), offer.Buyer.Bytes())

	offerType := "SELL"
	if !isSellOffer {
		offerType = "BUY"
	}
	cutils.Section("CANCEL OFFER INFO")
	cutils.PrintKeyValue("Offer ID", offerID)
	cutils.PrintKeyValue("Token ID", offer.TokenId.String())
	cutils.PrintKeyValue("Type", offerType)

	gasLimit := cutils.GasLimitHighComplexity
	txopts := cutils.CreateTransactOpts(net, acc, big.NewInt(0), gasLimit)
	cutils.PrintTxSubmitting("CancelOffer", big.NewInt(0), gasLimit, cutils.AdjustGasPrice(net.GasPrice))

	var tx *types.Transaction
	if isSellOffer {
		tx, err = marketCtrct.CancelSellOffer(txopts, big.NewInt(offerID))
	} else {
		tx, err = marketCtrct.CancelBuyOffer(txopts, big.NewInt(offerID))
	}
	if !cutils.PrintTxResultAndWait(net.Client, tx, err) {
		os.Exit(1)
	}
}
