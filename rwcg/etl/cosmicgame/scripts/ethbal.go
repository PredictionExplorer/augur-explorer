package main

import (
	"os"
	"fmt"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	//. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/cosmicgame"
)
const (
)
var (
	RPC_URL string
)
func fmt_eth(wei *big.Int) string {
	ether := new(big.Float).SetInt(wei)
	eth_value := new(big.Float).Quo(ether, big.NewFloat(1e18))
	return eth_value.Text('f', 18) // 18 decimal places to match Ethereum precision
}

func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	var addr string
	if len(os.Args) < 2 {
		fmt.Printf(
			"Usage: \n\t\t%v [address]\n\t\t"+
			"Gets balance of a user (ETH)\n\n",os.Args[0],
		)
		os.Exit(1)
	} else {
		addr = os.Args[1]
	}
	user_addr := common.HexToAddress(addr)

	balance,err := eclient.BalanceAt(context.Background(),user_addr,nil)
	if err != nil {
		fmt.Printf("Error at balanceAt()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	fmt.Printf("Address %v balance: %v\n",user_addr.String(),fmt_eth(balance))
}
