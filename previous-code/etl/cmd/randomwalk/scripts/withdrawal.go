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
	CONTRACT_ADDR string = "0x895a6F444BE4ba9d124F61DF736605792B35D66b"
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
			"Usage: \n\t\t%v [rwalk_addr]\n\t\t"+
			"Gets withdrawal amount from  RandomWalk contract\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	var copts bind.CallOpts
	rwalk_addr := common.HexToAddress(os.Args[1])
	fmt.Printf("Calling to contract at %v\n",rwalk_addr.String())

	rwalk_ctrct,err := NewRWalk(rwalk_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate RWalk contract: %v\n",err)
		os.Exit(1)
	}

	wamount,err := rwalk_ctrct.WithdrawalAmount(&copts)
	if err != nil {
		fmt.Printf("Error at WithdrawalAmount()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	fmt.Printf("Withdrawal amount = %v\n",wamount.String())
}
