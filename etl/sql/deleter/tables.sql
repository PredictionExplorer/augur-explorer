CREATE TABLE d_status(
	last_evtlog_id			BIGINT DEFAULT 0,
	contract_addr			TEXT,
	info					TEXT,
	block_num				BIGINT DEFAULT 0
);
