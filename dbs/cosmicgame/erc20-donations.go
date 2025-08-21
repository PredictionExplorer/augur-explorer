package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_erc20_donations_by_round_detailed(round_num int64) []p.CGERC20Donation {

	var query string
	query = "SELECT "+
				"tok.id,"+
				"tok.evtlog_id,"+
				"tok.block_num,"+
				"tok.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM tok.time_stamp)::BIGINT,"+
				"tok.time_stamp,"+
				"tok.round_num,"+
				"tok.donor_aid,"+
				"da.addr, "+
				"tokaddr.address_id,"+
				"tokaddr.addr, "+
				"tok.amount, "+
				"tok.amount/1e18, "+
				"p.winner_aid,"+
				"wa.addr "+
			"FROM "+sw.S.SchemaName()+".cg_erc20_donation tok "+
				"INNER JOIN cg_prize_claim p ON p.round_num=tok.round_num "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tok.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON tok.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address tokaddr ON tok.token_aid=tokaddr.address_id "+
				"LEFT JOIN address wa ON wa.address_id = p.winner_aid "+
			"WHERE tok.round_num= $1 " +
			"ORDER BY tok.id DESC"
	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGERC20Donation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGERC20Donation
		var null_winner_addr sql.NullString
		var null_winner_aid sql.NullInt64
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.RoundNum,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.TokenAid,
			&rec.TokenAddr,
			&rec.Amount,
			&rec.AmountEth,
			&null_winner_aid,
			&null_winner_addr,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_winner_aid.Valid { rec.WinnerAid=null_winner_aid.Int64 }
		if null_winner_addr.Valid { rec.WinnerAddr = null_winner_addr.String }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_erc20_donations_by_round_summarized(round_num int64) []p.CGSummarizedERC20Donation {

	var query string
	query = "WITH claim AS ("+
				"SELECT SUM(amount) total,round_num,token_aid,winner_aid "+
				"FROM cg_donated_tok_claimed GROUP BY round_num,token_aid,winner_aid "+
			") " + 
			"SELECT "+
				"p.id,"+
				"p.evtlog_id,"+
				"p.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,"+
				"p.time_stamp,"+
				"dt20.round_num,"+
				"tokaddr.address_id,"+
				"tokaddr.addr, "+
				"dt20.total_amount, "+
				"dt20.total_amount/1e18, "+
				"claim.total, "+
				"claim.total/1e18, "+
				"dt20.total_amount-claim.total,"+
				"(dt20.total_amount-claim.total)/1e18,"+
				"p.winner_aid,"+
				"wa.addr, "+
				"dt20.claimed "+
			"FROM "+sw.S.SchemaName()+".cg_erc20_donation_stats dt20 "+
				"INNER JOIN cg_prize_claim p ON p.round_num=dt20.round_num "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=p.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address tokaddr ON dt20.token_aid=tokaddr.address_id "+
				"LEFT JOIN claim ON (claim.token_aid=dt20.token_aid AND dt20.round_num=claim.round_num) "+
				"LEFT JOIN address wa ON wa.address_id = p.winner_aid "+
			"WHERE p.round_num= $1 " +
			"ORDER BY dt20.token_aid DESC"
	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGSummarizedERC20Donation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGSummarizedERC20Donation
		var null_winner_addr sql.NullString
		var null_winner_aid sql.NullInt64
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.RoundNum,
			&rec.TokenAid,
			&rec.TokenAddr,
			&rec.AmountDonated,
			&rec.AmountDonatedEth,
			&rec.AmountClaimed,
			&rec.AmountClaimedEth,
			&rec.DonateClaimDiff,
			&rec.DonateClaimDiffEth,
			&null_winner_aid,
			&null_winner_addr,
			&rec.Claimed,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_winner_aid.Valid { rec.WinnerAid=null_winner_aid.Int64 }
		if null_winner_addr.Valid { rec.WinnerAddr = null_winner_addr.String }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_erc20_donations_global(offset, limit int) []p.CGERC20Donation {

	var query string
	query = "SELECT "+
				"tok.id,"+
				"tok.evtlog_id,"+
				"tok.block_num,"+
				"tok.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM tok.time_stamp)::BIGINT,"+
				"tok.time_stamp,"+
				"tok.round_num,"+
				"tok.donor_aid,"+
				"da.addr, "+
				"tokaddr.address_id,"+
				"tokaddr.addr, "+
				"tok.amount, "+
				"tok.amount/1e18, "+
				"p.winner_aid,"+
				"wa.addr  "+
			"FROM "+sw.S.SchemaName()+".cg_erc20_donation tok "+
				"INNER JOIN cg_prize_claim p ON p.round_num=tok.round_num "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tok.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON tok.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address tokaddr ON tok.token_aid=tokaddr.address_id "+
				"LEFT JOIN address wa ON wa.address_id = p.winner_aid "+
			"ORDER BY tok.id DESC " +
			"OFFSET $1 LIMIT $2"
	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGERC20Donation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGERC20Donation
		var null_winner_addr sql.NullString
		var null_winner_aid sql.NullInt64
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.RoundNum,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.TokenAid,
			&rec.TokenAddr,
			&rec.Amount,
			&rec.AmountEth,
			&null_winner_aid,
			&null_winner_addr,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_winner_aid.Valid { rec.WinnerAid=null_winner_aid.Int64 }
		if null_winner_addr.Valid { rec.WinnerAddr = null_winner_addr.String }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_erc20_donation_info(id int64) (bool,p.CGERC20Donation) {

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.round_num,"+
				"d.donor_aid,"+
				"da.addr, "+
				"toka.address_id,"+
				"toka.addr "+
			"FROM "+sw.S.SchemaName()+".cg_erc20_donation d "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address toka ON d.token_aid=toka.address_id "+
			"WHERE d.id=$1"

	row := sw.S.Db().QueryRow(query,id)
	var err error
	var rec p.CGERC20Donation
	rec.RecordId = id
	err=row.Scan(
		&rec.EvtLogId,
		&rec.BlockNum,
		&rec.TxId,
		&rec.TxHash,
		&rec.TimeStamp,
		&rec.DateTime,
		&rec.RoundNum,
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.TokenAid,
		&rec.TokenAddr,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_ERC2_donation_info(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return true,rec
}
