package cosmicgame

import (
	"context"
	"fmt"
	"math/big"

	"github.com/jackc/pgx/v5"
)

// ParticipantKind scopes a participant-directory cursor to one endpoint.
type ParticipantKind string

const (
	ParticipantBidders            ParticipantKind = "bidders"
	ParticipantWinners            ParticipantKind = "winners"
	ParticipantDonors             ParticipantKind = "donors"
	ParticipantCSTStakers         ParticipantKind = "cstStakers"
	ParticipantRandomWalkStakers  ParticipantKind = "randomWalkStakers"
	ParticipantDualStakers        ParticipantKind = "dualStakers"
	ParticipantCsTokenHolders     ParticipantKind = "cosmicSignatureHolders"
	ParticipantCosmicTokenHolders ParticipantKind = "cosmicTokenHolders"
)

// ParticipantPageCursor identifies the final row of a participant page.
// SortValue is an exact, canonical, non-negative integer.
type ParticipantPageCursor struct {
	Kind      ParticipantKind
	SortValue string
	AddressID int64
}

type BidderParticipantRecord struct {
	BidderAid int64
	Address   string
	BidCount  int64
	MaxBidWei string
}

type WinnerParticipantRecord struct {
	WinnerAid          int64
	Address            string
	PrizeCount         int64
	MaxMainPrizeETHWei string
	TotalETHWonWei     string
	CSTPrizeCount      int64
	NFTPrizeCount      int64
	UnclaimedNFTCount  int64
	TotalETHSpentWei   string
}

type DonorParticipantRecord struct {
	DonorAid        int64
	Address         string
	DonationCount   int64
	TotalDonatedWei string
}

type CSTStakerParticipantRecord struct {
	StakerAid          int64
	Address            string
	StakedTokenCount   int64
	StakeActionCount   int64
	UnstakeActionCount int64
	TotalRewardWei     string
	UnclaimedRewardWei string
}

type RandomWalkStakerParticipantRecord struct {
	StakerAid          int64
	Address            string
	StakedTokenCount   int64
	StakeActionCount   int64
	UnstakeActionCount int64
	MintedTokenCount   int64
}

type DualStakerParticipantRecord struct {
	StakerAid                    int64
	Address                      string
	TotalStakedTokenCount        int64
	CSTStakedTokenCount          int64
	CSTStakeActionCount          int64
	CSTUnstakeActionCount        int64
	CSTTotalRewardWei            string
	CSTUnclaimedRewardWei        string
	RandomWalkStakedTokenCount   int64
	RandomWalkStakeActionCount   int64
	RandomWalkUnstakeActionCount int64
	RandomWalkMintedTokenCount   int64
}

func validateParticipantPage(
	op string,
	kind ParticipantKind,
	after *ParticipantPageCursor,
	limit int,
) error {
	if limit <= 0 || limit > maxStatisticsPageLimit {
		return fmt.Errorf("%s: invalid limit", op)
	}
	if after == nil {
		return nil
	}
	if after.Kind != kind ||
		after.AddressID < 1 ||
		!validParticipantSortValue(kind, after.SortValue) {
		return fmt.Errorf("%s: invalid cursor", op)
	}
	return nil
}

func canonicalParticipantSortValue(value string) bool {
	parsed, ok := new(big.Int).SetString(value, 10)
	return ok && parsed.Sign() >= 0 && parsed.String() == value
}

func validParticipantSortValue(kind ParticipantKind, value string) bool {
	parsed, ok := new(big.Int).SetString(value, 10)
	if !ok || parsed.Sign() < 0 || parsed.String() != value {
		return false
	}
	switch kind {
	case ParticipantBidders, ParticipantWinners,
		ParticipantRandomWalkStakers, ParticipantDualStakers,
		ParticipantCsTokenHolders:
		return parsed.IsInt64()
	case ParticipantDonors, ParticipantCSTStakers,
		ParticipantCosmicTokenHolders:
		return true
	default:
		return false
	}
}

func participantInt64SortValue(after *ParticipantPageCursor) int64 {
	if after == nil {
		return 0
	}
	value, _ := new(big.Int).SetString(after.SortValue, 10)
	return value.Int64()
}

func trimParticipantPage[T any](records []T, limit int) ([]T, bool) {
	if len(records) <= limit {
		return records, false
	}
	return records[:limit], true
}

