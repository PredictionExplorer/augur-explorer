CREATE OR REPLACE FUNCTION on_swf_hist_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_swf_hist_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	-- if a record at some block-height is deleted, subsequent records become invalid
	DELETE from swf_hist WHERE
		(
			block_num > OLD.block_num AND
		) OR (
			block_num = OLD.block_num AND
			tx_index > OLD.tx_index
		) OR (
			block_num = OLD.block_num AND
			tx_index = OLD.tx_index AND
			log_index > OLD.log_index
		)
	;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_swf_hist_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_swf_hist_delete() RETURNS trigger AS  $$
DECLARE
BEGIN
	-- if a record at some block-height is deleted, subsequent records become invalid
	DELETE from swf_hist WHERE
		(
			block_num > OLD.block_num AND
		) OR (
			block_num = OLD.block_num AND
			tx_index > OLD.tx_index
		) OR (
			block_num = OLD.block_num AND
			tx_index = OLD.tx_index AND
			log_index > OLD.log_index
		)
	;
	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_poolhist_insert() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE pool_hist SET
		total_fees = (total_fees + NEW.swap_fee),
		total_swasp = (total_swaps + 1)
	WHERE pool_id = NEW.pool_id;

	RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION on_poolhist_delete() RETURNS trigger AS  $$
DECLARE
BEGIN

	UPDATE pool_hist SET
		total_fees = (total_fees - OLD.swap_fee),
		total_swasp = (total_swaps - 1)
	WHERE pool_id = OLD.pool_id;

	RETURN OLD;
END;
$$ LANGUAGE plpgsql;
