-- +goose NO TRANSACTION
-- +goose Up
-- API v2 user activity pages one wallet's token transfers, marketing
-- rewards and owned Cosmic Signature tokens. The transfer ledgers key on
-- (wallet, evtlog_id DESC) per side and merge the two bounded scans; the
-- owned-token directory keys on (cur_owner_aid, token_id) ascending.
CREATE INDEX CONCURRENTLY cg_erc721_transfer_from_evt_idx
	ON cg_erc721_transfer (from_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_erc721_transfer_to_evt_idx
	ON cg_erc721_transfer (to_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_erc20_transfer_from_evt_idx
	ON cg_erc20_transfer (from_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_erc20_transfer_to_evt_idx
	ON cg_erc20_transfer (to_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_mkt_reward_marketer_evt_idx
	ON cg_mkt_reward (marketer_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_mint_event_owner_token_idx
	ON cg_mint_event (cur_owner_aid, token_id);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_mint_event_owner_token_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_mkt_reward_marketer_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_erc20_transfer_to_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_erc20_transfer_from_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_erc721_transfer_to_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_erc721_transfer_from_evt_idx;
