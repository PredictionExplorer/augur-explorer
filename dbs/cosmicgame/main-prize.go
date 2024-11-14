package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_prize_claims(offset,limit int) []p.CGPrizeRec {

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
				"p.amount, "+
				"p.amount/1e18 amount_eth, " +
				"p.prize_num,"+
				"p.token_id,"+
				"m.seed,"+
				"s.total_bids, "+
				"s.total_nft_donated, "+
				"s.total_raffle_eth_deposits,"+
				"s.total_raffle_eth_deposits/1e18 eth_deposits,"+
				"s.total_raffle_nfts, "+
				"d.donation_amount,"+
				"d.donation_amount/1e18 AS amount_eth, "+
				"d.charity_addr, "+
				"dp.amount,"+
				"dp.amount/1e18, "+
				"dp.amount_per_staker,"+
				"dp.amount_per_staker/1e18i, "+
				"dp.deposit_num, "+
				"dp.num_staked_nfts "+
			"FROM "+sw.S.SchemaName()+".cg_prize_claim p "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
				"LEFT JOIN cg_mint_event m ON m.token_id=p.prize_num "+
				"LEFT JOIN cg_round_stats s ON p.prize_num=s.round_num "+
				"LEFT JOIN cg_eth_deposit dp ON dp.round_num=p.prize_num " +
				"LEFT JOIN LATERAL (" +
					"SELECT d.evtlog_id,d.amount donation_amount,cha.addr charity_addr "+
						"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
						"JOIN "+sw.S.SchemaName()+".address cha ON d.contract_aid=cha.address_id "+
				") d ON p.donation_evt_id=d.evtlog_id "+
			"ORDER BY p.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var null_seed sql.NullString
	records := make([]p.CGPrizeRec,0, 256)
	defer rows.Close()
	var null_dep_amount,null_dep_amount_per_tok sql.NullString
	var null_dep_amount_eth,null_dep_amount_per_token_eth sql.NullFloat64
	var null_dep_deposit_num,null_num_staked_nfts sql.NullInt64
	for rows.Next() {
		var rec p.CGPrizeRec
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.Amount,
			&rec.AmountEth,
			&rec.PrizeNum,
			&rec.TokenId,
			&null_seed,
			&rec.RoundStats.TotalBids,
			&rec.RoundStats.TotalDonatedNFTs,
			&rec.RoundStats.TotalRaffleEthDeposits,
			&rec.RoundStats.TotalRaffleEthDepositsEth,
			&rec.RoundStats.TotalRaffleNFTs,
			&rec.CharityAmount,
			&rec.CharityAmountETH,
			&rec.CharityAddress,
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
		if null_seed.Valid { rec.Seed = null_seed.String } else {rec.Seed = "???"}
		if null_dep_amount.Valid { rec.StakingDepositAmount = null_dep_amount.String }
		if null_dep_amount_eth.Valid { rec.StakingDepositAmountEth = null_dep_amount_eth.Float64 }
		if null_dep_amount_per_tok.Valid { rec.StakingPerToken = null_dep_amount_per_tok.String }
		if null_dep_amount_per_token_eth.Valid { rec.StakingPerTokenEth = null_dep_amount_per_token_eth.Float64 }
		if null_dep_deposit_num.Valid { rec.StakingDepositNum = null_dep_deposit_num.Int64} else {rec.StakingDepositNum = -1}
		if null_num_staked_nfts.Valid { rec.StakingNumStakedTokens = null_num_staked_nfts.Int64 }

		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_prize_info(prize_num int64) (bool,p.CGPrizeRec) {

	var rec p.CGPrizeRec
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
				"p.amount, "+
				"p.amount/1e18 amount_eth, " +
				"p.prize_num,"+
				"p.token_id,"+
				"m.seed, "+
				"s.total_bids,"+
				"s.total_nft_donated, "+
				"s.total_raffle_eth_deposits, "+
				"s.total_raffle_eth_deposits/1e18,"+
				"s.total_raffle_nfts, "+
				"d.donation_amount, "+
				"d.donation_amount/1e+18,"+
				"d.charity_addr, "+
				"dp.amount, "+
				"dp.amount/1e18,"+
				"dp.amount_per_staker,"+
				"dp.amount_per_staker/1e18, "+
				"dp.deposit_num, "+
				"dp.num_staked_nfts, "+
				"endu.erc721_token_id, "+
				"end_a.addr, "+
				"top.erc721_token_id,"+
				"top_a.addr, "+
				"endu.erc20_amount,"+
				"endu.erc20_amount/1e18, "+
				"top.erc20_amount,"+
				"top.erc20_amount/1e18, "+
				"s.donations_round_count,"+
				"s.donations_round_total,"+
				"s.donations_round_total/1e18 "+
			"FROM "+sw.S.SchemaName()+".cg_prize_claim p "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
				"LEFT JOIN cg_mint_event m ON m.token_id=p.token_id "+
				"LEFT JOIN cg_eth_deposit dp ON dp.round_num=p.prize_num " +
				"LEFT JOIN cg_round_stats s ON s.round_num=p.prize_num "+
				"LEFT JOIN cg_winner ws ON p.winner_aid=ws.winner_aid "+
				"LEFT JOIN cg_endurance_winner endu ON endu.round_num=p.prize_num "+
				"LEFT JOIN cg_stellar_winner top ON top.round_num=p.prize_num "+
				"LEFT JOIN address end_a ON endu.winner_aid=end_a.address_id "+
				"LEFT JOIN address top_a ON top.winner_aid=top_a.address_id "+
				"LEFT JOIN LATERAL (" +
					"SELECT d.evtlog_id,d.amount donation_amount,cha.addr charity_addr "+
						"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
						"JOIN "+sw.S.SchemaName()+".address cha ON d.contract_aid=cha.address_id "+
				") d ON p.donation_evt_id=d.evtlog_id "+
			"WHERE p.prize_num=$1"

	row := sw.S.Db().QueryRow(query,prize_num)
	var null_seed sql.NullString
	var null_dep_amount,null_dep_amount_per_tok sql.NullString
	var null_dep_amount_eth,null_dep_amount_per_token_eth sql.NullFloat64
	var null_dep_deposit_num,null_num_staked_nfts sql.NullInt64
	var null_endurance_tid,null_topbidder_tid sql.NullInt64
	var null_endurance_addr,null_topbidder_addr,null_endurance_erc20_amount,null_topbidder_erc20_amount sql.NullString
	var null_endurance_erc20_amount_float,null_topbidder_erc20_amount_float sql.NullFloat64
	err := row.Scan(
		&rec.EvtLogId,
		&rec.BlockNum,
		&rec.TxId,
		&rec.TxHash,
		&rec.TimeStamp,
		&rec.DateTime,
		&rec.WinnerAid,
		&rec.WinnerAddr,
		&rec.Amount,
		&rec.AmountEth,
		&rec.PrizeNum,
		&rec.TokenId,
		&null_seed,
		&rec.RoundStats.TotalBids,
		&rec.RoundStats.TotalDonatedNFTs,
		&rec.RoundStats.TotalRaffleEthDeposits,
		&rec.RoundStats.TotalRaffleEthDepositsEth,
		&rec.RoundStats.TotalRaffleNFTs,
		&rec.CharityAmount,
		&rec.CharityAmountETH,
		&rec.CharityAddress,
		&null_dep_amount,
		&null_dep_amount_eth,
		&null_dep_amount_per_tok,
		&null_dep_amount_per_token_eth,
		&null_dep_deposit_num,
		&null_num_staked_nfts,
		&null_endurance_tid,
		&null_endurance_addr,
		&null_topbidder_tid,
		&null_topbidder_addr,
		&null_endurance_erc20_amount,
		&null_endurance_erc20_amount_float,
		&null_topbidder_erc20_amount,
		&null_topbidder_erc20_amount_float,
		&rec.RoundStats.TotalDonatedCount,
		&rec.RoundStats.TotalDonatedAmount,
		&rec.RoundStats.TotalDonatedAmountEth,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_seed.Valid { rec.Seed = null_seed.String } else {rec.Seed = "???"}

	raffle_nft_winners := sw.Get_raffle_nft_winners_by_round(prize_num)
	raffle_eth_deposits := sw.Get_prize_deposits_by_round(prize_num)

	rec.RaffleNFTWinners = raffle_nft_winners
	rec.RaffleETHDeposits = raffle_eth_deposits

	if null_dep_amount.Valid { rec.StakingDepositAmount = null_dep_amount.String }
	if null_dep_amount_eth.Valid { rec.StakingDepositAmountEth = null_dep_amount_eth.Float64 }
	if null_dep_amount_per_tok.Valid { rec.StakingPerToken = null_dep_amount_per_tok.String }
	if null_dep_amount_per_token_eth.Valid { rec.StakingPerTokenEth = null_dep_amount_per_token_eth.Float64 }
	if null_dep_deposit_num.Valid { rec.StakingDepositNum = null_dep_deposit_num.Int64} else {rec.StakingDepositNum = -1}
	if null_num_staked_nfts.Valid { rec.StakingNumStakedTokens = null_num_staked_nfts.Int64 }
	if null_endurance_tid.Valid { rec.EnduranceWinnerAddr = null_endurance_addr.String; rec.EnduranceERC721TokenId=null_endurance_tid.Int64 }
	if null_topbidder_tid.Valid { rec.StellarWinnerAddr = null_topbidder_addr.String; rec.StellarERC721TokenId=null_topbidder_tid.Int64 }
	if null_endurance_erc20_amount.Valid { rec.EnduranceERC20Amount = null_endurance_erc20_amount.String; rec.EnduranceERC20AmountEth = null_endurance_erc20_amount_float.Float64 }
	if null_topbidder_erc20_amount.Valid { rec.StellarERC20Amount = null_topbidder_erc20_amount.String; rec.StellarERC20AmountEth = null_topbidder_erc20_amount_float.Float64 }

	return true,rec
}
