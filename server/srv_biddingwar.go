package main
import (
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"encoding/json"
	"math/big"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"

	. "github.com/PredictionExplorer/augur-explorer/dbs/biddingwar" // bidding war types
	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/primitives/biddingwar"

)
const (
	CONTRACT_CONSTANTS_REFRESH_TIME		= 5*60	// seconds
	CONTRACT_VARIABLES_REFRESH_TIME		= 15	// seconds
)
var (
	cosmic_game_addr				common.Address
	cosmic_signature_addr		common.Address
	cosmic_token_addr			common.Address
	charity_wallet_addr			common.Address

	// contract constants (variables not frequently modified, and only by the owner)
	price_increase				string
	charity_addr				common.Address
	charity_percentage			int64
	token_reward				string
	prize_percentage			int64
	raffle_percentage			int64
	time_increase				string
	raffle_eth_winners			int64		// numRaffleWinnersPerRound
	raffle_nft_winners			int64		// numRaffleNFTWinnersPerRound
	raffle_holder_winners		int64		// numHolderNFTWinnersPerRound

	// contract variables (variables usually modified by a bid())
	bid_price					string
	bid_price_eth				float64
	prize_claim_date			int64
	prize_amount				string
	prize_amount_eth			float64
	round_num					int64
	nanoseconds_extra			string
	last_bidder					common.Address
	charity_balance				string
	charity_balance_eth			float64
	round_start_ts				int64

	// contract counters	(collected via DB)
	bw_stats					BWStatistics

	arb_storagew				SQLStorageWrapper
)
type rpcBlock struct {
    Timestamp         string      `json:"timestamp"`
}

