// Dumps balancer pool info
package main

import (
	"os"
	"fmt"
	"math/big"
	//"strconv"
	//"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"


	. "github.com/PredictionExplorer/augur-explorer/primitives"
	//. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
)
var (
	RPC_URL string
	token_addr		common.Address
)
func main() {

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Printf("Usage: \n\t\t%v [pool_addr]\n\n\t\tShows pool related data",os.Args[0])
		os.Exit(1)
	}
	var copts = new(bind.CallOpts)

	bpool_addr := common.HexToAddress(os.Args[1])

	bpool,err := NewBPool(bpool_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate BPool contract: %v\n",err)
		os.Exit(1)
	}
	is_public,err := bpool.IsPublicSwap(copts)
	if err != nil {
		fmt.Printf("Error in IsPublicSwap: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Is public: %v\n",is_public)

	is_finalized,err := bpool.IsFinalized(copts)
	if err != nil {
		fmt.Printf("Error in IsFinalized: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Is finalized: %v\n",is_finalized)

	num_tokens,err := bpool.GetNumTokens(copts)
	if err != nil {
		fmt.Printf("Error in GetNumTokens: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Num tokens: %v\n",num_tokens)

	tokens,err := bpool.GetCurrentTokens(copts)
	if err != nil {
		fmt.Printf("Error in GetCurrentTokens: %v\n",err)
		os.Exit(1)
	}
	divisor:=big.NewInt(0)
	divisor.SetString("10000000000000000",10)
	fmt.Printf("Token addresses:\n")
	for i:=0; i<len(tokens); i++ {
		weight,err := bpool.GetNormalizedWeight(copts,tokens[i])
		if err != nil {
			fmt.Printf("Error in GetSwapFee: %v\n",err)
			os.Exit(1)
		}
		compact_num:= big.NewInt(0)
		reminder := big.NewInt(0)
		compact_num.QuoRem(weight,divisor,reminder)
		fmt.Printf("\t%v (%v.%018s %%)\n",tokens[i].String(),compact_num.String(),reminder.String())
	}
	controller,err := bpool.GetController(copts)
	if err != nil {
		fmt.Printf("Error in GetController: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Controller: %v\n",controller.String())

	swap_fee,err := bpool.GetSwapFee(copts)
	if err != nil {
		fmt.Printf("Error in GetSwapFee: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Swap Fee: %v\n",swap_fee.String())

}
