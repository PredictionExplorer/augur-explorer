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
	DBG_SWAP_LOOP =	"f5b431d95ff8c2e6f40d1ab0bdd3f039c0cec559bfbe2b354dcd8e90c52d5929"
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
			fmt.Printf("DBG_SWAP_HEAD {\n")
			fmt.Printf("\tPool addr:            %v\n",eth_evt.Pool.String())
			fmt.Printf("\tZeroForOne:           %v\n",eth_evt.ZeroForOne)
			fmt.Printf("\tAmount:               %v\n",eth_evt.Amount.String())
			fmt.Printf("\tSqrtStartPrice:       %v\n",eth_evt.SqrtStartPrice.String())
			fmt.Printf("\tSqrtPriceLimit:       %v\n",eth_evt.SqrtPriceLimit.String())
			fmt.Printf("\tSlot0SqrtPriceX96:    %v\n",eth_evt.Slot0sqrtpricex96.String())
			fmt.Printf("}\n")
		}
		if bytes.Equal(log.Topics[0].Bytes(),evt_swap_loop) {
			var eth_evt IUniswapV3PoolEventsDBGSWAPLOOP
			err := dbg_abi.UnpackIntoInterface(&eth_evt,"DBG_SWAP_LOOP",log.Data)
			if err != nil {
				fmt.Printf("Error unpacking event DBG_SWAP_LOOP: %v\n",err)
				os.Exit(1)
			}
			eth_evt.Pool = common.BytesToAddress(log.Topics[1][12:])
			fmt.Printf("DBG_SWAP_LOOP {\n")
			fmt.Printf("\tPool addr:            %v\n",eth_evt.Pool.String())
			fmt.Printf("\tSqrtPriceX96:         %v\n",eth_evt.SqrtPriceX96.String())
			fmt.Printf("\tSqrtPriceStartX96:    %v\n",eth_evt.SqrtPriceStartX96.String())
			fmt.Printf("\tSqrtPriceNextX96:     %v\n",eth_evt.SqrtPriceNextX96.String())
			fmt.Printf("\tTick:                 %v\n",eth_evt.Tick)
			fmt.Printf("\tTickCumulative:       %v\n",eth_evt.TickCumulative)
			fmt.Printf("\tInitialized:          %v\n",eth_evt.Initialized)
			fmt.Printf("\tStepAmountIn:         %v\n",eth_evt.StepAmountIn.String())
			fmt.Printf("\tStepAmountOut:        %v\n",eth_evt.StepAmountOut.String())
			fmt.Printf("\tAmountProcessed:      %v\n",eth_evt.AmountProcessed.String())
			fmt.Printf("\tFeeAmount:            %v\n",eth_evt.FeeAmount.String())
			fmt.Printf("\tLiquidity:            %v\n",eth_evt.Liquidity.String())
			fmt.Printf("\tExactInput:           %v\n",eth_evt.ExactInput)
			fmt.Printf("\tFeeGrrowthGlobalX128: %v\n",eth_evt.FeeGrowthGlobalX128.String())
			fmt.Printf("}\n")
		}
		if bytes.Equal(log.Topics[0].Bytes(),evt_swap_tail) {
			var eth_evt IUniswapV3PoolEventsDBGSWAPTAIL
			err := dbg_abi.UnpackIntoInterface(&eth_evt,"DBG_SWAP_TAIL",log.Data)
			if err != nil {
				fmt.Printf("Error unpacking event DBG_SWAP_TAIL: %v\n",err)
				os.Exit(1)
			}
			eth_evt.Pool = common.BytesToAddress(log.Topics[1][12:])
			fmt.Printf("DBG_SWAP_TAIL {\n")
			fmt.Printf("\tPool addr:            %v\n",eth_evt.Pool.String())
			fmt.Printf("\tFeeGrowthGlobal0X128: %v\n",eth_evt.FeeGrowthGlobal0X128.String())
			fmt.Printf("\tFeeGrowthGlobal1X128: %v\n",eth_evt.FeeGrowthGlobal1X128.String())
			fmt.Printf("\tTick:                 %v\n",eth_evt.Tick)
			fmt.Printf("\tSqrtPriceX96:         %v\n",eth_evt.SqrtPriceX96.String())
			fmt.Printf("\tLiquidity:            %v\n",eth_evt.Liquidity.String())
			fmt.Printf("\tLiquidityDiff:        %v\n",eth_evt.LiquidityDiff.String())
			fmt.Printf("}\n")
		}
	}
}
