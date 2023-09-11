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

	if len(os.Args) < 3 {
		fmt.Printf(
			"Usage: \n\t\t%v [token_addr] [owner] [operator]\n\t\t"+
			"Gets ERC721 approved for all status (operator level approval)\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	var copts bind.CallOpts
	cosmic_sig_addr := common.HexToAddress(os.Args[1])
	fmt.Printf("Calling to contract at %v\n",cosmic_sig_addr.String())
	owner := common.HexToAddress(os.Args[2])
	operator := common.HexToAddress(os.Args[3])

	cosmic_sig_ctrct,err := NewCosmicSignature(cosmic_sig_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicSignature contract: %v\n",err)
		os.Exit(1)
	}

	is_approved_for_all,err := cosmic_sig_ctrct.IsApprovedForAll(&copts,owner,operator)
	if err != nil {
		fmt.Printf("Error at IsApprovedForAll()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	fmt.Printf("Owner: %v\n",owner.String())
	fmt.Printf("Operator:  %v\n",operator.String())
	fmt.Printf("Approved: %v\n",is_approved_for_all)
}
