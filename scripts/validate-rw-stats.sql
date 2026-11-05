\set ON_ERROR_STOP on
\pset pager off
\echo ''
\echo '=== RandomWalk DB validator (marketplace statistics, section G only) ==='
\echo ''

BEGIN;
SET LOCAL client_min_messages = warning;

CREATE TEMP TABLE vr (
	id         BIGSERIAL PRIMARY KEY,
	section    TEXT,
	check_name TEXT,
	severity   TEXT,    -- ERROR | WARN
	violations BIGINT,
	details    TEXT
) ON COMMIT DROP;



-- ===========================================================================
-- SECTION G: RandomWalk marketplace statistics
-- ===========================================================================

-- Purchases resolved to their offer (price / token / offer type).
CREATE TEMP TABLE rwb ON COMMIT DROP AS
SELECT b.id, b.evtlog_id, b.contract_aid, b.offer_id, b.buyer_aid, b.seller_aid,
       o.rwalk_aid, o.token_id, o.price, o.otype
FROM rw_item_bought b
JOIN rw_new_offer o ON o.contract_aid=b.contract_aid AND o.offer_id=b.offer_id;

-- G1: offer active flag = no purchase and no cancellation for the offer.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. RandomWalk market','rw_new_offer.active = offer neither bought nor canceled','ERROR',COUNT(*),
       LEFT(STRING_AGG('offer '||v.offer_id,',' ORDER BY v.offer_id),300)
FROM (
	SELECT o.offer_id FROM rw_new_offer o
	WHERE o.active IS DISTINCT FROM NOT (
	      EXISTS (SELECT 1 FROM rw_item_bought b WHERE b.contract_aid=o.contract_aid AND b.offer_id=o.offer_id)
	   OR EXISTS (SELECT 1 FROM rw_offer_canceled c WHERE c.contract_aid=o.contract_aid AND c.offer_id=o.offer_id))
) v;

-- G2: rw_stats per RandomWalk token contract.
-- (total_withdrawals has no writer - not checked.)
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. RandomWalk market','rw_stats total_num_toks/money_accumulated/total_vol/total_num_trades','ERROR',COUNT(*),
       LEFT(STRING_AGG(
	'rwalk_aid='||v.rwalk_aid
	|| CASE WHEN v.s_tok <> v.e_tok THEN ' toks '||v.s_tok||'/'||v.e_tok ELSE '' END
	|| CASE WHEN v.s_mon <> v.e_mon THEN ' money '||v.s_mon||'/'||v.e_mon ELSE '' END
	|| CASE WHEN v.s_trd <> v.e_trd THEN ' trades '||v.s_trd||'/'||v.e_trd ELSE '' END
	|| CASE WHEN v.s_vol <> v.e_vol THEN ' vol '||v.s_vol||'/'||v.e_vol ELSE '' END
	,'; ' ORDER BY v.rwalk_aid),1000)
FROM (
	WITH mint AS (SELECT contract_aid rwalk_aid, COUNT(*) c, COALESCE(SUM(price),0) p
	              FROM rw_mint_evt GROUP BY 1),
	trade AS (SELECT rwalk_aid, COUNT(*) c, COALESCE(SUM(price),0) p FROM rwb GROUP BY 1)
	SELECT COALESCE(s.rwalk_aid, COALESCE(m.rwalk_aid, t.rwalk_aid)) AS rwalk_aid,
	       COALESCE(s.total_num_toks,0)    s_tok, COALESCE(m.c,0) e_tok,
	       COALESCE(s.money_accumulated,0) s_mon, COALESCE(m.p,0) e_mon,
	       COALESCE(s.total_num_trades,0)  s_trd, COALESCE(t.c,0) e_trd,
	       COALESCE(s.total_vol,0)         s_vol, COALESCE(t.p,0) e_vol
	FROM rw_stats s
	FULL JOIN mint  m ON m.rwalk_aid=s.rwalk_aid
	FULL JOIN trade t ON t.rwalk_aid=COALESCE(s.rwalk_aid, m.rwalk_aid)
	WHERE COALESCE(s.total_num_toks,0)    <> COALESCE(m.c,0)
	   OR COALESCE(s.money_accumulated,0) <> COALESCE(m.p,0)
	   OR COALESCE(s.total_num_trades,0)  <> COALESCE(t.c,0)
	   OR COALESCE(s.total_vol,0)         <> COALESCE(t.p,0)
) v;

