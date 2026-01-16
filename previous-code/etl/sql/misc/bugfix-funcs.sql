-- Stored procedures to fix data errors
CREATE OR REPLACE FUNCTION fix_defi_stats() RETURNS void AS  $$
-- fixes values in `defi_stats` (error happens when you forget to truncate
--		the table at initialization)
DECLARE
	v_rec RECORD;
	v_num_recs BIGINT;
	v_count_bad_uniswap BIGINT;
	v_count_bad_balancer BIGINT;
	v_count_good_uniswap BIGINT;
	v_count_good_balancer BIGINT;
BEGIN
	
	v_count_bad_uniswap := 0;
	v_count_bad_balancer := 0;
	v_count_good_uniswap := 0;
	v_count_good_balancer := 0;
	FOR v_rec IN (SELECT * FROM defi_stats)
	LOOP
		SELECT count(*) AS total FROM agtx_evt WHERE account_aid=v_rec.aid AND defi_platform = 1
			INTO v_num_recs;
		IF v_num_recs IS NOT NULL THEN
			IF v_num_recs = v_rec.uniswap_swaps THEN
				v_count_good_uniswap := v_count_good_uniswap + 1 ;
			ELSE
				v_count_bad_uniswap := v_count_bad_uniswap + 1 ;
				RAISE NOTICE 'UNISWAP: Account % stored value is % swaps but should have %v swaps',v_rec.aid,v_rec.uniswap_swaps,v_num_recs;
				UPDATE defi_stats SET uniswap_swaps = v_num_recs WHERE aid=v_rec.aid;
			END IF;
		END IF;
		SELECT count(*) AS total FROM agtx_evt WHERE account_aid=v_rec.aid AND defi_platform = 1
			INTO v_num_recs;
		IF v_num_recs IS NOT NULL THEN
			IF v_num_recs = v_rec.balancer_swaps THEN
				v_count_good_balancer := v_count_good_balancer + 1 ;
			ELSE
				v_count_bad_balancer := v_count_bad_balancer + 1 ;
				RAISE NOTICE 'BALANCER: Account % stored value is % swaps but should have %v swaps',v_rec.aid,v_rec.balancer_swaps,v_num_recs;
				UPDATE defi_stats SET balancer_swaps = v_num_recs WHERE aid=v_rec.aid;
			END IF;
		END IF;
	END LOOP;
	RAISE NOTICE 'Summary report:';
	RAISE NOTICE 'Uniswap swaps: good records: % , bad records: %',v_count_good_uniswap,v_count_bad_uniswap;
	RAISE NOTICE 'Balancer swaps: good records: % , bad records: %',v_count_good_balancer,v_count_bad_balancer ;
END;
$$ LANGUAGE plpgsql;
