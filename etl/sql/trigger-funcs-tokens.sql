CREATE OR REPLACE FUNCTION on_dai_transf_insert() RETURNS trigger AS  $$
DECLARE
	v_aid bigint;
	v_cnt numeric;
	v_augur bool;	-- true if this transfer is made to Augur Wallet account
	v_internal bool;
BEGIN

	v_augur := false;
	SELECT aid FROM ustats WHERE aid = NEW.from_aid INTO v_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt > 0 THEN
		v_augur := true;
	END IF;
	INSERT INTO dai_bal(block_num,tx_id,dai_transf_id,aid,amount,augur,internal)
			VALUES(NEW.block_num,NEW.tx_id,NEW.id,NEW.from_aid,-NEW.amount,v_augur,NEW.from_internal);


	v_augur := false;
	SELECT aid FROM ustats WHERE aid = NEW.to_aid INTO v_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt > 0 THEN
		v_augur := true;
	END IF;
	INSERT INTO dai_bal(block_num,tx_id,dai_transf_id,aid,amount,augur,internal)
			VALUES(NEW.block_num,NEW.tx_id,NEW.id,NEW.to_aid,NEW.amount,v_augur,NEW.to_internal);

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_dai_transf_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	DELETE FROM dai_bal WHERE dai_transf_id = OLD.id;
	RETURN OLD;

END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_dai_bal_insert() RETURNS trigger AS  $$
DECLARE
	v_aid bigint;
	v_cnt numeric;
	v_augur bool;
BEGIN

	v_augur := false;
	SELECT aid FROM ustats WHERE aid = NEW.aid INTO v_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt > 0 THEN
		v_augur := true;
	END IF;

	IF v_augur THEN
		if NEW.internal IS FALSE THEN
			UPDATE block AS b
				SET cash_flow = (cash_flow + NEW.amount)
				WHERE	b.block_num = NEW.block_num;
		END IF;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_dai_bal_delete() RETURNS trigger AS  $$
DECLARE
	v_augur bool;	-- true if this transfer is made to Augur Wallet account
	v_aid bigint;
	v_cnt numeric;
BEGIN

	v_augur := false;
	SELECT aid FROM ustats WHERE aid = OLD.aid INTO v_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt > 0 THEN
		v_augur := true;
	END IF;
	IF v_augur THEN
		IF OLD.internal IS FALSE THEN
			UPDATE block AS b
				SET cash_flow = (cash_flow - OLD.amount)
				WHERE	b.block_num = OLD.block_num;
		END IF;
	END IF;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_dai_bal_update() RETURNS trigger AS  $$
DECLARE
	v_augur bool;	-- true if this transfer is made to Augur Wallet account
	v_aid bigint;
	v_cnt numeric;
BEGIN
	-- Noute: this trigger only calculates block.cash_flow. For another process
	--			use another trigger function
	-- cash_flow calculation starts
	IF NEW.internal != OLD.internal THEN
		RAISE EXCEPTION 'Changing dai_bal.internal field not possible. Delete the whole block';
	END IF;
	IF OLD.augur != NEW.augur THEN -- this update is coming from ustats table
		IF NEW.augur THEN
			IF NEW.internal IS FALSE THEN
				UPDATE block AS b
					SET cash_flow = (cash_flow + NEW.amount)
					WHERE	b.block_num = NEW.block_num;
			END IF;
		END IF;
	END IF;
	-- cash flow calculation ends
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_af_wrapper_insert() RETURNS trigger AS  $$
DECLARE
	v_rec record;
BEGIN

	-- there might be Balancer swaps and Uniswap swaps of this token
	-- since insert into 'agtx' table has pure statistical purpose,
	-- we can insert as many times as we want, dupliates are prevented with
	-- unique index key
	-- this function will have effect only if token tables are dropped and
	--	then rebuilt again. This is not the normal exeuction flow

	-- Balancer
	FOR v_rec IN 
		SELECT * FROM bswap WHERE token_in_aid=v_rec.wrapper_aid OR token_out_aid=v_rec.wrapper_aid
	LOOP
		PERFORM insert_agtx_event(
			v_rec.tx_id,v_rec.evtlog_id,v_rec.block_num,v_rec.caller_aid,NEW.market_aid,2,1
		);
	END LOOP;

	-- Uniswap
	FOR v_rec IN 
		SELECT * FROM uswap2 WHERE token_aid=v_rec.wrapper_aid
	LOOP
		PERFORM insert_agtx_event(
			v_rec.tx_id,v_rec.evtlog_id,v_rec.block_num,v_rec.recipient_aid,NEW.market_aid,1,1
		);
	END LOOP;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_af_wrapper_delete() RETURNS trigger AS  $$
DECLARE
	v_rec record;
BEGIN

	FOR v_rec IN 
		SELECT * FROM bswap WHERE token_in_aid=v_rec.wrapper_aid OR token_out_aid=v_rec.wrapper_aid
	LOOP
		PERFORM delete_agtx_event(v_rec.tx_id,v_rec.evtlog_id);
	END LOOP;

	FOR v_rec IN 
		SELECT * FROM uswap2 WHERE token_aid=v_rec.wrapper_aid
	LOOP
		PERFORM delete_agtx_event(v_rec.tx_id,v_rec.evtlog_id);
	END LOOP;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
