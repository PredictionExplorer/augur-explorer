/// API v1
package main
import (
	//"fmt"
	"time"

	"net/http"
	"github.com/gin-gonic/gin"

	. "github.com/PredictionExplorer/augur-explorer/primitives"
)
func api_stats_main_statistics_arbitrum(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	api_get_network_statistics(c,"st_arb")
}
func api_stats_main_statistics_main_net(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	api_get_network_statistics(c,"st_arb")
}
func api_get_network_statistics(c *gin.Context,schema_name string) {

	var records []BigStatRec

	success,ini,fin := parse_timeframe_ini_fin(c,HTTP)
	if !success {
		return
	}
	if schema_name == "st_arb" {
		if  !augur_srv.arbitrum_initialized() {
			respond_error_json(c,"Database link for Arbitrum wasn't configured")
			return
		}
	} else {
		if schema_name == "st_eth" {
			if  !augur_srv.main_net_initialized() {
				respond_error_json(c,"Database link for MainNet wasn't configured")
				return
			}
		} else {
			respond_error_json(c,"Unknown database schema")
			return
		}
	}
	records = augur_srv.db_arbitrum.Bigstats_get_statistics_by_period(schema_name,ini,fin)
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
func api_stats_get_timeframe_ranges(c *gin.Context) {

	var MainNetTsIni,MainNetTsFin int64
	var ArbTsIni,ArbTsFin int64
	if  augur_srv.arbitrum_initialized() {
		entry := augur_srv.db_arbitrum.Bigstats_get_timeframe_range("st_arb")
		ArbTsIni = entry.TsIni
		ArbTsFin = entry.TsFin
	}
	if augur_srv.main_net_initialized() {
		entry := augur_srv.db_main_net.Bigstats_get_timeframe_range("st_eth")
		MainNetTsIni = entry.TsIni
		MainNetTsFin = entry.TsFin
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"MainNetTsIni":MainNetTsIni,
		"MainNetTsFin":MainNetTsFin,
		"ArbTsIni":ArbTsIni,
		"ArbTsFin":ArbTsFin,
	})

}
