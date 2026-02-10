// Makes a bid at CosmicGame current round
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
	if len(os.Args) != 2 {
		cutils.PrintUsage(os.Args[0],
			"[cosmicgame_contract_addr]",
			"Makes a bid at CosmicGame current round",
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

	// Get current round info
	copts := cutils.CreateCallOpts()

	roundNum, err := cosmicGame.RoundNum(copts)
	if err != nil {
		cutils.Fatal("Error getting round number: %v", err)
	}

	bidPrice, err := cosmicGame.GetNextEthBidPrice(copts)
	if err != nil {
		cutils.Fatal("Error getting bid price: %v", err)
	}

	lastBidder, err := cosmicGame.LastBidderAddress(copts)
	if err != nil {
		cutils.Fatal("Error getting last bidder: %v", err)
	}

	totalBids, err := cosmicGame.GetTotalNumBids(copts, roundNum)
	if err != nil {
		cutils.Fatal("Error getting total bids: %v", err)
	}

	cutils.Section("ROUND INFO")
	cutils.PrintKeyValue("Round Number", roundNum.String())
	cutils.PrintKeyValue("Total Bids This Round", totalBids.String())
	cutils.PrintKeyValue("Last Bidder", lastBidder.String())
	cutils.PrintKeyValueEth("Next Bid Price", bidPrice)

	// Check if account has enough balance
	if acc.Balance.Cmp(bidPrice) < 0 {
		cutils.Fatal("Insufficient balance. Need %s ETH, have %s ETH",
			cutils.WeiToEth(bidPrice), cutils.WeiToEth(acc.Balance))
	}

	// Create and submit transaction
	cutils.PrintTxSubmitting("BidWithEth", bidPrice, cutils.GasLimitBid, net.GasPrice)

	txopts := cutils.CreateTransactOpts(net, acc, bidPrice, cutils.GasLimitBid)

	tx, err := cosmicGame.BidWithEth(txopts, big.NewInt(-1), "")
	cutils.PrintTxResult(tx, err)

	if err != nil {
		os.Exit(1)
	}
}
