// Revokes ERC20 token spending (sets allowance to 0)
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
	if len(os.Args) < 4 {
		cutils.PrintUsage(os.Args[0],
			"[private_key] [erc20_token_addr] [spender_addr]",
			"Revokes allowance (sets to 0) for spender to spend your ERC20 tokens",
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

	// Parse addresses
	tokenAddr := common.HexToAddress(os.Args[2])
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

	currentAllowance, err := erc20.Allowance(copts, acc.Address, spenderAddr)
	if err != nil {
		cutils.Fatal("Error getting current allowance: %v", err)
	}

	cutils.Section("TOKEN INFO")
	cutils.PrintKeyValue("Token Address", tokenAddr.String())
	cutils.PrintKeyValue("Token Symbol", symbol)
	cutils.PrintKeyValue("Spender Address", spenderAddr.String())
	cutils.PrintKeyValue("Current Allowance", currentAllowance.String())

	cutils.Section("REVOKE INFO")
	cutils.PrintKeyValue("New Allowance", "0 (revoking)")

	// Create and submit transaction
	cutils.PrintTxSubmitting("Approve (revoke to 0)", nil, cutils.GasLimitERC20Approve, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitERC20Approve)

	tx, err := erc20.Approve(txopts, spenderAddr, big.NewInt(0))
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
