package dbs
import (
	"os"
	"fmt"
	"database/sql"
	p "github.com/PredictionExplorer/augur-explorer/primitives"
)
func (ss *SQLStorage) Get_randomwalk_processing_status() p.RandomWalkProcStatus {

	var output p.RandomWalkProcStatus
	var null_id,null_block sql.NullInt64

	var query string
	for {
		query = "SELECT last_evt_id,last_block FROM rw_proc_status"

		res := ss.db.QueryRow(query)
		err := res.Scan(&null_id,&null_block)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				query = "INSERT INTO rw_proc_status DEFAULT VALUES"
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
		output.LastIdProcessed = null_id.Int64
	}
	if null_block.Valid {
		output.LastBlockNum = null_block.Int64
	}
	return output
}
func (ss *SQLStorage) Update_randomwalk_process_status(status *p.RandomWalkProcStatus) {

	var query string
	query = "UPDATE rw_proc_status SET last_evt_id = $1,last_block=$2"

	_,err := ss.db.Exec(query,status.LastIdProcessed,status.LastBlockNum)
	if (err!=nil) {
		ss.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
}
func (ss *SQLStorage) Get_randomwalk_contract_addresses() p.RW_ContractAddresses {

	var output p.RW_ContractAddresses
	var query string
	query = "SELECT "+
				"marketplace_addr,randomwalk_addr," +
				"mp_a.address_id,rw_a.address_id "+
			"FROM rw_contracts ca " +
				"JOIN address mp_a ON mp.marketplace_addr=mp_a.addr " +
				"JOIN address rw_a ON rw.ransomwalk_addr=rw_a.addr "

	res := ss.db.QueryRow(query)
	err := res.Scan(
		&output.MarketPlace,
		&output.RandomWalk,
		&output.MarketPlaceAid,
		&output.RandomWalkAid,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return output
		}
		ss.Log_msg(fmt.Sprintf("Get_randomwalk_contract_addresses() failed: %v, q=%v",err,query))
		os.Exit(1)
	}
	return output
}
