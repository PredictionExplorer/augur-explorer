package cosmicgame

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// UserEventPageCursor identifies the last immutable event returned by a
// newest-first user history page.
type UserEventPageCursor struct {
	EventLogID int64
}

func (c *UserEventPageCursor) valid() bool {
	return c == nil || c.EventLogID >= 1
}

// truncatePage applies the LIMIT n+1 keyset convention: the extra row only
// proves another page exists and is never returned.
func truncatePage[T any](records []T, limit int) ([]T, bool) {
	if len(records) > limit {
		return records[:limit], true
	}
	return records, false
}

// UserPrizePageCursor identifies the last prize returned by UserPrizesPage.
// The triple is unique per winner by cg_prize's primary key.
type UserPrizePageCursor struct {
	Round       int64
	PrizeType   int64
	WinnerIndex int64
}

// userPrizesWhere filters the cg_prize registry join to the prizes one
// winner actually won. The round-wide staking allocation (ptype 15) has no
// single winner and never matches.
const userPrizesWhere = `
		WHERE (
			(p.ptype IN (0,1,2) AND pc.winner_aid = $1) OR
			(p.ptype IN (3,4) AND lw.winner_aid = $1) OR
			(p.ptype IN (5,6) AND ew.winner_aid = $1) OR
			(p.ptype IN (7,8,9) AND cw.winner_aid = $1) OR
			(p.ptype = 10 AND rew.winner_aid = $1) OR
			(p.ptype IN (11,12) AND rnw_bidder.winner_aid = $1) OR
			(p.ptype IN (13,14) AND rnw_rwalk.winner_aid = $1)
		)`

