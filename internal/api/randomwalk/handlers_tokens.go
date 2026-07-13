package randomwalk

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Token list sequential (API).
func (a *API) handleTokenListSeq(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	rwalkAid := addrs.RandomWalkAid
	tokens, err := a.repo.MintedTokensSequentially(c.Request.Context(), rwalkAid, 0, 10000000000)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":       1,
		"error":        "",
		"MintedTokens": tokens,
	})
}

// Token list by period (API).
func (a *API) handleTokenListPeriod(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	rwalkAid := addrs.RandomWalkAid
	success, ini, fin := common.ParseTimeframeIniFin(c, JSON)
	if !success {
		return
	}
	tokens, err := a.repo.MintedTokensByPeriod(c.Request.Context(), rwalkAid, ini, fin)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":       1,
		"error":        "",
		"MintedTokens": tokens,
		"InitTs":       ini,
		"FinTs":        fin,
		"RWalkAid":     rwalkAid,
	})
}

// Token info (API).
func (a *API) handleTokenInfo(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	rwalkAid := addrs.RandomWalkAid
	pTokenID := c.Param("token_id")
	var tokenID int64
	if len(pTokenID) > 0 {
		var success bool
		tokenID, success = common.ParseIntFromRemoteOrError(c, HTTP, &pTokenID)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'token_id' parameter is not set")
		return
	}
	tokenInfo, err := a.repo.TokenInfo(c.Request.Context(), rwalkAid, tokenID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			// Byte-identical legacy error text (pinned by the parity golden
			// errors__missing_rw_token).
			common.RespondErrorJSON(c, "Error during query execution: "+legacyNoRowsText)
			return
		}
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":    1,
		"error":     "",
		"TokenInfo": tokenInfo,
	})
}

// Token history (API).
func (a *API) handleTokenHistory(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pTokenID := c.Param("token_id")
	var tokenID int64
	if len(pTokenID) > 0 {
		var success bool
		tokenID, success = common.ParseIntFromRemoteOrError(c, JSON, &pTokenID)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'token_id' parameter is not set")
		return
	}
	addrs, ok := a.rwContractAddrs(c)
	if !ok {
		return
	}
	rwalkAid := addrs.RandomWalkAid
	pRwalkAddr := addrs.RandomWalk
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	history, err := a.repo.TokenFullHistory(c.Request.Context(), rwalkAid, tokenID, offset, limit)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":       1,
		"error":        "",
		"TokenId":      tokenID,
		"TokenHistory": history,
		"RWalkAddr":    pRwalkAddr,
		"RWalkAid":     rwalkAid,
	})
}

// Token name history (API).
func (a *API) handleTokenNameHistory(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pTokenID := c.Param("token_id")
	var tokenID int64
	if len(pTokenID) > 0 {
		var success bool
		tokenID, success = common.ParseIntFromRemoteOrError(c, JSON, &pTokenID)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'token_id' parameter is not set")
		return
	}
	nameChanges, err := a.repo.TokenNameChanges(c.Request.Context(), tokenID)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":           1,
		"error":            "",
		"TokenNameChanges": nameChanges,
	})
}

// Tokens by user (API).
func (a *API) handleTokensByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !a.dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	pUserAid := c.Param("user_aid")
	var userAid int64
	if len(pUserAid) > 0 {
		var err error
		userAid, err = strconv.ParseInt(pUserAid, 10, 64)
		if err != nil {
			if (len(pUserAid) != 40) && (len(pUserAid) != 42) {
				common.RespondErrorJSON(c, "Can't resolve user identifier to valid address ID or address hex")
				return
			}
			userAid, err = a.store.LookupAddressID(c.Request.Context(), pUserAid)
			if err != nil {
				common.RespondErrorJSON(c, "Cant find provided user")
				return
			}
		}
	} else {
		common.RespondErrorJSON(c, "'user_aid' parameter is not set")
		return
	}
	userAddr, err := a.store.AddressByID(c.Request.Context(), userAid)
	if err != nil {
		if !errors.Is(err, store.ErrNotFound) {
			a.respondStoreError(c, err)
			return
		}
		common.RespondErrorJSON(c, "Address lookup on user_aid failed")
		return
	}
	userTokens, err := a.repo.TokensByUser(c.Request.Context(), userAid)
	if err != nil {
		a.respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":     1,
		"error":      "",
		"UserTokens": userTokens,
		"UserAid":    userAid,
		"UserAddr":   userAddr,
	})
}
