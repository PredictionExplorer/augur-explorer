package biddingwar

import (
	"os"
	"fmt"

	//p "github.com/PredictionExplorer/augur-explorer/primitives/biddingwar"
)
// Note: these deletes are per transaction, therefore multiple records will be deleted
//			if transaction contains more than one event of the same kind,
//			it is done this way because INSERTs are per block, so, functions must be
//			called before processing each transaction
func (sw *SQLStorageWrapper) Delete_prize_claim_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_prize_claim WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_bid(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_bid WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donation(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_donation WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donation_received(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_donation_received  WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donation_sent(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_donation_sent WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_nft_donation_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_nft_donation WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_charity_updated(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_charity_updated WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_token_name(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_token_name WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_mint_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_mint_event WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_raffle_deposit(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_raffle_deposit WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_raffle_withdrawal(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_raffle_withdrawal WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_raffle_nft_winner(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_raffle_nft_winner WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_raffle_nft_claimed(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_raffle_nft_claimed WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donated_nft_claimed(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_donated_nft_claimed WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_token_transfer_event(evtlog_id int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_transfer WHERE evtlog_id=$1"
	_,err := sw.S.Db().Exec(query,evtlog_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
