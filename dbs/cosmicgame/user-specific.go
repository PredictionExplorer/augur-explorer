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
				"p.tokens_count,"+
				"p.unclaimed_nfts, "+
				"p.tokens_count, "+
				"trs.erc721_num_transfers, "+
				"d.count_donations,"+
				"d.total_eth_donated/1e18 "+
			"FROM address a "+
				"LEFT JOIN cg_bidder b ON b.bidder_aid=a.address_id "+
				"LEFT JOIN cg_winner p ON p.winner_aid=a.address_id "+
				"LEFT JOIN cg_donor d ON d.donor_aid=a.address_id "+
				"LEFT JOIN cg_raffle_winner_stats rw ON rw.winner_aid=a.address_id "+
				"LEFT JOIN cg_raffle_nft_winner_stats rn ON rn.winner_aid=a.address_id "+
				"LEFT JOIN cg_transfer_stats trs ON trs.user_aid=a.address_id "+
			"WHERE a.address_id=$1"

	var rec p.CGUserInfo
	var null_num_bids,null_prizes_count sql.NullInt64
	var null_max_bid,null_max_win sql.NullFloat64
	var null_raffle_sum_winnings,null_raffle_sum_withdrawal sql.NullFloat64
	var null_raffles_count,null_raffle_nft_won,null_reward_nfts sql.NullInt64
	var null_unclaimed_nfts,null_total_tokens sql.NullInt64
	var null_erc721_transfs sql.NullInt64
	var null_count_donations sql.NullInt64
	var null_total_eth_donated sql.NullFloat64


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
		&null_reward_nfts,
		&null_unclaimed_nfts,
		&null_total_tokens,
		&null_erc721_transfs,
		&null_count_donations,
		&null_total_eth_donated,
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
	if null_raffle_nft_won.Valid { rec.RaffleNFTsCount = null_raffle_nft_won.Int64 }
	if null_reward_nfts.Valid { rec.RewardNFTsCount = null_reward_nfts.Int64 }
	if null_unclaimed_nfts.Valid { rec.UnclaimedNFTs = null_unclaimed_nfts.Int64 }
	if null_total_tokens.Valid { rec.TotalCSTokensWon= null_total_tokens.Int64 }
	if null_erc721_transfs.Valid { rec.CosmicSignatureNumTransfers = null_erc721_transfs.Int64 }
	if null_count_donations.Valid { rec.TotalDonatedCount = null_count_donations.Int64 }
	if null_total_eth_donated.Valid { rec.TotalDonatedAmountEth = null_total_eth_donated.Float64 }

	// CST staking info moved to /ct/summary/by_user endpoint
	
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
			if err != sql.ErrNoRows {
				sw.S.Log_msg(fmt.Sprintf("Error in staker_rwalk query in Get_user_info(): %v, q=%v",err,query))
				os.Exit(1)
			}
		}
		if null_total_tokens_staked.Valid { rec.StakingStatisticsRWalk.TotalTokensStaked = null_total_tokens_staked.Int64 }
		if null_num_stake_actions.Valid { rec.StakingStatisticsRWalk.TotalNumStakeActions = null_num_stake_actions.Int64 }
		if null_num_unstake_actions.Valid { rec.StakingStatisticsRWalk.TotalNumUnstakeActions = null_num_unstake_actions.Int64 }
		if null_num_tokens_minted.Valid { rec.StakingStatisticsRWalk.TotalTokensMinted = null_num_tokens_minted.Int64 }
	}
	return true,rec
}
func (sw *SQLStorageWrapper) Get_prize_claims_by_user(winner_aid int64) []p.CGRoundRec {

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
				"p.round_num,"+
				"p.token_id,"+
				"m.seed,"+
				"s.total_bids,"+
				"s.total_nft_donated, "+
				"s.total_raffle_eth_deposits,"+
				"s.total_raffle_eth_deposits/1e18 eth_deposits,"+
				"s.total_raffle_nfts, "+
				"COALESCE(d.donation_amount,0),"+
				"COALESCE(d.donation_amount,0)/1e+18, "+
				"COALESCE(d.charity_addr,'0x0')"+
			"FROM "+sw.S.SchemaName()+".cg_prize_claim p "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
				"LEFT JOIN cg_mint_event m ON m.token_id=p.token_id "+
			"LEFT JOIN cg_round_stats s ON p.round_num=s.round_num "+
			"LEFT JOIN ("+
				"SELECT round_num, SUM(amount) as donation_amount, STRING_AGG(DISTINCT cha.addr, ', ') as charity_addr "+
					"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
					"LEFT JOIN "+sw.S.SchemaName()+".address cha ON d.contract_aid=cha.address_id "+
					"WHERE round_num >= 0 "+
					"GROUP BY round_num "+
			") d ON p.round_num = d.round_num "+
			"WHERE winner_aid=$1 "+
			"ORDER BY p.id DESC"

	rows,err := sw.S.Db().Query(query,winner_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	var null_seed sql.NullString
	records := make([]p.CGRoundRec,0, 32)
	defer rows.Close()
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
			&rec.MainPrize.EthAmount,
			&rec.MainPrize.EthAmountEth,
			&rec.RoundNum,
			&rec.MainPrize.NftTokenId,
			&null_seed,
			&rec.RoundStats.TotalBids,
			&rec.RoundStats.TotalDonatedNFTs,
			&rec.RoundStats.TotalRaffleEthDeposits,
			&rec.RoundStats.TotalRaffleEthDepositsEth,
			&rec.RoundStats.TotalRaffleNFTs,
			&rec.CharityDeposit.CharityAmount,
			&rec.CharityDeposit.CharityAmountETH,
			&rec.CharityDeposit.CharityAddress,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_seed.Valid { rec.MainPrize.Seed = null_seed.String } else {rec.MainPrize.Seed = "???"}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_bids_by_user(bidder_aid int64) []p.CGBidRec {

	query := sw.buildBidSelectQuery("b.bidder_aid=$1", "b.id DESC", "")

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
		var null_donated_erc20_addr,null_donated_erc20_amount sql.NullString
		var null_donated_erc20_amount_eth sql.NullFloat64
		err=rows.Scan(
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
			&rec.BidderAid,
		&rec.BidderAddr,
		&rec.EthPrice,
		&rec.EthPriceEth,
		&rec.CstPrice,
		&rec.CstPriceEth,
		&rec.RWalkNFTId,
		&null_token_id,
		&null_tok_addr,
		&null_token_uri,
		&rec.Message,
		&rec.RoundNum,
		&rec.CSTReward,
		&rec.CSTRewardEth,
		&rec.BidType,
		&rec.PrizeTime,
		&rec.PrizeTimeDate,
		&rec.TimeUntilPrize,
		&rec.BidPosition,
		&null_donated_erc20_addr,
		&null_donated_erc20_amount,
		&null_donated_erc20_amount_eth,
	)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.NFTDonationTokenId = -1
		if null_token_id.Valid { rec.NFTDonationTokenId=null_token_id.Int64 }
		if null_tok_addr.Valid { rec.NFTDonationTokenAddr = null_tok_addr.String }
		if null_token_uri.Valid { rec.NFTTokenURI = null_token_uri.String }
		if null_donated_erc20_addr.Valid { rec.DonatedERC20TokenAddr = null_donated_erc20_addr.String }
		if null_donated_erc20_amount.Valid { rec.DonatedERC20TokenAmount = null_donated_erc20_amount.String }
		if null_donated_erc20_amount_eth.Valid { rec.DonatedERC20TokenAmountEth = null_donated_erc20_amount_eth.Float64 }
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
				"JOIN "+sw.S.SchemaName()+".cg_prize_claim p ON p.round_num=d.round_num "+
				"LEFT JOIN cg_donated_nft_claimed c ON c.idx=d.idx "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=d.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address nft ON d.token_aid=nft.address_id "+
			"WHERE p.winner_aid=$1 AND p.round_num IS NOT NULL  AND c.idx IS NULL " +
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
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
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
				"p.winner_idx, "+
				"p.is_rwalk,"+
				"p.is_staker "+
			"FROM "+sw.S.SchemaName()+".cg_raffle_nft_prize p "+
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
func (sw *SQLStorageWrapper) Get_prize_deposits_chrono_warrior_by_user(winner_aid int64) []p.CGPrizeDepositRec {

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
			"p.eth_amount/1e18 amount_eth "+
		"FROM "+sw.S.SchemaName()+".cg_chrono_warrior_prize p "+
				"LEFT JOIN transaction t ON t.id=p.tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"WHERE p.winner_aid = $1 " +
			"ORDER BY p.id DESC"
	fmt.Printf("q = %v\n",query)
	fmt.Printf("user_aid= %v\n",winner_aid)
	rows,err := sw.S.Db().Query(query,winner_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGPrizeDepositRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGPrizeDepositRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.RoundNum,
			&rec.Amount,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.RecordType = 2
		records = append(records,rec)
	}
	return records

}
func (sw *SQLStorageWrapper) Get_prize_deposits_raffle_eth_by_user(winner_aid int64) []p.CGPrizeDepositRec {

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
			"FROM "+sw.S.SchemaName()+".cg_raffle_eth_prize p "+
				"LEFT JOIN transaction t ON t.id=p.tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"WHERE p.winner_aid = $1 " +
			"ORDER BY p.id DESC"

	rows,err := sw.S.Db().Query(query,winner_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGPrizeDepositRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGPrizeDepositRec
		err=rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.RoundNum,
			&rec.Amount,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.RecordType = 1
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
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
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
	query := sw.buildNFTSelectQuery("m.cur_owner_aid=$1", "m.id DESC", "OFFSET $2 LIMIT $3")
	rows, err := sw.S.Db().Query(query, user_aid, offset, limit)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	
	records := make([]p.CGCosmicSignatureMintRec, 0, 64)
	defer rows.Close()
	for rows.Next() {
		rec, err := scanNFTRecord(rows)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
			os.Exit(1)
		}
		records = append(records, rec)
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
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
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
			"FROM "+sw.S.SchemaName()+".cg_erc721_transfer t "+
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
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
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
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
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
				"p.round_num, "+
				"m.token_name, "+
				"EXTRACT(EPOCH FROM a.time_stamp)::BIGINT,"+
				"a.time_Stamp,"+
				"st.stake_action_id "+
			"FROM "+sw.S.SchemaName()+".cg_staked_token_cst st "+
				"LEFT JOIN cg_mint_event m ON st.token_id=m.token_id "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
				"LEFT JOIN cg_nft_staked_cst a ON a.action_id=st.stake_action_id "+
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
			&rec.TokenInfo.Tx.EvtLogId,
			&rec.TokenInfo.Tx.BlockNum,
			&rec.TokenInfo.Tx.TxId,
			&rec.TokenInfo.Tx.TxHash,
			&rec.TokenInfo.Tx.TimeStamp,
			&rec.TokenInfo.Tx.DateTime,
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
				"st.stake_action_id, "+
				"st.token_id "+
			"FROM "+sw.S.SchemaName()+".cg_staked_token_rwalk st "+
				"LEFT JOIN cg_mint_event m ON st.token_id=m.token_id "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
				"LEFT JOIN cg_nft_staked_rwalk a ON a.action_id=st.stake_action_id "+
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
			&rec.StakeActionId,
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
				"w.winner_idx,"+
				"w.round_num,"+
				"w.winner_aid,"+
				"wa.addr "+
			"FROM cg_raffle_nft_prize w "+
				"LEFT JOIN transaction t ON t.id=w.tx_id "+
				"LEFT JOIN address wa ON w.winner_aid=wa.address_id "+
			"WHERE is_rwalk=TRUE AND is_staker=TRUE AND w.winner_aid=$1 "+
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
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
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
func (sw *SQLStorageWrapper) Get_staking_cst_mints_by_user(user_aid int64) []p.CGRaffleNFTWinnerRec {

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
			"FROM cg_raffle_nft_prize w "+
				"LEFT JOIN transaction t ON t.id=w.tx_id "+
				"LEFT JOIN address wa ON w.winner_aid=wa.address_id "+
			"WHERE is_rwalk=FALSE AND is_staker=TRUE AND w.winner_aid=$1 "+
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
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
			&rec.TokenId,
			&rec.WinnerIndex,
			&rec.RoundNum,
			&rec.WinnerAid,
			&rec.WinnerAddr,
		)
		rec.IsRWalk = false
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
					"-1 AS usts,"+
					"TO_TIMESTAMP(0) AS unstake_time,"+
					"s.action_id,"+
					"s.token_id,"+
					"s.num_staked_nfts, "+
					"s.claimed "+
				"FROM "+sw.S.SchemaName()+".cg_nft_staked_cst s "+
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
					"TO_TIMESTAMP(0) AS unstake_time,"+
					"u.action_id,"+
					"u.token_id,"+
					"u.num_staked_nfts, "+
					"'F' AS claimed "+
				"FROM "+sw.S.SchemaName()+".cg_nft_unstaked_cst u "+
					"LEFT JOIN transaction tx ON tx.id=u.tx_id " +
					"LEFT JOIN cg_nft_staked_cst s ON u.action_id=s.action_id "+
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
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
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
					"-1 AS usts,"+
					"TO_TIMESTAMP(0) AS unstake_time,"+
					"s.action_id,"+
					"s.token_id,"+
					"s.num_staked_nfts "+
				"FROM "+sw.S.SchemaName()+".cg_nft_staked_rwalk s "+
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
					"u.num_staked_nfts "+
				"FROM "+sw.S.SchemaName()+".cg_nft_unstaked_rwalk u "+
					"LEFT JOIN transaction tx ON tx.id=u.tx_id " +
					"LEFT JOIN cg_nft_staked_rwalk s ON u.action_id=s.action_id "+
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
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
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
func (sw *SQLStorageWrapper) Get_user_notif_red_box_rewards(winner_aid int64) p.CGClaimInfo {

	var output p.CGClaimInfo
	var query string
	query = "SELECT " +
				"s.amount_sum,"+ 
				"s.amount_sum/1e18, " +
				"w.unclaimed_nfts  " +
			"FROM cg_raffle_winner_stats s " +
				"LEFT JOIN cg_winner w ON s.winner_aid=w.winner_aid "+
			"WHERE s.winner_aid = $1"

	row := sw.S.Db().QueryRow(query,winner_aid)
	var err error
	var null_wei sql.NullString
	var null_eth sql.NullFloat64
	var null_nfts sql.NullInt64

	err=row.Scan(&null_wei,&null_eth,&null_nfts);
	if err != nil {
		if err == sql.ErrNoRows {
		} else {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	if null_eth.Valid {
		output.ETHRaffleToClaim = null_eth.Float64
	}
	if null_wei.Valid {
		output.ETHRaffleToClaimWei = null_wei.String
	}
	if null_nfts.Valid {
		output.NumDonatedNFTToClaim = null_nfts.Int64
	}

	var null_staking_rewards sql.NullFloat64
	query = "SELECT unclaimed_reward/1e18 FROM cg_staker_cst WHERE staker_aid=$1"
	row = sw.S.Db().QueryRow(query,winner_aid)
	err=row.Scan(&null_staking_rewards);
	if err != nil {
		if err == sql.ErrNoRows {
			return output;
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	if null_staking_rewards.Valid {
		output.UnclaimedStakingReward = null_staking_rewards.Float64
	}
	query = "SELECT "+
				"p.round_num,"+
				"d.token_aid,"+
				"ta.addr,"+
				"d.total_amount, "+
				"d.total_amount/1e18 "+
			"FROM cg_prize_claim p "+
				"JOIN cg_erc20_donation_stats d ON d.round_num=p.round_num AND claimed='F' "+
				"LEFT JOIN address ta ON d.token_aid=ta.address_id "+
			"WHERE p.winner_aid=$1 "

	rows,err := sw.S.Db().Query(query,winner_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	output.DonatedERC20Tokens = make([]p.ERC20DonatedTokensInfo,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.ERC20DonatedTokensInfo
		err=rows.Scan(
			&rec.RoundNum,
			&rec.TokenAid,
			&rec.TokenAddr,
			&rec.Amount,
			&rec.AmountEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		output.DonatedERC20Tokens = append(output.DonatedERC20Tokens,rec)
	}
	return output
}
func (sw *SQLStorageWrapper) Get_erc20_donated_prizes_erc20_by_winner(user_aid int64) []p.CGSummarizedERC20Donation {

	var query string
	query = "WITH claim AS ("+
				"SELECT SUM(amount) total,round_num,token_aid,winner_aid "+
				"FROM cg_donated_tok_claimed GROUP BY round_num,token_aid,winner_aid "+
			") " + 
			"SELECT "+
				"p.id,"+
				"p.evtlog_id,"+
				"p.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,"+
				"p.time_stamp,"+
				"dt20.round_num,"+
				"tokaddr.address_id,"+
				"tokaddr.addr, "+
				"dt20.total_amount, "+
				"dt20.total_amount/1e18, "+
				"COALESCE(claim.total,0), "+
				"COALESCE(claim.total,0)/1e18, "+
				"dt20.total_amount-COALESCE(claim.total,0),"+
				"(dt20.total_amount-COALESCE(claim.total,0))/1e18,"+
				"dt20.winner_aid,"+
				"wa.addr, "+
				"dt20.claimed "+
			"FROM "+sw.S.SchemaName()+".cg_erc20_donation_stats dt20 "+
				"INNER JOIN cg_prize_claim p ON p.round_num=dt20.round_num "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=p.tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address tokaddr ON dt20.token_aid=tokaddr.address_id "+
				"LEFT JOIN claim ON (claim.token_aid=dt20.token_aid AND dt20.round_num=claim.round_num) "+
				"LEFT JOIN address wa ON wa.address_id = claim.winner_aid "+
			"WHERE p.winner_aid = $1 " +
			"ORDER BY dt20.token_aid"
//	fmt.Printf("q=%v\nuser_aid=%v\n",query,user_aid)
	rows,err := sw.S.Db().Query(query,user_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGSummarizedERC20Donation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGSummarizedERC20Donation
		var null_winner_addr sql.NullString
		var null_winner_aid sql.NullInt64
		err=rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			&rec.Tx.DateTime,
			&rec.RoundNum,
			&rec.TokenAid,
			&rec.TokenAddr,
			&rec.AmountDonated,
			&rec.AmountDonatedEth,
			&rec.AmountClaimed,
			&rec.AmountClaimedEth,
			&rec.DonateClaimDiff,
			&rec.DonateClaimDiffEth,
			&null_winner_aid,
			&null_winner_addr,
			&rec.Claimed,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_winner_aid.Valid { rec.WinnerAid=null_winner_aid.Int64 }
		if null_winner_addr.Valid { rec.WinnerAddr = null_winner_addr.String }
		records = append(records,rec)
	}
	return records
}
