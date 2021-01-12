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
	"github.com/wealdtech/go-ens/v3/contracts/reverseresolver"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	REVERSE_REG_V2_ADDR         = "0xA2C122BE93b0074270ebeE7f6b7292C7deB45047"
	REVERSE_REG_V1_ADDR         = "0x5fBb459C49BB06083C33109fA4f14810eC2Cf358"
)
var (
	RPC_URL =			os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient				*ethclient.Client
	rpcclient			*rpc.Client
	Info                *log.Logger
	storage             *SQLStorage
	market_order_id		int64 = 0
)
func ens_direct_reverse_lookup(ctrct_addr common.Address,account_addr common.Address) string {

	rreg,err := reverseresolver.NewContract(ctrct_addr,eclient)
	if err != nil {
		fmt.Printf("Error on NewContract(): %v\n",err)
		os.Exit(1)
	}
	node,err := NameHash(fmt.Sprintf("%s.addr.reverse", account_addr.Hex()[2:]))
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	var copts = new(bind.CallOpts)
	name,err := rreg.Name(copts,node)
	if err != nil {
		fmt.Printf("Error on Name(): %+v\n",err)
		os.Exit(1)
	}
	return name
}
func main() {

	rpcclient, err := rpc.DialContext(context.Background(), RPC_URL)
	if err != nil {
		fmt.Printf("Error: %v\n",err)
		os.Exit(1)
	}
	eclient = ethclient.NewClient(rpcclient)

	Info = log.New(os.Stdout,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)
	storage = Connect_to_storage(&market_order_id,Info)

	new_resolver := common.HexToAddress(REVERSE_REG_V2_ADDR)
	old_resolver := common.HexToAddress(REVERSE_REG_V1_ADDR)

	addresses := storage.Get_addrs_with_reverse_name()
	for _,addr_str := range addresses {
		fmt.Printf("\n")
		addr := common.HexToAddress(addr_str)
		//	node,err := NameHash(fmt.Sprintf("%s.addr.reverse", address.Hex()[2:]))
		//	if err != nil {
		//		fmt.Printf("Error: %v\n",err)
		//		os.Exit(1)
		//	}
		name1 := ens_direct_reverse_lookup(new_resolver,addr)
		name2 := ens_direct_reverse_lookup(old_resolver,addr)

		len_name1 := len(name1)
		len_name2 := len(name2)
		var status_str string = "failure"
		if (len_name1 > 0) || (len_name2 > 0 ) {
			status_str = "success"
		}
		fmt.Printf(
			"status %v: %v:    %2v (old)  %2v (new) :  %24v %24v\n",
			status_str,addr_str,len_name1,len_name2,name1,name2,
		)
		domain, err := ens.ReverseResolve(eclient, addr)
		if err != nil {
			fmt.Printf("\tAddr %48v\tError: %v\n",addr_str,err)
			continue
		}
		fmt.Printf("\tAddr: %48v\t%v\n",addr_str,domain)
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
		fmt.Printf("\tgo-ens: Name found: label %v , word %v\n",label_str,word)
		exists := storage.Label_exists_in_ens_labels(label_str)
		if exists {
			fmt.Printf("\tThis label already exists in 'ens_labels', not inserting\n")
		}
		node,err := NameHash(fmt.Sprintf("%s.addr.reverse", addr.Hex()[2:]))
		if err != nil {
			fmt.Printf("Error: %v\n",err)
			os.Exit(1)
		}
		node_str := hex.EncodeToString(node[:])
		fmt.Printf("\tNode: %v\n",node_str)
		exists = storage.Reverse_lookup_registration_exists(addr_str,node_str)
		fmt.Printf("\tnode exists in DB: %v\n",exists)
		if !exists {
			fmt.Printf("\tINSERT: word %v , label %v can be inserted\n",word,label_str)
		}
	}
}
