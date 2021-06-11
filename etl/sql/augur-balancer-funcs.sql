CREATE OR REPLACE FUNCTION on_bswap_insert_augur() RETURNS trigger AS  $$
DECLARE
	v_wrapper_aid bigint;
	v_market_aid bigint;
	v_cnt numeric;
BEGIN
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
CREATE OR REPLACE FUNCTION on_bswap_delete_augur() RETURNS trigger AS  $$
DECLARE
BEGIN

	PERFORM delete_agtx_event(OLD.tx_id,OLD.evtlog_id);
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
