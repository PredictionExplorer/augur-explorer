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
		"GlobalStakingActionsRWalk" : actions,
	})
}
func api_cosmic_game_staking_actions_rwalk_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
    p_user_addr:= c.Param("user_addr")
    if len(p_user_addr) == 0 {
        respond_error_json(c,"'user_addr' parameter is not set")
        return
    }

    user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
    if err != nil {
        respond_error_json(c,"Provided address wasn't found")
        return
    } 
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	actions := arb_storagew.Get_staking_actions_rwalk_by_user(user_aid,offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAid" : user_aid,
		"UserAddr" : p_user_addr,
		"Offset" : offset,
		"Limit" : limit,
		"UserStakingActionsRWalk" : actions,
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
func api_cosmic_game_staking_rwalk_mints_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

    success,offset,limit := parse_offset_limit_params_json(c)
    if !success {
        return
    }
	mints := arb_storagew.Get_staking_rwalk_mints_global(offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"StakingRWalkRewardsMints" : mints,
	})
}
func api_cosmic_game_staking_rwalk_mints_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

    p_user_addr:= c.Param("user_addr")
    if len(p_user_addr) == 0 {
        respond_error_json(c,"'user_addr' parameter is not set")
        return
    }

    user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
    if err != nil {
        respond_error_json(c,"Provided address wasn't found")
        return
    } 

	mints := arb_storagew.Get_staking_rwalk_mints_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RWalkStakingRewardMints" : mints,
	})
}
func api_cosmic_game_staked_tokens_rwalk_by_user(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error_json(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		respond_error_json(c,"Provided address wasn't found")
		return
	}
	tokens := arb_storagew.Get_staked_tokens_rwalk_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"StakedTokensRWalk" : tokens,
	})
}
func api_cosmic_game_staked_tokens_rwalk_global(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	tokens := arb_storagew.Get_staked_tokens_rwalk_global()
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"StakedTokensRWalk" : tokens,
	})
}
