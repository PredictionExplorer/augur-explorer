package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_stake_action_cst_info(action_id int64) (bool,p.CGStakeUnstakeCombined) {

	var rec p.CGStakeUnstakeCombined
	var query string
	query = "SELECT " +
				// stake action fields
				"st.id,"+
				"st.evtlog_id,"+
				"st.block_num,"+
				"ts.id,"+
				"ts.tx_hash,"+
				"EXTRACT(EPOCH FROM st.time_stamp)::BIGINT,"+
				"st.time_stamp,"+
				"st.action_id,"+
				"st.token_id,"+
				"st.num_staked_nfts,"+
				"st.staker_aid,"+
				"sa.addr,"+

				//unstake action fields
				"u.id,"+
				"u.evtlog_id,"+
				"u.block_num,"+
				"tu.id,"+
				"tu.tx_hash,"+
				"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT,"+
				"u.time_stamp,"+
				"u.action_id,"+
				"u.token_id, "+
				"u.num_staked_nfts, "+
				"u.reward,"+
				"u.reward/1e18,"+
				"u.staker_aid, "+
				"ua.addr "+
			"FROM "+sw.S.SchemaName()+".cg_nft_staked_cst st "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction ts ON ts.id=st.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address sa ON st.staker_aid=sa.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_nft_unstaked_cst u ON st.action_id=u.action_id " +
				"LEFT JOIN "+sw.S.SchemaName()+".transaction tu ON tu.id=u.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address ua ON u.staker_aid=ua.address_id "+
			"WHERE st.action_id=$1"

	row := sw.S.Db().QueryRow(query,action_id)
	var err error
	var null_record_id,null_evtlog_id,null_tx_id,null_unstake_ts,null_action_id sql.NullInt64
	var null_block_num,null_token_id,null_num_staked_nfts,null_staker_aid sql.NullInt64
	var null_unstake_date,null_tx_hash,null_staker_addr,null_reward sql.NullString
	var null_reward_eth sql.NullFloat64
	err=row.Scan(
		// stake action fields
		&rec.Stake.RecordId,
		&rec.Stake.EvtLogId,
		&rec.Stake.BlockNum,
		&rec.Stake.TxId,
		&rec.Stake.TxHash,
		&rec.Stake.TimeStamp,
		&rec.Stake.DateTime,
		&rec.Stake.ActionId,
		&rec.Stake.TokenId,
		&rec.Stake.NumStakedNFTs,
		&rec.Stake.StakerAid,
		&rec.Stake.StakerAddr,
		// unstake action fields
		&null_record_id,
		&null_evtlog_id,
		&null_block_num,
		&null_tx_id,
		&null_tx_hash,
		&null_unstake_ts,
		&null_unstake_date,
		&null_action_id,
		&null_token_id,
		&null_num_staked_nfts,
		&null_reward,
		&null_reward_eth,
		&null_staker_aid,
		&null_staker_addr,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_record_id.Valid { rec.Unstake.RecordId = null_record_id.Int64 }
	if null_evtlog_id.Valid { rec.Unstake.EvtLogId = null_evtlog_id.Int64 }
	if null_block_num.Valid { rec.Unstake.BlockNum = null_block_num.Int64 }
	if null_tx_id.Valid { rec.Unstake.TxId = null_tx_id.Int64 }
	if null_tx_hash.Valid { rec.Unstake.TxHash = null_tx_hash.String }
	if null_unstake_ts.Valid { rec.Unstake.TimeStamp = null_unstake_ts.Int64 }
	if null_unstake_date.Valid { rec.Unstake.DateTime = null_unstake_date.String }
	if null_action_id.Valid { rec.Unstake.ActionId = null_action_id.Int64 }
	if null_token_id.Valid { rec.Unstake.TokenId = null_token_id.Int64 }
	if null_num_staked_nfts.Valid { rec.Unstake.NumStakedNFTs = null_num_staked_nfts.Int64 }
	if null_reward.Valid { rec.Unstake.RewardAmount = null_reward.String }
	if null_reward_eth.Valid { rec.Unstake.RewardAmountEth = null_reward_eth.Float64 }
	if null_staker_aid.Valid { rec.Unstake.StakerAid = null_staker_aid.Int64 }
	if null_staker_addr.Valid { rec.Unstake.StakerAddr = null_staker_addr.String }
	return true,rec
}
func (sw *SQLStorageWrapper) Get_stake_action_rwalk_info(action_id int64) (bool,p.CGStakeUnstakeCombined) {

	var rec p.CGStakeUnstakeCombined
	var query string
	query = "SELECT " +
				// stake action fields
				"st.id,"+
				"st.evtlog_id,"+
				"st.block_num,"+
				"ts.id,"+
				"ts.tx_hash,"+
				"EXTRACT(EPOCH FROM st.time_stamp)::BIGINT,"+
				"st.time_stamp,"+
				"st.action_id,"+
				"st.token_id,"+
				"st.num_staked_nfts,"+
				"st.staker_aid,"+
				"sa.addr,"+

				//unstake action fields
				"u.id,"+
				"u.evtlog_id,"+
				"u.block_num,"+
				"tu.id,"+
				"tu.tx_hash,"+
				"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT,"+
				"u.time_stamp,"+
				"u.action_id,"+
				"u.token_id, "+
				"u.num_staked_nfts, "+
				"u.staker_aid, "+
				"ua.addr "+
			"FROM "+sw.S.SchemaName()+".cg_nft_staked_rwalk st "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction ts ON ts.id=st.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address sa ON st.staker_aid=sa.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_nft_unstaked_rwalk u ON st.action_id=u.action_id " +
				"LEFT JOIN "+sw.S.SchemaName()+".transaction tu ON tu.id=u.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address ua ON u.staker_aid=ua.address_id "+
			"WHERE st.action_id=$1"

	row := sw.S.Db().QueryRow(query,action_id)
	var err error
	var null_record_id,null_evtlog_id,null_tx_id,null_unstake_ts,null_action_id sql.NullInt64
	var null_block_num,null_token_id,null_num_staked_nfts,null_staker_aid sql.NullInt64
	var null_unstake_date,null_tx_hash,null_staker_addr sql.NullString
	err=row.Scan(
		// stake action fields
		&rec.Stake.RecordId,
		&rec.Stake.EvtLogId,
		&rec.Stake.BlockNum,
		&rec.Stake.TxId,
		&rec.Stake.TxHash,
		&rec.Stake.TimeStamp,
		&rec.Stake.DateTime,
		&rec.Stake.ActionId,
		&rec.Stake.TokenId,
		&rec.Stake.NumStakedNFTs,
		&rec.Stake.StakerAid,
		&rec.Stake.StakerAddr,
		// unstake action fields
		&null_record_id,
		&null_evtlog_id,
		&null_block_num,
		&null_tx_id,
		&null_tx_hash,
		&null_unstake_ts,
		&null_unstake_date,
		&null_action_id,
		&null_token_id,
		&null_num_staked_nfts,
		&null_staker_aid,
		&null_staker_addr,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_record_id.Valid { rec.Unstake.RecordId = null_record_id.Int64 }
	if null_evtlog_id.Valid { rec.Unstake.EvtLogId = null_evtlog_id.Int64 }
	if null_block_num.Valid { rec.Unstake.BlockNum = null_block_num.Int64 }
	if null_tx_id.Valid { rec.Unstake.TxId = null_tx_id.Int64 }
	if null_tx_hash.Valid { rec.Unstake.TxHash = null_tx_hash.String }
	if null_unstake_ts.Valid { rec.Unstake.TimeStamp = null_unstake_ts.Int64 }
	if null_unstake_date.Valid { rec.Unstake.DateTime = null_unstake_date.String }
	if null_action_id.Valid { rec.Unstake.ActionId = null_action_id.Int64 }
	if null_token_id.Valid { rec.Unstake.TokenId = null_token_id.Int64 }
	if null_num_staked_nfts.Valid { rec.Unstake.NumStakedNFTs = null_num_staked_nfts.Int64 }
	if null_staker_aid.Valid { rec.Unstake.StakerAid = null_staker_aid.Int64 }
	if null_staker_addr.Valid { rec.Unstake.StakerAddr = null_staker_addr.String }
	return true,rec
}
func (sw *SQLStorageWrapper) Get_staking_rewards_to_be_claimed(user_aid int64) []p.CGRewardToClaim {

	var query string
	query = "WITH rwd AS ("+
				"SELECT "+
					"COUNT(id) AS num_toks_collected,"+
					"SUM(reward) AS collected_reward," +
					"SUM(reward)/1e18 AS collected_reward_eth,"+
					"deposit_id "+
				"FROM cg_claim_reward "+
				"WHERE staker_aid=$1 "+
				"GROUP BY deposit_id "+
			") "+
			"SELECT "+
				"d.id,"+
				"d.evtlog_id,"+
				"d.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"EXTRACT(EPOCH FROM d.deposit_time)::BIGINT,"+
				"d.deposit_time,"+
				"d.deposit_id,"+
				"d.num_staked_nfts,"+
				"d.amount,"+
				"d.amount/1e18,"+
				"sd.tokens_staked,"+
				"sd.amount_to_claim,"+
				"sd.amount_to_claim/1e18,"+
				"sd.amount_to_claim - COALESCE(rwd.collected_reward,0),"+
				"(sd.amount_to_claim - COALESCE(rwd.collected_reward,0))/1e18, "+
				"d.amount/d.num_staked_nfts,"+
				"(d.amount/d.num_staked_nfts)/1e18 "+
			"FROM "+sw.S.SchemaName()+".cg_staker_deposit sd "+
				"INNER JOIN cg_eth_deposit d ON sd.deposit_id=d.deposit_id "+
				"INNER JOIN transaction tx ON tx.id=d.tx_id " +
				"LEFT JOIN rwd ON rwd.deposit_id=sd.deposit_id "+
			"WHERE (sd.staker_aid = $1) AND (sd.amount_to_claim != COALESCE(rwd.collected_reward,0)) " +
			"ORDER BY d.id DESC "
	rows,err := sw.S.Db().Query(query,user_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRewardToClaim,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRewardToClaim
		var null_collected sql.NullString
		var null_collected_eth sql.NullFloat64
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DepositTimeStamp,
			&rec.DepositDate,
			&rec.DepositId,
			&rec.NumStakedNFTs,
			&rec.DepositAmount,
			&rec.DepositAmountEth,
			&rec.YourTokensStaked,
			&rec.YourClaimableAmount,
			&rec.YourClaimableAmountEth,
			&null_collected,
			&null_collected_eth,
			&rec.AmountPerToken,
			&rec.AmountPerTokenEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_collected.Valid {rec.YourClaimableAmount = null_collected.String }
		if null_collected_eth.Valid { rec.YourClaimableAmountEth = null_collected_eth.Float64 }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_staking_rewards_collected(user_aid int64,offset,limit int) []p.CGCollectedReward {

	var query string
	query = "WITH rwd AS ("+
				"SELECT "+
					"COUNT(token_id) AS num_toks_collected,"+
					"SUM(reward) AS collected_reward," +
					"SUM(reward)/1e18 AS collected_reward_eth,"+
					"deposit_id "+
				"FROM cg_st_reward "+
				"WHERE staker_aid=$1 AND collected='T' "+
				"GROUP BY deposit_id "+
			") "+
			"SELECT "+
				"d.id,"+
				"d.evtlog_id,"+
				"d.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"EXTRACT(EPOCH FROM d.deposit_time)::BIGINT,"+
				"d.deposit_time,"+
				"d.deposit_id,"+
				"d.num_staked_nfts,"+
				"d.amount,"+
				"d.amount/1e18,"+
				"sd.tokens_staked,"+
				"sd.amount_to_claim,"+
				"sd.amount_to_claim/1e18,"+
				"rwd.num_toks_collected,"+
				"d.amount/num_staked_nfts,"+
				"(d.amount/num_staked_nfts)/1e18, "+
				"modulo,"+
				"modulo/1e+18, "+
				"d.round_num, "+
				"rwd.collected_reward,"+
				"rwd.collected_reward_eth "+
			"FROM "+sw.S.SchemaName()+".cg_staker_deposit sd "+
				"INNER JOIN cg_eth_deposit d ON sd.deposit_id=d.deposit_id "+
				"INNER JOIN transaction tx ON tx.id=d.tx_id " +
				"INNER JOIN rwd ON rwd.deposit_id=sd.deposit_id "+
			"WHERE sd.staker_aid=$1 "+
			"ORDER BY d.id DESC " +
			"OFFSET $2 LIMIT $3"

	rows,err := sw.S.Db().Query(query,user_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCollectedReward,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCollectedReward	
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DepositTimeStamp,
			&rec.DepositDate,
			&rec.DepositId,
			&rec.NumStakedNFTs,
			&rec.TotalDepositAmount,
			&rec.TotalDepositAmountEth,
			&rec.YourTokensStaked,
			&rec.YourAmountToClaim,
			&rec.YourAmountToClaimEth,
			&rec.NumTokensCollected,
			&rec.DepositAmountPerToken,
			&rec.DepositAmountPerTokenEth,
			&rec.Modulo,
			&rec.ModuloF64,
			&rec.RoundNum,
			&rec.YourCollectedAmount,
			&rec.YourCollectedAmountEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if rec.YourAmountToClaimEth == rec.YourCollectedAmountEth { rec.FullyClaimed = true }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_staked_tokens_cst_global() []p.CGStakedTokenCSTRec {

	var query string
	query = "SELECT "+
				"m.id,"+
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
				"a.evtlog_id,"+
				"a.block_num,"+
				"a.action_id,"+
				"EXTRACT(EPOCH FROM a.time_stamp)::BIGINT,"+
				"a.time_stamp,"+
				"sa.addr,"+
				"sa.address_id "+
			"FROM "+sw.S.SchemaName()+".cg_staked_token_cst st "+
				"LEFT JOIN cg_mint_event m ON st.token_id=m.token_id "+
				"LEFT JOIN transaction t ON t.id=m.tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
				"LEFT JOIN cg_nft_staked_cst a ON a.action_id=st.stake_action_id "+
				"LEFT JOIN address sa ON a.staker_aid = sa.address_id "+
			"ORDER BY m.token_id"

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGStakedTokenCSTRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGStakedTokenCSTRec 
		var null_prize_num sql.NullInt64
		err=rows.Scan(
			&rec.TokenInfo.RecordId,
			&rec.TokenInfo.EvtLogId,
			&rec.TokenInfo.BlockNum,
			&rec.TokenInfo.TxId,
			&rec.TokenInfo.TxHash,
			&rec.TokenInfo.TimeStamp,
			&rec.TokenInfo.DateTime,
			&rec.TokenInfo.WinnerAid,
			&rec.TokenInfo.WinnerAddr,
			&rec.TokenInfo.CurOwnerAid,
			&rec.TokenInfo.CurOwnerAddr,
			&rec.TokenInfo.Seed,
			&rec.TokenInfo.TokenId,
			&rec.TokenInfo.RoundNum,
			&null_prize_num,
			&rec.TokenInfo.TokenName,
			&rec.StakeEvtLogId,
			&rec.StakeBlockNum,
			&rec.StakeActionId,
			&rec.StakeTimeStamp,
			&rec.StakeDateTime,
			&rec.UserAddr,
			&rec.UserAid,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_prize_num.Valid { rec.TokenInfo.RecordType = 3 } else {rec.TokenInfo.RecordType = 1 }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_staked_tokens_rwalk_global() []p.CGStakedTokenRWalkRec {

	var query string
	query = "SELECT "+
				"a.evtlog_id,"+
				"a.block_num,"+
				"a.action_id,"+
				"EXTRACT(EPOCH FROM a.time_stamp)::BIGINT,"+
				"a.time_Stamp,"+
				"sa.addr,"+
				"sa.address_id, "+
				"st.token_id "+
			"FROM "+sw.S.SchemaName()+".cg_staked_token_rwalk st "+
				"LEFT JOIN cg_mint_event m ON st.token_id=m.token_id "+
				"LEFT JOIN transaction t ON t.id=m.tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
				"LEFT JOIN cg_nft_staked_rwalk a ON a.action_id=st.stake_action_id "+
				"LEFT JOIN address sa ON a.staker_aid = sa.address_id "+
			"ORDER BY m.token_id"

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGStakedTokenRWalkRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGStakedTokenRWalkRec 
		err=rows.Scan(
			&rec.StakeEvtLogId,
			&rec.StakeBlockNum,
			&rec.StakeActionId,
			&rec.StakeTimeStamp,
			&rec.StakeDateTime,
			&rec.UserAddr,
			&rec.UserAid,
			&rec.StakedTokenId,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_action_ids_for_deposit(deposit_id int64,user_aid int64) []p.CGActionIdsForDeposit {

	records := make([]p.CGActionIdsForDeposit,0, 16)
	cur_ts := sw.S.Get_last_block_timestamp()
	var query string
	query = "SELECT EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS ts FROM cg_eth_deposit d WHERE deposit_id=$1"
	row := sw.S.Db().QueryRow(query,deposit_id)
	var null_ts sql.NullInt64
	err:=row.Scan(&null_ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_action_ids_for_deposit(): %v, q=%v",err,query))
		os.Exit(1)
	}
	query = "SELECT "+
				"a.id,"+
				"a.action_id, "+
				"a.token_id, "+
				"EXTRACT(epoch FROM a.time_stamp)::BIGINT action_ts,"+
//				"r.deposit_id, "+
				"d.amount_per_staker, "+
				"d.amount_per_staker/1e18 amount_eth "+
			"FROM "+sw.S.SchemaName()+".cg_nft_staked_cst a "+
				"JOIN cg_eth_deposit d ON d.deposit_id=$3 "+
				"LEFT JOIN cg_nft_staked_cst u ON a.action_id=u.action_id "+
//PENDING				"LEFT JOIN cg_claim_reward r ON (a.action_id=r.action_id) AND (r.deposit_id=$3) AND (r.staker_aid=a.staker_aid)" +
			"WHERE "+
				"(a.staker_aid = $1) AND ("+
					"("+
						"(a.action_id < $2) AND (u.id IS NULL)"+
					") OR "+
						"(" + 
							"(a.action_id<$2 AND "+
							"(u.id IS NOT NULL) AND "+
							"($2<=u.action_id) "+
						")"+
					")"+
				") " +
			"ORDER BY a.action_id "
	rows,err := sw.S.Db().Query(query,user_aid,null_ts.Int64,deposit_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.CGActionIdsForDeposit
		var null_deposit_id sql.NullInt64
		err=rows.Scan(
			&rec.RecordId,
			&rec.StakeActionId,
			&rec.TokenId,
			&rec.StakeActionTimeStamp,
//			&null_deposit_id,			// PENDING for refactoring
			&rec.Amount,
			&rec.AmountEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.DepositId = deposit_id
		rec.UserAid = user_aid
		if null_deposit_id.Valid {rec.Claimed = true }
		rec.CurChainTimeStamp = cur_ts
		rec.TimeStampDiff = -1
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_action_ids_for_deposit_with_claim_info(deposit_id int64,user_aid int64) []p.CGActionIdsForDepositWithClaimInfo {

	records := make([]p.CGActionIdsForDepositWithClaimInfo,0, 16)
	var query string
	query = "SELECT EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS ts FROM cg_eth_deposit d WHERE deposit_id=$1"
	row := sw.S.Db().QueryRow(query,deposit_id)
	var null_ts sql.NullInt64
	err:=row.Scan(&null_ts)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return records
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_action_ids_for_deposit(): %v, q=%v",err,query))
		os.Exit(1)
	}

	query = "SELECT "+
				"r.id,"+
				"a.action_id, "+
				"a.token_id, "+
				"r.deposit_id, "+
				"r.block_num, "+
				"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT,"+
				"r.time_stamp,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"r.reward,"+
				"r.reward/1e18 "+
			"FROM "+sw.S.SchemaName()+".cg_nft_staked_cst a "+
				"JOIN cg_claim_reward r ON (a.action_id=r.action_id) AND (r.deposit_id=$3) AND (r.staker_aid=a.staker_aid)" +
				"LEFT JOIN cg_nft_unstaked_cst u ON a.action_id=u.action_id "+
				"LEFT JOIN transaction tx ON tx.id=r.tx_id " +
			"WHERE "+
				"(a.staker_aid = $1) AND ("+
					"("+
						"(a.time_stamp < TO_TIMESTAMP($2)) AND (u.id IS NULL)"+
					") OR "+
						"(" + 
							"(a.time_stamp<TO_TIMESTAMP($2) AND "+
							"(u.id IS NOT NULL) AND "+
							"(TO_TIMESTAMP($2)<=u.time_stamp) "+
						")"+
					")"+
				") " +
			"ORDER BY r.evtlog_id DESC "

	rows,err := sw.S.Db().Query(query,user_aid,null_ts.Int64,deposit_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.CGActionIdsForDepositWithClaimInfo
		var null_deposit_id sql.NullInt64
		var null_rwd_block_num,null_rwd_timestamp,null_rwd_tx_id sql.NullInt64
		var null_rwd_datetime,null_rwd_tx_hash,null_reward sql.NullString
		var null_reward_eth sql.NullFloat64
		err=rows.Scan(
			&rec.RecordId,
			&rec.StakeActionId,
			&rec.TokenId,
			&null_deposit_id,
			&null_rwd_block_num,
			&null_rwd_timestamp,
			&null_rwd_datetime,
			&null_rwd_tx_id,
			&null_rwd_tx_hash,
			&null_reward,
			&null_reward_eth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.DepositId = deposit_id
		rec.UserAid = user_aid
		if null_deposit_id.Valid {rec.Claimed = true }
		if null_rwd_block_num.Valid {rec.ClaimBlockNum = null_rwd_block_num.Int64}
		if null_rwd_timestamp.Valid {rec.ClaimTimeStamp = null_rwd_timestamp.Int64}
		if null_rwd_datetime.Valid {rec.ClaimDateTime = null_rwd_datetime.String}
		if null_rwd_tx_id.Valid {rec.ClaimTxId = null_rwd_tx_id.Int64}
		if null_rwd_tx_hash.Valid {rec.ClaimTxHash = null_rwd_tx_hash.String}
		if null_reward.Valid {rec.ClaimRewardAmount = null_reward.String}
		if null_reward_eth.Valid {rec.ClaimRewardAmountEth = null_reward_eth.Float64}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_global_staking_rewards(offset,limit int) []p.CGCollectedReward {

	records := make([]p.CGCollectedReward,0, 32)
	var query string
	query = "WITH rwd AS ("+
				"SELECT "+
					"COUNT(id) AS num_toks_collected,"+
					"SUM(reward) AS collected_reward," +
					"SUM(reward)/1e18 AS collected_reward_eth,"+
					"deposit_id, "+
					"staker_aid "+
				"FROM cg_claim_reward "+
				"GROUP BY staker_aid,deposit_id "+
			") "+
			"SELECT "+
				"d.id,"+
				"d.evtlog_id,"+
				"d.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"EXTRACT(EPOCH FROM d.deposit_time)::BIGINT,"+
				"d.deposit_time,"+
				"d.deposit_id,"+
				"d.num_staked_nfts,"+
				"d.amount,"+
				"d.amount/1e18,"+
				"sd.tokens_staked,"+
				"sd.amount_to_claim,"+
				"sd.amount_to_claim/1e18,"+
				"COALESCE(rwd.num_toks_collected,0),"+
				"d.round_num, "+
				"COALESCE(rwd.collected_reward,0),"+
				"COALESCE(rwd.collected_reward_eth,0), "+
				"sd.staker_aid,"+
				"sa.addr "+
			"FROM "+sw.S.SchemaName()+".cg_staker_deposit sd "+
				"INNER JOIN cg_eth_deposit d ON sd.deposit_id=d.deposit_id "+
				"INNER JOIN transaction tx ON tx.id=d.tx_id " +
				"LEFT JOIN rwd ON (rwd.deposit_id=sd.deposit_id) AND (rwd.staker_aid=sd.staker_aid) "+
				"INNER JOIN address sa ON sd.staker_aid = sa.address_id "+
			"ORDER BY d.id DESC,sd.staker_aid " +
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCollectedReward	
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DepositTimeStamp,
			&rec.DepositDate,
			&rec.DepositId,
			&rec.NumStakedNFTs,
			&rec.TotalDepositAmount,
			&rec.TotalDepositAmountEth,
			&rec.YourTokensStaked,
			&rec.YourAmountToClaim,
			&rec.YourAmountToClaimEth,
			&rec.NumTokensCollected,
			&rec.RoundNum,
			&rec.YourCollectedAmount,
			&rec.YourCollectedAmountEth,
			&rec.StakerAid,
			&rec.StakerAddr,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if rec.YourAmountToClaimEth == rec.YourCollectedAmountEth { rec.FullyClaimed = true }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_staking_cst_rewards_by_round(round_num int64) []p.CGEthDepositAsReward {
	
	records := make([]p.CGEthDepositAsReward,0, 32)
	var query string
	query = "SELECT "+
				"d.id,"+
				"d.evtlog_id,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS tstmp, "+
				"d.time_stamp AS date_time, "+
				"d.block_num,"+
				"d.tx_id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.deposit_time)::BIGINT, "+
				"d.deposit_time, "+
				"d.num_staked_nfts,"+
				"d.amount,"+
				"d.amount/1e18 AS amount_eth,"+
				"d.amount_per_staker,"+
				"d.amount_per_staker/1e18 AS amount_eth, "+
				"sd.staker_aid, "+
				"sa.addr,"+
				"sd.tokens_staked,"+
				"sd.amount_to_claim, "+
				"sd.amount_to_claim/1e18  "+
			"FROM cg_staker_deposit sd "+
				"LEFT JOIN cg_eth_deposit d ON sd.deposit_id=d.deposit_id "+
				"LEFT JOIN address sa ON sd.staker_aid = sa.address_id "+
				"LEFT JOIN transaction t ON t.id=d.tx_id "+
			"WHERE d.round_num=$1 "

	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.CGEthDepositAsReward
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.DepositTimeStamp,
			&rec.DepositDate,
			&rec.NumStakedNFTsTotal,
			&rec.Amount,
			&rec.AmountEth,
			&rec.AmountPerToken,
			&rec.AmountPerTokenEth,
			&rec.StakerAid,
			&rec.StakerAddr,
			&rec.StakerNumStakedNFTs,
			&rec.StakerAmount,
			&rec.StakerAmountEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_global_staking_cst_history(offset,limit int) []p.CGStakingCSTHistoryRec {

	var query string
	query = "("+
				"SELECT "+
					"0 AS action_type,"+
					"s.id,"+
					"s.evtlog_id,"+
					"s.block_num,"+
					"tx.id,"+
					"tx.tx_hash,"+
					"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT,"+
					"s.time_stamp,"+
					"-1 AS usts,"+
					"TO_TIMESTAMP(0) AS unstake_time,"+
					"s.action_id,"+
					"s.token_id,"+
					"s.num_staked_nfts, "+
					"s.staker_aid, "+
					"sa.addr staker_addr "+
				"FROM "+sw.S.SchemaName()+".cg_nft_staked_cst s "+
					"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
					"LEFT JOIN address sa ON s.staker_aid=sa.address_id "+
			") UNION ALL ("+
				"SELECT "+
					"1 AS action_type,"+
					"u.id,"+
					"u.evtlog_id,"+
					"u.block_num,"+
					"tx.id,"+
					"tx.tx_hash,"+
					"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT AS usts,"+
					"u.time_stamp,"+
					"0 AS usts,"+
					"TO_TIMESTAMP(0) AS unstake_time,"+
					"u.action_id,"+
					"u.token_id,"+
					"u.num_staked_nfts, "+
					"u.staker_aid," +
					"ua.addr staker_addr "+
				"FROM "+sw.S.SchemaName()+".cg_nft_unstaked_cst u "+
					"LEFT JOIN transaction tx ON tx.id=u.tx_id " +
					"LEFT JOIN address ua ON u.staker_aid=ua.address_id "+
			") ORDER BY evtlog_id DESC " +
			"OFFSET $1 LIMIT $2 "

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGStakingCSTHistoryRec,0, 16)
	accum_staked_nfts := int64(0)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGStakingCSTHistoryRec
		err=rows.Scan(
			&rec.ActionType,
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.UnstakeTimeStamp,
			&rec.UnstakeDate,
			&rec.ActionId,
			&rec.TokenId,
			&rec.NumStakedNFTs,
			&rec.StakerAid,
			&rec.StakerAddr,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		accum_staked_nfts = accum_staked_nfts + rec.NumStakedNFTs
		rec.AccumNumStakedNFTs = accum_staked_nfts
//		rec.LastBlockTS = last_ts		DISCONTINUED, removal pending
//		rec.UnstakeExpirationDiff = -1	DISCONTINUED, removal pending
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_global_staking_rwalk_history(offset,limit int) []p.CGStakingRWalkHistoryRec {

	last_ts := sw.S.Get_last_block_timestamp()
	var query string
	query = "("+
				"SELECT "+
					"0 AS action_type,"+
					"s.id,"+
					"s.evtlog_id,"+
					"s.block_num,"+
					"tx.id,"+
					"tx.tx_hash,"+
					"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT,"+
					"s.time_stamp,"+
					"-1 AS usts,"+
					"TO_TIMESTAMP(0) AS unstake_time,"+
					"s.action_id,"+
					"s.token_id,"+
					"s.num_staked_nfts, "+
					"s.staker_aid, "+
					"sa.addr staker_addr "+
				"FROM "+sw.S.SchemaName()+".cg_nft_staked_rwalk s "+
					"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
					"LEFT JOIN address sa ON s.staker_aid=sa.address_id "+
				"ORDER BY s.id DESC " +
				"OFFSET $1 LIMIT $2 "+
			") UNION ALL ("+
				"SELECT "+
					"1 AS action_type,"+
					"u.id,"+
					"u.evtlog_id,"+
					"u.block_num,"+
					"tx.id,"+
					"tx.tx_hash,"+
					"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT AS usts,"+
					"u.time_stamp,"+
					"0 AS usts,"+
					"TO_TIMESTAMP(0) AS unnstake_time,"+
					"u.action_id,"+
					"u.token_id,"+
					"u.num_staked_nfts, "+
					"u.staker_aid," +
					"ua.addr staker_addr "+
				"FROM "+sw.S.SchemaName()+".cg_nft_unstaked_rwalk u "+
					"LEFT JOIN transaction tx ON tx.id=u.tx_id " +
					"LEFT JOIN address ua ON u.staker_aid=ua.address_id "+
					"LEFT JOIN cg_nft_staked_rwalk s ON u.action_id=s.action_id "+
				"ORDER BY u.id DESC " +
				"OFFSET $1 LIMIT $2 "+
			") order by evtlog_id DESC"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGStakingRWalkHistoryRec,0, 16)
	accum_staked_nfts := int64(0)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGStakingRWalkHistoryRec
		err=rows.Scan(
			&rec.ActionType,
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.UnstakeTimeStamp,
			&rec.UnstakeDate,
			&rec.ActionId,
			&rec.TokenId,
			&rec.NumStakedNFTs,
			&rec.StakerAid,
			&rec.StakerAddr,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		accum_staked_nfts = accum_staked_nfts + rec.NumStakedNFTs
		rec.AccumNumStakedNFTs = accum_staked_nfts
		rec.LastBlockTS = last_ts
		rec.UnstakeExpirationDiff = -1
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_staking_rwalk_mints_global(offset,limit int) []p.CGRaffleNFTWinnerRec {

	var query string
	query = "SELECT "+
				"w.id,"+
				"w.evtlog_id,"+
				"w.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM w.time_stamp)::BIGINT,"+
				"w.time_stamp,"+
				"w.token_id,"+
				"w.winner_idx,"+
				"w.round_num,"+
				"w.winner_aid,"+
				"wa.addr "+
			"FROM cg_raffle_nft_winner w "+
				"LEFT JOIN transaction t ON t.id=w.tx_id "+
				"LEFT JOIN address wa ON w.winner_aid=wa.address_id "+
			"WHERE (is_rwalk=TRUE) AND (is_staker=TRUE) "+
			"ORDER BY w.evtlog_id DESC " +
			"OFFSET $1 LIMIT $2"
	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleNFTWinnerRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleNFTWinnerRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TokenId,
			&rec.WinnerIndex,
			&rec.RoundNum,
			&rec.WinnerAid,
			&rec.WinnerAddr,
		)
		rec.IsRWalk = true
		rec.IsStaker = true
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_staking_cst_mints_global(offset,limit int) []p.CGRaffleNFTWinnerRec {

	var query string
	query = "SELECT "+
				"w.id,"+
				"w.evtlog_id,"+
				"w.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM w.time_stamp)::BIGINT,"+
				"w.time_stamp,"+
				"w.token_id,"+
				"w.winner_idx,"+
				"w.round_num,"+
				"w.winner_aid,"+
				"wa.addr "+
			"FROM cg_raffle_nft_winner w "+
				"LEFT JOIN transaction t ON t.id=w.tx_id "+
				"LEFT JOIN address wa ON w.winner_aid=wa.address_id "+
			"WHERE (is_rwalk=FALSE) AND (is_staker=TRUE) "+
			"ORDER BY w.evtlog_id DESC " +
			"OFFSET $1 LIMIT $2"
	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleNFTWinnerRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleNFTWinnerRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TokenId,
			&rec.WinnerIndex,
			&rec.RoundNum,
			&rec.WinnerAid,
			&rec.WinnerAddr,
		)
		rec.IsRWalk = true
		rec.IsStaker = true
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_staking_cst_history_by_user(user_aid int64) []p.CGCombinedDepositRewardRec {

	var query string
		query = 
			"SELECT "+
				"sa_id,sa_evtlog_id,sa_block_num,sa_tx_id,sa_tx_hash,sa_time_stamp,sa_date_time, "+
				"sa_action_id,sa_token_id,sa_num_staked_nfts,"+
				"ua_id,ua_evtlog_id,ua_block_num,ua_tx_id,ua_tx_hash,ua_time_stamp,ua_date_time, "+
				"ua_action_id,ua_token_id,ua_num_staked_nfts,ua_reward,ua_reward_eth,ua_unpaid_deposit,"+
				"d.id,d.evtlog_id,d.block_num,tx.id,tx.tx_hash,EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,d.time_stamp,"+
				"d.deposit_id,d.round_num,d.num_staked_nfts,d.amount,d.amount/1e18,amount_per_staker/1e18,"+
				"str.reward,"+
				"str.reward/1e18,"+
				"str.collected "+
			"FROM "+sw.S.SchemaName()+".cg_st_reward str "+
				"INNER JOIN cg_eth_deposit d ON str.deposit_id=d.deposit_id "+
				"INNER JOIN transaction tx ON tx.id=d.tx_id " +
				"INNER JOIN LATERAL ("+
					"SELECT "+
						"sa.id sa_id,sa.evtlog_id sa_evtlog_id,sa.block_num sa_block_num,satx.id sa_tx_id,satx.tx_hash sa_tx_hash,EXTRACT(EPOCH FROM sa.time_stamp)::BIGINT sa_time_stamp,sa.time_stamp sa_date_time, "+
						"sa.action_id sa_action_id,sa.token_id sa_token_id,sa.num_staked_nfts sa_num_staked_nfts,"+
						"ua.id ua_id,ua.evtlog_id ua_evtlog_id,ua.block_num ua_block_num,uatx.id ua_tx_id,uatx.tx_hash ua_tx_hash,EXTRACT(EPOCH FROM ua.time_stamp)::BIGINT ua_time_stamp,ua.time_stamp ua_date_time, "+
						"ua.action_id ua_action_id,ua.token_id ua_token_id,ua.num_staked_nfts ua_num_staked_nfts,ua.reward ua_reward,ua.reward/1e18 ua_reward_eth,ua.unpaid_deposit ua_unpaid_deposit "+
					"FROM cg_nft_staked_cst sa " +
						"LEFT JOIN cg_nft_unstaked_cst ua ON ua.action_id=sa.action_id "+
						"LEFT JOIN transaction satx ON satx.id=sa.tx_id "+
						"LEFT JOIN transaction uatx ON uatx.id=ua.tx_id "+
				") a ON a.sa_action_id=str.action_id "+
			"WHERE str.staker_aid=$1 "+
			"ORDER BY d.id ASC,sa_action_id DESC "		// Note: sort order is ASC because we are accumularting the sum in golang (between rows)

	fmt.Printf("user_id = %v\n",user_aid)
	fmt.Printf("query = %v\n",query)
	rows,err := sw.S.Db().Query(query,user_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var rec p.CGCombinedDepositRewardRec
	var cur_deposit_id int64 = -1
	var your_tokens_staked int64 = 0
	var num_tokens_collected int64 = 0
	var your_claimable_amount float64 = 0
	var your_claimed_amount float64 = 0
	var fully_claimed bool = true
	records := make([]p.CGCombinedDepositRewardRec,0, 16)
	actions := make([]p.CGNftStakeUnstakeCombined,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec_row p.CGNftStakeUnstakeCombined
		var null_record_id,null_action_id,null_evtlog_id,null_block_num,null_tx_id,null_token_id,null_num_staked_nfts,null_unpaid_dep_idx,null_time_stamp sql.NullInt64
		var null_tx_hash,null_reward,null_date_time sql.NullString
		var null_reward_eth sql.NullFloat64
		var record_id,evtlog_id,block_num,tx_id,deposit_id,time_stamp,dep_round,num_staked_nfts int64
		var tx_hash,date_time,deposit_amount string
		var dep_amount_eth,amount_per_token_eth float64
		err=rows.Scan(
			&rec_row.Stake.RecordId,&rec_row.Stake.EvtLogId,&rec_row.Stake.BlockNum,&rec_row.Stake.TxId,&rec_row.Stake.TxHash,&rec_row.Stake.TimeStamp,&rec_row.Stake.DateTime,
			&rec_row.Stake.ActionId,&rec_row.Stake.TokenId,&rec_row.Stake.NumStakedNFTs,
			&null_record_id,&null_evtlog_id,&null_block_num,&null_tx_id,&null_tx_hash,&null_time_stamp,&null_date_time,
			&null_action_id,&null_token_id,&null_num_staked_nfts,&null_reward,&null_reward_eth,&null_unpaid_dep_idx,
			&record_id,&evtlog_id,&block_num,&tx_id,&tx_hash,&time_stamp,&date_time,
			&deposit_id,&dep_round,&num_staked_nfts,&deposit_amount,&dep_amount_eth,&amount_per_token_eth,
			&rec_row.Reward,&rec_row.RewardEth,
			&rec_row.Claimed,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_record_id.Valid { rec_row.Unstake.RecordId = null_record_id.Int64; rec_row.Unstake.StakerAid = user_aid; }
		if null_evtlog_id.Valid { rec_row.Unstake.EvtLogId = null_evtlog_id.Int64 }
		if null_block_num.Valid { rec_row.Unstake.BlockNum = null_block_num.Int64 }
		if null_tx_id.Valid { rec_row.Unstake.TxId = null_tx_id.Int64 }
		if null_tx_hash.Valid { rec_row.Unstake.TxHash = null_tx_hash.String }
		if null_time_stamp.Valid { rec_row.Unstake.TimeStamp = null_time_stamp.Int64 }
		if null_date_time.Valid { rec_row.Unstake.DateTime = null_date_time.String }
		if null_action_id.Valid { rec_row.Unstake.ActionId = null_action_id.Int64 }
		if null_token_id.Valid { rec_row.Unstake.TokenId = null_token_id.Int64 }
		if null_num_staked_nfts.Valid { rec_row.Unstake.NumStakedNFTs = null_num_staked_nfts.Int64 }
		if null_unpaid_dep_idx.Valid { rec_row.Unstake.MaxUnpaidDepositIndex = null_unpaid_dep_idx.Int64 }
		if null_reward.Valid { rec_row.Unstake.RewardAmount = null_reward.String }
		if null_reward_eth.Valid { rec_row.Unstake.RewardAmountEth = null_reward_eth.Float64 }
		if deposit_id != cur_deposit_id {
			if cur_deposit_id != -1 {
				rec.Actions = actions;
				rec.YourTokensStaked = your_tokens_staked
				rec.YourClaimableAmountEth = your_claimable_amount
				rec.FullyClaimed = fully_claimed
				rec.NumTokensCollected = num_tokens_collected
				rec.ClaimedAmountEth = your_claimed_amount
				records = append(records,rec)

				fully_claimed = true
				your_tokens_staked = 0
				your_claimable_amount = 0
				actions = make([]p.CGNftStakeUnstakeCombined,0, 16)
			}
			rec.RecordId = record_id
			rec.EvtLogId = evtlog_id
			rec.BlockNum = block_num
			rec.TxId = tx_id
			rec.TxHash = tx_hash
			rec.TimeStamp = time_stamp
			rec.DateTime  = date_time
			rec.DepositId = deposit_id
			rec.DepositRoundNum = dep_round
			rec.NumStakedNFTs = num_staked_nfts
			rec.DepositId = deposit_id 
			rec.DepositAmount = deposit_amount
			rec.DepositAmountEth = dep_amount_eth
			rec.AmountPerTokenEth = amount_per_token_eth
			cur_deposit_id = deposit_id
		}
		your_tokens_staked += 1
		if !rec_row.Claimed {
			your_claimable_amount += rec_row.RewardEth
			fully_claimed = false
		} else {
			num_tokens_collected += 1
			your_claimed_amount += rec_row.RewardEth
		}
		actions = append(actions,rec_row)
	}
	if your_tokens_staked > 0 {
		rec.Actions = actions;
		rec.YourTokensStaked = your_tokens_staked
		rec.YourClaimableAmountEth = your_claimable_amount
		rec.FullyClaimed = fully_claimed
		rec.NumTokensCollected = num_tokens_collected
		rec.ClaimedAmountEth = your_claimed_amount
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_staking_reward_paid_records(user_aid int64) []p.CGRewardPaidRec {

	records := make([]p.CGRewardPaidRec,0, 16)
	var query string

	query = "SELECT "+
				"r.id,"+
				"r.action_id, "+
				"r.token_id, "+
				"r.block_num, "+
				"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT,"+
				"r.time_stamp,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"r.reward,"+
				"r.reward/1e18, "+
				"r.unpaid_dep_id "+
			"FROM "+sw.S.SchemaName()+".cg_reward_paid r "+
				"LEFT JOIN address sa ON r.staker_aid=sa.address_id "+
				"LEFT JOIN transaction tx ON tx.id=r.tx_id " +
			"WHERE "+
				"r.staker_aid = $1 "+
			"ORDER BY r.evtlog_id DESC "

	rows,err := sw.S.Db().Query(query,user_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRewardPaidRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.ActionId,
			&rec.TokenId,
			&rec.BlockNum,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TxId,
			&rec.TxHash,
			&rec.RewardAmount,
			&rec.RewardAmountEth,
			&rec.UnpaidDepositIndex,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.StakerAid = user_aid
		records = append(records,rec)
	}
	return records
}
