package cosmicgame

import (
	"os"
	"fmt"

	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_system_mode_change_event_list(offset,limit int) []p.CGSystemModeRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"s.id,"+
				"s.evtlog_id,"+
				"s.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT,"+
				"s.time_stamp, "+
				"s.sysmode "+
			"FROM "+sw.S.SchemaName()+".cg_adm_sysmode s "+
				"LEFT JOIN transaction t ON t.id=s.tx_id "+
			"ORDER BY s.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGSystemModeRec,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGSystemModeRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.SystemMode,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
