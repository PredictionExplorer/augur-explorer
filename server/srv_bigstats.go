package main
import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"

)
func stats_index_page(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	caddrs := augur_srv.db_arbitrum.Get_randomwalk_contract_addresses()
	top5tokens := augur_srv.db_arbitrum.Get_top5_traded_tokens()
	c.HTML(http.StatusOK, "rw_index.html", gin.H{
		"ContractAddresses":caddrs,
		"Top5Tokens":top5tokens,
	})
}
func stats_main_statistics(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,ini,fin := parse_timeframe_ini_fin(c,HTTP)
	if !success {
		return
	}

	records := augur_srv.db_arbitrum.Bigstats_get_statistics_by_period("st_arb",ini,fin)
	c.HTML(http.StatusOK, "bigstats_statistics.html", gin.H{
		"Statistics" : records,
		"IniTs": ini,
		"FinTs": fin,
		"IniDate" : time.Unix(int64(ini),0).String(),
		"FinDate":time.Unix(int64(fin),0).String(),
	})
}
