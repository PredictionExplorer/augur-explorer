// Burns CosmicSignatureToken (CST) from the sender's account
package main

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	// Usage check
	if len(os.Args) < 4 {
		cutils.PrintUsage(os.Args[0],
			"[private_key] [cosmictoken_addr] [amount]",
			"Burns CosmicSignatureToken (CST) from the sender's account",
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

	// Parse parameters
	tokenAddr := common.HexToAddress(os.Args[2])

	amount := big.NewInt(0)
	_, success := amount.SetString(os.Args[3], 10)
	if !success {
		cutils.Fatal("Invalid amount provided: %s", os.Args[3])
	}

	// Contract setup
	cstToken, err := NewCosmicSignatureToken(tokenAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicSignatureToken: %v", err)
	}

	// Get token info
	copts := cutils.CreateCallOpts()

	symbol, err := cstToken.Symbol(copts)
	if err != nil {
		symbol = "CST"
	}

	balance, err := cstToken.BalanceOf(copts, acc.Address)
	if err != nil {
		cutils.Fatal("Error getting balance: %v", err)
	}

	cutils.Section("TOKEN INFO")
	cutils.PrintKeyValue("Token Address", tokenAddr.String())
	cutils.PrintKeyValue("Token Symbol", symbol)

	cutils.Section("BURN INFO")
	cutils.PrintKeyValue("Your Balance Before", cutils.WeiToEth(balance)+" "+symbol)
	cutils.PrintKeyValue("Amount to Burn", cutils.WeiToEth(amount)+" "+symbol)
	cutils.PrintKeyValue("Balance After Burn", cutils.WeiToEth(new(big.Int).Sub(balance, amount))+" "+symbol)

	// Check balance
	if balance.Cmp(amount) < 0 {
		cutils.Fatal("Insufficient balance. Trying to burn %s %s, have %s %s",
			cutils.WeiToEth(amount), symbol, cutils.WeiToEth(balance), symbol)
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("Burn", nil, cutils.GasLimitContractCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitContractCall)

	tx, err := cstToken.Burn(txopts, amount)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
