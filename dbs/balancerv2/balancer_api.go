package balancerv2

import (
	"os"
	"fmt"

	"database/sql"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/balancerv2"
)
func (sw *SQLStorageWrapper) Get_pool_info(pool_id string) p.PoolInfo {

	var output p.PoolInfo
	var query string
	query = "SELECT " +
				"block_num,"+
				"pa.addr,"+
				"specialization "+
			"FROM "+ss.S.SchemaName()+".pool_reg p "+
				"JOIN "+ss.S.SchemaName()+".addr pa ON p.pool_aid=pa.address_id "+
			"WHERE p.pool_id=$1"


	row := sw.S.Db().QueryRow(query,pool_id)
	var err error
	err=row.Scan(
		&output.BlockNum,
		&output.PoolAddr,
		&output.Specialization,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_pool_info(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return output
}
func (sw *SQLStorageWrapper) Get_pool_total_swaps(pool_id string) int64 {

	var query string
	query = "SELECT "+
					"count(*) totswaps "+
				"FROM "+ss.S.SchemaName()+".swap "+
				"WHERE pool_id=$1"

	row := sw.S.Db().QueryRow(query,pool_id)
	var err error
	var total_swaps int64
	err=row.Scan(&total_swaps)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_pool_info(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return total_swaps
}
func (sw *SQLStorageWrapper) Get_pool_tokens(pool_aid int64) {

	var query string
	query = "SELECT " +
				"bb.aid,"+
				"ta.addr,"+
				"bb.balance "+
			"FROM "+sw.S.SchemaName()+".bpt_bal bb "+
				"JOIN "+sw.S.SchemaName()+".addr ta ON bb.aid=ta.address_id"+
			"WHERE bb.pool_aid=$1"
				

}
