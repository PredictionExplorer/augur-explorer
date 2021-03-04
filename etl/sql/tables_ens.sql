CREATE TABLE ens_node( -- strictly ENS data about a node
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	cur_owner_aid		BIGINT DEFAULT 0,	-- current owner
	cur_owner_evt		BIGINT DEFAULT 0,	-- evtlog_id of the last update of the owner aid
	time_stamp			TIMESTAMPTZ,
	label				TEXT,
	node				TEXT,
	fqdn				TEXT,			-- fully qualified domain name hash
	fqdn_words			TEXT DEFAULT '',
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(fqdn)
);
CREATE TABLE active_name( -- ENS names that are currently active (i.e. haven't expired)
	id					BIGSERIAL PRIMARY KEY,
	ensname_id			BIGINT NOT NULL, -- latest `ens_name.id` field
	expires				TIMESTAMPTZ NOT NULL,
	name				TEXT,
	label				TEXT NOT NULL,
	node				TEXT NOT NULL,
	fqdn				TEXT NOT NULL UNIQUE
);
CREATE TABLE ens_name (-- this is a complementary table to ens_onode (with more non-ENS info about node)
	id					BIGSERIAL PRIMARY KEY,
	owner_aid			BIGINT NOT NULL,
	expires				TIMESTAMPTZ,
	label				TEXT NOT NULL,
	node				TEXT NOT NULL,
	fqdn				TEXT NOT NULL,
	name				TEXT,	-- human name
	pubkey				TEXT DEFAULT '',
	content_hash		TEXT DEFAULT '',
	cost				DECIMAL(32,18)
);
CREATE TABLE ens_name_reg1(-- ENS NameRegistered1 event (signature ca6abbe9)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	owner_aid			BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	expires				TIMESTAMPTZ,
	label				TEXT NOT NULL,
	node				TEXT NOT NULL,
	fqdn				TEXT NOT NULL,
	name				TEXT,
	tx_hash				TEXT NOT NULL,
	cost				DECIMAL(32,18),
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(tx_hash,label)
);
CREATE TABLE ens_name_reg2(-- ENS NAmeRegistered2 event (signature b3d98796)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	owner_aid			BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	expires				TIMESTAMPTZ,
	label				TEXT NOT NULL,
	node				TEXT NOT NULL,
	fqdn				TEXT NOT NULL,
	tx_hash				TEXT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(tx_hash,fqdn)
);
CREATE TABLE ens_name_reg3(-- ENS NAmeRegistered3 event (signature 0f0c27ad)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	caller_aid			BIGINT NOT NULL,
	beneficiary_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	created_date		TIMESTAMPTZ,
	subdomain			TEXT NOT NULL,
	label				TEXT NOT NULL,
	node				TEXT NOT NULL,
	fqdn				TEXT NOT NULL,
	tx_hash				TEXT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(tx_hash,fqdn)
);
CREATE TABLE IF NOT EXISTS ens_label ( -- label <=> real world name, mapping
	label				TEXT UNIQUE,
	word				TEXT UNIQUE
);
CREATE TABLE ens_new_owner(
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	owner_aid			BIGINT NOT NULL,
	tx_hash				TEXT NOT NULL,
	label				TEXT NOT NULL,
	node				TEXT NOT NULL,
	fqdn				TEXT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE
);
CREATE TABLE ens_addr1 (-- AddrChanged event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	tx_hash				TEXT NOT NULL,
	fqdn				TEXT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE
);
CREATE TABLE ens_addr2(-- AddressChanged event (the event with coin type field)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	coin_type			INT NOT NULL,
	tx_hash				TEXT NOT NULL,
	fqdn				TEXT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE
);
CREATE TABLE ens_new_resolver(
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	aid					BIGINT NOT NULL,
	tx_hash				TEXT NOT NULL,
	node				TEXT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE
);
CREATE TABLE ens_hash_inval(	-- HashInvalidated event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	reg_date			TIMESTAMPTZ,
	tx_hash				TEXT NOT NULL,
	hash				TEXT NOT NULL,
	name				TEXT NOT NULL,
	value				DECIMAL(32,18),
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE
);
CREATE TABLE ens_hash_reg (	-- HashRegistered event
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	owner_id			BIGINT,
	time_stamp			TIMESTAMPTZ,
	reg_date			TIMESTAMPTZ,
	tx_hash				TEXT NOT NULL,
	hash				TEXT NOT NULL,
	value				DECIMAL(32,18),
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE
);
CREATE TABLE ens_reg_transf ( -- Transfer event on the ENS Registry contract
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	aid					BIGINT NOT NULL,
	tx_hash				TEXT NOT NULL,
	node				TEXT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE
);
CREATE TABLE ens_text_chg (
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT,
	block_num			BIGINT,			-- this is just a copy (for easy data management)
	tx_id				BIGINT,
	contract_aid		BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	tx_hash				TEXT NOT NULL,
	node				TEXT NOT NULL,
	key					TEXT NOT NULL,
	value				TEXT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE
);
CREATE TABLE ens_text_key (
	node				TEXT NOT NULL,
	key					TEXT NOT NULL,
	value				TEXT NOT NULL,
	PRIMARY KEY(node,key)
);
CREATE TABLE ens_text (	-- all keys althogether
	node				TEXT NOT NULL,
	num_keys			INT DEFAULT 0,
	all_keys			JSONB NOT NULL,
	PRIMARY KEY(node)
);
CREATE TABLE ens_status (
	block_num_limit		BIGINT DEFAULT 10543755, -- limit for initial load
	--block_num_limit		BIGINT DEFAULT 11650046, -- limit for initial load
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
