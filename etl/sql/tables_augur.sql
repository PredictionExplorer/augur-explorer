CREATE TABLE register_contract (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	addr				TEXT NOT NULL,
	key					TEXT NOT NULL
);
-- Universe: The container contract for Augur Service
CREATE TABLE universe (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	universe_id			BIGINT NOT NULL,
	parent_id			BIGINT DEFAULT 0,
	creation_ts			TIMESTAMPTZ DEFAULT TO_TIMESTAMP(0),
	validity_bond		DECIMAL(64,18) DEFAULT 0.0,
	noshow_bond		DECIMAL(64,18) DEFAULT 0.0,
	universe_addr		TEXT NOT NULL UNIQUE,		-- Ethereum address of the Universe contract
	payout_numerators	TEXT DEFAULT ''
);
-- Market category
CREATE TABLE category (
	cat_id				BIGSERIAL	PRIMARY KEY,
	total_markets		BIGINT DEFAULT 0,
	category			TEXT NOT NULL UNIQUE		-- includes parent category too (comma separated list)
);
-- Market
CREATE TABLE market (
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL UNIQUE,-- address ID of the Market
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	cat_id				BIGINT NOT NULL,			-- category id
	universe_id			BIGSERIAL NOT NULL,			-- reference to universe table
	creator_aid			BIGINT NOT NULL,			-- address ID of account creating the market
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
	lo_price			DECIMAL(78,18) DEFAULT 0,
	hi_price			DECIMAL(78,18) DEFAULT 0,
--DISCONTINUED	prices				TEXT NOT NULL,				-- range of prices the Market can take
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
	aid					BIGINT NOT NULL,			-- Address of the user who created the order (Creator)
	fill_aid			BIGINT NOT NULL,			-- address of the filler; source: AugurTrading.sol:24
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
-- Report, submitted by Market Creator
CREATE TABLE report (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,			-- User's address (EOA) of the Reporter
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
	volume				DECIMAL(64,18) NOT NULL,
	outcome_vols		TEXT NOT NULL,		-- this his not numeric because it is not queried (archive only)
	ins_timestamp		TIMESTAMPTZ NOT NULL
);
CREATE TABLE outcome_vol (	-- this is the (accumulated) volume per outcome (indexed) upd. on VolumeChanged
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	total_trades		BIGINT DEFAULT 0,
	total_oorders		BIGINT DEFAULT 0,
	outcome_idx			SMALLINT NOT NULL,
	volume				DECIMAL(64,18) DEFAULT 0.0,
	last_price			DECIMAL(24,18) DEFAULT 0.0,
	highest_bid			DECIMAL(64,18) DEFAULT 0.0,	-- highest BID price , updated from open orders
	lowest_ask			DECIMAL(64,18) DEFAULT 0.0,	-- lowest ASK price, updated from open orders
	cur_spread			DECIMAL(64,18) DEFAULT 0.0,	-- spread from open orders (lowest_ask - highest bid)
	price_estimate		DECIMAL(64,18) DEFAULT 0.0  -- calculated using trigger update_price_estimate()
);
CREATE TABLE cancel_0x ( -- events canceling 0x orders
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	otype				SMALLINT NOT NULL,-- 0: BID, 1: ASK
	price				DECIMAL(32,18) NOT NULL,
	order_hash			CHAR(66) NULL UNIQUE
);
CREATE TABLE tproceeds (	-- table to store TradingProceedsClaimed event (User has claimed his funds)
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	num_shares			DECIMAL(36,18) NOT NULL,
	num_payout_tok		DECIMAL(36,18) NOT NULL,
	fees				DECIMAL(36,18) NOT NULL
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
-- Statistics, automatically accumulated for the main page
CREATE TABLE main_stats (
	id					BIGSERIAL PRIMARY KEY,
	universe_id			BIGINT NOT NULL UNIQUE,
	markets_count		BIGINT DEFAULT 0,	-- counter of all the markets for this Universe
	yesno_count			BIGINT DEFAULT 0,	-- counter for Yes/No markets
	categ_count			BIGINT DEFAULT 0,	-- counter for Categorical markets
	scalar_count		BIGINT DEFAULT 0,	-- counter for Scalar markets
	active_count		BIGINT DEFAULT 0,	-- counter for not-finalized markets
	total_accounts		BIGINT DEFAULT 0,	-- total number of accounts that have Augur-related activity
	money_at_stake		DECIMAL(64,18) DEFAULT 0.0,		-- amount in ETH
	trades_count		BIGINT DEFAULT 0	-- total amount of trades
);
CREATE TABLE trd_mkt_stats (	-- trade statistics per User and per Market
	id					BIGSERIAL PRIMARY KEY,
	aid					BIGINT NOT NULL,
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
	aid					BIGINT NOT NULL,
	market_aid			BIGINT NOT NULL
);
CREATE TABLE ustats (	-- statistics per User account
	-- Notte: not only this table is for statistics, but it keeps important link between EOA and Wallet contract
	aid					BIGINT PRIMARY KEY,		-- Externally Owned ACcount (EOA) address for this user
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
CREATE TABLE eoa_wallet (	-- Table linking EOA and Wallet addresses
	eoa_aid				BIGINT PRIMARY KEY,
	wallet_aid			BIGINT NOT NULL
);
CREATE TABLE profit_loss ( -- captures ProfitLossChanged event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,			-- addressID of Ethereum account that is doing the operation
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
	aid					BIGINT NOT NULL,
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
	aid		            BIGINT PRIMARY KEY,
	total_trades		BIGINT DEFAULT 0,
	top_profit          DECIMAL(5,2) DEFAULT 100.0,    -- position of the user in profits accumulated over lifetime
	top_trades          DECIMAL(5,2) DEFAULT 100.0,    -- position of the user in number of accumulated trades
	top_volume			DECIMAL(5,2) DEFAULT 100.0,	   -- position of the user in accumulated trading volume
	profit				DECIMAL(32,18) DEFAULT 0.0,
	volume				DECIMAL(32,18) DEFAULT 0.0
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
	wallet_aid			BIGINT DEFAULT 0,   -- Uniswap can exchange tokens for EOAs, so Wallet id will be 0
	to_aid				BIGINT NOT NULL,	-- destination contract address to where the transaciton is sent
	referral_aid		BIGINT DEFAULT 0,	-- address of referral account (referral_aid will get commissions on TXs)
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
	aid					BIGINT NOT NULL,
	success				BOOLEAN NOT NULL,
	funding_success		BOOLEAN NOT NULL
);
CREATE TABLE augur_flag ( -- collection of signs required to consider an account as enabled for Augur trading
	-- when all flags are TRUE , we insert a record into 'ustats' table meaning that this is an Augur account
	aid					BIGINT PRIMARY KEY,
	act_block_num		BIGINT,					-- Block number when activation happened (not always gathered, 0 if unknown but active)
	ap_0xtrade_on_cash	BOOLEAN DEFAULT FALSE,	-- Approval for ZeroXTrade at Cash (DAI) contract
	ap_fill_on_cash		BOOLEAN DEFAULT FALSE,	-- Approval for FillOrder contract at Cash (DAI) contract
	ap_fill_on_shtok	BOOLEAN DEFAULT FALSE,	-- ApprovalForAll for FillOrder at ShareToken contract
	set_referrer		BOOLEAN DEFAULT FALSE	-- Affiliates::setReferrer() tx input (informative only, not obligatory)
);
CREATE TABLE oorders (	-- contains open orders made on 0x Mesh network, later they are converted into 'mktord` records
	id					BIGSERIAL PRIMARY KEY,
	otype				SMALLINT NOT NULL,			-- enum:  0 => BID, 1 => ASK
	outcome_idx			SMALLINT NOT NULL,
	opcode				SMALLINT NOT NULL,			-- operation; 0: CREATED, 1: AUTOEXPIRED, 2: USER-CANCELLED, 3: FILLED DB SYNC
	market_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,			-- address of the account that created the order
	price				DECIMAL(32,18) NOT NULL,
	initial_amount		DECIMAL(32,18) NOT NULL,	-- when partially filled, this keeps the original amount
	amount				DECIMAL(32,18) NOT NULL,
	evt_timestamp		TIMESTAMPTZ NOT NULL,		-- 0x Mesh event timestamp
	srv_timestamp		TIMESTAMPTZ NOT NULL,		-- Postgres Server timestamp (not blockchain timestamp)
	expiration			TIMESTAMPTZ NOT NULL,
	order_hash			CHAR(66) NULL UNIQUE
);
CREATE TABLE oostats (	-- open order statistics per User
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	num_bids			INT DEFAULT 0,				-- number of total BID orders for this EOA
	num_asks			INT DEFAULT 0,				-- number of total ASK orders for this EOA
	num_cancel			INT DEFAULT 0				-- number of cancelled orders
);
CREATE TABLE pl_debug (-- Profit loss data for debugging, scanned after Block has been processed
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	profit_loss			DECIMAL(64,36) DEFAULT 0.0,
	frozen_funds		DECIMAL(64,36) DEFAULT 0.0,
	net_position		DECIMAL(32,18) DEFAULT 0.0,
	avg_price			DECIMAL(32,20) DEFAULT 0.0
);
CREATE TABLE augur_proc_status (-- Augur Tradign process status
	last_tx_id			BIGINT DEFAULT 0
);
CREATE table tok_transf (	-- Tokens Transferred event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
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
	outside_augur_ui	BOOLEAN DEFAULT false, -- true if the transfer was not made by Augur UI (fontend)
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
CREATE TABLE mkt_words(-- search tokens for searching markets by description/category
	id					BIGSERIAL PRIMARY KEY,
	market_aid			BIGINT,
	cat_id				BIGINT,
	tok_type			SMALLINT DEFAULT 0,				-- 0-market 1 - category
	tokens				TSVECTOR
);
CREATE TABLE defi_stats( -- statistics of DeFi trading
	aid					BIGINT NOT NULL,
	balancer_swaps		BIGINT DEFAULT 0,
	uniswap_swaps		BIGINT DEFAULT 0
);
CREATE TABLE agtx_evt(	-- augur transaction event (for accumulating Augur-related events)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	ref_id				BIGINT NOT NULL,	-- reference id (event id or market order id)
	block_num			BIGINT NOT NULL,	-- this is just a copy (for easy data management)
	account_aid			BIGINT DEFAULT 0,	-- if account is created, the address of the account
	market_aid			BIGINT DEFAULT 0,	-- if market is related, the market address
	defi_platform		INT	DEFAULT 0,		-- 0: Unknown, 1: Uniswap, 2: Balancer
	--evt_type's: 0:Outside Augur-UI and outside DeFI trading (i.e. Unknown), 1: DeFi trading 2: Market Created,
	-- 3: Trade, 4: Report, 5: Claim profits. 6:Other Augur UI
	evt_type			INT DEFAULT 0,
	PRIMARY KEY(tx_id,ref_id)
);
CREATE TABLE agtx( -- augur transaction (transaction related to Augur trading (transfer of market shares of any kind)
	tx_id				BIGINT PRIMARY KEY REFERENCES transaction(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			 -- this is just a copy (for easy data management)
	num_events			INT DEFAULT 0,		-- number of augur events in this transaction
	num_tx				INT DEFAULT 0,		-- counter for number of Ethereum Transactions
	num_agtx_evts		INT DEFAULT 0,		-- counter for number of Augur events 
	num_defi_evts		INT DEFAULT 0,		-- counter for number of swaps on DeFi platforms
	num_other_evts		INT DEFAULT 0,		-- counter for number of other type of events
	num_bal_swaps		INT DEFAULT 0,		-- counter for how many balancer swaps were made
	num_uni_swaps		INT DEFAULT 0,		-- counter for how many uniswap v2 swaps were made
	gas_used			BIGINT NOT NULL,	-- copy from 'transaction' table
	gas_price			DECIMAL(64,18) NOT NULL	-- copy from 'transaction' table
);
CREATE TABLE agblk( -- augur block (a block which has Augur-related data , data derived from 'agtx' table)
	block_num			BIGINT PRIMARY KEY REFERENCES block(block_num) ON DELETE CASCADE,
	-- the following counters are accumulated from 'agtx' table using triggers
	num_events			INT DEFAULT 0,		-- number of augur events in this block
	num_tx				INT DEFAULT 0,		-- counter for number of Ethereum Transactions
	num_agtx_evts		INT DEFAULT 0,		-- counter for number of transactions (Augur UI related)
	num_defi_evts		INT DEFAULT 0,		-- counter for number of trades on DeFi platforms
	num_other_evts		INT DEFAULT 0,		-- counter for number of other type of transactions (like direct sharetoken transfs)
	num_bal_swaps		INT DEFAULT 0,
	num_uni_swaps		INT DEFAULT 0
);
CREATE TABLE val_bond_chg( -- validity bond changed event (sig: 69af68e3)
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	universe_id			BIGINT NOT NULL,
	bond_value			DECIMAL(64,18)
);
CREATE TABLE noshow_bond_chg( -- validity bond changed event (sig: 69af68e3)
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	universe_id			BIGINT NOT NULL,
	bond_value			DECIMAL(64,18)
);
CREATE table crowdsourcer_created (			--
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	market_aid			BIGINT NOT NULL,
	dispute_aid			BIGINT NOT NULL,
	dispute_round		INT NOT NULL,
	payout_numerators	TEXT DEFAULT '',
	size				DECIMAL(64,18)
);
CREATE table crowdsourcer_completed (	--
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	market_aid			BIGINT NOT NULL,
	crowdsrc_aid		BIGINT NOT NULL,
	next_win_start		TIMESTAMPTZ DEFAULT TO_TIMESTAMP(0),
	next_win_end		TIMESTAMPTZ DEFAULT TO_TIMESTAMP(0),
	dispute_round		INT NOT NULL,
	pacing_on			BOOLEAN NOT NULL,
	payout_numerators	TEXT DEFAULT '',
	tot_rep_payout		DECIMAL(64,18),
	tot_rep_market		DECIMAL(64,18)
);
CREATE table crowdsourcer_redeemed ( -- DisputeCrowdsourcerRedeemed event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	reporter_aid		BIGINT NOT NULL,
	crowdsourcer_aid	BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	amount				DECIMAL(64,18) NOT NULL,
	rep					DECIMAL(64,18) NOT NULL,
	payout_numerators	TEXT DEFAULT ''
);
CREATE table dispute_window (
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	universe_id			BIGINT NOT NULL,
	wid					BIGINT NOT NULL, -- window ID (not the address of the contract)
	window_aid			BIGINT NOT NULL, -- address of Window contract
	start_time			TIMESTAMPTZ NOT NULL,
	end_time			TIMESTAMPTZ NOT NULL,
	initial				BOOLEAN NOT NULL
);
CREATE table rep_stake_chg ( -- DesignatedReportStakeChanged event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	universe_id			BIGINT NOT NULL,
	rep_stake			DECIMAL(64,18)
);
CREATE TABLE cset_buy (-- CompleteSetPurchases event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	num_sets			DECIMAL(64,18)
);
CREATE TABLE cset_sell (--CompleteSetSold event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	num_sets			DECIMAL(64,18),
	fees				DECIMAL(64,18)
);
CREATE TABLE irep_redeem ( -- InitialReporterRedeemed event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	market_aid			BIGINT NOT NULL,
	reporter_aid		BIGINT NOT NULL,
	ini_rep_aid			BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	amount				DECIMAL(64,18) NOT NULL,
	rep					DECIMAL(64,18) NOT NULL,
	payout_numerators	TEXT DEFAULT ''
);
CREATE TABLE reporter_disavowed ( -- ReportingParticipantDisavowed event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	reporter_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL
);
CREATE TABLE reporting_fee ( -- ReportingFeeChanged event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	fee_divisor			DECIMAL NOT NULL
);
CREATE TABLE rep_tok_redeem ( -- ParticipationTokensRedeemed event
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	dispwin_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	ptokens				DECIMAL(64,18),
	fee_payout			DECIMAL(64,18)
);

