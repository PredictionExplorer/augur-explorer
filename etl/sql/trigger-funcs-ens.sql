CREATE OR REPLACE FUNCTION on_ens_name_reg1_insert_before() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	IF NEW.fqdn='' THEN
		RAISE EXCEPTION 'Attempt to INSERT ens_name with empty fqdn in ens_reg_name1';
	END IF;
	UPDATE active_name SET
		ensname_id = NEW.id,
		expires = NEW.expires
		WHERE fqdn=NEW.fqdn;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO ens_name(owner_aid,expires,label,node,fqdn,name,cost)
			VALUES(NEW.owner_aid,NEW.expires,NEW.label,NEW.node,NEW.fqdn,NEW.name,NEW.cost);
		INSERT INTO active_name(ensname_id,expires,name,label,node,fqdn)
			VALUES(NEW.id,NEW.expires,NEW.name,NEW.label,NEW.node,NEW.fqdn);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_name_reg2_insert_before() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	IF NEW.fqdn='' THEN
		RAISE EXCEPTION 'Attempt to INSERT ens_name with empty fqdn in ens_name_reg2';
	END IF;
	UPDATE active_name SET
		ensname_id = NEW.id,
		expires = NEW.expires
		WHERE fqdn=NEW.fqdn;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO ens_name(owner_aid,expires,label,node,fqdn)-- event v2 doesn't have cost/name fields
			VALUES(NEW.owner_aid,NEW.expires,NEW.label,NEW.node,NEW.fqdn);
		INSERT INTO active_name(ensname_id,expires,label,node,fqdn)
			VALUES(NEW.id,NEW.expires,NEW.label,NEW.node,NEW.fqdn);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_name_renewed_insert_before() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	IF NEW.fqdn='' THEN
		RAISE EXCEPTION 'Attempt to INSERT ens_name with empty fqdn in ens_name_renewed';
	END IF;
	UPDATE active_name SET
		ensname_id = NEW.id,
		expires = NEW.expires
		WHERE fqdn=NEW.fqdn;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO ens_name(owner_aid,expires,label,node,fqdn,name,cost)
			VALUES(NEW.owner_aid,NEW.expires,NEW.label,NEW.node,NEW.fqdn,NEW.name,NEW.cost);
		INSERT INTO active_name(ensname_id,expires,name,label,node,fqdn)
			VALUES(NEW.id,NEW.expires,NEW.name,NEW.label,NEW.node,NEW.fqdn);
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_new_owner_insert() RETURNS trigger AS  $$
DECLARE
	v_prev_timestamp timestamptz;
	v_cnt numeric;
	v_zero_aid BIGINT;
BEGIN

	INSERT INTO ens_node(evtlog_id,block_num,tx_id,contract_aid,time_stamp,label,node,fqdn)
		VALUES(NEW.evtlog_id,NEW.block_num,NEW.tx_id,NEW.contract_aid,NEW.time_stamp,NEW.label,NEW.node,NEW.fqdn)
		ON CONFLICT DO NOTHING;
	-- fixes active_name records that do not have label/node set
	UPDATE active_name SET	label = NEW.label,node = NEW.node WHERE fqdn=NEW.fqdn AND label IS NULL;
	INSERT INTO name_ownership(tx_hash,owner_aid,fqdn)
		VALUES(NEW.tx_hash,NEW.owner_aid,NEW.fqdn) ON CONFLICT DO NOTHING;
	SELECT address_id FROM address WHERE addr='0x0000000000000000000000000000000000000000' INTO v_zero_aid;
	IF NEW.owner_aid = v_zero_aid  THEN
		UPDATE ens_name SET inactive=TRUE WHERE fqdn=NEW.fqdn;
		UPDATE ens_name SET unregistered = TRUE WHERE fqdn = NEW.fqdn;
	ELSE
		UPDATE ens_name SET unregistered = FALSE WHERE fqdn = NEW.fqdn;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_addr_changed1_insert() RETURNS trigger AS  $$
DECLARE
	v_zero_aid BIGINT;
BEGIN

	INSERT INTO name_address(fqdn,aid,coin_type) VALUES(NEW.fqdn,NEW.aid,60) 
		ON CONFLICT DO NOTHING;
	SELECT address_id FROM address WHERE addr='0x0000000000000000000000000000000000000000' INTO v_zero_aid;
	IF NEW.aid = v_zero_aid  THEN
		UPDATE ens_name SET inactive=TRUE WHERE fqdn=NEW.fqdn;
		UPDATE ens_name SET no_address = TRUE WHERE fqdn=NEW.fqdn;
	ELSE
		UPDATE ens_name SET no_address = FALSE WHERE fqdn=NEW.fqdn;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_addr_changed2_insert() RETURNS trigger AS  $$
DECLARE
	v_zero_aid BIGINT;
