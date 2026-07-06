package randomwalk

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
)

// Token list sequential (API)
func apiRwalkTokenListSeq(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs := rwContractAddrs()
	rwalk_aid := addrs.RandomWalkAid
	tokens := rw_storagew.Get_minted_tokens_sequentially(rwalk_aid, 0, 10000000000)
	c.JSON(http.StatusOK, gin.H{
		"status":       1,
		"error":        "",
		"MintedTokens": tokens,
	})
}

// Token list by period (API)
func apiRwalkTokenListPeriod(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs := rwContractAddrs()
	rwalk_aid := addrs.RandomWalkAid
	success, ini, fin := common.ParseTimeframeIniFin(c, JSON)
	if !success {
		return
	}
	tokens := rw_storagew.Get_minted_tokens_by_period(rwalk_aid, ini, fin)
	c.JSON(http.StatusOK, gin.H{
		"status":       1,
		"error":        "",
		"MintedTokens": tokens,
		"InitTs":       ini,
		"FinTs":        fin,
		"RWalkAid":     rwalk_aid,
	})
}

// Token info (API)
func apiRwalkTokenInfo(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	addrs := rwContractAddrs()
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
	token_info, err := rw_storagew.Get_rwalk_token_info(rwalk_aid, token_id)
	if err != nil {
		common.RespondErrorJSON(c, fmt.Sprintf("Error during query execution: %v", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    1,
		"error":     "",
		"TokenInfo": token_info,
	})
}

// Token history (API)
func apiRwalkTokenHistory(c *gin.Context) {
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
	addrs := rwContractAddrs()
	rwalk_aid := addrs.RandomWalkAid
	p_rwalk_addr := addrs.RandomWalk
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	history := rw_storagew.Get_token_full_history(rwalk_aid, token_id, offset, limit)
	c.JSON(http.StatusOK, gin.H{
		"status":       1,
		"error":        "",
		"TokenId":      token_id,
		"TokenHistory": history,
		"RWalkAddr":    p_rwalk_addr,
		"RWalkAid":     rwalk_aid,
	})
}

// Token name history (API)
func apiRwalkTokenNameHistory(c *gin.Context) {
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
	name_changes := rw_storagew.Get_name_changes_for_token(token_id)
	c.JSON(http.StatusOK, gin.H{
		"status":           1,
		"error":            "",
		"TokenNameChanges": name_changes,
	})
}

// Tokens by user (API)
func apiRwalkTokensByUser(c *gin.Context) {
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
			user_aid, err = rw_storagew.S.Nonfatal_lookup_address_id(p_user_aid)
			if err != nil {
				common.RespondErrorJSON(c, "Cant find provided user")
				return
			}
		}
	} else {
		common.RespondErrorJSON(c, "'user_aid' parameter is not set")
		return
	}
	user_addr, err := rw_storagew.S.Lookup_address(user_aid)
	if err != nil {
		common.RespondErrorJSON(c, "Address lookup on user_aid failed")
		return
	}
	user_tokens := rw_storagew.Get_random_walk_tokens_by_user(user_aid)
	c.JSON(http.StatusOK, gin.H{
		"status":     1,
		"error":      "",
		"UserTokens": user_tokens,
		"UserAid":    user_aid,
		"UserAddr":   user_addr,
	})
}
