// Sets the delay duration before round activation
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
			"[private_key] [cosmicgame_contract_addr] [seconds]",
			"Sets the delay duration before next round activation",
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

	// Parse seconds parameter
	seconds, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing seconds parameter: %v", err)
	}

	// Contract setup
	cosmicGameAddr := common.HexToAddress(os.Args[2])
	cutils.PrintContractInfo("CosmicGame Address", cosmicGameAddr)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}

	// Get current value
	copts := cutils.CreateCallOpts()

	currentDelay, err := cosmicGame.DelayDurationBeforeRoundActivation(copts)
	if err != nil {
		cutils.Fatal("Error getting current delay: %v", err)
	}

	owner, err := cosmicGame.Owner(copts)
	if err != nil {
		cutils.Fatal("Error getting contract owner: %v", err)
	}

	cutils.Section("CURRENT STATE")
	cutils.PrintKeyValue("Contract Owner", owner.String())
	cutils.PrintKeyValueDuration("Current Delay", currentDelay.Int64())
	cutils.PrintKeyValueDuration("New Delay", seconds)

	// Check ownership
	if acc.Address != owner {
		cutils.Section("WARNING")
		cutils.PrintKeyValue("Your Address", acc.Address.String())
		cutils.PrintKeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("SetDelayDurationBeforeRoundActivation", nil, cutils.GasLimitAdminCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall)

	tx, err := cosmicGame.SetDelayDurationBeforeRoundActivation(txopts, big.NewInt(seconds))
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
