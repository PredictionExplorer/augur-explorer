// Package marketplace provides collection-scoped reads over the shared
// RandomWalk marketplace event tables. A Scope always identifies both the
// marketplace contract and the traded NFT collection, preventing one
// collection's offers from leaking into another collection's API.
package marketplace

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// Scope identifies one NFT collection on one marketplace contract.
type Scope struct {
	MarketplaceAid int64
	CollectionAid  int64
}

func (s Scope) valid() bool {
	return s.MarketplaceAid > 0 && s.CollectionAid > 0
}

// Repo executes collection-scoped marketplace queries on the shared store.
type Repo struct {
	store *store.Store
}

// NewRepo returns a marketplace repository backed by st.
func NewRepo(st *store.Store) *Repo {
	return &Repo{store: st}
}

func (r *Repo) q(ctx context.Context) store.Querier {
	return r.store.Querier(ctx)
}

// ResolveScope resolves deployment addresses into the database-local address
// IDs used by the marketplace tables.
func (r *Repo) ResolveScope(
	ctx context.Context,
	marketplaceAddress string,
	collectionAddress string,
) (Scope, error) {
	marketplaceAid, err := r.store.LookupAddressID(ctx, marketplaceAddress)
	if err != nil {
		return Scope{}, fmt.Errorf("resolve marketplace address: %w", err)
	}
	collectionAid, err := r.store.LookupAddressID(ctx, collectionAddress)
	if err != nil {
		return Scope{}, fmt.Errorf("resolve marketplace collection: %w", err)
	}
	return Scope{MarketplaceAid: marketplaceAid, CollectionAid: collectionAid}, nil
}

// LegacyOffer is the complete v1 offer row shared by the RandomWalk
// compatibility surface and the Cosmic Signature compatibility projection.
type LegacyOffer struct {
	ID              int64
	EventLogID      int64
	BlockNumber     int64
	TransactionID   int64
	TransactionHash string
	TimeStamp       int64
	DateTime        string
	LegacyDateTime  string
	OfferID         int64
	OfferType       int
	SellerAid       int64
	SellerAddress   string
	BuyerAid        int64
	BuyerAddress    string
	TokenID         int64
	Active          bool
	PriceETH        float64
	ProfitETH       *float64
	MarketplaceAid  int64
	Marketplace     string
	CollectionAid   int64
	Collection      string
	WasCanceled     bool
}

func legacyOrderClause(orderBy int) string {
	switch orderBy {
	case 1:
		return " ORDER BY o.price DESC"
	case 2:
		return " ORDER BY o.price ASC"
	default:
		return " ORDER BY o.id"
	}
}

// ActiveOffersLegacy returns every active offer in the scope using the v1
// numeric ordering selector (1 price descending, 2 ascending, otherwise
// insertion order).
func (r *Repo) ActiveOffersLegacy(
	ctx context.Context,
	scope Scope,
	orderBy int,
) ([]LegacyOffer, error) {
	const op = "marketplace active offers"
	if !scope.valid() {
		return nil, fmt.Errorf("%s: invalid scope", op)
	}
	query := `SELECT
			o.id,
			o.evtlog_id,
			o.block_num,
			o.tx_id,
			tx.tx_hash,
			EXTRACT(EPOCH FROM o.time_stamp)::BIGINT,
			o.time_stamp,
			o.time_stamp::TEXT,
			o.offer_id,
			o.otype,
			o.seller_aid,
			sa.addr,
			o.buyer_aid,
			ba.addr,
			o.token_id,
			o.active,
			o.price/1e+18,
			o.rwalk_aid
		FROM rw_new_offer o
			INNER JOIN transaction tx ON tx.id=o.tx_id
			INNER JOIN address sa ON sa.address_id=o.seller_aid
			INNER JOIN address ba ON ba.address_id=o.buyer_aid
		WHERE o.active AND o.contract_aid=$1 AND o.rwalk_aid=$2` +
		legacyOrderClause(orderBy)
	scan := func(rows pgx.Rows, record *LegacyOffer) error {
		return rows.Scan(
			&record.ID,
			&record.EventLogID,
			&record.BlockNumber,
			&record.TransactionID,
			&record.TransactionHash,
			&record.TimeStamp,
			store.TimeText(&record.DateTime),
			&record.LegacyDateTime,
			&record.OfferID,
			&record.OfferType,
			&record.SellerAid,
			&record.SellerAddress,
			&record.BuyerAid,
			&record.BuyerAddress,
			&record.TokenID,
			&record.Active,
			&record.PriceETH,
			&record.CollectionAid,
		)
	}
	return store.QueryList(ctx, r.q(ctx), op, 16, query, scan,
		scope.MarketplaceAid, scope.CollectionAid)
}

