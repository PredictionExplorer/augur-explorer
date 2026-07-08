package cosmicgame

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// UserInfo returns the aggregate profile of one address (bids, prizes,
// raffle winnings, donations, transfer counts, RandomWalk staking totals),
// or store.ErrNotFound when the address id does not exist.
func (r *Repo) UserInfo(ctx context.Context, userAid int64) (p.CGUserInfo, error) {
	const op = "user info"
	query := "SELECT " +
		"a.address_id," +
		"a.addr, " +
		"b.num_bids, " +
		"b.max_bid/1e18 AS max_bid," +
		"p.prizes_count," +
		"p.max_win_amount/1e18 max_win, " +
		"rw.amount_sum/1e18 raffle_win_sum, " +
		"rw.withdrawal_sum/1e18 withdrawal_sum, " +
		"rw.raffles_count, " +
		"rn.num_won raffle_nft_won, " +
		"p.erc721_count," +
		"p.unclaimed_nfts, " +
		"p.erc721_count, " +
		"trs.erc721_num_transfers, " +
		"d.count_donations," +
		"d.total_eth_donated/1e18 " +
		"FROM address a " +
		"LEFT JOIN cg_bidder b ON b.bidder_aid=a.address_id " +
		"LEFT JOIN cg_winner p ON p.winner_aid=a.address_id " +
		"LEFT JOIN cg_donor d ON d.donor_aid=a.address_id " +
		"LEFT JOIN cg_raffle_winner_stats rw ON rw.winner_aid=a.address_id " +
		"LEFT JOIN cg_raffle_nft_winner_stats rn ON rn.winner_aid=a.address_id " +
		"LEFT JOIN cg_transfer_stats trs ON trs.user_aid=a.address_id " +
		"WHERE a.address_id=$1"

	var rec p.CGUserInfo
	var nullNumBids, nullPrizesCount sql.NullInt64
	var nullMaxBid, nullMaxWin sql.NullFloat64
	var nullRaffleSumWinnings, nullRaffleSumWithdrawal sql.NullFloat64
	var nullRafflesCount, nullRaffleNftWon, nullRewardNfts sql.NullInt64
	var nullUnclaimedNfts, nullTotalTokens sql.NullInt64
	var nullErc721Transfs sql.NullInt64
	var nullCountDonations sql.NullInt64
	var nullTotalEthDonated sql.NullFloat64

	err := r.pool().QueryRow(ctx, query, userAid).Scan(
		&rec.AddressId,
		&rec.Address,
		&nullNumBids,
		&nullMaxBid,
		&nullPrizesCount,
		&nullMaxWin,
		&nullRaffleSumWinnings,
		&nullRaffleSumWithdrawal,
		&nullRafflesCount,
		&nullRaffleNftWon,
		&nullRewardNfts,
		&nullUnclaimedNfts,
		&nullTotalTokens,
		&nullErc721Transfs,
		&nullCountDonations,
		&nullTotalEthDonated,
	)
	if err != nil {
		return p.CGUserInfo{}, store.WrapError(op, err)
	}
	if nullNumBids.Valid {
		rec.NumBids = nullNumBids.Int64
	}
	if nullPrizesCount.Valid {
		rec.NumPrizes = nullPrizesCount.Int64
	}
	if nullMaxBid.Valid {
		rec.MaxBidAmount = nullMaxBid.Float64
	}
	if nullMaxWin.Valid {
		rec.MaxWinAmount = nullMaxWin.Float64
	}
	if nullRaffleSumWinnings.Valid {
		rec.SumRaffleEthWinnings = nullRaffleSumWinnings.Float64
	}
	if nullRaffleSumWithdrawal.Valid {
		rec.SumRaffleEthWithdrawal = nullRaffleSumWithdrawal.Float64
	}
	if nullRafflesCount.Valid {
		rec.NumRaffleEthWinnings = nullRafflesCount.Int64
	}
	if nullRaffleNftWon.Valid {
		rec.RaffleNFTsCount = nullRaffleNftWon.Int64
	}
	if nullRewardNfts.Valid {
		rec.RewardNFTsCount = nullRewardNfts.Int64
	}
	if nullUnclaimedNfts.Valid {
		rec.UnclaimedNFTs = nullUnclaimedNfts.Int64
	}
	if nullTotalTokens.Valid {
		rec.TotalCSTokensWon = nullTotalTokens.Int64
	}
	if nullErc721Transfs.Valid {
		rec.CosmicSignatureNumTransfers = nullErc721Transfs.Int64
	}
	if nullCountDonations.Valid {
		rec.TotalDonatedCount = nullCountDonations.Int64
	}
	if nullTotalEthDonated.Valid {
		rec.TotalDonatedAmountEth = nullTotalEthDonated.Float64
	}

	// RandomWalk staking totals live in their own row; a user who never
	// staked has none and keeps the zero values (CST staking info moved to
	// the /ct/summary/by_user endpoint).
	query = "SELECT " +
		"s.total_tokens_staked," +
		"s.num_stake_actions," +
		"s.num_unstake_actions," +
		"s.num_tokens_minted " +
		"FROM cg_staker_rwalk s " +
		"WHERE staker_aid=$1"
	var nullTotalTokensStaked, nullNumStakeActions, nullNumUnstakeActions, nullNumTokensMinted sql.NullInt64
	err = r.pool().QueryRow(ctx, query, userAid).Scan(
		&nullTotalTokensStaked,
		&nullNumStakeActions,
		&nullNumUnstakeActions,
		&nullNumTokensMinted,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return p.CGUserInfo{}, store.WrapError(op+": rwalk staking totals", err)
	}
	if nullTotalTokensStaked.Valid {
		rec.StakingStatisticsRWalk.TotalTokensStaked = nullTotalTokensStaked.Int64
	}
	if nullNumStakeActions.Valid {
		rec.StakingStatisticsRWalk.TotalNumStakeActions = nullNumStakeActions.Int64
	}
	if nullNumUnstakeActions.Valid {
		rec.StakingStatisticsRWalk.TotalNumUnstakeActions = nullNumUnstakeActions.Int64
	}
	if nullNumTokensMinted.Valid {
		rec.StakingStatisticsRWalk.TotalTokensMinted = nullNumTokensMinted.Int64
	}
	return rec, nil
}

