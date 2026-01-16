// Dumps shares of a User in a market
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
	"github.com/ethereum/go-ethereum/core/types"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	CHAIN_ID		int64 = 80001
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

	if len(os.Args) < 4 {
		fmt.Printf(
			"Usage: \n\t\t%v [market_factory] [market_id] \n\t\t"+
			"Shows owned shares of a user in a market\n\n",os.Args[0],
		)
		os.Exit(1)
	}
	var copts = new(bind.CallOpts)

	market_factory_addr := common.HexToAddress(os.Args[1])
	market_id,err := strconv.ParseInt(os.Args[2],10,64)
	if err != nil {
		fmt.Printf("Error parsing market_id field: %v\n",err)
		os.Exit(1)
	}
	user_addr := common.HexToAddress(os.Args[3])

	market_factory,err := NewSportsLinkMarketFactory(market_factory_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Market Factory contract: %v\n",err)
		os.Exit(1)
	}

	collateral_addr,err := market_factory.Collateral(copts)
	erc20_token,err := NewOwnedERC20(collateral_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC20 contract for the collateral %v: %v\n",collateral_addr.String(),err)
		os.Exit(1)
	}
	decimals,err:=erc20_token.Decimals(copts)
	if err!=nil {
		fmt.Printf("Error during call Decimals() for the collateral: %v\n",err)
	}
	balance,err:=erc20_token.BalanceOf(copts,user_addr)
	if err!=nil {
		fmt.Printf("Error during call BalanceOf() for the collateral: %v\n",err)
	}
	divisor:=big.NewInt(0)
	if decimals == 0 {
		divisor = big.NewInt(1)	//to avoid divide by 0 error
	} else {
		multiplier_str := strings.Repeat("0",int(decimals))
		multiplier_str = "1" + multiplier_str
		divisor.SetString(multiplier_str,10)
	}
	compact_balance := big.NewInt(0)
	reminder := big.NewInt(0)
	compact_balance.QuoRem(balance,divisor,reminder)
	fmt.Printf(
		"Collateral balance : %v (%v.%018s)\n",
		balance.String(),
		compact_balance.String(),
		reminder.String(),
	)

	big_market_id := big.NewInt(market_id)
	market_obj,err:=market_factory.GetMarket(copts,big_market_id)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}

	for i:=0 ; i<len(market_obj.ShareTokens); i++ {
		token_addr := market_obj.ShareTokens[i]
		erc20_token,err := NewOwnedERC20(token_addr,eclient)
		if err!=nil {
			fmt.Printf("Failed to instantiate ERC20 contract for %v: %v\n",token_addr.String(),err)
			os.Exit(1)
		}
		symbol,err:=erc20_token.Symbol(copts)
		if err!=nil {
			fmt.Printf("Error during call Symbol(): %v\n",err)
		}
		decimals,err:=erc20_token.Decimals(copts)
		if err!=nil {
			fmt.Printf("Error during call Decimals(): %v\n",err)
		}
		var copts = new(bind.CallOpts)
		balance,err:=erc20_token.BalanceOf(copts,user_addr)
		if err!=nil {
			fmt.Printf("Error during call BalanceOf(): %v\n",err)
		}
		divisor:=big.NewInt(0)
		if decimals == 0 {
			divisor = big.NewInt(1)	//to avoid divide by 0 error
		} else {
			multiplier_str := strings.Repeat("0",int(decimals))
			multiplier_str = "1" + multiplier_str
			divisor.SetString(multiplier_str,10)
		}
		compact_balance := big.NewInt(0)
		reminder := big.NewInt(0)
		compact_balance.QuoRem(balance,divisor,reminder)
		fmt.Printf(
			"%v [%v] : %v (%v.%018s)\n",
			token_addr.String(),
			symbol,
			balance.String(),
			compact_balance.String(),
			reminder.String(),
		)
	}

}
