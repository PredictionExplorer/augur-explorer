// 
package main

import (
	"os"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"


)
const (
	CHAIN_ID		int64 = 1234
)
var (
	RPC_URL string
	token_addr		common.Address
)
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) != 6 {
		fmt.Printf(
			"Usage: \n\t\t%v [contract_addr] [factory_addr] [token0_addr] [token1_addr] [fee]\n\n"+
			"\t\tCalculates pool address\n",
			os.Args[0],
		)
		os.Exit(1)
	}

	contract_addr := common.HexToAddress(os.Args[1])
	factory_addr := common.HexToAddress(os.Args[2])
	token0_addr := common.HexToAddress(os.Args[3])
	token1_addr := common.HexToAddress(os.Args[4])
	fee_str := os.Args[5]

	fee := big.NewInt(0)
	_,success := fee.SetString(fee_str,10)
	if !success {
		fmt.Printf("Incorrect fee number provided on the command line")
		os.Exit(1)
	}

	ctrct,err := NewGetAddr(contract_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate contract: %v\n",err)
		os.Exit(1)
	}

	var pool_key GetAddrPoolKey
	pool_key.Token0=token0_addr
	pool_key.Token1=token1_addr
	pool_key.Fee=big.NewInt(0).Set(fee)

	fmt.Printf("Addr0 %v\n",pool_key.Token0.String())
	fmt.Printf("Addr1 %v\n",pool_key.Token1.String())
	fmt.Printf("Fee %v\n",pool_key.Fee.String())
	fmt.Printf("Factory %v\n",factory_addr.String())

	var copts = new(bind.CallOpts)
	pool_addr,err:=ctrct.ComputeAddress(copts,factory_addr,pool_key)
	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Pool address = %v\n",pool_addr.String())
}
