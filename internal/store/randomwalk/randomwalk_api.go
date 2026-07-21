// API read queries for the RandomWalk explorer: offers, mints, trading
// history, statistics and token/user detail views.

package randomwalk

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/jackc/pgx/v5"

	rwmodel "github.com/PredictionExplorer/augur-explorer/internal/model/randomwalk"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
	"github.com/PredictionExplorer/augur-explorer/internal/timefmt"
)

func profitFromNull(nf sql.NullFloat64) rwmodel.JSONNullFloat64 {
	if !nf.Valid {
		return rwmodel.JSONNullFloat64{}
	}
	p := nf.Float64
	if math.IsNaN(p) || math.IsInf(p, 0) {
		return rwmodel.JSONNullFloat64{}
	}
	return rwmodel.JSONNullFloat64{Valid: true, Value: p}
}

// =============================================================================
// OFFER QUERIES
// =============================================================================

// activeOffersOrderClause maps the numeric order_by API parameter onto a
// whitelisted ORDER BY clause (1: price high→low; 2: price low→high;
// anything else: insertion order). Request input never reaches the SQL text.
func activeOffersOrderClause(orderBy int) string {
	switch orderBy {
	case 1:
		return " ORDER BY o.price DESC"
	case 2:
		return " ORDER BY o.price ASC"
	default:
		return " ORDER BY o.id"
	}
}

// ActiveOffers returns all open marketplace offers for the RandomWalk
// contract, ordered per the whitelisted orderBy selector.
func (r *Repo) ActiveOffers(ctx context.Context, rwalkAid, marketAid int64, orderBy int) ([]rwmodel.Offer, error) {
	query := `SELECT
			o.id,
			o.evtlog_id,
			o.block_num,
			o.tx_id,
			tx.tx_hash,
			EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,
			o.time_stamp,
			o.offer_id,
			o.otype,
			o.seller_aid,
			sa.addr seller_addr,
			o.buyer_aid,
			ba.addr buyer_addr,
			o.token_id,
			o.active,
			o.price/1e+18 price,
			o.rwalk_aid
		FROM rw_new_offer o
			JOIN transaction tx ON o.tx_id=tx.id
			JOIN address sa ON o.seller_aid=sa.address_id
			JOIN address ba ON o.buyer_aid=ba.address_id
		WHERE (active = 't') AND (o.rwalk_aid=$1) AND (o.contract_aid=$2)` +
		activeOffersOrderClause(orderBy)
	return queryList(ctx, r, "active offers", 16, query, func(rows pgx.Rows, rec *rwmodel.Offer) error {
		return rows.Scan(
			&rec.Id,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			store.TimeText(&rec.DateTime),
			&rec.OfferId,
			&rec.OfferType,
			&rec.SellerAid,
			&rec.SellerAddr,
			&rec.BuyerAid,
			&rec.BuyerAddr,
			&rec.TokenId,
			&rec.Active,
			&rec.Price,
			&rec.RWalkAid,
		)
	}, rwalkAid, marketAid)
}

// =============================================================================
// TOKEN MINT QUERIES
// =============================================================================

// #nosec G101 -- SQL column list, not a credential
const mintedTokensSelect = `SELECT
			t.block_num,
			EXTRACT(EPOCH FROM t.time_stamp)::BIGINT as ts,
			t.time_stamp,
			t.contract_aid,
			ca.addr,
			t.token_id,
			t.owner_aid,
			oa.addr minter_addr,
			seed seed_hex,
			seed_num,
			price/1e+18,
			tx.tx_hash
		FROM rw_mint_evt t
			LEFT JOIN address ca ON t.contract_aid=ca.address_id
			LEFT JOIN address oa ON t.owner_aid=oa.address_id
			LEFT JOIN transaction tx ON t.tx_id=tx.id `

func scanMintedToken(rows pgx.Rows, rec *rwmodel.TokenMint) error {
	return rows.Scan(
		&rec.BlockNum,
		&rec.TimeStamp,
		store.TimeText(&rec.DateTime),
		&rec.ContractAid,
		&rec.ContractAddr,
		&rec.TokenId,
		&rec.MinterAid,
		&rec.MinterAddr,
		&rec.Seed,
		&rec.SeedNum,
		&rec.Price,
		&rec.TxHash,
	)
}

// MintedTokensByPeriod returns mints inside [iniTs, finTs).
func (r *Repo) MintedTokensByPeriod(ctx context.Context, rwalkAid int64, iniTs, finTs int) ([]rwmodel.TokenMint, error) {
	query := mintedTokensSelect +
		`WHERE (t.time_stamp >= TO_TIMESTAMP($1)) AND (t.time_stamp<TO_TIMESTAMP($2)) AND t.contract_aid=$3`
	return queryList(ctx, r, "minted tokens by period", 32, query, scanMintedToken, iniTs, finTs, rwalkAid)
}

// MintedTokensSequentially returns mints newest first, paginated.
func (r *Repo) MintedTokensSequentially(ctx context.Context, rwalkAid int64, offset, limit int) ([]rwmodel.TokenMint, error) {
	query := mintedTokensSelect +
		`WHERE contract_aid=$1 ORDER by t.id DESC OFFSET $2 LIMIT $3`
	return queryList(ctx, r, "minted tokens sequentially", 32, query, scanMintedToken, rwalkAid, offset, limit)
}

// =============================================================================
// TRADING HISTORY
// =============================================================================