// BidderParticipantsPage returns bidders ordered by bid count, then address ID.
func (r *Repo) BidderParticipantsPage(
	ctx context.Context,
	after *ParticipantPageCursor,
	limit int,
) ([]BidderParticipantRecord, bool, error) {
	const op = "bidder participants page"
	if err := validateParticipantPage(op, ParticipantBidders, after, limit); err != nil {
		return nil, false, err
	}
	query := `SELECT b.bidder_aid,a.addr,b.num_bids,COALESCE(b.max_bid,0)::TEXT
		FROM cg_bidder b
		LEFT JOIN address a ON a.address_id=b.bidder_aid
		WHERE b.num_bids > 0
		ORDER BY b.num_bids DESC,b.bidder_aid ASC
		LIMIT $1`
	args := []any{limit + 1}
	if after != nil {
		query = `SELECT b.bidder_aid,a.addr,b.num_bids,COALESCE(b.max_bid,0)::TEXT
			FROM cg_bidder b
			LEFT JOIN address a ON a.address_id=b.bidder_aid
			WHERE b.num_bids > 0
				AND (b.num_bids < $1
					OR (b.num_bids = $1 AND b.bidder_aid > $2))
			ORDER BY b.num_bids DESC,b.bidder_aid ASC
			LIMIT $3`
		args = []any{participantInt64SortValue(after), after.AddressID, limit + 1}
	}
	scan := func(rows pgx.Rows, rec *BidderParticipantRecord) error {
		return rows.Scan(&rec.BidderAid, &rec.Address, &rec.BidCount, &rec.MaxBidWei)
	}
	records, err := queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore := trimParticipantPage(records, limit)
	return records, hasMore, nil
}

const canonicalWinnerPrizesCTE = `WITH prize_winners AS (
		SELECT p.ptype,
			COALESCE(pc.winner_aid,rew.winner_aid,rnw.winner_aid,
				ew.winner_aid,lw.winner_aid,cw.winner_aid) AS winner_aid
		FROM cg_prize p
		LEFT JOIN cg_prize_claim pc
			ON p.round_num=pc.round_num AND p.ptype IN (0,1,2)
		LEFT JOIN cg_raffle_eth_prize rew
			ON p.round_num=rew.round_num AND p.winner_index=rew.winner_idx AND p.ptype=10
		LEFT JOIN cg_raffle_nft_prize rnw
			ON p.round_num=rnw.round_num AND p.winner_index=rnw.winner_idx
				AND p.ptype IN (11,12,13,14)
				AND ((p.ptype IN (11,12) AND rnw.is_rwalk=false)
					OR (p.ptype IN (13,14) AND rnw.is_rwalk=true))
		LEFT JOIN cg_endurance_prize ew
			ON p.round_num=ew.round_num AND p.ptype IN (5,6)
		LEFT JOIN cg_lastcst_prize lw
			ON p.round_num=lw.round_num AND p.ptype IN (3,4)
		LEFT JOIN cg_chrono_warrior_prize cw
			ON p.round_num=cw.round_num AND p.winner_index=cw.winner_index
				AND p.ptype IN (7,8,9)
		WHERE p.ptype != 15
	), winner_prizes AS (
		SELECT winner_aid,
			COUNT(*)::BIGINT AS prize_count,
			COUNT(*) FILTER (WHERE ptype IN (1,4,6,8,11,13))::BIGINT AS cst_prize_count,
			COUNT(*) FILTER (WHERE ptype IN (2,3,5,9,12,14))::BIGINT AS nft_prize_count
		FROM prize_winners
		WHERE winner_aid IS NOT NULL
		GROUP BY winner_aid
	), main_winnings AS (
		SELECT winner_aid,MAX(COALESCE(amount,0)) AS max_main_prize
		FROM cg_prize_claim
		GROUP BY winner_aid
	), eth_winnings AS (
		SELECT winner_aid,SUM(amount) AS total_eth_won
		FROM (
			SELECT winner_aid,COALESCE(amount,0) AS amount FROM cg_prize_claim
			UNION ALL
			SELECT winner_aid,amount FROM cg_raffle_eth_prize
			UNION ALL
			SELECT winner_aid,eth_amount FROM cg_chrono_warrior_prize
		) amounts
		GROUP BY winner_aid
	), unclaimed_nfts AS (
		SELECT pc.winner_aid,COUNT(*)::BIGINT AS unclaimed_nft_count
		FROM cg_prize_claim pc
		JOIN cg_nft_donation d ON d.round_num=pc.round_num
		LEFT JOIN cg_donated_nft_claimed c
			ON c.round_num=d.round_num AND c.idx=d.idx
		WHERE c.id IS NULL
		GROUP BY pc.winner_aid
	)`

