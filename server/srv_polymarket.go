package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func poly_buysell_operations(c *gin.Context) {

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

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	market_info,err := augur_srv.db_matic.Get_poly_market_info(market_id)
	if err != nil {
		respond_error(c,"Market not found")
		return
	}
	operations := augur_srv.db_matic.Get_polymarkets_buysell_operations(fpmm_aid,0,1000000)

	var js_outcomes_history JSOutcomes
	for outc:=0; outc<int(market_info.OutcomeSlotCount); outc++ {
		prices:= augur_srv.db_matic.Get_poly_market_outcome_price_history(fpmm_aid,int32(outc))
		js_prices := build_js_polymarkets_outcome_price_history(&prices)
		js_outcomes_history.OutcomesDataJS  = append(js_outcomes_history.OutcomesDataJS,js_prices)
	}
	prices:= augur_srv.db_matic.Get_poly_market_outcome_price_history(fpmm_aid,0)
	price0 := build_js_polymarkets_outcome_price_history(&prices)
	prices= augur_srv.db_matic.Get_poly_market_outcome_price_history(fpmm_aid,1)
	price1 := build_js_polymarkets_outcome_price_history(&prices)

	c.HTML(http.StatusOK, "buysell_operations.html", gin.H{
		"BuySellOperations" : operations,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"Prices" : js_outcomes_history,
		"Price0" : price0,
		"Price1" : price1,
	})
}
func poly_liquidity_operations(c *gin.Context) {

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

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	operations := augur_srv.db_matic.Get_polymarkets_liquidity_operations(fpmm_aid,0,1000000)

	c.HTML(http.StatusOK, "liquidity_operations.html", gin.H{
		"LiquidityOperations" : operations,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
	})
}
func poly_market_info(c *gin.Context) {

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
	info,err := augur_srv.db_matic.Get_poly_market_info(market_id)
	if err != nil {
		respond_error(c,"Market not found")
		return
	}
	c.HTML(http.StatusOK, "market_info.html", gin.H{
		"MarketInfo" : info,
		"MarketId" : market_id,
	})
}
func poly_market_stats(c *gin.Context) {

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
	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}
	stats,_:= augur_srv.db_matic.Get_poly_market_stats(fpmm_aid)
	c.HTML(http.StatusOK, "market_stats.html", gin.H{
		"MarketStats" : stats,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
	})
}
func poly_liq_hist_global(c *gin.Context) {

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	liq_hist := augur_srv.db_matic.Get_polymarket_global_liquidity_history(init_ts,fin_ts,interval_secs)

	c.HTML(http.StatusOK, "global_liquidity.html", gin.H{
		"GlobalLiquidityHistory" : liq_hist,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
	})
}
func poly_market_liquidity_periods(c *gin.Context) {

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
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
		respond_error(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	liq_hist := augur_srv.db_matic.Get_polymarket_market_liquidity_history(fpmm_aid,init_ts,fin_ts,interval_secs)

	c.HTML(http.StatusOK, "market_liquidity_by_periods.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"MarketLiquidityHistory" : liq_hist,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
	})
}
func poly_user_list(c *gin.Context) {

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

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	users_list := augur_srv.db_matic.Get_polymarkets_market_user_list(fpmm_aid)

	c.HTML(http.StatusOK, "market_userlist.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"Users" : users_list,
	})

}
func poly_market_trader_operations(c *gin.Context) {

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
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,false,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'user_aid' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	trade_list := augur_srv.db_matic.Get_poly_market_trader_operations(fpmm_aid,user_aid,offset,limit)

	c.HTML(http.StatusOK, "market_trader_operations.html", gin.H{
		"MarketId" : market_id,
		"UserAid" : user_aid,
		"ContractAid" : fpmm_aid,
		"TraderOperations" : trade_list,
	})
}
func poly_market_funder_operations(c *gin.Context) {

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
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,false,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'user_aid' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	liq_operation_list := augur_srv.db_matic.Get_poly_market_funder_operations(fpmm_aid,user_aid,offset,limit)

	c.HTML(http.StatusOK, "market_funder_operations.html", gin.H{
		"MarketId" : market_id,
		"UserAid" : user_aid,
		"ContractAid" : fpmm_aid,
		"FunderOperations" : liq_operation_list,
	})
}
func poly_market_open_positions(c *gin.Context) {

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

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	open_positions,prices := augur_srv.db_matic.Get_poly_market_open_positions(fpmm_aid)

	c.HTML(http.StatusOK, "market_open_positions.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"OpenPositions" : open_positions,
		"Prices" : prices,
	})
}
func poly_market_user_open_positions(c *gin.Context) {

	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,false,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'user_aid' parameter is not set")
		return
	}

	user_open_positions := augur_srv.db_matic.Get_poly_user_open_positions(user_aid)

	c.HTML(http.StatusOK, "market_user_open_positions.html", gin.H{
		"UserAid": user_aid,
		"UserOpenPositions" :user_open_positions,
	})
}
func poly_market_funder_share_ratio(c *gin.Context) {

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

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	share_ratios := augur_srv.db_matic.Get_poly_liquidity_provider_share_ratio(fpmm_aid)

	c.HTML(http.StatusOK, "market_funders_share_ratio.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"ShareRatios" : share_ratios,
	})
}
func poly_markets_listing(c *gin.Context) {

	p_status := c.Param("status")
	var status int64
	if len(p_status) > 0 {
		var success bool
		status,success = parse_int_from_remote_or_error(c,HTTP,&p_status)
		if !success {
			return
		}
	} else {
		// the default is status = 0
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
		// the default is sort = 0
	}

	category := c.Query("c")

	markets_listing := augur_srv.db_matic.Get_polymarkets_markets(int(status),int(sort),category)
	num_elts := len(markets_listing)
	c.HTML(http.StatusOK, "market_listing.html", gin.H{
		"Markets" : markets_listing,
		"QueryingStatus" : status,
		"NumElts" : num_elts,
	})
}
func poly_top_users(c *gin.Context) {

	top_profit_makers := augur_srv.db_matic.Get_polymarket_top_profit_makers()
	top_trade_makers := augur_srv.db_matic.Get_polymarket_top_trade_makers()
	top_volume_makers := augur_srv.db_matic.Get_polymarket_top_volume_makers()
	c.HTML(http.StatusOK, "poly_top_users.html", gin.H{
			"title": "Top 100 Users of Polymarket Markets",
			"ProfitMakers" : top_profit_makers,
			"TradeMakers" : top_trade_makers,
			"VolumeMakers" : top_volume_makers,
	})
}
func poly_market_payout_redemptions(c *gin.Context) {

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

	condition_id := augur_srv.db_matic.Get_condition_id(market_id)
	if len(condition_id) == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	payout_redemptions := augur_srv.db_matic.Get_polymarket_market_redemptions(condition_id,0,1000000)

	c.HTML(http.StatusOK, "market_redemptions.html", gin.H{
		"MarketId" : market_id,
		"PayoutRedemptions" : payout_redemptions,
	})
}
func poly_market_categories(c *gin.Context) {

	categories := augur_srv.db_matic.Get_polymarket_categories()
	c.HTML(http.StatusOK, "categories.html", gin.H{
		"MarketCategories" : categories,
	})
}
func poly_market_erc1155_transfers(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,HTTP,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	erc1155_transfers := augur_srv.db_matic.Get_polymarket_erc1155_transfers(fpmm_aid,offset,limit)

	c.HTML(http.StatusOK, "erc1155_transfers.html", gin.H{
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"ERC1155Transfers" : erc1155_transfers,
	})
}
func poly_market_open_interest_history(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,HTTP,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}
	condition_id := augur_srv.db_matic.Get_condition_id(market_id)
	caddrs := augur_srv.db_matic.Get_polymarket_contract_addresses()
	if caddrs.CondTokAid == 0 {
		respond_error(c,"Conditional token Aid is zero")
		return
	}
	totals,oi_hist := augur_srv.db_matic.Get_polymarket_open_interst_history_v2(
		caddrs.USDCAid,
		caddrs.CondTokAid,
		fpmm_aid,
		condition_id,
		offset,
		limit,
	)

	c.HTML(http.StatusOK, "open_interest_history.html", gin.H{
		"MarketId" : market_id,
		"MktMkrAid" : fpmm_aid,
		"OIHistory" : oi_hist,
		"CondTokAid" : caddrs.CondTokAid,
		"USDCAid" : caddrs.USDCAid,
		"Totals" : totals,
	})
}
func poly_market_search(c *gin.Context) {

	p_keyword:= c.Query("q")
	if len(p_keyword) == 0 {
		respond_error(c,"'q' parameter is not set")
		return
	}

	results := augur_srv.db_matic.Search_polymarket_keywords(p_keyword)

	c.HTML(http.StatusOK, "search_results.html", gin.H{
		"Keywords" : p_keyword,
		"SearchResults" : results,
	})
}
func poly_user_info(c *gin.Context) {

	p_user := c.Param("user")
	_,user_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}

	user_info,_ := augur_srv.db_matic.Get_polymarket_user_info(user_aid)

	c.HTML(http.StatusOK, "user_info.html", gin.H{
		"User": p_user,
		"UserInfo" :user_info,
	})
}
