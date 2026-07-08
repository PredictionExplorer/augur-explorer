package cosmicgame

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// bidSelectBase is the unified SELECT for bid lists: the bid row joined with
// its transaction, the bidder address and the per-bid NFT/ERC-20 donations.
const bidSelectBase = "SELECT b.evtlog_id,b.block_num,t.id,t.tx_hash," +
	"EXTRACT(EPOCH FROM b.time_stamp)::BIGINT,b.time_stamp," +
	"b.bidder_aid,ba.addr,b.eth_price,b.eth_price/1e18 AS eth_price_eth, " +
	"b.cst_price, b.cst_price/1e18 AS cst_price_eth, b.rwalk_nft_id," +
	"d.token_id,d.tok_addr, d.token_uri, b.msg, b.round_num, " +
	"b.cst_reward, b.cst_reward/1e18, " +
	"b.bid_cst_reward_amount, (CASE WHEN b.bid_cst_reward_amount >= 0 THEN b.bid_cst_reward_amount/1e18 ELSE -1 END), " +
	"b.cst_dutch_auction_duration, (CASE WHEN b.cst_dutch_auction_duration >= 0 THEN b.cst_dutch_auction_duration::bigint ELSE -1 END), " +
	"b.bid_type, " +
	"EXTRACT(EPOCH FROM b.prize_time)::BIGINT AS prize_time_ts, b.prize_time, " +
	"GREATEST(0, EXTRACT(EPOCH FROM b.prize_time)::BIGINT - EXTRACT(EPOCH FROM NOW())::BIGINT) AS time_until_prize, " +
	"b.bid_position, d2.tok_addr, d2.amount, d2.amount/1e18 " +
	"FROM cg_bid b " +
	"LEFT JOIN transaction t ON t.id=tx_id " +
	"LEFT JOIN address ba ON b.bidder_aid=ba.address_id " +
	"LEFT JOIN LATERAL (SELECT d.bid_id,token_id,token_aid,ta.addr tok_addr,d.token_uri " +
	"FROM cg_nft_donation d " +
	"JOIN address ta ON d.token_aid=ta.address_id) d ON b.id=d.bid_id " +
	"LEFT JOIN LATERAL (SELECT d.bid_id,token_id,token_aid,ta.addr tok_addr,d.amount " +
	"FROM cg_erc20_donation d " +
	"JOIN address ta ON d.token_aid=ta.address_id) d2 ON b.id=d2.bid_id "

// The bid query builder only accepts clauses from these whitelists. Every
// clause is a compile-time constant in this package; the whitelist turns a
// slipped-through dynamic string (the classic ORDER-BY injection) into an
// immediate error instead of SQL. TestBidSelectQueryWhitelists exercises
// every entry and the rejection path.
var (
	bidWhereWhitelist = map[string]bool{
		"":                true,
		"b.evtlog_id=$1":  true,
		"b.round_num=$1":  true,
		"b.bidder_aid=$1": true,
		"b.round_num=$1 AND b.msg IS NOT NULL AND TRIM(b.msg) <> ''": true,
	}
	bidOrderWhitelist = map[string]bool{
		"":                    true,
		"b.id ASC":            true,
		"b.id DESC":           true,
		"b.bid_position ASC":  true,
		"b.bid_position DESC": true,
	}
	bidPagingWhitelist = map[string]bool{
		"":                   true,
		"OFFSET $1 LIMIT $2": true,
		"OFFSET $2 LIMIT $3": true,
	}
)

// bidSelectQuery composes the unified bid SELECT from whitelisted clauses.
func bidSelectQuery(whereClause, orderBy, paging string) (string, error) {
	if !bidWhereWhitelist[whereClause] {
		return "", fmt.Errorf("bid query: WHERE clause not whitelisted: %q", whereClause)
	}
	if !bidOrderWhitelist[orderBy] {
		return "", fmt.Errorf("bid query: ORDER BY clause not whitelisted: %q", orderBy)
	}
	if !bidPagingWhitelist[paging] {
		return "", fmt.Errorf("bid query: paging clause not whitelisted: %q", paging)
	}
	query := bidSelectBase
	if whereClause != "" {
		query += "WHERE " + whereClause + " "
	}
	if orderBy != "" {
		query += "ORDER BY " + orderBy + " "
	}
	if paging != "" {
		query += paging
	}
	return query, nil
}

