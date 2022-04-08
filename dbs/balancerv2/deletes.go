package balancerv2

import (
	"fmt"
	"os"
	_  "github.com/lib/pq"

)
func (sw *SQLStorageWrapper) Delete_balance_changes(block_num int64) {

	var err error
	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".tok_bal WHERE block_num=$1"
	sw.S.Info.Printf("deleting tok_bal records for block %v\n",block_num)
	_,err = sw.S.Db().Exec(query,block_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v, block_num=%v",err,query,block_num))
		os.Exit(1)
	}
}
func (sw *SQLStorageWrapper) Delete_swap_history(block_num int64) {

	var err error
	var query string
	query = "DELETE FROM "+sw.S.SchemaName()+".swf_hist WHERE block_num=$1"
	_,err = sw.S.Db().Exec(query,block_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v, block_num=%v",err,query,block_num))
		os.Exit(1)
	}
}
