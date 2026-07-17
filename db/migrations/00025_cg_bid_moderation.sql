-- +goose Up
-- A bid ban is an addressable active resource in API v2, so at most one row
-- may exist for a bid. Historical duplicates are semantically equivalent;
-- retain the newest row before enforcing the invariant.
DELETE FROM cg_banned_bids AS older
USING cg_banned_bids AS newer
WHERE older.bid_id = newer.bid_id
  AND older.id < newer.id;

DROP INDEX IF EXISTS idx_cg_banned_bids_bid_id;
CREATE UNIQUE INDEX idx_cg_banned_bids_bid_id
	ON cg_banned_bids (bid_id);

-- +goose Down
DROP INDEX IF EXISTS idx_cg_banned_bids_bid_id;
CREATE INDEX idx_cg_banned_bids_bid_id
	ON cg_banned_bids (bid_id);
