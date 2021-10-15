CREATE TABLE rw_new_offer(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	offer_id		BIGINT NOT NULL
	seller_aid		BIGINT NOT NULL,
	buyer_aid		BIGINT NOT NULL,
	token_id		TEXT NOT NULL,
	active			BOOLEAN,
	price			DECIMAL,
	UNIQUE(offer_id)
);
CREATE TABLE rw_item_bought(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	offer_id		BIGINT NOT NULL
);
CREATE TABLE rw_offer_cancelled(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	offer_id		BIGINT NOT NULL
);
CREATE TABLE rw_withdrawal (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	aid				BIGINT NOT NULL,
	token_id		TEXT NOT NULL,
	amount			DECIMAL
);
CREATE TABLE rw_withdrawal (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	token_id		TEXT NOT NULL,
	new_name		TEXT
);
CREATE TABLE rw_proc_status (
	last_evt_id             BIGINT DEFAULT 0,
	last_block              BIGINT DEFAULT 0 -- used when getting event logs via ethclient.FilterLogs
);
CREATE TABLE rw_contracts (
	marketplace_addr		TEXT DEFAULT '',
	randomwalk_addr			TEXT DEFAULT ''
);
