-- Balancer v2 tables
CREATE TABLE block (	-- bigstats block tracking table
	block_num			BIGINT NOT NULL UNIQUE,
	num_tx				BIGINT DEFAULT 0,
	time_stamp			TIMESTAMPTZ NOT NULL,
	block_hash			CHAR(66) NOT NULL PRIMARY KEY,
	parent_hash			CHAR(66) NOT NULL
);
CREATE TABLE addr (
	address_id			BIGSERIAL	PRIMARY KEY,
	addr				TEXT NOT NULL UNIQUE		-- 20 byte Ethereum address , stored as 42 hex string (0x+addr)
);
CREATE TABLE config(
	chain_id			BIGINT DEFAULT 0, --Arbitrum: 42161
	last_block			BIGINT DEFAULT 0,
	starting_block		BIGINT DEFAULT 0
);
CREATE TABLE swap ( -- Swap() event
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	tx_index			INT NOT NULL,
	log_index			INT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	pool_id				BIGINT NOT NULL,
	token_in_aid		BIGINT NOT NULL,
	token_out_aid		BIGINT NOT NULL,
	amount_in			DECIMAL,
	amount_out			DECIMAL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE pool_reg (	-- PoolRegistered event
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	tx_index			INT NOT NULL,
	log_index			INT NOT NULL,
	pool_id				BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	specialization		INT NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE pool_created (	-- PoolCreated event
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	tx_index			INT NOT NULL,
	log_index			INT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE tokens_reg (	-- TokensRegistered event
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	tx_index			INT NOT NULL,
	log_index			INT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	pool_id				BIGINT NOT NULL,
	tokens				TEXT NOT NULL,
	managers			TEXT NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE tokens_unreg (	-- TokensUnregistered event
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	tx_index			INT NOT NULL,
	log_index			INT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	pool_id				BIGINT NOT NULL,
	tokens				TEXT NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE ibalance (	-- InternalBalanceChanged event
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	tx_index			INT NOT NULL,
	log_index			INT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	delta				DECIMAL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE ebal_transf (	-- ExternalBalanceTransfer event
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	tx_index			INT NOT NULL,
	log_index			INT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	delta				DECIMAL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE pool_bal (	-- PoolBalanceChanged event
	block_num			BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp			TIMESTAMPTZ NOT NULL,
	tx_index			INT NOT NULL,
	log_index			INT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	pool_id				BIGINT NOT NULL,
	tokens				TEXT,
	deltas				TEXT,
	fee_amounts			TEXT,
	PRIMARY KEY(block_num,tx_index,log_index)
);