// SaleHistoryLegacy returns completed, non-canceled sales in the scope,
// oldest first by purchase time, with offset pagination.
func (r *Repo) SaleHistoryLegacy(
	ctx context.Context,
	scope Scope,
	offset int,
	limit int,
) ([]LegacyOffer, error) {
	const op = "marketplace sale history"
	if !scope.valid() || offset < 0 || limit < 0 {
		return nil, fmt.Errorf("%s: invalid scope or bounds", op)
	}
	query := `SELECT
			o.id,
			o.evtlog_id,
			o.block_num,
			ib.tx_id,
			tx.tx_hash,
			EXTRACT(EPOCH FROM ib.time_stamp)::BIGINT,
			ib.time_stamp,
			ib.time_stamp::TEXT,
			o.offer_id,
			o.otype,
			o.seller_aid,
			sa.addr,
			o.buyer_aid,
			ba.addr,
			o.token_id,
			o.active,
			can.id,
			o.price/1e+18,
			o.profit/1e+18,
			o.contract_aid,
			ca.addr,
			o.rwalk_aid,
			rwa.addr
		FROM rw_new_offer o
			INNER JOIN rw_item_bought ib
				ON ib.contract_aid=o.contract_aid AND ib.offer_id=o.offer_id
			INNER JOIN transaction tx ON tx.id=ib.tx_id
			INNER JOIN address sa ON sa.address_id=o.seller_aid
			INNER JOIN address ba ON ba.address_id=o.buyer_aid
			INNER JOIN address ca ON ca.address_id=o.contract_aid
			INNER JOIN address rwa ON rwa.address_id=o.rwalk_aid
			LEFT JOIN rw_offer_canceled can
				ON can.contract_aid=o.contract_aid AND can.offer_id=o.offer_id
		WHERE NOT o.active AND o.contract_aid=$3 AND o.rwalk_aid=$4 AND can.id IS NULL
		ORDER BY ib.time_stamp, o.id
		OFFSET $1 LIMIT $2`
	scan := func(rows pgx.Rows, record *LegacyOffer) error {
		var canceledID sql.NullInt64
		var profit sql.NullFloat64
		if err := rows.Scan(
			&record.ID,
			&record.EventLogID,
			&record.BlockNumber,
			&record.TransactionID,
			&record.TransactionHash,
			&record.TimeStamp,
			store.TimeText(&record.DateTime),
			&record.LegacyDateTime,
			&record.OfferID,
			&record.OfferType,
			&record.SellerAid,
			&record.SellerAddress,
			&record.BuyerAid,
			&record.BuyerAddress,
			&record.TokenID,
			&record.Active,
			&canceledID,
			&record.PriceETH,
			&profit,
			&record.MarketplaceAid,
			&record.Marketplace,
			&record.CollectionAid,
			&record.Collection,
		); err != nil {
			return err
		}
		record.WasCanceled = canceledID.Valid
		if profit.Valid {
			value := profit.Float64
			record.ProfitETH = &value
		}
		return nil
	}
	return store.QueryList(ctx, r.q(ctx), op, 16, query, scan,
		offset, limit, scope.MarketplaceAid, scope.CollectionAid)
}

