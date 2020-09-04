// Dumps ShareTokens that an account currently holds
package main

import (
	"os"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"


	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
const (
//	SHARE_TOKEN_ADDR_PROD = "0x9e4799ff2023819b1272eee430eadf510eDF85f0"	// MainNet
	SHARE_TOKEN_ADDR_PROD = "0xE60c9fe85aEE7B4848a97271dA8c86323CdFb897"	// Dev
)
var (
	share_token_addr common.Address = common.HexToAddress(SHARE_TOKEN_ADDR_PROD)
	RPC_URL string
)
func main() {

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		Fatalf("Can't connect to ETH RPC: %v\n",err)
	}

	ctrct_share_token,err := NewShareToken(share_token_addr,eclient)
	if err!=nil {
		Fatalf("Failed to instantiate ShareToken contract: %v\n",err)
	}

	if len(os.Args) < 4 {
		fmt.Printf("Usage: \n\t\t%v [market_address] [ethereum_address_of_augur_shares_token_holder] [outcome_idx]\n\n(outcome_idx is integer, addresses can be 0x prepended or without (40 or 42 chars)). This executable have been configured for Main Net. ShareToken contract addr: %v",os.Args[0],SHARE_TOKEN_ADDR_PROD)
		Fatalf("Aborting.")
	}

	market_addr := common.HexToAddress(os.Args[1])
	holder_addr := common.HexToAddress(os.Args[2])
	outcome_idx,err := strconv.ParseInt(os.Args[3],10,32)
	if err!=nil {
		Fatalf("Bad integer for 'outcome_idx': %v\n",err)
	}
	big_outcome_idx:=big.NewInt(outcome_idx)

	var copts = new(bind.CallOpts)
	balance,err:=ctrct_share_token.BalanceOfMarketOutcome(copts,market_addr,big_outcome_idx,holder_addr)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	} else {
		fmt.Printf("%v\n",balance.String())
	}
}
