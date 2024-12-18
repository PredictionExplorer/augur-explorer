package cosmicgame

import (
	"os"
	"fmt"

)
// Note: these deletes are per transaction, therefore multiple records will be deleted
//			if transaction contains more than one event of the same kind,
//			it is done this way because INSERTs are per block, so, functions must be
//			called before processing each transaction
func (sw *SQLStorageWrapper) Delete_prize_claim_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_prize_claim WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_bid(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_bid WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donation_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_donation WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donation_with_info_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_donation_wi WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donation_received(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_donation_received  WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donation_sent(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_donation_sent WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_erc20_donated_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_erc20_donation WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_nft_donation_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_nft_donation WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_charity_updated(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_charity_updated WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_token_name(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_token_name WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_mint_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_mint_event WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_prize_deposit(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_prize_deposit WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_prize_withdrawal(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_prize_withdrawal WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_raffle_nft_winner(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_raffle_nft_winner WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_raffle_eth_winner(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_raffle_eth_winner WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_endurance_winner(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_endurance_winner WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_lastcst_bidder_winner(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_lastcst_winner WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donated_token_claimed(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_donated_tok_claimed WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donated_nft_claimed(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_donated_nft_claimed WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_stake_action_cst_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_stake_action_cst WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_unstake_action_cst_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_unstake_action_cst WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_eth_deposit_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_eth_deposit WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
/* DISCONTINUED, removal pending
func (sw *SQLStorageWrapper) Delete_claim_reward_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_claim_reward WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}*/
func (sw *SQLStorageWrapper) Delete_stake_action_rwalk_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_stake_action_rwalk WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_unstake_action_rwalk_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_unstake_action_rwalk WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_marketing_reward_sent_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_mkt_reward WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_signature_transfer_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_transfer WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_token_transfer_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_erc20_transfer WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_charity_percentage_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_charity_pcent WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_prize_percentage_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_prize_pcent WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_raffle_percentage_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_raffle_pcent WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_num_raffle_eth_winners_bidding_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_raf_eth_bidding WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_num_raffle_nft_winners_bidding_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_raf_nft_bidding WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_num_raffle_nft_winners_staking_cst_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_raf_nft_staking_cst WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_num_raffle_nft_winners_staking_rwalk_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_raf_nft_staking_rwalk WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_staking_percentage_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_stake_pcent WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_chrono_percentage_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_chrono_pcent WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_system_mode_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_sysmode WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_charity_address_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_charity_addr WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_random_walk_address_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_rwalk_addr WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_raffle_wallet_address_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_raffle_addr WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_staking_wallet_cst_address_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_staking_cst_addr WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_staking_wallet_rwalk_address_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_staking_rwalk_addr WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_marketing_wallet_address_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_marketing_addr WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_token_address_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_costok_addr WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_signature_address_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_cossig_addr WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_upgraded_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_upgraded WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_time_increase_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_time_inc WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_timeout_claimprize_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_timeout_claimprize WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_price_increase_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_price_inc WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_nanoseconds_extra_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_nanosec_extra WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_initial_seconds_until_prize_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_inisecprize WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_initial_bid_amount_fraction_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_bidfraction WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_activation_time_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_acttime WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_ethcst_bid_ratio_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_ethcst WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_round_start_cst_auction_length_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_auclen WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
	sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_erc20_reward_multiplier_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_erc_rwd_mul WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
/* DISCONTINUED
func (sw *SQLStorageWrapper) Delete_starting_bid_price_st_min_limit_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_cst_min_lim WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}*/
func (sw *SQLStorageWrapper) Delete_marketing_reward_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_mkt_reward WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_erc20_token_reward_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_erc20_reward WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_max_message_length_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_msg_len WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_token_generation_script_url_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_script_url WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_base_uri_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_base_uri_cs WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_nft_staked_cst_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_nft_staked_cst WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_nft_staked_rwalk_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_nft_staked_rwalk WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_nft_unstaked_rwalk_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_nft_unstaked_rwalk WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_nft_unstaked_cst_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_nft_unstaked_cst WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_reward_paid_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_reward_paid WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_ownership_transferred_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_ownership WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_initialized_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_initialized WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_chrono_warrior_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_chrono_warrior WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cst_min_limit_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_cst_min_limit WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_fund_transfer_failed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_fund_transfer_err WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_erc20_transfer_failed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_erc20_transfer_err WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_funds_transferred_to_charity_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_funds_to_charity WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_delay_duration_before_next_round_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_delay_duration WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_round_started_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_round_started WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