// PrizeClaimsByUser returns the main prizes claimed by one winner, newest
// first, each with its round statistics and charity deposit summary.
func (r *Repo) PrizeClaimsByUser(ctx context.Context, winnerAid int64) ([]p.CGRoundRec, error) {
	query := "SELECT " +
		"p.evtlog_id," +
		"p.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT," +
		"p.time_stamp," +
		"p.winner_aid," +
		"wa.addr," +
		"p.amount, " +
		"p.amount/1e18 amount_eth, " +
		"p.round_num," +
		"p.token_id," +
		"m.seed," +
		"s.total_bids," +
		"s.total_nft_donated, " +
		"s.total_raffle_eth_deposits," +
		"s.total_raffle_eth_deposits/1e18 eth_deposits," +
		"s.total_raffle_nfts, " +
		"COALESCE(d.donation_amount,0)," +
		"COALESCE(d.donation_amount,0)/1e+18, " +
		"COALESCE(d.charity_addr,'0x0')" +
		"FROM cg_prize_claim p " +
		"LEFT JOIN transaction t ON t.id=tx_id " +
		"LEFT JOIN address wa ON p.winner_aid=wa.address_id " +
		"LEFT JOIN cg_mint_event m ON m.token_id=p.token_id " +
		"LEFT JOIN cg_round_stats s ON p.round_num=s.round_num " +
		"LEFT JOIN (" +
		"SELECT round_num, SUM(amount) as donation_amount, STRING_AGG(DISTINCT cha.addr, ', ') as charity_addr " +
		"FROM cg_donation_received d " +
		"LEFT JOIN address cha ON d.contract_aid=cha.address_id " +
		"WHERE round_num >= 0 " +
		"GROUP BY round_num " +
		") d ON p.round_num = d.round_num " +
		"WHERE winner_aid=$1 " +
		"ORDER BY p.id DESC"
	scan := func(rows pgx.Rows, rec *p.CGRoundRec) error {
		var nullSeed sql.NullString
		err := rows.Scan(
			&rec.ClaimPrizeTx.Tx.EvtLogId,
			&rec.ClaimPrizeTx.Tx.BlockNum,
			&rec.ClaimPrizeTx.Tx.TxId,
			&rec.ClaimPrizeTx.Tx.TxHash,
			&rec.ClaimPrizeTx.Tx.TimeStamp,
			store.TimeText(&rec.ClaimPrizeTx.Tx.DateTime),
			&rec.MainPrize.WinnerAid,
			&rec.MainPrize.WinnerAddr,
			&rec.MainPrize.EthAmount,
			&rec.MainPrize.EthAmountEth,
			&rec.RoundNum,
			&rec.MainPrize.NftTokenId,
			&nullSeed,
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
			return err
		}
		if nullSeed.Valid {
			rec.MainPrize.Seed = nullSeed.String
		} else {
			rec.MainPrize.Seed = "???"
		}
		return nil
	}
	return queryList(ctx, r, "prize claims by user", 32, query, scan, winnerAid)
}

// BidsByUser returns every bid of one bidder, newest first.
func (r *Repo) BidsByUser(ctx context.Context, bidderAid int64) ([]p.CGBidRec, error) {
	return bidList(ctx, r, "bids by user", "b.bidder_aid=$1", "b.id DESC", "", 32, bidderAid)
}

