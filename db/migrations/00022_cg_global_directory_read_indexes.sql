-- +goose NO TRANSACTION
-- +goose Up
-- API v2 global directories page one token's renames and ownership history
-- on (token_id, evtlog_id DESC), the named-token directory on token_id over
-- the named subset, and the Cosmic Token balance directory on the
-- (cur_balance DESC, owner_aid) keyset over positive balances.
CREATE INDEX CONCURRENTLY cg_token_name_token_evt_idx
	ON cg_token_name (token_id, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_erc721_transfer_token_evt_idx
	ON cg_erc721_transfer (token_id, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_mint_event_named_token_idx
	ON cg_mint_event (token_id)
	WHERE LENGTH(token_name) > 0;

CREATE INDEX CONCURRENTLY cg_costok_owner_balance_idx
	ON cg_costok_owner (cur_balance DESC, owner_aid)
	WHERE cur_balance > 0;

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_costok_owner_balance_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_mint_event_named_token_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_erc721_transfer_token_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_token_name_token_evt_idx;
