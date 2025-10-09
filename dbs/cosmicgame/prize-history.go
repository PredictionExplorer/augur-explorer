package cosmicgame

import (
	"os"
	"fmt"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_prize_history_detailed_by_user(winner_aid int64,offset,limit int) []p.CGPrizeHistory {
	
	var query string
	query = "SELECT "+
				"record_type,"+
				"evtlog_id,"+
				"tstmp,"+
				"date_time,"+
				"block_num,"+
				"tx_id,"+
				"tx_hash,"+
				"round_num,"+
				"amount,"+
				"amount_eth,"+
				"token_addr,"+
				"token_id," +
				"token_uri,"+
				"winner_index, "+
				"claimed "+
			"FROM (" +
				"(" +
					"SELECT "+
						"0 AS record_type,"+
						"rd.evtlog_id,"+
						"EXTRACT(EPOCH FROM rd.time_stamp)::BIGINT AS tstmp, "+
						"rd.time_stamp AS date_time, "+
						"rd.block_num,"+
						"rd.tx_id,"+
						"t.tx_hash,"+
						"rd.round_num,"+
						"rd.amount, "+
						"rd.amount/1e18 AS amount_eth,"+
						"'' AS token_addr, "+
						"-1 AS token_id,"+
						"'' AS token_uri,"+
						"rd.winner_index, "+
						"rd.claimed "+
					"FROM cg_prize_deposit rd "+
						"LEFT JOIN transaction t ON t.id=rd.tx_id "+
						"INNER JOIN cg_raffle_eth_winner rew ON (rew.round_num=rd.round_num AND rew.winner_idx=rd.winner_index) "+
					"WHERE rd.winner_aid=$1  "+
				") UNION ALL (" +
					"SELECT "+
						"1 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"'' AS token_addr, "+
						"rn.token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed "+
					"FROM cg_raffle_nft_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
					"WHERE (rn.winner_aid=$1) AND (is_rwalk=FALSE) "+
				") UNION ALL (" +
					"SELECT "+
						"2 AS record_type,"+
						"d.evtlog_id,"+
						"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS tstmp, "+
						"p.time_stamp AS date_time, "+
						"p.block_num,"+
						"p.tx_id,"+
						"t.tx_hash,"+
						"p.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"ta.addr token_addr, " +
						"d.token_id,"+
						"d.token_uri,"+
						"d.idx winner_index,"+
						"c.id IS NOT NULL as claimed "+
					"FROM cg_prize_claim p "+
						"JOIN cg_nft_donation d ON p.round_num=d.round_num "+ 
						"LEFT JOIN transaction t ON t.id=p.tx_id "+
						"LEFT JOIN address ta ON d.token_aid=ta.address_id "+
						"LEFT JOIN cg_donated_nft_claimed c ON (c.round_num=p.round_num) AND (d.idx=c.idx) "+
					"WHERE p.winner_aid=$1 "+
				") UNION ALL (" +
					"SELECT "+
						"3 AS record_type,"+
						"p.evtlog_id,"+
						"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS tstmp, "+
						"p.time_stamp AS date_time, "+
						"p.block_num,"+
						"p.tx_id,"+
						"t.tx_hash,"+
						"p.round_num,"+
						"p.amount,"+
						"p.amount/1e18 AS amount_eth,"+
						"'' AS token_addr, " +
						"p.token_id,"+
						"'' AS token_uri,"+
						"-1 AS winner_index,"+
						"'T' AS claimed "+
					"FROM cg_prize_claim p "+
						"LEFT JOIN transaction t ON t.id=p.tx_id "+
						"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
					"WHERE p.winner_aid=$1 "+
				") UNION ALL (" +
					"WITH rwd AS ("+
						"SELECT "+
							"COUNT(token_id) AS num_toks_collected,"+
							"SUM(reward) AS collected_reward," +
							"SUM(reward)/1e18 AS collected_reward_eth,"+
							"deposit_id, "+
							"staker_aid "+
						"FROM cg_st_reward "+
						"GROUP BY staker_aid,deposit_id "+
					") "+
					"SELECT "+
						"4 AS record_type,"+
						"d.evtlog_id,"+
						"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
						"d.time_stamp,"+
						"d.block_num,"+
						"tx.id,"+
						"tx.tx_hash,"+
						"d.round_num, "+
						"sd.amount_to_claim,"+
						"sd.amount_to_claim/1e18,"+
						"'' AS token_addr, "+
						"-1 AS token_id,"+
						"'' AS token_uri,"+
						"-1 AS winner_index, "+
						"CASE "+
							"WHEN "+
								"COALESCE(rwd.collected_reward,0)=COALESCE(sd.amount_to_claim,0)"+
								"THEN TRUE "+
								"ELSE FALSE "+
						"END AS claimed "+
					"FROM "+sw.S.SchemaName()+".cg_staker_deposit sd "+
						"INNER JOIN cg_eth_deposit d ON sd.deposit_id=d.deposit_id "+
						"INNER JOIN transaction tx ON tx.id=d.tx_id " +
						"LEFT JOIN rwd ON (rwd.deposit_id=sd.deposit_id) AND (rwd.staker_aid=sd.staker_aid) "+
						"INNER JOIN address sa ON sd.staker_aid = sa.address_id "+
					"WHERE sd.staker_aid=$1 "+
					"ORDER BY d.id DESC,sd.staker_aid " +
				") UNION ALL (" +
					"SELECT "+
						"5 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"'' AS token_addr, "+
						"rn.token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed "+
					"FROM cg_raffle_nft_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
					"WHERE (rn.winner_aid=$1) AND (is_rwalk=TRUE) AND (is_staker=TRUE) "+
				") UNION ALL (" +
					"SELECT "+
						"6 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"'' AS token_addr, "+
						"rn.token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed "+
					"FROM cg_raffle_nft_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
					"WHERE (rn.winner_aid=$1) AND (is_rwalk=FALSE) AND (is_staker=TRUE) "+
				") UNION ALL (" +
					"SELECT "+
						"7 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"'' AS token_addr, "+
						"rn.erc721_token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed "+
					"FROM cg_endurance_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
					"WHERE (rn.winner_aid=$1) "+
				") UNION ALL (" +
					"SELECT "+
						"8 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"'' AS token_addr, "+
						"rn.erc721_token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed "+
					"FROM cg_lastcst_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
					"WHERE (rn.winner_aid=$1) "+
				") UNION ALL (" +
					"SELECT "+
						"9 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"erc20_amount AS amount,"+
						"erc20_amount/1e18 AS amount_eth,"+
						"'' AS token_addr, "+
						"-1 AS token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed "+
					"FROM cg_endurance_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
					"WHERE (rn.winner_aid=$1) "+
				") UNION ALL (" +
					"SELECT "+
						"10 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"erc20_amount AS amount,"+
						"erc20_amount/1e18 AS amount_eth,"+
						"'' AS token_addr, "+
						"-1 AS token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed "+
					"FROM cg_lastcst_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
					"WHERE (rn.winner_aid=$1) "+
				") UNION ALL (" +
					"SELECT "+
						"11 AS record_type,"+
						"d.evtlog_id,"+
						"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS tstmp, "+
						"p.time_stamp AS date_time, "+
						"p.block_num,"+
						"p.tx_id,"+
						"t.tx_hash,"+
						"p.round_num,"+
						"d.amount AS amount,"+
						"d.amount/1e18 AS amount_eth,"+
						"ta.addr token_addr, " +
						"-1 AS token_id,"+
						"'' AS token_uri,"+
						"-1 AS winner_index,"+
						"c.id IS NOT NULL as claimed "+
					"FROM cg_prize_claim p "+
						"JOIN cg_erc20_donation d ON p.round_num=d.round_num "+ 
						"LEFT JOIN transaction t ON t.id=p.tx_id "+
						"LEFT JOIN address ta ON d.token_aid=ta.address_id "+
						"LEFT JOIN cg_donated_tok_claimed c ON (c.round_num=p.round_num) AND (c.token_aid=d.token_aid)"+
					"WHERE p.winner_aid=$1 "+
				") UNION ALL (" +
					"SELECT "+
						"12 AS record_type,"+
						"cw.evtlog_id,"+
						"EXTRACT(EPOCH FROM rd.time_stamp)::BIGINT AS tstmp, "+
						"rd.time_stamp AS date_time, "+
						"rd.block_num,"+
						"rd.tx_id,"+
						"t.tx_hash,"+
						"rd.round_num,"+
						"rd.amount AS amount,"+
						"rd.amount/1e18 AS amount_eth,"+
						"'' token_addr, " +
						"-1 AS token_id,"+
						"'' AS token_uri,"+
						"rd.winner_index,"+
						"'T' as claimed "+
					"FROM cg_prize_deposit rd "+
						"LEFT JOIN transaction t ON t.id=rd.tx_id "+
						"INNER JOIN cg_chrono_warrior cw ON (cw.round_num=rd.round_num AND cw.winner_index=rd.winner_index) "+
					"WHERE rd.winner_aid=$1 "+
				") "+
			") everything " +
			"ORDER BY evtlog_id DESC " +
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
				"record_type,"+
				"evtlog_id,"+
				"tstmp,"+
				"date_time,"+
				"block_num,"+
				"tx_id,"+
				"tx_hash,"+
				"round_num,"+
				"amount,"+
				"amount_eth,"+
				"token_addr,"+
				"token_id," +
				"token_uri,"+
				"winner_index, "+
				"claimed, "+
				"winner_addr,"+
				"winner_aid "+
			"FROM (" +
				"(" +
					"SELECT "+
						"0 AS record_type,"+
						"rd.evtlog_id,"+
						"EXTRACT(EPOCH FROM rd.time_stamp)::BIGINT AS tstmp, "+
						"rd.time_stamp AS date_time, "+
						"rd.block_num,"+
						"rd.tx_id,"+
						"t.tx_hash,"+
						"rd.round_num,"+
						"rd.amount, "+
						"rd.amount/1e18 AS amount_eth,"+
						"'' AS token_addr, "+
						"-1 AS token_id,"+
						"'' AS token_uri,"+
						"-1 AS winner_index, "+
						"rd.claimed, "+
						"wa.addr winner_addr," +
						"rd.winner_aid "+
					"FROM cg_prize_deposit rd "+
						"LEFT JOIN transaction t ON t.id=rd.tx_id "+
						"LEFT JOIN address wa ON rd.winner_aid=wa.address_id "+
				") UNION ALL (" +
					"SELECT "+
						"1 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"'' AS token_addr, "+
						"rn.token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed, "+
						"wa.addr winner_addr,"+
						"rn.winner_aid "+
					"FROM cg_raffle_nft_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
						"LEFT JOIN address wa ON rn.winner_aid=wa.address_id "+
					"WHERE (is_rwalk=FALSE) AND (is_staker=FALSE) " +
				") UNION ALL (" +
					"SELECT "+
						"2 AS record_type,"+
						"d.evtlog_id,"+
						"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS tstmp, "+
						"p.time_stamp AS date_time, "+
						"p.block_num,"+
						"p.tx_id,"+
						"t.tx_hash,"+
						"p.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"ta.addr token_addr, " +
						"d.token_id,"+
						"d.token_uri,"+
						"d.idx winner_index,"+
						"c.id IS NOT NULL as claimed, "+
						"wa.addr winner_addr,"+
						"p.winner_aid "+
					"FROM cg_prize_claim p "+
						"JOIN cg_nft_donation d ON p.round_num=d.round_num "+ 
						"LEFT JOIN transaction t ON t.id=p.tx_id "+
						"LEFT JOIN address ta ON d.token_aid=ta.address_id "+
						"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
						"LEFT JOIN cg_donated_nft_claimed c ON (c.round_num=p.round_num) AND (d.idx=c.idx) "+
				") UNION ALL (" +
					"SELECT "+
						"3 AS record_type,"+
						"p.evtlog_id,"+
						"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS tstmp, "+
						"p.time_stamp AS date_time, "+
						"p.block_num,"+
						"p.tx_id,"+
						"t.tx_hash,"+
						"p.round_num,"+
						"p.amount,"+
						"p.amount/1e18 AS amount_eth,"+
						"'' AS token_addr, " +
						"p.token_id,"+
						"'' AS token_uri,"+
						"-1 AS winner_index,"+
						"'T' AS claimed, "+
						"wa.addr winner_addr,"+
						"p.winner_aid "+
					"FROM cg_prize_claim p "+
						"LEFT JOIN transaction t ON t.id=p.tx_id "+
						"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
				") UNION ALL (" +
					"WITH rwd AS ("+
						"SELECT "+
							"COUNT(token_id) AS num_toks_collected,"+
							"SUM(reward) AS collected_reward," +
							"SUM(reward)/1e18 AS collected_reward_eth,"+
							"deposit_id, "+
							"staker_aid "+
						"FROM cg_st_reward "+
						"GROUP BY staker_aid,deposit_id "+
					") "+
					"SELECT "+
						"4 AS record_type,"+
						"d.evtlog_id,"+
						"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
						"d.time_stamp,"+
						"d.block_num,"+
						"tx.id,"+
						"tx.tx_hash,"+
						"d.round_num, "+
						"sd.amount_to_claim,"+
						"sd.amount_to_claim/1e18,"+
						"'' AS token_addr, "+
						"-1 AS token_id,"+
						"'' AS token_uri,"+
						"-1 AS winner_index, "+
						"CASE "+
							"WHEN "+
								"COALESCE(rwd.collected_reward,0)=COALESCE(sd.amount_to_claim,0)"+
								"THEN TRUE "+
								"ELSE FALSE "+
						"END AS claimed, "+
						"sa.addr winner_addr,"+
						"sd.staker_aid winner_aid "+
					"FROM "+sw.S.SchemaName()+".cg_staker_deposit sd "+
						"INNER JOIN cg_eth_deposit d ON sd.deposit_id=d.deposit_id "+
						"INNER JOIN transaction tx ON tx.id=d.tx_id " +
						"LEFT JOIN rwd ON (rwd.deposit_id=sd.deposit_id) AND (rwd.staker_aid=sd.staker_aid) "+
						"INNER JOIN address sa ON sd.staker_aid = sa.address_id "+
					"ORDER BY d.id DESC,sd.staker_aid " +
				") UNION ALL (" +
					"SELECT "+
						"5 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"'' AS token_addr, "+
						"rn.token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed, "+
						"wa.addr winner_addr,"+
						"rn.winner_aid "+
					"FROM cg_raffle_nft_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
						"LEFT JOIN address wa ON rn.winner_aid=wa.address_id "+
					"WHERE (is_rwalk=TRUE) AND (is_staker=TRUE) "+
				") UNION ALL (" +
					"SELECT "+
						"6 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"'' AS token_addr, "+
						"rn.token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed, "+
						"wa.addr winner_addr,"+
						"rn.winner_aid "+
					"FROM cg_raffle_nft_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
						"LEFT JOIN address wa ON rn.winner_aid=wa.address_id "+
					"WHERE (is_rwalk=FALSE) AND (is_staker=TRUE) "+
				") UNION ALL (" +
					"SELECT "+
						"7 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"'' AS token_addr, "+
						"rn.erc721_token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed, "+
						"wa.addr winner_addr,"+
						"rn.winner_aid "+
					"FROM cg_endurance_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
						"LEFT JOIN address wa ON rn.winner_aid=wa.address_id "+
				") UNION ALL (" +
					"SELECT "+
						"8 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"'' AS token_addr, "+
						"rn.erc721_token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed, "+
						"wa.addr winner_addr,"+
						"rn.winner_aid "+
					"FROM cg_lastcst_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
						"LEFT JOIN address wa ON rn.winner_aid=wa.address_id "+
				") UNION ALL (" +
					"SELECT "+
						"9 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"erc20_amount AS amount,"+
						"erc20_amount/1e18 AS amount_eth,"+
						"'' AS token_addr, "+
						"-1 as token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed, "+
						"wa.addr winner_addr,"+
						"rn.winner_aid "+
					"FROM cg_endurance_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
						"LEFT JOIN address wa ON rn.winner_aid=wa.address_id "+
				") UNION ALL (" +
					"SELECT "+
						"10 AS record_type,"+
						"rn.evtlog_id,"+
						"EXTRACT(EPOCH FROM rn.time_stamp)::BIGINT AS tstmp, "+
						"rn.time_stamp AS date_time, "+
						"rn.block_num,"+
						"rn.tx_id,"+
						"t.tx_hash,"+
						"rn.round_num,"+
						"erc20_amount AS amount,"+
						"erc20_amount/1e18 AS amount_eth,"+
						"'' AS token_addr, "+
						"-1 as token_id," +
						"'' AS token_uri,"+
						"rn.winner_idx, "+
						"'T' AS claimed, "+
						"wa.addr winner_addr,"+
						"rn.winner_aid "+
					"FROM cg_lastcst_winner rn "+
						"LEFT JOIN transaction t ON t.id=rn.tx_id "+
						"LEFT JOIN address wa ON rn.winner_aid=wa.address_id "+
				") UNION ALL (" +
					"SELECT "+
						"11 AS record_type,"+
						"d.evtlog_id,"+
						"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS tstmp, "+
						"p.time_stamp AS date_time, "+
						"p.block_num,"+
						"p.tx_id,"+
						"t.tx_hash,"+
						"p.round_num,"+
						"d.amount AS amount,"+
						"d.amount/1e18 AS amount_eth,"+
						"ta.addr token_addr, " +
						"-1 AS token_id,"+
						"'' AS token_uri,"+
						"-1 AS winner_index,"+
						"c.id IS NOT NULL as claimed, "+
						"wa.addr winner_addr,"+
						"p.winner_aid "+
					"FROM cg_prize_claim p "+
						"JOIN cg_erc20_donation d ON p.round_num=d.round_num "+ 
						"LEFT JOIN transaction t ON t.id=p.tx_id "+
						"LEFT JOIN address ta ON d.token_aid=ta.address_id "+
						"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
						"LEFT JOIN cg_donated_tok_claimed c ON (c.round_num=p.round_num) AND (c.token_aid=d.token_aid)"+
				") UNION ALL (" +
					"SELECT "+
						"12 AS record_type,"+
						"cw.evtlog_id,"+
						"EXTRACT(EPOCH FROM rd.time_stamp)::BIGINT AS tstmp, "+
						"rd.time_stamp AS date_time, "+
						"rd.block_num,"+
						"rd.tx_id,"+
						"t.tx_hash,"+
						"rd.round_num,"+
						"rd.amount, "+
						"rd.amount/1e18 AS amount_eth,"+
						"'' AS token_addr, "+
						"-1 AS token_id,"+
						"'' AS token_uri,"+
						"rd.winner_index, "+
						"'T' AS claimed, "+
						"wa.addr winner_addr," +
						"rd.winner_aid "+
					"FROM cg_prize_deposit rd "+
						"LEFT JOIN transaction t ON t.id=rd.tx_id "+
						"INNER JOIN cg_chrono_warrior cw ON (cw.round_num=rd.round_num AND cw.winner_index=rd.winner_index) "+
						"LEFT JOIN address wa ON rd.winner_aid=wa.address_id "+
				") "+
			") everything " +
			"ORDER BY evtlog_id DESC " +
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
