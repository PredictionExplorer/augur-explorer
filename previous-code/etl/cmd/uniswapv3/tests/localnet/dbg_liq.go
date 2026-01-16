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
	DBG_MOD_POS=	"d1067085947d0663c9eae886f8c165cbbccbc686928d7742d4241b4fb310228f"
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
			fmt.Printf("DBG_MOD_POS {\n")
			fmt.Printf("\tOwner:                          %v\n",eth_evt.Owner.String())
			fmt.Printf("\tTickLower:                      %v\n",eth_evt.TickLower)
			fmt.Printf("\tTickUpper:                      %v\n",eth_evt.TickUpper)
			fmt.Printf("\tSlot0Tick:                      %v\n",eth_evt.Slot0Tick.String())
			fmt.Printf("\tLiquidityDelta:                 %v\n",eth_evt.LiquidityDelta.String())
			fmt.Printf("\tLiquidityBefore:                %v\n",eth_evt.LiquidityBefore.String())
			fmt.Printf("\tAmount0:                        %v\n",eth_evt.Amount0.String())
			fmt.Printf("\tAmount1:                        %v\n",eth_evt.Amount1.String())
			fmt.Printf("\tSqrtPriceX96:                   %v\n",eth_evt.SqrtPriceX96.String())
			fmt.Printf("}\n")
		}
		if bytes.Equal(log.Topics[0].Bytes(),evt_upd_pos) {
			var eth_evt IUniswapV3PoolEventsDBGUPDPOS
			err := dbg_abi.UnpackIntoInterface(&eth_evt,"DBG_UPD_POS",log.Data)
			if err != nil {
				fmt.Printf("Error unpacking event DBG_UPD_POS: %v\n",err)
				os.Exit(1)
			}
			fmt.Printf("DBG_UPD_POS {\n")
			fmt.Printf("\tOwner:                          %v\n",eth_evt.Owner.String())
			fmt.Printf("\tTickLower:                      %v\n",eth_evt.TickLower)
			fmt.Printf("\tTickUpper:                      %v\n",eth_evt.TickUpper)
			fmt.Printf("\tTick:                           %v\n",eth_evt.Tick)
			fmt.Printf("\tLiquidityDelta:                 %v\n",eth_evt.LiquidityDelta.String())
			fmt.Printf("\tfeeGrowthGlobal0X128Before:     %v\n",eth_evt.FeeGrowthGlobal0X128Before.String())
			fmt.Printf("\tfeeGrowthGlobal1X128Before:     %v\n",eth_evt.FeeGrowthGlobal1X128Before.String())
			fmt.Printf("\tfeeGrowthInside0X128:           %v\n",eth_evt.FeeGrowthInside0X128)
			fmt.Printf("\tfeeGrowthInside1X128:           %v\n",eth_evt.FeeGrowthInside1X128)
			fmt.Printf("\tFlippedLower:                   %v\n",eth_evt.FlippedLower)
			fmt.Printf("\tFlippedUpper:                   %v\n",eth_evt.FlippedUpper)
			fmt.Printf("}\n")
		}
	}
}