// UserPrizesPage returns at most limit prizes won by userAid, newest round
// first with the stable (ptype, winner_index) order inside each round. A nil
// cursor starts at the newest round.
func (r *Repo) UserPrizesPage(
	ctx context.Context,
	userAid int64,
	after *UserPrizePageCursor,
	limit int,
) (records []cgmodel.CGPrizeHistory, hasMore bool, err error) {
	const op = "user prizes page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}

	query := allPrizesSelect + userPrizesWhere + `
		ORDER BY p.round_num DESC, p.ptype, p.winner_index
		LIMIT $2`
	args := []any{userAid, limit + 1}
	if after != nil {
		if after.Round < 0 || after.PrizeType < 0 || after.PrizeType > 15 || after.WinnerIndex < 0 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		query = allPrizesSelect + userPrizesWhere + `
			AND (p.round_num < $2 OR (p.round_num = $2 AND (p.ptype, p.winner_index) > ($3, $4)))
			ORDER BY p.round_num DESC, p.ptype, p.winner_index
			LIMIT $5`
		args = []any{userAid, after.Round, after.PrizeType, after.WinnerIndex, limit + 1}
	}

	records, err = queryList(ctx, r, op, limit+1, query, scanPrizeHistoryRow, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserDepositWithdrawalRecord is the PrizesWallet withdrawal that claimed a
// deposit, resolved through cg_prize_deposit.withdrawal_id.
type UserDepositWithdrawalRecord struct {
	EventLogID      int64
	TxHash          string
	DateTime        string
	BeneficiaryAddr string
}

// UserRaffleEthDepositRecord is the exact-wei PrizesWallet ledger projection
// used by the v2 user deposit resource.
type UserRaffleEthDepositRecord struct {
	Tx              cgmodel.Transaction
	RoundNum        int64
	WinnerIndex     int64
	WinnerAid       int64
	WinnerAddr      string
	EthAmountWei    string
	IsChronoWarrior bool
	Claimed         bool
	Withdrawal      *UserDepositWithdrawalRecord
}

const userRaffleEthDepositsSelect = `SELECT
			p.evtlog_id,
			p.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,
			p.time_stamp,
			p.winner_aid,
			wa.addr,
			p.winner_index,
			p.round_num,
			p.amount::TEXT,
			(cw.round_num IS NOT NULL),
			p.claimed,
			w.evtlog_id,
			wt.tx_hash,
			w.time_stamp,
			ba.addr
		FROM cg_prize_deposit p
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
			LEFT JOIN cg_chrono_warrior_prize cw
				ON (p.round_num = cw.round_num AND p.winner_index = cw.winner_index)
			LEFT JOIN cg_prize_withdrawal w ON w.evtlog_id = p.withdrawal_id
			LEFT JOIN transaction wt ON wt.id = w.tx_id
			LEFT JOIN address ba ON w.beneficiary_aid = ba.address_id`

func scanUserRaffleEthDeposit(rows pgx.Rows, rec *UserRaffleEthDepositRecord) error {
	var (
		withdrawalEvt  sql.NullInt64
		withdrawalHash sql.NullString
		withdrawalTime string
		beneficiary    sql.NullString
	)
	if err := rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.WinnerAid,
		&rec.WinnerAddr,
		&rec.WinnerIndex,
		&rec.RoundNum,
		&rec.EthAmountWei,
		&rec.IsChronoWarrior,
		&rec.Claimed,
		&withdrawalEvt,
		&withdrawalHash,
		store.NullTimeText(&withdrawalTime),
		&beneficiary,
	); err != nil {
		return err
	}
	if withdrawalEvt.Valid {
		rec.Withdrawal = &UserDepositWithdrawalRecord{
			EventLogID:      withdrawalEvt.Int64,
			TxHash:          withdrawalHash.String,
			DateTime:        withdrawalTime,
			BeneficiaryAddr: beneficiary.String,
		}
	}
	return nil
}

// UserRaffleEthDepositsPage returns at most limit PrizesWallet ETH deposits
// credited to userAid, newest first. claimed narrows the ledger to claimed
// (true) or unclaimed (false) deposits; nil returns both.
func (r *Repo) UserRaffleEthDepositsPage(
	ctx context.Context,
	userAid int64,
	claimed *bool,
	after *UserEventPageCursor,
	limit int,
) (records []UserRaffleEthDepositRecord, hasMore bool, err error) {
	const op = "user raffle eth deposits page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	filter := "WHERE p.winner_aid=$1"
	args := []any{userAid}
	if claimed != nil {
		filter += fmt.Sprintf(" AND p.claimed=$%d", len(args)+1)
		args = append(args, *claimed)
	}
	if after != nil {
		filter += fmt.Sprintf(" AND p.evtlog_id < $%d", len(args)+1)
		args = append(args, after.EventLogID)
	}
	query := fmt.Sprintf(`%s
		%s
		ORDER BY p.evtlog_id DESC
		LIMIT $%d`, userRaffleEthDepositsSelect, filter, len(args)+1)
	args = append(args, limit+1)

	records, err = queryList(ctx, r, op, limit+1, query, scanUserRaffleEthDeposit, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserRaffleNftWinRecord is the exact-amount raffle NFT win projection used
// by the v2 user resource. Unlike the round-scoped winner resource it spans
// all three pools, so both pool flags travel with each row.
type UserRaffleNftWinRecord struct {
	Tx           cgmodel.Transaction
	RoundNum     int64
	WinnerIndex  int64
	WinnerAid    int64
	WinnerAddr   string
	TokenID      int64
	CstAmountWei string
	IsRWalk      bool
	IsStaker     bool
}

func scanUserRaffleNftWin(rows pgx.Rows, rec *UserRaffleNftWinRecord) error {
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
		&rec.TokenID,
		&rec.WinnerIndex,
		&rec.CstAmountWei,
		&rec.IsRWalk,
		&rec.IsStaker,
	)
}

// UserRaffleNftWinsPage returns at most limit raffle NFT wins of userAid,
// newest first by immutable event-log ID.
func (r *Repo) UserRaffleNftWinsPage(
	ctx context.Context,
	userAid int64,
	after *UserEventPageCursor,
	limit int,
) (records []UserRaffleNftWinRecord, hasMore bool, err error) {
	const op = "user raffle nft wins page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
			p.evtlog_id,
			p.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM p.time_stamp)::BIGINT,
			p.time_stamp,
			p.winner_aid,
			wa.addr,
			p.round_num,
			p.token_id,
			p.winner_idx,
			p.cst_amount::TEXT,
			p.is_rwalk,
			p.is_staker
		FROM cg_raffle_nft_prize p
			LEFT JOIN transaction t ON t.id=p.tx_id
			LEFT JOIN address wa ON p.winner_aid=wa.address_id
		WHERE p.winner_aid=$1`
	args := []any{userAid}
	if after != nil {
		query += fmt.Sprintf(" AND p.evtlog_id < $%d", len(args)+1)
		args = append(args, after.EventLogID)
	}
	query += fmt.Sprintf(`
		ORDER BY p.evtlog_id DESC
		LIMIT $%d`, len(args)+1)
	args = append(args, limit+1)

	records, err = queryList(ctx, r, op, limit+1, query, scanUserRaffleNftWin, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// EthDonationsByUserPage returns at most limit plain and info-carrying ETH
// donations made by donorAid, newest first across all rounds. It reuses the
// v2 round donation projection; the round travels with each row.
func (r *Repo) EthDonationsByUserPage(
	ctx context.Context,
	donorAid int64,
	after *UserEventPageCursor,
	limit int,
) (records []RoundEthDonationRecord, hasMore bool, err error) {
	const op = "eth donations by user page"
	if donorAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	args := []any{donorAid, limit + 1}
	if after != nil {
		args = []any{donorAid, after.EventLogID, limit + 1}
	}
	records, err = queryList(
		ctx,
		r,
		op,
		limit+1,
		userEthDonationsPageSQL(after != nil),
		scanRoundEthDonation,
		args...,
	)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// userEthDonationsPageSQL mirrors roundEthDonationsPageSQL with the donor
// filter replacing the round filter: both event tables are bounded before
// the newest-first merge.
func userEthDonationsPageSQL(after bool) string {
	filter := "WHERE d.donor_aid = $1"
	limitPlaceholder := "$2"
	if after {
		filter += " AND d.evtlog_id < $2"
		limitPlaceholder = "$3"
	}
	return fmt.Sprintf(`SELECT
			donation_kind,
			evtlog_id,
			block_num,
			tx_id,
			tx_hash,
			ts,
			date_time,
			donor_addr,
			amount_wei,
			round_num,
			contract_record_id,
			data
		FROM (
			(SELECT
				'plain'::TEXT AS donation_kind,
				d.evtlog_id,
				d.block_num,
				t.id AS tx_id,
				t.tx_hash,
				EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS ts,
				d.time_stamp AS date_time,
				da.addr AS donor_addr,
				d.amount::TEXT AS amount_wei,
				d.round_num,
				NULL::BIGINT AS contract_record_id,
				NULL::TEXT AS data
			FROM cg_eth_donated d
				LEFT JOIN transaction t ON t.id=d.tx_id
				LEFT JOIN address da ON da.address_id=d.donor_aid
			%s
			ORDER BY d.evtlog_id DESC
			LIMIT %s)
			UNION ALL
			(SELECT
				'withInfo'::TEXT AS donation_kind,
				d.evtlog_id,
				d.block_num,
				t.id AS tx_id,
				t.tx_hash,
				EXTRACT(EPOCH FROM d.time_stamp)::BIGINT AS ts,
				d.time_stamp AS date_time,
				da.addr AS donor_addr,
				d.amount::TEXT AS amount_wei,
				d.round_num,
				d.record_id AS contract_record_id,
				dj.data
			FROM cg_eth_donated_wi d
				LEFT JOIN cg_donation_json dj ON dj.record_id=d.record_id
				LEFT JOIN transaction t ON t.id=d.tx_id
				LEFT JOIN address da ON da.address_id=d.donor_aid
			%s
			ORDER BY d.evtlog_id DESC
			LIMIT %s)
		) donations
		ORDER BY evtlog_id DESC
		LIMIT %s`, filter, limitPlaceholder, filter, limitPlaceholder, limitPlaceholder)
}

// ERC20DonationsByUserPage returns at most limit ERC-20 donation events made
// by donorAid, newest first across all rounds.
func (r *Repo) ERC20DonationsByUserPage(
	ctx context.Context,
	donorAid int64,
	after *UserEventPageCursor,
	limit int,
) (records []RoundERC20DonationRecord, hasMore bool, err error) {
	const op = "erc20 donations by user page"
	if donorAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := roundERC20DonationsSelect + `
		WHERE tok.donor_aid=$1
		ORDER BY tok.evtlog_id DESC
		LIMIT $2`
	args := []any{donorAid, limit + 1}
	if after != nil {
		query = roundERC20DonationsSelect + `
			WHERE tok.donor_aid=$1 AND tok.evtlog_id < $2
			ORDER BY tok.evtlog_id DESC
			LIMIT $3`
		args = []any{donorAid, after.EventLogID, limit + 1}
	}

	records, err = queryList(ctx, r, op, limit+1, query, scanRoundERC20Donation, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// NFTDonationsByUserPage returns at most limit NFT donation events made by
// donorAid, newest first across all rounds.
func (r *Repo) NFTDonationsByUserPage(
	ctx context.Context,
	donorAid int64,
	after *UserEventPageCursor,
	limit int,
) (records []RoundNFTDonationRecord, hasMore bool, err error) {
	const op = "nft donations by user page"
	if donorAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := roundNFTDonationsSelect + `
		WHERE d.donor_aid=$1
		ORDER BY d.evtlog_id DESC
		LIMIT $2`
	args := []any{donorAid, limit + 1}
	if after != nil {
		query = roundNFTDonationsSelect + `
			WHERE d.donor_aid=$1 AND d.evtlog_id < $2
			ORDER BY d.evtlog_id DESC
			LIMIT $3`
		args = []any{donorAid, after.EventLogID, limit + 1}
	}

	records, err = queryList(ctx, r, op, limit+1, query, scanRoundNFTDonation, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserDonatedNftClaimRecord is the claim event that took a donated NFT out
// of the prize wallet.
type UserDonatedNftClaimRecord struct {
	EventLogID  int64
	TxHash      string
	DateTime    string
	ClaimerAid  int64
	ClaimerAddr string
}

// UserDonatedNftRecord pairs one NFT donation event with its claim state.
// A row belongs to a wallet when the wallet won the donation's round
// (RoundWinnerAid) or claimed the NFT after the winner's timeout
// (Claim.ClaimerAid).
type UserDonatedNftRecord struct {
	Tx             cgmodel.Transaction
	RoundNum       int64
	DonorAddr      string
	TokenAddr      string
	TokenID        int64
	DonationIndex  int64
	TokenURI       string
	RoundWinnerAid int64
	Claimed        bool
	Claim          *UserDonatedNftClaimRecord
}

const userDonatedNftsSelect = `SELECT
			d.evtlog_id,
			d.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM d.time_stamp)::BIGINT,
			d.time_stamp,
			d.round_num,
			da.addr,
			token.addr,
			d.token_id,
			d.idx,
			d.token_uri,
			COALESCE(pc.winner_aid, 0),
			(c.idx IS NOT NULL),
			c.evtlog_id,
			ct.tx_hash,
			c.time_stamp,
			c.winner_aid,
			ca.addr
		FROM cg_nft_donation d
			LEFT JOIN transaction t ON t.id=d.tx_id
			LEFT JOIN address da ON da.address_id=d.donor_aid
			LEFT JOIN address token ON token.address_id=d.token_aid
			LEFT JOIN cg_prize_claim pc ON pc.round_num=d.round_num
			LEFT JOIN cg_donated_nft_claimed c ON c.idx=d.idx
			LEFT JOIN transaction ct ON ct.id=c.tx_id
			LEFT JOIN address ca ON ca.address_id=c.winner_aid`

func scanUserDonatedNft(rows pgx.Rows, rec *UserDonatedNftRecord) error {
	var (
		claimEvt    sql.NullInt64
		claimHash   sql.NullString
		claimTime   string
		claimerAid  sql.NullInt64
		claimerAddr sql.NullString
	)
	if err := rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.RoundNum,
		&rec.DonorAddr,
		&rec.TokenAddr,
		&rec.TokenID,
		&rec.DonationIndex,
		&rec.TokenURI,
		&rec.RoundWinnerAid,
		&rec.Claimed,
		&claimEvt,
		&claimHash,
		store.NullTimeText(&claimTime),
		&claimerAid,
		&claimerAddr,
	); err != nil {
		return err
	}
	if claimEvt.Valid {
		rec.Claim = &UserDonatedNftClaimRecord{
			EventLogID:  claimEvt.Int64,
			TxHash:      claimHash.String,
			DateTime:    claimTime,
			ClaimerAid:  claimerAid.Int64,
			ClaimerAddr: claimerAddr.String,
		}
	}
	return nil
}

// UserDonatedNftsPage returns at most limit donated NFTs that belong to
// userAid's claim surface: donations from rounds the user won plus donations
// the user claimed elsewhere. claimed narrows to claimed (true) or unclaimed
// (false) donations; nil returns both. Newest donation first.
func (r *Repo) UserDonatedNftsPage(
	ctx context.Context,
	userAid int64,
	claimed *bool,
	after *UserEventPageCursor,
	limit int,
) (records []UserDonatedNftRecord, hasMore bool, err error) {
	const op = "user donated nfts page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	filter := "WHERE (pc.winner_aid=$1 OR c.winner_aid=$1)"
	args := []any{userAid}
	if claimed != nil {
		if *claimed {
			filter += " AND c.idx IS NOT NULL"
		} else {
			filter += " AND c.idx IS NULL"
		}
	}
	if after != nil {
		filter += fmt.Sprintf(" AND d.evtlog_id < $%d", len(args)+1)
		args = append(args, after.EventLogID)
	}
	query := fmt.Sprintf(`%s
		%s
		ORDER BY d.evtlog_id DESC
		LIMIT $%d`, userDonatedNftsSelect, filter, len(args)+1)
	args = append(args, limit+1)

	records, err = queryList(ctx, r, op, limit+1, query, scanUserDonatedNft, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// UserDonatedErc20ClaimRecord is the most recent claim event of one
// round-token donated ERC-20 entitlement.
type UserDonatedErc20ClaimRecord struct {
	EventLogID      int64
	TxHash          string
	DateTime        string
	ClaimerAid      int64
	ClaimerAddr     string
	AmountBaseUnits string
}

// UserDonatedErc20Record summarizes one (round, token) donated ERC-20
// entitlement with exact base-unit totals. TokenAid is the keyset
// tie-breaker inside a round and never reaches the wire.
type UserDonatedErc20Record struct {
	RoundNum           int64
	TokenAid           int64
	TokenAddr          string
	DonatedBaseUnits   string
	ClaimedBaseUnits   string
	RemainingBaseUnits string
	LastClaim          *UserDonatedErc20ClaimRecord
}

// UserDonatedErc20PageCursor identifies the last summary returned by
// UserDonatedErc20Page. The pair is cg_erc20_donation_stats' primary key.
type UserDonatedErc20PageCursor struct {
	Round    int64
	TokenAid int64
}

// userDonatedErc20Select derives the exact totals from the two canonical
// sources: cg_erc20_donation_stats.total_amount is trigger-decremented on
// every claim (it holds the remaining amount), so the donated total is the
// remainder plus the sum of claim events.
const userDonatedErc20Select = `WITH claim_totals AS (
			SELECT round_num, token_aid, SUM(amount) AS total
			FROM cg_donated_tok_claimed
			GROUP BY round_num, token_aid
		), last_claims AS (
			SELECT DISTINCT ON (round_num, token_aid)
				round_num, token_aid, evtlog_id, tx_id, time_stamp, winner_aid, amount
			FROM cg_donated_tok_claimed
			ORDER BY round_num, token_aid, evtlog_id DESC
		)
		SELECT
			s.round_num,
			s.token_aid,
			ta.addr,
			(s.total_amount + COALESCE(ct.total, 0))::TEXT,
			COALESCE(ct.total, 0)::TEXT,
			s.total_amount::TEXT,
			lc.evtlog_id,
			lt.tx_hash,
			lc.time_stamp,
			lc.winner_aid,
			ca.addr,
			lc.amount::TEXT
		FROM cg_erc20_donation_stats s
			LEFT JOIN address ta ON ta.address_id=s.token_aid
			LEFT JOIN claim_totals ct
				ON (ct.round_num=s.round_num AND ct.token_aid=s.token_aid)
			LEFT JOIN last_claims lc
				ON (lc.round_num=s.round_num AND lc.token_aid=s.token_aid)
			LEFT JOIN transaction lt ON lt.id=lc.tx_id
			LEFT JOIN address ca ON ca.address_id=lc.winner_aid`

func scanUserDonatedErc20(rows pgx.Rows, rec *UserDonatedErc20Record) error {
	var (
		claimEvt    sql.NullInt64
		claimHash   sql.NullString
		claimTime   string
		claimerAid  sql.NullInt64
		claimerAddr sql.NullString
		claimAmount sql.NullString
	)
	if err := rows.Scan(
		&rec.RoundNum,
		&rec.TokenAid,
		&rec.TokenAddr,
		&rec.DonatedBaseUnits,
		&rec.ClaimedBaseUnits,
		&rec.RemainingBaseUnits,
		&claimEvt,
		&claimHash,
		store.NullTimeText(&claimTime),
		&claimerAid,
		&claimerAddr,
		&claimAmount,
	); err != nil {
		return err
	}
	if claimEvt.Valid {
		rec.LastClaim = &UserDonatedErc20ClaimRecord{
			EventLogID:      claimEvt.Int64,
			TxHash:          claimHash.String,
			DateTime:        claimTime,
			ClaimerAid:      claimerAid.Int64,
			ClaimerAddr:     claimerAddr.String,
			AmountBaseUnits: claimAmount.String,
		}
	}
	return nil
}

// UserDonatedErc20Page returns at most limit donated ERC-20 entitlement
// summaries that belong to userAid's claim surface: rounds the user won plus
// round-token pairs the user claimed elsewhere. Newest round first with the
// stable token order inside each round.
func (r *Repo) UserDonatedErc20Page(
	ctx context.Context,
	userAid int64,
	after *UserDonatedErc20PageCursor,
	limit int,
) (records []UserDonatedErc20Record, hasMore bool, err error) {
	const op = "user donated erc20 page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}

	filter := `WHERE (
			EXISTS (
				SELECT 1 FROM cg_prize_claim pc
				WHERE pc.round_num=s.round_num AND pc.winner_aid=$1
			) OR EXISTS (
				SELECT 1 FROM cg_donated_tok_claimed c
				WHERE c.round_num=s.round_num
					AND c.token_aid=s.token_aid
					AND c.winner_aid=$1
			)
		)`
	args := []any{userAid}
	if after != nil {
		if after.Round < 0 || after.TokenAid < 1 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		filter += fmt.Sprintf(
			" AND (s.round_num < $%d OR (s.round_num = $%d AND s.token_aid > $%d))",
			len(args)+1, len(args)+1, len(args)+2)
		args = append(args, after.Round, after.TokenAid)
	}
	query := fmt.Sprintf(`%s
		%s
		ORDER BY s.round_num DESC, s.token_aid
		LIMIT $%d`, userDonatedErc20Select, filter, len(args)+1)
	args = append(args, limit+1)

	records, err = queryList(ctx, r, op, limit+1, query, scanUserDonatedErc20, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}
