CREATE OR REPLACE FUNCTION on_uswap1_insert() RETURNS trigger AS  $$
DECLARE
	v_token0_aid		bigint;
	v_token1_aid		bigint;
	v_wrapper_aid		bigint;
	v_market_aid		bigint;
BEGIN

	SELECT token0_aid,token1_aid
		FROM upair WHERE pair_aid = NEW.pair_aid INTO v_token0_aid,v_token1_aid;
	IF v_token0_aid IS NULL THEN
		RAISE EXCEPTION 'upair record with id % not found',NEW.pair_aid;
	END IF;

	IF NEW.amount0_in > 0 THEN
		INSERT INTO uswap2(evtlog_id,block_num,tx_id,time_stamp,uswap1_id,aid,token_aid,amount)
			VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.time_stamp,NEW.id,NEW.recipient_aid,v_token0_aid,-NEW.amount0_in);
	END IF;
	IF NEW.amount1_in > 0 THEN
		INSERT INTO uswap2(evtlog_id,block_num,tx_id,time_stamp,uswap1_id,aid,token_aid,amount)
			VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.time_stamp,NEW.id,NEW.recipient_aid,v_token1_aid,-NEW.amount1_in);
	END IF;
	IF NEW.amount0_out > 0 THEN
		INSERT INTO uswap2(evtlog_id,block_num,tx_id,time_stamp,uswap1_id,aid,token_aid,amount)
			VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.time_stamp,NEW.id,NEW.recipient_aid,v_token0_aid,NEW.amount0_out);
	END IF;
	IF NEW.amount1_out > 0 THEN
		INSERT INTO uswap2(evtlog_id,block_num,tx_id,time_stamp,uswap1_id,aid,token_aid,amount)
			VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.time_stamp,NEW.id,NEW.recipient_aid,v_token1_aid,NEW.amount1_out);
	END IF;

	UPDATE upair SET total_swaps = (total_swaps + 1) WHERE pair_aid=NEW.pair_aid;

	SELECT wrapper_aid,v_market_aid FROM af_wrapper
		WHERE wrapper_aid IN(v_token0_aid,v_token1_aid) LIMIT 1
		INTO v_wrapper_aid,v_market_aid;
	IF v_wrapper_aid IS NOT NULL THEN
		INSERT INTO agtx(tx_id,block_num,account_aid,market_aid,tx_type)
			VALUES(NEW.tx_id,NEW.block_num,NEW.recipient_aid,v_market_aid,1)
			ON CONFLICT DO NOTHING;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_uswap1_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	DELETE FROM uswap2 WHERE uswap1_id=OLD.id;
	UPDATE upair SET total_swaps = (total_swaps - 1) WHERE pair_aid=OLD.pair_aid;
	DELETE FROM agtx WHERE tx_id=OLD.tx_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_uswap2_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	INSERT INTO uswap_stats(aid,token_aid) VALUES(NEW.aid,NEW.token_aid) 
		ON CONFLICT DO NOTHING;
	UPDATE uswap_stats SET
		num_swaps = (num_swaps + 1),
		volume = (volume + ABS(NEW.amount))
		WHERE aid=NEW.aid AND token_aid=NEW.token_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_uswap2_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE uswap_stats SET
		num_swaps = (num_swaps - 1),
		volume = (volume - ABS(OLD.amount))
		WHERE aid=OLD.aid AND token_aid=OLD.token_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
