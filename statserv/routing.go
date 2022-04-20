package main

import (

	"github.com/gin-gonic/gin"
)
func set_routing_html(r *gin.Engine) {

	//r.GET("/black/markets.html",markets)
	r.GET("/balancer/v2_poolinfo.html",bal_v2_poolinfo)

}
func set_routing_api(r *gin.Engine) {

	//r.GET("/api/active_market_ids",a1_active_market_ids)
}
