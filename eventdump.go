/// Utinity to dump Event logs

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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	MARKET_CREATED = "ea17ae24b0d40ea7962a6d832db46d1f81eaec1562946d0830d1c21d4c000ec1"
)
var (
	evt_market_created,_ = hex.DecodeString(MARKET_CREATED)
)
type MarketCreatedEvt struct {
	Universe             common.Address
	EndTime              *big.Int
	ExtraInfo            string
	Market               common.Address
	MarketCreator        common.Address
	DesignatedReporter   common.Address
	FeePerCashInAttoCash *big.Int
	Prices               []*big.Int
	MarketType           uint8
	NumTicks             *big.Int
	Outcomes             [][32]byte
	NoShowBond           *big.Int
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}
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

	abi_data, err := ioutil.ReadFile("./abis/main-abis/Augur.abi")
	check(err)
	abi_rdr := bytes.NewReader(abi_data)
	check(err)
	contract_abi,err:=abi.JSON(abi_rdr)
	_  = contract_abi

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
//											fmt.Printf("\t\t\t\tMATCH! Data len = %v\n",len(rcpt.Logs[i].Data))
											var mevt MarketCreatedEvt
											err := contract_abi.Unpack(&mevt,"MarketCreated",rcpt.Logs[i].Data)
											if err != nil {
												fmt.Printf("Event decode error: %v",err)
											} else {
												fmt.Printf("Block %v: MarketCreated event found",block.Number())
												mevt.Dump()
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
