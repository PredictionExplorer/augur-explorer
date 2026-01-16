// Shows allowance of a user in an ERC20 token
package main

import (
	"os"
	"fmt"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"

)
const (
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

	if len(os.Args) !=2 {
		fmt.Printf("Usage: \n\t\t%v [contract_addr]\n\n\t\tGets code hash of a contract\n",os.Args[0])
		os.Exit(1)
	}
	contract_addr := common.HexToAddress(os.Args[1])
	code, err := eclient.CodeAt(context.Background(),contract_addr,nil)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	hash := crypto.Keccak256Hash(code)
	fmt.Printf("Code hash: %v\n",hash.String())
	fmt.Printf("Code length: %v\n",len(code))
}