// UnclaimedDonatedNFTsByUser returns the donated NFTs a main-prize winner
// has not claimed yet, newest first.
func (r *Repo) UnclaimedDonatedNFTsByUser(ctx context.Context, winnerAid int64) ([]p.CGNFTDonation, error) {
	query := "SELECT " +
		"d.id," +
		"d.evtlog_id," +
		"d.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM d.time_stamp)::BIGINT," +
		"d.time_stamp," +
		"d.round_num," +
		"d.donor_aid," +
		"da.addr, " +
		"d.token_id, " +
		"d.idx," +
		"nft.address_id," +
		"nft.addr, " +
		"d.token_uri " +
		"FROM cg_nft_donation d " +
		"JOIN cg_prize_claim p ON p.round_num=d.round_num " +
		"LEFT JOIN cg_donated_nft_claimed c ON c.idx=d.idx " +
		"LEFT JOIN transaction t ON t.id=d.tx_id " +
		"LEFT JOIN address da ON d.donor_aid=da.address_id " +
		"LEFT JOIN address nft ON d.token_aid=nft.address_id " +
		"WHERE p.winner_aid=$1 AND p.round_num IS NOT NULL  AND c.idx IS NULL " +
		"ORDER BY d.evtlog_id DESC "
	scan := func(rows pgx.Rows, rec *p.CGNFTDonation) error {
		return rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.RoundNum,
			&rec.DonorAid,
			&rec.DonorAddr,
			&rec.NFTTokenId,
			&rec.Index,
			&rec.TokenAddressId,
			&rec.TokenAddr,
			&rec.NFTTokenURI,
		)
	}
	return queryList(ctx, r, "unclaimed donated nfts by user", 256, query, scan, winnerAid)
}

// RaffleNFTWinningsByUser returns the raffle NFT prizes won by one address,
// newest first.
func (r *Repo) RaffleNFTWinningsByUser(ctx context.Context, winnerAid int64) ([]p.CGRaffleNFTWinnerRec, error) {
	query := "SELECT " +
		"p.evtlog_id," +
		"p.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT," +
		"p.time_stamp," +
		"p.winner_aid," +
		"wa.addr," +
		"p.round_num, " +
		"p.token_id," +
		"p.winner_idx, " +
		"p.is_rwalk," +
		"p.is_staker " +
		"FROM cg_raffle_nft_prize p " +
		"LEFT JOIN transaction t ON t.id=p.tx_id " +
		"LEFT JOIN address wa ON p.winner_aid=wa.address_id " +
		"WHERE p.winner_aid=$1 " +
		"ORDER BY p.evtlog_id DESC "
	scan := func(rows pgx.Rows, rec *p.CGRaffleNFTWinnerRec) error {
		return rows.Scan(
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.RoundNum,
			&rec.TokenId,
			&rec.WinnerIndex,
			&rec.IsRWalk,
			&rec.IsStaker,
		)
	}
	return queryList(ctx, r, "raffle nft winnings by user", 256, query, scan, winnerAid)
}

// PrizeDepositsChronoWarriorByUser returns one winner's chrono warrior
// prizes (record type 2), newest first.
func (r *Repo) PrizeDepositsChronoWarriorByUser(ctx context.Context, winnerAid int64) ([]p.CGPrizeDepositRec, error) {
	query := "SELECT " +
		"p.id," +
		"p.evtlog_id," +
		"p.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT," +
		"p.time_stamp," +
		"p.winner_aid," +
		"wa.addr," +
		"p.round_num," +
		"p.eth_amount/1e18 amount_eth " +
		"FROM cg_chrono_warrior_prize p " +
		"LEFT JOIN transaction t ON t.id=p.tx_id " +
		"LEFT JOIN address wa ON p.winner_aid=wa.address_id " +
		"WHERE p.winner_aid = $1 " +
		"ORDER BY p.id DESC"
	scan := func(rows pgx.Rows, rec *p.CGPrizeDepositRec) error {
		err := rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.RoundNum,
			&rec.Amount,
		)
		if err != nil {
			return err
		}
		rec.RecordType = 2
		return nil
	}
	return queryList(ctx, r, "prize deposits chrono warrior by user", 32, query, scan, winnerAid)
}

// PrizeDepositsRaffleEthByUser returns one winner's raffle ETH prizes
// (record type 1), newest first.
func (r *Repo) PrizeDepositsRaffleEthByUser(ctx context.Context, winnerAid int64) ([]p.CGPrizeDepositRec, error) {
	query := "SELECT " +
		"p.id," +
		"p.evtlog_id," +
		"p.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT," +
		"p.time_stamp," +
		"p.winner_aid," +
		"wa.addr," +
		"p.round_num," +
		"p.amount/1e18 amount_eth " +
		"FROM cg_raffle_eth_prize p " +
		"LEFT JOIN transaction t ON t.id=p.tx_id " +
		"LEFT JOIN address wa ON p.winner_aid=wa.address_id " +
		"WHERE p.winner_aid = $1 " +
		"ORDER BY p.id DESC"
	scan := func(rows pgx.Rows, rec *p.CGPrizeDepositRec) error {
		err := rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.RoundNum,
			&rec.Amount,
		)
		if err != nil {
			return err
		}
		rec.RecordType = 1
		return nil
	}
	return queryList(ctx, r, "prize deposits raffle eth by user", 32, query, scan, winnerAid)
}

