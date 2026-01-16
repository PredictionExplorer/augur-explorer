// 
package main

import (
	"os"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"


)
const (
	CHAIN_ID		int64 = 1234
)
var (
	RPC_URL string
	token_addr		common.Address
)
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Printf(
			"Usage: \n\t\t%v [contract_addr]\n\n"+
			"\t\tGets factory address\n",
			os.Args[0],
		)
		os.Exit(1)
	}

	contract_addr := common.HexToAddress(os.Args[1])
	ctrct,err := NewNonfungiblePositionManager(contract_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate contract: %v\n",err)
		os.Exit(1)
	}

	var copts = new(bind.CallOpts)
	factory_addr,err:=ctrct.Factory(copts)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Factory address = %v\n",factory_addr.String())
}
