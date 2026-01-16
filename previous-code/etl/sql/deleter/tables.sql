CREATE TABLE d_contracts(	-- contracts which must have data
	contract_addr			TEXT PRIMARY KEY,
	info					TEXT DEFAULT ''
);
CREATE TABLE d_status(
	last_block_num			BIGINT DEFAULT 0
);