// DonatedNFTClaimsByUser returns the donated NFTs one winner has claimed,
// newest first.
func (r *Repo) DonatedNFTClaimsByUser(ctx context.Context, winnerAid int64) ([]p.CGDonatedNFTClaimRec, error) {
	query := "SELECT " +
		"c.id," +
		"c.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM c.time_stamp)::BIGINT," +
		"c.time_stamp," +
		"c.round_num," +
		"ta.addr," +
		"c.token_id, " +
		"c.idx, " +
		"c.winner_aid," +
		"wa.addr, " +
		"da.addr, " +
		"d.token_uri " +
		"FROM cg_donated_nft_claimed c " +
		"LEFT JOIN transaction t ON t.id=c.tx_id " +
		"LEFT JOIN address ta ON c.token_aid=ta.address_id " +
		"LEFT JOIN address wa ON c.winner_aid=wa.address_id " +
		"LEFT JOIN cg_nft_donation d ON d.idx=c.idx " +
		"LEFT JOIN address da ON d.donor_aid=da.address_id " +
		"WHERE c.winner_aid=$1 " +
		"ORDER BY c.id DESC "
	scan := func(rows pgx.Rows, rec *p.CGDonatedNFTClaimRec) error {
		return rows.Scan(
			&rec.RecordId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.RoundNum,
			&rec.TokenAddr,
			&rec.NFTTokenId,
			&rec.WinnerIndex,
			&rec.WinnerAid,
			&rec.WinnerAddr,
			&rec.DonorAddr,
			&rec.NFTTokenURI,
		)
	}
	return queryList(ctx, r, "donated nft claims by user", 256, query, scan, winnerAid)
}

// CosmicSignatureTokensByUser returns the Cosmic Signature NFTs currently
// owned by one address, newest first. limit 0 means no effective limit.
func (r *Repo) CosmicSignatureTokensByUser(ctx context.Context, userAid int64, offset, limit int) ([]p.CGCosmicSignatureMintRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := nftListSelectSQL + `
		WHERE m.cur_owner_aid=$1
		ORDER BY m.id DESC
		OFFSET $2 LIMIT $3`
	return queryList(ctx, r, "cosmic signature tokens by user", 64, query, scanNFTListRecord, userAid, offset, limit)
}

// CosmicTokenTransfersByUser returns the ERC-20 Cosmic Token transfers one
// address sent or received, newest first. limit 0 means no effective limit.
func (r *Repo) CosmicTokenTransfersByUser(ctx context.Context, userAid int64, offset, limit int) ([]p.CGERC20TransferRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := "SELECT " +
		"t.id," +
		"t.evtlog_id," +
		"t.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT," +
		"t.time_stamp," +
		"t.from_aid," +
		"fa.addr," +
		"t.to_aid," +
		"ta.addr," +
		"t.otype, " +
		"t.value," +
		"t.value/1e18 " +
		"FROM cg_erc20_transfer t " +
		"LEFT JOIN transaction tx ON tx.id=t.tx_id " +
		"LEFT JOIN address fa ON t.from_aid=fa.address_id " +
		"LEFT JOIN address ta ON t.to_aid=ta.address_id " +
		"WHERE (t.from_aid=$1) OR (t.to_aid=$1) " +
		"ORDER BY t.id DESC " +
		"OFFSET $2 LIMIT $3"
	scan := func(rows pgx.Rows, rec *p.CGERC20TransferRec) error {
		return rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.FromAid,
			&rec.FromAddr,
			&rec.ToAid,
			&rec.ToAddr,
			&rec.TransferType,
			&rec.Value,
			&rec.ValueFloat,
		)
	}
	return queryList(ctx, r, "cosmic token transfers by user", 256, query, scan, userAid, offset, limit)
}

// CosmicSignatureTransfersByUser returns the ERC-721 Cosmic Signature
// transfers one address sent or received, newest first. limit 0 means no
// effective limit.
func (r *Repo) CosmicSignatureTransfersByUser(ctx context.Context, userAid int64, offset, limit int) ([]p.CGTransfer, error) {
	if limit == 0 {
		limit = 1000000
	}
	query := "SELECT " +
		"t.id," +
		"t.evtlog_id," +
		"t.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM t.time_stamp)::BIGINT," +
		"t.time_stamp," +
		"t.from_aid," +
		"fa.addr," +
		"t.to_aid," +
		"ta.addr," +
		"t.otype, " +
		"t.token_id " +
		"FROM cg_erc721_transfer t " +
		"LEFT JOIN transaction tx ON tx.id=t.tx_id " +
		"LEFT JOIN address fa ON t.from_aid=fa.address_id " +
		"LEFT JOIN address ta ON t.to_aid=ta.address_id " +
		"WHERE (t.from_aid=$1) OR (t.to_aid=$1) " +
		"ORDER BY t.id DESC " +
		"OFFSET $2 LIMIT $3"
	scan := func(rows pgx.Rows, rec *p.CGTransfer) error {
		return rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.FromAid,
			&rec.FromAddr,
			&rec.ToAid,
			&rec.ToAddr,
			&rec.TransferType,
			&rec.TokenId,
		)
	}
	return queryList(ctx, r, "cosmic signature transfers by user", 256, query, scan, userAid, offset, limit)
}

