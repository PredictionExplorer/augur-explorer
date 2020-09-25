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
	"bytes"
	//"encoding/hex"

	"github.com/0xProject/0x-mesh/rpc"
	"github.com/0xProject/0x-mesh/zeroex"
	"github.com/0xProject/0x-mesh/common/types"
	"github.com/plaid/go-envvar/envvar"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/dbs"

)
const (
	DEFAULT_SYNC_INTERVAL_SECS int64 = 60*10
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

	caddrs *ContractAddresses
	last_sync_time int64
	adecoder *zeroex.AssetDataDecoder

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
func fetch_and_sync_orders() {

	ohash_map := make(map[string]struct{})
	var augur_count int = 0
	var insert_count int = 0
	var update_count int = 0
	var deleted_count int = 0
	var orders_total int = 0
	var anomalies_count int = 0
	var page_size int = 256
	var page_num int = 0
	var done bool = false
	for !done {
		orders2sync,err := rpcclient.GetOrders(page_num,page_size,"")
		if err == nil {
			acnt,icnt,ucnt:=sync_orders(orders2sync,&ohash_map)
			insert_count = insert_count + icnt
			update_count = update_count + ucnt
			augur_count = augur_count + acnt
			if len(orders2sync.OrdersInfos) == 0 {
				done = true
			}
			orders_total = orders_total + len(orders2sync.OrdersInfos)
		} else {
			done = true
		}
		page_num++
	}
	db_orders := storage.Get_all_open_order_hashes()
	for i:=0 ; i<len(db_orders) ; i++ {
		_, exists := ohash_map[db_orders[i]];
		if exists {
			// ok
		} else {
			storage.Delete_open_0x_order(db_orders[i],time.Now().Unix(),OOOpCodeSyncProcess)
			Info.Printf(
				"Order %v doesn't exist in Mesh Node, but does exist in the DB. Deleting. (DB_DIRTY_OORDERS)",
				db_orders[i],
			)
			anomalies_count++
			deleted_count++
		}
	}
	Info.Printf(
		"Order sync process complete. scanned: %v orders. " +
		"Augur-related %v. Inserted %v. Updated %v. Deleted %v.\n",
		orders_total,augur_count,insert_count,update_count,deleted_count,
	)
}
func get_ospec(order *zeroex.SignedOrder,order_hash *string) (ZxMeshOrderSpec,error) {

	var copts = new(bind.CallOpts)
	adata,err := zerox_contract.DecodeAssetData(copts,order.MakerAssetData)
	if err!=nil {
		Error.Printf("couldn't decode asset data for order %v : %v\n",*order_hash,err)
		return ZxMeshOrderSpec{},err
	}
	unpacked_id,err := zerox_contract.UnpackTokenId(copts,adata.TokenIds[0])
	if err!=nil {
		Error.Printf("Unpack token id failed for order %v: %v\n",*order_hash,err)
		return ZxMeshOrderSpec{},err
	}
	return unpacked_id,err
}
/*discontinued
func oo_insert(order_hash *string,order *zeroex.SignedOrder,fillable_amount *big.Int,timestamp int64) error {

	ctx := context.Background()
	var copts = new(bind.CallOpts)
	adata,err := zerox_contract.DecodeAssetData(copts,order.MakerAssetData)
	if err!=nil {
		Error.Printf("couldn't decode asset data for order %v : %v\n",*order_hash,err)
		return err
	}
	unpacked_id,err := zerox_contract.UnpackTokenId(copts,adata.TokenIds[0])
	if err!=nil {
		Error.Printf("Unpack token id failed for order %v: %v\n",*order_hash,err)
		return err
	}
	num:=big.NewInt(int64(owner_fld_offset))
	key:=common.BigToHash(num)
	eoa,err := eclient.StorageAt(ctx,order.MakerAddress,key,nil)
	Info.Printf("oo_insert: order_hash=%v\n",*order_hash)
	Info.Printf("oo insert: maker=%v eoa=%v; err=%v\n",order.MakerAddress.String(),hex.EncodeToString(eoa[:]),err)
	var eoa_addr_str string
	if err == nil {
		eoa_addr_str = common.BytesToAddress(eoa[12:]).String()
	} else {
		Info.Printf(
			"ethclient::StorageAt() failed for order %v, maker addr %v: %v. " +
			"Order will be inserted without EOA link. (ETH_STORAGE_FAIL)",
			order.MakerAddress.String(),*order_hash,err,
		)
		return err
	}
	err = storage.Insert_open_order(order_hash,order,fillable_amount,eoa_addr_str,&unpacked_id,OOOpCodeCreated,timestamp)
	return err
}
*/
func order_belongs_to_augur(order *zeroex.SignedOrder) bool {

	// detecting if the order belongs to Augur Platform and it does if MakerAssetData
	// contains ZeroXTrade contract address
	var zeroex_addr_offset int = 4+32*11+4	// offset found after looking into hex data
	if len(order.MakerAssetData) < (zeroex_addr_offset + 32) {
		return false // MakerAssetData is too small, this is not Augur's order
	}
	from_offset := zeroex_addr_offset + 12	// real start of the address within big.Int
	to_offset := from_offset + 20
	possible_zeroex_addr_bytes:=order.MakerAssetData[from_offset:to_offset]
	if !bytes.Equal(possible_zeroex_addr_bytes,caddrs.ZeroxTrade.Bytes()) {
		return false
	}
	var order_adata zeroex.MultiAssetData
	err := adecoder.Decode(order.MakerAssetData, &order_adata)
	if err!=nil {
		Info.Printf("Assed data decode error: %v\n",err)
		return false
	}
	return true
}
func sync_orders(response *types.GetOrdersResponse,ohash_map *map[string]struct{}) (int,int,int) {
	// routine to synchronize orders on 0x Mesh Network with the table `oorders` in postgres
	//	Executed on startup and every 10 minutes

	var insert_count int = 0
	var update_count int = 0
	var augur_count int = 0

	for i:=0 ; i<len(response.OrdersInfos); i++ {
		order_info := response.OrdersInfos[i]
		if order_info != nil {
			if !order_belongs_to_augur(order_info.SignedOrder) {
				continue
			}
			augur_count++
			order_hash := order_info.OrderHash.String()
			var empty struct{}
			(*ohash_map)[order_hash]=empty
			time_stamp:=response.SnapshotTimestamp.Unix()
			var new_timestamp int64
			new_timestamp = order_info.SignedOrder.Salt.Int64()/1000 // Salt usually contains timestamp
			if new_timestamp > 1595894451	 { // 28 July (Augur v2 release date)
				time_stamp = new_timestamp
			}
			ospec,err := get_ospec(order_info.SignedOrder,&order_hash)
			if err!=nil {
				Info.Printf("Error decoding market data: %v\n",err)
				Error.Printf("Error decoding market data: %v\n",err)
			} else {
				DumpOrderSpec(Info,&ospec)
				storage.Insert_0x_mesh_order_event(time_stamp,order_info,&ospec,nil,MeshEvtAdded)
			}
			/*discontinued
			amount := order_info.FillableTakerAssetAmount
			retval,bad_amount := storage.Update_oo_fillable_amount(order_hash,amount,order_info.SignedOrder)
			if retval == 2 { // order doesn't exist
				err := oo_insert(&order_hash,order_info.SignedOrder,order_info.FillableTakerAssetAmount,0)
				if err!=nil {
					// nothing
					Info.Printf("Error inserting open order %v: %v\n",order_hash,err)
					Error.Printf("Error inserting open order %v: %v\n",order_hash,err)
				} else {
					insert_count++
					Info.Printf("Inserted open order %v\n",order_hash)
				}
			}
			if retval == 1 {
				Info.Printf(
					"Order %v had incorrect amount, fixed. (bad amount=%v, good amount=%v) (BAD_AMOUNT)",
					order_hash,bad_amount,amount.String(),
				)
				update_count++
			}
			*/
		}
	}
	return augur_count,insert_count,update_count
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
	caddrs = &caddrs_obj
	if caddrs.ChainId == 0 {
		log.Fatalf("ChainID = 0, db is not initialized")
	}
	adecoder = zeroex.NewAssetDataDecoder()

	eclient, err = ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	Info.Printf("ZeroX contract = %v\n",caddrs.ZeroxTrade.String())
	zerox_contract, err = NewZeroX(
			common.HexToAddress(
				caddrs.ZeroxTrade.String(),
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
	fetch_and_sync_orders()
	last_sync_time=time.Now().Unix()
	//var copts = new(bind.CallOpts)
	for {
		select {
		case orderEvents := <-orderEventsChan:
			for _, orderEvent := range orderEvents {
				if !order_belongs_to_augur(orderEvent.SignedOrder) {
					//Info.Printf("Event listener, skipped non-augur: %v\n",orderEvent.OrderHash.String())
					continue
				}
				//order_hash:=orderEvent.OrderHash.String()
				Info.Printf("--------------------------------------------------\n")
				Info.Printf("Order event arrived in state=%+v:\n",orderEvent.EndState)
	//			Info.Printf("Order Hash: %v\n",order_hash)
	//			Info.Printf("FillableTakerAssetAmount: %v\n",orderEvent.FillableTakerAssetAmount)
	//			Info.Printf("Timestamp: %v\n",orderEvent.Timestamp)
				// store the event in the DB
				var order_info types.OrderInfo
				order_info.OrderHash.SetBytes(orderEvent.OrderHash.Bytes())
				order_info.SignedOrder = orderEvent.SignedOrder
				order_info.FillableTakerAssetAmount = new(big.Int)
				order_info.FillableTakerAssetAmount.Set(orderEvent.FillableTakerAssetAmount)
				event_code := Get_mesh_event_code(orderEvent.EndState)
				order_hash := orderEvent.OrderHash.String()
				ospec,err := get_ospec(order_info.SignedOrder,&order_hash)
				if err!=nil {
					Info.Printf("Error decoding market data: %v\n",err)
					Error.Printf("Error decoding market data: %v\n",err)
					continue
				}
				Dump_0x_mesh_order(Info,&order_info)
				DumpOrderSpec(Info,&ospec)
				switch orderEvent.EndState {
				case zeroex.ESOrderAdded,
					zeroex.ESOrderExpired,
					zeroex.ESOrderFillabilityIncreased,
					zeroex.ESOrderBecameUnfunded,
					zeroex.ESStoppedWatching,
					zeroex.ESOrderUnexpired:
					storage.Insert_0x_mesh_order_event(
						orderEvent.Timestamp.Unix(),
						&order_info,
						&ospec,
						nil,//amount_filled
						event_code,
					)
				}
/* discontinued
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
						err:=oo_insert(&order_hash,orderEvent.SignedOrder,orderEvent.FillableTakerAssetAmount,orderEvent.Timestamp.Unix())
						if err!=nil {
							Info.Printf("Error inserting order %v: %v\n",order_hash,err)
						}
					case zeroex.ESOrderExpired:
						storage.Delete_open_0x_order(orderEvent.OrderHash.String(),OOOpCodeExpired)
					case zeroex.ESOrderCancelled:
						storage.Delete_open_0x_order(orderEvent.OrderHash.String(),OOOpCodeCancelledByUser)
					case zeroex.ESOrderFullyFilled:
						// FULLY FILLED event: quantity of the order matches filling quantity
						storage.Delete_open_0x_order(orderEvent.OrderHash.String(),OOOpCodeNone)
					case zeroex.ESOrderFilled:
						// FILLED event: partial order fill
						storage.Update_0x_order_on_partial_fill(orderEvent)
					// the following are rare events, so we don't implement them, just do a resync
					case zeroex.ESOrderFillabilityIncreased,
						zeroex.ESOrderBecameUnfunded,
						zeroex.ESStoppedWatching,
						zeroex.ESOrderUnexpired:
						// do a re-sync
						fetch_and_sync_orders()
				}
*/
			}
			cur_time := time.Now().Unix()
			time_diff := cur_time - last_sync_time
			if time_diff >= DEFAULT_SYNC_INTERVAL_SECS {
				fetch_and_sync_orders()
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
