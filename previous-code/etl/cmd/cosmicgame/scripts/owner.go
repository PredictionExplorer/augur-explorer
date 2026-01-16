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
			"Usage: \n\t\t%v [contract_addr] [tokenid]\n\t\t"+
			"Gets ownership of a contract\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	var copts bind.CallOpts
	ctrct_addr := common.HexToAddress(os.Args[1])
	fmt.Printf("Calling to contract at %v\n",ctrct_addr.String())

	ownable_ctrct,err := NewOwnable(ctrct_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicSignature contract: %v\n",err)
		os.Exit(1)
	}

	owner_addr,err := ownable_ctrct.Owner(&copts)
	if err != nil {
		fmt.Printf("Error at Owner()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	fmt.Printf("Owner: %v\n",owner_addr.String())
}
