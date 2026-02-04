// Gets the owner of an Ownable contract
package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	// Usage check
	if len(os.Args) < 2 {
		cutils.PrintUsage(os.Args[0],
			"[contract_addr]",
			"Gets the owner of an Ownable contract",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)"},
		)
		os.Exit(1)
	}

	// Connect to network
	net, err := cutils.ConnectToRPC()
	if err != nil {
		cutils.Fatal("Network connection failed: %v", err)
	}
	cutils.PrintNetworkInfo(net)

	// Contract setup
	contractAddr := common.HexToAddress(os.Args[1])
	cutils.PrintContractInfo("Contract Address", contractAddr)

	ownable, err := NewOwnable(contractAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate Ownable contract: %v", err)
	}

	// Get owner
	copts := cutils.CreateCallOpts()

	owner, err := ownable.Owner(copts)
	if err != nil {
		cutils.Fatal("Error calling Owner(): %v", err)
	}

	// Get owner's balance for additional info
	ownerBalance, err := cutils.GetBalance(net, owner)
	if err != nil {
		ownerBalance = nil
	}

	cutils.Section("OWNERSHIP INFO")
	cutils.PrintKeyValue("Owner Address", owner.String())
	if ownerBalance != nil {
		cutils.PrintKeyValueEth("Owner Balance", ownerBalance)
	}
}
