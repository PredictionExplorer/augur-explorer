//Update addr.reverse (Reverse) ENS names
package main

import (
	"os"
	"fmt"
	"log"
	"context"
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wealdtech/go-ens/v3"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)

var (
	RPC_URL =			os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient				*ethclient.Client
	rpcclient			*rpc.Client
	Info                *log.Logger
	storage             *SQLStorage
	market_order_id		int64 = 0
)
func main() {

	rpcclient, err := rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	eclient = ethclient.NewClient(rpcclient)

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)

	addresses := storage.Get_addrs_with_reverse_name()
	for _,addr_str := range addresses {
		addr := common.HexToAddress(addr_str)
		//	node,err := NameHash(fmt.Sprintf("%s.addr.reverse", address.Hex()[2:]))
		//	if err != nil {
		//		fmt.Printf("Error: %v\n",err)
		//		os.Exit(1)
		//	}
		domain, err := ens.ReverseResolve(eclient, addr)
		if err != nil {
			fmt.Printf("Addr %48v\tError: %v\n",addr_str,err)
			continue
		}
		fmt.Printf("Addr: %48v\t%v\n",addr_str,domain)
		parsed := strings.Split(domain,".")
		if len(parsed) == 0 {
			fmt.Printf("Error: name is of length=0\n")
			continue
		}
		word := parsed[0]
		label,err := LabelHash(word)
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			continue
		}
		label_str := hex.EncodeToString(label[:])
		fmt.Printf("Name found: label %v , word %v\n",label_str,word)
	}
}
