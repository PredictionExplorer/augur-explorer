CREATE OR REPLACE FUNCTION on_bjoin_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	INSERT INTO bholder(pool_aid,holder_aid)
		VALUES(NEW.pool_aid,NEW.caller_aid)
		ON CONFLICT DO NOTHING;
	UPDATE bholder SET amount = amount + NEW.amount_in
		WHERE pool_aid=NEW.pool_aid AND holder_aid=NEW.caller_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bjoin_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bholder SET amount = amount - OLD.amount_in
		WHERE pool_aid=OLD.pool_aid AND holder_aid=OLD.caller_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bexit_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	INSERT INTO bholder(pool_aid,holder_aid)
		VALUES(NEW.pool_aid,NEW.caller_aid)
		ON CONFLICT DO NOTHING;
	UPDATE bholder SET amount = amount - NEW.amount_out
		WHERE pool_aid=NEW.pool_aid AND holder_aid=NEW.caller_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bexit_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bholder SET amount = amount + OLD.amount_out
		WHERE pool_aid=OLD.pool_aid AND holder_aid=OLD.caller_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bswap_insert() RETURNS trigger AS  $$
DECLARE
	v_wrapper_aid bigint;
	v_market_aid bigint;
	v_cnt numeric;
BEGIN
	UPDATE bpool SET num_swaps = num_swaps + 1
		WHERE pool_aid=NEW.pool_aid;
	UPDATE b_swaps_per_pair SET num_swaps = num_swaps + 1
		WHERE pool_aid=NEW.pool_aid AND token1_aid=NEW.token1_aid AND token2_aid=NEW.token2_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO b_swaps_per_pair(pool_aid,token1_aid,token2_aid,num_swaps,last_price_block,last_amount1,last_amount2)
		VALUES(NEW.pool_aid,NEW.token1_aid,NEW.token2_aid,1,NEW.block_num,NEW.token1_amount,NEW.token2_amount);
	END IF;
	UPDATE b_swaps_per_pair SET
			last_price_block = NEW.block_num,
			last_amount1 = NEW.token1_amount,
			last_amount2 = NEW.token2_amount
		WHERE
				pool_aid=NEW.pool_aid AND
				token1_aid=NEW.token1_aid AND
				token2_aid=NEW.token2_aid AND
				last_price_block < NEW.block_num;
	SELECT wrapper_aid,market_aid FROM af_wrapper
		WHERE wrapper_aid IN(NEW.token_in_aid,NEW.token_out_aid) LIMIT 1
		INTO v_wrapper_aid,v_market_aid;
	IF v_wrapper_aid IS NOT NULL THEN
		PERFORM insert_agtx_event(
			NEW.tx_id,NEW.evtlog_id,NEW.block_num,NEW.caller_aid,v_market_aid,2,1
		);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bswap_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bpool SET num_swaps = num_swaps - 1
		WHERE pool_aid=OLD.pool_aid;
	UPDATE b_swaps_per_pair SET num_swaps = num_swaps - 1
		WHERE pool_aid=NEW.pool_aid AND token1_aid=NEW.token1_aid AND token2_aid=NEW.token2_aid;
	PERFORM delete_agtx_event(OLD.tx_id,OLD.evtlog_id);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bholder_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE bpool SET num_holders = num_holders + 1
		WHERE pool_aid=NEW.pool_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bholder_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bpool SET num_holders = num_holders - 1
		WHERE pool_aid=OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_bind_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	INSERT INTO btoken(evtlog_id,block_num,tx_id,time_stamp,pool_aid,token_aid,denorm,balance)
		VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.time_stamp,NEW.pool_aid,NEW.token_aid,NEW.denorm,NEW.balance)
		ON CONFLICT DO NOTHING;
	UPDATE bpool
		SET
			num_tokens = (num_tokens + 1),
			total_weight = (total_weight + NEW.denorm)
		WHERE pool_aid=NEW.pool_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_bind_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	DELETE FROM btoken WHERE pool_aid=OLD.pool_aid AND token_aid=OLD.token_aid;
	UPDATE bpool
		SET
			num_tokens = (num_tokens - 1),
			total_weight = (total_weight - OLD.denorm)
		WHERE pool_aid=OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_unbind_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	DELETE FROM btoken WHERE pool_aid=NEW.pool_aid AND token_aid=NEW.token_aid;
	UPDATE bpool
		SET
			num_tokens = (num_tokens - 1),
			total_weight = (total_weight - NEW.saved_denorm)
		WHERE pool_aid=NEW.pool_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_unbind_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	INSERT INTO btoken(block_num,tx_id,time_stamp,pool_aid,token_aid,denorm,balance)
		VALUES(OLD.block_num,OLD,tx_id,OLD.time_stamp,OLD.pool_aid,OLD.token_aid,OLD.saved_denorm,OLD.saved_balance)
		ON CONFLICT DO NOTHING;
	UPDATE bpool
		SET
			num_tokens = (num_tokens + 1),
			total_weight = (total_weight + OLD.saved_denorm)
		WHERE pool_aid=OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_rebind_insert() RETURNS trigger AS  $$