// FloorPriceETH returns the cheapest active sell offer in ETH. noOffers is
// true for an empty order book.
func (r *Repo) FloorPriceETH(
	ctx context.Context,
	scope Scope,
) (noOffers bool, floorPrice float64, offerID int64, tokenID int64, err error) {
	const op = "marketplace floor price"
	if !scope.valid() {
		return false, 0, 0, 0, fmt.Errorf("%s: invalid scope", op)
	}
	var nullableFloor sql.NullFloat64
	var nullableOfferID, nullableTokenID sql.NullInt64
	err = r.q(ctx).QueryRow(ctx, `SELECT
			o.price/1e+18,
			o.offer_id,
			o.token_id
		FROM rw_new_offer o
		WHERE o.active AND o.otype=1 AND o.contract_aid=$1 AND o.rwalk_aid=$2
		ORDER BY o.price, o.evtlog_id
		LIMIT 1`, scope.MarketplaceAid, scope.CollectionAid).Scan(
		&nullableFloor,
		&nullableOfferID,
		&nullableTokenID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return true, 0, 0, 0, nil
		}
		return false, 0, 0, 0, store.WrapError(op, err)
	}
	return false, nullableFloor.Float64, nullableOfferID.Int64, nullableTokenID.Int64, nil
}

// EventTx identifies one immutable indexed marketplace event.
type EventTx struct {
	EventLogID int64
	BlockNum   int64
	TxID       int64
	TxHash     string
	TimeStamp  int64
	DateTime   string
}

// OfferSort selects the live order-book order.
type OfferSort string

const (
	// OfferSortNewest orders by descending event-log ID.
	OfferSortNewest OfferSort = "newest"
	// OfferSortOldest orders by ascending event-log ID.
	OfferSortOldest OfferSort = "oldest"
	// OfferSortPriceAsc orders by exact price then event-log ID.
	OfferSortPriceAsc OfferSort = "priceAsc"
	// OfferSortPriceDesc orders by descending exact price then event-log ID.
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

// OfferPageCursor identifies the last live offer returned.
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

// EventPageCursor identifies the last event in a newest-first ledger page.
type EventPageCursor struct {
	EventLogID int64
}

func (c *EventPageCursor) valid() bool {
	return c == nil || c.EventLogID >= 1
}

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

// ActiveOffersPage returns one bounded keyset page of active offers.
func (r *Repo) ActiveOffersPage(
	ctx context.Context,
	scope Scope,
	sort OfferSort,
	after *OfferPageCursor,
	limit int,
) (records []OfferRecord, hasMore bool, err error) {
	const op = "marketplace active offers page"
	if !scope.valid() || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid scope or limit", op)
	}
	if !sort.valid() {
		return nil, false, fmt.Errorf("%s: invalid sort", op)
	}
	if !after.validFor(sort) {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}

	query := `SELECT
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
			mka.addr
		FROM rw_new_offer o
			INNER JOIN transaction t ON t.id=o.tx_id
			INNER JOIN address mka ON mka.address_id=
				CASE WHEN o.otype=1 THEN o.seller_aid ELSE o.buyer_aid END
		WHERE o.active AND o.contract_aid=$1 AND o.rwalk_aid=$2`
	args := []any{scope.MarketplaceAid, scope.CollectionAid}
	var orderBy string
	switch sort {
	case OfferSortNewest:
		orderBy = "ORDER BY o.evtlog_id DESC"
		if after != nil {
			args = append(args, after.EventLogID)
			query += fmt.Sprintf(" AND o.evtlog_id < $%d", len(args))
		}
	case OfferSortOldest:
		orderBy = "ORDER BY o.evtlog_id"
		if after != nil {
			args = append(args, after.EventLogID)
			query += fmt.Sprintf(" AND o.evtlog_id > $%d", len(args))
		}
	case OfferSortPriceAsc:
		orderBy = "ORDER BY o.price, o.evtlog_id"
		if after != nil {
			args = append(args, after.PriceWei, after.EventLogID)
			query += fmt.Sprintf(
				" AND (o.price > $%d::NUMERIC OR (o.price = $%d::NUMERIC AND o.evtlog_id > $%d))",
				len(args)-1, len(args)-1, len(args))
		}
	case OfferSortPriceDesc:
		orderBy = "ORDER BY o.price DESC, o.evtlog_id"
		if after != nil {
			args = append(args, after.PriceWei, after.EventLogID)
			query += fmt.Sprintf(
				" AND (o.price < $%d::NUMERIC OR (o.price = $%d::NUMERIC AND o.evtlog_id > $%d))",
				len(args)-1, len(args)-1, len(args))
		}
	}
	args = append(args, limit+1)
	query += fmt.Sprintf("\n\t\t%s\n\t\tLIMIT $%d", orderBy, len(args))

	scan := func(rows pgx.Rows, record *OfferRecord) error {
		return rows.Scan(
			&record.ListTx.EventLogID,
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
		)
	}
	records, err = store.QueryList(ctx, r.q(ctx), op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		return records[:limit], true, nil
	}
	return records, false, nil
}

