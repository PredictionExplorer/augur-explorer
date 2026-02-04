// Gets ERC20 token total supply
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
			"[erc20_contract_addr]",
			"Gets ERC20 token total supply",
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
	contractAddr := common.HexToAddress(os.Args[1])

	// Contract setup
	erc20, err := NewERC20(contractAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate ERC20 contract: %v", err)
	}

	// Get token info
	copts := cutils.CreateCallOpts()

	name, err := erc20.Name(copts)
	if err != nil {
		name = "UNKNOWN"
	}

	symbol, err := erc20.Symbol(copts)
	if err != nil {
		symbol = "UNKNOWN"
	}

	decimals, err := erc20.Decimals(copts)
	if err != nil {
		decimals = 18
	}

	totalSupply, err := erc20.TotalSupply(copts)
	if err != nil {
		cutils.Fatal("Error getting total supply: %v", err)
	}

	cutils.Section("TOKEN INFO")
	cutils.PrintKeyValue("Contract Address", contractAddr.String())
	cutils.PrintKeyValue("Name", name)
	cutils.PrintKeyValue("Symbol", symbol)
	cutils.PrintKeyValue("Decimals", decimals)

	cutils.Section("SUPPLY INFO")
	cutils.PrintKeyValue("Total Supply (raw)", totalSupply.String())
	cutils.PrintKeyValue("Total Supply", cutils.FormatTokenAmount(totalSupply, decimals, symbol))
}