// scanBidRow scans one bidSelectBase row. The donation columns come from the
// LATERAL joins and are NULL for bids without donations; NFTDonationTokenId
// keeps the legacy -1 default in that case.
func scanBidRow(rows pgx.Rows, rec *p.CGBidRec) error {
	var nullTokenID sql.NullInt64
	var nullTokAddr, nullTokenURI sql.NullString
	var nullDonatedERC20Addr, nullDonatedERC20Amount sql.NullString
	var nullDonatedERC20AmountEth sql.NullFloat64

	err := rows.Scan(
		&rec.Tx.EvtLogId,
		&rec.Tx.BlockNum,
		&rec.Tx.TxId,
		&rec.Tx.TxHash,
		&rec.Tx.TimeStamp,
		store.TimeText(&rec.Tx.DateTime),
		&rec.BidderAid,
		&rec.BidderAddr,
		&rec.EthPrice,
		&rec.EthPriceEth,
		&rec.CstPrice,
		&rec.CstPriceEth,
		&rec.RWalkNFTId,
		&nullTokenID,
		&nullTokAddr,
		&nullTokenURI,
		&rec.Message,
		&rec.RoundNum,
		&rec.CSTReward,
		&rec.CSTRewardEth,
		&rec.BidCstRewardAmount,
		&rec.BidCstRewardAmountEth,
		&rec.CstDutchAuctionDuration,
		&rec.CstDutchAuctionDurationInt,
		&rec.BidType,
		&rec.PrizeTime,
		store.TimeText(&rec.PrizeTimeDate),
		&rec.TimeUntilPrize,
		&rec.BidPosition,
		&nullDonatedERC20Addr,
		&nullDonatedERC20Amount,
		&nullDonatedERC20AmountEth,
	)
	if err != nil {
		return err
	}

	rec.NFTDonationTokenId = -1
	if nullTokenID.Valid {
		rec.NFTDonationTokenId = nullTokenID.Int64
	}
	if nullTokAddr.Valid {
		rec.NFTDonationTokenAddr = nullTokAddr.String
	}
	if nullTokenURI.Valid {
		rec.NFTTokenURI = nullTokenURI.String
	}
	if nullDonatedERC20Addr.Valid {
		rec.DonatedERC20TokenAddr = nullDonatedERC20Addr.String
	}
	if nullDonatedERC20Amount.Valid {
		rec.DonatedERC20TokenAmount = nullDonatedERC20Amount.String
	}
	if nullDonatedERC20AmountEth.Valid {
		rec.DonatedERC20TokenAmountEth = nullDonatedERC20AmountEth.Float64
	}
	return nil
}

// bidList runs a whitelisted bid query and scans the rows.
func bidList(ctx context.Context, r *Repo, op, whereClause, orderBy, paging string, capHint int, args ...any) ([]p.CGBidRec, error) {
	query, err := bidSelectQuery(whereClause, orderBy, paging)
	if err != nil {
		return nil, store.WrapError(op, err)
	}
	return queryList(ctx, r, op, capHint, query, scanBidRow, args...)
}

// BidIDByEvtlog returns the cg_bid row id of the bid inserted for evtlogID,
// or store.ErrNotFound when that event has no bid.
func (r *Repo) BidIDByEvtlog(ctx context.Context, evtlogID int64) (int64, error) {
	var id int64
	err := r.pool().QueryRow(ctx, "SELECT b.id FROM cg_bid b WHERE b.evtlog_id=$1", evtlogID).Scan(&id)
	if err != nil {
		return 0, store.WrapError("bid id by evtlog", err)
	}
	return id, nil
}

// Bids returns all bids, newest first. limit 0 means no effective limit.
func (r *Repo) Bids(ctx context.Context, offset, limit int) ([]p.CGBidRec, error) {
	if limit == 0 {
		limit = 1000000
	}
	return bidList(ctx, r, "bids", "", "b.id DESC", "OFFSET $1 LIMIT $2", 256, offset, limit)
}

// BidInfo returns the bid of one BidPlaced event log, or store.ErrNotFound
// when the event has no bid row.
func (r *Repo) BidInfo(ctx context.Context, evtlogID int64) (p.CGBidRec, error) {
	const op = "bid info"
	recs, err := bidList(ctx, r, op, "b.evtlog_id=$1", "", "", 1, evtlogID)
	if err != nil {
		return p.CGBidRec{}, err
	}
	if len(recs) == 0 {
		return p.CGBidRec{}, store.WrapError(op, pgx.ErrNoRows)
	}
	return recs[0], nil
}

// EvtlogIDByRoundAndBidPosition returns the evtlog_id of the bid at
// bidPosition within roundNum, or store.ErrNotFound.
func (r *Repo) EvtlogIDByRoundAndBidPosition(ctx context.Context, roundNum, bidPosition int64) (int64, error) {
	var evtlogID int64
	err := r.pool().QueryRow(ctx,
		"SELECT b.evtlog_id FROM cg_bid b WHERE b.round_num=$1 AND b.bid_position=$2",
		roundNum, bidPosition).Scan(&evtlogID)
	if err != nil {
		return 0, store.WrapError("evtlog id by round and bid position", err)
	}
	return evtlogID, nil
}

// BidsWithMessageByRound returns the bids of one round that carry a
// non-empty message, ordered by bid position (descending when sortDesc).
// limit 0 defaults to 1000.
func (r *Repo) BidsWithMessageByRound(ctx context.Context, roundNum int64, sortDesc bool, offset, limit int) ([]p.CGBidRec, error) {
	if limit == 0 {
		limit = 1000
	}
	orderBy := "b.bid_position ASC"
	if sortDesc {
		orderBy = "b.bid_position DESC"
	}
	return bidList(ctx, r, "bids with message by round",
		"b.round_num=$1 AND b.msg IS NOT NULL AND TRIM(b.msg) <> ''",
		orderBy, "OFFSET $2 LIMIT $3", 32, roundNum, offset, limit)
}

