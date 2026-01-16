// Dumps ERC20 token info of a User address
package main

import (
	"os"
	"fmt"
	"math/big"
	"strings"

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

	RPC_URL = os.Getenv("AUGUR_ETH_NODE_RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Printf("Usage: \n\t\t%v [contract_addr] [user_addr]\n\n\t\tShows amount of tokens for an ERC20 tokens of [user_addr] account\n\n",os.Args[0])
		os.Exit(1)
	}
	contract_addr := common.HexToAddress(os.Args[1])
	user_addr := common.HexToAddress(os.Args[2])
	erc20_token,err := NewOwnedERC20(contract_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC20 contract: %v\n",err)
		os.Exit(1)
	}

	var copts = new(bind.CallOpts)
	balance,err:=erc20_token.BalanceOf(copts,user_addr)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}
	decimals,err:=erc20_token.Decimals(copts)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
	}
	symbol,err:=erc20_token.Symbol(copts)
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
	fmt.Printf("Amount of %v tokens: %v\n",symbol,balance.String())
	compact_balance := big.NewInt(0)
	reminder := big.NewInt(0)
	compact_balance.QuoRem(balance,divisor,reminder)
	fmt.Printf("Amount of %v tokens: %v.%018s (with decimal point applied)\n",symbol,compact_balance.String(),reminder.String())
}
