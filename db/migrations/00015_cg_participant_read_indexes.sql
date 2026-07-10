-- +goose NO TRANSACTION
-- +goose Up
-- API v2 participant directories use stable descending aggregate keysets
-- with the participant address ID as the ascending tie-breaker.
CREATE INDEX CONCURRENTLY cg_bidder_participant_page_idx
	ON cg_bidder (num_bids DESC, bidder_aid)
	WHERE num_bids > 0;

CREATE INDEX CONCURRENTLY cg_donor_participant_page_idx
	ON cg_donor (COALESCE(total_eth_donated, 0) DESC, donor_aid)
	WHERE count_donations > 0;

CREATE INDEX CONCURRENTLY cg_staker_cst_participant_page_idx
	ON cg_staker_cst (COALESCE(total_reward, 0) DESC, staker_aid)
	WHERE num_stake_actions > 0;

CREATE INDEX CONCURRENTLY cg_staker_rwalk_participant_page_idx
	ON cg_staker_rwalk (COALESCE(total_tokens_staked, 0) DESC, staker_aid)
	WHERE num_stake_actions > 0;

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_staker_rwalk_participant_page_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_staker_cst_participant_page_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_donor_participant_page_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_bidder_participant_page_idx;
