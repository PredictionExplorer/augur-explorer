CREATE TABLE erc1155_tok (
	token_id			BIGSERIAL PRIMARY KEY,
	contract_aid		BIGINT NOT NULL,
	token_id_hex		TEXT NOT NULL,
	num_holders			INT DEFAULT 0,
	total_supply		DECIMAL DEFAULT 0,
	UNIQUE(contract_aid,token_id)
);
CREATE TABLE erc1155_holder (
	contract_aid		BIGINT NOT NULL,
	token_id			BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	cur_balance			DECIMAL DEFAULT 0,
	PRIMARY KEY(contract_aid,token_id,aid)
);
CREATE table erc1155_transf (	-- transfers of ERC1155  tokens
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	time_stamp			TIMESTAMPTZ,
	tx_id				BIGINT NOT NULL,
	contract_aid		BIGINT NOT NULL,
	operator_aid		BIGINT NOT NULL,
	token_id			BIGINT NOT NULL,
	from_aid			BIGINT DEFAULT 0,
	to_aid				BIGINT DEFAULT 0,
	op_type				INT DEFAULT 0,-- 0: regular transfer, 1:mint, 2: burn
	amount				DECIMAL DEFAULT 0.0,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE table erc1155_batch (	-- transfer batch event of ERC1155 tokens
	id					BIGSERIAL PRIMARY KEY,
	evtlog_id			BIGINT NOT NULL,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	time_stamp			TIMESTAMPTZ,
	tx_id				BIGINT NOT NULL,
	contract_aid		BIGINT NOT NULL,
	operator_aid		BIGINT NOT NULL,
	from_aid			BIGINT DEFAULT 0,
	to_aid				BIGINT DEFAULT 0,
	op_type				INT DEFAULT 0,-- 0: regular transfer, 1:mint, 2: burn
	token_ids			TEXT NOT NULL,
	amounts				TEXT NOT NULL,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE table erc1155_bal (	-- token balance
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,
	time_stamp			TIMESTAMPTZ,
	tx_id				BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	contract_aid		BIGINT NOT NULL,
	token_id			BIGINT NOT NULL,
	parent_id			BIGINT NOT NULL REFERENCES erc20_transf(id) ON DELETE CASCADE,
	processed			BOOLEAN DEFAULT false,	-- true if balances have been calculated
	balance				DECIMAL DEFAULT 0.0,
	amount				DECIMAL DEFAULT 0.0
);
CREATE TABLE erc1155_uri ( -- URI
	id                  BIGSERIAL PRIMARY KEY,
	evtlog_id           BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num           BIGINT NOT NULL,            -- this is just a copy (for easy data management)
	tx_id               BIGINT NOT NULL,
	time_stamp          TIMESTAMPTZ NOT NULL,
	contract_aid        BIGINT NOT NULL,
	value               TEXT DEFAULT '',
	token_type_id		TEXT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE erc1155_proc_status (-- DAI processing status
	last_evt_id			BIGINT DEFAULT 0 --id of last event log processed
);
