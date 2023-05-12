package main

import (
	"os"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
)
var (
	RPC_URL string
)
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Printf(
			"Usage: \n\t\t%v [cosmic_game_addr]\n\t\t"+
			"Gets CosmicGame read only variables\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	var copts bind.CallOpts
	cosmic_game_addr := common.HexToAddress(os.Args[1])
	fmt.Printf("Calling to contract at %v\n",cosmic_game_addr.String())

	cosmic_game_ctrct,err := NewCosmicGame(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
		os.Exit(1)
	}

	time_until_prize,err := cosmic_game_ctrct.TimeUntilPrize(&copts)
	if err != nil {
		fmt.Printf("Error at TimeUntilPrize()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	fmt.Printf("Time until prize = %v\n",time_until_prize.Int64())
}
