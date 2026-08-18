-- +goose Up
-- V3 (v3-2026-07-24 contracts) replaced the last-bidder bid CST reward
-- percentage split with an all-to-previous-bidder reward, and replaced the
-- stored CST Dutch auction duration with a direct price decline multiplier.
-- The retired LastBidderBidCstRewardAmountPercentageChanged event can no
-- longer be emitted (V3 never reached production, so its table is empty);
-- the two new admin events get their own history tables.

DROP TABLE cg_adm_last_bidder_reward_pct;

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

CREATE TABLE cg_adm_last_bidder_reward_pct (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	new_value		DECIMAL NOT NULL CHECK (new_value BETWEEN 0 AND 100),
	UNIQUE (evtlog_id)
);
