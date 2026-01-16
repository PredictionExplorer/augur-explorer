// Makes a bid
package main

import (
	"os"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	CHAIN_ID		int64 = 31337
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

	if len(os.Args) != 2 {
		fmt.Printf("Usage: \n\t\t%v [contract_addr]\n\n\t\tMakes bid using CST tokens\n",os.Args[0])
		os.Exit(1)
	}

	cosmic_game_addr := common.HexToAddress(os.Args[1])

	cosmic_game_ctrct,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
		os.Exit(1)
	}

	var copts bind.CallOpts
	cst_price,err := cosmic_game_ctrct.GetCurrentBidPriceCST(&copts)
	if err != nil {
		fmt.Printf("Error at currentCSTPrice()(): %v\n",err)
		fmt.Printf("Aborting\n")
		os.Exit(1)
	}

	fmt.Printf("CSt price = %v\n",cst_price.String())
	mult := big.NewInt(2)
	cst_price.Mul(cst_price,mult)
	fmt.Printf("CST bid amount: %v\n",cst_price.String())

	f_divisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
	f_bid_price := big.NewFloat(0.0).SetInt(cst_price)
	f_quo := big.NewFloat(0.0).Quo(f_bid_price,f_divisor)
	bid_price_eth,_ := f_quo.Float64()

	fmt.Printf("Bid price in CST ETH = %v\n",bid_price_eth)
}
