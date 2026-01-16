package main

import (
	"os"
	"fmt"
	"strconv"
	"math/big"
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
	block_num_to_query		int64 = 0
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
			"Usage: \n\t\t%v [pool_id] [[block_num]]\n\t\t"+
			"Gets token list and balances of tokens from Balancer v2 Vault contract (block number is optional)\n\n",os.Args[0],
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
	if len(os.Args) == 3 {// block number is provided
		var err error
		block_num_to_query,err = strconv.ParseInt(os.Args[2],10,64)
		if err != nil {
			fmt.Sprintf("Error parsing block number: %v\n",err)
			os.Exit(1)
		}
	}
	var for_block string
	var copts bind.CallOpts
	if block_num_to_query > 0 {
		copts.BlockNumber = big.NewInt(block_num_to_query)
		for_block = fmt.Sprintf(" for block %v", block_num_to_query)
	}
	vault_addr := common.HexToAddress(CONTRACT_ADDR)
	
	fmt.Printf("Calling Vault contract at %v%v\n",vault_addr.String(),for_block)

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
