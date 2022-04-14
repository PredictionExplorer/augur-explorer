package main

import (
	"os"
	"fmt"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	CONTRACT_ADDR string = "0xBA12222222228d8Ba445958a75a0704d566BF2C8"
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
			"Usage: \n\t\t%v [pool_id]\n\t\t"+
			"Gets token list and balances of tokens from Balancer v2 Vault contract\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	pool_id_bytes,err := hex.DecodeString(os.Args[1])
	if err != nil {
		fmt.Printf("Can't decode pool id hex: %v\n",err)
		os.Exit(1)
	}
	var pool_id_32b [32]byte
	for i:=0;i<32;i++ {
		pool_id_32b[i] = pool_id_bytes[i]
	}

	var copts bind.CallOpts
	vault_addr := common.HexToAddress(CONTRACT_ADDR)
	fmt.Printf("Calling Vault contract at %v\n",vault_addr.String())

	vault_ctrct,err := NewBalancerV2Vault(vault_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Vault contract: %v\n",err)
		os.Exit(1)
	}

	result,err := vault_ctrct.GetPoolTokens(&copts,pool_id_32b)
	if err != nil {
		fmt.Printf("Error at GetPoolTokens()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	tokens := result.Tokens
	balances := result.Balances
	for i:=0;i<len(tokens);i++ {
		fmt.Printf("\t%v : \t%v\n",tokens[i].String(),balances[i].String())
	}
}
