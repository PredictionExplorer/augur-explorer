package cosmicgame

import (
	"errors"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func (a *API) handleDonationsCgSimpleList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	donations, err := a.repo.SimpleEthDonations(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"DirectCGDonations": donations,
		"Offset":            offset,
		"Limit":             limit,
	})
}

func (a *API) handleDonationsCgSimpleByRound(c *httpx.Context) {
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

	donations, err := a.repo.SimpleEthDonationsByRound(c.Request.Context(), roundNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"DirectCGDonations": donations,
	})
}

func (a *API) handleDonationsCgWithInfoList(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	donations, err := a.repo.EthDonationsWithInfo(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"DirectCGDonations": donations,
		"Offset":            offset,
		"Limit":             limit,
	})
}

func (a *API) handleDonationsCgWithInfoByRound(c *httpx.Context) {
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

	donations, err := a.repo.EthDonationsWithInfoByRound(c.Request.Context(), roundNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":            reqStatus,
		"error":             errStr,
		"DirectCGDonations": donations,
		"RoundNum":          roundNum,
	})
}

func (a *API) handleDonationsCgWithInfoRecordInfo(c *httpx.Context) {
	if !a.dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	pRecordID := c.Param("record_id")
	var recordID int64
	if len(pRecordID) > 0 {
		var success bool
		recordID, success = common.ParseIntFromRemoteOrError(c, HTTP, &pRecordID)
		if !success {
			return
		}
	}
	recordInfo, err := a.repo.EthDonationWithInfoRecord(c.Request.Context(), recordID)
	if err != nil && !errors.Is(err, store.ErrNotFound) {
		a.respondStoreError(c, err)
		return
	}
	// The legacy layer served the zero-value record for unknown ids (the
	// parity goldens pin that shape); ErrNotFound keeps exactly that.
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":      reqStatus,
		"error":       errStr,
		"ETHDonation": recordInfo,
		"RecordId":    recordID,
	})
}

func (a *API) handleDonationsByUser(c *httpx.Context) {
	if !a.dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "CombinedDonationRecords": []any{},
			"UserAddr": pUserAddr, "UserAid": int64(0),
		})
		return
	}

	donations, err := a.repo.EthDonationsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                  reqStatus,
		"error":                   errStr,
		"CombinedDonationRecords": donations,
		"UserAddr":                pUserAddr,
		"UserAid":                 userAid,
	})
}

func (a *API) handleDonationsCgBothByRound(c *httpx.Context) {
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

	donations, err := a.repo.EthDonationsByRound(c.Request.Context(), roundNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":              reqStatus,
		"error":               errStr,
		"CosmicGameDonations": donations,
		"RoundNum":            roundNum,
	})
}

func (a *API) handleDonationsCgBothAll(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	donations, err := a.repo.EthDonations(c.Request.Context())
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":              reqStatus,
		"error":               errStr,
		"CosmicGameDonations": donations,
	})
}

func (a *API) handleDonationsErc20ByRoundDetailed(c *httpx.Context) {
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

	donations, err := a.repo.ERC20DonationsByRoundDetailed(c.Request.Context(), roundNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                        reqStatus,
		"error":                         errStr,
		"DonationsERC20ByRoundDetailed": donations,
		"RoundNum":                      roundNum,
	})
}

func (a *API) handleDonationsErc20ByRoundAll(c *httpx.Context) {
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

	donations, err := a.repo.ERC20DonationsByRoundAll(c.Request.Context(), roundNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                   reqStatus,
		"error":                    errStr,
		"DonationsERC20ByRoundAll": donations,
		"RoundNum":                 roundNum,
	})
}

func (a *API) handleDonationsErc20ByRoundSummarized(c *httpx.Context) {
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

	donations, err := a.repo.ERC20DonationsByRoundSummarized(c.Request.Context(), roundNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                          reqStatus,
		"error":                           errStr,
		"DonationsERC20ByRoundSummarized": donations,
		"RoundNum":                        roundNum,
	})
}

func (a *API) handleDonationsErc20ByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "DonatedPrizesERC20ByWinner": []any{},
			"UserAddr": pUserAddr, "UserAid": int64(0),
		})
		return
	}

	donatedPrizes, err := a.repo.ERC20DonatedPrizesByWinner(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                     reqStatus,
		"error":                      errStr,
		"DonatedPrizesERC20ByWinner": donatedPrizes,
		"UserAddr":                   pUserAddr,
		"UserAid":                    userAid,
	})
}

func (a *API) handleDonationsErc20Global(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	donations, err := a.repo.ERC20Donations(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":         reqStatus,
		"error":          errStr,
		"DonationsERC20": donations,
		"Offset":         offset,
		"Limit":          limit,
	})
}

func (a *API) handleDonatedErc20Info(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	pRecordID := c.Param("record_id")
	var recordID int64
	if len(pRecordID) > 0 {
		var success bool
		recordID, success = common.ParseIntFromRemoteOrError(c, JSON, &pRecordID)
		if !success {
			return
		}
	}
	nftdonation, err := a.repo.ERC20DonationInfo(c.Request.Context(), recordID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "Record not found")
			return
		}
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":        1,
		"error":         "",
		"ERC20Donation": nftdonation,
	})
}

func (a *API) handleDonationsErc20DonatedByUser(c *httpx.Context) {
	// DONOR PERSPECTIVE: Returns ERC20 tokens this user DONATED (not won)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "DonationsERC20ByDonor": []any{},
			"UserAddr": pUserAddr, "UserAid": int64(0),
		})
		return
	}

	donations, err := a.repo.ERC20DonationsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":                reqStatus,
		"error":                 errStr,
		"DonationsERC20ByDonor": donations,
		"UserAddr":              pUserAddr,
		"UserAid":               userAid,
	})
}

func (a *API) handleErc20ClaimsGlobal(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	pOffset := c.Param("offset")
	pLimit := c.Param("limit")

	if len(pOffset) == 0 || len(pLimit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	claims, err := a.repo.ERC20DonationClaims(c.Request.Context(), offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":      reqStatus,
		"error":       errStr,
		"ERC20Claims": claims,
		"Offset":      offset,
		"Limit":       limit,
	})
}

func (a *API) handleErc20ClaimsByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAddr := c.Param("user_addr")
	userAid, err := a.store.LookupAddressID(c.Request.Context(), pUserAddr)
	if err != nil {
		c.JSON(http.StatusOK, httpx.H{
			"status": 1, "error": "", "ERC20ClaimsByWinner": []any{},
			"UserAddr": pUserAddr, "UserAid": int64(0),
		})
		return
	}

	claims, err := a.repo.ERC20DonationClaimsByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":              reqStatus,
		"error":               errStr,
		"ERC20ClaimsByWinner": claims,
		"UserAddr":            pUserAddr,
		"UserAid":             userAid,
	})
}

func (a *API) handleErc20ClaimsByRound(c *httpx.Context) {
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

	claims, err := a.repo.ERC20DonationClaimsByRound(c.Request.Context(), roundNum)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	reqStatus := 1
	errStr := ""
	c.JSON(http.StatusOK, httpx.H{
		"status":             reqStatus,
		"error":              errStr,
		"ERC20ClaimsByRound": claims,
		"RoundNum":           roundNum,
	})
}