// MarketingRewardHistoryByUser returns one page of a marketer's reward
// history, newest first.
func (r *Repo) MarketingRewardHistoryByUser(ctx context.Context, userAid int64, offset, limit int) ([]p.CGMarketingRewardRec, error) {
	query := "SELECT " + marketingRewardColumns + `
		WHERE r.marketer_aid = $1
		ORDER BY r.id DESC
		OFFSET $2 LIMIT $3`
	return queryList(ctx, r, "marketing reward history by user", 16, query, scanMarketingReward, userAid, offset, limit)
}

// StakedTokensCstByUser returns the Cosmic Signature tokens one address
// currently has staked, with mint provenance.
func (r *Repo) StakedTokensCstByUser(ctx context.Context, userAid int64) ([]p.CGStakedTokenCSTRec, error) {
	query := "SELECT " +
		"m.id," +
		"m.evtlog_id," +
		"m.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM m.time_stamp)::BIGINT," +
		"m.time_stamp," +
		"m.owner_aid," +
		"wa.addr," +
		"m.cur_owner_aid," +
		"oa.addr," +
		"m.seed, " +
		"m.token_id," +
		"m.round_num," +
		"p.round_num, " +
		"m.token_name, " +
		"EXTRACT(EPOCH FROM a.time_stamp)::BIGINT," +
		"a.time_Stamp," +
		"st.stake_action_id " +
		"FROM cg_staked_token_cst st " +
		"LEFT JOIN cg_mint_event m ON st.token_id=m.token_id " +
		"LEFT JOIN transaction t ON t.id=tx_id " +
		"LEFT JOIN address wa ON m.owner_aid=wa.address_id " +
		"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id " +
		"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id " +
		"LEFT JOIN cg_nft_staked_cst a ON a.action_id=st.stake_action_id " +
		"WHERE st.staker_aid=$1 " +
		"ORDER BY m.token_id"
	scan := func(rows pgx.Rows, rec *p.CGStakedTokenCSTRec) error {
		var nullPrizeNum sql.NullInt64
		err := rows.Scan(
			&rec.TokenInfo.RecordId,
			&rec.TokenInfo.Tx.EvtLogId,
			&rec.TokenInfo.Tx.BlockNum,
			&rec.TokenInfo.Tx.TxId,
			&rec.TokenInfo.Tx.TxHash,
			&rec.TokenInfo.Tx.TimeStamp,
			store.TimeText(&rec.TokenInfo.Tx.DateTime),
			&rec.TokenInfo.WinnerAid,
			&rec.TokenInfo.WinnerAddr,
			&rec.TokenInfo.CurOwnerAid,
			&rec.TokenInfo.CurOwnerAddr,
			&rec.TokenInfo.Seed,
			&rec.TokenInfo.TokenId,
			&rec.TokenInfo.RoundNum,
			&nullPrizeNum,
			&rec.TokenInfo.TokenName,
			&rec.StakeTimeStamp,
			store.TimeText(&rec.StakeDateTime),
			&rec.TokenInfo.StakeActionId,
		)
		if err != nil {
			return err
		}
		if nullPrizeNum.Valid {
			rec.TokenInfo.RecordType = 3
		} else {
			rec.TokenInfo.RecordType = 1
		}
		return nil
	}
	return queryList(ctx, r, "staked tokens cst by user", 16, query, scan, userAid)
}

// StakedTokensRwalkByUser returns the RandomWalk tokens one address
// currently has staked.
func (r *Repo) StakedTokensRwalkByUser(ctx context.Context, userAid int64) ([]p.CGStakedTokenRWalkRec, error) {
	query := "SELECT " +
		"a.action_id," +
		"EXTRACT(EPOCH FROM a.time_stamp)::BIGINT," +
		"a.time_Stamp," +
		"st.stake_action_id, " +
		"st.token_id " +
		"FROM cg_staked_token_rwalk st " +
		"LEFT JOIN cg_mint_event m ON st.token_id=m.token_id " +
		"LEFT JOIN transaction t ON t.id=tx_id " +
		"LEFT JOIN address wa ON m.owner_aid=wa.address_id " +
		"LEFT JOIN address oa ON m.cur_owner_aid=oa.address_id " +
		"LEFT JOIN cg_prize_claim p ON m.token_id=p.token_id " +
		"LEFT JOIN cg_nft_staked_rwalk a ON a.action_id=st.stake_action_id " +
		"WHERE st.staker_aid=$1 " +
		"ORDER BY m.token_id"
	scan := func(rows pgx.Rows, rec *p.CGStakedTokenRWalkRec) error {
		// The stake action id is scanned twice (a.action_id and
		// st.stake_action_id name the same action), mirroring the legacy
		// column list.
		return rows.Scan(
			&rec.StakeActionId,
			&rec.StakeTimeStamp,
			store.TimeText(&rec.StakeDateTime),
			&rec.StakeActionId,
			&rec.StakedTokenId,
		)
	}
	return queryList(ctx, r, "staked tokens rwalk by user", 16, query, scan, userAid)
}

