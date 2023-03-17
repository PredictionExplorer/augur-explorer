package main
import (
	"time"

	"net/http"
	"github.com/gin-gonic/gin"

)
func api_biddingwar_dashboard(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
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
func api_biddingwar_prize_claims(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	prizes := arb_storagew.Get_prize_claims(offset,limit)


	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"PrizeClaims" : prizes,
	})
}
func api_biddingwar_bids(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	bids := arb_storagew.Get_bids(offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Bids" : bids,
		"Offset" : offset,
		"Limit" : limit,
	})
}
