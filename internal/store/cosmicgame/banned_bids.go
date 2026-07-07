package cosmicgame

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	p "github.com/PredictionExplorer/augur-explorer/internal/primitives/cosmicgame"
	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// BannedBids returns all rows from cg_banned_bids ordered by id.
func (r *Repo) BannedBids(ctx context.Context) ([]p.CGBannedBidRec, error) {
	query := "SELECT id, bid_id, user_addr, created_at FROM cg_banned_bids ORDER BY id"
	scan := func(rows pgx.Rows, rec *p.CGBannedBidRec) error {
		return rows.Scan(&rec.Id, &rec.BidId, &rec.UserAddr, &rec.CreatedAt)
	}
	return queryList(ctx, r, "banned bids", 8, query, scan)
}

// InsertBannedBid adds a row to cg_banned_bids with created_at set to the
// current Unix timestamp. Banning the same bid twice yields store.ErrConflict.
func (r *Repo) InsertBannedBid(ctx context.Context, bidID int64, userAddr string) error {
	query := "INSERT INTO cg_banned_bids (bid_id, user_addr, created_at) VALUES ($1, $2, $3)"
	_, err := r.pool().Exec(ctx, query, bidID, userAddr, time.Now().Unix())
	return store.WrapError("insert banned bid", err)
}

// DeleteBannedBid removes all bans of the given bid id (unban).
func (r *Repo) DeleteBannedBid(ctx context.Context, bidID int64) error {
	query := "DELETE FROM cg_banned_bids WHERE bid_id = $1"
	_, err := r.pool().Exec(ctx, query, bidID)
	return store.WrapError("delete banned bid", err)
}
