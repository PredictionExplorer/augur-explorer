// Gets contracts code and saves to file
package main

import (
	"os"
	"fmt"
	"math/big"
	"strconv"
	"context"

	"github.com/ethereum/go-ethereum/ethclient"

//	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
var (
	RPC_URL string
)
func main() {

	if len(os.Args) !=3 {
		fmt.Printf("Usage: %v [block_num] [tx_index]\n",os.Args[0])
		os.Exit(1)
	}

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC, please set RPC_URL env variable : %v\n",err)
		os.Exit(0)
	}

	block_num ,err := strconv.ParseInt(os.Args[1],10,64)
	if err != nil {
		fmt.Printf("Can't convert block to number: %v\n",err)
		os.Exit(1)
	}
	big_block_num := big.NewInt(block_num)

	header,err := eclient.HeaderByNumber(context.Background(),big_block_num)
	if err!=nil {
		fmt.Printf("Can't get block's hash by number: %v\n",err)
		os.Exit(1)
	}

	tx_index, err := strconv.ParseInt(os.Args[2],10,64)
	if err != nil {
		fmt.Printf("Can't convert tx index to number: %v\n",err)
		os.Exit(1)
	}

	tx,err := eclient.TransactionInBlock(context.Background(),header.Hash(),uint(tx_index))
	if err == nil {
		fmt.Printf("tx hash is %v\n",tx.Hash().String())
	} else {
		fmt.Printf("Error at RPC: %v\n",err)
	}
}
