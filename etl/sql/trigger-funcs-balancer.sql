CREATE OR REPLACE FUNCTION on_bjoin_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	INSERT INTO bholder(pool_aid,holder_aid)
		VALUES(NEW.pool_aid,NEW.caller_aid);
	UPDATE bholder SET amount = amount + NEW.amount
		WHERE pool_aid=NEW.pool_aid AND holder_aid=NEW.caller_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bjoin_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bholder SET amount = amount - OLD.amount
		WHERE pool_aid=OLD.pool_aid AND holder_aid=OLD.caller_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bexit_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	INSERT INTO bholder(pool_aid,holder_aid)
		VALUES(NEW.pool_aid,NEW.caller_aid);
	UPDATE bholder SET amount = amount - NEW.amount
		WHERE pool_aid=NEW.pool_aid AND holder_aid=NEW.caller_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bexit_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bholder SET amount = amount + OLD.amount
		WHERE pool_aid=OLD.pool_aid AND holder_aid=OLD.caller_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bswap_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE bpool SET num_swaps = num_swaps + 1
		WHERE pool_aid=NEW.pool_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bswap_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bpool SET num_swaps = num_swaps -1
		WHERE pool_aid=OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bholder_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE bpool SET num_holders = num_holders + 1
		WHERE pool_aid=NEW.pool_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bswap_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bpool SET num_holders = num_holders - 1
		WHERE pool_aid=OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bbind_insert() RETURNS trigger AS  $$
DECLARE
BEGIN
	UPDATE bpool SET num_tokens = num_tokens + 1
		WHERE pool_aid=NEW.pool_aid;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bbind_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bpool SET num_tokens = num_tokens - 1
		WHERE pool_aid=OLD.pool_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
