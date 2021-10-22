package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func rwalk_current_offers(c *gin.Context) {

	p_order_by := c.Param("order_by")
	var order_by int64
	if len(p_order_by) > 0 {
		var success bool
		order_by,success = parse_int_from_remote_or_error(c,HTTP,&p_order_by)
		if !success {
			return
		}
	} else {
		respond_error(c,"'order_by' parameter is not set")
		return
	}
	offers := augur_srv.db_arbitrum.Get_active_offers(int(order_by))

	c.HTML(http.StatusOK, "current_offers.html", gin.H{
		"Offers" : offers,
	})
}
func rwalk_token_list_seq(c *gin.Context) {

	tokens:= augur_srv.db_arbitrum.Get_minted_tokens_sequentially(0,10000000000)

	c.HTML(http.StatusOK, "tokens_minted.html", gin.H{
		"MintedTokens" : tokens,
	})
}
func rwalk_token_list_period(c *gin.Context) {

	success,ini,fin := parse_timeframe_ini_fin(c,HTTP)
	if !success {
		return
	}
	tokens := augur_srv.db_arbitrum.Get_minted_tokens_by_period(ini,fin)

	c.HTML(http.StatusOK, "tokens_minted.html", gin.H{
		"MintedTokens" : tokens,
		"InitTs": ini,
		"FinTs":fin,
	})
}


