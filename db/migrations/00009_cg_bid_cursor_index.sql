-- +goose NO TRANSACTION
-- +goose Up
-- API v2 pages round bids by this immutable composite. Keep the index
-- non-unique: the contract should assign one bid_position per round, but the
-- evtlog_id tie-breaker makes reads deterministic even if historical data is
-- corrupt and avoids turning deployment into an unreviewed data cleanup.
CREATE INDEX CONCURRENTLY cg_bid_round_position_evtlog_idx
	ON cg_bid (round_num, bid_position, evtlog_id);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_bid_round_position_evtlog_idx;