-- G3: rw_mkt_stats per marketplace contract: volume/trades plus open-order
-- counters (offer +1, purchase/cancel -1 per the offer's type).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. RandomWalk market','rw_mkt_stats total_vol/total_num_trades/open sell+buy orders','ERROR',COUNT(*),
       LEFT(STRING_AGG(
	'contract_aid='||v.contract_aid
	|| CASE WHEN v.s_trd  <> v.e_trd  THEN ' trades '||v.s_trd||'/'||v.e_trd ELSE '' END
	|| CASE WHEN v.s_vol  <> v.e_vol  THEN ' vol '||v.s_vol||'/'||v.e_vol ELSE '' END
	|| CASE WHEN v.s_sell <> v.e_sell THEN ' sell_orders '||v.s_sell||'/'||v.e_sell ELSE '' END
	|| CASE WHEN v.s_buy  <> v.e_buy  THEN ' buy_orders '||v.s_buy||'/'||v.e_buy ELSE '' END
	,'; ' ORDER BY v.contract_aid),1000)
FROM (
	WITH trade AS (SELECT contract_aid, COUNT(*) c, COALESCE(SUM(price),0) p FROM rwb GROUP BY 1),
	offers AS (
		SELECT o.contract_aid,
		       COUNT(*) FILTER (WHERE o.otype=1 AND o.active)  AS open_sell,
		       COUNT(*) FILTER (WHERE o.otype<>1 AND o.active) AS open_buy
		FROM (SELECT o.contract_aid, o.otype,
		             NOT (EXISTS (SELECT 1 FROM rw_item_bought b
		                          WHERE b.contract_aid=o.contract_aid AND b.offer_id=o.offer_id)
		               OR EXISTS (SELECT 1 FROM rw_offer_canceled c
		                          WHERE c.contract_aid=o.contract_aid AND c.offer_id=o.offer_id)) AS active
		      FROM rw_new_offer o) o
		GROUP BY 1)
	SELECT COALESCE(s.contract_aid, COALESCE(t.contract_aid, o.contract_aid)) AS contract_aid,
	       COALESCE(s.total_num_trades,0)  s_trd,  COALESCE(t.c,0)         e_trd,
	       COALESCE(s.total_vol,0)         s_vol,  COALESCE(t.p,0)         e_vol,
	       COALESCE(s.total_sell_orders,0) s_sell, COALESCE(o.open_sell,0) e_sell,
	       COALESCE(s.total_buy_orders,0)  s_buy,  COALESCE(o.open_buy,0)  e_buy
	FROM rw_mkt_stats s
	FULL JOIN trade  t ON t.contract_aid=s.contract_aid
	FULL JOIN offers o ON o.contract_aid=COALESCE(s.contract_aid, t.contract_aid)
	WHERE COALESCE(s.total_num_trades,0)  <> COALESCE(t.c,0)
	   OR COALESCE(s.total_vol,0)         <> COALESCE(t.p,0)
	   OR COALESCE(s.total_sell_orders,0) <> COALESCE(o.open_sell,0)
	   OR COALESCE(s.total_buy_orders,0)  <> COALESCE(o.open_buy,0)
) v;

-- G4: rw_token state: trade counters, latest owner (rw_transfer), latest
-- price (mint or last purchase) and latest name.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. RandomWalk market','rw_token num_trades/total_vol/cur_owner/last_price/last_name','ERROR',COUNT(*),
       LEFT(STRING_AGG('token '||v.token_id,',' ORDER BY v.token_id),300)
FROM (
	SELECT t.token_id FROM rw_token t
	LEFT JOIN (SELECT rwalk_aid, token_id, COUNT(*) c, SUM(price) p FROM rwb GROUP BY 1,2) tr
	       ON tr.rwalk_aid=t.rwalk_aid AND tr.token_id=t.token_id
	LEFT JOIN LATERAL (SELECT x.to_aid FROM rw_transfer x
	                   WHERE x.contract_aid=t.rwalk_aid AND x.token_id=t.token_id
	                   ORDER BY x.evtlog_id DESC LIMIT 1) own ON TRUE
	LEFT JOIN LATERAL (
	        SELECT p.price FROM (
	            SELECT m.evtlog_id, m.price FROM rw_mint_evt m
	            WHERE m.contract_aid=t.rwalk_aid AND m.token_id=t.token_id
	            UNION ALL
	            SELECT b.evtlog_id, b.price FROM rwb b
	            WHERE b.rwalk_aid=t.rwalk_aid AND b.token_id=t.token_id) p
	        ORDER BY p.evtlog_id DESC LIMIT 1) pr ON TRUE
	LEFT JOIN LATERAL (SELECT n.new_name FROM rw_token_name n
	                   WHERE n.contract_aid=t.rwalk_aid AND n.token_id=t.token_id
	                   ORDER BY n.evtlog_id DESC LIMIT 1) nm ON TRUE
	WHERE COALESCE(t.num_trades,0) <> COALESCE(tr.c,0)
	   OR COALESCE(t.total_vol,0)  <> COALESCE(tr.p,0)
	   OR (own.to_aid IS NOT NULL AND t.cur_owner_aid <> own.to_aid)
	   OR COALESCE(t.last_price,0) <> COALESCE(pr.price,0)
	   OR COALESCE(t.last_name,'') <> COALESCE(nm.new_name,'')
) v;

-- G5: rw_user_stats per (contract, user): trade count and volume (buyer and
-- seller both count; self-trades once).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. RandomWalk market','rw_user_stats total_num_trades/total_vol','ERROR',COUNT(*),
       LEFT(STRING_AGG('user_aid='||v.user_aid,',' ORDER BY v.user_aid),300)
FROM (
	WITH trade AS (
		SELECT rwalk_aid, user_aid, COUNT(*) c, SUM(price) p FROM (
			SELECT rwalk_aid, buyer_aid user_aid, price FROM rwb
			UNION ALL
			SELECT rwalk_aid, seller_aid, price FROM rwb WHERE seller_aid <> buyer_aid) x
		GROUP BY 1,2)
	SELECT COALESCE(s.user_aid, e.user_aid) AS user_aid
	FROM rw_user_stats s
	FULL JOIN trade e ON e.rwalk_aid=s.rwalk_aid AND e.user_aid=s.user_aid
	WHERE COALESCE(s.total_num_trades,0) <> COALESCE(e.c,0)
	   OR COALESCE(s.total_vol,0)        <> COALESCE(e.p,0)
) v;

-- G5b: rw_user_stats.total_num_toks must equal the user's mint count.
-- Before migration 00028 the mint trigger silently dropped mints by users
-- with no rw_user_stats row yet; databases with such legacy drift are
-- repaired by scripts/repair-rw-stats.sql.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. RandomWalk market','rw_user_stats total_num_toks = mint count','ERROR',COUNT(*),
       LEFT(STRING_AGG('user_aid='||v.user_aid||' (stored '||v.stored||', minted '||v.minted||')',
                       ',' ORDER BY v.user_aid),300)
FROM (
	WITH mint AS (SELECT contract_aid rwalk_aid, owner_aid user_aid, COUNT(*) c
	              FROM rw_mint_evt GROUP BY 1,2)
	SELECT COALESCE(s.user_aid, m.user_aid) AS user_aid,
	       COALESCE(s.total_num_toks,0) stored, COALESCE(m.c,0) minted
	FROM rw_user_stats s
	FULL JOIN mint m ON m.rwalk_aid=s.rwalk_aid AND m.user_aid=s.user_aid
	WHERE COALESCE(s.total_num_toks,0) <> COALESCE(m.c,0)
) v;

-- G6 (WARN): seller profit. rw_new_offer.profit is only set when the seller
-- held a tracked cost basis (price_bought), so user total_profit must equal
-- the sum of profit over offers they sold. Order-dependent bookkeeping ->
-- WARN.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. RandomWalk market','rw_user_stats total_profit = sum of sold-offer profits','WARN',COUNT(*),
       LEFT(STRING_AGG('user_aid='||v.user_aid,',' ORDER BY v.user_aid),300)
FROM (
	WITH prof AS (
		SELECT b.rwalk_aid, b.seller_aid user_aid, SUM(COALESCE(o.profit,0)) p
		FROM rwb b JOIN rw_new_offer o ON o.contract_aid=b.contract_aid AND o.offer_id=b.offer_id
		GROUP BY 1,2)
	SELECT COALESCE(s.user_aid, p.user_aid) AS user_aid
	FROM rw_user_stats s
	FULL JOIN prof p ON p.rwalk_aid=s.rwalk_aid AND p.user_aid=s.user_aid
	WHERE COALESCE(s.total_profit,0) <> COALESCE(p.p,0)
) v;

-- G7: rw_user_rwtok cost basis (price_bought): last event wins - mint or
-- purchase sets it to the paid price for the new owner, a sale clears it for
-- the seller.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. RandomWalk market','rw_user_rwtok price_bought matches mint/buy/sell replay','ERROR',COUNT(*),
       LEFT(STRING_AGG('user '||v.user_aid||' token '||v.token_id,',' ),300)
FROM (
	WITH ev AS (
		SELECT contract_aid rwalk_aid, owner_aid user_aid, token_id, evtlog_id, price
		FROM rw_mint_evt
		UNION ALL
		SELECT rwalk_aid, buyer_aid, token_id, evtlog_id, price FROM rwb
		UNION ALL
		SELECT rwalk_aid, seller_aid, token_id, evtlog_id, NULL FROM rwb),
	exp AS (
		SELECT DISTINCT ON (rwalk_aid, user_aid, token_id)
		       rwalk_aid, user_aid, token_id, price
		FROM ev ORDER BY rwalk_aid, user_aid, token_id, evtlog_id DESC)
	SELECT COALESCE(t.user_aid,e.user_aid) user_aid, COALESCE(t.token_id,e.token_id) token_id
	FROM rw_user_rwtok t
	FULL JOIN exp e ON e.rwalk_aid=t.rwalk_aid AND e.user_aid=t.user_aid AND e.token_id=t.token_id
	WHERE t.user_aid IS NULL OR e.user_aid IS NULL
	   OR t.price_bought IS DISTINCT FROM e.price
) v;


-- ===========================================================================
-- REPORT
-- ===========================================================================

\echo ''
\echo '--- Check results ---------------------------------------------------------'
SELECT
	CASE WHEN violations = 0 THEN 'PASS'
	     WHEN severity = 'WARN' THEN 'WARN' ELSE 'FAIL' END AS status,
	section, check_name, violations
FROM vr ORDER BY id;

\echo ''
\echo '--- Failing checks (details) ----------------------------------------------'
SELECT CASE WHEN severity='WARN' THEN 'WARN' ELSE 'FAIL' END AS status,
       section, check_name, violations, details
FROM vr WHERE violations > 0 ORDER BY id;

\echo ''
\echo '--- Summary ----------------------------------------------------------------'
SELECT
	COUNT(*)                                                            AS checks_run,
	COUNT(*) FILTER (WHERE violations = 0)                              AS passed,
	COUNT(*) FILTER (WHERE violations > 0 AND severity = 'ERROR')       AS failed,
	COUNT(*) FILTER (WHERE violations > 0 AND severity = 'WARN')        AS warnings
FROM vr;

SELECT COUNT(*) > 0 AS has_errors
FROM vr WHERE violations > 0 AND severity = 'ERROR' \gset

ROLLBACK;

\if :has_errors
\warn 'STATS VALIDATION FAILED: at least one ERROR-severity check found violations.'
-- Force a non-zero psql exit status for cron/CI usage.
DO $$ BEGIN RAISE EXCEPTION 'statistics validation failed'; END $$;
\else
\echo 'STATS VALIDATION OK: no ERROR-severity violations found.'
\endif
