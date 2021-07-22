CREATE TABLE pol_cond_prep (	-- table to store ConditionPreparation event of conditional token (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	oracle_aid			BIGINT NOT NULL,
	condition_id		TEXT NOT NULL,
	question_id			TEXT NOT NULL,
	outcome_slot_count	DECIMAL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_cond_res (-- ConditionResolution, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	oracle_aid			BIGINT NOT NULL,
	condition_id		TEXT NOT NULL,
	question_id			TEXT NOT NULL,
	outcome_slot_count	DECIMAL,
	payout_numerators	TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_pos_split ( -- PositionSplit, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	stakeholder_aid		BIGINT NOT NULL,
	collateral_aid		BIGINT NOT NULL,
	parent_coll_id		BIGINT NOT NULL, -- parent collection id
	condition_id		TEXt NOT NULL,
	partition			TEXT NOT NULL,
	amount				DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_pos_merge ( -- PositionMerge, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	stakeholder_aid		BIGINT NOT NULL,
	collateral_aid		BIGINT NOT NULL,
	parent_coll_id		BIGINT NOT NULL, -- parent collection id
	condition_id		TEXT NOT NULL,
	partition			TEXT NOT NULL,
	amount				DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_pay_redem (-- PayoutRedemption, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	redeemer_aid		BIGINT NOT NULL,
	collateral_aid		TEXT NOT NULL,
	parent_coll_id		TEXT NOT NULL,
	condition_id		TEXT NOT NULL,
	index_sets			TEXT NOT NULL,
	payout				DECIMAL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_uri ( -- URI, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	value				TEXT,
	uri_id				DECIMAL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_fund_add ( -- FPMMFundAdded event of contract FixedProductMarketMaker
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	funder_aid			BIGINT NOT NULL,
	amounts_added		TEXT NOT NULL,
	shares_minted		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_fund_rem ( -- FPMMFundRemoved event of contract FixedProductMarketMaker
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	funder_aid			BIGINT NOT NULL,
	amounts_removed		TEXT NOT NULL,
	shares_burnt		DECIMAL NOT NULL,
	collateral_removed	DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_buysell ( -- FPMMBuy/FPMMSell event of contract FixedProductMarketMaker
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	user_aid			BIGINT NOT NULL,
	op_type				SMALLINT NOT NULL,	-- 0- buy, 1 - sell
	outcome_idx			SMALLINT NOT NULL,
	collateral_amount	DECIMAL NOT NULL,
	fee_amount			DECIMAL NOT NULL,
	token_amount		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
/* DISCONTINUED
CREATE TABLE pol_sell ( -- FPMMSell event of contract FixedProductMarketMaker
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	seller_aid			BIGINT NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	return_amount		DECIMAL NOT NULL,
	fee_amount			DECIMAL NOT NULL,
	tokens_sold			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
*/
CREATE TABLE pol_mkt_stats ( -- market statistics
	market_id				BIGINT PRIMARY KEY,
	open_interest			DECIMAL,
	num_liquidity_ops		INT DEFAULT 0, -- number of liquidity addition/deletions
	num_trades				INT DEFAULT 0,
	total_volume			DECIMAL,
	total_fees				DECIMAL		-- sums amount of fees paid for this market
);
CREATE TABLE update_needed (	-- used to flag the market fetching process (polysync) to update markets
	market_update		BOOLEAN DEFAULT FALSE
);
CREATE table pol_market ( -- As received from https://strapi-matic.poly.market/markets
	id							BIGSERIAL PRIMARY KEY,
	market_id					BIGINT NOT NULL,
	question					TEXT NOT NULL,
	condition_id				TEXT NOT NULL,
	slug						TEXT NOT NULL,
	twitter_card_image			TEXT DEFAULT '',
	resolution_source			TEXT NOT NULL,
	end_date					TEXT NOT NULL,
	end_date_ts					TIMESTAMPTZ NOT NULL,
	category					TEXT NOT NULL,
	amm_type					TEXT NOT NULL,
	sponsor_name				TEXT NOT NULL,
	sponsor_image				TEXT NOT NULL,
	start_date					TEXT NOT NULL,
	start_date_ts				TIMESTAMPTZ NOT NULL,
	x_axis_value				BIGINT DEFAULT 0,
	y_axis_value				BIGINT DEFAULT 0,
	denomination_token			TEXT NOT NULL,
	fee							DECIMAL NOT NULL,
	image						TEXT NOT NULL,
	icon						TEXT NOT NULL,
	lower_bound					DECIMAL DEFAULT 0,
	upper_bound					DECIMAL DEFAULT 0,
	description					TEXT NOT NULL,
	tags						TEXT DEFAULT '',
	outcomes					TEXT NOT NULL,
	outcome_prices				TEXT NOT NULL,
	volume						DECIMAL NOT NULL,
	active						BOOLEAN DEFAULT TRUE,
	market_type					TEXT NOT NULL,
	market_type_code			SMALLINT NOT NULL,
	format_type					TEXT DEFAULT '',
	lower_bound_date			TEXT NOT NULL,
	lower_bound_ts				TIMESTAMPTZ NOT NULL,
	upper_bound_date			TEXT NOT NULL,
	upper_bound_ts				TIMESTAMPTZ NOT NULL,
	closed						BOOLEAN DEFAULT FALSE,
	mkt_mkr_aid					BIGINT NOT NULL,-- Market Maker Address
	created_at_date				TEXT NOT NULL,
	created_at_ts				TIMESTAMPTZ NOT NULL,
	updated_at_date				TEXT NOT NULL,
	updated_at_ts				TIMESTAMPTZ NOT NULL,
	closed_time					TEXT NOT NULL,
	closed_time_ts				TIMESTAMPTZ NOT NULL,
	wide_format					BOOLEAN NOT NULL,
	new							BOOLEAN NOT NULL,
	sent_discord				BOOLEAN NOT NULL,
	mailchimp_tag				TEXT DEFAULT '',
	featured					BOOLEAN NOT NULL,
	submitted_by				TEXT DEFAULT '',
	subcategory					TEXT DEFAULT '',
	category_mailchimp_tag		TEXT DEFAULT '',
	use_cases					TEXT DEFAULT '',
	liquidity					DECIMAL DEFAULT 0,

	UNIQUE(market_id)
);
CREATE TABLE poly_proc_status (
	last_evt_id			BIGINT DEFAULT 0
);
