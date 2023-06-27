package main
import (
	"time"
	"os"
	"fmt"
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common"

	"net/http"
	"github.com/gin-gonic/gin"

)
func api_biddingwar_dashboard(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	caddrs := arb_storagew.Get_cosmic_game_contract_addrs()
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"BiddingWarAddr":cosmic_game_addr,
		"CosmicSignatureAddr":cosmic_signature_addr,
		"CosmicSignatureTokenAddr":cosmic_token_addr,
		"CharityWalletAddr":charity_wallet_addr,
		"BidPrice":bid_price,
		"BidPriceEth":bid_price_eth,
		"PrizeClaimDate":time.Unix(prize_claim_date,0).Format(time.RFC822),
		"PrizeClaimTs":prize_claim_date,
		"CurRoundNum":round_num+1,
		"CurNumBids" : bw_stats.CurNumBids,
		"PrizeAmount" : prize_amount,
		"PrizeAmountEth" : prize_amount_eth,
		"TotalPrizes": bw_stats.TotalPrizes,
		"TotalPrizesPaidAmountEth": bw_stats.TotalPrizesPaidAmountEth,
		"LastBidderAddr":last_bidder.String(),
		"NumVoluntaryDonations":bw_stats.NumVoluntaryDonations,
		"SumVoluntaryDonationsEth":bw_stats.SumVoluntaryDonationsEth,
		"NumRwalkTokensUsed":bw_stats.NumRwalkTokensUsed,
		"PriceIncrease" : price_increase,
		"TimeIncrease" : time_increase,
		"NanosecondsExtra" : nanoseconds_extra,
		"TokenReward" : token_reward,
		"PrizePercentage" : prize_percentage,
		"RafflePercentage" : raffle_percentage,
		"NumRaffleEthWinners" : raffle_eth_winners,
		"NumRaffleNFTWinners" : raffle_nft_winners,
		"CharityAddr" : charity_addr.String(),
		"CharityPercentage" : charity_percentage,
		"CharityBalance": charity_balance,
		"CharityBalanceEth": charity_balance_eth,
		"NumDonatedNFTs" : bw_stats.NumDonatedNFTs,
		"ContractAddrs" : caddrs,
		"MainStats" : bw_stats,
		"TsRoundStart" : round_start_ts,
	})
}
func api_biddingwar_prize_list(c *gin.Context) {

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
func api_biddingwar_bid_list(c *gin.Context) {

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
func api_biddingwar_bid_list_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,JSON,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'round_num' parameter is not set")
		return
	}
	p_sort:= c.Param("sort")
	var sort int64
	if len(p_sort) > 0 {
		var success bool
		sort,success = parse_int_from_remote_or_error(c,JSON,&p_sort)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'sort' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	bids := arb_storagew.Get_bids_by_round_num(round_num,int(sort),offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RoundNum" : round_num,
		"Offset" : offset,
		"Limit" : limit,
		"Sort" : sort,
		"BidsByRound" : bids,
	})
}
func api_biddingwar_bid_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error_json(c,"'user_addr' parameter is not set")
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
func api_biddingwar_charity_donations(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	biddingwar_aid,err :=arb_storagew.S.Nonfatal_lookup_address_id(cosmic_game_addr.String())
	if err != nil {
		Error.Printf("BiddingWar contract address doesn't exist in the DB, aborting server")
		os.Exit(1)
	}

	donations := arb_storagew.Get_charity_donations(biddingwar_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CharityDonations" : donations,
	})
}
func api_biddingwar_donations_eth(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	donations := arb_storagew.Get_donations_to_biddingwar()
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CharityDonations" : donations,
	})
}
func api_biddingwar_user_unique_bidders(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	unique_bidders := arb_storagew.Get_unique_bidders()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UniqueBidders" : unique_bidders,
	})
}
func api_biddingwar_user_unique_winners(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	unique_winners := arb_storagew.Get_unique_winners()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UniqueWinners" : unique_winners,
	})
}
func api_biddingwar_donations_nft_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	nft_donations := arb_storagew.Get_NFT_donations(offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"NFTDonations" : nft_donations,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func api_biddingwar_nft_donation_stats(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	nft_donation_stats := arb_storagew.Get_NFT_donation_stats()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"NFTDonationStats" : nft_donation_stats,
	})
}
func api_biddingwar_record_counters(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	record_counters := arb_storagew.Get_record_counters()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RecordCounters" : record_counters,
	})
}
func api_biddingwar_donated_nft_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_record_id:= c.Param("record_id")
	var record_id int64
	if len(p_record_id) > 0 {
		var success bool
		record_id,success = parse_int_from_remote_or_error(c,JSON,&p_record_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'record_id' parameter is not set")
		return
	}
	found,nftdonation := arb_storagew.Get_NFT_donation_info(record_id)
	var req_status int = 1
	var err_str string = ""
	if !found {
		respond_error_json(c,"Record not found")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": req_status,
			"error" : err_str,
			"NFTDonation" : nftdonation,
		})
	}
}
func api_biddingwar_raffle_deposits_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	deposits := arb_storagew.Get_raffle_eth_deposits_list(offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RaffleDeposits" : deposits,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func api_biddingwar_raffle_deposits_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,JSON,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'round_num' parameter is not set")
		return
	}

	deposits := arb_storagew.Get_raffle_deposits_by_round(round_num)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RaffleDeposits" : deposits,
		"RoundNum" : round_num,
	})
}
func api_biddingwar_raffle_nft_winners_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	winners := arb_storagew.Get_raffle_nft_winners(offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RaffleNFTWinners" : winners,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func api_biddingwar_raffle_nft_winners_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,JSON,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'round_num' parameter is not set")
		return
	}

	winners := arb_storagew.Get_raffle_nft_winners_by_round(round_num)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RaffleNFTWinners" : winners,
		"RoundNum" : round_num,
	})
}
func api_biddingwar_user_raffle_nft_winnings(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error_json(c,"'user_addr' parameter is not set")
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

	winnings := arb_storagew.Get_raffle_nft_winnings_by_user(user_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserRaffleNFTWinnings" : winnings,
		"UserInfo" : user_info,
	})
}
func api_biddingwar_raffle_nft_claims(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	claims := arb_storagew.Get_raffle_nft_claims(offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RaffleNFTClaims" : claims,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func api_biddingwar_user_raffle_deposits(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error_json(c,"'user_addr' parameter is not set")
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

	deposits := arb_storagew.Get_raffle_deposits_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserRaffleDeposits" : deposits,
		"UserInfo" : user_info,
	})
}
func api_biddingwar_user_raffle_nft_claims(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error_json(c,"'user_addr' parameter is not set")
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

	deposits := arb_storagew.Get_raffle_nft_claims_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserRaffleNFTClaims" : deposits,
		"UserInfo" : user_info,
	})
}
func api_biddingwar_nft_donations_by_prize(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_prize_num := c.Param("prize_num")
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
	nft_donations := arb_storagew.Get_nft_donations_by_prize(prize_num)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"NFTDonations" : nft_donations,
		"PrizeNum": prize_num,
	})
}
func api_biddingwar_cosmic_signature_token_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	tokens := arb_storagew.Get_cosmic_signature_nft_list(offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CosmicSignatureTokenList" : tokens,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func api_biddingwar_cosmic_signature_token_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_token_id:= c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id,success = parse_int_from_remote_or_error(c,JSON,&p_token_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'token_id' parameter is not set")
		return
	}

	record_found,token_info := arb_storagew.Get_cosmic_signature_token_info(token_id)
	if !record_found {
		respond_error_json(c,"record not found")
		return
	}

	var req_status int = 1
	var err_str string = ""

	if token_info.PrizeNum > -1 {
		_,prize_info := arb_storagew.Get_prize_info(token_info.PrizeNum)
		c.JSON(http.StatusOK, gin.H{
			"status": req_status,
			"error" : err_str,
			"TokenInfo" : token_info,
			"PrizeInfo" : prize_info,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": req_status,
			"error" : err_str,
			"TokenInfo" : token_info,
		})
	}
}
func api_biddingwar_donated_nft_claims_all(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	claims := arb_storagew.Get_donated_nft_claims(offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"DonatedNFTClaims" : claims,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func api_biddingwar_donated_nft_claims_by_user(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error_json(c,"'user_addr' parameter is not set")
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
	claims := arb_storagew.Get_donated_nft_claims_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"DonatedNFTClaims" : claims,
		"UserInfo" : user_info,
	})
}
func api_biddingwar_time_current(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	var raw json.RawMessage
	err := rpcclient.CallContext(context.Background(), &raw,"eth_getBlockByNumber", "pending",true)
	if err != nil {
		respond_error_json(c,fmt.Sprintf("%v",err))
		return
	}
	var rpcobj map[string]interface{}
	rpcobj = make(map[string]interface{})
	err = json.Unmarshal(raw,&rpcobj)
	if err != nil {
		respond_error_json(c,fmt.Sprintf("Error decoding JSON: %v",err))
		return
	}

	ts_hex := rpcobj["timestamp"].(string)
	ts,err := hexutil.DecodeUint64(ts_hex)
	if err !=nil {
		respond_error_json(c,fmt.Sprintf("Error decoding timestamp from hex: %v",err))
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CurrentTimeStamp": ts,
	})
}
func api_biddingwar_time_until_prize(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	const time_until_prize_sig string = "0x8b1329e0"
	var raw json.RawMessage
	args := map[string]interface{}{
		"to": cosmic_game_addr.String(),
		"data":time_until_prize_sig,
	}
	err := rpcclient.CallContext(context.Background(), &raw,"eth_call", args,"pending")
	if err != nil {
		respond_error_json(c,fmt.Sprintf("%v",err))
		return
	}
	var ts_hex string
	err = json.Unmarshal(raw,&ts_hex)
	if err != nil {
		respond_error_json(c,fmt.Sprintf("Error decoding JSON: %v",err))
		return
	}
	ts_big := common.HexToHash(ts_hex).Big()
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"TimeUntilPrize": ts_big.Int64(),
	})
}
