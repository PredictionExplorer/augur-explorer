package cosmicgame

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func api_cosmic_game_staking_action_cst_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}

	p_action_id := c.Param("action_id")
	var action_id int64
	if len(p_action_id) > 0 {
		var success bool
		action_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_action_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'action_id' parameter is not set")
		return
	}
	action_info, err := arbRepo.StakeActionCstInfo(c.Request.Context(), action_id)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "record not found")
		return
	}
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                    req_status,
		"error":                     err_str,
		"CombinedStakingRecordInfo": action_info,
	})
}
func api_cosmic_game_staking_actions_cst_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	actions, err := arbRepo.GlobalStakingCstHistory(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"Offset":            offset,
		"Limit":             limit,
		"StakingCSTActions": actions,
	})
}
func api_cosmic_game_staking_cst_actions_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "Offset": offset, "Limit": limit,
			"UserAddr": p_user_addr, "UserAid": int64(0), "StakingCSTActions": []interface{}{},
		})
		return
	}
	actions, err := arbRepo.StakingActionsCstByUser(c.Request.Context(), user_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"Offset":            offset,
		"Limit":             limit,
		"UserAddr":          p_user_addr,
		"UserAid":           user_aid,
		"StakingCSTActions": actions,
	})
}
func api_cosmic_game_user_unique_stakers_cst(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	unique_stakers, err := arbRepo.UniqueStakersCst(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"UniqueStakersCST": unique_stakers,
	})
}
func api_cosmic_game_staking_cst_rewards_to_claim_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		// Address not in DB yet (e.g. new wallet) — return 200 with empty list so UI works
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": int64(0), "UnclaimedEthDeposits": []interface{}{},
		})
		return
	}
	deposits, err := arbRepo.StakingRewardsToBeClaimed(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":               req_status,
		"error":                err_str,
		"UserAddr":             p_user_addr,
		"UserAid":              user_aid,
		"UnclaimedEthDeposits": deposits,
	})
}
func api_cosmic_game_staked_tokens_cst_by_user(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		// Address not in DB yet (e.g. new wallet) — return 200 with empty list so UI and bidding still work
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": int64(0), "StakedTokensCST": []interface{}{},
		})
		return
	}
	tokens, err := arbRepo.StakedTokensCstByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":          req_status,
		"error":           err_str,
		"UserAddr":        p_user_addr,
		"UserAid":         user_aid,
		"StakedTokensCST": tokens,
	})
}
func api_cosmic_game_staked_tokens_cst_global(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	tokens, err := arbRepo.StakedTokensCstGlobal(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":          req_status,
		"error":           err_str,
		"StakedTokensCST": tokens,
	})
}
func api_cosmic_game_staking_cst_rewards_collected_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		// Address not in DB yet — return 200 with empty list so UI works
		success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "Offset": offset, "Limit": limit,
			"UserAddr": p_user_addr, "UserAid": int64(0), "CollectedStakingCSTRewards": []interface{}{},
		})
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	actions, err := arbRepo.StakingRewardsCollected(c.Request.Context(), user_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                     req_status,
		"error":                      err_str,
		"Offset":                     offset,
		"Limit":                      limit,
		"UserAddr":                   p_user_addr,
		"UserAid":                    user_aid,
		"CollectedStakingCSTRewards": actions,
	})
}
func api_cosmic_game_staking_cst_rewards_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	rewards, err := arbRepo.StakingCstRewardsByRound(c.Request.Context(), round_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":   req_status,
		"error":    err_str,
		"RoundNum": round_num,
		"Rewards":  rewards,
	})
}
func api_cosmic_game_staking_cst_rewards_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	rewards, err := arbRepo.GlobalStakingRewards(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"StakingCSTRewards": rewards,
	})
}
func api_cosmic_game_staking_cst_mints_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	mints, err := arbRepo.StakingCstMintsGlobal(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                 req_status,
		"error":                  err_str,
		"Offset":                 offset,
		"Limit":                  limit,
		"StakingCSTRewardsMints": mints,
	})
}
func api_cosmic_game_staking_cst_mints_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}

	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "CSTStakingRewardMints": []interface{}{},
		})
		return
	}

	mints, err := arbRepo.StakingCstMintsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                req_status,
		"error":                 err_str,
		"CSTStakingRewardMints": mints,
	})
}
func api_cosmic_game_staking_cst_rewards_action_ids_by_deposit(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	p_deposit_id := c.Param("deposit_id")
	var deposit_id int64
	if len(p_deposit_id) > 0 {
		var success bool
		deposit_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_deposit_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'deposit_id' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": int64(0),
			"DepositId": deposit_id, "ActionIdsWithClaimInfo": []interface{}{},
		})
		return
	}
	action_ids, err := arbRepo.ActionIDsForDepositWithClaimInfo(c.Request.Context(), deposit_id, user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                 req_status,
		"error":                  err_str,
		"UserAddr":               p_user_addr,
		"UserAid":                user_aid,
		"DepositId":              deposit_id,
		"ActionIdsWithClaimInfo": action_ids,
	})
}
func api_cosmic_game_staking_cst_by_user_by_deposit_rewards(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		// Address not in DB yet — return 200 with empty list so UI works
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAid": int64(0), "UserAddr": p_user_addr, "RewardsByDeposit": []interface{}{},
		})
		return
	}

	history, err := arbRepo.StakingCstUserDepositRewards(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":           req_status,
		"error":            err_str,
		"UserAid":          user_aid,
		"UserAddr":         p_user_addr,
		"RewardsByDeposit": history,
	})
}
func api_cosmic_game_staking_cst_by_user_by_token_rewards(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		// Address not in DB yet — return 200 with empty list so UI works
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": int64(0), "RewardsByToken": []interface{}{},
		})
		return
	}

	rewards, err := arbRepo.StakingCstUserTokenRewards(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":         req_status,
		"error":          err_str,
		"UserAddr":       p_user_addr,
		"UserAid":        user_aid,
		"RewardsByToken": rewards,
	})
}
func api_cosmic_game_staking_cst_by_user_by_token_rewards_details(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'token_id' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": int64(0),
			"TokenId": token_id, "RewardsByTokenDetails": []interface{}{},
		})
		return
	}

	rewards, err := arbRepo.StakingCstUserTokenRewardDetails(c.Request.Context(), user_aid, token_id)
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                req_status,
		"error":                 err_str,
		"UserAddr":              p_user_addr,
		"UserAid":               user_aid,
		"TokenId":               token_id,
		"RewardsByTokenDetails": rewards,
	})
}
