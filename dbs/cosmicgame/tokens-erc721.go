package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_cosmic_signature_nft_list(offset,limit int) []p.CGCosmicSignatureMintRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"m.evtlog_id,"+
				"m.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT,"+
				"m.time_stamp,"+
				"m.owner_aid,"+
				"wa.addr,"+
				"m.cur_owner_aid,"+
				"oa.addr,"+
				"m.seed, "+
				"m.token_id,"+
				"m.token_name,"+
				"m.round_num,"+
				"p.prize_num "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event m "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
			"ORDER BY m.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicSignatureMintRec,0, 64)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicSignatureMintRec
		var null_prize_num sql.NullInt64
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.CurOwnerAid,
			&rec.CurOwnerAddr,
			&rec.Seed,
			&rec.TokenId,
			&rec.TokenName,
			&rec.RoundNum,
			&null_prize_num,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_prize_num.Valid { rec.RecordType = 3 } else {rec.RecordType = 1 }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_cosmic_signature_token_info(token_id int64) (bool,p.CGCosmicSignatureMintRec) {

	var query string
	query = "SELECT "+
				"m.evtlog_id,"+
				"m.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT,"+
				"m.time_stamp,"+
				"m.owner_aid,"+
				"wa.addr,"+
				"m.cur_owner_aid,"+
				"oa.addr,"+
				"m.seed, "+
				"m.token_id,"+
				"m.round_num,"+
				"p.prize_num, "+
				"m.token_name, "+
				"m.staked,"+
				"m.staked_owner_aid,"+
				"sa.addr,"+
				"st.action_id,"+
				"EXTRACT(EPOCH FROM st.unstake_time)::BIGINT,"+
				"st.unstake_time, "+
				"u.id, "+
				"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT,"+
				"u.time_stamp "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event m "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
				"LEFT JOIN cg_stake_action st ON m.stake_action_id=st.id "+
				"LEFT JOIN cg_unstake_action u ON st.action_id=u.action_id "+
				"LEFT JOIN address sa ON m.staked_owner_aid = sa.address_id "+
			"WHERE m.token_id=$1"

	var rec p.CGCosmicSignatureMintRec
	var err error
	var null_prize_num,null_unstake_id,null_action_id sql.NullInt64
	var null_ue_timestamp,null_au_timestamp sql.NullInt64
	var null_staked_owner_addr,null_ue_datetime,null_au_datetime sql.NullString
	row := sw.S.Db().QueryRow(query,token_id)
	err=row.Scan(
		&rec.EvtLogId,
		&rec.BlockNum,
		&rec.TxId,
		&rec.TxHash,
		&rec.TimeStamp,
		&rec.DateTime,
		&rec.WinnerAid,
		&rec.WinnerAddr,
		&rec.CurOwnerAid,
		&rec.CurOwnerAddr,
		&rec.Seed,
		&rec.TokenId,
		&rec.RoundNum,
		&null_prize_num,
		&rec.TokenName,
		&rec.Staked,
		&rec.StakedOwnerAid,
		&null_staked_owner_addr,
		&null_action_id,
		&null_ue_timestamp,
		&null_ue_datetime,
		&null_unstake_id,
		&null_au_timestamp,
		&null_au_datetime,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_prize_num.Valid { rec.RecordType = 3 } else {rec.RecordType = 1 }
	if null_unstake_id.Valid { rec.WasUnstaked = true } 
	if null_action_id.Valid { rec.StakeActionId = null_action_id.Int64 }
	if null_ue_timestamp.Valid { rec.UnstakeElligibleTimeStamp = null_ue_timestamp.Int64 }
	if null_ue_datetime.Valid { rec.UnstakeElligibleDateTime=null_ue_datetime.String }
	if null_au_timestamp.Valid { rec.ActualUnstakeTimeStamp = null_au_timestamp.Int64 }
	if null_au_datetime.Valid { rec.ActualUnstakeDateTime = null_au_datetime.String }
	if null_staked_owner_addr.Valid { rec.StakedOwnerAddr = null_staked_owner_addr.String }
	return true,rec
}
func (sw *SQLStorageWrapper) Get_cosmic_signature_token_name_history(token_id int64) []p.CGTokenName {

	var query string
	query = "SELECT "+
				"n.evtlog_id,"+
				"n.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM n.time_stamp)::BIGINT,"+
				"n.time_stamp,"+
				"n.token_id,"+
				"n.token_name "+
			"FROM "+sw.S.SchemaName()+".cg_token_name n "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
			"WHERE n.token_id=$1 "+
			"ORDER BY n.id DESC "

	rows,err := sw.S.Db().Query(query,token_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGTokenName,0, 64)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGTokenName
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TokenId,
			&rec.TokenName,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_cst_ownership_transfers(token_id int64,offset,limit int) []p.CGTransfer {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"t.id,"+
				"t.evtlog_id,"+
				"t.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,"+
				"t.time_stamp,"+
				"t.from_aid,"+
				"fa.addr,"+
				"t.to_aid,"+
				"ta.addr,"+
				"t.token_id,"+
				"t.otype "+
			"FROM "+sw.S.SchemaName()+".cg_transfer t "+
				"LEFT JOIN transaction tx ON tx.id=t.tx_id "+
				"LEFT JOIN address fa ON t.from_aid=fa.address_id "+
				"LEFT JOIN address ta ON t.to_aid=ta.address_id "+
			"WHERE t.token_id=$1 "+
			"ORDER BY t.id DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := sw.S.Db().Query(query,token_id,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGTransfer,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGTransfer
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.FromAid,
			&rec.FromAddr,
			&rec.ToAid,
			&rec.ToAddr,
			&rec.TokenId,
			&rec.TransferType,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_cosmic_signature_token_distribution() []p.CGCSTokenDistributionRec {

	var query string
	query = "SELECT "+
				"m.cur_owner_aid, "+
				"a.addr, "+
				"count(m.cur_owner_aid) AS counter "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event m "+
				"LEFT JOIN address a ON m.cur_owner_aid=a.address_id "+
			"GROUP BY m.cur_owner_aid,a.addr "+
			"ORDER BY counter DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCSTokenDistributionRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCSTokenDistributionRec
		err=rows.Scan(
			&rec.OwnerAid,
			&rec.OwnerAddr,
			&rec.NumTokens,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Search_token_by_name(name string) []p.CGTokenSearchResult {

	name = "%" + name + "%"
	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,"+
				"t.time_stamp,"+
				"t.token_id,"+
				"t.token_name "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event t "+
			"WHERE t.token_name ILIKE  $1 "+
			"ORDER BY t.token_id"

	rows,err := sw.S.Db().Query(query,name)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGTokenSearchResult,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGTokenSearchResult
		err=rows.Scan(
			&rec.MintTimeStamp,
			&rec.MintDateTime,
			&rec.TokenId,
			&rec.TokenName,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_named_tokens() []p.CGTokenSearchResult {

	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,"+
				"t.time_stamp,"+
				"t.token_id,"+
				"t.token_name "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event t "+
			"WHERE LENGTH(t.token_name)>0 "+
			"ORDER BY t.token_name"

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGTokenSearchResult,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGTokenSearchResult
		err=rows.Scan(
			&rec.MintTimeStamp,
			&rec.MintDateTime,
			&rec.TokenId,
			&rec.TokenName,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
