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
			multiaddr,err := storage.Get_last_name_address(node_entry.FQDN)
			owner_addr := multiaddr.OwnerAddr
			_ = owner_addr
			assigned_addr := multiaddr.AddrChgAddr
			if err!=nil {
				fmt.Printf("SQL error at owner/assigned query for node %v: %v\n",node_entry.FQDN,err)
				os.Exit(1)
			}
			nr_addr,_,err := storage.Get_last_new_resolver_name_addr(node_entry.FQDN)
			if err != nil {
				fmt.Printf("SQL error at resolver/addr for node %v: %v\n",node_entry.FQDN,err)
				os.Exit(1)
			}
			_ = nr_addr
			//name_addr := common.HexToAddress(name_addr_str)
			//owner := common.HexToAddress(owner_addr)
			assigned := common.HexToAddress(assigned_addr)
			_ = assigned
			unregistered,no_resolver,no_address,err := storage.ENS_name_resolution_status(node_entry.FQDN)
			if err!=nil {
				fmt.Printf(
					"Data incorrect: no corresponding ENS name record found for %v : %v\n",
					node_entry.FQDN,err,
				)
			}

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
			fmt.Printf("\t Owner: %v\n",owner_addr_ens.String())
			if bytes.Equal(owner_addr_ens.Bytes(),zeroaddr.Bytes()) {
				if unregistered {
					fmt.Printf("\tmarked as unregistered. ok\n")
					continue;
				}
				fmt.Printf(
					"\tResolution for node %v is incorrect: name should be marked as unregistered in the DB" +
					"(owner address is 0x0)\n",
					node_entry.FQDN,
				)
				continue;
			}
			resolver_addr, err := regstr.Contract.Resolver(nil,nodehash)
			fmt.Printf("\tResolver addr = %v\n",resolver_addr.String())
			if err!=nil {
				fmt.Printf(
					"\tResolver() call on node %v failed: %v (resolver addr is %v)",
					node_entry.FQDN,err,resolver_addr.String(),
				)
				os.Exit(1)
			}
			if bytes.Equal(resolver_addr.Bytes(),zeroaddr.Bytes()) {
				fmt.Printf("\tNode %v has resolver at 0x0 addr\n",node_entry.FQDN)
				if no_resolver {
					fmt.Printf("\tno_resolver set to true, ok\n")
					continue;
				}
				fmt.Printf(
					"\tResolution for node %v is incorrect: name should be market with " +
					"no_resolver = true in the DB\n",
					node_entry.FQDN,
				)
				continue
				//os.Exit(1)
			}
			resolver_ctrct,err := resolver.NewContract(resolver_addr,eclient)
			if err!=nil {
				fmt.Printf("NewContract() for resolver failed (node=%v): %v\n",node_entry.FQDN,err)
				os.Exit(1)
			}
			looked_up,err := resolver_ctrct.Addr(nil,nodehash)
			if err != nil {
				if err.Error() == "no contract code at given address" {
					fmt.Printf("\tnode %v has resolver set to EOA account\n",node_entry.FQDN)
					continue
				}
				fmt.Printf("resolver's Address() call on node %v failed: %v (resolver addr is %v)\n",node_entry.FQDN,err,resolver_addr.String())
				continue
			}
			fmt.Printf(
				"\tlookup at Resolver %v : \tnode addr = %v\n",
				resolver_addr.String(),	looked_up.String(),
			)
			if bytes.Equal(looked_up.Bytes(),zeroaddr.Bytes()) {
				if no_resolver || no_address {
					fmt.Printf("\tname doesn't have resolver/address set. ok\n")
					continue;
				}
				fmt.Printf(
					"\tResolution for node %v is incorrect: looked up addr is 0x0, " +
					"but the DB hasn't it marked as no resolver/address flag\n",
					node_entry.FQDN,
				)
				continue
			}
			db_addr,field_code,err := storage.Resolve_ens_name(node_entry.FQDN)
			if err != nil {
				fmt.Printf("Error at DB: %v\n",err)
				os.Exit(1)
			}
			if looked_up.String()!=db_addr {
				fmt.Printf(
					"\t Resolution for node %v is incorrect, address mismatch\n" +
					"\tmust be %v\n" +
					"\tstored in DB is %v (field code=%v)\n",
					node_entry.FQDN,looked_up.String(),db_addr,field_code,
				)
			}
			/* Temporarily disabled
			if len(assigned_addr) > 0 {
				if !bytes.Equal(assigned.Bytes(),looked_up[:]) {
					fmt.Printf("\tmultiaddr=%v, looked up=%v\n",multiaddr.NewResAddr,looked_up.String())
					if multiaddr.NewResAddr == looked_up.String() {
						// the address was changed by NewResolver event , ok
					} else {
						fmt.Printf(
							"\tResolution for node %v is incorrect, address mismatch!\n"+
							"\tmust be %v\n\towner : %v\n\tassigned : %v\n",
							node_entry.FQDN,looked_up.String(),owner_addr,assigned_addr,
						)
					}
				} else {
					fmt.Printf("\tnode ok: addr = %v\n",assigned_addr)
				}
			} else {
				if multiaddr.NewResAddr == looked_up.String() {
						// the address was changed by NewResolver event , ok
				} else {
					fmt.Printf(
						"\tResolution for node %v is incorrect\n\tmust be %v\n\towner : %v\n\tassigned : %v\n",
						node_entry.FQDN,looked_up.String(),owner_addr,assigned_addr,
					)
				}
			}
			*/
		}
		cur_id = cur_id + lot_size
	}
	fmt.Printf("Exiting\n")
}
