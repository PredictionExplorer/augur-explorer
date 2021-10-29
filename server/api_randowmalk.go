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

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

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

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

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

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

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
func api_rwalk_sale_history(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	sales := augur_srv.db_arbitrum.Get_sale_history(offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Sales" : sales,
	})
}
func api_rwalk_token_history(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id,success = parse_int_from_remote_or_error(c,JSON,&p_token_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'token_id' parameter is not set")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	history := augur_srv.db_arbitrum.Get_token_full_history(token_id,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"TokenId" : token_id,
		"TokenHistory" : history,
	})
}
func api_rwalk_trading_volume_by_period(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	vol_hist := augur_srv.db_arbitrum.Get_randomwalk_trading_volume_by_period(init_ts,fin_ts,interval_secs)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"VolumeHistory" : vol_hist,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
	})
}
func api_rwalk_token_name_history(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id,success = parse_int_from_remote_or_error(c,JSON,&p_token_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'token_id' parameter is not set")
		return
	}
	name_changes := augur_srv.db_arbitrum.Get_name_changes_for_token(token_id)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"TokenNameChanges" : name_changes,
	})
}
func api_rwalk_global_stats(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	stats := augur_srv.db_arbitrum.Get_global_stats()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"GlobalStats" : stats,
	})
}
func api_rwalk_tokens_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,JSON,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'user_aid' parameter is not set")
		return
	}
	user_addr,err := augur_srv.db_arbitrum.Lookup_address(user_aid)
	if err != nil {
		respond_error_json(c,"Address lookup on user_aid failed")
		return
	}
	user_tokens := augur_srv.db_arbitrum.Get_random_walk_tokens_by_user(user_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserTokens" : user_tokens,
		"UserAid" : user_aid,
		"UserAddr" : user_addr,
	})
}
