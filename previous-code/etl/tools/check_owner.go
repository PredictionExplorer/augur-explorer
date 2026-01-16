// Gets Owner address (EOA) of a Wallet contract in Augur Platform
// If owner exists (i.e. Wallet contract is deployed, returns address of the wallet contract)
// If no wallet contract exists, returns Zero-address
package main

import (
	"os"
	"fmt"
	"math/big"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
var (
	RPC_URL string
)
func main() {

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		Fatalf("Can't connect to ETH RPC, please set AUGUR_ETH_NODE_RPC_URL env variable : %v\n",err)
	}

	if len(os.Args) < 2 {
		fmt.Printf("Usage: \n\t\t%v [wallet_contract_addr]\n",os.Args[0])
		os.Exit(1)
	}

	num:=big.NewInt(int64(OWNER_FIELD_OFFSET))
	key:=common.BigToHash(num)
	wallet_addr := common.HexToAddress(os.Args[1])

	eoa,err := eclient.StorageAt(context.Background(),wallet_addr,key,nil)
	if err == nil {
		eth_addr := common.BytesToAddress(eoa[12:])
		fmt.Printf("%v\n",eth_addr.String())
	} else {
		Fatalf("Error at RPC: %v\n",err)
	}
}
