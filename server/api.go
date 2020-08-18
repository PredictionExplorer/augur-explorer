/// API v1
package main
import (
	"fmt"
	"strconv"
	"strings"

	"net/http"
	"github.com/gin-gonic/gin"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)

func a1_active_market_ids(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	off_str := c.Query("off")
	var off int = 0
	var err error
	if len(off_str) > 0 {
		off, err = strconv.Atoi(off_str)
		if err != nil {
			c.JSON(422,gin.H{
				"MarketIDs": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad offset parameter: %v",err),
			})
			return
		}
	}
	p_sort := c.Query("sort")
	var sort int = 0
	if len(p_sort) > 0 {
		sort, err = strconv.Atoi(p_sort)
		if err != nil {
			c.JSON(422,gin.H{
				"MarketIDs": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad sort parameter: %v",err),
			})
			return
		}
	}
	ids := augur_srv.storage.Get_active_market_ids(sort,off,1000000)
	var status int = 1
	c.JSON(200,gin.H{
		"MarketIDs": ids,
		"status":status,
		"error":"",
	})
}
func a1_market_card(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var err error
	var market_aid int64 = 0
	p_market_aid := c.Param("market_aid")
	if len(p_market_aid ) > 0 {
		market_aid, err = strconv.ParseInt(p_market_aid,10,64)
		if err != nil {
			c.JSON(422,gin.H{
				"MarketInfo": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad integer for market_aid parameter: %v",err),
			})
			return
		}
	} else {
		c.JSON(422,gin.H{
			"MarketInfo": make([]int64,0,0),
			"status":0,
			"error":fmt.Sprintf("market_aid parameter wasn't provided"),
		})
		return
	}
	mkt_info,err := augur_srv.storage.Get_market_card_data(market_aid)
	fmt.Printf("mkt_info=%+v",mkt_info)
	var status int = 0
	var err_str string = ""
	if err == nil {
		status = 1
	} else {
		err_str = err.Error()
	}
	c.JSON(200,gin.H{
		"MarketInfo": mkt_info,
		"status":status,
		"error":err_str,
	})
}
func a1_active_markets(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	p_start := c.Param("start")
	p_num_rows := c.Param("num_rows")

	var start,num_rows int
	var err error
	if len(p_start) > 0 {
		start, err = strconv.Atoi(p_start)
		if err != nil {
			c.JSON(422,gin.H{
				"Markets": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad number for 'start' parameter: %v",err),
			})
			return
		}
	} else {
		c.JSON(422,gin.H{
			"Markets": make([]int64,0,0),
			"status":0,
			"error":fmt.Sprintf("Empty 'start' parameter for page offset"),
		})
	}
	if len(p_num_rows) > 0 {
		num_rows, err = strconv.Atoi(p_num_rows)
		if err != nil {
			c.JSON(422,gin.H{
				"Markets": make([]int64,0,0),
				"status":0,
				"error":fmt.Sprintf("Bad number for 'num_rows' parameter: %v",err),
			})
			return
		}
		if num_rows < 1 {
			num_rows = DEFAULT_MARKET_ROWS_LIMIT
		}
	} else {
		c.JSON(422,gin.H{
			"Markets": make([]int64,0,0),
			"status":0,
			"error":fmt.Sprintf("Empty 'num_rows' parameter for number of records to return"),
		})
	}

	markets := augur_srv.storage.Get_active_market_list(start,num_rows)
	var status int = 1
	c.JSON(200,gin.H{
		"Markets": markets,
		"status":status,
		"error":"",
	})
}
func a1_market_info(c *gin.Context) {

	p_market := c.Param("market")
	if len(p_market) > 0 {
		market_aid, err := strconv.ParseInt(p_market,10,64)
		if err == nil {
			p_market,err = augur_srv.storage.Lookup_address(market_aid)
			if err!=nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"status":0,
					"error":fmt.Sprintf("Market with ID=%v not found",market_aid),
				})
				return
			}
		}
	} else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":0,
			"error":fmt.Sprintf("Empty 'market' parameter for market lookup"),
		})
		return
	}
	market_addr,valid:=is_address_valid(c,true,p_market)
	if !valid {
		return
	}
	market_info,err := augur_srv.storage.Get_market_info(market_addr,0,false)
	if err != nil {
		show_market_not_found_error(c,true,&market_addr)
		return
	}
	complete_and_output_market_info(c,true,market_info)
}
func a1_multiple_market_cards(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	markets := make([]InfoMarket,0,64)
	p_market_aid_list := c.Param("market_aid_list")
	aid_list := strings.Split(p_market_aid_list,",")
	for i,aid_str := range aid_list {
		if len(aid_str ) > 0 {
			market_aid, err := strconv.ParseInt(aid_str,10,64)
			if err != nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"Markets": make([]InfoMarket,0,0),
					"status":0,
					"error":fmt.Sprintf("Bad integer for market_aid parameter at position %v : %v",i,err),
				})
				return
			}
			mkt_info,err := augur_srv.storage.Get_market_card_data(market_aid)
			if err!=nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"Markets": make([]InfoMarket,0,0),
					"status":0,
					"error":fmt.Sprintf("Error getting market by ID (id=%v): %v",market_aid,err),
				})
				return
			}
			markets = append(markets,mkt_info)
		} else {
			c.JSON(http.StatusBadRequest,gin.H{
				"Markets": make([]InfoMarket,0,0),
				"status":0,
				"error":fmt.Sprintf("market_aid at position %v is empty",i+1),
			})
			return
		}
	}
	var status int = 0
	var err_str string = ""
	c.JSON(200,gin.H{
		"Markets": markets,
		"status":status,
		"error":err_str,
	})
}
