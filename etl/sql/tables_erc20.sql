CREATE table erc20_transf (	-- transfers of ERC20 tokens
	id					BIGSERIAL PRIMARY KEY,
--	dai_proc_id			BIGINT NOT NULL REFERENCES dai_proc(id) ON DELETE CASCADE,
	evtlog_id			BIGINT NOT NULL,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL,
	contract_aid		BIGINT NOT NULL,
	from_aid			BIGINT DEFAULT 0,
	to_aid				BIGINT DEFAULT 0,
	amount				DECIMAL DEFAULT 0.0,
	FOREIGN KEY(evtlog_id) REFERENCES evt_log(id) ON DELETE CASCADE,
	UNIQUE(evtlog_id)
);
CREATE table erc20_bal (	-- token balance
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,
	tx_id				BIGINT NOT NULL,
	aid					BIGINT NOT NULL,
	contract_aid		BIGINT NOT NULL,
	parent_id			BIGINT NOT NULL REFERENCES erc20_transf(id) ON DELETE CASCADE,
	processed			BOOLEAN DEFAULT false,	-- true if balances have been calculated
	balance				DECIMAL DEFAULT 0.0,
	amount				DECIMAL DEFAULT 0.0
);
CREATE TABLE erc20_proc_status (-- DAI processing status
	last_evt_id			BIGINT DEFAULT 0 --id of last event log processed
);
