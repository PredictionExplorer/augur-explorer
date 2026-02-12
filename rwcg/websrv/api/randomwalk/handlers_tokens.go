package randomwalk

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/rwcg/websrv/api/common"
)

// Token list sequential (API + HTML)
func apiRwalkTokenListSeq(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "NTF address wasn't found in the 'address' table")
		return
	}
	tokens := rw_storagew.Get_minted_tokens_sequentially(rwalk_aid, 0, 10000000000)
	c.JSON(http.StatusOK, gin.H{
		"status":       1,
		"error":        "",
		"MintedTokens": tokens,
	})
}

func rwalk_token_list_seq(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondError(c, "NTF address wasn't found in the 'address' table")
		return
	}
	tokens := rw_storagew.Get_minted_tokens_sequentially(rwalk_aid, 0, 10000000000)
	fin_ts := int(time.Now().Unix())
	interval := int(2 * 24 * 60 * 60)
	init_ts := fin_ts - interval
	c.HTML(http.StatusOK, "rw_tokens_minted.html", gin.H{
		"MintedTokens": tokens,
		"RWalkAddr":    p_rwalk_addr,
		"RWalkAid":     rwalk_aid,
		"InitTs":       init_ts,
		"FinTs":        fin_ts,
		"Interval":     interval,
	})
}

// Token list by period (API + HTML)
func apiRwalkTokenListPeriod(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "NTF address wasn't found in the 'address' table")
		return
	}
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

func rwalk_token_list_period(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondError(c, "NTF address wasn't found in the 'address' table")
		return
	}
	success, ini, fin := common.ParseTimeframeIniFin(c, HTTP)
	if !success {
		return
	}
	tokens := rw_storagew.Get_minted_tokens_by_period(rwalk_aid, ini, fin)
	c.HTML(http.StatusOK, "rw_tokens_minted_period.html", gin.H{
		"MintedTokens": tokens,
		"RWalkAddr":    p_rwalk_addr,
		"RWalkAid":     rwalk_aid,
		"InitTs":       ini,
		"FinTs":        fin,
		"InitDate":     time.Unix(int64(ini), 0).String(),
		"FinDate":      time.Unix(int64(fin), 0).String(),
	})
}

// Token info (API + HTML)
func apiRwalkTokenInfo(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of NFT token address in the Db has failed")
		return
	}
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

func rwalk_token_info(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondError(c, "Lookup of NFT token address in the Db has failed")
		return
	}
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, HTTP, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondError(c, "'token_id' parameter is not set")
		return
	}
	token_info, err := rw_storagew.Get_rwalk_token_info(rwalk_aid, token_id)
	if err != nil {
		common.RespondError(c, fmt.Sprintf("Error during query execution: %v", err))
		return
	}
	c.HTML(http.StatusOK, "rw_token_info.html", gin.H{
		"TokenInfo": token_info,
		"RWalkAddr": p_rwalk_addr,
		"RWalkAid":  rwalk_aid,
	})
}

// Token history (API + HTML)
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
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondErrorJSON(c, "Lookup of 'rwalk_addr' failed, address doesn't exist")
		return
	}
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

func rwalk_token_history(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, HTTP, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondError(c, "'token_id' parameter is not set")
		return
	}
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondError(c, "Lookup of 'rwalk_addr' failed, address doesn't exist")
		return
	}
	offset := int(0)
	limit := int(100000)
	history := rw_storagew.Get_token_full_history(rwalk_aid, token_id, offset, limit)
	token_info, err := rw_storagew.Get_rwalk_token_info(rwalk_aid, token_id)
	if err != nil {
		fmt.Printf("Error getting token info for token_id=%v, rwalk_aid=%v : %v\n", token_id, rwalk_aid, err)
	}
	c.HTML(http.StatusOK, "rw_token_history.html", gin.H{
		"TokenId":      token_id,
		"TokenHistory": history,
		"RWalkAddr":    p_rwalk_addr,
		"RWalkAid":     rwalk_aid,
		"TokenInfo":    token_info,
	})
}

