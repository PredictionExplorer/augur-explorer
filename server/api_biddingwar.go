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
func api_biddingwar_bid_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_evtlog_id := c.Param("evtlog_id")
	var evtlog_id int64
	if len(p_evtlog_id) > 0 {
		var success bool
		evtlog_id,success = parse_int_from_remote_or_error(c,JSON,&p_evtlog_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'evtlog_id' parameter is not set")
		return
	}
	record_found,bid_info := arb_storagew.Get_bid_info(evtlog_id)
	if !record_found {
		respond_error_json(c,"record not found")
	} else {
		var req_status int = 1
		var err_str string = ""
		c.JSON(http.StatusOK, gin.H{
			"status": req_status,
			"error" : err_str,
			"BidInfo" : bid_info,
		})
	}
} 
func api_biddingwar_prize_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_prize_num:= c.Param("prize_num")
	var prize_num int64
	if len(p_prize_num) > 0 {
		var success bool
		prize_num,success = parse_int_from_remote_or_error(c,JSON,&p_prize_num)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'prize_num' parameter is not set")
		return
	}
	record_found,prize_info := arb_storagew.Get_prize_info(prize_num)
	if !record_found {
		respond_error_json(c,"record not found")
	} else {
		var req_status int = 1
		var err_str string = ""
		c.JSON(http.StatusOK, gin.H{
			"status": req_status,
			"error" : err_str,
			"PrizeInfo" : prize_info,
		})
	}
} 
func api_biddingwar_user_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		respond_error_json(c,"Provided address wasn't found")
		return
	}

	found, user_info := arb_storagew.Get_user_info(user_aid)
	if !found {
		respond_error_json(c,"Provided address wasn't found")
		return
	}
	bids := arb_storagew.Get_bids_by_user(user_aid)
	prizes := arb_storagew.Get_prize_claims_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserInfo" : user_info,
		"Bids" : bids,
		"Prizes" : prizes,
	})
}
