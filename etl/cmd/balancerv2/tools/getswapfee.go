package main

import (
	"os"
	"fmt"
	"strings"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	CONTRACT_ADDR string = "0xBA12222222228d8Ba445958a75a0704d566BF2C8"
)
var (
	RPC_URL string
	fee_collection_abi *abi.ABI
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

	abi_parsed := strings.NewReader(FeeCollectionABI)
	abi,err := abi.JSON(abi_parsed)
	if err != nil {
		fmt.Printf("Can't parse FeeCollection ABI")
		os.Exit(1)
	}
	fee_collection_abi = &abi

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

	pool_addr,specialization,err := vault_ctrct.GetPool(&copts,pool_id_32b)
	if err != nil {
		fmt.Printf("Error at GetPool()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}
	fmt.Printf("Pool address: %v\n",pool_addr.String())
	fmt.Printf("Specialization: %v\n",specialization)

	pool_ctrct,err := NewGetSwapFee(pool_addr,eclient)
	if err != nil {
		fmt.Printf("Can't instantiate pool contract:%v\n",err)
		os.Exit(1)
	}
	result,err := pool_ctrct.GetSwapFeePercentage(&copts)
	if err != nil {
		fmt.Printf("Call to GetSwapFeePercentage() failed: %v\n",err)
		os.Exit(1)
	}
	fee_percentage_str := result.String()
	fmt.Printf("Pool swap fee: %v\n",fee_percentage_str)
}
