CREATE OR REPLACE FUNCTION on_item_bought_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN

	UPDATE rw_new_offer SET active=FALSE WHERE offer_id=NEW.offer_id;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'Offer % not found',NEW.offer_id;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_item_bought_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