// TradingHistory returns the merged offer/sale/cancel timeline of the
// marketplace, ordered by the real (most specific) event timestamp.
func (r *Repo) TradingHistory(ctx context.Context, contractAid int64, offset, limit int) ([]rwmodel.TradingHistoryLog, error) {
	query := "SELECT " +
		"record_id," + "evtlog_id," + "block_num," + "tx_id, " + "offer_ts," + "offer_date," +
		"offer_id," + "otype," + "seller_aid," + "seller_addr," + "buyer_aid," + "buyer_addr," +
		"token_id," + "active," + "cancel_id," + "price, " + "profit, " + "contract_aid," + "contract_addr," +
		"rwalk_aid," + "rwalk_addr, " + "itembought_ts," + "itembought_date, " +
		"canceled_ts," + "canceled_date, " + "real_ts, " + "real_date " +
		"FROM (" +
		"(" +
		"SELECT " +
		"o.id AS record_id," +
		"o.evtlog_id AS evtlog_id," +
		"o.block_num AS block_num," +
		"o.tx_id AS tx_id, " +
		"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as offer_ts," +
		"o.time_stamp offer_date," + "o.offer_id offer_id," + "o.otype," +
		"o.seller_aid," + "sa.addr seller_addr," + "o.buyer_aid," + "ba.addr buyer_addr," +
		"o.token_id," + "o.active," + "NULL AS cancel_id," + "o.price/1e+18 price, " +
		"o.profit/1e+18 profit, " + "o.contract_aid," + "ca.addr contract_addr," +
		"o.rwalk_aid," + "rwa.addr rwalk_addr, " +
		"NULL AS itembought_ts," +
		"NULL AS itembought_date, " +
		"NULL AS canceled_ts," +
		"NULL AS canceled_date, " +
		"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT  AS real_ts, " +
		"o.time_stamp AS real_date " +
		"FROM " +
		"rw_new_offer o " +
		"JOIN transaction tx ON o.tx_id=tx.id " +
		"JOIN address sa ON o.seller_aid=sa.address_id " +
		"JOIN address ba ON o.buyer_aid=ba.address_id " +
		"JOIN address ca ON o.contract_aid=ca.address_id " +
		"JOIN address rwa ON o.rwalk_aid=rwa.address_id " +
		"WHERE o.contract_aid=$3 " +
		") UNION ALL (" +
		"SELECT " +
		"COALESCE(ib.id,can.id,o.id) AS record_id," +
		"COALESCE(ib.evtlog_id,can.evtlog_id,o.evtlog_id) AS evtlog_id," +
		"COALESCE(ib.block_num,can.block_num,o.block_num) AS block_num," +
		"COALESCE(ib.tx_id,can.tx_id,o.tx_id) AS tx_id, " +
		"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as offer_ts," +
		"o.time_stamp offer_date," + "o.offer_id offer_id," + "o.otype," +
		"o.seller_aid," + "sa.addr seller_addr," + "o.buyer_aid," + "ba.addr buyer_addr," +
		"o.token_id," + "o.active," + "can.id cancel_id," + "o.price/1e+18 price, " +
		"o.profit/1e+18 profit, " + "o.contract_aid," + "ca.addr contract_addr," +
		"o.rwalk_aid," + "rwa.addr rwalk_addr, " +
		"EXTRACT(EPOCH FROM ib.time_stamp)::BIGINT as itembought_ts," +
		"ib.time_stamp itembought_date, " +
		"EXTRACT(EPOCH FROM can.time_stamp)::BIGINT as canceled_ts," +
		"can.time_stamp canceled_date, " +
		"COALESCE(" +
		"EXTRACT(EPOCH FROM ib.time_stamp)::BIGINT," +
		"EXTRACT(EPOCH FROM can.time_stamp)::BIGINT," +
		"EXTRACT(EPOCH FROM o.time_stamp)::BIGINT" +
		") AS real_ts, " +
		"COALESCE(" +
		"ib.time_stamp," +
		"can.time_stamp," +
		"o.time_stamp" +
		") AS real_date " +
		"FROM " +
		"rw_new_offer o " +
		"JOIN transaction tx ON o.tx_id=tx.id " +
		"JOIN address sa ON o.seller_aid=sa.address_id " +
		"JOIN address ba ON o.buyer_aid=ba.address_id " +
		"JOIN address ca ON o.contract_aid=ca.address_id " +
		"JOIN address rwa ON o.rwalk_aid=rwa.address_id " +
		"LEFT JOIN rw_offer_canceled can ON (can.contract_aid=o.contract_aid) AND (can.offer_id=o.offer_id) " +
		"LEFT JOIN rw_item_bought ib ON (ib.contract_aid=o.contract_aid) AND (ib.offer_id=o.offer_id) " +
		"WHERE o.contract_aid=$3 AND ((can.id IS NOT NULL) OR (ib.id IS NOT NULL)) " +
		")" +
		") recs " +
		"ORDER BY real_ts " +
		"OFFSET $1 LIMIT $2"

	return queryList(ctx, r, "trading history", 16, query, func(rows pgx.Rows, rec *rwmodel.TradingHistoryLog) error {
		var nullProfit sql.NullFloat64
		var nullCanID sql.NullInt64
		var nullBoughtTs, nullCancelTs sql.NullInt64
		if err := rows.Scan(
			&rec.Id,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TimeStamp,
			store.TimeText(&rec.DateTime),
			&rec.OfferId,
			&rec.OfferType,
			&rec.SellerAid,
			&rec.SellerAddr,
			&rec.BuyerAid,
			&rec.BuyerAddr,
			&rec.TokenId,
			&rec.Active,
			&nullCanID,
			&rec.Price,
			&nullProfit,
			&rec.ContractAid,
			&rec.ContractAddr,
			&rec.RWalkAid,
			&rec.RWalkAddr,
			&nullBoughtTs,
			store.NullTimeText(&rec.ItemBoughtDate),
			&nullCancelTs,
			store.NullTimeText(&rec.CanceledDate),
			&rec.RealTs,
			store.TimeText(&rec.RealDate),
		); err != nil {
			return err
		}
		rec.Profit = profitFromNull(nullProfit)
		if nullCanID.Valid {
			rec.WasCanceled = true
		}
		if nullCancelTs.Valid {
			rec.CanceledTs = nullCancelTs.Int64
			timeCanceled := time.Unix(rec.CanceledTs, 0)
			timeOffered := time.Unix(rec.TimeStamp, 0)
			rec.CanceledDuration = timefmt.DurationToString(timefmt.TimeDifference(timeOffered, timeCanceled))
		}
		if nullBoughtTs.Valid {
			rec.WasBought = true
			rec.ItemBoughtTs = nullBoughtTs.Int64
			timeBought := time.Unix(rec.ItemBoughtTs, 0)
			timeOffered := time.Unix(rec.TimeStamp, 0)
			rec.BoughtDuration = timefmt.DurationToString(timefmt.TimeDifference(timeOffered, timeBought))
		}
		return nil
	}, offset, limit, contractAid)
}

// =============================================================================
// STATISTICS QUERIES
// =============================================================================

// RandomWalkStats returns contract-level statistics (volume, trades, mints,
// withdrawals, unique users, last mint price). Missing statistic rows leave
// the corresponding fields at zero, matching the legacy soft path.
func (r *Repo) RandomWalkStats(ctx context.Context, rwalkAid int64) (rwmodel.RWalkStats, error) {
	const op = "random walk stats"
	var output rwmodel.RWalkStats
	err := r.q(ctx).QueryRow(ctx, `SELECT
			total_vol/1e+18,
			total_num_trades,
			total_num_toks,
			total_withdrawals,
			money_accumulated/2e+18 withdrawal_amount
		FROM rw_stats
		WHERE rwalk_aid = $1`, rwalkAid).Scan(
		&output.TradingVol,
		&output.NumTrades,
		&output.TokensMinted,
		&output.NumWithdrawals,
		&output.WithdrawalAmount,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return output, nil
		}
		return output, store.WrapError(op, err)
	}
	var nUniqUsers sql.NullInt64
	err = r.q(ctx).QueryRow(ctx, "SELECT count(*) AS total FROM rw_uranks").Scan(&nUniqUsers)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return output, nil
		}
		return output, store.WrapError(op+": unique users", err)
	}
	if nUniqUsers.Valid {
		output.UniqueUsers = nUniqUsers.Int64
	}
	var nLastPrice sql.NullFloat64
	err = r.q(ctx).QueryRow(ctx, "SELECT price/1e+18 price FROM rw_mint_evt ORDER BY id DESC LIMIT 1").Scan(&nLastPrice)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return output, nil
		}
		return output, store.WrapError(op+": last minted price", err)
	}
	if nLastPrice.Valid {
		output.LastMintedPrice = nLastPrice.Float64
	}
	return output, nil
}

