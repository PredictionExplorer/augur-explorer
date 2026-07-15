-- +goose NO TRANSACTION
-- +goose Up
-- API v2 user staking histories page one wallet's stake/unstake events,
-- currently staked tokens and reward ledgers. The action histories key on
-- (staker, evtlog_id DESC); the staked-token collections key on
-- (staker, token_id); the smallest reward units are read per
-- (staker, deposit) and per (staker, token); deposit existence and the
-- staker-deposit join resolve through cg_staking_eth_deposit.deposit_id.
CREATE INDEX CONCURRENTLY cg_nft_staked_cst_staker_evt_idx
	ON cg_nft_staked_cst (staker_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_nft_unstaked_cst_staker_evt_idx
	ON cg_nft_unstaked_cst (staker_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_nft_staked_rwalk_staker_evt_idx
	ON cg_nft_staked_rwalk (staker_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_nft_unstaked_rwalk_staker_evt_idx
	ON cg_nft_unstaked_rwalk (staker_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_staked_token_cst_staker_token_idx
	ON cg_staked_token_cst (staker_aid, token_id);

CREATE INDEX CONCURRENTLY cg_staked_token_rwalk_staker_token_idx
	ON cg_staked_token_rwalk (staker_aid, token_id);

CREATE INDEX CONCURRENTLY cg_st_reward_staker_deposit_action_idx
	ON cg_st_reward (staker_aid, deposit_id, action_id);

CREATE INDEX CONCURRENTLY cg_st_reward_staker_token_deposit_idx
	ON cg_st_reward (staker_aid, token_id, deposit_id);

CREATE INDEX CONCURRENTLY cg_staking_eth_deposit_deposit_idx
	ON cg_staking_eth_deposit (deposit_id);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_staking_eth_deposit_deposit_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_st_reward_staker_token_deposit_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_st_reward_staker_deposit_action_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_staked_token_rwalk_staker_token_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_staked_token_cst_staker_token_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_nft_unstaked_rwalk_staker_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_nft_staked_rwalk_staker_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_nft_unstaked_cst_staker_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_nft_staked_cst_staker_evt_idx;
