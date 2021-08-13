CREATE OR REPLACE FUNCTION on_erc1155_transf_insert() RETURNS trigger AS  $$
DECLARE
	v_aid bigint;
	v_cnt numeric;
BEGIN

	INSERT INTO erc1155_bal(block_num,tx_id,parent_id,contract_aid,token_id,time_stamp,aid,amount)
			VALUES(NEW.block_num,NEW.tx_id,NEW.id,NEW.contract_aid,NEW.token_id,NEW.time_stamp,NEW.from_aid,-NEW.amount);

	INSERT INTO erc1155_bal(block_num,tx_id,parent_id,contract_aid,token_id,time_stamp,aid,amount)
			VALUES(NEW.block_num,NEW.tx_id,NEW.id,NEW.contract_aid,NEW.token_id,NEW.time_stamp,NEW.to_aid,NEW.amount);
	IF NEW.op_type = 1 THEN -- mint
		UPDATE erc1155_tok SET total_supply=(total_supply + NEW.amount)
			WHERE contract_aid=NEW.contract_aid AND token_id=NEW.token_id;
	END IF;
	IF NEW.op_type = 2 THEN -- burn
		UPDATE erc1155_tok SET total_supply=(total_supply - NEW.amount)
			WHERE contract_aid=NEW.contract_aid AND token_id=NEW.token_id;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc1155_transf_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	IF OLD.op_type = 1 THEN -- mint
		UPDATE erc1155_tok SET total_supply=(total_supply - OLD.amount)
			WHERE contract_aid=OLD.contract_aid AND token_id=OLD.token_id;
	END IF;
	IF OLD.op_type = 2 THEN -- burn
		UPDATE erc1155_tok SET total_supply=(total_supply + OLD.amount)
			WHERE contract_aid=OLD.contract_aid AND token_id=OLD.token_id;
	END IF;
	DELETE FROM erc1155_bal WHERE parent_id = OLD.id;
	RETURN OLD;

END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc1155_bal_insert() RETURNS trigger AS  $$
DECLARE
	v_cur_balance		DECIMAL;
BEGIN

	INSERT INTO erc1155_holder(contract_aid,token_id,aid)
		VALUES(NEW.contract_aid,NEW.token_id,NEW.aid)
		ON CONFLICT DO NOTHING;
	UPDATE erc1155_holder SET cur_balance = (cur_balance + NEW.amount)
		WHERE contract_aid=NEW.contract_aid AND aid=NEW.aid AND token_id=NEW.token_id
		RETURNING cur_balance INTO v_cur_balance;
	UPDATE erc1155_bal SET balance = v_cur_balance WHERE id=NEW.id;
	IF v_cur_balance = 0 THEN
		UPDATE erc1155_tok SET num_holders = (num_holders - 1)
			WHERE contract_aid=NEW.contract_aid AND token_id = NEW.token_id;
	END IF;
	IF v_cur_balance > 0 THEN -- address 0x0 is excluded here
		UPDATE erc1155_tok SET num_holders = (num_holders + 1)
			WHERE contract_aid=NEW.contract_aid AND token_id=NEW.token_id;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc1155_bal_delete() RETURNS trigger AS  $$
DECLARE
	v_cur_balance DECIMAL;
BEGIN

	UPDATE erc1155_holder SET cur_balance = (cur_balance - OLD.amount)
		WHERE contract_aid=OLD.contract_aid AND aid=OLD.aid AND token_id=OLD.token_id
		RETURNING cur_balance INTO v_cur_balance;
	IF v_cur_balance = 0 THEN
		UPDATE erc1155_tok SET num_holders = (num_holders - 1)
			WHERE contract_aid=OLD.contract_aid AND token_id = OLD.token_id;
	END IF;
	IF v_cur_balance > 0 THEN -- address 0x0 is excluded here
		UPDATE erc1155_tok SET num_holders = (num_holders + 1)
			WHERE contract_aid=OLD.contract_aid AND token_id=OLD.token_id;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc1155_batch_insert() RETURNS trigger AS  $$
DECLARE
	v_aid bigint;
	v_cnt numeric;
	v_token_id_str TEXT;
	v_token_id BIGINT;
	v_amounts TEXT[];
	v_amount DECIMAL;
	v_counter BIGINT;
BEGIN

	v_counter := 1;
	v_amounts := STRING_TO_ARRAY(NEW.amounts,',');
	FOREACH v_token_id_str IN ARRAY STRING_TO_ARRAY(NEW.token_ids,',')
		LOOP
			v_token_id := v_token_id_str::BIGINT;
			v_amount := v_amounts[v_counter]::DECIMAL;
	--		RAISE NOTICE 'v_amounts = %',v_amounts;
	--		RAISE NOTICE 'v_amount = %',v_amount;
			INSERT INTO erc1155_bal(block_num,tx_id,batch_id,contract_aid,token_id,time_stamp,aid,amount)
					VALUES(NEW.block_num,NEW.tx_id,NEW.id,NEW.contract_aid,v_token_id,NEW.time_stamp,NEW.from_aid,-v_amount);

			INSERT INTO erc1155_bal(block_num,tx_id,batch_id,contract_aid,token_id,time_stamp,aid,amount)
					VALUES(NEW.block_num,NEW.tx_id,NEW.id,NEW.contract_aid,v_token_id,NEW.time_stamp,NEW.to_aid,v_amount);
			IF NEW.op_type = 1 THEN -- mint
				UPDATE erc1155_tok SET total_supply=(total_supply + v_amount)
					WHERE contract_aid=NEW.contract_aid AND token_id=v_token_id;
			END IF;
			IF NEW.op_type = 2 THEN -- burn
				UPDATE erc1155_tok SET total_supply=(total_supply - v_amount)
					WHERE contract_aid=NEW.contract_aid AND token_id=v_token_id;
			END IF;
			v_counter := (v_counter + 1);
		END LOOP;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc1155_batch_delete() RETURNS trigger AS  $$
DECLARE
	v_token_id_str TEXT;
	v_token_id BIGINT;
	v_amounts TEXT[];
	v_amount DECIMAL;
	v_counter BIGINT;
BEGIN

	v_counter := 1;
	v_amounts := STRING_TO_ARRAY(OLD.amounts,',');
	FOREACH v_token_id_str IN ARRAY STRING_TO_ARRAY(OLD.token_ids,',')
		LOOP
			v_token_id := v_token_id_str::BIGINT;
			v_amount := v_amounts[v_counter]::DECIMAL;
			IF OLD.op_type = 1 THEN -- mint
				UPDATE erc1155_tok SET total_supply=(total_supply - v_amount)
					WHERE contract_aid=OLD.contract_aid AND token_id=v_token_id;
			END IF;
			IF OLD.op_type = 2 THEN -- burn
				UPDATE erc1155_tok SET total_supply=(total_supply + v_amount)
					WHERE contract_aid=OLD.contract_aid AND token_id=v_token_id;
			END IF;
			v_counter := (v_counter + 1);
		END LOOP;
	DELETE FROM erc1155_bal WHERE batch_id = OLD.id;
	RETURN OLD;

END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION on_erc1155_bal_update() RETURNS trigger AS  $$
DECLARE
BEGIN
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc1155_holder_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_erc1155_holder_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
