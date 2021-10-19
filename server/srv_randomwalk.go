package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func rw_current_offers(c *gin.Context) {

	p_market_id := c.Param("market_id")
	var market_id int64
	if len(p_market_id) > 0 {
		var success bool
		market_id,success = parse_int_from_remote_or_error(c,false,&p_market_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'market_id' parameter is not set")
		return
	}

	fpmm_aid := augur_srv.db_matic.Get_fpmm_contract_aid(market_id)
	if fpmm_aid == 0 {
		respond_error(c,"Polymarket with this ID wasn't found")
		return
	}

	market_info,err := augur_srv.db_matic.Get_poly_market_info(market_id)
	if err != nil {
		respond_error(c,"Market not found")
		return
	}
	operations := augur_srv.db_matic.Get_polymarkets_buysell_operations(fpmm_aid,0,1000000)

	var js_outcomes_history JSOutcomes
	for outc:=0; outc<int(market_info.OutcomeSlotCount); outc++ {
		prices:= augur_srv.db_matic.Get_poly_market_outcome_price_history(fpmm_aid,int32(outc))
		js_prices := build_js_polymarkets_outcome_price_history(&prices)
		js_outcomes_history.OutcomesDataJS  = append(js_outcomes_history.OutcomesDataJS,js_prices)
	}
	prices:= augur_srv.db_matic.Get_poly_market_outcome_price_history(fpmm_aid,0)
	price0 := build_js_polymarkets_outcome_price_history(&prices)
	prices= augur_srv.db_matic.Get_poly_market_outcome_price_history(fpmm_aid,1)
	price1 := build_js_polymarkets_outcome_price_history(&prices)

	c.HTML(http.StatusOK, "buysell_operations.html", gin.H{
		"BuySellOperations" : operations,
		"MarketId" : market_id,
		"ContractAid" : fpmm_aid,
		"Prices" : js_outcomes_history,
		"Price0" : price0,
		"Price1" : price1,
	})
}

