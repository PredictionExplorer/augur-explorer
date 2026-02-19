// Sets delayDurationBeforeRoundActivation, then claims prize
// This ensures the new round activates after the specified delay
package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	cutils.ParseInfoFlag()
	// Usage check
	if len(os.Args) < 2 || len(os.Args) > 3 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicgame_contract_addr] [delay_seconds]",
			"Sets delayDurationBeforeRoundActivation, then claims prize. delay_seconds defaults to 60 if not provided",
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

	// Parse delay seconds
	delaySeconds := int64(60)
	if len(os.Args) == 3 {
		delaySeconds, err = strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			cutils.Fatal("Error parsing delay_seconds: %v", err)
		}
	}

	// Contract setup
	cosmicGameAddr := common.HexToAddress(os.Args[1])
	cutils.PrintContractInfo("CosmicGame Address", cosmicGameAddr)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}

	// Get current state
	copts := cutils.CreateCallOpts()

	currentDelay, err := cosmicGame.DelayDurationBeforeRoundActivation(copts)
	if err != nil {
		cutils.Fatal("Error getting current delay: %v", err)
	}

	prizeAmount, err := cosmicGame.GetMainEthPrizeAmount(copts)
	if err != nil {
		cutils.Fatal("Error getting prize amount: %v", err)
	}

	cutils.Section("CURRENT STATE")
	cutils.PrintKeyValueDuration("Current Delay", currentDelay.Int64())
	cutils.PrintKeyValueDuration("New Delay To Set", delaySeconds)
	cutils.PrintKeyValueEth("Prize Amount", prizeAmount)

	// Step 1: Set delayDurationBeforeRoundActivation
	cutils.Section("STEP 1: SET DELAY")
	cutils.PrintTxSubmitting("SetDelayDurationBeforeRoundActivation", nil, cutils.GasLimitAdminCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall)

	tx1, err := cosmicGame.SetDelayDurationBeforeRoundActivation(txopts, big.NewInt(delaySeconds))
	cutils.PrintTxResult(tx1, err)

	if err != nil {
		cutils.Fatal("Failed to set delay, aborting")
	}

	// Wait for tx to be mined
	fmt.Printf("\nWaiting 2 seconds for tx to be mined...\n")
	time.Sleep(2 * time.Second)

	// Refresh account nonce
	acc, err = cutils.PrepareAccount(net, cutils.MustGetPkeyHex())
	if err != nil {
		cutils.Fatal("Account refresh failed: %v", err)
	}

	// Refresh gas price
	net2, err := cutils.ConnectToRPC()
	if err != nil {
		cutils.Fatal("Network refresh failed: %v", err)
	}

	// Step 2: Claim the prize
	cutils.Section("STEP 2: CLAIM PRIZE")
	cutils.PrintTxSubmitting("ClaimMainPrize", nil, cutils.GasLimitClaimPrize, net2.GasPrice)

	txopts2 := cutils.CreateTransactOpts(net2, acc, nil, cutils.GasLimitClaimPrize)

	tx2, err := cosmicGame.ClaimMainPrize(txopts2)
	cutils.PrintTxResult(tx2, err)

	if err != nil {
		os.Exit(1)
	}

	cutils.Section("SUMMARY")
	cutils.PrintKeyValueDuration("Delay set to", delaySeconds)
	cutils.PrintKeyValue("Status", "Prize claimed successfully")
	cutils.PrintKeyValue("Note", fmt.Sprintf("New round will activate %d seconds after the claim", delaySeconds))
}
