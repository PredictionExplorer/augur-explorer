CREATE TABLE upair( -- uniswap pair (PairCreated event)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pair_aid			BIGINT NOT NULL,
	token0_aid			BIGINT NOT NULL,
	token1_aid			BIGINT NOT NULL,
	total_swaps			BIGINT DEFAULT 0,	--total number of swaps ocurred
	pair_seq			INT NOT NULL UNIQUE,-- sequential number of Pair creation
	UNIQUE(token0_aid,token1_aid)
);
CREATE TABLE uswap1( -- original swap event as it comes from Uniswap Pair contract
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	pair_aid			BIGINT NOT NULL,
	sender_aid			BIGINT NOT NULL,
	recipient_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	amount0_in			DECIMAL(64,18) NOT NULL,
	amount1_in			DECIMAL(64,18) NOT NULL,
	amount0_out			DECIMAL(64,18) NOT NULL,
	amount1_out			DECIMAL(64,18) NOT NULL
);
CREATE TABLE uswap2( -- shorter version of swap, splitting tokens for easier data querying
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	uswap1_id			BIGINT NOT NULL,	-- parent record in 'uswap' table
	aid					BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	amount				DECIMAL(64,18) NOT NULL -- negative means 'in', positive 'out'
);
CREATE TABLE uswap_stats(
	aid					BIGINT PRIMARY KEY,
	token_aid			BIGINT NOT NULL,
	num_swaps			BIGINT DEFAULT 0,
	volume				DECIMAL(128,18) DEFAULT 0.0,
	UNIQUE(aid,token_aid)
);
CREATE TABLE u_slippage (
	pair_aid			BIGINT NOT NULL,
	upd_block_num		BIGINT NOT NULL,
	token_in			BIGINT NOT NULL,
	token_out			BIGINT NOT NULL,
	slippage			DECIMAL(64,18),
	amount_in			DECIMAL(64,18),
	amount_out			DECIMAL(64,18),
	PRIMARY KEY(pair_aid,token_in,token_out)
);
CREATE TABLE uniswap_status (
	last_evt_id			BIGINT DEFAULT 0	-- event id (latest processed)
);
CREATE TABLE uniswap_contracts ( -- this is Uniswap V2
	uniswap_router01	TEXT DEFAULT '0xf164fC0Ec4E93095b804a4795bBe1e041497b92a', -- Version 1 of Router
	uniswap_router02	TEXT DEFAULT '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D', -- Version 2 of Router
	uniswap_factory		TEXT DEFAULT '0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f', -- Factory contract (V2)
	dai_token_addr		TEXT DEFAULT '0x6B175474E89094C44Da98b954EedeAC495271d0F',
	weth_addr			TEXT DEFAULT '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2'
);
