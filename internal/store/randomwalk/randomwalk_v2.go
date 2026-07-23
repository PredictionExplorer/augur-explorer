// API v2 read surface of the RandomWalk domain: bounded keyset pages over
// the mint directory, per-token event ledgers, the marketplace order book
// and its immutable offer/trade history, wallet-scoped projections, exact
// statistics snapshots and bounded chart series. Every amount is an exact
// wei string produced by ::TEXT casts; nothing here divides through 1e18.

package randomwalk

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
	marketstore "github.com/PredictionExplorer/augur-explorer/internal/store/marketplace"
)

// truncatePage reports whether more records than limit exist and trims the
// probe row every keyset page fetches.
func truncatePage[T any](records []T, limit int) ([]T, bool) {
	if len(records) > limit {
		return records[:limit], true
	}
	return records, false
}

// EventTx is the shared transaction identity of one indexed RandomWalk
// event: the immutable event-log ID, its block, transaction hash and
// timestamp (Unix seconds plus the RFC 3339 rendering used by mappers).
type EventTx struct {
	EvtLogID  int64
	BlockNum  int64
	TxID      int64
	TxHash    string
	TimeStamp int64
	DateTime  string
}

// EventPageCursor identifies the last immutable event returned by a
// newest-first ledger page.
type EventPageCursor struct {
	EventLogID int64
}

func (c *EventPageCursor) valid() bool {
	return c == nil || c.EventLogID >= 1
}

// =============================================================================
// TOKEN DIRECTORY
// =============================================================================

// TokenSort selects the order of the token directory.
type TokenSort string

// Token directory orders: the immutable ascending mint order and the
// mutable most-traded ranking (trade count descending, token ID ascending).
const (
	TokenSortByID     TokenSort = "tokenId"
	TokenSortByTrades TokenSort = "mostTraded"
)

func (s TokenSort) valid() bool {
	return s == TokenSortByID || s == TokenSortByTrades
}

// TokenFilter narrows the token directory. At most one of NamedOnly and
// NameContains may be set; NameContains matches the current token name
// case-insensitively with ILIKE wildcards escaped, so the term is always
// literal. MintedFrom/MintedUntil bound the mint timestamp as a half-open
// [from, until) window of Unix seconds.
type TokenFilter struct {
	NamedOnly    bool
	NameContains string
	MintedFrom   *int64
	MintedUntil  *int64
}

func (f TokenFilter) valid() bool {
	if f.NamedOnly && f.NameContains != "" {
		return false
	}
	if f.MintedFrom != nil && *f.MintedFrom < 0 {
		return false
	}
	if f.MintedUntil != nil && *f.MintedUntil < 1 {
		return false
	}
	return true
}

// TokenPageCursor identifies the last token returned by a directory page.
// TradeCount participates only in the TokenSortByTrades order.
type TokenPageCursor struct {
	TokenID    int64
	TradeCount int64
}

func (c *TokenPageCursor) valid() bool {
	return c == nil || (c.TokenID >= 0 && c.TradeCount >= 0)
}

// escapeLikePattern makes term literal inside an ILIKE pattern by escaping
// the backslash escape character itself and both wildcards.
func escapeLikePattern(term string) string {
	replacer := strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`)
	return replacer.Replace(term)
}

// TokenRecord is one minted RandomWalk token: mint provenance plus the live
// owner, name and exact trading state.
type TokenRecord struct {
	MintTx           EventTx
	TokenID          int64
	MinterAid        int64
	MinterAddr       string
	Seed             string
	SeedNum          string
	MintPriceWei     string
	CurOwnerAid      int64
	CurOwnerAddr     string
	TokenName        string
	LastPriceWei     string
	TradeCount       int64
	TradingVolumeWei string
}

// tokenProjectionSQL joins each mint event to the live rw_token row; the
// callers append WHERE/ORDER BY/paging.
// #nosec G101 -- a SQL projection, not credentials.
const tokenProjectionSQL = `SELECT
		m.evtlog_id,
		m.block_num,
		t.id,
		t.tx_hash,
		EXTRACT(EPOCH FROM m.time_stamp)::BIGINT,
		m.time_stamp,
		m.token_id,
		m.owner_aid,
		ma.addr,
		m.seed,
		m.seed_num::TEXT,
		m.price::TEXT,
		tk.cur_owner_aid,
		oa.addr,
		COALESCE(tk.last_name, ''),
		tk.last_price::TEXT,
		tk.num_trades,
		tk.total_vol::TEXT
	FROM rw_mint_evt m
		INNER JOIN rw_token tk ON tk.rwalk_aid=m.contract_aid AND tk.token_id=m.token_id
		INNER JOIN transaction t ON t.id=m.tx_id
		INNER JOIN address ma ON ma.address_id=m.owner_aid
		INNER JOIN address oa ON oa.address_id=tk.cur_owner_aid`

func scanTokenRecord(rows pgx.Rows, record *TokenRecord) error {
	return rows.Scan(
		&record.MintTx.EvtLogID,
		&record.MintTx.BlockNum,
		&record.MintTx.TxID,
		&record.MintTx.TxHash,
		&record.MintTx.TimeStamp,
		store.TimeText(&record.MintTx.DateTime),
		&record.TokenID,
		&record.MinterAid,
		&record.MinterAddr,
		&record.Seed,
		&record.SeedNum,
		&record.MintPriceWei,
		&record.CurOwnerAid,
		&record.CurOwnerAddr,
		&record.TokenName,
		&record.LastPriceWei,
		&record.TradeCount,
		&record.TradingVolumeWei,
	)
}

// TokensPage returns at most limit minted tokens in the requested order,
// optionally narrowed by TokenFilter. The tokenId order pages the immutable
// mint sequence; the mostTraded ranking is live and can move between pages.
func (r *Repo) TokensPage(
	ctx context.Context,
	filter TokenFilter,
	sort TokenSort,
	after *TokenPageCursor,
	limit int,
) (records []TokenRecord, hasMore bool, err error) {
	const op = "rwalk tokens page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	if !filter.valid() {
		return nil, false, fmt.Errorf("%s: contradictory filter", op)
	}
	if !sort.valid() {
		return nil, false, fmt.Errorf("%s: invalid sort", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return nil, false, err
	}

	args := []any{addrs.RandomWalkAid}
	conditions := []string{"m.contract_aid=$1"}
	if filter.NamedOnly {
		conditions = append(conditions, "LENGTH(tk.last_name) > 0")
	}
	if filter.NameContains != "" {
		args = append(args, "%"+escapeLikePattern(filter.NameContains)+"%")
		conditions = append(conditions,
			fmt.Sprintf(`tk.last_name ILIKE $%d ESCAPE '\'`, len(args)))
	}
	if filter.MintedFrom != nil {
		args = append(args, *filter.MintedFrom)
		conditions = append(conditions,
			fmt.Sprintf("m.time_stamp >= TO_TIMESTAMP($%d)", len(args)))
	}
	if filter.MintedUntil != nil {
		args = append(args, *filter.MintedUntil)
		conditions = append(conditions,
			fmt.Sprintf("m.time_stamp < TO_TIMESTAMP($%d)", len(args)))
	}

	orderBy := "ORDER BY m.token_id"
	if sort == TokenSortByTrades {
		orderBy = "ORDER BY tk.num_trades DESC, tk.token_id"
	}
	if after != nil {
		if sort == TokenSortByTrades {
			args = append(args, after.TradeCount, after.TokenID)
			conditions = append(conditions, fmt.Sprintf(
				"(tk.num_trades < $%d OR (tk.num_trades = $%d AND tk.token_id > $%d))",
				len(args)-1, len(args)-1, len(args)))
		} else {
			args = append(args, after.TokenID)
			conditions = append(conditions, fmt.Sprintf("m.token_id > $%d", len(args)))
		}
	}

	args = append(args, limit+1)
	query := tokenProjectionSQL +
		"\n\tWHERE " + strings.Join(conditions, " AND ") +
		fmt.Sprintf("\n\t%s\n\tLIMIT $%d", orderBy, len(args))

	records, err = queryList(ctx, r, op, limit+1, query, scanTokenRecord, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// TokenDetailRecord is one minted token plus its naming recency.
type TokenDetailRecord struct {
	TokenRecord

	NameChangeTs   int64
	NameChangeText string
}

// TokenDetailV2 returns one minted token with mint provenance, live state
// and the most recent rename timestamp, or store.ErrNotFound for a token
// the collection never minted.
func (r *Repo) TokenDetailV2(ctx context.Context, tokenID int64) (TokenDetailRecord, error) {
	const op = "rwalk token detail v2"
	if tokenID < 0 {
		return TokenDetailRecord{}, fmt.Errorf("%s: invalid token id", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return TokenDetailRecord{}, err
	}

	query := tokenProjectionSQL + `
	WHERE m.contract_aid=$1 AND m.token_id=$2
	ORDER BY m.evtlog_id DESC
	LIMIT 1`
	rows, err := r.q(ctx).Query(ctx, query, addrs.RandomWalkAid, tokenID)
	if err != nil {
		return TokenDetailRecord{}, store.WrapError(op, err)
	}
	defer rows.Close()
	var record TokenDetailRecord
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return TokenDetailRecord{}, store.WrapError(op, err)
		}
		return TokenDetailRecord{}, store.WrapError(op, pgx.ErrNoRows)
	}
	if err := scanTokenRecord(rows, &record.TokenRecord); err != nil {
		return TokenDetailRecord{}, store.WrapError(op, err)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return TokenDetailRecord{}, store.WrapError(op, err)
	}

	var nameTs sql.NullInt64
	var nameText string
	err = r.q(ctx).QueryRow(ctx, `SELECT
			EXTRACT(EPOCH FROM time_stamp)::BIGINT,
			time_stamp
		FROM rw_token_name
		WHERE contract_aid=$1 AND token_id=$2
		ORDER BY evtlog_id DESC
		LIMIT 1`, addrs.RandomWalkAid, tokenID).Scan(&nameTs, store.TimeText(&nameText))
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return TokenDetailRecord{}, store.WrapError(op+": last rename", err)
	}
	if nameTs.Valid {
		record.NameChangeTs = nameTs.Int64
		record.NameChangeText = nameText
	}
	return record, nil
}

