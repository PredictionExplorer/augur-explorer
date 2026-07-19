package cosmicgame

import (
	"errors"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func (a *API) handleStakingActionCstInfo(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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
	actionInfo, err := a.repo.StakeActionCstInfo(c.Request.Context(), actionID)
	if errors.Is(err, store.ErrNotFound) {
		common.RespondErrorJSON(c, "record not found")
		return
	}
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                    reqStatus,
		"error":                     errStr,
		"CombinedStakingRecordInfo": actionInfo,
	})
}

func (a *API) handleStakingActionsCstGlobal(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	actions, err := a.repo.GlobalStakingCstHistory(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"Offset":            offset,
		"Limit":             limit,
		"StakingCSTActions": actions,
	})
}

func (a *API) handleStakingCstActionsByUser(c *httpx.Context) {
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
			"status": 1, "error": "", "Offset": offset, "Limit": limit,
			"UserAddr": pUserAddr, "UserAid": int64(0), "StakingCSTActions": []any{},
		})
		return
	}
	actions, err := a.repo.StakingActionsCstByUser(c.Request.Context(), userAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"Offset":            offset,
		"Limit":             limit,
		"UserAddr":          pUserAddr,
		"UserAid":           userAid,
		"StakingCSTActions": actions,
	})
}

func (a *API) handleUserUniqueStakersCst(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	uniqueStakers, err := a.repo.UniqueStakersCst(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"UniqueStakersCST": uniqueStakers,
	})
}

func (a *API) handleStakingCstRewardsToClaimByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		// Address not in DB yet (e.g. new wallet) — return 200 with empty list so UI works
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": int64(0), "UnclaimedEthDeposits": []any{},
		})
		return
	}
	deposits, err := a.repo.StakingRewardsToBeClaimed(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":               reqStatus,
		"error":                errStr,
		"UserAddr":             pUserAddr,
		"UserAid":              userAid,
		"UnclaimedEthDeposits": deposits,
	})
}

func (a *API) handleStakedTokensCstByUser(c *httpx.Context) {
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
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": int64(0), "StakedTokensCST": []any{},
		})
		return
	}
	tokens, err := a.repo.StakedTokensCstByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":          reqStatus,
		"error":           errStr,
		"UserAddr":        pUserAddr,
		"UserAid":         userAid,
		"StakedTokensCST": tokens,
	})
}

func (a *API) handleStakedTokensCstGlobal(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	tokens, err := a.repo.StakedTokensCstGlobal(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":          reqStatus,
		"error":           errStr,
		"StakedTokensCST": tokens,
	})
}

func (a *API) handleStakingCstRewardsCollectedByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		// Address not in DB yet — return 200 with empty list so UI works
		success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "Offset": offset, "Limit": limit,
			"UserAddr": pUserAddr, "UserAid": int64(0), "CollectedStakingCSTRewards": []any{},
		})
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	actions, err := a.repo.StakingRewardsCollected(c.Request.Context(), userAid, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                     reqStatus,
		"error":                      errStr,
		"Offset":                     offset,
		"Limit":                      limit,
		"UserAddr":                   pUserAddr,
		"UserAid":                    userAid,
		"CollectedStakingCSTRewards": actions,
	})
}

func (a *API) handleStakingCstRewardsByRound(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pRoundNum := c.Param("round_num")
	var roundNum int64
	if len(pRoundNum) > 0 {
		var success bool
		roundNum, success = common.ParseIntFromRemoteOrError(c, JSON, &pRoundNum)
		if !success {
			return
		}
	}

	rewards, err := a.repo.StakingCstRewardsByRound(c.Request.Context(), roundNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":   reqStatus,
		"error":    errStr,
		"RoundNum": roundNum,
		"Rewards":  rewards,
	})
}

func (a *API) handleStakingCstRewardsGlobal(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	rewards, err := a.repo.GlobalStakingRewards(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"StakingCSTRewards": rewards,
	})
}

func (a *API) handleStakingCstMintsGlobal(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	mints, err := a.repo.StakingCstMintsGlobal(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                 reqStatus,
		"error":                  errStr,
		"Offset":                 offset,
		"Limit":                  limit,
		"StakingCSTRewardsMints": mints,
	})
}

func (a *API) handleStakingCstMintsByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pUserAddr := c.Param("user_addr")

	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "CSTStakingRewardMints": []any{},
		})
		return
	}

	mints, err := a.repo.StakingCstMintsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                reqStatus,
		"error":                 errStr,
		"CSTStakingRewardMints": mints,
	})
}

func (a *API) handleStakingCstRewardsActionIDsByDeposit(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	pDepositID := c.Param("deposit_id")
	var depositID int64
	if len(pDepositID) > 0 {
		var success bool
		depositID, success = common.ParseIntFromRemoteOrError(c, JSON, &pDepositID)
		if !success {
			return
		}
	}
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": int64(0),
			"DepositId": depositID, "ActionIdsWithClaimInfo": []any{},
		})
		return
	}
	actionIDs, err := a.repo.ActionIDsForDepositWithClaimInfo(c.Request.Context(), depositID, userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                 reqStatus,
		"error":                  errStr,
		"UserAddr":               pUserAddr,
		"UserAid":                userAid,
		"DepositId":              depositID,
		"ActionIdsWithClaimInfo": actionIDs,
	})
}

func (a *API) handleStakingCstByUserByDepositRewards(c *httpx.Context) {
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
			"status": 1, "error": "", "UserAid": int64(0), "UserAddr": pUserAddr, "RewardsByDeposit": []any{},
		})
		return
	}

	history, err := a.repo.StakingCstUserDepositRewards(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":           reqStatus,
		"error":            errStr,
		"UserAid":          userAid,
		"UserAddr":         pUserAddr,
		"RewardsByDeposit": history,
	})
}

func (a *API) handleStakingCstByUserByTokenRewards(c *httpx.Context) {
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
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": int64(0), "RewardsByToken": []any{},
		})
		return
	}

	rewards, err := a.repo.StakingCstUserTokenRewards(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":         reqStatus,
		"error":          errStr,
		"UserAddr":       pUserAddr,
		"UserAid":        userAid,
		"RewardsByToken": rewards,
	})
}

func (a *API) handleStakingCstByUserByTokenRewardsDetails(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	pTokenID := c.Param("token_id")
	var tokenID int64
	if len(pTokenID) > 0 {
		var success bool
		tokenID, success = common.ParseIntFromRemoteOrError(c, JSON, &pTokenID)
		if !success {
			return
		}
	}
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "UserAddr": pUserAddr, "UserAid": int64(0),
			"TokenId": tokenID, "RewardsByTokenDetails": []any{},
		})
		return
	}

	rewards, err := a.repo.StakingCstUserTokenRewardDetails(c.Request.Context(), userAid, tokenID)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}

	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                reqStatus,
		"error":                 errStr,
		"UserAddr":              pUserAddr,
		"UserAid":               userAid,
		"TokenId":               tokenID,
		"RewardsByTokenDetails": rewards,
	})
}