// StakingRwalkMintsByUser returns the Cosmic Signature tokens minted to one
// RandomWalk-staker raffle winner, newest first.
func (r *Repo) StakingRwalkMintsByUser(ctx context.Context, userAid int64) ([]p.CGRaffleNFTWinnerRec, error) {
	query := "SELECT " +
		"w.id," +
		"w.evtlog_id," +
		"w.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM w.time_stamp)::BIGINT," +
		"w.time_stamp," +
		"w.token_id," +
		"w.winner_idx," +
		"w.round_num," +
		"w.winner_aid," +
		"wa.addr " +
		"FROM cg_raffle_nft_prize w " +
		"LEFT JOIN transaction t ON t.id=w.tx_id " +
		"LEFT JOIN address wa ON w.winner_aid=wa.address_id " +
		"WHERE is_rwalk=TRUE AND is_staker=TRUE AND w.winner_aid=$1 " +
		"ORDER BY w.evtlog_id DESC"
	scan := func(rows pgx.Rows, rec *p.CGRaffleNFTWinnerRec) error {
		if err := scanUserStakingMint(rows, rec); err != nil {
			return err
		}
		rec.IsRWalk = true
		rec.IsStaker = true
		return nil
	}
	return queryList(ctx, r, "staking rwalk mints by user", 16, query, scan, userAid)
}

// StakingCstMintsByUser returns the Cosmic Signature tokens minted to one
// CST-staker raffle winner, newest first.
func (r *Repo) StakingCstMintsByUser(ctx context.Context, userAid int64) ([]p.CGRaffleNFTWinnerRec, error) {
	query := "SELECT " +
		"w.id," +
		"w.evtlog_id," +
		"w.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM w.time_stamp)::BIGINT," +
		"w.time_stamp," +
		"w.token_id," +
		"w.winner_idx," +
		"w.round_num," +
		"w.winner_aid," +
		"wa.addr " +
		"FROM cg_raffle_nft_prize w " +
		"LEFT JOIN transaction t ON t.id=w.tx_id " +
		"LEFT JOIN address wa ON w.winner_aid=wa.address_id " +
		"WHERE is_rwalk=FALSE AND is_staker=TRUE AND w.winner_aid=$1 " +
		"ORDER BY w.evtlog_id DESC"
	scan := func(rows pgx.Rows, rec *p.CGRaffleNFTWinnerRec) error {
		if err := scanUserStakingMint(rows, rec); err != nil {
			return err
		}
		rec.IsRWalk = false
		rec.IsStaker = true
		return nil
	}
	return queryList(ctx, r, "staking cst mints by user", 16, query, scan, userAid)
}

// scanUserStakingMint scans the by-user staking-mint column list (which,
// unlike the global variant, has no cst_amount columns).
func scanUserStakingMint(rows pgx.Rows, rec *p.CGRaffleNFTWinnerRec) error {
	return rows.Scan(
		&rec.RecordId,
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.TokenId,
		&rec.WinnerIndex,
		&rec.RoundNum,
		&rec.WinnerAid,
		&rec.WinnerAddr,
	)
}

// StakingActionsCstByUser returns one page of a user's CST stake/unstake
// actions (newest first); NumStakedNFTs carries the user's running staked
// count recomputed over the returned page.
func (r *Repo) StakingActionsCstByUser(ctx context.Context, userAid int64, offset, limit int) ([]p.CGStakeActionCSTRec, error) {
	query := "(" +
		"SELECT " +
		"0 AS action_type," +
		"s.id," +
		"s.evtlog_id," +
		"s.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT," +
		"s.time_stamp," +
		"-1 AS usts," +
		"TO_TIMESTAMP(0) AS unstake_time," +
		"s.action_id," +
		"s.token_id," +
		"s.num_staked_nfts, " +
		"s.claimed " +
		"FROM cg_nft_staked_cst s " +
		"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
		"WHERE (s.staker_aid=$1) " +
		"OFFSET $2 LIMIT $3 " +
		") UNION ALL (" +
		"SELECT " +
		"1 AS action_type," +
		"u.id," +
		"u.evtlog_id," +
		"u.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT AS usts," +
		"u.time_stamp," +
		"0 AS usts," +
		"TO_TIMESTAMP(0) AS unstake_time," +
		"u.action_id," +
		"u.token_id," +
		"u.num_staked_nfts, " +
		"'F' AS claimed " +
		"FROM cg_nft_unstaked_cst u " +
		"LEFT JOIN transaction tx ON tx.id=u.tx_id " +
		"LEFT JOIN cg_nft_staked_cst s ON u.action_id=s.action_id " +
		"WHERE (u.staker_aid=$1) " +
		"OFFSET $2 LIMIT $3 " +
		") ORDER BY evtlog_id DESC"
	scan := func(rows pgx.Rows, rec *p.CGStakeActionCSTRec) error {
		return rows.Scan(
			&rec.ActionType,
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.UnstakeTimeStamp,
			store.TimeText(&rec.UnstakeDate),
			&rec.ActionId,
			&rec.TokenId,
			&rec.NumStakedNFTs,
			&rec.Claimed,
		)
	}
	records, err := queryList(ctx, r, "staking actions cst by user", 16, query, scan, userAid, offset, limit)
	if err != nil {
		return nil, err
	}
	var accumNumTokens int64
	for i := len(records) - 1; i >= 0; i-- {
		if records[i].ActionType == 0 {
			accumNumTokens++
		} else {
			accumNumTokens--
		}
		records[i].NumStakedNFTs = accumNumTokens
	}
	return records, nil
}

