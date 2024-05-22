package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

)
func api_cosmic_game_staking_action_rwalk_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_action_id := c.Param("action_id")
	var action_id int64
	if len(p_action_id) > 0 {
		var success bool
		action_id,success = parse_int_from_remote_or_error(c,HTTP,&p_action_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'action_id' parameter is not set")
		return
	}
	record_found,action_info := arb_storagew.Get_stake_action_rwalk_info(action_id)
	if !record_found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided action_id wasn't found"),
		})
	} else {
		c.HTML(http.StatusOK, "cg_stake_action_info.html", gin.H{
			"CombinedStakingRecordInfo" : action_info,
		})
	}
} 
func api_cosmic_game_staking_actions_rwalk_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	actions := arb_storagew.Get_global_staking_rwalk_history(offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offset" : offset,
		"Limit" : limit,
		"StakingActions" : actions,
	})
}
func api_cosmic_game_user_unique_stakers_rwalk(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	unique_stakers := arb_storagew.Get_unique_stakers_rwalk()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UniqueStakersRWalk" : unique_stakers,
	})
}
