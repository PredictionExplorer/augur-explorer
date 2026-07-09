package randomwalk

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// User info (API)
func apiRwalkUserInfo(c *gin.Context) {
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
	p_rwalk_addr := addrs.RandomWalk
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid, success = common.ParseIntFromRemoteOrError(c, JSON, &p_user_aid)
		if !success {
			return
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
	user_info, dberr := rwRepo.UserInfo(c.Request.Context(), user_aid, rwalk_aid)
	var dberr_str string
	if dberr != nil {
		// A user without stats rows keeps the HTTP 200 + DBError wire shape
		// (byte-identical legacy text); real failures answer 500.
		if !errors.Is(dberr, store.ErrNotFound) {
			respondStoreError(c, dberr)
			return
		}
		dberr_str = legacyNoRowsText
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    1,
		"error":     "",
		"UserInfo":  user_info,
		"UserAid":   user_aid,
		"UserAddr":  user_addr,
		"RWalkAddr": p_rwalk_addr,
		"RWalkAid":  rwalk_aid,
		"DBError":   dberr_str,
	})
}
