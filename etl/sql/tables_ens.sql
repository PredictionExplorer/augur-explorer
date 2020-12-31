CREATE TABLE ens_node(
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	time_stamp			TIMESTAMPTZ,
	label				TEXT,
	node				TEXT,
	fqdn				TEXT,			-- fully qualified domain name hash
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(fqdn)
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
	name				TEXT,
	label				TEXT,
	fqdn				TEXT UNIQUE
);
CREATE TABLE ens_label ( -- label <=> real world name, mapping
	label				TEXT UNIQUE,
	word				TEXT UNIQUE
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
	node				TEXT NOT NULL,
	fqdn				TEXT NOT NULL
);
CREATE TABLE ens_new_resolver(
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	time_stamp			TIMESTAMPTZ,
	aid					BIGINT NOT NULL,
	tx_hash				TEXT NOT NULL,
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
CREATE TABLE alexa_top1m(	-- Alexa's top 1M domain names, about 700k records
	name				TEXT,
	hash				TEXT UNIQUE	-- label hash
);
CREATE TABLE en_prop_names(	-- English proper names (list of 61k words)
	word				TEXT,
	hash				TEXT UNIQUE	-- label hash
);
CREATE TABLE email_tokens( -- Words extracted from 300million emails list dataset
	token				TEXT,
	hash				TEXT UNIQUE	-- label hash
);
CREATE TABLE pwd_db ( -- 36 million record password database
	password			TEXT,
	hash				TEXT UNIQUE	-- label hash
);
