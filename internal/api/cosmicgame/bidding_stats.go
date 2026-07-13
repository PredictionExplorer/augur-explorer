package cosmicgame

import (
	"net/http"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	cgdb "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

const recentSpikeWindowSecs = 30 * 24 * 3600

func (a *API) handleBiddingActivity(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
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

	buckets, err := a.repo.BidFrequencyByPeriod(c.Request.Context(), initTs, finTs, intervalSecs)
	if err != nil {
		a.respondStoreError(c, err)
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

	c.JSON(http.StatusOK, httpx.H{
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

func (a *API) handleBiddingFrequency(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
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

	buckets, err := a.repo.BidFrequencyByPeriod(c.Request.Context(), initTs, finTs, intervalSecs)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":           1,
		"error":            "",
		"InitTs":           initTs,
		"FinTs":            finTs,
		"Interval":         intervalSecs,
		"FrequencyHistory": buckets,
	})
}

func (a *API) handleBiddingTopActivePeriods(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
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

	topBidders, periods, err := a.repo.TopBidderActivePeriods(c.Request.Context(), topN, initTs, finTs, gapHours, minBids)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
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

// handleBidTypeRatio serves the per-interval bid-type composition
// used by the 100% stacked area chart. Params are read from the query string:
//
//	from_ts       unix seconds, start of range (default 0)
//	to_ts         unix seconds, end of range   (default now / 2147483647)
//	intervalSecs sampling window size         (default 86400 = 1 day)
func (a *API) handleBidTypeRatio(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	fromTs := cgdb.ParseOptionalIntQuery(c.Query("from_ts"), 0)
	toTs := cgdb.ParseOptionalIntQuery(c.Query("to_ts"), 2147483647)
	intervalSecs := cgdb.ParseOptionalIntQuery(c.Query("interval_secs"), 86400)
	if intervalSecs <= 0 {
		intervalSecs = 86400
	}

	buckets, err := a.repo.BidTypeRatioByPeriod(c.Request.Context(), fromTs, toTs, intervalSecs)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":       1,
		"error":        "",
		"FromTs":       fromTs,
		"ToTs":         toTs,
		"Interval":     intervalSecs,
		"RatioHistory": buckets,
	})
}

func (a *API) handleBiddingTimeBounds(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	minTs, maxTs, err := a.repo.BidTimeBounds(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status": 1,
		"error":  "",
		"MinTs":  minTs,
		"MaxTs":  maxTs,
	})
}
