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
CREATE table tok_transf (	-- Tokens Transferred event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	market_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	from_aid			BIGINT NOT NULL,
	to_aid				BIGINT NOT NULL,
	token_type			SMALLINT DEFAULT 0,
	value				DECIMAL(64,32) DEFAULT 0.0
);
CREATE table tbc (			-- Token Balance Changed event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	market_aid			BIGINT NOT NULL,
	owner_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	token_type			SMALLINT DEFAULT 0,
	outcome				SMALLINT NOT NULL,
	balance				DECIMAL(64,32) DEFAULT 0.0
);
CREATE table stbc (			-- Share Token Balance Changed event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	market_aid			BIGINT NOT NULL,
	account_aid			BIGINT NOT NULL,
	outcome_idx			SMALLINT NOT NULL,
	balance				DECIMAL(64,32) DEFAULT 0.0
);
-- Balances of Share tokens per Market (accumulated data, one record per account)
CREATE TABLE sbalances (
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			 -- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	account_aid			BIGINT NOT NULL,			-- address id of the User(holder of the shares)
	market_aid			BIGINT NOT NULL,			-- market id of the Market these shares blong
	num_transfers		BIGINT DEFAULT 0,			-- counter for tracking now many transfers we had
	outcome_idx			SMALLINT NOT NULL,				-- market outcome (index)
	balance				DECIMAL(24,18) NOT NULL		-- balance of shares (bigint as string)
);
CREATE TABLE etl_tokens ( -- ETL process state variables to import tokens from Geth 
--single record table
	last_id_dai			BIGINT DEFAULT 0,
	last_id_rep			BIGINT DEFAULT 0, -- Rep V2 token
	last_id_stok		BIGINT DEFAULT 0, -- ShareToken ERC20 transfer
	last_id_stbc		BIGINT DEFAULT 0 -- ShareTokenBalance changed
);
CREATE TABLE chain_reorg_dai ( -- stores chain reorg events
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,
	hash				CHAR(66) NOT NULL
);
CREATE TABLE dai_proc_status (-- DAI processing status
	last_block			BIGINT DEFAULT 0
--	last_id_dai			BIGINT DEFAULT 0 --id of last event log processed
);
