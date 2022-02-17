CREATE TABLE ep_swap (	-- bigstats block tracking table
	tx_hash				TEXT PRIMARY KEY,
	time_stamp			TIMESTAMPTZ NOT NULL,
	block_num			BIGINT NOT NULL,
	log_idx				INT NOT NULL,
	token_code			INT NOT NULL, -- 0 - UDSC , 1 - DAI
	sender				TEXT NOT NULL,
	recipient			TEXT NOT NULL,
	amount0				DECIMAL,
	amount1				DECIMAL,
	sqrt_price			DECIMAL,
	liquidity			DECIMAL,
	tick				DECIMAL,
	PRIMARY KEY(tx_hash,log_idx)
);
