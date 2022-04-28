package main

import (

	"github.com/gin-gonic/gin"
)
func set_routing_html(r *gin.Engine) {

	//r.GET("/black/markets.html",markets)
	r.GET("/balancerv2/poolinfo.html",bal_v2_poolinfo)
	r.GET("/balancerv2/tokbalhistory/:pool_aid/:tok_aid",bal_v2_pool_token_history)
	r.GET("/balancerv2/pool_profits/:pool_aid/:ini_ts/:fin_ts",bal_v2_pool_profits_in_swaps)
	r.GET("/balancerv2/top_pools/:tf_code/:ini_ts/:fin_ts",bal_v2_top_pools)
}
func set_routing_api(r *gin.Engine) {

	//r.GET("/api/active_market_ids",a1_active_market_ids)
}
