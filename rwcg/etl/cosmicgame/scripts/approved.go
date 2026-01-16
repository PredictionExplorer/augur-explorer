package main

import (
	"os"
	"fmt"
	"strconv"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
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
			"Usage: \n\t\t%v [erc721_addr] [tokenid]\n\t\t"+
			"Gets ERC721 approved status status (single (token level) approval)\n\n",os.Args[0],
		)
		os.Exit(1)
	}

	var copts bind.CallOpts
	erc721_addr := common.HexToAddress(os.Args[1])
	fmt.Printf("Calling to contract at %v\n",erc721_addr.String())
	tokenid_str := os.Args[2]
	token_id,err := strconv.ParseInt(tokenid_str,10,64)
	if err != nil {
		fmt.Printf("error parsing tokenid: %v\n",err)
		os.Exit(1)
	}

	erc721_ctrct,err := NewCosmicSignatureNft(erc721_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC721 contract: %v\n",err)
		os.Exit(1)
	}

	operator,err := erc721_ctrct.GetApproved(&copts,big.NewInt(token_id))
	if err != nil {
		fmt.Printf("Error at GetApproved()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	fmt.Printf("Contract: %v\n",erc721_addr.String())
	fmt.Printf("Token id: %v\n",token_id);
	fmt.Printf("Operator: %v\n",operator.String())
}
