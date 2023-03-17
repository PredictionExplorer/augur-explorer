/// API v1
package main
import (
	"fmt"

	"net/http"
	"github.com/gin-gonic/gin"

)
func a1_arbitrum_augur_pools(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	pools := augur_srv.db_matic.Get_arbitrum_augur_pools()
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"error" : err_str,
		"AMMPools" : pools,
	})
}
func a1_arbitrum_markets_sports(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_status := c.Param("status")
	var status int64
	if len(p_status) > 0 {
		var success bool
		status,success = parse_int_from_remote_or_error(c,false,&p_status)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'status' parameter is not set")
		return
	}
	p_sort := c.Param("sort")
	var sort int64
	if len(p_sort) > 0 {
		var success bool
		sort ,success = parse_int_from_remote_or_error(c,false,&p_sort)
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
	contract_addrs := augur_srv.db_matic.Get_arbitrum_augur_factory_aids(&amm_contracts)
	total_rows,markets := augur_srv.db_matic.Get_sport_markets(status,sort,offset,limit,&amm_constants,contract_addrs)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MStatus" : status,
		"Markets" : markets,
		"Offset" : offset,
		"Limit" : limit,
		"TotalRows" : total_rows,
	})
}
func a1_arbitrum_liquidity_changed(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("factory_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'contract_aid' parameter is not set")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Market with market_id=%v has error: %v",err))
		return
	}
	total_rows,lchanges := augur_srv.db_matic.Get_liquidity_change_events(
		contract_aid,market_id,offset,limit,
	)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId":market_id,
		"MarketInfo" : market,
		"LiquidityChanges" : lchanges,
		"Offset" : offset,
		"Limit" : limit,
		"TotalRows" : total_rows,
	})
}
func a1_arbitrum_shares_swapped(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("contract_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'contract_aid' parameter is not set")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	total_rows,swaps:= augur_srv.db_matic.Get_shares_swapped(
		&amm_constants,contract_aid,market_id,offset,limit,
	)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId":market_id,
		"Swaps" : swaps,
		"Offset" : offset,
		"Limit" : limit,
		"TotalRows" : total_rows,
	})
}
func a1_arbitrum_market_info_sports(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("factory_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'contract_aid' parameter is not set")
		return
	}

	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Market with market_id=%v has error: %v",err))
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId":market_id,
		"MarketInfo" : market,
	})
}
func a1_arbitrum_market_liquidity_providers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("factory_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'factory_aid' parameter is not set")
		return
	}

	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Market with market_id=%v has error: %v",err))
		return
	}
	pool_aid,err := augur_srv.db_matic.Get_market_pool_aid(contract_aid,market_id)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Pool wasn't found in the database for this market: %v",err))
		return
	}
	providers:= augur_srv.db_matic.Get_pool_holder_distribution(pool_aid)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Market with market_id=%v has error: %v",err))
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId":market_id,
		"MarketInfo" : market,
		"PoolTokenHolderDistribution" : providers,
	})
}
func a1_arbitrum_market_outside_augur_shares_burned(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("factory_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'factory_aid' parameter is not set")
		return
	}

	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Market with market_id=%v has error: %v",err))
		return
	}
	operations := augur_srv.db_matic.Get_outside_augur_shares_burned(contract_aid,market_id,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId":market_id,
		"MarketInfo" : market,
		"SharesBurnedOperations" : operations,
	})
}
func a1_arbitrum_market_outside_augur_shares_minted(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("factory_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'factory_aid' parameter is not set")
		return
	}

	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Market with market_id=%v has error: %v",err))
		return
	}
	operations := augur_srv.db_matic.Get_outside_augur_shares_minted(contract_aid,market_id,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId":market_id,
		"MarketInfo" : market,
		"SharesMintedOperations" : operations,
	})
}
func a1_arbitrum_market_outside_augur_balancer_swaps(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("factory_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'factory_aid' parameter is not set")
		return
	}

	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	pool_aid,err := augur_srv.db_matic.Get_market_pool_aid(contract_aid,market_id)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Pool wasn't found in the database for this market: %v",err))
		return
	}
	pool_addr,_ := augur_srv.db_matic.Lookup_address(pool_aid)
	balancer_swaps := augur_srv.db_matic.Get_outside_augur_balancer_swaps(pool_aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId":market_id,
		"MarketInfo" : market,
		"PoolAid": pool_aid,
		"PoolAddr" : pool_addr,
		"BalancerSwaps" : balancer_swaps,
	})

}
func a1_arbitrum_market_outside_augur_erc20_transfers(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_contract_aid := c.Param("factory_aid")
	var contract_aid int64
	if len(p_contract_aid) > 0 {
		var success bool
		contract_aid,success = parse_int_from_remote_or_error(c,false,&p_contract_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'factory_aid' parameter is not set")
		return
	}

	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error_json(c,fmt.Sprintf("Market with market_id=%v couldn't be located, error: %v",market_id,err))
		return
	}
	transfers := augur_srv.db_matic.Get_erc20_transfers_outside_augur(contract_aid,market_id,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId":market_id,
		"MarketInfo" : market,
		"ERC20Transfers" : transfers,
	})
}
