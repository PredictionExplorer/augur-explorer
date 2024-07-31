package main

import (
	"os"
	"fmt"
	"math/big"

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

	var cg_addr string
	if len(os.Args) < 2 {
		fmt.Printf(
			"Usage: \n\t\t%v [cosmic_game_addr]\n\t\t"+
			"Dumps donate with info records\n\n",os.Args[0],
		)
		fmt.Printf("Setting default cosmic game contract address to 0x5FbDB2315678afecb367f032d93F642f64180aa3\n")
		cg_addr = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
	} else {
		cg_addr = os.Args[1]
	}
	var copts bind.CallOpts
	cosmic_game_addr := common.HexToAddress(cg_addr)
	fmt.Printf("Calling to contract at %v\n",cosmic_game_addr.String())

	cosmic_game_ctrct,err := NewCosmicGame(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
		os.Exit(1)
	}

	num_recs,err := cosmic_game_ctrct.DonateWithInfoNumRecords(&copts)
	if err != nil {
		fmt.Printf("Error at DonateWithInfoNumRecords()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	n := num_recs.Int64()
	if n == 0 {
		fmt.Printf("No donate with info records , exiting\n")
		os.Exit(1)
	}
	for i:=int64(0);i<n;i++ {
		rec,err := cosmic_game_ctrct.DonationInfoRecords(&copts,big.NewInt(i))
		if err != nil {
			fmt.Printf("Error calling DonationInfoRecords() :%v\n",err)
		} else {
			fmt.Printf("Record %v\n%v\n",i,rec)
		}
	}
}
