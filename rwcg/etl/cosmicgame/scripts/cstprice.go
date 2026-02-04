// Gets the current CST bid price from CosmicGame
package main

import (
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
			"Gets the current CST bid price from CosmicGame",
			map[string]string{"RPC_URL": "Ethereum RPC endpoint (required)"},
		)
		os.Exit(1)
	}

	// Connect to network
	net, err := cutils.ConnectToRPC()
	if err != nil {
		cutils.Fatal("Network connection failed: %v", err)
	}
	cutils.PrintNetworkInfo(net)

	// Contract setup
	cosmicGameAddr := common.HexToAddress(os.Args[1])
	cutils.PrintContractInfo("CosmicGame Address", cosmicGameAddr)

	cosmicGame, err := NewCosmicSignatureGame(cosmicGameAddr, net.Client)
	if err != nil {
		cutils.Fatal("Failed to instantiate CosmicGame: %v", err)
	}

	// Get CST price info
	copts := cutils.CreateCallOpts()

	cstPrice, err := cosmicGame.GetNextCstBidPrice(copts)
	if err != nil {
		cutils.Fatal("Error getting CST bid price: %v", err)
	}

	ethPrice, err := cosmicGame.GetNextEthBidPrice(copts)
	if err != nil {
		cutils.Fatal("Error getting ETH bid price: %v", err)
	}

	// Dutch auction info
	cstAuctionDuration, cstAuctionElapsed, err := cosmicGame.GetCstDutchAuctionDurations(copts)
	if err != nil {
		cutils.Fatal("Error getting CST Dutch auction durations: %v", err)
	}

	cstDutchBeginPrice, err := cosmicGame.CstDutchAuctionBeginningBidPrice(copts)
	if err != nil {
		cutils.Fatal("Error getting CST Dutch auction begin price: %v", err)
	}

	cutils.Section("BID PRICES")
	cutils.PrintKeyValue("CST Bid Price", cutils.WeiToEth(cstPrice)+" CST")
	cutils.PrintKeyValueEth("ETH Bid Price", ethPrice)

	cutils.Section("CST DUTCH AUCTION")
	cutils.PrintKeyValue("Auction Duration", cstAuctionDuration.String()+" seconds")
	cutils.PrintKeyValue("Auction Elapsed", cstAuctionElapsed.String()+" seconds")
	cutils.PrintKeyValue("Begin Price", cutils.WeiToEth(cstDutchBeginPrice)+" CST")
	cutils.PrintKeyValue("Current Price", cutils.WeiToEth(cstPrice)+" CST")
}
