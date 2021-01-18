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
	"github.com/wealdtech/go-ens/v3/contracts/reverseresolver"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
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
	fmt.Printf("Name from go-ens module: %v\n",domain)

	//old_resolver := ens.NewReverseResolverAt(backend, address)
	ctrct_addr := common.HexToAddress(REVERSE_REG_V2_ADDR)
	fmt.Printf("Querying contract directly (contract addr=%v)\n",ctrct_addr.String())
	rreg,err := reverseresolver.NewContract(ctrct_addr,eclient)
	if err != nil {
		fmt.Printf("Error on NewContract(): %v\n",err)
		os.Exit(1)
	}
	var copts = new(bind.CallOpts)
	name,err := rreg.Name(copts,node)
	if err != nil {
		fmt.Printf("Error on Name(): %+v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Name from ENS V2 ReverseRegistrar: %v\n",name)

	fmt.Printf("\n\n")
	fmt.Printf("Now we are going to check the OLD reverse registrar\n")
	ctrct_addr = common.HexToAddress(REVERSE_REG_V1_ADDR)
	fmt.Printf("Querying contract directly (contract addr=%v)\n",ctrct_addr.String())
	rreg,err = reverseresolver.NewContract(ctrct_addr,eclient)
	if err != nil {
		fmt.Printf("Error on NewContract(): %v\n",err)
		os.Exit(1)
	}
	name,err = rreg.Name(copts,node)
	if err != nil {
		fmt.Printf("Error on Name(): %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Name from ENS V1 (OLD) ReverseRegistrar: %+v\n",name)
}
