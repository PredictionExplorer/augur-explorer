CREATE TABLE af_wrapper ( -- Augur Foundry wrapper (wraps ShareToken to ERC20 contract)
	id					BIGSERIAL PRIMARY KEY,
	block_num			BIGINT NOT NULL,			-- this is just a copy (for easy data management)
	tx_id				BIGINT NOT NULL REFERENCES transaction(id) ON DELETE CASCADE,
	token_id			TEXT,	-- hex encode token id (ShareToken format)
	wrapper_aid			BIGINT NOT NULL,	-- address_id of ERC20 contract wrapping this share token
	time_stamp			BIGINT DEFAULT 0,	-- timestamp copied from block
	name				TEXT,
	symbol				TEXT,
	decimals			INT DEFAULT 0
);
