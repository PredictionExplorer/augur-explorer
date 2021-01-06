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
CREATE OR REPLACE FUNCTION update_root_ens_human_names() RETURNS BIGINT AS  $$
-- updates 'ens_node.fqdn_words' field and returns number of new names found
-- this func must be called recursively until return value returns 0
DECLARE
	v_rec				RECORD;
	v_updated_count		BIGINT;
	v_word				TEXT;
	v_word_stored		TEXT;
	v_word_parent		TEXT;
	v_id				BIGINT;
BEGIN

	v_updated_count := 0;
	FOR v_rec IN (SELECT * FROM ens_node WHERE fqdn_words='')
	LOOP
		SELECT word FROM ens_label WHERE v_rec.label = label INTO v_word;
		IF v_word IS NOT NULL THEN
			-- check root node existence
			v_id:=NULL;
			SELECT id,fqdn_words
				FROM ens_node
				WHERE
					node='0000000000000000000000000000000000000000000000000000000000000000' AND
					label=v_rec.label AND
					fqdn_words = ''
				INTO v_id,v_word_stored;
			IF v_id IS NOT NULL THEN
				IF v_word_stored IS NOT NULL THEN
					-- this is a TLD node (like .eth or .test)
					UPDATE ens_node SET fqdn_words = v_word WHERE id=v_id;
					v_updated_count := v_updated_count + 1;
				END IF;
			END IF;
		END IF;
	END LOOP;
	RETURN v_updated_count;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION update_ens_human_names() RETURNS BIGINT AS  $$
-- updates 'ens_node.fqdn_words' field and returns number of new names found
-- this func must be called recursively until return value returns 0
DECLARE
	v_rec				RECORD;
	v_updated_count		BIGINT;
	v_word				TEXT;
	v_word_parent		TEXT;
	v_id_parent			BIGINT;
BEGIN

	v_updated_count := 0;
	FOR v_rec IN (SELECT * FROM ens_node WHERE fqdn_words='')
	LOOP
		SELECT word FROM ens_label WHERE v_rec.label = label INTO v_word;
		IF v_word IS NOT NULL THEN
			-- check parent root node eixstence
			v_id_parent := NULL;
			SELECT id,fqdn_words FROM ens_node
				WHERE
					fqdn=v_rec.node
				INTO v_id_parent,v_word_parent;
			IF v_id_parent IS NOT NULL THEN
				IF v_word_parent IS NOT NULL THEN
					IF LENGTH(v_word_parent) > 0 THEN
						UPDATE ens_node
							SET fqdn_words = CONCAT(v_word,'.',v_word_parent)
							WHERE id=v_rec.id;
					END IF;
				END IF;
			END IF;
		END IF;
	END LOOP;
	RETURN v_updated_count;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_reg_transfer_insert() RETURNS trigger AS  $$
DECLARE
	v_node_id			BIGINT;
BEGIN

	IF NEW.evtlog_id IS NULL THEN
		NEW.evtlog_id := 0;
	END IF;
	UPDATE ens_node
		SET cur_owner_aid = NEW.aid,
			cur_owner_evt = NEW.evtlog_id
		WHERE fqdn=NEW.node AND (cur_owner_evt >= NEW.evtlog_id);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_reg_transfer_delete() RETURNS trigger AS  $$
DECLARE
	v_cur_owner_aid		BIGINT;
	v_cur_owner_evt		BIGINT;
BEGIN

	IF OLD.evtlog_id IS NULL THEN
		OLD.evtlog_id := 0;
	END IF;
	SELECT evtlog_id,aid FROM ens_reg_transf 
		WHERE node = OLD.node
		ORDER BY id DESC LIMIT 1
		INTO v_cur_owner_evt,v_cur_owner_aid;

	UPDATE ens_node
		SET cur_owner_aid = v_cur_owner_aid,
			cur_owner_evt = v_cur_owner_evt
		WHERE fqdn = OLD.node;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
