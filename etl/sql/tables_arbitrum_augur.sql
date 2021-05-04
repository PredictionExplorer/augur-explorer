-- Tables with information about Augur events at Arbitrum chain
CREATE TABLE aa_caddrs ( -- Addresses of contracts for Arbitrum Augur
	chain_id			BIGINT DEFAULT 0,
	amm_factory			TEXT DEFAULT '0xa513E6E4b8f2a923D98304ec87F64353C4D5C853',
	hatchery_reg		TEXT DEFAULT '0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9'
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
	spot_price			DECIMAL(64,18) NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
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
	creator_aid			BIGINT NOT NULL,
	event_id			BIGINT NOT NULL,
	home_team_id		BIGINT NOT NULL,
	away_team_id		BIGINT NOT NULL,
	score				BIGINT NOT NULL,
	market_type			INT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
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
	settlement_fee		DECIMAL(64,18) NOT NULL,
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
	receiver_aid		BIGINT NOT NULL,
	amount				DECIMAL(64,18),
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


