// Sets raffleTotalEthPrizeAmountForBiddersPercentage (percentage of funds for raffle ETH prizes)
package main

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	cutils.ParseInfoFlag()
	// Usage check
	if len(os.Args) < 3 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicgame_contract_addr] [percentage]",
			"Sets raffleTotalEthPrizeAmountForBiddersPercentage (percentage of funds for raffle ETH prizes)",
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

	percentageVal := big.NewInt(0)
	_, success := percentageVal.SetString(os.Args[2], 10)
	if !success {
		cutils.Fatal("Invalid percentage value provided: %s", os.Args[2])
	}

	// Contract setup
	cutils.PrintContractInfo("CosmicGame Address", cosmicGameAddr)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}

	// Get current value
	copts := cutils.CreateCallOpts()

	currentValue, err := cosmicGame.RaffleTotalEthPrizeAmountForBiddersPercentage(copts)
	if err != nil {
		cutils.Fatal("Error getting current value: %v", err)
	}

	cutils.Section("RAFFLE ETH PRIZE PERCENTAGE CONFIG")
	cutils.PrintKeyValue("Current Value", currentValue.String()+"%")
	cutils.PrintKeyValue("New Value", percentageVal.String()+"%")

	// Create and submit transaction
	cutils.PrintTxSubmitting("SetRaffleTotalEthPrizeAmountForBiddersPercentage", nil, cutils.GasLimitAdminCall, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, nil, cutils.GasLimitAdminCall)

	tx, err := cosmicGame.SetRaffleTotalEthPrizeAmountForBiddersPercentage(txopts, percentageVal)
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
