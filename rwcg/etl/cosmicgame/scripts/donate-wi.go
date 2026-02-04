// Makes a donation with info (JSON metadata) to CosmicGame contract
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
	if len(os.Args) != 5 {
		cutils.PrintUsage(os.Args[0],
			"[private_key] [cosmicgame_contract_addr] [amount_wei] [json_data]",
			"Makes an ETH donation with info (JSON metadata) to CosmicGame contract",
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
	cosmicGameAddr := common.HexToAddress(os.Args[2])

	donationAmount := big.NewInt(0)
	_, success := donationAmount.SetString(os.Args[3], 10)
	if !success {
		cutils.Fatal("Invalid amount provided: %s", os.Args[3])
	}

	jsonData := os.Args[4]

	// Contract setup
	cutils.PrintContractInfo("CosmicGame Address", cosmicGameAddr)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}

	// Get current contract balance
	contractBalance, err := cutils.GetBalance(net, cosmicGameAddr)
	if err != nil {
		cutils.Fatal("Error getting contract balance: %v", err)
	}

	cutils.Section("DONATION INFO")
	cutils.PrintKeyValueEth("Contract Current Balance", contractBalance)
	cutils.PrintKeyValueEth("Donation Amount", donationAmount)
	cutils.PrintKeyValue("JSON Data", jsonData)

	// Check if account has enough balance
	if acc.Balance.Cmp(donationAmount) < 0 {
		cutils.Fatal("Insufficient balance. Need %s ETH, have %s ETH",
			cutils.WeiToEth(donationAmount), cutils.WeiToEth(acc.Balance))
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("DonateWithInfo", donationAmount, cutils.GasLimitDonate, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, donationAmount, cutils.GasLimitDonate)

	tx, err := cosmicGame.DonateEthWithInfo(txopts, jsonData)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
