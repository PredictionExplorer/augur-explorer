// Augur ETL: Converts Augur Data to SQL database
package main

import (
	"fmt"
	"context"
	"log"
	"math/big"
	"bytes"
	"io/ioutil"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)
const (
	MARKET_CREATED = "ea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1"
	MARKET_OI_CHANGED = "213a05b9ad8567c2f8fa868e7375e5bf30e69add0dbb5913ca8a3e58c815c268"
)
var (
	evt_market_created,_ = hex.DecodeString(MARKET_CREATED)
	evt_market_oi_changed,_ = hex.DecodeString(MARKET_OI_CHANGED)
)
func bigint_ptr_slice_to_str(data *[]*big.Int,separator string) string {
	var output bytes.Buffer
	length := len(*data)
	for i:=0 ; i< length ; i++  {
		if i>0 {
			output.WriteString(separator)
		}
		output.WriteString((*data)[i].String())
	}
	return output.String()
}
func outcomes_to_str(outcomes *[][32]byte,separator string) string {
	var output bytes.Buffer
	length := len(*outcomes)
	for i:=0 ; i<length ; i++ {
		if i>0 {
			output.WriteString(separator)
		}
		s := hex.EncodeToString((*outcomes)[i][:])
		output.WriteString(s)
	}
	return output.String()
}
func addresses_to_str(addresses *[]common.Address,separator string) string {
	var output bytes.Buffer
	length := len(*addresses)
	for i:=0 ; i<length ; i++ {
		if i>0 {
			output.WriteString(separator)
		}
		output.WriteString((*addresses)[i].String())
	}
	return output.String()
}
func (evt *MarketCreatedEvt) Dump() {	// dumps struct to stdout for debugging
	fmt.Printf("MarketCreated {\n")
	fmt.Printf("\tUniverse: %v\n",evt.Universe.String())
	fmt.Printf("\tEndTime: %v\n",evt.EndTime)
	fmt.Printf("\tExtraInfo: %v\n",evt.ExtraInfo)
	fmt.Printf("\tMarket: %v\n",evt.Market.String())
	fmt.Printf("\tMarketCreator: %v\n",evt.MarketCreator.String())
	fmt.Printf("\tDesignatedReporter: %v\n",evt.DesignatedReporter.String())
	fmt.Printf("\tFeePerCashInAttoCash: %v\n",evt.FeePerCashInAttoCash);
	prices := bigint_ptr_slice_to_str(&evt.Prices,",")
	fmt.Printf("\tPrices: %v\n",prices)
	fmt.Printf("\tMarketType: %v\n",evt.MarketType)
	fmt.Printf("\tNumTicks: %v\n",evt.NumTicks)
	outcomes := outcomes_to_str(&evt.Outcomes,",")
	fmt.Printf("\tOutcomes: %v\n",outcomes)
	fmt.Printf("\tNoShowBond: %v\n",evt.NoShowBond)
	fmt.Printf("\tTimestamp: %v\n",evt.Timestamp)
	fmt.Printf("}\n")
}
func (evt *MarketOIChangedEvt) Dump() {	// dumps struct to stdout for debugging

	fmt.Printf("MarketOIChanged {\n")
	fmt.Printf("\tUniverse: %v",evt.Universe.String())
	fmt.Printf("\tMarket: %v",evt.Market.String())
	fmt.Printf("\tMarket Open Interest: %v",evt.MarketOI.String())
	fmt.Printf("}\n")
}
func (evt *OrderEvt) Dump() { // dumps struct to stdout for debugging

	fmt.Printf("OrderEvnet {")
	fmt.Printf("\tUniverse: %v",evt.Universe.String())
	fmt.Printf("\tMarket: %v",evt.Market.String())
	fmt.Printf("\tEventType: %v",evt.EventType)
	fmt.Printf("\tOrderType: %v",evt.OrderType)
	fmt.Printf("\tOrderId: %v",evt.OrderId)
	fmt.Printf("\tTradeGroupId: %v",evt.TradeGroupId)
	fmt.Printf("\tAddressData: %v",addresses_to_str(&evt.AddressData,","))
	uintdata := bigint_ptr_slice_to_str(&evt.Uint256Data,",")
	fmt.Printf("\tUint256data: %v",uintdata)
	fmt.Printf("}\n")
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	client, err := ethclient.Dial("http://:::8545")
	if err != nil {
		log.Fatal(err)
	}

	storage := connect_to_storage()
	abi_data, err := ioutil.ReadFile("./abis/main-abis/Augur.abi")
	check(err)
	abi_rdr := bytes.NewReader(abi_data)
	check(err)
	contract_abi,err:=abi.JSON(abi_rdr)
	_  = contract_abi

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
									fmt.Printf("\t\t\tlog %v for contract %v\n",i,rcpt.Logs[i].Address.String())
									num_topics := len(rcpt.Logs[i].Topics)
									for j:=0; j < num_topics ; j++ {
										fmt.Printf("\t\t\t\tLog %v Topic %v , %v \n",i,j,rcpt.Logs[i].Topics[j].String())
										if 0 == bytes.Compare(rcpt.Logs[i].Topics[j].Bytes(),evt_market_created) {
											var mevt MarketCreatedEvt
											err := contract_abi.Unpack(&mevt,"MarketCreated",rcpt.Logs[i].Data)
											if err != nil {
												fmt.Printf("Event decode error: %v",err)
											} else {
												fmt.Printf("Block %v: MarketCreated event found",block.Number())
												mevt.Dump()
												storage.insert_market_created_evt(&mevt)
											}
										}
										if 0 == bytes.Compare(rcpt.Logs[i].Topics[j].Bytes(),evt_market_oi_changed) {
											var mevt MarketOIChangedEvt
											err := contract_abi.Unpack(&mevt,"MarketOIChanged",rcpt.Logs[i].Data)
											if err != nil {
												fmt.Printf("Event decode error: %v",err)
											} else {
												fmt.Printf("Block %v: MarketOIChanged found",block.Number())
												mevt.Dump()
												// Topics[1] contains market address, we extract it from 32 byte array (20 bytes addr)
												mbytes := common.BytesToAddress(rcpt.Logs[i].Topics[2][12:])
												market_addr := mbytes.String()
												fmt.Printf("Addr = %v\n",market_addr)
												storage.insert_market_oi_changed_evt(&mevt,market_addr)
											}
										}
									}
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
