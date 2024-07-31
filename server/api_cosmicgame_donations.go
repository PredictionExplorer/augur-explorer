package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func api_cosmic_game_donations_cg_simple_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_simple_list(offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"DirectCGDonations" : donations,
		"Offset": offset,
		"Limit": limit,
	})
}
func api_cosmic_game_donations_cg_simple_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,JSON,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'round_num' parameter is not set")
		return
	}

	donations := arb_storagew.Get_donations_to_cosmic_game_simple_by_round(round_num)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"DirectCGDonations" : donations,
	})
}
func api_cosmic_game_donations_cg_with_info_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}

	donations := arb_storagew.Get_donations_to_cosmic_game_with_info_simple_list(offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"DirectCGDonations" : donations,
		"Offset": offset,
		"Limit": limit,
	})
}
func api_cosmic_game_donations_cg_with_info_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,JSON,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'round_num' parameter is not set")
		return
	}

	donations := arb_storagew.Get_donations_to_cosmic_game_with_info_by_round(round_num)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"DirectCGDonations" : donations,
		"RoundNum":round_num,
	})
}
