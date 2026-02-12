package randomwalk

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// User info (API + HTML)
func apiRwalkUserInfo(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of NFT token failed")
		return
	}
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
	user_addr, err := rw_storagew.S.Lookup_address(user_aid)
	if err != nil {
		common.RespondErrorJSON(c, "Address lookup on user_aid failed")
		return
	}
	user_info, dberr := rw_storagew.Get_rwalk_user_info(user_aid, rwalk_aid)
	var dberr_str string
	if dberr != nil {
		dberr_str = dberr.Error()
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

func rwalk_user_info(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondError(c, "Lookup of NFT token failed")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var success bool
		user_aid, success = common.ParseIntFromRemoteOrError(c, HTTP, &p_user_aid)
		if !success {
			return
		}
	} else {
		common.RespondError(c, "'user_aid' parameter is not set")
		return
	}
	user_addr, err := rw_storagew.S.Lookup_address(user_aid)
	if err != nil {
		common.RespondError(c, "Address lookup on user_aid failed")
		return
	}
	user_info, dberr := rw_storagew.Get_rwalk_user_info(user_aid, rwalk_aid)
	var dberr_string string
	if dberr != nil {
		dberr_string = dberr.Error()
	}
	c.HTML(http.StatusOK, "rw_user_info.html", gin.H{
		"UserInfo": user_info,
		"UserAid":  user_aid,
		"UserAddr": user_addr,
		"DBError":  dberr_string,
	})
}
