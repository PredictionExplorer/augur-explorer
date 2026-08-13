-- +goose Up
-- The RandomWalk stats triggers silently lose counts when an event is
-- processed before the target stats row exists:
--   * on_mint_event_insert(): the rw_stats insert-if-missing fallback only
--     set total_num_toks, so the FIRST mint of a contract never reached
--     money_accumulated;
--   * on_mint_event_insert(): the rw_user_stats total_num_toks increment was
--     a plain UPDATE with no insert-if-missing fallback, so every mint by a
--     user with no prior marketplace activity was dropped from the counter;
--   * on_new_offer_insert(): the rw_mkt_stats order counters were plain
--     UPDATEs with no fallback, and the rw_mkt_stats row is only created by
--     the first purchase, so offers placed before the first trade of a
--     marketplace contract were dropped from total_sell/buy_orders.
-- All are visible as drift flagged by scripts/validate-stats.sql; existing
-- rows are backfilled by scripts/repair-rw-stats.sql.

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
		INSERT INTO rw_stats(rwalk_aid,total_num_toks,money_accumulated)
			VALUES(NEW.contract_aid,1,NEW.price);
	END IF;
	UPDATE rw_user_stats
		SET total_num_toks = (total_num_toks + 1)
		WHERE rwalk_aid=NEW.contract_aid AND user_aid=NEW.owner_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_user_stats(rwalk_aid,user_aid,total_num_toks)
			VALUES(NEW.contract_aid,NEW.owner_aid,1);
	END IF;
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
CREATE OR REPLACE FUNCTION on_new_offer_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN

	IF NEW.otype = 1 THEN
		UPDATE rw_mkt_stats SET 
				total_sell_orders = (total_sell_orders + 1)
			WHERE contract_aid=NEW.contract_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO rw_mkt_stats(contract_aid,total_sell_orders)
				VALUES(NEW.contract_aid,1);
		END IF;
	ELSE
		UPDATE rw_mkt_stats SET 
				total_buy_orders = (total_buy_orders + 1)
			WHERE contract_aid=NEW.contract_aid;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			INSERT INTO rw_mkt_stats(contract_aid,total_buy_orders)
				VALUES(NEW.contract_aid,1);
		END IF;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose Down
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
