-- Uniswap v3 tables
CREATE TABLE block (    -- bigstats block tracking table
	block_num           BIGINT NOT NULL UNIQUE,
	num_tx              BIGINT DEFAULT 0,
	time_stamp          TIMESTAMPTZ NOT NULL,
	block_hash          CHAR(66) NOT NULL PRIMARY KEY,
	parent_hash         CHAR(66) NOT NULL
);
CREATE TABLE addr (
	address_id          BIGSERIAL   PRIMARY KEY,
	addr                TEXT NOT NULL UNIQUE        -- 20 byte Ethereum address , stored as 42 hex string (0x+addr)
);
CREATE TABLE config(
	chain_id            BIGINT DEFAULT 0, --Arbitrum: 42161
	last_block          BIGINT DEFAULT 0,
	starting_block      BIGINT DEFAULT 0,
	factory_addr        TEXT DEFAULT '0x1F98431c8aD98523631AE4a59f267346ea31F984',
	nft_pos_mgr_addr	TEXT DEFAULT '0xC36442b4a4522E871399CD717aBDD847Ab11FE88'-- NonfungiblePositionManager
);
CREATE TABLE pool_created ( -- PoolCreated event
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	pool_aid            BIGINT NOT NULL,
	token0_aid			BIGINT NOT NULL,
	token1_aid			BIGINT NOT NULL,
	fee					DECIMAL NOT NULL,
	tick_spacing		DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE initialize( -- Initialize event
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	sqrt_pricex96		DECIMAL NOT NULL,
	tick				DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE mint( -- Mint event
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	sender_aid			BIGINT NOT NULL,
	owner_aid			BIGINT NOT NULL,
	tick_lower			DECIMAL NOT NULL,
	tick_upper			DECIMAL NOT NULL,
	amount				DECIMAL NOT NULL,
	amount0				DECIMAL NOT NULL,
	amount1				DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE collect( -- Collect event
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	owner_aid			BIGINT NOT NULL,
	recipient_aid		BIGINT NOT NULL,
	tick_lower			DECIMAL NOT NULL,
	tick_upper			DECIMAL NOT NULL,
	amount0				DECIMAL NOT NULL,
	amount1				DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE collect_nfpm( -- Collect  (Periphery) event (from NonFungiblePositionManager contract)
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	recipient_aid		BIGINT NOT NULL,
	token_id			DECIMAL NOT NULL,
	amount0				DECIMAL NOT NULL,
	amount1				DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE burn( -- Burn event
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	owner_aid			BIGINT NOT NULL,
	tick_lower			DECIMAL NOT NULL,
	tick_upper			DECIMAL NOT NULL,
	amount				DECIMAL NOT NULL,
	amount0				DECIMAL NOT NULL,
	amount1				DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE swap( -- Swap event
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	sender_aid			BIGINT NOT NULL,
	recipient_aid		BIGINT NOT NULL,
	amount0				DECIMAL NOT NULL,
	amount1				DECIMAL NOT NULL,
	sqrt_pricex96		DECIMAL NOT NULL,
	liquidity			DECIMAL NOT NULL,
	tick				DECIMAL NOT NULL,
	fee					DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE flash( -- Flash event
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	sender_aid			BIGINT NOT NULL,
	recipient_aid		BIGINT NOT NULL,
	amount0				DECIMAL NOT NULL,
	amount1				DECIMAL NOT NULL,
	paid0				DECIMAL NOT NULL,
	paid1				DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE iocn ( -- IncreaseObservationCardinalityNext event
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	next_old			INT NOT NULL,
	next_new			INT NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE set_fee_proto( -- SetFeeProtocol event
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	fee_protocol0_old	INT NOT NULL,
	fee_protocol1_old	INT NOT NULL,
	fee_protocol0_new	INT NOT NULL,
	fee_protocol1_new	INT NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE collect_prot ( -- CollectProtocol event
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	sender_aid			BIGINT NOT NULL,
	recipient_aid		BIGINT NOT NULL,
	amount0				DECIMAL NOT NULL,
	amount1				DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE inc_liq( -- IncreaseLiquidity event (from Periphery contract)
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	token_id			DECIMAL NOT NULL,
	liquidity			DECIMAL NOT NULL,
	amount0				DECIMAL NOT NULL,
	amount1				DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE dec_liq( -- DecreaseLiquidity event (from Periphery contract)
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,
	contract_aid        BIGINT NOT NULL,
	token_id			DECIMAL NOT NULL,
	liquidity			DECIMAL NOT NULL,
	amount0				DECIMAL NOT NULL,
	amount1				DECIMAL NOT NULL,
	PRIMARY KEY(block_num,tx_index,log_index)
);
CREATE TABLE dbg_swap_loop(
	block_num           BIGINT NOT NULL REFERENCES block(block_num) ON DELETE CASCADE,
	time_stamp          TIMESTAMPTZ NOT NULL,
	tx_index            INT NOT NULL,
	log_index           INT NOT NULL,	-- this will be changing when sliding through ticks
										--	(it isnt log index of MainNet, but log index of local states)
										--	(the block num and tx_index is of MainNet though)
	contract_aid        BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	tick				INT NOT NULL,
	sqrt_price			DECIMAL,
	price				DECIMAL,	-- decoded price (decoded from fixed point 96 binary format)
	liquidity			DECIMAL,
	step_amount_in		DECIMAL,
	step_amount_out		DECIMAL,
	fee_amount			DECIMAL,
	fee_growthX128		DECIMAL,	-- variable name in .sol files is FeeGrowthGlobalX128
	fee_growth			DECIMAL
);

