// Shows allowance of a user in an ERC20 token
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

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Printf("Usage: \n\t\t%v [contract_addr] [owner_addr] [spender_addr]\n\n\t\tShows allowance of ERC20 token\n",os.Args[0])
		os.Exit(1)
	}
	contract_addr := common.HexToAddress(os.Args[1])
	owner_addr := common.HexToAddress(os.Args[2])
	spender_addr := common.HexToAddress(os.Args[3])
	erc20_token,err := NewERC20(contract_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate ERC20 contract: %v\n",err)
		os.Exit(1)
	}

	var copts = new(bind.CallOpts)
	allowance,err:=erc20_token.Allowance(copts,owner_addr,spender_addr)
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
	fmt.Printf("Amount of %v allowed tokens: %v\n",symbol,allowance.String())
	compact_allowance := big.NewInt(0)
	reminder := big.NewInt(0)
	compact_allowance.QuoRem(allowance,divisor,reminder)
	fmt.Printf("Amount of %v allowed tokens: %v.%018s (with decimal point applied)\n",symbol,compact_allowance.String(),reminder.String())
}
