
package uniswapv3

import (
	"os"
	"fmt"

	"database/sql"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/uniswapv3"
)
type SQLStorageWrapper struct {
	S					*SQLStorage
}
func (sw *SQLStorageWrapper) Uniswap_get_contract_addrs() p.UniV3ContractAddrs {

	var query string
	query = "SELECT factory_addr,nft_pos_mgr_addr FROM "+sw.S.SchemaName()+".config"
	row := sw.S.Db().QueryRow(query)
	var factory_addr,nft_pos_mgr_addr string
	var err error
	err=row.Scan(&factory_addr,&nft_pos_mgr_addr);
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Uniswap_get_contract_addrs(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var output p.UniV3ContractAddrs
	output.FactoryAddr = factory_addr
	output.NFTPosMgrAddr=nft_pos_mgr_addr
	return output
}
func (sw *SQLStorageWrapper) Get_uniswap_v3_pool_aid(pool_addr string) int64 {

	var query string
	query = "SELECT p.pool_aid FROM pool_created p "+
				"JOIN addr a ON p.pool_aid=a.address_id "+
				"WHERE a.addr=$1"

	row := sw.S.Db().QueryRow(query,pool_addr)
	var err error
	var pool_aid int64
	err=row.Scan(&pool_aid);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_uniswap_v3_pool_aid(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return pool_aid
}