// OfferOutcomePurchase is the purchase event that filled an offer.
type OfferOutcomePurchase struct {
	Tx         EventTx
	BuyerAid   int64
	BuyerAddr  string
	SellerAid  int64
	SellerAddr string
}

// OfferHistoryRecord is one offer creation plus its current outcome.
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
		&record.ListTx.EventLogID,
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
	record.ProfitWei = profit.String
	if boughtEvt.Valid {
		record.Purchase = &OfferOutcomePurchase{
			Tx: EventTx{
				EventLogID: boughtEvt.Int64,
				BlockNum:   boughtBlock.Int64,
				TxID:       boughtTxID.Int64,
				TxHash:     boughtHash.String,
				TimeStamp:  boughtTs.Int64,
				DateTime:   boughtText,
			},
			BuyerAid:   buyerAid.Int64,
			BuyerAddr:  buyerAddr.String,
			SellerAid:  sellerAid.Int64,
			SellerAddr: sellerAddr.String,
		}
	}
	if cancelEvt.Valid {
		record.Cancellation = &EventTx{
			EventLogID: cancelEvt.Int64,
			BlockNum:   cancelBlock.Int64,
			TxID:       cancelTxID.Int64,
			TxHash:     cancelHash.String,
			TimeStamp:  cancelTs.Int64,
			DateTime:   cancelText,
		}
	}
	return nil
}

