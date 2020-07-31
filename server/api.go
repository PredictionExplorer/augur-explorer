/// API v1
package main
import (
	"fmt"
	"strconv"

	//"net/http"
	"github.com/gin-gonic/gin"
)

func a1_active_markets(c *gin.Context) {

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
				"error":fmt.Sprintf("Bad offset parameter: ",err),
			})
			return
		}
	}
	ids := augur_srv.storage.Get_active_market_ids(off,DEFAULT_MARKET_ROWS_LIMIT)
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
				"error":fmt.Sprintf("Bad integer for market_aid parameter: ",err),
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
