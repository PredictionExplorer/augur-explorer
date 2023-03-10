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
func (sw *SQLStorageWrapper) Delete_prize_claim_event(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_prize_claim WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_bid(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_bid WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donation(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_donation WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donation_received(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_donation_received  WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_donation_sent(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_donation_sent WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_charity_updated(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_charity_updated WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_token_name(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_token_name WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_mint_event(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".bw_mint_event WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
