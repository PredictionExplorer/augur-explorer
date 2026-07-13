package randomwalk

import (
	"net/http"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// Token stats (API)
func (a *API) handleTokenStats(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	rwalkAid := addrs.RandomWalkAid
	pRwalkAddr := addrs.RandomWalk
	stats, err := a.repo.RandomWalkStats(c.Request.Context(), rwalkAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":     1,
		"error":      "",
		"TokenStats": stats,
		"RWalkAid":   rwalkAid,
		"RWalkAddr":  pRwalkAddr,
	})
}

// Market stats (API)
func (a *API) handleMarketStats(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	marketAid := addrs.MarketPlaceAid
	pMarketAddr := addrs.MarketPlace
	stats, err := a.repo.MarketStats(c.Request.Context(), marketAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":      1,
		"error":       "",
		"MarketStats": stats,
		"MarketAid":   marketAid,
		"MarketAddr":  pMarketAddr,
	})
}

// Trading volume by period (API)
func (a *API) handleTradingVolumeByPeriod(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	marketAid := addrs.MarketPlaceAid
	success, initTs, finTs, intervalSecs := common.ParseTimeframeParams(c)
	if !success {
		return
	}
	volHist, err := a.repo.MarketTradingVolumeByPeriod(c.Request.Context(), marketAid, initTs, finTs, intervalSecs)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":        1,
		"error":         "",
		"VolumeHistory": volHist,
		"InitTs":        initTs,
		"FinTs":         finTs,
		"Interval":      intervalSecs,
	})
}

// Mint intervals (API)
func (a *API) handleMintIntervals(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	rwalkAid := addrs.RandomWalkAid
	pRwalkAddr := addrs.RandomWalk
	mintIntervals, err := a.repo.MintIntervals(c.Request.Context(), rwalkAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":        1,
		"error":         "",
		"MintIntervals": mintIntervals,
		"RWalkAid":      rwalkAid,
		"RWalkAddr":     pRwalkAddr,
	})
}

// Withdrawal chart (API)
func (a *API) handleWithdrawalChart(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	rwalkAid := addrs.RandomWalkAid
	pRwalkAddr := addrs.RandomWalk
	withdrawalEntries, err := a.repo.WithdrawalChart(c.Request.Context(), rwalkAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	withdrawalData := common.BuildJSRandomwalkWithdrawalChart(&withdrawalEntries)
	rwalkStats, err := a.repo.RandomWalkStats(c.Request.Context(), rwalkAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":             1,
		"error":              "",
		"WithdrawalEntries":  withdrawalEntries,
		"WithdrawalData":     withdrawalData,
		"ContractStatistics": rwalkStats,
		"RWalkAid":           rwalkAid,
		"RWalkAddr":          pRwalkAddr,
	})
}

// Floor price over time (API)
func (a *API) handleFloorPriceOverTime(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	rwalkAid := addrs.RandomWalkAid
	pRwalkAddr := addrs.RandomWalk
	marketAid := addrs.MarketPlaceAid
	pMarketAddr := addrs.MarketPlace
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
	priceEntries, err := a.repo.FloorPriceByPeriod(c.Request.Context(), rwalkAid, marketAid, ini, fin, interval)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	priceData := common.BuildJSFloorPriceData(&priceEntries)
	rwalkStats, err := a.repo.RandomWalkStats(c.Request.Context(), rwalkAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":             1,
		"error":              "",
		"PriceEntries":       priceEntries,
		"PriceData":          priceData,
		"ContractStatistics": rwalkStats,
		"RWalkAid":           rwalkAid,
		"RWalkAddr":          pRwalkAddr,
		"MarketAid":          marketAid,
		"MarketAddr":         pMarketAddr,
		"InitTs":             ini,
		"FinTs":              fin,
		"Interval":           interval,
	})
}

// Top 5 traded tokens (API only)
func (a *API) handleTop5TradedTokens(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	top5toks, err := a.repo.Top5TradedTokens(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":           1,
		"error":            "",
		"Top5TradedTokens": top5toks,
	})
}

// Mint report (API)
func (a *API) handleMintReport(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	records, err := a.repo.MintReport(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":  1,
		"error":   "",
		"Records": records,
	})
}
