// Package randomwalk - API v1 handlers for RandomWalk NFT
package randomwalk

import (
	"fmt"
	"strconv"
	"time"

	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

const (
	JSON = true
	HTTP = false
)

func apiRwalkCurrentOffers(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		rwalk_aid = 0
	}
	p_market_addr := c.Param("market_addr")
	market_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		market_aid = 0
	}

	p_order_by := c.Param("order_by")
	var order_by int64
	if len(p_order_by) > 0 {
		var success bool
		order_by, success = common.ParseIntFromRemoteOrError(c, JSON, &p_order_by)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'order_by' parameter is not set")
		return
	}

	offers := rw_storagew.Get_active_offers(rwalk_aid, market_aid, int(order_by))

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":    req_status,
		"error":     err_str,
		"Offers":    offers,
		"RWalkAid":  rwalk_aid,
		"MarketAid": market_aid,
	})
}

func apiRwalkFloorPrice(c *gin.Context) {

	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		rwalk_aid = -1
	}
	p_market_addr := c.Param("market_addr")
	market_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		market_aid = -1
	}

	_, floor_price, _, _, err := rw_storagew.Get_floor_price(rwalk_aid, market_aid)
	var db_err string
	if err != nil {
		db_err = err.Error()
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":     req_status,
		"error":      err_str,
		"FloorPrice": floor_price,
		"DBError":    db_err,
		"MarketAddr": p_market_addr,
		"RWalkAddr":  p_rwalk_addr,
		"RWalkAid":   rwalk_aid,
		"MarketAid":  market_aid,
	})
}

func apiRwalkTokenListSeq(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "NTF address wasn't found in the 'address' table")
		return
	}

	tokens := rw_storagew.Get_minted_tokens_sequentially(rwalk_aid, 0, 10000000000)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":       req_status,
		"error":        err_str,
		"MintedTokens": tokens,
	})
}

func apiRwalkTokenListPeriod(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "NTF address wasn't found in the 'address' table")
		return
	}

	success, ini, fin := common.ParseTimeframeIniFin(c, JSON)
	if !success {
		return
	}
	tokens := rw_storagew.Get_minted_tokens_by_period(rwalk_aid, ini, fin)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":       req_status,
		"error":        err_str,
		"MintedTokens": tokens,
		"InitTs":       ini,
		"FinTs":        fin,
		"RWalkAid":     rwalk_aid,
	})
}

func apiRwalkTradingHistory(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	var market_aid int64 = 0
	if p_market_addr != "0x0000000000000000000000000000000000000000" {
		var err error
		market_aid, err = rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
		if err != nil {
			common.RespondErrorJSON(c, "Market address doesn't exist in the database")
			return
		}
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	sales := rw_storagew.Get_trading_history(market_aid, offset, limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":     req_status,
		"error":      err_str,
		"Sales":      sales,
		"MarketAid":  market_aid,
		"MarketAddr": p_market_addr,
	})
}

func apiRwalkSaleHistory(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	var market_aid int64 = 0
	if p_market_addr != "0x0000000000000000000000000000000000000000" {
		var err error
		market_aid, err = rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
		if err != nil {
			common.RespondError(c, "Market address doesn't exist in the database")
			return
		}
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	sales := rw_storagew.Get_sale_history(market_aid, offset, limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":          req_status,
		"error":           err_str,
		"Trading":         sales,
		"MarketPlaceAddr": p_market_addr,
		"MarketPlaceAid":  market_aid,
	})
}

func apiRwalkTokenHistory(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'token_id' parameter is not set")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of 'rwalk_addr' failed, address doesn't exist")
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	history := rw_storagew.Get_token_full_history(rwalk_aid, token_id, offset, limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":       req_status,
		"error":        err_str,
		"TokenId":      token_id,
		"TokenHistory": history,
		"RWalkAddr":    p_rwalk_addr,
		"RWalkAid":     rwalk_aid,
	})
}

func apiRwalkTradingVolumeByPeriod(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Market address doesn't exist in the database")
		return
	}
	success, init_ts, fin_ts, interval_secs := common.ParseTimeframeParams(c)
	if !success {
		return
	}

	vol_hist := rw_storagew.Get_market_trading_volume_by_period(market_aid, init_ts, fin_ts, interval_secs)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":        req_status,
		"error":         err_str,
		"VolumeHistory": vol_hist,
		"InitTs":        init_ts,
		"FinTs":         fin_ts,
		"Interval":      interval_secs,
	})
}

func apiRwalkTokenNameHistory(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondError(c, "'token_id' parameter is not set")
		return
	}
	name_changes := rw_storagew.Get_name_changes_for_token(token_id)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"TokenNameChanges": name_changes,
	})
}

func apiRwalkTokenStats(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of NFT token address in the Db has failed")
		return
	}
	stats := rw_storagew.Get_random_walk_stats(rwalk_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":     req_status,
		"error":      err_str,
		"TokenStats": stats,
		"RWalkAid":   rwalk_aid,
		"RWalkAddr":  p_rwalk_addr,
	})
}

func apiRwalkMarketStats(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of Market address in the DB has failed")
		return
	}
	stats := rw_storagew.Get_market_stats(market_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":      req_status,
		"error":       err_str,
		"MarketStats": stats,
		"MarketAid":   market_aid,
		"MarketAddr":  p_market_addr,
	})
}

