// Sets mainPrizeTimeIncrementInMicroSeconds to achieve desired time increment per bid
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
			"[private_key] [cosmicgame_contract_addr] [time_increment_seconds]",
			"Sets mainPrizeTimeIncrementInMicroSeconds so that each bid extends the time until main prize by the specified seconds",
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
	desiredSeconds, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing time_increment_seconds: %v", err)
	}
	if desiredSeconds <= 0 {
		cutils.Fatal("time_increment_seconds must be positive")
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

	currentMicroseconds, err := cosmicGame.MainPrizeTimeIncrementInMicroSeconds(copts)
	if err != nil {
		cutils.Fatal("Error reading mainPrizeTimeIncrementInMicroSeconds: %v", err)
	}
	currentSeconds := new(big.Int).Div(currentMicroseconds, big.NewInt(1000000))

	owner, err := cosmicGame.Owner(copts)
	if err != nil {
		cutils.Fatal("Error getting contract owner: %v", err)
	}

	// Calculate new value: desired_seconds * 1,000,000
	newMicroseconds := new(big.Int).Mul(big.NewInt(desiredSeconds), big.NewInt(1000000))

	cutils.Section("CURRENT STATE")
	cutils.PrintKeyValue("Contract Owner", owner.String())
	cutils.PrintKeyValue("Current Microseconds", currentMicroseconds.String())
	cutils.PrintKeyValueDuration("Current Time Increment", currentSeconds.Int64())

	cutils.Section("NEW VALUES")
	cutils.PrintKeyValue("New Microseconds", newMicroseconds.String())
	cutils.PrintKeyValueDuration("New Time Increment", desiredSeconds)
	cutils.PrintKeyValue("Formula", "timeIncrement (seconds) = microseconds / 1,000,000")

	// Check ownership
	if acc.Address != owner {
		cutils.Section("WARNING")
		cutils.PrintKeyValue("Your Address", acc.Address.String())
		cutils.PrintKeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("SetMainPrizeTimeIncrementInMicroSeconds", nil, cutils.GasLimitAdminCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall)

	tx, err := cosmicGame.SetMainPrizeTimeIncrementInMicroSeconds(txopts, newMicroseconds)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
