package cosmicgame

import (
	"os"
	"fmt"
	"database/sql"


	p "github.com/PredictionExplorer/augur-explorer/primitives/cosmicgame"
)
func (sw *SQLStorageWrapper) Get_cosmic_game_statistics() p.CGStatistics {

	var stats p.CGStatistics
	var query string
	query = "SELECT "+
				"num_vol_donations, "+
				"vol_donations_total/1e18 as voluntary_donations_sum,"+
				"num_cg_donations,"+
				"cg_donations_total/1e18,"+
				"num_withdrawals,"+
				"sum_withdrawals/1e18,"+
				"num_bids," +
				"cur_num_bids,"+
				"num_wins, "+
				"num_rwalk_used, "+
				"num_mints, "+
				"total_raffle_eth_deposits/1e18, "+
				"total_raffle_eth_withdrawn/1e18, "+
				"total_nft_donated "+
			"FROM cg_glob_stats LIMIT 1"

	row := sw.S.Db().QueryRow(query)
	var err error
	err=row.Scan(
		&stats.NumVoluntaryDonations,
		&stats.SumVoluntaryDonationsEth,
		&stats.NumCosmicGameDonations,
		&stats.SumCosmicGameDonationsEth,
		&stats.NumWithdrawals,
		&stats.SumWithdrawals,
		&stats.TotalBids,
		&stats.CurNumBids,
		&stats.TotalPrizes,
		&stats.NumRwalkTokensUsed,
		&stats.NumCSTokenMints,
		&stats.TotalRaffleEthDeposits,
		&stats.TotalRaffleEthWithdrawn,
		&stats.TotalNFTDonated,
	)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var null_bidders sql.NullInt64
	query = "SELECT "+
				"COUNT(*) AS total "+
			"FROM cg_bidder " +
			"WHERE num_bids > 0 "
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(&null_bidders)
	if null_bidders.Valid  { stats.NumUniqueBidders = uint64(null_bidders.Int64) }
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var null_winners sql.NullInt64
	var null_sum_wei sql.NullString
	var null_sum_eth sql.NullFloat64
	query = "SELECT "+
				"COUNT(*) AS total,"+
				"SUM(prizes_sum) AS sum_wei,"+
				"SUM(prizes_sum)/1e18 AS sum_eth "+
				"FROM cg_winner " +
				"WHERE prizes_count > 0"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_winners,
		&null_sum_wei,
		&null_sum_eth,
	)
	if null_winners.Valid { stats.NumUniqueWinners = uint64(null_winners.Int64) }
	if null_sum_wei.Valid { stats.TotalPrizesPaidAmountWei = null_sum_wei.String }
	if null_sum_eth.Valid { stats.TotalPrizesPaidAmountEth = null_sum_eth.Float64 }
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	var null_donated_nfts sql.NullInt64
	query = "SELECT "+
				"SUM(num_donated) as total FROM cg_nft_stats"
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_donated_nfts,
	)
	stats.NumDonatedNFTs=uint64(null_donated_nfts.Int64)
	query = "SELECT count(*) AS total FROM cg_mint_event "+
			"WHERE LENGTH(token_name) > 0 "
	var null_named_tokens sql.NullInt64
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_named_tokens.Int64,
	)
	stats.TotalNamedTokens = null_named_tokens.Int64

	query = "SELECT count(winner_aid) AS total FROM cg_raffle_Winner_stats "+
			"WHERE amount_sum > 0 "
	var null_num_users_missing_withdrawal sql.NullInt64
	row = sw.S.Db().QueryRow(query)
	err=row.Scan(
		&null_num_users_missing_withdrawal.Int64,
	)
	stats.NumWinnersWithPendingRaffleWithdrawal = null_num_users_missing_withdrawal.Int64

	stats.DonatedTokenDistribution = sw.Get_donated_token_distribution();
	return stats
}
func (sw *SQLStorageWrapper) Get_cosmic_game_round_statistics(round_num int64) p.CGRoundStats {

	var stats p.CGRoundStats
	var query string
	query = "SELECT "+
				"round_num, "+
				"total_bids,"+
				"total_nft_donated," +
				"total_raffle_eth_deposits,"+
				"total_raffle_eth_deposits/1e18,"+
				"total_raffle_nfts "+
			"FROM cg_round_stats WHERE round_num=$1"

	row := sw.S.Db().QueryRow(query,round_num)
	var err error
	err=row.Scan(
		&stats.RoundNum,
		&stats.TotalBids,
		&stats.TotalDonatedNFTs,
		&stats.TotalRaffleEthDeposits,
		&stats.TotalRaffleEthDepositsEth,
		&stats.TotalRaffleNFTs,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return stats
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_cosmic_game_round_statistics(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return stats
}
func (sw *SQLStorageWrapper) Get_bids(offset,limit int) []p.CGBidRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
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
				"b.msg, "+
				"b.round_num, "+
				"b.num_cst_tokens, "+
				"b.num_cst_tokens/1e18, "+
				"b.bid_type "+
			"FROM "+sw.S.SchemaName()+".cg_bid b "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address ba ON b.bidder_aid=ba.address_id "+
				"LEFT JOIN LATERAL (" +
					"SELECT d.bid_id,token_id,token_aid,ta.addr tok_addr "+
						"FROM "+sw.S.SchemaName()+".cg_nft_donation d "+
						"JOIN "+sw.S.SchemaName()+".address ta ON d.token_aid=ta.address_id "+
				") d ON b.id=d.bid_id "+
			"ORDER BY b.id DESC "+
			"OFFSET $1 LIMIT $2"
	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGBidRec,0, 256)
	defer rows.Close()
	var null_token_id sql.NullInt64
	var null_tok_addr sql.NullString
	for rows.Next() {
		var rec p.CGBidRec
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
		records = append(records,rec)
	}
	return records
}
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
				"d.charity_addr "+
			"FROM "+sw.S.SchemaName()+".cg_prize_claim p "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
				"LEFT JOIN cg_mint_event m ON m.token_id=p.prize_num "+
				"LEFT JOIN cg_round_stats s ON p.prize_num=s.round_num "+
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
func (sw *SQLStorageWrapper) Get_bid_info(evtlog_id int64) (bool,p.CGBidRec) {

	var rec p.CGBidRec
	var query string
	query = "SELECT " +
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
			"WHERE b.evtlog_id=$1"

	row := sw.S.Db().QueryRow(query,evtlog_id)
	var err error
	var null_token_id sql.NullInt64
	var null_tok_addr,null_token_uri sql.NullString
	err=row.Scan(
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
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	rec.NFTDonationTokenId = -1
	if null_token_id.Valid { rec.NFTDonationTokenId=null_token_id.Int64 }
	if null_tok_addr.Valid { rec.NFTDonationTokenAddr = null_tok_addr.String }
	if null_token_uri.Valid { rec.NFTTokenURI = null_token_uri.String }
	return true,rec
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
				"d.charity_addr "+
			"FROM "+sw.S.SchemaName()+".cg_prize_claim p "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
				"LEFT JOIN cg_mint_event m ON m.token_id=p.token_id "+
				"LEFT JOIN cg_round_stats s ON s.round_num=p.prize_num "+
				"LEFT JOIN cg_winner ws ON p.winner_aid=ws.winner_aid "+
				"LEFT JOIN LATERAL (" +
					"SELECT d.evtlog_id,d.amount donation_amount,cha.addr charity_addr "+
						"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
						"JOIN "+sw.S.SchemaName()+".address cha ON d.contract_aid=cha.address_id "+
				") d ON p.donation_evt_id=d.evtlog_id "+
			"WHERE p.prize_num=$1"

	row := sw.S.Db().QueryRow(query,prize_num)
	var null_seed sql.NullString
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
	raffle_eth_deposits := sw.Get_raffle_deposits_by_round(prize_num)

	rec.RaffleNFTWinners = raffle_nft_winners
	rec.RaffleETHDeposits = raffle_eth_deposits
	return true,rec
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
func (sw *SQLStorageWrapper) Get_bids_by_round_num(round_num int64,sort,offset,limit int) []p.CGBidRec {

	var order_by string = " ASC "
	if sort == 1 {
		order_by = " DESC "
	}
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
			"WHERE b.round_num=$1 "+
			"ORDER BY b.id "+order_by+" OFFSET $2 LIMIT $3"

	rows,err := sw.S.Db().Query(query,round_num,offset,limit)
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
func (sw *SQLStorageWrapper) Get_unclaimed_raffle_eth_deposits(winner_aid int64,offset,limit int) []p.CGRaffleDepositRec {

	var query string
	query = 
			"SELECT "+
				"rd.id,"+
				"rd.evtlog_id,"+
				"rd.block_num,"+
				"rd.tx_id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM rd.time_stamp)::BIGINT AS tstmp, "+
				"rd.time_stamp AS date_time, "+
				"wa.addr,"+
				"rd.winner_aid,"+
				"rd.round_num,"+
				"rd.amount/1e18 AS amount_eth,"+
				"rd.claimed, "+
				"EXTRACT(EPOCH FROM rw.time_stamp)::BIGINT AS tstmp, "+
				"rw.time_stamp "+
			"FROM cg_raffle_deposit rd "+
				"LEFT JOIN cg_raffle_withdrawal rw ON rw.evtlog_id=rd.withdrawal_id "+
				"LEFT JOIN transaction t ON t.id=rd.tx_id "+
				"LEFT JOIN address wa ON rd.winner_aid = wa.address_id "+
			"WHERE rd.winner_aid=$1 AND rd.claimed='F' " +
			"ORDER BY rd.id DESC "+
			"OFFSET $2 LIMIT $3"
	rows,err := sw.S.Db().Query(query,winner_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleDepositRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleDepositRec
		var null_ts sql.NullInt64
		var null_date sql.NullString
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.WinnerAddr,
			&rec.WinnerAid,
			&rec.RoundNum,
			&rec.Amount,
			&rec.Claimed,
			&null_ts,
			&null_date,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if null_ts.Valid { rec.ClaimTimeStamp = null_ts.Int64 }
		if null_date.Valid { rec.ClaimDateTime = null_date.String }
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_claim_history_detailed(winner_aid int64,offset,limit int) []p.CGRaffleHistory {
	
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
						"-1 AS winner_index, "+
						"rd.claimed "+
					"FROM cg_raffle_deposit rd "+
						"LEFT JOIN transaction t ON t.id=rd.tx_id "+
					"WHERE rd.winner_aid=$1 "+
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
					"WHERE rn.winner_aid=$1 "+
				") UNION ALL (" +
					"SELECT "+
						"2 AS record_type,"+
						"d.evtlog_id,"+
						"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS tstmp, "+
						"p.time_stamp AS date_time, "+
						"p.block_num,"+
						"p.tx_id,"+
						"t.tx_hash,"+
						"p.prize_num,"+
						"0 AS amount,"+
						"0 AS amount_eth,"+
						"ta.addr token_addr, " +
						"d.token_id,"+
						"d.token_uri,"+
						"d.idx winner_index,"+
						"c.id IS NOT NULL as claimed "+
					"FROM cg_prize_claim p "+
						"JOIN cg_nft_donation d ON p.prize_num=d.round_num "+ 
						"LEFT JOIN transaction t ON t.id=p.tx_id "+
						"LEFT JOIN address ta ON d.token_aid=ta.address_id "+
						"LEFT JOIN cg_donated_nft_claimed c ON (c.round_num=p.prize_num) AND (d.idx=c.idx) "+
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
						"p.prize_num,"+
						"p.amount,"+
						"p.amount/1e18 AS amount_eth,"+
						"'' AS token_addr, " +
						"p.token_id,"+
						"'' AS token_uri,"+
						"-1 AS winner_index,"+
						"'T' AS claimed "+
					"FROM cg_prize_claim p "+
						"LEFT JOIN transaction t ON t.id=p.tx_id "+
					"WHERE p.winner_aid=$1 "+
				") UNION ALL (" +
					"SELECT "+
						"4 AS record_type,"+
						"d.evtlog_id,"+
						"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS tstmp, "+
						"d.time_stamp AS date_time, "+
						"d.block_num,"+
						"d.tx_id,"+
						"t.tx_hash,"+
						"d.deposit_num AS round_num,"+
						"r.reward, "+
						"r.reward/1e18 AS amount_eth,"+
						"'' AS token_addr, "+
						"-1 AS token_id,"+
						"'' AS token_uri,"+
						"-1 AS winner_index, "+
						"'T' AS claimed "+
					"FROM cg_eth_deposit d "+
						"LEFT JOIN transaction t ON t.id=d.tx_id "+
						"LEFT JOIN cg_claim_reward r ON d.deposit_num=r.deposit_id "+
					"WHERE (r.staker_aid=$1) AND (r.id IS NOT NULL) "+
				") "+
			") everything " +
			"ORDER BY evtlog_id DESC " +
			"OFFSET $2 LIMIT $3"

	rows,err := sw.S.Db().Query(query,winner_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleHistory,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleHistory
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
func (sw *SQLStorageWrapper) Get_claim_history_detailed_global(offset,limit int) []p.CGRaffleHistory {
	
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
					"FROM cg_raffle_deposit rd "+
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
				") UNION ALL (" +
					"SELECT "+
						"2 AS record_type,"+
						"d.evtlog_id,"+
						"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS tstmp, "+
						"p.time_stamp AS date_time, "+
						"p.block_num,"+
						"p.tx_id,"+
						"t.tx_hash,"+
						"p.prize_num,"+
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
						"JOIN cg_nft_donation d ON p.prize_num=d.round_num "+ 
						"LEFT JOIN transaction t ON t.id=p.tx_id "+
						"LEFT JOIN address ta ON d.token_aid=ta.address_id "+
						"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
						"LEFT JOIN cg_donated_nft_claimed c ON (c.round_num=p.prize_num) AND (d.idx=c.idx) "+
				") UNION ALL (" +
					"SELECT "+
						"3 AS record_type,"+
						"p.evtlog_id,"+
						"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT AS tstmp, "+
						"p.time_stamp AS date_time, "+
						"p.block_num,"+
						"p.tx_id,"+
						"t.tx_hash,"+
						"p.prize_num,"+
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
					"SELECT "+
						"4 AS record_type,"+
						"d.evtlog_id,"+
						"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS tstmp, "+
						"d.time_stamp AS date_time, "+
						"d.block_num,"+
						"d.tx_id,"+
						"t.tx_hash,"+
						"d.deposit_num AS round_num,"+
						"r.reward, "+
						"r.reward/1e18 AS amount_eth,"+
						"'' AS token_addr, "+
						"-1 AS token_id,"+
						"'' AS token_uri,"+
						"-1 AS winner_index, "+
						"'T' AS claimed, "+
						"wa.addr winner_addr,"+
						"r.staker_aid "+
					"FROM cg_eth_deposit d "+
						"LEFT JOIN transaction t ON t.id=d.tx_id "+
						"LEFT JOIN cg_claim_reward r ON d.deposit_num=r.deposit_id "+
						"LEFT JOIN address wa ON r.staker_aid=wa.address_id "+
					"WHERE (r.id IS NOT NULL) "+
				") "+
			") everything " +
			"ORDER BY evtlog_id DESC " +
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleHistory,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRaffleHistory
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
				"trs.erc721_num_transfers, "+
				"st.total_tokens_staked,"+
				"st.num_stake_actions,"+
				"total_reward,"+
				"total_reward/1e18,"+
				"unclaimed_reward,"+
				"unclaimed_reward/1e18 "+
			"FROM address a "+
				"LEFT JOIN cg_bidder b ON b.bidder_aid=a.address_id "+
				"LEFT JOIN cg_winner p ON p.winner_aid=a.address_id "+
				"LEFT JOIN cg_raffle_winner_stats rw ON rw.winner_aid=a.address_id "+
				"LEFT JOIN cg_raffle_nft_winner_stats rn ON rn.winner_aid=a.address_id "+
				"LEFT JOIN cg_transfer_stats trs ON trs.user_aid=a.address_id "+
				"LEFT JOIN cg_staker st ON st.staker_aid=a.address_id "+
			"WHERE a.address_id=$1"

	var rec p.CGUserInfo
	var null_num_bids,null_prizes_count sql.NullInt64
	var null_max_bid,null_max_win sql.NullFloat64
	var null_raffle_sum_winnings,null_raffle_sum_withdrawal sql.NullFloat64
	var null_raffles_count,null_raffle_nft_won sql.NullInt64
	var null_unclaimed_nfts,null_total_tokens sql.NullInt64
	var null_erc20_transfs,null_erc721_transfs sql.NullInt64
	var null_total_tokens_staked,null_num_stake_actions sql.NullInt64
	var null_total_reward,null_unclaimed_reward sql.NullString
	var null_total_reward_eth,null_unclaimed_reward_eth sql.NullFloat64


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
		&null_total_tokens_staked,
		&null_num_stake_actions,
		&null_total_reward,
		&null_total_reward_eth,
		&null_unclaimed_reward,
		&null_unclaimed_reward_eth,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_user_info(): %v, q=%v",err,query))
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
	if null_total_tokens_staked.Valid { rec.StakingStatistics.TotalTokensStaked = null_total_tokens_staked.Int64 }
	if null_num_stake_actions.Valid { rec.StakingStatistics.TotalNumStakeActions = null_num_stake_actions.Int64 }
	if null_total_reward.Valid { rec.StakingStatistics.TotalReward = null_total_reward.String }
	if null_total_reward_eth.Valid { rec.StakingStatistics.TotalRewardEth = null_total_reward_eth.Float64 }
	if null_unclaimed_reward.Valid { rec.StakingStatistics.UnclaimedReward = null_unclaimed_reward.String }
	if null_unclaimed_reward_eth.Valid { rec.StakingStatistics.UnclaimedRewardEth = null_unclaimed_reward_eth.Float64 }

	return true,rec
}
func (sw *SQLStorageWrapper) Get_charity_donations(cosmicgame_aid int64) []p.CGCharityDonation{

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr,"+
				"d.amount, "+
				"d.amount/1e18 amount_eth,  " +
				"d.round_num "+
			"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"ORDER BY d.id DESC"
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCharityDonation,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCharityDonation
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.Amount,
			&rec.AmountEth,
			&rec.RoundNum,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		if rec.DonorAid == cosmicgame_aid { rec.IsVoluntary = false } else {rec.IsVoluntary=true}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_charity_donations_from_cosmic_game(cosmicgame_aid int64) []p.CGCharityDonation{

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr,"+
				"d.amount, "+
				"d.amount/1e18 amount_eth,  " +
				"d.round_num "+
			"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"WHERE donor_aid = $1 "+
			"ORDER BY d.id DESC"
	rows,err := sw.S.Db().Query(query,cosmicgame_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCharityDonation,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCharityDonation
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.Amount,
			&rec.AmountEth,
			&rec.RoundNum,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_charity_donations_voluntary(cosmicgame_aid int64) []p.CGCharityDonation{

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr,"+
				"d.amount, "+
				"d.amount/1e18 amount_eth,  " +
				"d.round_num "+
			"FROM "+sw.S.SchemaName()+".cg_donation_received d "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"WHERE donor_aid != $1 "+
			"ORDER BY d.id DESC"
	rows,err := sw.S.Db().Query(query,cosmicgame_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCharityDonation,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCharityDonation
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.Amount,
			&rec.AmountEth,
			&rec.RoundNum,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		rec.IsVoluntary=true
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_charity_wallet_withdrawals() []p.CGCharityWithdrawal {

	var query string
	query = "SELECT "+
				"w.id,"+
				"w.evtlog_id,"+
				"w.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM w.time_stamp)::BIGINT,"+
				"w.time_stamp,"+
				"ca.addr,"+
				"w.amount,"+
				"w.amount/1e18 "+
			"FROM "+sw.S.SchemaName()+".cg_donation_sent w "+
				"LEFT JOIN transaction tx ON tx.id=w.tx_id "+
				"LEFT JOIN address ca ON w.charity_aid=ca.address_id "+
			"ORDER BY w.id DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCharityWithdrawal,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCharityWithdrawal
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DestinationAddr,
			&rec.Amount,
			&rec.AmountEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_donations_to_cosmic_game() []p.CGCosmicGameDonation{

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr,"+
				"d.amount, "+
				"d.amount/1e18 amount_eth  " +
			"FROM "+sw.S.SchemaName()+".cg_donation d "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id "+
			"ORDER BY d.id DESC"
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicGameDonation,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicGameDonation
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.Amount,
			&rec.AmountEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_unique_bidders() []p.CGUniqueBidder {

	var query string
	query = "SELECT "+
				"b.bidder_aid,"+
				"a.addr,"+
				"b.num_bids,"+
				"b.max_bid,"+
				"b.max_bid/1e18 max_bid_eth "+
			"FROM "+sw.S.SchemaName()+".cg_bidder b "+
				"LEFT JOIN address a ON b.bidder_aid=a.address_id " +
			"ORDER BY num_bids DESC "
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGUniqueBidder,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGUniqueBidder
		err=rows.Scan(
			&rec.BidderAid,
			&rec.BidderAddr,
			&rec.NumBids,
			&rec.MaxBidAmount,
			&rec.MaxBidAmountEth,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_unique_winners() []p.CGUniqueWinner {

	var query string
	query = "SELECT "+
				"w.winner_aid,"+
				"a.addr,"+
				"w.prizes_count,"+
				"w.max_win_amount,"+
				"w.max_win_amount/1e18 max_win_eth, "+
				"w.prizes_sum/1e18 prizes_sum_eth "+
			"FROM "+sw.S.SchemaName()+".cg_winner w "+
				"LEFT JOIN address a ON w.winner_aid=a.address_id " +
			"WHERE w.prizes_count > 0 " +
			"ORDER BY prizes_count DESC "
	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGUniqueWinner,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGUniqueWinner
		err=rows.Scan(
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.PrizesCount,
			&rec.MaxWinAmount,
			&rec.MaxWinAmountEth,
			&rec.PrizesSum,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_NFT_donations(offset,limit int) []p.GNFTDonation{

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"d.id,"+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr, "+
				"d.token_id, "+
				"d.round_num,"+
				"nft.address_id,"+
				"d.idx,"+
				"nft.addr, "+
				"d.token_uri "+
			"FROM "+sw.S.SchemaName()+".cg_nft_donation d "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address nft ON d.token_aid=nft.address_id "+
			"ORDER BY d.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.GNFTDonation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.GNFTDonation
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.NFTTokenId,
			&rec.RoundNum,
			&rec.TokenAddressId,
			&rec.Index,
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
func (sw *SQLStorageWrapper) Get_NFT_donation_stats() []p.GNFTDonationStats {

	var query string
	query = "SELECT "+
				"s.contract_aid,"+
				"a.addr,"+
				"s.num_donated "+
			"FROM "+sw.S.SchemaName()+".cg_nft_stats s " +
				"LEFT JOIN address a ON s.contract_aid=a.address_id "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.GNFTDonationStats,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.GNFTDonationStats
		err=rows.Scan(
			&rec.TokenAddressId,
			&rec.TokenAddress,
			&rec.NumDonations,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_record_counters() p.CGRecordCounters {

	var output p.CGRecordCounters
	var null_total_bids,null_total_prizes,null_total_tok_donations sql.NullInt64
	var query string
	query = "SELECT count(*) AS total FROM "+sw.S.SchemaName()+".cg_bid"
	res := sw.S.Db().QueryRow(query)
	err := res.Scan(&null_total_bids)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	} else { output.TotalBids = null_total_bids.Int64 }

	query = "SELECT count(*) AS total FROM "+sw.S.SchemaName()+".cg_prize_claim"
	res = sw.S.Db().QueryRow(query)
	err = res.Scan(&null_total_prizes)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	} else { output.TotalPrizes= null_total_prizes.Int64 }

	query = "SELECT count(*) AS total FROM "+sw.S.SchemaName()+".cg_nft_donation"
	res = sw.S.Db().QueryRow(query)
	err = res.Scan(&null_total_tok_donations)
	if (err!=nil) {
		if err != sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v q=%v",err,query))
			os.Exit(1)
		}
	} else { output.TotalDonatedNFTs = null_total_tok_donations.Int64 }

	return output
}
func (sw *SQLStorageWrapper) Get_NFT_donation_info(id int64) (bool,p.GNFTDonation) {

	var query string
	query = "SELECT "+
				"d.evtlog_id,"+
				"d.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"d.donor_aid,"+
				"da.addr, "+
				"d.token_id, "+
				"nft.address_id,"+
				"nft.addr, "+
				"d.token_uri "+
			"FROM "+sw.S.SchemaName()+".cg_nft_donation d "+
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address nft ON d.token_aid=nft.address_id "+
			"WHERE d.id=$1"

	row := sw.S.Db().QueryRow(query,id)
	var err error
	var rec p.GNFTDonation
	rec.RecordId = id
	err=row.Scan(
		&rec.EvtLogId,
		&rec.BlockNum,
		&rec.TxId,
		&rec.TxHash,
		&rec.TimeStamp,
		&rec.DateTime,
		&rec.DonorAid,
		&rec.DonorAddr,
		&rec.NFTTokenId,
		&rec.TokenAddressId,
		&rec.TokenAddr,
		&rec.NFTTokenURI,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_NFT_donation_info(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return true,rec
}
func (sw *SQLStorageWrapper) Get_raffle_eth_deposits_list(offset,limit int) []p.CGRaffleDepositRec {

	if limit == 0 { limit = 1000000 }
	var query string
	query = "SELECT "+
				"p.id,"+
				"p.evtlog_id,"+
				"p.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,"+
				"p.time_stamp,"+
				"p.winner_aid,"+
				"wa.addr,"+
				"p.round_num, "+
				"p.amount/1e18 amount_eth "+
			"FROM "+sw.S.SchemaName()+".cg_raffle_deposit p "+
				"LEFT JOIN transaction t ON t.id=p.tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"ORDER BY p.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRaffleDepositRec,0, 256)
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
func (sw *SQLStorageWrapper) Get_raffle_nft_winners_by_round(round_num int64) []p.CGRaffleNFTWinnerRec {

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
				"p.winner_idx "+
			"FROM "+sw.S.SchemaName()+".cg_raffle_nft_winner p "+
				"LEFT JOIN transaction t ON t.id=p.tx_id "+
				"LEFT JOIN address wa ON p.winner_aid=wa.address_id "+
			"WHERE p.round_num=$1 " +
			"ORDER BY p.id DESC"

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
func (sw *SQLStorageWrapper) Get_unclaimed_donated_nft_by_user(winner_aid int64) []p.GNFTDonation {

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
	records := make([]p.GNFTDonation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.GNFTDonation
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
				"p.winner_idx "+
			"FROM "+sw.S.SchemaName()+".cg_raffle_nft_winner p "+
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
func (sw *SQLStorageWrapper) Get_raffle_deposits_by_round(round_num int64) []p.CGRaffleDepositRec {

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
			"WHERE p.round_num = $1 " +
			"ORDER BY p.id DESC "

	rows,err := sw.S.Db().Query(query,round_num)
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
func (sw *SQLStorageWrapper) Get_donated_nft_claims(offset,limit int) []p.CGDonatedNFTClaimRec {

	if limit == 0 { limit = 1000000 }
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
				"LEFT JOIN cg_nft_donation d ON c.idx=d.idx "+
				"LEFT JOIN address da ON d.donor_aid=da.address_id " +
			"ORDER BY c.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
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
func (sw *SQLStorageWrapper) Get_user_global_winnings(winner_aid int64) p.CGClaimInfo {

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
			return output;
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
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
	return output
}
func (sw *SQLStorageWrapper) Get_num_prize_claims() int64 {

	var query string
	query = "SELECT num_wins FROM cg_glob_stats"
	row := sw.S.Db().QueryRow(query)
	var err error
	var null_num_claims sql.NullInt64
	err=row.Scan(&null_num_claims);
	if err != nil {
		if err == sql.ErrNoRows {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
	}
	return null_num_claims.Int64
}
func (sw *SQLStorageWrapper) Get_nft_donations_by_prize(prize_num int64) []p.GNFTDonation {

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
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address nft ON d.token_aid=nft.address_id "+
			"WHERE d.round_num= $1 " +
			"ORDER BY d.id DESC"

	rows,err := sw.S.Db().Query(query,prize_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.GNFTDonation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.GNFTDonation
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
func (sw *SQLStorageWrapper) Get_unclaimed_donated_nfts_by_prize(prize_num int64) []p.GNFTDonation {

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
				"LEFT JOIN "+sw.S.SchemaName()+".transaction t ON t.id=tx_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address da ON d.donor_aid=da.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".address nft ON d.token_aid=nft.address_id "+
				"LEFT JOIN "+sw.S.SchemaName()+".cg_donated_nft_claimed dc ON d.idx = dc.idx "+
			"WHERE d.round_num= $1 AND dc.idx IS NULL " +
			"ORDER BY d.id DESC"

	rows,err := sw.S.Db().Query(query,prize_num)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.GNFTDonation,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.GNFTDonation
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
func (sw *SQLStorageWrapper) Get_cosmic_signature_nft_list(offset,limit int) []p.CGCosmicSignatureMintRec {

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
			"ORDER BY m.id DESC "+
			"OFFSET $1 LIMIT $2"

	rows,err := sw.S.Db().Query(query,offset,limit)
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
func (sw *SQLStorageWrapper) Get_cosmic_signature_token_info(token_id int64) (bool,p.CGCosmicSignatureMintRec) {

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
				"m.round_num,"+
				"p.prize_num, "+
				"m.token_name "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event m "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
				"LEFT JOIN address wa ON m.owner_aid=wa.address_id "+
				"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id "+
				"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id "+
			"WHERE m.token_id=$1"

	var rec p.CGCosmicSignatureMintRec
	var err error
	var null_prize_num sql.NullInt64
	row := sw.S.Db().QueryRow(query,token_id)
	err=row.Scan(
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
		&rec.RoundNum,
		&null_prize_num,
		&rec.TokenName,
	)
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return false,rec
		}
		sw.S.Log_msg(fmt.Sprintf("DB Error: %v, q=%v",err,query))
		os.Exit(1)
	}
	if null_prize_num.Valid { rec.RecordType = 3 } else {rec.RecordType = 1 }
	return true,rec
}
func (sw *SQLStorageWrapper) Get_cosmic_signature_token_name_history(token_id int64) []p.CGTokenName {

	var query string
	query = "SELECT "+
				"n.evtlog_id,"+
				"n.block_num,"+
				"t.id,"+
				"t.tx_hash,"+
				"EXTRACT(EPOCH FROM n.time_stamp)::BIGINT,"+
				"n.time_stamp,"+
				"n.token_id,"+
				"n.token_name "+
			"FROM "+sw.S.SchemaName()+".cg_token_name n "+
				"LEFT JOIN transaction t ON t.id=tx_id "+
			"WHERE n.token_id=$1 "+
			"ORDER BY n.id DESC "

	rows,err := sw.S.Db().Query(query,token_id)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGTokenName,0, 64)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGTokenName
		err=rows.Scan(
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.TokenId,
			&rec.TokenName,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_round_start_timestamp(round_num uint64) int64  {

	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT "+
			"FROM cg_bid b "+
			"WHERE round_num=$1 "+
			"ORDER BY b.id LIMIT 1"

	var null_ts sql.NullInt64
	row := sw.S.Db().QueryRow(query,round_num)
	var err error
	err=row.Scan(&null_ts);
	if (err!=nil) {
		if err == sql.ErrNoRows {
			return 0
		}
		sw.S.Log_msg(fmt.Sprintf("Error in Get_round_start_timestamp(): %v, q=%v",err,query))
		os.Exit(1)
	}
	return null_ts.Int64
}
func (sw *SQLStorageWrapper) Get_cst_ownership_transfers(token_id int64,offset,limit int) []p.CGTransfer {

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
				"t.token_id,"+
				"t.otype "+
			"FROM "+sw.S.SchemaName()+".cg_transfer t "+
				"LEFT JOIN transaction tx ON tx.id=t.tx_id "+
				"LEFT JOIN address fa ON t.from_aid=fa.address_id "+
				"LEFT JOIN address ta ON t.to_aid=ta.address_id "+
			"WHERE t.token_id=$1 "+
			"ORDER BY t.id DESC "+
			"OFFSET $2 LIMIT $3"

	rows,err := sw.S.Db().Query(query,token_id,offset,limit)
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
			&rec.TokenId,
			&rec.TransferType,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_cosmic_signature_token_distribution() []p.CGCSTokenDistributionRec {

	var query string
	query = "SELECT "+
				"m.cur_owner_aid, "+
				"a.addr, "+
				"count(m.cur_owner_aid) AS counter "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event m "+
				"LEFT JOIN address a ON m.cur_owner_aid=a.address_id "+
			"GROUP BY m.cur_owner_aid,a.addr "+
			"ORDER BY counter DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCSTokenDistributionRec,0, 32)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCSTokenDistributionRec
		err=rows.Scan(
			&rec.OwnerAid,
			&rec.OwnerAddr,
			&rec.NumTokens,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_random_walk_tokens_in_bids() []p.CGRWalkUsed {

	var query string
	query = "SELECT "+
				"b.id,"+
				"b.evtlog_id,"+
				"b.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT,"+
				"b.time_stamp,"+
				"b.round_num,"+
				"b.bidder_aid,"+
				"ba.addr,"+
				"b.rwalk_nft_id "+
			"FROM "+sw.S.SchemaName()+".cg_bid b "+
				"LEFT JOIN transaction tx ON tx.id=b.tx_id "+
				"LEFT JOIN address ba ON b.bidder_aid=ba.address_id "+
			"WHERE b.rwalk_nft_id != -1 "+
			"ORDER BY b.id DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGRWalkUsed,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRWalkUsed
		err=rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			&rec.DateTime,
			&rec.RoundNum,
			&rec.BidderAid,
			&rec.BidderAddr,
			&rec.RWalkTokenId,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Search_token_by_name(name string) []p.CGTokenSearchResult {

	name = "%" + name + "%"
	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,"+
				"t.time_stamp,"+
				"t.token_id,"+
				"t.token_name "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event t "+
			"WHERE t.token_name ILIKE  $1 "+
			"ORDER BY t.token_id"

	rows,err := sw.S.Db().Query(query,name)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGTokenSearchResult,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGTokenSearchResult
		err=rows.Scan(
			&rec.MintTimeStamp,
			&rec.MintDateTime,
			&rec.TokenId,
			&rec.TokenName,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_named_tokens() []p.CGTokenSearchResult {

	var query string
	query = "SELECT "+
				"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT,"+
				"t.time_stamp,"+
				"t.token_id,"+
				"t.token_name "+
			"FROM "+sw.S.SchemaName()+".cg_mint_event t "+
			"WHERE LENGTH(t.token_name)>0 "+
			"ORDER BY t.token_name"

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGTokenSearchResult,0, 256)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGTokenSearchResult
		err=rows.Scan(
			&rec.MintTimeStamp,
			&rec.MintDateTime,
			&rec.TokenId,
			&rec.TokenName,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_donated_token_distribution() []p.CGDonatedTokenDistrRec {

	var query string
	query = "SELECT "+
				"ca.addr,"+
				"ns.num_donated "+
			"FROM "+sw.S.SchemaName()+".cg_nft_stats ns "+
				"LEFT JOIN address ca ON ns.contract_aid=ca.address_id "+
			"ORDER BY ns.num_donated DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGDonatedTokenDistrRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGDonatedTokenDistrRec
		err=rows.Scan(
			&rec.ContractAddr,
			&rec.NumDonatedTokens,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	return records
}
func (sw *SQLStorageWrapper) Get_cosmic_token_holders() []p.CGCosmicTokenHolderRec {

	var query string
	query = "SELECT "+
				"o.owner_aid,"+
				"oa.addr,"+
				"o.cur_balance,"+
				"o.cur_balance/1e18 " +
			"FROM "+sw.S.SchemaName()+".cg_costok_owner o "+
				"LEFT JOIN address oa ON o.owner_aid=oa.address_id "+
			"ORDER BY o.cur_balance DESC "

	rows,err := sw.S.Db().Query(query)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGCosmicTokenHolderRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGCosmicTokenHolderRec
		err=rows.Scan(
			&rec.OwnerAid,
			&rec.OwnerAddr,
			&rec.Balance,
			&rec.BalanceFloat,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
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
func (sw *SQLStorageWrapper) Get_staking_rewards_to_be_claimed(user_aid int64) []p.CGRewardToClaim {

	var query string
	query = "SELECT "+
				"d.id,"+
				"d.evtlog_id,"+
				"d.block_num,"+
				"tx.id,"+
				"tx.tx_hash,"+
				"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,"+
				"d.time_stamp,"+
				"EXTRACT(EPOCH FROM d.deposit_time)::BIGINT,"+
				"d.deposit_time,"+
				"d.deposit_num,"+
				"d.num_staked_nfts,"+
				"d.amount,"+
				"d.amount/1e18,"+
				"sd.tokens_staked,"+
				"sd.amount_to_claim,"+
				"sd.amount_to_claim/1e18,"+
				"amount/num_staked_nfts,"+
				"(amount/num_staked_nfts)/1e18, "+
				"modulo,"+
				"modulo/1e+18 "+
			"FROM "+sw.S.SchemaName()+".cg_eth_deposit d "+
				"LEFT JOIN transaction tx ON tx.id=d.tx_id " +
				"LEFT JOIN cg_staker_deposit sd ON d.deposit_num=sd.deposit_id "+
				"LEFT JOIN cg_claim_reward r ON (d.deposit_num=r.deposit_id) AND (sd.staker_aid=r.staker_aid) "+
			"WHERE (sd.staker_aid = $1) AND (r.id IS NULL)" +
			"ORDER BY d.id DESC "

	rows,err := sw.S.Db().Query(query,user_aid)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	//fmt.Printf("user_aid=%v\n",user_aid)
	//fmt.Printf("q = %v\n",query)
	records := make([]p.CGRewardToClaim,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGRewardToClaim
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
			&rec.AmountPerToken,
			&rec.AmountPerTokenEth,
			&rec.Modulo,
			&rec.ModuloF64,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
			os.Exit(1)
		}
		records = append(records,rec)
	}
	fmt.Printf("len = %v\n",len(records))
	return records
}
func (sw *SQLStorageWrapper) Get_staking_actions(user_aid int64,offset,limit int) []p.CGStakeActionRec {

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
				"FROM "+sw.S.SchemaName()+".cg_stake_action s "+
					"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
				"WHERE (s.staker_aid=$1) " +
				"ORDER BY s.id DESC " +
				"OFFSET $2 LIMIT $3 "+
			") UNION ALL ("+
				"SELECT "+
					"1 AS action_type,"+
					"s.id,"+
					"s.evtlog_id,"+
					"s.block_num,"+
					"tx.id,"+
					"tx.tx_hash,"+
					"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT AS usts,"+
					"time_stamp,"+
					"0 AS usts,"+
					"TO_TIMESTAMP(0) AS unnstake_time,"+
					"s.action_id,"+
					"s.token_id,"+
					"s.num_staked_nfts, "+
					"'F' AS claimed "+
				"FROM "+sw.S.SchemaName()+".cg_unstake_action s "+
					"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
				"WHERE (s.staker_aid=$1) " +
				"ORDER BY s.id DESC " +
				"OFFSET $2 LIMIT $3 "+
			")"

	rows,err := sw.S.Db().Query(query,user_aid,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGStakeActionRec,0, 16)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGStakeActionRec
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
	return records
}
func (sw *SQLStorageWrapper) Get_global_staking_history(offset,limit int) []p.CGStakingHistoryRec {

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
					"s.staker_aid, "+
					"sa.addr staker_addr "+
				"FROM "+sw.S.SchemaName()+".cg_stake_action s "+
					"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
					"LEFT JOIN address sa ON s.staker_aid=sa.address_id "+
				"ORDER BY s.id DESC " +
				"OFFSET $1 LIMIT $2 "+
			") UNION ALL ("+
				"SELECT "+
					"1 AS action_type,"+
					"s.id,"+
					"s.evtlog_id,"+
					"s.block_num,"+
					"tx.id,"+
					"tx.tx_hash,"+
					"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT AS usts,"+
					"time_stamp,"+
					"0 AS usts,"+
					"TO_TIMESTAMP(0) AS unnstake_time,"+
					"s.action_id,"+
					"s.token_id,"+
					"s.num_staked_nfts, "+
					"s.staker_aid," +
					"sa.addr staker_addr "+
				"FROM "+sw.S.SchemaName()+".cg_unstake_action s "+
					"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
					"LEFT JOIN address sa ON s.staker_aid=sa.address_id "+
				"ORDER BY s.id DESC " +
				"OFFSET $1 LIMIT $2 "+
			") order by evtlog_id DESC"

	rows,err := sw.S.Db().Query(query,offset,limit)
	if (err!=nil) {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)",err,query))
		os.Exit(1)
	}
	records := make([]p.CGStakingHistoryRec,0, 16)
	accum_staked_nfts := int64(0)
	defer rows.Close()
	for rows.Next() {
		var rec p.CGStakingHistoryRec
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
		records = append(records,rec)
	}
	return records
}
