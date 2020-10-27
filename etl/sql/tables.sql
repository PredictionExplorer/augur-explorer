-- Blockchain tables
CREATE TABLE block (
	block_num			BIGINT NOT NULL UNIQUE,
	num_tx				BIGINT DEFAULT 0,
	ts					TIMESTAMPTZ NOT NULL,
	cash_flow			DECIMAL(64,18) DEFAULT 0.0,
	block_hash			TEXT NOT NULL PRIMARY KEY,
	parent_hash			TEXT NOT NULL
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
	tx_hash				TEXT NOT NULL UNIQUE
);
-- Universe: The container contract for Augur Service
CREATE TABLE universe (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	universe_id			BIGINT NOT NULL,
	parent_id			BIGINT DEFAULT 0,
	creation_ts			TIMESTAMPTZ DEFAULT TO_TIMESTAMP(0),
	universe_addr		TEXT NOT NULL UNIQUE,		-- Ethereum address of the Universe contract
	payout_numerators	TEXT DEFAULT ''
);
CREATE TABLE address (
	address_id			BIGSERIAL	PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- block number at which this address was created
	tx_id				BIGINT NOT NULL,			-- transaction at which this address was created
	addr				TEXT NOT NULL UNIQUE		-- 20 byte Ethereum address , stored as 42 hex string (0x+addr)
);
-- Market category
CREATE TABLE category (
	cat_id				BIGSERIAL	PRIMARY KEY,
	total_markets		BIGINT DEFAULT 0,
	category			TEXT NOT NULL UNIQUE		-- includes parent category too (comma separated list)
);
-- Market
CREATE TABLE market (
	market_aid			BIGINT NOT NULL PRIMARY KEY,-- address ID of the Market
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	cat_id				BIGINT NOT NULL,			-- category id
	universe_id			BIGSERIAL NOT NULL,			-- reference to universe table
	wallet_aid			BIGINT NOT NULL,			-- address ID of the contract wallet of the User
	eoa_aid				BIGINT NOT NULL,			-- address ID of the User (EOA) who created the market
	reporter_aid		BIGINT NOT NULL,			-- address ID of the User who will report on the outcome
	end_time			TIMESTAMPTZ NOT NULL,		-- when the Market expires
	num_ticks			BIGINT NOT NULL,			-- maximum price range (number of intervals)
	create_timestamp	TIMESTAMPTZ NOT NULL,
	total_trades		BIGINT DEFAULT 0,			-- current number of trades that took place
	total_oorders		BIGINT DEFAULT 0,			-- number of open orders for this market
	winning_outcome		SMALLINT DEFAULT -1,		-- outcome decided by MarketFinalized event
	designated_outcome	SMALLINT DEFAULT -1,		-- outcome submitted by Designated Reported
	initial_outcome		SMALLINT DEFAULT -1,		-- first report that was submitted
	-- Status lookup codes  0=>Traded,1=>Reporting,3=>Reported,4=>Disputing,5=>Finalized,6=>Finalized as invalid
	status				SMALLINT DEFAULT 0,
	market_type			SMALLINT NOT NULL,			-- enum: 0:YES_NO | 1:CATEGORICAL | 2:SCALAR
	money_at_stake		DECIMAL(64,18) DEFAULT 0.0,	-- accumulated money bet on outcomes
	open_interest		DECIMAL(64,18) DEFAULT 0.0,	-- amount of shares created
	fee					DECIMAL(64,18) NOT NULL,	-- fee to be paid to Market creator as percentage of transaction
	prices				TEXT NOT NULL,				-- range of prices the Market can take
	extra_info			TEXT NOT NULL,				-- specific market metadata (JSON format)
	outcomes			TEXT NOT NULL,				-- possible outcomes of the market
	winning_payouts		TEXT DEFAULT '',
	fin_timestamp		TIMESTAMPTZ DEFAULT TO_TIMESTAMP(0),
	no_show_bond		DECIMAL(64,18),				-- $ penalty to the Creator for failing to emit report
	validity_bond		DECIMAL DEFAULT 0.0,		-- fee returned to creator if market isnt invalid
	cur_volume			DECIMAL(64,18) DEFAULT 0.0	-- this is the total volume (for all outcomes althogether)
);
-- Market Order (BUY/SELL request made by the User via GUI)
CREATE TABLE mktord (-- in this table only 'Fill' type orders are stored (Create/Cancel are temporary)
	id					BIGSERIAL PRIMARY KEY,
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	eoa_aid				BIGINT NOT NULL,			-- Address of the user who created the order (Creator)
	wallet_aid			BIGINT NOT NULL,			-- address of the creator
	eoa_fill_aid		BIGINT NOT NULL,			-- address of the filler; source: AugurTrading.sol:24
	wallet_fill_aid		BIGINT NOT NULL,			-- address of the Wallet address of the filler
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	time_stamp			TIMESTAMPTZ NOT NULL,
	oaction				SMALLINT NOT NULL,			-- order action:  0=>Create, 1=>Cancel, 2=>Fill
													-- Create: User posts a BID or ASK execpting to be filed
													-- Fill: User buys or sells existing (Created) order
													-- Cancel: User removes active order (BID/ASK)
	otype				SMALLINT NOT NULL,			-- enum:  0 => BID, 1 => ASK
	outcome_idx			SMALLINT NOT NULL,
	price				DECIMAL(24,18) NOT NULL,
	amount				DECIMAL(24,18) NOT NULL,
	token_refund		DECIMAL(24,18) NOT NULL,
	shares_refund		DECIMAL(24,18) NOT NULL,
	fees				DECIMAL(24,18) NOT NULL,
	amount_filled		DECIMAL(24,18) NOT NULL,
	shares_escrowed		TEXT NOT NULL,
	tokens_escrowed		TEXT NOT NULL,
	trade_group			TEXT NOT NULL,			-- User defined group label to identify multiple trades
	order_hash			TEXT NOT NULL
);
CREATE TABLE mesh_evt ( -- Events received from 0x Mesh network. source: github.com/0xProject/0x-mesh/zeroex
	id						BIGSERIAL PRIMARY KEY,
	eoa_aid					BIGINT DEFAULT 0,	-- can be 0 if address isn't registered yet
	wallet_aid				BIGINT DEFAULT 0,	-- can be 0 if address isn't registered yet
-- Event fields:
	time_stamp				TIMESTAMPTZ NOT NULL,
	fillable_amount			DECIMAL(32,18) NOT NULL,
	evt_code				SMALLINT NOT NULL,
-- Augur fields:
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
	expiration_time			TIMESTAMPTZ NOT NULL,
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
	min_ask				DECIMAL(32,18) NOT NULL
);
CREATE TABLE oorders (	-- contains open orders made on 0x Mesh network, later they are converted into 'mktord` records
	id					BIGSERIAL PRIMARY KEY,
	otype				SMALLINT NOT NULL,			-- enum:  0 => BID, 1 => ASK
	outcome_idx			SMALLINT NOT NULL,
	opcode				SMALLINT NOT NULL,			-- operation; 0: CREATED, 1: AUTOEXPIRED, 2: USER-CANCELLED, 3: FILLED DB SYNC
	market_aid			BIGINT NOT NULL,
	wallet_aid			BIGINT NOT NULL,			-- address of the Wallet Contract of the EOA
	eoa_aid				BIGINT NOT NULL,			-- address of EOA (Externally Owned Account, the real User)
	price				DECIMAL(32,18) NOT NULL,
	initial_amount		DECIMAL(32,18) NOT NULL,	-- when partially filled, this keeps the original amount
	amount				DECIMAL(32,18) NOT NULL,
	evt_timestamp		TIMESTAMPTZ NOT NULL,		-- 0x Mesh event timestamp
	srv_timestamp		TIMESTAMPTZ NOT NULL,		-- Postgres Server timestamp (not blockchain timestamp)
	expiration			TIMESTAMPTZ NOT NULL,
	order_hash			CHAR(66) NULL UNIQUE
);
CREATE TABLE oohist ( -- open order history
	id					BIGSERIAL PRIMARY KEY,
	mktord_id			BIGINT DEFAULT NULL REFERENCES mktord(id) ON DELETE CASCADE, -- used only for Fill events
	otype				SMALLINT NOT NULL,			-- enum:  0 => BID, 1 => ASK
	outcome_idx			SMALLINT NOT NULL,
	opcode				SMALLINT NOT NULL,			-- operation; 0: CREATED, 1: AUTOEXPIRED, 2: USER-CANCELLED
	market_aid			BIGINT NOT NULL,
	eoa_aid				BIGINT NOT NULL,			-- address of EOA (Externally Owned Account, the real User)
	wallet_aid			BIGINT NOT NULL,			-- address of the Wallet Contract of the EOA
	price_estimate		DECIMAL(32,18) DEFAULT 0.0,
	price				DECIMAL(32,18) NOT NULL,
	initial_amount		DECIMAL(32,18) NOT NULL,	-- initial amount order was created
	amount			DECIMAL(32,18) NOT NULL,		-- amount remaining to be filled
	evt_timestamp		TIMESTAMPTZ DEFAULT to_timestamp(0),-- 0x Mesh event timestamp
	srv_timestamp		TIMESTAMPTZ DEFAULT to_timestamp(0),-- Postgres Server timestamp (not blockchain timestamp)
	expiration			TIMESTAMPTZ DEFAULT to_timestamp(0),
	order_hash			CHAR(66)					-- Order Hash (github.com/0x-mesh/zeroex/order.go:Order.hash)
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
CREATE TABLE ooconfig ( -- configuration for spread calculation
	spread_threshold	DECIMAL(64,18) DEFAULT 110.0,	-- Reasonable spread to calculate Price Estimate
	osize_threshold		DECIMAL(64,18) DEFAULT 0.0		-- Order size to calculate Price Estimate
);
-- Report, submitted by Market Creator
CREATE TABLE report (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	eoa_aid				BIGINT NOT NULL,			-- User's address (EOA) of the Reporter
	wallet_aid			BIGINT NOT NULL,			-- Wallet's contract address of the Reporter
	ini_reporter_aid	BIGINT DEFAULT 0,
	disputed_aid		BIGINT DEFAULT 0,
	dispute_round		BIGINT DEFAULT 1,
	outcome_idx			SMALLINT NOT NULL,
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
	total_trades		BIGINT DEFAULT 0,
	volume				DECIMAL(24,18) NOT NULL,
	outcome_vols		TEXT NOT NULL,		-- this his not numeric because it is not queried (archive only)
	ins_timestamp		TIMESTAMPTZ NOT NULL
);
CREATE TABLE outcome_vol (	-- this is the (accumulated) volume per outcome (indexed) upd. on VolumeChanged
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL REFERENCES market(market_aid) ON DELETE CASCADE,
	total_trades		BIGINT DEFAULT 0,
	total_oorders		BIGINT DEFAULT 0,
	outcome_idx			SMALLINT NOT NULL,
	volume				DECIMAL(24,18) DEFAULT 0.0,
	last_price			DECIMAL(24,18) DEFAULT 0.0,
	highest_bid			DECIMAL(64,18) DEFAULT 0.0,	-- highest BID price , updated from open orders
	lowest_ask			DECIMAL(64,18) DEFAULT 0.0,	-- lowest ASK price, updated from open orders
	cur_spread			DECIMAL(64,18) DEFAULT 0.0,	-- spread from open orders (lowest_ask - highest bid)
	price_estimate		DECIMAL(64,18) DEFAULT 0.0  -- calculated using trigger update_price_estimate()
);
CREATE table oi_chg ( -- open interest changed event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	ts_inserted			TIMESTAMPTZ NOT NULL, -- timestamp
	oi					DECIMAL(64,18) NOT NULL
);
CREATE TABLE mkt_fin (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	winning_outcome		SMALLINT DEFAULT 0,
	fin_timestamp		TIMESTAMPTZ NOT NULL,
	winning_payouts		TEXT NOT NULL
);
CREATE TABLE last_block (	-- the value in this table is guaranteeing integrity in the data up to last block
	block_num			BIGINT	NOT NULL	-- last block processed by the ETL
);
CREATE table dai_transf (	-- transfers of DAI tokens (deposits/withdrawals of funds)
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	from_aid			BIGINT DEFAULT 0,
	to_aid				BIGINT DEFAULT 0,
	from_internal		BOOLEAN DEFAULT false,
	to_internal			BOOLEAN DEFAULT false,
	amount				DECIMAL(64,18) DEFAULT 0.0
);
CREATE table dai_bal (	-- DAI token balance
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	dai_transf_id		BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	processed			BOOLEAN DEFAULT false,	-- true if balances have been calculated
	augur				BOOLEAN DEFAULT false,	-- true if the user has account on Augur Platform
	internal			BOOLEAN DEFAULT false,	-- true if it is an exchange between Agur's contracts
	balance				DECIMAL(64,18) DEFAULT 0.0,
	amount				DECIMAL(64,18) DEFAULT 0.0
);
CREATE table rep_transf (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	from_aid			BIGINT DEFAULT 0,
	to_aid				BIGINT DEFAULT 0,
	amount				DECIMAL(32,18) DEFAULT 0.0
);
CREATE table tok_transf (	-- Tokens Transferred event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	from_aid			BIGINT NOT NULL,
	to_aid				BIGINT NOT NULL,
	token_type			SMALLINT DEFAULT 0,
	value				DECIMAL(64,32) DEFAULT 0.0
);
CREATE table tbc (			-- Token Balance Changed event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	owner_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	token_type			SMALLINT DEFAULT 0,
	outcome				SMALLINT NOT NULL,
	balance				DECIMAL(64,32) DEFAULT 0.0
);
CREATE table stbc (			-- Share Token Balance Changed event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	account_aid			BIGINT NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	balance				DECIMAL(64,32) DEFAULT 0.0
);
-- Balances of Share tokens per Market (accumulated data, one record per account)
CREATE TABLE sbalances (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			 -- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	account_aid			BIGINT NOT NULL,			-- address id of the User(holder of the shares)
	market_aid			BIGINT NOT NULL,			-- market id of the Market these shares blong
	num_transfers		BIGINT DEFAULT 0,			-- counter for tracking now many transfers we had
	outcome_idx			SMALLINT NOT NULL,				-- market outcome (index)
	balance				DECIMAL(24,18) NOT NULL		-- balance of shares (bigint as string)
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
	money_at_stake		DECIMAL(64,18) DEFAULT 0.0,		-- amount in ETH
	trades_count		BIGINT DEFAULT 0	-- total amount of trades
);
CREATE TABLE trd_mkt_stats (	-- trade statistics per User and per Market
	id					BIGSERIAL PRIMARY KEY,
	eoa_aid				BIGINT NOT NULL,
	wallet_aid			BIGINT NOT NULL,
	market_aid			BIGINT NOT NULL,
	total_trades		BIGINT DEFAULT 0,
	total_reports		BIGINT DEFAULT 0,
	volume_traded		DECIMAL(64,18) DEFAULT 0.0,
	profit_loss			DECIMAL(32,18) DEFAULT 0.0,
	report_profits		DECIMAL(32,18) DEFAULT 0.0,
	aff_profits			DECIMAL(32,18) DEFAULT 0.0,
	frozen_funds		DECIMAL(32,18) DEFAULT 0.0
);
CREATE TABLE mkts_traded (	-- just a simple link to calculate how many markets a User has traded
	eoa_aid				BIGINT NOT NULL,
	market_aid			BIGINT NOT NULL
);
CREATE TABLE ustats (	-- statistics per User account
	-- Notte: not only this table is for statistics, but it keeps important link between EOA and Wallet contract
	eoa_aid				BIGINT PRIMARY KEY,		-- Externally Owned ACcount (EOA) address for this user
	wallet_aid			BIGINT NOT NULL,	-- Wallet Contract address id
	total_trades		BIGINT DEFAULT 0,
	markets_created		BIGINT DEFAULT 0,
	markets_traded		BIGINT DEFAULT 0,
	withdraw_reqs		BIGINT DEFAULT 0,
	deposit_reqs		BIGINT DEFAULT 0,
	total_reports		BIGINT DEFAULT 0,
	total_designated	BIGINT DEFAULT 0,			-- total reports submitted as designated reporter
	volume_traded		DECIMAL(64,18) DEFAULT 0.0,
	profit_loss			DECIMAL(32,18) DEFAULT 0.0,
	report_profits		DECIMAL(32,18) DEFAULT 0.0,
	aff_profits			DECIMAL(32,18) DEFAULT 0.0,	-- affiliate commissions earned
	money_at_stake		DECIMAL(32,18) DEFAULT 0.0, -- how much has this User bet on Augur mkts
	total_withdrawn		DECIMAL(32,18) DEFAULT 0.0,
	total_deposited		DECIMAL(32,18) DEFAULT 0.0,
	validity_bonds		DECIMAL DEFAULT 0.0,	-- sum of all validity bonds (market creation bond)
	rep_frozen			DECIMAL(32,18) DEFAULT 0.0,	-- amount of REP tokens frozen for all (participated) markets
	-- Gas usage statistics per user:
	-- values contain Gas Used , accumulated
	gtrading			DECIMAL DEFAULT 0,
	greporting			DECIMAL DEFAULT 0,
	gmarkets			DECIMAL DEFAULT 0,
	-- values contain Gas Price , accumulated
	geth_trading		DECIMAL(64,18) DEFAULT 0.0,
	geth_reporting		DECIMAL(64,18) DEFAULT 0.0,
	geth_markets		DECIMAL(64,18) DEFAULT 0.0
);
CREATE TABLE profit_loss ( -- captures ProfitLossChanged event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	eoa_aid				BIGINT NOT NULL,
	wallet_aid			BIGINT NOT NULL,
	mktord_id			BIGINT DEFAULT 0,			-- this is the id of the market order generated this PL
	outcome_idx			SMALLINT NOT NULL,
	closed_position		SMALLINT DEFAULT 0,			-- 0 - open position, 1 - closed position
	-- noote: the following decimal precisions depend on precision of Augur events , inserted in db.go
	net_position		DECIMAL(32,18) DEFAULT 0.0,
	avg_price			DECIMAL(32,20) DEFAULT 0.0,
	frozen_funds		DECIMAL(64,36) DEFAULT 0.0,
	realized_profit		DECIMAL(64,36) DEFAULT 0.0,	-- this is the field copied directly from Augur' Event Log
	realized_cost		DECIMAL(64,36) DEFAULT 0.0,
	immediate_profit	DECIMAL(64,36) DEFAULT 0.0,	-- the profit on position direction change (or position size update)
	immediate_ff		DECIMAL(64,36) DEFAULT 0.0,	-- frozen funds held for current position (not the accumulated)
	time_stamp			TIMESTAMPTZ NOT NULL
);
CREATE TABLE claim_funds (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	eoa_aid				BIGINT NOT NULL,
	market_aid			BIGINT NOT NULL,
	outcome_idx			BIGINT NOT NULL,
	last_pl_id			BIGINT NOT NULL,			-- last id of profit loss that updated this record
	claim_ts			TIMESTAMPTZ DEFAULT to_timestamp(0),		-- timestamp of claim action
	claim_status		SMALLINT DEFAULT 0,			-- 0:nothing to claim;1:unclaimed but existent;2:claimed
	autocalculated		BOOLEAN DEFAULT FALSE,		-- true if PL was automatically calculated (not by PL event)
	final_profit		DECIMAL(64,18) DEFAULT 0.0,
	unfrozen_funds		DECIMAL(64,18) DEFAULT 0.0	-- amount of funds removed from frozen funds
);
CREATE TABLE uranks (   -- User Rankings (how this user ranks against each other, ex: Top 13% in profit made
	eoa_aid             BIGINT PRIMARY KEY,
	total_trades		BIGINT DEFAULT 0,
	top_profit          DECIMAL(5,2) DEFAULT 100.0,    -- position of the user in profits accumulated over lifetime
	top_trades          DECIMAL(5,2) DEFAULT 100.0,    -- position of the user in number of accumulated trades
	top_volume			DECIMAL(5,2) DEFAULT 100.0,	   -- position of the user in accumulated trading volume
	profit				DECIMAL(32,18) DEFAULT 0.0,
	volume				DECIMAL(32,18) DEFAULT 0.0
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
CREATE TABLE register_contract (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	addr				TEXT NOT NULL,
	key					TEXT NOT NULL
);
CREATE TABLE unique_addrs (	-- Unique addresses per day, statistics
	day					DATE PRIMARY KEY,
	num_addrs			BIGINT DEFAULT 0
);
CREATE TABLE gas_spent (-- global gas spent
	day					DATE PRIMARY KEY,
	num_trading			BIGINT DEFAULT 0,		--number of trading transactions for that day
	num_reporting		BIGINT DEFAULT 0,
	num_markets			BIGINT DEFAULT 0,
	num_total			BIGINT DEFAULT 0,
	-- values contain raw Gas used
	trading				DECIMAL DEFAULT 0,
	reporting			DECIMAL DEFAULT 0,
	markets				DECIMAL DEFAULT 0,
	total				DECIMAL DEFAULT 0,
	-- values contain Gas Price , accumulated
	eth_trading			DECIMAL(64,18) DEFAULT 0.0,
	eth_reporting		DECIMAL(64,18) DEFAULT 0.0,
	eth_markets			DECIMAL(64,18) DEFAULT 0.0,
	eth_total			DECIMAL(64,18) DEFAULT 0.0
);
CREATE TABLE exec_wtx (	-- stores contract calls of input with sig=78dc0eed (executeTransactionStatus)
-- source: AugurWalletRegistry.sol:executeWalletTransaction()
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	eoa_aid				BIGINT NOT NULL,
	wallet_aid			BIGINT DEFAULT 0,	-- Uniswap can exchange tokens for EOAs, so Wallet id will be 0
	to_aid				BIGINT NOT NULL,
	referral_aid		BIGINT DEFAULT 0,	-- address of referral account (referral_aid will get commissions on TXs of eoa_aid)
	value				DECIMAL(64,18) DEFAULT 0.0,
	payment				DECIMAL(64,18) DEFAULT 0.0,
	desired_signer_bal	DECIMAL(64,18) DEFAULT 0.0,	-- desiredSignerBalance
	max_xchg_rate_dai	DECIMAL(64,18) DEFAULT 0.0,	-- maxExchangeRateInDai
	input_sig			CHAR(8),			-- hex encoded first 4 bytes of call_data (indexed field)
	fingerprint			TEXT NOT NULL,		-- hex-encoded 32 byte (64char) value of Browser fingerprint
	call_data			TEXT DEFAULT '',	-- hex-encoded input to contract in 'to' field
	revert_on_failure	BOOLEAN DEFAULT FALSE
);
CREATE TABLE agtx_status (-- Augur transaction status (used to track Gas fees for all interactions with Augur
-- this table stores the result of the call registered in `exec_wtx` table
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	eoa_aid				BIGINT NOT NULL,
	wallet_aid			BIGINT NOT NULL,
	success				BOOLEAN NOT NULL,
	funding_success		BOOLEAN NOT NULL
);
CREATE TABLE augur_flag ( -- collection of signs required to consider an account as enabled for Augur trading
	-- when all flags are TRUE , we insert a record into 'ustats' table with eoa_aid=wallet_aid=aid
	aid					BIGINT PRIMARY KEY,
	act_block_num		BIGINT,					-- Block number when activation happened
	ap_0xtrade_on_cash	BOOLEAN DEFAULT FALSE,	-- Approval for ZeroXTrade at Cash (DAI) contract
	ap_fill_on_cash		BOOLEAN DEFAULT FALSE,	-- Approval for FillOrder contract at Cash (DAI) contract
	ap_fill_on_shtok	BOOLEAN DEFAULT FALSE,	-- ApprovalForAll for FillOrder at ShareToken contract
	set_referrer		BOOLEAN DEFAULT FALSE	-- Affiliates::setReferrer() tx input (informative only, not obligatory)
);
CREATE TABLE pl_debug (-- Profit loss data for debugging, scanned after Block has been processed
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	wallet_aid			BIGINT NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	profit_loss			DECIMAL(64,36) DEFAULT 0.0,
	frozen_funds		DECIMAL(64,36) DEFAULT 0.0,
	net_position		DECIMAL(32,18) DEFAULT 0.0,
	avg_price			DECIMAL(32,20) DEFAULT 0.0
);
