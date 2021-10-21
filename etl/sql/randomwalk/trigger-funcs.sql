CREATE OR REPLACE FUNCTION on_item_bought_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
	v_rwalk_aid				BIGINT;
BEGIN

	UPDATE rw_new_offer SET active=FALSE WHERE offer_id=NEW.offer_id;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'Offer % not found',NEW.offer_id;
	END IF;

	SELECT rwalk_aid FROM rw_new_offer WHERE offer_id=NEW.offer_id INTO v_rwalk_aid;
	IF v_rwalk_aid IS NULL THEN
		RAISE EXCEPTION 'Offer %v not found when looking up for rwalk_aid',NEW.offer_id;
	END IF;
	UPDATE rw_stats
		SET		total_vol = (total_vol +1),
				total_num_ops = (total_num_ops +1)
		WHERE rwalk_aid = v_rwalk_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO rw_stats(rwalk_aid,total_vol,num_ops)
			VALUES(v_rwalk_aid,NEW.price);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_item_bought_delete() RETURNS trigger AS  $$
DECLARE
	v_rwalk_aid				BIGINT;
BEGIN

	SELECT rwalk_aid FROM rw_new_offer WHERE offer_id=OLD.offer_id INTO v_rwalk_aid;
	IF v_rwalk_aid IS NULL THEN
		RETURN OLD;
	END IF;
	UPDATE rw_stats
		SET		total_vol = (total_vol -1),
				total_num_ops = (total_num_ops -1)
		WHERE rwalk_aid = v_rwalk_aid;
	UPDATE rw_new_offer SET active=TRUE WHERE offer_id=OLD.offer_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
