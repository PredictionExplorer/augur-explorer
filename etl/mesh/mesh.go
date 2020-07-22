// +build !js

// demo/add_order is a short program that adds an order to 0x Mesh via RPC
package main

import (
	"context"
	"math/big"
	"os"
	"log"
	"fmt"
	"encoding/hex"

	"github.com/0xProject/0x-mesh/rpc"
	"github.com/0xProject/0x-mesh/zeroex"
	"github.com/plaid/go-envvar/envvar"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/accounts/abi"

	. "augur-extractor/primitives"
	. "augur-extractor/dbs"

)

var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	ctrct_zerox *ZeroX
	market_order_id int64 = 0
	fill_order_id int64 = 0

	Error   *log.Logger
	Info    *log.Logger
)
type clientEnvVars struct {
	// RPCAddress is the address of the 0x Mesh node to communicate with.
	WSRPCAddress string `envvar:"WS_RPC_ADDR"`
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
	Info = log.New(logfile,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	fname = fmt.Sprintf("%v/mesh_error.log",log_dir)
	logfile, err = os.OpenFile(fname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!=nil {
		fmt.Printf("Can't start: %v\n",err)
		os.Exit(1)
	}
	Error = log.New(logfile,"ERROR: ",log.Ldate|log.Ltime|log.Lshortfile)

	if len(RPC_URL) == 0 {
		Fatalf("Configuration error: RPC URL of Ethereum node is not set."+
				" Please set AUGUR_ETH_NODE_RPC environment variable")
	}

	storage := Connect_to_storage(&market_order_id,Info)
	caddrs_obj,err := storage.Get_contract_addresses()
	if err != nil {
		log.Fatalf("Can't find contract addresses in 'contract_addresses' table")
	}

	ethclient, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	Info.Printf("ZeroX contract = %v\n",caddrs_obj.Zerox.String())
	zerox_contract, err := NewZeroX(
			common.HexToAddress(
				caddrs_obj.Zerox.String(),
			),
			ethclient,
	)
	if err != nil {
		log.Fatalf("Failed to instantiate a ZeroX contract: %v", err)
	}

	env := clientEnvVars{}
	if err := envvar.Parse(&env); err != nil {
		panic(err)
	}

	client, err := rpc.NewClient(env.WSRPCAddress)
	if err != nil {
		log.Fatal("could not create client %v",err)
	}

	ctx := context.Background()
	orderEventsChan := make(chan []*zeroex.OrderEvent, 8000)
	clientSubscription, err := client.SubscribeToOrders(ctx, orderEventsChan)
	if err != nil {
		log.Fatal("Couldn't set up OrderStream subscription")
	}
	defer clientSubscription.Unsubscribe()

	var copts = new(bind.CallOpts)
	for {
		select {
		case orderEvents := <-orderEventsChan:
			for _, orderEvent := range orderEvents {
				fmt.Printf("Order event arrived in state=%+v:\n",orderEvent.EndState)
				fmt.Printf("Order Hash: %v\n",orderEvent.OrderHash.String())
				fmt.Printf("%+v\n",orderEvent)
				fmt.Println()
				ad:=hex.EncodeToString(orderEvent.SignedOrder.Order.MakerAssetData)
				fmt.Printf("decoding asset data: %v\n",ad)
				adata,err := zerox_contract.DecodeAssetData(copts,orderEvent.SignedOrder.Order.MakerAssetData)
				if err!=nil {
					fmt.Printf("couldn't decode asset data: %v\n",err)
				} else {
					unpacked_id,err := zerox_contract.UnpackTokenId(copts,adata.TokenIds[0])
					if err!=nil {
						fmt.Printf("Unpack token id failed: %v\n",err)
					} else {
						num:=big.NewInt(int64(2))	// 1 is the offset at Storage where EOA is stored
						key:=common.BigToHash(num)
						eoa,err := ethclient.StorageAt(ctx,orderEvent.SignedOrder.Order.MakerAddress,key,nil)
						var eoa_addr_str string
						if err == nil {
							eoa_addr_str = common.BytesToAddress(eoa[12:]).String()
						}
						if orderEvent.EndState == zeroex.ESOrderAdded {
							storage.Insert_open_order(orderEvent,eoa_addr_str,&unpacked_id)
						}
						if orderEvent.EndState == zeroex.ESOrderExpired {
							storage.Delete_open_0x_order(orderEvent.OrderHash.String())
						}
						if orderEvent.EndState == zeroex.ESOrderFullyFilled {
							// FULLY FILLED event: quantity of the order matches filling quantity
							storage.Delete_open_0x_order(orderEvent.OrderHash.String())
						}
						if orderEvent.EndState == zeroex.ESOrderFilled {
							// FILLED event: partial order fill
							storage.Update_0x_order_on_partial_fill(orderEvent)
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
			}
		case err := <-clientSubscription.Err():
			log.Fatal(err)
		}
	}
}
