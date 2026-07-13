package cosmicgame

import (
	"errors"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func (a *API) handleStakingActionRwalkInfo(c *httpx.Context) {

	if !a.dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}

	pActionID := c.Param("action_id")
	var actionID int64
	if len(pActionID) > 0 {
		var success bool
		actionID, success = common.ParseIntFromRemoteOrError(c, JSON, &pActionID)
		if !success {
			return
		}
	}
	actionInfo, err := a.repo.StakeActionRwalkInfo(c.Request.Context(), actionID)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "record not found")
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	var reqStatus int = 1
	var errStr string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                         reqStatus,
		"error":                          errStr,
		"CombinedRWalkStakingRecordInfo": actionInfo,
	})
}
func (a *API) handleStakingActionsRwalkGlobal(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	actions, err := a.repo.GlobalStakingRwalkHistory(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	var reqStatus int = 1
	var errStr string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                    reqStatus,
		"error":                     errStr,
		"Offset":                    offset,
		"Limit":                     limit,
		"GlobalStakingActionsRWalk": actions,
	})
}
func (a *API) handleStakingActionsRwalkByUser(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAid": int64(0), "UserAddr": pUserAddr,
			"Offset": offset, "Limit": limit, "UserStakingActionsRWalk": []interface{}{},
		})
		return
	}
	actions, err := a.repo.StakingActionsRwalkByUser(c.Request.Context(), userAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	var reqStatus int = 1
	var errStr string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                  reqStatus,
		"error":                   errStr,
		"UserAid":                 userAid,
		"UserAddr":                pUserAddr,
		"Offset":                  offset,
		"Limit":                   limit,
		"UserStakingActionsRWalk": actions,
	})
}
func (a *API) handleUserUniqueStakersRwalk(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	uniqueStakers, err := a.repo.UniqueStakersRwalk(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	var reqStatus int = 1
	var errStr string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":             reqStatus,
		"error":              errStr,
		"UniqueStakersRWalk": uniqueStakers,
	})
}
func (a *API) handleStakingRwalkMintsGlobal(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	mints, err := a.repo.StakingRwalkMintsGlobal(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	var reqStatus int = 1
	var errStr string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                   reqStatus,
		"error":                    errStr,
		"StakingRWalkRewardsMints": mints,
	})
}
func (a *API) handleStakingRwalkMintsByUser(c *httpx.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pUserAddr := c.Param("user_addr")

	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		// Address not in DB yet — return 200 with empty list so UI works
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "RWalkStakingRewardMints": []interface{}{},
		})
		return
	}

	mints, err := a.repo.StakingRwalkMintsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	var reqStatus int = 1
	var errStr string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                  reqStatus,
		"error":                   errStr,
		"RWalkStakingRewardMints": mints,
	})
}
func (a *API) handleStakedTokensRwalkByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		// Address not in DB yet (e.g. new wallet) — return 200 with empty list so UI and bidding still work
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": int64(0), "StakedTokensRWalk": []interface{}{},
		})
		return
	}
	tokens, err := a.repo.StakedTokensRwalkByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	var reqStatus int = 1
	var errStr string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"UserAddr":          pUserAddr,
		"UserAid":           userAid,
		"StakedTokensRWalk": tokens,
	})
}
func (a *API) handleStakedTokensRwalkGlobal(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	tokens, err := a.repo.StakedTokensRwalkGlobal(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	var reqStatus int = 1
	var errStr string = ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"StakedTokensRWalk": tokens,
	})
}
