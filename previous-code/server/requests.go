package main
import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"

)
func respond_error(c *gin.Context,error_text string) {

	c.HTML(http.StatusBadRequest, "error.html", gin.H{
		"title": "Augur Markets: Error",
		"ErrDescr": error_text,
	})
}
func respond_error_json(c *gin.Context,error_text string) {

	c.JSON(http.StatusBadRequest, gin.H{
		"status": 0,
		"error": error_text,
	})
}
func json_validate_and_lookup_address_or_aid(c *gin.Context,p_addr *string) (string,int64,bool) {
	// Note: this function transforms address in checksumed format
	var aid int64 = 0
	if len(*p_addr) > 0 {
		aid, err := strconv.ParseInt(*p_addr,10,64)
		if err == nil {
			var addr string
			addr,err = augur_srv.db_augur.Lookup_address(aid)
			if err!=nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"status":0,
					"error":fmt.Sprintf("Address with ID=%v not found",aid),
				})
				return "",aid,false
			}
			return addr,aid,true
		} else {
			aid = 0
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("Empty 'address' parameter for lookup"),
		})
		return "",0,false
	}
	address,valid:=is_address_valid(c,true,*p_addr)
	if !valid {
		return "",0,false
	}
	aid,err := augur_srv.db_augur.Nonfatal_lookup_address_id(address)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("Address not found in the DB"),
		})
		return "",0,false
	}
	return address,aid,true
}
func show_market_not_found_error(c *gin.Context,json_output bool,addr *string) {

	if json_output {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 0,
			"error": fmt.Sprintf("Market with address %v wasn't found",*addr),
		})
	} else {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Augur Markets: Error",
			"ErrDescr": fmt.Sprintf("Market with address %v wasn't found",*addr),
		})
	}
}
func parse_int_from_remote_or_error(c *gin.Context,json_output bool,ascii_int *string) (int64,bool) {
	p, err := strconv.ParseInt(*ascii_int,10,64)
	if err != nil {
		if json_output {
			c.JSON(http.StatusBadRequest, gin.H{
				"status" : 0,
				"error": fmt.Sprintf("Can't parse integer parameter : ",err),
			})
		} else {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"title": "Augur Markets: Error",
				"ErrDescr": fmt.Sprintf("Can't parse integer parameter : ",err),
			})
		}
		return 0,false
	}
	return p,true
}
func parse_timeframe_ini_fin(c *gin.Context,json_output bool) (bool,int,int) {

	var err error
	p_init_ts := c.Param("init_ts")
	var init_ts int = 0
	if len(p_init_ts) > 0 {
		init_ts, err = strconv.Atoi(p_init_ts)
		if err != nil {
			if json_output {
				c.JSON(http.StatusBadRequest,gin.H{
					"status":0,
					"error":fmt.Sprintf("Bad 'init_ts' parameter: %v",err),
				})
			} else {
				c.HTML(http.StatusBadRequest,"error.html",gin.H{
					"title":"Error",
					"ErrDescr":fmt.Sprintf("Bad 'init_ts' parameter: %v",err),
				})
			}
			return false,0,0
		}
	} else {
		if json_output {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("'init_ts' parameter wasn't provided: %v",err),
			})
		} else {
			c.HTML(http.StatusBadRequest,"error.html",gin.H{
				"title":"Error",
				"ErrDescr":fmt.Sprintf("'init_ts' parameter wasn't provided: %v",err),
			})

		}
		return false,0,0
	}

	p_fin_ts := c.Param("fin_ts")
	var fin_ts int = 0
	if len(p_fin_ts) > 0 {
		fin_ts, err = strconv.Atoi(p_fin_ts)
		if err != nil {
			if json_output {
				c.JSON(http.StatusBadRequest,gin.H{
					"status":0,
					"error":fmt.Sprintf("'fin_ts' parameter: %v",err),
				})
			} else {
				c.HTML(http.StatusBadRequest,"error.html",gin.H{
					"title":"Error",
					"ErrDescr":fmt.Sprintf("'fin_ts' parameter: %v",err),
				})
			}
			return false,0,0
		}
	} else {
		if json_output {
			c.JSON(http.StatusBadRequest,gin.H{
				"status":0,
				"error":fmt.Sprintf("'fin_ts' parameter wasn't provided: %v",err),
			})
		} else {
			c.HTML(http.StatusBadRequest,"error.html",gin.H{
				"title":"Error",
				"ErrDescr":fmt.Sprintf("'fin_ts' parameter wasn't provided: %v",err),
			})
		}
		return false,0,0
	}
	if fin_ts == 0 {
		fin_ts = 2147483647
	}
	return true,init_ts,fin_ts
}
func parse_offset_limit_params_html(c *gin.Context) (bool,int,int) {

	var err error
	p_offset := c.Param("offset")
	var offset int = 0
	if len(p_offset) > 0 {
		offset, err = strconv.Atoi(p_offset)
		if err != nil {
			c.HTML(http.StatusBadRequest,"error.html",gin.H{
				"title":"Error",
				"ErrDescr":fmt.Sprintf("Bad 'offset' parameter: %v",err),
			})
			return false,0, 0
		}
	}

	p_limit := c.Param("limit")
	var limit int = 0
	if len(p_limit) > 0 {
		limit, err = strconv.Atoi(p_limit)
		if err != nil {
			c.HTML(http.StatusBadRequest,"error.html",gin.H{
				"title":"Error",
				"ErrDescr":fmt.Sprintf("Bad 'limit' parameter: %v",err),
			})
			return false,0, 0
		}
	}
	return true,offset,limit
}
