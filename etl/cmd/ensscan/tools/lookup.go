// ENS Reverse lookup by addr

package main

import (
	"os"
	"fmt"
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/wealdtech/go-ens/v3"

)
const (
	REVERSE_REG_V2_ADDR         = "0xA2C122BE93b0074270ebeE7f6b7292C7deB45047"
	REVERSE_REG_V1_ADDR         = "0x5fBb459C49BB06083C33109fA4f14810eC2Cf358"
)
var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
)
func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %v [name]\n")
		os.Exit(1)
	}
	rpcclient, err := rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	eclient = ethclient.NewClient(rpcclient)

	domain := os.Args[1]
	address, err := ens.Resolve(eclient, domain)

	if err!=nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Address: %v\n",address.String())

}
