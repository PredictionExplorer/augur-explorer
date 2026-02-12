package randomwalk

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// Index / dashboard (HTML only)
func rwalk_index_page(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	caddrs := rw_storagew.Get_randomwalk_contract_addresses()
	top5tokens := rw_storagew.Get_top5_traded_tokens()
	c.HTML(http.StatusOK, "rw_index.html", gin.H{
		"ContractAddresses": caddrs,
		"Top5Tokens":        top5tokens,
	})
}

// Token stats (API + HTML)
func apiRwalkTokenStats(c *gin.Context) {
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
	stats := rw_storagew.Get_random_walk_stats(rwalk_aid)
	c.JSON(http.StatusOK, gin.H{
		"status":    1,
		"error":     "",
		"TokenStats": stats,
		"RWalkAid":   rwalk_aid,
		"RWalkAddr":  p_rwalk_addr,
	})
}

func rwalk_token_stats(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondError(c, "Lookup of NFT token address in the Db has failed")
		return
	}
	stats := rw_storagew.Get_random_walk_stats(rwalk_aid)
	c.HTML(http.StatusOK, "rw_token_stats.html", gin.H{
		"TokenStats": stats,
		"RWalkAid":   rwalk_aid,
		"RWalkAddr":  p_rwalk_addr,
	})
}

// Market stats (API + HTML)
func apiRwalkMarketStats(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of Market address in the DB has failed")
		return
	}
	stats := rw_storagew.Get_market_stats(market_aid)
	c.JSON(http.StatusOK, gin.H{
		"status":      1,
		"error":       "",
		"MarketStats": stats,
		"MarketAid":   market_aid,
		"MarketAddr":  p_market_addr,
	})
}

func rwalk_market_stats(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		common.RespondError(c, "Market address doesn't exist in the database")
		return
	}
	stats := rw_storagew.Get_market_stats(market_aid)
	c.HTML(http.StatusOK, "rw_market_stats.html", gin.H{
		"MarketStats": stats,
		"MarketAid":   market_aid,
		"MarketAddr":  p_market_addr,
	})
}

// Trading volume by period (API + HTML)
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
	c.JSON(http.StatusOK, gin.H{
		"status":       1,
		"error":        "",
		"VolumeHistory": vol_hist,
		"InitTs":        init_ts,
		"FinTs":         fin_ts,
		"Interval":      interval_secs,
	})
}

func rwalk_trading_volume_by_period(c *gin.Context) {
	success, init_ts, fin_ts, interval_secs := common.ParseTimeframeParams(c)
	if !success {
		return
	}
	p_market_addr := c.Param("market_addr")
	market_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		common.RespondError(c, "Market address doesn't exist in the database")
		return
	}
	vol_hist := rw_storagew.Get_market_trading_volume_by_period(market_aid, init_ts, fin_ts, interval_secs)
	volume_data := common.BuildJSRandomwalkVolumeHistory(&vol_hist)
	c.HTML(http.StatusOK, "rw_volume_history.html", gin.H{
		"VolumeHistory": vol_hist,
		"VolumeData":    volume_data,
		"InitTs":        init_ts,
		"FinTs":         fin_ts,
		"Interval":      interval_secs,
		"MarketAddr":    p_market_addr,
		"MarketAid":     market_aid,
	})
}

// Mint intervals (API + HTML)
func apiRwalkMintIntervals(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of NFT token failed")
		return
	}
	mint_intervals := rw_storagew.Get_rwalk_mint_intervals(rwalk_aid)
	c.JSON(http.StatusOK, gin.H{
		"status":        1,
		"error":         "",
		"MintIntervals": mint_intervals,
		"RWalkAid":      rwalk_aid,
		"RWalkAddr":     p_rwalk_addr,
	})
}

func rwalk_mint_intervals(c *gin.Context) {
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
	mint_intervals := rw_storagew.Get_rwalk_mint_intervals(rwalk_aid)
	mint_data := common.BuildJSRandomwalkMintIntervals(&mint_intervals)
	c.HTML(http.StatusOK, "rw_mint_intervals.html", gin.H{
		"MintIntervals":    mint_intervals,
		"MintIntervalData": mint_data,
		"RWalkAid":         rwalk_aid,
		"RWalkAddr":        p_rwalk_addr,
	})
}

// Withdrawal chart (API + HTML)
func apiRwalkWithdrawalChart(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of NFT token failed")
		return
	}
	withdrawal_entries := rw_storagew.Get_rwalk_withdrawal_chart(rwalk_aid)
	withdrawal_data := common.BuildJSRandomwalkWithdrawalChart(&withdrawal_entries)
	rwalk_stats := rw_storagew.Get_random_walk_stats(rwalk_aid)
	c.JSON(http.StatusOK, gin.H{
		"status":            1,
		"error":              "",
		"WithdrawalEntries":  withdrawal_entries,
		"WithdrawalData":     withdrawal_data,
		"ContractStatistics": rwalk_stats,
		"RWalkAid":           rwalk_aid,
		"RWalkAddr":          p_rwalk_addr,
	})
}

func rwalk_withdrawal_chart(c *gin.Context) {
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
	c.HTML(http.StatusOK, "rw_withdrawal_chart.html", gin.H{
		"WithdrawalEntries":  withdrawal_entries,
		"WithdrawalData":     withdrawal_data,
		"ContractStatistics": rwalk_stats,
		"RWalkAid":           rwalk_aid,
		"RWalkAddr":          p_rwalk_addr,
	})
}

// Floor price over time (API + HTML)
func apiRwalkFloorPriceOverTime(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
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
	c.JSON(http.StatusOK, gin.H{
		"status":            1,
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

func rwalk_floor_price_over_time(c *gin.Context) {
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
	p_market_addr := c.Param("market_addr")
	market_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_market_addr)
	if err != nil {
		common.RespondError(c, "Market address doesn't exist in the database")
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
	c.HTML(http.StatusOK, "rw_floor_price_over_time.html", gin.H{
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

// Top users (HTML only)
func rwalk_top_users(c *gin.Context) {
	top_profit_makers := rw_storagew.Get_randomwalk_top_profit_makers()
	top_trade_makers := rw_storagew.Get_randomwalk_top_trade_makers()
	top_volume_makers := rw_storagew.Get_randomwalk_top_volume_makers()
	c.HTML(http.StatusOK, "rw_top_users.html", gin.H{
		"title":         "Top 100 Users of RandomWalk Token",
		"ProfitMakers":  top_profit_makers,
		"TradeMakers":  top_trade_makers,
		"VolumeMakers": top_volume_makers,
	})
}

// Top 5 traded tokens (API only)
func apiRwalkTop5TradedTokens(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	top5toks := rw_storagew.Get_top5_traded_tokens()
	c.JSON(http.StatusOK, gin.H{
		"status":            1,
		"error":             "",
		"Top5TradedTokens":  top5toks,
	})
}

// Mint report (API + HTML)
func apiRwalkMintReport(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	records := rw_storagew.Get_mint_report()
	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"error":   "",
		"Records": records,
	})
}

func rwalk_mint_report(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	records := rw_storagew.Get_mint_report()
	c.HTML(http.StatusOK, "rw_mint_report.html", gin.H{
		"Records": records,
	})
}
