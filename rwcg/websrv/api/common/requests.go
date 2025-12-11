package common

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, errorText string) {
	c.HTML(http.StatusBadRequest, "error.html", gin.H{
		"title":    "RWCG: Error",
		"ErrDescr": errorText,
	})
}

func RespondErrorJSON(c *gin.Context, errorText string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": 0,
		"error":  errorText,
	})
}

func JSONValidateAndLookupAddressOrAid(c *gin.Context, pAddr *string) (string, int64, bool) {
	// Note: this function transforms address in checksumed format
	var aid int64 = 0
	if len(*pAddr) > 0 {
		aid, err := strconv.ParseInt(*pAddr, 10, 64)
		if err == nil {
			var addr string
			addr, err = Ctx.Db.Lookup_address(aid)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": 0,
					"error":  fmt.Sprintf("Address with ID=%v not found", aid),
				})
				return "", aid, false
			}
			return addr, aid, true
		} else {
			aid = 0
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 0,
			"error":  fmt.Sprintf("Empty 'address' parameter for lookup"),
		})
		return "", 0, false
	}
	address, valid := IsAddressValid(c, true, *pAddr)
	if !valid {
		return "", 0, false
	}
	aid, err := Ctx.Db.Nonfatal_lookup_address_id(address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 0,
			"error":  fmt.Sprintf("Address not found in the DB"),
		})
		return "", 0, false
	}
	return address, aid, true
}

func ShowMarketNotFoundError(c *gin.Context, jsonOutput bool, addr *string) {
	if jsonOutput {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 0,
			"error":  fmt.Sprintf("Market with address %v wasn't found", *addr),
		})
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title":    "RWCG: Error",
			"ErrDescr": fmt.Sprintf("Market with address %v wasn't found", *addr),
		})
	}
}

func ParseIntFromRemoteOrError(c *gin.Context, jsonOutput bool, asciiInt *string) (int64, bool) {
	p, err := strconv.ParseInt(*asciiInt, 10, 64)
	if err != nil {
		if jsonOutput {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 0,
				"error":  fmt.Sprintf("Can't parse integer parameter : ", err),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title":    "RWCG: Error",
				"ErrDescr": fmt.Sprintf("Can't parse integer parameter : ", err),
			})
		}
		return 0, false
	}
	return p, true
}

func ParseTimeframeIniFin(c *gin.Context, jsonOutput bool) (bool, int, int) {
	var err error
	pInitTs := c.Param("init_ts")
	var initTs int = 0
	if len(pInitTs) > 0 {
		initTs, err = strconv.Atoi(pInitTs)
		if err != nil {
			if jsonOutput {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": 0,
					"error":  fmt.Sprintf("Bad 'init_ts' parameter: %v", err),
				})
			} else {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title":    "Error",
					"ErrDescr": fmt.Sprintf("Bad 'init_ts' parameter: %v", err),
				})
			}
			return false, 0, 0
		}
	} else {
		if jsonOutput {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 0,
				"error":  fmt.Sprintf("'init_ts' parameter wasn't provided: %v", err),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title":    "Error",
				"ErrDescr": fmt.Sprintf("'init_ts' parameter wasn't provided: %v", err),
			})
		}
		return false, 0, 0
	}

	pFinTs := c.Param("fin_ts")
	var finTs int = 0
	if len(pFinTs) > 0 {
		finTs, err = strconv.Atoi(pFinTs)
		if err != nil {
			if jsonOutput {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": 0,
					"error":  fmt.Sprintf("'fin_ts' parameter: %v", err),
				})
			} else {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title":    "Error",
					"ErrDescr": fmt.Sprintf("'fin_ts' parameter: %v", err),
				})
			}
			return false, 0, 0
		}
	} else {
		if jsonOutput {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 0,
				"error":  fmt.Sprintf("'fin_ts' parameter wasn't provided: %v", err),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title":    "Error",
				"ErrDescr": fmt.Sprintf("'fin_ts' parameter wasn't provided: %v", err),
			})
		}
		return false, 0, 0
	}
	if finTs == 0 {
		finTs = 2147483647
	}
	return true, initTs, finTs
}

func ParseOffsetLimitParamsHTML(c *gin.Context) (bool, int, int) {
	var err error
	pOffset := c.Param("offset")
	var offset int = 0
	if len(pOffset) > 0 {
		offset, err = strconv.Atoi(pOffset)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title":    "Error",
				"ErrDescr": fmt.Sprintf("Bad 'offset' parameter: %v", err),
			})
			return false, 0, 0
		}
	}

	pLimit := c.Param("limit")
	var limit int = 0
	if len(pLimit) > 0 {
		limit, err = strconv.Atoi(pLimit)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title":    "Error",
				"ErrDescr": fmt.Sprintf("Bad 'limit' parameter: %v", err),
			})
			return false, 0, 0
		}
	}
	return true, offset, limit
}
