CREATE OR REPLACE FUNCTION on_swf_hist_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_swf_hist_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	-- if a record at some block-height is deleted, subsequent records become invalid
	DELETE from swf_hist WHERE
		(
			block_num > OLD.block_num
		) OR (
			block_num = OLD.block_num AND
			tx_index > OLD.tx_index
		) OR (
			block_num = OLD.block_num AND
			tx_index = OLD.tx_index AND
			log_index > OLD.log_index
		)
	;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_poolhist_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE pool_hist SET
		total_fees = (total_fees + NEW.swap_fee),
		total_swasp = (total_swaps + 1)
	WHERE pool_id = NEW.pool_id;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_poolhist_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE pool_hist SET
		total_fees = (total_fees - OLD.swap_fee),
		total_swasp = (total_swaps - 1)
	WHERE pool_id = OLD.pool_id;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_tokbal_insert() RETURNS trigger AS  $$
DECLARE
	v_prev_bal DECIMAL;
	v_new_bal DECIMAL;
BEGIN

	SELECT balance
		FROM tok_bal
		WHERE tok_aid = NEW.tok_aid AND pool_aid=NEW.pool_aid AND id<NEW.id
		ORDER BY id DESC LIMIT 1
		INTO v_prev_bal;

	IF v_prev_bal IS NULL THEN
		v_prev_bal := 0;
	END IF;
	v_new_bal := v_prev_bal + NEW.amount * NEW.op_sign;
	UPDATE tok_bal SET balance = v_new_bal WHERE id=NEW.id;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_tokbal_delete() RETURNS trigger AS  $$
DECLARE
BEGIN


	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bpt_transf_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	UPDATE bpt_bal SET balance = (balance - NEW.amount)
		WHERE pool_aid=NEW.pool_aid AND aid=NEW.from_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO bpt_bal(aid,pool_aid,balance)
			VALUES(NEW.from_aid,NEW.pool_aid,-NEW.amount);
	END IF;

	UPDATE bpt_bal SET balance = (balance + NEW.amount)
		WHERE pool_aid=NEW.pool_aid AND aid=NEW.to_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO bpt_bal(aid,pool_aid,balance)
			VALUES(NEW.to_aid,NEW.pool_aid,NEW.amount);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bpt_transf_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE bpt_bal SET balance = (balance + NEW.amount)
		WHERE pool_aid=NEW.pool_aid AND aid=NEW.from_aid;
	UPDATE bpt_bal SET balance = (balance - NEW.amount)
		WHERE pool_aid=NEW.pool_aid AND aid=NEW.to_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_tokens_reg_insert() RETURNS trigger AS  $$
DECLARE
	v_token_addr		TEXT;
	v_pool_aid			BIGINT;
	v_tok_aid			BIGINT;
BEGIN

	IF LENGTH(NEW.tokens) > 0 THEN
		FOREACH v_token_addr IN ARRAY STRING_TO_ARRAY(NEW.tokens,',')
		LOOP
			SELECT pool_aid FROM pool_reg WHERE pool_id=NEW.pool_id INTO v_pool_aid;
			IF v_pool_aid IS NULL THEN
				RAISE EXCEPTION 'Cant locate pool_aid from pool_reg table by pool_id';
			END IF;
			SELECT address_id FROM addr WHERE addr = v_token_addr INTO v_tok_aid;
			IF v_tok_aid IS NULL THEN
				INSERT INTO addr(addr) VALUES(v_token_addr) RETURNING address_id INTO v_tok_aid;
				IF v_tok_aid IS NULL THEN
					RAISE EXCEPTION 'Failed to insert addres % in on_tokens_reg_insert trigger ',v_token_addr;
				END IF;
			END IF;
			INSERT INTO bptok(pool_aid,tok_aid,block_num_reg,tx_idx_reg)
				VALUES(v_pool_aid,v_tok_aid,NEW.block_num,NEW.tx_index)
				ON CONFLICT DO NOTHING;
		END LOOP;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_tokens_reg_delete() RETURNS trigger AS  $$
DECLARE
BEGIN


	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