// MarketStats returns marketplace-level volume and trade counts; a missing
// stats row yields zeros.
func (r *Repo) MarketStats(ctx context.Context, marketAid int64) (rwmodel.MarketStats, error) {
	var output rwmodel.MarketStats
	err := r.q(ctx).QueryRow(ctx, `SELECT
			total_vol/1e+18,
			total_num_trades
		FROM rw_mkt_stats
		WHERE contract_aid = $1`, marketAid).Scan(
		&output.TradingVol,
		&output.NumTrades,
	)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return output, store.WrapError("market stats", err)
	}
	return output, nil
}

// TokenFullHistory returns every event touching one token (mint, offers,
// cancellations, buys, name changes, wallet transfers) in chronological
// order.
func (r *Repo) TokenFullHistory(ctx context.Context, rwalkAid, tokenID int64, offset, limit int) ([]rwmodel.FullHistoryEntry, error) {
	query := "SELECT " +
		"block_num," +
		"EXTRACT(EPOCH FROM time_stamp)::BIGINT as ts," +
		"time_stamp," +
		"contract_aid," +
		"contract_addr," +
		//--------Mint
		"token_id," +
		"owner_aid," +
		"owner_addr," +
		"seed," +
		"seed_num::TEXT," +
		//--------NewOffer
		"seller_aid," +
		"seller_addr," +
		"buyer_aid," +
		"buyer_addr," +
		"otype," +
		"offer_id," +
		"active, " +
		"price, " +
		//--------Offer Canceled
		"offer_canceled_id," +
		//--------Item Bought
		"item_bought_id," +
		//--------Token Name
		"token_name, " +
		//--------Transfer
		"transfer_id " +
		"FROM (" +
		"(" +
		"SELECT " +
		"t.block_num," +
		"EXTRACT(EPOCH FROM time_stamp)::BIGINT as ts," +
		"time_stamp," +
		"contract_aid," +
		//---------Mint
		"ca.addr contract_addr," +
		"token_id," +
		"owner_aid," +
		"oa.addr owner_addr," +
		"seed," +
		"seed_num," +
		//--------NewOffer
		"NULL AS seller_aid," +
		"NULL AS seller_addr," +
		"NULL AS buyer_aid," +
		"NULL AS buyer_addr," +
		"NULL as otype," +
		"CAST(NULL AS BIGINT) AS offer_id," +
		"NULL AS active," +
		"price/1e+18 AS price, " +
		//---------Offer Canceled
		"CAST(NULL AS BIGINT) AS offer_canceled_id," +
		//---------Item Bought
		"CAST(NULL AS BIGINT) AS item_bought_id," +
		//---------Token Name
		"CAST(NULL AS TEXT ) as token_name," +
		//---------TransferId
		"CAST(NULL AS BIGINT) as transfer_id " +
		"FROM rw_mint_evt t " +
		"LEFT JOIN address ca ON contract_aid=ca.address_id " +
		"LEFT JOIN address oa ON owner_aid=oa.address_id " +
		"WHERE (token_id=$1) AND (contract_aid=$2) " +
		"ORDER BY id" +
		") " +
		"UNION ALL" +
		"(" +
		"SELECT " +
		"t.block_num," +
		"EXTRACT(EPOCH FROM time_stamp)::BIGINT as ts," +
		"time_stamp," +
		"contract_aid," +
		"ca.addr contract_addr," +
		//---------Mint
		"token_id," +
		"NULL AS owner_aid," +
		"NULL AS owner_addr," +
		"NULL AS seed," +
		"NULL AS seed_num," +
		//---------NewOffer
		"seller_aid," +
		"sa.addr seller_addr," +
		"buyer_aid," +
		"ba.addr buyer_addr," +
		"otype," +
		"t.offer_id," +
		"active," +
		"price/1e+18 price, " +
		//---------Offer Canceled
		"CAST(NULL AS BIGINT) AS offer_canceled_id," +
		//---------Item Bought
		"CAST(NULL AS BIGINT) AS item_bought_id," +
		//---------Token Name
		"CAST(NULL AS TEXT) as token_name," +
		//---------Transfer
		"CAST(NULL AS BIGINT) as transfer_id " +
		"FROM rw_new_offer t " +
		"LEFT JOIN address ca ON contract_aid=ca.address_id " +
		"LEFT JOIN address sa ON seller_aid=sa.address_id " +
		"LEFT JOIN address ba ON buyer_aid=ba.address_id " +
		"WHERE (token_id=$1) AND (rwalk_aid=$2) " +
		"ORDER BY id" +
		")" +
		"UNION ALL" +
		"(" +
		"SELECT " +
		"c.block_num," +
		"EXTRACT(EPOCH FROM c.time_stamp)::BIGINT as ts," +
		"c.time_stamp," +
		"c.contract_aid," +
		"ca.addr contract_addr," +
		//---------Mint
		"token_id," +
		"NULL AS owner_aid," +
		"NULL AS owner_addr," +
		"NULL AS seed," +
		"NULL AS seed_num," +
		//---------NewOffer
		"seller_aid," +
		"sa.addr seller_addr," +
		"buyer_aid," +
		"ba.addr buyer_addr," +
		"o.otype as otype," +
		"o.offer_id," +
		"NULL AS active," +
		"o.price/1e+18 AS price," +
		//---------Offer Canceled
		"c.id offer_canceled_id," +
		//---------Item Bought
		"CAST(NULL AS BIGINT) AS item_bought_id," +
		//---------Token Name
		"CAST(NULL AS TEXT) as token_name," +
		//---------Transfer
		"CAST(NULL AS BIGINT) as transfer_id " +
		"FROM rw_offer_canceled c " +
		"JOIN rw_new_offer o ON (c.offer_id=o.offer_id AND c.contract_aid=o.contract_aid) " +
		"LEFT JOIN address ca ON c.contract_aid=ca.address_id " +
		"LEFT JOIN address sa ON o.seller_aid=sa.address_id " +
		"LEFT JOIN address ba ON o.buyer_aid=ba.address_id " +
		"WHERE (o.token_id=$1) AND (o.rwalk_aid=$2) " +
		"ORDER BY c.id " +
		")" +
		"UNION ALL" +
		"(" +
		"SELECT " +
		"b.block_num," +
		"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT as ts," +
		"b.time_stamp," +
		"b.contract_aid," +
		"ca.addr contract_addr," +
		//---------Mint
		"token_id," +
		"NULL AS owner_aid," +
		"NULL AS owner_addr," +
		"NULL AS seed," +
		"NULL AS seed_num," +
		//---------NewOffer
		"b.seller_aid," +
		"sa.addr seller_addr," +
		"b.buyer_aid," +
		"ba.addr buyer_addr," +
		"o.otype as otype," +
		"o.offer_id," +
		"NULL AS active," +
		"o.price/1e+18 AS price," +
		//---------Offer Canceled
		"CAST(NULL AS BIGINT) offer_canceled_id," +
		//---------Item Bought
		"b.id AS item_bought_id," +
		//---------Token Name
		"CAST(NULL AS TEXT) as token_name," +
		//---------Transfer
		"CAST(NULL AS BIGINT) as transfer_id " +
		"FROM rw_item_bought b " +
		"JOIN rw_new_offer o ON (b.offer_id=o.offer_id AND b.contract_aid=o.contract_aid) " +
		"LEFT JOIN address ca ON b.contract_aid=ca.address_id " +
		"LEFT JOIN address sa ON b.seller_aid=sa.address_id " +
		"LEFT JOIN address ba ON b.buyer_aid=ba.address_id " +
		"WHERE (o.token_id=$1) AND (o.rwalk_aid=$2) " +
		"ORDER BY b.id " +
		")" +
		"UNION ALL" +
		"(" +
		"SELECT " +
		"n.block_num," +
		"EXTRACT(EPOCH FROM n.time_stamp)::BIGINT as ts," +
		"n.time_stamp," +
		"n.contract_aid," +
		"ca.addr contract_addr," +
		//---------Mint
		"token_id," +
		"NULL AS owner_aid," +
		"NULL AS owner_addr," +
		"NULL AS seed," +
		"NULL AS seed_num," +
		//---------NewOffer
		"CAST(NULL AS BIGINT) AS seller_aid," +
		"NULL AS seller_addr," +
		"CAST(NULL AS BIGINT) AS buyer_aid," +
		"NULL AS buyer_addr," +
		"NULL AS otype," +
		"NULL AS offer_id," +
		"NULL AS active," +
		"NULL AS price," +
		//---------Offer Canceled
		"CAST(NULL AS BIGINT) offer_canceled_id," +
		//---------Item Bought
		"NULL AS item_bought_id," +
		//---------Token Name
		"n.new_name token_name," +
		//---------TransferID
		"CAST(NULL AS BIGINT) transfer_id " +
		"FROM rw_token_name n " +
		"LEFT JOIN address ca ON n.contract_aid=ca.address_id " +
		"WHERE (n.token_id=$1) AND (n.contract_aid=$2) " +
		"ORDER BY n.id " +
		")" +
		"UNION ALL" +
		"(" +
		"SELECT " +
		"tr.block_num," +
		"EXTRACT(EPOCH FROM tr.time_stamp)::BIGINT as ts," +
		"tr.time_stamp," +
		"tr.contract_aid," +
		"ca.addr contract_addr," +
		//---------Mint
		"tr.token_id," +
		"NULL AS owner_aid," +
		"NULL AS owner_addr," +
		"NULL AS seed," +
		"NULL AS seed_num," +
		//---------NewOffer
		"tr.from_aid AS seller_aid," +
		"fa.addr AS seller_addr," +
		"tr.to_aid AS buyer_aid," +
		"ta.addr AS buyer_addr," +
		"NULL AS otype," +
		"NULL AS offer_id," +
		"NULL AS active," +
		"NULL AS price," +
		//---------Offer Canceled
		"CAST(NULL AS BIGINT) offer_canceled_id," +
		//---------Item Bought
		"NULL AS item_bought_id," +
		//---------Token Name
		"NULL as token_name," +
		//---------Transfer
		"tr.id AS transfer_id " +
		"FROM rw_transfer tr " +
		"LEFT JOIN address ca ON tr.contract_aid=ca.address_id " +
		"LEFT JOIN address fa ON tr.from_aid=fa.address_id " +
		"LEFT JOIN address ta ON tr.to_aid=ta.address_id " +
		"LEFT JOIN rw_new_offer off ON tr.tx_id=off.tx_id " +
		"LEFT JOIN rw_mint_evt mint ON (tr.tx_id=mint.tx_id AND mint.token_id=$1) " +
		"LEFT JOIN rw_item_bought item ON tr.tx_id=item.tx_id " +
		"LEFT JOIN rw_offer_canceled cancel ON tr.tx_id=cancel.tx_id " +
		"LEFT JOIN rw_token_name name ON (tr.tx_id=name.tx_id AND name.token_id=$1) " +
		"LEFT JOIN evt_log elog ON ((tr.tx_id=elog.tx_id) AND (" +
		"elog.topic0_sig='55076e90' OR " + // new offer
		"elog.topic0_sig='caacc56f' OR " + // item bought
		"elog.topic0_sig='0ff09947' OR " + // offer canceled
		"elog.topic0_sig='8ad5e159' OR " + // token name
		"elog.topic0_sig='ad2bc79f' OR " + // mint event
		"elog.topic0_sig='ddf252ad' " + // transfer event
		") AND (elog.id=tr.evtlog_id) AND " +
		"(fa.addr!='0x0000000000000000000000000000000000000000')" +
		"AND (mint.token_id=$1))" +
		"WHERE (tr.token_id=$1) AND (tr.contract_aid=$2) " +
		"AND (off.id IS NULL) " +
		"AND (" +
		"(mint.id IS NULL) OR " +
		"(" +
		"(mint.id IS NOT NULL) AND (" +
		"(fa.addr!='0x0000000000000000000000000000000000000000')" +
		")" +
		")" +
		")" +
		"AND (item.id IS NULL) " +
		"AND (cancel.id IS NULL) " +
		"AND (name.id IS NULL) " +
		"ORDER BY tr.id " +
		")" +
		") AS data " + // FROM
		"ORDER BY ts " +
		"OFFSET $3 LIMIT $4"

	return queryList(ctx, r, "token full history", 32, query, scanFullHistoryEntry, tokenID, rwalkAid, offset, limit)
}

