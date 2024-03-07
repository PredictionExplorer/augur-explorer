package cosmicgame

import (
	"os"
	"fmt"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_cosmic_token_holders() []p.CGCosmicTokenHolderRec {

	var query string
	query = "SELECT "+
				"o.owner_aid,"+
				"oa.addr,"+
				"o.cur_balance,"+
				"o.cur_balance/1e18 " +
			"FROM "+sw.S.SchemaName()+".cg_costok_owner o "+
				"LEFT JOIN address oa ON o.owner_aid=oa.address_id "+
			"ORDER BY o.cur_balance DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicTokenHolderRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicTokenHolderRec
		err=rows.Scan(
			&rec.OwnerAid,
			&rec.OwnerAddr,
			&rec.Balance,
			&rec.BalanceFloat,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
