package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_unclaimed_prize_eth_deposits(winner_aid int64,offset,limit int) []p.CGPrizeDepositRec {

	var query string
	query = 
			"SELECT "+
				"rd.id,"+
				"rd.evtlog_id,"+
				"rd.block_num,"+
				"rd.tx_id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM rd.time_stamp)::BIGINT AS tstmp, "+
				"rd.time_stamp AS date_time, "+
			"wa.addr,"+
			"rd.winner_aid,"+
		"rd.winner_index,"+
		"rd.round_num,"+
		"rd.amount/1e18 AS amount_eth,"+
		"rd.claimed, "+
		"EXTRACT(EPOCH FROM rw.time_stamp)::BIGINT AS tstmp, "+
		"rw.time_stamp, "+
		"CASE WHEN cw.round_num IS NOT NULL THEN 7 ELSE 10 END AS record_type "+
	"FROM cg_prize_deposit rd "+
		"LEFT JOIN cg_prize_withdrawal rw ON rw.evtlog_id=rd.withdrawal_id "+
		"LEFT JOIN transaction t ON t.id=rd.tx_id "+
		"LEFT JOIN address wa ON rd.winner_aid = wa.address_id "+
		"LEFT JOIN cg_chrono_warrior_prize cw ON (rd.round_num = cw.round_num AND rd.winner_index = cw.winner_index) "+
	"WHERE rd.winner_aid=$1 AND rd.claimed='F' " +
			"ORDER BY rd.id DESC "+
			"OFFSET $2 LIMIT $3"
	rows,err := sw.S.Db().Query(query,winner_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGPrizeDepositRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGPrizeDepositRec
		var null_ts sql.NullInt64
		var null_date sql.NullString
		err=rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		&rec.Tx.DateTime,
		&rec.WinnerAddr,
		&rec.WinnerAid,
		&rec.WinnerIndex,
		&rec.RoundNum,
	&rec.Amount,
	&rec.Claimed,
	&null_ts,
	&null_date,
	&rec.RecordType,
)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	if null_ts.Valid { rec.ClaimTimeStamp = null_ts.Int64 }
	if null_date.Valid { rec.ClaimDateTime = null_date.String }
	records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_prize_eth_deposits_list(offset,limit int) []p.CGPrizeDepositRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"p.id,"+
				"p.evtlog_id,"+
				"p.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,"+
				"p.time_stamp,"+
			"p.winner_aid,"+
			"wa.addr,"+
			"p.winner_index,"+
			"p.round_num, "+
			"p.amount/1e18 amount_eth "+
		"FROM "+sw.S.SchemaName()+".cg_prize_deposit p "+
			"LEFT JOIN transaction t ON t.id=p.tx_id "+
			"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
		"ORDER BY p.id DESC "+
		"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGPrizeDepositRec,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGPrizeDepositRec
		err=rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		&rec.Tx.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.WinnerIndex,
			&rec.RoundNum,
			&rec.Amount,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_prize_deposits_by_round(round_num int64) []p.CGPrizeDepositRec {

	var query string
	query =  "SELECT " +
			"p.id,"+
			"p.evtlog_id,"+
			"p.block_num,"+
			"t.id,"+
			"t.tx_hash,"+
			"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,"+
			"p.time_stamp,"+
		"p.winner_aid,"+
		"wa.addr,"+
		"p.winner_index,"+
		"p.round_num,"+
		"p.amount/1e18 amount_eth "+
	"FROM "+sw.S.SchemaName()+".cg_prize_deposit p "+
		"INNER JOIN "+sw.S.SchemaName()+".cg_prize pr ON (pr.round_num = p.round_num AND pr.winner_index = p.winner_index AND pr.ptype = 10) "+
		"LEFT JOIN transaction t ON t.id=p.tx_id "+
		"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
	"WHERE p.round_num = $1 " +
	"ORDER BY p.winner_index "

	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGPrizeDepositRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGPrizeDepositRec
		err=rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		&rec.Tx.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.WinnerIndex,
			&rec.RoundNum,
			&rec.Amount,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records

}
