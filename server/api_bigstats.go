/// API v1
package main
import (
	//"fmt"
	"time"

	"net/http"
	"github.com/gin-gonic/gin"

)
func api_stats_main_statistics(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,ini,fin := parse_timeframe_ini_fin(c,HTTP)
	if !success {
		return
	}

	records := augur_srv.db_arbitrum.Bigstats_get_statistics_by_period("st_arb",ini,fin)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Statistics" : records,
		"IniTs": ini,
		"FinTs": fin,
		"IniDate" : time.Unix(int64(ini),0).String(),
		"FinDate":time.Unix(int64(fin),0).String(),
	})
}
