package main
import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"

)
func rwalk_index_page(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	caddrs := augur_srv.db_arbitrum.Get_randomwalk_contract_addresses()
	top5tokens := augur_srv.db_arbitrum.Get_top5_traded_tokens()
	c.HTML(http.StatusOK, "rw_index.html", gin.H{
		"ContractAddresses":caddrs,
		"Top5Tokens":top5tokens,
	})
}
func rwalk_current_offers(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
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
	offers := augur_srv.db_arbitrum.Get_active_offers(rwalk_aid,market_aid,int(order_by))

	c.HTML(http.StatusOK, "rw_current_offers.html", gin.H{
		"Offers" : offers,
		"RWalkAid": rwalk_aid,
		"MarketAid": market_aid,
	})
}
func rwalk_floor_price(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
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
	c.HTML(http.StatusOK, "rw_floor_price.html", gin.H{
		"FloorPrice" : floor_price,
		"DBError": db_err,
		"MarketAddr":p_market_addr,
		"RWalkAddr":p_rwalk_addr,
		"RWalkAid": rwalk_aid,
		"MarketAid": market_aid,
	})
}
func rwalk_token_list_seq(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"NTF address wasn't found in the 'address' table")
		return
	}
	tokens:= augur_srv.db_arbitrum.Get_minted_tokens_sequentially(rwalk_aid,0,10000000000)

	fin_ts := int(time.Now().Unix())
	interval := int(2 * 24 * 60* 60)
	init_ts := fin_ts - interval
	c.HTML(http.StatusOK, "rw_tokens_minted.html", gin.H{
		"MintedTokens" : tokens,
		"RWalkAddr" : p_rwalk_addr,
		"RWalkAid" : rwalk_aid,
		"InitTs":init_ts,
		"FinTs":fin_ts,
		"Interval":interval,
	})
}
func rwalk_token_list_period(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"NTF address wasn't found in the 'address' table")
		return
	}
	success,ini,fin := parse_timeframe_ini_fin(c,HTTP)
	if !success {
		return
	}
	tokens := augur_srv.db_arbitrum.Get_minted_tokens_by_period(rwalk_aid,ini,fin)

	c.HTML(http.StatusOK, "rw_tokens_minted_period.html", gin.H{
		"MintedTokens" : tokens,
		"RWalkAddr" : p_rwalk_addr,
		"RWalkAid" : rwalk_aid,
		"InitTs": ini,
		"FinTs":fin,
		"InitDate" : time.Unix(int64(ini),0).String(),
		"FinDate":time.Unix(int64(fin),0).String(),
	})
}
func rwalk_trading_history(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
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

	c.HTML(http.StatusOK, "rw_trading_history.html", gin.H{
		"Sales" : sales,
	})
}
func rwalk_token_stats(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"Lookup of NFT token address in the Db has failed")
		return
	}
	stats := augur_srv.db_arbitrum.Get_random_walk_stats(rwalk_aid)

	c.HTML(http.StatusOK, "rw_token_stats.html", gin.H{
		"TokenStats" : stats,
		"RWalkAid": rwalk_aid,
		"RWalkAddr" : p_rwalk_addr,
	})
}
func rwalk_market_stats(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		respond_error(c,"Market address doesn't exist in the database")
		return
	}
	stats := augur_srv.db_arbitrum.Get_market_stats(market_aid)

	c.HTML(http.StatusOK, "rw_market_stats.html", gin.H{
		"MarketStats" : stats,
		"MarketAid": market_aid,
		"MarketAddr" : p_market_addr,
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
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"Lookup of 'rwalk_addr' failed, address doesn't exist")
	}
	offset := int(0) ; limit:= int(100000)
	/*success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}*/
	history := augur_srv.db_arbitrum.Get_token_full_history(rwalk_aid,token_id,offset,limit)
	token_info,err := augur_srv.db_arbitrum.Get_rwalk_token_info(rwalk_aid,token_id)
	if err != nil {
		fmt.Printf("Error getting token info for token_id=%v, rwalk_aid=%v : %v\n",token_id,rwalk_aid,err)
	}

	c.HTML(http.StatusOK, "rw_token_history.html", gin.H{
		"TokenId" : token_id,
		"TokenHistory" : history,
		"RWalkAddr" : p_rwalk_addr,
		"RWalkAid" : rwalk_aid,
		"TokenInfo" : token_info,
	})
}
func rwalk_trading_volume_by_period(c *gin.Context) {

	success,init_ts,fin_ts,interval_secs := parse_timeframe_params(c)
	if !success {
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		respond_error(c,"Market address doesn't exist in the database")
		return
	}

	vol_hist := augur_srv.db_arbitrum.Get_market_trading_volume_by_period(market_aid,init_ts,fin_ts,interval_secs)
	volume_data := build_js_randomwalk_volume_history(&vol_hist)
	c.HTML(http.StatusOK, "rw_volume_history.html", gin.H{
		"VolumeHistory" : vol_hist,
		"VolumeData" : volume_data,
		"InitTs" : init_ts,
		"FinTs" : fin_ts,
		"Interval" : interval_secs,
		"MarketAddr" : p_market_addr,
		"MarketAid" : market_aid,
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
func rwalk_tokens_by_user(c *gin.Context) {

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
		respond_error(c,"Address lookup on user_aid failed")
		return
	}
	user_tokens := augur_srv.db_arbitrum.Get_random_walk_tokens_by_user(user_aid)

	c.HTML(http.StatusOK, "rw_tokens_by_user.html", gin.H{
		"UserTokens" : user_tokens,
		"UserAid" : user_aid,
		"UserAddr" : user_addr,
	})
}
func rwalk_trading_history_by_user(c *gin.Context) {

	if !augur_srv.arbitrum_initialized() {
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
		respond_error(c,fmt.Sprintf("Address lookup on user_aid %v failed: %v",user_aid,err))
		return
	}
	user_trading := augur_srv.db_arbitrum.Get_trading_history_by_user(user_aid)

	c.HTML(http.StatusOK, "rw_trading_by_user.html", gin.H{
		"UserTrading" : user_trading,
		"UserAid" : user_aid,
		"UserAddr" : user_addr,
	})
}
func rwalk_user_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid,err := augur_srv.db_arbitrum.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		respond_error(c,"Lookup of NFT token failed")
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
		respond_error(c,"Address lookup on user_aid failed")
		return
	}
	user_info,err := augur_srv.db_arbitrum.Get_rwalk_user_info(user_aid,rwalk_aid)
	if err != nil {
		respond_error(c,fmt.Sprintf("Statistics record for this user in token %v wasn't found",p_rwalk_addr))
		return
	}

	c.HTML(http.StatusOK, "rw_user_info.html", gin.H{
		"UserInfo" : user_info,
		"UserAid" : user_aid,
		"UserAddr" : user_addr,
	})
}
