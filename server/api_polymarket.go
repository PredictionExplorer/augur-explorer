/// API v1
package main
import (
	"fmt"
	"strconv"

	"net/http"
	"github.com/gin-gonic/gin"

)
func a1_poly_buysell_operations(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	operations := augur_srv.db_matic.Get_polymarkets_buysell_operations(fpmm_aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"TradingOperations" : operations,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
	})
}
func a1_poly_liquidity_operations(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	operations := augur_srv.db_matic.Get_polymarkets_liquidity_operations(fpmm_aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"LiquidityOperations" : operations,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
	})
}
func a1_poly_market_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}

	info,err := augur_srv.db_matic.Get_poly_market_info(market_id)
	if err != nil {
		respond_error_json(c,"Market not found")
		return
	}
	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	prices := augur_srv.db_matic.Calculate_prices(fpmm_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketInfo" : info,
		"Prices" : prices,
		"MarketId" : market_id,
	})
}
func a1_poly_market_stats(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error_json(c,"Polymarket with this ID wasn't found")
		return
	}
	stats,_:= augur_srv.db_matic.Get_poly_market_stats(fpmm_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketStats" : stats,
		"MarketId" : market_id,
	})
}
func a1_poly_unique_users(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	success,init_ts,fin_ts := parse_timeframe_ini_fin(c,JSON)
	if !success {
		return
	}

	stats := augur_srv.db_matic.Get_polymarkets_unique_users_stats(init_ts,fin_ts)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UniquePolymarketPlatformUsers" : stats,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
	})

}
func a1_poly_liq_hist_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	liq_hist := augur_srv.db_matic.Get_polymarket_global_liquidity_history(init_ts,fin_ts,interval_secs)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"GlobalLiquidityHistory" : liq_hist,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
	})
}
func a1_poly_market_liquidity_periods(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error_json(c,"Polymarket with this ID wasn't found")
		return
	}

	liq_hist := augur_srv.db_matic.Get_polymarket_market_liquidity_history(fpmm_aid,init_ts,fin_ts,interval_secs)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"MarketLiquidityHistory" : liq_hist,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
	})
}
func a1_poly_trade_hist_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	trade_hist := augur_srv.db_matic.Get_polymarket_global_trading_history(init_ts,fin_ts,interval_secs)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"GlobalTradeHistory" : trade_hist,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
	})
}
func a1_poly_market_trading_periods(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error_json(c,"Polymarket with this ID wasn't found")
		return
	}

	trade_hist := augur_srv.db_matic.Get_polymarket_market_trading_history(fpmm_aid,init_ts,fin_ts,interval_secs)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"MarketTradingHistory" : trade_hist,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
	})
}
func a1_poly_datafeed(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	evtlog_id := augur_srv.db_matic.Get_data_feed_status()
	new_evtlog_id,data_feed := augur_srv.db_matic.Polymarkets_data_feed(evtlog_id)
	augur_srv.db_matic.Update_data_feed_status(new_evtlog_id)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"DataFeed" : data_feed,
		"LastEvtId" : new_evtlog_id,
	})
}
func a1_poly_user_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error_json(c,"Polymarket with this ID wasn't found")
		return
	}

	users_list := augur_srv.db_augur.Get_polymarkets_market_user_list(fpmm_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"Users" : users_list,
	})
}
func a1_poly_trader_operations(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,true,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'user_aid' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_augur.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error_json(c,"Polymarket with this ID wasn't found")
		return
	}

	trade_list := augur_srv.db_augur.Get_poly_market_trader_operations(fpmm_aid,user_aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId" : market_id,
		"UserAid" : user_aid,
		"ContractAid" : fpmm_aid,
		"TraderOperations" : trade_list,
	})
}
func a1_poly_funder_operations(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,true,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'user_aid' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_augur.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error_json(c,"Polymarket with this ID wasn't found")
		return
	}

	liq_operation_list := augur_srv.db_augur.Get_poly_market_funder_operations(fpmm_aid,user_aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId" : market_id,
		"UserAid" : user_aid,
		"ContractAid" : fpmm_aid,
		"FunderOperations" : liq_operation_list,
	})
}
func a1_poly_markets_listing(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_status := c.Param("status")
	var status int64
	if len(p_status) > 0 {
		var success bool
		status,success = parse_int_from_remote_or_error(c,true,&p_status)
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
		sort,success = parse_int_from_remote_or_error(c,true,&p_sort)
		if !success {
			return
		}
	} else {
		// the default is status = 0
	}
	category := c.Query("c")
	markets_listing := augur_srv.db_augur.Get_polymarkets_markets(int(status),int(sort),category)
	num_elts := len(markets_listing)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Markets" : markets_listing,
		"QueryingStatus" : status,
		"Sort" : sort,
		"NumElts" : num_elts,
	})
}
func a1_poly_market_open_positions(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_augur.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error_json(c,"Polymarket with this ID wasn't found")
		return
	}

	open_positions,prices := augur_srv.db_augur.Get_poly_market_open_positions(fpmm_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"OpenPositions" : open_positions,
		"Prices" : prices,
	})
}
func a1_poly_market_user_open_positions(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,true,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'user_aid' parameter is not set")
		return
	}

	user_open_positions := augur_srv.db_augur.Get_poly_user_open_positions(user_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAid": user_aid,
		"UserOpenPositions" :user_open_positions,
	})
}
func a1_poly_market_funder_share_ratio(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error_json(c,"Polymarket with this ID wasn't found")
		return
	}

	share_ratios := augur_srv.db_matic.Get_poly_liquidity_provider_share_ratio(fpmm_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"ShareRatios" : share_ratios,
	})
}
func a1_poly_market_price_history(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,true,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	p_outcome := c.Param("outcome")
	var outcome int64
	if len(p_outcome) > 0 {
		var success bool
		outcome,success = parse_int_from_remote_or_error(c,true,&p_outcome)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'outcome' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error_json(c,"Polymarket with this ID wasn't found")
		return
	}

	prices:= augur_srv.db_matic.Get_poly_market_outcome_price_history(fpmm_aid,int32(outcome))

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"PriceHistory" : prices,
		"OutcomeIdx" : outcome,
	})
}
func a1_poly_market_buysell_info(c*gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_id:= c.Param("id")
	var id int64
	if len(p_id) > 0 {
		var success bool
		id,success = parse_int_from_remote_or_error(c,true,&p_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'id' parameter is not set")
		return
	}

	var req_status int = 1
	var err_str string = ""

	op_info,err := augur_srv.db_matic.Get_buysell_operation_info(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": req_status,
			"error" : err_str,
			"NotFound" : true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"OperationId" : id,
		"OperationInfo" : op_info,
		"NotFound" : false,
	})
}
func a1_poly_top_users(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var err error
	p_sort := c.Query("sort")
	var sort int = 0
	if len(p_sort) > 0 {
		sort, err = strconv.Atoi(p_sort)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"MarketIDs": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad 'sort' parameter: %v",err),
			})
			return
		}
	}
	ord_str := c.Query("ord")
	var ord int = 0
	if len(ord_str) > 0 {
		ord, err = strconv.Atoi(ord_str)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"UserRanks": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad 'order' parameter: %v",err),
			})
			return
		}
	}

	user_ranks := augur_srv.db_matic.Get_polymarket_user_ranks(sort,ord)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"UserRanks" : user_ranks,
			"status": status,
			"error": err_str,
	})
}
func a1_poly_market_payout_redemptions(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,JSON,&p_market_id)
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

	condition_id := augur_srv.db_matic.Get_condition_id(market_id)
	if len(condition_id) == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	payout_redemptions := augur_srv.db_matic.Get_polymarket_market_redemptions(condition_id,offset,limit)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"status": status,
		"error": err_str,
		"MarketId" : market_id,
		"PayoutRedemptions" : payout_redemptions,
	})
}
func a1_poly_categories(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	categories := augur_srv.db_matic.Get_polymarket_categories()
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"status": status,
		"error": err_str,
		"Categories" : categories,
	})

}
func a1_poly_market_erc1155_transfers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,JSON,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
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

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"status": status,
		"error": err_str,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"ERC1155Transfers" : erc1155_transfers,
	})
}
func a1_market_open_interest_history(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,JSON,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'market_id' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error_json(c,"Polymarket with this ID wasn't found")
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

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"status": status,
		"error": err_str,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"OIHistory" : oi_hist,
		"CondTokAid" : caddrs.CondTokAid,
		"USDCAid" : caddrs.USDCAid,
		"Totals" : totals,
	})
}
func a1_poly_market_search(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_keyword:= c.Query("q")
	if len(p_keyword) == 0 {
		respond_error_json(c,"'q' parameter is not set")
		return
	}

	results := augur_srv.db_matic.Search_polymarket_keywords(p_keyword)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"status": status,
		"error": err_str,
		"Keywords" : p_keyword,
		"SearchResults" : results,
	})
}
func a1_poly_user_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_user := c.Param("user")
	_,user_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}

	user_info,_ := augur_srv.db_matic.Get_polymarket_user_info(user_aid)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"status": status,
		"error": err_str,
		"User": p_user,
		"UserInfo" :user_info,
	})
}
func a1_poly_user_traded_markets(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_user := c.Param("user")
	_,user_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}

	user_info,_ := augur_srv.db_matic.Get_polymarket_user_info(user_aid)
	markets := augur_srv.db_matic.Get_polymarket_markets_by_user(user_aid)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"status": status,
		"error": err_str,
		"User": p_user,
		"UserInfo" :user_info,
		"MarketList":markets,
	})
}
