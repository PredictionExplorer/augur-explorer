// Augur ETL: Converts Augur Data to SQL database
package main

import (
	"fmt"
	"context"
	"log"
	"math/big"
//	"bytes"
//	"io/ioutil"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/ethclient"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)
const (
	MARKET_CREATED = "ea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1"
	MARKET_OI_CHANGED = "213a05b9ad8567c2f8fa868e7375e5bf30e69add0dbb5913ca8a3e58c815c268"
	MARKET_ORDER = "9bab1368a1ed530afaad9c630ba75e6a5c1efa9f6af0139d6cda2b6af6aa801e"
)
var (
	evt_market_created,_ = hex.DecodeString(MARKET_CREATED)
	evt_market_oi_changed,_ = hex.DecodeString(MARKET_OI_CHANGED)
	evt_market_order,_ = hex.DecodeString(MARKET_ORDER)
	storage *SQLStorage
	augur_abi *abi.ABI
	trading_abi *abi.ABI
)
func main() {
	client, err := ethclient.Dial("http://:::8545")
	if err != nil {
		log.Fatal(err)
	}

	storage = connect_to_storage()

	augur_init()
	//ctx, _:= context.WithTimeout(context.Background(), 3 * time.Second)
	ctx := context.Background()
	latestBlock, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatal("oops:", err)
	}
	log.Println("latest block:", latestBlock.Number())

	var bnum_high uint64 = latestBlock.Number().Uint64()
	var bnum uint64 = 1;
	for ; bnum<bnum_high; bnum++ {
		big_bnum:=big.NewInt(int64(bnum))
	//	ctx, _:= context.WithTimeout(context.Background(), 3 * time.Second)
		block, _ := client.BlockByNumber(ctx,big_bnum)
		if block != nil {
			num_transactions, err := client.TransactionCount(ctx,block.Hash())
			if err != nil {
				fmt.Printf("block error: %v \n",err)
			} else {
				if num_transactions > 0 {
//					fmt.Printf("block: %v %v %v transactions\n",block.Hash(),block.Number(),num_transactions)
					for tnum:=0 ; tnum < int(num_transactions) ; tnum++ {
						tx , err := client.TransactionInBlock(ctx,block.Hash(),uint(tnum))
						if err != nil {
							fmt.Printf("Error: %v",err)
						} else {
							fmt.Printf("\ttx: %v\n",tx.Hash().String())
							rcpt,err := client.TransactionReceipt(ctx,tx.Hash())
							if err != nil {
								fmt.Printf("Error: %v",err)
							} else {
								num_logs := len(rcpt.Logs)
//								fmt.Printf("\t\tReceipt logs: %v entries",num_logs)
								for i:=0 ; i<num_logs ; i++ {
									fmt.Printf("\t\t\tlog %v for contract %v (%v of %v items)\n",i,rcpt.Logs[i].Address.String(),(i+1),len(rcpt.Logs))
									process_event(rcpt.Logs[i])
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("latest block = %v\n",bnum_high)
}
