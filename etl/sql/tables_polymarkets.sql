CREATE TABLE pol_cond_prep (	-- table to store ConditionPreparation event of conditional token (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
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
	evtlog_id			BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
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
	evtlog_id			BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	stakeholder_aid		BIGINT NOT NULL,
	collateral_aid		BIGINT NOT NULL,
	parent_coll_id		TEXT NOT NULL, -- parent collection id
	condition_id		TEXt NOT NULL,
	partition			TEXT NOT NULL,
	amount				DECIMAL NOT NULL,
	-- The following are linked ERC1155 transfers
	tok_ids				TEXT NOT NULL,	-- transferred token IDs (comma separated)
	tok_froms			TEXT NOT NULL,  -- the From field (comma separated)
	tok_tos				TEXT NOT NULL,	-- the To fields (comma separated)
	tok_amounts			TEXT NOT NULL,	-- Amount fileds (comma separated)
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_pos_merge ( -- PositionMerge, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	stakeholder_aid		BIGINT NOT NULL,
	collateral_aid		BIGINT NOT NULL,
	parent_coll_id		TEXT NOT NULL, -- parent collection id
	condition_id		TEXT NOT NULL,
	partition			TEXT NOT NULL,
	amount				DECIMAL NOT NULL,
	-- The following are linked ERC1155 transfers
	tok_ids				TEXT NOT NULL,	-- transferred token IDs (comma separated)
	tok_froms			TEXT NOT NULL,  -- the From field (comma separated)
	tok_tos				TEXT NOT NULL,	-- the To fields (comma separated)
	tok_amounts			TEXT NOT NULL,	-- Amount fileds (comma separated)
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_pay_redem (-- PayoutRedemption, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
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
	-- The following are linked ERC1155 transfers
	tok_ids				TEXT NOT NULL,	-- transferred token IDs (comma separated)
	tok_froms			TEXT NOT NULL,  -- the From field (comma separated)
	tok_tos				TEXT NOT NULL,	-- the To fields (comma separated)
	tok_amounts			TEXT NOT NULL,	-- Amount fileds (comma separated)
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_tok_id_ops ( -- Token IDs that correspond to position merge/position split/payout redeem
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	tx_id				BIGINT,
	parent_split_id		BIGINT REFERENCES pol_pos_split(id) ON DELETE CASCADE,
	parent_merge_id		BIGINT REFERENCES pol_pos_merge(id) ON DELETE CASCADE,
	parent_redeem_id	BIGINT REFERENCES pol_pay_redem(id) ON DELETE CASCADE,
	contract_aid		BIGINT,
	outcome_idx			INT NOT NULL,
	condition_id		TEXT NOT NULL, -- will be null for FixedProductMarkeMaker, and non-null for cond. token
	token_id_hex		TEXT NOT NULL,
	token_from			TEXT NOT NULL,
	token_to			TEXT NOT NULL,
	token_amount		DECIMAL NOT NULL,
	UNIQUE(evtlog_id,token_id_hex)

);
CREATE TABLE pol_tok_ids (	-- table that collects only unique token_IDs per market
	contract_aid		BIGINT NOT NULL,
	outcome_idx			INT NOT NULL,
	token_id_hex		TEXT PRIMARY KEY
);
CREATE TABLE pol_uri ( -- URI, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	value				TEXT,
	uri_id				DECIMAL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_fund_addrem ( -- FPMMFundAdded event of contract FixedProductMarketMaker
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	funder_aid			BIGINT NOT NULL,
	op_type				SMALLINT NOT NULL, -- 0 - Add, 1 - Remove
	amounts				TEXT NOT NULL,		-- amounts added or removed, comma separated
	sum_amounts			DECIMAL NOT NULL,	-- all the amounts summed
	collateral_removed	DECIMAL DEFAULT 0,
	norm_collateral		DECIMAL DEFAULT 0,	-- normalized collateral amount (negative for puts, positive for gets)
	shares				DECIMAL NOT NULL,	-- shares minted or burned
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_buysell ( -- FPMMBuy/FPMMSell event of contract FixedProductMarketMaker
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	user_aid			BIGINT NOT NULL,
	op_type				SMALLINT NOT NULL,	-- 0- buy, 1 - sell
	outcome_idx			SMALLINT NOT NULL,
	collateral_amount	DECIMAL NOT NULL,	-- amount as it comes from the event log (types.Log)
	normalized_amount	DECIMAL DEFAULT 0,	-- negative for deposits, positive for withdrawals
	fee_amount			DECIMAL NOT NULL,
	token_amount		DECIMAL NOT NULL,
	accum_collateral	DECIMAL DEFAULT 0,
	UNIQUE(evtlog_id)
);
/*
CREATE TABLE pol_oi_hist ( -- Open interest change history
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	user_aid			BIGINT NOT NULL,
	market_id			INT NOT NULL,
	
	PRIMARY KEY(user_aid,market_id,condition_id)
);*/
CREATE TABLE update_needed (	-- used to flag the market fetching process (polysync) to update markets
	market_update		BOOLEAN DEFAULT FALSE
);
CREATE table pol_market ( -- As received from https://strapi-matic.poly.market/markets
	id							BIGSERIAL PRIMARY KEY,
	market_id					INT NOT NULL,
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
CREATE TABLE pol_mkt_words(-- search tokens for searching markets by description/title
	id					BIGSERIAL PRIMARY KEY,
	market_id			BIGINT NOT NULL,
	contract_aid		BIGINT NOT NULL,
	tok_type			SMALLINT DEFAULT 0,				-- 0-description 1 - title
	tokens				TSVECTOR
);
CREATE TABLE pol_proc_status (
	last_evt_id				BIGINT DEFAULT 0,
	last_block				BIGINT DEFAULT 0 -- used when getting event logs via ethclient.FilterLogs
);
CREATE TABLE pol_mkt_stats ( -- market statistics
	contract_aid			BIGINT PRIMARY KEY,
	num_liq_ops				INT DEFAULT 0, -- number of liquidity addition/deletions
	num_trades				INT DEFAULT 0,
	open_interest			DECIMAL DEFAULT 0,	-- BUYs + Add of liquidity totalled
	total_volume			DECIMAL DEFAULT 0,
	total_fees				DECIMAL DEFAULT 0,		-- sums amount of fees paid for this market
	total_liquidity			DECIMAL DEFAULT 0		-- amount of USDC held (without fees) (without BUYs)
);
CREATE TABLE pol_ustats ( -- user statistics
	user_aid				BIGINT PRIMARY KEY,
	reg_time_stamp			TIMESTAMPTZ NOT NULL,
	markets_count			INT DEFAULT 0, -- total count of markets traded
	tot_trades				INT DEFAULT 0, -- total amount of buy/sell operations
	tot_liq_ops				INT DEFAULT 0, -- total amount of liquidity add/remove operations
	tot_volume				DECIMAL DEFAULT 0, -- total trading volume for this user in collateral
	tot_liq_given			DECIMAL DEFAULT 0, -- total of invested liequidity
	tot_fees				DECIMAL DEFAULT 0, -- total fees paid
	profit					DECIMAL DEFAULT 0,	-- profit of the user made in collateral
	UNIQUE(user_aid)
);
CREATE TABLE pol_ustats_mkt (-- user statistics per specific market
	user_aid				BIGINT NOT NULL,
	contract_aid			INT NOT NULL,	-- Fixed Product Market Maker
	tot_trades				INT DEFAULT 0,
	tot_liq_ops				INT DEFAULT 0,
	tot_volume				DECIMAL DEFAULT 0,
	tot_liq_given			DECIMAL DEFAULT 0, -- total of invested liequidity
	tot_fees				DECIMAL DEFAULT 0,  -- accumulated amount of fees paid by this user
	profit					DECIMAL DEFAULT 0, -- profits made by the user in terms of collateral token
	UNIQUE(contract_aid,user_aid)
);
CREATE TABLE pol_unique_addrs (	-- Unique addresses per day, statistics
	day					DATE PRIMARY KEY,
	num_addrs			INT DEFAULT 0,	-- Total number of uinique users (traders + liquidity providers)
	num_funders			INT DEFAULT 0,	-- Number of unique liquidity providers (pol_addrem table)
	num_traders			INT DEFAULT 0	-- Number of unique traders (pol_buysell table)
);
CREATE TABLE pol_data_feed (
	last_evt_id			BIGINT DEFAULT 0	-- stores last event ID returned by data feed
);
CREATE TABLE pol_uranks (   -- User Rankings (how this user ranks against each other, ex: Top 13% in profit made
	aid		            BIGINT PRIMARY KEY,
	total_trades		BIGINT DEFAULT 0,
	top_profit          DECIMAL(5,2) DEFAULT 100.0,    -- position of the user in profits accumulated over lifetime
	top_trades          DECIMAL(5,2) DEFAULT 100.0,    -- position of the user in number of accumulated trades
	top_volume			DECIMAL(5,2) DEFAULT 100.0,	   -- position of the user in accumulated trading volume
	profit				DECIMAL(32,18) DEFAULT 0.0,
	volume				DECIMAL(32,18) DEFAULT 0.0
);
CREATE TABLE pol_contracts (
	cond_tok_addr		TEXT DEFAULT '0x4D97DCd97eC945f40cF65F87097ACe5EA0476045',
	usdc_addr			TEXT DEFAULT '0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174'
);
