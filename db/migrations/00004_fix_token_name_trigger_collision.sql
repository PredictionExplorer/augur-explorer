-- +goose Up
-- Migrations 00002 (cosmicgame) and 00003 (randomwalk) both defined trigger
-- functions named on_token_name_insert()/on_token_name_delete(). Because
-- 00003 runs after 00002, its CREATE OR REPLACE silently overwrote the
-- CosmicGame bodies, so the cg_token_name_insert trigger executed the
-- RandomWalk body and every INSERT INTO cg_token_name failed with
--   ERROR: record "new" has no field "new_name"
-- (found by the API parity suite's fixture seeding). Give each project its
-- own unambiguous function pair and rebind the triggers.

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_cg_token_name_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_mint_event SET token_name = NEW.token_name WHERE token_id=NEW.token_id;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_cg_token_name_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE cg_mint_event SET token_name = '' WHERE token_id = OLD.token_id;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_rw_token_name_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN

	UPDATE rw_token SET last_name=NEW.new_name
		WHERE token_id=NEW.token_id AND rwalk_aid=NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'Token ID % not found',NEW.token_id;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_rw_token_name_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

DROP TRIGGER cg_token_name_insert ON cg_token_name;
DROP TRIGGER cg_token_name_delete ON cg_token_name;
CREATE TRIGGER cg_token_name_insert AFTER INSERT ON cg_token_name FOR EACH ROW EXECUTE PROCEDURE on_cg_token_name_insert();
CREATE TRIGGER cg_token_name_delete AFTER DELETE ON cg_token_name FOR EACH ROW EXECUTE PROCEDURE on_cg_token_name_delete();

DROP TRIGGER rw_token_name_insert ON rw_token_name;
DROP TRIGGER rw_token_name_delete ON rw_token_name;
CREATE TRIGGER rw_token_name_insert AFTER INSERT ON rw_token_name FOR EACH ROW EXECUTE PROCEDURE on_rw_token_name_insert();
CREATE TRIGGER rw_token_name_delete AFTER DELETE ON rw_token_name FOR EACH ROW EXECUTE PROCEDURE on_rw_token_name_delete();

-- Remove the ambiguous shared-name functions entirely.
DROP FUNCTION on_token_name_insert();
DROP FUNCTION on_token_name_delete();

-- +goose Down
-- Restore the (buggy) shared-name binding exactly as 00003 left it.
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_token_name_insert() RETURNS trigger AS  $$
DECLARE
	v_cnt                   NUMERIC;
BEGIN

	UPDATE rw_token SET last_name=NEW.new_name
		WHERE token_id=NEW.token_id AND rwalk_aid=NEW.contract_aid;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	IF v_cnt = 0 THEN
		RAISE EXCEPTION 'Token ID % not found',NEW.token_id;
	END IF;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION on_token_name_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

DROP TRIGGER cg_token_name_insert ON cg_token_name;
DROP TRIGGER cg_token_name_delete ON cg_token_name;
CREATE TRIGGER cg_token_name_insert AFTER INSERT ON cg_token_name FOR EACH ROW EXECUTE PROCEDURE on_token_name_insert();
CREATE TRIGGER cg_token_name_delete AFTER DELETE ON cg_token_name FOR EACH ROW EXECUTE PROCEDURE on_token_name_delete();

DROP TRIGGER rw_token_name_insert ON rw_token_name;
DROP TRIGGER rw_token_name_delete ON rw_token_name;
CREATE TRIGGER rw_token_name_insert AFTER INSERT ON rw_token_name FOR EACH ROW EXECUTE PROCEDURE on_token_name_insert();
CREATE TRIGGER rw_token_name_delete AFTER DELETE ON rw_token_name FOR EACH ROW EXECUTE PROCEDURE on_token_name_delete();

DROP FUNCTION on_rw_token_name_insert();
DROP FUNCTION on_rw_token_name_delete();
DROP FUNCTION on_cg_token_name_insert();
DROP FUNCTION on_cg_token_name_delete();
