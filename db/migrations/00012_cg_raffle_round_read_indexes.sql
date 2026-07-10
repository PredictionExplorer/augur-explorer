-- +goose NO TRANSACTION
-- +goose Up
-- API v2 pages completed-round raffle payout records by immutable public
-- winner/event identities rather than database row ids.
CREATE INDEX CONCURRENTLY cg_prize_deposit_round_winner_evt_idx
	ON cg_prize_deposit (round_num, winner_index, evtlog_id);

CREATE INDEX CONCURRENTLY cg_raffle_nft_prize_round_pool_winner_evt_idx
	ON cg_raffle_nft_prize (round_num, is_staker, winner_idx, evtlog_id);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_raffle_nft_prize_round_pool_winner_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_prize_deposit_round_winner_evt_idx;
