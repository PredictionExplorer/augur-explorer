package main
import (
	"fmt"
	"time"
	"os"
	"strconv"
	"encoding/csv"
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
func rwalk_daily_stats_arbitrum(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	success,ini,fin := parse_timeframe_ini_fin(c,HTTP)
	if !success {
		return
	}

	c.HTML(http.StatusOK, "rw_current_offers.html", gin.H{
		"Offers" : offers,
		"RWalkAid": rwalk_aid,
		"RWalkAddr" : p_rwalk_addr,
		"MarketAid": market_aid,
		"MarketAddr" : p_market_addr,
	})
}
