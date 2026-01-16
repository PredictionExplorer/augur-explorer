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
				"COALESCE(pc.evtlog_id, lw.evtlog_id, ew.evtlog_id, cw.evtlog_id, rew.evtlog_id, rnw_bidder.evtlog_id, rnw_rwalk.evtlog_id) AS evtlog_id,"+
				"COALESCE(EXTRACT(EPOCH FROM pc.time_stamp)::BIGINT, EXTRACT(EPOCH FROM lw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM cw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rnw_bidder.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rnw_rwalk.time_stamp)::BIGINT) AS tstmp,"+
				"COALESCE(pc.time_stamp, lw.time_stamp, ew.time_stamp, cw.time_stamp, rew.time_stamp, rnw_bidder.time_stamp, rnw_rwalk.time_stamp) AS date_time,"+
				"COALESCE(pc.block_num, lw.block_num, ew.block_num, cw.block_num, rew.block_num, rnw_bidder.block_num, rnw_rwalk.block_num) AS block_num,"+
				"COALESCE(tc.id, tlw.id, tew.id, tcw.id, trew.id, trnw_bidder.id, trnw_rwalk.id) AS tx_id,"+
				"COALESCE(tc.tx_hash, tlw.tx_hash, tew.tx_hash, tcw.tx_hash, trew.tx_hash, trnw_bidder.tx_hash, trnw_rwalk.tx_hash) AS tx_hash,"+
				"p.round_num,"+
				"CASE "+
					"WHEN p.ptype = 0 THEN pc.amount "+
					"WHEN p.ptype = 1 THEN pc.cst_amount "+
					"WHEN p.ptype = 4 THEN lw.erc20_amount "+
					"WHEN p.ptype = 6 THEN ew.erc20_amount "+
					"WHEN p.ptype = 7 THEN cw.eth_amount "+
					"WHEN p.ptype = 8 THEN cw.cst_amount "+
					"WHEN p.ptype = 10 THEN rew.amount "+
					"WHEN p.ptype = 11 THEN rnw_bidder.cst_amount "+
					"WHEN p.ptype = 13 THEN rnw_rwalk.cst_amount "+
					"ELSE '0' "+
				"END AS amount,"+
				"CASE "+
					"WHEN p.ptype = 0 THEN pc.amount/1e18 "+
					"WHEN p.ptype = 1 THEN pc.cst_amount/1e18 "+
					"WHEN p.ptype = 4 THEN lw.erc20_amount/1e18 "+
					"WHEN p.ptype = 6 THEN ew.erc20_amount/1e18 "+
					"WHEN p.ptype = 7 THEN cw.eth_amount/1e18 "+
					"WHEN p.ptype = 8 THEN cw.cst_amount/1e18 "+
					"WHEN p.ptype = 10 THEN rew.amount/1e18 "+
					"WHEN p.ptype = 11 THEN rnw_bidder.cst_amount/1e18 "+
					"WHEN p.ptype = 13 THEN rnw_rwalk.cst_amount/1e18 "+
					"ELSE 0 "+
				"END AS amount_eth,"+
				"'' AS token_addr,"+
				"CASE "+
					"WHEN p.ptype = 2 THEN pc.token_id "+
					"WHEN p.ptype = 3 THEN lw.erc721_token_id "+
					"WHEN p.ptype = 5 THEN ew.erc721_token_id "+
					"WHEN p.ptype = 9 THEN cw.nft_id "+
					"WHEN p.ptype = 12 THEN rnw_bidder.token_id "+
					"WHEN p.ptype = 14 THEN rnw_rwalk.token_id "+
					"ELSE -1 "+
				"END AS token_id,"+
			"'' AS token_uri,"+
		"p.winner_index,"+
		"CASE WHEN p.ptype = 10 THEN pd.claimed ELSE TRUE END AS claimed,"+
		"FALSE AS is_timeout_claim "+
	"FROM "+sw.S.SchemaName()+".cg_prize p "+
		"LEFT JOIN "+sw.S.SchemaName()+".cg_prize_claim pc ON (p.round_num = pc.round_num AND p.ptype IN (0,1,2)) "+
		"LEFT JOIN "+sw.S.SchemaName()+".transaction tc ON tc.id = pc.tx_id "+
		"LEFT JOIN "+sw.S.SchemaName()+".cg_lastcst_prize lw ON (p.round_num = lw.round_num AND p.winner_index = lw.winner_idx AND p.ptype IN (3,4)) "+
		"LEFT JOIN "+sw.S.SchemaName()+".transaction tlw ON tlw.id = lw.tx_id "+
		"LEFT JOIN "+sw.S.SchemaName()+".cg_endurance_prize ew ON (p.round_num = ew.round_num AND p.winner_index = ew.winner_idx AND p.ptype IN (5,6)) "+
		"LEFT JOIN "+sw.S.SchemaName()+".transaction tew ON tew.id = ew.tx_id "+
		"LEFT JOIN "+sw.S.SchemaName()+".cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index AND p.ptype IN (7,8,9)) "+
		"LEFT JOIN "+sw.S.SchemaName()+".transaction tcw ON tcw.id = cw.tx_id "+
		"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_eth_prize rew ON (p.round_num = rew.round_num AND p.winner_index = rew.winner_idx AND p.ptype = 10) "+
		"LEFT JOIN "+sw.S.SchemaName()+".transaction trew ON trew.id = rew.tx_id "+
		"LEFT JOIN "+sw.S.SchemaName()+".cg_prize_deposit pd ON (p.round_num = pd.round_num AND p.winner_index = pd.winner_index AND p.ptype = 10) "+
		"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_nft_prize rnw_bidder ON (p.round_num = rnw_bidder.round_num AND p.winner_index = rnw_bidder.winner_idx AND p.ptype IN (11,12) AND rnw_bidder.is_rwalk = false) "+
		"LEFT JOIN "+sw.S.SchemaName()+".transaction trnw_bidder ON trnw_bidder.id = rnw_bidder.tx_id "+
		"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_nft_prize rnw_rwalk ON (p.round_num = rnw_rwalk.round_num AND p.winner_index = rnw_rwalk.winner_idx AND p.ptype IN (13,14) AND rnw_rwalk.is_rwalk = true) "+
		"LEFT JOIN "+sw.S.SchemaName()+".transaction trnw_rwalk ON trnw_rwalk.id = rnw_rwalk.tx_id "+
		"WHERE ("+
				"(p.ptype IN (0,1,2) AND pc.winner_aid = $1) OR "+
				"(p.ptype IN (3,4) AND lw.winner_aid = $1) OR "+
				"(p.ptype IN (5,6) AND ew.winner_aid = $1) OR "+
				"(p.ptype IN (7,8,9) AND cw.winner_aid = $1) OR "+
				"(p.ptype = 10 AND rew.winner_aid = $1) OR "+
				"(p.ptype IN (11,12) AND rnw_bidder.winner_aid = $1) OR "+
				"(p.ptype IN (13,14) AND rnw_rwalk.winner_aid = $1)"+
			") "+
		
		"UNION "+
		
		// Add timeout-claimed ETH prizes (where user is beneficiary but not winner)
		"SELECT "+
			"18 AS record_type,"+  // ETH timeout claim (different from normal type 10)
			"pw.evtlog_id,"+
			"EXTRACT(EPOCH FROM pw.time_stamp)::BIGINT AS tstmp,"+
			"pw.time_stamp AS date_time,"+
			"pw.block_num,"+
			"t.id AS tx_id,"+
			"t.tx_hash,"+
			"pw.round_num,"+
			"pw.amount,"+
			"pw.amount/1e18 AS amount_eth,"+
			"'' AS token_addr,"+
			"-1 AS token_id,"+
			"'' AS token_uri,"+
			"0 AS winner_index,"+  // Winner index not available in withdrawal table
			"TRUE AS claimed,"+
			"TRUE AS is_timeout_claim "+
		"FROM "+sw.S.SchemaName()+".cg_prize_withdrawal pw "+
			"JOIN "+sw.S.SchemaName()+".transaction t ON t.id = pw.tx_id "+
		"WHERE pw.beneficiary_aid = $1 AND pw.beneficiary_aid != pw.winner_aid "+
		
		"UNION "+
		
		// Add timeout-claimed donated NFTs (where user is beneficiary but not round winner)
		"SELECT "+
			"16 AS record_type,"+  // Custom type for timeout-claimed donated NFT
			"nc.evtlog_id,"+
			"EXTRACT(EPOCH FROM nc.time_stamp)::BIGINT AS tstmp,"+
			"nc.time_stamp AS date_time,"+
			"nc.block_num,"+
			"t.id AS tx_id,"+
			"t.tx_hash,"+
			"nc.round_num,"+
			"'0' AS amount,"+
			"0 AS amount_eth,"+
			"wa_token.addr AS token_addr,"+
			"nc.token_id,"+
			"'' AS token_uri,"+
			"nc.idx AS winner_index,"+  // Using NFT index
			"TRUE AS claimed,"+
			"TRUE AS is_timeout_claim "+
		"FROM "+sw.S.SchemaName()+".cg_donated_nft_claimed nc "+
			"JOIN "+sw.S.SchemaName()+".transaction t ON t.id = nc.tx_id "+
			"JOIN "+sw.S.SchemaName()+".cg_prize_claim pc ON pc.round_num = nc.round_num "+
			"JOIN "+sw.S.SchemaName()+".address wa_token ON nc.token_aid = wa_token.address_id "+
		"WHERE nc.winner_aid = $1 AND nc.winner_aid != pc.winner_aid "+
		
		"UNION "+
		
		// Add timeout-claimed donated ERC20 (where user is beneficiary but not round winner)
		"SELECT "+
			"17 AS record_type,"+  // Custom type for timeout-claimed donated ERC20
			"tc.evtlog_id,"+
			"EXTRACT(EPOCH FROM tc.time_stamp)::BIGINT AS tstmp,"+
			"tc.time_stamp AS date_time,"+
			"tc.block_num,"+
			"t.id AS tx_id,"+
			"t.tx_hash,"+
			"tc.round_num,"+
			"tc.amount AS amount,"+
			"tc.amount/1e18 AS amount_eth,"+
			"wa_token.addr AS token_addr,"+
			"-1 AS token_id,"+
			"'' AS token_uri,"+
			"0 AS winner_index,"+
			"TRUE AS claimed,"+
			"TRUE AS is_timeout_claim "+
		"FROM "+sw.S.SchemaName()+".cg_donated_tok_claimed tc "+
			"JOIN "+sw.S.SchemaName()+".transaction t ON t.id = tc.tx_id "+
			"JOIN "+sw.S.SchemaName()+".cg_prize_claim pc ON pc.round_num = tc.round_num "+
			"JOIN "+sw.S.SchemaName()+".address wa_token ON tc.token_aid = wa_token.address_id "+
		"WHERE tc.winner_aid = $1 AND tc.winner_aid != pc.winner_aid "+
		
		"ORDER BY round_num DESC, winner_index, record_type "+
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
			&rec.Tx.EvtLogId,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.RoundNum,
			&rec.Amount,
			&rec.AmountEth,
			&rec.TokenAddress,
			&rec.TokenId,
			&rec.TokenURI,
			&rec.WinnerIndex,
			&rec.Claimed,
			&rec.IsTimeoutClaim,
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
				"COALESCE(pc.evtlog_id, lw.evtlog_id, ew.evtlog_id, cw.evtlog_id, rew.evtlog_id, rnw_bidder.evtlog_id, rnw_rwalk.evtlog_id, ed.evtlog_id) AS evtlog_id,"+
				"COALESCE(EXTRACT(EPOCH FROM pc.time_stamp)::BIGINT, EXTRACT(EPOCH FROM lw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM cw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rnw_bidder.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rnw_rwalk.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ed.time_stamp)::BIGINT) AS tstmp,"+
				"COALESCE(pc.time_stamp, lw.time_stamp, ew.time_stamp, cw.time_stamp, rew.time_stamp, rnw_bidder.time_stamp, rnw_rwalk.time_stamp, ed.time_stamp) AS date_time,"+
				"COALESCE(pc.block_num, lw.block_num, ew.block_num, cw.block_num, rew.block_num, rnw_bidder.block_num, rnw_rwalk.block_num, ed.block_num) AS block_num,"+
				"COALESCE(tc.id, tlw.id, tew.id, tcw.id, trew.id, trnw_bidder.id, trnw_rwalk.id, ted.id) AS tx_id,"+
				"COALESCE(tc.tx_hash, tlw.tx_hash, tew.tx_hash, tcw.tx_hash, trew.tx_hash, trnw_bidder.tx_hash, trnw_rwalk.tx_hash, ted.tx_hash) AS tx_hash,"+
				"p.round_num,"+
				"CASE "+
					"WHEN p.ptype = 0 THEN pc.amount "+
					"WHEN p.ptype = 1 THEN pc.cst_amount "+
					"WHEN p.ptype = 4 THEN lw.erc20_amount "+
					"WHEN p.ptype = 6 THEN ew.erc20_amount "+
					"WHEN p.ptype = 7 THEN cw.eth_amount "+
					"WHEN p.ptype = 8 THEN cw.cst_amount "+
					"WHEN p.ptype = 10 THEN rew.amount "+
					"WHEN p.ptype = 11 THEN rnw_bidder.cst_amount "+
					"WHEN p.ptype = 13 THEN rnw_rwalk.cst_amount "+
					"WHEN p.ptype = 15 THEN ed.deposit_amount "+
					"ELSE '0' "+
				"END AS amount,"+
				"CASE "+
					"WHEN p.ptype = 0 THEN pc.amount/1e18 "+
					"WHEN p.ptype = 1 THEN pc.cst_amount/1e18 "+
					"WHEN p.ptype = 4 THEN lw.erc20_amount/1e18 "+
					"WHEN p.ptype = 6 THEN ew.erc20_amount/1e18 "+
					"WHEN p.ptype = 7 THEN cw.eth_amount/1e18 "+
					"WHEN p.ptype = 8 THEN cw.cst_amount/1e18 "+
					"WHEN p.ptype = 10 THEN rew.amount/1e18 "+
					"WHEN p.ptype = 11 THEN rnw_bidder.cst_amount/1e18 "+
					"WHEN p.ptype = 13 THEN rnw_rwalk.cst_amount/1e18 "+
					"WHEN p.ptype = 15 THEN ed.deposit_amount/1e18 "+
					"ELSE 0 "+
				"END AS amount_eth,"+
				"'' AS token_addr,"+
				"CASE "+
					"WHEN p.ptype = 2 THEN pc.token_id "+
					"WHEN p.ptype = 3 THEN lw.erc721_token_id "+
					"WHEN p.ptype = 5 THEN ew.erc721_token_id "+
					"WHEN p.ptype = 9 THEN cw.nft_id "+
					"WHEN p.ptype = 12 THEN rnw_bidder.token_id "+
					"WHEN p.ptype = 14 THEN rnw_rwalk.token_id "+
					"ELSE -1 "+
				"END AS token_id,"+
			"'' AS token_uri,"+
			"p.winner_index,"+
		"CASE WHEN p.ptype = 10 THEN pd.claimed ELSE TRUE END AS claimed,"+
		"FALSE AS is_timeout_claim,"+
		"CASE WHEN p.ptype = 15 THEN '(All CS NFT Stakers)' ELSE COALESCE(wa_pc.addr, wa_lw.addr, wa_ew.addr, wa_cw.addr, wa_rew.addr, wa_rnw_bidder.addr, wa_rnw_rwalk.addr, '') END AS winner_addr,"+
			"COALESCE(pc.winner_aid, lw.winner_aid, ew.winner_aid, cw.winner_aid, rew.winner_aid, rnw_bidder.winner_aid, rnw_rwalk.winner_aid, 0) AS winner_aid "+
		"FROM "+sw.S.SchemaName()+".cg_prize p "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_prize_claim pc ON (p.round_num = pc.round_num AND p.ptype IN (0,1,2)) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction tc ON tc.id = pc.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_pc ON pc.winner_aid = wa_pc.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_lastcst_prize lw ON (p.round_num = lw.round_num AND p.winner_index = lw.winner_idx AND p.ptype IN (3,4)) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction tlw ON tlw.id = lw.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_lw ON lw.winner_aid = wa_lw.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_endurance_prize ew ON (p.round_num = ew.round_num AND p.winner_index = ew.winner_idx AND p.ptype IN (5,6)) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction tew ON tew.id = ew.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_ew ON ew.winner_aid = wa_ew.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_chrono_warrior_prize cw ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index AND p.ptype IN (7,8,9)) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction tcw ON tcw.id = cw.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_cw ON cw.winner_aid = wa_cw.address_id "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_eth_prize rew ON (p.round_num = rew.round_num AND p.winner_index = rew.winner_idx AND p.ptype = 10) "+
			"LEFT JOIN "+sw.S.SchemaName()+".transaction trew ON trew.id = rew.tx_id "+
			"LEFT JOIN "+sw.S.SchemaName()+".address wa_rew ON rew.winner_aid = wa_rew.address_id "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_prize_deposit pd ON (p.round_num = pd.round_num AND p.winner_index = pd.winner_index AND p.ptype = 10) "+
			"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_nft_prize rnw_bidder ON (p.round_num = rnw_bidder.round_num AND p.winner_index = rnw_bidder.winner_idx AND p.ptype IN (11,12) AND rnw_bidder.is_rwalk = false) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction trnw_bidder ON trnw_bidder.id = rnw_bidder.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_rnw_bidder ON rnw_bidder.winner_aid = wa_rnw_bidder.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_nft_prize rnw_rwalk ON (p.round_num = rnw_rwalk.round_num AND p.winner_index = rnw_rwalk.winner_idx AND p.ptype IN (13,14) AND rnw_rwalk.is_rwalk = true) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction trnw_rwalk ON trnw_rwalk.id = rnw_rwalk.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address wa_rnw_rwalk ON rnw_rwalk.winner_aid = wa_rnw_rwalk.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_staking_eth_deposit ed ON (p.round_num = ed.round_num AND p.ptype = 15) "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction ted ON ted.id = ed.tx_id "+
			
			"UNION "+
			
		// Add timeout-claimed ETH prizes (global view)
		"SELECT "+
			"18 AS record_type,"+  // ETH timeout claim
				"pw.evtlog_id,"+
				"EXTRACT(EPOCH FROM pw.time_stamp)::BIGINT AS tstmp,"+
				"pw.time_stamp AS date_time,"+
				"pw.block_num,"+
				"t.id AS tx_id,"+
				"t.tx_hash,"+
				"pw.round_num,"+
				"pw.amount,"+
				"pw.amount/1e18 AS amount_eth,"+
				"'' AS token_addr,"+
				"-1 AS token_id,"+
				"'' AS token_uri,"+
				"0 AS winner_index,"+
				"TRUE AS claimed,"+
				"TRUE AS is_timeout_claim,"+
				"wa_winner.addr AS winner_addr,"+
				"pw.winner_aid "+
			"FROM "+sw.S.SchemaName()+".cg_prize_withdrawal pw "+
				"JOIN "+sw.S.SchemaName()+".transaction t ON t.id = pw.tx_id "+
				"JOIN "+sw.S.SchemaName()+".address wa_winner ON pw.winner_aid = wa_winner.address_id "+
			"WHERE pw.beneficiary_aid != pw.winner_aid "+
			
			"UNION "+
			
			// Add timeout-claimed donated NFTs (global view)
			"SELECT "+
				"16 AS record_type,"+
				"nc.evtlog_id,"+
				"EXTRACT(EPOCH FROM nc.time_stamp)::BIGINT AS tstmp,"+
				"nc.time_stamp AS date_time,"+
				"nc.block_num,"+
				"t.id AS tx_id,"+
				"t.tx_hash,"+
				"nc.round_num,"+
				"'0' AS amount,"+
				"0 AS amount_eth,"+
				"wa_token.addr AS token_addr,"+
				"nc.token_id,"+
				"'' AS token_uri,"+
				"nc.idx AS winner_index,"+
				"TRUE AS claimed,"+
				"TRUE AS is_timeout_claim,"+
				"wa_beneficiary.addr AS winner_addr,"+  // Show beneficiary as winner for these records
				"nc.winner_aid "+
			"FROM "+sw.S.SchemaName()+".cg_donated_nft_claimed nc "+
				"JOIN "+sw.S.SchemaName()+".transaction t ON t.id = nc.tx_id "+
				"JOIN "+sw.S.SchemaName()+".cg_prize_claim pc ON pc.round_num = nc.round_num "+
				"JOIN "+sw.S.SchemaName()+".address wa_token ON nc.token_aid = wa_token.address_id "+
				"JOIN "+sw.S.SchemaName()+".address wa_beneficiary ON nc.winner_aid = wa_beneficiary.address_id "+
			"WHERE nc.winner_aid != pc.winner_aid "+
			
			"UNION "+
			
			// Add timeout-claimed donated ERC20 (global view)
			"SELECT "+
				"17 AS record_type,"+
				"tc.evtlog_id,"+
				"EXTRACT(EPOCH FROM tc.time_stamp)::BIGINT AS tstmp,"+
				"tc.time_stamp AS date_time,"+
				"tc.block_num,"+
				"t.id AS tx_id,"+
				"t.tx_hash,"+
				"tc.round_num,"+
				"tc.amount AS amount,"+
				"tc.amount/1e18 AS amount_eth,"+
				"wa_token.addr AS token_addr,"+
				"-1 AS token_id,"+
				"'' AS token_uri,"+
				"0 AS winner_index,"+
				"TRUE AS claimed,"+
				"TRUE AS is_timeout_claim,"+
				"wa_beneficiary.addr AS winner_addr,"+  // Show beneficiary as winner for these records
				"tc.winner_aid "+
			"FROM "+sw.S.SchemaName()+".cg_donated_tok_claimed tc "+
				"JOIN "+sw.S.SchemaName()+".transaction t ON t.id = tc.tx_id "+
				"JOIN "+sw.S.SchemaName()+".cg_prize_claim pc ON pc.round_num = tc.round_num "+
				"JOIN "+sw.S.SchemaName()+".address wa_token ON tc.token_aid = wa_token.address_id "+
				"JOIN "+sw.S.SchemaName()+".address wa_beneficiary ON tc.winner_aid = wa_beneficiary.address_id "+
			"WHERE tc.winner_aid != pc.winner_aid "+
			
			"ORDER BY round_num DESC, winner_index, record_type "+
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
			&rec.Tx.EvtLogId,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.RoundNum,
			&rec.Amount,
			&rec.AmountEth,
			&rec.TokenAddress,
			&rec.TokenId,
			&rec.TokenURI,
			&rec.WinnerIndex,
			&rec.Claimed,
			&rec.IsTimeoutClaim,
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
