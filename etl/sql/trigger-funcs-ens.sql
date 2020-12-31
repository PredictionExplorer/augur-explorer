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
CREATE OR REPLACE FUNCTION on_ens_new_owner_insert() RETURNS trigger AS  $$
DECLARE
	v_prev_timestamp timestamptz;
	v_cnt numeric;
BEGIN

	INSERT INTO ens_node(evtlog_id,block_num,tx_id,time_stamp,label,node,fqdn)
		VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.time_stamp,NEW.label,NEW.node,NEW.fqdn)
		ON CONFLICT DO NOTHING;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_email_tokens_insert() RETURNS trigger AS  $$
DECLARE
	v_rec	RECORD;
BEGIN
	-- we use the trigger to prevent 'duplicate index' error by returning NULL on duplicates
	SELECT * FROM email_tokens WHERE hash=NEW.hash INTO v_rec;
	IF v_rec IS NOT NULL THEN
		RETURN NULL;
	END IF;
	BEGIN 
		INSERT INTO email_tokens(token,hash) VALUES(NEW.token,NEW.hash)
			ON CONFLICT DO NOTHING;
	EXCEPTION
		WHEN OTHERS THEN
			NULL;
	END;
	RETURN NULL;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_pwd_db_insert() RETURNS trigger AS  $$
DECLARE
	v_rec	RECORD;
BEGIN
	SELECT * FROM pwd_db WHERE hash=NEW.hash INTO v_rec;
	IF v_rec IS NOT NULL THEN
		RETURN NULL;
	END IF;
	BEGIN 
		INSERT INTO pwd_db(password,hash) VALUES(NEW.password,NEW.hash)
			ON CONFLICT DO NOTHING;
	EXCEPTION
		WHEN OTHERS THEN
			RETURN NULL;
	END;
	RETURN NULL;
END;
$$ LANGUAGE plpgsql;
