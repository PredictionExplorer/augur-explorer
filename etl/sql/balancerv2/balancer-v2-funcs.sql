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
	v_new_bal := v_prev_bal + NEW.amount;
	UPDATE tok_bal SET balance = v_new_bal WHERE id=NEW.id;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_tokbal_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	DELETE from tok_bal
	WHERE
		(
			block_num > OLD.block_num
		) OR (
			(block_num = OLD.block_num) AND
			(tx_index > OLD.tx_index)
		) OR (
			(block_num = OLD.block_num) AND
			(tx_index = OLD.tx_index) AND
			(log_index > OLD.log_index)
		);

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
