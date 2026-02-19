// Claims staking reward by unstaking from StakingWalletCosmicSignatureNft
// Note: In the current contract version, unstaking also claims rewards
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
			"[staking_wallet_addr] [stake_action_id]",
			"Claims staking reward by unstaking (Unstake) from StakingWalletCosmicSignatureNft",
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

	stakeActionID, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing stake_action_id: %v", err)
	}

	// Contract setup
	cutils.PrintContractInfo("StakingWallet Address", stakingWalletAddr)

	stakingWallet, err := NewStakingWalletCosmicSignatureNft(stakingWalletAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate StakingWalletCosmicSignatureNft: %v", err)
	}

	// Get staking wallet balance for info
	walletBalance, err := cutils.GetBalance(net, stakingWalletAddr)
	if err != nil {
		walletBalance = nil
	}

	cutils.Section("CLAIM INFO")
	cutils.PrintKeyValue("Stake Action ID", stakeActionID)
	if walletBalance != nil {
		cutils.PrintKeyValueEth("Staking Wallet Balance", walletBalance)
	}
	cutils.PrintKeyValue("Note", "Unstaking also claims accumulated rewards")

	// Create and submit transaction
	cutils.PrintTxSubmitting("Unstake (claim rewards)", nil, cutils.GasLimitContractCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitContractCall)

	tx, err := stakingWallet.Unstake(txopts, big.NewInt(stakeActionID))
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