BEGIN

	INSERT INTO name_address(fqdn,aid,coin_type) VALUES(NEW.fqdn,NEW.aid,NEW.coin_type)
		ON CONFLICT DO NOTHING;
	SELECT address_id FROM address WHERE addr='0x0000000000000000000000000000000000000000' INTO v_zero_aid;
	IF NEW.aid = v_zero_aid  THEN
		UPDATE ens_name SET inactive=TRUE WHERE fqdn=NEW.fqdn;
		UPDATE ens_name SET no_address = TRUE WHERE fqdn=NEW.fqdn;
	ELSE
		UPDATE ens_name SET no_address = FALSE WHERE fqdn=NEW.fqdn;
	END IF;
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
						v_updated_count := v_updated_count + 1;
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
	v_zero_aid			BIGINT;
BEGIN

	IF NEW.evtlog_id IS NULL THEN
		NEW.evtlog_id := 0;
	END IF;
	UPDATE ens_node
		SET cur_owner_aid = NEW.aid,
			cur_owner_evt = NEW.evtlog_id
		WHERE fqdn=NEW.node AND (cur_owner_evt >= NEW.evtlog_id);
	INSERT INTO name_ownership(tx_hash,owner_aid,fqdn)
		VALUES(NEW.tx_hash,NEW.aid,NEW.node) ON CONFLICT DO NOTHING;

	SELECT address_id FROM address WHERE addr='0x0000000000000000000000000000000000000000' INTO v_zero_aid;
	IF NEW.aid = v_zero_aid  THEN
		UPDATE ens_name SET inactive=TRUE WHERE fqdn=NEW.node;
		UPDATE ens_name SET unregistered=TRUE WHERE fqdn=NEW.node;
	ELSE
		UPDATE ens_name SET unregistered=FALSE WHERE fqdn=NEW.node;
	END IF;

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
CREATE OR REPLACE FUNCTION on_ens_text_chg_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt numeric;
BEGIN

	UPDATE ens_text_key SET value = NEW.value
		WHERE node = NEW.node AND key = NEW.key;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		INSERT INTO ens_text_key(node,key,value)
			VALUES(NEW.node,NEW.key,NEW.value);
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_text_chg_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	-- we do not resotre keys on delete, too complicated for such a very rare event
	--	i.e. User sends a transaction , then fork occurs on the chain which causes block
	--	deletion, then the User	cancels the transaction by sending another (empty) transaction
	--	with higher Gas, and this must happen before his transaction is included in the block
	--	by the miner. This table doesn't have event_ids or block numbers, so it can be updated
	--	on any block
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_text_key_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt NUMERIC;
	v_val TEXT;
BEGIN
	UPDATE ens_text SET
			all_keys = JSONB_SET(all_keys, ARRAY[NEW.key::TEXT], ('"'||NEW.value||'"')::JSONB)
		WHERE node = NEW.node;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		--v_val := CONCAT('{"',NEW.key,'":"',NEW.value,'"}');
		v_val := CONCAT('{"',NEW.key,'":"',NEW.value,'"}');
		INSERT INTO ens_text(node,all_keys)
			VALUES(NEW.node,v_val::JSONB);
	END IF;
	UPDATE ens_text SET num_keys = (num_keys + 1) WHERE node=NEW.node;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_text_key_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION ens_name_list_by_addr(aid BIGINT) RETURNS text AS  $$
-- Returns all active names for an address (resolution from resolver)
DECLARE
BEGIN


	RETURN '';
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION ens_name_list_by_ownership(aid BIGINT) RETURNS text AS  $$
-- Returns all active names for an address (resolution from resolver)
DECLARE
BEGIN

	RETURN '';
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_new_resolver_insert() RETURNS trigger AS  $$
DECLARE
	v_zero_aid BIGINT;
BEGIN

	IF NEW.evtlog_id IS NULL THEN
		NEW.evtlog_id := 0;
	END IF;
	SELECT address_id FROM address WHERE addr='0x0000000000000000000000000000000000000000' INTO v_zero_aid;
	IF NEW.name_aid = v_zero_aid THEN -- when changing Resolvers, the name address got deleted
		UPDATE ens_name SET inactive=TRUE WHERE fqdn=NEW.node;
		UPDATE ens_name SET no_resolver = TRUE WHERE fqdn=NEW.node;
	ELSE
		UPDATE ens_name SET no_resolver = FALSE WHERE fqdn=NEW.node;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_ens_rstr_transf_insert() RETURNS trigger AS  $$
DECLARE
	v_zero_aid BIGINT;
BEGIN

	IF NEW.evtlog_id IS NULL THEN
		NEW.evtlog_id := 0;
	END IF;
	SELECT address_id FROM address WHERE addr='0x0000000000000000000000000000000000000000' INTO v_zero_aid;
	IF NEW.to_aid = v_zero_aid THEN -- when changing Resolvers, the name address got deleted
		UPDATE ens_name SET inactive=TRUE WHERE fqdn=NEW.fqdn;
		UPDATE ens_name SET unregistered = TRUE WHERE fqdn=NEW.fqdn;
	ELSE
		UPDATE ens_name SET unregistered = FALSE WHERE fqdn=NEW.fqdn;
	END IF;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

