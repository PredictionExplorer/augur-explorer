/// API v1
package main
import (
	"fmt"
	"strconv"
	"strings"
	"encoding/hex"

	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ethereum/go-ethereum/common"

	ens "github.com/wealdtech/go-ens/v3"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)

func a1_active_market_ids(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	off_str := c.Query("off")
	var off int = 0
	var err error
	if len(off_str) > 0 {
		off, err = strconv.Atoi(off_str)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"MarketIDs": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad offset parameter: %v",err),
			})
			return
		}
	}
	p_sort := c.Query("sort")
	var sort int = 0
	if len(p_sort) > 0 {
		sort, err = strconv.Atoi(p_sort)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"MarketIDs": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad sort parameter: %v",err),
			})
			return
		}
	}
	p_all := c.Query("all")
	var all int = 1
	if len(p_all) > 0 {
		all , err = strconv.Atoi(p_all)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"MarketIDs": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad 'all' parameter: %v",err),
			})
			return
		}
	}
	p_fin := c.Query("fin")
	var fin int = 0
	if len(p_fin) > 0 {
		fin , err = strconv.Atoi(p_fin)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"MarketIDs": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad 'fin' parameter: %v",err),
			})
			return
		}
	}
	p_alive:= c.Query("alive")
	var alive int = 0
	if len(p_alive) > 0 {
		alive , err = strconv.Atoi(p_alive)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"MarketIDs": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad 'alive' parameter: %v",err),
			})
			return
		}
	}
	p_inval_thresh := c.Query("it")
	var inval_thresh int = 0
	if len(p_inval_thresh) > 0 {
		inval_thresh , err = strconv.Atoi(p_inval_thresh)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"MarketIDs": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad 'it' (invalid_threshold) parameter: %v",err),
			})
			return
		}
	}
	ids := augur_srv.storage.Get_active_market_ids(sort,all,fin,alive,inval_thresh,off,1000000)
	var status int = 1
	c.JSON(http.StatusOK,gin.H{
		"MarketIDs": ids,
		"status":status,
		"error":"",
	})
}
func a1_market_card(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var market_aid int64 = 0
	p_market_aid := c.Param("market_aid")
	if len(p_market_aid ) > 0 {
		market_aid, err = strconv.ParseInt(p_market_aid,10,64)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"MarketInfo": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad integer for market_aid parameter: %v",err),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"MarketInfo": make([]int64,0,0),
			"status":0,
			"error":fmt.Sprintf("market_aid parameter wasn't provided"),
		})
		return
	}
	mkt_info,err := augur_srv.storage.Get_market_card_data(market_aid)
	var status int = 0
	var err_str string = ""
	if err == nil {
		status = 1
	} else {
		err_str = err.Error()
	}
	c.JSON(http.StatusOK,gin.H{
		"MarketInfo": mkt_info,
		"status":status,
		"error":err_str,
	})
}
func a1_active_markets(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_start := c.Param("start")
	p_num_rows := c.Param("num_rows")

	var start,num_rows int
	var err error
	if len(p_start) > 0 {
		start, err = strconv.Atoi(p_start)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"Markets": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad number for 'start' parameter: %v",err),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"Markets": make([]int64,0,0),
			"status":0,
			"error":fmt.Sprintf("Empty 'start' parameter for page offset"),
		})
	}
	if len(p_num_rows) > 0 {
		num_rows, err = strconv.Atoi(p_num_rows)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"Markets": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad number for 'num_rows' parameter: %v",err),
			})
			return
		}
		if num_rows < 1 {
			num_rows = DEFAULT_MARKET_ROWS_LIMIT
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"Markets": make([]int64,0,0),
			"status":0,
			"error":fmt.Sprintf("Empty 'num_rows' parameter for number of records to return"),
		})
	}

	markets := augur_srv.storage.Get_active_market_list(start,num_rows)
	var status int = 1
	c.JSON(http.StatusOK,gin.H{
		"Markets": markets,
		"status":status,
		"error":"",
	})
}
func a1_market_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	p_market := c.Param("market")
	if len(p_market) > 0 {
		market_aid, err := strconv.ParseInt(p_market,10,64)
		if err == nil {
			p_market,err = augur_srv.storage.Lookup_address(market_aid)
			if err!=nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"status":0,
					"error":fmt.Sprintf("Market with ID=%v not found",market_aid),
				})
				return
			}
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("Empty 'market' parameter for market lookup"),
		})
		return
	}
	market_addr,valid:=is_address_valid(c,true,p_market)
	if !valid {
		return
	}
	market_info,err := augur_srv.storage.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,true,&market_addr)
		return
	}
	complete_and_output_market_info(c,true,market_info)
}
func a1_multiple_market_cards(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	markets := make([]InfoMarket,0,64)
	p_market_aid_list := c.Param("market_aid_list")
	aid_list := strings.Split(p_market_aid_list,",")
	for i,aid_str := range aid_list {
		if len(aid_str ) > 0 {
			market_aid, err := strconv.ParseInt(aid_str,10,64)
			if err != nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"Markets": make([]InfoMarket,0,0),
					"status":0,
					"error":fmt.Sprintf("Bad integer for market_aid parameter at position %v : %v",i,err),
				})
				return
			}
			mkt_info,err := augur_srv.storage.Get_market_card_data(market_aid)
			if err!=nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"Markets": make([]InfoMarket,0,0),
					"status":0,
					"error":fmt.Sprintf("Error getting market by ID (id=%v): %v",market_aid,err),
				})
				return
			}
			markets = append(markets,mkt_info)
		} else {
			c.JSON(http.StatusBadRequest,gin.H{
				"Markets": make([]InfoMarket,0,0),
				"status":0,
				"error":fmt.Sprintf("market_aid at position %v is empty",i+1),
			})
			return
		}
	}
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"Markets": markets,
		"status":status,
		"error":err_str,
	})
}
func a1_user_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,eoa_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	Info.Printf("eoa_aid=%v\n",eoa_aid)
	gas_spent,_ := augur_srv.storage.Get_gas_spent_for_user(eoa_aid)
	user_info,err := augur_srv.storage.Get_user_info(eoa_aid)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"UserInfo" : user_info,
			"GasSpent" : gas_spent,
		})
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("User not found"),
		})
	}
}
func a1_user_funds(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	user_addr,_,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	serve_user_funds_v2(c,&user_addr)
}
func a1_user_traded_markets(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,eoa_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	p_active_flag := c.Param("active")
	var active_flag int = 1
	if len(p_active_flag) > 0 {
		var err error
		active_flag, err = strconv.Atoi(p_active_flag)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"Markets": make([]InfoMarket,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad 'active' parameter: %v",err),
			})
			return
		}
	}
	user_markets := augur_srv.storage.Get_traded_markets_for_user(eoa_aid,active_flag)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"Markets": user_markets,
		"status": status,
		"error": err_str,
	})
}
func a1_user_created_markets(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,eoa_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	created_markets := augur_srv.storage.Get_created_markets_for_user(eoa_aid)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"Markets": created_markets,
		"status": status,
		"error": err_str,
	})
}
func a1_user_reports(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,eoa_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	user_reports := augur_srv.storage.Get_user_reports(eoa_aid,10000000)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"Reports": user_reports ,
		"status": status ,
		"error": err_str,
	})
}
func a1_user_trades_for_market(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,eoa_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}

	p_market := c.Param("market")
	_,market_aid,success := json_validate_and_lookup_address_or_aid(c,&p_market)
	if !success {
		return
	}
	var status int = 1
	var err_str string = ""
	trades := augur_srv.storage.Get_user_trades_for_market(eoa_aid,market_aid)
	c.JSON(http.StatusOK,gin.H{
		"UTrades" : trades,
		"status": status,
		"error": err_str,
	})
}
func a1_user_profit_loss(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,eoa_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	pl_entries := augur_srv.storage.Get_profit_loss(eoa_aid)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"ProfitLossEntries" : pl_entries,
		"status": status,
		"error": err_str,
	})
}
func a1_user_open_positions(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,eoa_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	open_pos_entries := augur_srv.storage.Get_open_positions(eoa_aid)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"OpenPositionEntries" : open_pos_entries,
		"status": status,
		"error": err_str,
	})
}
func a1_user_open_orders(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,eoa_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	open_orders := augur_srv.storage.Get_user_open_orders(eoa_aid)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"OpenOrders" : open_orders,
		"status": status,
		"error": err_str,
	})
}
func a1_market_open_orders(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_market := c.Param("market")
	_,market_aid,success := json_validate_and_lookup_address_or_aid(c,&p_market)
	if !success {
		return
	}

	var err error
	p_outcome := c.Param("outcome")
	var outcome int
	if len(p_outcome) > 0 {
		outcome , err = strconv.Atoi(p_outcome)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("Bad outcome parameter: %v",err),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("Outcome parameter wasn't provided: %v",err),
		})
		return
	}

	mdepth,last_oo_id := augur_srv.storage.Get_mkt_depth(market_aid,outcome)
	num_orders:=len(mdepth.Bids) + len(mdepth.Asks)
	c.JSON(http.StatusOK,gin.H{
			"LastOOID": last_oo_id,
			"NumOrders" : num_orders,
			"Bids": mdepth.Bids,
			"Asks": mdepth.Asks,
			"OutcomeIdx" : outcome,
			"status": 1 ,
			"error": "",
	})
}
func a1_top_users(c *gin.Context) {

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

	user_ranks := augur_srv.storage.Get_user_ranks(sort,ord)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"UserRanks" : user_ranks,
			"status": status,
			"error": err_str,
	})
}
func a1_stats_main(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	stats := augur_srv.storage.Get_main_stats()

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"MainStats" : stats,
			"status": status,
			"error": err_str,
	})
}
func a1_stats_cashflow(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	cash_flow_entries := augur_srv.storage.Get_cash_flow()

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"CashFlow" : cash_flow_entries,
			"status": status,
			"error": err_str,
	})
}
func a1_stats_gasusage(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	gas_usage := augur_srv.storage.Get_gas_usage_global()

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"GasUsage" : gas_usage,
			"status": status,
			"error": err_str,
	})
}
func a1_stats_gasused(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	gas_usage := augur_srv.storage.Get_gas_used()

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"GasUsed" : gas_usage,
			"status": status,
			"error": err_str,
	})
}
func a1_stats_txcost(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	success,init_ts,fin_ts := parse_timeframe_ini_fin(c)
	if !success {
		return
	}
	tx_cost := augur_srv.storage.Get_txcost(init_ts,fin_ts)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"TxCost" : tx_cost ,
			"status": status,
			"error": err_str,
	})
}
func a1_stats_uniqaddr(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	uniq_addrs := augur_srv.storage.Get_unique_addresses()

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"UniqueAddresses" : uniq_addrs,
			"status": status,
			"error": err_str,
	})
}
func a1_price_history_zoomed(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_market := c.Param("market")
	market_addr,market_aid,success := json_validate_and_lookup_address_or_aid(c,&p_market)
	if !success {
		return
	}

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}
	price_history := augur_srv.storage.Get_zoomed_price_history(
		market_addr,market_aid,init_ts,fin_ts,interval_secs,
	)
	c.JSON(http.StatusOK,gin.H{
		"PriceHistory": price_history,
		"status": 1 ,
		"error": "",
	})
}
func a1_stats_accum_trades(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}
	trades := augur_srv.storage.Get_accumulated_trades_all_markets(int(init_ts),int(fin_ts),int(interval_secs))

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"AccumulatedTrades" : trades,
			"status": status,
			"error": err_str,
	})
}
func a1_stats_accum_oi(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}
	oi := augur_srv.storage.Get_accumulated_open_interest_all_markets_v3(init_ts,fin_ts,interval_secs)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"AccumulatedOpenInterest" : oi,
			"status": status,
			"error": err_str,
	})
}
func a1_stats_gasused_accum(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}
	gasused  := augur_srv.storage.Get_gasused_accum(init_ts,fin_ts,interval_secs)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"GasUsed" : gasused,
			"status": status,
			"error": err_str,
	})
}
func a1_stats_txcost_accum(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}
	txcost  := augur_srv.storage.Get_txcost_accum(init_ts,fin_ts,interval_secs)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"TxCost" : txcost,
			"status": status,
			"error": err_str,
	})
}
func a1_mkt_trade_history(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_market := c.Param("market")
	mkt_addr,market_aid,success := json_validate_and_lookup_address_or_aid(c,&p_market)
	if !success {
		return
	}
	lo_price,err := augur_srv.storage.Get_market_lo_price(market_aid)
	if err!=nil {
		c.JSON(http.StatusOK,gin.H{
				"status": 0,
				"error": err.Error(),
		})
		return
	}

	price_history := augur_srv.storage.Get_full_price_history(mkt_addr,market_aid,lo_price)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"UTrades" : price_history,
		"status": status,
		"error": err_str,
	})
}
func a1_wrapped_tokens(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_market:= c.Param("market")
	market_addr,_,success := json_validate_and_lookup_address_or_aid(c,&p_market)
	if !success {
		return
	}
	market_info,err := augur_srv.storage.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,false,&market_addr)
		return
	}
	wrappers := augur_srv.storage.Get_wrapped_tokens_for_market(market_info.MktAid)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"MarketInfo": market_info,
		"WrappedContracts": wrappers,
		"status": status ,
		"error": err_str,
	})
}
func a1_wrapped_token_transfers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_address:= c.Param("address")
	_,wrapper_aid,success := json_validate_and_lookup_address_or_aid(c,&p_address)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	wrapper_info,_ := augur_srv.storage.Get_wrapped_token_info(wrapper_aid)
	market_info,_ := augur_srv.storage.Get_market_info(wrapper_info.MktAddr,wrapper_info.OutcomeIdx,true)
	transfers,total_rows := augur_srv.storage.Get_wrapped_token_transfers(wrapper_aid,offset,limit)
	c.JSON(http.StatusOK, gin.H{
			"MarketInfo" : market_info,
			"TokenInfo" : wrapper_info,
			"TotalRows" : total_rows,
			"Offset" : offset,
			"Limit" : limit,
			"WrappedTransfers" : transfers,
	})
}
func a1_pool_swaps(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")


	p_address:= c.Param("address")
	_,pool_aid,success := json_validate_and_lookup_address_or_aid(c,&p_address)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	pool_info,_ := augur_srv.storage.Get_pool_info(pool_aid)
	swaps := augur_srv.storage.Get_pool_swaps(pool_aid,offset,limit)
	c.JSON(http.StatusOK, gin.H{
			"PoolInfo" : pool_info,
			"PoolSwaps" : swaps,
	})
}
func a1_market_share_token_balance_changes(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_market := c.Param("market")
	_,market_aid,success := json_validate_and_lookup_address_or_aid(c,&p_market)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	balance_changes,nr := augur_srv.storage.Outside_augur_share_balance_changes(market_aid,offset,limit)
	c.JSON(http.StatusOK,gin.H{
			"OutsideAugurBalanceChanges": balance_changes,
			"TotalRows" : nr,
			"Offset" : offset,
			"Limit" : limit,
			"status": 1 ,
			"error": "",
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
	volume := augur_srv.storage.Get_balancer_volume(market_aid,outcome_idx,init_ts,fin_ts,interval_secs)

	c.JSON(http.StatusOK,gin.H{
			"AllPoolsVolume": volume,
			"MktAid" : market_aid,
			"MktAddr" : p_market,
			"OutcomeIdx" : outcome_idx,
			"status": 1 ,
			"error": "",
	})

}
func a1_wrapped_token_volume(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_wrapper:= c.Param("address")
	_,wrapper_aid,success := json_validate_and_lookup_address_or_aid(c,&p_wrapper)
	if !success {
		return
	}
	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}
	volume := augur_srv.storage.Get_wrapped_transfers_volume(wrapper_aid,init_ts,fin_ts,interval_secs)

	c.JSON(http.StatusOK,gin.H{
			"Volume": volume,
			"WrapperAid" : wrapper_aid,
			"WrapperAddr" : p_wrapper,
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

	pairs := augur_srv.storage.Get_market_uniswap_pairs(market_aid)
	c.JSON(http.StatusOK,gin.H{
		"MktAid": market_aid,
		"MktAddr" :market_addr,
		"MarketUniswapPairs": pairs,
		"status": 1,
		"error": "",
	})
}
func a1_uniswap_pair_swaps(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_address:= c.Param("address")
	_,pair_aid,success := json_validate_and_lookup_address_or_aid(c,&p_address)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	pair_info,_ := augur_srv.storage.Get_uniswap_pair_info(pair_aid)
	swaps := augur_srv.storage.Get_uniswap_swaps(pair_aid,offset,limit)
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
	volume := augur_srv.storage.Get_uniswap_volume(market_aid,outcome_idx,init_ts,fin_ts,interval_secs)

	c.JSON(http.StatusOK,gin.H{
			"AllPairsVolume": volume,
			"MktAid" : market_aid,
			"MktAddr" : p_market,
			"OutcomeIdx" : outcome_idx,
			"status": 1 ,
			"error": "",
	})

}
func a1_search(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	keyword := c.Query("q")
	if len(keyword) == 0 {
		c.JSON(http.StatusOK,gin.H{
			"status": 1 ,
			"error": "Empty keyword",
		})
		return
	}
	sro := execute_search(keyword)
	var status int = 1
	if len(sro.ErrStr) > 0 {
		status = 0
	}
	c.JSON(http.StatusOK,gin.H{
			"SearchResult": sro,
			"status": status ,
			"error": sro.ErrStr,
	})
}
func a1_categories(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	categories := augur_srv.storage.Get_categories()

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
			"Categories" : categories,
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

	pool_info,_ := augur_srv.storage.Get_pool_info(pool_aid)
	token1_info,_ := augur_srv.storage.Get_bpool_token_info(pool_aid,token1_aid)
	token2_info,_ := augur_srv.storage.Get_bpool_token_info(pool_aid,token2_aid)
	prices := augur_srv.storage.Get_balancer_token_prices(pool_aid,token1_aid,token2_aid,init_ts,fin_ts,interval_secs)
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
	pair_info,_:= augur_srv.storage.Get_uniswap_pair_info(pair_aid)
	prices := augur_srv.storage.Get_uniswap_token_prices(pair_aid,bool_inverse,init_ts,fin_ts,interval_secs)
	c.JSON(http.StatusOK, gin.H{
			"PairInfo" : pair_info,
			"Prices" : prices,
			"InitTimeStamp": init_ts,
			"FinTimeSTamp": fin_ts,
			"Interval" : interval_secs,
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

	swap,err := augur_srv.storage.Get_uniswap_swap_by_id(id)
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

	swap,err := augur_srv.storage.Get_balancer_swap_by_id(id)
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
	pool_info,_ := augur_srv.storage.Get_pool_info(pool_aid)

	tokens := augur_srv.storage.Get_balancer_latest_slippages(pool_aid)
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
	einf,err := augur_srv.storage.Get_erc20_info(p_tok_in)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
	pair_info,err := augur_srv.storage.Get_uniswap_pair_info(pair_aid)
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
	pair_info,err := augur_srv.storage.Get_uniswap_pair_info(pair_aid)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
	//amount_to_trade := "100";
	//slippages,err := produce_uniswap_slippages(&pair_info,amount_to_trade)

	slippages := augur_srv.storage.Get_uniswap_latest_slippages(pair_aid)
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
func a1_wrapped_shtoken_balances(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,eoa_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	shtoken_balances := augur_srv.storage.Get_wrapped_shtoken_balances(eoa_aid)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK,gin.H{
		"WrappedShareTokenBalances" : shtoken_balances,
		"status": status,
		"error": err_str,
	})
}
func a1_user_wrapped_token_transfers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,user_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	p_wrapper:= c.Param("wrapper")
	_,wrapper_aid,success := json_validate_and_lookup_address_or_aid(c,&p_wrapper)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	wrapper_info,_ := augur_srv.storage.Get_wrapped_token_info(wrapper_aid)
	market_info,_ := augur_srv.storage.Get_market_info(wrapper_info.MktAddr,wrapper_info.OutcomeIdx,true)
	total_rows,transfers:= augur_srv.storage.Get_user_wrapped_shtoken_transfers(user_aid,wrapper_aid,offset,limit)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"MarketInfo" : market_info,
			"TokenInfo" : wrapper_info,
			"Offset" : offset,
			"Limit" : limit,
			"Transfers" : transfers,
			"TotalRows" : total_rows,
			"status": status,
			"error": err_str,
	})
}
func a1_ens_reverse_lookup(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_address:= c.Param("address")
/*	address_str,aid,success := json_validate_and_lookup_address_or_aid(c,&p_address)
	if !success {
		return
	}*/
	addr := common.HexToAddress(p_address)
	name, err := ens.ReverseResolve(rpcclient, addr)
	Info.Printf("reverse lookup of %v, name=%v\n",addr.String(),name)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"Addr" : p_address,
			"Name" : name,
			"status": status,
			"error": err_str,
	})
}
func a1_whats_new_augur(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var err error
	p_code := c.Param("code")
	var code int = 0
	if len(p_code) > 0 {
		code , err = strconv.Atoi(p_code)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error": err.Error(),
			})
			return
		}
	}
	block_num_from,block_num_to,err := augur_srv.storage.Get_block_range_for_whats_new(WhatsNewAugurCode(code))
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
	Info.Printf("from_block=%v, to_block=%v\n",block_num_from,block_num_to)
	block_info,err := augur_srv.storage.Get_block_info(int64(block_num_from),int64(block_num_to))
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"BlockInfo" : block_info,
		"status": status,
		"error": err_str,
	})
}
func a1_user_uniswap_swaps(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,user_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	user_info,err := augur_srv.storage.Get_user_info(user_aid)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": fmt.Sprintf("Error getting UserInfo: %v",err.Error()),
		})
	}
	swaps,total_recs := augur_srv.storage.Get_user_uniswap_swaps(user_aid,offset,limit)
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
func a1_user_balancer_swaps(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,user_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	user_info,err := augur_srv.storage.Get_user_info(user_aid)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": fmt.Sprintf("Error getting UserInfo: %v",err.Error()),
		})
	}
	swaps,total_rows := augur_srv.storage.Get_user_balancer_swaps(user_aid,offset,limit)
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
func a1_user_ens_names(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,user_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}

	user_info,err := augur_srv.storage.Get_user_info(user_aid)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": fmt.Sprintf("Error getting UserInfo: %v",err.Error()),
		})
	}
	ens_names,total_rows := augur_srv.storage.Get_user_ens_names(user_aid,offset,limit)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"status": status,
			"error" : err_str,
			"UserInfo" : user_info,
			"ENS_Names" : ens_names,
			"TotalRows" : total_rows,
	})
}
func a1_node_text_key_value_pairs(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	node := c.Param("node")
	fqdn,key_value_pairs:= augur_srv.storage.Get_node_text_key_values(node)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"error" : err_str,
		"Node" : node,
		"FullName" : fqdn,
		"KeyValuePairs" : key_value_pairs,
	})
}
func a1_augur_foundry_contracts(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	wrappers:= augur_srv.storage.Get_augur_foundry_wrapper_list()
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"error" : err_str,
		"ERC20MarketOutcomeWrappers" : wrappers,
	})
}
func a1_transaction_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_hash := c.Param("hash")
	if len(p_hash)==66 {
		p_hash = p_hash[2:]
	}
	hash_bytes,err := hex.DecodeString(p_hash)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": fmt.Sprintf("Error decoding hash hex: %v",err.Error()),
		})
		return
	}

	hash := common.BytesToHash(hash_bytes)
	hash_str := hash.String()
	tx_info,err := augur_srv.storage.Get_transaction(hash_str)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": fmt.Sprintf("Error in DB query: %v",err.Error()),
		})
		return
	}
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"error" : err_str,
		"TransactionInfo" : tx_info,
	})
}
func a1_block_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var block_num int64
	var err error
	p_block_num:= c.Param("block_num")
	if len(p_block_num) > 0 {
		block_num, err = strconv.ParseInt(p_block_num,10,64)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("Bad integer for block_num parameter: %v",err),
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("block_num parameter wasn't provided"),
		})
		return
	}
	block_info,err := augur_srv.storage.Get_block_info(block_num,block_num)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("block_num parameter wasn't provided"),
		})
		return
	}
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"BlockInfo" : block_info,
			"status": status,
			"error": err_str,
	})

}
func a1_reporting_table(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_market := c.Param("market")
	market_addr,market_aid,success := json_validate_and_lookup_address_or_aid(c,&p_market)
	if !success {
		return
	}
	market_info,err := augur_srv.storage.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,true,&market_addr)
		return
	}
	reporting_status,err := augur_srv.storage.Get_reporting_table(market_aid)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": err.Error(),
		})
		return
	}

	round_table,num_outcomes,outcomes,scalar_vals := augur_srv.storage.Get_round_table(market_aid)
	initial_report_redemption := augur_srv.storage.Get_initial_report_redeemed_record(market_aid)
	redeemed_participants := augur_srv.storage.Get_redeemed_participants(market_aid)

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"error" : err_str,
		"MarketInfo" : market_info,
		"ReportingTable" : reporting_status,
		"RoundTable" : round_table,
		"NumOutcomes" : num_outcomes,
		"Outcomes" : outcomes,
		"ScalarValues" : scalar_vals,
		"RedeemIniReporter" : initial_report_redemption,
		"RedeemedParticipants" : redeemed_participants,
	})
}
func a1_user_rep_profit_loss(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_user:= c.Param("user")
	_,user_aid,success := json_validate_and_lookup_address_or_aid(c,&p_user)
	if !success {
		return
	}

	user_info,err := augur_srv.storage.Get_user_info(user_aid)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error": fmt.Sprintf("Error getting UserInfo: %v",err.Error()),
		})
	}
	rep_profits := augur_srv.storage.Get_user_report_profits(user_aid)
	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
			"status": status,
			"error" : err_str,
			"UserInfo" : user_info,
			"RepProfits" : rep_profits,
			"RepLosses" : 0,
	})
}
func a1_noshow_bond_prices(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	bond_prices := augur_srv.storage.Get_noshow_bond_price_history()

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"error" : err_str,
		"NoShowBondPrices" : bond_prices,
	})
}
func a1_validity_bond_prices(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	bond_prices := augur_srv.storage.Get_validity_bond_price_history()

	var status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"error" : err_str,
		"ValidityBondPrices" : bond_prices,
	})
}
