-- +goose NO TRANSACTION
-- +goose Up
-- API v2 lists completed rounds by this stable descending composite. Keep
-- the index non-unique because historical integrity is not a deployment-time
-- migration concern; evtlog_id remains the deterministic tie-breaker.
CREATE INDEX CONCURRENTLY cg_prize_claim_round_evtlog_idx
	ON cg_prize_claim (round_num DESC, evtlog_id DESC);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_prize_claim_round_evtlog_idx;
