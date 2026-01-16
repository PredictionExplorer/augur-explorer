package uniswapv3

import (
	"fmt"
	"os"
	_  "github.com/lib/pq"

)
func (sw *SQLStorageWrapper) Delete_pool_created(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".pool_created WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_pool_initialize(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".initialize WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_pool_mint(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".mint WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_pool_burn(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".burn WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_pool_swap(block_num,tx_index int64) {

	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".swap WHERE block_num=$1 AND tx_index=$2"
	_,err := sw.S.Db().Exec(query,block_num,tx_index)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
