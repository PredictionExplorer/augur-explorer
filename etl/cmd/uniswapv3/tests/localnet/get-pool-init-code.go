// Shows allowance of a user in an ERC20 token
package main

import (
	"os"
	"fmt"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

)
const (
)
var (
	RPC_URL string
	token_addr		common.Address
)
func main() {

	code,err := hex.DecodeString(UniswapV3PoolBin[2:])
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	hash := crypto.Keccak256Hash(code)
	fmt.Printf("Code hash: %v\n",hash.String())
	fmt.Printf("Code length: %v\n",len(code))
}
