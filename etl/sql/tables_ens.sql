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
CREATE TABLE ens_name(
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	time_stamp			TIMESTAMPTZ,
	owner_aid			BIGINT NOT NULL,
	label				TEXT,
	cost				DECIMAL(32,18),
);
CREATE TABLE ens_status (
	block_num_limit		BIGINT DEFAULT 10543755, -- limit for initial load
	last_evt_id			BIGINT DEFAULT 0	-- event id (latest processed)
);
