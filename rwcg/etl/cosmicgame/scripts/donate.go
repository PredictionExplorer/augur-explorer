// Makes a donation to CosmicGame contract
package main

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	cutils.ParseInfoFlag()
	// Usage check
	if len(os.Args) != 3 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicgame_contract_addr] [amount_wei]",
			"Makes an ETH donation to CosmicGame contract",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)", "PKEY_HEX": "64-char hex private key, no 0x prefix (required)"},
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
	acc, err := cutils.PrepareAccount(net, cutils.MustGetPkeyHex())
	if err != nil {
		cutils.Fatal("Account setup failed: %v", err)
	}
	cutils.PrintAccountInfo(acc)

	// Parse donation amount
	donationAmount := big.NewInt(0)
	_, success := donationAmount.SetString(os.Args[2], 10)
	if !success {
		cutils.Fatal("Invalid amount provided: %s", os.Args[2])
	}

	// Contract setup
	cosmicGameAddr := common.HexToAddress(os.Args[1])
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
	cutils.PrintKeyValueEth("Contract Balance After", new(big.Int).Add(contractBalance, donationAmount))

	// Check if account has enough balance
	if acc.Balance.Cmp(donationAmount) < 0 {
		cutils.Fatal("Insufficient balance. Need %s ETH, have %s ETH",
			cutils.WeiToEth(donationAmount), cutils.WeiToEth(acc.Balance))
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("DonateEth", donationAmount, cutils.GasLimitDonate, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, donationAmount, cutils.GasLimitDonate)

	tx, err := cosmicGame.DonateEth(txopts)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
