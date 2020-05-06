-- Blockchain tables
CREATE TABLE block (
	block_num			BIGINT NOT NULL UNIQUE,
	num_tx				BIGINT DEFAULT 0,
	block_hash			TEXT NOT NULL PRIMARY KEY,
	parent_hash			TEXT NOT NULL
);
CREATE TABLE transaction (	-- we're only storing transactions related to Augur platform
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	tx_hash				TEXT NOT NULL UNIQUE
);
-- Universe: The container contract for Augur Service
CREATE TABLE universe (
	universe_id			BIGSERIAL PRIMARY KEY,
	universe_addr		TEXT NOT NULL UNIQUE		-- Ethereum address of the Universe contract
);
CREATE TABLE address (
	address_id			BIGSERIAL	PRIMARY KEY,
	addr				TEXT NOT NULL UNIQUE		-- 20 byte Ethereum address , stored as 42 hex string (0x+addr)
);
-- Market category
CREATE TABLE category (
	cat_id				BIGSERIAL	PRIMARY KEY,
	category			TEXT NOT NULL UNIQUE		-- includes parent category too (comma separated list)
);
-- Market
CREATE TABLE market (
	market_aid			BIGINT NOT NULL PRIMARY KEY,-- address ID of the Market
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	cat_id				BIGINT NOT NULL,			-- category id
	universe_id			BIGSERIAL NOT NULL,			-- reference to universe table
	creator_aid			BIGINT NOT NULL,			-- address ID of market creator (not real User)
	signer_aid			BIGINT NOT NULL,			-- address ID of the User who signed transaction (real  creator)
	reporter_aid		BIGINT NOT NULL,			-- address ID of the User who will report on the outcome
	end_time			BIGINT NOT NULL,			-- when the Market expires
	max_ticks			BIGINT NOT NULL,			-- maximum price range (number of intervals)
	create_timestamp	BIGINT NOT NULL,
	market_type			SMALLINT NOT NULL,			-- Market type enum: 0:YES_NO | 1:CATEGORICAL | 2:SCALAR
	open_interest		TEXT DEFAULT '',			-- amount of shares created
	fee					TEXT NOT NULL,				-- fee to be paid to Market creator as percentage of transaction
	prices				TEXT NOT NULL,			-- range of prices the Market can take
	extra_info			TEXT NOT NULL,				-- specific market metadata (JSON format)
	outcomes			TEXT NOT NULL,				-- possible outcomes of the market
	winning_payouts		TEXT DEFAULT '',
	fin_timestamp		BIGINT DEFAULT 0,
	no_show_bond		TEXT NOT NULL,				-- $ penalty to the Creator for failing to emit report
	cur_volume			TEXT DEFAULT ''
);
-- Balances of Share tokens per Market (accumulated data, one record per account)
CREATE TABLE sbalances (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			 -- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	account_aid			BIGINT NOT NULL,			-- address id of the User(holder of the shares)
	market_aid			BIGINT NOT NULL,			-- market id of the Market these shares blong
	outcome				TEXT NOT NULL,				-- market outcome
	balance				TEXT NOT NULL				-- balance of shares (bigint as string)
);
-- Market Order (BUY/SELL request made by the User via GUI)
CREATE TABLE mktord (-- in this table only 'Fill' type orders are stored (Create/Cancel are temporary)
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL,
	signer_aid			BIGINT NOT NULL,			-- Address of the user who signes the transaction
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	oaction				SMALLINT NOT NULL,			-- order action:  0=>Create, 1=>Cancel, 2=>Fill
													-- Create: User posts a BID or ASK execpting to be filed
													-- Fill: User buys or sells existing (Created) order
													-- Cancel: User removes active order (BID/ASK)
	otype				SMALLINT NOT NULL,			-- enum:  0 => BID, 1 => ASK
	creator_aid			BIGINT NOT NULL,			-- address of the creator
	filler_aid			BIGINT NOT NULL,			-- address of the filler; source: AugurTrading.sol:24
	price				BIGINT NOT NULL,
	amount				BIGINT NOT NULL,
	outcome				BIGINT NOT NULL,
	token_refund		TEXT NOT NULL,
	shares_refund		TEXT NOT NULL,
	fees				TEXT NOT NULL,
	amount_filled		TEXT NOT NULL,
	time_stamp			BIGINT NOT NULL,
	shares_escrowed		TEXT NOT NULL,
	tokens_escrowed		TEXT NOT NULL,
	trade_group			TEXT NOT NULL,			-- User defined group label to identify multiple trades
	order_id			TEXT NOT NULL
);
CREATE TABLE oorders (	-- open orders table mirrors `mktord` table, it's used only for active orders
	-- this table is currently disabled until 0x Mesh trading is integrated
	id					BIGSERIAL PRIMARY KEY,
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGSERIAL NOT NULL,
	otype				SMALLINT NOT NULL,			-- enum:  0 => BID, 1 => ASK
	creator_aid			BIGINT NOT NULL,			-- address of the creator
	price				BIGINT NOT NULL,
	amount				BIGINT NOT NULL,
	outcome				BIGINT NOT NULL,
	time_stamp			BIGINT NOT NULL,
	order_id			TEXT NOT NULL
);
-- Report, submitted by Market Creator
CREATE TABLE report (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	reporter_aid		BIGINT NOT NULL,
	signer_aid			BIGINT NOT NULL,			-- transaction signer (the User who is submitting the report)
	ini_reporter_aid	BIGINT DEFAULT 0,
	disputed_aid		BIGINT DEFAULT 0,
	dispute_round		BIGINT DEFAULT 1,
	is_initial			BOOLEAN DEFAULT false,
	is_designated		BOOLEAN DEFAULT false,
	amount_staked		TEXT NOT NULL,
	pnumerators			TEXT NOT NULL,		-- payout numerators
	description			TEXT DEFAULT '',
	current_stake		TEXT DEFAULT '',
	stake_remaining		TEXT DEFAULT '',
	next_win_start		BIGINT,
	next_win_end		BIGINT,
	rpt_timestamp		BIGINT
);
-- Volume
CREATE TABLE volume (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	volume				TEXT NOT NULL,
	outcome_vols		TEXT NOT NULL,
	ins_timestamp		BIGINT NOT NULL
);
CREATE table oi_chg ( -- open interest changed event
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL REFERENCES market(market_aid) ON DELETE CASCADE,
	ts_inserted			BIGINT NOT NULL, -- timestamp
	oi					TEXT NOT NULL
);
CREATE TABLE mkt_fin (
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL REFERENCES market(market_aid) ON DELETE CASCADE,
	fin_timestamp		BIGINT NOT NULL,
	winning_payouts		TEXT NOT NULL
);
CREATE TABLE last_block (
	block_num			BIGINT	NOT NULL	-- last block processed by the ETL
);

