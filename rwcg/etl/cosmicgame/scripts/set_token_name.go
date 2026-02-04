// Sets NFT name for a CosmicSignatureNft token
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
			"[private_key] [cosmicsignaturenft_contract_addr] [token_id] [name (optional)]",
			"Sets NFT name for a CosmicSignatureNft token. If name is omitted, sets to empty string.",
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

	// Parse parameters
	cosmicSigAddr := common.HexToAddress(os.Args[2])

	tokenId := big.NewInt(0)
	_, success := tokenId.SetString(os.Args[3], 10)
	if !success {
		cutils.Fatal("Invalid token_id value provided: %s", os.Args[3])
	}

	// Name is optional, defaults to empty string
	nftName := ""
	if len(os.Args) >= 5 {
		nftName = os.Args[4]
	}

	// Contract setup
	cutils.PrintContractInfo("CosmicSignatureNft Address", cosmicSigAddr)

	cosmicSig, err := NewCosmicSignatureNft(cosmicSigAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicSignatureNft: %v", err)
	}

	cutils.Section("NFT NAME CONFIG")
	cutils.PrintKeyValue("Token ID", tokenId.String())
	if nftName == "" {
		cutils.PrintKeyValue("New Name", "(empty)")
	} else {
		cutils.PrintKeyValue("New Name", nftName)
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("SetNftName", nil, cutils.GasLimitAdminCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall)

	tx, err := cosmicSig.SetNftName(txopts, tokenId, nftName)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
