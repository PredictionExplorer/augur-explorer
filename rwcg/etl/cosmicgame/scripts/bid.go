// Makes a bid at CosmicGame current round
package main

import (
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
	cutils "github.com/PredictionExplorer/augur-explorer/rwcg/etl/cosmicgame/scripts/common"
)

// pickContractArg returns the argument that looks like an Ethereum address (0x + 40 hex chars).
// If one arg is given, it is returned. If two are given (e.g. tx hash + address), returns the address.
func pickContractArg(args []string) string {
	if len(args) == 0 {
		return ""
	}
	if len(args) == 1 {
		return args[0]
	}
	// Two args: pick the one that looks like an address (40 hex chars; 64 = tx hash)
	for _, a := range args {
		trimmed := strings.TrimPrefix(a, "0x")
		if len(trimmed) == 40 {
			if strings.HasPrefix(a, "0x") {
				return a
			}
			return "0x" + trimmed
		}
	}
	return args[1]
}

func main() {
	cutils.ParseInfoFlag()
	// Usage check: need at least one arg (contract address)
	if len(os.Args) < 2 {
		cutils.PrintUsage(os.Args[0],
			"[-i] [cosmicgame_contract_addr]",
			"Makes a bid at CosmicGame current round. Use -i for detailed output.",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)", "PKEY_HEX": "64-char hex private key, no 0x prefix (required)"},
		)
		os.Exit(1)
	}

	contractArg := pickContractArg(os.Args[1:])
	if contractArg == "" || !common.IsHexAddress(contractArg) {
		cutils.PrintUsage(os.Args[0],
			"[-i] [cosmicgame_contract_addr]",
			"Makes a bid at CosmicGame current round. Use -i for detailed output.",
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
	cosmicGameAddr := common.HexToAddress(contractArg)
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
