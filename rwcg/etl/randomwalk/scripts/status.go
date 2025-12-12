package main

import (
	"os"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/rwcg/contracts/randomwalk"
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

	if len(os.Args) < 2 {
		fmt.Printf(
			"Usage: \n\t\t%v [rwalk_addr]\n\t\t"+
			"Reads variables from  RandomWalk contract\n\n",os.Args[0],
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

	next_token_id,err := rwalk_ctrct.NextTokenId(&copts)
	if err != nil {
		fmt.Printf("Error at NexttokenId()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	time_remaining,err := rwalk_ctrct.TimeUntilWithdrawal(&copts)
	if err != nil {
		fmt.Printf("Error at timeUntilWithdrawal()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	withdrawal_amount,err := rwalk_ctrct.WithdrawalAmount(&copts)
	if err != nil {
		fmt.Printf("Error at withdrawalAmount()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	num_withdrawals,err := rwalk_ctrct.NumWithdrawals(&copts)
	if err != nil {
		fmt.Printf("Error at numWithdrawals()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	base_uri,err := rwalk_ctrct.TokenURI(&copts,big.NewInt(0))
	if err != nil {
		fmt.Printf("Error at TokenURI()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	fmt.Printf("Next token ID = %v\n",next_token_id.Int64())
	fmt.Printf("Time remaining: %v\n",time_remaining.Int64())
	fmt.Printf("Withdrawal amount: %v\n",fmt_eth(withdrawal_amount))
	fmt.Printf("Num withdrawals: %v\n",num_withdrawals.Int64())
	fmt.Printf("Base uri: %v\n",base_uri)
}
