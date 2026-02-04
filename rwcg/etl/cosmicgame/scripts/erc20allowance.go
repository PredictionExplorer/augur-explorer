// Gets ERC20 token allowance
package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	// Usage check
	if len(os.Args) < 4 {
		cutils.PrintUsage(os.Args[0],
			"[erc20_token_addr] [owner_addr] [spender_addr]",
			"Gets the ERC20 token allowance for a spender",
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
	tokenAddr := common.HexToAddress(os.Args[1])
	ownerAddr := common.HexToAddress(os.Args[2])
	spenderAddr := common.HexToAddress(os.Args[3])

	// Contract setup
	erc20, err := NewERC20(tokenAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate ERC20 contract: %v", err)
	}

	// Get token info
	copts := cutils.CreateCallOpts()

	symbol, err := erc20.Symbol(copts)
	if err != nil {
		symbol = "UNKNOWN"
	}

	decimals, err := erc20.Decimals(copts)
	if err != nil {
		decimals = 18
	}

	allowance, err := erc20.Allowance(copts, ownerAddr, spenderAddr)
	if err != nil {
		cutils.Fatal("Error getting allowance: %v", err)
	}

	ownerBalance, err := erc20.BalanceOf(copts, ownerAddr)
	if err != nil {
		cutils.Fatal("Error getting owner balance: %v", err)
	}

	cutils.Section("TOKEN INFO")
	cutils.PrintKeyValue("Token Address", tokenAddr.String())
	cutils.PrintKeyValue("Token Symbol", symbol)
	cutils.PrintKeyValue("Decimals", decimals)

	cutils.Section("ALLOWANCE INFO")
	cutils.PrintKeyValue("Owner", ownerAddr.String())
	cutils.PrintKeyValue("Spender", spenderAddr.String())
	cutils.PrintKeyValue("Owner Balance (raw)", ownerBalance.String())
	cutils.PrintKeyValue("Owner Balance", cutils.FormatTokenAmount(ownerBalance, decimals, symbol))
	cutils.PrintKeyValue("Allowance (raw)", allowance.String())
	cutils.PrintKeyValue("Allowance", cutils.FormatTokenAmount(allowance, decimals, symbol))

	// Check if unlimited
	maxUint256 := cutils.MaxUint256()
	if allowance.Cmp(maxUint256) == 0 {
		cutils.PrintKeyValue("Status", "UNLIMITED (MAX_UINT256)")
	} else if allowance.Cmp(ownerBalance) >= 0 {
		cutils.PrintKeyValue("Status", "Sufficient for full balance")
	} else {
		cutils.PrintKeyValue("Status", "Limited allowance")
	}
}