func apiRwalkTokensByUser(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}

	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var err error
		user_aid, err = strconv.ParseInt(p_user_aid, 10, 64)
		if err != nil {
			if (len(p_user_aid) != 40) && (len(p_user_aid) != 42) {
				common.RespondErrorJSON(c, "Can't resolve user identifier to valid address ID or address hex")
				return
			} else {
				user_aid, err = rw_storagew.S.Nonfatal_lookup_address_id(p_user_aid)
				if err != nil {
					common.RespondErrorJSON(c, "Cant find provided user")
					return
				}
			}
		}
	} else {
		common.RespondErrorJSON(c, "'user_aid' parameter is not set")
		return
	}

	user_addr, err := rw_storagew.S.Lookup_address(user_aid)
	if err != nil {
		common.RespondErrorJSON(c, "Address lookup on user_aid failed")
		return
	}
	user_tokens := rw_storagew.Get_random_walk_tokens_by_user(user_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":     req_status,
		"error":      err_str,
		"UserTokens": user_tokens,
		"UserAid":    user_aid,
		"UserAddr":   user_addr,
	})
}

func apiRwalkTradingHistoryByUser(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid, success = common.ParseIntFromRemoteOrError(c, HTTP, &p_user_aid)
		if !success {
			return
		}
	} else {
		common.RespondError(c, "'user_aid' parameter is not set")
		return
	}
	user_addr, err := rw_storagew.S.Lookup_address(user_aid)
	if err != nil {
		common.RespondErrorJSON(c, "Address lookup on user_aid failed")
		return
	}
	user_trading := rw_storagew.Get_trading_history_by_user(user_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":      req_status,
		"error":       err_str,
		"UserTrading": user_trading,
		"UserAid":     user_aid,
		"UserAddr":    user_addr,
	})
}

func apiRwalkUserInfo(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of NFT token failed")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid, success = common.ParseIntFromRemoteOrError(c, JSON, &p_user_aid)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'user_aid' parameter is not set")
		return
	}
	user_addr, err := rw_storagew.S.Lookup_address(user_aid)
	if err != nil {
		common.RespondErrorJSON(c, "Address lookup on user_aid failed")
		return
	}
	user_info, dberr := rw_storagew.Get_rwalk_user_info(user_aid, rwalk_aid)

	var dberr_str string
	if dberr != nil {
		dberr_str = dberr.Error()
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":    req_status,
		"error":     err_str,
		"UserInfo":  user_info,
		"UserAid":   user_aid,
		"UserAddr":  user_addr,
		"RWalkAddr": p_rwalk_addr,
		"RWalkAid":  rwalk_aid,
		"DBError":   dberr_str,
	})
}

func apiRwalkTop5TradedTokens(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	top5toks := rw_storagew.Get_top5_traded_tokens()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":          req_status,
		"error":           err_str,
		"Top5TradedTokens": top5toks,
	})
}

func apiRwalkMintIntervals(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of NFT token failed")
		return
	}
	mint_intervals := rw_storagew.Get_rwalk_mint_intervals(rwalk_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":        req_status,
		"error":         err_str,
		"MintIntervals": mint_intervals,
		"RWalkAid":      rwalk_aid,
		"RWalkAddr":     p_rwalk_addr,
	})
}

func apiRwalkWithdrawalChart(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondError(c, "Lookup of NFT token failed")
		return
	}
	withdrawal_entries := rw_storagew.Get_rwalk_withdrawal_chart(rwalk_aid)
	withdrawal_data := common.BuildJSRandomwalkWithdrawalChart(&withdrawal_entries)
	rwalk_stats := rw_storagew.Get_random_walk_stats(rwalk_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":             req_status,
		"error":              err_str,
		"WithdrawalEntries":  withdrawal_entries,
		"WithdrawalData":     withdrawal_data,
		"ContractStatistics": rwalk_stats,
		"RWalkAid":           rwalk_aid,
		"RWalkAddr":          p_rwalk_addr,
	})
}

func apiRwalkFloorPriceOverTime(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of NFT token failed")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Market address doesn't exist in the database")
		return
	}
	success, ini, fin, interval := common.ParseTimeframeParams(c)
	if !success {
		return
	}
	if ini == 0 {
		ini = 1636676049
	}
	if fin == 0 {
		fin = int(time.Now().Unix())
	}
	if interval == 0 || interval == 2147483647 {
		interval = 24 * 60 * 60
	}
	price_entries := rw_storagew.Get_rwalk_floor_price_for_periods(rwalk_aid, market_aid, ini, fin, interval)
	price_data := common.BuildJSFloorPriceData(&price_entries)
	rwalk_stats := rw_storagew.Get_random_walk_stats(rwalk_aid)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":             req_status,
		"error":              err_str,
		"PriceEntries":       price_entries,
		"PriceData":          price_data,
		"ContractStatistics": rwalk_stats,
		"RWalkAid":           rwalk_aid,
		"RWalkAddr":          p_rwalk_addr,
		"MarketAid":          market_aid,
		"MarketAddr":         p_market_addr,
		"InitTs":             ini,
		"FinTs":              fin,
		"Interval":           interval,
	})
}

func apiRwalkTokenInfo(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of NFT token address in the Db has failed")
		return
	}
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, HTTP, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'token_id' parameter is not set")
		return
	}
	token_info, err := rw_storagew.Get_rwalk_token_info(rwalk_aid, token_id)
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Error during query execution: %v", err))
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":    req_status,
		"error":     err_str,
		"TokenInfo": token_info,
	})
}

func apiRwalkMintReport(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	records := rw_storagew.Get_mint_report()
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":  req_status,
		"error":   err_str,
		"Records": records,
	})
}
