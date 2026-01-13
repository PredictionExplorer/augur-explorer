// Displays current ERC20 allowance for a spender
package main

import (
	"os"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
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

	if len(os.Args) < 4 {
		fmt.Printf(
			"Usage: \n\t\t%v [erc20_token_addr] [owner_addr] [spender_addr]\n\t\t"+
			"Displays current allowance for spender to spend owner's ERC20 tokens\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	token_addr := common.HexToAddress(os.Args[1])
	owner_addr := common.HexToAddress(os.Args[2])
	spender_addr := common.HexToAddress(os.Args[3])

	erc20_ctrct,err := NewERC20(token_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC20 contract: %v\n",err)
		os.Exit(1)
	}

	var copts bind.CallOpts
	allowance, err := erc20_ctrct.Allowance(&copts, owner_addr, spender_addr)
	if err!=nil {
		fmt.Printf("Error querying allowance: %v\n",err)
		os.Exit(1)
	}

	fmt.Printf("Token address:   %v\n", token_addr.String())
	fmt.Printf("Owner address:   %v\n", owner_addr.String())
	fmt.Printf("Spender address: %v\n", spender_addr.String())
	fmt.Printf("Allowance:       %v\n", allowance.String())
}

