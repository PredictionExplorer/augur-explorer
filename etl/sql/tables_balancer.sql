CREATE TABLE bpool ( -- Balancer Pool creation event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL UNIQUE,
	caller_aid			BIGINT NOT NULL,
	controller_aid		BIGINT NOT NULL,
	num_swaps			BIGINT DEFAULT 0,
	num_holders			BIGINT DEFAULT 0,
	num_tokens			BIGINT DEFAULT 0,
	went_public			BIGINT DEFAULT 0,-- block number of when the pool went public
	was_finalized		BIGINT DEFAULT 0,-- block number of when the pool was finalized
	total_weight		INT DEFAULT 0,
	is_public			BOOLEAN DEFAULT FALSE,
	went_public_ts		TIMESTAMPTZ DEFAULT TO_TIMESTAMP(0),
	finalized_ts		TIMESTAMPTZ DEFAULT TO_TIMESTAMP(0),
	swap_fee			DECIMAL(64,18) DEFAULT 0.0,
	usd_liquidity		DECIMAL(64,18) DEFAULT 0.0 -- Pool's liquidity calculated in US dollars
);
CREATE TABLE btoken ( -- Token contained in Balancer pool
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	denorm				DECIMAL(32,18) DEFAULT 0.0,
	balance				DECIMAL(64,18) DEFAULT 0.0
);
CREATE TABLE bjoin ( -- Join event to join balancer pool
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	caller_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	amount_in			DECIMAL(64,18)
);
CREATE TABLE bexit ( -- Join event to join balancer pool
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	caller_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	amount_out			DECIMAL(64,18)
);
CREATE TABLE bholder ( -- User who holds liquidity for the pool
	pool_aid			BIGINT NOT NULL,
	holder_aid			BIGINT NOT NULL,
	amount				DECIMAL(64,18) DEFAULT 0.0,
	PRIMARY KEY(pool_aid,holder_aid)
);
CREATE TABLE bswap ( -- Balancer swap events
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	caller_aid			BIGINT NOT NULL,
	token_in_aid		BIGINT NOT NULL,
	token_out_aid		BIGINT NOT NULL,
	amount_in			DECIMAL(64,18),
	amount_out			DECIMAL(64,18)
);
CREATE TABLE b_set_swap_fee (-- setSwapFee() calls
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	fee					DECIMAL(32,18)
);
CREATE TABLE b_set_controller (-- setController() calls
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	controller_aid		BIGINT NOT NULL,
	old_controller_aid	BIGINT NOT NULL
);
CREATE TABLE b_set_public ( -- setPublicSwap() calls
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	is_public			BOOLEAN NOT NULL,
	old_is_public		BOOLEAN NOT NULL,
	old_went_public		BIGINT NOT NULL,
	old_went_public_ts	BIGINT NOT NULL
);
CREATE TABLE b_finalized ( -- finazlie() calls
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL
);
CREATE TABLE b_bind (-- Balancer bind() function calls
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,	-- duplicated events occur due to additional call to rebind()
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL, -- token address linked to this pool
	denorm				DECIMAL(32,18) NOT NULL,
	balance				DECIMAL(64,18) NOT NULL,
	UNIQUE(pool_aid,token_aid)
);
CREATE TABLE b_unbind (-- Balancer unbind() function calls
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL, -- token address linked to this pool
	saved_denorm		DECIMAL(32,18) NOT NULL,
	saved_balance		DECIMAL(64,18) NOT NULL,
	UNIQUE(pool_aid,token_aid)
);
CREATE TABLE b_rebind (-- Balancer rebind() function calls
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL, -- token address linked to this pool
	denorm				DECIMAL(32,18) NOT NULL,
	balance				DECIMAL(64,18) NOT NULL,
	saved_denorm		DECIMAL(32,18) NOT NULL,
	saved_balance		DECIMAL(64,18) NOT NULL
);
CREATE TABLE b_gulp (-- Balancer gulp() function calls
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL UNIQUE REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	pool_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL, -- token address linked to this pool
	abs_balance			DECIMAL(64,18) DEFAULT 0.0 -- absorbed balance
);
CREATE TABLE balancer_status (
	last_evt_id			BIGINT DEFAULT 0	-- event id
);
CREATE TABLE balancer_contracts (
	pool_token			TEXT DEFAULT '0x6B74Fb4E4b3B177b8e95ba9fA4c3a3121d22fbfB',
	factory				TEXT DEFAULT '0x9424B1412450D0f8Fc2255FAf6046b98213B76Bd',
	xchg_proxy			TEXT DEFAULT '0x3E66B66Fd1d0b02fDa6C811Da9E0547970DB2f21'
);