const winnerParticipantsBase = canonicalWinnerPrizesCTE + `
	SELECT wp.winner_aid,a.addr,wp.prize_count,
		COALESCE(m.max_main_prize,0)::TEXT,
		COALESCE(e.total_eth_won,0)::TEXT,
		wp.cst_prize_count,wp.nft_prize_count,
		COALESCE(u.unclaimed_nft_count,0)::BIGINT,
		COALESCE(b.total_eth_spent,0)::TEXT
	FROM winner_prizes wp
	LEFT JOIN address a ON a.address_id=wp.winner_aid
	LEFT JOIN main_winnings m ON m.winner_aid=wp.winner_aid
	LEFT JOIN eth_winnings e ON e.winner_aid=wp.winner_aid
	LEFT JOIN unclaimed_nfts u ON u.winner_aid=wp.winner_aid
	LEFT JOIN cg_bidder b ON b.bidder_aid=wp.winner_aid`

// WinnerParticipantsPage returns winners ordered by canonical prize-registry
// count, then address ID. It deliberately does not trust the replay-sensitive
// cg_winner counters.
func (r *Repo) WinnerParticipantsPage(
	ctx context.Context,
	after *ParticipantPageCursor,
	limit int,
) ([]WinnerParticipantRecord, bool, error) {
	const op = "winner participants page"
	if err := validateParticipantPage(op, ParticipantWinners, after, limit); err != nil {
		return nil, false, err
	}
	query := winnerParticipantsBase + `
		ORDER BY wp.prize_count DESC,wp.winner_aid ASC
		LIMIT $1`
	args := []any{limit + 1}
	if after != nil {
		query = winnerParticipantsBase + `
			WHERE wp.prize_count < $1
				OR (wp.prize_count = $1 AND wp.winner_aid > $2)
			ORDER BY wp.prize_count DESC,wp.winner_aid ASC
			LIMIT $3`
		args = []any{participantInt64SortValue(after), after.AddressID, limit + 1}
	}
	scan := func(rows pgx.Rows, rec *WinnerParticipantRecord) error {
		return rows.Scan(
			&rec.WinnerAid,
			&rec.Address,
			&rec.PrizeCount,
			&rec.MaxMainPrizeETHWei,
			&rec.TotalETHWonWei,
			&rec.CSTPrizeCount,
			&rec.NFTPrizeCount,
			&rec.UnclaimedNFTCount,
			&rec.TotalETHSpentWei,
		)
	}
	records, err := queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore := trimParticipantPage(records, limit)
	return records, hasMore, nil
}

// DonorParticipantsPage returns ETH donors ordered by exact donated amount.
func (r *Repo) DonorParticipantsPage(
	ctx context.Context,
	after *ParticipantPageCursor,
	limit int,
) ([]DonorParticipantRecord, bool, error) {
	const op = "donor participants page"
	if err := validateParticipantPage(op, ParticipantDonors, after, limit); err != nil {
		return nil, false, err
	}
	query := `SELECT d.donor_aid,a.addr,d.count_donations,
			COALESCE(d.total_eth_donated,0)::TEXT
		FROM cg_donor d
		LEFT JOIN address a ON a.address_id=d.donor_aid
		WHERE d.count_donations > 0
		ORDER BY COALESCE(d.total_eth_donated,0) DESC,d.donor_aid ASC
		LIMIT $1`
	args := []any{limit + 1}
	if after != nil {
		query = `SELECT d.donor_aid,a.addr,d.count_donations,
				COALESCE(d.total_eth_donated,0)::TEXT
			FROM cg_donor d
			LEFT JOIN address a ON a.address_id=d.donor_aid
			WHERE d.count_donations > 0
				AND (COALESCE(d.total_eth_donated,0) < $1::NUMERIC
					OR (COALESCE(d.total_eth_donated,0) = $1::NUMERIC
						AND d.donor_aid > $2))
			ORDER BY COALESCE(d.total_eth_donated,0) DESC,d.donor_aid ASC
			LIMIT $3`
		args = []any{after.SortValue, after.AddressID, limit + 1}
	}
	scan := func(rows pgx.Rows, rec *DonorParticipantRecord) error {
		return rows.Scan(&rec.DonorAid, &rec.Address, &rec.DonationCount, &rec.TotalDonatedWei)
	}
	records, err := queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore := trimParticipantPage(records, limit)
	return records, hasMore, nil
}

