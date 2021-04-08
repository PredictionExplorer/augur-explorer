package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_arbitrum_augur_contract_addresses() (p.AA_ContractAddrs,error) {

	var query string
	query="SELECT " +
				"amm_factory,hatchery_reg "+
			"FROM aa_caddrs";
	row := ss.db.QueryRow(query)
	var c_addrs p.AA_ContractAddrs
	var err error
	var (
		amm_factory string
		hatchery_reg string
	)
	err=row.Scan(
		&amm_factory,&hatchery_reg,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_arbitrum_augur_contract_addresses(): %v",err))
			os.Exit(1)
		}
		return c_addrs,err
	}
	c_addrs.AMM_Factory=common.HexToAddress(amm_factory)
	c_addrs.HatcheryRegistry=common.HexToAddress(hatchery_reg)
	return c_addrs,nil
}
