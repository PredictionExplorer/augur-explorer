// +build !js

// demo/add_order is a short program that adds an order to 0x Mesh via RPC
package main

import (
	"context"
	"math/big"
	"os"
	"io"
	"runtime"
	"fmt"
	"encoding/hex"
//	"strings"

	"github.com/0xProject/0x-mesh/rpc"
	"github.com/0xProject/0x-mesh/zeroex"
	"github.com/plaid/go-envvar/envvar"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
)
type clientEnvVars struct {
	// RPCAddress is the address of the 0x Mesh node to communicate with.
	WSRPCAddress string `envvar:"WS_RPC_ADDR"`
}
func Fatalf(format string, args ...interface{}) {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}
	fmt.Fprintf(w, "Fatal: "+format+"\n", args...)
	os.Exit(1)
}

func main() {

	if len(RPC_URL) == 0 {
		Fatalf("Configuration error: RPC URL of Ethereum node is not set."+
				" Please set AUGUR_ETH_NODE_RPC environment variable")
	}
	ethclient, err := ethclient.Dial(RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	zerox_contract, err := NewZeroX(common.HexToAddress("0x6749E370e7B1955FFa924F4f75f5F12653C7512C"), ethclient)
	if err != nil {
		log.Fatalf("Failed to instantiate a ZeroX contract: %v", err)
	}
	awallet_contract, err := NewAugurWallet(common.HexToAddress("0xcc165aa8353BcCe882C14677aD78B20b55C0A5ED"), ethclient)
	if err != nil {
		log.Fatalf("Failed to instantiate a AugurWalletRegistry contract: %v", err)
	}
	storage := connect_to_storage()
	log.SetFormatter(&log.JSONFormatter{})

	env := clientEnvVars{}
	if err := envvar.Parse(&env); err != nil {
		panic(err)
	}

	client, err := rpc.NewClient(env.WSRPCAddress)
	if err != nil {
		log.WithError(err).Fatal("could not create client")
	}

	ctx := context.Background()
	orderEventsChan := make(chan []*zeroex.OrderEvent, 8000)
	clientSubscription, err := client.SubscribeToOrders(ctx, orderEventsChan)
	if err != nil {
		log.WithError(err).Fatal("Couldn't set up OrderStream subscription")
	}
	defer clientSubscription.Unsubscribe()

	var copts = new(bind.CallOpts)
	copts.Pending = true
	for {
		select {
		case orderEvents := <-orderEventsChan:
			for _, orderEvent := range orderEvents {
				/*
				log.WithFields(log.Fields{
					"event": orderEvent,
				}).Printf("received order event")
				*/
				adata,err := zerox_contract.DecodeAssetData(copts,orderEvent.SignedOrder.Order.MakerAssetData)
				if err!=nil {
					fmt.Printf("couldn't decode asset data: %v\n",err)
				} else {
					//fmt.Printf("asset data = %+v\n",adata)
					//fmt.Printf("Token Address = %v\n",adata.TokenAddress.String())
					//token_ids := bigint_ptr_slice_to_str(&adata.TokenIds,",")
					//fmt.Printf("Token IDs: %v\n",token_ids)
					//token_values := bigint_ptr_slice_to_str(&adata.TokenValues,",")
					//fmt.Printf("Token values: %v\n",token_values)
					unpacked_id,err := zerox_contract.UnpackTokenId(copts,adata.TokenIds[0])
					if err!=nil {
						fmt.Printf("Unpack token id failed: %v\n",err)
					} else {
						//fmt.Printf("Token ID data = %+v\n",unpacked_id)
						//fmt.Printf("\tMarket = %v\n",unpacked_id.Market.String())
						num:=big.NewInt(int64(1))	// 1 is the offset at Storage where EOA is stored
						key:=common.BigToHash(num)
						eoa,err := ethclient.StorageAt(ctx,orderEvent.SignedOrder.Order.MakerAddress,key,nil)
						var eoa_addr_str string
						if err == nil {
							eoa_addr_str = common.BytesToAddress(eoa[12:]).String()
						}
						storage.insert_open_order(&orderEvent.SignedOrder.Order,eoa_addr_str,&unpacked_id)
					}
				}

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
				//fmt.Printf("Encoded order len = %v\n",len(ordr_blob))
				//fmt.Printf("Encoded order: %v\n",hex.EncodeToString(ordr_blob))
				msg_hash,err := awallet_contract.GetMessageHash(copts,ordr_blob[:])
				if err != nil {
					fmt.Printf("getMessageHash() produced error: %v\n",err)
				}
				//fmt.Printf("msg hash = %v\n",hex.EncodeToString(msg_hash[:]))
				//fmt.Printf("signature = %v\n",hex.EncodeToString(orderEvent.SignedOrder.Signature))
				//fmt.Printf("sig type = %d\n",orderEvent.SignedOrder.Signature[65])
				public_key, err := crypto.Ecrecover(msg_hash[:], orderEvent.SignedOrder.Signature[0:65])
				if err != nil {
					fmt.Printf("couldn't extract public key: %v\n",err)
				} else {
					fmt.Printf("public key : %v\n",hex.EncodeToString(public_key))
				}
			}
		case err := <-clientSubscription.Err():
			log.Fatal(err)
		}
	}
}
/*
Pending for deletion
				fmt.Printf("getting storage of ctrct %v\n",orderEvent.SignedOrder.Order.MakerAddress.String())
				for i:=0 ; i< 6 ; i++  {
					num:=big.NewInt(int64(i))
					key:=common.BigToHash(num)
					var1,err := ethclient.StorageAt(ctx,orderEvent.SignedOrder.Order.MakerAddress,key,nil)
					if err != nil {
						fmt.Printf("getting storage %v: %v\n",i,err)
					} else {
						fmt.Printf("storage value key=%v is: %v\n",i,hex.EncodeToString(var1))
					}
				}
*/
