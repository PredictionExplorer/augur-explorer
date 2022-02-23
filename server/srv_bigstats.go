package main
import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"

)
func stats_index_page(c *gin.Context) {

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
	c.HTML(http.StatusOK, "bigstats_index.html", gin.H{
		"MainNetTsIni":MainNetTsIni,
		"MainNetTsFin":MainNetTsFin,
		"ArbTsIni":ArbTsIni,
		"ArbTsFin":ArbTsFin,
	})
}
func stats_main_statistics_main_net(c *gin.Context) {

	stats_get_network_statistics(c,"st_eth")
}
func stats_main_statistics_arbitrum(c *gin.Context) {

	stats_get_network_statistics(c,"st_arb")
}
func stats_get_network_statistics(c *gin.Context,schema_name string) {
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	if schema_name == "st_arb" {
		if  !augur_srv.arbitrum_initialized() {
			respond_error(c,"Database link for Arbitrum wasn't configured")
			return
		}
	} else {
		if schema_name == "st_eth" {
			if  !augur_srv.main_net_initialized() {
				respond_error(c,"Database link for MainNet wasn't configured")
				return
			}
		} else {
			respond_error(c,"Unknown database schema")
			return
		}
	}

	success,ini,fin := parse_timeframe_ini_fin(c,HTTP)
	if !success {
		return
	}

	fmt.Printf("schema_name=%v ini=%v fin=%v\n",schema_name,ini,fin)
	records := augur_srv.db_arbitrum.Bigstats_get_statistics_by_period(schema_name,"ethprice",ini,fin)
	c.HTML(http.StatusOK, "bigstats_statistics.html", gin.H{
		"Statistics" : records,
		"IniTs": ini,
		"FinTs": fin,
		"IniDate" : time.Unix(int64(ini),0).String(),
		"FinDate":time.Unix(int64(fin),0).String(),
	})
}
func stats_get_timeframe_range(c *gin.Context) {

}
