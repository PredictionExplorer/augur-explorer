// Gets ETH balance of an address
package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"

	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	// Usage check
	if len(os.Args) < 2 {
		cutils.PrintUsage(os.Args[0],
			"[address]",
			"Gets ETH balance of an address",
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

	// Parse address
	userAddr := common.HexToAddress(os.Args[1])

	// Get balance
	balance, err := cutils.GetBalance(net, userAddr)
	if err != nil {
		cutils.Fatal("Error getting balance: %v", err)
	}

	cutils.Section("BALANCE INFO")
	cutils.PrintKeyValue("Address", userAddr.String())
	cutils.PrintKeyValue("Balance (wei)", balance.String())
	cutils.PrintKeyValueEth("Balance", balance)
}
