CREATE TABLE ens_node(
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	time_stamp			TIMESTAMPTZ,
	label				TEXT UNIQUE,
	node				TEXT,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE
);
CREATE TABLE ens_name( -- NameRegistered_v1 event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	time_stamp			TIMESTAMPTZ,
	expires				TIMESTAMPTZ,
	owner_aid			BIGINT NOT NULL,
	label				TEXT,
	name				TEXT,
	tx_hash				TEXT NOT NULL,
	cost				DECIMAL(32,18),
	UNIQUE(tx_hash,label)
);
CREATE TABLE active_name( -- ENS names that are currently active (i.e. haven't expired)
	id					BIGSERIAL PRIMARY KEY,
	ensname_id			BIGINT NOT NULL, -- latest `ens_name.id` field
	expires				TIMESTAMPTZ NOT NULL,
	prev_expires		TIMESTAMPTZ,
	name				TEXT UNIQUE,
	label				TEXT
);
CREATE TABLE ens_new_owner(
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	time_stamp			TIMESTAMPTZ,
	owner_aid			BIGINT NOT NULL,
	tx_hash				TEXT NOT NULL,
	label				TEXT NOT NULL,
	node				TEXT NOT NULL
);
CREATE TABLE ens_hash_inval(	-- HashInvalidated event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	time_stamp			TIMESTAMPTZ,
	reg_date			TIMESTAMPTZ,
	tx_hash				TEXT NOT NULL,
	hash				TEXT NOT NULL,
	name				TEXT NOT NULL,
	value				DECIMAL(32,18)
);
CREATE TABLE ens_status (
	block_num_limit		BIGINT DEFAULT 10543755, -- limit for initial load
	last_evt_id			BIGINT DEFAULT 0	-- event id (latest processed)
);
