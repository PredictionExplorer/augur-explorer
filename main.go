// Augur ETL: Converts Augur Data to SQL database
package main

import (
	"fmt"
	"context"
	"log"
	"math/big"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi"
)
const (
	// ToDo: get these signatures from the abi files
	MARKET_CREATED = "ea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1"
	MARKET_OI_CHANGED = "213a05b9ad8567c2f8fa868e7375e5bf30e69add0dbb5913ca8a3e58c815c268"
	MARKET_ORDER = "9bab1368a1ed530afaad9c630ba75e6a5c1efa9f6af0139d6cda2b6af6aa801e"
	MARKET_FINALIZED = "6d39632c2dc10305bf5771cfff4af1851f07c03ea27b821cad382466bdf7a21f"
	INITIAL_REPORT_SUBMITTED = "c3ebb227c22e7644e9bef8822009f746a72c86f239760124d67fdc2c302b3115"
	MARKET_VOLUME_CHANGED = "e9f0af820300e73bae76c8e76943abe7fbb4224b49cb133e2dadc6f71acf6370"
)
var (
	// these evt_ variables are here for speed to avoid calculation of Keccak256
	//		on each bytes.Compare() operation
	evt_market_created,_ = hex.DecodeString(MARKET_CREATED)
	evt_market_oi_changed,_ = hex.DecodeString(MARKET_OI_CHANGED)
	evt_market_order,_ = hex.DecodeString(MARKET_ORDER)
	evt_market_finalized,_ = hex.DecodeString(MARKET_FINALIZED)
	evt_initial_report_submitted,_ = hex.DecodeString(INITIAL_REPORT_SUBMITTED)
	evt_market_volume_changed,_ = hex.DecodeString(MARKET_VOLUME_CHANGED)

	storage *SQLStorage
	augur_abi *abi.ABI
	trading_abi *abi.ABI
)
func main() {
	//client, err := ethclient.Dial("http://:::8545")
	client, err := ethclient.Dial("http://192.168.1.102:18545")
	if err != nil {
		log.Fatal(err)
	}

	storage = connect_to_storage()

	augur_init()
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
		block, _ := client.BlockByNumber(ctx,big_bnum)
		if block != nil {
			fmt.Printf("stats_block:%v gas limit = %d ,gas used = %d\n",block.Number(),block.GasLimit,block.GasUsed)
			num_transactions, err := client.TransactionCount(ctx,block.Hash())
			if err != nil {
				fmt.Printf("block error: %v \n",err)
			} else {
				if num_transactions > 0 {
					fmt.Printf("block: %v %v %v transactions\n",block.Hash().String(),block.Number(),num_transactions)
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
								sequencer := new(EventSequencer)
								num_logs := len(rcpt.Logs)
								for i:=0 ; i<num_logs ; i++ {
									fmt.Printf(
										"\t\t\tlog %v for contract %v (%v of %v items)\n",
										rcpt.Logs[i].Topics[0].String(),rcpt.Logs[i].Address.String(),(i+1),len(rcpt.Logs))
									sequencer.append_event(rcpt.Logs[i])
								}
								ordered_list := sequencer.get_ordered_event_list()
								num_logs = len(ordered_list)
								for i:=0 ; i < num_logs ; i++ {
									fmt.Printf(
										"\t\t\tprocessing log with sig %v for contract %v\n",
										ordered_list[i].Topics[0].String(),
										ordered_list[i].Address.String())
									process_event(ordered_list[i])
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
