package main

import (

	"github.com/gin-gonic/gin"
)
func set_routing_html(r *gin.Engine) {

	//r.GET("/black/markets.html",markets)
	r.GET("/balancerv2/poolinfo.html",bal_v2_poolinfo)
	r.GET("/balancerv2/tokbalhistory/:pool_aid/:tok_aid",bal_v2_pool_token_history)

}
func set_routing_api(r *gin.Engine) {

	//r.GET("/api/active_market_ids",a1_active_market_ids)
}
