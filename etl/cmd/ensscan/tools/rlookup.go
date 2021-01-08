// ENS Reverse lookup by addr

package main

import (
	"os"
	"fmt"
	"context"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)

var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
)
func main() {

	if len(os.Args) != 2 {
		fmt.Printf("usage: %v [address]\n")
		os.Exit(1)
	}
	rpcclient, err := rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	eclient = ethclient.NewClient(rpcclient)

	address := common.HexToAddress(os.Args[1])
	node,err := NameHash(fmt.Sprintf("%s.addr.reverse", address.Hex()[2:]))
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Node: %v\n",hex.EncodeToString(node[:]))
	domain, err := ens.ReverseResolve(eclient, address)
	fmt.Printf("Name = %v\n",domain)
}