// Token name history (API + HTML)
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

func rwalk_token_name_history(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_token_id := c.Param("token_id")
	var token_id int64
	if len(p_token_id) > 0 {
		var success bool
		token_id, success = common.ParseIntFromRemoteOrError(c, HTTP, &p_token_id)
		if !success {
			return
		}
	} else {
		common.RespondError(c, "'token_id' parameter is not set")
		return
	}
	name_changes := rw_storagew.Get_name_changes_for_token(token_id)
	c.HTML(http.StatusOK, "rw_token_names.html", gin.H{
		"TokenNameChanges": name_changes,
	})
}

// Tokens by user (API + HTML)
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

func rwalk_tokens_by_user(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_user_aid := c.Param("user_aid")
	var user_aid int64
	if len(p_user_aid) > 0 {
		var err error
		user_aid, err = strconv.ParseInt(p_user_aid, 10, 64)
		if err != nil {
			if (len(p_user_aid) != 40) && (len(p_user_aid) != 42) {
				common.RespondError(c, "Can't resolve user identifier to valid address ID or address hex")
				return
			}
			user_aid, err = rw_storagew.S.Nonfatal_lookup_address_id(p_user_aid)
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
	user_tokens := rw_storagew.Get_random_walk_tokens_by_user(user_aid)
	c.HTML(http.StatusOK, "rw_tokens_by_user.html", gin.H{
		"UserTokens": user_tokens,
		"UserAid":    user_aid,
		"UserAddr":   user_addr,
	})
}

// Token CSV export (HTML only)
func rwalk_token_csv_export(c *gin.Context) {
	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	c.Writer.Header().Set("Cache-Control", "must-revalidate")
	c.Writer.Header().Set("Pragma", "must-revalidate")
	c.Writer.Header().Set("Content-type", "application/vnd.ms-excel")
	c.Writer.Header().Set("Content-disposition", "attachment; filename=mints.csv")
	p_rwalk_addr := c.Param("rwalk_addr")
	rwalk_aid, err := rw_storagew.S.Nonfatal_lookup_address_id(p_rwalk_addr)
	if err != nil {
		common.RespondError(c, "NTF address wasn't found in the 'address' table")
		return
	}
	data := rw_storagew.Get_minted_tokens_for_CSV(rwalk_aid)
	header := []string{
		"BlockNum", "TimeStamp", "DateTime", "ContractAddr", "TokenId",
		"MinterAddr", "SeedHex", "SeedNum", "PriceMinted", "TxHash",
		"NumTrades", "TotalVolume", "LastPrice", "TokenName", "CurOwner",
	}
	fname := "/tmp/mints.csv"
	f, err := os.Create(fname)
	if err != nil {
		common.RespondError(c, fmt.Sprintf("Cant create file: %v\n", err.Error()))
		return
	}
	w := csv.NewWriter(f)
	if err = w.Write(header); err != nil {
		common.RespondError(c, fmt.Sprintf("Error during header write to csv: %v\n", err.Error()))
		f.Close()
		return
	}
	for i := 0; i < len(data); i++ {
		rec := &data[i]
		row := []string{
			fmt.Sprintf("%d", rec.BlockNum),
			fmt.Sprintf("%d", rec.TimeStamp),
			rec.DateTime,
			rec.ContractAddr,
			fmt.Sprintf("%v", rec.TokenId),
			rec.MinterAddr,
			rec.Seed,
			fmt.Sprintf("%s", rec.SeedNum),
			fmt.Sprintf("%f", rec.Price),
			rec.TxHash,
			fmt.Sprintf("%d", rec.NumTrades),
			fmt.Sprintf("%f", rec.TotalVolume),
			fmt.Sprintf("%f", rec.LastPrice),
			rec.LastName,
			rec.LastOwner,
		}
		if err = w.Write(row); err != nil {
			common.RespondError(c, fmt.Sprintf("Error during write to csv: %v\n", err.Error()))
			f.Close()
			return
		}
	}
	w.Flush()
	f.Close()
	c.File(fname)
}
