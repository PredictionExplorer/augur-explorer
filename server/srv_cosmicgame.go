package main
import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"math"
	"strings"
	"io/ioutil"
	"encoding/json"
	"math/big"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"

	. "github.com/PredictionExplorer/augur-explorer/dbs/cosmicgame"
	. "github.com/PredictionExplorer/augur-explorer/contracts"
	. "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"

)
const (
	CONTRACT_CONSTANTS_REFRESH_TIME		= 5*60	// seconds
	CONTRACT_VARIABLES_REFRESH_TIME		= 5	// seconds
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
	staking_percentage			int64
	time_increase				string
	initial_seconds				int64
	timeout_claim				int64
	roundstart_auclen			int64
	raffle_eth_winners_bidding	int64		// numRaffleETHWinnersBidding
	raffle_nft_winners_bidding	int64		// numRaffleNFTWinnersBidding
	raffle_nft_winners_staking_rwalk	int64	// numRaffleNFTWinnersStakingRWalk

	// contract variables (variables usually modified by a bid())
	bid_price					string
	bid_price_eth				float64
	prize_claim_date			int64
	prize_amount				string
	prize_amount_eth			float64
	raffle_amount				string
	raffle_amount_eth			float64
	round_num					int64
	mainprize_microseconds_inc	string
	last_bidder					common.Address
	last_bidder_bid_time		int64
	charity_balance				string
	charity_balance_eth			float64
	round_start_ts				int64
	endurance_champ_addr		string
	endurance_duration			int64
	chrono_warrior_addr			string
	chrono_warrior_duration		int64
	lastcst_bidder_addr			string

	// contract counters	(collected via DB)
	bw_stats					CGStatistics

	arb_storagew				SQLStorageWrapper
)
type rpcBlock struct {
    Timestamp         string      `json:"timestamp"`
}

