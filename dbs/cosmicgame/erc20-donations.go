package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_erc20_donations_by_round(round_num int64) []p.CGERC20Donation {

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
				"tc.winner_aid,"+
				"wa.addr "+
			"FROM "+sw.S.SchemaName()+".cg_erc20_donation tok "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tok.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON tok.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address tokaddr ON tok.token_aid=tokaddr.address_id "+
				"LEFT JOIN cg_donated_tok_claimed tc ON (tc.token_aid=tok.token_aid AND tok.round_num=tc.round_num)"+
				"LEFT JOIN address wa ON wa.address_id = tc.winner_aid "+
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
		if null_winner_aid.Valid { rec.Claimed = true; rec.WinnerAid=null_winner_aid.Int64 }
		if null_winner_addr.Valid { rec.WinnerAddr = null_winner_addr.String }
		records = append(records,rec)
	}
	return records
}
