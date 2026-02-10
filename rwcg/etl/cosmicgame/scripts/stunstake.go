// Unstakes a token from StakingWalletCosmicSignatureNft
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
	if len(os.Args) != 3 {
		cutils.PrintUsage(os.Args[0],
			"[staking_wallet_addr] [action_id]",
			"Unstakes a token from StakingWalletCosmicSignatureNft",
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
	stakingWalletAddr := common.HexToAddress(os.Args[1])

	actionID, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing action_id: %v", err)
	}

	// Contract setup
	cutils.PrintContractInfo("StakingWallet Address", stakingWalletAddr)

	stakingWallet, err := NewStakingWalletCosmicSignatureNft(stakingWalletAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate StakingWalletCosmicSignatureNft: %v", err)
	}

	cutils.Section("UNSTAKE INFO")
	cutils.PrintKeyValue("Action ID", actionID)

	// Create and submit transaction
	cutils.PrintTxSubmitting("Unstake", nil, cutils.GasLimitContractCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitContractCall)

	tx, err := stakingWallet.Unstake(txopts, big.NewInt(actionID))
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
