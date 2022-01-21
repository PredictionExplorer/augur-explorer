CREATE TABLE bs_block (	-- bigstats block tracking table
	block_num			BIGINT NOT NULL UNIQUE,
	num_tx				BIGINT DEFAULT 0,
	ts					TIMESTAMPTZ NOT NULL,
	total_eth			DECIMAL(64,18) DEFAULT 0.0,
	block_hash			CHAR(66) NOT NULL PRIMARY KEY,
	parent_hash			CHAR(66) NOT NULL
);
CREATE TABLE bs_addr (
	address_id			BIGSERIAL	PRIMARY KEY,
	addr				TEXT NOT NULL UNIQUE,		-- 20 byte Ethereum address , stored as 42 hex string (0x+addr)
	is_contract			BOOLEAN
);
CREATE TABLE bs_stats(	-- statistics accumulator table (per block)
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL REFERENCES bs_block(block_num) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	total_eth			DECIMAL,
	UNIQUE(block_num)
);
CREATE TABLE bs_period(
	id					SERIAL PRIMARY KEY,
	time_stamp			TIMESTAMPTZ NOT NULL,
	duration_sec		BIGINT DEFAULT 86400,
	--- statistics
	unique_addrs_eoa	BIGINT DEFAULT 0,
	unique_addrs_code	BIGINT DEFAULT 0,-- contract accounts
	eth_transferred		DECIMAL,
	UNIQUE(time_stamp)
);
CREATE TABLE bs_log(
	block_num			BIGINT NOT NULL,
	tx_index			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	PRIMARY KEY(block_num,tx_index,aid)
);
CREATE TABLE tx_short( -- short version of tx
	block_num			BIGINT NOT NULL REFERENCES bs_block(block_num) ON DELETE CASCADE,
	tx_index			BIGINT NOT NULL,
	tx_fee				DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index)
);
CREATE TABLE bs_config(
	chain_id			BIGINT DEFAULT 0,
	last_block			BIGINT DEFAULT 0
);
