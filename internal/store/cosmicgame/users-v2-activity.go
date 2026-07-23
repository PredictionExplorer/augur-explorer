package cosmicgame

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// CosmicSignatureMintSource identifies the prize event that minted one
// Cosmic Signature token. Every mint comes from exactly one source; a token
// matching zero or several sources is a data inconsistency the scan rejects.
type CosmicSignatureMintSource string

// The six prize families that mint Cosmic Signature tokens. Raffle mints
// split by pool: bidder raffles, RandomWalk-staker raffles and Cosmic
// Signature staker raffles.
const (
	MintSourceMainPrize          CosmicSignatureMintSource = "mainPrize"
	MintSourceBidderRaffle       CosmicSignatureMintSource = "bidderRaffle"
	MintSourceRandomWalkStaker   CosmicSignatureMintSource = "randomWalkStakerRaffle"
	MintSourceCosmicSigStaker    CosmicSignatureMintSource = "cosmicSignatureStakerRaffle"
	MintSourceEnduranceChampion  CosmicSignatureMintSource = "enduranceChampion"
	MintSourceLastCstBidder      CosmicSignatureMintSource = "lastCstBidder"
	MintSourceChronoWarriorPrize CosmicSignatureMintSource = "chronoWarrior"
)

// UserTokenPageCursor identifies the last token returned by the ascending
// owned-token directory.
type UserTokenPageCursor struct {
	TokenID int64
}

func (c *UserTokenPageCursor) valid() bool {
	return c == nil || c.TokenID >= 0
}

// UserOwnedTokenRecord is one Cosmic Signature token the indexed ERC-721
// ledger currently attributes to one wallet, with its mint provenance and
// live staking-wallet membership.
type UserOwnedTokenRecord struct {
	MintTx     cgmodel.Transaction
	OwnerAid   int64
	TokenID    int64
	MintRound  int64
	Seed       string
	TokenName  string
	WinnerAddr string
	MintSource CosmicSignatureMintSource
	Staked     bool
}

