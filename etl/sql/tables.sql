-- Layer1 tables
-- Blockchain tables
CREATE TABLE block (
	block_num			BIGINT NOT NULL UNIQUE,
	num_tx				BIGINT DEFAULT 0,
	ts					TIMESTAMPTZ NOT NULL,
	cash_flow			DECIMAL(64,18) DEFAULT 0.0,
	block_hash			CHAR(66) NOT NULL PRIMARY KEY,
	parent_hash			CHAR(66) NOT NULL
);
CREATE TABLE address (
	address_id			BIGSERIAL	PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- block number at which this address was created
	tx_id				BIGINT NOT NULL,			-- transaction at which this address was created
	addr				TEXT NOT NULL UNIQUE		-- 20 byte Ethereum address , stored as 42 hex string (0x+addr)
);
CREATE TABLE transaction (	-- we're only storing transactions related to Augur platform
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	from_aid			BIGINT DEFAULT 0,
	to_aid				BIGINT DEFAULT 0,
	gas_used			BIGINT DEFAULT 0,
	tx_index			INT DEFAULT 0,
	ctrct_create		BOOLEAN DEFAULT FALSE,	-- true if To = nil
	value				DECIMAL(64,18) DEFAULT 0.0,
	gas_price			DECIMAL(64,18) DEFAULT 0.0,
	tx_hash				CHAR(66) NOT NULL UNIQUE,
	input_sig			CHAR(10)	-- input signature (first 4 bytes of Transaction::Data(), hex encoded)
);
CREATE TABLE tx_input ( -- holds transaction input but only for those transactions that we store
	-- since 'data' is holding binary data, this table will grow big if we don't store specific records
	-- records in this table are created by Layer2 functions (specific purpose processing functions)
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	data				TEXT DEFAULT '' -- hex-encoded 0x prefixed core/types.go::Transaction::Data()
);
CREATE TABLE chain_reorg ( -- stores chain reorg events, used by Layer2 to rebuild data on modified blocks
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,
	hash				CHAR(66) NOT NULL
);
CREATE TABLE evt_log (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	contract_aid		BIGINT NOT NULL, -- copied for easy data management
	topic0_sig			CHAR(8) NOT NULL,-- 4 bytes (8 hex chars)  from Topics[0] (the event signature)
	log_rlp				bytea NOT NULL -- RLP encoded (core/types.log:RLPEncode()) event log data
--	data				TEXT NOT NULL
);
CREATE TABLE evt_topic (	-- stores indexed topics of Ethereum Event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,
	tx_id				BIGINT NOT NULL,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	contract_aid		BIGINT NOT NULL,
	pos					SMALLINT NOT NULL,
	value				CHAR(64)	-- hex encoded value of the topic
);
CREATE TABLE abi_funcs (-- Ethereum function signature of the contract method
	id					BIGSERIAL PRIMARY KEY,
	signature			CHAR(8) UNIQUE,
	func_name			TEXT NOT NULL,
	contracts			TEXT NOT NULL
);
CREATE TABLE abi_events (-- Ethereum event signature of the contract event
	id					BIGSERIAL PRIMARY KEY,
	signature			CHAR(8) UNIQUE,
	evt_name			TEXT NOT NULL,
	contracts			TEXT NOT NULL
);
-- Layer2 tables
CREATE TABLE mesh_evt ( -- Events received from 0x Mesh network. source: github.com/0xProject/0x-mesh/zeroex
	id						BIGSERIAL PRIMARY KEY,
	aid						BIGINT DEFAULT 0,	-- can be 0 if address isn't registered yet
-- Event fields:
	time_stamp				TIMESTAMPTZ NOT NULL,
	fillable_amount			DECIMAL(32,18) NOT NULL,
	evt_code				SMALLINT NOT NULL,
-- Augur fields:
	mktord_id				BIGINT DEFAULT NULL, -- the DELETE trigger in 'mktord' does the deletion
	market_aid				BIGINT NOT NULL,
	outcome_idx				SMALLINT NOT NULL,
	otype					SMALLINT NOT NULL,-- 0: BID, 1: ASK
	price					DECIMAL(32,18) NOT NULL,
-- Fill fields:
	amount_fill				DECIMAL(32,18) DEFAULT 0.0,
-- `Order` struct follows:
	order_hash				CHAR(66) NOT NULL,
	chain_id				INT NOT NULL,
	exchange_addr			CHAR(42) NOT NULL,
	maker_addr				CHAR(42) NOT NULL,
	maker_asset_data		TEXT NOT NULL,	-- hex encoded
	maker_fee_asset_data	TEXT NOT NULL,	-- hex encoded
	maker_asset_amount		DECIMAL(32,18) NOT NULL,
	maker_fee				DECIMAL(32,18) NOT NULL,
	taker_address			CHAR(42) NOT NULL,
	taker_asset_data		TEXT NOT NULL,
	taker_fee_asset_data	TEXT NOT NULL,
	taker_asset_amount		DECIMAL(32,18) NOT NULL,
	taker_fee				DECIMAL(32,18) NOT NULL,
	sender_address			CHAR(42) NOT NULL,
	fee_recipient_address	CHAR(42) NOT NULL,
	expiration_time			TIMESTAMPTZ(3) NOT NULL,
	salt					TEXT NOT NULL, -- big.Int as string
	signature				TEXT
);
CREATE TABLE mesh_status (
	last_id_processed	BIGINT DEFAULT 0
);
CREATE TABLE depth_state ( -- the state market depth at any given point in time, used to calculate 
	id					BIGSERIAL PRIMARY KEY,
	meshevt_id			BIGINT NOT NULL REFERENCES mesh_evt(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	otype				SMALLINT NOT NULL,
	order_hash			CHAR(66),
	price				DECIMAL(32,18) NOT NULL,
	amount				DECIMAL(32,18) NOT NULL,
	ini_ts				TIMESTAMPTZ NOT NULL,
	fin_ts				TIMESTAMPTZ NOT NULL
);
CREATE TABLE mesh_link ( -- links two mesh events (ex. one event cancel order created in another event)
	id					BIGSERIAL PRIMARY KEY,
	depthst_id			BIGINT NOT NULL REFERENCES depth_state(id) ON DELETE CASCADE,
	meshevt_id			BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	order_hash			CHAR(66) NOT NULL
);
CREATE TABLE price_estimate (
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL,
	meshevt_id			BIGINT NOT NULL REFERENCES mesh_evt(id) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	bid_state_id		BIGINT, -- will be NULL if there is are no orders (fake orders used)
	ask_state_id		BIGINT,	-- will be NULL if there is are no orders (fake orders used)
	outcome_idx			SMALLINT NOT NULL,
	spread				DECIMAL(32,18) NOT NULL,
	price_est			DECIMAL(32,18) NOT NULL,
	wprice_est			DECIMAL(32,18),-- weighted price estimate (taking volume into consideration)
	wmax_bid			DECIMAL(32,18),
	wmin_ask			DECIMAL(32,18),
	max_bid				DECIMAL(32,18) NOT NULL,
	min_ask				DECIMAL(32,18) NOT NULL,
	wbid_size			DECIMAL(64,18),
	wask_size			DECIMAL(64,18)
);
CREATE TABLE last_block (	-- the value in this table is guaranteeing integrity in the data up to last block
	block_num			BIGINT	NOT NULL	-- last block processed by the ETL
);
CREATE TABLE contract_addresses ( -- Addresses of contracts that compose Augur Platform
	-- format for contract address comment -> [key]:[description]
	-- the Key is used to Augur.sol::lookup() function
	upload_block		BIGINT DEFAULT 0,
	chain_id			BIGINT DEFAULT 0,
	augur				TEXT DEFAULT '',-- Augur: Augur Main contract
	augur_trading		TEXT DEFAULT '',-- AugurTrading: Augur Trading contract
	profit_loss			TEXT DEFAULT '',-- ProfitLoss: Profit Loss contract
	dai_cash			TEXT DEFAULT '',-- Cash: Cash/CashFaucet (local testnet)
	zerox_trade			TEXT DEFAULT '',-- ZeroXTrade: ZeroX Trade
	zerox_xchg			TEXT DEFAULT '',-- Exchange: 0x Exchange
	rep_token			TEXT DEFAULT '',-- REPv2: Reuptation token
	wallet_reg			TEXT DEFAULT '',-- AugurWalletRegistry: Wallet registry v1
	wallet_reg2			TEXT DEFAULT '',-- AugurWalletRegistryV2: Wallet registry v2
	fill_order			TEXT DEFAULT '',-- FillOrder: FillOrder.sol contract
	eth_xchg			TEXT DEFAULT '',-- EthExchange: Uniswap v2 contract
	share_token			TEXT DEFAULT '',-- ShareToken: ShareToken.sol contract
	universe			TEXT DEFAULT '',-- Universe: This holds the Genesis Universe contract
	create_order		TEXT DEFAULT '',-- CreateOrder:
	leg_rep_token		TEXT DEFAULT '',-- LegacyReputationToken:
	buy_part_tok		TEXT DEFAULT '',-- BuyParticipationTokens:
	redeem_stake		TEXT DEFAULT '',-- RedeemStake:
	warp_sync			TEXT DEFAULT '',-- WarpSync:
	hot_loading			TEXT DEFAULT '',-- HotLoading:
	affiliates			TEXT DEFAULT '',-- Affiliates:
	affiliate_val		TEXT DEFAULT '',-- AffiliateValidator:
	ctime				TEXT DEFAULT '',-- Time:
	cancel_order		TEXT DEFAULT '',-- CancelOrder:
	orders				TEXT DEFAULT '',-- Orders:
	sim_trade			TEXT DEFAULT '',-- SiimulateTrade:
	trade				TEXT DEFAULT '',-- Trade:
	oi_cash				TEXT DEFAULT '',-- OICash:
	uniswap_v2_fact		TEXT DEFAULT '',-- UniswapV2Factory:
	uniswap_v2_r2		TEXT DEFAULT '',-- UniswapV2Router02:
	audit_funds			TEXT DEFAULT '',-- AuditFunds:
	weth9				TEXT DEFAULT '',-- WETH9:
	usdc				TEXT DEFAULT '',-- USDC:
	usdt				TEXT DEFAULT '',-- USDT:
	relay_hub_v2		TEXT DEFAULT '',-- RelayHubV2:
	account_loader		TEXT DEFAULT '' -- AccountLoader
);
CREATE TABLE ooconfig ( -- configuration for spread calculation
	spread_threshold	DECIMAL(64,18) DEFAULT 110.0,	-- Reasonable spread to calculate Price Estimate
	osize_threshold		DECIMAL(64,18) DEFAULT 0.0		-- Order size to calculate Price Estimate
);
