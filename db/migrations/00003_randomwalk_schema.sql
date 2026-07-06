-- +goose Up
CREATE TABLE rw_token(
	token_id		BIGINT NOT NULL,
	rwalk_aid		BIGINT NOT NULL,
	cur_owner_aid	BIGINT NOT NULL,
	seed_hex		TEXT DEFAULT NULL,
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
	msg_text					TEXT DEFAULT 'New token '
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

-- RandomWalk token ranking (pairwise "cool or not" style), migrated from Python games/tokens.rating.
-- Apply after randomwalk base tables (needs rw_contracts for default rwalk_aid in app logic).

CREATE TABLE IF NOT EXISTS rw_ranking_match (
	id				BIGSERIAL PRIMARY KEY,
	nft1			BIGINT NOT NULL,
	nft2			BIGINT NOT NULL,
	nft1_won		BOOLEAN NOT NULL,
	voter_aid		BIGINT REFERENCES address(address_id),
	created_at		TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_rw_ranking_match_nft1 ON rw_ranking_match(nft1);
CREATE INDEX IF NOT EXISTS idx_rw_ranking_match_nft2 ON rw_ranking_match(nft2);

CREATE TABLE IF NOT EXISTS rw_ranking_vote_nonce (
	nonce		TEXT PRIMARY KEY,
	expires_at	TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_rw_ranking_vote_nonce_expires ON rw_ranking_vote_nonce(expires_at);

CREATE UNIQUE INDEX IF NOT EXISTS idx_rw_ranking_match_voter_pair
	ON rw_ranking_match (voter_aid, LEAST(nft1, nft2), GREATEST(nft1, nft2))
	WHERE voter_aid IS NOT NULL;

-- Elo-style rating per token (single RandomWalk collection).
CREATE TABLE IF NOT EXISTS rw_token_ranking (
	token_id		BIGINT PRIMARY KEY,
	rating			DOUBLE PRECISION NOT NULL DEFAULT 1200,
	updated_at		TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Wallet-signed beauty votes: voter address id + one vote per voter per unordered token pair.
-- Apply after token_ranking.sql and address table exist.

ALTER TABLE rw_ranking_match ADD COLUMN IF NOT EXISTS voter_aid BIGINT REFERENCES address(address_id);

CREATE TABLE IF NOT EXISTS rw_ranking_vote_nonce (
	nonce		TEXT PRIMARY KEY,
	expires_at	TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_rw_ranking_vote_nonce_expires ON rw_ranking_vote_nonce(expires_at);

CREATE UNIQUE INDEX IF NOT EXISTS idx_rw_ranking_match_voter_pair
	ON rw_ranking_match (voter_aid, LEAST(nft1, nft2), GREATEST(nft1, nft2))
	WHERE voter_aid IS NOT NULL;

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_item_bought_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
	v_rwalk_aid				BIGINT;
	v_token_id				BIGINT;
	v_msg_sender_aid		BIGINT;
	v_offer_type			SMALLINT;
	v_price					DECIMAL;
	v_price_bought			DECIMAL;
BEGIN

	UPDATE rw_new_offer
		SET active=FALSE
		WHERE offer_id=NEW.offer_id AND contract_aid=NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'Offer % not found',NEW.offer_id;
	END IF;

	SELECT rwalk_aid,token_id,price,otype
		FROM rw_new_offer
		WHERE offer_id=NEW.offer_id AND contract_aid=NEW.contract_aid
		INTO v_rwalk_aid,v_token_id,v_price,v_offer_type;
	IF v_rwalk_aid IS NULL THEN
		RAISE EXCEPTION 'Offer %v not found when looking up for rwalk_aid',NEW.offer_id;
	END IF;
	SELECT from_aid FROM transaction WHERE id=NEW.tx_id INTO v_msg_sender_aid;
	IF v_offer_type = 1::SMALLINT THEN -- sell
		UPDATE rw_new_offer SET buyer_aid=NEW.buyer_aid
			WHERE offer_id=NEW.offer_id AND contract_aid=NEW.contract_aid;
	ELSE
		UPDATE rw_new_offer SET seller_aid=NEW.seller_aid
			WHERE offer_id=NEW.offer_id AND contract_aid=NEW.contract_aid;
	END IF;
	UPDATE rw_stats
		SET		total_vol = (total_vol + v_price),
				total_num_trades = (total_num_trades +1)
		WHERE rwalk_aid = v_rwalk_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_stats(rwalk_aid,total_vol,total_num_trades)
			VALUES(v_rwalk_aid,v_price,1);
	END IF;
	UPDATE rw_mkt_stats
		SET		total_vol = (total_vol + v_price),
				total_num_trades = (total_num_trades +1)
		WHERE contract_aid = NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_mkt_stats(contract_aid,total_vol,total_num_trades)
			VALUES(NEW.contract_aid,v_price,1);
	END IF;
	UPDATE rw_token SET
			last_price=v_price,
			num_trades=(num_trades+1),
			total_vol=(total_vol+v_price)
		WHERE token_id=v_token_id AND rwalk_aid=v_rwalk_aid;
	UPDATE rw_user_stats SET
			total_num_trades = (total_num_trades+1),
			total_vol = (total_vol+v_price)
		WHERE user_aid=NEW.buyer_aid AND rwalk_aid=v_rwalk_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_user_stats(rwalk_aid,user_aid,total_num_trades,total_vol)
			VALUES(v_rwalk_aid,NEW.buyer_aid,1,v_price);
	END IF;
	IF NEW.buyer_aid != NEW.seller_aid THEN
		UPDATE rw_user_stats SET
				total_num_trades = (total_num_trades+1),
				total_vol = (total_vol+v_price)
			WHERE user_aid=NEW.seller_aid AND rwalk_aid=v_rwalk_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO rw_user_stats(rwalk_aid,user_aid,total_num_trades,total_vol)
				VALUES(v_rwalk_aid,NEW.seller_aid,1,v_price);
		END IF;
	END IF;
	IF v_offer_type = 1 THEN
		UPDATE rw_mkt_stats SET
				total_sell_orders = (total_sell_orders - 1)
			WHERE contract_aid=NEW.contract_aid;
	ELSE
		UPDATE rw_mkt_stats SET
				total_buy_orders = (total_buy_orders - 1)
			WHERE contract_aid=NEW.contract_aid;
	END IF;

	SELECT price_bought FROM rw_user_rwtok
		WHERE rwalk_aid=v_rwalk_aid AND user_aid=NEW.seller_aid AND token_id=v_token_id
		INTO v_price_bought;
	IF v_price_bought IS NOT NULL THEN
		UPDATE rw_new_offer SET profit = (v_price - v_price_bought)
		WHERE contract_aid=NEW.contract_aid AND offer_id=NEW.offer_id;
		UPDATE rw_user_stats SET total_profit = (total_profit + (v_price - v_price_bought))
			WHERE user_aid=NEW.seller_aid AND rwalk_aid=v_rwalk_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO rw_user_stats(rwalk_aid,user_aid,total_profit)
				VALUES(v_rwalk_aid,NEW.seller_aid,v_price - v_price_bought);
		END IF;
	END IF;
	UPDATE rw_user_rwtok
		SET price_bought = NULL
		WHERE rwalk_aid=v_rwalk_aid AND user_aid=NEW.seller_aid AND token_id=v_token_id;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_user_rwtok(rwalk_aid,user_aid,token_id,price_bought)
			VALUES(v_rwalk_aid,NEW.seller_aid,v_token_id,NULL);
	END IF;
	UPDATE rw_user_rwtok
		SET price_bought = v_price
		WHERE rwalk_aid=v_rwalk_aid AND user_aid=NEW.buyer_aid AND token_id=v_token_id;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_user_rwtok(rwalk_aid,user_aid,token_id,price_bought)
			VALUES(v_rwalk_aid,NEW.buyer_aid,v_token_id,v_price);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_item_bought_delete() RETURNS trigger AS  $$
DECLARE
	v_rwalk_aid				BIGINT;
	v_token_id				BIGINT;
	v_buyer_aid				BIGINT;
	v_seller_aid			BIGINT;
	v_price					DECIMAL;
	v_offer_type			SMALLINT;
BEGIN

	SELECT rwalk_aid,token_id,buyer_aid,seller_aid,price,otype
		FROM rw_new_offer
		WHERE offer_id=OLD.offer_id AND contract_aid=OLD.contract_aid
		INTO v_rwalk_aid,v_token_id,v_buyer_aid,v_seller_aid,v_price,v_offer_type;
	IF v_rwalk_aid IS NULL THEN
		RETURN OLD;
	END IF;
	UPDATE rw_stats
		SET		total_vol = (total_vol - v_price),
				total_num_trades = (total_num_trades -1)
		WHERE rwalk_aid = v_rwalk_aid;
	UPDATE rw_mkt_stats
		SET		total_vol = (total_vol - v_price),
				total_num_trades = (total_num_trades -1)
		WHERE contract_aid = NEW.contract_aid;

	UPDATE rw_new_offer SET active=TRUE WHERE offer_id=OLD.offer_id AND contract_aid=OLD.contract_aid;
	UPDATE rw_token SET
			num_trades=(num_trades - 1),
			total_vol=(total_vol-v_price)
		WHERE token_id=v_token_id AND rwalk_aid=v_rwalk_aid;
	UPDATE rw_user_stats SET
			total_num_trades = (total_num_trades-1),
			total_vol = (total_vol-v_price)
		WHERE user_aid=v_buyer_aid AND rwalk_aid=v_rwalk_aid;
	if v_buyer_aid != v_seller_aid THEN 
		UPDATE rw_user_stats SET
				total_num_trades = (total_num_trades-1),
				total_vol = (total_vol-v_price)
			WHERE user_aid=v_seller_aid AND rwalk_aid=v_rwalk_aid;
	END IF;
	IF v_offer_type = 1 THEN
		UPDATE rw_mkt_stats SET
				total_sell_orders = (total_sell_orders + 1)
			WHERE contract_aid=OLD.contract_aid;
	ELSE
		UPDATE rw_mkt_stats SET
				total_buy_orders = (total_buy_orders + 1)
			WHERE contract_aid=OLD.contract_aid;
	END IF;
	UPDATE rw_user_rwtok
		SET price_bought = v_price
		WHERE rwalk_aid=v_rwalk_aid AND user_aid=NEW.seller_aid AND token_id=v_token_id;
	UPDATE rw_user_rwtok
		SET price_bought = NULL
		WHERE rwalk_aid=v_rwalk_aid AND user_aid=NEW.buyer_aid AND token_id=v_token_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_token_name_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN

	UPDATE rw_token SET last_name=NEW.new_name
		WHERE token_id=NEW.token_id AND rwalk_aid=NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'Token ID % not found',NEW.token_id;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_token_name_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_mint_event_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN

	UPDATE rw_token SET	-- UPDATE is used because the record is inserted during Transfer event
			seed_hex=NEW.seed,
			seed_num=NEW.seed_num,
			last_price=NEW.price
		WHERE token_id=NEW.token_id AND rwalk_aid=NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_token(rwalk_aid,token_id,cur_owner_aid,seed_hex,seed_num,last_price)
			VALUES(NEW.contract_aid,NEW.token_id,NEW.owner_aid,NEW.seed,NEW.seed_num,NEW.price);
	END IF;
	UPDATE rw_stats SET 
			total_num_toks = (total_num_toks +  1),
			money_accumulated = (money_accumulated + NEW.price)
		WHERE rwalk_aid=NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_stats(rwalk_aid,total_num_toks)
			VALUES(NEW.contract_aid,1);
	END IF;
	UPDATE rw_user_stats
		SET total_num_toks = (total_num_toks + 1)
		WHERE rwalk_aid=NEW.contract_aid AND user_aid=NEW.owner_aid;
	UPDATE rw_user_rwtok
		SET price_bought = NEW.price
		WHERE rwalk_aid=NEW.contract_aid AND user_aid=NEW.owner_aid AND token_id=NEW.token_id;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_user_rwtok(rwalk_aid,user_aid,token_id,price_bought)
			VALUES(NEW.contract_aid,NEW.owner_aid,NEW.token_id,NEW.price);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_mint_event_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE rw_user_stats
		SET total_num_toks = (total_num_toks - 1)
		WHERE rwalk_aid=OLD.contract_aid AND user_aid=OLD.owner_aid;
	UPDATE rw_stats SET
			total_num_toks = (total_num_toks - 1 ),
			money_accumulated = (money_accumulated - OLD.price)
		WHERE rwalk_aid=OLD.contract_aid;
	UPDATE rw_user_rwtok
		SET price_bought=NULL
		WHERE rwalk_aid=OLD.contract_aid AND user_aid=OLD.owner_aid AND token_id=OLD.token_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_offer_canceled_insert() RETURNS trigger AS  $$
DECLARE
	v_offer_type			SMALLINT;
	v_cnt                   NUMERIC;
BEGIN

	SELECT otype FROM rw_new_offer
		WHERE offer_id=NEW.offer_id AND contract_aid=NEW.contract_aid
		INTO v_offer_type;
	UPDATE rw_new_offer SET active=FALSE
		WHERE offer_id=NEW.offer_id AND contract_aid=NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'Offer % not found',NEW.offer_id;
	END IF;
	IF v_offer_type = 1 THEN
		UPDATE rw_mkt_stats SET
				total_sell_orders = (total_sell_orders - 1)
			WHERE contract_aid=NEW.contract_aid;
	ELSE
		UPDATE rw_mkt_stats SET
				total_buy_orders = (total_buy_orders - 1)
			WHERE contract_aid=NEW.contract_aid;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_offer_canceled_delete() RETURNS trigger AS  $$
DECLARE
	v_offer_type			SMALLINT;
BEGIN

	SELECT otype FROM rw_new_offer
		WHERE offer_id=OLD.offer_id AND contract_aid=OLD.contract_aid
		INTO v_offer_type;
	IF v_offer_Type = 1 THEN
		UPDATE rw_mkt_stats SET 
				total_sell_orders = (total_sell_orders + 1)
			WHERE contract_aid=OLD.contract_aid;
	ELSE
		UPDATE rw_mkt_stats SET 
				total_buy_orders = (total_buy_orders + 1)
			WHERE contract_aid=OLD.contract_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_rw_transfer_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN

	UPDATE rw_token SET
			cur_owner_aid = NEW.to_aid
		WHERE token_id=NEW.token_id AND rwalk_aid=NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		-- Use NULL, not default '' : UNIQUE(seed_hex) allows many NULLs but only one empty string,
		-- so mint/burn placeholder rows must not all share ''.
		INSERT INTO rw_token(rwalk_aid,token_id,cur_owner_aid,seed_hex)
			VALUES(NEW.contract_aid,NEW.token_id,NEW.to_aid,NULL);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_rw_transfer_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	
	-- we do not restore previous token because there will be an INSERT anyway
	-- since the transaction was already signed and will be processed in the future
	-- and any possible failure of this transaction will be an extremely rare event
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_new_offer_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN

	IF NEW.otype = 1 THEN
		UPDATE rw_mkt_stats SET 
				total_sell_orders = (total_sell_orders + 1)
			WHERE contract_aid=NEW.contract_aid;
	ELSE
		UPDATE rw_mkt_stats SET 
				total_buy_orders = (total_buy_orders + 1)
			WHERE contract_aid=NEW.contract_aid;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_new_offer_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF OLD.otype = 1 THEN
		UPDATE rw_mkt_stats SET 
				total_sell_orders = (total_sell_orders - 1)
			WHERE contract_aid=OLD.contract_aid;
	ELSE
		UPDATE rw_mkt_stats SET 
				total_buy_orders = (total_buy_orders - 1)
			WHERE contract_aid=OLD.contract_aid;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_new_offer_update() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN
	IF OLD.profit != NEW.profit THEN
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER rw_item_bougnt_insert AFTER INSERT ON rw_item_bought FOR EACH ROW EXECUTE PROCEDURE on_item_bought_insert();
CREATE TRIGGER rw_item_bought_delete AFTER DELETE ON rw_item_bought FOR EACH ROW EXECUTE PROCEDURE on_item_bought_delete();
CREATE TRIGGER rw_token_name_insert AFTER INSERT ON rw_token_name FOR EACH ROW EXECUTE PROCEDURE on_token_name_insert();
CREATE TRIGGER rw_token_name_delete AFTER DELETE ON rw_token_name FOR EACH ROW EXECUTE PROCEDURE on_token_name_delete();
CREATE TRIGGER rw_mint_event_insert AFTER INSERT ON rw_mint_evt FOR EACH ROW EXECUTE PROCEDURE on_mint_event_insert();
CREATE TRIGGER rw_mint_event_delete AFTER DELETE ON rw_mint_evt FOR EACH ROW EXECUTE PROCEDURE on_mint_event_delete();
CREATE TRIGGER rw_offer_canceled_insert AFTER INSERT ON rw_offer_canceled FOR EACH ROW EXECUTE PROCEDURE on_offer_canceled_insert();
CREATE TRIGGER rw_offer_canceled_delete AFTER DELETE ON rw_offer_canceled FOR EACH ROW EXECUTE PROCEDURE on_offer_canceled_delete();
CREATE TRIGGER rw_transfer_insert AFTER INSERT ON rw_transfer FOR EACH ROW EXECUTE PROCEDURE on_rw_transfer_insert();
CREATE TRIGGER rw_transfer_delete AFTER DELETE ON rw_transfer FOR EACH ROW EXECUTE PROCEDURE on_rw_transfer_delete();
CREATE TRIGGER rw_new_offer_insert AFTER INSERT ON rw_new_offer FOR EACH ROW EXECUTE PROCEDURE on_new_offer_insert();
CREATE TRIGGER rw_new_offer_delete AFTER DELETE ON rw_new_offer FOR EACH ROW EXECUTE PROCEDURE on_new_offer_delete();
CREATE TRIGGER rw_new_offer_update AFTER UPDATE on rw_new_offer FOR EACH ROW EXECUTE PROCEDURE on_new_offer_update();

CREATE INDEX rw_ofact_idx			ON	rw_new_offer			(active);
CREATE INDEX rw_trsf_tx_idx			ON	rw_transfer				(tx_id);

-- +goose Down
DROP TABLE IF EXISTS
	rw_ranking_vote_nonce,
	rw_ranking_match,
	rw_token_ranking,
	rw_token,
	rw_new_offer,
	rw_item_bought,
	rw_offer_canceled,
	rw_withdrawal,
	rw_token_name,
	rw_mint_evt,
	rw_transfer,
	rw_notif_status,
	rw_messaging_status,
	rw_stats,
	rw_mkt_stats,
	rw_user_stats,
	rw_user_rwtok,
	rw_uranks,
	rw_proc_status,
	rw_contracts
CASCADE;

