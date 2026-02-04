// Gets ERC20 token balance for a user
package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	// Usage check
	if len(os.Args) < 3 {
		cutils.PrintUsage(os.Args[0],
			"[erc20_contract_addr] [user_addr]",
			"Shows ERC20 token balance for a user address",
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

	// Parse addresses
	contractAddr := common.HexToAddress(os.Args[1])
	userAddr := common.HexToAddress(os.Args[2])

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

	balance, err := erc20.BalanceOf(copts, userAddr)
	if err != nil {
		cutils.Fatal("Error getting balance: %v", err)
	}

	// Get user's ETH balance for context
	ethBalance, err := cutils.GetBalance(net, userAddr)
	if err != nil {
		ethBalance = nil
	}

	cutils.Section("TOKEN INFO")
	cutils.PrintKeyValue("Contract Address", contractAddr.String())
	cutils.PrintKeyValue("Name", name)
	cutils.PrintKeyValue("Symbol", symbol)
	cutils.PrintKeyValue("Decimals", decimals)
	cutils.PrintKeyValue("Total Supply (raw)", totalSupply.String())
	cutils.PrintKeyValue("Total Supply", cutils.FormatTokenAmount(totalSupply, decimals, symbol))

	cutils.Section("USER BALANCE")
	cutils.PrintKeyValue("User Address", userAddr.String())
	cutils.PrintKeyValue("Balance (raw)", balance.String())
	cutils.PrintKeyValue("Balance", cutils.FormatTokenAmount(balance, decimals, symbol))
	if ethBalance != nil {
		cutils.PrintKeyValueEth("ETH Balance", ethBalance)
	}
}
