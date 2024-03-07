package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_unclaimed_raffle_eth_deposits(winner_aid int64,offset,limit int) []p.CGRaffleDepositRec {

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
				"rd.round_num,"+
				"rd.amount/1e18 AS amount_eth,"+
				"rd.claimed, "+
				"EXTRACT(EPOCH FROM rw.time_stamp)::BIGINT AS tstmp, "+
				"rw.time_stamp "+
			"FROM cg_raffle_deposit rd "+
				"LEFT JOIN cg_raffle_withdrawal rw ON rw.evtlog_id=rd.withdrawal_id "+
				"LEFT JOIN transaction t ON t.id=rd.tx_id "+
				"LEFT JOIN address wa ON rd.winner_aid = wa.address_id "+
			"WHERE rd.winner_aid=$1 AND rd.claimed='F' " +
			"ORDER BY rd.id DESC "+
			"OFFSET $2 LIMIT $3"
	rows,err := sw.S.Db().Query(query,winner_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleDepositRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleDepositRec
		var null_ts sql.NullInt64
		var null_date sql.NullString
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WinnerAddr,
			&rec.WinnerAid,
			&rec.RoundNum,
			&rec.Amount,
			&rec.Claimed,
			&null_ts,
			&null_date,
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
func (sw *SQLStorageWrapper) Get_raffle_eth_deposits_list(offset,limit int) []p.CGRaffleDepositRec {

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
				"p.round_num, "+
				"p.amount/1e18 amount_eth "+
			"FROM "+sw.S.SchemaName()+".cg_raffle_deposit p "+
				"LEFT JOIN transaction t ON t.id=p.tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"ORDER BY p.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleDepositRec,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleDepositRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
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
func (sw *SQLStorageWrapper) Get_raffle_deposits_by_round(round_num int64) []p.CGRaffleDepositRec {

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
				"p.round_num,"+
				"p.amount/1e18 amount_eth "+
			"FROM "+sw.S.SchemaName()+".cg_raffle_deposit p "+
				"LEFT JOIN transaction t ON t.id=p.tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"WHERE p.round_num = $1 " +
			"ORDER BY p.id DESC "

	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleDepositRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleDepositRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
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
