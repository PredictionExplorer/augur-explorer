package cosmicgame

import (
	"os"
	"fmt"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func query_nft_winners(is_staker bool) string {

	var query string
	var staking_condition ="'F'"
	if is_staker {
		staking_condition = "'T'"
	}
	query = "SELECT "+
				"p.evtlog_id,"+
				"p.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,"+
				"p.time_stamp,"+
				"p.winner_aid,"+
				"wa.addr,"+
				"p.round_num, "+
				"p.token_id, "+
				"p.cst_amount, "+
				"p.cst_amount/1e18 cst_amount_eth, "+
				"p.winner_idx, "+
				"p.is_rwalk,"+
				"p.is_staker "+
			"FROM cg_raffle_nft_prize p "+
				"LEFT JOIN transaction t ON t.id=p.tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"WHERE p.round_num=$1 AND p.is_staker= " + staking_condition +
			"ORDER BY p.id DESC"
	return query
}
func (sw *SQLStorageWrapper) Get_raffle_nft_winners_by_round(round_num int64,is_staker bool) []p.CGRaffleNFTWinnerRec {

	var query string
	query = query_nft_winners(is_staker)
	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleNFTWinnerRec,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleNFTWinnerRec
		err=rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		&rec.Tx.DateTime,
		&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.RoundNum,
			&rec.TokenId,
			&rec.CstAmount,
			&rec.CstAmountEth,
			&rec.WinnerIndex,
			&rec.IsRWalk,
			&rec.IsStaker,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_raffle_nft_winners(offset,limit int) []p.CGRaffleNFTWinnerRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"p.evtlog_id,"+
				"p.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,"+
				"p.time_stamp,"+
				"p.winner_aid,"+
				"wa.addr,"+
				"p.round_num, "+
				"p.token_id, "+
				"p.cst_amount, "+
				"p.cst_amount/1e18 cst_amount_eth, "+
				"p.winner_idx, "+
				"p.is_rwalk,"+
				"p.is_staker "+
			"FROM "+sw.S.SchemaName()+".cg_raffle_nft_prize p "+
				"LEFT JOIN transaction t ON t.id=p.tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"ORDER BY p.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleNFTWinnerRec,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleNFTWinnerRec
		err=rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		&rec.Tx.DateTime,
		&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.RoundNum,
			&rec.TokenId,
			&rec.CstAmount,
			&rec.CstAmountEth,
			&rec.WinnerIndex,
			&rec.IsRWalk,
			&rec.IsStaker,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
