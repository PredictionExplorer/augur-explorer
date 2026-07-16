-- +goose NO TRANSACTION
-- +goose Up
-- API v2 RandomWalk resources page the mint directory by (contract, token),
-- rank it by live trade count, page per-token event ledgers and the global
-- offer/trade/withdrawal ledgers by immutable event-log ID, serve the live
-- order book by price, resolve offer outcomes by (contract, offer) and scan
-- wallet-scoped offers and live ownership directly.
CREATE INDEX CONCURRENTLY rw_mint_evt_contract_token_idx
	ON rw_mint_evt (contract_aid, token_id);

CREATE INDEX CONCURRENTLY rw_mint_evt_owner_idx
	ON rw_mint_evt (owner_aid, contract_aid);

CREATE INDEX CONCURRENTLY rw_token_trades_rank_idx
	ON rw_token (rwalk_aid, num_trades DESC, token_id);

CREATE INDEX CONCURRENTLY rw_token_owner_idx
	ON rw_token (cur_owner_aid, rwalk_aid, token_id);

CREATE INDEX CONCURRENTLY rw_token_name_token_evt_idx
	ON rw_token_name (contract_aid, token_id, evtlog_id DESC);

CREATE INDEX CONCURRENTLY rw_transfer_token_evt_idx
	ON rw_transfer (contract_aid, token_id, evtlog_id DESC);

CREATE INDEX CONCURRENTLY rw_new_offer_token_evt_idx
	ON rw_new_offer (rwalk_aid, token_id, evtlog_id DESC);

CREATE INDEX CONCURRENTLY rw_new_offer_contract_evt_idx
	ON rw_new_offer (contract_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY rw_new_offer_active_price_idx
	ON rw_new_offer (contract_aid, price, evtlog_id)
	WHERE active;

CREATE INDEX CONCURRENTLY rw_new_offer_seller_evt_idx
	ON rw_new_offer (seller_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY rw_new_offer_buyer_evt_idx
	ON rw_new_offer (buyer_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY rw_new_offer_sell_time_idx
	ON rw_new_offer (contract_aid, time_stamp)
	INCLUDE (price)
	WHERE otype = 1;

CREATE INDEX CONCURRENTLY rw_item_bought_offer_idx
	ON rw_item_bought (contract_aid, offer_id);

CREATE INDEX CONCURRENTLY rw_item_bought_contract_evt_idx
	ON rw_item_bought (contract_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY rw_item_bought_time_idx
	ON rw_item_bought (contract_aid, time_stamp);

CREATE INDEX CONCURRENTLY rw_offer_canceled_offer_idx
	ON rw_offer_canceled (contract_aid, offer_id);

CREATE INDEX CONCURRENTLY rw_withdrawal_contract_evt_idx
	ON rw_withdrawal (contract_aid, evtlog_id DESC);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS rw_withdrawal_contract_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_offer_canceled_offer_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_item_bought_time_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_item_bought_contract_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_item_bought_offer_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_new_offer_sell_time_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_new_offer_buyer_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_new_offer_seller_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_new_offer_active_price_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_new_offer_contract_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_new_offer_token_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_transfer_token_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_token_name_token_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_token_owner_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_token_trades_rank_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_mint_evt_owner_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_mint_evt_contract_token_idx;
