package randomwalk

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Token list sequential (API)
func apiRwalkTokenListSeq(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	rwalk_aid := addrs.RandomWalkAid
	tokens, err := rwRepo.MintedTokensSequentially(c.Request.Context(), rwalk_aid, 0, 10000000000)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":       1,
		"error":        "",
		"MintedTokens": tokens,
	})
}

// Token list by period (API)
func apiRwalkTokenListPeriod(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	rwalk_aid := addrs.RandomWalkAid
	success, ini, fin := common.ParseTimeframeIniFin(c, JSON)
	if !success {
		return
	}
	tokens, err := rwRepo.MintedTokensByPeriod(c.Request.Context(), rwalk_aid, ini, fin)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":       1,
		"error":        "",
		"MintedTokens": tokens,
		"InitTs":       ini,
		"FinTs":        fin,
		"RWalkAid":     rwalk_aid,
	})
}

// Token info (API)
func apiRwalkTokenInfo(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	rwalk_aid := addrs.RandomWalkAid
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, HTTP, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'token_id' parameter is not set")
		return
	}
	token_info, err := rwRepo.TokenInfo(c.Request.Context(), rwalk_aid, token_id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			// Byte-identical legacy error text (pinned by the parity golden
			// errors__missing_rw_token).
			common.RespondErrorJSON(c, "Error during query execution: "+legacyNoRowsText)
			return
		}
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":    1,
		"error":     "",
		"TokenInfo": token_info,
	})
}

// Token history (API)
func apiRwalkTokenHistory(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
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
	addrs, ok := rwContractAddrs(c)
	if !ok {
		return
	}
	rwalk_aid := addrs.RandomWalkAid
	p_rwalk_addr := addrs.RandomWalk
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	history, err := rwRepo.TokenFullHistory(c.Request.Context(), rwalk_aid, token_id, offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":       1,
		"error":        "",
		"TokenId":      token_id,
		"TokenHistory": history,
		"RWalkAddr":    p_rwalk_addr,
		"RWalkAid":     rwalk_aid,
	})
}

// Token name history (API)
func apiRwalkTokenNameHistory(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
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
	name_changes, err := rwRepo.TokenNameChanges(c.Request.Context(), token_id)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":           1,
		"error":            "",
		"TokenNameChanges": name_changes,
	})
}

// Tokens by user (API)
func apiRwalkTokensByUser(c *httpx.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var err error
		user_aid, err = strconv.ParseInt(p_user_aid, 10, 64)
		if err != nil {
			if (len(p_user_aid) != 40) && (len(p_user_aid) != 42) {
				common.RespondErrorJSON(c, "Can't resolve user identifier to valid address ID or address hex")
				return
			}
			user_aid, err = rwStore.LookupAddressID(c.Request.Context(), p_user_aid)
			if err != nil {
				common.RespondErrorJSON(c, "Cant find provided user")
				return
			}
		}
	} else {
		common.RespondErrorJSON(c, "'user_aid' parameter is not set")
		return
	}
	user_addr, err := rwStore.AddressByID(c.Request.Context(), user_aid)
	if err != nil {
		common.RespondErrorJSON(c, "Address lookup on user_aid failed")
		return
	}
	user_tokens, err := rwRepo.TokensByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":     1,
		"error":      "",
		"UserTokens": user_tokens,
		"UserAid":    user_aid,
		"UserAddr":   user_addr,
	})
}
