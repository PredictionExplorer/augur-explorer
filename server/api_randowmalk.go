/// API v1
package main
import (
	"fmt"
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

	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		rwalk_aid = 0
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		market_aid = 0
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

	offers := augur_srv.db_arbitrum.Get_active_offers(rwalk_aid,market_aid,int(order_by))

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offers" : offers,
		"RWalkAid": rwalk_aid,
		"MarketAid": market_aid,
	})
}
func api_rwalk_floor_price(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		rwalk_aid = -1
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		market_aid = -1
	}

	floor_price,err := augur_srv.db_arbitrum.Get_floor_price(rwalk_aid,market_aid)
	var db_err string
	if err != nil { db_err = err.Error() }

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"FloorPrice" : floor_price,
		"DBError": db_err,
		"MarketAddr":p_market_addr,
		"RWalkAddr":p_rwalk_addr,
		"RWalkAid": rwalk_aid,
		"MarketAid": market_aid,
	})
}
func api_rwalk_token_list_seq(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error_json(c,"NTF address wasn't found in the 'address' table")
		return
	}

	tokens := augur_srv.db_arbitrum.Get_minted_tokens_sequentially(rwalk_aid,0,10000000000)

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

	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error_json(c,"NTF address wasn't found in the 'address' table")
		return
	}

	success,ini,fin := parse_timeframe_ini_fin(c,JSON)
	if !success {
		return
	}
	tokens := augur_srv.db_arbitrum.Get_minted_tokens_by_period(rwalk_aid,ini,fin)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MintedTokens" : tokens,
		"InitTs": ini,
		"FinTs":fin,
		"RWalkAid":rwalk_aid,
	})
}
func api_rwalk_sale_history(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		respond_error(c,"Market address doesn't exist in the database")
		return
	}

	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	sales := augur_srv.db_arbitrum.Get_trading_history(market_aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Sales" : sales,
		"MarketAid" : market_aid,
		"MarketAddr" : p_market_addr,
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
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error_json(c,"Lookup of 'rwalk_addr' failed, address doesn't exist")
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	history := augur_srv.db_arbitrum.Get_token_full_history(rwalk_aid,token_id,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"TokenId" : token_id,
		"TokenHistory" : history,
		"RWalkAddr" : p_rwalk_addr,
		"RWalkAid" : rwalk_aid,
	})
}
func api_rwalk_trading_volume_by_period(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		respond_error_json(c,"Market address doesn't exist in the database")
		return
	}
	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}

	vol_hist := augur_srv.db_arbitrum.Get_market_trading_volume_by_period(market_aid,init_ts,fin_ts,interval_secs)
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
func api_rwalk_token_stats(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error_json(c,"Lookup of NFT token address in the Db has failed")
		return
	}
	stats := augur_srv.db_arbitrum.Get_random_walk_stats(rwalk_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"TokenStats" : stats,
		"RWalkAid": rwalk_aid,
		"RWalkAddr" : p_rwalk_addr,
	})
}
func api_rwalk_market_stats(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		respond_error_json(c,"Lookup of Market address in the DB has failed")
		return
	}
	stats := augur_srv.db_arbitrum.Get_market_stats(market_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MarketStats" : stats,
		"MarketAid": market_aid,
		"MarketAddr" : p_market_addr,
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
func api_rwalk_trading_history_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid,success = parse_int_from_remote_or_error(c,HTTP,&p_user_aid)
		if !success {
			return
		}
	} else {
		respond_error(c,"'user_aid' parameter is not set")
		return
	}
	user_addr,err := augur_srv.db_arbitrum.Lookup_address(user_aid)
	if err != nil {
		respond_error_json(c,"Address lookup on user_aid failed")
		return
	}
	user_trading := augur_srv.db_arbitrum.Get_trading_history_by_user(user_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserTrading" : user_trading,
		"UserAid" : user_aid,
		"UserAddr" : user_addr,
	})
}
func api_rwalk_user_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error_json(c,"Lookup of NFT token failed")
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
	user_info,err := augur_srv.db_arbitrum.Get_rwalk_user_info(user_aid,rwalk_aid)
	if err != nil {
		respond_error_json(c,fmt.Sprintf("Statistics record for this user in token %v wasn't found",p_rwalk_addr))
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserInfo" : user_info,
		"UserAid" : user_aid,
		"UserAddr" : user_addr,
		"RWalkAddr" : p_rwalk_addr,
		"RWalkAid" : rwalk_aid,
	})
}
func api_rwalk_top5_traded_tokens(c *gin.Context) {

	if !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	top5toks := augur_srv.db_arbitrum.Get_top5_traded_tokens()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Top5TradedTokens" : top5toks,
	})
}
