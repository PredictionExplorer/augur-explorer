-- +goose Up
-- Solidity commit 0bc80af0 (branch v3-2026-07-24) reverted the V3 CST Dutch
-- auction duration to the exact V2 behavior: cstBidPriceDeclineMultiplier and
-- its change divisor no longer exist on the contract, so the two events whose
-- history these tables hold can never be emitted again. V3 is not in
-- production, so the tables are empty. The V2 duration events
-- (CstDutchAuctionDurationChanged, CstDutchAuctionDurationChangeDivisorChanged)
-- are live again on V3 and keep using their existing V2 tables.

DROP TABLE cg_adm_cst_price_decline_mul_div;
DROP TABLE cg_adm_cst_price_decline_mul;

-- +goose Down
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
