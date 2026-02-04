// Dumps ETH donation with info records from CosmicGame
package main

import (
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

func main() {
	// Usage check
	var cgAddr string
	if len(os.Args) < 2 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicgame_contract_addr]",
			"Dumps ETH donation with info records from CosmicGame",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)"},
		)
		cutils.Section("DEFAULT ADDRESS")
		cgAddr = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
		cutils.PrintKeyValue("Using default", cgAddr)
	} else {
		cgAddr = os.Args[1]
	}

	// Connect to network
	net, err := cutils.ConnectToRPC()
	if err != nil {
		cutils.Fatal("Network connection failed: %v", err)
	}
	cutils.PrintNetworkInfo(net)

	// Contract setup
	cosmicGameAddr := common.HexToAddress(cgAddr)
	cutils.PrintContractInfo("CosmicGame Address", cosmicGameAddr)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}

	// Get number of records
	copts := cutils.CreateCallOpts()

	numRecs, err := cosmicGame.NumEthDonationWithInfoRecords(copts)
	if err != nil {
		cutils.Fatal("Error getting NumEthDonationWithInfoRecords: %v", err)
	}

	cutils.Section("DONATION RECORDS")
	cutils.PrintKeyValue("Total Records", numRecs.String())

	n := numRecs.Int64()
	if n == 0 {
		cutils.PrintKeyValue("Note", "No ETH donation with info records found")
		return
	}

	// Iterate through records
	for i := int64(0); i < n; i++ {
		rec, err := cosmicGame.EthDonationWithInfoRecords(copts, big.NewInt(i))
		if err != nil {
			cutils.PrintKeyValue("Error at record "+string(rune(i)), err.Error())
		} else {
			cutils.Section("RECORD " + big.NewInt(i).String())
			cutils.PrintKeyValue("Round Number", rec.RoundNum.String())
			cutils.PrintKeyValue("Donor Address", rec.DonorAddress.String())
			cutils.PrintKeyValueEth("Amount", rec.Amount)
			cutils.PrintKeyValue("Data", rec.Data)
		}
	}
}
