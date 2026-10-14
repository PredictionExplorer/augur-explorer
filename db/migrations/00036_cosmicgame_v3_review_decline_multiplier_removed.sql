-- +goose Up
-- Solidity commit 383e0650 (branch claude/v3-contract-review-dc3c65) removed
-- the configurable CST bid price decline multiplier from the V3 contracts:
-- the CST Dutch auction decline rate now derives from the bid CST reward rate
-- (bidCstRewardAmountMultiplier / mainPrizeTimeIncrementInMicroSeconds), so
-- cstBidPriceDeclineMultiplier, its change divisor, their setters and the
-- CstBidPriceDeclineMultiplierChanged / CstBidPriceDeclineMultiplierChangeDivisorChanged
-- events no longer exist. V3 never reached production, so the tables are
-- empty (same situation as migration 00030, which dropped the same pair for
-- an earlier contract revision before 00032 re-created them).

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
