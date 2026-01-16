/// API v1
package main
import (
	"fmt"
	"strconv"

	"net/http"
	"github.com/gin-gonic/gin"

)
func a1_user_uniswap_swaps(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,user_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	user_info,err := augur_srv.db_augur.Get_user_info(user_aid)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": fmt.Sprintf("Error getting UserInfo: %v",err.Error()),
		})
	}
	swaps,total_recs := augur_srv.db_augur.Get_user_uniswap_swaps(user_aid,offset,limit)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"UserInfo" : user_info,
		"UserSwaps" : swaps,
		"TotalRows" : total_recs,
		"status": status,
		"error": err_str,
	})
}
func a1_uniswap_calculate_slippage(c *gin.Context) {
	// Calculates slippage for swapping single token in the Pair
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_pair := c.Param("pair")
	pair_addr,pair_aid,success := json_validate_and_lookup_address_or_aid(c,&p_pair)
	if !success {
		return
	}
	p_tok_in := c.Param("tok_in")
	tok_in,_,success := json_validate_and_lookup_address_or_aid(c,&p_tok_in)
	if !success {
		return
	}
	einf,err := augur_srv.db_augur.Get_erc20_info(p_tok_in)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
	pair_info,err := augur_srv.db_augur.Get_uniswap_pair_info(pair_aid)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
	p_amount:= c.Param("amount")
	amount := fmt.Sprintf("%v%0*d",p_amount,einf.Decimals,0)
	slippage,token_amount_out,err := uniswap_calc_slippage(pair_addr,tok_in,amount)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
	var dec1,dec2 *int
	if pair_info.Token1Addr == p_tok_in { // figure out which token is the divisor
		dec1 = &pair_info.Token0Decimals
		dec2 = &pair_info.Token1Decimals
	} else {
		dec1 = &pair_info.Token1Decimals
		dec2 = &pair_info.Token0Decimals
	}
	uniswap_correct_for_difference_in_decimals(slippage,*dec1,*dec2)
	var status int = 1
	var err_str string = ""
	var amount_out_str string = "?"
	var slippage_str string = "?"
	if slippage != nil {
		slippage_str = slippage.String()
	}
	if token_amount_out != nil {
		amount_out_str = token_amount_out.String()
	}
	c.JSON(http.StatusOK, gin.H{
			"status": status,
			"error": err_str,
			"Slippage" : slippage_str,
			"AmountOut" : amount_out_str,
			"AmountToTrade": p_amount,
			"AmountToTradeWei" : amount,
	})
}
func a1_uniswap_slippage(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_pair:= c.Param("pair")
	_,pair_aid,success := json_validate_and_lookup_address_or_aid(c,&p_pair)
	if !success {
		return
	}
	pair_info,err := augur_srv.db_augur.Get_uniswap_pair_info(pair_aid)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
	//amount_to_trade := "100";
	//slippages,err := produce_uniswap_slippages(&pair_info,amount_to_trade)

	slippages := augur_srv.db_augur.Get_uniswap_latest_slippages(pair_aid)
	var amount_to_trade float64 = 0.0
	if len(slippages) > 0 {
		amount_to_trade = slippages[0].AmountIn
	}

	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"status": status,
			"error": err_str,
			"PairInfo" : pair_info,
			"AmountToTrade" : amount_to_trade,
			"TokenSlippages" : slippages,
	})
}
func a1_single_uniswap_swap(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_id := c.Param("id")
	var id int64
	var err error
	id, err = strconv.ParseInt(p_id,10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("Bad integer for 'id' parameter: %v",err),
		})
		return
	}

	swap,err := augur_srv.db_augur.Get_uniswap_swap_by_id(id)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"UniswapSwap" : swap,
			"Id": id,
			"status": status,
			"error": err_str,
	})
}
func a1_upair_price_history(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_pair:= c.Param("pair")
	_,pair_aid,success := json_validate_and_lookup_address_or_aid(c,&p_pair)
	if !success {
		return
	}

	var err error
	p_inverse := c.Param("inverse")
	var inverse int = 0
	if len(p_inverse) > 0 {
		inverse, err = strconv.Atoi(p_inverse)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("'inverse' parameter: %v",err),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("'inverse' parameter wasn't provided: %v",err),
		})
		return
	}

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	bool_inverse := false
	if inverse > 0 {
		bool_inverse = true
	}
	pair_info,_:= augur_srv.db_augur.Get_uniswap_pair_info(pair_aid)
	prices := augur_srv.db_augur.Get_uniswap_token_prices(pair_aid,bool_inverse,init_ts,fin_ts,interval_secs)
	c.JSON(http.StatusOK, gin.H{
			"PairInfo" : pair_info,
			"Prices" : prices,
			"InitTimeStamp": init_ts,
			"FinTimeSTamp": fin_ts,
			"Interval" : interval_secs,
	})
}
func a1_uniswap_pair_swaps(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_address:= c.Param("address")
	_,pair_aid,success := json_validate_and_lookup_address_or_aid(c,&p_address)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	pair_info,_ := augur_srv.db_augur.Get_uniswap_pair_info(pair_aid)
	swaps := augur_srv.db_augur.Get_uniswap_swaps(pair_aid,offset,limit)
	c.JSON(http.StatusOK, gin.H{
			"PairInfo" : pair_info,
			"PairSwaps" : swaps,
	})
}
func a1_uniswap_volume(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market := c.Param("market")
	_,market_aid,success := json_validate_and_lookup_address_or_aid(c,&p_market)
	if !success {
		return
	}
	success,outcome_idx:= parse_outcome_param(c)
	if !success {
		return
	}
	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}
	volume := augur_srv.db_augur.Get_uniswap_volume(market_aid,outcome_idx,init_ts,fin_ts,interval_secs)

	c.JSON(http.StatusOK,gin.H{
			"AllPairsVolume": volume,
			"MktAid" : market_aid,
			"MktAddr" : p_market,
			"OutcomeIdx" : outcome_idx,
			"status": 1 ,
			"error": "",
	})

}
func a1_market_uniswap_pairs(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_market:= c.Param("market")
	market_addr,market_aid,success := json_validate_and_lookup_address_or_aid(c,&p_market)
	if !success {
		return
	}

	pairs := augur_srv.db_augur.Get_market_uniswap_pairs(market_aid)
	c.JSON(http.StatusOK,gin.H{
		"MktAid": market_aid,
		"MktAddr" :market_addr,
		"MarketUniswapPairs": pairs,
		"status": 1,
		"error": "",
	})
}
