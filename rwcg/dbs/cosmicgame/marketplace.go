package cosmicgame

import (
	"database/sql"
	"fmt"
	"os"
)

// CGMarketOffer is a marketplace offer (or completed sale) for a Cosmic
// Signature NFT, read from the shared rw_new_offer / rw_item_bought tables.
// Those tables are populated by the RandomWalk marketplace ETL, which indexes
// every NFT collection traded on the marketplace contract (the traded NFT is
// stored per-row in rw_new_offer.rwalk_aid), not just RandomWalk. Prices are
// converted to ETH.
type CGMarketOffer struct {
	OfferId    int64
	OfferType  int // 0 = buy bid, 1 = sell listing
	TokenId    int64
	SellerAddr string
	BuyerAddr  string
	Active     bool
	Price      float64
	BlockNum   int64
	TimeStamp  int64
	DateTime   string
}

// Get_marketplace_addr returns the marketplace contract address recorded in
// rw_contracts (written by the RandomWalk ETL feeding this database). Returned
// as a hex string so the caller resolves it to a per-database address id.
func (sw *SQLStorageWrapper) Get_marketplace_addr() (string, error) {
	var addr string
	err := sw.S.Db().QueryRow("SELECT marketplace_addr FROM rw_contracts LIMIT 1").Scan(&addr)
	return addr, err
}

// Get_marketplace_active_offers returns active offers (both sell listings and
// buy bids) for the given NFT collection (nft_aid) on the given marketplace
// (market_aid). order_by: 0 = by offer id, 1 = price desc, 2 = price asc.
func (sw *SQLStorageWrapper) Get_marketplace_active_offers(nft_aid int64, market_aid int64, order_by int) []CGMarketOffer {
	records := make([]CGMarketOffer, 0, 16)

	order_by_mod := " ORDER BY o.id"
	if order_by == 1 {
		order_by_mod = " ORDER BY o.price DESC"
	}
	if order_by == 2 {
		order_by_mod = " ORDER BY o.price ASC"
	}

	query := "SELECT " +
		"o.offer_id," +
		"o.otype," +
		"o.token_id," +
		"sa.addr," +
		"ba.addr," +
		"o.active," +
		"o.price/1e+18," +
		"o.block_num," +
		"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT," +
		"o.time_stamp::text " +
		"FROM rw_new_offer o " +
		"JOIN address sa ON o.seller_aid=sa.address_id " +
		"JOIN address ba ON o.buyer_aid=ba.address_id " +
		"WHERE (o.active='t') " +
		fmt.Sprintf(" AND (o.rwalk_aid=%v) ", nft_aid) +
		fmt.Sprintf(" AND (o.contract_aid=%v) ", market_aid) +
		order_by_mod

	rows, err := sw.S.Db().Query(query)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec CGMarketOffer
		err = rows.Scan(
			&rec.OfferId,
			&rec.OfferType,
			&rec.TokenId,
			&rec.SellerAddr,
			&rec.BuyerAddr,
			&rec.Active,
			&rec.Price,
			&rec.BlockNum,
			&rec.TimeStamp,
			&rec.DateTime,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v", err, query))
			os.Exit(1)
		}
		records = append(records, rec)
	}
	return records
}

// Get_marketplace_floor_price returns the lowest active sell price for a
// collection on the marketplace. no_offers is true when there are no active
// listings.
func (sw *SQLStorageWrapper) Get_marketplace_floor_price(nft_aid int64, market_aid int64) (no_offers bool, floor_price float64, offer_id int64, token_id int64, err error) {
	var n_floor sql.NullFloat64
	var n_offer, n_token sql.NullInt64
	query := "SELECT o.price/1e+18, o.offer_id, o.token_id " +
		"FROM rw_new_offer o " +
		"WHERE (o.active='t') AND (o.otype=1) " +
		fmt.Sprintf(" AND (o.rwalk_aid=%v) ", nft_aid) +
		fmt.Sprintf(" AND (o.contract_aid=%v) ", market_aid) +
		"ORDER BY o.price ASC LIMIT 1"

	res := sw.S.Db().QueryRow(query)
	err = res.Scan(&n_floor, &n_offer, &n_token)
	if err != nil {
		if err == sql.ErrNoRows {
			no_offers = true
			err = nil
			return
		}
		sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v", err, query))
		os.Exit(1)
	}
	floor_price = n_floor.Float64
	offer_id = n_offer.Int64
	token_id = n_token.Int64
	return
}

// Get_marketplace_sale_history returns completed sales (bought, not canceled)
// for a collection on the marketplace, oldest-first. Timestamps come from the
// rw_item_bought (purchase) event, not the offer creation time.
func (sw *SQLStorageWrapper) Get_marketplace_sale_history(nft_aid int64, market_aid int64, offset, limit int) []CGMarketOffer {
	records := make([]CGMarketOffer, 0, 16)

	query := "SELECT " +
		"o.offer_id," +
		"o.otype," +
		"o.token_id," +
		"sa.addr," +
		"ba.addr," +
		"o.active," +
		"o.price/1e+18," +
		"o.block_num," +
		"EXTRACT(EPOCH FROM ib.time_stamp)::BIGINT," +
		"ib.time_stamp::text " +
		"FROM rw_new_offer o " +
		"JOIN rw_item_bought ib ON ib.contract_aid=o.contract_aid AND ib.offer_id=o.offer_id " +
		"JOIN address sa ON o.seller_aid=sa.address_id " +
		"JOIN address ba ON o.buyer_aid=ba.address_id " +
		"LEFT JOIN rw_offer_canceled can ON (can.contract_aid=o.contract_aid) AND (can.offer_id=o.offer_id) " +
		"WHERE (o.active='f') AND (can.id IS NULL) " +
		fmt.Sprintf(" AND (o.rwalk_aid=%v) ", nft_aid) +
		fmt.Sprintf(" AND (o.contract_aid=%v) ", market_aid) +
		"ORDER BY ib.time_stamp ASC, o.id ASC OFFSET $1 LIMIT $2"

	rows, err := sw.S.Db().Query(query, offset, limit)
	if err != nil {
		sw.S.Log_msg(fmt.Sprintf("DB error: %v (query=%v)", err, query))
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var rec CGMarketOffer
		err = rows.Scan(
			&rec.OfferId,
			&rec.OfferType,
			&rec.TokenId,
			&rec.SellerAddr,
			&rec.BuyerAddr,
			&rec.Active,
			&rec.Price,
			&rec.BlockNum,
			&rec.TimeStamp,
			&rec.DateTime,
		)
		if err != nil {
			sw.S.Log_msg(fmt.Sprintf("DB error: %v, q=%v", err, query))
			os.Exit(1)
		}
		records = append(records, rec)
	}
	return records
}
