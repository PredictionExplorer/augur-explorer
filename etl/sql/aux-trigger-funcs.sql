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
CREATE OR REPLACE FUNCTION dai_balances_from_ustats()
RETURNS TABLE(f_aid bigint,f_balance DECIMAL) AS $$
DECLARE
	v_stat_rec record;
BEGIN

	FOR v_stat_rec IN 
		SELECT * FROM ustats
	LOOP
		SELECT aid,balance FROM dai_bal AS db WHERE db.aid = v_stat_rec.wallet_aid
				ORDER BY db.block_num DESC,db.id DESC LIMIT 1 INTO f_aid,f_balance;
		RETURN NEXT;
	END LOOP;
END;
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION test_insert()
RETURNS void AS $$
DECLARE
	v_cnt numeric;
	v_field integer;
BEGIN

	INSERT INTO test(field) VALUES(1) ON CONFLICT DO NOTHING RETURNING field  INTO v_field ;
	IF v_field IS NULL THEN
		RAISE NOTICE 'field is null';
	END IF;
	GET DIAGNOSTICS v_cnt = ROW_COUNT;
	RAISE NOTICE 'row count is % v_id=%',v_cnt,v_field;

END;
$$ LANGUAGE plpgsql;
