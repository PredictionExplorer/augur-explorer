package randomwalk

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// Trading history (API)
func apiRwalkTradingHistory(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs := rwContractAddrs()
	p_market_addr := addrs.MarketPlace
	var market_aid int64
	if p_market_addr == "0x0000000000000000000000000000000000000000" {
		market_aid = 0
	} else {
		market_aid = addrs.MarketPlaceAid
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

// Sale history (API)
func apiRwalkSaleHistory(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs := rwContractAddrs()
	p_market_addr := addrs.MarketPlace
	var market_aid int64
	if p_market_addr == "0x0000000000000000000000000000000000000000" {
		market_aid = 0
	} else {
		market_aid = addrs.MarketPlaceAid
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

// Trading history by user (API)
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
