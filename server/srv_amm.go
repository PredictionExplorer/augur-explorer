package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

)
func arbitrum_markets_sports(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'status' parameter is not set")
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
		respond_error(c,"'sort' parameter is not set")
		return
	}
	contract_addrs := augur_srv.db_matic.Get_arbitrum_augur_factory_aids(&amm_contracts)
	fmt.Printf("contract_addrs = %+v\n",contract_addrs)
	total_rows,markets := augur_srv.db_matic.Get_sport_markets(status,sort,0,10000000,&amm_constants,contract_addrs)
	c.HTML(http.StatusOK, "arbitrum_markets_sports.html", gin.H{
		"Status" : status,
		"Markets" : markets,
		"TotalRows" : total_rows,
	})
}
func arbitrum_liquidity_changed(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	p_factory_aid:= c.Param("factory_aid")
	var factory_aid int64
	if len(p_factory_aid) > 0 {
		var success bool
		factory_aid,success = parse_int_from_remote_or_error(c,false,&p_factory_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'factory_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,factory_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	total_rows,lchanges := augur_srv.db_matic.Get_liquidity_change_events(
		factory_aid,market_id,0,10000000,
	)
	c.HTML(http.StatusOK, "augur_amm/liquidity_changed.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"LiquidityChanges" : lchanges,
		"TotalRows" : total_rows,
	})
}
func arbitrum_shares_swapped(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'market_id' parameter is not set")
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
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}

	total_rows,swaps:= augur_srv.db_matic.Get_shares_swapped(
		&amm_constants,contract_aid,market_id,0,10000000,
	)
	c.HTML(http.StatusOK, "augur_amm/shares_swapped.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"Swaps" : swaps,
		"TotalRows" : total_rows,
	})
}
func amm_user_swaps(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,p_user)
	if !valid {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	aid,err := augur_srv.db_matic.Nonfatal_lookup_address_id(user_addr)
	if err != nil {
		aid = 0
	}
	total_rows,swaps := augur_srv.db_matic.Get_amm_user_swaps(&amm_constants,aid,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/user_swaps.html", gin.H{
		"Swaps" : swaps,
		"TotalRows" : total_rows,
		"User":p_user,
		"UserAid":aid,
	})
}
func amm_user_liquidity(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,p_user)
	if !valid {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	aid,err := augur_srv.db_matic.Nonfatal_lookup_address_id(user_addr)
	if err != nil {
		aid = 0
	}
	total_rows,liquidity := augur_srv.db_matic.Get_amm_user_liquidity(&amm_constants,aid,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/user_liquidity.html", gin.H{
		"Liquidity" : liquidity,
		"TotalRows" : total_rows,
		"User": p_user,
		"UserAid": aid,
	})
}
func arbitrum_market_info(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'market_id' parameter is not set")
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
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}

	c.HTML(http.StatusOK, "augur_amm/market_info.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
	})
}
func arbitrum_market_liquidity_providers(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'market_id' parameter is not set")
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
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	pool_aid,err := augur_srv.db_matic.Get_market_pool_aid(contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Pool wasn't found in the database for this market: %v",err))
		return
	}
	providers:= augur_srv.db_matic.Get_pool_holder_distribution(pool_aid)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	js_tok_distr := build_js_token_holder_distribution(&providers)

	c.HTML(http.StatusOK, "augur_amm/liquidity_providers_distrib.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"PoolTokenHolderDistribution" : providers,
		"JSTokenHolderDistribution" : js_tok_distr,
	})
}
func arbitrum_market_outside_augur_shares_burned(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'market_id' parameter is not set")
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
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	offset := int(0) ; limit:= int(100000)
	operations := augur_srv.db_matic.Get_outside_augur_shares_burned(contract_aid,market_id,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/outside_augur_shares_bruned.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"SharesBurnedOperations" : operations,
	})

}
func arbitrum_market_outside_augur_shares_minted(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'market_id' parameter is not set")
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
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	offset := int(0) ; limit:= int(100000)
	operations := augur_srv.db_matic.Get_outside_augur_shares_minted(contract_aid,market_id,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/outside_augur_shares_minted.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"SharesMintedOperations" : operations,
	})

}
func arbitrum_market_outside_augur_balancer_swaps(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'market_id' parameter is not set")
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
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v has error: %v",market_id,err))
		return
	}
	pool_aid,err := augur_srv.db_matic.Get_market_pool_aid(contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Pool wasn't found in the database for this market: %v",err))
		return
	}
	pool_addr,_ := augur_srv.db_matic.Lookup_address(pool_aid)
	offset:=int(0);limit:=int(1000000000)
	balancer_swaps := augur_srv.db_matic.Get_outside_augur_balancer_swaps(pool_aid,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/balancer_swaps_outside_augur.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"PoolAid": pool_aid,
		"PoolAddr" : pool_addr,
		"BalancerSwaps" : balancer_swaps,
	})

}
func arbitrum_market_outside_augur_erc20_transfers(c *gin.Context) {

	if  !augur_srv.matic_initialized() {
		respond_error(c,"Database link wasn't configured")
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
		respond_error(c,"'market_id' parameter is not set")
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
		respond_error(c,"'contract_aid' parameter is not set")
		return
	}
	market,err := augur_srv.db_matic.Get_sport_market_info(&amm_constants,contract_aid,market_id)
	if err!=nil {
		respond_error(c,fmt.Sprintf("Market with market_id=%v couldn't be located, error: %v",market_id,err))
		return
	}
	offset:=int(0);limit:=int(1000000000)
	transfers := augur_srv.db_matic.Get_erc20_transfers_outside_augur(contract_aid,market_id,offset,limit)

	c.HTML(http.StatusOK, "augur_amm/erc20_transfers_outside_augur.html", gin.H{
		"MarketId":market_id,
		"MarketInfo" : market,
		"ERC20Transfers" : transfers,
	})
}
