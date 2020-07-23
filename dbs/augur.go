// Data Base Storage
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
