
package uniswapv3

import (
	"os"
	"fmt"

//	"database/sql"
	. "github.com/PredictionExplorer/augur-explorer/dbs"
	_  "github.com/lib/pq"

	p "github.com/PredictionExplorer/augur-explorer/primitives/uniswapv3"
)
type SQLStorageWrapper struct {
	S					*SQLStorage
}
func (sw *SQLStorageWrapper) Uniswap_get_contract_addrs() p.UniV3ContractAddrs {

	var query string
	query = "SELECT factory_addr FROM "+sw.S.SchemaName()+".config"
	row := sw.S.Db().QueryRow(query)
	var factory_addr string
	var err error
	err=row.Scan(&factory_addr);
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Uniswap_get_contract_addrs(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var output p.UniV3ContractAddrs
	output.FactoryAddr = factory_addr
	return output
}
