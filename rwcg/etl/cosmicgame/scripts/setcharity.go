// Sets the charity address to receive charity deposits
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
	if len(os.Args) != 3 {
		cutils.PrintUsage(os.Args[0],
			"[charity_wallet_addr] [new_charity_addr]",
			"Sets the charity address to receive charity deposits",
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

	// Parse addresses
	charityWalletAddr := common.HexToAddress(os.Args[1])
	newCharityAddr := common.HexToAddress(os.Args[2])

	// Contract setup
	cutils.PrintContractInfo("CharityWallet Address", charityWalletAddr)

	charityWallet, err := NewCharityWallet(charityWalletAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CharityWallet: %v", err)
	}

	// Get current charity info
	copts := cutils.CreateCallOpts()

	currentCharity, err := charityWallet.CharityAddress(copts)
	if err != nil {
		cutils.Fatal("Error getting current charity address: %v", err)
	}

	cutils.Section("CHARITY CHANGE")
	cutils.PrintKeyValue("Current Charity Address", currentCharity.String())
	cutils.PrintKeyValue("New Charity Address", newCharityAddr.String())

	// Create and submit transaction
	cutils.PrintTxSubmitting("SetCharityAddress", nil, cutils.GasLimitAdminCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall)

	tx, err := charityWallet.SetCharityAddress(txopts, newCharityAddr)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
