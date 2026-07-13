package randomwalk

import (
	"errors"
	"net/http"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// User info (API).
func (a *API) handleUserInfo(c *httpx.Context) {
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
	pRwalkAddr := addrs.RandomWalk
	pUserAid := c.Param("user_aid")
	var userAid int64
	if len(pUserAid) > 0 {
		var success bool
		userAid, success = common.ParseIntFromRemoteOrError(c, JSON, &pUserAid)
		if !success {
			return
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
	userInfo, dberr := a.repo.UserInfo(c.Request.Context(), userAid, rwalkAid)
	var dberrStr string
	if dberr != nil {
		// A user without stats rows keeps the HTTP 200 + DBError wire shape
		// (byte-identical legacy text); real failures answer 500.
		if !errors.Is(dberr, store.ErrNotFound) {
			a.respondStoreError(c, dberr)
			return
		}
		dberrStr = legacyNoRowsText
	}
	c.JSON(http.StatusOK, httpx.H{
		"status":    1,
		"error":     "",
		"UserInfo":  userInfo,
		"UserAid":   userAid,
		"UserAddr":  userAddr,
		"RWalkAddr": pRwalkAddr,
		"RWalkAid":  rwalkAid,
		"DBError":   dberrStr,
	})
}
