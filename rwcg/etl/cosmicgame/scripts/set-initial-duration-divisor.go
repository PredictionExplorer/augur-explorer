// Sets the initial duration until main prize divisor
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
			"[cosmicgame_contract_addr] [divisor]",
			"Sets the initial duration until main prize divisor (e.g., 100 = 1% bump on first bid)",
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

	// Parse divisor parameter
	divisor, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		cutils.Fatal("Error parsing divisor parameter: %v", err)
	}
	if divisor <= 0 {
		cutils.Fatal("Divisor must be positive")
	}

	// Contract setup
	cosmicGameAddr := common.HexToAddress(os.Args[1])
	cutils.PrintContractInfo("CosmicGame Address", cosmicGameAddr)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}

	// Get current value
	copts := cutils.CreateCallOpts()

	currentDivisor, err := cosmicGame.InitialDurationUntilMainPrizeDivisor(copts)
	if err != nil {
		cutils.Fatal("Error getting current divisor: %v", err)
	}

	owner, err := cosmicGame.Owner(copts)
	if err != nil {
		cutils.Fatal("Error getting contract owner: %v", err)
	}

	cutils.Section("CURRENT STATE")
	cutils.PrintKeyValue("Contract Owner", owner.String())
	cutils.PrintKeyValue("Current Divisor", currentDivisor.String())
	cutils.PrintKeyValue("Current Percentage", cutils.ConvertToPercentage(currentDivisor))

	cutils.Section("NEW VALUES")
	cutils.PrintKeyValue("New Divisor", divisor)
	cutils.PrintKeyValue("New Percentage", cutils.ConvertToPercentage(big.NewInt(divisor)))
	cutils.PrintKeyValue("Formula", "percentage = 100 / divisor")

	// Check ownership
	if acc.Address != owner {
		cutils.Section("WARNING")
		cutils.PrintKeyValue("Your Address", acc.Address.String())
		cutils.PrintKeyValue("Note", "You are NOT the contract owner. Transaction will likely fail.")
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("SetInitialDurationUntilMainPrizeDivisor", nil, cutils.GasLimitAdminCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall)

	tx, err := cosmicGame.SetInitialDurationUntilMainPrizeDivisor(txopts, big.NewInt(divisor))
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}




