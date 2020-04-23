-- Universe: The container contract for Augur Service
CREATE TABLE universe (
	universe_id			BIGSERIAL PRIMARY KEY,
	universe_addr		TEXT NOT NULL				-- Ethereum address of the Universe contract
);
CREATE TABLE address (
	address_id			BIGSERIAL	PRIMARY KEY,
	addr				TEXT NOT NULL UNIQUE		-- 20 byte Ethereum address , stored as 42 hex string (0x+addr)
);
-- Market
CREATE TABLE market (
	id					BIGSERIAL	PRIMARY KEY,	-- unique id (postgres internal)
	universe_id			BIGSERIAL NOT NULL,			-- reference to universe table
	market_aid			BIGINT NOT NULL UNIQUE,		-- address ID of the Market
	creator_aid			BIGINT NOT NULL,			-- address ID of the User who created the Market
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
	fin_timestmap		BIGINT DEFAULT 0,
	no_show_bond		TEXT NOT NULL,				-- $ penalty to the Creator for failing to emit report
	cur_volume			TEXT DEFAULT ''
);
-- Balances of Share tokens per Market (accumulated data, one record per account)
CREATE TABLE sbalances (
	sb_id				BIGSERIAL PRIMARY KEY,
	address_id			BIGINT NOT NULL,			-- address id of the User(holder of the shares)
	market_aid			BIGINT NOT NULL,			-- market id of the Market these shares blong
	outcome				TEXT NOT NULL,				-- market outcome
	balance				TEXT NOT NULL				-- balance of shares (bigint as string)
);
-- Market Order (BUY/SELL request made by the User via GUI)
CREATE TABLE mktord (
	mktord_id			BIGSERIAL PRIMARY KEY,
	market_aid			BIGSERIAL NOT NULL,
	evt_type			SMALLINT NOT NULL,			-- enum:  0 => Create, 1 => Cancel, 2 => Fill
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
-- Initial Report, submitted by Market Creator
CREATE TABLE ireport (
	ir_id				BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL,
	reporter_aid		BIGINT NOT NULL,
	ini_reporter_aid	BIGINT NOT NULL,
	is_designated		BOOLEAN NOT NULL,
	amount_staked		TEXT NOT NULL,
	pnumerators			TEXT NOT NULL,		-- payout numerators
	description			TEXT DEFAULT '',
	next_win_start		BIGINT,
	next_win_end		BIGINT,
	rpt_timestamp		BIGINT
);
-- Volume
CREATE TABLE volume (
	vol_id				BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL,
	volume				TEXT NOT NULL,
	outcome_vols		TEXT NOT NULL,
	ins_timestamp		BIGINT NOT NULL
)
