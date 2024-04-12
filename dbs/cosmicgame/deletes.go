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
func (sw *SQLStorageWrapper) Delete_donation(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_donation WHERE evtlog_id=$1"
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
func (sw *SQLStorageWrapper) Delete_raffle_deposit(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_raffle_deposit WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_raffle_withdrawal(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_raffle_withdrawal WHERE evtlog_id=$1"
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
func (sw *SQLStorageWrapper) Delete_raffle_nft_claimed(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_raffle_nft_claimed WHERE evtlog_id=$1"
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
func (sw *SQLStorageWrapper) Delete_stake_action_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_stake_action WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_unstake_action_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_unstake_action WHERE evtlog_id=$1"
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
func (sw *SQLStorageWrapper) Delete_claim_reward_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_claim_reward WHERE evtlog_id=$1"
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
func (sw *SQLStorageWrapper) Delete_cosmic_game_num_raffle_eth_winners_per_round_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_raf_eth_winners WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_num_raffle_nft_winners_per_round_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_raf_nft_winners WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_cosmic_game_num_raffle_nft_holders_per_round_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_raf_nft_holders WHERE evtlog_id=$1"
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
func (sw *SQLStorageWrapper) Delete_cosmic_game_staking_wallet_address_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_staking_addr WHERE evtlog_id=$1"
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
func (sw *SQLStorageWrapper) Delete_business_logic_address_changed_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".cg_adm_blogic_addr WHERE evtlog_id=$1"
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
