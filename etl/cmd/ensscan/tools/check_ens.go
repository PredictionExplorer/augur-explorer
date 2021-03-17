package main

import (
	"os"
	"fmt"
	"context"
	"bytes"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/wealdtech/go-ens/v3"
	"github.com/wealdtech/go-ens/v3/contracts/resolver"

	. "github.com/PredictionExplorer/augur-explorer/dbs"
)
const (
	REVERSE_REG_V2_ADDR         = "0xA2C122BE93b0074270ebeE7f6b7292C7deB45047"
	REVERSE_REG_V1_ADDR         = "0x5fBb459C49BB06083C33109fA4f14810eC2Cf358"
)
var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient *ethclient.Client
	rpcclient *rpc.Client

	Info				*log.Logger
	storage				*SQLStorage
	market_order_id int64 = 0
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

	var zeroaddr = common.HexToAddress("00")
	cur_id:=int64(0)
	if len(os.Args) > 1 {
		cur_id_converted,err:=strconv.Atoi(os.Args[1])
		if err!=nil {
			fmt.Printf("invalid starting id %v: %v\n",os.Args[1],err)
			os.Exit(1)
		}
		cur_id = int64(cur_id_converted)
	}
	lot_size:=int64(50)
	for {
		lot := storage.Get_node_lot(cur_id,lot_size)
		if len(lot) == 0 {
			fmt.Printf("Process finished")
			os.Exit(0)
		}
		for i:=0 ; i<len(lot); i++ {
			node_entry := lot[i]
			owner_addr,assigned_addr,err := storage.Get_last_owner_addr(node_entry.FQDN)
			if err!=nil {
				fmt.Printf("Error at node %v: %v\n",node_entry.FQDN,err)
				os.Exit(1)
			}
			owner := common.HexToAddress(owner_addr)
			assigned := common.HexToAddress(assigned_addr)
			
			nodehash:= common.HexToHash("0x"+node_entry.FQDN)
			fmt.Printf("%v : %v : id=%v\n",node_entry.FQDN_Words,node_entry.FQDN,node_entry.Id)
			regstr, err := ens.NewRegistry(eclient)
			if err!=nil {
				fmt.Printf("NewRegistry error: %v\n",err)
				os.Exit(1)
			}
			owner_addr_ens,err := regstr.Contract.Owner(nil,nodehash)
			if err!=nil {
				fmt.Printf("Owner error: %v\n",err)
				os.Exit(1)
			}
			if bytes.Equal(owner_addr_ens.Bytes(),zeroaddr.Bytes()) {
				fmt.Printf("Node %v unregistered\n",node_entry.FQDN)
				continue;
				//os.Exit(1)
			}
			resolver_addr, err := regstr.Contract.Resolver(nil,nodehash)
			if err!=nil {
				fmt.Printf("Resolver() call on node %v failed: %v (resolver addr is %v)",node_entry.FQDN,err,resolver_addr.String())
				os.Exit(1)
			}
			if bytes.Equal(resolver_addr.Bytes(),zeroaddr.Bytes()) {
				fmt.Printf("Node %v has resolver at 0x0 addr\n",node_entry.FQDN)
				continue;
				//os.Exit(1)
			}
			resolver_ctrct,err := resolver.NewContract(resolver_addr,eclient)
			if err!=nil {
				fmt.Printf("NewContract() for resolver failed (node=%v): %v\n",node_entry.FQDN,err)
				os.Exit(1)
			}
			looked_up,err := resolver_ctrct.Addr(nil,nodehash)
			if err != nil {
				fmt.Printf("resolver's Address() call on node %v failed: %v (resolver addr is %v)\n",node_entry.FQDN,err,resolver_addr.String())
				os.Exit(1)
			}
			fmt.Printf("\tlookup at Resolver: node addr = %v\n",looked_up.String())
/*
			looked_up,err := ens.Resolve(eclient, node_entry.FQDN)
			if err!=nil {
				fmt.Printf("ENS lookup error for node %v : %v\n",node_entry.FQDN,err)
				os.Exit(1)
			}*/
			if len(assigned_addr) > 0 {
				if !bytes.Equal(assigned.Bytes(),looked_up[:]) {
					fmt.Printf(
						"Resolution for node %v is incorrect!\n\tmust be %v\n\towner : %v\n\tassigned : %v\n",
						node_entry.FQDN,looked_up.String(),owner_addr,assigned_addr,
					)
					//os.Exit(1)
				} else {
					fmt.Printf("\tnode ok: addr = %v\n",assigned_addr)
				}
			} else {
				if len(owner_addr) >0 {
					if !bytes.Equal(owner.Bytes(),looked_up[:]) {
						fmt.Printf(
							"Resolution for node %v is incorrect\n\tmust be %v\n\towner : %v\n\tassigned : %v\n",
							node_entry.FQDN,looked_up.String(),owner_addr,assigned_addr,
						)
						//os.Exit(1)
					} else {
						fmt.Printf("Node %v : ok\n",node_entry.FQDN)
					}
				} else {
					fmt.Printf("Resolution incorrect: owner address wasn't found in the DB for node %v\n",node_entry.FQDN)
					//os.Exit(1)
				}
			}
		}
		cur_id = cur_id + lot_size
	}
}