// scanFullHistoryEntry reads one row of the TokenFullHistory UNION and
// converts it into the typed record variant selected by the non-NULL
// discriminator columns.
func scanFullHistoryEntry(rows pgx.Rows, rec *rwmodel.FullHistoryEntry) error {
	var (
		blockNum        sql.NullInt64
		timestamp       sql.NullInt64
		datetime        sql.NullString
		contractAid     sql.NullInt64
		contractAddr    sql.NullString
		tokenID         sql.NullInt64
		ownerAid        sql.NullInt64
		ownerAddr       sql.NullString
		seed            sql.NullString
		seedNum         sql.NullString
		price           sql.NullFloat64
		sellerAid       sql.NullInt64
		sellerAddr      sql.NullString
		buyerAid        sql.NullInt64
		buyerAddr       sql.NullString
		otype           sql.NullInt64
		offerID         sql.NullInt64
		active          sql.NullBool
		offerCanceledID sql.NullInt64
		itemBoughtID    sql.NullInt64
		tokenName       sql.NullString
		transferID      sql.NullInt64
	)
	if err := rows.Scan(
		&blockNum,
		&timestamp,
		&datetime,
		&contractAid,
		&contractAddr,
		&tokenID,
		&ownerAid,
		&ownerAddr,
		&seed,
		&seedNum,
		&sellerAid,
		&sellerAddr,
		&buyerAid,
		&buyerAddr,
		&otype,
		&offerID,
		&active,
		&price,
		&offerCanceledID,
		&itemBoughtID,
		&tokenName,
		&transferID,
	); err != nil {
		return err
	}
	if seed.Valid {
		rec.Record = rwmodel.HistEntryMint{
			BlockNum:     blockNum.Int64,
			TimeStamp:    timestamp.Int64,
			DateTime:     datetime.String,
			ContractAid:  contractAid.Int64,
			ContractAddr: contractAddr.String,
			OwnerAid:     ownerAid.Int64,
			OwnerAddr:    ownerAddr.String,
			TokenId:      tokenID.Int64,
			SeedHex:      seed.String,
			SeedNum:      seedNum.String,
			Price:        price.Float64,
		}
		rec.RecordType = 1
	}
	if otype.Valid {
		rec.Record = rwmodel.HistEntryOffer{
			BlockNum:     blockNum.Int64,
			TimeStamp:    timestamp.Int64,
			DateTime:     datetime.String,
			ContractAid:  contractAid.Int64,
			ContractAddr: contractAddr.String,
			TokenId:      tokenID.Int64,
			SellerAid:    sellerAid.Int64,
			SellerAddr:   sellerAddr.String,
			BuyerAid:     buyerAid.Int64,
			BuyerAddr:    buyerAddr.String,
			OfferType:    int(otype.Int64),
			OfferId:      offerID.Int64,
			Active:       active.Bool,
			Price:        price.Float64,
		}
		rec.RecordType = 2
	}
	if offerCanceledID.Valid {
		iface := rwmodel.HistEntryOfferCanceled{
			BlockNum:        blockNum.Int64,
			TimeStamp:       timestamp.Int64,
			DateTime:        datetime.String,
			ContractAid:     contractAid.Int64,
			ContractAddr:    contractAddr.String,
			TokenId:         tokenID.Int64,
			OfferCanceledId: offerCanceledID.Int64,
			OfferType:       int(otype.Int64),
			OfferId:         offerID.Int64,
			SellerAid:       sellerAid.Int64,
			SellerAddr:      sellerAddr.String,
			BuyerAid:        buyerAid.Int64,
			BuyerAddr:       buyerAddr.String,
			Price:           price.Float64,
		}
		if iface.OfferType == 0 { // BUY
			iface.Aid = sellerAid.Int64
			iface.Address = sellerAddr.String
		} else {
			iface.Aid = buyerAid.Int64
			iface.Address = buyerAddr.String
		}
		rec.Record = iface
		rec.RecordType = 3
	}
	if itemBoughtID.Valid {
		iface := rwmodel.HistEntryItemBought{
			BlockNum:     blockNum.Int64,
			TimeStamp:    timestamp.Int64,
			DateTime:     datetime.String,
			ContractAid:  contractAid.Int64,
			ContractAddr: contractAddr.String,
			TokenId:      tokenID.Int64,
			ItemBoughtId: itemBoughtID.Int64,
			OfferType:    int(otype.Int64),
			OfferId:      offerID.Int64,
			SellerAid:    sellerAid.Int64,
			SellerAddr:   sellerAddr.String,
			BuyerAid:     buyerAid.Int64,
			BuyerAddr:    buyerAddr.String,
			Price:        price.Float64,
		}
		if iface.OfferType == 0 { // BUY
			iface.Aid = sellerAid.Int64
			iface.Address = sellerAddr.String
		} else {
			iface.Aid = buyerAid.Int64
			iface.Address = buyerAddr.String
		}
		rec.Record = iface
		rec.RecordType = 4
	}
	if tokenName.Valid {
		rec.Record = rwmodel.HistEntryTokenName{
			BlockNum:     blockNum.Int64,
			TimeStamp:    timestamp.Int64,
			DateTime:     datetime.String,
			ContractAid:  contractAid.Int64,
			ContractAddr: contractAddr.String,
			TokenId:      tokenID.Int64,
			TokenName:    tokenName.String,
		}
		rec.RecordType = 5
	}
	if transferID.Valid {
		rec.Record = rwmodel.HistEntryTransfer{
			BlockNum:     blockNum.Int64,
			TimeStamp:    timestamp.Int64,
			DateTime:     datetime.String,
			ContractAid:  contractAid.Int64,
			ContractAddr: contractAddr.String,
			TokenId:      tokenID.Int64,
			TransferId:   transferID.Int64,
			FromAid:      sellerAid.Int64,
			FromAddr:     sellerAddr.String,
			ToAid:        buyerAid.Int64,
			ToAddr:       buyerAddr.String,
		}
		rec.RecordType = 6
	}
	return nil
}

