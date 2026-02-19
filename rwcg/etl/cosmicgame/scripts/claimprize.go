// Claims the main prize from CosmicGame
package main

import (
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	cutils.ParseInfoFlag()
	// Usage check
	if len(os.Args) != 2 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicgame_contract_addr]",
			"Claims the main prize from CosmicGame",
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

	// Contract setup
	cosmicGameAddr := common.HexToAddress(os.Args[1])
	cutils.PrintContractInfo("CosmicGame Address", cosmicGameAddr)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}

	// Get current state info
	copts := cutils.CreateCallOpts()

	roundNum, err := cosmicGame.RoundNum(copts)
	if err != nil {
		cutils.Fatal("Error getting round number: %v", err)
	}

	lastBidder, err := cosmicGame.LastBidderAddress(copts)
	if err != nil {
		cutils.Fatal("Error getting last bidder: %v", err)
	}

	prizeAmount, err := cosmicGame.GetMainEthPrizeAmount(copts)
	if err != nil {
		cutils.Fatal("Error getting prize amount: %v", err)
	}

	durationUntilPrize, err := cosmicGame.GetDurationUntilMainPrize(copts)
	if err != nil {
		cutils.Fatal("Error getting duration until prize: %v", err)
	}

	cutils.Section("PRIZE INFO")
	cutils.PrintKeyValue("Round Number", roundNum.String())
	cutils.PrintKeyValue("Last Bidder", lastBidder.String())
	cutils.PrintKeyValueEth("Prize Amount", prizeAmount)
	cutils.PrintKeyValueDuration("Time Until Prize", durationUntilPrize.Int64())

	// Check if caller is the last bidder
	if acc.Address != lastBidder {
		cutils.Section("WARNING")
		cutils.PrintKeyValue("Your Address", acc.Address.String())
		cutils.PrintKeyValue("Last Bidder", lastBidder.String())
		cutils.PrintKeyValue("Note", "You are NOT the last bidder. Claim may fail unless timeout has passed.")
	}

	// Check if prize can be claimed
	if durationUntilPrize.Int64() > 0 {
		cutils.Section("WARNING")
		cutils.PrintKeyValue("Status", "Prize is NOT yet claimable")
		cutils.PrintKeyValueDuration("Wait Time Remaining", durationUntilPrize.Int64())
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("ClaimMainPrize", nil, cutils.GasLimitClaimPrize, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitClaimPrize)

	tx, err := cosmicGame.ClaimMainPrize(txopts)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
