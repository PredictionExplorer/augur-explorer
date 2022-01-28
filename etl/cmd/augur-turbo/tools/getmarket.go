// Dumps Market struct of AbstractMarketFactory contract
package main

import (
	"os"
	"fmt"
	"math/big"
	"strconv"

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
		fmt.Printf("Usage: \n\t\t%v [factory_addr] [market_id]\n\n\t\tShows 'Market' sturct given marketID",os.Args[0])
		os.Exit(1)
	}
	factory_addr := common.HexToAddress(os.Args[1])
	market_id,err := strconv.ParseInt(os.Args[2],10,64)
	if err != nil {
		fmt.Printf("Error parsing market_id field: %v\n",err)
		os.Exit(1)
	}
	factory,err := NewSportsLinkMarketFactory(factory_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate Market Factory contract: %v\n",err)
		os.Exit(1)
	}

	var copts = new(bind.CallOpts)
	big_market_id := big.NewInt(market_id)
	market_obj,err:=factory.GetMarket(copts,big_market_id)
	if err!=nil {
		fmt.Printf("Error during call: %v\n",err)
		os.Exit(1)
	}

	is_resolved,err:=factory.IsMarketResolved(copts,big_market_id)
	if err!=nil {
		fmt.Printf("Error during IsResolved() call: %v\n",err)
		os.Exit(1)
	}
	sharefactor,err:=factory.ShareFactor(copts)
	if err!=nil {
		fmt.Printf("Error during ShareFactor() call: %v\n",err)
		os.Exit(1)
	}
	feepot,err:=factory.FeePot(copts)
	if err!=nil {
		fmt.Printf("Error during FeePot() call: %v\n",err)
		os.Exit(1)
	}
	collateral,err:=factory.Collateral(copts)
	if err!=nil {
		fmt.Printf("Error during Collateral() call: %v\n",err)
		os.Exit(1)
	}


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
	fmt.Printf("\tIsResolved: %v\n",is_resolved)
	fmt.Printf("\tShareFactor: %v\n",sharefactor)
	fmt.Printf("\tFeePot: %v\n",feepot.String())
	fmt.Printf("\tCollateral: %v\n",collateral.String())

	fmt.Printf("}\n")
}
