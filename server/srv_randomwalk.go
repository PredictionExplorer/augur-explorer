package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func rwalk_index_page(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	c.HTML(http.StatusOK, "rw_index.html", gin.H{
	})
}
func rwalk_current_offers(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
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

	c.HTML(http.StatusOK, "rw_current_offers.html", gin.H{
		"Offers" : offers,
	})
}
func rwalk_token_list_seq(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	tokens:= augur_srv.db_arbitrum.Get_minted_tokens_sequentially(0,10000000000)

	c.HTML(http.StatusOK, "rw_tokens_minted.html", gin.H{
		"MintedTokens" : tokens,
	})
}
func rwalk_token_list_period(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	success,ini,fin := parse_timeframe_ini_fin(c,HTTP)
	if !success {
		return
	}
	tokens := augur_srv.db_arbitrum.Get_minted_tokens_by_period(ini,fin)

	c.HTML(http.StatusOK, "rw_tokens_minted.html", gin.H{
		"MintedTokens" : tokens,
		"InitTs": ini,
		"FinTs":fin,
	})
}
func rwalk_sale_history(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	sales := augur_srv.db_arbitrum.Get_sale_history(offset,limit)

	c.HTML(http.StatusOK, "rw_sale_history.html", gin.H{
		"Sales" : sales,
	})
}
func rwalk_global_stats(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	stats := augur_srv.db_arbitrum.Get_global_stats()

	c.HTML(http.StatusOK, "rw_global_stats.html", gin.H{
		"GlobalStats" : stats,
	})
}
func rwalk_token_history(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id,success = parse_int_from_remote_or_error(c,HTTP,&p_token_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'token_id' parameter is not set")
		return
	}
	offset := int(0) ; limit:= int(100000)
	/*success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}*/
	history := augur_srv.db_arbitrum.Get_token_full_history(token_id,offset,limit)

	c.HTML(http.StatusOK, "rw_token_history.html", gin.H{
		"TokenId" : token_id,
		"TokenHistory" : history,
	})
}
func rwalk_trading_volume_by_period(c *gin.Context) {

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	vol_hist := augur_srv.db_arbitrum.Get_randomwalk_trading_volume_by_period(init_ts,fin_ts,interval_secs)
	volume_data := build_js_randomwalk_volume_history(&vol_hist)
	c.HTML(http.StatusOK, "rw_volume_history.html", gin.H{
		"VolumeHistory" : vol_hist,
		"VolumeData" : volume_data,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
	})
}
func rwalk_token_name_history(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id,success = parse_int_from_remote_or_error(c,HTTP,&p_token_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'token_id' parameter is not set")
		return
	}
	name_changes := augur_srv.db_arbitrum.Get_name_changes_for_token(token_id)

	c.HTML(http.StatusOK, "rw_token_names.html", gin.H{
		"TokenNameChanges" : name_changes,
	})
}
