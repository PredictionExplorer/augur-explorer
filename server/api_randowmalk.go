/// API v1
package main
import (
	//"fmt"
	//"strconv"

	"net/http"
	"github.com/gin-gonic/gin"

)
func api_rwalk_current_offers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_order_by := c.Param("order_by")
	var order_by int64
	if len(p_order_by) > 0 {
		var success bool
		order_by,success = parse_int_from_remote_or_error(c,JSON,&p_order_by)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'order_by' parameter is not set")
		return
	}

	offers := augur_srv.db_arbitrum.Get_active_offers(int(order_by))

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offers" : offers,
	})
}
func api_rwalk_token_list_seq(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	tokens := augur_srv.db_arbitrum.Get_minted_tokens_sequentially(0,10000000000)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MintedTokens" : tokens ,
	})
}
func api_rwalk_token_list_period(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	success,ini,fin := parse_timeframe_ini_fin(c,JSON)
	if !success {
		return
	}
	tokens := augur_srv.db_arbitrum.Get_minted_tokens_by_period(ini,fin)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MintedTokens" : tokens,
		"InitTs": ini,
		"FinTs":fin,
	})
}

