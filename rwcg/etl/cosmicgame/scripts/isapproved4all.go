// Gets ERC721 isApprovedForAll status (operator level approval)
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
			"[erc721_token_addr] [owner_addr] [operator_addr]",
			"Gets ERC721 isApprovedForAll status (operator level approval)",
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
	operatorAddr := common.HexToAddress(os.Args[3])

	// Contract setup
	erc721, err := NewCosmicSignatureNft(tokenAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate ERC721 contract: %v", err)
	}

	// Get approval status
	copts := cutils.CreateCallOpts()

	isApproved, err := erc721.IsApprovedForAll(copts, ownerAddr, operatorAddr)
	if err != nil {
		cutils.Fatal("Error calling IsApprovedForAll(): %v", err)
	}

	// Get owner's token balance for context
	balance, err := erc721.BalanceOf(copts, ownerAddr)
	if err != nil {
		balance = nil
	}

	cutils.Section("TOKEN INFO")
	cutils.PrintKeyValue("Token Contract", tokenAddr.String())

	cutils.Section("APPROVAL STATUS")
	cutils.PrintKeyValue("Owner", ownerAddr.String())
	cutils.PrintKeyValue("Operator", operatorAddr.String())
	if balance != nil {
		cutils.PrintKeyValue("Owner's Token Balance", balance.String())
	}
	cutils.PrintKeyValue("Is Approved For All", isApproved)

	if isApproved {
		cutils.PrintKeyValue("Status", "APPROVED - Operator can transfer all owner's tokens")
	} else {
		cutils.PrintKeyValue("Status", "NOT APPROVED - Operator cannot transfer owner's tokens")
	}
}
