// Sends a new offer to the marketplace
package main

import (
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
	if len(os.Args) != 6 {
		cutils.PrintUsage(os.Args[0],
			"[BUY|SELL] [market_addr] [nft_addr] [token_id] [price_wei]",
			"Creates a BUY or SELL offer on the RandomWalk marketplace.",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)", "PKEY_HEX": "64-char hex private key, no 0x prefix (required)"},
		)
		os.Exit(1)
	}

	method := os.Args[1]
	if method != "BUY" && method != "SELL" {
		cutils.Fatal("Invalid operation: must be BUY or SELL")
	}
	marketAddr := common.HexToAddress(os.Args[2])
	nftAddr := common.HexToAddress(os.Args[3])
	tokenID, err := strconv.ParseInt(os.Args[4], 10, 64)
	if err != nil {
		cutils.Fatal("Invalid token_id: %v", err)
	}
	amount := big.NewInt(0)
	if _, ok := amount.SetString(os.Args[5], 10); !ok {
		cutils.Fatal("Invalid price_wei: %s", os.Args[5])
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

	cutils.Section("OFFER INFO")
	cutils.PrintKeyValue("Type", method)
	cutils.PrintKeyValue("Market", marketAddr.String())
	cutils.PrintKeyValue("NFT", nftAddr.String())
	cutils.PrintKeyValue("Token ID", tokenID)
	cutils.PrintKeyValueEth("Price", amount)

	gasLimit := cutils.GasLimitHighComplexity
	value := big.NewInt(0)
	if method == "BUY" {
		value.Set(amount)
	}
	txopts := cutils.CreateTransactOpts(net, acc, value, gasLimit)
	cutils.PrintTxSubmitting("Make"+method+"Offer", value, gasLimit, cutils.AdjustGasPrice(net.GasPrice))

	var tx *types.Transaction
	if method == "BUY" {
		tx, err = marketCtrct.MakeBuyOffer(txopts, nftAddr, big.NewInt(tokenID))
	} else {
		tx, err = marketCtrct.MakeSellOffer(txopts, nftAddr, big.NewInt(tokenID), amount)
	}
	if !cutils.PrintTxResultAndWait(net.Client, tx, err) {
		os.Exit(1)
	}
}
