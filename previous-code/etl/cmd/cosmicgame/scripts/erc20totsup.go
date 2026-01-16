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
			"Usage: \n\t\t%v [erc20_addr]\n\t\t"+
			"Gets erc20 totalsupply value\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	var copts bind.CallOpts
	caddr := common.HexToAddress(os.Args[1])
	fmt.Printf("Creating contract instance for address %v\n",caddr.String())

	ctrct,err := NewERC20(caddr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC20 contract: %v\n",err)
		os.Exit(1)
	}

	totsup,err := ctrct.TotalSupply(&copts)
	if err != nil {
		fmt.Printf("Error at TotalSupply()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	fmt.Printf("Total Supply: %v\n",totsup.String())
}
