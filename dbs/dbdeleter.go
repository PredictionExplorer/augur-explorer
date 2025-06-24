package dbs

import (
	"os"
	"fmt"
//	"database/sql"
	 "github.com/ethereum/go-ethereum/common"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_deleter_status() []p.Deleter_status{

	var query string
	query = "SELECT "+
				"a.addr," +
				"a.address_id, "+
				"s.block_num, "+
				"s.info "+
			"FROM d_status AS s " +
				"JOIN address a ON a.addr=d.contract_addr "+
			"WHERE a.addr = d_status.contract_addr"

	rows,err := ss.db.Query(query)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.Deleter_status,0,101)
	defer rows.Close()
	for rows.Next() {
		var rec p.Deleter_status
		err=rows.Scan(
			&rec.ContractAddr,
			&rec.ContractAid,
			&rec.LastBlockNum,
			&rec.Info,
		)
		if err != nil {
			ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
		rec.ContractEthAddr = common.HexToAddress(rec.ContractAddr)
		records = append(records,rec)
	}
	return records
}
