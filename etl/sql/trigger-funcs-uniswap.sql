CREATE OR REPLACE FUNCTION on_uswap1_insert() RETURNS trigger AS  $$
DECLARE
	v_token0_aid		bigint;
	v_token1_aid		bigint;
BEGIN

	SELECT token0_aid,token1_aid
		FROM upair WHERE pair_aid = NEW.pair_aid INTO v_token0_aid,v_token1_aid;
	IF v_token0_aid IS NULL THEN
		RAISE EXCEPTION 'upair record with id % not found',NEW.pair_aid;
	END IF;

	IF NEW.amount0_in > 0 THEN
		INSERT INTO uswap2(evtlog_id,block_num,tx_id,recipient_aid,time_stamp,uswap1_id,aid,token_aid,amount)
			VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.time_stamp,NEW.id,NEW.aid,v_token0_aid,-NEW.amount0_in);
	END IF;
	IF NEW.amount1_in > 0 THEN
		INSERT INTO uswap2(evtlog_id,block_num,tx_id,recipient_aid,time_stamp,uswap1_id,aid,token_aid,amount)
			VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.time_stamp,NEW.id,NEW.aid,v_token1_aid,-NEW.amount1_in);
	END IF;
	IF NEW.amount0_out > 0 THEN
		INSERT INTO uswap2(evtlog_id,block_num,tx_id,recipient_aid,time_stamp,uswap1_id,aid,token_aid,amount)
			VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.time_stamp,NEW.id,NEW.aid,v_token0_aid,NEW.amount0_out);
	END IF;
	IF NEW.amount0_out > 0 THEN
		INSERT INTO uswap2(evtlog_id,block_num,tx_id,recipient_aid,time_stamp,uswap1_id,aid,token_aid,amount)
			VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.time_stamp,NEW.id,NEW.aid,v_token1_aid,NEW.amount1_out);
	END IF;

	UPDATE upair SET total_swaps = (total_swaps + 1) WHERE pair_aid=NEW.pair_aid;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_uswap1_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	DELETE FROM uswap2 WHERE uswap1_id=OLD.id;
	UPDATE upair SET total_swaps = (total_swaps - 1) WHERE pair_aid=NEW.pair_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
