-- +goose Up
-- The v3.1 contracts (Cosmic-Signature-v3.1-2026-08-19) reinstate the CST bid
-- price decline multiplier mechanics that migration 00030 reverted: V3 prices
-- CST bids with cstBidPriceDeclineMultiplier (wei/second) instead of the V2
-- cstDutchAuctionDuration, and emits CstBidPriceDeclineMultiplierChanged /
-- CstBidPriceDeclineMultiplierChangeDivisorChanged admin events. Re-create the
-- two history tables dropped by 00030 (they were empty when dropped; V3 never
-- reached production).

CREATE TABLE cg_adm_cst_price_decline_mul (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	new_value		DECIMAL NOT NULL,
	UNIQUE (evtlog_id)
);

CREATE TABLE cg_adm_cst_price_decline_mul_div (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	new_value		DECIMAL NOT NULL,
	UNIQUE (evtlog_id)
);

-- +goose Down
DROP TABLE cg_adm_cst_price_decline_mul_div;
DROP TABLE cg_adm_cst_price_decline_mul;
