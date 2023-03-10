CREATE TABLE bw_prize_claim(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	prize_num		BIGINT NOT NULL,
	winner_aid		BIGINT NOT NULL,
	amount			DECIMAL DEFAULT 0,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_bid (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	bidder_aid		BIGINT NOT NULL,
	rwalk_nft_id	BIGINT NOT NULL,	--token_id of RandomWalk, if present
	bid_price		DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_donation (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	donor_aid		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_donation_received (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	donor_aid		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_donation_sent (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	charity_aid		BIGINT NOT NULL,
	amount			DECIMAL NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_charity_updated (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	charity_aid		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_token_name (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	token_name		TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_mint_event(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	owner_aid		BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	seed			TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE bw_contracts (
	bidding_war_addr		TEXT,
	cosmic_signature_addr	TEXT,
	cosmic_token_addr		TEXT,
	charity_wallet_addr		TEXT
);
CREATE TABLE rw_proc_status (
	last_evt_id             BIGINT DEFAULT 0
);
