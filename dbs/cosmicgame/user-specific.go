package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_user_info(user_aid int64) (bool,p.CGUserInfo) {

	var query string
	query = "SELECT "+
				"a.address_id,"+
				"a.addr, "+
				"b.num_bids, "+
				"b.max_bid/1e18 AS max_bid,"+
				"p.prizes_count,"+
				"p.max_win_amount/1e18 max_win, "+
				"rw.amount_sum/1e18 raffle_win_sum, "+
				"rw.withdrawal_sum/1e18 withdrawal_sum, "+
				"rw.raffles_count, "+
				"rn.num_won raffle_nft_won, "+
				"p.unclaimed_nfts, "+
				"p.tokens_count, "+
				"trs.erc20_num_transfers, "+
				"trs.erc721_num_transfers "+
			"FROM address a "+
				"LEFT JOIN cg_bidder b ON b.bidder_aid=a.address_id "+
				"LEFT JOIN cg_winner p ON p.winner_aid=a.address_id "+
				"LEFT JOIN cg_raffle_winner_stats rw ON rw.winner_aid=a.address_id "+
				"LEFT JOIN cg_raffle_nft_winner_stats rn ON rn.winner_aid=a.address_id "+
				"LEFT JOIN cg_transfer_stats trs ON trs.user_aid=a.address_id "+
				"LEFT JOIN cg_staker_cst st ON st.staker_aid=a.address_id "+
			"WHERE a.address_id=$1"

	var rec p.CGUserInfo
	var null_num_bids,null_prizes_count sql.NullInt64
	var null_max_bid,null_max_win sql.NullFloat64
	var null_raffle_sum_winnings,null_raffle_sum_withdrawal sql.NullFloat64
	var null_raffles_count,null_raffle_nft_won sql.NullInt64
	var null_unclaimed_nfts,null_total_tokens sql.NullInt64
	var null_erc20_transfs,null_erc721_transfs sql.NullInt64


	row := sw.S.Db().QueryRow(query,user_aid)
	var err error
	err=row.Scan(
		&rec.AddressId,
		&rec.Address,
		&null_num_bids,
		&null_max_bid,
		&null_prizes_count,
		&null_max_win,
		&null_raffle_sum_winnings,
		&null_raffle_sum_withdrawal,
		&null_raffles_count,
		&null_raffle_nft_won,
		&null_unclaimed_nfts,
		&null_total_tokens,
		&null_erc20_transfs,
		&null_erc721_transfs,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("Error in main query of Get_user_info(): %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_num_bids.Valid { rec.NumBids = null_num_bids.Int64 }
	if null_prizes_count.Valid { rec.NumPrizes = null_prizes_count.Int64 }
	if null_max_bid.Valid { rec.MaxBidAmount = null_max_bid.Float64 }
	if null_max_win.Valid { rec.MaxWinAmount = null_max_win.Float64 }
	if null_raffle_sum_winnings.Valid { rec.SumRaffleEthWinnings = null_raffle_sum_winnings.Float64 }
	if null_raffle_sum_withdrawal.Valid { rec.SumRaffleEthWithdrawal = null_raffle_sum_withdrawal.Float64 }
	if null_raffles_count.Valid { rec.NumRaffleEthWinnings = null_raffles_count.Int64 }
	if null_raffle_nft_won.Valid { rec.RaffleNFTWon = null_raffle_nft_won.Int64 }
	if null_unclaimed_nfts.Valid { rec.UnclaimedNFTs = null_unclaimed_nfts.Int64 }
	if null_total_tokens.Valid { rec.TotalCSTokensWon= null_total_tokens.Int64 }
	if null_erc20_transfs.Valid { rec.CosmicTokenNumTransfers = null_erc20_transfs.Int64 }
	if null_erc721_transfs.Valid { rec.CosmicSignatureNumTransfers = null_erc721_transfs.Int64 }

	query = "SELECT "+
				"s.total_tokens_staked,"+
				"s.num_stake_actions,"+
				"s.num_unstake_actions,"+
				"s.total_reward,"+
				"total_reward/1e18,"+
				"unclaimed_reward,"+
				"unclaimed_reward/1e18, "+
				"num_tokens_minted "+
			"FROM cg_staker_cst s "+
			"WHERE staker_aid=$1"
	row = sw.S.Db().QueryRow(query,user_aid)
	{
		// we use a code block because null_*** variables have same names in both code blocks, to ensure they are empty
		var err error
		var null_total_tokens_staked,null_num_stake_actions,null_num_unstake_actions,null_num_tokens_minted sql.NullInt64
		var null_total_reward,null_unclaimed_reward sql.NullString
		var null_total_reward_eth,null_unclaimed_reward_eth sql.NullFloat64
		err=row.Scan(
			&null_total_tokens_staked,
			&null_num_stake_actions,
			&null_num_unstake_actions,
			&null_total_reward,
			&null_total_reward_eth,
			&null_unclaimed_reward,
			&null_unclaimed_reward_eth,
			&null_num_tokens_minted,
		)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				return false,rec
			}
			sw.S.Log_msg(fmt.Sprintf("Error in staker_cst query in Get_user_info(): %v, q=%v",err,query))
			os.Exit(1)
		}
		if null_total_tokens_staked.Valid { rec.StakingStatistics.CSTStakingInfo.TotalTokensStaked = null_total_tokens_staked.Int64 }
		if null_num_stake_actions.Valid { rec.StakingStatistics.CSTStakingInfo.TotalNumStakeActions = null_num_stake_actions.Int64 }
		if null_num_unstake_actions.Valid { rec.StakingStatistics.CSTStakingInfo.TotalNumUnstakeActions = null_num_unstake_actions.Int64 }
		if null_total_reward.Valid { rec.StakingStatistics.CSTStakingInfo.TotalReward = null_total_reward.String }
		if null_total_reward_eth.Valid { rec.StakingStatistics.CSTStakingInfo.TotalRewardEth = null_total_reward_eth.Float64 }
		if null_unclaimed_reward.Valid { rec.StakingStatistics.CSTStakingInfo.UnclaimedReward = null_unclaimed_reward.String }
		if null_unclaimed_reward_eth.Valid { rec.StakingStatistics.CSTStakingInfo.UnclaimedRewardEth = null_unclaimed_reward_eth.Float64 }
		if null_num_tokens_minted.Valid { rec.StakingStatistics.CSTStakingInfo.TotalTokensMinted = null_num_tokens_minted.Int64 }
	}
	query = "SELECT "+
				"s.total_tokens_staked,"+
				"s.num_stake_actions,"+
				"s.num_unstake_actions,"+
				"s.num_tokens_minted "+
			"FROM cg_staker_rwalk s "+
			"WHERE staker_aid=$1"
	{
		// we use a code block because null_*** variables have same names in both code blocks, to ensure they are empty
		row := sw.S.Db().QueryRow(query,user_aid)
		var err error
		var null_total_tokens_staked,null_num_stake_actions,null_num_unstake_actions,null_num_tokens_minted sql.NullInt64
		err=row.Scan(
			&null_total_tokens_staked,
			&null_num_stake_actions,
			&null_num_unstake_actions,
			&null_num_tokens_minted,
		)
		if (err!=nil) {
			if err == sql.ErrNoRows {
				return false,rec
			}
			sw.S.Log_msg(fmt.Sprintf("Error in staker_rwalk query in Get_user_info(): %v, q=%v",err,query))
			os.Exit(1)
		}
		if null_total_tokens_staked.Valid { rec.StakingStatistics.RWalkStakingInfo.TotalTokensStaked = null_total_tokens_staked.Int64 }
		if null_num_stake_actions.Valid { rec.StakingStatistics.RWalkStakingInfo.TotalNumStakeActions = null_num_stake_actions.Int64 }
		if null_num_unstake_actions.Valid { rec.StakingStatistics.RWalkStakingInfo.TotalNumUnstakeActions = null_num_unstake_actions.Int64 }
		if null_num_tokens_minted.Valid { rec.StakingStatistics.RWalkStakingInfo.TotalTokensMinted = null_num_tokens_minted.Int64 }
	}
	return true,rec
}
func (sw *SQLStorageWrapper) Get_prize_claims_by_user(winner_aid int64) []p.CGPrizeRec {

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
				"s.total_bids,"+
				"s.total_nft_donated, "+
				"s.total_raffle_eth_deposits,"+
				"s.total_raffle_eth_deposits/1e18 eth_deposits,"+
				"s.total_raffle_nfts, "+
				"d.donation_amount,"+
				"d.donation_amount/1e+18, "+
				"d.charity_addr "+
			"FROM "+sw.S.SchemaName()+".cg_prize_claim p "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
				"LEFT JOIN cg_mint_event m ON m.token_id=p.token_id "+
				"LEFT JOIN cg_round_stats s ON p.prize_num=s.round_num "+
				"LEFT JOIN LATERAL (" +
					"SELECT d.evtlog_id,d.amount donation_amount,cha.addr charity_addr "+
						"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
						"JOIN "+sw.S.SchemaName()+".address cha ON d.contract_aid=cha.address_id "+
				") d ON p.donation_evt_id=d.evtlog_id "+
			"WHERE winner_aid=$1 "+
			"ORDER BY p.id DESC"

	rows,err := sw.S.Db().Query(query,winner_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var null_seed sql.NullString
	records := make([]p.CGPrizeRec,0, 32)
	defer rows.Close()
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
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_seed.Valid { rec.Seed = null_seed.String } else {rec.Seed = "???"}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_bids_by_user(bidder_aid int64) []p.CGBidRec {

	var query string
	query =  "SELECT " +
				"b.evtlog_id,"+
				"b.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT,"+
				"b.time_stamp,"+
				"b.bidder_aid,"+
				"ba.addr,"+
				"b.bid_price,"+
				"b.bid_price/1e18 bid_price_eth, " +
				"b.rwalk_nft_id,"+
				"b.erc20_amount,"+
				"b.erc20_amount/1e18 erc20_amount_eth, "+
				"d.token_id,"+
				"d.tok_addr, "+
				"d.token_uri, "+
				"b.msg, "+
				"b.round_num, "+
				"b.num_cst_tokens, "+
				"b.num_cst_tokens/1e18, "+
				"b.bid_type "+
			"FROM "+sw.S.SchemaName()+".cg_bid b "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address ba ON b.bidder_aid=ba.address_id "+
				"LEFT JOIN LATERAL (" +
					"SELECT d.bid_id,token_id,token_aid,ta.addr tok_addr,d.token_uri "+
						"FROM "+sw.S.SchemaName()+".cg_nft_donation d "+
						"JOIN "+sw.S.SchemaName()+".address ta ON d.token_aid=ta.address_id "+
				") d ON b.id=d.bid_id "+
			"WHERE b.bidder_aid=$1 "+
			"ORDER BY b.id DESC"

	rows,err := sw.S.Db().Query(query,bidder_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGBidRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGBidRec
		var null_token_id sql.NullInt64
		var null_tok_addr,null_token_uri sql.NullString
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.BidderAid,
			&rec.BidderAddr,
			&rec.BidPrice,
			&rec.BidPriceEth,
			&rec.RWalkNFTId,
			&rec.ERC20_Amount,
			&rec.ERC20_AmountEth,
			&null_token_id,
			&null_tok_addr,
			&null_token_uri,
			&rec.Message,
			&rec.RoundNum,
			&rec.NumCSTTokens,
			&rec.NumCSTTokensEth,
			&rec.BidType,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.NFTDonationTokenId = -1
		if null_token_id.Valid { rec.NFTDonationTokenId=null_token_id.Int64 }
		if null_tok_addr.Valid { rec.NFTDonationTokenAddr = null_tok_addr.String }
		if null_token_uri.Valid { rec.NFTTokenURI = null_token_uri.String }
		records = append(records,rec)
	}
	return records

}
func (sw *SQLStorageWrapper) Get_unclaimed_donated_nft_by_user(winner_aid int64) []p.CGNFTDonation {

	var query string
	query = "SELECT "+
				"d.id,"+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.round_num,"+
				"d.donor_aid,"+
				"da.addr, "+
				"d.token_id, "+
				"d.idx,"+
				"nft.address_id,"+
				"nft.addr, "+
				"d.token_uri "+
			"FROM "+sw.S.SchemaName()+".cg_nft_donation d "+
				"JOIN "+sw.S.SchemaName()+".cg_prize_claim p ON p.prize_num=d.round_num "+
				"LEFT JOIN cg_donated_nft_claimed c ON c.idx=d.idx "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=d.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address nft ON d.token_aid=nft.address_id "+
			"WHERE p.winner_aid=$1 AND p.prize_num IS NOT NULL  AND c.idx IS NULL " +
			"ORDER BY d.evtlog_id DESC "

	rows,err := sw.S.Db().Query(query,winner_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGNFTDonation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGNFTDonation
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
			&rec.NFTTokenId,
			&rec.Index,
			&rec.TokenAddressId,
			&rec.TokenAddr,
			&rec.NFTTokenURI,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_unclaimed_token_ids(winner_aid int64) []int64 {

	var query string
	query = "SELECT "+
				"p.token_id, "+
			"FROM "+sw.S.SchemaName()+".cg_raffle_nft_winner w "+
				"LEFT JOIN cg_raffle_nft_claimed c ON c.nft_winner_evtlog_id=w.evtlog_id "+
			"WHERE w.winner_aid=$1  AND c.nft_winner_vetlog_id IS NULL "+
			"ORDER BY w.id"

	rows,err := sw.S.Db().Query(query,winner_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]int64,0, 256)
	defer rows.Close()
	for rows.Next() {
		var token_id int64
		err=rows.Scan(&token_id)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,token_id)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_raffle_nft_winnings_by_user(winner_aid int64) []p.CGRaffleNFTWinnerRec {

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
				"p.token_id,"+
				"p.winner_idx "+
			"FROM "+sw.S.SchemaName()+".cg_raffle_nft_winner p "+
				"LEFT JOIN transaction t ON t.id=p.tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"WHERE p.winner_aid=$1 "+
			"ORDER BY p.evtlog_id DESC "

	rows,err := sw.S.Db().Query(query,winner_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleNFTWinnerRec,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleNFTWinnerRec
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.RoundNum,
			&rec.TokenId,
			&rec.WinnerIndex,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_raffle_deposits_by_user(winner_aid int64) []p.CGRaffleDepositRec {

	var query string
	query =  "SELECT " +
				"p.id,"+
				"p.evtlog_id,"+
				"p.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,"+
				"p.time_stamp,"+
				"p.winner_aid,"+
				"wa.addr,"+
				"p.round_num,"+
				"p.amount/1e18 amount_eth "+
			"FROM "+sw.S.SchemaName()+".cg_raffle_deposit p "+
				"LEFT JOIN transaction t ON t.id=p.tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"WHERE p.winner_aid = $1 " +
			"ORDER BY p.id DESC"

	rows,err := sw.S.Db().Query(query,winner_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleDepositRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleDepositRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.RoundNum,
			&rec.Amount,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records

}
func (sw *SQLStorageWrapper) Get_donated_nft_claims_by_user(winner_aid int64) []p.CGDonatedNFTClaimRec {

	var query string
	query = "SELECT "+
				"c.id,"+
				"c.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM c.time_stamp)::BIGINT,"+
				"c.time_stamp,"+
				"c.round_num,"+
				"ta.addr,"+
				"c.token_id, "+
				"c.idx, "+
				"c.winner_aid,"+
				"wa.addr, "+
				"da.addr, "+
				"d.token_uri "+
			"FROM "+sw.S.SchemaName()+".cg_donated_nft_claimed c "+
				"LEFT JOIN transaction t ON t.id=c.tx_id "+
				"LEFT JOIN address ta ON c.token_aid=ta.address_id "+
				"LEFT JOIN address wa ON c.winner_aid=wa.address_id "+
				"LEFT JOIN cg_nft_donation d ON d.idx=c.idx "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"WHERE c.winner_aid=$1 "+
			"ORDER BY c.id DESC "

	rows,err := sw.S.Db().Query(query,winner_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGDonatedNFTClaimRec,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGDonatedNFTClaimRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.RoundNum,
			&rec.TokenAddr,
			&rec.NFTTokenId,
			&rec.WinnerIndex,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.DonorAddr,
			&rec.NFTTokenURI,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_cosmic_signature_nft_list_by_user(user_aid int64,offset,limit int) []p.CGCosmicSignatureMintRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
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
				"m.token_name,"+
				"m.round_num,"+
				"p.prize_num "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event m "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
			"WHERE m.cur_owner_aid=$1 "+
			"ORDER BY m.id DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := sw.S.Db().Query(query,user_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicSignatureMintRec,0, 64)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicSignatureMintRec
		var null_prize_num sql.NullInt64
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.CurOwnerAid,
			&rec.CurOwnerAddr,
			&rec.Seed,
			&rec.TokenId,
			&rec.TokenName,
			&rec.RoundNum,
			&null_prize_num,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_prize_num.Valid { rec.RecordType = 3 } else {rec.RecordType = 1 }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_cosmic_token_transfers_by_user(user_aid int64,offset,limit int) []p.CGERC20TransferRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"t.id,"+
				"t.evtlog_id,"+
				"t.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,"+
				"t.time_stamp,"+
				"t.from_aid,"+
				"fa.addr,"+
				"t.to_aid,"+
				"ta.addr,"+
				"t.otype, "+
				"t.value,"+
				"t.value/1e18 "+ 
			"FROM "+sw.S.SchemaName()+".cg_erc20_transfer t "+
				"LEFT JOIN transaction tx ON tx.id=t.tx_id "+
				"LEFT JOIN address fa ON t.from_aid=fa.address_id "+
				"LEFT JOIN address ta ON t.to_aid=ta.address_id "+
			"WHERE (t.from_aid=$1) OR (t.to_aid=$1) "+
			"ORDER BY t.id DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := sw.S.Db().Query(query,user_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGERC20TransferRec,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGERC20TransferRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.FromAid,
			&rec.FromAddr,
			&rec.ToAid,
			&rec.ToAddr,
			&rec.TransferType,
			&rec.Value,
			&rec.ValueFloat,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_cosmic_signature_transfers_by_user(user_aid int64,offset,limit int) []p.CGTransfer {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"t.id,"+
				"t.evtlog_id,"+
				"t.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,"+
				"t.time_stamp,"+
				"t.from_aid,"+
				"fa.addr,"+
				"t.to_aid,"+
				"ta.addr,"+
				"t.otype, "+
				"t.token_id "+
			"FROM "+sw.S.SchemaName()+".cg_transfer t "+
				"LEFT JOIN transaction tx ON tx.id=t.tx_id "+
				"LEFT JOIN address fa ON t.from_aid=fa.address_id "+
				"LEFT JOIN address ta ON t.to_aid=ta.address_id "+
			"WHERE (t.from_aid=$1) OR (t.to_aid=$1) "+
			"ORDER BY t.id DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := sw.S.Db().Query(query,user_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGTransfer,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGTransfer
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.FromAid,
			&rec.FromAddr,
			&rec.ToAid,
			&rec.ToAddr,
			&rec.TransferType,
			&rec.TokenId,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_marketing_reward_history_by_user(user_aid int64,offset,limit int) []p.CGMarketingRewardRec {

	var query string
	query = "SELECT "+
					"r.id,"+
					"r.evtlog_id,"+
					"r.block_num,"+
					"tx.id,"+
					"tx.tx_hash,"+
					"EXTRACT(EPOCH FROM r.time_stamp)::BIGINT,"+
					"r.time_stamp,"+
					"r.amount,"+
					"r.amount/1e18,"+
					"r.marketer_aid,"+
					"ma.addr "+
				"FROM "+sw.S.SchemaName()+".cg_mkt_reward r "+
					"LEFT JOIN transaction tx ON tx.id=r.tx_id " +
					"LEFT JOIN address ma ON r.marketer_aid=ma.address_id "+
				"WHERE r.marketer_aid = $1 "+
				"ORDER BY r.id DESC " +
				"OFFSET $2 LIMIT $3 "

	rows,err := sw.S.Db().Query(query,user_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGMarketingRewardRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGMarketingRewardRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.Amount,
			&rec.AmountEth,
			&rec.MarketerAid,
			&rec.MarketerAddr,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_staked_tokens_cst_by_user(user_aid int64) []p.CGStakedTokenCSTRec {

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
				"EXTRACT(EPOCH FROM a.time_stamp)::BIGINT,"+
				"a.time_Stamp,"+
				"EXTRACT(EPOCH FROM a.unstake_time)::BIGINT,"+
				"a.unstake_time, "+
				"st.stake_action_id "+
			"FROM "+sw.S.SchemaName()+".cg_staked_token_cst st "+
				"LEFT JOIN cg_mint_event m ON st.token_id=m.token_id "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
				"LEFT JOIN cg_stake_action_cst a ON a.action_id=st.stake_action_id "+
			"WHERE st.staker_aid=$1 "+
			"ORDER BY m.token_id"

	rows,err := sw.S.Db().Query(query,user_aid)
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
			&rec.StakeTimeStamp,
			&rec.StakeDateTime,
			&rec.UnstakeTimeStamp,
			&rec.UnstakeDateTime,
			&rec.TokenInfo.StakeActionId,
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
func (sw *SQLStorageWrapper) Get_staked_tokens_rwalk_by_user(user_aid int64) []p.CGStakedTokenRWalkRec {

	var query string
	query = "SELECT "+
				"a.action_id,"+
				"EXTRACT(EPOCH FROM a.time_stamp)::BIGINT,"+
				"a.time_Stamp,"+
				"EXTRACT(EPOCH FROM a.unstake_time)::BIGINT,"+
				"a.unstake_time, "+
				"st.stake_action_id, "+
				"st.token_id "+
			"FROM "+sw.S.SchemaName()+".cg_staked_token_rwalk st "+
				"LEFT JOIN cg_mint_event m ON st.token_id=m.token_id "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
				"LEFT JOIN cg_stake_action_rwalk a ON a.action_id=st.stake_action_id "+
			"WHERE st.staker_aid=$1 "+
			"ORDER BY m.token_id"

	rows,err := sw.S.Db().Query(query,user_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGStakedTokenRWalkRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGStakedTokenRWalkRec 
		err=rows.Scan(
			&rec.StakeActionId,
			&rec.StakeTimeStamp,
			&rec.StakeDateTime,
			&rec.UnstakeTimeStamp,
			&rec.UnstakeDateTime,
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
func (sw *SQLStorageWrapper) Get_staking_rwalk_mints_by_user(user_aid int64) []p.CGRaffleNFTWinnerRec {

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
				"w.winner_index,"+
				"w.round_num,"+
				"w.winner_aid,"+
				"wa.addr "+
			"FROM cg_raffle_nft_winner w "+
				"LEFT JOIN w."+
				"LEFT JOIN transaction t ON t.id=w.tx_id "+
				"LEFT JOIN address wa ON w.winner_aid=wa.address_id "+
			"WHERE is_rwalk=TRUE AND is_staker=TRUE AMD w.winner_aid=$1 "+
			"ORDER BY w.evtlog_id DESC"
	rows,err := sw.S.Db().Query(query,user_aid)
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
func (sw *SQLStorageWrapper) Get_staking_actions_cst_by_user(user_aid int64,offset,limit int) []p.CGStakeActionCSTRec {

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
					"EXTRACT(EPOCH FROM s.unstake_time)::BIGINT AS usts,"+
					"s.unstake_time,"+
					"s.action_id,"+
					"s.token_id,"+
					"s.num_staked_nfts, "+
					"s.claimed "+
				"FROM "+sw.S.SchemaName()+".cg_stake_action_cst s "+
					"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
				"WHERE (s.staker_aid=$1) " +
				"OFFSET $2 LIMIT $3 "+
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
					"'F' AS claimed "+
				"FROM "+sw.S.SchemaName()+".cg_unstake_action_cst u "+
					"LEFT JOIN transaction tx ON tx.id=u.tx_id " +
					"LEFT JOIN cg_stake_action_cst s ON u.action_id=s.action_id "+
				"WHERE (u.staker_aid=$1) " +
				"OFFSET $2 LIMIT $3 "+
			") ORDER BY evtlog_id DESC"

	rows,err := sw.S.Db().Query(query,user_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGStakeActionCSTRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGStakeActionCSTRec
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
			&rec.Claimed,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	var accum_num_tokens int64
	for i:=len(records) - 1 ; i >= 0 ; i-- {
		if records[i].ActionType == 0 {
			accum_num_tokens = accum_num_tokens + 1
		} else {
			accum_num_tokens = accum_num_tokens - 1
		}
		records[i].NumStakedNFTs = accum_num_tokens
	}
	return records
}
func (sw *SQLStorageWrapper) Get_staking_actions_rwalk_by_user(user_aid int64,offset,limit int) []p.CGStakeActionRWalkRec {

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
					"EXTRACT(EPOCH FROM s.unstake_time)::BIGINT AS usts,"+
					"s.unstake_time,"+
					"s.action_id,"+
					"s.token_id,"+
					"s.num_staked_nfts, "+
				"FROM "+sw.S.SchemaName()+".cg_stake_action_rwalk s "+
					"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
				"WHERE (s.staker_aid=$1) " +
				"OFFSET $2 LIMIT $3 "+
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
					"'F' AS claimed "+
				"FROM "+sw.S.SchemaName()+".cg_unstake_action_rwalk u "+
					"LEFT JOIN transaction tx ON tx.id=u.tx_id " +
					"LEFT JOIN cg_stake_action_rwalk s ON u.action_id=s.action_id "+
				"WHERE (u.staker_aid=$1) " +
				"OFFSET $2 LIMIT $3 "+
			") ORDER BY evtlog_id DESC"

	rows,err := sw.S.Db().Query(query,user_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGStakeActionRWalkRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGStakeActionRWalkRec
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
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	var accum_num_tokens int64
	for i:=len(records) - 1 ; i >= 0 ; i-- {
		if records[i].ActionType == 0 {
			accum_num_tokens = accum_num_tokens + 1
		} else {
			accum_num_tokens = accum_num_tokens - 1
		}
		records[i].NumStakedNFTs = accum_num_tokens
	}
	return records
}
