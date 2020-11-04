CREATE TABLE bpool ( -- Balancer Pool creation event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL UNIQUE,
	caller_aid			BIGINT NOT NULL,
	controller_aid		BIGINT DEFAULT 0,
	num_swaps			BIGINT DEFAULT 0,
	num_holders			BIGINT DEFAULT 0,
	num_tokens			BIGINT DEFAULT 0,
	went_public			BIGINT DEFAULT 0,-- block number of when the pool went public
	was_finalied		BIGINT DEFAULT 0,-- block number of when the pool was finalized
	swap_fee			DECIMAL(64,18) DEFAULT 0.0
);
CREATE TABLE bbind (-- Balancer bind function calls
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL, -- token address linked to this pool
	balance				DECIMAL(64,18) NOT NULL,
	denorm				BIGINT NOT NULL
);
CREATE TABLE bjoin ( -- Join event to join balancer pool
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	caller_aid			BIGINT NOT NULL,
	token_aid			BIGINT NOT NULL,
	amount_in			DECIMAL(64,18)
);
CREATE TABLE bexit ( -- Join event to join balancer pool
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			BIGINT NOT NULL,
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
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			BIGINT NOT NULL,
	pool_aid			BIGINT NOT NULL,
	caller_aid			BIGINT NOT NULL,
	token_in_aid		BIGINT NOT NULL,
	token_out_aid		BIGINT NOT NULL,
	amount_in			DECIMAL(64,18),
	amount_out			DECIMAL(64,18)
);
CREATE TABLE balancer_status (
	last_evt_id			BIGINT DEFAULT 0	-- event id
);
CREATE TABLE balancer_contracts (
	pool_token			TEXT DEFAULT '0x6B74Fb4E4b3B177b8e95ba9fA4c3a3121d22fbfB',
	factory				TEXT DEFAULT '0x9424B1412450D0f8Fc2255FAf6046b98213B76Bd',
	xchg_proxy			TEXT DEFAULT '0x3E66B66Fd1d0b02fDa6C811Da9E0547970DB2f21'
);
