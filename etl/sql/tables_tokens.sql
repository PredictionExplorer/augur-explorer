--Layer2 tables:  Token processing layer
CREATE table dai_transf (	-- transfers of DAI tokens (deposits/withdrawals of funds)
	id					BIGSERIAL PRIMARY KEY,
--	dai_proc_id			BIGINT NOT NULL REFERENCES dai_proc(id) ON DELETE CASCADE,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	from_aid			BIGINT DEFAULT 0,
	to_aid				BIGINT DEFAULT 0,
	from_internal		BOOLEAN DEFAULT false,
	to_internal			BOOLEAN DEFAULT false,
	amount				DECIMAL(64,18) DEFAULT 0.0
);
CREATE table dai_bal (	-- DAI token balance
	id					BIGSERIAL PRIMARY KEY,
	dai_transf_id		BIGINT NOT NULL REFERENCES dai_transf(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,
	tx_id				BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	processed			BOOLEAN DEFAULT false,	-- true if balances have been calculated
	augur				BOOLEAN DEFAULT false,	-- true if the user has account on Augur Platform
	internal			BOOLEAN DEFAULT false,	-- true if it is an exchange between Agur's contracts
	balance				DECIMAL(64,18) DEFAULT 0.0,
	amount				DECIMAL(64,18) DEFAULT 0.0
);
--CREATE table dai_proc ( -- DAI processing marker used for cascading delets in case of block goes away
--	id					BIGSERIAL PRIMARY KEY,
--	block_num			BIGINT NOT NULL,
--	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
--	evtlog_id			BIGINT NOT NULL UNIQUE
--);
CREATE table rep_transf (
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	from_aid			BIGINT DEFAULT 0,
	to_aid				BIGINT DEFAULT 0,
	amount				DECIMAL(32,18) DEFAULT 0.0
);
CREATE TABLE etl_tokens ( -- ETL process state variables to import tokens from Geth 
--single record table
	last_id_dai			BIGINT DEFAULT 0,
	last_id_rep			BIGINT DEFAULT 0, -- Rep V2 token
	last_id_stok		BIGINT DEFAULT 0, -- ShareToken ERC20 transfer
	last_id_stbc		BIGINT DEFAULT 0 -- ShareTokenBalance changed
);
CREATE TABLE token_proc_status (-- DAI processing status
	last_evt_id			BIGINT DEFAULT 0 --id of last event log processed
);
CREATE TABLE af_wrapper ( -- Augur Foundry wrapper (wraps ShareToken to ERC20 contract)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	wrapper_aid			BIGINT NOT NULL,	-- address_id of ERC20 contract wrapping this share token
	market_aid			BIGINT NOT NULL,
	last_evt_id			BIGINT DEFAULT 0,	-- the event ID of ERC20 Transfer event processed last time
	outcome_idx			INT NOT NULL,
	decimals			INT DEFAULT 0,
	time_stamp			BIGINT DEFAULT 0,	-- timestamp copied from block
	token_id			TEXT,	-- hex encode token id (ShareToken format)
	name				TEXT DEFAULT '',
	symbol				TEXT DEFAULT '',
	UNIQUE(wrapper_aid)
);
CREATE TABLE wstok_transf ( -- ERC20 Wrapped ShareToken transfer
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	wrapper_aid			BIGINT NOT NULL,			-- foreign key to af_wrapper.wrapper_aid
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	from_aid			BIGINT DEFAULT 0,
	to_aid				BIGINT DEFAULT 0,
	amount				DECIMAL(32,18) DEFAULT 0.0,
	balance				DECIMAL(32,18) DEFAULT 0.0
);
CREATE TABLE af_addr (
	augur_foundry_addr	TEXT DEFAULT '0x87876F172087E2fb5838E655DC6A929dC2Dcf85c'
);
CREATE TABLE af_status (
	last_evt_id			BIGINT DEFAULT 0	-- event id
);

