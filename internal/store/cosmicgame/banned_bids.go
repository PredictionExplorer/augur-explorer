package cosmicgame

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	cgmodel "github.com/PredictionExplorer/augur-explorer/internal/model/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

func scanBannedBid(rows pgx.Rows, rec *cgmodel.CGBannedBidRec) error {
	return rows.Scan(&rec.Id, &rec.BidId, &rec.UserAddr, &rec.CreatedAt)
}

// BannedBids returns all rows from cg_banned_bids ordered by id.
func (r *Repo) BannedBids(ctx context.Context) ([]cgmodel.CGBannedBidRec, error) {
	query := "SELECT id, bid_id, user_addr, created_at FROM cg_banned_bids ORDER BY id"
	return queryList(ctx, r, "banned bids", 8, query, scanBannedBid)
}

// BannedBidPageCursor identifies the final row returned by a newest-first
// active-ban page.
type BannedBidPageCursor struct {
	ID int64
}

// BannedBidsPage returns at most limit active bid bans newest first. A nil
// cursor starts at the newest row; hasMore reports whether another page exists.
func (r *Repo) BannedBidsPage(
	ctx context.Context,
	after *BannedBidPageCursor,
	limit int,
) (records []cgmodel.CGBannedBidRec, hasMore bool, err error) {
	const op = "banned bids page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: limit must be positive", op)
	}

	query := `SELECT id, bid_id, user_addr, created_at
		FROM cg_banned_bids
		ORDER BY id DESC
		LIMIT $1`
	args := []any{limit + 1}
	if after != nil {
		if after.ID < 1 {
			return nil, false, fmt.Errorf("%s: invalid cursor", op)
		}
		query = `SELECT id, bid_id, user_addr, created_at
			FROM cg_banned_bids
			WHERE id < $1
			ORDER BY id DESC
			LIMIT $2`
		args = []any{after.ID, limit + 1}
	}

	records, err = queryList(ctx, r, op, limit+1, query, scanBannedBid, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// BidderAddressForBid returns the indexed bidder address for a bid database
// identifier, or store.ErrNotFound when the bid does not exist.
func (r *Repo) BidderAddressForBid(ctx context.Context, bidID int64) (string, error) {
	if bidID < 1 {
		return "", fmt.Errorf("bidder address for bid: invalid bid id")
	}
	var address string
	err := r.pool().QueryRow(ctx, `SELECT a.addr
		FROM cg_bid AS b
		JOIN address AS a ON a.address_id = b.bidder_aid
		WHERE b.id = $1`, bidID).Scan(&address)
	if err != nil {
		return "", store.WrapError("bidder address for bid", err)
	}
	return address, nil
}

// CreateBannedBid creates one active ban and returns its persisted row.
// Banning an already-banned bid yields store.ErrConflict.
func (r *Repo) CreateBannedBid(
	ctx context.Context,
	bidID int64,
	userAddr string,
	bannedAt time.Time,
) (cgmodel.CGBannedBidRec, error) {
	if bidID < 1 || userAddr == "" || bannedAt.IsZero() || bannedAt.Unix() < 0 {
		return cgmodel.CGBannedBidRec{}, fmt.Errorf("create banned bid: invalid input")
	}
	var rec cgmodel.CGBannedBidRec
	err := r.pool().QueryRow(ctx, `INSERT INTO cg_banned_bids (bid_id, user_addr, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, bid_id, user_addr, created_at`,
		bidID, userAddr, bannedAt.Unix(),
	).Scan(&rec.Id, &rec.BidId, &rec.UserAddr, &rec.CreatedAt)
	if err != nil {
		return cgmodel.CGBannedBidRec{}, store.WrapError("create banned bid", err)
	}
	return rec, nil
}

// InsertBannedBid adds a row to cg_banned_bids with created_at set to the
// current Unix timestamp. Banning the same bid twice yields store.ErrConflict.
func (r *Repo) InsertBannedBid(ctx context.Context, bidID int64, userAddr string) error {
	_, err := r.CreateBannedBid(ctx, bidID, userAddr, time.Now())
	return err
}

// RemoveBannedBid removes the active ban for bidID. The boolean reports
// whether a row existed.
func (r *Repo) RemoveBannedBid(ctx context.Context, bidID int64) (bool, error) {
	if bidID < 1 {
		return false, fmt.Errorf("remove banned bid: invalid bid id")
	}
	tag, err := r.pool().Exec(ctx, "DELETE FROM cg_banned_bids WHERE bid_id = $1", bidID)
	if err != nil {
		return false, store.WrapError("remove banned bid", err)
	}
	return tag.RowsAffected() > 0, nil
}

// DeleteBannedBid removes all bans of the given bid id (unban).
func (r *Repo) DeleteBannedBid(ctx context.Context, bidID int64) error {
	_, err := r.pool().Exec(ctx, "DELETE FROM cg_banned_bids WHERE bid_id = $1", bidID)
	return store.WrapError("delete banned bid", err)
}