// CollectionTokenExists reports whether the RandomWalk collection minted
// tokenID.
func (r *Repo) CollectionTokenExists(ctx context.Context, tokenID int64) (bool, error) {
	const op = "rwalk collection token exists"
	if tokenID < 0 {
		return false, fmt.Errorf("%s: invalid token id", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return false, err
	}
	var exists bool
	err = r.q(ctx).QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM rw_mint_evt WHERE contract_aid=$1 AND token_id=$2)",
		addrs.RandomWalkAid, tokenID).Scan(&exists)
	if err != nil {
		return false, store.WrapError(op, err)
	}
	return exists, nil
}

// TokenNameChangeRecord is one naming event of one token.
type TokenNameChangeRecord struct {
	Tx       EventTx
	TokenID  int64
	NewName  string
	OwnerAid int64
	Owner    string
}

// TokenNameChangesPageV2 returns at most limit renames of one token, newest
// first by immutable event-log ID.
func (r *Repo) TokenNameChangesPageV2(
	ctx context.Context,
	tokenID int64,
	after *EventPageCursor,
	limit int,
) (records []TokenNameChangeRecord, hasMore bool, err error) {
	const op = "rwalk token name changes page v2"
	if tokenID < 0 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid token or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return nil, false, err
	}

	query := `SELECT
			n.evtlog_id,
			n.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM n.time_stamp)::BIGINT,
			n.time_stamp,
			n.token_id,
			COALESCE(n.new_name, ''),
			t.from_aid,
			oa.addr
		FROM rw_token_name n
			INNER JOIN transaction t ON t.id=n.tx_id
			INNER JOIN address oa ON oa.address_id=t.from_aid
		WHERE n.contract_aid=$1 AND n.token_id=$2`
	args := []any{addrs.RandomWalkAid, tokenID}
	if after != nil {
		args = append(args, after.EventLogID)
		query += fmt.Sprintf(" AND n.evtlog_id < $%d", len(args))
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
		ORDER BY n.evtlog_id DESC
		LIMIT $%d`, len(args))
	scan := func(rows pgx.Rows, record *TokenNameChangeRecord) error {
		return rows.Scan(
			&record.Tx.EvtLogID,
			&record.Tx.BlockNum,
			&record.Tx.TxID,
			&record.Tx.TxHash,
			&record.Tx.TimeStamp,
			store.TimeText(&record.Tx.DateTime),
			&record.TokenID,
			&record.NewName,
			&record.OwnerAid,
			&record.Owner,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// =============================================================================
// PER-TOKEN EVENT LEDGER
// =============================================================================

// TokenEventKind discriminates the per-token provenance ledger rows.
type TokenEventKind string

// Per-token event kinds. Transfers that merely mirror the mint or a
// purchase in the same transaction are represented once, by the mint or
// purchase row.
const (
	TokenEventMint          TokenEventKind = "mint"
	TokenEventTransfer      TokenEventKind = "transfer"
	TokenEventNameChange    TokenEventKind = "nameChange"
	TokenEventListed        TokenEventKind = "listed"
	TokenEventOfferCanceled TokenEventKind = "offerCanceled"
	TokenEventPurchase      TokenEventKind = "purchase"
)

// TokenEventRecord is one row of the per-token provenance ledger. The
// populated optional fields follow Kind: mint fills Minter/Seed/SeedNum/
// PriceWei; transfer fills From/To; nameChange fills NewName (HasNewName
// distinguishes a cleared name from absence); the marketplace kinds fill
// OfferID/OfferType/PriceWei plus MakerAddr (listed, offerCanceled) or
// BuyerAddr and SellerAddr (purchase).
type TokenEventRecord struct {
	Tx         EventTx
	Kind       TokenEventKind
	TokenID    int64
	MinterAddr string
	Seed       string
	SeedNum    string
	PriceWei   string
	FromAddr   string
	ToAddr     string
	NewName    string
	HasNewName bool
	OfferID    int64
	HasOffer   bool
	OfferType  int16
	MakerAddr  string
	BuyerAddr  string
	SellerAddr string
}

// tokenEventsBranchSQL renders one branch of the per-token ledger merge.
// Every branch projects the same wide column list; kind, table and the
// optional columns are compile-time literals.
func tokenEventsBranch(kind, from, where, projection string, cursorArg, limitArg int) string {
	filter := where
	if cursorArg > 0 {
		filter += fmt.Sprintf(" AND e.evtlog_id < $%d", cursorArg)
	}
	return fmt.Sprintf(`(SELECT
			'%s'::TEXT AS event_kind,
			e.evtlog_id,
			e.block_num,
			t.id AS tx_id,
			t.tx_hash,
			EXTRACT(EPOCH FROM e.time_stamp)::BIGINT AS ts,
			e.time_stamp AS date_time,
			%s
		FROM %s
			INNER JOIN transaction t ON t.id=e.tx_id
		WHERE %s
		ORDER BY e.evtlog_id DESC
		LIMIT $%d)`, kind, projection, from, filter, limitArg)
}

// tokenEventsPageSQL merges the six per-token event sources newest first.
// Placeholders: $1 rwalkAid, $2 tokenID, $3 marketAid, then the optional
// cursor and finally the limit.
func tokenEventsPageSQL(withCursor bool) string {
	cursorArg, limitArg := 0, 4
	if withCursor {
		cursorArg, limitArg = 4, 5
	}
	branches := []string{
		tokenEventsBranch(string(TokenEventMint),
			"rw_mint_evt e",
			"e.contract_aid=$1 AND e.token_id=$2",
			`ma.addr AS minter_addr, e.seed, e.seed_num::TEXT AS seed_num,
			e.price::TEXT AS price_wei, NULL::TEXT AS from_addr, NULL::TEXT AS to_addr,
			NULL::TEXT AS new_name, NULL::BIGINT AS offer_id, NULL::SMALLINT AS otype,
			NULL::TEXT AS maker_addr, NULL::TEXT AS buyer_addr, NULL::TEXT AS seller_addr`,
			cursorArg, limitArg),
		tokenEventsBranch(string(TokenEventTransfer),
			"rw_transfer e",
			`e.contract_aid=$1 AND e.token_id=$2 AND e.otype <> 1
			AND NOT EXISTS (SELECT 1 FROM rw_item_bought ib
				INNER JOIN rw_new_offer po
					ON po.contract_aid=ib.contract_aid AND po.offer_id=ib.offer_id
				WHERE ib.tx_id=e.tx_id AND po.token_id=e.token_id AND po.rwalk_aid=e.contract_aid)`,
			`NULL::TEXT AS minter_addr, NULL::TEXT AS seed, NULL::TEXT AS seed_num,
			NULL::TEXT AS price_wei, fa.addr AS from_addr, ta.addr AS to_addr,
			NULL::TEXT AS new_name, NULL::BIGINT AS offer_id, NULL::SMALLINT AS otype,
			NULL::TEXT AS maker_addr, NULL::TEXT AS buyer_addr, NULL::TEXT AS seller_addr`,
			cursorArg, limitArg),
		tokenEventsBranch(string(TokenEventNameChange),
			"rw_token_name e",
			"e.contract_aid=$1 AND e.token_id=$2",
			`NULL::TEXT AS minter_addr, NULL::TEXT AS seed, NULL::TEXT AS seed_num,
			NULL::TEXT AS price_wei, NULL::TEXT AS from_addr, NULL::TEXT AS to_addr,
			COALESCE(e.new_name, '') AS new_name, NULL::BIGINT AS offer_id, NULL::SMALLINT AS otype,
			NULL::TEXT AS maker_addr, NULL::TEXT AS buyer_addr, NULL::TEXT AS seller_addr`,
			cursorArg, limitArg),
		tokenEventsBranch(string(TokenEventListed),
			"rw_new_offer e",
			"e.rwalk_aid=$1 AND e.token_id=$2 AND e.contract_aid=$3",
			`NULL::TEXT AS minter_addr, NULL::TEXT AS seed, NULL::TEXT AS seed_num,
			e.price::TEXT AS price_wei, NULL::TEXT AS from_addr, NULL::TEXT AS to_addr,
			NULL::TEXT AS new_name, e.offer_id, e.otype,
			mka.addr AS maker_addr, NULL::TEXT AS buyer_addr, NULL::TEXT AS seller_addr`,
			cursorArg, limitArg),
		tokenEventsBranch(string(TokenEventOfferCanceled),
			"rw_offer_canceled e",
			"e.contract_aid=$3 AND o.rwalk_aid=$1 AND o.token_id=$2",
			`NULL::TEXT AS minter_addr, NULL::TEXT AS seed, NULL::TEXT AS seed_num,
			o.price::TEXT AS price_wei, NULL::TEXT AS from_addr, NULL::TEXT AS to_addr,
			NULL::TEXT AS new_name, e.offer_id, o.otype,
			mka.addr AS maker_addr, NULL::TEXT AS buyer_addr, NULL::TEXT AS seller_addr`,
			cursorArg, limitArg),
		tokenEventsBranch(string(TokenEventPurchase),
			"rw_item_bought e",
			"e.contract_aid=$3 AND o.rwalk_aid=$1 AND o.token_id=$2",
			`NULL::TEXT AS minter_addr, NULL::TEXT AS seed, NULL::TEXT AS seed_num,
			o.price::TEXT AS price_wei, NULL::TEXT AS from_addr, NULL::TEXT AS to_addr,
			NULL::TEXT AS new_name, e.offer_id, o.otype,
			NULL::TEXT AS maker_addr, ba.addr AS buyer_addr, sa.addr AS seller_addr`,
			cursorArg, limitArg),
	}
	// Attach the per-branch joins the projections reference.
	branches[0] = strings.Replace(branches[0],
		"INNER JOIN transaction t ON t.id=e.tx_id",
		`INNER JOIN transaction t ON t.id=e.tx_id
			INNER JOIN address ma ON ma.address_id=e.owner_aid`, 1)
	branches[1] = strings.Replace(branches[1],
		"INNER JOIN transaction t ON t.id=e.tx_id",
		`INNER JOIN transaction t ON t.id=e.tx_id
			INNER JOIN address fa ON fa.address_id=e.from_aid
			INNER JOIN address ta ON ta.address_id=e.to_aid`, 1)
	branches[3] = strings.Replace(branches[3],
		"INNER JOIN transaction t ON t.id=e.tx_id",
		`INNER JOIN transaction t ON t.id=e.tx_id
			INNER JOIN address mka ON mka.address_id=
				CASE WHEN e.otype=1 THEN e.seller_aid ELSE e.buyer_aid END`, 1)
	branches[4] = strings.Replace(branches[4],
		"INNER JOIN transaction t ON t.id=e.tx_id",
		`INNER JOIN transaction t ON t.id=e.tx_id
			INNER JOIN rw_new_offer o
				ON o.contract_aid=e.contract_aid AND o.offer_id=e.offer_id
			INNER JOIN address mka ON mka.address_id=
				CASE WHEN o.otype=1 THEN o.seller_aid ELSE o.buyer_aid END`, 1)
	branches[5] = strings.Replace(branches[5],
		"INNER JOIN transaction t ON t.id=e.tx_id",
		`INNER JOIN transaction t ON t.id=e.tx_id
			INNER JOIN rw_new_offer o
				ON o.contract_aid=e.contract_aid AND o.offer_id=e.offer_id
			INNER JOIN address ba ON ba.address_id=e.buyer_aid
			INNER JOIN address sa ON sa.address_id=e.seller_aid`, 1)

	limitPlaceholder := fmt.Sprintf("$%d", limitArg)
	return `SELECT
			event_kind, evtlog_id, block_num, tx_id, tx_hash, ts, date_time,
			minter_addr, seed, seed_num, price_wei, from_addr, to_addr,
			new_name, offer_id, otype, maker_addr, buyer_addr, seller_addr
		FROM (` +
		strings.Join(branches, "\n\t\tUNION ALL\n\t\t") + `
		) events
		ORDER BY evtlog_id DESC
		LIMIT ` + limitPlaceholder
}

func scanTokenEvent(rows pgx.Rows, record *TokenEventRecord) error {
	var (
		minter, seed, seedNum, price sql.NullString
		fromAddr, toAddr, newName    sql.NullString
		maker, buyer, seller         sql.NullString
		offerID                      sql.NullInt64
		offerType                    sql.NullInt16
	)
	if err := rows.Scan(
		&record.Kind,
		&record.Tx.EvtLogID,
		&record.Tx.BlockNum,
		&record.Tx.TxID,
		&record.Tx.TxHash,
		&record.Tx.TimeStamp,
		store.TimeText(&record.Tx.DateTime),
		&minter,
		&seed,
		&seedNum,
		&price,
		&fromAddr,
		&toAddr,
		&newName,
		&offerID,
		&offerType,
		&maker,
		&buyer,
		&seller,
	); err != nil {
		return err
	}
	record.MinterAddr = minter.String
	record.Seed = seed.String
	record.SeedNum = seedNum.String
	record.PriceWei = price.String
	record.FromAddr = fromAddr.String
	record.ToAddr = toAddr.String
	record.NewName = newName.String
	record.HasNewName = newName.Valid
	record.OfferID = offerID.Int64
	record.HasOffer = offerID.Valid
	record.OfferType = offerType.Int16
	record.MakerAddr = maker.String
	record.BuyerAddr = buyer.String
	record.SellerAddr = seller.String
	return nil
}

// TokenEventsPage returns at most limit provenance events of one token,
// newest first by immutable event-log ID across all six sources.
func (r *Repo) TokenEventsPage(
	ctx context.Context,
	tokenID int64,
	after *EventPageCursor,
	limit int,
) (records []TokenEventRecord, hasMore bool, err error) {
	const op = "rwalk token events page"
	if tokenID < 0 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid token or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return nil, false, err
	}
	args := []any{addrs.RandomWalkAid, tokenID, addrs.MarketPlaceAid}
	if after != nil {
		args = append(args, after.EventLogID)
	}
	args = append(args, limit+1)
	records, err = queryList(
		ctx, r, op, limit+1, tokenEventsPageSQL(after != nil), scanTokenEvent, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// =============================================================================
// MARKETPLACE ORDER BOOK AND LEDGERS
// =============================================================================

// OfferSort selects the order of the live order book.
type OfferSort string

// Order-book orders: the immutable listing order in both directions and the
// live price ranking with an ascending event-log tie-break.
const (
	OfferSortNewest    OfferSort = "newest"
	OfferSortOldest    OfferSort = "oldest"
	OfferSortPriceAsc  OfferSort = "priceAsc"
	OfferSortPriceDesc OfferSort = "priceDesc"
)

func (s OfferSort) valid() bool {
	switch s {
	case OfferSortNewest, OfferSortOldest, OfferSortPriceAsc, OfferSortPriceDesc:
		return true
	default:
		return false
	}
}

// OfferPageCursor identifies the last offer returned by an order-book page.
// PriceWei participates only in the two price orders and must be a
// non-negative integer string there.
type OfferPageCursor struct {
	EventLogID int64
	PriceWei   string
}

func (c *OfferPageCursor) validFor(sort OfferSort) bool {
	if c == nil {
		return true
	}
	if c.EventLogID < 1 {
		return false
	}
	switch sort {
	case OfferSortPriceAsc, OfferSortPriceDesc:
		return validWeiString(c.PriceWei)
	case OfferSortNewest, OfferSortOldest:
		return c.PriceWei == ""
	default:
		return false
	}
}

// validWeiString reports whether value is a plain non-negative decimal
// integer (an exact wei amount).
func validWeiString(value string) bool {
	if value == "" {
		return false
	}
	for _, r := range value {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// OfferRecord is one currently active marketplace offer.
type OfferRecord struct {
	ListTx    EventTx
	OfferID   int64
	OfferType int16
	TokenID   int64
	PriceWei  string
	MakerAid  int64
	MakerAddr string
}

// ActiveOffersPage returns at most limit active offers in the requested
// order. The book is live: offers leave it when bought or canceled.
func (r *Repo) ActiveOffersPage(
	ctx context.Context,
	sort OfferSort,
	after *OfferPageCursor,
	limit int,
) (records []OfferRecord, hasMore bool, err error) {
	const op = "rwalk active offers page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !sort.valid() {
		return nil, false, fmt.Errorf("%s: invalid sort", op)
	}
	if !after.validFor(sort) {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return nil, false, err
	}
	marketSort := map[OfferSort]marketstore.OfferSort{
		OfferSortNewest:    marketstore.OfferSortNewest,
		OfferSortOldest:    marketstore.OfferSortOldest,
		OfferSortPriceAsc:  marketstore.OfferSortPriceAsc,
		OfferSortPriceDesc: marketstore.OfferSortPriceDesc,
	}[sort]
	var marketAfter *marketstore.OfferPageCursor
	if after != nil {
		marketAfter = &marketstore.OfferPageCursor{
			EventLogID: after.EventLogID,
			PriceWei:   after.PriceWei,
		}
	}
	marketRecords, hasMore, err := r.marketplace.ActiveOffersPage(
		ctx,
		marketstore.Scope{
			MarketplaceAid: addrs.MarketPlaceAid,
			CollectionAid:  addrs.RandomWalkAid,
		},
		marketSort,
		marketAfter,
		limit,
	)
	if err != nil {
		return nil, false, err
	}
	records = make([]OfferRecord, 0, len(marketRecords))
	for i := range marketRecords {
		record := marketRecords[i]
		records = append(records, OfferRecord{
			ListTx:    randomWalkEventTx(record.ListTx),
			OfferID:   record.OfferID,
			OfferType: record.OfferType,
			TokenID:   record.TokenID,
			PriceWei:  record.PriceWei,
			MakerAid:  record.MakerAid,
			MakerAddr: record.MakerAddr,
		})
	}
	return records, hasMore, nil
}

func randomWalkEventTx(record marketstore.EventTx) EventTx {
	return EventTx{
		EvtLogID:  record.EventLogID,
		BlockNum:  record.BlockNum,
		TxID:      record.TxID,
		TxHash:    record.TxHash,
		TimeStamp: record.TimeStamp,
		DateTime:  record.DateTime,
	}
}

// OfferOutcomePurchase is the purchase event that filled an offer, with the
// trigger-backfilled trading parties.
type OfferOutcomePurchase struct {
	Tx         EventTx
	BuyerAid   int64
	BuyerAddr  string
	SellerAid  int64
	SellerAddr string
}

// OfferHistoryRecord is one offer-creation event plus its current outcome:
// still active, or the purchase or cancellation that closed it. ProfitWei
// is the signed tracked seller profit ("" when the marketplace tracked no
// position).
type OfferHistoryRecord struct {
	ListTx       EventTx
	OfferID      int64
	OfferType    int16
	TokenID      int64
	PriceWei     string
	MakerAid     int64
	MakerAddr    string
	Active       bool
	ProfitWei    string
	Purchase     *OfferOutcomePurchase
	Cancellation *EventTx
}

const offerHistorySelectSQL = `SELECT
		o.evtlog_id,
		o.block_num,
		t.id,
		t.tx_hash,
		EXTRACT(EPOCH FROM o.time_stamp)::BIGINT,
		o.time_stamp,
		o.offer_id,
		o.otype,
		o.token_id,
		o.price::TEXT,
		CASE WHEN o.otype=1 THEN o.seller_aid ELSE o.buyer_aid END,
		mka.addr,
		o.active,
		o.profit::TEXT,
		ib.evtlog_id,
		ib.block_num,
		ibt.id,
		ibt.tx_hash,
		EXTRACT(EPOCH FROM ib.time_stamp)::BIGINT,
		ib.time_stamp,
		ib.buyer_aid,
		iba.addr,
		ib.seller_aid,
		ibs.addr,
		can.evtlog_id,
		can.block_num,
		cant.id,
		cant.tx_hash,
		EXTRACT(EPOCH FROM can.time_stamp)::BIGINT,
		can.time_stamp
	FROM rw_new_offer o
		INNER JOIN transaction t ON t.id=o.tx_id
		INNER JOIN address mka ON mka.address_id=
			CASE WHEN o.otype=1 THEN o.seller_aid ELSE o.buyer_aid END
		LEFT JOIN rw_item_bought ib
			ON ib.contract_aid=o.contract_aid AND ib.offer_id=o.offer_id
		LEFT JOIN transaction ibt ON ibt.id=ib.tx_id
		LEFT JOIN address iba ON iba.address_id=ib.buyer_aid
		LEFT JOIN address ibs ON ibs.address_id=ib.seller_aid
		LEFT JOIN rw_offer_canceled can
			ON can.contract_aid=o.contract_aid AND can.offer_id=o.offer_id
		LEFT JOIN transaction cant ON cant.id=can.tx_id`

func scanOfferHistory(rows pgx.Rows, record *OfferHistoryRecord) error {
	var (
		profit                 sql.NullString
		boughtEvt, boughtBlock sql.NullInt64
		boughtTxID             sql.NullInt64
		boughtHash             sql.NullString
		boughtTs               sql.NullInt64
		boughtText             string
		buyerAid, sellerAid    sql.NullInt64
		buyerAddr, sellerAddr  sql.NullString
		cancelEvt, cancelBlock sql.NullInt64
		cancelTxID             sql.NullInt64
		cancelHash             sql.NullString
		cancelTs               sql.NullInt64
		cancelText             string
	)
	if err := rows.Scan(
		&record.ListTx.EvtLogID,
		&record.ListTx.BlockNum,
		&record.ListTx.TxID,
		&record.ListTx.TxHash,
		&record.ListTx.TimeStamp,
		store.TimeText(&record.ListTx.DateTime),
		&record.OfferID,
		&record.OfferType,
		&record.TokenID,
		&record.PriceWei,
		&record.MakerAid,
		&record.MakerAddr,
		&record.Active,
		&profit,
		&boughtEvt,
		&boughtBlock,
		&boughtTxID,
		&boughtHash,
		&boughtTs,
		store.NullTimeText(&boughtText),
		&buyerAid,
		&buyerAddr,
		&sellerAid,
		&sellerAddr,
		&cancelEvt,
		&cancelBlock,
		&cancelTxID,
		&cancelHash,
		&cancelTs,
		store.NullTimeText(&cancelText),
	); err != nil {
		return err
	}
	if profit.Valid {
		record.ProfitWei = profit.String
	}
	if boughtEvt.Valid {
		record.Purchase = &OfferOutcomePurchase{
			Tx: EventTx{
				EvtLogID:  boughtEvt.Int64,
				BlockNum:  boughtBlock.Int64,
				TxID:      boughtTxID.Int64,
				TxHash:    boughtHash.String,
				TimeStamp: boughtTs.Int64,
				DateTime:  boughtText,
			},
			BuyerAid:   buyerAid.Int64,
			BuyerAddr:  buyerAddr.String,
			SellerAid:  sellerAid.Int64,
			SellerAddr: sellerAddr.String,
		}
	}
	if cancelEvt.Valid {
		record.Cancellation = &EventTx{
			EvtLogID:  cancelEvt.Int64,
			BlockNum:  cancelBlock.Int64,
			TxID:      cancelTxID.Int64,
			TxHash:    cancelHash.String,
			TimeStamp: cancelTs.Int64,
			DateTime:  cancelText,
		}
	}
	return nil
}

func (r *Repo) offerHistoryPage(
	ctx context.Context,
	op string,
	userAid *int64,
	after *EventPageCursor,
	limit int,
) (records []OfferHistoryRecord, hasMore bool, err error) {
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	if userAid != nil && *userAid < 1 {
		return nil, false, fmt.Errorf("%s: invalid address id", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return nil, false, err
	}
	if userAid == nil {
		var marketAfter *marketstore.EventPageCursor
		if after != nil {
			marketAfter = &marketstore.EventPageCursor{EventLogID: after.EventLogID}
		}
		marketRecords, marketHasMore, err := r.marketplace.OfferHistoryPage(
			ctx,
			marketstore.Scope{
				MarketplaceAid: addrs.MarketPlaceAid,
				CollectionAid:  addrs.RandomWalkAid,
			},
			marketAfter,
			limit,
		)
		if err != nil {
			return nil, false, err
		}
		records = make([]OfferHistoryRecord, 0, len(marketRecords))
		for i := range marketRecords {
			records = append(records, randomWalkOfferHistory(marketRecords[i]))
		}
		return records, marketHasMore, nil
	}

	query := offerHistorySelectSQL + `
	WHERE o.contract_aid=$1 AND o.rwalk_aid=$2`
	args := []any{addrs.MarketPlaceAid, addrs.RandomWalkAid}
	if userAid != nil {
		// Match the wallet as the offer maker or as a recorded trading
		// party; the placeholder counterparty column of an unfilled offer
		// never matches.
		args = append(args, *userAid)
		query += fmt.Sprintf(` AND (
			(o.otype=1 AND o.seller_aid=$%d) OR
			(o.otype=0 AND o.buyer_aid=$%d) OR
			(ib.id IS NOT NULL AND (ib.buyer_aid=$%d OR ib.seller_aid=$%d))
		)`, len(args), len(args), len(args), len(args))
	}
	if after != nil {
		args = append(args, after.EventLogID)
		query += fmt.Sprintf(" AND o.evtlog_id < $%d", len(args))
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
	ORDER BY o.evtlog_id DESC
	LIMIT $%d`, len(args))

	records, err = queryList(ctx, r, op, limit+1, query, scanOfferHistory, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

func randomWalkOfferHistory(record marketstore.OfferHistoryRecord) OfferHistoryRecord {
	output := OfferHistoryRecord{
		ListTx:    randomWalkEventTx(record.ListTx),
		OfferID:   record.OfferID,
		OfferType: record.OfferType,
		TokenID:   record.TokenID,
		PriceWei:  record.PriceWei,
		MakerAid:  record.MakerAid,
		MakerAddr: record.MakerAddr,
		Active:    record.Active,
		ProfitWei: record.ProfitWei,
	}
	if record.Purchase != nil {
		output.Purchase = &OfferOutcomePurchase{
			Tx:         randomWalkEventTx(record.Purchase.Tx),
			BuyerAid:   record.Purchase.BuyerAid,
			BuyerAddr:  record.Purchase.BuyerAddr,
			SellerAid:  record.Purchase.SellerAid,
			SellerAddr: record.Purchase.SellerAddr,
		}
	}
	if record.Cancellation != nil {
		cancellation := randomWalkEventTx(*record.Cancellation)
		output.Cancellation = &cancellation
	}
	return output
}

// OfferHistoryPage returns at most limit offer-creation events with their
// outcomes, newest first by immutable event-log ID.
func (r *Repo) OfferHistoryPage(
	ctx context.Context,
	after *EventPageCursor,
	limit int,
) ([]OfferHistoryRecord, bool, error) {
	return r.offerHistoryPage(ctx, "rwalk offer history page", nil, after, limit)
}

// UserOffersPage returns at most limit offers the wallet made or filled,
// newest first by immutable event-log ID.
func (r *Repo) UserOffersPage(
	ctx context.Context,
	userAid int64,
	after *EventPageCursor,
	limit int,
) ([]OfferHistoryRecord, bool, error) {
	return r.offerHistoryPage(ctx, "rwalk user offers page", &userAid, after, limit)
}

// TradeRecord is one completed purchase: the rw_item_bought event joined to
// the offer it filled. ProfitWei is the signed tracked seller profit (""
// when untracked).
type TradeRecord struct {
	Tx         EventTx
	OfferID    int64
	OfferType  int16
	TokenID    int64
	PriceWei   string
	BuyerAid   int64
	BuyerAddr  string
	SellerAid  int64
	SellerAddr string
	ProfitWei  string
}

// TradesPage returns at most limit completed purchases, newest first by the
// purchase event's immutable event-log ID.
func (r *Repo) TradesPage(
	ctx context.Context,
	after *EventPageCursor,
	limit int,
) (records []TradeRecord, hasMore bool, err error) {
	const op = "rwalk trades page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return nil, false, err
	}
	var marketAfter *marketstore.EventPageCursor
	if after != nil {
		marketAfter = &marketstore.EventPageCursor{EventLogID: after.EventLogID}
	}
	marketRecords, hasMore, err := r.marketplace.TradesPage(
		ctx,
		marketstore.Scope{
			MarketplaceAid: addrs.MarketPlaceAid,
			CollectionAid:  addrs.RandomWalkAid,
		},
		marketAfter,
		limit,
	)
	if err != nil {
		return nil, false, err
	}
	records = make([]TradeRecord, 0, len(marketRecords))
	for i := range marketRecords {
		record := marketRecords[i]
		records = append(records, TradeRecord{
			Tx:         randomWalkEventTx(record.Tx),
			OfferID:    record.OfferID,
			OfferType:  record.OfferType,
			TokenID:    record.TokenID,
			PriceWei:   record.PriceWei,
			BuyerAid:   record.BuyerAid,
			BuyerAddr:  record.BuyerAddr,
			SellerAid:  record.SellerAid,
			SellerAddr: record.SellerAddr,
			ProfitWei:  record.ProfitWei,
		})
	}
	return records, hasMore, nil
}

