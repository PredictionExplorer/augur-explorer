/// API v1
package main
import (
	"fmt"
	"strconv"

	"net/http"
	"github.com/gin-gonic/gin"

)
func a1_user_balancer_swaps(c *gin.Context) {

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
	swaps,total_rows := augur_srv.db_augur.Get_user_balancer_swaps(user_aid,offset,limit)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"status": status,
			"error" : err_str,
			"UserInfo" : user_info,
			"PoolSwaps" : swaps,
			"TotalRows" : total_rows,
	})
}
func a1_balancer_calculate_slippage(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_pool := c.Param("pool")
	pool_addr,_,success := json_validate_and_lookup_address_or_aid(c,&p_pool)
	if !success {
		return
	}
	p_tok_in := c.Param("tok_in")
	tok_in,_,success := json_validate_and_lookup_address_or_aid(c,&p_tok_in)
	if !success {
		return
	}
	p_tok_out := c.Param("tok_out")
	tok_out,_,success := json_validate_and_lookup_address_or_aid(c,&p_tok_out)
	if !success {
		return
	}
	p_amount:= c.Param("amount")
	slippage,token_amount_out,err := balancer_calc_slippage(pool_addr,tok_in,tok_out,p_amount)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
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
	})
}
func a1_pool_slippage(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_pool:= c.Param("pool")
	_,pool_aid,success := json_validate_and_lookup_address_or_aid(c,&p_pool)
	if !success {
		return
	}
	pool_info,_ := augur_srv.db_augur.Get_pool_info(pool_aid)

	tokens := augur_srv.db_augur.Get_balancer_latest_slippages(pool_aid)
	var amount_to_trade float64 = 0.0
	if len(tokens) > 0 {
		amount_to_trade = tokens[0].AmountIn
	}

	//tokens := produce_pool_slippages(amount_to_trade,pool_aid)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"status": status,
			"error": err_str,
			"PoolInfo" : pool_info,
			"TokenSlippages" : tokens,
			"AmountToTrade" : amount_to_trade,
	})
}
func a1_single_balancer_swap(c *gin.Context) {

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

	swap,err := augur_srv.db_augur.Get_balancer_swap_by_id(id)
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
			"BalancerSwap" : swap,
			"Id": id,
			"status": status,
			"error": err_str,
	})
}
func a1_pool_price_history(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_pool:= c.Param("pool")
	_,pool_aid,success := json_validate_and_lookup_address_or_aid(c,&p_pool)
	if !success {
		return
	}

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	p_token1:= c.Param("token1")
	_,token1_aid,success := json_validate_and_lookup_address_or_aid(c,&p_token1)
	if !success {
		return
	}
	p_token2:= c.Param("token2")
	_,token2_aid,success := json_validate_and_lookup_address_or_aid(c,&p_token2)
	if !success {
		return
	}

	pool_info,_ := augur_srv.db_augur.Get_pool_info(pool_aid)
	token1_info,_ := augur_srv.db_augur.Get_bpool_token_info(pool_aid,token1_aid)
	token2_info,_ := augur_srv.db_augur.Get_bpool_token_info(pool_aid,token2_aid)
	prices := augur_srv.db_augur.Get_balancer_token_prices(pool_aid,token1_aid,token2_aid,init_ts,fin_ts,interval_secs)
	c.JSON(http.StatusOK, gin.H{
			"PoolInfo" : pool_info,
			"Token1Info" : token1_info,
			"Token2Info" : token2_info,
			"Prices" : prices,
			"InitTimeStamp": init_ts,
			"FinTimeSTamp": fin_ts,
			"Interval" : interval_secs,
	})
}
func a1_balancer_volume(c *gin.Context) {

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
	volume := augur_srv.db_augur.Get_balancer_volume(market_aid,outcome_idx,init_ts,fin_ts,interval_secs)

	c.JSON(http.StatusOK,gin.H{
			"AllPoolsVolume": volume,
			"MktAid" : market_aid,
			"MktAddr" : p_market,
			"OutcomeIdx" : outcome_idx,
			"status": 1 ,
			"error": "",
	})

}
func a1_pool_swaps(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")


	p_address:= c.Param("address")
	_,pool_aid,success := json_validate_and_lookup_address_or_aid(c,&p_address)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	pool_info,_ := augur_srv.db_augur.Get_pool_info(pool_aid)
	swaps := augur_srv.db_augur.Get_pool_swaps(pool_aid,offset,limit)
	c.JSON(http.StatusOK, gin.H{
			"PoolInfo" : pool_info,
			"PoolSwaps" : swaps,
	})
}
