/// API v1
package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func a1_amm_user_swaps(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,p_user)
	if !valid {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	aid,err := augur_srv.db_matic.Nonfatal_lookup_address_id(user_addr)
	if err != nil {
		aid = 0
	}
	total_rows,swaps := augur_srv.db_matic.Get_amm_user_swaps(&amm_constants,aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Swaps" : swaps,
		"TotalRows" : total_rows,
		"User":p_user,
		"UserAid":aid,
	})
}
func a1_amm_user_liquidity(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	if  !augur_srv.matic_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_user := c.Param("user")
	user_addr,valid := is_address_valid(c,false,p_user)
	if !valid {
		return
	}
	success,offset,limit := parse_offset_limit_params(c)
	if !success {
		return
	}
	aid,err := augur_srv.db_matic.Nonfatal_lookup_address_id(user_addr)
	if err != nil {
		aid = 0
	}
	total_rows,liquidity := augur_srv.db_matic.Get_amm_user_liquidity(&amm_constants,aid,offset,limit)

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Liquidity" : liquidity,
		"TotalRows" : total_rows,
		"User": p_user,
		"UserAid": aid,
	})
}
