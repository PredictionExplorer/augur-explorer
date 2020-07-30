package dbs

import (
	"fmt"
	"os"
	"database/sql"
	_  "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_contract_addresses() (p.ContractAddresses,error) {

	var query string
	query="SELECT	 augur,augur_trading,profit_loss,dai_cash,rep_token,zerox,wallet_reg,fill_order," +
					"eth_xchg,share_token,universe FROM contract_addresses";
	row := ss.db.QueryRow(query)
	var c_addresses p.ContractAddresses
	var err error
	var (
		augur string
		augur_trading string
		pl string
		dai string
		rep string
		zerox string
		walletreg string
		fill_order string
		eth_xchg string
		share_token string
		universe string
	)
	err=row.Scan(
		&augur,&augur_trading,&pl,&dai,&rep,&zerox,&walletreg,&fill_order,&eth_xchg,
		&share_token,&universe,
	);
	if (err!=nil) {
		if err == sql.ErrNoRows {
		} else {
			ss.Log_msg(fmt.Sprintf("Error in Get_contract_addresses(): %v",err))
			os.Exit(1)
		}
		return c_addresses,err
	}
	c_addresses.Augur=common.HexToAddress(augur)
	c_addresses.AugurTrading=common.HexToAddress(augur_trading)
	c_addresses.PL=common.HexToAddress(pl)
	c_addresses.Dai=common.HexToAddress(dai)
	c_addresses.Reputation=common.HexToAddress(rep)
	c_addresses.Zerox=common.HexToAddress(zerox)
	c_addresses.WalletReg=common.HexToAddress(walletreg)
	c_addresses.FillOrder=common.HexToAddress(fill_order)
	c_addresses.EthXchg=common.HexToAddress(eth_xchg)
	c_addresses.ShareToken=common.HexToAddress(share_token)
	c_addresses.Universe= common.HexToAddress(universe)
	return c_addresses,nil
}
func (ss *SQLStorage) Get_augur_blocks(market_aid int64) []int64 {

	var where_cond string = ""
	if market_aid > 0 {
		where_cond = fmt.Sprintf(" WHERE market_aid = %v ",market_aid)
	}
	var query string = ""
	query = "SELECT DISTINCT block_num FROM (" +
				"(SELECT DISTINCT block_num FROM mktord " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM market " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM mkt_fin " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM claim_funds " + where_cond + ") " +
					"UNION ALL" +
				"(SELECT DISTINCT block_num FROM report " + where_cond + ") " +
			") as block_numbers ORDER BY block_num"
	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]int64,0,4096)

	defer rows.Close()
	for rows.Next() {
		var block_num int64
		err=rows.Scan(&block_num)
		if err!=nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v, q=%v",err,query))
			os.Exit(1)
		}
		records = append(records,block_num)
	}
	return records
}
