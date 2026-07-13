package common

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

func ParseTimeframeParams(c *httpx.Context) (bool, int, int, int) {
	var err error
	pInitTs := c.Param("init_ts")
	var initTs int
	if len(pInitTs) > 0 {
		initTs, err = strconv.Atoi(pInitTs)
		if err != nil {
			c.JSON(http.StatusBadRequest, httpx.H{
				"status": 0,
				"error":  fmt.Sprintf("Bad 'init_ts' parameter: %v", err),
			})
			return false, 0, 0, 0
		}
	} else {
		c.JSON(http.StatusBadRequest, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("'init_ts' parameter wasn't provided: %v", err),
		})
		return false, 0, 0, 0
	}

	pFinTs := c.Param("fin_ts")
	var finTs int
	if len(pFinTs) > 0 {
		finTs, err = strconv.Atoi(pFinTs)
		if err != nil {
			c.JSON(http.StatusBadRequest, httpx.H{
				"status": 0,
				"error":  fmt.Sprintf("'fin_ts' parameter: %v", err),
			})
			return false, 0, 0, 0
		}
	} else {
		c.JSON(http.StatusBadRequest, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("'fin_ts' parameter wasn't provided: %v", err),
		})
		return false, 0, 0, 0
	}
	if finTs == 0 {
		finTs = 2147483647
	}

	pIntervalSecs := c.Param("interval_secs")
	var intervalSecs int
	if len(pIntervalSecs) > 0 {
		intervalSecs, err = strconv.Atoi(pIntervalSecs)
		if err != nil {
			c.JSON(http.StatusBadRequest, httpx.H{
				"status": 0,
				"error":  fmt.Sprintf("Bad 'interval_secs' parameter: %v", err),
			})
			return false, 0, 0, 0
		}
	} else {
		c.JSON(http.StatusBadRequest, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("'interval_secs' parameter wasn't provided: %v", err),
		})
		return false, 0, 0, 0
	}
	if intervalSecs == 0 {
		intervalSecs = finTs - initTs // we can't divide by 0
	}
	return true, initTs, finTs, intervalSecs
}

func ParseOffsetLimitParamsJSON(c *httpx.Context) (bool, int, int) {
	var err error
	pOffset := c.Param("offset")
	var offset int
	if len(pOffset) > 0 {
		offset, err = strconv.Atoi(pOffset)
		if err != nil {
			c.JSON(http.StatusBadRequest, httpx.H{
				"status": 0,
				"error":  fmt.Sprintf("Bad 'offset' parameter: %v", err),
			})
			return false, 0, 0
		}
	} else {
		c.JSON(http.StatusBadRequest, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("'offset' parameter wasn't provided: %v", err),
		})
		return false, 0, 0
	}

	pLimit := c.Param("limit")
	var limit int
	if len(pLimit) > 0 {
		limit, err = strconv.Atoi(pLimit)
		if err != nil {
			c.JSON(http.StatusBadRequest, httpx.H{
				"status": 0,
				"error":  fmt.Sprintf("'limit' parameter: %v", err),
			})
			return false, 0, 0
		}
	} else {
		c.JSON(http.StatusBadRequest, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("'limit' parameter wasn't provided: %v", err),
		})
		return false, 0, 0
	}
	if limit == 0 {
		limit = 20
	}
	return true, offset, limit
}

// ParseInitFinTsParams reads init_ts and fin_ts path params (no interval_secs).
func ParseInitFinTsParams(c *httpx.Context) (bool, int, int) {
	initTs, err := strconv.Atoi(c.Param("init_ts"))
	if err != nil {
		c.JSON(http.StatusBadRequest, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("Bad 'init_ts' parameter: %v", err),
		})
		return false, 0, 0
	}
	finTs, err := strconv.Atoi(c.Param("fin_ts"))
	if err != nil {
		c.JSON(http.StatusBadRequest, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("Bad 'fin_ts' parameter: %v", err),
		})
		return false, 0, 0
	}
	if finTs == 0 {
		finTs = 2147483647
	}
	return true, initTs, finTs
}
