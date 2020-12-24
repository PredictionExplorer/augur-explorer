CREATE OR REPLACE FUNCTION on_ens_name_insert_before() RETURNS trigger AS  $$
DECLARE
	v_prev_timestamp timestamptz;
	v_cnt numeric;
BEGIN

	SELECT expires FROM active_name WHERE label = NEW.label INTO v_prev_timestamp;
	UPDATE active_name SET
		ensname_id = NEW.id,
		expires = NEW.expires
		WHERE label=NEW.label;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO active_name(ensname_id,expires,prev_expires,name,label)
		VALUES(NEW.id,NEW.expires,v_prev_timestamp,NEW.name,NEW.label);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_name_insert_after() RETURNS trigger AS  $$
DECLARE
BEGIN

	-- Noote: the process that execute DELETEs in 'active_name' table is running outside

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_name_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE active_name SET
		expires = OLD.prev_expires
		WHERE label=OLD.label;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
