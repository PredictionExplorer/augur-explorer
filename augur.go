package main

import (
	"fmt"
//	"context"
//	"log"
//	"math/big"
	"bytes"
//	"encoding/hex"
	"io/ioutil"

//	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)
func load_abi(fname string) *abi.ABI {

	abi_data, err := ioutil.ReadFile(fname)
	check(err)
	abi_rdr := bytes.NewReader(abi_data)
	check(err)
	abi,err := abi.JSON(abi_rdr)
	check(err)
	return &abi
}
func augur_init() {

	// Augur service involves 39 contracts in total. We only use a few of them

	// Load main Agur contract ABI
	augur_abi = load_abi("./abis/main-abis/Augur.abi")

	// Load AugurTrading contract
	trading_abi = load_abi("./abis/trading-abis/AugurTrading.abi")
}
func process_event(log *types.Log) {

	if log == nil {
		Fatalf("process_event() received null pointer")
	}
	num_topics := len(log.Topics)
	for j:=0; j < num_topics ; j++ {
		fmt.Printf("\t\t\t\tLog Topic %v , %v \n",j,log.Topics[j].String())
		if 0 == bytes.Compare(log.Topics[j].Bytes(),evt_market_order) {
			var mevt MktOrderEvt
			fmt.Printf("log.Data= %+v\n",log.Data)
			err := trading_abi.Unpack(&mevt,"OrderEvent",log.Data)
			if err != nil {
				Fatalf("Event decode error: %v",err)
			} else {
				fmt.Printf("Block %v: OrderEvent event found\n",log.BlockNumber)
				mevt.Dump()
				bytes := common.BytesToAddress(log.Topics[1][12:])	// extract universe addr
				a_universe := bytes.String()
				bytes = common.BytesToAddress(log.Topics[2][12:])	// extract market addr
				a_market := bytes.String()
				storage.insert_market_order_evt(a_universe,a_market,&mevt)
			}
		}
		if 0 == bytes.Compare(log.Topics[j].Bytes(),evt_market_created) {
			var mevt MarketCreatedEvt
			err := augur_abi.Unpack(&mevt,"MarketCreated",log.Data)
			if err != nil {
				Fatalf("Event decode error: %v",err)
			} else {
				fmt.Printf("Block %v: MarketCreated event found\n",log.BlockNumber)
				mevt.Dump()
				storage.insert_market_created_evt(&mevt)
			}
		}
		if 0 == bytes.Compare(log.Topics[j].Bytes(),evt_market_oi_changed) {
			var mevt MarketOIChangedEvt
			err := augur_abi.Unpack(&mevt,"MarketOIChanged",log.Data)
			if err != nil {
				Fatalf("Event decode error: %v",err)
			} else {
				fmt.Printf("Block %v: MarketOIChanged found\n",log.BlockNumber)
				mevt.Dump()
				// Topics[2] contains market address, we extract it from 32 byte array (20 bytes addr)
				mbytes := common.BytesToAddress(log.Topics[2][12:])	// extract addr from 32byte array
				market_addr := mbytes.String()
				storage.insert_market_oi_changed_evt(&mevt,market_addr)
			}
		}
	}
}
/*
	if 0 == bytes.Compare(rcpt.Logs[i].Topics[j].Bytes(),evt_market_created) {
		var mevt MarketCreatedEvt
		err := contract_abi.Unpack(&mevt,"MarketCreated",rcpt.Logs[i].Data)
		if err != nil {
			Fatalf("Event decode error: %v",err)
		} else {
			fmt.Printf("Block %v: MarketCreated event found",block.Number())
			mevt.Dump()
			storage.insert_market_created_evt(&mevt)
		}
		return
	}
	if 0 == bytes.Compare(rcpt.Logs[i].Topics[j].Bytes(),evt_market_oi_changed) {
		var mevt MarketOIChangedEvt
		err := contract_abi.Unpack(&mevt,"MarketOIChanged",rcpt.Logs[i].Data)
		if err != nil {
			Faltalf("Event decode error: %v",err)
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
*/
