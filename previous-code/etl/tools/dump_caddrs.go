// Dumps ShareTokens that an account currently holds
package main

import (
	"os"
	"fmt"
	"log"
	
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"


	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
const (
)
var (
	RPC_URL string
)
func main() {

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		Fatalf("Can't connect to ETH RPC: %v\n",err)
	}
	if len(os.Args) < 2 {
		fmt.Printf("Usage: \n\t\t%v [augur_contract_address] \n",os.Args[0])
		os.Exit(1)
	}

	augur_contract_addr := common.HexToAddress(os.Args[1])

	Info := log.New(os.Stdout,"",0)
	caddrs,err := Get_contract_addresses_from_net(augur_contract_addr,eclient)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		caddrs.Dump(Info)
		os.Exit(1)
	}
	caddrs.Dump(Info)
	os.Exit(0)
}
