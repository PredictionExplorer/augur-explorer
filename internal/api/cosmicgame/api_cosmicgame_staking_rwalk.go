package cosmicgame

import (
	"errors"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func api_cosmic_game_staking_action_rwalk_info(c *httpx.Context) {

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
	}
	action_info, err := arbRepo.StakeActionRwalkInfo(c.Request.Context(), action_id)
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
	c.JSON(http.StatusOK, httpx.H{
		"status":                         req_status,
		"error":                          err_str,
		"CombinedRWalkStakingRecordInfo": action_info,
	})
}
func api_cosmic_game_staking_actions_rwalk_global(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	actions, err := arbRepo.GlobalStakingRwalkHistory(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                    req_status,
		"error":                     err_str,
		"Offset":                    offset,
		"Limit":                     limit,
		"GlobalStakingActionsRWalk": actions,
	})
}
func api_cosmic_game_staking_actions_rwalk_by_user(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAid": int64(0), "UserAddr": p_user_addr,
			"Offset": offset, "Limit": limit, "UserStakingActionsRWalk": []interface{}{},
		})
		return
	}
	actions, err := arbRepo.StakingActionsRwalkByUser(c.Request.Context(), user_aid, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                  req_status,
		"error":                   err_str,
		"UserAid":                 user_aid,
		"UserAddr":                p_user_addr,
		"Offset":                  offset,
		"Limit":                   limit,
		"UserStakingActionsRWalk": actions,
	})
}
func api_cosmic_game_user_unique_stakers_rwalk(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	unique_stakers, err := arbRepo.UniqueStakersRwalk(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}

	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":             req_status,
		"error":              err_str,
		"UniqueStakersRWalk": unique_stakers,
	})
}
func api_cosmic_game_staking_rwalk_mints_global(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	mints, err := arbRepo.StakingRwalkMintsGlobal(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                   req_status,
		"error":                    err_str,
		"StakingRWalkRewardsMints": mints,
	})
}
func api_cosmic_game_staking_rwalk_mints_by_user(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_user_addr := c.Param("user_addr")

	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		// Address not in DB yet — return 200 with empty list so UI works
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "RWalkStakingRewardMints": []interface{}{},
		})
		return
	}

	mints, err := arbRepo.StakingRwalkMintsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                  req_status,
		"error":                   err_str,
		"RWalkStakingRewardMints": mints,
	})
}
func api_cosmic_game_staked_tokens_rwalk_by_user(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	user_aid, err := arbStore.LookupAddressID(c.Request.Context(), p_user_addr)
	if err != nil {
		// Address not in DB yet (e.g. new wallet) — return 200 with empty list so UI and bidding still work
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": p_user_addr, "UserAid": int64(0), "StakedTokensRWalk": []interface{}{},
		})
		return
	}
	tokens, err := arbRepo.StakedTokensRwalkByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            req_status,
		"error":             err_str,
		"UserAddr":          p_user_addr,
		"UserAid":           user_aid,
		"StakedTokensRWalk": tokens,
	})
}
func api_cosmic_game_staked_tokens_rwalk_global(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	tokens, err := arbRepo.StakedTokensRwalkGlobal(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            req_status,
		"error":             err_str,
		"StakedTokensRWalk": tokens,
	})
}
