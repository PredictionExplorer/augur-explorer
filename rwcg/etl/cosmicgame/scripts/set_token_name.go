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
	if len(os.Args) < 3 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicsignaturenft_contract_addr] [token_id] [name (optional)]",
			"Sets NFT name for a CosmicSignatureNft token. If name is omitted, sets to empty string.",
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
	cosmicSigAddr := common.HexToAddress(os.Args[1])

	tokenId := big.NewInt(0)
	_, success := tokenId.SetString(os.Args[2], 10)
	if !success {
		cutils.Fatal("Invalid token_id value provided: %s", os.Args[2])
	}

	// Name is optional, defaults to empty string
	nftName := ""
	if len(os.Args) >= 4 {
		nftName = os.Args[3]
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
