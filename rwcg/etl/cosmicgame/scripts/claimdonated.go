// Claims a donated NFT from the PrizesWallet
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
	if len(os.Args) != 3 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicgame_contract_addr] [donated_nft_index]",
			"Claims a donated NFT from the PrizesWallet",
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
	cosmicGameAddr := common.HexToAddress(os.Args[1])

	nftIndex, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing donated_nft_index: %v", err)
	}

	// Get PrizesWallet address from CosmicGame
	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}

	copts := cutils.CreateCallOpts()

	prizesWalletAddr, err := cosmicGame.PrizesWallet(copts)
	if err != nil {
		cutils.Fatal("Failed to get PrizesWallet address: %v", err)
	}

	prizesWallet, err := NewPrizesWallet(prizesWalletAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate PrizesWallet: %v", err)
	}

	// Get donation info
	numDonatedNfts, err := prizesWallet.NextDonatedNftIndex(copts)
	if err != nil {
		cutils.Fatal("Error getting next donated NFT index: %v", err)
	}

	cutils.Section("PRIZES WALLET INFO")
	cutils.PrintKeyValue("CosmicGame Address", cosmicGameAddr.String())
	cutils.PrintKeyValue("PrizesWallet Address", prizesWalletAddr.String())
	cutils.PrintKeyValue("Total Donated NFTs", numDonatedNfts.String())
	cutils.PrintKeyValue("Claiming NFT Index", nftIndex)

	// Validate index
	if nftIndex < 0 || nftIndex >= numDonatedNfts.Int64() {
		cutils.Fatal("Invalid NFT index %d. Valid range: 0 to %d", nftIndex, numDonatedNfts.Int64()-1)
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("ClaimDonatedNft", nil, cutils.GasLimitContractCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitContractCall)

	tx, err := prizesWallet.ClaimDonatedNft(txopts, big.NewInt(nftIndex))
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
