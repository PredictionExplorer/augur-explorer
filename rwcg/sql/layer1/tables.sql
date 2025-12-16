
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
	num_logs			INT DEFAULT 0,
	ctrct_create		BOOLEAN DEFAULT FALSE,	-- true if To = nil
	value				DECIMAL(80,18) DEFAULT 0.0,
	gas_price			DECIMAL(80,18) DEFAULT 0.0,
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
	log_index			INT NOT NULL DEFAULT 0, -- log index within the block
	log_rlp				bytea NOT NULL, -- RLP encoded (core/types.log:RLPEncode()) event log data
	UNIQUE(block_num, log_index) -- prevent duplicate event insertions
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
