// +build !js

// demo/add_order is a short program that adds an order to 0x Mesh via RPC
package main

import (
	"context"
	"math/big"
	"os"
	"log"
	"fmt"
	"time"
	//"encoding/hex"

	"github.com/0xProject/0x-mesh/rpc"
	"github.com/0xProject/0x-mesh/zeroex"
	"github.com/0xProject/0x-mesh/common/types"
	"github.com/plaid/go-envvar/envvar"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/accounts/abi"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"

)
const (
	DEFAULT_SYNC_INTERVAL_SECS int64 = 60*1
)
var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	ctrct_zerox *ZeroX
	market_order_id int64 = 0
	fill_order_id int64 = 0
	storage *SQLStorage
	owner_fld_offset int64 = 2	// offset to AugurContract::owner field obtained with eth_getStorage()
	zerox_contract *ZeroX

	eclient *ethclient.Client
	rpcclient *rpc.Client

	last_sync_time int64

	Error   *log.Logger
	Info    *log.Logger
)
type clientEnvVars struct {
	// RPCAddress is the address of the 0x Mesh node to communicate with.
	WSRPCAddress string `envvar:"WS_RPC_ADDR"`
}
func Dump_stats(s *types.Stats,output *log.Logger) {
	output.Printf("Version: %v\n",s.Version)
	output.Printf("PubSubTopic: %v\n",s.PubSubTopic)
	output.Printf("Rendezvous: %v\n",s.Rendezvous)
	output.Printf("Secondary rendezvous: %+v\n",s.SecondaryRendezvous)
	output.Printf("PeerId: %v\n",s.PeerID)
	output.Printf("Ethereum ChainID: %v\n",s.EthereumChainID)
	output.Printf("LatestBlock: %v\n",s.LatestBlock.Number)
	output.Printf("NumPeers: %v\n",s.NumPeers)
	output.Printf("NumOrders: %v\n",s.NumOrders)
	output.Printf("NumOrdersIncludingRemoved: %v\n",s.NumOrdersIncludingRemoved)
	output.Printf("NumPinnedOrders: %v\b",s.NumPinnedOrders)
	output.Printf("MaxExpirationTime: %v\n",s.MaxExpirationTime)
	output.Printf("StartOfCurrentUTCDay: %v\n",s.StartOfCurrentUTCDay)
	output.Printf("EthRPCRequestsSentInCurrentUTCDay: %v\n",s.EthRPCRequestsSentInCurrentUTCDay)
	output.Printf("EthRPCRateLimitExpiredRequests: %v\n",s.EthRPCRateLimitExpiredRequests)
}
func fetch_orders() (*types.GetOrdersResponse, error) {
	orders2sync,err := rpcclient.GetOrders(0,256*1024,"")
	return orders2sync,err
}
func oo_insert(order_hash *string,order *zeroex.SignedOrder,timestamp int64) {

	ctx := context.Background()
	var copts = new(bind.CallOpts)
	adata,err := zerox_contract.DecodeAssetData(copts,order.MakerAssetData)
	if err!=nil {
		Error.Printf("couldn't decode asset data for order %v : %v\n",order_hash,err)
		return
	}
	unpacked_id,err := zerox_contract.UnpackTokenId(copts,adata.TokenIds[0])
	if err!=nil {
		Error.Printf("Unpack token id failed for order %v: %v\n",order_hash,err)
		return
	}
	num:=big.NewInt(int64(owner_fld_offset))
	key:=common.BigToHash(num)
	eoa,err := eclient.StorageAt(ctx,order.MakerAddress,key,nil)
	var eoa_addr_str string
	if err == nil {
		eoa_addr_str = common.BytesToAddress(eoa[12:]).String()
	} else {
		Info.Printf(
			"ethclient::StorageAt() failed for order %v, maker addr %v: %v. " +
			"Order will be inserted without EOA link. (ETH_STORAGE_FAIL)",
			order.MakerAddress.String(),*order_hash,err,
		)
	}
	storage.Insert_open_order(order_hash,order,eoa_addr_str,&unpacked_id,timestamp)
}
func sync_orders(response *types.GetOrdersResponse) {
	// routine to synchronize orders on 0x Mesh Network with the table `oorders` in postgres
	//	Executed on startup and every 10 minutes

	Info.Printf("Syncing orders with Postgres DB: num_orders=%v\n",len(response.OrdersInfos))
	var anomalies_count int = 0

	ohash_map := make(map[string]struct{})

	for i:=0 ; i<len(response.OrdersInfos); i++ {
		order_info := response.OrdersInfos[i]
		if order_info != nil {
			order_hash := order_info.OrderHash.String()
			var empty struct{}
			ohash_map[order_hash]=empty
			amount := order_info.FillableTakerAssetAmount
			retval,bad_amount := storage.Update_open_order(order_hash,amount,order_info.SignedOrder)
			if retval == 2 { // order doesn't exist
				oo_insert(&order_hash,order_info.SignedOrder,0)
				Info.Printf("Inserted open order %v\n",order_hash)
				anomalies_count++
			}
			if retval == 1 {
				Info.Printf(
					"Order %v had incorrect amount, fixed. (bad amount=%v, good amount=%v) (BAD_AMOUNT)",
					bad_amount,amount.String(),
				)
				anomalies_count++
			}
		}
	}
	db_orders := storage.Get_all_open_order_hashes()
	for i:=0 ; i<len(db_orders) ; i++ {
		_, exists := ohash_map[db_orders[i]];
		if exists {
			// ok
		} else {
			storage.Delete_open_0x_order(db_orders[i])
			Info.Printf(
				"Order %v doesn't exist in Mesh Node, but does exist in the DB. Deleting. (DB_DIRTY_OORDERS)",
				db_orders[i],
			)
			anomalies_count++
		}
	}
	var anomalies_str string = ""
	if anomalies_count > 0 {
		anomalies_str = fmt.Sprintf(" Anomalies: %v (ANOMALIES_FOUND)",anomalies_count)
	}
	Info.Printf("Order sync process complete.%v",anomalies_str)
}
func main() {

	log_dir:=fmt.Sprintf("%v/%v",os.Getenv("HOME"),DEFAULT_LOG_DIR)
	os.MkdirAll(log_dir, os.ModePerm)

	fname:=fmt.Sprintf("%v/mesh_info.log",log_dir)
	logfile, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Info = log.New(logfile,"INFO: ",log.Ldate|log.Ltime/*|log.Lshortfile*/)

	fname = fmt.Sprintf("%v/mesh_error.log",log_dir)
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Error = log.New(logfile,"ERROR: ",log.Ldate|log.Ltime/*|log.Lshortfile*/)

	if len(RPC_URL) == 0 {
		Fatalf("Configuration error: RPC URL of Ethereum node is not set."+
				" Please set AUGUR_ETH_NODE_RPC environment variable")
	}

	storage = Connect_to_storage(&market_order_id,Info)
	caddrs_obj,err := storage.Get_contract_addresses()
	if err != nil {
		log.Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}

	eclient, err = ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	Info.Printf("ZeroX contract = %v\n",caddrs_obj.Zerox.String())
	zerox_contract, err = NewZeroX(
			common.HexToAddress(
				caddrs_obj.Zerox.String(),
			),
			eclient,
	)
	if err != nil {
		log.Fatalf("Failed to instantiate a ZeroX contract: %v", err)
	}

	env := clientEnvVars{}
	if err := envvar.Parse(&env); err != nil {
		panic(err)
	}

	rpcclient, err = rpc.NewClient(env.WSRPCAddress)
	if err != nil {
		log.Fatal("could not create client %v",err)
	}

	ctx := context.Background()

	stats,err := rpcclient.GetStats()
	if err == nil {
		Info.Printf("Connected to server %v\n",env.WSRPCAddress)
		//fmt.Printf("0x Mesh server stats: %+v\n",stats)
		Info.Printf("Dumping server statistics...\n")
		Dump_stats(stats,Info)
	} else {
		fmt.Printf("Connection error: %v\n",err)
		os.Exit(1)
	}
	orderEventsChan := make(chan []*zeroex.OrderEvent, 8000)
	clientSubscription, err := rpcclient.SubscribeToOrders(ctx, orderEventsChan)
	if err != nil {
		log.Fatal("Couldn't set up OrderStream subscription")
	}
	defer clientSubscription.Unsubscribe()
	Info.Printf("Subscribed to events successfully..., 0x Mesh Listener started.\n")
	orders2sync,err := fetch_orders()
	if err == nil {
		sync_orders(orders2sync)
	} else {
		Error.Printf("Order sync at startup failed with error: %v\n",err)
	}
	last_sync_time=time.Now().Unix()
	var copts = new(bind.CallOpts)
	for {
		select {
		case orderEvents := <-orderEventsChan:
			for _, orderEvent := range orderEvents {
				order_hash:=orderEvent.OrderHash.String()
				Info.Printf("--------------------------------------------------\n")
				Info.Printf("Order event arrived in state=%+v:\n",orderEvent.EndState)
				Info.Printf("Order Hash: %v\n",order_hash)
				Info.Printf("FillableTakerAssetAmount: %v\n",orderEvent.FillableTakerAssetAmount)
				Info.Printf("Timestamp: %v\n",orderEvent.Timestamp)
				adata,err := zerox_contract.DecodeAssetData(copts,orderEvent.SignedOrder.Order.MakerAssetData)
				if err==nil {
					unpacked_position,err := zerox_contract.UnpackTokenId(copts,adata.TokenIds[0])
					if err!=nil {
						Info.Printf("Market: %v\n",unpacked_position.Market.String())
						Info.Printf("Outcome: %v\n",unpacked_position.Outcome)
						Info.Printf("Type: %v\n",unpacked_position.Type)
						Info.Printf("Price: %v\n",unpacked_position.Price)
					} else {
						Error.Printf("Couldn't decode position data for order %v: %v\n",order_hash,err)
					}
				} else {
					Error.Printf("Failed to decode market data for order %v: %v\n",order_hash,err)
				}
				switch orderEvent.EndState {
					case zeroex.ESOrderAdded:
						oo_insert(&order_hash,orderEvent.SignedOrder,orderEvent.Timestamp.Unix())
					case zeroex.ESOrderExpired,
						zeroex.ESOrderCancelled:
						storage.Delete_open_0x_order(orderEvent.OrderHash.String())
					case zeroex.ESOrderFullyFilled:
						// FULLY FILLED event: quantity of the order matches filling quantity
						storage.Delete_open_0x_order(orderEvent.OrderHash.String())
					case zeroex.ESOrderFilled:
						// FILLED event: partial order fill
						storage.Update_0x_order_on_partial_fill(orderEvent)

					// the following are rare events, so we don't implement them, just do a resync
					case zeroex.ESOrderFillabilityIncreased,
						zeroex.ESOrderBecameUnfunded,
						zeroex.ESStoppedWatching,
						zeroex.ESOrderUnexpired:
						// do a re-sync
						orders2sync,err := fetch_orders()
						if err == nil {
							sync_orders(orders2sync)
						}
				}
			}
			cur_time := time.Now().Unix()
			time_diff := cur_time - last_sync_time
			if time_diff >= DEFAULT_SYNC_INTERVAL_SECS {
				orders2sync,err := fetch_orders()
				if err == nil {
					sync_orders(orders2sync)
				} else {
					Error.Printf("Order sync at startup failed with error: %v\n",err)
				}
				last_sync_time=time.Now().Unix()
			}

		case err := <-clientSubscription.Err():
			log.Fatal(err)
		}
	}
}
/* This code should be used to extract public key from the signature, but it is not working. Todo: fix it
				ohash, err := orderEvent.SignedOrder.Order.ComputeOrderHash()
				if err !=nil {
					fmt.Printf("can't compute order's hash: %v\n",err)
				}
				//fmt.Printf("computed order hash: %v\n",ohash.String())
				xchg_ordr := new(IExchangeOrder)
				xchg_ordr.MakerAddress			= orderEvent.SignedOrder.Order.MakerAddress
				xchg_ordr.TakerAddress			= orderEvent.SignedOrder.Order.TakerAddress
				xchg_ordr.FeeRecipientAddress	= orderEvent.SignedOrder.Order.FeeRecipientAddress
				xchg_ordr.SenderAddress         = orderEvent.SignedOrder.Order.SenderAddress
				xchg_ordr.MakerAssetAmount      = orderEvent.SignedOrder.Order.MakerAssetAmount
				xchg_ordr.TakerAssetAmount      = orderEvent.SignedOrder.Order.TakerAssetAmount
				xchg_ordr.MakerFee              = orderEvent.SignedOrder.Order.MakerFee
				xchg_ordr.TakerFee              = orderEvent.SignedOrder.Order.TakerFee
				xchg_ordr.ExpirationTimeSeconds = orderEvent.SignedOrder.Order.ExpirationTimeSeconds
				xchg_ordr.Salt                  = orderEvent.SignedOrder.Order.Salt
				xchg_ordr.MakerAssetData        = orderEvent.SignedOrder.Order.MakerAssetData
				xchg_ordr.TakerAssetData        = orderEvent.SignedOrder.Order.TakerAssetData
				xchg_ordr.MakerFeeAssetData     = orderEvent.SignedOrder.Order.MakerFeeAssetData
				xchg_ordr.TakerFeeAssetData     = orderEvent.SignedOrder.Order.TakerFeeAssetData

				ordr_blob,err := zerox_contract.EncodeEIP1271OrderWithHash(copts,*xchg_ordr,ohash)
				if err != nil {
					fmt.Printf("EncodeEIP1271 produced error: %v\n",err)
				}
				msg_hash,err := awallet_contract.GetMessageHash(copts,ordr_blob[:])
				if err != nil {
					fmt.Printf("getMessageHash() produced error: %v\n",err)
				}
				public_key, err := crypto.Ecrecover(msg_hash[:], orderEvent.SignedOrder.Signature[0:65])
				if err != nil {
					fmt.Printf("couldn't extract public key: %v\n",err)
				} else {
					fmt.Printf("public key : %v\n",hex.EncodeToString(public_key))
				}
*/
