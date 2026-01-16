// fetches bidder's last bid timestamp
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
)
func fmt_eth(wei *big.Int) string {
    ether := new(big.Float).SetInt(wei)
    eth_value := new(big.Float).Quo(ether, big.NewFloat(1e18))
    return eth_value.Text('f', 18) // 18 decimal places to match Ethereum precision
}
func convert_to_percentage(in *big.Int) (float64) {

	one := big.NewFloat(1)
	hundred := big.NewFloat(100)
	divisor_float := new(big.Float).SetInt(in)
	increase_fraction := new(big.Float).Quo(one,divisor_float)
	increase_percent := new(big.Float).Mul(increase_fraction, hundred)
	out,_ := increase_percent.Float64()
	return out
}
func main() {

	RPC_URL = os.Getenv("RPC_URL")
	eclient, err := ethclient.Dial(RPC_URL)
	if err!=nil {
		fmt.Printf("Can't connect to ETH RPC: %v\n",err)
		os.Exit(1)
	}

	var cg_addr string
	if len(os.Args) < 4 {
		fmt.Printf(
			"Usage: \n\t\t%v [cosmic_game_addr] [round_num] [bidder_address]\n\t\t"+
			"Gets bid timestamp (queried by bid position)\n\n",os.Args[0],
		)
		return
	} else {
		cg_addr = os.Args[1]
	}

    num_str := os.Args[2]
    round,err := strconv.ParseInt(num_str,10,64)
    if err != nil {
        fmt.Printf("error parsing round parameter: %v\n",err)
        os.Exit(1)
    }   
	bidder_addr := common.HexToAddress(os.Args[3])

	var copts bind.CallOpts
	cosmic_game_addr := common.HexToAddress(cg_addr)
	fmt.Printf("Calling to contract at %v\n",cosmic_game_addr.String())

	cosmic_game_ctrct,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err!=nil {
		fmt.Printf("Failed to instantiate CosmicGame contract: %v\n",err)
		os.Exit(1)
	}

	binfo,err := cosmic_game_ctrct.BiddersInfo(&copts,big.NewInt(round),bidder_addr)
	if err!=nil {
		fmt.Printf("Failed to get biddersInfo(round,bidder_addr): %v\n",err)
		os.Exit(1)
	}
	
	fmt.Printf("for bid at round %v , for bidder %v :\n",round,bidder_addr)
	fmt.Printf("totalSpendEthAmount = %v (%v ETH)\n",binfo.TotalSpentEthAmount.String(),fmt_eth(binfo.TotalSpentEthAmount))
	fmt.Printf("totalSpentCstAmount = %v (%v ETH)\n",binfo.TotalSpentCstAmount.String(),fmt_eth(binfo.TotalSpentCstAmount))
	fmt.Printf("lastBidTimeStamp = %v\n",binfo.LastBidTimeStamp);
}
