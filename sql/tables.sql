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
	creator_aid			BIGINT NOT NULL,			-- address ID of the contract wallet of the User
	signer_aid			BIGINT NOT NULL,			-- address ID of the User (EOA) who created the market
	reporter_aid		BIGINT NOT NULL,			-- address ID of the User who will report on the outcome
	end_time			TIMESTAMPTZ NOT NULL,			-- when the Market expires
	max_ticks			BIGINT NOT NULL,			-- maximum price range (number of intervals)
	create_timestamp	TIMESTAMPTZ NOT NULL,
	-- Status lookup codes  0=>Traded,1=>Reporting,3=>Reported,4=>Disputing,5=>Finalized,6=>Finalized as invalid
	status				SMALLINT DEFAULT 0,
	market_type			SMALLINT NOT NULL,			-- Market type enum: 0:YES_NO | 1:CATEGORICAL | 2:SCALAR
	open_interest		DECIMAL(24,18) DEFAULT 0.0,		-- amount of shares created
	fee					DECIMAL(24,18) NOT NULL,		-- fee to be paid to Market creator as percentage of transaction
	prices				TEXT NOT NULL,				-- range of prices the Market can take
	extra_info			TEXT NOT NULL,				-- specific market metadata (JSON format)
	outcomes			TEXT NOT NULL,				-- possible outcomes of the market
	winning_payouts		TEXT DEFAULT '',
	fin_timestamp		TIMESTAMPTZ DEFAULT TO_TIMESTAMP(0),
	no_show_bond		TEXT NOT NULL,				-- $ penalty to the Creator for failing to emit report
	cur_volume			DECIMAL(24,18) DEFAULT 0.0	-- this is the total volume (for all outcomes althogether)
);
-- Balances of Share tokens per Market (accumulated data, one record per account)
CREATE TABLE sbalances (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			 -- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	account_aid			BIGINT NOT NULL,			-- address id of the User(holder of the shares)
	market_aid			BIGINT NOT NULL,			-- market id of the Market these shares blong
	outcome_idx			SMALLINT NOT NULL,				-- market outcome (index)
	balance				DECIMAL(24,18) NOT NULL		-- balance of shares (bigint as string)
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
	price				DECIMAL(24,18) NOT NULL,
	amount				DECIMAL(24,18) NOT NULL,
	outcome				SMALLINT NOT NULL,
	token_refund		DECIMAL(24,18) NOT NULL,
	shares_refund		DECIMAL(24,18) NOT NULL,
	fees				DECIMAL(24,18) NOT NULL,
	amount_filled		DECIMAL(24,18) NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	shares_escrowed		TEXT NOT NULL,
	tokens_escrowed		TEXT NOT NULL,
	trade_group			TEXT NOT NULL,			-- User defined group label to identify multiple trades
	order_id			TEXT NOT NULL
);
CREATE TABLE oorders (	-- open orders table mirrors `mktord` table, it's used only for active orders
	-- this table is currently disabled until 0x Mesh trading is integrated
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGSERIAL NOT NULL,
	otype				SMALLINT NOT NULL,			-- enum:  0 => BID, 1 => ASK
	outcome_idx			SMALLINT NOT NULL,
	wallet_aid			BIGINT NOT NULL,			-- address of the Wallet Contract of the EOA
	eoa_aid				BIGINT NOT NULL,			-- address of EOA (Externally Owned Account, the real User)
	price				BIGINT NOT NULL,
	amount				BIGINT NOT NULL,
	srv_timestamp		TIMESTAMPTZ NOT NULL,		-- Postgres Server timestamp (not blockchain timestamp)
	expiration			TIMESTAMPTZ NOT NULL,
	order_id			TEXT NOT NULL UNIQUE
);
CREATE TABLE oostats (	-- open order statistics per User
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL,
	eoa_aid				BIGINT NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	num_bids			INT DEFAULT 0,				-- number of total BID orders for this EOA
	num_asks			INT DEFAULT 0,				-- number of total ASK orders for this EOA
	num_cancel			INT DEFAULT 0				-- number of cancelled orders
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
	amount_staked		DECIMAL(24,18) NOT NULL,
	pnumerators			TEXT NOT NULL,		-- payout numerators
	description			TEXT DEFAULT '',
	current_stake		DECIMAL(24,18) DEFAULT 0.0,
	stake_remaining		DECIMAL(24,18) DEFAULT 0.0,
	next_win_start		TIMESTAMPTZ DEFAULT TO_TIMESTAMP(0),
	next_win_end		TIMESTAMPTZ DEFAULT TO_TIMESTAMP(0),
	rpt_timestamp		TIMESTAMPTZ NOT NULL
);
-- Volume
CREATE TABLE volume (	-- this is the VolumeChanged event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	volume				DECIMAL(24,18) NOT NULL,
	outcome_vols		TEXT NOT NULL,		-- this his not numeric because it is not queried (archive only)
	ins_timestamp		TIMESTAMPTZ NOT NULL
);
CREATE TABLE outcome_vol (	-- this is the (accumulated) volume per outcome (indexed) upd. on VolumeChanged
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL REFERENCES market(market_aid) ON DELETE CASCADE,
	outcome_idx			SMALLINT NOT NULL,
	volume				DECIMAL(24,18) NOT NULL,
	last_price			DECIMAL(24,18) DEFAULT 0.0
);
CREATE table oi_chg ( -- open interest changed event
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL REFERENCES market(market_aid) ON DELETE CASCADE,
	ts_inserted			TIMESTAMPTZ NOT NULL, -- timestamp
	oi					DECIMAL(24,18) NOT NULL
);
CREATE TABLE mkt_fin (
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL REFERENCES market(market_aid) ON DELETE CASCADE,
	fin_timestamp		TIMESTAMPTZ NOT NULL,
	winning_payouts		TEXT NOT NULL
);
CREATE TABLE last_block (
	block_num			BIGINT	NOT NULL	-- last block processed by the ETL
);
-- Statistics, automatically accumulated for the main page
CREATE TABLE main_stats (
	id					BIGSERIAL PRIMARY KEY,
	universe_id			BIGINT NOT NULL UNIQUE,
	markets_count		BIGINT DEFAULT 0,	-- counter of all the markets for this Universe
	yesno_count			BIGINT DEFAULT 0,	-- counter for Yes/No markets
	categ_count			BIGINT DEFAULT 0,	-- counter for Categorical markets
	scalar_count		BIGINT DEFAULT 0,	-- counter for Scalar markets
	active_count		BIGINT DEFAULT 0,	-- counter for not-finalized markets
	money_at_stake		DECIMAL(24,18) DEFAULT 0.0,		-- amount in ETH
	trades_count		BIGINT DEFAULT 0	-- total amount of trades
);
