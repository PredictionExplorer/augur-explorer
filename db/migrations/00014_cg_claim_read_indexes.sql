-- +goose NO TRANSACTION
-- +goose Up
-- API v2 merges claim events by immutable event-log identity within a round.
CREATE INDEX CONCURRENTLY cg_prize_withdrawal_round_evt_idx
	ON cg_prize_withdrawal (round_num, evtlog_id);

CREATE INDEX CONCURRENTLY cg_donated_nft_claimed_round_evt_idx
	ON cg_donated_nft_claimed (round_num, evtlog_id);

CREATE INDEX CONCURRENTLY cg_donated_tok_claimed_round_evt_idx
	ON cg_donated_tok_claimed (round_num, evtlog_id);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_donated_tok_claimed_round_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_donated_nft_claimed_round_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_prize_withdrawal_round_evt_idx;