// CSTStakerParticipantsPage returns CST stakers ordered by exact earned reward.
func (r *Repo) CSTStakerParticipantsPage(
	ctx context.Context,
	after *ParticipantPageCursor,
	limit int,
) ([]CSTStakerParticipantRecord, bool, error) {
	const op = "CST staker participants page"
	if err := validateParticipantPage(op, ParticipantCSTStakers, after, limit); err != nil {
		return nil, false, err
	}
	query := `SELECT s.staker_aid,a.addr,COALESCE(s.total_tokens_staked,0)::BIGINT,
			s.num_stake_actions,COALESCE(s.num_unstake_actions,0)::BIGINT,
			COALESCE(s.total_reward,0)::TEXT,COALESCE(s.unclaimed_reward,0)::TEXT
		FROM cg_staker_cst s
		LEFT JOIN address a ON a.address_id=s.staker_aid
		WHERE s.num_stake_actions > 0
		ORDER BY COALESCE(s.total_reward,0) DESC,s.staker_aid ASC
		LIMIT $1`
	args := []any{limit + 1}
	if after != nil {
		query = `SELECT s.staker_aid,a.addr,COALESCE(s.total_tokens_staked,0)::BIGINT,
				s.num_stake_actions,COALESCE(s.num_unstake_actions,0)::BIGINT,
				COALESCE(s.total_reward,0)::TEXT,COALESCE(s.unclaimed_reward,0)::TEXT
			FROM cg_staker_cst s
			LEFT JOIN address a ON a.address_id=s.staker_aid
			WHERE s.num_stake_actions > 0
				AND (COALESCE(s.total_reward,0) < $1::NUMERIC
					OR (COALESCE(s.total_reward,0) = $1::NUMERIC
						AND s.staker_aid > $2))
			ORDER BY COALESCE(s.total_reward,0) DESC,s.staker_aid ASC
			LIMIT $3`
		args = []any{after.SortValue, after.AddressID, limit + 1}
	}
	scan := func(rows pgx.Rows, rec *CSTStakerParticipantRecord) error {
		return rows.Scan(
			&rec.StakerAid,
			&rec.Address,
			&rec.StakedTokenCount,
			&rec.StakeActionCount,
			&rec.UnstakeActionCount,
			&rec.TotalRewardWei,
			&rec.UnclaimedRewardWei,
		)
	}
	records, err := queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore := trimParticipantPage(records, limit)
	return records, hasMore, nil
}

// RandomWalkStakerParticipantsPage returns RandomWalk stakers ordered by
// current staked-token count.
func (r *Repo) RandomWalkStakerParticipantsPage(
	ctx context.Context,
	after *ParticipantPageCursor,
	limit int,
) ([]RandomWalkStakerParticipantRecord, bool, error) {
	const op = "RandomWalk staker participants page"
	if err := validateParticipantPage(op, ParticipantRandomWalkStakers, after, limit); err != nil {
		return nil, false, err
	}
	query := `SELECT s.staker_aid,a.addr,COALESCE(s.total_tokens_staked,0)::BIGINT,
			s.num_stake_actions,COALESCE(s.num_unstake_actions,0)::BIGINT,
			COALESCE(s.num_tokens_minted,0)::BIGINT
		FROM cg_staker_rwalk s
		LEFT JOIN address a ON a.address_id=s.staker_aid
		WHERE s.num_stake_actions > 0
		ORDER BY COALESCE(s.total_tokens_staked,0) DESC,s.staker_aid ASC
		LIMIT $1`
	args := []any{limit + 1}
	if after != nil {
		query = `SELECT s.staker_aid,a.addr,COALESCE(s.total_tokens_staked,0)::BIGINT,
				s.num_stake_actions,COALESCE(s.num_unstake_actions,0)::BIGINT,
				COALESCE(s.num_tokens_minted,0)::BIGINT
			FROM cg_staker_rwalk s
			LEFT JOIN address a ON a.address_id=s.staker_aid
			WHERE s.num_stake_actions > 0
				AND (COALESCE(s.total_tokens_staked,0) < $1
					OR (COALESCE(s.total_tokens_staked,0) = $1
						AND s.staker_aid > $2))
			ORDER BY COALESCE(s.total_tokens_staked,0) DESC,s.staker_aid ASC
			LIMIT $3`
		args = []any{participantInt64SortValue(after), after.AddressID, limit + 1}
	}
	scan := func(rows pgx.Rows, rec *RandomWalkStakerParticipantRecord) error {
		return rows.Scan(
			&rec.StakerAid,
			&rec.Address,
			&rec.StakedTokenCount,
			&rec.StakeActionCount,
			&rec.UnstakeActionCount,
			&rec.MintedTokenCount,
		)
	}
	records, err := queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore := trimParticipantPage(records, limit)
	return records, hasMore, nil
}

