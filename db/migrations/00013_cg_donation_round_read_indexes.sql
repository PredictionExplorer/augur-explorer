-- +goose NO TRANSACTION
-- +goose Up
-- API v2 pages round donation events newest first by their immutable public
-- event-log identities. ETH donations live in two event-specific tables.
CREATE INDEX CONCURRENTLY cg_eth_donated_round_evt_idx
	ON cg_eth_donated (round_num, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_eth_donated_wi_round_evt_idx
	ON cg_eth_donated_wi (round_num, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_erc20_donation_round_evt_idx
	ON cg_erc20_donation (round_num, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_nft_donation_round_evt_idx
	ON cg_nft_donation (round_num, evtlog_id DESC);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_nft_donation_round_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_erc20_donation_round_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_eth_donated_wi_round_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_eth_donated_round_evt_idx;
