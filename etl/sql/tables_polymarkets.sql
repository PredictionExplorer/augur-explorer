CREATE TABLE pol_cond_prep (	-- table to store ConditionPreparation event of conditional token (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	oracle_aid			BIGINT NOT NULL,
	condition_id		TEXT NOT NULL,
	question_id			TEXT NOT NULL,
	outcome_slot_count	DECIMAL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_cond_res (-- ConditionResolution, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	oracle_aid			BIGINT NOT NULL,
	condition_id		TEXT NOT NULL,
	question_id			TEXT NOT NULL,
	outcome_slot_count	DECIMAL,
	payout_numerators	TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_pos_split ( -- PositionSplit, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	stakeholder_aid		BIGINT NOT NULL,
	collateral_aid		BIGINT NOT NULL,
	parent_coll_id		BIGINT NOT NULL, -- parent collection id
	condition_id		TEXt NOT NULL,
	partition			TEXT NOT NULL,
	amount				DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_pos_merge ( -- PositionMerge, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	stakeholder_aid		BIGINT NOT NULL,
	collateral_aid		BIGINT NOT NULL,
	parent_coll_id		BIGINT NOT NULL, -- parent collection id
	condition_id		TEXT NOT NULL,
	partition			TEXT NOT NULL,
	amount				DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE pol_uri ( -- URI, event of ConditionalToken (Gnosis)
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ NOT NULL,
	contract_aid		BIGINT NOT NULL,
	value				TEXT,
	uri_id				DECIMAL,
	UNIQUE(evtlog_id)
);
CREATE TABLE poly_proc_status (
	last_evt_id			BIGINT DEFAULT 0
);