func biddingwar_init() {

	if  !augur_srv.arbitrum_initialized() {
		err_str := fmt.Sprintf("biddingwar_init(): Database link wasn't configured")
		Info.Printf(err_str)
		Error.Printf(err_str)
	}
	arb_storagew.S=augur_srv.db_augur
	arb_storagew.S.Db_set_schema_name("public")
	bw_caddrs := arb_storagew.Get_cosmic_game_contract_addrs()
	cosmic_game_addr = common.HexToAddress(bw_caddrs.CosmicGameAddr)
	cosmic_signature_addr = common.HexToAddress(bw_caddrs.CosmicSignatureAddr)
	cosmic_token_addr = common.HexToAddress(bw_caddrs.CosmicTokenAddr)
	charity_wallet_addr = common.HexToAddress(bw_caddrs.CharityWalletAddr)
	do_reload_contract_variables()
	do_reload_database_variables()
	go reload_database_variables_goroutine()
	go reload_constants_goroutine()
	go reload_variables_goroutine()
}
func do_reload_contract_constants() {
	var copts bind.CallOpts
	code,err := eclient.CodeAt(context.Background(), cosmic_game_addr, nil)
	if (err != nil) {
		err_str := fmt.Sprintf("Can't instantiate Cosmic gane contract: %v\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
		fmt.Printf(err_str)
		os.Exit(1)
	}
	if len(code) == 0 {
		err_str := fmt.Sprintf("Can't instantiate Cosmic gane contract: no code at given address\n")
		Error.Printf(err_str)
		Info.Printf(err_str)
		fmt.Printf(err_str)
		os.Exit(1)
	}
	bwcontract,err := NewCosmicGame(cosmic_game_addr,eclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate BiddingWar contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
	} else {
		var err error
		var tmp_val *big.Int
		tmp_val,err = bwcontract.PriceIncrease(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PriceIncrease() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			price_increase = "error"
		} else { price_increase=tmp_val.String() }
		charity_addr,err = bwcontract.Charity(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at Charity() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		tmp_val,err = bwcontract.CharityPercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at Charity() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			charity_percentage = 0
		} else { charity_percentage = tmp_val.Int64() }
		tmp_val = big.NewInt(0);
		tmp_val.SetString("100000000000000000000",10)
		if err != nil {
			err_str := fmt.Sprintf("Error at TokenReward() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			token_reward = "error"
		} else { token_reward = tmp_val.String() }
		tmp_val,err = bwcontract.PrizePercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizePercentage() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			prize_percentage = -1
		} else { prize_percentage = tmp_val.Int64() }
		tmp_val,err = bwcontract.RafflePercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at RafflePercentage() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			raffle_percentage = -1
		} else { raffle_percentage = tmp_val.Int64() }
		tmp_val,err = bwcontract.TimeIncrease(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at TimeIncrease() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			time_increase = "error"
		} else { time_increase = tmp_val.String() }
		tmp_val,err = bwcontract.NumRaffleWinnersPerRound(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at numRaffleWinnersPerRound() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			raffle_eth_winners = -1 
		} else { raffle_eth_winners = tmp_val.Int64()}
		tmp_val,err = bwcontract.NumRaffleNFTWinnersPerRound(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at numRaffleNFTWinnersPerRound() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			raffle_nft_winners = -1
		} else { raffle_nft_winners = tmp_val.Int64() }
		tmp_val,err = bwcontract.NumHolderNFTWinnersPerRound(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at numRaffleNFTWinnersPerRound() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			raffle_holder_winners = -1
		} else { raffle_holder_winners = tmp_val.Int64() }
	}
}
func do_reload_contract_variables() {
	var copts bind.CallOpts
	bwcontract,err := NewCosmicGame(cosmic_game_addr,eclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate BiddingWar contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
	} else {
		var tmp_val *big.Int
		f_divisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
		tmp_val,err = bwcontract.GetBidPrice(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at GetBidPrice() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			bid_price = "error"
		} else {
			bid_price = tmp_val.String()
			f_bid_price := big.NewFloat(0.0).SetInt(tmp_val)
			f_quo := big.NewFloat(0.0).Quo(f_bid_price,f_divisor)
			bid_price_eth,_ = f_quo.Float64()
		}
		tmp_val,err = bwcontract.PrizeTime(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizeTime() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			prize_claim_date = -1
		} else { prize_claim_date = tmp_val.Int64() }
		tmp_val , err = bwcontract.PrizeAmount(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizeAmount() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			prize_amount = "error"
		} else {
			prize_amount = tmp_val.String()
			f_prize_amount:= big.NewFloat(0.0).SetInt(tmp_val)
			f_quo := big.NewFloat(0.0).Quo(f_prize_amount,f_divisor)
			prize_amount_eth,_ = f_quo.Float64()
		}
		tmp_val , err = bwcontract.RoundNum(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumPrizes() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			round_num = -1
		} else { round_num = tmp_val.Int64() }
		tmp_val,err = bwcontract.NanoSecondsExtra(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NanoSecondsExtra() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			nanoseconds_extra = "error"
		} else { nanoseconds_extra = tmp_val.String() }
		last_bidder,err = bwcontract.LastBidder(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at LastBidder() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		tmp_val,err = eclient.BalanceAt(context.Background(),charity_addr,nil)
		if err != nil {
			err_str := fmt.Sprintf("Error at BalanceAt() call for charity addr: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			charity_balance = "error"
		} else {
			charity_balance = tmp_val.String()
			f_charity_balance := big.NewFloat(0.0).SetInt(tmp_val)
			f_quo := big.NewFloat(0.0).Quo(f_charity_balance,f_divisor)
			charity_balance_eth,_ = f_quo.Float64()
		}
	}
}
func do_reload_database_variables() {
	bw_stats = arb_storagew.Get_biddingwar_statistics()
	round_start_ts = arb_storagew.Get_round_start_timestamp(bw_stats.TotalPrizes)
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
func reload_variables_goroutine() {
	for {
		do_reload_contract_variables()
		time.Sleep(CONTRACT_VARIABLES_REFRESH_TIME * time.Second)
	}
}
func reload_database_variables_goroutine() {
	for {
		do_reload_database_variables()
		time.Sleep(CONTRACT_VARIABLES_REFRESH_TIME * time.Second)
	}
}
func biddingwar_index_page(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	ts := time.Unix(round_start_ts,0)
	date_str := fmt.Sprintf("%v",ts);
	c.HTML(http.StatusOK, "bw_index.html", gin.H{
		"BiddingWarAddr":cosmic_game_addr,
		"CosmicSignatureAddr":cosmic_signature_addr,
		"CosmicSignatureTokenAddr":cosmic_token_addr,
		"CharityWalletAddr":charity_wallet_addr,
		"BidPrice":bid_price,
		"BidPriceEth":bid_price_eth,
		"PrizeClaimDate":time.Unix(prize_claim_date,0).Format(time.RFC822),
		"PrizeClaimTs":prize_claim_date,
		"CurRoundNum": round_num+1,
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
		"CharityAddr" : charity_addr.String(),
		"CharityPercentage" : charity_percentage,
		"NumRaffleEthWinners" : raffle_eth_winners,
		"NumRaffleNFTWinners" : raffle_nft_winners,
		"NumHolderNFTWinners" : raffle_holder_winners,
		"CharityBalance": charity_balance,
		"CharityBalanceEth": charity_balance_eth,
		"NumUniqueBidders" :  bw_stats.NumUniqueBidders,
		"NumUniqueWinners" : bw_stats.NumUniqueWinners,
		"NumDonatedNFTs" : bw_stats.NumDonatedNFTs,
		"MainStats" : bw_stats,
		"TsRoundStart" : round_start_ts,
		"DateRoundStart" : date_str,
	})
}
func biddingwar_prize_claims(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	prizes := arb_storagew.Get_prize_claims(offset,limit)

	c.HTML(http.StatusOK, "bw_prize_claims.html", gin.H{
		"PrizeClaims" : prizes,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func biddingwar_bids(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	bids := arb_storagew.Get_bids(offset,limit)

	c.HTML(http.StatusOK, "bw_bids.html", gin.H{
		"Bids" : bids,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func biddingwar_bid_list_by_round(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,HTTP,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'round_num' parameter is not set")
		return
	}
	p_sort:= c.Param("sort")
	var sort int64
	if len(p_sort) > 0 {
		var success bool
		sort,success = parse_int_from_remote_or_error(c,HTTP,&p_sort)
		if !success {
			return
		}
	} else {
		respond_error(c,"'sort' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}

	bids := arb_storagew.Get_bids_by_round_num(round_num,int(sort),offset,limit)
	c.HTML(http.StatusOK, "bw_bids_by_round_num.html", gin.H{
		"RoundNum" : round_num,
		"Offset" : offset,
		"Limit" : limit,
		"Sort" : sort,
		"BidsByRound" : bids,
	})
}
func biddingwar_bid_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_evtlog_id := c.Param("evtlog_id")
	var evtlog_id int64
	if len(p_evtlog_id) > 0 {
		var success bool
		evtlog_id,success = parse_int_from_remote_or_error(c,HTTP,&p_evtlog_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'evtlog_id' parameter is not set")
		return
	}
	record_found,bid_info := arb_storagew.Get_bid_info(evtlog_id)
	if !record_found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Provided evtlog_id wasn't found"),
		})
	} else {
		if len(bid_info.NFTTokenURI) > 0 {
			resp,err := http.Get(bid_info.NFTTokenURI)
			if err == nil {
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					var response map[string]interface{}
					response = make(map[string]interface{})
					err := json.Unmarshal(body,&response)
					if err == nil {
						image_url := response["image"].(string)
						bid_info.ImageURL=image_url
					}
				}
			}
		}
		c.HTML(http.StatusOK, "bw_bid_info.html", gin.H{
			"BidInfo" : bid_info,
		})
	}
} 
func biddingwar_prize_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_prize_num:= c.Param("prize_num")
	var prize_num int64
	if len(p_prize_num) > 0 {
		var success bool
		prize_num,success = parse_int_from_remote_or_error(c,HTTP,&p_prize_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'prize_num' parameter is not set")
		return
	}
	record_found,prize_info := arb_storagew.Get_prize_info(prize_num)
	if !record_found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Prize with provided number wasn't found"),
		})
	} else {
		c.HTML(http.StatusOK, "bw_prize_info.html", gin.H{
			"PrizeInfo" : prize_info,
		})
	}
}
func biddingwar_user_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}

	found, user_info := arb_storagew.Get_user_info(user_aid)
	if !found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	bids := arb_storagew.Get_bids_by_user(user_aid)
	prizes := arb_storagew.Get_prize_claims_by_user(user_aid)
	c.HTML(http.StatusOK, "bw_user_info.html", gin.H{
		"UserInfo" : user_info,
		"Bids" : bids,
		"Prizes" : prizes,
	})
}
func biddingwar_charity_donations(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	biddingwar_aid,err :=arb_storagew.S.Nonfatal_lookup_address_id(cosmic_game_addr.String())
	if err != nil {
		Error.Printf("BiddingWar contract address doesn't exist in the DB, aborting server")
		os.Exit(1)
	}
	donations := arb_storagew.Get_charity_donations(biddingwar_aid)
	c.HTML(http.StatusOK, "bw_charity_donations.html", gin.H{
		"CharityDonations" : donations,
	})
}
func biddingwar_donations_eth(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	donations := arb_storagew.Get_donations_to_biddingwar()
	c.HTML(http.StatusOK, "bw_donations_to_biddingwar.html", gin.H{
		"BiddingwarDonations" : donations,
	})
}
func biddingwar_unique_bidders(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	unique_bidders := arb_storagew.Get_unique_bidders()
	c.HTML(http.StatusOK, "bw_unique_bidders.html", gin.H{
		"UniqueBidders" : unique_bidders,
	})
}
func biddingwar_unique_winners(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	unique_winners:= arb_storagew.Get_unique_winners()
	c.HTML(http.StatusOK, "bw_unique_winners.html", gin.H{
		"UniqueWinners" : unique_winners,
	})
}
func biddingwar_donations_nft(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	nft_donations := arb_storagew.Get_NFT_donations(offset,limit)

	c.HTML(http.StatusOK, "bw_nft_donations.html", gin.H{
		"NFTDonations" : nft_donations,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func biddingwar_nft_donation_stats(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	nft_donation_stats := arb_storagew.Get_NFT_donation_stats()
	c.HTML(http.StatusOK, "bw_nft_donation_stats.html", gin.H{
		"NFTDonationStats" : nft_donation_stats,
	})
}
func biddingwar_donations_nft_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_record_id:= c.Param("record_id")
	var record_id int64
	if len(p_record_id) > 0 {
		var success bool
		record_id,success = parse_int_from_remote_or_error(c,HTTP,&p_record_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'record_id' parameter is not set")
		return
	}
	found,nftdonation := arb_storagew.Get_NFT_donation_info(record_id)
	if !found {
		respond_error(c,"Database link wasn't configured")
	} else {
		c.HTML(http.StatusOK, "bw_donated_nft_info.html", gin.H{
			"NFTDonation" : nftdonation,
		})
	}
}
func biddingwar_raffle_deposits_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	deposits := arb_storagew.Get_raffle_eth_deposits_list(offset,limit)

	c.HTML(http.StatusOK, "bw_raffle_deposits.html", gin.H{
		"RaffleDeposits" : deposits,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func biddingwar_raffle_deposits_by_round(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,HTTP,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'round_num' parameter is not set")
		return
	}
	deposits := arb_storagew.Get_raffle_deposits_by_round(round_num)

	c.HTML(http.StatusOK, "bw_raffle_deposits_by_round.html", gin.H{
		"RaffleDeposits" : deposits,
		"RoundNum" : round_num,
	})
}
func biddingwar_raffle_nft_winners_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	winners := arb_storagew.Get_raffle_nft_winners(offset,limit)

	c.HTML(http.StatusOK, "bw_raffle_nft_winners.html", gin.H{
		"RaffleNFTWinners" : winners,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func biddingwar_raffle_nft_winners_by_round(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,HTTP,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'round_num' parameter is not set")
		return
	}
	winners := arb_storagew.Get_raffle_nft_winners_by_round(round_num)

	c.HTML(http.StatusOK, "bw_raffle_nft_winners_by_round.html", gin.H{
		"RaffleNFTWinners" : winners,
		"RoundNum" : round_num,
	})
}
/* DISCONTINUED, removal pending
func biddingwar_raffle_nft_claims(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	deposits := arb_storagew.Get_raffle_nft_claims(offset,limit)

	c.HTML(http.StatusOK, "bw_raffle_nft_claims.html", gin.H{
		"RaffleNFTClaims" : deposits,
		"Offset" : offset,
		"Limit" : limit,
	})
}*/
func biddingwar_raffle_deposits_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	found, user_info := arb_storagew.Get_user_info(user_aid)
	if !found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}

	deposits := arb_storagew.Get_raffle_deposits_by_user(user_aid)

	c.HTML(http.StatusOK, "bw_user_raffle_deposits.html", gin.H{
		"UserRaffleDeposits" : deposits,
		"UserInfo" : user_info,
	})
}
/* DISCONTINUED  , removal pending
func biddingwar_raffle_nft_claims_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	found, user_info := arb_storagew.Get_user_info(user_aid)
	if !found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}

	claims := arb_storagew.Get_raffle_nft_claims_by_user(user_aid)

	c.HTML(http.StatusOK, "bw_user_raffle_nft_claims.html", gin.H{
		"UserRaffleNFTClaims" : claims,
		"UserInfo" : user_info,
	})
}*/
func biddingwar_user_raffle_nft_winnings(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	found, user_info := arb_storagew.Get_user_info(user_aid)
	if !found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	winnings := arb_storagew.Get_raffle_nft_winnings_by_user(user_aid)
	fmt.Printf("winnings len = %v\n",len(winnings))

	c.HTML(http.StatusOK, "bw_user_raffle_nft_winnings.html", gin.H{
		"UserRaffleNFTWinnings" : winnings,
		"UserInfo" : user_info,
	})
}
func biddingwar_nft_donations_by_prize(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_prize_num:= c.Param("prize_num")
	var prize_num int64
	if len(p_prize_num) > 0 {
		var success bool
		prize_num,success = parse_int_from_remote_or_error(c,HTTP,&p_prize_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'prize_num' parameter is not set")
		return
	}
	nft_donations := arb_storagew.Get_nft_donations_by_prize(prize_num)
	c.HTML(http.StatusOK, "bw_nft_donations_by_prize.html", gin.H{
		"NFTDonations" : nft_donations,
		"PrizeNum": prize_num,
	})
}
func biddingwar_cosmic_signature_token_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	tokens := arb_storagew.Get_cosmic_signature_nft_list(offset,limit)

	c.HTML(http.StatusOK, "bw_cosmic_sig_token_list.html", gin.H{
		"CosmicSignatureTokenList" : tokens,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func biddingwar_cosmic_signature_token_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_token_id:= c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id,success = parse_int_from_remote_or_error(c,HTTP,&p_token_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'token_id' parameter is not set")
		return
	}

	record_found,token_info := arb_storagew.Get_cosmic_signature_token_info(token_id)
	if !record_found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Cosmic Game: Error",
			"ErrDescr": fmt.Sprintf("Prize with provided token_id wasn't found"),
		})
		return
	}
	if token_info.PrizeNum > -1 {
		_,prize_info := arb_storagew.Get_prize_info(token_info.PrizeNum)
		c.HTML(http.StatusOK, "bw_cosmic_sig_token_info.html", gin.H{
			"TokenInfo" : token_info,
			"PrizeInfo" : prize_info,
		})
	} else {
		c.HTML(http.StatusOK, "bw_cosmic_sig_token_info.html", gin.H{
			"TokenInfo" : token_info,
		})
	}
}
func biddingwar_donated_nft_claims_all(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	claims := arb_storagew.Get_donated_nft_claims(offset,limit)
	c.HTML(http.StatusOK, "bw_donated_nft_claims.html", gin.H{
		"DonatedNFTClaims" : claims,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func biddingwar_donated_nft_claims_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}

	found, user_info := arb_storagew.Get_user_info(user_aid)
	if !found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	claims := arb_storagew.Get_donated_nft_claims_by_user(user_aid)
	c.HTML(http.StatusOK, "bw_donated_nft_claims_by_user.html", gin.H{
		"DonatedNFTClaims" : claims,
		"UserInfo" : user_info,
	})
}
func biddingwar_time_current(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	var raw json.RawMessage
	err := rpcclient.CallContext(context.Background(), &raw,"eth_getBlockByNumber", "pending",true)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "CosmicGame: Error",
			"ErrDescr": fmt.Sprintf("%v",err),
		})
		return
	}
	var rpcobj map[string]interface{}
	rpcobj = make(map[string]interface{})
	err = json.Unmarshal(raw,&rpcobj)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "CosmicGame: Error",
			"ErrDescr": fmt.Sprintf("Error decoding JSON: %v",err),
		})
		return
	}

	ts_hex := rpcobj["timestamp"].(string)
	ts,err := hexutil.DecodeUint64(ts_hex)
	if err !=nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "CosmicGame: Error",
			"ErrDescr": fmt.Sprintf("Error decoding timestamp from hex: %v",err),
		})
		return
	}
	c.HTML(http.StatusOK, "bw_cur_ts.html", gin.H{
		"CurrentTimeStamp": ts,
	})
}
func biddingwar_time_until_prize(c *gin.Context) {

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
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "CosmicGame: Error",
			"ErrDescr": fmt.Sprintf("%v",err),
		})
		return
	}
	var ts_hex string
	err = json.Unmarshal(raw,&ts_hex)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "CosmicGame: Error",
			"ErrDescr": fmt.Sprintf("Error decoding JSON: %v",err),
		})
		return
	}
	ts_big := common.HexToHash(ts_hex).Big()
	c.HTML(http.StatusOK, "bw_time_until_prize.html", gin.H{
		"TimeUntilPrize": ts_big.Int64(),
	})
}
func biddingwar_user_global_winnings(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}

	found, user_info := arb_storagew.Get_user_info(user_aid)
	if !found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	claim_info := arb_storagew.Get_user_global_winnings(user_aid)
	c.HTML(http.StatusOK, "bw_user_global_winnings.html", gin.H{
		"Winnings" : claim_info,
		"UserInfo" : user_info,
	})
}
/* DISCONTINUED, removal pending
func biddingwar_unclaimed_token_list_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}

	token_list := arb_storagew.Get_unclaimed_token_ids(user_aid)
	c.HTML(http.StatusOK, "bw_unclaimed_token_list_by_user.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"Tokens" : token_list,
	})
}*/
