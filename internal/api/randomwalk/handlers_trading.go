package randomwalk

import (
	"errors"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Trading history (API).
func (a *API) handleTradingHistory(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	pMarketAddr := addrs.MarketPlace
	var marketAid int64
	if pMarketAddr == "0x0000000000000000000000000000000000000000" {
		marketAid = 0
	} else {
		marketAid = addrs.MarketPlaceAid
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	sales, err := a.repo.TradingHistory(c.Request.Context(), marketAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":     1,
		"error":      "",
		"Sales":      sales,
		"MarketAid":  marketAid,
		"MarketAddr": pMarketAddr,
	})
}

// Sale history (API).
func (a *API) handleSaleHistory(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	pMarketAddr := addrs.MarketPlace
	var marketAid int64
	if pMarketAddr == "0x0000000000000000000000000000000000000000" {
		marketAid = 0
	} else {
		marketAid = addrs.MarketPlaceAid
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	sales, err := a.repo.SaleHistory(c.Request.Context(), marketAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":          1,
		"error":           "",
		"Trading":         sales,
		"MarketPlaceAddr": pMarketAddr,
		"MarketPlaceAid":  marketAid,
	})
}

// Trading history by user (API).
func (a *API) handleTradingHistoryByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAid := c.Param("user_aid")
	var userAid int64
	if len(pUserAid) > 0 {
		var success bool
		userAid, success = common.ParseIntFromRemoteOrError(c, HTTP, &pUserAid)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'user_aid' parameter is not set")
		return
	}
	userAddr, err := a.store.AddressByID(c.Request.Context(), userAid)
	if err != nil {
		if !errors.Is(err, store.ErrNotFound) {
			a.respondStoreError(c, err)
			return
		}
		common.RespondErrorJSON(c, "Address lookup on user_aid failed")
		return
	}
	userTrading, err := a.repo.TradingHistoryByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":      1,
		"error":       "",
		"UserTrading": userTrading,
		"UserAid":     userAid,
		"UserAddr":    userAddr,
	})
}
