// Transfers CosmicSignatureNFT to another address
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
	cutils.ParseInfoFlag()
	// Usage check
	if len(os.Args) < 4 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicsig_addr] [recipient_addr] [token_id]",
			"Transfers a CosmicSignatureNFT to another address",
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
	nftAddr := common.HexToAddress(os.Args[1])
	recipientAddr := common.HexToAddress(os.Args[2])

	tokenIDNum, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		cutils.Fatal("Invalid token_id: %s", os.Args[3])
	}
	tokenID := big.NewInt(tokenIDNum)

	// Contract setup
	nft, err := NewCosmicSignatureNft(nftAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicSignatureNft: %v", err)
	}

	// Get token info
	copts := cutils.CreateCallOpts()

	currentOwner, err := nft.OwnerOf(copts, tokenID)
	if err != nil {
		cutils.Fatal("Error getting token owner (token may not exist): %v", err)
	}

	senderBalance, err := nft.BalanceOf(copts, acc.Address)
	if err != nil {
		senderBalance = nil
	}

	cutils.Section("NFT INFO")
	cutils.PrintKeyValue("Contract Address", nftAddr.String())
	cutils.PrintKeyValue("Token ID", tokenID.String())

	cutils.Section("TRANSFER INFO")
	cutils.PrintKeyValue("Current Owner", currentOwner.String())
	cutils.PrintKeyValue("From (you)", acc.Address.String())
	cutils.PrintKeyValue("To", recipientAddr.String())
	if senderBalance != nil {
		cutils.PrintKeyValue("Your NFT Balance", senderBalance.String())
	}

	// Check ownership
	if acc.Address != currentOwner {
		cutils.Section("WARNING")
		cutils.PrintKeyValue("Note", "You are NOT the current owner. Transfer may fail unless you are approved.")
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("TransferFrom", nil, cutils.GasLimitContractCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitContractCall)

	tx, err := nft.TransferFrom(txopts, acc.Address, recipientAddr, tokenID)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
