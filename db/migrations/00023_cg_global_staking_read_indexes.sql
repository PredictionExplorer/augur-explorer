-- +goose NO TRANSACTION
-- +goose Up
-- API v2 global staking resolves action lifecycles by contract action ID,
-- pages per-round allocations by (deposit_id DESC, staker_aid), and pages
-- the two staker-raffle pools by immutable event-log ID.
CREATE UNIQUE INDEX CONCURRENTLY cg_nft_staked_cst_action_idx
	ON cg_nft_staked_cst (action_id);

CREATE UNIQUE INDEX CONCURRENTLY cg_nft_unstaked_cst_action_idx
	ON cg_nft_unstaked_cst (action_id);

CREATE UNIQUE INDEX CONCURRENTLY cg_nft_staked_rwalk_action_idx
	ON cg_nft_staked_rwalk (action_id);

CREATE UNIQUE INDEX CONCURRENTLY cg_nft_unstaked_rwalk_action_idx
	ON cg_nft_unstaked_rwalk (action_id);

CREATE INDEX CONCURRENTLY cg_staking_eth_deposit_round_deposit_idx
	ON cg_staking_eth_deposit (round_num, deposit_id DESC);

CREATE INDEX CONCURRENTLY cg_staker_deposit_deposit_staker_idx
	ON cg_staker_deposit (deposit_id DESC, staker_aid);

CREATE INDEX CONCURRENTLY cg_st_reward_deposit_collected_idx
	ON cg_st_reward (deposit_id, collected)
	INCLUDE (reward);

CREATE INDEX CONCURRENTLY cg_raffle_nft_prize_staker_pool_evt_idx
	ON cg_raffle_nft_prize (is_rwalk, evtlog_id DESC)
	WHERE is_staker=TRUE;

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_raffle_nft_prize_staker_pool_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_st_reward_deposit_collected_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_staker_deposit_deposit_staker_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_staking_eth_deposit_round_deposit_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_nft_unstaked_rwalk_action_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_nft_staked_rwalk_action_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_nft_unstaked_cst_action_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_nft_staked_cst_action_idx;
