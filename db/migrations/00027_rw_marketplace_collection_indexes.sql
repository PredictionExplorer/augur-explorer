-- +goose NO TRANSACTION
-- +goose Up
-- The marketplace contract serves multiple NFT collections. Collection-first
-- keys keep offer-history and live price/floor scans proportional to the
-- requested collection instead of every offer on the marketplace.
CREATE INDEX CONCURRENTLY rw_new_offer_collection_evt_idx
	ON rw_new_offer (contract_aid, rwalk_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY rw_new_offer_collection_active_price_idx
	ON rw_new_offer (contract_aid, rwalk_aid, price, evtlog_id)
	WHERE active;

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS rw_new_offer_collection_active_price_idx;
DROP INDEX CONCURRENTLY IF EXISTS rw_new_offer_collection_evt_idx;
