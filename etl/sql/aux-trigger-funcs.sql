CREATE OR REPLACE FUNCTION check_chain() RETURNS text AS  $$
DECLARE
	v_block_num bigint;
	v_first_block bigint;
	v_cnt numeric;
	v_rec record;
BEGIN

	SELECT block_num FROM block ORDER BY block_num LIMIT 1 INTO v_first_block;
	FOR v_rec IN (	SELECT block_num,parent_hash,block_hash FROM block order by block_num) LOOP
		SELECT block_num FROM block WHERE block_hash=v_rec.parent_hash INTO v_block_num;
		GET DIAGNOSTICS v_cnt = ROW_COUNT;
		IF v_cnt = 0 THEN
			IF v_first_block != v_rec.block_num THEN
				RAISE EXCEPTION 'No parent block with hash % for block num % (first block = %)',v_rec.parent_hash,v_rec.block_num,v_first_block;
			END IF;
		END IF;
	END LOOP;
	RETURN 'Chain has correct parent block for all blocks';
END;
$$ LANGUAGE plpgsql;