// MarketTradingVolumeByPeriod buckets marketplace sales into interval-second
// periods between initTs and finTs, carrying an accumulated volume that
// includes sales before the window.
func (r *Repo) MarketTradingVolumeByPeriod(ctx context.Context, contractAid int64, initTs, finTs, interval int) ([]rwmodel.VolumeHistory, error) {
	const op = "market trading volume by period"
	var initialVolume sql.NullFloat64
	err := r.q(ctx).QueryRow(ctx, `SELECT sum(price)/1e+18 AS accum_vol FROM rw_item_bought b
			JOIN rw_new_offer o ON o.offer_id=b.offer_id
			WHERE (b.time_stamp < TO_TIMESTAMP($1)) AND (o.contract_aid=$2)`,
		initTs, contractAid).Scan(&initialVolume)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, store.WrapError(op+": initial volume", err)
	}
	query := "WITH periods AS (" +
		"SELECT * FROM (" +
		"SELECT " +
		"generate_series AS start_ts," +
		"TO_TIMESTAMP(EXTRACT(EPOCH FROM generate_series) + $3) AS end_ts " +
		"FROM (" +
		"SELECT * " +
		"FROM generate_series(" +
		"TO_TIMESTAMP($1)," +
		"TO_TIMESTAMP($2)," +
		"TO_TIMESTAMP($3)-TO_TIMESTAMP(0)) " +
		") AS i" +
		") AS data " +
		") " +
		"SELECT " +
		"COALESCE(COUNT(b.id),0) as num_rows, " +
		"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
		"SUM(b.price)/1e+18 as volume " +
		"FROM periods AS p " +
		"LEFT JOIN LATERAL ( " +
		"SELECT b.id,b.time_stamp,o.price " +
		"FROM rw_item_bought b " +
		"JOIN rw_new_offer o ON b.offer_id=o.offer_id " +
		"WHERE b.contract_aid=$4" +
		") b ON " +
		"(p.start_ts <= b.time_stamp) AND " +
		"(b.time_stamp < p.end_ts) " +
		"GROUP BY start_ts " +
		"ORDER BY start_ts"

	accumVol := 0.0
	if initialVolume.Valid {
		accumVol = initialVolume.Float64
	}
	return queryList(ctx, r, op, 8, query, func(rows pgx.Rows, rec *rwmodel.VolumeHistory) error {
		var sumVolume sql.NullFloat64
		var numRows int
		if err := rows.Scan(&numRows, &rec.StartTs, &sumVolume); err != nil {
			return err
		}
		if sumVolume.Valid {
			rec.Volume = sumVolume.Float64
			rec.NumOperations = int64(numRows)
			accumVol += rec.Volume
		}
		rec.VolumeAccum = accumVol
		return nil
	}, initTs, finTs, interval, contractAid)
}

