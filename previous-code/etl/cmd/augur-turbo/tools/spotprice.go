// Dumps spot price at Balancer
package main

import (
	"os"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"


	. "github.com/PredictionExplorer/augur-explorer/primitives"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
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

	if len(os.Args) < 5 {
		fmt.Printf("Usage: \n\t\t%v [amm_factory_addr] [market_factory] [market_id] [outcome1] [outcome2]\n\n\t\tShows spot price for swapping outcomes\n\n",os.Args[0])
		os.Exit(1)
	}
	var copts = new(bind.CallOpts)

	amm_factory_addr := common.HexToAddress(os.Args[1])
	market_factory_addr := common.HexToAddress(os.Args[2])
	market_id,err := strconv.ParseInt(os.Args[3],10,64)
	if err != nil {
		fmt.Printf("Error parsing market_id field: %v\n",err)
		os.Exit(1)
	}
	outc1,err := strconv.ParseInt(os.Args[4],10,64)
	if err != nil {
		fmt.Printf("Error parsing outcome1 field: %v\n",err)
		os.Exit(1)
	}
	outc2,err := strconv.ParseInt(os.Args[5],10,64)
	if err != nil {
		fmt.Printf("Error parsing outcome1 field: %v\n",err)
		os.Exit(1)
	}
	amm_factory,err := NewAMMFactory(amm_factory_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate AMMFactory contract: %v\n",err)
		os.Exit(1)
	}

	market_factory,err := NewSportsLinkMarketFactory(market_factory_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Market Factory contract: %v\n",err)
		os.Exit(1)
	}
	sharefactor,err := market_factory.ShareFactor(copts)
	if err!=nil {
		fmt.Printf("Error getting Share factor: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("ShareFactor = %v\n",sharefactor.String())

	big_market_id := big.NewInt(market_id)
	market_obj,err:=market_factory.GetMarket(copts,big_market_id)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}

	if outc1 >= int64(len(market_obj.ShareTokens)) {
		fmt.Printf("Outcome 1 is larger than the size of outcomes (%v)\n",outc1)
		os.Exit(1)
	}
	if outc2 >= int64(len(market_obj.ShareTokens)) {
		fmt.Printf("Outcome 2 is larger than the size of outcomes (%v)\n",outc2)
		os.Exit(1)
	}

	token1 := market_obj.ShareTokens[outc1]
	token2 := market_obj.ShareTokens[outc2]

	erc20_token1,err := NewOwnedERC20(token1,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC20 contract for %v: %v\n",token1.String(),err)
		os.Exit(1)
	}
	symbol1,err:=erc20_token1.Symbol(copts)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}

	erc20_token2,err := NewOwnedERC20(token2,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC20 contract for %v: %v\n",token2.String(),err)
		os.Exit(1)
	}
	symbol2,err:=erc20_token2.Symbol(copts)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}
	decimals,err:=erc20_token1.Decimals(copts)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}

	bpool_addr,err := amm_factory.Pools(copts,market_factory_addr,big_market_id)
	if err!=nil {
		fmt.Printf("Failed to retrieve Balancer Pool (BPool) contract from market factory: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Params:\n")
	fmt.Printf("\tAMM factory:    %v\n",amm_factory_addr.String())
	fmt.Printf("\tMarket factory: %v\n",market_factory_addr.String())
	fmt.Printf("\tMarket ID:      %v\n",market_id)
	fmt.Printf("\tOutcome IN:     %v (%v)\n",outc1,symbol1)
	fmt.Printf("\tOutcome OUT:    %v (%v)\n",outc2,symbol2)
	fmt.Printf("\n")

	fmt.Printf("Pool addr: %v\n",bpool_addr.String())
	fmt.Printf("Market {\n")
	fmt.Printf("\tSettlement Address: %v\n",market_obj.SettlementAddress.String())
	fmt.Printf("\tShareTokens:\n")
	for i:=0; i<len(market_obj.ShareTokens) ; i++ {
		fmt.Printf("\t\t%v\n",market_obj.ShareTokens[i].String())
	}
	fmt.Printf("\tEndTime: %v\n",market_obj.EndTime.String())
	fmt.Printf("\tWinner: %v\n",market_obj.Winner.String())
	fmt.Printf("\tSettlement Fee: %v\n",market_obj.SettlementFee.String())
	fmt.Printf("\tProtocol Fee: %v\n",market_obj.ProtocolFee.String())
	fmt.Printf("\tStakerFee: %v\n",market_obj.StakerFee.String())
	fmt.Printf("\tCreation Timestamp: %v\n",market_obj.CreationTimestamp.String())
	fmt.Printf("}\n")

	bpool,err := NewBPool(bpool_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate BPool contract: %v\n",err)
		os.Exit(1)
	}


	spot_price,err:=bpool.GetSpotPrice(copts,token1,token2)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}

	divisor:=big.NewInt(0)
	if decimals == 0 {
		divisor = big.NewInt(1)	//to avoid divide by 0 error
	} else {
		multiplier_str := strings.Repeat("0",int(decimals))
		multiplier_str = "1" + multiplier_str
		divisor.SetString(multiplier_str,10)
	}
	fmt.Printf("decimals=%v, divisor=%v\n",decimals,divisor.String())
	compact_price:= big.NewInt(0)
	reminder := big.NewInt(0)
	compact_price.QuoRem(spot_price,divisor,reminder)

	fmt.Printf("Spot price:\n")
	fmt.Printf(
		"Swap of \n\t%v (outcome %v (%v)) for \n\t%v (outcome %v (%v)) = %v\n",
		token1.String(),outc1,symbol1,
		token2.String(),outc2,symbol2,
		spot_price.String(),
	)
	fmt.Printf(
		"Price with floating point: %v.%018s \n",compact_price.String(),reminder.String(),
	)
}
