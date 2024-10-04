package cosmicgame

import (
	"os"
	"fmt"
	"math"

	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_system_mode_change_event_list(offset,limit int) []p.CGSystemModeRec {

	if limit == 0 { limit = 1000000 }
	var add_deployment_events bool
	if offset == -1 {
		add_deployment_events = true
		offset = 0
	}
	var query string
	query = "SELECT "+
				"s.id,"+
				"s.evtlog_id,"+
				"s.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT,"+
				"s.time_stamp, "+
				"s.sysmode "+
			"FROM "+sw.S.SchemaName()+".cg_adm_sysmode s "+
				"LEFT JOIN transaction t ON t.id=s.tx_id "+
			"ORDER BY s.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGSystemModeRec,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGSystemModeRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.SystemMode,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	var next_evtlog int64 = math.MaxInt64;
	for i:=0; i<len(records); i++ {
		r := records[i];
		r.NextEvtLogId = next_evtlog
		next_evtlog = r.EvtLogId
		records[i]=r
	}
	if add_deployment_events {
		if len(records) > 0 {
			var rec p.CGSystemModeRec
			rec.RecordId = -1
			rec.EvtLogId = -1
			rec.BlockNum = -1
			rec.TxId = -1
			rec.SystemMode = -1
			rec.NextEvtLogId = records[len(records)-1].EvtLogId
			records = append(records,rec)
		}
	}
	return records
}
func (sw *SQLStorageWrapper) Get_admin_events_in_range(evtlog_start,evtlog_end int64) []p.CGAdminEvent {

	var query string
	query = "SELECT "+
				"record_type,"+
				"record_id,"+
				"evtlog_id,"+
				"block_num,"+
				"tx_id,"+
				"tx_hash,"+
				"ts,"+
				"date_time,"+
				"addr_value, "+ 
				"int_value "+
				"float_value "+
				"string_value "+
			"FROM ("+
				"("+
					"SELECT "+
						"1 AS record_type,"+			// CharityPercentageChanged
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"r.percentage AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_charity_pcent r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"2 AS record_type,"+			// PrizePercentageChanged
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"r.percentage AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_prize_pcent r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"3 AS record_type,"+			// RafflePercentageChanged
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"r.percentage AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_raffle_pcent r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"4 AS record_type,"+			//  StakingPercentageChanged
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"r.percentage AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_stake_pcent r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"5 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"r.num_winners AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_raf_eth_bidding r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"6 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"r.num_winners AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_raf_nft_bidding r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"7 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"r.num_winners AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_raf_nft_staking_cst r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"8 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"r.num_winners AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_raf_nft_staking_rwalk r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"9 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value, "+
						"0 AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_charity_addr r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.new_charity_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"10 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value, "+
						"0 AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_rwalk_addr r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.new_rwalk_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"11 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value, "+
						"0 AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_raffle_addr r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.new_raffle_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"12 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value, "+
						"0 AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_staking_cst_addr r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.new_staking_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"13 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value, "+
						"0 AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_staking_rwalk_addr r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.new_staking_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"14 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value, "+
						"0 AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_marketing_addr r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.new_marketing_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"15 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value, "+
						"0 AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_costok_addr r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.new_costok_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"16 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value, "+
						"0 AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_cossig_addr r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.new_cossig_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"17 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"a.addr AS addr_value, "+
						"0 AS int_value, " +
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_upgraded r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"LEFT JOIN address a ON a.address_id = r.implementation_aid "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"18 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value, "+
						"r.new_time_inc AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_time_inc r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"19 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value, "+
						"r.new_timeout AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_timeout_claimprize r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"20 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value, "+
						"r.new_price_increase AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_price_inc r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"21 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value, "+
						"r.new_nanoseconds AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_nanosec_extra r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"22 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value, "+
						"r.new_inisec AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_inisecprize r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"23 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value, "+
						"r.new_fraction AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_bidfraction r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"24 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value, "+
						"r.new_atime AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_acttime r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"25 AS record_type,"+
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value, "+
						"r.new_len AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_auclen r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"26 AS record_type,"+			//  Erc20RewardMultiplierChanged
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"0 AS int_value, "+
						"r.new_multiplier/1e18 AS float_value, "+
						"r.new_multiplier AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_erc_rwd_mul r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"27 AS record_type,"+			//  StartingBidPriceCSTMinLimitChanged
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"0 AS int_value, "+
						"r.new_price/1e18 AS float_value, "+
						"r.new_price AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_cst_min_lim r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"28 AS record_type,"+			// MarketingRewardChanged
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"0 AS int_value, "+
						"r.new_reward/1e18 AS float_value, "+
						"r.new_reward AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_mkt_reward r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"29 AS record_type,"+			// TokenRewardChanged
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"0 AS int_value, "+
						"r.new_reward/1e18 AS float_value, "+
						"r.new_reward AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_erc20_reward r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"30 AS record_type,"+			// MaxMessageLengthChanged
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"r.new_length AS int_value, "+
						"0 AS float_value, "+
						"'' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_msg_len r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"31 AS record_type,"+			// TokenGenerationScriptURLEvent
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"0 AS int_value, "+
						"0 AS float_value, "+
						"'new_uri' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_script_url r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				") UNION ALL ("+
					"SELECT "+
						"32 AS record_type,"+			// BaseURI
						"r.id record_id,"+
						"r.evtlog_id,"+
						"r.block_num,"+
						"t.id tx_id,"+
						"t.tx_hash,"+
						"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT ts,"+
						"r.time_stamp AS date_time, "+
						"'' AS addr_value," +
						"0 AS int_value, "+
						"0 AS float_value, "+
						"'new_uri' AS string_value "+
					"FROM "+sw.S.SchemaName()+".cg_adm_base_uri_cs r "+
					"LEFT JOIN transaction t ON t.id=r.tx_id "+
					"WHERE (r.evtlog_id>$1) AND (r.evtlog_id<$2) "+
				")" +
			") everything "+
			"ORDER BY evtlog_id "

	rows,err := sw.S.Db().Query(query,evtlog_start,evtlog_end)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGAdminEvent,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGAdminEvent
		err=rows.Scan(
			&rec.RecordType,
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.AddressValue,
			&rec.IntegerValue,
			&rec.FloatValue,
			&rec.StringValue,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
