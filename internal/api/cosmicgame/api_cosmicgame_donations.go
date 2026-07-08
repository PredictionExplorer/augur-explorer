package cosmicgame

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func api_cosmic_game_donations_cg_simple_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}
	donations, err := arbRepo.SimpleEthDonations(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"DirectCGDonations": donations,
		"Offset":            offset,
		"Limit":             limit,
	})
}
func api_cosmic_game_donations_cg_simple_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	donations, err := arbRepo.SimpleEthDonationsByRound(c.Request.Context(), round_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"DirectCGDonations": donations,
	})
}
func api_cosmic_game_donations_cg_with_info_list(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	donations, err := arbRepo.EthDonationsWithInfo(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"DirectCGDonations": donations,
		"Offset":            offset,
		"Limit":             limit,
	})
}
func api_cosmic_game_donations_cg_with_info_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	donations, err := arbRepo.EthDonationsWithInfoByRound(c.Request.Context(), round_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":            req_status,
		"error":             err_str,
		"DirectCGDonations": donations,
		"RoundNum":          round_num,
	})
}
func api_cosmic_game_donations_cg_with_info_record_info(c *gin.Context) {

	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_record_id := c.Param("record_id")
	var record_id int64
	if len(p_record_id) > 0 {
		var success bool
		record_id, success = common.ParseIntFromRemoteOrError(c, HTTP, &p_record_id)
		if !success {
			return
		}
	} else {
		common.RespondError(c, "'record_id' parameter is not set")
		return
	}
	record_info, err := arbRepo.EthDonationWithInfoRecord(c.Request.Context(), record_id)
	if err != nil && !errors.Is(err, store.ErrNotFound) {
		respondStoreError(c, err)
		return
	}
	// The legacy layer served the zero-value record for unknown ids (the
	// parity goldens pin that shape); ErrNotFound keeps exactly that.
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":      req_status,
		"error":       err_str,
		"ETHDonation": record_info,
		"RecordId":    record_id,
	})
}
func api_cosmic_game_donations_by_user(c *gin.Context) {

	if !dbInitialized() {
		common.RespondError(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "CombinedDonationRecords": []interface{}{},
			"UserAddr": p_user_addr, "UserAid": int64(0),
		})
		return
	}

	donations, err := arbRepo.EthDonationsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                  req_status,
		"error":                   err_str,
		"CombinedDonationRecords": donations,
		"UserAddr":                p_user_addr,
		"UserAid":                 user_aid,
	})
}
func api_cosmic_game_donations_cg_both_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	donations, err := arbRepo.EthDonationsByRound(c.Request.Context(), round_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":              req_status,
		"error":               err_str,
		"CosmicGameDonations": donations,
		"RoundNum":            round_num,
	})
}
func api_cosmic_game_donations_cg_both_all(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	donations, err := arbRepo.EthDonations(c.Request.Context())
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":              req_status,
		"error":               err_str,
		"CosmicGameDonations": donations,
	})
}
func api_cosmic_game_donations_erc20_by_round_detailed(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	donations, err := arbRepo.ERC20DonationsByRoundDetailed(c.Request.Context(), round_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                        req_status,
		"error":                         err_str,
		"DonationsERC20ByRoundDetailed": donations,
		"RoundNum":                      round_num,
	})
}
func api_cosmic_game_donations_erc20_by_round_all(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	donations, err := arbRepo.ERC20DonationsByRoundAll(c.Request.Context(), round_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                   req_status,
		"error":                    err_str,
		"DonationsERC20ByRoundAll": donations,
		"RoundNum":                 round_num,
	})
}
func api_cosmic_game_donations_erc20_by_round_summarized(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	donations, err := arbRepo.ERC20DonationsByRoundSummarized(c.Request.Context(), round_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                          req_status,
		"error":                           err_str,
		"DonationsERC20ByRoundSummarized": donations,
		"RoundNum":                        round_num,
	})
}
func api_cosmic_game_donations_erc20_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "DonatedPrizesERC20ByWinner": []interface{}{},
			"UserAddr": p_user_addr, "UserAid": int64(0),
		})
		return
	}

	donated_prizes, err := arbRepo.ERC20DonatedPrizesByWinner(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                     req_status,
		"error":                      err_str,
		"DonatedPrizesERC20ByWinner": donated_prizes,
		"UserAddr":                   p_user_addr,
		"UserAid":                    user_aid,
	})
}
func api_cosmic_game_donations_erc20_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	success, offset, limit := common.ParseOffsetLimitParamsJSON(c)
	if !success {
		return
	}

	donations, err := arbRepo.ERC20Donations(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":         req_status,
		"error":          err_str,
		"DonationsERC20": donations,
		"Offset":         offset,
		"Limit":          limit,
	})
}
func api_cosmic_game_donated_erc20_info(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	p_record_id := c.Param("record_id")
	var record_id int64
	if len(p_record_id) > 0 {
		var success bool
		record_id, success = common.ParseIntFromRemoteOrError(c, JSON, &p_record_id)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'record_id' parameter is not set")
		return
	}
	nftdonation, err := arbRepo.ERC20DonationInfo(c.Request.Context(), record_id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			common.RespondErrorJSON(c, "Record not found")
			return
		}
		respondStoreError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":        1,
		"error":         "",
		"ERC20Donation": nftdonation,
	})
}
func api_cosmic_game_donations_erc20_donated_by_user(c *gin.Context) {
	// DONOR PERSPECTIVE: Returns ERC20 tokens this user DONATED (not won)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "DonationsERC20ByDonor": []interface{}{},
			"UserAddr": p_user_addr, "UserAid": int64(0),
		})
		return
	}

	donations, err := arbRepo.ERC20DonationsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":                req_status,
		"error":                 err_str,
		"DonationsERC20ByDonor": donations,
		"UserAddr":              p_user_addr,
		"UserAid":               user_aid,
	})
}
func api_cosmic_game_erc20_claims_global(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}

	var offset, limit int
	p_offset := c.Param("offset")
	p_limit := c.Param("limit")

	if len(p_offset) == 0 || len(p_limit) == 0 {
		offset = 0
		limit = 10000
	} else {
		var success bool
		success, offset, limit = common.ParseOffsetLimitParamsJSON(c)
		if !success {
			return
		}
	}

	claims, err := arbRepo.ERC20DonationClaims(c.Request.Context(), offset, limit)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":      req_status,
		"error":       err_str,
		"ERC20Claims": claims,
		"Offset":      offset,
		"Limit":       limit,
	})
}
func api_cosmic_game_erc20_claims_by_user(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_user_addr := c.Param("user_addr")
	if len(p_user_addr) == 0 {
		common.RespondErrorJSON(c, "'user_addr' parameter is not set")
		return
	}
	user_aid, err := arb_storagew.S.Nonfatal_lookup_address_id(p_user_addr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 1, "error": "", "ERC20ClaimsByWinner": []interface{}{},
			"UserAddr": p_user_addr, "UserAid": int64(0),
		})
		return
	}

	claims, err := arbRepo.ERC20DonationClaimsByUser(c.Request.Context(), user_aid)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":              req_status,
		"error":               err_str,
		"ERC20ClaimsByWinner": claims,
		"UserAddr":            p_user_addr,
		"UserAid":             user_aid,
	})
}
func api_cosmic_game_erc20_claims_by_round(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	if !dbInitialized() {
		common.RespondErrorJSON(c, "Database link wasn't configured")
		return
	}
	p_round_num := c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num, success = common.ParseIntFromRemoteOrError(c, JSON, &p_round_num)
		if !success {
			return
		}
	} else {
		common.RespondErrorJSON(c, "'round_num' parameter is not set")
		return
	}

	claims, err := arbRepo.ERC20DonationClaimsByRound(c.Request.Context(), round_num)
	if err != nil {
		respondStoreError(c, err)
		return
	}
	var req_status int = 1
	var err_str string = ""
	c.JSON(http.StatusOK, gin.H{
		"status":             req_status,
		"error":              err_str,
		"ERC20ClaimsByRound": claims,
		"RoundNum":           round_num,
	})
}
