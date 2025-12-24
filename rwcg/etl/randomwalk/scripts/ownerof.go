package main

import (
	"os"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/randomwalk"
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

	if len(os.Args) < 3 {
		fmt.Printf(
			"Usage: \n\t\t%v [rwalk_addr] [token_id]\n\t\t"+
			"Gets ownership of a token from  RandomWalk contract\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	var copts bind.CallOpts
	rwalk_addr := common.HexToAddress(os.Args[1])
	token_id,err := strconv.ParseInt(os.Args[2],10,64)
	if err != nil {
		fmt.Printf("Error parsing token_id field: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Calling to contract at %v\n",rwalk_addr.String())

	rwalk_ctrct,err := NewRWalk(rwalk_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate RWalk contract: %v\n",err)
		os.Exit(1)
	}

	owner,err := rwalk_ctrct.OwnerOf(&copts,big.NewInt(token_id))
	if err != nil {
		fmt.Printf("Error at OwnerOf()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	fmt.Printf("Owner  = %v\n",owner.String())
}
