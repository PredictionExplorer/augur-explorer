// Gets ERC721 single token approval status
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
	if len(os.Args) < 3 {
		cutils.PrintUsage(os.Args[0],
			"[erc721_contract_addr] [token_id]",
			"Gets ERC721 approved status (single token level approval)",
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

	// Parse parameters
	erc721Addr := common.HexToAddress(os.Args[1])
	tokenID, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing token ID: %v", err)
	}

	// Contract setup
	erc721, err := NewCosmicSignatureNft(erc721Addr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate ERC721 contract: %v", err)
	}

	// Get approval info
	copts := cutils.CreateCallOpts()

	operator, err := erc721.GetApproved(copts, big.NewInt(tokenID))
	if err != nil {
		cutils.Fatal("Error at GetApproved(): %v", err)
	}

	owner, err := erc721.OwnerOf(copts, big.NewInt(tokenID))
	if err != nil {
		cutils.Fatal("Error at OwnerOf(): %v", err)
	}

	cutils.Section("TOKEN INFO")
	cutils.PrintKeyValue("Contract Address", erc721Addr.String())
	cutils.PrintKeyValue("Token ID", tokenID)

	cutils.Section("APPROVAL STATUS")
	cutils.PrintKeyValue("Token Owner", owner.String())
	cutils.PrintKeyValue("Approved Operator", operator.String())

	// Check if approved to zero address (no approval)
	zeroAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")
	if operator == zeroAddr {
		cutils.PrintKeyValue("Status", "NOT APPROVED - No operator set for this token")
	} else {
		cutils.PrintKeyValue("Status", "APPROVED - Operator can transfer this token")
	}
}
