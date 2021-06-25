-- Tables with information about Augur events at Arbitrum chain
CREATE TABLE aa_caddrs ( -- Addresses of contracts for Arbitrum Augur
	chain_id			BIGINT DEFAULT 0,
	amm_factory			TEXT DEFAULT '0x8860542771F787dD8B2c8f9134662751DE2F664f',
	--sports_factory		TEXT DEFAULT '0x43D9f2d22f1306D012251d032a5B67553FE4aA82',
	sportsball1			TEXT DEFAULT '0xBC8C695dd045FBfe81C353Fd88E3bedE45C2855D',
	sportsball2			TEXT DEFAULT '0x1ac5742415c071f376C81F6e2A7fE56eA19fb3dF',
	mma					TEXT DEFAULT '0xb2a568C444C6B74D10f7cf66bEcfeAF88a94808a',
	trusted_factory		TEXT DEFAULT '0x4117A1F75Dfe784F315AabF7dB8caf86Fc10653b'
);
CREATE TABLE aa_factory (
	factory_aid			BIGINT PRIMARY KEY,
	market_type			TINYINT DEFAULT 0, -- market types to be defined (pending)
	factory_addr		TEXT -- copy to facilitate testing
);
CREATE TABLE aa_proc_status (-- Arbitrum Augur process status
	last_evt_id			BIGINT DEFAULT 0
);
CREATE TABLE aa_pool_created (
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	pool_aid			BIGINT NOT NULL,
	factory_aid			BIGINT NOT NULL,-- Market Factory
	creator_aid			BIGINT NOT NULL,
	market_id			BIGINT NOT NULL,
	token_rcpt_aid		BIGINT NOT NULL,
	balances			TEXT,			-- balances of sharetokens, updated on liquidity/swap events
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_liquidity_changed (-- LiquidityChanged event 
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	market_id			BIGINT NOT NULL,
	factory_aid			BIGINT NOT NULL,
	user_aid			BIGINT NOT NULL,
	recipient_aid		BIGINT NOT NULL,
	collateral			DECIMAL(64,18) NOT NULL,
	lp_tokens			DECIMAL(64,18) NOT NULL,
	shares_returned		TEXT,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_price_market (-- PriceMarketCreated event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	end_time			TIMESTAMPTZ,
	creator_aid			BIGINT NOT NULL,
	aa_mkt_id			BIGINT NOT NULL REFERENCES aa_market(id),
	spot_price			DECIMAL(64,18) NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_market ( -- AbstractMarketFactory object type , parent of Sports,Trusted,Price markets
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,-- this is factory_aid
	time_stamp			TIMESTAMPTZ,
	created_time		TIMESTAMPTZ,
	end_time			TIMESTAMPTZ,
	market_id			BIGINT NOT NULL,
	collateral_aid		BIGINT NOT NULL,	-- usually USDC contract (Cash)
	protocol_aid		BIGINT NOT NULL,
	settlement_aid		BIGINT NOT NULL,
	winner_aid			BIGINT DEFAULT 0,
	sharefactor			DECIMAL NOT NULL,
	settlement_fee		DECIMAL NOT NULL,
	protocol_fee		DECIMAL NOT NULL,
	liquidity			DECIMAL(64,18) DEFAULT 0.0,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_shtok ( -- Market ShareToken (OwnedShareToken.sol)
	id					BIGSERIAL PRIMARY KEY,
	parent_id			BIGSERIAL NOT NULL,
	token_aid			BIGSERIAL NOT NULL,	-- this  should match a record in erc20_info
	FOREIGN KEY(parent_id) REFERENCES aa_market(id) ON DELETE CASCADE,
}
CREATE TABLE aa_sports_market (
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	end_time			TIMESTAMPTZ,
	start_time			TIMESTAMPTZ,	-- estimatedStartTime
	market_id			BIGINT NOT NULL,
	aa_mkt_id			BIGINT NOT NULL REFERENCES aa_market(id),
	creator_aid			BIGINT NOT NULL,
	event_id			BIGINT NOT NULL,
	home_team_id		BIGINT NOT NULL,
	away_team_id		BIGINT NOT NULL,
	market_type			INT NOT NULL,
	value0				DECIMAL NOT NULL, -- SportsLinkMarketFactory.sol::MarketDetail::value0
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_outcome ( --Outcome object, links Token contract
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	market_id			BIGINT NOT NULL, -- copied for simplicity
	aa_mkt_id			BIGINT NOT NULL REFERENCES aa_market(id),
	token_aid			BIGINT NOT NULL, -- address of ERC20 token for this outcome
	symbol				TEXT DEFAULT '',
	name				TEXT DEFAULT ''
);
CREATE TABLE aa_last_price (--populated using swap events from Balancer contracts
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	in_aid				BIGINT NOT NULL, -- address id of token contract for swapping IN
	out_aid				BIGINT NOT NULL, -- address id of token contract for swapping OUT
	outc_in				DECIMAL(64,18) DEFAULT 0.0, -- Outcome in (token1 amount)
	outc_out			DECIMAL(64,18) DEFAULT 0.0, -- Outcome out (token2 amount)
);
CREATE TABLE aa_trusted_market (
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	end_time			TIMESTAMPTZ,
	market_id			BIGINT NOT NULL,
	aa_mkt_id			BIGINT NOT NULL REFERENCES aa_market(id),
	creator_aid			BIGINT NOT NULL,
	descr				TEXT NOT NULL,
	outcomes			TEXT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_shares_minted (-- SharesMinted
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	aid					BIGINT NOT NULL,
	market_id			TEXT NOT NULL,
	amount				DECIMAL(64,18),
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_shares_burned (-- SharesBurned
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	aid					BIGINT NOT NULL,
	market_id			TEXT NOT NULL,
	amount				DECIMAL(64,18),
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_shares_swapped (-- SharesSwapped
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	market_id			BIGINT NOT NULL,
	factory_aid			BIGINT NOT NULL,
	user_aid			BIGINT NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	collateral			DECIMAL(64,18) NOT NULL,
	shares				DECIMAL(64,18) NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_sfee_claimed (-- Settlement Fee Claimed event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	settlement_aid		BIGINT NOT NULL,
	receiver_aid		BIGINT NOT NULL,
	amount				DECIMAL(64,18) NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_pfee_claimed (-- Protocol Fee Claimed event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	protocol_aid		BIGINT NOT NULL,
	amount				DECIMAL(64,18) NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_proto_chg(-- Protocol Changed
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	protocol_aid		BIGINT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_pfee_chg(-- Protocol Fee Changed
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	protocol_fee		DECIMAL(64,18) NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_sfee_chg(-- Settlement Fee Changed
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	settlement_fee		DECIMAL(64,18) NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_stk_fee_chg(-- StakerFee Changed
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	staker_fee			DECIMAL(64,18) NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_winclaim (-- WinningsClaimed
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	market_id			BIGINT NOT NULL,
	win_outc_aid		BIGINT NOT NULL, -- winning outcome addr
	receiver_aid		BIGINT NOT NULL,
	amount				DECIMAL(64,18),
	settlement_fee		DECIMAL(64,18),
	payout				DECIMAL(64,18),
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_mkt_resolved (-- MarketResolved
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	market_id			BIGINT NOT NULL,
	winner_aid			BIGINT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE TABLE aa_feepot_trsf (-- FeePot transfer event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	from_aid			BIGINT NOT NULL,
	to_aid				BIGINT NOT NULL,
	value				DECIMAL(64,18),
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);


