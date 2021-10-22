CREATE TABLE rw_token(
	token_id		BIGINT PRIMARY KEY,
	seed_hex		TEXT,
	seed_num		DECIMAL,
	last_name		TEXT,
	last_price		DECIMAL,
	num_trades		BIGINT,
	total_vol		DECIMAL,	-- total trading volume
	UNIQUE(seed_hex)
);
CREATE TABLE rw_new_offer(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	rwalk_aid		BIGINT NOT NULL,	-- the address of randomwalk token contract
	offer_id		BIGINT NOT NULL,
	seller_aid		BIGINT NOT NULL,
	buyer_aid		BIGINT NOT NULL,
	otype			SMALLINT NOT NULL, --0-buy offer,1-sell offer
	token_id		TEXT NOT NULL,
	active			BOOLEAN,
	price			DECIMAL,
	UNIQUE(offer_id),
	UNIQUE(evtlog_id)
);
CREATE TABLE rw_item_bought(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	offer_id		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE rw_offer_canceled(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	offer_id		BIGINT NOT NULL,
	UNIQUE(evtlog_id)
);
CREATE TABLE rw_withdrawal (
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	aid				BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	amount			DECIMAL,
	UNIQUE(evtlog_id)
);
CREATE TABLE rw_token_name(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	new_name		TEXT,
	UNIQUE(evtlog_id)
);
CREATE TABLE rw_mint_evt(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	owner_aid		BIGINT NOT NULL,
	seed			TEXT NOT NULL,
	seed_num		DECIMAL,		-- seed as numeric 256 bit integer
	price			DECIMAL
);
CREATE TABLE rw_transfer(
	id				BIGSERIAL PRIMARY KEY,
	evtlog_id		BIGINT REFERENCES evt_log(id) ON DELETE CASCADE,
	block_num		BIGINT NOT NULL,
	tx_id			BIGINT NOT NULL,
	time_stamp		TIMESTAMPTZ NOT NULL,
	contract_aid	BIGINT NOT NULL,
	token_id		BIGINT NOT NULL,
	from_aid		BIGINT NOT NULL,
	to_aid			BIGINT NOT NULL,
	otype			SMALLINT NOT NULL,-- 0-regular transfer,1-Mint,2-Burn
	UNIQUE(evtlog_id)
);
CREATE TABLE rw_stats(
	rwalk_aid				BIGINT PRIMARY KEY,
	total_vol				DECIMAL DEFAULT 0,		-- total volume
	total_num_trades		BIGINT DEFAULT 0,		-- total count of trade operations made
	total_num_toks			BIGINT DEFAULT 0,		-- total count of tokens registered
	total_withdrawals		BIGINT DEFAULT 0
);
CREATE TABLE rw_user_stats (
	rwalk_aid				BIGINT PRIMARY KEY,
	user_aid				BIGINT NOT NULL,
	total_vol				DECIMAL DEFAULT 0,		-- total volume
	total_num_trades		BIGINT DEFAULT 0,		-- total count of tokens traded by user
	total_num_toks			BIGINT DEFAULT 0,		-- total count of tokens minted by user
	total_withdrawals		BIGINT DEFAULT 0,
	UNIQUE(rwalk_aid,user_aid)
);
CREATE TABLE rw_proc_status (
	last_evt_id             BIGINT DEFAULT 0,
	last_block              BIGINT DEFAULT 0 -- used when getting event logs via ethclient.FilterLogs
);
CREATE TABLE rw_contracts (
	marketplace_addr		TEXT DEFAULT '0x70cf513E1fE1C481144f7FF967036eb05A6bc045',
	randomwalk_addr			TEXT DEFAULT '0x332E5e3dE89cDe8131aCCdd09ecbd51Ea4B9b0E3'
	--marketplace_addr		TEXT DEFAULT '0x728A419D264532442ea9CF639ec6a766f64840d6',
	--randomwalk_addr			TEXT DEFAULT '0x27fAFD053dD7e4E5349F90bd32c8233D3d3c0235'
);
