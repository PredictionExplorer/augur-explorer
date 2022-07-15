// Dumps DBG_SWAP_HEAD, DBG_SWAP_LOOP and DBG_SWAP_TAIL events
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
	DBG_SWAP_HEAD =	"884bf0b67c5a7a033ec93aeab39846c4e697715e51811c8b9499b2666468c207"
	DBG_SWAP_LOOP =	"5cdbca1f8631135a1302301500a65ffc83d0b7c1177e309dfc6a14e4ea6c58f3"
	DBG_SWAP_TAIL = "507cf18262378ccee2abf0e6b16ba3a87afe50f610f6d98fec691ae6ab12a584"
)
var (
	RPC_URL string
	dbg_abi *abi.ABI
	evt_swap_head,_ = hex.DecodeString(DBG_SWAP_HEAD)
	evt_swap_loop,_ = hex.DecodeString(DBG_SWAP_LOOP)
	evt_swap_tail,_ = hex.DecodeString(DBG_SWAP_TAIL)
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
		if bytes.Equal(log.Topics[0].Bytes(),evt_swap_head) {
			var eth_evt IUniswapV3PoolEventsDBGSWAPHEAD
			err := dbg_abi.UnpackIntoInterface(&eth_evt,"DBG_SWAP_HEAD",log.Data)
			if err != nil {
				fmt.Printf("Error unpacking event DBG_SWAP_HEAD: %v\n",err)
				os.Exit(1)
			}
			eth_evt.Pool = common.BytesToAddress(log.Topics[1][12:])
			fmt.Printf("Pool addr:            %v\n",eth_evt.Pool.String())
			fmt.Printf("ZeroForOne:           %v\n",eth_evt.ZeroForOne)
			fmt.Printf("Amount:               %v\n",eth_evt.Amount.String())
			fmt.Printf("SqrtStartPrice:       %v\n",eth_evt.SqrtStartPrice.String())
			fmt.Printf("SqrtPriceLimit:       %v\n",eth_evt.SqrtPriceLimit.String())
			fmt.Printf("Slot0SqrtPriceX96:    %v\n",eth_evt.Slot0sqrtpricex96.String())
		}
		if bytes.Equal(log.Topics[0].Bytes(),evt_swap_loop) {
			var eth_evt IUniswapV3PoolEventsDBGSWAPLOOP
			err := dbg_abi.UnpackIntoInterface(&eth_evt,"DBG_SWAP_LOOP",log.Data)
			if err != nil {
				fmt.Printf("Error unpacking event DBG_SWAP_LOOP: %v\n",err)
				os.Exit(1)
			}
			eth_evt.Pool = common.BytesToAddress(log.Topics[1][12:])
			fmt.Printf("Pool addr:            %v\n",eth_evt.Pool.String())
			fmt.Printf("SqrtPriceX96:         %v\n",eth_evt.SqrtPriceX96.String())
			fmt.Printf("SqrtPriceStartX96:    %v\n",eth_evt.SqrtPriceStartX96.String())
			fmt.Printf("SqrtPriceNextX96:     %v\n",eth_evt.SqrtPriceNextX96.String())
			fmt.Printf("Tick:                 %v\n",eth_evt.Tick)
			fmt.Printf("TickCumulative:       %v\n",eth_evt.TickCumulative)
			fmt.Printf("Initialized:          %v\n",eth_evt.Initialized)
			fmt.Printf("StepAmountIn:         %v\n",eth_evt.StepAmountIn.String())
			fmt.Printf("StepAmountOut:        %v\n",eth_evt.StepAmountOut.String())
			fmt.Printf("AmountProcessed:      %v\n",eth_evt.AmountProcessed.String())
			fmt.Printf("FeeAmount:            %v\n",eth_evt.FeeAmount.String())
			fmt.Printf("Liquidity:            %v\n",eth_evt.Liquidity.String())
			fmt.Printf("ExactInput:           %v\n",eth_evt.ExactInput)
			fmt.Printf("FeeGrrowthGlobalX128: %v\n",eth_evt.FeeGrowthGlobalX128.String())
		}
		if bytes.Equal(log.Topics[0].Bytes(),evt_swap_tail) {
			var eth_evt IUniswapV3PoolEventsDBGSWAPTAIL
			err := dbg_abi.UnpackIntoInterface(&eth_evt,"DBG_SWAP_TAIL",log.Data)
			if err != nil {
				fmt.Printf("Error unpacking event DBG_SWAP_TAIL: %v\n",err)
				os.Exit(1)
			}
			eth_evt.Pool = common.BytesToAddress(log.Topics[1][12:])
			fmt.Printf("Pool addr:            %v\n",eth_evt.Pool.String())
			fmt.Printf("FeeGrowthGlobal0X128: %v\n",eth_evt.FeeGrowthGlobal0X128.String())
			fmt.Printf("FeeGrowthGlobal1X128: %v\n",eth_evt.FeeGrowthGlobal1X128.String())
			fmt.Printf("Tick:                 %v\n",eth_evt.Tick)
			fmt.Printf("SqrtPriceX96:         %v\n",eth_evt.SqrtPriceX96.String())
			fmt.Printf("Liquidity:            %v\n",eth_evt.Liquidity.String())
			fmt.Printf("LiquidityDiff:        %v\n",eth_evt.LiquidityDiff.String())
		}
	}
}
