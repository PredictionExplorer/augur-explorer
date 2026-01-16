// Gets test metadata associated to a ENS name

package main

import (
	"os"
	"fmt"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/wealdtech/go-ens/v3"
	"github.com/wealdtech/go-ens/v3/contracts/resolver"

)
var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client
)
func main() {

	if len(os.Args) != 3 {
		fmt.Printf(
			"usage: %v [domain/node] [key]\n\tDomain is text, ex:vitalik.eth\n\tNode is hash " +
			"(64-char long hex, without '0x' prefix!)\n",os.Args[0],
		)
		os.Exit(1)
	}
	rpcclient, err := rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	eclient = ethclient.NewClient(rpcclient)

	domain := os.Args[1]
	key := os.Args[2]

	if len(domain) == 64 {// its a hash
		node_input := common.HexToHash("0x" + domain)
		var node [32]byte
		copy(node[:],node_input.Bytes())
		registry,err := ens.NewRegistry(eclient)
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		resolver_addr,err := registry.Contract.Resolver(nil,node)
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		resolver_ctrct,err := resolver.NewContract(resolver_addr, eclient)
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		text,err := resolver_ctrct.Text(nil,node,key)
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		fmt.Printf("Text metadata for this node is:\n\t%v => %v\n",key,text)

	} else {
		r,err := ens.NewResolver(eclient,domain)
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		
		text,err := r.Text(key)
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		fmt.Printf("Text metadata for ENS name is:\n\t%v => %v\n",key,text)
	}
}
