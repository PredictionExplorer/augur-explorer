package main
import (
	"net/http"
	"github.com/gin-gonic/gin"

)
func api_cosmic_game_staking_action_cst_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_action_id := c.Param("action_id")
	var action_id int64
	if len(p_action_id) > 0 {
		var success bool
		action_id,success = parse_int_from_remote_or_error(c,JSON,&p_action_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'action_id' parameter is not set")
		return
	}
	record_found,action_info := arb_storagew.Get_stake_action_cst_info(action_id)
	if !record_found {
		respond_error_json(c,"record not found")
	} else {
		var req_status int = 1
		var err_str string = ""
		c.JSON(http.StatusOK, gin.H{
			"status": req_status,
			"error" : err_str,
			"CombinedStakingRecordInfo" : action_info,
		})
	}
} 
func api_cosmic_game_staking_actions_cst_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	actions := arb_storagew.Get_global_staking_cst_history(offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offset" : offset,
		"Limit" : limit,
		"StakingCSTActions" : actions,
	})
}
func api_cosmic_game_staking_cst_actions_by_user(c *gin.Context) {

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
	actions := arb_storagew.Get_staking_actions_cst_by_user(user_aid,offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offset" : offset,
		"Limit" : limit,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"StakingCSTActions" : actions,
	})
}
func api_cosmic_game_user_unique_stakers_cst(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}

	unique_stakers := arb_storagew.Get_unique_stakers_cst()

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UniqueStakersCST" : unique_stakers,
	})
}
func api_cosmic_game_staking_cst_rewards_to_claim_by_user(c *gin.Context) {

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
	deposits := arb_storagew.Get_staking_rewards_to_be_claimed(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"UnclaimedEthDeposits" : deposits,
	})
}
func api_cosmic_game_staked_tokens_cst_by_user(c *gin.Context) {
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
	tokens := arb_storagew.Get_staked_tokens_cst_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"StakedTokensCST" : tokens,
	})
}
func api_cosmic_game_staked_tokens_cst_global(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	tokens := arb_storagew.Get_staked_tokens_cst_global()
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"StakedTokensCST" : tokens,
	})
}
func api_cosmic_game_staking_cst_rewards_collected_by_user(c *gin.Context) {

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
	actions := arb_storagew.Get_staking_rewards_collected(user_aid,offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offset" : offset,
		"Limit" : limit,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"CollectedStakingCSTRewards" : actions,
	})
}
func api_cosmic_game_staking_cst_rewards_by_round(c *gin.Context) {

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

	rewards := arb_storagew.Get_staking_cst_rewards_by_round(round_num)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"RoundNum" : round_num,
		"Rewards" : rewards,
	})
}
func api_cosmic_game_staking_cst_rewards_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	rewards := arb_storagew.Get_global_staking_rewards(offset, limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offset" : offset,
		"Limit" : limit,
		"StakingCSTRewards" : rewards,
	})
}
func api_cosmic_game_staking_cst_mints_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error_json(c,"Database link wasn't configured")
		return
	}
	success,offset,limit := parse_offset_limit_params_json(c)
	if !success {
		return
	}
	mints := arb_storagew.Get_staking_cst_mints_global(offset,limit)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"Offset" : offset,
		"Limit" : limit,
		"StakingCSTRewardsMints" : mints,
	})
}
func api_cosmic_game_staking_cst_mints_by_user(c *gin.Context) {

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

	mints := arb_storagew.Get_staking_cst_mints_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"CSTStakingRewardMints" : mints,
	})
}
func api_cosmic_game_staking_cst_rewards_action_ids_by_deposit(c *gin.Context) {

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
	p_deposit_id := c.Param("deposit_id")
	var deposit_id int64
	if len(p_deposit_id) > 0 {
		var success bool
		deposit_id,success = parse_int_from_remote_or_error(c,JSON,&p_deposit_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'deposit_id' parameter is not set")
		return
	}
	action_ids := arb_storagew.Get_action_ids_for_deposit(deposit_id,user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"DepositId" : deposit_id,
		"ActionIds" : action_ids,
	})
}
func api_cosmic_game_staking_cst_rewards_action_ids_by_deposit_with_claim_info(c *gin.Context) {

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
	p_deposit_id := c.Param("deposit_id")
	var deposit_id int64
	if len(p_deposit_id) > 0 {
		var success bool
		deposit_id,success = parse_int_from_remote_or_error(c,JSON,&p_deposit_id)
		if !success {
			return
		}
	} else {
		respond_error_json(c,"'deposit_id' parameter is not set")
		return
	}
	action_ids := arb_storagew.Get_action_ids_for_deposit_with_claim_info(deposit_id,user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"DepositId" : deposit_id,
		"ActionIdsWithClaimInfo" : action_ids,
	})
}
func api_cosmic_game_staking_cst_history_by_user(c *gin.Context) {

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

	history := arb_storagew.Get_staking_cst_history_by_user(user_aid)
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status": req_status,
		"error" : err_str,
		"UserAid":user_aid,
		"UserAddr":p_user_addr,
		"UserCstStakingHistory" : history,
	})
}
