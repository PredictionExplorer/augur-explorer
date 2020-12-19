package dbs

import (
	"fmt"
	"os"

	"database/sql"
	_  "github.com/lib/pq"

//	"github.com/ethereum/go-ethereum/common"

	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_ens_processing_status() p.EnsProcStatus {

	var output p.EnsProcStatus
	var null_id sql.NullInt64

	var query string
	for {
		query = "SELECT block_num_limit,last_evt_id FROM uniswap_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO ens_status DEFAULT VALUES"
				_,err := ss.db.Exec(query)
				if (err!=nil) {
					ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
					os.Exit(1)
				}
			} else {
				ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
				os.Exit(1)
			}
		} else {
			break
		}
	}
	if null_id.Valid {
		output.LastEvtId = null_id.Int64
	}
	return output
}
