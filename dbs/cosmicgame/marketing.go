package cosmicgame

import (
	"os"
	"fmt"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_marketing_reward_history_global(offset,limit int) []p.CGMarketingRewardRec {

	var query string
	query = "SELECT "+
					"r.id,"+
					"r.evtlog_id,"+
					"r.block_num,"+
					"tx.id,"+
					"tx.tx_hash,"+
					"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT,"+
					"r.time_stamp,"+
					"r.amount,"+
					"r.amount/1e18,"+
					"r.marketer_aid,"+
					"ma.addr "+
				"FROM "+sw.S.SchemaName()+".cg_mkt_reward r "+
					"LEFT JOIN transaction tx ON tx.id=r.tx_id " +
					"LEFT JOIN address ma ON r.marketer_aid=ma.address_id "+
				"ORDER BY r.id DESC " +
				"OFFSET $1 LIMIT $2 "

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGMarketingRewardRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGMarketingRewardRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
			&rec.Amount,
			&rec.AmountEth,
			&rec.MarketerAid,
			&rec.MarketerAddr,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