// BidsByRound returns one page of a round's bids (ascending insertion order
// unless sort is 1) together with the round's total bid count.
func (r *Repo) BidsByRound(ctx context.Context, roundNum int64, sort, offset, limit int) ([]p.CGBidRec, int64, error) {
	const op = "bids by round"
	orderBy := "b.id ASC"
	if sort == 1 {
		orderBy = "b.id DESC"
	}
	records, err := bidList(ctx, r, op, "b.round_num=$1", orderBy, "OFFSET $2 LIMIT $3", 32, roundNum, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	var totalRows int64
	err = r.pool().QueryRow(ctx,
		"SELECT count(*) AS total_rows FROM cg_bid b WHERE round_num=$1", roundNum).Scan(&totalRows)
	if err != nil {
		return nil, 0, store.WrapError(op+": total count", err)
	}
	return records, totalRows, nil
}

// BidCountForRound returns the number of bids in cg_bid for the given round.
// Use this for the dashboard so the "Bids This Round" count matches the bid
// list for that round.
func (r *Repo) BidCountForRound(ctx context.Context, roundNum int64) (int64, error) {
	var n int64
	err := r.pool().QueryRow(ctx, "SELECT count(*) FROM cg_bid WHERE round_num=$1", roundNum).Scan(&n)
	if err != nil {
		return 0, store.WrapError("bid count for round", err)
	}
	return n, nil
}

// LastCstBidEvtlogForBidder returns the evtlog_id of the latest CST bid in a
// round by address (case-insensitive), or store.ErrNotFound.
func (r *Repo) LastCstBidEvtlogForBidder(ctx context.Context, roundNum int64, bidderAddr string) (int64, error) {
	query := `SELECT b.evtlog_id FROM cg_bid b
		JOIN address ba ON b.bidder_aid=ba.address_id
		WHERE b.round_num=$1 AND lower(ba.addr)=lower($2) AND b.cst_price > 0
		ORDER BY b.time_stamp DESC LIMIT 1`
	var evtlogID int64
	err := r.pool().QueryRow(ctx, query, roundNum, bidderAddr).Scan(&evtlogID)
	if err != nil {
		return 0, store.WrapError("last cst bid evtlog for bidder", err)
	}
	return evtlogID, nil
}

// RoundStartTimestamp returns the epoch timestamp of a round's first bid, or
// 0 when the round has no bids yet (the callers treat 0 as "unknown").
func (r *Repo) RoundStartTimestamp(ctx context.Context, roundNum uint64) (int64, error) {
	query := `SELECT
			EXTRACT(EPOCH FROM b.time_stamp)::BIGINT
		FROM cg_bid b
		WHERE round_num=$1
		ORDER BY b.id LIMIT 1`
	var ts int64
	err := r.pool().QueryRow(ctx, query, roundNum).Scan(&ts)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, store.WrapError("round start timestamp", err)
	}
	return ts, nil
}

// RandomWalkTokensUsedInBids returns every bid that consumed a RandomWalk
// NFT for its discount, newest first.
func (r *Repo) RandomWalkTokensUsedInBids(ctx context.Context) ([]p.CGRWalkUsed, error) {
	query := `SELECT
			b.id,
			b.evtlog_id,
			b.block_num,
			tx.id,
			tx.tx_hash,
			EXTRACT(EPOCH FROM b.time_stamp)::BIGINT,
			b.time_stamp,
			b.round_num,
			b.bidder_aid,
			ba.addr,
			b.rwalk_nft_id
		FROM cg_bid b
			LEFT JOIN transaction tx ON tx.id=b.tx_id
			LEFT JOIN address ba ON b.bidder_aid=ba.address_id
		WHERE b.rwalk_nft_id != -1
		ORDER BY b.id DESC`
	scan := func(rows pgx.Rows, rec *p.CGRWalkUsed) error {
		return rows.Scan(
			&rec.RecordId,
			&rec.EvtLogId,
			&rec.BlockNum,
			&rec.TxId,
			&rec.TxHash,
			&rec.TimeStamp,
			store.TimeText(&rec.DateTime),
			&rec.RoundNum,
			&rec.BidderAid,
			&rec.BidderAddr,
			&rec.RWalkTokenId,
		)
	}
	return queryList(ctx, r, "random walk tokens used in bids", 16, query, scan)
}

// BidRowIDByEvtlogID returns the cg_bid row id for a BidPlaced evtlog_id, or
// 0 when the event carries no bid (a pure Donate() call) — 0 is the value
// the ETL persists in that case.
func (r *Repo) BidRowIDByEvtlogID(ctx context.Context, bidEvtlogID int64) (int64, error) {
	var id int64
	err := r.pool().QueryRow(ctx, "SELECT id FROM cg_bid WHERE evtlog_id=$1", bidEvtlogID).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, store.WrapError("bid row id by evtlog id", err)
	}
	return id, nil
}
