package randomwalk

import (
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// Trading history (API)
func apiRwalkTradingHistory(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
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
	sales, err := rwRepo.TradingHistory(c.Request.Context(), market_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":     1,
		"error":      "",
		"Sales":      sales,
		"MarketAid":  market_aid,
		"MarketAddr": p_market_addr,
	})
}

// Sale history (API)
func apiRwalkSaleHistory(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
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
	sales, err := rwRepo.SaleHistory(c.Request.Context(), market_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":          1,
		"error":           "",
		"Trading":         sales,
		"MarketPlaceAddr": p_market_addr,
		"MarketPlaceAid":  market_aid,
	})
}

// Trading history by user (API)
func apiRwalkTradingHistoryByUser(c *httpx.Context) {
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
	user_addr, err := rwStore.AddressByID(c.Request.Context(), user_aid)
	if err != nil {
		common.RespondErrorJSON(c, "Address lookup on user_aid failed")
		return
	}
	user_trading, err := rwRepo.TradingHistoryByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":      1,
		"error":       "",
		"UserTrading": user_trading,
		"UserAid":     user_aid,
		"UserAddr":    user_addr,
	})
}
