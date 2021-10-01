CREATE OR REPLACE FUNCTION on_erc20_transf_insert() RETURNS trigger AS  $$
DECLARE
	v_aid bigint;
	v_cnt numeric;
BEGIN

	INSERT INTO erc20_tok(contract_aid) VALUES(NEW.contract_aid) ON CONFLICT DO NOTHING;
	INSERT INTO erc20_bal(block_num,tx_id,parent_id,contract_aid,time_stamp,aid,amount)
			VALUES(NEW.block_num,NEW.tx_id,NEW.id,NEW.contract_aid,NEW.time_stamp,NEW.from_aid,-NEW.amount);

	INSERT INTO erc20_bal(block_num,tx_id,parent_id,contract_aid,time_stamp,aid,amount)
			VALUES(NEW.block_num,NEW.tx_id,NEW.id,NEW.contract_aid,NEW.time_stamp,NEW.to_aid,NEW.amount);

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc20_transf_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	DELETE FROM erc20_bal WHERE parent_id = OLD.id;
	RETURN OLD;

END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc20_bal_insert() RETURNS trigger AS  $$
DECLARE
	v_cur_balance		DECIMAL;
BEGIN

	INSERT INTO erc20_holder(contract_aid,aid)
		VALUES(NEW.contract_aid,NEW.aid)
		ON CONFLICT DO NOTHING;
	UPDATE erc20_holder SET cur_balance = (cur_balance + NEW.amount)
		WHERE contract_aid=NEW.contract_aid AND aid=NEW.aid
		RETURNING cur_balance INTO v_cur_balance;
	UPDATE erc20_bal SET balance = v_cur_balance WHERE id=NEW.id;
	IF NEW.amount != 0 THEN
		IF v_cur_balance = 0 THEN
			UPDATE erc20_tok SET num_holders = (num_holders - 1)
				WHERE contract_aid=NEW.contract_aid ;
		ELSE
			UPDATE erc20_tok SET num_holders = (num_holders + 1)
				WHERE contract_aid=NEW.contract_aid;
		END IF;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc20_bal_delete() RETURNS trigger AS  $$
DECLARE
	v_cur_balance		DECIMAL;
BEGIN

	UPDATE erc20_holder SET cur_balance = (cur_balance - OLD.amount)
		WHERE contract_aid=OLD.contract_aid AND aid=OLD.aid
		RETURNING cur_balance INTO v_cur_balance;
	IF OLD.amount != 0 THEN
		IF v_cur_balance = 0 THEN
			UPDATE erc20_tok SET num_holders = (num_holders - 1)
				WHERE contract_aid=OLD.contract_aid;
		ELSE
			UPDATE erc20_tok SET num_holders = (num_holders + 1)
				WHERE contract_aid=OLD.contract_aid;
		END IF;
	END IF;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc20_bal_update() RETURNS trigger AS  $$
DECLARE
BEGIN
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
