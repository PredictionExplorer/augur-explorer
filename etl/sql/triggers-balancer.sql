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
CREATE OR REPLACE FUNCTION on_bjoin_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE bholder SET amount = amount + OLD.amount
		WHERE pool_aid=OLD.pool_aid AND holder_aid=OLD.caller_aid;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