// DualStakerParticipantsPage returns current dual stakers ordered by the
// combined number of staked tokens.
func (r *Repo) DualStakerParticipantsPage(
	ctx context.Context,
	after *ParticipantPageCursor,
	limit int,
) ([]DualStakerParticipantRecord, bool, error) {
	const op = "dual staker participants page"
	if err := validateParticipantPage(op, ParticipantDualStakers, after, limit); err != nil {
		return nil, false, err
	}
	const selectSQL = `SELECT c.staker_aid,a.addr,
			(COALESCE(c.total_tokens_staked,0)::NUMERIC+
				COALESCE(r.total_tokens_staked,0)::NUMERIC)::BIGINT,
			COALESCE(c.total_tokens_staked,0)::BIGINT,
			COALESCE(c.num_stake_actions,0)::BIGINT,
			COALESCE(c.num_unstake_actions,0)::BIGINT,
			COALESCE(c.total_reward,0)::TEXT,COALESCE(c.unclaimed_reward,0)::TEXT,
			COALESCE(r.total_tokens_staked,0)::BIGINT,
			COALESCE(r.num_stake_actions,0)::BIGINT,
			COALESCE(r.num_unstake_actions,0)::BIGINT,
			COALESCE(r.num_tokens_minted,0)::BIGINT
		FROM cg_staker_cst c
		JOIN cg_staker_rwalk r ON r.staker_aid=c.staker_aid
		LEFT JOIN address a ON a.address_id=c.staker_aid`
	query := selectSQL + `
		WHERE COALESCE(c.total_tokens_staked,0) > 0
			AND COALESCE(r.total_tokens_staked,0) > 0
		ORDER BY (COALESCE(c.total_tokens_staked,0)::NUMERIC+
			COALESCE(r.total_tokens_staked,0)::NUMERIC) DESC,c.staker_aid ASC
		LIMIT $1`
	args := []any{limit + 1}
	if after != nil {
		query = selectSQL + `
			WHERE COALESCE(c.total_tokens_staked,0) > 0
				AND COALESCE(r.total_tokens_staked,0) > 0
				AND ((COALESCE(c.total_tokens_staked,0)::NUMERIC+
						COALESCE(r.total_tokens_staked,0)::NUMERIC) < $1::NUMERIC
					OR ((COALESCE(c.total_tokens_staked,0)::NUMERIC+
							COALESCE(r.total_tokens_staked,0)::NUMERIC) = $1::NUMERIC
						AND c.staker_aid > $2))
			ORDER BY (COALESCE(c.total_tokens_staked,0)::NUMERIC+
				COALESCE(r.total_tokens_staked,0)::NUMERIC) DESC,c.staker_aid ASC
			LIMIT $3`
		args = []any{participantInt64SortValue(after), after.AddressID, limit + 1}
	}
	scan := func(rows pgx.Rows, rec *DualStakerParticipantRecord) error {
		return rows.Scan(
			&rec.StakerAid,
			&rec.Address,
			&rec.TotalStakedTokenCount,
			&rec.CSTStakedTokenCount,
			&rec.CSTStakeActionCount,
			&rec.CSTUnstakeActionCount,
			&rec.CSTTotalRewardWei,
			&rec.CSTUnclaimedRewardWei,
			&rec.RandomWalkStakedTokenCount,
			&rec.RandomWalkStakeActionCount,
			&rec.RandomWalkUnstakeActionCount,
			&rec.RandomWalkMintedTokenCount,
		)
	}
	records, err := queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore := trimParticipantPage(records, limit)
	return records, hasMore, nil
}
