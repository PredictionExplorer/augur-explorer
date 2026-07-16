package cosmicgame

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// GlobalTokenFilter narrows the global Cosmic Signature directory. At most
// one of NamedOnly and NameContains may be set; NameContains matches the
// current token name case-insensitively with ILIKE wildcards escaped, so
// the term is always literal.
type GlobalTokenFilter struct {
	NamedOnly    bool
	NameContains string
}

func (f GlobalTokenFilter) valid() bool {
	return !f.NamedOnly || f.NameContains == ""
}

// GlobalTokenPageCursor identifies the last token returned by the
// descending global token directory.
type GlobalTokenPageCursor struct {
	TokenID int64
}

func (c *GlobalTokenPageCursor) valid() bool {
	return c == nil || c.TokenID >= 0
}

// GlobalTokenRecord is one minted Cosmic Signature token in the global
// directory: mint provenance plus the live current owner, name and
// staking-wallet membership.
type GlobalTokenRecord struct {
	MintTx       cgmodel.Transaction
	TokenID      int64
	MintRound    int64
	Seed         string
	TokenName    string
	WinnerAddr   string
	CurOwnerAid  int64
	CurOwnerAddr string
	MintSource   CosmicSignatureMintSource
	Staked       bool
}

// escapeLikePattern makes term literal inside an ILIKE pattern by escaping
// the backslash escape character itself and both wildcards.
func escapeLikePattern(term string) string {
	replacer := strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`)
	return replacer.Replace(term)
}

// globalTokenProjectionSQL resolves each token's provenance through scalar
// subqueries rather than one wide join: the planner prices five correlated
// subplans far below a ten-relation join search, and a prize table holding
// several rows for one (token, round) fails loudly ("more than one row
// returned by a subquery") instead of silently duplicating directory rows.
// #nosec G101 -- a SQL projection, not credentials.
const globalTokenProjectionSQL = `
		m.evtlog_id,
		m.block_num,
		t.id,
		t.tx_hash,
		EXTRACT(EPOCH FROM m.time_stamp)::BIGINT,
		m.time_stamp,
		m.token_id,
		m.round_num,
		m.seed,
		COALESCE(m.token_name, ''),
		wa.addr,
		m.cur_owner_aid,
		oa.addr,
		EXISTS(SELECT 1 FROM cg_staked_token_cst st WHERE st.token_id=m.token_id),
		EXISTS(SELECT 1 FROM cg_prize_claim pc WHERE pc.token_id=m.token_id),
		(SELECT rnw.is_rwalk FROM cg_raffle_nft_prize rnw
			WHERE rnw.token_id=m.token_id AND rnw.round_num=m.round_num),
		(SELECT rnw.is_staker FROM cg_raffle_nft_prize rnw
			WHERE rnw.token_id=m.token_id AND rnw.round_num=m.round_num),
		EXISTS(SELECT 1 FROM cg_endurance_prize endu
			WHERE endu.erc721_token_id=m.token_id AND endu.round_num=m.round_num),
		EXISTS(SELECT 1 FROM cg_lastcst_prize stel
			WHERE stel.erc721_token_id=m.token_id AND stel.round_num=m.round_num),
		EXISTS(SELECT 1 FROM cg_chrono_warrior_prize cw
			WHERE cw.nft_id=m.token_id AND cw.round_num=m.round_num)`

// globalTokenSelectSQL is the directory page query; callers append
// WHERE/ORDER BY/paging.
const globalTokenSelectSQL = `SELECT` + globalTokenProjectionSQL + `
	FROM cg_mint_event m
		LEFT JOIN transaction t ON t.id=m.tx_id
		LEFT JOIN address wa ON wa.address_id=m.owner_aid
		LEFT JOIN address oa ON oa.address_id=m.cur_owner_aid`

func scanGlobalToken(rows pgx.Rows, rec *GlobalTokenRecord) error {
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
		&rec.TokenID,
		&rec.MintRound,
		&rec.Seed,
		&rec.TokenName,
		&rec.WinnerAddr,
		&rec.CurOwnerAid,
		&rec.CurOwnerAddr,
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

// CosmicSignatureTokensGlobalPage returns at most limit minted Cosmic
// Signature tokens, newest mint first by token ID, optionally narrowed to
// named tokens or a case-insensitive name search. Mints are append-only,
// so pages are stable; owner, name and staking columns are live.
func (r *Repo) CosmicSignatureTokensGlobalPage(
	ctx context.Context,
	filter GlobalTokenFilter,
	after *GlobalTokenPageCursor,
	limit int,
) (records []GlobalTokenRecord, hasMore bool, err error) {
	const op = "cosmic signature tokens global page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	if !filter.valid() {
		return nil, false, fmt.Errorf("%s: contradictory filter", op)
	}

	query := globalTokenSelectSQL
	var conditions []string
	var args []any
	if filter.NamedOnly {
		conditions = append(conditions, "LENGTH(m.token_name) > 0")
	}
	if filter.NameContains != "" {
		args = append(args, "%"+escapeLikePattern(filter.NameContains)+"%")
		conditions = append(conditions,
			fmt.Sprintf(`m.token_name ILIKE $%d ESCAPE '\'`, len(args)))
	}
	if after != nil {
		args = append(args, after.TokenID)
		conditions = append(conditions, fmt.Sprintf("m.token_id < $%d", len(args)))
	}
	if len(conditions) > 0 {
		query += "\n\tWHERE " + strings.Join(conditions, " AND ")
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
	ORDER BY m.token_id DESC
	LIMIT $%d`, len(args))

	records, err = queryList(ctx, r, op, limit+1, query, scanGlobalToken, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// GlobalTokenStake is the stake currently locking one token.
type GlobalTokenStake struct {
	StakeActionID int64
	StakerAid     int64
	StakerAddr    string
	StakedAt      int64
	StakedAtText  string
}

// GlobalTokenDetailRecord is one minted token with provenance, current
// owner and, while staked, the locking stake action.
type GlobalTokenDetailRecord struct {
	GlobalTokenRecord
	CurrentStake *GlobalTokenStake
}

// CosmicSignatureTokenDetailV2 returns one minted token with mint
// provenance, current owner, name and live staking state, or
// store.ErrNotFound for an unknown token ID.
func (r *Repo) CosmicSignatureTokenDetailV2(
	ctx context.Context,
	tokenID int64,
) (GlobalTokenDetailRecord, error) {
	const op = "cosmic signature token detail v2"
	if tokenID < 0 {
		return GlobalTokenDetailRecord{}, fmt.Errorf("%s: invalid token id", op)
	}

	query := `SELECT` + globalTokenProjectionSQL + `,
			st.stake_action_id,
			st.staker_aid,
			sta.addr,
			EXTRACT(EPOCH FROM sa.time_stamp)::BIGINT,
			sa.time_stamp
		FROM cg_mint_event m
			LEFT JOIN transaction t ON t.id=m.tx_id
			LEFT JOIN address wa ON wa.address_id=m.owner_aid
			LEFT JOIN address oa ON oa.address_id=m.cur_owner_aid
			LEFT JOIN cg_staked_token_cst st ON st.token_id=m.token_id
			LEFT JOIN address sta ON sta.address_id=st.staker_aid
			LEFT JOIN cg_nft_staked_cst sa ON sa.action_id=st.stake_action_id
		WHERE m.token_id=$1`

	var (
		rec           GlobalTokenDetailRecord
		isMainPrize   bool
		raffleIsRWalk sql.NullBool
		raffleStaker  sql.NullBool
		isEndurance   bool
		isLastCst     bool
		isChrono      bool
		stakeActionID sql.NullInt64
		stakerAid     sql.NullInt64
		stakerAddr    sql.NullString
		stakedAt      sql.NullInt64
		stakedAtText  string
	)
	err := r.pool().QueryRow(ctx, query, tokenID).Scan(
		&rec.MintTx.EvtLogId,
		&rec.MintTx.BlockNum,
		&rec.MintTx.TxId,
		&rec.MintTx.TxHash,
		&rec.MintTx.TimeStamp,
		store.TimeText(&rec.MintTx.DateTime),
		&rec.TokenID,
		&rec.MintRound,
		&rec.Seed,
		&rec.TokenName,
		&rec.WinnerAddr,
		&rec.CurOwnerAid,
		&rec.CurOwnerAddr,
		&rec.Staked,
		&isMainPrize,
		&raffleIsRWalk,
		&raffleStaker,
		&isEndurance,
		&isLastCst,
		&isChrono,
		&stakeActionID,
		&stakerAid,
		&stakerAddr,
		&stakedAt,
		store.NullTimeText(&stakedAtText),
	)
	if err != nil {
		return GlobalTokenDetailRecord{}, store.WrapError(op, err)
	}

	source, err := deriveMintSource(
		rec.TokenID, isMainPrize, raffleIsRWalk, raffleStaker, isEndurance, isLastCst, isChrono)
	if err != nil {
		return GlobalTokenDetailRecord{}, store.WrapError(op, err)
	}
	rec.MintSource = source

	if rec.Staked {
		if !stakeActionID.Valid || !stakerAid.Valid || !stakerAddr.Valid ||
			!stakedAt.Valid || stakedAtText == "" {
			return GlobalTokenDetailRecord{}, store.WrapError(op,
				fmt.Errorf("token %d is staked but misses its stake action", rec.TokenID))
		}
		rec.CurrentStake = &GlobalTokenStake{
			StakeActionID: stakeActionID.Int64,
			StakerAid:     stakerAid.Int64,
			StakerAddr:    stakerAddr.String,
			StakedAt:      stakedAt.Int64,
			StakedAtText:  stakedAtText,
		}
	}
	return rec, nil
}

// TokenEventPageCursor identifies the last event returned by one
// token-scoped, newest-first event page.
type TokenEventPageCursor struct {
	EventLogID int64
}

func (c *TokenEventPageCursor) valid() bool {
	return c == nil || c.EventLogID >= 1
}

// TokenNameChangeRecord is one rename of one token. An empty NewName
// cleared the token's name.
type TokenNameChangeRecord struct {
	Tx           cgmodel.Transaction
	TokenID      int64
	NewName      string
	ChangedByAid int64
	ChangedBy    string
}

// TokenNameHistoryPage returns at most limit renames of one token, newest
// first by immutable event-log ID, with the wallet that performed each
// rename.
func (r *Repo) TokenNameHistoryPage(
	ctx context.Context,
	tokenID int64,
	after *TokenEventPageCursor,
	limit int,
) (records []TokenNameChangeRecord, hasMore bool, err error) {
	const op = "token name history page"
	if tokenID < 0 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid token id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
			n.evtlog_id,
			n.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM n.time_stamp)::BIGINT,
			n.time_stamp,
			n.token_id,
			n.token_name,
			t.from_aid,
			a.addr
		FROM cg_token_name n
			LEFT JOIN transaction t ON t.id=n.tx_id
			LEFT JOIN address a ON a.address_id=t.from_aid
		WHERE n.token_id=$1`
	args := []any{tokenID}
	if after != nil {
		args = append(args, after.EventLogID)
		query += fmt.Sprintf(" AND n.evtlog_id < $%d", len(args))
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
		ORDER BY n.evtlog_id DESC
		LIMIT $%d`, len(args))

	scan := func(rows pgx.Rows, rec *TokenNameChangeRecord) error {
		return rows.Scan(
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.TokenID,
			&rec.NewName,
			&rec.ChangedByAid,
			&rec.ChangedBy,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// TokenTransferRecord is one ERC-721 transfer of one token. TransferType
// is the raw otype column (0 transfer, 1 mint, 2 burn).
type TokenTransferRecord struct {
	Tx           cgmodel.Transaction
	TokenID      int64
	FromAid      int64
	FromAddr     string
	ToAid        int64
	ToAddr       string
	TransferType int64
}

// TokenTransfersPage returns at most limit ERC-721 transfers of one token,
// newest first by immutable event-log ID, the mint included.
func (r *Repo) TokenTransfersPage(
	ctx context.Context,
	tokenID int64,
	after *TokenEventPageCursor,
	limit int,
) (records []TokenTransferRecord, hasMore bool, err error) {
	const op = "token transfers page"
	if tokenID < 0 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid token id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
			tr.evtlog_id,
			tr.block_num,
			tx.id,
			tx.tx_hash,
			EXTRACT(EPOCH FROM tr.time_stamp)::BIGINT,
			tr.time_stamp,
			tr.token_id,
			tr.from_aid,
			fa.addr,
			tr.to_aid,
			ta.addr,
			tr.otype
		FROM cg_erc721_transfer tr
			LEFT JOIN transaction tx ON tx.id=tr.tx_id
			LEFT JOIN address fa ON fa.address_id=tr.from_aid
			LEFT JOIN address ta ON ta.address_id=tr.to_aid
		WHERE tr.token_id=$1`
	args := []any{tokenID}
	if after != nil {
		args = append(args, after.EventLogID)
		query += fmt.Sprintf(" AND tr.evtlog_id < $%d", len(args))
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
		ORDER BY tr.evtlog_id DESC
		LIMIT $%d`, len(args))

	scan := func(rows pgx.Rows, rec *TokenTransferRecord) error {
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
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// CosmicSignatureHolderRecord is one wallet's current Cosmic Signature
// token count.
type CosmicSignatureHolderRecord struct {
	OwnerAid   int64
	Address    string
	TokenCount int64
}

// CosmicSignatureHoldersPage returns wallets ordered by the number of
// Cosmic Signature tokens they currently hold, then address ID. Holdings
// change on transfers, so ranks are weakly consistent across pages.
func (r *Repo) CosmicSignatureHoldersPage(
	ctx context.Context,
	after *ParticipantPageCursor,
	limit int,
) ([]CosmicSignatureHolderRecord, bool, error) {
	const op = "cosmic signature holders page"
	if err := validateParticipantPage(op, ParticipantCsTokenHolders, after, limit); err != nil {
		return nil, false, err
	}
	query := `SELECT h.owner_aid,a.addr,h.token_count
		FROM (
			SELECT m.cur_owner_aid AS owner_aid,COUNT(*)::BIGINT AS token_count
			FROM cg_mint_event m
			GROUP BY m.cur_owner_aid
		) h
		LEFT JOIN address a ON a.address_id=h.owner_aid
		ORDER BY h.token_count DESC,h.owner_aid ASC
		LIMIT $1`
	args := []any{limit + 1}
	if after != nil {
		query = `SELECT h.owner_aid,a.addr,h.token_count
			FROM (
				SELECT m.cur_owner_aid AS owner_aid,COUNT(*)::BIGINT AS token_count
				FROM cg_mint_event m
				GROUP BY m.cur_owner_aid
			) h
			LEFT JOIN address a ON a.address_id=h.owner_aid
			WHERE h.token_count < $1
				OR (h.token_count = $1 AND h.owner_aid > $2)
			ORDER BY h.token_count DESC,h.owner_aid ASC
			LIMIT $3`
		args = []any{participantInt64SortValue(after), after.AddressID, limit + 1}
	}
	scan := func(rows pgx.Rows, rec *CosmicSignatureHolderRecord) error {
		return rows.Scan(&rec.OwnerAid, &rec.Address, &rec.TokenCount)
	}
	records, err := queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore := trimParticipantPage(records, limit)
	return records, hasMore, nil
}

// CosmicTokenHolderRecord is one wallet's current exact Cosmic Token
// base-unit balance.
type CosmicTokenHolderRecord struct {
	OwnerAid   int64
	Address    string
	BalanceWei string
}

// CosmicTokenHoldersPage returns wallets with a positive Cosmic Token
// balance ordered by balance, then address ID, with exact base-unit
// strings. Balances change on transfers, so ranks are weakly consistent
// across pages.
func (r *Repo) CosmicTokenHoldersPage(
	ctx context.Context,
	after *ParticipantPageCursor,
	limit int,
) ([]CosmicTokenHolderRecord, bool, error) {
	const op = "cosmic token holders page"
	if err := validateParticipantPage(op, ParticipantCosmicTokenHolders, after, limit); err != nil {
		return nil, false, err
	}
	query := `SELECT o.owner_aid,a.addr,o.cur_balance::TEXT
		FROM cg_costok_owner o
		LEFT JOIN address a ON a.address_id=o.owner_aid
		WHERE o.cur_balance > 0
		ORDER BY o.cur_balance DESC,o.owner_aid ASC
		LIMIT $1`
	args := []any{limit + 1}
	if after != nil {
		query = `SELECT o.owner_aid,a.addr,o.cur_balance::TEXT
			FROM cg_costok_owner o
			LEFT JOIN address a ON a.address_id=o.owner_aid
			WHERE o.cur_balance > 0
				AND (o.cur_balance < $1::NUMERIC
					OR (o.cur_balance = $1::NUMERIC AND o.owner_aid > $2))
			ORDER BY o.cur_balance DESC,o.owner_aid ASC
			LIMIT $3`
		args = []any{after.SortValue, after.AddressID, limit + 1}
	}
	scan := func(rows pgx.Rows, rec *CosmicTokenHolderRecord) error {
		return rows.Scan(&rec.OwnerAid, &rec.Address, &rec.BalanceWei)
	}
	records, err := queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore := trimParticipantPage(records, limit)
	return records, hasMore, nil
}

// CosmicTokenTopHolderRecord is one of the largest Cosmic Token holders
// inside the statistics snapshot.
type CosmicTokenTopHolderRecord struct {
	OwnerAid   int64  `json:"aid"`
	Address    string `json:"addr"`
	BalanceWei string `json:"balance"`
}

// CosmicTokenStatisticsRecord is the game-wide exact Cosmic Token position
// from one database snapshot.
type CosmicTokenStatisticsRecord struct {
	TotalSupplyWei             string
	HolderCount                int64
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
	TopHolders                 []CosmicTokenTopHolderRecord
}

// CosmicTokenStatisticsV2 computes the game-wide exact Cosmic Token
// position in one query, so every field — the top-holder list included —
// comes from the same database snapshot. Sums skip the -1 "not a CST bid"
// sentinels exactly like the per-user summary.
func (r *Repo) CosmicTokenStatisticsV2(ctx context.Context) (CosmicTokenStatisticsRecord, error) {
	const op = "cosmic token statistics v2"

	query := `WITH supply AS (
			SELECT
				COUNT(*)::BIGINT AS holders,
				COALESCE(SUM(cur_balance), 0) AS wei
			FROM cg_costok_owner WHERE cur_balance > 0
		), bidding AS (
			SELECT
				COALESCE(SUM(cst_reward) FILTER (WHERE cst_reward > 0), 0) AS earned,
				COALESCE(SUM(cst_price) FILTER (WHERE cst_price > 0), 0) AS consumed
			FROM cg_bid
		), main_prizes AS (
			SELECT COALESCE(SUM(cst_amount) FILTER (WHERE cst_amount > 0), 0) AS earned
			FROM cg_prize_claim
		), raffles AS (
			SELECT COALESCE(SUM(cst_amount) FILTER (WHERE cst_amount > 0), 0) AS earned
			FROM cg_raffle_nft_prize
		), chrono AS (
			SELECT COALESCE(SUM(cst_amount) FILTER (WHERE cst_amount > 0), 0) AS earned
			FROM cg_chrono_warrior_prize
		), endurance AS (
			SELECT COALESCE(SUM(erc20_amount) FILTER (WHERE erc20_amount > 0), 0) AS earned
			FROM cg_endurance_prize
		), lastcst AS (
			SELECT COALESCE(SUM(erc20_amount) FILTER (WHERE erc20_amount > 0), 0) AS earned
			FROM cg_lastcst_prize
		), marketing AS (
			SELECT COALESCE(SUM(amount), 0) AS earned
			FROM cg_mkt_reward
		), activity AS (
			SELECT
				COUNT(*)::BIGINT AS transfers,
				COUNT(*) FILTER (WHERE otype = 1)::BIGINT AS mints,
				COUNT(*) FILTER (WHERE otype = 2)::BIGINT AS burns
			FROM cg_erc20_transfer
		), top AS (
			SELECT o.owner_aid, a.addr, o.cur_balance
			FROM cg_costok_owner o
				LEFT JOIN address a ON a.address_id=o.owner_aid
			WHERE o.cur_balance > 0
			ORDER BY o.cur_balance DESC, o.owner_aid ASC
			LIMIT 10
		), top_json AS (
			SELECT COALESCE(jsonb_agg(jsonb_build_object(
				'aid', t.owner_aid,
				'addr', t.addr,
				'balance', t.cur_balance::TEXT
			) ORDER BY t.cur_balance DESC, t.owner_aid ASC), '[]'::jsonb) AS holders
			FROM top t
		)
		SELECT
			s.wei::TEXT,
			s.holders,
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
			a.burns,
			tj.holders
		FROM supply s
			CROSS JOIN bidding bd
			CROSS JOIN main_prizes mp
			CROSS JOIN raffles rf
			CROSS JOIN chrono ch
			CROSS JOIN endurance en
			CROSS JOIN lastcst lc
			CROSS JOIN marketing mk
			CROSS JOIN activity a
			CROSS JOIN top_json tj`

	var record CosmicTokenStatisticsRecord
	var topHolders []byte
	err := r.pool().QueryRow(ctx, query).Scan(
		&record.TotalSupplyWei,
		&record.HolderCount,
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
		&topHolders,
	)
	if err != nil {
		return CosmicTokenStatisticsRecord{}, store.WrapError(op, err)
	}
	if err := json.Unmarshal(topHolders, &record.TopHolders); err != nil {
		return CosmicTokenStatisticsRecord{}, store.WrapError(op+": top holders", err)
	}
	return record, nil
}

// SupplyChangePageCursor identifies the last bid returned by the ascending
// per-bid supply ledger.
type SupplyChangePageCursor struct {
	EventLogID int64
}

func (c *SupplyChangePageCursor) valid() bool {
	return c == nil || c.EventLogID >= 1
}

// SupplyChangeRecord is one bid's exact effect on the Cosmic Token supply
// with the running total after the bid.
type SupplyChangeRecord struct {
	Tx             cgmodel.Transaction
	BidderAid      int64
	BidderAddr     string
	BidType        int64
	MintedWei      string
	BurnedWei      string
	NetWei         string
	TotalSupplyWei string
}

// CosmicTokenSupplyByBidPage returns at most limit per-bid supply changes,
// oldest first by immutable event-log ID, with exact running totals. The
// running total resumes from an aggregate over everything at or before the
// cursor, so later pages do not rescan their prefix rows.
func (r *Repo) CosmicTokenSupplyByBidPage(
	ctx context.Context,
	after *SupplyChangePageCursor,
	limit int,
) (records []SupplyChangeRecord, hasMore bool, err error) {
	const op = "cosmic token supply by bid page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	const netExpression = `GREATEST(COALESCE(b.cst_reward, 0), 0)
		- CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END`
	selectList := `SELECT
			b.evtlog_id,
			b.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM b.time_stamp)::BIGINT,
			b.time_stamp,
			b.bidder_aid,
			ba.addr,
			b.bid_type,
			GREATEST(COALESCE(b.cst_reward, 0), 0)::TEXT,
			(CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END)::TEXT,
			(` + netExpression + `)::TEXT,`

	var query string
	var args []any
	if after == nil {
		query = selectList + `
			SUM(` + netExpression + `)
				OVER (ORDER BY b.evtlog_id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW)::TEXT
		FROM cg_bid b
			LEFT JOIN transaction t ON t.id=b.tx_id
			LEFT JOIN address ba ON ba.address_id=b.bidder_aid
		ORDER BY b.evtlog_id
		LIMIT $1`
		args = []any{limit + 1}
	} else {
		query = `WITH base AS (
			SELECT COALESCE(SUM(` + netExpression + `), 0) AS supply
			FROM cg_bid b WHERE b.evtlog_id <= $1
		)
		` + selectList + `
			((SELECT supply FROM base) + SUM(` + netExpression + `)
				OVER (ORDER BY b.evtlog_id ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW))::TEXT
		FROM cg_bid b
			LEFT JOIN transaction t ON t.id=b.tx_id
			LEFT JOIN address ba ON ba.address_id=b.bidder_aid
		WHERE b.evtlog_id > $1
		ORDER BY b.evtlog_id
		LIMIT $2`
		args = []any{after.EventLogID, limit + 1}
	}

	scan := func(rows pgx.Rows, rec *SupplyChangeRecord) error {
		return rows.Scan(
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.BidderAid,
			&rec.BidderAddr,
			&rec.BidType,
			&rec.MintedWei,
			&rec.BurnedWei,
			&rec.NetWei,
			&rec.TotalSupplyWei,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// MaxSupplyDailyWindowDays caps the daily supply window; one row per day
// bounds the response at 2,000 rows.
const MaxSupplyDailyWindowDays = 2000

// DailySupplyRecord is one UTC calendar day's exact Cosmic Token supply
// movement with the running total at the end of the day.
type DailySupplyRecord struct {
	Date           string
	BidCount       int64
	MintedWei      string
	BurnedWei      string
	NetWei         string
	TotalSupplyWei string
}

// CosmicTokenSupplyDaily returns one row per UTC day with bids inside the
// half-open [from, to) window: exact minted, burned and net base units and
// the running total supply computed over all history up to that day. The
// window may span at most MaxSupplyDailyWindowDays days.
func (r *Repo) CosmicTokenSupplyDaily(
	ctx context.Context,
	from, to time.Time,
) ([]DailySupplyRecord, error) {
	const op = "cosmic token supply daily"
	if !from.Before(to) {
		return nil, fmt.Errorf("%s: empty window", op)
	}
	if to.Sub(from) > MaxSupplyDailyWindowDays*24*time.Hour {
		return nil, fmt.Errorf("%s: window exceeds %d days", op, MaxSupplyDailyWindowDays)
	}

	query := `WITH daily AS (
			SELECT
				DATE(b.time_stamp) AS bid_date,
				COUNT(*)::BIGINT AS num_bids,
				SUM(GREATEST(COALESCE(b.cst_reward, 0), 0)) AS mint_amt,
				SUM(CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END) AS burn_amt,
				SUM(GREATEST(COALESCE(b.cst_reward, 0), 0)
					- CASE WHEN b.bid_type = 2 AND b.cst_price > 0 THEN b.cst_price ELSE 0 END) AS net_amt
			FROM cg_bid b
			GROUP BY DATE(b.time_stamp)
		), with_totals AS (
			SELECT
				d.bid_date,
				d.num_bids,
				d.mint_amt, d.burn_amt, d.net_amt,
				SUM(d.net_amt) OVER (ORDER BY d.bid_date
					ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW) AS total_supply
			FROM daily d
		)
		SELECT
			w.bid_date::TEXT,
			w.num_bids,
			w.mint_amt::TEXT,
			w.burn_amt::TEXT,
			w.net_amt::TEXT,
			w.total_supply::TEXT
		FROM with_totals w
		WHERE w.bid_date >= $1::DATE AND w.bid_date < $2::DATE
		ORDER BY w.bid_date`

	scan := func(rows pgx.Rows, rec *DailySupplyRecord) error {
		return rows.Scan(
			&rec.Date,
			&rec.BidCount,
			&rec.MintedWei,
			&rec.BurnedWei,
			&rec.NetWei,
			&rec.TotalSupplyWei,
		)
	}
	return queryList(ctx, r, op, 64, query, scan,
		from.UTC().Format("2006-01-02"), to.UTC().Format("2006-01-02"))
}

// MarketingRewardRecord is one MarketingWallet reward with the rewarded
// wallet and the exact Cosmic Token base-unit amount.
type MarketingRewardRecord struct {
	Tx           cgmodel.Transaction
	MarketerAid  int64
	MarketerAddr string
	AmountWei    string
}

// MarketingRewardsGlobalPage returns at most limit marketing rewards,
// newest first by immutable event-log ID, with the rewarded wallet on each
// row.
func (r *Repo) MarketingRewardsGlobalPage(
	ctx context.Context,
	after *UserEventPageCursor,
	limit int,
) (records []MarketingRewardRecord, hasMore bool, err error) {
	const op = "marketing rewards global page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
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
			a.addr,
			m.amount::TEXT
		FROM cg_mkt_reward m
			LEFT JOIN transaction t ON t.id=m.tx_id
			LEFT JOIN address a ON a.address_id=m.marketer_aid`
	var args []any
	if after != nil {
		args = append(args, after.EventLogID)
		query += fmt.Sprintf(`
		WHERE m.evtlog_id < $%d`, len(args))
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
		ORDER BY m.evtlog_id DESC
		LIMIT $%d`, len(args))

	scan := func(rows pgx.Rows, rec *MarketingRewardRecord) error {
		return rows.Scan(
			&rec.Tx.EvtLogId,
			&rec.Tx.BlockNum,
			&rec.Tx.TxId,
			&rec.Tx.TxHash,
			&rec.Tx.TimeStamp,
			store.TimeText(&rec.Tx.DateTime),
			&rec.MarketerAid,
			&rec.MarketerAddr,
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
