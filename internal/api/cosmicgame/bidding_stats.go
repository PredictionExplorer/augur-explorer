package cosmicgame

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	cgdb "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const recentSpikeWindowSecs = 30 * 24 * 3600

func api_cosmic_game_bidding_activity(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, initTs, finTs, intervalSecs := common.ParseTimeframeParams(c)
	if !success {
		return
	}
	if intervalSecs <= 0 {
		intervalSecs = 3600
	}

	buckets, err := arbRepo.BidFrequencyByPeriod(c.Request.Context(), initTs, finTs, intervalSecs)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	spikes := cgdb.DetectBidSpikes(buckets, intervalSecs)

	recentSpikeIndex := int64(-1)
	nowTs := time.Now().Unix()
	for _, spike := range spikes {
		if spike.StartTs >= nowTs-recentSpikeWindowSecs {
			recentSpikeIndex = int64(spike.Index)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":           1,
		"error":            "",
		"InitTs":           initTs,
		"FinTs":            finTs,
		"Interval":         intervalSecs,
		"FrequencyHistory": buckets,
		"Spikes":           spikes,
		"RecentSpikeIndex": recentSpikeIndex,
		"RecentWindowSecs": recentSpikeWindowSecs,
	})
}

func api_cosmic_game_bidding_frequency(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, initTs, finTs, intervalSecs := common.ParseTimeframeParams(c)
	if !success {
		return
	}
	if intervalSecs <= 0 {
		intervalSecs = 86400
	}

	buckets, err := arbRepo.BidFrequencyByPeriod(c.Request.Context(), initTs, finTs, intervalSecs)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":           1,
		"error":            "",
		"InitTs":           initTs,
		"FinTs":            finTs,
		"Interval":         intervalSecs,
		"FrequencyHistory": buckets,
	})
}

func api_cosmic_game_bidding_top_active_periods(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, initTs, finTs := common.ParseInitFinTsParams(c)
	if !success {
		return
	}

	topN := cgdb.ParseOptionalIntQuery(c.Param("n"), 20)
	gapHours := cgdb.ParseOptionalIntQuery(c.Query("gap_hours"), 6)
	minBids := cgdb.ParseOptionalIntQuery(c.Query("min_bids"), 2)

	topBidders, periods, err := arbRepo.TopBidderActivePeriods(c.Request.Context(), topN, initTs, finTs, gapHours, minBids)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":        1,
		"error":         "",
		"InitTs":        initTs,
		"FinTs":         finTs,
		"TopN":          topN,
		"GapHours":      gapHours,
		"MinBids":       minBids,
		"TopBidders":    topBidders,
		"ActivePeriods": periods,
	})
}

// api_cosmic_game_bid_type_ratio serves the per-interval bid-type composition
// used by the 100% stacked area chart. Params are read from the query string:
//
//	from_ts       unix seconds, start of range (default 0)
//	to_ts         unix seconds, end of range   (default now / 2147483647)
//	interval_secs sampling window size         (default 86400 = 1 day)
func api_cosmic_game_bid_type_ratio(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	fromTs := cgdb.ParseOptionalIntQuery(c.Query("from_ts"), 0)
	toTs := cgdb.ParseOptionalIntQuery(c.Query("to_ts"), 2147483647)
	intervalSecs := cgdb.ParseOptionalIntQuery(c.Query("interval_secs"), 86400)
	if intervalSecs <= 0 {
		intervalSecs = 86400
	}

	buckets, err := arbRepo.BidTypeRatioByPeriod(c.Request.Context(), fromTs, toTs, intervalSecs)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":       1,
		"error":        "",
		"FromTs":       fromTs,
		"ToTs":         toTs,
		"Interval":     intervalSecs,
		"RatioHistory": buckets,
	})
}

func api_cosmic_game_bidding_time_bounds(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	minTs, maxTs, err := arbRepo.BidTimeBounds(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"error":  "",
		"MinTs":  minTs,
		"MaxTs":  maxTs,
	})
}
