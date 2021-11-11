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
CREATE OR REPLACE FUNCTION on_token_name_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
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
	UPDATE rw_stats SET total_num_toks = (total_num_toks +  1) WHERE rwalk_aid=NEW.contract_aid;
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
CREATE OR REPLACE FUNCTION on_mint_event_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE rw_user_stats
		SET total_num_toks = (total_num_toks - 1)
		WHERE rwalk_aid=OLD.contract_aid AND user_aid=OLD.owner_aid;
	UPDATE rw_stats SET total_num_toks = (total_num_toks - 1 ) WHERE rwalk_aid=OLD.contract_aid;
	UPDATE rw_user_rwtok
		SET price_bought=NULL
		WHERE rwalk_aid=OLD.contract_aid AND user_aid=OLD.owner_aid AND token_id=OLD.token_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
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
CREATE OR REPLACE FUNCTION on_rw_transfer_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN

	UPDATE rw_token SET
			cur_owner_aid = NEW.to_aid
		WHERE token_id=NEW.token_id AND rwalk_aid=NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_token(rwalk_aid,token_id,cur_owner_aid)
			VALUES(NEW.contract_aid,NEW.token_id,NEW.to_aid);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_rw_transfer_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	
	-- we do not restore previous token because there will be an INSERT anyway
	-- since the transaction was already signed and will be processed in the future
	-- and any possible failure of this transaction will be an extremely rare event
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
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
CREATE OR REPLACE FUNCTION on_new_offer_update() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN
	IF OLD.profit != NEW.profit THEN
		UPDATE rw_user_stats SET total_profit = (total_profit + (NEW.profit - COALESCE(OLD.profit,0)))
		WHERE user_aid=NEW.seller_aid AND rwalk_aid=NEW.rwalk_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO rw_user_stats(rwalk_aid,user_aid,total_profit)
				VALUES(NEW.rwalk_aid,NEW.seller_aid,(NEW.profit - COALESCE(OLD.profit,0)));
		END IF;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
