package main
import (
	"fmt"
	"time"
	"math/big"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	bwt "github.com/PredictionExplorer/augur-explorer/primitives/biddingwar" // bidding war types
	. "github.com/PredictionExplorer/augur-explorer/dbs/biddingwar" // bidding war types
	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/primitives/biddingwar"
)
const (
	CONTRACT_CONSTANTS_REFRESH_TIME		= 5*60	// seconds
)
var (
	biddingwar_addr				common.Address
	cosmic_signature_addr		common.Address
	cosmic_token_addr			common.Address
	charity_wallet_addr			common.Address

	// contract constants (variables not frequently modified, and only by the owner)
	price_increase				*big.Int = big.NewInt(0)
	charity_addr				common.Address
	charity_percentage			*big.Int = big.NewInt(0)
	token_reward				*big.Int = big.NewInt(0)
	prize_percentage			*big.Int = big.NewInt(0)
	time_increase				*big.Int

	// contract variables (variables usually modified by a bid())
	bid_price					*big.Int = big.NewInt(0)
	bid_price_eth				float64
	prize_claim_date			*big.Int = big.NewInt(0)	// timestamp (Unix)
	prize_amount				*big.Int = big.NewInt(0)
	prize_amount_eth			float64
	num_prizes					*big.Int = big.NewInt(0)
	total_prizes_amount_paid	*big.Int = big.NewInt(0)
	nanoseconds_extra			*big.Int = big.NewInt(0)
	last_bidder					common.Address
	charity_balance				*big.Int = big.NewInt(0)
	charity_balance_eth			float64

	// contract counters	(collected via DB)
	/* DISCONTINUED
	num_voluntary_donations		uint64
	num_rwalk_tokens_used		uint64
	total_bids					uint64
	num_unique_bidders			uint64
	num_unique_winners			uint64
	total_prizes_paid_eth		float64	// in ETH
	total_prizes_paid_wei		string
	*/
	bw_stats					BWStatistics

	arb_storagew				SQLStorageWrapper
)
func biddingwar_init() {

	if  !augur_srv.arbitrum_initialized() {
		err_str := fmt.Sprintf("biddingwar_init(): Database link wasn't configured")
		Info.Printf(err_str)
		Error.Printf(err_str)
	}
	arb_storagew.S=augur_srv.db_augur
	arb_storagew.S.Db_set_schema_name("public")
	bw_caddrs := arb_storagew.Get_biddingwar_contract_addrs()
	biddingwar_addr = common.HexToAddress(bw_caddrs.BiddingWarAddr)
	cosmic_signature_addr = common.HexToAddress(bw_caddrs.CosmicSignatureAddr)
	cosmic_token_addr = common.HexToAddress(bw_caddrs.CosmicSignatureTokenAddr)
	charity_wallet_addr = common.HexToAddress(bw_caddrs.CharityWalletAddr)
	do_reload_contract_variables()
	do_reload_database_variables()
	go reload_constants_goroutine()
}
func do_reload_contract_constants() {
	var copts bind.CallOpts
	bwcontract,err := NewBiddingWar(biddingwar_addr,rpcclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate BiddingWar contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
	} else {
		var err error
		price_increase,err = bwcontract.PriceIncrease(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PriceIncrease() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		charity_addr,err = bwcontract.Charity(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at Charity() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		charity_percentage,err = bwcontract.CharityPercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at Charity() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		token_reward,err = bwcontract.TokenReward(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at TokenReward() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		prize_percentage,err = bwcontract.PrizePercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizePercentage() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		time_increase,err = bwcontract.TimeIncrease(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at TimeIncrease() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
	}
}
func do_reload_contract_variables() {
	var copts bind.CallOpts
	bwcontract,err := NewBiddingWar(biddingwar_addr,rpcclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate BiddingWar contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
	} else {
		f_divisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
		bid_price,err = bwcontract.GetBidPrice(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at GetBidPrice() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		} else {
			f_bid_price := big.NewFloat(0.0).SetInt(bid_price)
			f_quo := big.NewFloat(0.0).Quo(f_bid_price,f_divisor)
			bid_price_eth,_ = f_quo.Float64()
		}
		prize_claim_date ,err = bwcontract.PrizeTime(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizeTime() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		prize_amount, err = bwcontract.PrizeAmount(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizeAmount() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		} else {
			f_prize_amount:= big.NewFloat(0.0).SetInt(prize_amount)
			f_quo := big.NewFloat(0.0).Quo(f_prize_amount,f_divisor)
			prize_amount_eth,_ = f_quo.Float64()
		}
		num_prizes, err = bwcontract.NumPrizes(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumPrizes() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		nanoseconds_extra,err = bwcontract.NanoSecondsExtra(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NanoSecondsExtra() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		last_bidder,err = bwcontract.LastBidder(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at LastBidder() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		charity_balance,err = rpcclient.BalanceAt(context.Background(),charity_addr,nil)
		if err != nil {
			err_str := fmt.Sprintf("Error at BalanceAt() call for charity addr: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		} else {
			f_charity_balance := big.NewFloat(0.0).SetInt(charity_balance)
			f_quo := big.NewFloat(0.0).Quo(f_charity_balance,f_divisor)
			charity_balance_eth,_ = f_quo.Float64()
		}

	}
}
func do_reload_database_variables() {
	/* DISCONTINUED
	// fetches accumulators (statistics) and other counters calculated by the DB but not by the contract
	// total_bids :=
	// num_unique_bidders :=
	// num_unique_winners :=
	// total_prizes_paid :=
	// num_voluntary_donations :=
	// num_rwalk_tokens_used :=
	*/
	bw_stats = arb_storagew.Get_biddingwar_statistics()
}
func reload_constants_goroutine() {
	// we will load contract constants up web requests but to avoid having to restart
	// the server every time a constant changes we will have a refresh of the constants
	// every few minutes
	for {
		do_reload_contract_constants()
		time.Sleep(CONTRACT_CONSTANTS_REFRESH_TIME * time.Second)
	}
}
func biddingwar_index_page(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	c.HTML(http.StatusOK, "bw_index.html", gin.H{
		"BiddingWarAddr":biddingwar_addr,
		"CosmicSignatureAddr":cosmic_signature_addr,
		"CosmicSignatureTokenAddr":cosmic_token_addr,
		"CharityWalletAddr":charity_wallet_addr,
		"BidPrice":bid_price.String(),
		"BidPriceEth":bid_price_eth,
		"PrizeClaimDate":time.Unix(prize_claim_date.Int64(),0).Format(time.RFC822),
		"PrizeClaimTs":prize_claim_date.Int64(),
		"CurRoundNum":num_prizes.Int64()+1,
		"CurNumBids" : bw_stats.CurNumBids,
		"PrizeAmount" : prize_amount.Int64(),
		"PrizeAmountEth" : prize_amount_eth,
		"TotalPrizes": bw_stats.TotalPrizes,
		"TotalPrizesPaidAmountEth": bw_stats.TotalPrizesPaidAmountEth,
		"LastBidderAddr":last_bidder.String(),
		"NumVoluntaryDonations":bw_stats.NumVoluntaryDonations,
		"SumVoluntaryDonationsEth":bw_stats.SumVoluntaryDonationsEth,
		"NumRwalkTokensUsed":bw_stats.NumRwalkTokensUsed,
		"PriceIncrease" : price_increase.Int64(),
		"TimeIncrease" : time_increase.Int64(),
		"NanosecondsExtra" : nanoseconds_extra.Int64(),
		"TokenReward" : token_reward.Int64(),
		"PrizePercentage" : prize_percentage.Int64(),
		"CharityAddr" : charity_addr.String(),
		"CharityPercentage" : charity_percentage.Int64(),
		"CharityBalance": charity_balance.String(),
		"CharityBalanceEth": charity_balance_eth,
	})
}
