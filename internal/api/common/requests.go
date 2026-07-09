package common

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// RespondError reports a request error as JSON. It used to render the
// server-side HTML error page; the HTML explorer was removed, so it now
// emits the same payload as RespondErrorJSON (signature kept for call sites).
func RespondError(c *httpx.Context, errorText string) {
	c.JSON(http.StatusBadRequest, httpx.H{
		"status": 0,
		"error":  errorText,
	})
}

func RespondErrorJSON(c *httpx.Context, errorText string) {
	c.JSON(http.StatusBadRequest, httpx.H{
		"status": 0,
		"error":  errorText,
	})
}

// RespondInternalErrorJSON reports a server-side failure (typically a
// database error) in the legacy error envelope without leaking internal
// details to the client. Handlers log the underlying error before calling it.
func RespondInternalErrorJSON(c *httpx.Context) {
	c.JSON(http.StatusInternalServerError, httpx.H{
		"status": 0,
		"error":  "Internal server error",
	})
}

// ParseIntFromRemoteOrError parses an integer request parameter. The
// jsonOutput flag used to select between JSON and HTML error rendering;
// HTML pages are gone, so both branches now emit the same JSON error.
func ParseIntFromRemoteOrError(c *httpx.Context, jsonOutput bool, asciiInt *string) (int64, bool) {
	p, err := strconv.ParseInt(*asciiInt, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("Can't parse integer parameter: %v", err),
		})
		return 0, false
	}
	return p, true
}

// ParseTimeframeIniFin parses the init_ts/fin_ts path parameters. The
// jsonOutput flag used to select between JSON and HTML error rendering;
// HTML pages are gone, so both branches now emit the same JSON error.
func ParseTimeframeIniFin(c *httpx.Context, jsonOutput bool) (bool, int, int) {
	var err error
	pInitTs := c.Param("init_ts")
	var initTs int = 0
	if len(pInitTs) > 0 {
		initTs, err = strconv.Atoi(pInitTs)
		if err != nil {
			c.JSON(http.StatusBadRequest, httpx.H{
				"status": 0,
				"error":  fmt.Sprintf("Bad 'init_ts' parameter: %v", err),
			})
			return false, 0, 0
		}
	} else {
		c.JSON(http.StatusBadRequest, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("'init_ts' parameter wasn't provided: %v", err),
		})
		return false, 0, 0
	}

	pFinTs := c.Param("fin_ts")
	var finTs int = 0
	if len(pFinTs) > 0 {
		finTs, err = strconv.Atoi(pFinTs)
		if err != nil {
			c.JSON(http.StatusBadRequest, httpx.H{
				"status": 0,
				"error":  fmt.Sprintf("'fin_ts' parameter: %v", err),
			})
			return false, 0, 0
		}
	} else {
		c.JSON(http.StatusBadRequest, httpx.H{
			"status": 0,
			"error":  fmt.Sprintf("'fin_ts' parameter wasn't provided: %v", err),
		})
		return false, 0, 0
	}
	if finTs == 0 {
		finTs = 2147483647
	}
	return true, initTs, finTs
}
