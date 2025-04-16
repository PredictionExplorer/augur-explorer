package main
import (
	"time"
	"os"
	"fmt"
	"math/big"
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"net/http"
	"github.com/gin-gonic/gin"

	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
func api_cosmic_game_dashboard(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	caddrs := arb_storagew.Get_cosmic_game_contract_addrs()
	cur_round_stats := arb_storagew.Get_cosmic_game_round_statistics(round_num);
	cg_balance := get_cosmic_game_contract_balance()
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CosmicGameAddr": cosmic_game_addr,
		"CosmicGameBalanceEth": cg_balance,
		"CosmicSignatureAddr": cosmic_signature_addr,
		"CosmicSignatureTokenAddr": cosmic_token_addr,
		"CharityWalletAddr": charity_wallet_addr,
		"BidPrice": bid_price,
		"BidPriceEth": bid_price_eth,
		"PrizeClaimDate": time.Unix(prize_claim_date,0).Format(time.RFC822),
		"PrizeClaimTs": prize_claim_date,
		"CurRoundNum": round_num,
		"CurNumBids" : bw_stats.CurNumBids,
		"PrizeAmount" : prize_amount,
		"PrizeAmountEth" : prize_amount_eth,
		"RaffleAmount" : raffle_amount,
		"RaffleAmountEth" : raffle_amount_eth,
		"StakingAmount" : staking_amount,
		"StakingAmountEth" : staking_amount_eth,
		"TotalPrizes": bw_stats.TotalPrizes,
		"TotalPrizesPaidAmountEth": bw_stats.TotalPrizesPaidAmountEth,
		"TotalEthDonatedAmount" : bw_stats.TotalEthDonatedAmount,
		"TotalEthDonatedAmountEth" : bw_stats.TotalEthDonatedAmountEth,
		"LastBidderAddr":last_bidder.String(),
		"NumVoluntaryDonations":bw_stats.NumVoluntaryDonations,
		"SumVoluntaryDonationsEth":bw_stats.SumVoluntaryDonationsEth,
		"NumRwalkTokensUsed":bw_stats.NumRwalkTokensUsed,
		"PriceIncrease" : price_increase,
		"TimeIncrease" : time_increase,
		"MainPrizeTimeIncrementInMicroSeconds" : mainprize_microseconds_inc,
		"InitialSecondsUntilPrize" : initial_seconds,
		"TimeoutClaimPrize" : timeout_claim,
		"RoundStartCSTAuctionLength" : roundstart_auclen,
		"TokenReward" : token_reward,
		"PrizePercentage" : prize_percentage,
		"RafflePercentage" : raffle_percentage,
		"StakignPercentage" : staking_percentage,
		"CharityAddr" : charity_addr.String(),
		"CharityPercentage" : charity_percentage,
		"CharityBalance": charity_balance,
		"CharityBalanceEth": charity_balance_eth,
		"NumRaffleEthWinnersBidding" : raffle_eth_winners_bidding,
		"NumRaffleNFTWinnersBidding" : raffle_nft_winners_bidding,
		"NumRaffleNFTWinnersStakingRWalk" : raffle_nft_winners_staking_rwalk,
		"NumUniqueBidders" :  bw_stats.NumUniqueBidders,
		"NumUniqueWinners" : bw_stats.NumUniqueWinners,
		"NumUniqueDonors" : bw_stats.NumUniqueDonors,
		"NumUniqueStakersCST" : bw_stats.NumUniqueStakersCST,
		"NumUniqueStakersRWalk" : bw_stats.NumUniqueStakersRWalk,
		"NumUniqueStakersBoth" : bw_stats.NumUniqueStakersBoth,
		"NumDonatedNFTs" : bw_stats.NumDonatedNFTs,
		"MainStats" : bw_stats,
		"CurRoundStats" : cur_round_stats,
		"TsRoundStart" : round_start_ts,
		"ContractAddrs" : caddrs,
	})
}
func api_cosmic_game_prize_list(c *gin.Context) {

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
		"Rounds" : prizes,
	})
}
func api_cosmic_game_bid_list(c *gin.Context) {

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
func api_cosmic_game_bid_list_by_round(c *gin.Context) {

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

	bids,total_rows := arb_storagew.Get_bids_by_round_num(round_num,int(sort),offset,limit)
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
		"TotalRows" : total_rows,
	})
}
func api_cosmic_game_bid_info(c *gin.Context) {

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
func api_cosmic_game_round_info(c *gin.Context) {

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
			"RoundInfo" : prize_info,
		})
	}
} 
func api_cosmic_game_user_info(c *gin.Context) {

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
func api_cosmic_game_charity_cosmicgame_deposits(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	cosmicgame_aid,err :=arb_storagew.S.Nonfatal_lookup_address_id(cosmic_game_addr.String())
	if err != nil {
		Error.Printf("CosmicGame contract address doesn't exist in the DB, aborting server")
		os.Exit(1)
	}

	donations := arb_storagew.Get_charity_donations_from_cosmic_game(cosmicgame_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CharityDonations" : donations,
	})
}
func api_cosmic_game_charity_voluntary_deposits(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	cosmicgame_aid,err :=arb_storagew.S.Nonfatal_lookup_address_id(cosmic_game_addr.String())
	if err != nil {
		Error.Printf("CosmicGame contract address doesn't exist in the DB, aborting server")
		os.Exit(1)
	}

	donations := arb_storagew.Get_charity_donations_voluntary(cosmicgame_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CharityDonations" : donations,
	})
}
func api_cosmic_game_charity_donations_deposits(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	cosmicgame_aid,err :=arb_storagew.S.Nonfatal_lookup_address_id(cosmic_game_addr.String())
	if err != nil {
		Error.Printf("CosmicGame contract address doesn't exist in the DB, aborting server")
		os.Exit(1)
	}

	donations := arb_storagew.Get_charity_donations(cosmicgame_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CharityDonations" : donations,
	})
}
func api_cosmic_game_charity_donations_withdrawals(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	withdrawals := arb_storagew.Get_charity_wallet_withdrawals()
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CharityWithdrawals" : withdrawals,
	})
}
func api_cosmic_game_user_unique_bidders(c *gin.Context) {

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
func api_cosmic_game_user_unique_winners(c *gin.Context) {

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
func api_cosmic_game_user_unique_donors(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	unique_donors := arb_storagew.Get_unique_donors()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UniqueDonors" : unique_donors,
	})
}
func api_cosmic_game_donations_nft_list(c *gin.Context) {

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
func api_cosmic_game_nft_donation_stats(c *gin.Context) {

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
func api_cosmic_game_record_counters(c *gin.Context) {

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
func api_cosmic_game_donated_nft_info(c *gin.Context) {

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
func api_cosmic_game_prize_deposits_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	deposits := arb_storagew.Get_prize_eth_deposits_list(offset,limit)

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
func api_cosmic_game_prize_deposits_by_round(c *gin.Context) {

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

	deposits := arb_storagew.Get_prize_deposits_by_round(round_num)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RaffleDeposits" : deposits,
		"RoundNum" : round_num,
	})
}
func api_cosmic_game_raffle_nft_winners_list(c *gin.Context) {

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
func api_cosmic_game_raffle_nft_winners_by_round(c *gin.Context) {

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

	winners_raffle := arb_storagew.Get_raffle_nft_winners_by_round(round_num,false)
	winners_staking := arb_storagew.Get_raffle_nft_winners_by_round(round_num,true)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RaffleNFTWinners" : winners_raffle,
		"StakingNFTWinners" : winners_staking,
		"RoundNum" : round_num,
	})
}
func api_cosmic_game_user_raffle_nft_winnings(c *gin.Context) {

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
func api_cosmic_game_prize_deposits_raffle_eth_by_user(c *gin.Context) {

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

	deposits := arb_storagew.Get_prize_deposits_raffle_eth_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserRaffleDeposits" : deposits,
		"UserInfo" : user_info,
	})
}
func api_cosmic_game_prize_deposits_chrono_warrior_by_user(c *gin.Context) {

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

	deposits := arb_storagew.Get_prize_deposits_chrono_warrior_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserChronoWarriorDeposits" : deposits,
		"UserInfo" : user_info,
	})
}
func api_cosmic_game_nft_donations_by_prize(c *gin.Context) {

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
		"RoundNum": prize_num,
	})
}
func api_cosmic_game_cosmic_signature_token_list(c *gin.Context) {

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
func api_cosmic_game_cosmic_signature_token_info(c *gin.Context) {

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

	if token_info.RecordType == 3 {
		_,prize_info := arb_storagew.Get_prize_info(token_info.RoundNum)
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
func api_cosmic_game_donated_nft_claims_all(c *gin.Context) {

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
func api_cosmic_game_donated_nft_claims_by_user(c *gin.Context) {
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
func api_cosmic_game_time_current(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
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
func api_cosmic_game_time_until_prize(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
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
func api_cosmic_game_prize_cur_round_time(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	var copts bind.CallOpts
	bwcontract,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err != nil {
		respond_error_json(c,fmt.Sprintf("Error during call: can't instantiate CG contract: %v",err))
		return
	}
	prize_time,err := bwcontract.MainPrizeTime(&copts)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Error during call: %v",err))
		return
	} else {
		var req_status int = 1
		var err_str string = ""
		c.JSON(http.StatusOK, gin.H{
			"status": req_status,
			"error" : err_str,
			"CurRoundPrizeTime" : prize_time,
		})
	}
} 
func api_cosmic_game_user_global_winnings(c *gin.Context) {

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

	claim_info := arb_storagew.Get_user_notif_red_box_rewards(user_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Winnings" : claim_info,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
	})
}
func api_cosmic_game_prize_history_detail_by_user(c *gin.Context) {

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
	found, _ := arb_storagew.Get_user_info(user_aid)
	if !found {
		respond_error_json(c,"Provided address wasn't found")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	claim_history := arb_storagew.Get_prize_history_detailed_by_user(user_aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"USerPrizeHistory" : claim_history,
	})
}
func api_cosmic_game_global_claim_history_detail(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	claim_history := arb_storagew.Get_claim_history_detailed_global(offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"GlobalPrizeHistory" : claim_history,
	})
}
func api_cosmic_game_unclaimed_donated_nfts_by_user(c *gin.Context) {

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
	found, _ := arb_storagew.Get_user_info(user_aid)
	if !found {
		respond_error_json(c,"Provided address wasn't found")
		return
	}

	nfts := arb_storagew.Get_unclaimed_donated_nft_by_user(user_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UnclaimedDonatedNFTs" : nfts,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
	})
}
func api_cosmic_game_unclaimed_donated_nfts_by_prize(c *gin.Context) {

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


	nft_donations := arb_storagew.Get_unclaimed_donated_nfts_by_prize(prize_num)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"NFTDonations" : nft_donations,
		"RoundNum": prize_num,
	})
}
func api_cosmic_game_unclaimed_prize_deposits_by_user(c *gin.Context) {

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
	found, _ := arb_storagew.Get_user_info(user_aid)
	if !found {
		respond_error_json(c,"Provided address wasn't found")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	deposits := arb_storagew.Get_unclaimed_prize_eth_deposits(user_aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"UnclaimedDeposits" : deposits,
	})
}
func api_cosmic_game_cosmic_signature_token_list_by_user(c *gin.Context) {

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
	found, _ := arb_storagew.Get_user_info(user_aid)
	if !found {
		respond_error_json(c,"Provided address wasn't found")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	user_tokens := arb_storagew.Get_cosmic_signature_nft_list_by_user(user_aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"UserTokens" : user_tokens,
	})
}
func api_cosmic_game_token_name_history(c *gin.Context) {

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

	var req_status int = 1
	var err_str string = ""

	tokname_history := arb_storagew.Get_cosmic_signature_token_name_history(token_id)
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"TokenId" : token_id,
		"TokenNameHistory" : tokname_history,
	})
}
func api_cosmic_game_token_name_search(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_name:= c.Param("name")
	if len(p_name) > 0 {
	} else {
		respond_error_json(c,"'search_text' parameter is not set")
		return
	}

	results := arb_storagew.Search_token_by_name(p_name)
	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"SearchText" : p_name,
		"TokenNameSearchResults" : results ,
	})
}
func api_cosmic_game_named_tokens_only(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	var req_status int = 1
	var err_str string = ""

	results := arb_storagew.Get_named_tokens()
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"NamedTokens" : results ,
	})
}
func api_cosmic_game_token_ownership_transfers(c *gin.Context) {

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
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	var req_status int = 1
	var err_str string = ""

	transfers := arb_storagew.Get_cst_ownership_transfers(token_id,offset,limit)
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offset" : offset,
		"Limit" : limit,
		"TokenId" : token_id,
		"TokenTransfers" : transfers,
	})
}
func api_cosmic_game_cs_token_distribution(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	distribution := arb_storagew.Get_cosmic_signature_token_distribution()

	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CosmicSignatureTokenDistribution" : distribution,
	})
}
func api_cosmic_game_user_balances(c *gin.Context) {

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

	addr := common.HexToAddress(p_user_addr)
	user_eth_bal,err := eclient.BalanceAt(context.Background(),addr,nil)
	if err != nil {
		err_str := fmt.Sprintf("Error at BalanceAt() call for addr: %v\n",err)
		respond_error_json(c,err_str)
		return
	}
	ct_contract,err := NewERC20(cosmic_token_addr,eclient);
	if err != nil {
		err_str := fmt.Sprintf("Error at instantiation of ERC20 contract: %v\n",err)
		respond_error_json(c,err_str)
		return
	}
	var copts bind.CallOpts
	ct_balance,err := ct_contract.BalanceOf(&copts,addr)
	if err != nil {
		err_str := fmt.Sprintf("Error at BalanceOf() call: %v\n",err)
		respond_error_json(c,err_str)
		return
	}

	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"ETH_Balance" : user_eth_bal.String(),
		"CosmicTokenBalance" : ct_balance.String(),
	})
}
func api_cosmic_game_cosmic_token_balances(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	balances := arb_storagew.Get_cosmic_token_holders()
	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CosmicTokenBalances" : balances,
	})
}
func api_cosmic_game_cosmic_token_transfers_by_user(c *gin.Context) {

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
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	transfers := arb_storagew.Get_cosmic_token_transfers_by_user(user_aid,offset,limit)
	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"Offset" : offset,
		"Limit" : limit,
		"CosmicTokenTransfers" : transfers,
	})
}
func api_cosmic_game_cosmic_signature_transfers_by_user(c *gin.Context) {

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
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	transfers := arb_storagew.Get_cosmic_signature_transfers_by_user(user_aid,offset,limit)
	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"Offset" : offset,
		"Limit" : limit,
		"CosmicSignatureTransfers" : transfers,
	})
}
func api_cosmic_game_used_rwalk_nfts(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	used_nfts := arb_storagew.Get_random_walk_tokens_in_bids()
	var req_status int = 1
	var err_str string = ""

	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UsedRwalkNFTs" : used_nfts,
	})
}
func api_cosmic_game_marketing_rewards_global(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	rewards := arb_storagew.Get_marketing_reward_history_global(offset, limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offset" : offset,
		"Limit" : limit,
		"MarketingRewards" : rewards,
	})
}
func api_cosmic_game_marketing_rewards_by_user(c *gin.Context) {
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
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	rewards := arb_storagew.Get_marketing_reward_history_by_user(user_aid,offset, limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offset" : offset,
		"Limit" : limit,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"UserMarketingRewards" : rewards,
	})
}
func api_cosmic_game_get_cst_price(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var copts bind.CallOpts
	contract,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate CosmicGame contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
		respond_error_json(c,err_str)
	} else {
		cst_price,err := contract.GetNextCstBidPrice(&copts,big.NewInt(0));
		if err != nil {
			Error.Printf(err.Error())
			Info.Printf(err.Error())
			respond_error(c,err.Error());
		} else {
			auction_duration,seconds_elapsed,err := contract.GetCstDutchAuctionDurations(&copts);
			if err != nil {
				Error.Printf(err.Error())
				Info.Printf(err.Error())
				respond_error(c,err.Error());
			} else {
				var req_status int = 1
				var err_str string = ""
				c.JSON(http.StatusOK, gin.H{
					"status": req_status,
					"error" : err_str,
					"CSTPrice": cst_price.String(),
					"SecondsElapsed" : seconds_elapsed.String(),
					"AuctionDuration" : auction_duration.String(),
				})
			}
		}
	}
}
func api_cosmic_game_get_eth_price(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var copts bind.CallOpts
	contract,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate CosmicGame contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
		respond_error_json(c,err_str)
	} else {
		eth_price,err := contract.GetNextEthBidPrice(&copts,big.NewInt(0));
		if err != nil {
			Error.Printf(err.Error())
			Info.Printf(err.Error())
			respond_error(c,err.Error());
		} else {
			auction_duration,seconds_elapsed,err := contract.GetEthDutchAuctionDurations(&copts);
			if err != nil {
				Error.Printf(err.Error())
				Info.Printf(err.Error())
				respond_error(c,err.Error());
			} else {
				var req_status int = 1
				var err_str string = ""
				c.JSON(http.StatusOK, gin.H{
					"status": req_status,
					"error" : err_str,
				"ETHPrice": eth_price.String(),
					"SecondsElapsed" : seconds_elapsed.String(),
					"AuctionDuration" : auction_duration.String(),
				})
			}
		}
	}
}
func api_cosmic_game_sysmode_changes(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	system_mode_changes := arb_storagew.Get_system_mode_change_event_list(offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offset" : offset,
		"Limit" : limit,
		"SystemModeChanges" : system_mode_changes,
	})
}
func api_cosmic_game_admin_events_in_range(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_evtlog_start:= c.Param("evtlog_start")
	var evtlog_start int64
	if len(p_evtlog_start) > 0 {
		var success bool
		evtlog_start,success = parse_int_from_remote_or_error(c,JSON,&p_evtlog_start)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'evtlog_start' parameter is not set")
		return
	}
	p_evtlog_end := c.Param("evtlog_end")
	var evtlog_end int64
	if len(p_evtlog_end) > 0 {
		var success bool
		evtlog_end,success = parse_int_from_remote_or_error(c,JSON,&p_evtlog_end)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'evtlog_end' parameter is not set")
		return
	}
	event_list := arb_storagew.Get_admin_events_in_range(evtlog_start,evtlog_end)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"AdminEvents" : event_list,
		"EvtLogIdStart" : evtlog_start,
		"EvtLogIdEnd" : evtlog_end,
	})
}
func api_cosmic_game_bid_special_winners(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"LastBidderAddress": last_bidder.String(),
		"LastBidderLastBidTime" : last_bidder_bid_time,
		"EnduranceChampionAddress": endurance_champ_addr,
		"EnduranceChampionDuration": endurance_duration,
		"ChronoWarriorAddress" : chrono_warrior_addr,
		"ChronoWarriorDuration" : chrono_warrior_duration,
		"LastCstBidderAddress" : lastcst_bidder_addr,
	})
}
