-- +goose NO TRANSACTION
-- +goose Up
-- API v2 pages one round's unified prize registry by this immutable
-- ascending composite. cg_prize's primary key uses winner_index before
-- ptype, so this read-order index avoids a sort while retaining the existing
-- schema identity.
CREATE INDEX CONCURRENTLY cg_prize_round_type_winner_idx
	ON cg_prize (round_num, ptype, winner_index);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_prize_round_type_winner_idx;