func cosmic_game_init() {

	if  !augur_srv.arbitrum_initialized() {
		err_str := fmt.Sprintf("cosmic_game_init(): Database link wasn't configured")
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
func get_cosmic_game_contract_balance() float64 {

	cg_eth_bal,err := eclient.BalanceAt(context.Background(),cosmic_game_addr,nil)
	if err != nil {
		err_str := fmt.Sprintf("Error at BalanceAt() call for cosmic game: %v\n",err)
		Info.Printf(err_str)
		Error.Printf(err_str)
		return math.NaN()
	}
	f_divisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
	f_balance := big.NewFloat(0.0).SetInt(cg_eth_bal)
	f_quo := big.NewFloat(0.0).Quo(f_balance,f_divisor)
	f_out,_ := f_quo.Float64()
	return f_out
}
func do_reload_contract_constants() {
	var copts bind.CallOpts
	code,err := eclient.CodeAt(context.Background(), cosmic_game_addr, nil)
	if (err != nil) {
		err_str := fmt.Sprintf("Can't instantiate Cosmic gane contract: %v\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
		fmt.Printf(err_str)
	}
	if len(code) == 0 {
		err_str := fmt.Sprintf("Can't instantiate Cosmic gane contract: no code at given address\n")
		Error.Printf(err_str)
		Info.Printf(err_str)
		fmt.Printf(err_str)
	}
	bwcontract,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate CosmicGame contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
	} else {
		var err error
		var tmp_val *big.Int
		tmp_val,err = bwcontract.EthBidPriceIncreaseDivisor(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PriceIncrease() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			price_increase = "error"
		} else { price_increase=tmp_val.String() }
		charity_addr,err = bwcontract.CharityAddress(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at Charity() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		tmp_val,err = bwcontract.CharityEthDonationAmountPercentage(&copts)
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
		tmp_val,err = bwcontract.MainEthPrizeAmountPercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizePercentage() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			prize_percentage = -1
		} else { prize_percentage = tmp_val.Int64() }
		tmp_val,err = bwcontract.RaffleTotalEthPrizeAmountForBiddersPercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at RafflePercentage() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			raffle_percentage = -1
		} else { raffle_percentage = tmp_val.Int64() }
		tmp_val,err = bwcontract.CosmicSignatureNftStakingTotalEthRewardAmountPercentage(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at StakingPercentage() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			staking_percentage = -1
		} else { staking_percentage = tmp_val.Int64() }
		tmp_val,err = bwcontract.MainPrizeTimeIncrementIncreaseDivisor(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at TimeIncrease() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			time_increase = "error"
		} else { time_increase = tmp_val.String() }
		tmp_val,err = bwcontract.NumRaffleEthPrizesForBidders(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumRaffleETHWinnersBidding() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			raffle_eth_winners_bidding = -1 
		} else { raffle_eth_winners_bidding = tmp_val.Int64()}
		tmp_val,err = bwcontract.NumRaffleCosmicSignatureNftsForBidders(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumRaffleNFTWinnersBidding() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			raffle_nft_winners_bidding = -1
		} else { raffle_nft_winners_bidding = tmp_val.Int64() }
		tmp_val,err = bwcontract.NumRaffleCosmicSignatureNftsForRandomWalkNftStakers(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumRaffleNFTWinnersStakingRWalk() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			raffle_nft_winners_staking_rwalk = -1
		} else { raffle_nft_winners_staking_rwalk = tmp_val.Int64() }
	}
}
func do_reload_contract_variables() {
	var copts bind.CallOpts
	bwcontract,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate CosmicGame contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
	} else {
		var tmp_val *big.Int
		f_divisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
		tmp_val,err = bwcontract.GetNextEthBidPrice(&copts,big.NewInt(0))
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
		tmp_val,err = bwcontract.GetDurationUntilMainPrize(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at PrizeTime() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			prize_claim_date = -1
		} else { prize_claim_date = tmp_val.Int64() }
		tmp_val , err = bwcontract.GetMainEthPrizeAmount(&copts)
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
		tmp_val , err = bwcontract.GetRaffleTotalEthPrizeAmountForBidders(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at RaffleAmount() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			raffle_amount = "error"
		} else {
			raffle_amount = tmp_val.String()
			f_raffle_amount:= big.NewFloat(0.0).SetInt(tmp_val)
			f_quo := big.NewFloat(0.0).Quo(f_raffle_amount,f_divisor)
			raffle_amount_eth,_ = f_quo.Float64()
		}
		tmp_val , err = bwcontract.RoundNum(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at NumPrizes() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			round_num = -1
		} else { round_num = tmp_val.Int64() }
		tmp_val,err = bwcontract.MainPrizeTimeIncrementInMicroSeconds(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at MainPrizeTimeIncrementInMicroseconds() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			mainprize_microseconds_inc = "error"
		} else { mainprize_microseconds_inc = tmp_val.String() }
		last_bidder,err = bwcontract.LastBidderAddress(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at LastBidder() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		}
		if round_num > -1 {
			tmp_bidder_info,err := bwcontract.BiddersInfo(&copts,big.NewInt(round_num),last_bidder)
			if err != nil {
				err_str := fmt.Sprintf("Error at BiddersInfo() call: %v\n",err)
				Error.Printf(err_str)
				Info.Printf(err_str)
				last_bidder_bid_time = -1
			} else { last_bidder_bid_time = tmp_bidder_info.LastBidTimeStamp.Int64() }
		}
		tmp_val,err = bwcontract.InitialDurationUntilMainPrizeDivisor(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at ImitialDurationUntilMainPrizeDivisor() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			initial_seconds = -1
		} else { initial_seconds = tmp_val.Int64() }
		tmp_val,err = bwcontract.TimeoutDurationToClaimMainPrize(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at TimeoutClaimPrize() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			timeout_claim = -1
		} else { timeout_claim = tmp_val.Int64() }
		tmp_val,err = bwcontract.CstDutchAuctionDurationDivisor(&copts)
		if err != nil {
			err_str := fmt.Sprintf("Error at CstDutchAuctionDurationDivisor() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
			roundstart_auclen = -1
		} else { roundstart_auclen = tmp_val.Int64() }
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
		//tmp_addr1, tmp_duration1,tmp_addr2,tmp_duration2, err := bwcontract.TryGetCurrentChampions(&copts);
		champs,err := bwcontract.TryGetCurrentChampions(&copts);
		if err != nil {
			err_str := fmt.Sprintf("Error at TryGetCurrentChampions() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		} else {
			endurance_champ_addr = champs.EnduranceChampionAddress.String()
			endurance_duration = champs.EnduranceChampionDuration.Int64()
			chrono_warrior_addr = champs.ChronoWarriorAddress.String()
			chrono_warrior_duration = champs.ChronoWarriorDuration.Int64()
		}
		tmp_addr, err := bwcontract.LastCstBidderAddress(&copts);
		if err != nil {
			err_str := fmt.Sprintf("Error at lastCstBidderAddress() call: %v\n",err)
			Error.Printf(err_str)
			Info.Printf(err_str)
		} else {
			lastcst_bidder_addr = tmp_addr.String()
		}
	}
}
func do_reload_database_variables() {
	bw_stats = arb_storagew.Get_cosmic_game_statistics()
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
func cosmic_game_index_page(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	cur_round_stats := arb_storagew.Get_cosmic_game_round_statistics(round_num);
	cg_balance := get_cosmic_game_contract_balance()
	ts := time.Unix(round_start_ts,0)
	date_str := fmt.Sprintf("%v",ts);
	record_counters := arb_storagew.Get_record_counters()
	c.HTML(http.StatusOK, "cg_index.html", gin.H{
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
		"NanosecondsExtra" : mainprize_microseconds_inc,
		"InitialSecondsUntilPrize" : initial_seconds,
		"TimeoutClaimPrize" : timeout_claim,
		"RoundStartCSTAuctionLength" : roundstart_auclen,
		"TokenReward" : token_reward,
		"PrizePercentage" : prize_percentage,
		"RafflePercentage" : raffle_percentage,
		"StakingPercentage" : staking_percentage,
		"CharityAddr" : charity_addr.String(),
		"CharityPercentage" : charity_percentage,
		"CharityBalance": charity_balance,
		"CharityBalanceEth": charity_balance_eth,
		"NumRaffleETHWinnersBidding" : raffle_eth_winners_bidding,
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
		"DateRoundStart" : date_str,
		"RecordCounters" : record_counters,
	})
}
func cosmic_game_prize_claims(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	prizes := arb_storagew.Get_prize_claims(offset,limit)

	c.HTML(http.StatusOK, "cg_rounds.html", gin.H{
		"Rounds" : prizes,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func cosmic_game_bids(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	bids := arb_storagew.Get_bids(offset,limit)

	c.HTML(http.StatusOK, "cg_bids.html", gin.H{
		"Bids" : bids,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func cosmic_game_bid_list_by_round(c *gin.Context) {

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

	bids,total_rows := arb_storagew.Get_bids_by_round_num(round_num,int(sort),offset,limit)
	c.HTML(http.StatusOK, "cg_bids_by_round_num.html", gin.H{
		"RoundNum" : round_num,
		"Offset" : offset,
		"Limit" : limit,
		"Sort" : sort,
		"BidsByRound" : bids,
		"TotalRows" : total_rows,
	})
}
func cosmic_game_bid_info(c *gin.Context) {

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
		c.HTML(http.StatusOK, "cg_bid_info.html", gin.H{
			"BidInfo" : bid_info,
		})
	}
} 
func cosmic_game_round_info(c *gin.Context) {

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
		nft_donations := arb_storagew.Get_nft_donations_by_prize(prize_num)
		c.HTML(http.StatusOK, "cg_round_info.html", gin.H{
			"RoundInfo" : prize_info,
			"DonatedNFTs" : nft_donations,
		})
	}
}
func cosmic_game_user_info(c *gin.Context) {

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
	ct_contract,err := NewERC20(cosmic_token_addr,eclient);
	if err != nil {
		err_str := fmt.Sprintf("Error at instantiation of ERC20 contract: %v\n",err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("%v",err_str),
		})
		return
	}
	addr := common.HexToAddress(p_user_addr)
	var copts bind.CallOpts
	ct_balance,err := ct_contract.BalanceOf(&copts,addr)
	if err != nil {
		err_str := fmt.Sprintf("Error at BalanceOf() call: %v\n",err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("%v",err_str),
		})
		return
	}
	f_divisor := big.NewFloat(0.0).SetInt(big.NewInt(1e18))
	f_balance_eth := big.NewFloat(0.0).SetInt(ct_balance)
	f_quo := big.NewFloat(0.0).Quo(f_balance_eth,f_divisor)
	bal_eth,_ := f_quo.Float64()

	bids := arb_storagew.Get_bids_by_user(user_aid)
	prizes := arb_storagew.Get_prize_claims_by_user(user_aid)
	c.HTML(http.StatusOK, "cg_user_info.html", gin.H{
		"UserInfo" : user_info,
		"Bids" : bids,
		"Prizes" : prizes,
		"CTBalance" : ct_balance.String(),
		"CTBalanceEth" : bal_eth,
	})
}
func cosmic_game_charity_donations_deposits(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	cosmicgame_aid,err :=arb_storagew.S.Nonfatal_lookup_address_id(cosmic_game_addr.String())
	if err != nil {
		Error.Printf("cosmic game contract address doesn't exist in the DB, aborting server")
		os.Exit(1)
	}
	donations := arb_storagew.Get_charity_donations(cosmicgame_aid)
	c.HTML(http.StatusOK, "cg_charity_donations_deposits.html", gin.H{
		"CharityDeposits" : donations,
	})
}
func cosmic_game_charity_cosmicgame_deposits(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	cosmicgame_aid,err :=arb_storagew.S.Nonfatal_lookup_address_id(cosmic_game_addr.String())
	if err != nil {
		Error.Printf("cosmic game contract address doesn't exist in the DB, aborting server")
		os.Exit(1)
	}
	donations := arb_storagew.Get_charity_donations_from_cosmic_game(cosmicgame_aid)
	c.HTML(http.StatusOK, "cg_charity_donations_cosmicgame_deposits.html", gin.H{
		"CharityDeposits" : donations,
	})
}
func cosmic_game_charity_voluntary_deposits(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	cosmicgame_aid,err :=arb_storagew.S.Nonfatal_lookup_address_id(cosmic_game_addr.String())
	if err != nil {
		Error.Printf("cosmic game contract address doesn't exist in the DB, aborting server")
		os.Exit(1)
	}
	donations := arb_storagew.Get_charity_donations_voluntary(cosmicgame_aid)
	c.HTML(http.StatusOK, "cg_charity_donations_voluntary_deposits.html", gin.H{
		"CharityDeposits" : donations,
	})
}
func cosmic_game_charity_donations_withdrawals(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	donations := arb_storagew.Get_charity_wallet_withdrawals()
	c.HTML(http.StatusOK, "cg_charity_donations_withdrawals.html", gin.H{
		"CharityWithdrawals" : donations,
	})
}
func cosmic_game_unique_bidders(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	unique_bidders := arb_storagew.Get_unique_bidders()
	c.HTML(http.StatusOK, "cg_unique_bidders.html", gin.H{
		"UniqueBidders" : unique_bidders,
	})
}
func cosmic_game_unique_winners(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	unique_winners:= arb_storagew.Get_unique_winners()
	c.HTML(http.StatusOK, "cg_unique_winners.html", gin.H{
		"UniqueWinners" : unique_winners,
	})
}
func cosmic_game_unique_donors(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	unique_donors := arb_storagew.Get_unique_donors()
	c.HTML(http.StatusOK, "cg_unique_donors.html", gin.H{
		"UniqueDonors" : unique_donors,
	})
}
func cosmic_game_donations_nft(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	nft_donations := arb_storagew.Get_NFT_donations(offset,limit)

	c.HTML(http.StatusOK, "cg_nft_donations.html", gin.H{
		"NFTDonations" : nft_donations,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func cosmic_game_nft_donation_stats(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	nft_donation_stats := arb_storagew.Get_NFT_donation_stats()
	c.HTML(http.StatusOK, "cg_nft_donation_stats.html", gin.H{
		"NFTDonationStats" : nft_donation_stats,
	})
}
func cosmic_game_donations_nft_info(c *gin.Context) {

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
		respond_error(c,"Record not found")
	} else {
		c.HTML(http.StatusOK, "cg_donated_nft_info.html", gin.H{
			"NFTDonation" : nftdonation,
		})
	}
}
func cosmic_game_prize_deposits_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	deposits := arb_storagew.Get_prize_eth_deposits_list(offset,limit)

	c.HTML(http.StatusOK, "cg_raffle_deposits.html", gin.H{
		"RaffleDeposits" : deposits,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func cosmic_game_prize_deposits_by_round(c *gin.Context) {

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
	deposits := arb_storagew.Get_prize_deposits_by_round(round_num)

	c.HTML(http.StatusOK, "cg_raffle_deposits_by_round.html", gin.H{
		"RaffleDeposits" : deposits,
		"RoundNum" : round_num,
	})
}
func cosmic_game_raffle_nft_winners_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	winners := arb_storagew.Get_raffle_nft_winners(offset,limit)

	c.HTML(http.StatusOK, "cg_raffle_nft_winners.html", gin.H{
		"RaffleNFTWinners" : winners,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func cosmic_game_raffle_nft_winners_by_round(c *gin.Context) {

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
	winners_raffle := arb_storagew.Get_raffle_nft_winners_by_round(round_num,false)
	winners_staking := arb_storagew.Get_raffle_nft_winners_by_round(round_num,true)

	c.HTML(http.StatusOK, "cg_raffle_nft_winners_by_round.html", gin.H{
		"RaffleNFTWinners" : winners_raffle,
		"StakingNFTWinners" : winners_staking,
		"RoundNum" : round_num,
	})
}
func cosmic_game_prize_deposits_raffle_eth_by_user(c *gin.Context) {

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

	deposits := arb_storagew.Get_prize_deposits_raffle_eth_by_user(user_aid)

	c.HTML(http.StatusOK, "cg_user_raffle_deposits.html", gin.H{
		"UserRaffleDeposits" : deposits,
		"UserInfo" : user_info,
	})
}
func cosmic_game_prize_deposits_chrono_warrior_by_user(c *gin.Context) {

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

	deposits := arb_storagew.Get_prize_deposits_chrono_warrior_by_user(user_aid)

	c.HTML(http.StatusOK, "cg_user_chrono_warrior_deposits.html", gin.H{
		"UserChronoWarriorDeposits" : deposits,
		"UserInfo" : user_info,
	})
}
func cosmic_game_user_raffle_nft_winnings(c *gin.Context) {

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

	c.HTML(http.StatusOK, "cg_user_raffle_nft_winnings.html", gin.H{
		"UserRaffleNFTWinnings" : winnings,
		"UserInfo" : user_info,
	})
}
func cosmic_game_nft_donations_by_prize(c *gin.Context) {

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
	c.HTML(http.StatusOK, "cg_nft_donations_by_prize.html", gin.H{
		"NFTDonations" : nft_donations,
		"RoundNum": prize_num,
	})
}
func cosmic_game_cosmic_signature_token_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	tokens := arb_storagew.Get_cosmic_signature_nft_list(offset,limit)

	c.HTML(http.StatusOK, "cg_cosmic_sig_token_list.html", gin.H{
		"CosmicSignatureTokenList" : tokens,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func cosmic_game_cosmic_signature_token_info(c *gin.Context) {

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
	tokname_history := arb_storagew.Get_cosmic_signature_token_name_history(token_id)
	transfers := arb_storagew.Get_cst_ownership_transfers(token_id,0, 0)
	if token_info.RecordType == 3 {
		_,prize_info := arb_storagew.Get_prize_info(token_info.RoundNum)
		c.HTML(http.StatusOK, "cg_cosmic_sig_token_info.html", gin.H{
			"TokenInfo" : token_info,
			"PrizeInfo" : prize_info,
			"TokenNameHistory" : tokname_history,
			"TokenTransfers": transfers,
		})
	} else {
		c.HTML(http.StatusOK, "cg_cosmic_sig_token_info.html", gin.H{
			"TokenInfo" : token_info,
			"TokenNameHistory" : tokname_history,
			"TokenTransfers" : transfers,
		})
	}
}
func cosmic_game_donated_nft_claims_all(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	claims := arb_storagew.Get_donated_nft_claims(offset,limit)
	c.HTML(http.StatusOK, "cg_donated_nft_claims.html", gin.H{
		"DonatedNFTClaims" : claims,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func cosmic_game_donated_nft_claims_by_user(c *gin.Context) {

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
	c.HTML(http.StatusOK, "cg_donated_nft_claims_by_user.html", gin.H{
		"DonatedNFTClaims" : claims,
		"UserInfo" : user_info,
	})
}
func cosmic_game_time_current(c *gin.Context) {

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
	c.HTML(http.StatusOK, "cg_cur_ts.html", gin.H{
		"CurrentTimeStamp": ts,
	})
}
func cosmic_game_time_until_prize(c *gin.Context) {

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
	c.HTML(http.StatusOK, "cg_time_until_prize.html", gin.H{
		"TimeUntilPrize": ts_big.Int64(),
	})
}
func cosmic_game_user_notif_red_box_rewards(c *gin.Context) {

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

	claim_info := arb_storagew.Get_user_notif_red_box_rewards(user_aid)
	c.HTML(http.StatusOK, "cg_notif_red_box.html", gin.H{
		"Winnings" : claim_info,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
	})
}
func cosmic_game_global_claim_history_detail(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}


	claim_history := arb_storagew.Get_claim_history_detailed_global(0, 1000000)
	c.HTML(http.StatusOK, "cg_prizes_history_global.html", gin.H{
		"GlobalPrizeHistory" : claim_history,
	})
}
func cosmic_game_prize_history_detail_by_user(c *gin.Context) {

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
	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}

	claim_history := arb_storagew.Get_prize_history_detailed_by_user(user_aid,offset,limit)
	c.HTML(http.StatusOK, "cg_user_claim_history_detail.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"ClaimHistory" : claim_history,
	})
}
func cosmic_game_unclaimed_donated_nfts_by_user(c *gin.Context) {

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

	nfts := arb_storagew.Get_unclaimed_donated_nft_by_user(user_aid)
	c.HTML(http.StatusOK, "cg_unclaimed_donated_nfts_by_user.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"UnclaimedDonatedNFTs" : nfts,
	})
}
func cosmic_game_unclaimed_donated_nfts_by_prize(c *gin.Context) {

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
	nft_donations := arb_storagew.Get_unclaimed_donated_nfts_by_prize(prize_num)
	c.HTML(http.StatusOK, "cg_nft_donations_by_prize.html", gin.H{
		"NFTDonations" : nft_donations,
		"RoundNum": prize_num,
	})
}
func cosmic_game_unclaimed_prize_deposits_by_user(c *gin.Context) {

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
	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}

	offset = 0; limit = 100000000;

	deposits := arb_storagew.Get_unclaimed_prize_eth_deposits(user_aid,offset,limit)
	c.HTML(http.StatusOK, "cg_user_unclaimed_raffle_eth_deposits.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"UnclaimedDeposits" : deposits,
	})
}
func cosmic_game_cosmic_signature_token_list_by_user(c *gin.Context) {

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
	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}

	offset = 0; limit = 100000000;

	user_tokens := arb_storagew.Get_cosmic_signature_nft_list_by_user(user_aid,offset,limit)
	c.HTML(http.StatusOK, "cg_cosmic_signature_tokens_by_user.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"UserTokens" : user_tokens,
	})
}
func cosmic_game_dev_donate_nft(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	cmd_str := fmt.Sprintf("%v/%v",os.Getenv("HOME"),"mint-artblocks.sh")
	cmd := exec.Command(cmd_str)
	buf := new(strings.Builder)
	err_buf:= new(strings.Builder)
	cmd.Stdout = buf
	cmd.Stderr = err_buf
	err := cmd.Run()
	output := buf.String()
	stderr := buf.String()
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("exec() failed: %v:\n%v\n%v",err,output,stderr),
		})
		return
	}
	c.HTML(http.StatusOK, "cg_dev_donate_nft.html", gin.H{
		"Output" : output,
	})

}
func cosmic_game_dev_funcs(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	c.HTML(http.StatusOK, "cg_dev_funcs.html", gin.H{
	})
}
func cosmic_game_token_name_history(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'token_id' parameter is not set")
		return
	}

	tokname_history := arb_storagew.Get_cosmic_signature_token_name_history(token_id)
	c.HTML(http.StatusOK, "cg_token_name_history.html", gin.H{
		"TokenId" : token_id,
		"TokenNameHistory" : tokname_history,
	})
}
func cosmic_game_token_name_search(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_name:= c.Param("name")
	if len(p_name) > 0 {
	} else {
		respond_error(c,"'name' parameter is not set")
		return
	}

	results := arb_storagew.Search_token_by_name(p_name)
	c.HTML(http.StatusOK, "cg_token_name_search_results.html", gin.H{
		"SearchText" : p_name,
		"TokenNameSearchResults" : results ,
	})
}
func cosmic_game_named_tokens_only(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	results := arb_storagew.Get_named_tokens()
	c.HTML(http.StatusOK, "cg_named_tokens_only.html", gin.H{
		"NamedTokens" : results ,
	})
}
func cosmic_game_token_ownership_transfers(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'token_id' parameter is not set")
		return
	}

	transfers := arb_storagew.Get_cst_ownership_transfers(token_id,0, 0)
	c.HTML(http.StatusOK, "cg_token_ownership_transfers.html", gin.H{
		"TokenId" : token_id,
		"TokenTransfers" : transfers,
	})
}
func cosmic_game_cs_token_distribution(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	distribution := arb_storagew.Get_cosmic_signature_token_distribution()

	c.HTML(http.StatusOK, "cg_cs_token_distribution.html", gin.H{
		"CosmicSignatureTokenDistribution" : distribution,
	})
}
func cosmic_game_user_balances(c *gin.Context) {

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

	addr := common.HexToAddress(p_user_addr)
	user_eth_bal,err := eclient.BalanceAt(context.Background(),addr,nil)
	if err != nil {
		err_str := fmt.Sprintf("Error at BalanceAt() call for addr: %v\n",err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("%v",err_str),
		})
		return
	}
	ct_contract,err := NewERC20(cosmic_token_addr,eclient);
	if err != nil {
		err_str := fmt.Sprintf("Error at instantiation of ERC20 contract: %v\n",err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("%v",err_str),
		})
		return
	}
	var copts bind.CallOpts
	ct_balance,err := ct_contract.BalanceOf(&copts,addr)
	if err != nil {
		err_str := fmt.Sprintf("Error at BalanceOf() call: %v\n",err)
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("%v",err_str),
		})
		return
	}

	c.HTML(http.StatusOK, "cg_user_balances.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"ETH_Balance" : user_eth_bal.String(),
		"CosmicTokenBalance" : ct_balance.String(),
	})
}
func cosmic_game_cosmic_token_balances(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	balances := arb_storagew.Get_cosmic_token_holders()
	c.HTML(http.StatusOK, "cg_cosmic_token_balances.html", gin.H{
		"CosmicTokenBalances" : balances,
	})
}
func cosmic_game_cosmic_token_transfers_by_user(c *gin.Context) {
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
	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	transfers := arb_storagew.Get_cosmic_token_transfers_by_user(user_aid,offset,limit)
	c.HTML(http.StatusOK, "cg_user_erc20_transfers.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"Offset" : offset,
		"Limit" : limit,
		"CosmicTokenTransfers" : transfers,
	})
}
func cosmic_game_cosmic_signature_transfers_by_user(c *gin.Context) {

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
	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	transfers := arb_storagew.Get_cosmic_signature_transfers_by_user(user_aid,offset,limit)
	c.HTML(http.StatusOK, "cg_user_erc721_transfers.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"Offset" : offset,
		"Limit" : limit,
		"CosmicSignatureTransfers" : transfers,
	})
}
func cosmic_game_used_rwalk_nfts(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	used_nfts := arb_storagew.Get_random_walk_tokens_in_bids()
	c.HTML(http.StatusOK, "cg_used_rwalk_tokens.html", gin.H{
		"UsedRwalkNFTs" : used_nfts,
	})
}
func cosmic_game_marketing_rewards_global(c *gin.Context) {
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	rewards := arb_storagew.Get_marketing_reward_history_global(0, 100000)
	c.HTML(http.StatusOK, "cg_marketing_rewards_global.html", gin.H{
		"MarketingRewards" : rewards,
	})
}
func cosmic_game_marketing_rewards_by_user(c *gin.Context) {
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
		respond_error(c,"Provided address wasn't found")
		return
	}
	rewards := arb_storagew.Get_marketing_reward_history_by_user(user_aid,0, 100000)
	c.HTML(http.StatusOK, "cg_marketing_rewards_by_user.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"UserMarketingRewards" : rewards,
	})
}
func cosmic_game_get_cst_price(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	var copts bind.CallOpts
	// Note: we are using BusinessLogic contract instead of CosmicGame because CurrentCSTPrice is a view 
	// function and will be called using 'Caller' interface instead of 'Transactor' interface, since
	// both function return a byte array of 32 bytes , this workaround will work, otherwise, we would
	// need to make explicit eth_call() method to CosmicGame contract (because the default is to transact
	// since the method is not declared as 'view')
	contract,err := NewCosmicSignatureGame(cosmic_game_addr,eclient)
	if err != nil {
		err_str := fmt.Sprintf("Can't instantiate CosmicGame contract: %v . Contract constants won't be fetched\n",err)
		Error.Printf(err_str)
		Info.Printf(err_str)
		respond_error(c,err_str)
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
				c.HTML(http.StatusOK, "cg_current_cst_price.html", gin.H{
					"CSTPrice": cst_price.String(),
					"SecondsElapsed" : seconds_elapsed.String(),
					"AuctionDuration" : auction_duration.String(),
				})
			}
		}
	}
}
func cosmic_game_sysmode_changes(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_html(c)
	if !success {
		return
	}
	system_mode_changes := arb_storagew.Get_system_mode_change_event_list(offset,limit)

	c.HTML(http.StatusOK, "cg_system_mode_changes.html", gin.H{
		"SystemModeChanges" : system_mode_changes,
		"Offset" : offset,
		"Limit" : limit,
	})
}
func cosmic_game_admin_events_in_range(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_evtlog_start:= c.Param("evtlog_start")
	var evtlog_start int64
	if len(p_evtlog_start) > 0 {
		var success bool
		evtlog_start,success = parse_int_from_remote_or_error(c,HTTP,&p_evtlog_start)
		if !success {
			return
		}
	} else {
		respond_error(c,"'evtlog_start' parameter is not set")
		return
	}
	p_evtlog_end := c.Param("evtlog_end")
	var evtlog_end int64
	if len(p_evtlog_end) > 0 {
		var success bool
		evtlog_end,success = parse_int_from_remote_or_error(c,HTTP,&p_evtlog_end)
		if !success {
			return
		}
	} else {
		respond_error(c,"'evtlog_end' parameter is not set")
		return
	}
	event_list := arb_storagew.Get_admin_events_in_range(evtlog_start,evtlog_end)

	c.HTML(http.StatusOK, "cg_system_admin_events_in_range.html", gin.H{
		"AdminEvents" : event_list,
		"EvtLogIdStart" : evtlog_start,
		"EvtLogIdEnd" : evtlog_end,
	})
}
func cosmic_game_bid_special_winners(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	c.HTML(http.StatusOK, "cg_current_special_winners.html", gin.H{
		"LastBidderAddress": last_bidder.String(),
		"LastBidderLastBidTime" : last_bidder_bid_time,
		"EnduranceChampionAddress": endurance_champ_addr,
		"EnduranceChampionDuration": endurance_duration,
		"LastCstBidderAddress" : lastcst_bidder_addr,
	})
}


