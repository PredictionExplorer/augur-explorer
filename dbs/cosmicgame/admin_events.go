package cosmicgame

import (
	"os"
	"fmt"
	"math"

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
	var next_evtlog int64 = math.MaxInt64;
	for i:=0; i<len(records); i++ {
		r := records[i];
		r.NextEvtLogId = next_evtlog
		next_evtlog = r.EvtLogId
		records[i]=r
	}
	return records
}
func (sw *SQLStorageWrapper) Get_admin_events_in_range(evtlog_start,evtlog_end int64) []p.CGAdminEvent {

	var query string
	query = "SELECT "+
				"record_type,"+
				"record_id,"+
				"evtlog_id,"+
				"block_num,"+
				"tx_id,"+
				"tx_hash,"+
				"ts,"+
				"date_time,"+
				"addr_value "+ 
			"FROM ("+
				"("+
					"SELECT "+
						"1 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_charity_addr r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.new_charity_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"2 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_rwalk_addr r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.new_rwalk_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				")" +
			") everything "+
			"ORDER BY evtlog_id "

			fmt.Printf("q= \n%v\n",query)
			fmt.Printf("evtlog_start = %v\n",evtlog_start)
			fmt.Printf("evtlog_end = %v\n",evtlog_end)
	rows,err := sw.S.Db().Query(query,evtlog_start,evtlog_end)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGAdminEvent,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGAdminEvent
		err=rows.Scan(
			&rec.RecordType,
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.AddressValue,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
