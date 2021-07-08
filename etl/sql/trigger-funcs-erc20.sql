CREATE OR REPLACE FUNCTION on_erc20_transf_insert() RETURNS trigger AS  $$
DECLARE
	v_aid bigint;
	v_cnt numeric;
BEGIN

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
BEGIN

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc20_bal_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc20_bal_update() RETURNS trigger AS  $$
DECLARE
BEGIN
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
