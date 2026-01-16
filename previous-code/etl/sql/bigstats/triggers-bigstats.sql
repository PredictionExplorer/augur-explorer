CREATE OR REPLACE FUNCTION on_bs_log_insert() RETURNS trigger AS  $$
DECLARE
	ts		BIGINT;
BEGIN

	ts := EXTRACT(EPOCH FROM NEW.time_stamp);
	ts := FLOOR(ts/86400);
	ts := ts * 8400;

	UPDATE bs_period SET
		

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_bs_log_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
