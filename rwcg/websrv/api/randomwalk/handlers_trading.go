package randomwalk

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// Trading history (API + HTML)
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
	c.JSON(http.StatusOK, gin.H{
		"status":    1,
		"error":     "",
		"Sales":      sales,
		"MarketAid":  market_aid,
		"MarketAddr": p_market_addr,
	})
}

func rwalk_trading_history(c *gin.Context) {
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
	success, offset, limit := common.ParseOffsetLimitParamsHTML(c)
	if !success {
		return
	}
	sales := rw_storagew.Get_trading_history(market_aid, offset, limit)
	c.HTML(http.StatusOK, "rw_trading_history.html", gin.H{
		"Trading":         sales,
		"MarketPlaceAddr": p_market_addr,
		"MarketPlaceAid":  market_aid,
	})
}

// Sale history (API + HTML)
func apiRwalkSaleHistory(c *gin.Context) {
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
	sales := rw_storagew.Get_sale_history(market_aid, offset, limit)
	c.JSON(http.StatusOK, gin.H{
		"status":          1,
		"error":           "",
		"Trading":         sales,
		"MarketPlaceAddr": p_market_addr,
		"MarketPlaceAid":  market_aid,
	})
}

func rwalk_sale_history(c *gin.Context) {
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
	offset := int(0)
	limit := int(100000)
	sales := rw_storagew.Get_sale_history(market_aid, offset, limit)
	c.HTML(http.StatusOK, "rw_sale_history.html", gin.H{
		"Trading":         sales,
		"MarketPlaceAddr": p_market_addr,
		"MarketPlaceAid":  market_aid,
	})
}

// Trading history by user (API + HTML)
func apiRwalkTradingHistoryByUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
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
		common.RespondErrorJSON(c, "'user_aid' parameter is not set")
		return
	}
	user_addr, err := rw_storagew.S.Lookup_address(user_aid)
	if err != nil {
		common.RespondErrorJSON(c, "Address lookup on user_aid failed")
		return
	}
	user_trading := rw_storagew.Get_trading_history_by_user(user_aid)
	c.JSON(http.StatusOK, gin.H{
		"status":      1,
		"error":       "",
		"UserTrading": user_trading,
		"UserAid":     user_aid,
		"UserAddr":    user_addr,
	})
}

func rwalk_trading_history_by_user(c *gin.Context) {
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
				common.RespondError(c, "Can't resolve user identifier to valid address ID or address hex")
				return
			}
			user_aid, err = rw_storagew.S.Nonfatal_lookup_address_id(p_user_aid)
			if err != nil {
				common.RespondError(c, "Cant find provided user")
				return
			}
		}
	} else {
		common.RespondError(c, "'user_aid' parameter is not set")
		return
	}
	user_addr, err := rw_storagew.S.Lookup_address(user_aid)
	if err != nil {
		common.RespondError(c, fmt.Sprintf("Address lookup on user_aid %v failed: %v", user_aid, err))
		return
	}
	user_trading := rw_storagew.Get_trading_history_by_user(user_aid)
	c.HTML(http.StatusOK, "rw_trading_by_user.html", gin.H{
		"UserTrading": user_trading,
		"UserAid":     user_aid,
		"UserAddr":    user_addr,
	})
}
