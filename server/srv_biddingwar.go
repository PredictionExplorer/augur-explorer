package main
import (
	"fmt"
	"time"
	"os"
	"strconv"
	"encoding/csv"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/abi/bind"
	bwt "github.com/PredictionExplorer/augur-explorer/primitives/biddingwar" // bidding war types
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (
	CONTRACT_CONSTANTS_REFRESH_TIME		= 5*60	// seconds
)
var (
	bw_caddrs					bwt.BiddingWarContractAddrs

	// contract constants (variables not frequently modified, and only by the owner)
	price_increase				*big.Int
	charity_addr				common.Address
	charity_percentage			*big.Int
	token_reward				*big.Int
	prize_percentage			*big.Int
	time_increase				*big.Int

	// contract variables (variables usually modified by a bid())
	bid_price					*big.Int
	prize_claim_date			*big.Int	// timestamp (Unix)
	prize_amount				*big.Int
	num_prizes					*big.Int
	total_prizes_amount_paid	*big.Int
	nanoseconds_extra			*big.Int
	last_bidder					common.Address
	charity_balance				*big.Int

	// contract counters	(collected via DB)
	num_voluntary_donations		uint64
	num_rwalk_tokens_used		uint64
	total_bids					uint64
	num_unique_bidders			uint64
	num_unique_winners			uint64
	total_prizes_paid			float64	// in ETH
)
func biddingwar_init() {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"biddingwar_init(): Database link wasn't configured")
		return
	}
	bw_caddrs = augur_srv.db_arbitrum.Get_biddingwar_contract_addrs()
	do_reload_contract_variables()
	go reload_constants_goroutine()
}
func do_reload_contract_constants() {
	bwcontract,err := NewBiddingWar(bw_caddrs.BiddingWarAddr,rpcclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate BiddingWar contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
	} else {
		price_increase,err := bwcontract.PriceIncrease(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PriceIncrease() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		charity_addr,err := bwcontract.Charity(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at Charity() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		charity_percentage,err := bwcontract.CharityPercentage(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at Charity() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		token_reward,err := bwcontract.TokenReward(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at TokenReward() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		prize_percentage,err := bwcontract.PrizePercentage(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizePercentage() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		time_increase,err := bwcontract.TimeIncrease(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at TimeIncrease() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
	}
}
func do_reload_contract_variables() {
	bwcontract,err := NewBiddingWar(bw_caddrs.BiddingWarAddr,rpcclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate BiddingWar contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
	} else {
		bid_price,err := bwcontract.GetBidPrice(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at GetBidPrice() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		prize_claim_date ,err := bwcontract.PrizeTime(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizeTime() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		prize_amount, err := bwcontract.PrizeAmount(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizeAmount() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		num_prizes, err := bwcontract.NumPrizes(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumPrizes() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		nanoseconds_extra,err := bwcontract.NanoSecondsExtra(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NanoSecondsExtra() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		last_bidder,err := bwcontract.LastBidder(copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at LastBidder() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		charity_balance,err := rpcclient.BalanceAt(charity_addr)
		if err != nil {
			err_str := fmt.Sprintf("Error at BalanceAt() call for charity addr: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
	}
}
func do_reload_database_variables() {
	// fetches accumulators (statistics) and other counters calculated by the DB but not by the contract
	// total_bids :=
	// num_unique_bidders :=
	// num_unique_winners :=
	// total_prizes_paid :=
	// num_voluntary_donations :=
	// num_rwalk_tokens_used :=
}
func reload_constants_goroutine() {
	// we will load contract constants up web requests but to avoid having to restart
	// the server every time a constant changes we will have a refresh of the constants
	// every few minutes
	copts bind.CallOpts
	for {
		do_reload_contract_constants()
		time.Sleep(CONTRACT_CONSTANTS_REFRESH_TIME * time.Second)
	}
}
func biddingwar_index_page(c *gin.Context) {

	c.HTML(http.StatusOK, "rw_index.html", gin.H{
		"ContractAddresses":caddrs,
	})
}
