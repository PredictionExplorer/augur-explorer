// Stakes a RandomWalk NFT in the StakingWalletRandomWalkNft
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
			"[private_key] [staking_wallet_addr] [token_id]",
			"Stakes a RandomWalk NFT in the StakingWalletRandomWalkNft",
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
	stakingWalletAddr := common.HexToAddress(os.Args[2])

	tokenID, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing token_id: %v", err)
	}

	// Contract setup
	cutils.PrintContractInfo("StakingWallet Address", stakingWalletAddr)

	stakingWallet, err := NewStakingWalletRandomWalkNft(stakingWalletAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate StakingWalletRandomWalkNft: %v", err)
	}

	cutils.Section("STAKE INFO")
	cutils.PrintKeyValue("Token ID", tokenID)
	cutils.PrintKeyValue("Note", "Make sure the token is approved for transfer to the staking wallet")

	// Create and submit transaction
	cutils.PrintTxSubmitting("Stake (RandomWalk)", nil, cutils.GasLimitContractCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitContractCall)

	tx, err := stakingWallet.Stake(txopts, big.NewInt(tokenID))
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