// UserCosmicSignatureTokensPage returns at most limit Cosmic Signature
// tokens currently owned by userAid in ascending token order. The
// collection is live membership: rows move between wallets on transfers.
func (r *Repo) UserCosmicSignatureTokensPage(
	ctx context.Context,
	userAid int64,
	after *UserTokenPageCursor,
	limit int,
) (records []UserOwnedTokenRecord, hasMore bool, err error) {
	const op = "user cosmic signature tokens page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
			m.evtlog_id,
			m.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM m.time_stamp)::BIGINT,
			m.time_stamp,
			m.cur_owner_aid,
			m.token_id,
			m.round_num,
			m.seed,
			COALESCE(m.token_name, ''),
			wa.addr,
			(st.token_id IS NOT NULL),
			(pc.token_id IS NOT NULL),
			rnw.is_rwalk,
			rnw.is_staker,
			(endu.erc721_token_id IS NOT NULL),
			(stel.erc721_token_id IS NOT NULL),
			(cw.nft_id IS NOT NULL)
		FROM cg_mint_event m
			LEFT JOIN transaction t ON t.id=m.tx_id
			LEFT JOIN address wa ON wa.address_id=m.owner_aid
			LEFT JOIN cg_staked_token_cst st ON st.token_id=m.token_id
			LEFT JOIN cg_prize_claim pc
				ON m.token_id>=pc.token_id AND m.token_id<pc.token_id+pc.num_cs_nfts
			LEFT JOIN cg_raffle_nft_prize rnw
				ON (rnw.token_id=m.token_id AND rnw.round_num=m.round_num)
			LEFT JOIN cg_endurance_prize endu
				ON (endu.erc721_token_id=m.token_id AND endu.round_num=m.round_num)
			LEFT JOIN cg_lastcst_prize stel
				ON (stel.erc721_token_id=m.token_id AND stel.round_num=m.round_num)
			LEFT JOIN cg_chrono_warrior_prize cw
				ON (cw.nft_id=m.token_id AND cw.round_num=m.round_num)
		WHERE m.cur_owner_aid=$1`
	args := []any{userAid}
	if after != nil {
		query += fmt.Sprintf(" AND m.token_id > $%d", len(args)+1)
		args = append(args, after.TokenID)
	}
	query += fmt.Sprintf(`
		ORDER BY m.token_id
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	records, err = queryList(ctx, r, op, limit+1, query, scanUserOwnedToken, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// scanUserOwnedToken derives the mint source from the per-prize-family
// joins, rejecting rows that match zero or several sources.
func scanUserOwnedToken(rows pgx.Rows, rec *UserOwnedTokenRecord) error {
	var (
		isMainPrize   bool
		raffleIsRWalk sql.NullBool
		raffleStaker  sql.NullBool
		isEndurance   bool
		isLastCst     bool
		isChrono      bool
	)
	if err := rows.Scan(
		&rec.MintTx.EvtLogId,
		&rec.MintTx.BlockNum,
		&rec.MintTx.TxId,
		&rec.MintTx.TxHash,
		&rec.MintTx.TimeStamp,
		store.TimeText(&rec.MintTx.DateTime),
		&rec.OwnerAid,
		&rec.TokenID,
		&rec.MintRound,
		&rec.Seed,
		&rec.TokenName,
		&rec.WinnerAddr,
		&rec.Staked,
		&isMainPrize,
		&raffleIsRWalk,
		&raffleStaker,
		&isEndurance,
		&isLastCst,
		&isChrono,
	); err != nil {
		return err
	}

	source, err := deriveMintSource(
		rec.TokenID, isMainPrize, raffleIsRWalk, raffleStaker, isEndurance, isLastCst, isChrono)
	if err != nil {
		return err
	}
	rec.MintSource = source
	return nil
}

// deriveMintSource resolves the per-prize-family join flags into the one
// prize source that minted the token, rejecting tokens that match zero or
// several sources.
func deriveMintSource(
	tokenID int64,
	isMainPrize bool,
	raffleIsRWalk, raffleStaker sql.NullBool,
	isEndurance, isLastCst, isChrono bool,
) (CosmicSignatureMintSource, error) {
	var sources []CosmicSignatureMintSource
	if isMainPrize {
		sources = append(sources, MintSourceMainPrize)
	}
	if raffleStaker.Valid {
		switch {
		case !raffleStaker.Bool:
			sources = append(sources, MintSourceBidderRaffle)
		case raffleIsRWalk.Valid && raffleIsRWalk.Bool:
			sources = append(sources, MintSourceRandomWalkStaker)
		default:
			sources = append(sources, MintSourceCosmicSigStaker)
		}
	}
	if isEndurance {
		sources = append(sources, MintSourceEnduranceChampion)
	}
	if isLastCst {
		sources = append(sources, MintSourceLastCstBidder)
	}
	if isChrono {
		sources = append(sources, MintSourceChronoWarriorPrize)
	}
	if len(sources) != 1 {
		return "", fmt.Errorf("token %d matches %d mint sources", tokenID, len(sources))
	}
	return sources[0], nil
}

// UserTransferDirection relates one transfer to the requested wallet.
type UserTransferDirection string

// The three wallet-relative transfer directions computed by the ledger
// queries; self covers transfers where both sides are the wallet.
const (
	UserTransferIn   UserTransferDirection = "in"
	UserTransferOut  UserTransferDirection = "out"
	UserTransferSelf UserTransferDirection = "self"
)

// UserCosmicSignatureTransferRecord is one ERC-721 Cosmic Signature
// transfer one wallet sent or received. TransferType is the raw otype
// column (0 transfer, 1 mint, 2 burn).
type UserCosmicSignatureTransferRecord struct {
	Tx           cgmodel.Transaction
	TokenID      int64
	FromAid      int64
	FromAddr     string
	ToAid        int64
	ToAddr       string
	TransferType int64
	Direction    UserTransferDirection
}

// UserCosmicTokenTransferRecord is one ERC-20 Cosmic Token transfer one
// wallet sent or received, with the exact base-unit amount.
type UserCosmicTokenTransferRecord struct {
	Tx           cgmodel.Transaction
	AmountWei    string
	FromAid      int64
	FromAddr     string
	ToAid        int64
	ToAddr       string
	TransferType int64
	Direction    UserTransferDirection
}

// userTransfersPageSQL merges the wallet's sent and received transfers
// newest first. Both sides are bounded index scans before the merge; UNION
// (not UNION ALL) collapses self-transfers that qualify on both sides into
// one row. valueColumn is a compile-time literal from this file.
func userTransfersPageSQL(table, valueColumn string, after bool) string {
	branch := func(side string) string {
		filter := fmt.Sprintf("WHERE t.%s = $1", side)
		limitPlaceholder := "$2"
		if after {
			filter += " AND t.evtlog_id < $2"
			limitPlaceholder = "$3"
		}
		return fmt.Sprintf(`(SELECT
				t.evtlog_id,
				t.block_num,
				tx.id AS tx_id,
				tx.tx_hash,
				EXTRACT(EPOCH FROM t.time_stamp)::BIGINT AS ts,
				t.time_stamp AS date_time,
				t.%s AS value,
				t.from_aid,
				fa.addr AS from_addr,
				t.to_aid,
				ta.addr AS to_addr,
				t.otype,
				CASE
					WHEN t.from_aid = $1 AND t.to_aid = $1 THEN 'self'
					WHEN t.from_aid = $1 THEN 'out'
					ELSE 'in'
				END AS direction
			FROM %s t
				LEFT JOIN transaction tx ON tx.id=t.tx_id
				LEFT JOIN address fa ON fa.address_id=t.from_aid
				LEFT JOIN address ta ON ta.address_id=t.to_aid
			%s
			ORDER BY t.evtlog_id DESC
			LIMIT %s)`, valueColumn, table, filter, limitPlaceholder)
	}
	limitPlaceholder := "$2"
	if after {
		limitPlaceholder = "$3"
	}
	return fmt.Sprintf(`SELECT
			evtlog_id,
			block_num,
			tx_id,
			tx_hash,
			ts,
			date_time,
			value,
			from_aid,
			from_addr,
			to_aid,
			to_addr,
			otype,
			direction
		FROM (
			%s
			UNION
			%s
		) transfers
		ORDER BY evtlog_id DESC
		LIMIT %s`, branch("from_aid"), branch("to_aid"), limitPlaceholder)
}

// UserCosmicSignatureTransfersPage returns at most limit ERC-721 Cosmic
// Signature transfers userAid sent or received, newest first by immutable
// event-log ID.
func (r *Repo) UserCosmicSignatureTransfersPage(
	ctx context.Context,
	userAid int64,
	after *UserEventPageCursor,
	limit int,
) (records []UserCosmicSignatureTransferRecord, hasMore bool, err error) {
	const op = "user cosmic signature transfers page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	args := []any{userAid, limit + 1}
	if after != nil {
		args = []any{userAid, after.EventLogID, limit + 1}
	}
	scan := func(rows pgx.Rows, rec *UserCosmicSignatureTransferRecord) error {
		return rows.Scan(
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.TokenID,
			&rec.FromAid,
			&rec.FromAddr,
			&rec.ToAid,
			&rec.ToAddr,
			&rec.TransferType,
			&rec.Direction,
		)
	}
	records, err = queryList(
		ctx, r, op, limit+1,
		userTransfersPageSQL("cg_erc721_transfer", "token_id", after != nil),
		scan, args...,
	)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserCosmicTokenTransfersPage returns at most limit ERC-20 Cosmic Token
// transfers userAid sent or received, newest first by immutable event-log
// ID, with exact base-unit amounts.
func (r *Repo) UserCosmicTokenTransfersPage(
	ctx context.Context,
	userAid int64,
	after *UserEventPageCursor,
	limit int,
) (records []UserCosmicTokenTransferRecord, hasMore bool, err error) {
	const op = "user cosmic token transfers page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	args := []any{userAid, limit + 1}
	if after != nil {
		args = []any{userAid, after.EventLogID, limit + 1}
	}
	scan := func(rows pgx.Rows, rec *UserCosmicTokenTransferRecord) error {
		return rows.Scan(
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.AmountWei,
			&rec.FromAid,
			&rec.FromAddr,
			&rec.ToAid,
			&rec.ToAddr,
			&rec.TransferType,
			&rec.Direction,
		)
	}
	records, err = queryList(
		ctx, r, op, limit+1,
		userTransfersPageSQL("cg_erc20_transfer", "value::TEXT", after != nil),
		scan, args...,
	)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserMarketingRewardRecord is one MarketingWallet reward paid to one
// wallet, with the exact Cosmic Token base-unit amount.
type UserMarketingRewardRecord struct {
	Tx          cgmodel.Transaction
	MarketerAid int64
	AmountWei   string
}

// UserMarketingRewardsPage returns at most limit marketing rewards paid to
// userAid, newest first by immutable event-log ID.
func (r *Repo) UserMarketingRewardsPage(
	ctx context.Context,
	userAid int64,
	after *UserEventPageCursor,
	limit int,
) (records []UserMarketingRewardRecord, hasMore bool, err error) {
	const op = "user marketing rewards page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
			m.evtlog_id,
			m.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM m.time_stamp)::BIGINT,
			m.time_stamp,
			m.marketer_aid,
			m.amount::TEXT
		FROM cg_mkt_reward m
			LEFT JOIN transaction t ON t.id=m.tx_id
		WHERE m.marketer_aid=$1`
	args := []any{userAid}
	if after != nil {
		query += fmt.Sprintf(" AND m.evtlog_id < $%d", len(args)+1)
		args = append(args, after.EventLogID)
	}
	query += fmt.Sprintf(`
		ORDER BY m.evtlog_id DESC
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	scan := func(rows pgx.Rows, rec *UserMarketingRewardRecord) error {
		return rows.Scan(
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.MarketerAid,
			&rec.AmountWei,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserCosmicTokenSummaryRecord is one wallet's exact Cosmic Token position:
// the indexed balance, base-unit earnings per source, the amount consumed
// in CST bids, the signed net game flow and the transfer activity counts.
type UserCosmicTokenSummaryRecord struct {
	BalanceWei                 string
	BiddingRewardsWei          string
	MainPrizesWei              string
	RafflePrizesWei            string
	ChronoWarriorPrizesWei     string
	EnduranceChampionPrizesWei string
	LastCstBidderPrizesWei     string
	MarketingRewardsWei        string
	TotalEarnedWei             string
	ConsumedInBidsWei          string
	NetWei                     string
	TransferCount              int64
	MintCount                  int64
	BurnCount                  int64
}

// UserCosmicTokenSummaryV2 computes the wallet's exact Cosmic Token
// position in one query, so every field comes from the same database
// snapshot. It replaces the v1 float summary; sums skip the -1 "not a CST
// bid" sentinels. A wallet without indexed activity gets the zero shape.
func (r *Repo) UserCosmicTokenSummaryV2(
	ctx context.Context,
	userAid int64,
) (UserCosmicTokenSummaryRecord, error) {
	const op = "user cosmic token summary v2"
	if userAid < 1 {
		return UserCosmicTokenSummaryRecord{}, fmt.Errorf("%s: invalid address id", op)
	}

	query := `WITH balance AS (
			SELECT COALESCE(SUM(cur_balance), 0) AS wei
			FROM cg_costok_owner WHERE owner_aid=$1
		), bidding AS (
			SELECT
				COALESCE(SUM(cst_reward) FILTER (WHERE cst_reward > 0), 0) AS earned,
				COALESCE(SUM(cst_price) FILTER (WHERE cst_price > 0), 0) AS consumed
			FROM cg_bid WHERE bidder_aid=$1
		), main_prizes AS (
			SELECT COALESCE(SUM(cst_amount) FILTER (WHERE cst_amount > 0), 0) AS earned
			FROM cg_prize_claim WHERE winner_aid=$1
		), raffles AS (
			SELECT COALESCE(SUM(cst_amount) FILTER (WHERE cst_amount > 0), 0) AS earned
			FROM cg_raffle_nft_prize WHERE winner_aid=$1
		), chrono AS (
			SELECT COALESCE(SUM(cst_amount) FILTER (WHERE cst_amount > 0), 0) AS earned
			FROM cg_chrono_warrior_prize WHERE winner_aid=$1
		), endurance AS (
			SELECT COALESCE(SUM(erc20_amount) FILTER (WHERE erc20_amount > 0), 0) AS earned
			FROM cg_endurance_prize WHERE winner_aid=$1
		), lastcst AS (
			SELECT COALESCE(SUM(erc20_amount) FILTER (WHERE erc20_amount > 0), 0) AS earned
			FROM cg_lastcst_prize WHERE winner_aid=$1
		), marketing AS (
			SELECT COALESCE(SUM(amount), 0) AS earned
			FROM cg_mkt_reward WHERE marketer_aid=$1
		), activity AS (
			SELECT
				COUNT(*)::BIGINT AS transfers,
				COUNT(*) FILTER (WHERE otype = 1)::BIGINT AS mints,
				COUNT(*) FILTER (WHERE otype = 2)::BIGINT AS burns
			FROM cg_erc20_transfer WHERE from_aid=$1 OR to_aid=$1
		)
		SELECT
			b.wei::TEXT,
			bd.earned::TEXT,
			mp.earned::TEXT,
			rf.earned::TEXT,
			ch.earned::TEXT,
			en.earned::TEXT,
			lc.earned::TEXT,
			mk.earned::TEXT,
			(bd.earned + mp.earned + rf.earned + ch.earned + en.earned + lc.earned + mk.earned)::TEXT,
			bd.consumed::TEXT,
			(bd.earned + mp.earned + rf.earned + ch.earned + en.earned + lc.earned + mk.earned
				- bd.consumed)::TEXT,
			a.transfers,
			a.mints,
			a.burns
		FROM balance b
			CROSS JOIN bidding bd
			CROSS JOIN main_prizes mp
			CROSS JOIN raffles rf
			CROSS JOIN chrono ch
			CROSS JOIN endurance en
			CROSS JOIN lastcst lc
			CROSS JOIN marketing mk
			CROSS JOIN activity a`

	var record UserCosmicTokenSummaryRecord
	err := r.q(ctx).QueryRow(ctx, query, userAid).Scan(
		&record.BalanceWei,
		&record.BiddingRewardsWei,
		&record.MainPrizesWei,
		&record.RafflePrizesWei,
		&record.ChronoWarriorPrizesWei,
		&record.EnduranceChampionPrizesWei,
		&record.LastCstBidderPrizesWei,
		&record.MarketingRewardsWei,
		&record.TotalEarnedWei,
		&record.ConsumedInBidsWei,
		&record.NetWei,
		&record.TransferCount,
		&record.MintCount,
		&record.BurnCount,
	)
	if err != nil {
		return UserCosmicTokenSummaryRecord{}, store.WrapError(op, err)
	}
	return record, nil
}

// UserPendingWinningsRecord is one wallet's claimable balances with exact
// wei amounts.
type UserPendingWinningsRecord struct {
	RaffleEthWei           string
	ChronoWarriorEthWei    string
	DonatedNftCount        int64
	StakingRewardWei       string
	DonatedErc20TokenCount int64
}

// UserPendingWinnings computes the wallet's unclaimed winnings in one
// query: PrizesWallet ETH split into raffle and chrono-warrior shares by
// the same chrono join the deposit ledger uses, the donated-NFT count, the
// unclaimed Cosmic Signature staking reward and the number of donated
// ERC-20 entitlements with a remaining balance in rounds the wallet won.
func (r *Repo) UserPendingWinnings(
	ctx context.Context,
	userAid int64,
) (UserPendingWinningsRecord, error) {
	const op = "user pending winnings"
	if userAid < 1 {
		return UserPendingWinningsRecord{}, fmt.Errorf("%s: invalid address id", op)
	}

	query := `WITH deposits AS (
			SELECT
				COALESCE(SUM(p.amount) FILTER (WHERE cw.round_num IS NULL), 0) AS raffle_wei,
				COALESCE(SUM(p.amount) FILTER (WHERE cw.round_num IS NOT NULL), 0) AS chrono_wei
			FROM cg_prize_deposit p
				LEFT JOIN cg_chrono_warrior_prize cw
					ON (cw.round_num = p.round_num AND cw.winner_index = p.winner_index)
			WHERE p.winner_aid=$1 AND NOT p.claimed
		), nfts AS (
			SELECT COALESCE(SUM(unclaimed_nfts), 0)::BIGINT AS pending
			FROM cg_winner WHERE winner_aid=$1
		), staking AS (
			SELECT COALESCE(SUM(unclaimed_reward), 0) AS wei
			FROM cg_staker_cst WHERE staker_aid=$1
		), erc20 AS (
			SELECT COUNT(*)::BIGINT AS tokens
			FROM cg_erc20_donation_stats s
			WHERE s.total_amount > 0 AND EXISTS (
				SELECT 1 FROM cg_prize_claim pc
				WHERE pc.round_num = s.round_num AND pc.winner_aid=$1
			)
		)
		SELECT
			d.raffle_wei::TEXT,
			d.chrono_wei::TEXT,
			n.pending,
			s.wei::TEXT,
			e.tokens
		FROM deposits d
			CROSS JOIN nfts n
			CROSS JOIN staking s
			CROSS JOIN erc20 e`

	var record UserPendingWinningsRecord
	err := r.q(ctx).QueryRow(ctx, query, userAid).Scan(
		&record.RaffleEthWei,
		&record.ChronoWarriorEthWei,
		&record.DonatedNftCount,
		&record.StakingRewardWei,
		&record.DonatedErc20TokenCount,
	)
	if err != nil {
		return UserPendingWinningsRecord{}, store.WrapError(op, err)
	}
	return record, nil
}
