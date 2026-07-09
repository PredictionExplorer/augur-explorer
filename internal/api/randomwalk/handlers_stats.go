package randomwalk

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// Token stats (API)
func apiRwalkTokenStats(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	rwalk_aid := addrs.RandomWalkAid
	p_rwalk_addr := addrs.RandomWalk
	stats, err := rwRepo.RandomWalkStats(c.Request.Context(), rwalk_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     1,
		"error":      "",
		"TokenStats": stats,
		"RWalkAid":   rwalk_aid,
		"RWalkAddr":  p_rwalk_addr,
	})
}

// Market stats (API)
func apiRwalkMarketStats(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	market_aid := addrs.MarketPlaceAid
	p_market_addr := addrs.MarketPlace
	stats, err := rwRepo.MarketStats(c.Request.Context(), market_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      1,
		"error":       "",
		"MarketStats": stats,
		"MarketAid":   market_aid,
		"MarketAddr":  p_market_addr,
	})
}

// Trading volume by period (API)
func apiRwalkTradingVolumeByPeriod(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	market_aid := addrs.MarketPlaceAid
	success, init_ts, fin_ts, interval_secs := common.ParseTimeframeParams(c)
	if !success {
		return
	}
	vol_hist, err := rwRepo.MarketTradingVolumeByPeriod(c.Request.Context(), market_aid, init_ts, fin_ts, interval_secs)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":        1,
		"error":         "",
		"VolumeHistory": vol_hist,
		"InitTs":        init_ts,
		"FinTs":         fin_ts,
		"Interval":      interval_secs,
	})
}

// Mint intervals (API)
func apiRwalkMintIntervals(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	rwalk_aid := addrs.RandomWalkAid
	p_rwalk_addr := addrs.RandomWalk
	mint_intervals, err := rwRepo.MintIntervals(c.Request.Context(), rwalk_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":        1,
		"error":         "",
		"MintIntervals": mint_intervals,
		"RWalkAid":      rwalk_aid,
		"RWalkAddr":     p_rwalk_addr,
	})
}

// Withdrawal chart (API)
func apiRwalkWithdrawalChart(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	rwalk_aid := addrs.RandomWalkAid
	p_rwalk_addr := addrs.RandomWalk
	withdrawal_entries, err := rwRepo.WithdrawalChart(c.Request.Context(), rwalk_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	withdrawal_data := common.BuildJSRandomwalkWithdrawalChart(&withdrawal_entries)
	rwalk_stats, err := rwRepo.RandomWalkStats(c.Request.Context(), rwalk_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":             1,
		"error":              "",
		"WithdrawalEntries":  withdrawal_entries,
		"WithdrawalData":     withdrawal_data,
		"ContractStatistics": rwalk_stats,
		"RWalkAid":           rwalk_aid,
		"RWalkAddr":          p_rwalk_addr,
	})
}

// Floor price over time (API)
func apiRwalkFloorPriceOverTime(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	rwalk_aid := addrs.RandomWalkAid
	p_rwalk_addr := addrs.RandomWalk
	market_aid := addrs.MarketPlaceAid
	p_market_addr := addrs.MarketPlace
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
	price_entries, err := rwRepo.FloorPriceByPeriod(c.Request.Context(), rwalk_aid, market_aid, ini, fin, interval)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	price_data := common.BuildJSFloorPriceData(&price_entries)
	rwalk_stats, err := rwRepo.RandomWalkStats(c.Request.Context(), rwalk_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":             1,
		"error":              "",
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

// Top 5 traded tokens (API only)
func apiRwalkTop5TradedTokens(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	top5toks, err := rwRepo.Top5TradedTokens(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":           1,
		"error":            "",
		"Top5TradedTokens": top5toks,
	})
}

// Mint report (API)
func apiRwalkMintReport(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	records, err := rwRepo.MintReport(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"error":   "",
		"Records": records,
	})
}
