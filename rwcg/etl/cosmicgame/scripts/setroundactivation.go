// Sets the round activation time
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
			"[private_key] [cosmicgame_contract_addr] [timestamp]",
			"Sets the round activation time (Unix timestamp)",
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

	// Parse timestamp parameter
	timestamp, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing timestamp parameter: %v", err)
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

	currentActivation, err := cosmicGame.RoundActivationTime(copts)
	if err != nil {
		cutils.Fatal("Error getting current activation time: %v", err)
	}

	owner, err := cosmicGame.Owner(copts)
	if err != nil {
		cutils.Fatal("Error getting contract owner: %v", err)
	}

	secsUntilCurrent := currentActivation.Int64() - int64(net.BlockTime)
	secsUntilNew := timestamp - int64(net.BlockTime)

	cutils.Section("CURRENT STATE")
	cutils.PrintKeyValue("Contract Owner", owner.String())
	cutils.PrintKeyValue("Current Block Time", net.BlockTime)
	cutils.PrintKeyValue("Current Activation Time", currentActivation.String())
	cutils.PrintKeyValueDuration("Time Until Current Activation", secsUntilCurrent)

	cutils.Section("NEW VALUES")
	cutils.PrintKeyValue("New Activation Time", timestamp)
	cutils.PrintKeyValueDuration("Time Until New Activation", secsUntilNew)

	// Check ownership
	if acc.Address != owner {
		cutils.Section("WARNING")
		cutils.PrintKeyValue("Your Address", acc.Address.String())
		cutils.PrintKeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("SetRoundActivationTime", nil, cutils.GasLimitAdminCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall)

	tx, err := cosmicGame.SetRoundActivationTime(txopts, big.NewInt(timestamp))
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
