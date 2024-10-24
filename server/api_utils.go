package main
import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"

)
func parse_timeframe_params(c *gin.Context) (bool,int,int,int) {

	var err error
	p_init_ts := c.Param("init_ts")
	var init_ts int = 0
	if len(p_init_ts) > 0 {
		init_ts, err = strconv.Atoi(p_init_ts)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("Bad 'init_ts' parameter: %v",err),
			})
			return false,0,0,0
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("'init_ts' parameter wasn't provided: %v",err),
		})
		return false,0,0,0
	}

	p_fin_ts := c.Param("fin_ts")
	var fin_ts int = 0
	if len(p_fin_ts) > 0 {
		fin_ts, err = strconv.Atoi(p_fin_ts)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("'fin_ts' parameter: %v",err),
			})
			return false,0,0,0
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("'fin_ts' parameter wasn't provided: %v",err),
		})
		return false,0,0,0
	}
	if fin_ts == 0 {
		fin_ts = 2147483647
	}

	p_interval_secs := c.Param("interval_secs")
	var interval_secs int = fin_ts - init_ts
	if len(p_interval_secs) > 0 {
		interval_secs, err = strconv.Atoi(p_interval_secs)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("Bad 'interval_secs' parameter: %v",err),
			})
			return false,0,0,0
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("'interval_secs' parameter wasn't provided: %v",err),
		})
		return false,0,0,0
	}
	if interval_secs == 0 {
		interval_secs = fin_ts - init_ts // we can't divide by 0
	}
	return true,init_ts,fin_ts,interval_secs
}
func parse_timeframe_ini_fin(c *gin.Context) (bool,int,int) {

	var err error
	p_init_ts := c.Param("init_ts")
	var init_ts int = 0
	if len(p_init_ts) > 0 {
		init_ts, err = strconv.Atoi(p_init_ts)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("Bad 'init_ts' parameter: %v",err),
			})
			return false,0,0
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("'init_ts' parameter wasn't provided: %v",err),
		})
		return false,0,0
	}

	p_fin_ts := c.Param("fin_ts")
	var fin_ts int = 0
	if len(p_fin_ts) > 0 {
		fin_ts, err = strconv.Atoi(p_fin_ts)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("'fin_ts' parameter: %v",err),
			})
			return false,0,0
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("'fin_ts' parameter wasn't provided: %v",err),
		})
		return false,0,0
	}
	if fin_ts == 0 {
		fin_ts = 2147483647
	}
	return true,init_ts,fin_ts
}
func parse_offset_limit_params(c *gin.Context) (bool,int,int) {

	var err error
	p_offset := c.Param("offset")
	var offset int = 0
	if len(p_offset) > 0 {
		offset, err = strconv.Atoi(p_offset)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("Bad 'offset' parameter: %v",err),
			})
			return false,0,0
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("'offset' parameter wasn't provided: %v",err),
		})
		return false,0,0
	}

	p_limit := c.Param("limit")
	var limit int = 0
	if len(p_limit) > 0 {
		limit, err = strconv.Atoi(p_limit)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("'limit' parameter: %v",err),
			})
			return false,0,0
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("'limit' parameter wasn't provided: %v",err),
		})
		return false,0,0
	}
	if limit == 0 {
		limit = 20
	}
	return true,offset,limit
}
func parse_outcome_param(c *gin.Context) (bool,int) {
	var err error
	p_outcome := c.Param("outcome")
	var outcome int
	if len(p_outcome) > 0 {
		outcome , err = strconv.Atoi(p_outcome)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("Bad outcome parameter: %v",err),
			})
			return false,0
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("Outcome parameter wasn't provided: %v",err),
		})
		return false,0
	}
	return true,outcome
}