// TokenNameChanges returns the naming history of one token, newest first.
func (r *Repo) TokenNameChanges(ctx context.Context, tokenID int64) ([]rwmodel.TokenNameRec, error) {
	query := `SELECT
			t.block_num,
			EXTRACT(EPOCH FROM t.time_stamp)::BIGINT as ts,
			t.time_stamp,
			t.contract_aid,
			ca.addr,
			t.new_name,
			tx.tx_hash,
			oa.address_id,
			oa.addr
		FROM rw_token_name t
			LEFT JOIN address ca ON t.contract_aid=ca.address_id
			LEFT JOIN transaction tx ON t.tx_id=tx.id
			LEFT JOIN address oa ON tx.from_aid=oa.address_id
		WHERE token_id = $1
		ORDER by t.id DESC`
	return queryList(ctx, r, "token name changes", 32, query, func(rows pgx.Rows, rec *rwmodel.TokenNameRec) error {
		if err := rows.Scan(
			&rec.BlockNum,
			&rec.TimeStamp,
			store.TimeText(&rec.DateTime),
			&rec.ContractAid,
			&rec.ContractAddr,
			&rec.TokenName,
			&rec.TxHash,
			&rec.OwnerAid,
			&rec.OwnerAddr,
		); err != nil {
			return err
		}
		rec.TokenId = tokenID
		return nil
	}, tokenID)
}

// TokensByUser returns the tokens currently owned by one user.
func (r *Repo) TokensByUser(ctx context.Context, userAid int64) ([]rwmodel.UserToken, error) {
	query := `SELECT
			t.token_id,
			seed_hex,
			seed_num,
			last_price/1e+18
		FROM rw_token t
		WHERE cur_owner_aid=$1
		ORDER BY token_id`
	return queryList(ctx, r, "tokens by user", 32, query, func(rows pgx.Rows, rec *rwmodel.UserToken) error {
		return rows.Scan(
			&rec.TokenId,
			&rec.Seed,
			&rec.SeedNum,
			&rec.Price,
		)
	}, userAid)
}

// FloorPrice returns the cheapest active sell offer. An empty order book is
// not an error: noOffers is true and err is nil.
func (r *Repo) FloorPrice(ctx context.Context, rwalkAid, marketAid int64) (noOffers bool, floorPrice float64, offerID, tokenID int64, err error) {
	var nFloorPrice sql.NullFloat64
	var nOfferID, nTokenID sql.NullInt64
	query := `SELECT
			o.price/1e+18 price,
			o.offer_id,
			o.token_id
		FROM rw_new_offer o
		WHERE (active = 't') AND (otype=1) AND (o.rwalk_aid=$1) AND (o.contract_aid=$2)
		ORDER BY o.price ASC
		LIMIT 1`
	err = r.q(ctx).QueryRow(ctx, query, rwalkAid, marketAid).Scan(
		&nFloorPrice,
		&nOfferID,
		&nTokenID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return true, 0, 0, 0, nil
		}
		return false, 0, 0, 0, store.WrapError("floor price", err)
	}
	return false, nFloorPrice.Float64, nOfferID.Int64, nTokenID.Int64, nil
}

// TradingHistoryByUser returns closed offers where the user was buyer or
// seller.
func (r *Repo) TradingHistoryByUser(ctx context.Context, userAid int64) ([]rwmodel.Offer, error) {
	query := `SELECT
			o.id,
			o.evtlog_id,
			o.block_num,
			o.tx_id,
			tx.tx_hash,
			EXTRACT(EPOCH FROM o.time_stamp)::BIGINT as ts,
			o.time_stamp,
			o.offer_id,
			o.otype,
			o.seller_aid,
			sa.addr seller_addr,
			o.buyer_aid,
			ba.addr buyer_addr,
			o.token_id,
			o.active,
			can.id,
			o.price/1e+18 price,
			o.profit/1e+18 profit,
			o.contract_aid,
			ca.addr,
			o.rwalk_aid,
			rwa.addr
		FROM rw_new_offer o
			JOIN transaction tx ON o.tx_id=tx.id
			JOIN address sa ON o.seller_aid=sa.address_id
			JOIN address ba ON o.buyer_aid=ba.address_id
			JOIN address ca ON o.contract_aid=ca.address_id
			JOIN address rwa ON o.rwalk_aid=rwa.address_id
			LEFT JOIN rw_offer_canceled can ON (can.contract_aid=o.contract_aid) AND (can.offer_id=o.offer_id)
		WHERE (active = 'f') AND ((o.buyer_aid=$1) OR (o.seller_aid=$1))
		ORDER BY o.id`
	return queryList(ctx, r, "trading history by user", 16, query, scanClosedOffer, userAid)
}

// scanClosedOffer reads the shared closed-offer row shape used by
// TradingHistoryByUser and SaleHistory.
func scanClosedOffer(rows pgx.Rows, rec *rwmodel.Offer) error {
	var nullProfit sql.NullFloat64
	var nullCanID sql.NullInt64
	if err := rows.Scan(
		&rec.Id,
		&rec.EvtLogId,
		&rec.BlockNum,
		&rec.TxId,
		&rec.TxHash,
		&rec.TimeStamp,
		store.TimeText(&rec.DateTime),
		&rec.OfferId,
		&rec.OfferType,
		&rec.SellerAid,
		&rec.SellerAddr,
		&rec.BuyerAid,
		&rec.BuyerAddr,
		&rec.TokenId,
		&rec.Active,
		&nullCanID,
		&rec.Price,
		&nullProfit,
		&rec.ContractAid,
		&rec.ContractAddr,
		&rec.RWalkAid,
		&rec.RWalkAddr,
	); err != nil {
		return err
	}
	rec.Profit = profitFromNull(nullProfit)
	if nullCanID.Valid {
		rec.WasCanceled = true
	}
	return nil
}

// UserInfo returns per-user trading statistics. A user without stats rows
// yields store.ErrNotFound with the partially-populated record (user aid and
// the marketplace-contract flag), matching the legacy soft-miss shape.
func (r *Repo) UserInfo(ctx context.Context, userAid, rwalkAid int64) (rwmodel.UserInfo, error) {
	const op = "rwalk user info"
	var output rwmodel.UserInfo
	var nullAid sql.NullInt64
	err := r.q(ctx).QueryRow(ctx,
		"SELECT contract_aid FROM rw_new_offer WHERE contract_aid=$1 LIMIT 1", userAid).Scan(&nullAid)
	if err == nil {
		output.IsMarketPlaceContract = true
	} else if !errors.Is(err, pgx.ErrNoRows) {
		return output, store.WrapError(op+": marketplace check", err)
	}
	output.UserAid = userAid
	err = r.q(ctx).QueryRow(ctx, `SELECT
			us.total_vol/1e+18,
			us.total_num_trades,
			us.total_num_toks,
			us.total_withdrawals
		FROM rw_user_stats us
		WHERE user_aid=$1 AND rwalk_aid=$2`, userAid, rwalkAid).Scan(
		&output.TotalVolume,
		&output.TotalNumTrades,
		&output.TotalMintedToks,
		&output.TotalNumWithdrawals,
	)
	if err != nil {
		return output, store.WrapError(op, err)
	}
	return output, nil
}