DECLARE
	v_old_denorm decimal;
	v_denorm_diff decimal;
	v_old_balance decimal;
	v_balance_diff decimal;
BEGIN
	SELECT denorm,balance FROM btoken WHERE pool_aid=NEW.pool_aid AND token_aid=NEW.token_aid
		INTO v_old_denorm,v_old_balance;
	IF v_old_denorm IS NULL THEN
		v_old_denorm := 0.0;
	END IF;
	IF v_old_balance IS NULL THEN
		v_old_balance := 0.0;
	END IF;
	v_denorm_diff := NEW.denorm - v_old_denorm;
	v_balance_diff := NEW.balance - v_old_balance;
	UPDATE btoken
		SET
			denorm = (denorm + v_denorm_diff),
			balance = (v_old_balance + v_balance_diff)
		WHERE pool_aid=NEW.pool_aid AND token_aid=NEW.token_aid;
	UPDATE bpool
		SET
			total_weight = (total_weight + v_denorm_diff)
		WHERE pool_aid=NEW.pool_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_rebind_delete() RETURNS trigger AS  $$
DECLARE
	v_old_denorm decimal;
	v_denorm_diff decimal;
	v_old_balance decimal;
	v_balance_diff decimal;
BEGIN
	SELECT denorm,balance FROM btoken WHERE pool_aid=OLD.pool_aid AND token_aid=OLD.token_aid
		INTO v_old_denorm,v_old_balance;
	IF v_old_denorm IS NULL THEN
		v_old_denorm := 0.0;
	END IF;
	IF v_old_balance IS NULL THEN
		v_old_balance := 0.0;
	END IF;
	v_denorm_diff := v_old_denorm - OLD.denorm;
	v_balance_diff := v_old_balance - OLD.balance;
	UPDATE btoken SET 
			denorm = (denorm + v_denorm_diff),
			balance = (v_old_balance + v_balance_diff)
		WHERE pool_aid=OLD.pool_aid AND token_aid=OLD.token_aid;
	UPDATE bpool
		SET
			total_weight = (total_weight + v_denorm_diff)
		WHERE pool_aid=OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_gulp_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE btoken
		SET balance = (balance + NEW.abs_balance)
	WHERE pool_aid = NEW.pool_aid AND token_aid = NEW.token_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_gulp_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE btoken
		SET balance = (balance + OLD.abs_balance)
	WHERE pool_aid = OLD.pool_aid AND token_aid = OLD.token_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_set_swap_fee_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE bpool
		SET
			swap_fee = NEW.fee
		WHERE pool_aid = NEW.pool_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_set_swap_fee_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bpool
		SET
			swap_fee = 0
		WHERE pool_aid = OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_set_controller_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE bpool
		SET
			controller_aid = NEW.controller_aid
		WHERE pool_aid = NEW.pool_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_set_controller_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bpool
		SET
			controller_aid = OLD.old_controller_aid
		WHERE pool_aid = OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_set_public_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	IF NEW.is_public THEN
		UPDATE bpool
			SET
				is_public = NEW.is_public,
				went_public = NEW.block_num,
				went_public_ts = NEW.time_stamp
			WHERE pool_aid = NEW.pool_aid;
	ELSE
		UPDATE bpool
			SET
				is_public = NEW.is_public,
				went_public = 0,
				went_public_ts = TO_TIMESTAMP(0)
			WHERE pool_aid = NEW.pool_aid;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_set_public_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bpool
		SET
			is_public = OLD.old_is_public,
			went_public = OLD.old_went_public,
			went_public_ts = OLD.old_went_public_ts
		WHERE pool_aid = OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_finalized_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE bpool
		SET
			was_finalized = NEW.block_num,
			finalized_ts = NEW.time_stamp,
			went_public = NEW.block_num,	-- 'finalized' implies 'public'
			went_public_ts = NEW.time_stamp
		WHERE pool_aid = NEW.pool_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_b_finalized_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bpool
		SET
			was_finalized = 0,
			finalized_ts = 0
		WHERE pool_aid = OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