// OfferHistoryPage returns one bounded newest-first offer ledger page.
func (r *Repo) OfferHistoryPage(
	ctx context.Context,
	scope Scope,
	after *EventPageCursor,
	limit int,
) (records []OfferHistoryRecord, hasMore bool, err error) {
	const op = "marketplace offer history page"
	if !scope.valid() || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid scope or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	query := offerHistorySelectSQL + `
	WHERE o.contract_aid=$1 AND o.rwalk_aid=$2`
	args := []any{scope.MarketplaceAid, scope.CollectionAid}
	if after != nil {
		args = append(args, after.EventLogID)
		query += fmt.Sprintf(" AND o.evtlog_id < $%d", len(args))
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
	ORDER BY o.evtlog_id DESC
	LIMIT $%d`, len(args))
	records, err = store.QueryList(ctx, r.q(ctx), op, limit+1, query, scanOfferHistory, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		return records[:limit], true, nil
	}
	return records, false, nil
}

// TradeRecord is one completed purchase joined to the offer it filled.
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

// TradesPage returns one bounded newest-first purchase ledger page.
func (r *Repo) TradesPage(
	ctx context.Context,
	scope Scope,
	after *EventPageCursor,
	limit int,
) (records []TradeRecord, hasMore bool, err error) {
	const op = "marketplace trades page"
	if !scope.valid() || limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid scope or limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	query := `SELECT
			ib.evtlog_id,
			ib.block_num,
			t.id,
			t.tx_hash,
			EXTRACT(EPOCH FROM ib.time_stamp)::BIGINT,
			ib.time_stamp,
			ib.offer_id,
			o.otype,
			o.token_id,
			o.price::TEXT,
			ib.buyer_aid,
			ba.addr,
			ib.seller_aid,
			sa.addr,
			o.profit::TEXT
		FROM rw_item_bought ib
			INNER JOIN rw_new_offer o
				ON o.contract_aid=ib.contract_aid AND o.offer_id=ib.offer_id
			INNER JOIN transaction t ON t.id=ib.tx_id
			INNER JOIN address ba ON ba.address_id=ib.buyer_aid
			INNER JOIN address sa ON sa.address_id=ib.seller_aid
		WHERE ib.contract_aid=$1 AND o.rwalk_aid=$2`
	args := []any{scope.MarketplaceAid, scope.CollectionAid}
	if after != nil {
		args = append(args, after.EventLogID)
		query += fmt.Sprintf(" AND ib.evtlog_id < $%d", len(args))
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
		ORDER BY ib.evtlog_id DESC
		LIMIT $%d`, len(args))
	scan := func(rows pgx.Rows, record *TradeRecord) error {
		var profit sql.NullString
		if err := rows.Scan(
			&record.Tx.EventLogID,
			&record.Tx.BlockNum,
			&record.Tx.TxID,
			&record.Tx.TxHash,
			&record.Tx.TimeStamp,
			store.TimeText(&record.Tx.DateTime),
			&record.OfferID,
			&record.OfferType,
			&record.TokenID,
			&record.PriceWei,
			&record.BuyerAid,
			&record.BuyerAddr,
			&record.SellerAid,
			&record.SellerAddr,
			&profit,
		); err != nil {
			return err
		}
		record.ProfitWei = profit.String
		return nil
	}
	records, err = store.QueryList(ctx, r.q(ctx), op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	if len(records) > limit {
		return records[:limit], true, nil
	}
	return records, false, nil
}

// FloorListingRecord is the cheapest currently active sell offer.
type FloorListingRecord struct {
	OfferID      int64
	TokenID      int64
	PriceWei     string
	ListedAtTs   int64
	ListedAtText string
}

// FloorPriceRecord is the active sell count and optional cheapest listing.
type FloorPriceRecord struct {
	ActiveSellOfferCount int64
	Floor                *FloorListingRecord
}

// FloorPrice returns the current exact sell-side floor for the scope.
func (r *Repo) FloorPrice(ctx context.Context, scope Scope) (FloorPriceRecord, error) {
	const op = "marketplace floor price v2"
	if !scope.valid() {
		return FloorPriceRecord{}, fmt.Errorf("%s: invalid scope", op)
	}
	var record FloorPriceRecord
	err := r.q(ctx).QueryRow(ctx, `SELECT COUNT(*)
		FROM rw_new_offer
		WHERE active AND otype=1 AND contract_aid=$1 AND rwalk_aid=$2`,
		scope.MarketplaceAid, scope.CollectionAid).Scan(&record.ActiveSellOfferCount)
	if err != nil {
		return FloorPriceRecord{}, store.WrapError(op+": count", err)
	}
	var floor FloorListingRecord
	err = r.q(ctx).QueryRow(ctx, `SELECT
			offer_id,
			token_id,
			price::TEXT,
			EXTRACT(EPOCH FROM time_stamp)::BIGINT,
			time_stamp
		FROM rw_new_offer
		WHERE active AND otype=1 AND contract_aid=$1 AND rwalk_aid=$2
		ORDER BY price, evtlog_id
		LIMIT 1`, scope.MarketplaceAid, scope.CollectionAid).Scan(
		&floor.OfferID,
		&floor.TokenID,
		&floor.PriceWei,
		&floor.ListedAtTs,
		store.TimeText(&floor.ListedAtText),
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return record, nil
		}
		return FloorPriceRecord{}, store.WrapError(op, err)
	}
	record.Floor = &floor
	return record, nil
}