// Top5TradedTokens returns the five most traded tokens.
func (r *Repo) Top5TradedTokens(ctx context.Context) ([]rwmodel.TopTradedToken, error) {
	query := `SELECT
			token_id,
			num_trades,
			seed_hex
		FROM rw_token t
		ORDER BY num_trades DESC, token_id ASC
		LIMIT 5`
	return queryList(ctx, r, "top 5 traded tokens", 16, query, func(rows pgx.Rows, rec *rwmodel.TopTradedToken) error {
		return rows.Scan(
			&rec.TokenId,
			&rec.TotalTrades,
			&rec.SeedHex,
		)
	})
}

// TokenInfo returns current ownership, naming and trading state of one
// token; a token without a rw_token row yields store.ErrNotFound.
func (r *Repo) TokenInfo(ctx context.Context, rwalkAid, tokenID int64) (rwmodel.TokenInfo, error) {
	const op = "rwalk token info"
	var output rwmodel.TokenInfo
	output.TokenId = tokenID
	err := r.q(ctx).QueryRow(ctx, `SELECT
			t.cur_owner_aid,
			oa.addr,
			seed_hex,
			seed_num,
			last_name,
			last_price/1e+18,
			t.total_vol/1e+18,
			t.num_trades
		FROM rw_token t
			LEFT JOIN address oa ON oa.address_id=t.cur_owner_aid
		WHERE t.rwalk_aid=$1 AND token_id=$2`, rwalkAid, tokenID).Scan(
		&output.CurOwnerAid,
		&output.CurOwnerAddr,
		&output.SeedHex,
		&output.SeedNum,
		&output.CurName,
		&output.LastPrice,
		&output.TotalVolume,
		&output.NumTrades,
	)
	if err != nil {
		return output, store.WrapError(op, err)
	}
	err = r.q(ctx).QueryRow(ctx, `SELECT
			EXTRACT(EPOCH FROM time_stamp)::BIGINT as ts,
			time_stamp
		FROM rw_token_name
		WHERE contract_aid=$1 AND token_id=$2
		ORDER BY id DESC
		LIMIT 1`, rwalkAid, tokenID).Scan(&output.LastNameUpdateTs, store.TimeText(&output.LastNameUpdateDate))
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return output, store.WrapError(op+": last name update", err)
	}
	return output, nil
}

// TokenMinted reports whether a mint event exists for tokenID (any
// contract). A missing token is not an error; the error return is reserved
// for real database failures.
func (r *Repo) TokenMinted(ctx context.Context, tokenID int64) (bool, error) {
	var got int64
	err := r.q(ctx).QueryRow(ctx, "SELECT token_id FROM rw_mint_evt m WHERE token_id=$1", tokenID).Scan(&got)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, store.WrapError("token minted check", err)
	}
	return true, nil
}

// MintIntervals returns mint number and time elapsed between consecutive
// mints (for a scatter plot).
func (r *Repo) MintIntervals(ctx context.Context, rwalkAid int64) ([]rwmodel.MintInterval, error) {
	query := `SELECT
			EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,
			token_id
		FROM rw_mint_evt m
		WHERE contract_aid = $1
		ORDER BY token_id`
	records, err := queryList(ctx, r, "mint intervals", 256, query, func(rows pgx.Rows, rec *rwmodel.MintInterval) error {
		if err := rows.Scan(&rec.TimeStamp, &rec.TokenId); err != nil {
			return err
		}
		rec.MintNumber = rec.TokenId // token ids are sequential
		return nil
	}, rwalkAid)
	if err != nil {
		return nil, err
	}
	var prevTs int64
	for i := range records {
		rec := &records[i]
		if prevTs > 0 {
			rec.Interval = rec.TimeStamp - prevTs
		}
		prevTs = rec.TimeStamp
	}
	return records, nil
}

// WithdrawalChart returns the cumulative withdrawable amount over time
// (half of every mint price accumulates).
func (r *Repo) WithdrawalChart(ctx context.Context, rwalkAid int64) ([]rwmodel.WithdrawalChartEntry, error) {
	query := `SELECT
			EXTRACT(EPOCH FROM m.time_stamp)::BIGINT as ts,
			time_stamp,
			price/2e+18 price
		FROM rw_mint_evt m
		WHERE contract_aid=$1`
	var withdrawalAmount float64
	return queryList(ctx, r, "withdrawal chart", 256, query, func(rows pgx.Rows, rec *rwmodel.WithdrawalChartEntry) error {
		if err := rows.Scan(
			&rec.TimeStamp,
			store.TimeText(&rec.DateTime),
			&rec.WithdrawalAmount,
		); err != nil {
			return err
		}
		withdrawalAmount += rec.WithdrawalAmount
		rec.WithdrawalAmount = withdrawalAmount
		return nil
	}, rwalkAid)
}

// SaleHistory returns completed sales ordered by purchase time (the
// rw_item_bought timestamp, so "latest sale" matches the most recent
// on-chain buy; the tx hash is the purchase transaction).
func (r *Repo) SaleHistory(ctx context.Context, contractAid int64, offset, limit int) ([]rwmodel.Offer, error) {
	query := `SELECT
			o.id,
			o.evtlog_id,
			o.block_num,
			ib.tx_id,
			tx.tx_hash,
			EXTRACT(EPOCH FROM ib.time_stamp)::BIGINT as ts,
			ib.time_stamp,
			o.offer_id,
			o.otype,
			o.seller_aid,
			sa.addr seller_addr,
			o.buyer_aid,
			ba.addr buyer_addr,
			o.token_id,
			o.active,
			can.id,
			o.price/1e+18 price,
			o.profit/1e+18 profit,
			o.contract_aid,
			ca.addr,
			o.rwalk_aid,
			rwa.addr
		FROM rw_new_offer o
			JOIN rw_item_bought ib ON ib.contract_aid=o.contract_aid AND ib.offer_id=o.offer_id
			JOIN transaction tx ON ib.tx_id=tx.id
			JOIN address sa ON o.seller_aid=sa.address_id
			JOIN address ba ON o.buyer_aid=ba.address_id
			JOIN address ca ON o.contract_aid=ca.address_id
			JOIN address rwa ON o.rwalk_aid=rwa.address_id
			LEFT JOIN rw_offer_canceled can ON (can.contract_aid=o.contract_aid) AND (can.offer_id=o.offer_id)
		WHERE (active = 'f') AND (o.contract_aid=$3) AND (can.id IS NULL)
			ORDER BY ib.time_stamp ASC, o.id ASC
		OFFSET $1 LIMIT $2`
	return queryList(ctx, r, "sale history", 16, query, scanClosedOffer, offset, limit, contractAid)
}

