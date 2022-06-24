// Shows Slot0 (global state) struct of the pool
package main

import (
	"os"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"


	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
)
var (
	RPC_URL string
	token_addr		common.Address
)
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Printf("Usage: \n\t\t%v [pool_addr]\n\n\t\tShows Slot0 struct of the pool\n",os.Args[0])
		os.Exit(1)
	}
	contract_addr := common.HexToAddress(os.Args[1])
	pool,err := NewUniswapV3Pool(contract_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate contract: %v\n",err)
		os.Exit(1)
	}
	var copts = new(bind.CallOpts)
	slot0,err:=pool.Slot0(copts)
	if err!=nil {
		fmt.Printf("Error during Slot0() call: %v\n",err)
		os.Exit(1)
	}
	liquidity,err := pool.Liquidity(copts)
	if err != nil {
		fmt.Printf("Error during Liquidity() call: %v\n",err)
		os.Exit(1)
	}
	feeGrowthGlobal0X128,err := pool.FeeGrowthGlobal0X128(copts)
	if err != nil {
		fmt.Printf("Error during feeGrowthGlobal0X128() call: %v\n",err)
		os.Exit(1)
	}
	feeGrowthGlobal1X128,err := pool.FeeGrowthGlobal1X128(copts)
	if err != nil {
		fmt.Printf("Error during feeGrowthGlobal1X128() call: %v\n",err)
		os.Exit(1)
	}
	tick_spacing,err := pool.TickSpacing(copts)
	if err != nil {
		fmt.Printf("Error during TickSpacing call: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Tick spacing:              \t%v\n",tick_spacing)
	fmt.Printf("SqrtPriceX96:              \t%v\n",slot0.SqrtPriceX96.String())
	fmt.Printf("Tick:                      \t%v\n",slot0.Tick.String())
	fmt.Printf("ObservationIndex:          \t%v\n",slot0.ObservationIndex)
	fmt.Printf("ObservationCardinalityNext:\t%v\n",slot0.ObservationCardinalityNext)
	fmt.Printf("FeeProtocol:               \t%v\n",slot0.FeeProtocol)
	fmt.Printf("Unlocked:                  \t%v\n",slot0.Unlocked)
	fmt.Printf("\n")
	fmt.Printf("Liquidity:                 \t%v\n",liquidity.String())
	fmt.Printf("feeGrowthGlobal1X128:      \t%v\n",feeGrowthGlobal1X128)
	fmt.Printf("feeGrowthGlobal0X128:      \t%v\n",feeGrowthGlobal0X128)
}
