// Gets the owner of a specific ERC721 token
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
			"Gets the owner of a specific ERC721 token",
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
	contractAddr := common.HexToAddress(os.Args[1])
	tokenID, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing token ID: %v", err)
	}

	// Contract setup
	erc721, err := NewERC721(contractAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate ERC721 contract: %v", err)
	}

	// Get token info
	copts := cutils.CreateCallOpts()

	owner, err := erc721.OwnerOf(copts, big.NewInt(tokenID))
	if err != nil {
		cutils.Fatal("Error calling OwnerOf(): %v (token may not exist)", err)
	}

	// Get owner's balance for context
	balance, err := erc721.BalanceOf(copts, owner)
	if err != nil {
		balance = nil
	}

	ownerEthBalance, err := cutils.GetBalance(net, owner)
	if err != nil {
		ownerEthBalance = nil
	}

	cutils.Section("TOKEN INFO")
	cutils.PrintKeyValue("Contract Address", contractAddr.String())
	cutils.PrintKeyValue("Token ID", tokenID)

	cutils.Section("OWNERSHIP INFO")
	cutils.PrintKeyValue("Owner Address", owner.String())
	if balance != nil {
		cutils.PrintKeyValue("Owner's Total NFTs", balance.String())
	}
	if ownerEthBalance != nil {
		cutils.PrintKeyValueEth("Owner's ETH Balance", ownerEthBalance)
	}
}
