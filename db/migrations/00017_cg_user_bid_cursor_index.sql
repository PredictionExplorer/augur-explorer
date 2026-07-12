-- +goose NO TRANSACTION
-- +goose Up
-- API v2 user bid histories page newest-first within one bidder.
CREATE INDEX CONCURRENTLY cg_bid_user_evtlog_idx
	ON cg_bid (bidder_aid, evtlog_id DESC);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_bid_user_evtlog_idx;
