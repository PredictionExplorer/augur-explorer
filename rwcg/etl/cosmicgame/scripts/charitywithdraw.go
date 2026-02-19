// Withdraws from charity wallet to the designated charity address
package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	cutils.ParseInfoFlag()
	// Usage check
	if len(os.Args) != 2 {
		cutils.PrintUsage(os.Args[0],
			"[charity_wallet_addr]",
			"Withdraws funds from CharityWallet to the designated charity address",
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

	// Contract setup
	charityWalletAddr := common.HexToAddress(os.Args[1])
	cutils.PrintContractInfo("CharityWallet Address", charityWalletAddr)

	charityWallet, err := NewCharityWallet(charityWalletAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CharityWallet: %v", err)
	}

	// Get charity info
	copts := cutils.CreateCallOpts()

	charityAddr, err := charityWallet.CharityAddress(copts)
	if err != nil {
		cutils.Fatal("Error getting charity address: %v", err)
	}

	walletBalance, err := cutils.GetBalance(net, charityWalletAddr)
	if err != nil {
		cutils.Fatal("Error getting wallet balance: %v", err)
	}

	cutils.Section("CHARITY INFO")
	cutils.PrintKeyValue("Charity Wallet", charityWalletAddr.String())
	cutils.PrintKeyValue("Charity Recipient", charityAddr.String())
	cutils.PrintKeyValueEth("Wallet Balance", walletBalance)

	// Create and submit transaction
	cutils.PrintTxSubmitting("Send (withdraw to charity)", nil, cutils.GasLimitContractCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitContractCall)

	// Get wallet balance to send all
	walletBalanceToSend, err := cutils.GetBalance(net, charityWalletAddr)
	if err != nil {
		cutils.Fatal("Error getting wallet balance for send: %v", err)
	}

	tx, err := charityWallet.Send(txopts, walletBalanceToSend)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