// FloorListingRecord is the cheapest currently active sell offer.
type FloorListingRecord struct {
	OfferID      int64
	TokenID      int64
	PriceWei     string
	ListedAtTs   int64
	ListedAtText string
}

// FloorPriceRecord is the live sell-side floor: the active sell-offer count
// and, when the book is non-empty, the cheapest listing.
type FloorPriceRecord struct {
	ActiveSellOfferCount int64
	Floor                *FloorListingRecord
}

// FloorPriceV2 returns the live sell-side floor. An empty order book is a
// valid result with a zero count and no floor listing.
func (r *Repo) FloorPriceV2(ctx context.Context) (FloorPriceRecord, error) {
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return FloorPriceRecord{}, err
	}
	marketRecord, err := r.marketplace.FloorPrice(ctx, marketstore.Scope{
		MarketplaceAid: addrs.MarketPlaceAid,
		CollectionAid:  addrs.RandomWalkAid,
	})
	if err != nil {
		return FloorPriceRecord{}, err
	}
	record := FloorPriceRecord{ActiveSellOfferCount: marketRecord.ActiveSellOfferCount}
	if marketRecord.Floor != nil {
		record.Floor = &FloorListingRecord{
			OfferID:      marketRecord.Floor.OfferID,
			TokenID:      marketRecord.Floor.TokenID,
			PriceWei:     marketRecord.Floor.PriceWei,
			ListedAtTs:   marketRecord.Floor.ListedAtTs,
			ListedAtText: marketRecord.Floor.ListedAtText,
		}
	}
	return record, nil
}

