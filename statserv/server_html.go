package main
import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"

	//. "github.com/PredictionExplorer/augur-explorer/dbs/balancerv2"
)
func main_page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Index Page",
	})
}
func bal_v2_pools_index_page(c *gin.Context) {

	//ts := time.Now().Unix() TEMPORARILY DISABLED , to activate after population of the DB ends
	ts := storagew.Get_first_last_swap_timestamp_all_pools(true)
	fmt.Printf("ts=%v\n",ts)
	ts = ts - 60*60*24*7	// discount a week, for demo purposes (to be removed)
	fmt.Printf("discounted ts = %v\n",ts)
	hourly_ts_ini := ts/(60*60)
	hourly_ts_ini = hourly_ts_ini * 60*60
	daily_ts_ini := ts/(60*60*24)
	daily_ts_ini = daily_ts_ini * 60*60*24
	weekly_ts_ini := ts/(60*60*24*7)
	weekly_ts_ini = weekly_ts_ini * 60*60*24*7
	hourly_ts_fin := hourly_ts_ini + 60*60
	daily_ts_fin := daily_ts_ini + 60*60*24
	weekly_ts_fin := weekly_ts_ini + 60*60*24*7


	c.HTML(http.StatusOK, "balancer_index.html", gin.H{
			"title": "Balancer Pools Home Page",
			"HourlyTsIni" : hourly_ts_ini,
			"HourlyTsFin" : hourly_ts_fin,
			"DailyTsIni" : daily_ts_ini,
			"DailyTsFin" : daily_ts_fin,
			"WeeklyTsIni" : weekly_ts_ini,
			"WeeklyTsFin" : weekly_ts_fin,
	})
}
func bal_v2_poolinfo(c *gin.Context) {

	pool_id := c.Query("pool_id")
	pool_aid,err := storagew.Lookup_pool_address_id(pool_id)
	if err != nil {
		respond_error_html(c,"No outcome provided")
		return
	}
	pool_info := storagew.Get_pool_info(pool_id)
	pool_tokens := storagew.Get_pool_registered_tokens(ethprice_storage,pool_aid,weth_aid)
	c.HTML(http.StatusOK, "poolinfo.html", gin.H{
			"title": "Balancer v2 Pool Info",
			"PoolInfo" : pool_info,
			"Tokens" : pool_tokens,
	})
}
func bal_v2_pool_token_history(c *gin.Context) {

	pool_aid,success := parse_integer_param_or_error(c,"pool_aid",PARAM_FORCED,FMT_HTML)
	if !success { return }
	pool_id,err := storagew.Lookup_pool_id_by_addr_id(pool_aid)
	if err != nil {
		respond_error_html(c,"Can't find pool id for provided pool address id")
		return
	}
	tok_aid,success := parse_integer_param_or_error(c,"tok_aid",PARAM_FORCED,FMT_HTML)
	if !success { return }
	tok_addr,err := storagew.S.Layer1_lookup_address(tok_aid)
	if err!=nil {
		fmt.Printf("err=%v\n",err)
		respond_error_html(c,"Token address ID not found")
		return
	}
	balances:=storagew.Get_pool_token_balance_history(pool_aid,tok_aid)
	c.HTML(http.StatusOK, "tokbalhistory.html", gin.H{
			"title": "Balancer v2 Pool Info",
			"PoolAid" : pool_aid,
			"PoolId" : pool_id,
			"Balances" : balances,
			"TokenAid" : tok_aid,
			"TokenAddr" : tok_addr,
	})
}
func bal_v2_pool_profits_in_swaps(c *gin.Context) {

	pool_aid,success := parse_integer_param_or_error(c,"pool_aid",PARAM_FORCED,FMT_HTML)
	if !success { return }

	ini_ts,success := parse_integer_param_or_error(c,"ini_ts",PARAM_FORCED,FMT_HTML)
	if !success { return }
	fin_ts,success := parse_integer_param_or_error(c,"fin_ts",PARAM_FORCED,FMT_HTML)
	if !success { return }

	ts := time.Unix(ini_ts,0)
	start_date := ts.String()
	ts = time.Unix(fin_ts,0)
	end_date := ts.String()
	profit := storagew.Get_pool_swap_fee_profits(pool_aid,ini_ts,fin_ts)
	c.HTML(http.StatusOK, "pool_swap_profits.html", gin.H{
		"title": "Balancer v2 Pool Info",
		"PoolAid" : pool_aid,
		"TotalProfitInSwapFees" : profit,
		"IniTs" : ini_ts,
		"FinTs" : fin_ts,
		"StartDate" : start_date,
		"EndDate" : end_date,
	})
}
func bal_v2_top_pools(c *gin.Context) {

	ini_ts,success := parse_integer_param_or_error(c,"ini_ts",PARAM_OPTIONAL,FMT_HTML)
	if !success { return }
	fin_ts,success := parse_integer_param_or_error(c,"fin_ts",PARAM_OPTIONAL,FMT_HTML)
	if !success { return }
	tf_code,success := parse_integer_param_or_error(c,"tf_code",PARAM_FORCED,FMT_HTML)
	if !success { return }

	if ini_ts == 0 {
		ini_ts = storagew.Get_first_last_swap_timestamp_all_pools(false)
	}
	if fin_ts == 0 {
		fin_ts = storagew.Get_first_last_swap_timestamp_all_pools(true)
	}
	ts := time.Unix(ini_ts,0)
	start_date := ts.String()
	ts = time.Unix(fin_ts,0)
	end_date := ts.String()

	top_pools := storagew.Get_top_profitable_pools_v2(tf_code,ini_ts,fin_ts)

	c.HTML(http.StatusOK, "top_pools.html", gin.H{
		"title": "Balancer v2 Top most profitable Pools",
		"TfCode" :tf_code,
		"IniTs" : ini_ts,
		"FinTs" : fin_ts,
		"StartDate" : start_date,
		"EndDate" : end_date,
		"Pools" : top_pools,
	})
}
func bal_v2_swap_history(c *gin.Context) {

	pool_aid,success := parse_integer_param_or_error(c,"pool_aid",PARAM_FORCED,FMT_HTML)
	if !success { return }
	offset,success := parse_integer_param_or_error(c,"offset",PARAM_OPTIONAL,FMT_HTML)
	if !success { return }
	limit,success := parse_integer_param_or_error(c,"limit",PARAM_OPTIONAL,FMT_HTML)
	if !success { return }

	if limit == 0 { limit = 1000 }
	swaps := storagew.Get_pool_swap_history_backwards(pool_aid,offset,limit)
	c.HTML(http.StatusOK, "swap_history.html", gin.H{
		"title": "Balancer v2 Swap History",
		"Swaps" : swaps,
	})
}
func bal_v2_pool_fees_weekly(c *gin.Context) {

	tf_code := int64(2)
	bal_v2_pool_fee_returns_by_timeframe(c,tf_code)
}
func bal_v2_pool_fees_daily(c *gin.Context) {

	tf_code := int64(1)
	bal_v2_pool_fee_returns_by_timeframe(c,tf_code)
}
func bal_v2_pool_fees_hourly(c *gin.Context) {

	tf_code := int64(0)
	bal_v2_pool_fee_returns_by_timeframe(c,tf_code)
}
func bal_v2_pool_fee_returns_by_timeframe(c *gin.Context,tf_code int64) {
	pool_aid,success := parse_integer_param_or_error(c,"pool_aid",PARAM_FORCED,FMT_HTML)
	if !success { return }

	offset,success := parse_integer_param_or_error(c,"offset",PARAM_OPTIONAL,FMT_HTML)
	if !success { return }
	limit,success := parse_integer_param_or_error(c,"limit",PARAM_OPTIONAL,FMT_HTML)
	if !success { return }

	if limit == 0 { limit = 100 }

	fee_returns := storagew.Get_pool_swap_fee_returns_by_timeframe_code(pool_aid,tf_code,offset,limit)

	var chart_js_data ChartJSData
	chart_js_data = build_js_fee_returns(tf_code,&fee_returns)

	c.HTML(http.StatusOK, "fee_returns.html", gin.H{
		"title": "Balancer v2 Fee Returns",
		"PoolAid" : pool_aid,
		"FeeReturns" : fee_returns,
		"JSParams": chart_js_data,
	})
}
func bal_v2_pool_liquidity_providers_distrib(c *gin.Context) {

	pool_aid,success := parse_integer_param_or_error(c,"pool_aid",PARAM_FORCED,FMT_HTML)
	if !success { return }

	liquidity := storagew.Get_pool_liquidity_provider_distrib(pool_aid)

	var chart_js_data ChartJSData
	chart_js_data = build_js_liquidity_distrib(&liquidity)

	c.HTML(http.StatusOK, "pool_fund_distribution.html", gin.H{
		"title": "Balancer v2 Fee Returns",
		"PoolAid" : pool_aid,
		"Liquidity" : liquidity,
		"JSParams": chart_js_data,
	})
}
