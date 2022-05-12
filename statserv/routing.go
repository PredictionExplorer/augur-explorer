package main

import (

	"github.com/gin-gonic/gin"
)
func set_routing_html(r *gin.Engine) {

	//r.GET("/black/markets.html",markets)
	r.GET("/balancerv2/poolinfo.html",bal_v2_poolinfo)
	r.GET("/balancerv2/tokbalhistory/:pool_aid/:tok_aid",bal_v2_pool_token_history)
	r.GET("/balancerv2/pool_profits/:pool_aid/:ini_ts/:fin_ts",bal_v2_pool_profits_in_swaps)
	r.GET("/balancerv2/pool_fees_weekly/:pool_aid/:offset/:limit",bal_v2_pool_fees_weekly)
	r.GET("/balancerv2/pool_fees_daily/:pool_aid/:offset/:limit",bal_v2_pool_fees_daily)
	r.GET("/balancerv2/pool_fees_hourly/:pool_aid/:offset/:limit",bal_v2_pool_fees_hourly)
	r.GET("/balancerv2/top_pools/:tf_code/:ini_ts/:fin_ts",bal_v2_top_pools)
	r.GET("/balancerv2/balancer_index.html",bal_v2_pools_index_page)
	r.GET("/balancerv2/swap_history/:pool_aid/:offset/:limit",bal_v2_swap_history)
}
func set_routing_api(r *gin.Engine) {

	//r.GET("/api/active_market_ids",a1_active_market_ids)
}
