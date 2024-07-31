package main
import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"

)
func cosmic_game_donations_cg_simple_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_simple_list(0, 10000)
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_simple_list.html", gin.H{
		"CosmicGameDonations" : donations,
	})
}
func cosmic_game_donations_cg_simple_by_round(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,HTTP,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'round_num' parameter is not set")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_simple_by_round(round_num)
	fmt.Printf("num records = %v\n",len(donations))
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_simple_by_round.html", gin.H{
		"CosmicGameDonations" : donations,
		"RoundNum": round_num,
	})
}
func cosmic_game_donations_cg_with_info_list(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_with_info_simple_list(0, 10000)
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_with_info_list.html", gin.H{
		"CosmicGameDonations" : donations,
	})
}
func cosmic_game_donations_cg_with_info_by_round(c *gin.Context) {

	if  !augur_srv.arbitrum_initialized() {
		respond_error(c,"Database link wasn't configured")
		return
	}
	p_round_num:= c.Param("round_num")
	var round_num int64
	if len(p_round_num) > 0 {
		var success bool
		round_num,success = parse_int_from_remote_or_error(c,HTTP,&p_round_num)
		if !success {
			return
		}
	} else {
		respond_error(c,"'round_num' parameter is not set")
		return
	}
	donations := arb_storagew.Get_donations_to_cosmic_game_with_info_by_round(round_num)
	c.HTML(http.StatusOK, "cg_donations_to_cosmicgame_with_info_by_round.html", gin.H{
		"CosmicGameDonations" : donations,
		"RoundNum": round_num,
	})
}
