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
				"p.prize_num, "+
				"stel.erc721_token_id,"+
				"endu.erc721_token_id, "+
				"rnw.is_staker, "+
				"rnw.id "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event m "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
				"LEFT JOIN cg_stellar_winner stel ON m.token_id=stel.erc721_token_id "+
				"LEFT JOIN cg_endurance_winner endu ON m.token_id=endu.erc721_token_id "+
				"LEFT JOIN cg_raffle_nft_winner rnw ON m.token_id=rnw.token_id "+
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
		var null_prize_num,null_raffle_id sql.NullInt64
		var null_staked sql.NullBool
		var null_endu_token_id,null_stel_token_id sql.NullInt64
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
			&null_stel_token_id,
			&null_endu_token_id,
			&null_staked,
			&null_raffle_id,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.RecordType = 3	// main prize
		if null_raffle_id.Valid { rec.RecordType = 1 }	// raffle NFT winer
		if null_staked.Valid { 
			if null_staked.Bool {rec.RecordType = 2 }	// nft won due to staking (RWalk)
		}
		if null_endu_token_id.Valid { rec.RecordType = 4 } // endurance champion
		if null_stel_token_id.Valid { rec.RecordType = 5 }	// stellar spender
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
				"st.staker_aid,"+
				"sta.addr,"+
				"sa.action_id,"+
				"EXTRACT(EPOCH FROM sa.time_stamp)::BIGINT,"+
				"sa.time_stamp,"+
				"u.id, "+
				"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT,"+
				"u.time_stamp, "+
				"stel.erc721_token_id,"+
				"endu.erc721_token_id, "+
				"rnw.is_staker, "+
				"rnw.id "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event m "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
				"LEFT JOIN cg_staked_token_cst st ON (m.token_id=st.token_id)"+
				"LEFT JOIN cg_nft_staked_cst sa ON sa.token_id = m.token_id "+
				"LEFT JOIN cg_nft_unstaked_cst u ON u.token_id=m.token_id "+
				"LEFT JOIN cg_stellar_winner stel ON m.token_id=stel.erc721_token_id "+
				"LEFT JOIN cg_endurance_winner endu ON m.token_id=endu.erc721_token_id "+
				"LEFT JOIN cg_raffle_nft_winner rnw ON m.token_id=rnw.token_id "+
				"LEFT JOIN address sta ON st.staker_aid = sta.address_id "+
			"WHERE m.token_id=$1"

	var rec p.CGCosmicSignatureMintRec
	var err error
	var null_prize_num,null_unstake_id,null_action_id,null_staker_aid,null_raffle_id sql.NullInt64
	var null_au_timestamp,null_sa_timestamp sql.NullInt64
	var null_staked_owner_addr,null_au_datetime,null_sa_datetime sql.NullString
	var null_staked sql.NullBool
	var null_endu_token_id,null_stel_token_id sql.NullInt64
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
		&null_staker_aid,
		&null_staked_owner_addr,
		&null_action_id,
		&null_sa_timestamp,
		&null_sa_datetime,
		&null_unstake_id,
		&null_au_timestamp,
		&null_au_datetime,
		&null_stel_token_id,
		&null_endu_token_id,
		&null_staked,
		&null_raffle_id,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	rec.RecordType = 3	// main prize
	if null_raffle_id.Valid { rec.RecordType = 1 }	// raffle NFT winer
	if null_staked.Valid { 
		if null_staked.Bool {rec.RecordType = 2 }	// nft won due to staking (RWalk)
	}
	if null_endu_token_id.Valid { rec.RecordType = 4 } // endurance champion
	if null_stel_token_id.Valid { rec.RecordType = 5 }	// stellar spender
	if null_unstake_id.Valid { rec.UnstakeActionId = null_unstake_id.Int64 } else {rec.UnstakeActionId = -1}
	if null_action_id.Valid { rec.StakeActionId = null_action_id.Int64 } else {rec.StakeActionId=-1}
	if null_au_timestamp.Valid { rec.ActualUnstakeTimeStamp = null_au_timestamp.Int64 }
	if null_au_datetime.Valid { rec.ActualUnstakeDateTime = null_au_datetime.String }
	if null_staked_owner_addr.Valid { rec.StakedOwnerAddr = null_staked_owner_addr.String }
	if null_staked.Valid { rec.Staked = null_staked.Bool } else {rec.Staked = false}
	if null_staker_aid.Valid { rec.StakedOwnerAid = null_staker_aid.Int64 } else { rec.StakedOwnerAid = -1 }
	if null_sa_timestamp.Valid { rec.StakeTimeStamp = null_sa_timestamp.Int64 } else { rec.StakeTimeStamp = -1}
	if null_sa_datetime.Valid { rec.StakeDateTime = null_sa_datetime.String} 
	if (rec.StakeActionId > -1) && (rec.UnstakeActionId > -1) { rec.WasUnstaked = true }
	if (rec.StakeActionId > -1) && (rec.UnstakeActionId == -1) { rec.Staked = true }
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