// StakingActionsRwalkByUser is StakingActionsCstByUser for RandomWalk stake
// actions (which carry no claimed flag).
func (r *Repo) StakingActionsRwalkByUser(ctx context.Context, userAid int64, offset, limit int) ([]p.CGStakeActionRWalkRec, error) {
	query := "(" +
		"SELECT " +
		"0 AS action_type," +
		"s.id," +
		"s.evtlog_id," +
		"s.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM s.time_stamp)::BIGINT," +
		"s.time_stamp," +
		"-1 AS usts," +
		"TO_TIMESTAMP(0) AS unstake_time," +
		"s.action_id," +
		"s.token_id," +
		"s.num_staked_nfts " +
		"FROM cg_nft_staked_rwalk s " +
		"LEFT JOIN transaction tx ON tx.id=s.tx_id " +
		"WHERE (s.staker_aid=$1) " +
		"OFFSET $2 LIMIT $3 " +
		") UNION ALL (" +
		"SELECT " +
		"1 AS action_type," +
		"u.id," +
		"u.evtlog_id," +
		"u.block_num," +
		"tx.id," +
		"tx.tx_hash," +
		"EXTRACT(EPOCH FROM u.time_stamp)::BIGINT AS usts," +
		"u.time_stamp," +
		"0 AS usts," +
		"TO_TIMESTAMP(0) AS unnstake_time," +
		"u.action_id," +
		"u.token_id," +
		"u.num_staked_nfts " +
		"FROM cg_nft_unstaked_rwalk u " +
		"LEFT JOIN transaction tx ON tx.id=u.tx_id " +
		"LEFT JOIN cg_nft_staked_rwalk s ON u.action_id=s.action_id " +
		"WHERE (u.staker_aid=$1) " +
		"OFFSET $2 LIMIT $3 " +
		") ORDER BY evtlog_id DESC"
	scan := func(rows pgx.Rows, rec *p.CGStakeActionRWalkRec) error {
		return rows.Scan(
			&rec.ActionType,
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.UnstakeTimeStamp,
			store.TimeText(&rec.UnstakeDate),
			&rec.ActionId,
			&rec.TokenId,
			&rec.NumStakedNFTs,
		)
	}
	records, err := queryList(ctx, r, "staking actions rwalk by user", 16, query, scan, userAid, offset, limit)
	if err != nil {
		return nil, err
	}
	var accumNumTokens int64
	for i := len(records) - 1; i >= 0; i-- {
		if records[i].ActionType == 0 {
			accumNumTokens++
		} else {
			accumNumTokens--
		}
		records[i].NumStakedNFTs = accumNumTokens
	}
	return records, nil
}

