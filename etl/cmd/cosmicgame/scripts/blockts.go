package main

import (
	"os"
	"fmt"
	"context"
	"time"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"

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
	for {
		bnum := big.NewInt(-1)
		block,err := eclient.BlockByNumber(context.Background(),bnum)
		if err != nil {
			fmt.Printf("Error in BlockByNumber(): %v\n",err)
		} else {
			fmt.Printf("timestamp = %v\n",block.Time)
		}
		time.Sleep(5 * time.Second)
	}
}
