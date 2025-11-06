package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_prize_claims(offset,limit int) []p.CGRoundRec {

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
				"p.timeout,"+
				"p.amount, "+
				"p.amount/1e18 amount_eth, " +
				"p.round_num,"+
				"p.token_id,"+
				"m.seed,"+
				"s.total_bids, "+
				"s.total_nft_donated, "+
				"s.num_erc20_donations, "+
				"s.total_raffle_eth_deposits,"+
				"s.total_raffle_eth_deposits/1e18 eth_deposits,"+
				"s.total_raffle_nfts, "+
				"d.donation_amount,"+
				"d.donation_amount/1e18 AS amount_eth, "+
				"d.charity_addr, "+
				"dp.deposit_amount,"+
				"dp.deposit_amount/1e18, "+
				"dp.amount_per_token,"+
				"dp.amount_per_token/1e18, "+
				"dp.deposit_id, "+
				"dp.num_staked_nfts "+
			"FROM "+sw.S.SchemaName()+".cg_prize_claim p "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
				"LEFT JOIN cg_mint_event m ON m.token_id=p.token_id "+
				"LEFT JOIN cg_round_stats s ON p.round_num=s.round_num "+
			"LEFT JOIN cg_staking_eth_deposit dp ON dp.round_num=p.round_num " +
			"LEFT JOIN ("+
				"SELECT round_num, SUM(amount) as donation_amount, STRING_AGG(DISTINCT cha.addr, ', ') as charity_addr "+
					"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
					"LEFT JOIN "+sw.S.SchemaName()+".address cha ON d.contract_aid=cha.address_id "+
					"WHERE round_num >= 0 "+
					"GROUP BY round_num "+
			") d ON p.round_num = d.round_num "+
			"ORDER BY p.id DESC "+
			"OFFSET $1 LIMIT $2"

			fmt.Printf("query - %v\n",query)
	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var null_seed sql.NullString
	records := make([]p.CGRoundRec,0, 256)
	defer rows.Close()
	var null_dep_amount,null_dep_amount_per_tok sql.NullString
	var null_dep_amount_eth,null_dep_amount_per_token_eth sql.NullFloat64
	var null_dep_deposit_num,null_num_staked_nfts sql.NullInt64
	var null_charity_amount,null_charity_addr sql.NullString
	var null_charity_amount_eth sql.NullFloat64
	for rows.Next() {
		var rec p.CGRoundRec
		err=rows.Scan(
			&rec.ClaimPrizeTx.Tx.EvtLogId,
			&rec.ClaimPrizeTx.Tx.BlockNum,
			&rec.ClaimPrizeTx.Tx.TxId,
			&rec.ClaimPrizeTx.Tx.TxHash,
			&rec.ClaimPrizeTx.Tx.TimeStamp,
			&rec.ClaimPrizeTx.Tx.DateTime,
			&rec.MainPrize.WinnerAid,
			&rec.MainPrize.WinnerAddr,
			&rec.MainPrize.TimeoutTs,
			&rec.MainPrize.EthAmount,
			&rec.MainPrize.EthAmountEth,
			&rec.RoundNum,
			&rec.MainPrize.NftTokenId,
			&null_seed,
			&rec.RoundStats.TotalBids,
			&rec.RoundStats.TotalDonatedNFTs,
			&rec.RoundStats.NumERC20Donations,
			&rec.RoundStats.TotalRaffleEthDeposits,
			&rec.RoundStats.TotalRaffleEthDepositsEth,
			&rec.RoundStats.TotalRaffleNFTs,
			&null_charity_amount,
			&null_charity_amount_eth,
			&null_charity_addr,
			&null_dep_amount,
			&null_dep_amount_eth,
			&null_dep_amount_per_tok,
			&null_dep_amount_per_token_eth,
			&null_dep_deposit_num,
			&null_num_staked_nfts,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_seed.Valid { rec.MainPrize.Seed = null_seed.String } else {rec.MainPrize.Seed = "???"}
		if null_charity_amount.Valid { rec.CharityDeposit.CharityAmount = null_charity_amount.String }
		if null_charity_amount_eth.Valid { rec.CharityDeposit.CharityAmountETH = null_charity_amount_eth.Float64 }
		if null_charity_addr.Valid { rec.CharityDeposit.CharityAddress = null_charity_addr.String }
		if null_dep_amount.Valid { rec.StakingDeposit.StakingDepositAmount = null_dep_amount.String }
		if null_dep_amount_eth.Valid { rec.StakingDeposit.StakingDepositAmountEth = null_dep_amount_eth.Float64 }
		if null_dep_amount_per_tok.Valid { rec.StakingDeposit.StakingPerToken = null_dep_amount_per_tok.String }
		if null_dep_amount_per_token_eth.Valid { rec.StakingDeposit.StakingPerTokenEth = null_dep_amount_per_token_eth.Float64 }
		if null_dep_deposit_num.Valid { rec.StakingDeposit.StakingDepositId = null_dep_deposit_num.Int64} else {rec.StakingDeposit.StakingDepositId = -1}
		if null_num_staked_nfts.Valid { rec.StakingDeposit.StakingNumStakedTokens = null_num_staked_nfts.Int64 }

		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_prize_info(round_num int64) (bool,p.CGRoundRec) {

	var rec p.CGRoundRec
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
			"p.timeout,"+
			"p.amount, "+
			"p.amount/1e18 amount_eth, " +
			"p.cst_amount, " +
			"p.cst_amount/1e18 cst_amount_eth, " +
			"p.round_num,"+
			"p.token_id,"+
			"m.seed, "+
			"s.total_bids,"+
			"s.total_nft_donated, "+
			"s.num_erc20_donations,"+
			"s.total_raffle_eth_deposits, "+
			"s.total_raffle_eth_deposits/1e18,"+
			"s.total_raffle_nfts, "+
			"d.donation_amount, "+
			"d.donation_amount/1e+18,"+
			"d.charity_addr, "+
			"dp.deposit_amount, "+
			"dp.deposit_amount/1e18,"+
			"dp.amount_per_token,"+
			"dp.amount_per_token/1e18, "+
			"dp.deposit_id, "+
			"dp.num_staked_nfts, "+
			"endu.erc721_token_id, "+
			"end_a.addr, "+
			"top.erc721_token_id,"+
			"top_a.addr, "+
			"w_a.addr,"+
			"endu.erc20_amount,"+
			"endu.erc20_amount/1e18, "+
		"top.erc20_amount,"+
		"top.erc20_amount/1e18, "+
		"w.eth_amount,"+
		"w.eth_amount/1e18, "+
		"w.cst_amount,"+
		"w.cst_amount/1e18, "+
		"w.nft_id,"+
		"s.donations_round_count,"+
		"s.donations_round_total,"+
		"s.donations_round_total/1e18, "+
		"s.param_window_start_time,"+
		"s.activation_time,"+
		"s.param_window_duration_seconds,"+
		"s.round_start_time,"+
		"s.round_end_time,"+
		"s.round_duration_seconds "+
		"FROM "+sw.S.SchemaName()+".cg_prize_claim p "+
			"LEFT JOIN transaction t ON t.id=tx_id "+
			"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"LEFT JOIN cg_mint_event m ON m.token_id=p.token_id "+
			"LEFT JOIN cg_staking_eth_deposit dp ON dp.round_num=p.round_num " +
			"LEFT JOIN cg_round_stats s ON s.round_num=p.round_num "+
			"LEFT JOIN cg_winner ws ON p.winner_aid=ws.winner_aid "+
			"LEFT JOIN cg_endurance_prize endu ON endu.round_num=p.round_num "+
			"LEFT JOIN cg_lastcst_prize top ON top.round_num=p.round_num "+
			"LEFT JOIN cg_chrono_warrior_prize w ON w.round_num = p.round_num "+
			"LEFT JOIN address end_a ON endu.winner_aid=end_a.address_id "+
			"LEFT JOIN address top_a ON top.winner_aid=top_a.address_id "+
		"LEFT JOIN address w_a ON w.winner_aid=w_a.address_id "+
		"LEFT JOIN ("+
			"SELECT round_num, SUM(amount) as donation_amount, STRING_AGG(DISTINCT cha.addr, ', ') as charity_addr "+
				"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
				"LEFT JOIN "+sw.S.SchemaName()+".address cha ON d.contract_aid=cha.address_id "+
				"WHERE round_num >= 0 "+
				"GROUP BY round_num "+
		") d ON p.round_num = d.round_num "+
	"WHERE p.round_num=$1"

	row := sw.S.Db().QueryRow(query,round_num)
	var null_seed sql.NullString
	var null_dep_amount,null_dep_amount_per_tok sql.NullString
	var null_dep_amount_eth,null_dep_amount_per_token_eth sql.NullFloat64
	var null_dep_deposit_num,null_num_staked_nfts sql.NullInt64
	var null_main_cst_amount sql.NullString
	var null_main_cst_amount_eth sql.NullFloat64
	var null_charity_amount,null_charity_addr sql.NullString
	var null_charity_amount_eth sql.NullFloat64
	var null_endurance_tid,null_lastcst_tid,null_warrior_nft_id sql.NullInt64
	var null_endurance_addr,null_lastcst_addr,null_warrior_addr sql.NullString
	var null_endurance_erc20_amount,null_lastcst_erc20_amount,null_warrior_eth_amount,null_warrior_cst_amount sql.NullString
	var null_endurance_erc20_amount_float,null_lastcst_erc20_amount_float,null_warrior_eth_amount_float,null_warrior_cst_amount_float sql.NullFloat64
	// Round timing fields (nullable)
	var null_param_window_start sql.NullString
	var null_activation_time sql.NullString
	var null_param_window_duration sql.NullInt64
	var null_round_start_time sql.NullString
	var null_round_end_time sql.NullString
	var null_round_duration sql.NullInt64
	err := row.Scan(
		&rec.ClaimPrizeTx.Tx.EvtLogId,
		&rec.ClaimPrizeTx.Tx.BlockNum,
		&rec.ClaimPrizeTx.Tx.TxId,
		&rec.ClaimPrizeTx.Tx.TxHash,
		&rec.ClaimPrizeTx.Tx.TimeStamp,
		&rec.ClaimPrizeTx.Tx.DateTime,
		&rec.MainPrize.WinnerAid,
		&rec.MainPrize.WinnerAddr,
		&rec.MainPrize.TimeoutTs,
		&rec.MainPrize.EthAmount,
		&rec.MainPrize.EthAmountEth,
		&null_main_cst_amount,
		&null_main_cst_amount_eth,
		&rec.RoundNum,
		&rec.MainPrize.NftTokenId,
		&null_seed,
		&rec.RoundStats.TotalBids,
		&rec.RoundStats.TotalDonatedNFTs,
		&rec.RoundStats.NumERC20Donations,
		&rec.RoundStats.TotalRaffleEthDeposits,
		&rec.RoundStats.TotalRaffleEthDepositsEth,
		&rec.RoundStats.TotalRaffleNFTs,
		&null_charity_amount,
		&null_charity_amount_eth,
		&null_charity_addr,
		&null_dep_amount,
		&null_dep_amount_eth,
		&null_dep_amount_per_tok,
		&null_dep_amount_per_token_eth,
		&null_dep_deposit_num,
		&null_num_staked_nfts,
		&null_endurance_tid,
		&null_endurance_addr,
		&null_lastcst_tid,
		&null_lastcst_addr,
		&null_warrior_addr,
		&null_endurance_erc20_amount,
		&null_endurance_erc20_amount_float,
		&null_lastcst_erc20_amount,
		&null_lastcst_erc20_amount_float,
		&null_warrior_eth_amount,
		&null_warrior_eth_amount_float,
		&null_warrior_cst_amount,
		&null_warrior_cst_amount_float,
		&null_warrior_nft_id,
		&rec.RoundStats.TotalDonatedCount,
		&rec.RoundStats.TotalDonatedAmount,
		&rec.RoundStats.TotalDonatedAmountEth,
		&null_param_window_start,
		&null_activation_time,
		&null_param_window_duration,
		&null_round_start_time,
		&null_round_end_time,
		&null_round_duration,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_seed.Valid { rec.MainPrize.Seed = null_seed.String } else {rec.MainPrize.Seed = "???"}
	if null_main_cst_amount.Valid { rec.MainPrize.CstAmount = null_main_cst_amount.String; rec.MainPrize.CstAmountEth = null_main_cst_amount_eth.Float64 }
	if null_charity_amount.Valid { rec.CharityDeposit.CharityAmount = null_charity_amount.String }
	if null_charity_amount_eth.Valid { rec.CharityDeposit.CharityAmountETH = null_charity_amount_eth.Float64 }
	if null_charity_addr.Valid { rec.CharityDeposit.CharityAddress = null_charity_addr.String }

	raffle_nft_winners := sw.Get_raffle_nft_winners_by_round(round_num,false)
	staking_nft_winners := sw.Get_raffle_nft_winners_by_round(round_num,true)
	raffle_eth_deposits := sw.Get_prize_deposits_by_round(round_num)
	all_prizes := sw.Get_all_prizes_for_round(round_num)

	rec.RaffleNFTWinners = raffle_nft_winners
	rec.StakingNFTWinners = staking_nft_winners
	rec.RaffleETHDeposits = raffle_eth_deposits
	rec.AllPrizes = all_prizes

	if null_dep_amount.Valid { rec.StakingDeposit.StakingDepositAmount = null_dep_amount.String }
	if null_dep_amount_eth.Valid { rec.StakingDeposit.StakingDepositAmountEth = null_dep_amount_eth.Float64 }
	if null_dep_amount_per_tok.Valid { rec.StakingDeposit.StakingPerToken = null_dep_amount_per_tok.String }
	if null_dep_amount_per_token_eth.Valid { rec.StakingDeposit.StakingPerTokenEth = null_dep_amount_per_token_eth.Float64 }
	if null_dep_deposit_num.Valid { rec.StakingDeposit.StakingDepositId = null_dep_deposit_num.Int64} else {rec.StakingDeposit.StakingDepositId = -1}
	if null_num_staked_nfts.Valid { rec.StakingDeposit.StakingNumStakedTokens = null_num_staked_nfts.Int64 }
	if null_endurance_tid.Valid { rec.EnduranceChampion.WinnerAddr = null_endurance_addr.String; rec.EnduranceChampion.NftTokenId=null_endurance_tid.Int64 }
	if null_lastcst_tid.Valid { rec.LastCstBidder.WinnerAddr = null_lastcst_addr.String; rec.LastCstBidder.NftTokenId=null_lastcst_tid.Int64 }
	if null_endurance_erc20_amount.Valid { rec.EnduranceChampion.CstAmount = null_endurance_erc20_amount.String; rec.EnduranceChampion.CstAmountEth = null_endurance_erc20_amount_float.Float64 }
	if null_lastcst_erc20_amount.Valid { rec.LastCstBidder.CstAmount = null_lastcst_erc20_amount.String; rec.LastCstBidder.CstAmountEth = null_lastcst_erc20_amount_float.Float64 }
	if null_warrior_eth_amount.Valid { rec.ChronoWarrior.EthAmount = null_warrior_eth_amount.String; rec.ChronoWarrior.EthAmountEth = null_warrior_eth_amount_float.Float64 }
	if null_warrior_cst_amount.Valid { rec.ChronoWarrior.CstAmount = null_warrior_cst_amount.String; rec.ChronoWarrior.CstAmountEth = null_warrior_cst_amount_float.Float64 }
	if null_warrior_nft_id.Valid { rec.ChronoWarrior.NftTokenId = null_warrior_nft_id.Int64 }
	if null_warrior_addr.Valid { rec.ChronoWarrior.WinnerAddr = null_warrior_addr.String }

	// Populate round timing fields
	if null_param_window_start.Valid { rec.RoundStats.ParamWindowStartTime = null_param_window_start.String }
	if null_activation_time.Valid { rec.RoundStats.ActivationTime = null_activation_time.String }
	if null_param_window_duration.Valid { rec.RoundStats.ParamWindowDurationSeconds = null_param_window_duration.Int64 }
	if null_round_start_time.Valid { rec.RoundStats.RoundStartTime = null_round_start_time.String }
	if null_round_end_time.Valid { rec.RoundStats.RoundEndTime = null_round_end_time.String }
	if null_round_duration.Valid { rec.RoundStats.RoundDurationSeconds = null_round_duration.Int64 }

	return true,rec
}
func (sw *SQLStorageWrapper) Get_all_prizes_for_round(round_num int64) []p.CGPrizeHistory {

	var query string
	query = "SELECT "+
				"p.ptype AS record_type,"+
			"COALESCE(pc.evtlog_id, rew.evtlog_id, rnw_bidder.evtlog_id, rnw_rwalk.evtlog_id, ew.evtlog_id, lw.evtlog_id, cw.evtlog_id, ed.evtlog_id) AS evtlog_id,"+
			"COALESCE(EXTRACT(EPOCH FROM pc.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rnw_bidder.time_stamp)::BIGINT, EXTRACT(EPOCH FROM rnw_rwalk.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ew.time_stamp)::BIGINT, EXTRACT(EPOCH FROM lw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM cw.time_stamp)::BIGINT, EXTRACT(EPOCH FROM ed.time_stamp)::BIGINT) AS tstmp,"+
			"COALESCE(pc.time_stamp, rew.time_stamp, rnw_bidder.time_stamp, rnw_rwalk.time_stamp, ew.time_stamp, lw.time_stamp, cw.time_stamp, ed.time_stamp) AS date_time,"+
			"COALESCE(pc.block_num, rew.block_num, rnw_bidder.block_num, rnw_rwalk.block_num, ew.block_num, lw.block_num, cw.block_num, ed.block_num) AS block_num,"+
			"COALESCE(tc.id, trew.id, trnw_bidder.id, trnw_rwalk.id, tew.id, tlw.id, tcw.id, ted.id) AS tx_id,"+
			"COALESCE(tc.tx_hash, trew.tx_hash, trnw_bidder.tx_hash, trnw_rwalk.tx_hash, tew.tx_hash, tlw.tx_hash, tcw.tx_hash, ted.tx_hash) AS tx_hash,"+
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
		"'T' AS claimed,"+
	"CASE WHEN p.ptype = 15 THEN '(All CS NFT Stakers)' ELSE COALESCE(wa_pc.addr, wa_rew.addr, wa_rnw_bidder.addr, wa_rnw_rwalk.addr, wa_ew.addr, wa_lw.addr, wa_cw.addr, '') END AS winner_addr,"+
		"COALESCE(pc.winner_aid, rew.winner_aid, rnw_bidder.winner_aid, rnw_rwalk.winner_aid, ew.winner_aid, lw.winner_aid, cw.winner_aid, 0) AS winner_aid "+
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
		"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_nft_prize rnw_bidder ON (p.round_num = rnw_bidder.round_num AND p.winner_index = rnw_bidder.winner_idx AND p.ptype IN (11,12) AND rnw_bidder.is_rwalk = false) "+
		"LEFT JOIN "+sw.S.SchemaName()+".transaction trnw_bidder ON trnw_bidder.id = rnw_bidder.tx_id "+
		"LEFT JOIN "+sw.S.SchemaName()+".address wa_rnw_bidder ON rnw_bidder.winner_aid = wa_rnw_bidder.address_id "+
		"LEFT JOIN "+sw.S.SchemaName()+".cg_raffle_nft_prize rnw_rwalk ON (p.round_num = rnw_rwalk.round_num AND p.winner_index = rnw_rwalk.winner_idx AND p.ptype IN (13,14) AND rnw_rwalk.is_rwalk = true) "+
		"LEFT JOIN "+sw.S.SchemaName()+".transaction trnw_rwalk ON trnw_rwalk.id = rnw_rwalk.tx_id "+
		"LEFT JOIN "+sw.S.SchemaName()+".address wa_rnw_rwalk ON rnw_rwalk.winner_aid = wa_rnw_rwalk.address_id "+
		"LEFT JOIN "+sw.S.SchemaName()+".cg_staking_eth_deposit ed ON (p.round_num = ed.round_num AND p.ptype = 15) "+
		"LEFT JOIN "+sw.S.SchemaName()+".transaction ted ON ted.id = ed.tx_id "+
		"WHERE p.round_num = $1 "+
			"ORDER BY p.ptype, p.winner_index"

	rows,err := sw.S.Db().Query(query,round_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGPrizeHistory,0, 64)
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