// FloorPriceByPeriod returns the minimum open sell-offer price per
// interval-second bucket; buckets without offers are omitted.
func (r *Repo) FloorPriceByPeriod(ctx context.Context, rwalkAid, marketAid int64, initTs, finTs, interval int) ([]rwmodel.FloorPrice, error) {
	const op = "floor price by period"
	query := "WITH periods AS (" +
		"SELECT * FROM (" +
		"SELECT " +
		"generate_series AS start_ts," +
		"TO_TIMESTAMP(EXTRACT(EPOCH FROM generate_series) + $3) AS end_ts " +
		"FROM (" +
		"SELECT * " +
		"FROM generate_series(" +
		"TO_TIMESTAMP($1)," +
		"TO_TIMESTAMP($2)," +
		"TO_TIMESTAMP($3)-TO_TIMESTAMP(0)) " +
		") AS i" +
		") AS data " +
		") " +
		"SELECT " +
		"ROUND(FLOOR(EXTRACT(EPOCH FROM start_ts)))::BIGINT as start_ts," +
		"MIN(o.price)/1e+18 as floor_price " +
		"FROM periods AS p " +
		"LEFT JOIN LATERAL ( " +
		"SELECT o.id,o.time_stamp,o.price " +
		"FROM rw_new_offer o " +
		"WHERE o.contract_aid=$4 AND o.rwalk_aid=$5 " +
		"AND otype = 1 " +
		") o ON " +
		"(p.start_ts <= o.time_stamp) AND " +
		"(o.time_stamp < p.end_ts) " +
		"GROUP BY start_ts " +
		"ORDER BY start_ts"

	rows, err := r.q(ctx).Query(ctx, query, initTs, finTs, interval, marketAid, rwalkAid)
	if err != nil {
		return nil, store.WrapError(op, err)
	}
	defer rows.Close()
	records := make([]rwmodel.FloorPrice, 0, 8)
	for rows.Next() {
		var rec rwmodel.FloorPrice
		var nullFloat sql.NullFloat64
		if err := rows.Scan(&rec.TimeStamp, &nullFloat); err != nil {
			return nil, store.WrapError(op, err)
		}
		if !nullFloat.Valid {
			continue // bucket without offers
		}
		rec.Price = nullFloat.Float64
		records = append(records, rec)
	}
	if err := rows.Err(); err != nil {
		return nil, store.WrapError(op, err)
	}
	return records, nil
}

// MintedTokensCSV returns every mint of the contract joined with current
// token state, in token order (feeds the CSV export).
func (r *Repo) MintedTokensCSV(ctx context.Context, rwalkAid int64) ([]rwmodel.TokenMintCSV, error) {
	query := `SELECT
			t.block_num,
			EXTRACT(EPOCH FROM t.time_stamp)::BIGINT as ts,
			t.time_stamp,
			t.contract_aid,
			ca.addr,
			t.token_id,
			t.owner_aid,
			oa.addr minter_addr,
			t.seed seed_hex,
			t.seed_num,
			price/1e+18,
			tx.tx_hash,
			tk.num_trades,
			tk.total_vol/1e+18 total_vol,
			tk.last_price/1e+18 last_price,
			tk.last_name,
			loa.addr
		FROM rw_mint_evt t
			LEFT JOIN address ca ON t.contract_aid=ca.address_id
			LEFT JOIN address oa ON t.owner_aid=oa.address_id
			LEFT JOIN transaction tx ON t.tx_id=tx.id
			LEFT JOIN rw_token tk ON (t.token_id=tk.token_id AND t.contract_aid=tk.rwalk_aid)
			LEFT JOIN address loa ON (tk.cur_owner_aid=loa.address_id)
		WHERE contract_aid=$1
		ORDER by t.token_id`
	return queryList(ctx, r, "minted tokens csv", 32, query, func(rows pgx.Rows, rec *rwmodel.TokenMintCSV) error {
		var nNumTrades sql.NullInt64
		var nLastOwnerAddr, nLastName sql.NullString
		var nTotalVol, nLastPrice sql.NullFloat64
		if err := rows.Scan(
			&rec.BlockNum,
			&rec.TimeStamp,
			store.TimeText(&rec.DateTime),
			&rec.ContractAid,
			&rec.ContractAddr,
			&rec.TokenId,
			&rec.MinterAid,
			&rec.MinterAddr,
			&rec.Seed,
			&rec.SeedNum,
			&rec.Price,
			&rec.TxHash,
			&nNumTrades,
			&nTotalVol,
			&nLastPrice,
			&nLastName,
			&nLastOwnerAddr,
		); err != nil {
			return err
		}
		if nNumTrades.Valid {
			rec.NumTrades = nNumTrades.Int64
		}
		if nTotalVol.Valid {
			rec.TotalVolume = nTotalVol.Float64
		}
		if nLastPrice.Valid {
			rec.LastPrice = nLastPrice.Float64
		}
		if nLastName.Valid {
			rec.LastName = nLastName.String
		}
		if nLastOwnerAddr.Valid {
			rec.LastOwner = nLastOwnerAddr.String
		}
		return nil
	}, rwalkAid)
}

func monthName(code int64) string {
	switch code {
	case 1:
		return "January"
	case 2:
		return "February"
	case 3:
		return "March"
	case 4:
		return "April"
	case 5:
		return "May"
	case 6:
		return "June"
	case 7:
		return "July"
	case 8:
		return "August"
	case 9:
		return "September"
	case 10:
		return "October"
	case 11:
		return "November"
	case 12:
		return "December"
	default:
		return "???"
	}
}

// MintReport aggregates mints per calendar month with the cumulative
// redeemable amount (half of everything deposited).
func (r *Repo) MintReport(ctx context.Context) ([]rwmodel.MintReportRec, error) {
	query := `WITH periods AS (
			SELECT
				s.m_start,
				date_trunc('month',s.m_start) + interval '1 month - 1 second' AS m_end,
				extract(year from s.m_start)*100+extract(month from s.m_start) AS yearmonth
			FROM generate_series('2021-11-01','2022-12-31',interval '1 month') AS s(m_start)
		)
		SELECT
			yearmonth,
			count(m.id) total_minted,
			sum(m.price) total_wei,
			sum(m.price)/1e18 total_eth
		FROM periods p
			JOIN rw_mint_evt m ON
				p.yearmonth = extract(year from m.time_stamp)*100+extract(month from m.time_stamp)
		GROUP BY p.yearmonth
		ORDER BY p.yearmonth`
	var sumDeposited float64
	return queryList(ctx, r, "mint report", 32, query, func(rows pgx.Rows, rec *rwmodel.MintReportRec) error {
		var nTotalMinted sql.NullInt64
		var nTotalWei sql.NullString
		var nTotalEth sql.NullFloat64
		var yearmonth int64
		if err := rows.Scan(
			&yearmonth,
			&nTotalMinted,
			&nTotalWei,
			&nTotalEth,
		); err != nil {
			return err
		}
		rec.Year = yearmonth / 100
		rec.Month = yearmonth % 100
		rec.MonthStr = monthName(rec.Month) + fmt.Sprintf(" %v", rec.Year)
		if nTotalMinted.Valid {
			rec.TotalMinted = nTotalMinted.Int64
		}
		if nTotalWei.Valid {
			rec.TotalWei = nTotalWei.String
		}
		if nTotalEth.Valid {
			rec.TotalEth = nTotalEth.Float64
			sumDeposited += rec.TotalEth
			rec.RedeemAmount = sumDeposited / 2
		}
		return nil
	})
}
