// Changes ownership of an Ownable contract
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
			"[contract_addr] [new_owner_addr]",
			"Changes ownership of an Ownable contract",
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
	contractAddr := common.HexToAddress(os.Args[1])
	newOwnerAddr := common.HexToAddress(os.Args[2])

	// Contract setup
	cutils.PrintContractInfo("Contract Address", contractAddr)

	ownable, err := NewOwnable(contractAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate Ownable contract: %v", err)
	}

	// Get current owner
	copts := cutils.CreateCallOpts()

	currentOwner, err := ownable.Owner(copts)
	if err != nil {
		cutils.Fatal("Error getting current owner: %v", err)
	}

	cutils.Section("OWNERSHIP TRANSFER")
	cutils.PrintKeyValue("Current Owner", currentOwner.String())
	cutils.PrintKeyValue("New Owner", newOwnerAddr.String())

	// Check ownership
	if acc.Address != currentOwner {
		cutils.Section("WARNING")
		cutils.PrintKeyValue("Your Address", acc.Address.String())
		cutils.PrintKeyValue("Note", "You are NOT the current owner. Transaction will fail.")
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("TransferOwnership", nil, cutils.GasLimitAdminCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall)

	tx, err := ownable.TransferOwnership(txopts, newOwnerAddr)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