// =============================================================================
// USERS
// =============================================================================

// UserAddressID resolves a canonical wallet address to its internal address
// ID, or store.ErrNotFound when the indexer has never seen the address.
func (r *Repo) UserAddressID(ctx context.Context, address string) (int64, error) {
	return r.store.LookupAddressID(ctx, address)
}

// UserProfileRecord is one wallet's exact RandomWalk activity aggregates.
type UserProfileRecord struct {
	Aid              int64
	Address          string
	MintedTokenCount int64
	OwnedTokenCount  int64
	TradeCount       int64
	TradingVolumeWei string
	ProfitWei        string
	WithdrawalCount  int64
}

// UserProfileV2 returns one wallet's exact aggregates; a wallet without
// RandomWalk activity yields the zero shape. Unknown address IDs yield
// store.ErrNotFound.
func (r *Repo) UserProfileV2(ctx context.Context, userAid int64) (UserProfileRecord, error) {
	const op = "rwalk user profile v2"
	if userAid < 1 {
		return UserProfileRecord{}, fmt.Errorf("%s: invalid address id", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return UserProfileRecord{}, err
	}
	// Counts come from the ledgers, not the legacy accumulators:
	// rw_user_stats.total_withdrawals is dead (nothing ever writes it) and
	// total_num_toks silently drops a wallet's mints made before its first
	// trade (the mint trigger has no insert-if-missing fallback). The
	// trade/volume/profit accumulators are trigger-maintained correctly.
	record := UserProfileRecord{Aid: userAid}
	err = r.q(ctx).QueryRow(ctx, `SELECT
			a.addr,
			(SELECT COUNT(*) FROM rw_mint_evt m
				WHERE m.owner_aid=a.address_id AND m.contract_aid=$2),
			COALESCE(us.total_num_trades, 0),
			COALESCE(us.total_vol, 0)::TEXT,
			COALESCE(us.total_profit, 0)::TEXT,
			(SELECT COUNT(*) FROM rw_withdrawal w
				WHERE w.aid=a.address_id AND w.contract_aid=$2),
			(SELECT COUNT(*) FROM rw_token tk
				WHERE tk.cur_owner_aid=a.address_id AND tk.rwalk_aid=$2)
		FROM address a
			LEFT JOIN rw_user_stats us
				ON us.user_aid=a.address_id AND us.rwalk_aid=$2
		WHERE a.address_id=$1`, userAid, addrs.RandomWalkAid).Scan(
		&record.Address,
		&record.MintedTokenCount,
		&record.TradeCount,
		&record.TradingVolumeWei,
		&record.ProfitWei,
		&record.WithdrawalCount,
		&record.OwnedTokenCount,
	)
	if err != nil {
		return UserProfileRecord{}, store.WrapError(op, err)
	}
	return record, nil
}

// OwnedTokenRecord is one token a wallet currently owns: live trading state
// plus mint provenance when the mint event is indexed.
type OwnedTokenRecord struct {
	TokenID          int64
	Seed             string
	SeedNum          string
	TokenName        string
	LastPriceWei     string
	TradeCount       int64
	TradingVolumeWei string
	HasMint          bool
	MintTs           int64
	MintText         string
	MintPriceWei     string
}

// UserTokensPage returns at most limit tokens the wallet currently owns in
// ascending token order. Ownership is live: rows move between wallets on
// transfer.
func (r *Repo) UserTokensPage(
	ctx context.Context,
	userAid int64,
	after *TokenPageCursor,
	limit int,
) (records []OwnedTokenRecord, hasMore bool, err error) {
	const op = "rwalk user tokens page"
	if userAid < 1 || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid address id or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return nil, false, err
	}

	query := `SELECT
			tk.token_id,
			COALESCE(tk.seed_hex, ''),
			COALESCE(tk.seed_num, 0)::TEXT,
			COALESCE(tk.last_name, ''),
			tk.last_price::TEXT,
			tk.num_trades,
			tk.total_vol::TEXT,
			m.id IS NOT NULL,
			COALESCE(EXTRACT(EPOCH FROM m.time_stamp)::BIGINT, 0),
			m.time_stamp,
			COALESCE(m.price, 0)::TEXT
		FROM rw_token tk
			LEFT JOIN rw_mint_evt m
				ON m.contract_aid=tk.rwalk_aid AND m.token_id=tk.token_id
		WHERE tk.cur_owner_aid=$1 AND tk.rwalk_aid=$2`
	args := []any{userAid, addrs.RandomWalkAid}
	if after != nil {
		args = append(args, after.TokenID)
		query += fmt.Sprintf(" AND tk.token_id > $%d", len(args))
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
		ORDER BY tk.token_id
		LIMIT $%d`, len(args))

	scan := func(rows pgx.Rows, record *OwnedTokenRecord) error {
		return rows.Scan(
			&record.TokenID,
			&record.Seed,
			&record.SeedNum,
			&record.TokenName,
			&record.LastPriceWei,
			&record.TradeCount,
			&record.TradingVolumeWei,
			&record.HasMint,
			&record.MintTs,
			store.NullTimeText(&record.MintText),
			&record.MintPriceWei,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// =============================================================================
// STATISTICS
// =============================================================================

// LastMintRecord is the most recent mint of the collection.
type LastMintRecord struct {
	TokenID    int64
	PriceWei   string
	MintTs     int64
	MintText   string
	MinterAid  int64
	MinterAddr string
}

// LatestWithdrawalRecord is the most recent mint-pool withdrawal.
type LatestWithdrawalRecord struct {
	AmountWei      string
	OccurredTs     int64
	OccurredText   string
	WithdrawerAid  int64
	WithdrawerAddr string
	TokenID        int64
}

// StatisticsRecord is one consistent snapshot of the exact collection,
// marketplace and withdrawal aggregates.
type StatisticsRecord struct {
	MintedCount            int64
	UniqueOwnerCount       int64
	TokenTradeCount        int64
	TokenTradingVolumeWei  string
	MintFundsWei           string
	LastMint               *LastMintRecord
	MarketTradeCount       int64
	MarketTradingVolumeWei string
	ActiveSellOfferCount   int64
	ActiveBuyOfferCount    int64
	WithdrawalCount        int64
	LatestWithdrawal       *LatestWithdrawalRecord
}

// StatisticsV2 returns the global RandomWalk statistics snapshot in one
// query. Missing aggregate rows contribute zeros, matching a freshly
// migrated database.
func (r *Repo) StatisticsV2(ctx context.Context) (StatisticsRecord, error) {
	const op = "rwalk statistics v2"
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return StatisticsRecord{}, err
	}

	record := StatisticsRecord{
		TokenTradingVolumeWei:  "0",
		MintFundsWei:           "0",
		MarketTradingVolumeWei: "0",
	}
	var (
		tokenVolume, mintFunds, marketVolume sql.NullString
		tokenTrades, mintedCount             sql.NullInt64
		withdrawalCount                      sql.NullInt64
		marketTrades                         sql.NullInt64
		lastMintToken, lastMintAid           sql.NullInt64
		lastMintTs                           sql.NullInt64
		lastMintPrice, lastMintAddr          sql.NullString
		lastMintText                         string
		lwAmount, lwAddr                     sql.NullString
		lwTs, lwAid, lwToken                 sql.NullInt64
		lwText                               string
	)
	// rw_stats.total_withdrawals is a dead accumulator (nothing ever writes
	// it), so the withdrawal count comes from the ledger itself.
	err = r.q(ctx).QueryRow(ctx, `SELECT
			s.total_num_toks,
			s.total_num_trades,
			s.total_vol::TEXT,
			s.money_accumulated::TEXT,
			(SELECT COUNT(*) FROM rw_withdrawal w WHERE w.contract_aid=$1),
			(SELECT COUNT(DISTINCT tk.cur_owner_aid) FROM rw_token tk WHERE tk.rwalk_aid=$1),
			ms.total_num_trades,
			ms.total_vol::TEXT,
			(SELECT COUNT(*) FROM rw_new_offer o
				WHERE o.active AND o.otype=1 AND o.contract_aid=$2 AND o.rwalk_aid=$1),
			(SELECT COUNT(*) FROM rw_new_offer o
				WHERE o.active AND o.otype=0 AND o.contract_aid=$2 AND o.rwalk_aid=$1),
			lm.token_id,
			lm.price::TEXT,
			EXTRACT(EPOCH FROM lm.time_stamp)::BIGINT,
			lm.time_stamp,
			lm.owner_aid,
			lma.addr,
			lw.amount::TEXT,
			EXTRACT(EPOCH FROM lw.time_stamp)::BIGINT,
			lw.time_stamp,
			lw.aid,
			lwa.addr,
			lw.token_id
		FROM (SELECT 1) one
			LEFT JOIN rw_stats s ON s.rwalk_aid=$1
			LEFT JOIN LATERAL (
				SELECT COUNT(*)::BIGINT AS total_num_trades,
					COALESCE(SUM(o.price), 0) AS total_vol
				FROM rw_item_bought ib
					INNER JOIN rw_new_offer o
						ON o.contract_aid=ib.contract_aid AND o.offer_id=ib.offer_id
				WHERE ib.contract_aid=$2 AND o.rwalk_aid=$1
			) ms ON TRUE
			LEFT JOIN LATERAL (SELECT * FROM rw_mint_evt m
				WHERE m.contract_aid=$1 ORDER BY m.evtlog_id DESC LIMIT 1) lm ON TRUE
			LEFT JOIN address lma ON lma.address_id=lm.owner_aid
			LEFT JOIN LATERAL (SELECT * FROM rw_withdrawal w
				WHERE w.contract_aid=$1 ORDER BY w.evtlog_id DESC LIMIT 1) lw ON TRUE
			LEFT JOIN address lwa ON lwa.address_id=lw.aid`,
		addrs.RandomWalkAid, addrs.MarketPlaceAid).Scan(
		&mintedCount,
		&tokenTrades,
		&tokenVolume,
		&mintFunds,
		&withdrawalCount,
		&record.UniqueOwnerCount,
		&marketTrades,
		&marketVolume,
		&record.ActiveSellOfferCount,
		&record.ActiveBuyOfferCount,
		&lastMintToken,
		&lastMintPrice,
		&lastMintTs,
		store.NullTimeText(&lastMintText),
		&lastMintAid,
		&lastMintAddr,
		&lwAmount,
		&lwTs,
		store.NullTimeText(&lwText),
		&lwAid,
		&lwAddr,
		&lwToken,
	)
	if err != nil {
		return StatisticsRecord{}, store.WrapError(op, err)
	}
	record.MintedCount = mintedCount.Int64
	record.TokenTradeCount = tokenTrades.Int64
	if tokenVolume.Valid {
		record.TokenTradingVolumeWei = tokenVolume.String
	}
	if mintFunds.Valid {
		record.MintFundsWei = mintFunds.String
	}
	record.WithdrawalCount = withdrawalCount.Int64
	record.MarketTradeCount = marketTrades.Int64
	if marketVolume.Valid {
		record.MarketTradingVolumeWei = marketVolume.String
	}
	if lastMintToken.Valid {
		record.LastMint = &LastMintRecord{
			TokenID:    lastMintToken.Int64,
			PriceWei:   lastMintPrice.String,
			MintTs:     lastMintTs.Int64,
			MintText:   lastMintText,
			MinterAid:  lastMintAid.Int64,
			MinterAddr: lastMintAddr.String,
		}
	}
	if lwAmount.Valid {
		record.LatestWithdrawal = &LatestWithdrawalRecord{
			AmountWei:      lwAmount.String,
			OccurredTs:     lwTs.Int64,
			OccurredText:   lwText,
			WithdrawerAid:  lwAid.Int64,
			WithdrawerAddr: lwAddr.String,
			TokenID:        lwToken.Int64,
		}
	}
	return record, nil
}

// VolumeBucketRecord is one zero-filled bucket of completed purchases.
type VolumeBucketRecord struct {
	BucketStart int64
	TradeCount  int64
	VolumeWei   string
}

// TradingVolumeSeries buckets purchase volume into intervalSecs-second
// buckets anchored at initTs across [initTs, finTs). Buckets without trades
// are zero-filled; baseVolumeWei is the exact volume traded before initTs.
// The caller bounds the bucket count before calling.
func (r *Repo) TradingVolumeSeries(
	ctx context.Context,
	initTs, finTs, intervalSecs int,
) (baseVolumeWei string, buckets []VolumeBucketRecord, err error) {
	const op = "rwalk trading volume series"
	if initTs < 0 || finTs <= initTs || intervalSecs < 1 {
		return "", nil, fmt.Errorf("%s: invalid window", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return "", nil, err
	}

	err = r.q(ctx).QueryRow(ctx, `SELECT COALESCE(SUM(o.price), 0)::TEXT
		FROM rw_item_bought ib
			INNER JOIN rw_new_offer o
				ON o.contract_aid=ib.contract_aid AND o.offer_id=ib.offer_id
		WHERE ib.contract_aid=$1 AND o.rwalk_aid=$2 AND ib.time_stamp < TO_TIMESTAMP($3)`,
		addrs.MarketPlaceAid, addrs.RandomWalkAid, initTs).Scan(&baseVolumeWei)
	if err != nil {
		return "", nil, store.WrapError(op+": base volume", err)
	}

	query := `WITH periods AS (
			SELECT generate_series AS start_ts
			FROM generate_series(
				TO_TIMESTAMP($1), TO_TIMESTAMP($2::bigint - 1), ($3 || ' seconds')::interval
			) AS generate_series
		), bucketed AS (
			SELECT DATE_BIN(($3 || ' seconds')::interval, ib.time_stamp, TO_TIMESTAMP($1)) AS start_ts,
				COUNT(*)::BIGINT AS trade_count,
				SUM(o.price) AS volume
			FROM rw_item_bought ib
				INNER JOIN rw_new_offer o
					ON o.contract_aid=ib.contract_aid AND o.offer_id=ib.offer_id
			WHERE ib.contract_aid=$4 AND o.rwalk_aid=$5
				AND ib.time_stamp >= TO_TIMESTAMP($1) AND ib.time_stamp < TO_TIMESTAMP($2)
			GROUP BY 1
		)
		SELECT
			FLOOR(EXTRACT(EPOCH FROM p.start_ts))::BIGINT,
			COALESCE(b.trade_count, 0)::BIGINT,
			COALESCE(b.volume, 0)::TEXT
		FROM periods p
			LEFT JOIN bucketed b ON b.start_ts = p.start_ts
		ORDER BY p.start_ts`
	scan := func(rows pgx.Rows, record *VolumeBucketRecord) error {
		return rows.Scan(&record.BucketStart, &record.TradeCount, &record.VolumeWei)
	}
	buckets, err = queryList(ctx, r, op, 256, query, scan,
		initTs, finTs, strconv.Itoa(intervalSecs), addrs.MarketPlaceAid, addrs.RandomWalkAid)
	if err != nil {
		return "", nil, err
	}
	return baseVolumeWei, buckets, nil
}

// FloorPointRecord is the cheapest sell listing created inside one bucket.
type FloorPointRecord struct {
	BucketStart int64
	FloorWei    string
}

// ListingFloorSeries returns the minimum sell-listing price per
// intervalSecs-second bucket anchored at initTs across [initTs, finTs).
// Buckets without new sell listings are omitted. The caller bounds the
// bucket count before calling.
func (r *Repo) ListingFloorSeries(
	ctx context.Context,
	initTs, finTs, intervalSecs int,
) ([]FloorPointRecord, error) {
	const op = "rwalk listing floor series"
	if initTs < 0 || finTs <= initTs || intervalSecs < 1 {
		return nil, fmt.Errorf("%s: invalid window", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return nil, err
	}
	query := `SELECT
			FLOOR(EXTRACT(EPOCH FROM
				DATE_BIN(($3 || ' seconds')::interval, o.time_stamp, TO_TIMESTAMP($1))
			))::BIGINT AS bucket_ts,
			MIN(o.price)::TEXT
		FROM rw_new_offer o
		WHERE o.contract_aid=$4 AND o.rwalk_aid=$5 AND o.otype=1
			AND o.time_stamp >= TO_TIMESTAMP($1) AND o.time_stamp < TO_TIMESTAMP($2)
		GROUP BY 1
		ORDER BY 1`
	scan := func(rows pgx.Rows, record *FloorPointRecord) error {
		return rows.Scan(&record.BucketStart, &record.FloorWei)
	}
	return queryList(ctx, r, op, 64, query, scan,
		initTs, finTs, strconv.Itoa(intervalSecs), addrs.MarketPlaceAid, addrs.RandomWalkAid)
}

// MonthlyMintRecord is one calendar month's exact mint aggregates.
type MonthlyMintRecord struct {
	Year      int64
	Month     int64
	MintCount int64
	MintedWei string
}

// MintReportV2 aggregates every mint per UTC calendar month, oldest first.
// The session time zone is pinned to UTC by the store pool.
func (r *Repo) MintReportV2(ctx context.Context) ([]MonthlyMintRecord, error) {
	const op = "rwalk mint report v2"
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return nil, err
	}
	query := `SELECT
			EXTRACT(YEAR FROM m.time_stamp)::BIGINT,
			EXTRACT(MONTH FROM m.time_stamp)::BIGINT,
			COUNT(*)::BIGINT,
			SUM(m.price)::TEXT
		FROM rw_mint_evt m
		WHERE m.contract_aid=$1
		GROUP BY 1, 2
		ORDER BY 1, 2`
	scan := func(rows pgx.Rows, record *MonthlyMintRecord) error {
		return rows.Scan(&record.Year, &record.Month, &record.MintCount, &record.MintedWei)
	}
	return queryList(ctx, r, op, 32, query, scan, addrs.RandomWalkAid)
}

// WithdrawalRecord is one exact mint-pool withdrawal.
type WithdrawalRecord struct {
	Tx             EventTx
	WithdrawerAid  int64
	WithdrawerAddr string
	TokenID        int64
	AmountWei      string
}

// WithdrawalsPage returns at most limit withdrawals, newest first by
// immutable event-log ID.
func (r *Repo) WithdrawalsPage(
	ctx context.Context,
	after *EventPageCursor,
	limit int,
) (records []WithdrawalRecord, hasMore bool, err error) {
	const op = "rwalk withdrawals page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	addrs, err := r.ContractAddrs(ctx)
	if err != nil {
		return nil, false, err
	}
	query := `SELECT
			w.evtlog_id,
			w.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM w.time_stamp)::BIGINT,
			w.time_stamp,
			w.aid,
			wa.addr,
			w.token_id,
			COALESCE(w.amount, 0)::TEXT
		FROM rw_withdrawal w
			INNER JOIN transaction t ON t.id=w.tx_id
			INNER JOIN address wa ON wa.address_id=w.aid
		WHERE w.contract_aid=$1`
	args := []any{addrs.RandomWalkAid}
	if after != nil {
		args = append(args, after.EventLogID)
		query += fmt.Sprintf(" AND w.evtlog_id < $%d", len(args))
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
		ORDER BY w.evtlog_id DESC
		LIMIT $%d`, len(args))
	scan := func(rows pgx.Rows, record *WithdrawalRecord) error {
		return rows.Scan(
			&record.Tx.EvtLogID,
			&record.Tx.BlockNum,
			&record.Tx.TxID,
			&record.Tx.TxHash,
			&record.Tx.TimeStamp,
			store.TimeText(&record.Tx.DateTime),
			&record.WithdrawerAid,
			&record.WithdrawerAddr,
			&record.TokenID,
			&record.AmountWei,
		)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}
