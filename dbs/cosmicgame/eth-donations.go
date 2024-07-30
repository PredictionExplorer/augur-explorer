package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_charity_donations(cosmicgame_aid int64) []p.CGCharityDonation{

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr,"+
				"d.amount, "+
				"d.amount/1e18 amount_eth,  " +
				"d.round_num "+
			"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"ORDER BY d.id DESC"
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCharityDonation,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCharityDonation
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.Amount,
			&rec.AmountEth,
			&rec.RoundNum,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if rec.DonorAid == cosmicgame_aid { rec.IsVoluntary = false } else {rec.IsVoluntary=true}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_charity_donations_from_cosmic_game(cosmicgame_aid int64) []p.CGCharityDonation{

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr,"+
				"d.amount, "+
				"d.amount/1e18 amount_eth,  " +
				"d.round_num "+
			"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"WHERE donor_aid = $1 "+
			"ORDER BY d.id DESC"
	rows,err := sw.S.Db().Query(query,cosmicgame_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCharityDonation,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCharityDonation
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.Amount,
			&rec.AmountEth,
			&rec.RoundNum,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_charity_wallet_withdrawals() []p.CGCharityWithdrawal {

	var query string
	query = "SELECT "+
				"w.id,"+
				"w.evtlog_id,"+
				"w.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM w.time_stamp)::BIGINT,"+
				"w.time_stamp,"+
				"ca.addr,"+
				"w.amount,"+
				"w.amount/1e18 "+
			"FROM "+sw.S.SchemaName()+".cg_donation_sent w "+
				"LEFT JOIN transaction tx ON tx.id=w.tx_id "+
				"LEFT JOIN address ca ON w.charity_aid=ca.address_id "+
			"ORDER BY w.id DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCharityWithdrawal,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCharityWithdrawal
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DestinationAddr,
			&rec.Amount,
			&rec.AmountEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_donations_to_cosmic_game_simple_list() []p.CGCosmicGameDonation{

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr,"+
				"d.amount, "+
				"d.amount/1e18 amount_eth  " +
			"FROM "+sw.S.SchemaName()+".cg_donation d "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"ORDER BY d.id DESC"
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicGameDonation,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicGameDonation
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.Amount,
			&rec.AmountEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_donation_received_evt_id(tx_id,starting_id int64,sig string) int64 {

	var query string 
	query = "SELECT "+
				"d.evtlog_id "+
			"FROM "+
				"evt_log e "+
				"LEFT JOIN cg_donation_received d ON e.id=d.evtlog_id "+
			"WHERE "+
				"(e.tx_id=$1) AND "+
				"(e.topic0_sig=$2) AND "+
				"(e.id<$3) "+
			"ORDER BY e.id DESC LIMIT 1"
	res := sw.S.Db().QueryRow(query,tx_id,sig,starting_id)
	var null_id sql.NullInt64
	err := res.Scan(&null_id)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return null_id.Int64
}
func (sw *SQLStorageWrapper) Get_charity_donations_voluntary(cosmicgame_aid int64) []p.CGCharityDonation{

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr,"+
				"d.amount, "+
				"d.amount/1e18 amount_eth,  " +
				"d.round_num "+
			"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"WHERE donor_aid != $1 "+
			"ORDER BY d.id DESC"
	rows,err := sw.S.Db().Query(query,cosmicgame_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCharityDonation,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCharityDonation
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.Amount,
			&rec.AmountEth,
			&rec.RoundNum,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.IsVoluntary=true
		records = append(records,rec)
	}
	return records
}
