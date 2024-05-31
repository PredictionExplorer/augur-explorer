package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

)
func cosmic_game_staking_cst_action_info(c *gin.Context) {

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
	record_found,action_info := arb_storagew.Get_stake_action_cst_info(action_id)
	if !record_found {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided action_id wasn't found"),
		})
	} else {
		c.HTML(http.StatusOK, "cg_stake_action_cst_info.html", gin.H{
			"CombinedCSTStakingRecordInfo" : action_info,
		})
	}
} 
func cosmic_game_staking_cst_rewards_by_round(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,HTTP,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'round_num' parameter is not set")
		return
	}

	rewards := arb_storagew.Get_staking_cst_rewards_by_round(round_num)
	c.HTML(http.StatusOK, "cg_staking_cst_deposits_by_round.html", gin.H{
		"RoundNum" : round_num,
		"Rewards" : rewards,
	})
}
func cosmic_game_staking_cst_rewards_global(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	rewards := arb_storagew.Get_global_staking_rewards(0, 1000000)
	c.HTML(http.StatusOK, "cg_staking_cst_rewards_global.html", gin.H{
		"StakingCSTRewards" : rewards,
	})
}
func cosmic_game_staking_cst_rewards_action_ids_by_deposit_with_claim_info(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		respond_error(c,"Provided address wasn't found")
		return
	}
	p_deposit_id := c.Param("deposit_id")
	var deposit_id int64
	if len(p_deposit_id) > 0 {
		var success bool
		deposit_id,success = parse_int_from_remote_or_error(c,HTTP,&p_deposit_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'deposit_id' parameter is not set")
		return
	}
	action_ids := arb_storagew.Get_action_ids_for_deposit_with_claim_info(deposit_id,user_aid)
	c.HTML(http.StatusOK, "cg_action_ids_by_deposit_with_claim_info.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"DepositId" : deposit_id,
		"ActionIdsWithClaimInfo" : action_ids,
	})
}
func cosmic_game_staking_cst_rewards_action_ids_by_deposit(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		respond_error(c,"Provided address wasn't found")
		return
	}
	p_deposit_id := c.Param("deposit_id")
	var deposit_id int64
	if len(p_deposit_id) > 0 {
		var success bool
		deposit_id,success = parse_int_from_remote_or_error(c,HTTP,&p_deposit_id)
		if !success {
			return
		}
	} else {
		respond_error(c,"'deposit_id' parameter is not set")
		return
	}
	action_ids := arb_storagew.Get_action_ids_for_deposit(deposit_id,user_aid)
	c.HTML(http.StatusOK, "cg_action_ids_by_deposit.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"DepositId" : deposit_id,
		"ActionIds" : action_ids,
	})
}
func cosmic_game_staked_tokens_cst_global(c *gin.Context) {
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	tokens := arb_storagew.Get_staked_tokens_cst_global()
	c.HTML(http.StatusOK, "cg_staked_tokens_cst_global.html", gin.H{
		"StakedTokensCST" : tokens,
	})
}
func cosmic_game_staked_tokens_cst_by_user(c *gin.Context) {
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		respond_error(c,"Provided address wasn't found")
		return
	}
	tokens := arb_storagew.Get_staked_tokens_cst_by_user(user_aid)
	c.HTML(http.StatusOK, "cg_staked_cst_tokens_by_user.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"StakedTokensCST" : tokens,
	})
}
func cosmic_game_staking_cst_actions_global(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	actions := arb_storagew.Get_global_staking_cst_history(0 ,100000)
	last_ts := arb_storagew.S.Get_last_block_timestamp()
	c.HTML(http.StatusOK, "cg_staking_cst_actions_global.html", gin.H{
		"StakingCSTActions" : actions,
		"LastTS" : last_ts,
	})
}
func cosmic_game_staking_actions_cst_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	actions := arb_storagew.Get_staking_actions_cst_by_user(user_aid,0 ,100000)
	c.HTML(http.StatusOK, "cg_staking_actions_cst_by_user.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"UserStakingActionsCST" : actions,
	})
}
func cosmic_game_staking_cst_rewards_collected_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		respond_error(c,"Provided address wasn't found")
		return
	}
	actions := arb_storagew.Get_staking_rewards_collected(user_aid,0, 1000000)
	c.HTML(http.StatusOK, "cg_staking_cst_rewards_collected_by_user.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"CollectedStakingCSTRewards" : actions,
	})
}
func cosmic_game_staking_cst_rewards_to_claim_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}

	p_user_addr:= c.Param("user_addr")
	if len(p_user_addr) == 0 {
		respond_error(c,"'user_addr' parameter is not set")
		return
	}
	user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"title": "Error",
			"ErrDescr": fmt.Sprintf("Provided address wasn't found"),
		})
		return
	}
	deposits := arb_storagew.Get_staking_rewards_to_be_claimed(user_aid)
	c.HTML(http.StatusOK, "cg_staking_cst_rewards_to_be_claimed_by_user.html", gin.H{
		"UserAddr" : p_user_addr,
		"UserAid" : user_aid,
		"UnclaimedEthDeposits" : deposits,
	})
}
func cosmic_game_unique_stakers_cst(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	unique_stakers := arb_storagew.Get_unique_stakers_cst()
	c.HTML(http.StatusOK, "cg_unique_stakers_cst.html", gin.H{
		"UniqueStakersCST" : unique_stakers,
	})
}
func cosmic_game_staking_cst_mints_global(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	offset:=int(0);limit:=int(100000)
	mints := arb_storagew.Get_staking_cst_mints_global(offset,limit)
	c.HTML(http.StatusOK, "cg_staking_cst_mints_global.html", gin.H{
		"StakingCSTMints" : mints,
	})
}
func cosmic_game_staking_cst_mints_by_user(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
    p_user_addr:= c.Param("user_addr")
    if len(p_user_addr) == 0 {
        respond_error(c,"'user_addr' parameter is not set")
        return
    }
    user_aid,err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
    if err != nil {
        c.HTML(http.StatusBadRequest, "error.html", gin.H{
            "title": "Error",
            "ErrDescr": fmt.Sprintf("Provided address wasn't found"),
        })
        return
    }

	mints := arb_storagew.Get_staking_cst_mints_by_user(user_aid)
	c.HTML(http.StatusOK, "cg_staking_cst_mints_by_user.html", gin.H{
		"UserAid":user_aid,
		"UserAddr":p_user_addr,
		"CSTStakingRewardMints" : mints,
	})
}
