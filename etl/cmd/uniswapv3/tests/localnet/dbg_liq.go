// Dumps DBG_MOD_POS and DBG_UPD_POS events
package main

import (
	"os"
	"fmt"
	"bytes"
	"context"
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	 "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"

)
const (
	DBG_MOD_POS=	"b05d03afcde6cc3b7059da906a709c3e3c68c3c665a90c35ae0ade8d0903f666"
	DBG_UPD_POS=	"2e894cce336ea758e6f37d72cafa9ef74e237e6e488a2f6ad4148c3fbca85f24"
)
var (
	RPC_URL string
	dbg_abi *abi.ABI
	evt_mod_pos,_ = hex.DecodeString(DBG_MOD_POS)
	evt_upd_pos,_ = hex.DecodeString(DBG_UPD_POS)
)
func main() {

	abi_parsed := strings.NewReader(IUniswapV3PoolEventsABI)
	abi,err := abi.JSON(abi_parsed)
	dbg_abi = &abi
	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Printf("Usage: \n\t\t%v [tx_hash]\n\n\t\tShows Uniswap v3 debug (DBG_*) events\n",os.Args[0])
		os.Exit(1)
	}
	tx_hash := common.HexToHash(os.Args[1])
	receipt,err := eclient.TransactionReceipt(context.Background(),tx_hash)
	if err!=nil {
		fmt.Printf("Receipt not found: %v\n",err)
		os.Exit(1)
	}
	logs := receipt.Logs
	for i:=0; i<len(logs); i++ {
		log := logs[i]
		if bytes.Equal(log.Topics[0].Bytes(),evt_mod_pos) {
			var eth_evt IUniswapV3PoolEventsDBGMODPOS
			err := dbg_abi.UnpackIntoInterface(&eth_evt,"DBG_MOD_POS",log.Data)
			if err != nil {
				fmt.Printf("Error unpacking event DBG_MOD_POS: %v\n",err)
				os.Exit(1)
			}
			fmt.Printf("Owner:                %v\n",eth_evt.Owner.String())
			fmt.Printf("TickLower:            %v\n",eth_evt.TickLower)
			fmt.Printf("TickUpper:            %v\n",eth_evt.TickUpper)
			fmt.Printf("Slot0Tick:            %v\n",eth_evt.Slot0Tick.String())
			fmt.Printf("LiquidityDelta:       %v\n",eth_evt.LiquidityDelta.String())
			fmt.Printf("LiquidityBefore:      %v\n",eth_evt.LiquidityBefore.String())
			fmt.Printf("Amount0:              %v\n",eth_evt.Amount0.String())
			fmt.Printf("Amount1:              %v\n",eth_evt.Amount1.String())
			fmt.Printf("SqrtPriceX96:         %v\n",eth_evt.SqrtPriceX96.String())
		}
		if bytes.Equal(log.Topics[0].Bytes(),evt_upd_pos) {
			var eth_evt IUniswapV3PoolEventsDBGUPDPOS
			err := dbg_abi.UnpackIntoInterface(&eth_evt,"DBG_UPD_POS",log.Data)
			if err != nil {
				fmt.Printf("Error unpacking event DBG_UPD_POS: %v\n",err)
				os.Exit(1)
			}
			fmt.Printf("Owner:                          %v\n",eth_evt.Owner.String())
			fmt.Printf("TickLower:                      %v\n",eth_evt.TickLower)
			fmt.Printf("TickUpper:                      %v\n",eth_evt.TickUpper)
			fmt.Printf("Tick:                           %v\n",eth_evt.Tick)
			fmt.Printf("LiquidityDelta:                 %v\n",eth_evt.LiquidityDelta.String())
			fmt.Printf("feeGrowthGlobal0X128Before:     %v\n",eth_evt.FeeGrowthGlobal0X128Before.String())
			fmt.Printf("feeGrowthGlobal1X128Before:     %v\n",eth_evt.FeeGrowthGlobal1X128Before.String())
			fmt.Printf("feeGrowthInside0X128:           %v\n",eth_evt.FeeGrowthInside0X128)
			fmt.Printf("feeGrowthInside1X128:           %v\n",eth_evt.FeeGrowthInside1X128)
			fmt.Printf("FlippedLower:                   %v\n",eth_evt.FlippedLower)
			fmt.Printf("FlippedUpper:                   %v\n",eth_evt.FlippedUpper)
		}
	}
}
