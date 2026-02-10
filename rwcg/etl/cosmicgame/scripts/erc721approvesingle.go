// Approves a specific ERC721 token for transfer by an operator
package main

import (
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	// Usage check
	if len(os.Args) != 4 {
		cutils.PrintUsage(os.Args[0],
			"[erc721_contract_addr] [operator_addr] [token_id]",
			"Approves a specific ERC721 token for transfer by an operator (single token approval)",
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

	// Parse parameters
	contractAddr := common.HexToAddress(os.Args[1])
	operatorAddr := common.HexToAddress(os.Args[2])

	tokenID, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing token_id: %v", err)
	}

	// Contract setup
	erc721, err := NewCosmicSignatureNft(contractAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate ERC721 contract: %v", err)
	}

	// Get current info
	copts := cutils.CreateCallOpts()

	currentOwner, err := erc721.OwnerOf(copts, big.NewInt(tokenID))
	if err != nil {
		cutils.Fatal("Error getting token owner: %v", err)
	}

	currentApproved, err := erc721.GetApproved(copts, big.NewInt(tokenID))
	if err != nil {
		cutils.Fatal("Error getting current approval: %v", err)
	}

	cutils.Section("APPROVAL INFO")
	cutils.PrintKeyValue("Contract Address", contractAddr.String())
	cutils.PrintKeyValue("Token ID", tokenID)
	cutils.PrintKeyValue("Token Owner", currentOwner.String())
	cutils.PrintKeyValue("Current Approved", currentApproved.String())
	cutils.PrintKeyValue("New Operator", operatorAddr.String())

	// Check ownership
	if acc.Address != currentOwner {
		cutils.Section("WARNING")
		cutils.PrintKeyValue("Your Address", acc.Address.String())
		cutils.PrintKeyValue("Note", "You are NOT the token owner. Approval will fail unless you are an approved operator.")
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("Approve (single token)", nil, cutils.GasLimitERC721Approve, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitERC721Approve)

	tx, err := erc721.Approve(txopts, operatorAddr, big.NewInt(tokenID))
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
