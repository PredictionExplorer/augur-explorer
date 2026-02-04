// Makes a bid using CST tokens at CosmicGame current round
package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	// Usage check
	if len(os.Args) != 3 {
		cutils.PrintUsage(os.Args[0],
			"[private_key] [cosmicgame_contract_addr]",
			"Makes a bid using CST tokens at CosmicGame current round",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)"},
		)
		os.Exit(1)
	}

	// Connect to network (chainID and gasPrice fetched from network)
	net, err := cutils.ConnectToRPC()
	if err != nil {
		cutils.Fatal("Network connection failed: %v", err)
	}
	cutils.PrintNetworkInfo(net)

	// Prepare account
	acc, err := cutils.PrepareAccount(net, os.Args[1])
	if err != nil {
		cutils.Fatal("Account setup failed: %v", err)
	}
	cutils.PrintAccountInfo(acc)

	// Contract setup
	cosmicGameAddr := common.HexToAddress(os.Args[2])
	cutils.PrintContractInfo("CosmicGame Address", cosmicGameAddr)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}

	// Get CST bid price
	copts := cutils.CreateCallOpts()

	roundNum, err := cosmicGame.RoundNum(copts)
	if err != nil {
		cutils.Fatal("Error getting round number: %v", err)
	}

	cstPrice, err := cosmicGame.GetNextCstBidPrice(copts)
	if err != nil {
		cutils.Fatal("Error getting CST bid price: %v", err)
	}

	// Get CST token address and check balance
	tokenAddr, err := cosmicGame.Token(copts)
	if err != nil {
		cutils.Fatal("Error getting token address: %v", err)
	}

	cstToken, err := NewERC20(tokenAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CST token: %v", err)
	}

	cstBalance, err := cstToken.BalanceOf(copts, acc.Address)
	if err != nil {
		cutils.Fatal("Error getting CST balance: %v", err)
	}

	cutils.Section("ROUND INFO")
	cutils.PrintKeyValue("Round Number", roundNum.String())
	cutils.PrintKeyValue("CST Token Address", tokenAddr.String())
	cutils.PrintKeyValue("Your CST Balance", cutils.WeiToEth(cstBalance)+" CST")
	cutils.PrintKeyValue("CST Bid Price", cutils.WeiToEth(cstPrice)+" CST")

	// Check if account has enough CST
	if cstBalance.Cmp(cstPrice) < 0 {
		cutils.Fatal("Insufficient CST balance. Need %s CST, have %s CST",
			cutils.WeiToEth(cstPrice), cutils.WeiToEth(cstBalance))
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("BidWithCst", nil, cutils.GasLimitBid, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitBid)

	tx, err := cosmicGame.BidWithCst(txopts, cstPrice, "")
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
