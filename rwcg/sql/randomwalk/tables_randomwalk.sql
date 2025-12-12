CREATE TABLE rw_token(
	token_id		BIGINT NOT NULL,
	rwalk_aid		BIGINT NOT NULL,
	cur_owner_aid	BIGINT NOT NULL,
	seed_hex		TEXT DEFAULT '',
	seed_num		DECIMAL DEFAULT 0,
	last_name		TEXT DEFAULT '',
	last_price		DECIMAL DEFAULT 0,
	num_trades		BIGINT DEFAULT 0,
	total_vol		DECIMAL DEFAULT 0,	-- total trading volume
	PRIMARY KEY(rwalk_aid,token_id),
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
	token_id		BIGINT NOT NULL,
	active			BOOLEAN,
	price			DECIMAL,
	profit			DECIMAL,		-- profit the seller made (if used MarketPlace contract for buy and sell operation)
	UNIQUE(contract_aid,offer_id),
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
	seller_aid		BIGINT NOT NULL,
	buyer_aid		BIGINT NOT NULL,
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
	total_withdrawals		BIGINT DEFAULT 0,
	money_accumulated		DECIMAL DEFAULT 0,		-- sum of all the mints (by price)
	UNIQUE(rwalk_aid)
);
CREATE TABLE rw_mkt_stats( -- statistis per market (can include many token contracts)
	contract_aid			BIGINT PRIMARY KEY,
	total_vol				DECIMAL DEFAULT 0,		-- total volume
	total_num_trades		BIGINT DEFAULT 0,		-- total count of trade operations made
	total_buy_orders		BIGINT DEFAULT 0,
	total_sell_orders		BIGINT DEFAULT 0,
	UNIQUE(contract_aid)
);
CREATE TABLE rw_user_stats (
	rwalk_aid				BIGINT,
	user_aid				BIGINT NOT NULL,
	total_vol				DECIMAL DEFAULT 0,		-- total volume
	total_num_trades		BIGINT DEFAULT 0,		-- total count of tokens traded by user
	total_num_toks			BIGINT DEFAULT 0,		-- total count of tokens minted by user
	total_withdrawals		BIGINT DEFAULT 0,
	total_profit			DECIMAL DEFAULT 0,
	PRIMARY KEY(rwalk_aid,user_aid)
);
CREATE TABLE rw_user_rwtok (	-- hold info of User-Token relation (to calculate profit)
									-- this profit is only available for trades made at MarketPlace
	rwalk_aid				BIGINT NOT NULL,
	user_aid                BIGINT NOT NULL,
	token_id                BIGINT NOT NULL,
	price_bought            DECIMAL NULL,   -- NOT NULL - position opened (price of the token when BUY order was executed)
											---NULL	 - no position for this token
	PRIMARY KEY(rwalk_aid,user_aid,token_id)
);
CREATE TABLE rw_uranks (   -- User Rankings (how this user ranks against each other, ex: Top 13% in profit made
	aid		            BIGINT PRIMARY KEY,
	total_trades		BIGINT DEFAULT 0,
	top_profit          DECIMAL(24,2) DEFAULT 100.0,    -- position of the user in profits accumulated over lifetime
	top_trades          DECIMAL(24,2) DEFAULT 100.0,    -- position of the user in number of accumulated trades
	top_volume			DECIMAL(24,2) DEFAULT 100.0,	   -- position of the user in accumulated trading volume
	profit				DECIMAL(64,18) DEFAULT 0.0,
	volume				DECIMAL(64,18) DEFAULT 0.0
);
CREATE TABLE rw_notif_status ( -- Status of Tweeter/Discord notifications
	last_token_id_tweeter		BIGINT DEFAULT 0,
	last_token_id_discord		BIGINT DEFAULT 0,
	msg_text					TEXT DEFALT 'New token '
);
CREATE TABLE rw_proc_status (
	last_evt_id             BIGINT DEFAULT 0,
	last_block              BIGINT DEFAULT 0 -- used when getting event logs via ethclient.FilterLogs
);
CREATE TABLE rw_contracts (
	--marketplace_addr		TEXT DEFAULT '0x70cf513E1fE1C481144f7FF967036eb05A6bc045',
	marketplace_addr		TEXT DEFAULT '0x52266bdbfa301803a62FCF7B3C946EF1c3f7691E',
	randomwalk_addr			TEXT DEFAULT '0x332E5e3dE89cDe8131aCCdd09ecbd51Ea4B9b0E3'
	--marketplace_addr		TEXT DEFAULT '0x728A419D264532442ea9CF639ec6a766f64840d6',
	--randomwalk_addr			TEXT DEFAULT '0x27fAFD053dD7e4E5349F90bd32c8233D3d3c0235'
);
CREATE TABLE rw_messaging_status (-- Status of the notification process
	last_tx_id			BIGINT DEFAULT 0,	-- last tx_id for which notification was sent successfuly
	last_evtlog_id		BIGINT DEFAULT 0,
	last_block_num		BIGINT DEFAULT 0,
	last_timestamp		BIGINT DEFAULT 0
);
