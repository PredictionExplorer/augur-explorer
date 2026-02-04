// Revokes ERC721 approval for all tokens (operator level disapproval)
package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	// Usage check
	if len(os.Args) != 4 {
		cutils.PrintUsage(os.Args[0],
			"[private_key] [erc721_contract_addr] [operator_addr]",
			"Revokes ERC721 approval for all tokens (setApprovalForAll to false)",
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
	contractAddr := common.HexToAddress(os.Args[2])
	operatorAddr := common.HexToAddress(os.Args[3])

	// Contract setup
	erc721, err := NewCosmicSignatureNft(contractAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate ERC721 contract: %v", err)
	}

	// Get current approval status
	copts := cutils.CreateCallOpts()

	currentApproval, err := erc721.IsApprovedForAll(copts, acc.Address, operatorAddr)
	if err != nil {
		cutils.Fatal("Error getting current approval status: %v", err)
	}

	cutils.Section("DISAPPROVAL INFO")
	cutils.PrintKeyValue("Contract Address", contractAddr.String())
	cutils.PrintKeyValue("Owner (you)", acc.Address.String())
	cutils.PrintKeyValue("Operator", operatorAddr.String())
	cutils.PrintKeyValue("Current Approval Status", currentApproval)
	cutils.PrintKeyValue("New Approval Status", false)

	// Create and submit transaction
	cutils.PrintTxSubmitting("SetApprovalForAll(false)", nil, cutils.GasLimitERC721Approve, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitERC721Approve)

	tx, err := erc721.SetApprovalForAll(txopts, operatorAddr, false)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
