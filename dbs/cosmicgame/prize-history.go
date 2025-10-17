package cosmicgame

import (
	"os"
	"fmt"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_prize_history_detailed_by_user(winner_aid int64,offset,limit int) []p.CGPrizeHistory {
	
	var query string
	query = "SELECT "+
				"p.ptype AS record_type,"+
				"COALESCE(pc.evtlog_id, rew.evtlog_id, rnw.evtlog_id, ew.evtlog_id, lw.evtlog_id, cw.evtlog_id, ed.evtlog_id) AS evtlog_id,"+
				"COALESCE(EXTRACT(EPOCH FROM pc.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rnw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM lw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM cw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ed.time_stamp)::BIGINT) AS tstmp,"+
				"COALESCE(pc.time_stamp, rew.time_stamp, rnw.time_stamp, ew.time_stamp, lw.time_stamp, cw.time_stamp, ed.time_stamp) AS date_time,"+
				"COALESCE(pc.block_num, rew.block_num, rnw.block_num, ew.block_num, lw.block_num, cw.block_num, ed.block_num) AS block_num,"+
				"COALESCE(tc.id, trew.id, trnw.id, tew.id, tlw.id, tcw.id, ted.id) AS tx_id,"+
				"COALESCE(tc.tx_hash, trew.tx_hash, trnw.tx_hash, tew.tx_hash, tlw.tx_hash, tcw.tx_hash, ted.tx_hash) AS tx_hash,"+
				"p.round_num,"+
				"CASE "+
					"WHEN p.ptype = 0 THEN pc.amount "+
					"WHEN p.ptype = 1 THEN pc.cst_amount "+
					"WHEN p.ptype = 3 THEN rew.amount "+
					"WHEN p.ptype IN (4,6) THEN rnw.cst_amount "+
					"WHEN p.ptype IN (9,16) THEN COALESCE(ew.erc20_amount, lw.erc20_amount) "+
					"WHEN p.ptype = 10 THEN cw.eth_amount "+
					"WHEN p.ptype = 11 THEN cw.cst_amount "+
					"WHEN p.ptype = 13 THEN ed.deposit_amount "+
					"ELSE '0' "+
				"END AS amount,"+
				"CASE "+
					"WHEN p.ptype = 0 THEN pc.amount/1e18 "+
					"WHEN p.ptype = 1 THEN pc.cst_amount/1e18 "+
					"WHEN p.ptype = 3 THEN rew.amount/1e18 "+
					"WHEN p.ptype IN (4,6) THEN rnw.cst_amount/1e18 "+
					"WHEN p.ptype IN (9,16) THEN COALESCE(ew.erc20_amount, lw.erc20_amount)/1e18 "+
					"WHEN p.ptype = 10 THEN cw.eth_amount/1e18 "+
					"WHEN p.ptype = 11 THEN cw.cst_amount/1e18 "+
					"WHEN p.ptype = 13 THEN ed.deposit_amount/1e18 "+
					"ELSE 0 "+
				"END AS amount_eth,"+
				"'' AS token_addr,"+
				"CASE "+
					"WHEN p.ptype = 2 THEN pc.token_id "+
					"WHEN p.ptype IN (5,7) THEN rnw.token_id "+
					"WHEN p.ptype = 8 THEN ew.erc721_token_id "+
					"WHEN p.ptype = 12 THEN cw.nft_id "+
					"WHEN p.ptype = 15 THEN lw.erc721_token_id "+
					"ELSE -1 "+
				"END AS token_id,"+
				"'' AS token_uri,"+
			"p.winner_index,"+
			"'T' AS claimed "+
		"FROM "+sw.S.SchemaName()+".cg_prize p "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_prize_claim pc ON (p.round_num = pc.round_num AND p.ptype IN (0,1,2)) "+
			"LEFT JOIN "+sw.S.SchemaName()+".transaction tc ON tc.id = pc.tx_id "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_eth_winner rew ON (p.round_num = rew.round_num AND p.winner_index = rew.winner_idx AND p.ptype = 3) "+
			"LEFT JOIN "+sw.S.SchemaName()+".transaction trew ON trew.id = rew.tx_id "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_nft_winner rnw ON (p.round_num = rnw.round_num AND p.winner_index = rnw.winner_idx AND p.ptype IN (4,5,6,7)) "+
			"LEFT JOIN "+sw.S.SchemaName()+".transaction trnw ON trnw.id = rnw.tx_id "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_endurance_winner ew ON (p.round_num = ew.round_num AND p.winner_index = ew.winner_idx AND p.ptype IN (8,9)) "+
			"LEFT JOIN "+sw.S.SchemaName()+".transaction tew ON tew.id = ew.tx_id "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_chrono_warrior cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index AND p.ptype IN (10,11,12)) "+
			"LEFT JOIN "+sw.S.SchemaName()+".transaction tcw ON tcw.id = cw.tx_id "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_eth_deposit ed ON (p.round_num = ed.round_num AND p.ptype = 13) "+
			"LEFT JOIN "+sw.S.SchemaName()+".transaction ted ON ted.id = ed.tx_id "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_lastcst_winner lw ON (p.round_num = lw.round_num AND p.winner_index = lw.winner_idx AND p.ptype IN (15,16)) "+
			"LEFT JOIN "+sw.S.SchemaName()+".transaction tlw ON tlw.id = lw.tx_id "+
		"WHERE ("+
				"(p.ptype IN (0,1,2) AND pc.winner_aid = $1) OR "+
				"(p.ptype = 3 AND rew.winner_aid = $1) OR "+
				"(p.ptype IN (4,5,6,7) AND rnw.winner_aid = $1) OR "+
				"(p.ptype IN (8,9) AND ew.winner_aid = $1) OR "+
				"(p.ptype IN (10,11,12) AND cw.winner_aid = $1) OR "+
				"(p.ptype IN (15,16) AND lw.winner_aid = $1)"+
			") "+
		"ORDER BY p.round_num DESC, p.winner_index, p.ptype "+
		"OFFSET $2 LIMIT $3"
	rows,err := sw.S.Db().Query(query,winner_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGPrizeHistory,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGPrizeHistory
		err=rows.Scan(
			&rec.RecordType,
			&rec.EvtLogId,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.RoundNum,
			&rec.Amount,
			&rec.AmountEth,
			&rec.TokenAddress,
			&rec.TokenId,
			&rec.TokenURI,
			&rec.WinnerIndex,
			&rec.Claimed,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_claim_history_detailed_global(offset,limit int) []p.CGPrizeHistory {
	
	var query string
	query = "SELECT "+
				"p.ptype AS record_type,"+
				"COALESCE(pc.evtlog_id, rew.evtlog_id, rnw.evtlog_id, ew.evtlog_id, lw.evtlog_id, cw.evtlog_id, ed.evtlog_id) AS evtlog_id,"+
				"COALESCE(EXTRACT(EPOCH FROM pc.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rnw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM lw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM cw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ed.time_stamp)::BIGINT) AS tstmp,"+
				"COALESCE(pc.time_stamp, rew.time_stamp, rnw.time_stamp, ew.time_stamp, lw.time_stamp, cw.time_stamp, ed.time_stamp) AS date_time,"+
				"COALESCE(pc.block_num, rew.block_num, rnw.block_num, ew.block_num, lw.block_num, cw.block_num, ed.block_num) AS block_num,"+
				"COALESCE(tc.id, trew.id, trnw.id, tew.id, tlw.id, tcw.id, ted.id) AS tx_id,"+
				"COALESCE(tc.tx_hash, trew.tx_hash, trnw.tx_hash, tew.tx_hash, tlw.tx_hash, tcw.tx_hash, ted.tx_hash) AS tx_hash,"+
				"p.round_num,"+
				"CASE "+
					"WHEN p.ptype = 0 THEN pc.amount "+
					"WHEN p.ptype = 1 THEN pc.cst_amount "+
					"WHEN p.ptype = 3 THEN rew.amount "+
					"WHEN p.ptype IN (4,6) THEN rnw.cst_amount "+
					"WHEN p.ptype IN (9,16) THEN COALESCE(ew.erc20_amount, lw.erc20_amount) "+
					"WHEN p.ptype = 10 THEN cw.eth_amount "+
					"WHEN p.ptype = 11 THEN cw.cst_amount "+
					"WHEN p.ptype = 13 THEN ed.deposit_amount "+
					"ELSE '0' "+
				"END AS amount,"+
				"CASE "+
					"WHEN p.ptype = 0 THEN pc.amount/1e18 "+
					"WHEN p.ptype = 1 THEN pc.cst_amount/1e18 "+
					"WHEN p.ptype = 3 THEN rew.amount/1e18 "+
					"WHEN p.ptype IN (4,6) THEN rnw.cst_amount/1e18 "+
					"WHEN p.ptype IN (9,16) THEN COALESCE(ew.erc20_amount, lw.erc20_amount)/1e18 "+
					"WHEN p.ptype = 10 THEN cw.eth_amount/1e18 "+
					"WHEN p.ptype = 11 THEN cw.cst_amount/1e18 "+
					"WHEN p.ptype = 13 THEN ed.deposit_amount/1e18 "+
					"ELSE 0 "+
				"END AS amount_eth,"+
				"'' AS token_addr,"+
				"CASE "+
					"WHEN p.ptype = 2 THEN pc.token_id "+
					"WHEN p.ptype IN (5,7) THEN rnw.token_id "+
					"WHEN p.ptype = 8 THEN ew.erc721_token_id "+
					"WHEN p.ptype = 12 THEN cw.nft_id "+
					"WHEN p.ptype = 15 THEN lw.erc721_token_id "+
					"ELSE -1 "+
				"END AS token_id,"+
				"'' AS token_uri,"+
				"p.winner_index,"+
			"'T' AS claimed,"+
			"CASE WHEN p.ptype = 13 THEN '(All CS NFT Stakers)' ELSE COALESCE(wa_pc.addr, wa_rew.addr, wa_rnw.addr, wa_ew.addr, wa_lw.addr, wa_cw.addr, '') END AS winner_addr,"+
			"COALESCE(pc.winner_aid, rew.winner_aid, rnw.winner_aid, ew.winner_aid, lw.winner_aid, cw.winner_aid, 0) AS winner_aid "+
			"FROM "+sw.S.SchemaName()+".cg_prize p "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_prize_claim pc ON (p.round_num = pc.round_num AND p.ptype IN (0,1,2)) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction tc ON tc.id = pc.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_pc ON pc.winner_aid = wa_pc.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_eth_winner rew ON (p.round_num = rew.round_num AND p.winner_index = rew.winner_idx AND p.ptype = 3) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction trew ON trew.id = rew.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_rew ON rew.winner_aid = wa_rew.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_nft_winner rnw ON (p.round_num = rnw.round_num AND p.winner_index = rnw.winner_idx AND p.ptype IN (4,5,6,7)) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction trnw ON trnw.id = rnw.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_rnw ON rnw.winner_aid = wa_rnw.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_endurance_winner ew ON (p.round_num = ew.round_num AND p.winner_index = ew.winner_idx AND p.ptype IN (8,9)) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction tew ON tew.id = ew.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_ew ON ew.winner_aid = wa_ew.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_chrono_warrior cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index AND p.ptype IN (10,11,12)) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction tcw ON tcw.id = cw.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_cw ON cw.winner_aid = wa_cw.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_eth_deposit ed ON (p.round_num = ed.round_num AND p.ptype = 13) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction ted ON ted.id = ed.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_lastcst_winner lw ON (p.round_num = lw.round_num AND p.winner_index = lw.winner_idx AND p.ptype IN (15,16)) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction tlw ON tlw.id = lw.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_lw ON lw.winner_aid = wa_lw.address_id "+
			"ORDER BY p.round_num DESC, p.winner_index, p.ptype "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGPrizeHistory,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGPrizeHistory
		err=rows.Scan(
			&rec.RecordType,
			&rec.EvtLogId,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.RoundNum,
			&rec.Amount,
			&rec.AmountEth,
			&rec.TokenAddress,
			&rec.TokenId,
			&rec.TokenURI,
			&rec.WinnerIndex,
			&rec.Claimed,
			&rec.WinnerAddr,
			&rec.WinnerAid,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
