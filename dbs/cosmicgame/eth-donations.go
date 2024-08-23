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
func (sw *SQLStorageWrapper) Get_donations_to_cosmic_game_simple_list(offset,limit int) []p.CGCosmicGameDonationSimple{

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
			"FROM "+sw.S.SchemaName()+".cg_donation d "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"ORDER BY d.id DESC " +
			"OFFSET $1 LIMIT $2"
	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicGameDonationSimple,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicGameDonationSimple
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
func (sw *SQLStorageWrapper) Get_donations_to_cosmic_game_simple_by_round(round_num int64) []p.CGCosmicGameDonationSimple{

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
				"d.amount/1e18 amount_eth, " +
				"d.round_num "+
			"FROM "+sw.S.SchemaName()+".cg_donation d "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"WHERE d.round_num = $1 "+
			"ORDER BY d.id DESC"
	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicGameDonationSimple,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicGameDonationSimple
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
func (sw *SQLStorageWrapper) Get_donations_to_cosmic_game_with_info_simple_list(offset,limit int) []p.CGCosmicGameDonationWithInfo{

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
				"d.round_num, "+
				"d.record_id,"+
				"dj.data "+
			"FROM "+sw.S.SchemaName()+".cg_donation_wi d "+
				"LEFT JOIN cg_donation_json dj ON dj.record_id=d.record_id "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"ORDER BY d.id DESC " +
			"OFFSET $1 LIMIT $2"
	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicGameDonationWithInfo,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicGameDonationWithInfo
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
			&rec.CGRecordId,
			&rec.DataJson,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_donations_to_cosmic_game_with_info_by_round(round_num int64) []p.CGCosmicGameDonationWithInfo{

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
				"d.round_num, "+
				"d.record_id,"+
				"dj.data "+
			"FROM "+sw.S.SchemaName()+".cg_donation_wi d "+
				"LEFT JOIN cg_donation_json dj ON dj.record_id=d.record_id "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"WHERE d.round_num=$1 "+
			"ORDER BY d.id DESC"
	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicGameDonationWithInfo,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicGameDonationWithInfo
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
			&rec.CGRecordId,
			&rec.DataJson,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_donations_to_cosmic_game_by_user(user_aid int64) []p.CGDonationCombinedRec {
	// returns both types: simple donation & donation-with-info
	var query string
	query = "SELECT "+
				"record_type,"+
				"evtlog_id,"+
				"block_num,"+
				"tx_id,"+
				"tx_hash,"+
				"ts," +
				"date_time,"+
				"donor_aid,"+
				"donor_addr,"+
				"amount,"+
				"amount_eth,"+
				"round_num,"+
				"record_id,"+
				"json_data "+
			"FROM ("+
				"(" +
					"SELECT "+
						"0 AS record_type,"+
						"d.evtlog_id,"+
						"d.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT ts,"+
						"d.time_stamp date_time,"+
						"d.donor_aid,"+
						"da.addr donor_addr,"+
						"d.amount, "+
						"d.amount/1e18 amount_eth, " +
						"d.round_num, "+
						"-1 AS record_id,"+
						"'' AS json_data "+
					"FROM "+sw.S.SchemaName()+".cg_donation d "+
						"LEFT JOIN transaction t ON t.id=tx_id "+
						"LEFT JOIN address da ON d.donor_aid=da.address_id "+
					"WHERE d.donor_aid = $1 "+
				") UNION ALL (" +
					"SELECT "+
						"1 AS record_type,"+
						"d.evtlog_id,"+
						"d.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT ts,"+
						"d.time_stamp date_time,"+
						"d.donor_aid,"+
						"da.addr donor_addr,"+
						"d.amount, "+
						"d.amount/1e18 amount_eth,  " +
						"d.round_num, "+
						"d.record_id,"+
						"dj.data json_data "+
					"FROM "+sw.S.SchemaName()+".cg_donation_wi d "+
						"LEFT JOIN cg_donation_json dj ON dj.record_id=d.record_id "+
						"LEFT JOIN transaction t ON t.id=tx_id "+
						"LEFT JOIN address da ON d.donor_aid=da.address_id "+
					"WHERE d.donor_aid = $1 " +
				")"+
			") donations " +
			"ORDER BY evtlog_id"
	rows,err := sw.S.Db().Query(query,user_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGDonationCombinedRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGDonationCombinedRec
		err=rows.Scan(
			&rec.RecordType,
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
			&rec.CGRecordId,
			&rec.DataJson,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_donation_with_info_record_info(record_id int64) p.CGCosmicGameDonationWithInfo {

	var query string
	query = "SELECT " +
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
				"d.round_num, "+
				"d.record_id,"+
				"dj.data "+
			"FROM "+sw.S.SchemaName()+".cg_donation_wi d "+
				"LEFT JOIN cg_donation_json dj ON dj.record_id=d.record_id "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"WHERE d.record_id=$1"
	res := sw.S.Db().QueryRow(query,record_id)
	var rec p.CGCosmicGameDonationWithInfo
	err:=res.Scan(
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
		&rec.CGRecordId,
		&rec.DataJson,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
		os.Exit(1)
	}
	return rec
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