// UserNotifRedBoxRewards returns the "red box" claimables of one winner:
// pending raffle and chrono-warrior ETH, unclaimed donated NFT count,
// unclaimed CST staking rewards and per-round donated ERC-20 balances.
// A user with no CST staking row keeps a nil DonatedERC20Tokens list,
// matching the legacy early return.
func (r *Repo) UserNotifRedBoxRewards(ctx context.Context, winnerAid int64) (p.CGClaimInfo, error) {
	const op = "user notif red box rewards"
	var output p.CGClaimInfo

	var nullRaffleWei sql.NullString
	var nullRaffleEth sql.NullFloat64
	query := "SELECT SUM(amount), SUM(amount)/1e18 FROM cg_prize_deposit " +
		"WHERE winner_aid = $1 AND winner_index < 4 AND claimed = false"
	err := r.pool().QueryRow(ctx, query, winnerAid).Scan(&nullRaffleWei, &nullRaffleEth)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return p.CGClaimInfo{}, store.WrapError(op+": raffle eth", err)
	}
	if nullRaffleEth.Valid {
		output.ETHRaffleToClaim = nullRaffleEth.Float64
	}
	if nullRaffleWei.Valid {
		output.ETHRaffleToClaimWei = nullRaffleWei.String
	}

	var nullNfts sql.NullInt64
	query = "SELECT unclaimed_nfts FROM cg_winner WHERE winner_aid = $1"
	err = r.pool().QueryRow(ctx, query, winnerAid).Scan(&nullNfts)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return p.CGClaimInfo{}, store.WrapError(op+": unclaimed nfts", err)
	}
	if nullNfts.Valid {
		output.NumDonatedNFTToClaim = nullNfts.Int64
	}

	var nullChronoWei sql.NullString
	var nullChronoEth sql.NullFloat64
	query = "SELECT SUM(amount), SUM(amount)/1e18 FROM cg_prize_deposit " +
		"WHERE winner_aid = $1 AND winner_index = 4 AND claimed = false"
	err = r.pool().QueryRow(ctx, query, winnerAid).Scan(&nullChronoWei, &nullChronoEth)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return p.CGClaimInfo{}, store.WrapError(op+": chrono warrior eth", err)
	}
	if nullChronoEth.Valid {
		output.ETHChronoWarriorToClaim = nullChronoEth.Float64
	}
	if nullChronoWei.Valid {
		output.ETHChronoWarriorToClaimWei = nullChronoWei.String
	}

	var nullStakingRewards sql.NullFloat64
	query = "SELECT unclaimed_reward/1e18 FROM cg_staker_cst WHERE staker_aid=$1"
	err = r.pool().QueryRow(ctx, query, winnerAid).Scan(&nullStakingRewards)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return output, nil
		}
		return p.CGClaimInfo{}, store.WrapError(op+": staking rewards", err)
	}
	if nullStakingRewards.Valid {
		output.UnclaimedStakingReward = nullStakingRewards.Float64
	}

	query = "SELECT " +
		"p.round_num," +
		"d.token_aid," +
		"ta.addr," +
		"d.total_amount, " +
		"d.total_amount/1e18 " +
		"FROM cg_prize_claim p " +
		"JOIN cg_erc20_donation_stats d ON d.round_num=p.round_num AND claimed='F' " +
		"LEFT JOIN address ta ON d.token_aid=ta.address_id " +
		"WHERE p.winner_aid=$1 "
	scan := func(rows pgx.Rows, rec *p.ERC20DonatedTokensInfo) error {
		return rows.Scan(
			&rec.RoundNum,
			&rec.TokenAid,
			&rec.TokenAddr,
			&rec.Amount,
			&rec.AmountEth,
		)
	}
	tokens, err := queryList(ctx, r, op+": donated erc20", 16, query, scan, winnerAid)
	if err != nil {
		return p.CGClaimInfo{}, err
	}
	output.DonatedERC20Tokens = tokens
	return output, nil
}

// ERC20DonatedPrizesByWinner returns, per round won by userAid and per
// donated ERC-20 token, the donated vs claimed amounts.
func (r *Repo) ERC20DonatedPrizesByWinner(ctx context.Context, userAid int64) ([]p.CGSummarizedERC20Donation, error) {
	query := "WITH claim AS (" +
		"SELECT SUM(amount) total,round_num,token_aid,winner_aid " +
		"FROM cg_donated_tok_claimed GROUP BY round_num,token_aid,winner_aid " +
		") " +
		"SELECT " +
		"p.id," +
		"p.evtlog_id," +
		"p.block_num," +
		"t.id," +
		"t.tx_hash," +
		"EXTRACT(EPOCH FROM p.time_stamp)::BIGINT," +
		"p.time_stamp," +
		"dt20.round_num," +
		"tokaddr.address_id," +
		"tokaddr.addr, " +
		"dt20.total_amount, " +
		"dt20.total_amount/1e18, " +
		"COALESCE(claim.total,0), " +
		"COALESCE(claim.total,0)/1e18, " +
		"dt20.total_amount-COALESCE(claim.total,0)," +
		"(dt20.total_amount-COALESCE(claim.total,0))/1e18," +
		"dt20.winner_aid," +
		"wa.addr, " +
		"dt20.claimed " +
		"FROM cg_erc20_donation_stats dt20 " +
		"INNER JOIN cg_prize_claim p ON p.round_num=dt20.round_num " +
		"LEFT JOIN transaction t ON t.id=p.tx_id " +
		"LEFT JOIN address tokaddr ON dt20.token_aid=tokaddr.address_id " +
		"LEFT JOIN claim ON (claim.token_aid=dt20.token_aid AND dt20.round_num=claim.round_num) " +
		"LEFT JOIN address wa ON wa.address_id = claim.winner_aid " +
		"WHERE p.winner_aid = $1 " +
		"ORDER BY dt20.token_aid"
	scan := func(rows pgx.Rows, rec *p.CGSummarizedERC20Donation) error {
		var nullWinnerAddr sql.NullString
		var nullWinnerAid sql.NullInt64
		err := rows.Scan(
			&rec.RecordId,
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.RoundNum,
			&rec.TokenAid,
			&rec.TokenAddr,
			&rec.AmountDonated,
			&rec.AmountDonatedEth,
			&rec.AmountClaimed,
			&rec.AmountClaimedEth,
			&rec.DonateClaimDiff,
			&rec.DonateClaimDiffEth,
			&nullWinnerAid,
			&nullWinnerAddr,
			&rec.Claimed,
		)
		if err != nil {
			return err
		}
		if nullWinnerAid.Valid {
			rec.WinnerAid = nullWinnerAid.Int64
		}
		if nullWinnerAddr.Valid {
			rec.WinnerAddr = nullWinnerAddr.String
		}
		return nil
	}
	return queryList(ctx, r, "erc20 donated prizes by winner", 256, query, scan, userAid)
}
