package main

import (
	"os"
	"fmt"
	"strconv"
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

	if len(os.Args) < 3 {
		fmt.Printf(
			"Usage: \n\t\t%v [cosmic_signature_addr] [tokenid]\n\t\t"+
			"Gets CosmicSignature approved for all status\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	var copts bind.CallOpts
	cosmic_sig_addr := common.HexToAddress(os.Args[1])
	fmt.Printf("Calling to contract at %v\n",cosmic_sig_addr.String())
	tokenid_str := os.Args[2]
	token_id,err := strconv.ParseInt(tokenid_str,10,64)
	if err != nil {
		fmt.Printf("error parsing tokenid: %v\n",err)
		os.Exit(1)
	}

	cosmic_sig_ctrct,err := NewCosmicSignature(cosmic_sig_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicSignature contract: %v\n",err)
		os.Exit(1)
	}

	ownerof,err := cosmic_sig_ctrct.OwnerOf(&copts,big.NewInt(token_id))
	if err != nil {
		fmt.Printf("Error at Ownerof()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	fmt.Printf("Owner: %v\n",ownerof.String())
}
