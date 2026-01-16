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
func fmt_eth(wei *big.Int) string {
    ether := new(big.Float).SetInt(wei)
    eth_value := new(big.Float).Quo(ether, big.NewFloat(1e18))
    return eth_value.Text('f', 18) // 18 decimal places to match Ethereum precision
}
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Printf(
			"Usage: \n\t\t%v [market_addr]\n\t\t"+
			"Reads variables from Market contract\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	var copts bind.CallOpts
	rwalk_addr := common.HexToAddress(os.Args[1])
	fmt.Printf("Calling to contract at %v\n",rwalk_addr.String())

	rwalk_ctrct,err := NewRWMarket(rwalk_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate RWMarket contract: %v\n",err)
		os.Exit(1)
	}

	num_offers,err := rwalk_ctrct.NumOffers(&copts)
	if err != nil {
		fmt.Printf("Error at NumOffers()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	fmt.Printf("NumOffers = %v\n",num_offers.Int64())
}
