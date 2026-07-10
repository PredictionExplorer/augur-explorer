-- +goose NO TRANSACTION
-- +goose Up
-- Bounded API v2 bidding analytics filter cg_bid by a half-open timestamp
-- range before assigning rows to zero-filled buckets.
CREATE INDEX CONCURRENTLY cg_bid_analytics_time_idx
	ON cg_bid (time_stamp);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_bid_analytics_time_idx;
