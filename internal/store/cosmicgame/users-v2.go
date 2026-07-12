package cosmicgame

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// UserProfileRecord is the exact, bounded projection used by the v2 user
// resource. Amounts retain their database integer precision as decimal text.
type UserProfileRecord struct {
	Address string

	BidCount         int64
	MaxETHBidWei     *string
	TotalETHSpentWei string
	TotalCSTSpentWei string

	PrizeCount         int64
	MaxMainPrizeETHWei string
	TotalETHWonWei     string
	CSTPrizeCount      int64
	NFTPrizeCount      int64
	UnclaimedNFTCount  int64

	RaffleETHPrizeCount int64
	RaffleETHTotalWei   string
	RaffleNFTPrizeCount int64
	RaffleCSTTotalWei   string

	ETHDonationCount int64
	ETHDonatedWei    string

	CosmicTokenTransferCount     int64
	CosmicSignatureTransferCount int64

	CSTStakedTokenCount   int64
	CSTStakeActionCount   int64
	CSTUnstakeActionCount int64
	CSTTotalRewardWei     string
	CSTUnclaimedRewardWei string

	RandomWalkStakedTokenCount   int64
	RandomWalkStakeActionCount   int64
	RandomWalkUnstakeActionCount int64
	RandomWalkMintedTokenCount   int64
}

// UserAddressID resolves an indexed wallet to its internal address ID.
func (r *Repo) UserAddressID(ctx context.Context, address string) (int64, error) {
	return r.store.LookupAddressID(ctx, address)
}

// UserProfile returns one exact, collection-free activity profile, or
// store.ErrNotFound when userAid has no address row.
func (r *Repo) UserProfile(ctx context.Context, userAid int64) (UserProfileRecord, error) {
	const op = "user profile"
	if userAid < 1 {
		return UserProfileRecord{}, fmt.Errorf("%s: invalid address id", op)
	}

	query := canonicalWinnerPrizesCTE + `,
	bid_stats AS (
		SELECT COUNT(*)::BIGINT AS bid_count,
			MAX(CASE WHEN eth_price >= 0 THEN eth_price END)::TEXT AS max_eth_bid,
			COALESCE(SUM(CASE WHEN eth_price > 0 THEN eth_price ELSE 0 END),0)::TEXT
				AS total_eth_spent,
			COALESCE(SUM(CASE WHEN cst_price > 0 THEN cst_price ELSE 0 END),0)::TEXT
				AS total_cst_spent
		FROM cg_bid
		WHERE bidder_aid=$1
	), raffle_eth AS (
		SELECT COUNT(*)::BIGINT AS prize_count,
			COALESCE(SUM(amount),0)::TEXT AS total_won
		FROM cg_raffle_eth_prize
		WHERE winner_aid=$1
	), raffle_nft AS (
		SELECT COUNT(*)::BIGINT AS prize_count,
			COALESCE(SUM(cst_amount),0)::TEXT AS total_cst
		FROM cg_raffle_nft_prize
		WHERE winner_aid=$1
	), transfer_stats AS (
		SELECT COALESCE(SUM(erc20_num_transfers),0)::BIGINT AS erc20_count,
			COALESCE(SUM(erc721_num_transfers),0)::BIGINT AS erc721_count
		FROM cg_transfer_stats
		WHERE user_aid=$1
	)
	SELECT a.addr,
		bs.bid_count,bs.max_eth_bid,bs.total_eth_spent,bs.total_cst_spent,
		COALESCE(wp.prize_count,0)::BIGINT,
		COALESCE(m.max_main_prize,0)::TEXT,
		COALESCE(e.total_eth_won,0)::TEXT,
		COALESCE(wp.cst_prize_count,0)::BIGINT,
		COALESCE(wp.nft_prize_count,0)::BIGINT,
		COALESCE(u.unclaimed_nft_count,0)::BIGINT,
		re.prize_count,re.total_won,rn.prize_count,rn.total_cst,
		COALESCE(d.count_donations,0)::BIGINT,
		COALESCE(d.total_eth_donated,0)::TEXT,
		ts.erc20_count,ts.erc721_count,
		COALESCE(sc.total_tokens_staked,0)::BIGINT,
		COALESCE(sc.num_stake_actions,0)::BIGINT,
		COALESCE(sc.num_unstake_actions,0)::BIGINT,
		COALESCE(sc.total_reward,0)::TEXT,
		COALESCE(sc.unclaimed_reward,0)::TEXT,
		COALESCE(sr.total_tokens_staked,0)::BIGINT,
		COALESCE(sr.num_stake_actions,0)::BIGINT,
		COALESCE(sr.num_unstake_actions,0)::BIGINT,
		COALESCE(sr.num_tokens_minted,0)::BIGINT
	FROM address a
	CROSS JOIN bid_stats bs
	CROSS JOIN raffle_eth re
	CROSS JOIN raffle_nft rn
	CROSS JOIN transfer_stats ts
	LEFT JOIN winner_prizes wp ON wp.winner_aid=a.address_id
	LEFT JOIN main_winnings m ON m.winner_aid=a.address_id
	LEFT JOIN eth_winnings e ON e.winner_aid=a.address_id
	LEFT JOIN unclaimed_nfts u ON u.winner_aid=a.address_id
	LEFT JOIN cg_donor d ON d.donor_aid=a.address_id
	LEFT JOIN cg_staker_cst sc ON sc.staker_aid=a.address_id
	LEFT JOIN cg_staker_rwalk sr ON sr.staker_aid=a.address_id
	WHERE a.address_id=$1`

	var record UserProfileRecord
	var maxETHBid sql.NullString
	err := r.pool().QueryRow(ctx, query, userAid).Scan(
		&record.Address,
		&record.BidCount,
		&maxETHBid,
		&record.TotalETHSpentWei,
		&record.TotalCSTSpentWei,
		&record.PrizeCount,
		&record.MaxMainPrizeETHWei,
		&record.TotalETHWonWei,
		&record.CSTPrizeCount,
		&record.NFTPrizeCount,
		&record.UnclaimedNFTCount,
		&record.RaffleETHPrizeCount,
		&record.RaffleETHTotalWei,
		&record.RaffleNFTPrizeCount,
		&record.RaffleCSTTotalWei,
		&record.ETHDonationCount,
		&record.ETHDonatedWei,
		&record.CosmicTokenTransferCount,
		&record.CosmicSignatureTransferCount,
		&record.CSTStakedTokenCount,
		&record.CSTStakeActionCount,
		&record.CSTUnstakeActionCount,
		&record.CSTTotalRewardWei,
		&record.CSTUnclaimedRewardWei,
		&record.RandomWalkStakedTokenCount,
		&record.RandomWalkStakeActionCount,
		&record.RandomWalkUnstakeActionCount,
		&record.RandomWalkMintedTokenCount,
	)
	if err != nil {
		return UserProfileRecord{}, store.WrapError(op, err)
	}
	if maxETHBid.Valid {
		record.MaxETHBidWei = &maxETHBid.String
	}
	return record, nil
}
