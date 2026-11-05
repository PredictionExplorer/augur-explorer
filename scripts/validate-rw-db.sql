-- ============================================================================
-- validate-rw-db.sql -- RandomWalk DB validator (event stream + mechanics)
--
-- Companion to validate-rw-stats.sql (statistics counters). This script
-- checks the RandomWalk event tables against the deterministic mechanics of
-- the deployed contracts (randomWalkNFT.sol, NFTMarket.sol) and against the
-- raw evt_log stream:
--
--   A. Event stream: evt_log linkage, per-topic completeness, unknown topics
--   B. RandomWalkNFT: dense token ids, exact mint-price chain, mint/transfer
--      pairing, ownership chains, withdrawal replay, name length, seeds
--   C. NFTMarket: dense offer ids, offer references, single termination,
--      party consistency, purchase transfers
--
-- Read-only: only TEMPORARY tables are created and the whole run is rolled
-- back. Exits non-zero (raised exception) if any ERROR-severity check fails.
-- Run:  psql -f validate-rw-db.sql
-- ============================================================================

\set ON_ERROR_STOP on
\pset pager off
\echo ''
\echo '=== RandomWalk DB validator (event stream and contract mechanics) ==='
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

-- Well-known address ids: the two contracts and the zero address.
CREATE TEMP TABLE rwc ON COMMIT DROP AS
SELECT
	(SELECT a.address_id FROM address a JOIN rw_contracts c
	    ON LOWER(a.addr)=LOWER(c.randomwalk_addr))   AS rwalk_aid,
	(SELECT a.address_id FROM address a JOIN rw_contracts c
	    ON LOWER(a.addr)=LOWER(c.marketplace_addr))  AS market_aid,
	(SELECT address_id FROM address
	    WHERE addr='0x0000000000000000000000000000000000000000') AS zero_aid;

-- ===========================================================================
-- SECTION A: event-stream integrity
-- ===========================================================================

-- A0: every domain row carries its evt_log id (the schema allows NULL, and a
-- NULL would silently escape every evt_log join below).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. Event stream','no NULL evtlog_id in domain tables','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.src||' id='||v.id,',' ORDER BY v.src,v.id),300)
FROM (
	          SELECT 'mint'     src, id FROM rw_mint_evt            WHERE evtlog_id IS NULL
	UNION ALL SELECT 'transfer',     id FROM rw_transfer            WHERE evtlog_id IS NULL
	UNION ALL SELECT 'name',         id FROM rw_token_name          WHERE evtlog_id IS NULL
	UNION ALL SELECT 'withdrawal',   id FROM rw_withdrawal          WHERE evtlog_id IS NULL
	UNION ALL SELECT 'offer',        id FROM rw_new_offer           WHERE evtlog_id IS NULL
	UNION ALL SELECT 'bought',       id FROM rw_item_bought         WHERE evtlog_id IS NULL
	UNION ALL SELECT 'canceled',     id FROM rw_offer_canceled      WHERE evtlog_id IS NULL
	UNION ALL SELECT 'approval',     id FROM rw_approval            WHERE evtlog_id IS NULL
	UNION ALL SELECT 'approval_all', id FROM rw_approval_for_all    WHERE evtlog_id IS NULL
	UNION ALL SELECT 'ownership',    id FROM rw_ownership_transferred WHERE evtlog_id IS NULL
) v;

-- A1: every domain row's (block_num, tx_id) matches its evt_log row.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. Event stream','domain rows match their evt_log block/tx','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.src||' evtlog_id='||v.evtlog_id,',' ORDER BY v.evtlog_id),300)
FROM (
	SELECT d.src, d.evtlog_id FROM (
		          SELECT 'mint'     src, evtlog_id, block_num, tx_id FROM rw_mint_evt
		UNION ALL SELECT 'transfer',       evtlog_id, block_num, tx_id FROM rw_transfer
		UNION ALL SELECT 'name',           evtlog_id, block_num, tx_id FROM rw_token_name
		UNION ALL SELECT 'withdrawal',     evtlog_id, block_num, tx_id FROM rw_withdrawal
		UNION ALL SELECT 'offer',          evtlog_id, block_num, tx_id FROM rw_new_offer
		UNION ALL SELECT 'bought',         evtlog_id, block_num, tx_id FROM rw_item_bought
		UNION ALL SELECT 'canceled',       evtlog_id, block_num, tx_id FROM rw_offer_canceled
		UNION ALL SELECT 'approval',       evtlog_id, block_num, tx_id FROM rw_approval
		UNION ALL SELECT 'approval_all',   evtlog_id, block_num, tx_id FROM rw_approval_for_all
		UNION ALL SELECT 'ownership',      evtlog_id, block_num, tx_id FROM rw_ownership_transferred
	) d
	JOIN evt_log e ON e.id=d.evtlog_id
	WHERE e.block_num<>d.block_num OR e.tx_id<>d.tx_id
) v;

-- A2: exactly one mint per token and per evt_log row (rw_mint_evt has no
-- UNIQUE constraints in the base schema).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. Event stream','rw_mint_evt has no duplicate token or evtlog rows','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.what,',' ORDER BY v.what),300)
FROM (
	SELECT 'token '||token_id AS what FROM rw_mint_evt
	GROUP BY contract_aid, token_id HAVING COUNT(*)>1
	UNION ALL
	SELECT 'evtlog '||evtlog_id FROM rw_mint_evt
	GROUP BY evtlog_id HAVING COUNT(*)>1
) v;

-- Expected topic0_sig -> domain table correspondence.
CREATE TEMP TABLE topicmap ON COMMIT DROP AS
SELECT * FROM (VALUES
	-- RandomWalkNFT events (ERROR when unindexed)
	('ad2bc79f','MintEvent',            'nft',    'core'),
	('ddf252ad','Transfer',             'nft',    'core'),
	('8ad5e159','TokenNameEvent',       'nft',    'core'),
	('a11b556a','WithdrawalEvent',      'nft',    'core'),
	-- inherited OpenZeppelin events (WARN: history predates handlers)
	('8c5be1e5','Approval',             'nft',    'inherited'),
	('17307eab','ApprovalForAll',       'nft',    'inherited'),
	('8be0079c','OwnershipTransferred', 'nft',    'inherited'),
	-- NFTMarket events
	('55076e90','NewOffer',             'market', 'core'),
	('caacc56f','ItemBought',           'market', 'core'),
	('0ff09947','OfferCanceled',        'market', 'core')
) t(sig, event, src, kind);

CREATE TEMP TABLE topic_counts ON COMMIT DROP AS
SELECT m.sig, m.event, m.kind,
	(SELECT COUNT(*) FROM evt_log e
	 WHERE e.topic0_sig=m.sig
	   AND e.contract_aid = CASE m.src WHEN 'nft' THEN (SELECT rwalk_aid FROM rwc)
	                                   ELSE (SELECT market_aid FROM rwc) END) AS in_evtlog,
	CASE m.event
		WHEN 'MintEvent'            THEN (SELECT COUNT(*) FROM rw_mint_evt)
		WHEN 'Transfer'             THEN (SELECT COUNT(*) FROM rw_transfer)
		WHEN 'TokenNameEvent'       THEN (SELECT COUNT(*) FROM rw_token_name)
		WHEN 'WithdrawalEvent'      THEN (SELECT COUNT(*) FROM rw_withdrawal)
		WHEN 'Approval'             THEN (SELECT COUNT(*) FROM rw_approval)
		WHEN 'ApprovalForAll'       THEN (SELECT COUNT(*) FROM rw_approval_for_all)
		WHEN 'OwnershipTransferred' THEN (SELECT COUNT(*) FROM rw_ownership_transferred)
		WHEN 'NewOffer'             THEN (SELECT COUNT(*) FROM rw_new_offer)
		WHEN 'ItemBought'           THEN (SELECT COUNT(*) FROM rw_item_bought)
		WHEN 'OfferCanceled'        THEN (SELECT COUNT(*) FROM rw_offer_canceled)
	END AS in_domain
FROM topicmap m;

-- A3: every stored core event log is decoded into its domain table.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. Event stream','core evt_log rows decoded into domain tables','ERROR',COUNT(*),
       LEFT(STRING_AGG(event||' evtlog/domain '||in_evtlog||'/'||in_domain,
                       '; ' ORDER BY event),300)
FROM topic_counts WHERE kind='core' AND in_evtlog<>in_domain;

-- A4: inherited-event history. Rows fetched before the 00035 handlers went
-- live sit in evt_log but were never decoded; this measures that gap.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. Event stream','inherited evt_log rows decoded (pre-00035 history -> WARN)','WARN',COUNT(*),
       LEFT(STRING_AGG(event||' evtlog/domain '||in_evtlog||'/'||in_domain,
                       '; ' ORDER BY event),300)
FROM topic_counts WHERE kind='inherited' AND in_evtlog<>in_domain;

-- A5: no unknown event signatures on either contract.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. Event stream','no unknown topic0 signatures on RW contracts','WARN',COUNT(*),
       LEFT(STRING_AGG(v.topic0_sig||' x'||v.c,',' ORDER BY v.c DESC),300)
FROM (
	SELECT e.topic0_sig, COUNT(*) c FROM evt_log e
	WHERE e.contract_aid IN ((SELECT rwalk_aid FROM rwc),(SELECT market_aid FROM rwc))
	  AND e.topic0_sig NOT IN (SELECT sig FROM topicmap)
	GROUP BY 1
) v;

-- ===========================================================================
-- SECTION B: RandomWalkNFT mechanics
-- ===========================================================================

-- B1: token ids are dense: 0..N-1, no gaps, no duplicates (nextTokenId+=1).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. RandomWalkNFT','token ids dense 0..N-1','ERROR',COUNT(*),
       LEFT(STRING_AGG('count='||v.c||' distinct='||v.d||' max='||v.m,','),300)
FROM (
	SELECT COUNT(*) c, COUNT(DISTINCT token_id) d, MAX(token_id) m
	FROM rw_mint_evt WHERE contract_aid=(SELECT rwalk_aid FROM rwc)
	HAVING COUNT(*)<>COUNT(DISTINCT token_id) OR COUNT(*)<>MAX(token_id)+1
) v;

-- B2: exact mint price chain: price starts at 10^15 wei and every mint pays
-- the EVM integer division prev*10011/10000 (getMintPrice()). DIV() is used
-- because NUMERIC '/' rounds its result, which breaks the chain by 1 wei
-- (first at token 2724 on the production data set).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. RandomWalkNFT','mint price = div(prev*10011,10000) chain from 0.001 ETH','ERROR',COUNT(*),
       LEFT(STRING_AGG('token '||v.token_id||' '||v.price||'/'||v.exp_price,
                       ',' ORDER BY v.token_id),300)
FROM (
	WITH RECURSIVE chain AS (
		SELECT 0::BIGINT AS token_id,
		       DIV(1000000000000000::NUMERIC*10011, 10000) AS exp_price
		UNION ALL
		SELECT token_id+1, DIV(exp_price*10011, 10000) FROM chain
		WHERE token_id < (SELECT MAX(token_id) FROM rw_mint_evt
		                  WHERE contract_aid=(SELECT rwalk_aid FROM rwc))
	)
	SELECT m.token_id, m.price, c.exp_price
	FROM rw_mint_evt m
	JOIN chain c ON c.token_id=m.token_id
	WHERE m.contract_aid=(SELECT rwalk_aid FROM rwc)
	  AND m.price IS DISTINCT FROM c.exp_price
) v;

-- B3: every mint pairs with exactly one mint-transfer (otype=1, from zero
-- address, to the mint owner) in the same transaction, and vice versa.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. RandomWalkNFT','mint <-> zero-address transfer pairing per tx','ERROR',COUNT(*),
       LEFT(STRING_AGG('token '||v.token_id,',' ORDER BY v.token_id),300)
FROM (
	SELECT COALESCE(m.token_id,t.token_id) AS token_id
	FROM (SELECT * FROM rw_mint_evt
	      WHERE contract_aid=(SELECT rwalk_aid FROM rwc)) m
	FULL JOIN (SELECT * FROM rw_transfer
	           WHERE contract_aid=(SELECT rwalk_aid FROM rwc) AND otype=1) t
	  ON t.token_id=m.token_id
	WHERE m.token_id IS NULL OR t.token_id IS NULL
	   OR t.tx_id<>m.tx_id
	   OR t.from_aid<>(SELECT zero_aid FROM rwc)
	   OR t.to_aid<>m.owner_aid
) v;

-- B4: per-token ownership chain: first transfer is the mint from the zero
-- address; every later transfer starts where the previous one ended.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. RandomWalkNFT','transfer chains unbroken (from = previous to)','ERROR',COUNT(*),
       LEFT(STRING_AGG(DISTINCT 'token '||v.token_id,',' ORDER BY 'token '||v.token_id),300)
FROM (
	SELECT token_id FROM (
		SELECT token_id, from_aid, otype,
		       LAG(to_aid)     OVER (PARTITION BY token_id ORDER BY evtlog_id) prev_to,
		       ROW_NUMBER()    OVER (PARTITION BY token_id ORDER BY evtlog_id) rn
		FROM rw_transfer WHERE contract_aid=(SELECT rwalk_aid FROM rwc)
	) t
	WHERE (rn=1 AND (otype<>1 OR from_aid<>(SELECT zero_aid FROM rwc)))
	   OR (rn>1 AND from_aid IS DISTINCT FROM prev_to)
) v;

-- B5a: withdrawal destination and token: withdraw() pays lastMinter and tags
-- the event with the last minted token, so aid = minter of token_id and
-- token_id = highest token minted before the withdrawal.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. RandomWalkNFT','withdrawal destination = last minter, token = last minted','ERROR',COUNT(*),
       LEFT(STRING_AGG('withdrawal evtlog='||v.evtlog_id,',' ORDER BY v.evtlog_id),300)
FROM (
	SELECT w.evtlog_id
	FROM rw_withdrawal w
	LEFT JOIN rw_mint_evt m
	  ON m.contract_aid=w.contract_aid AND m.token_id=w.token_id
	WHERE w.contract_aid=(SELECT rwalk_aid FROM rwc)
	  AND (m.owner_aid IS DISTINCT FROM w.aid
	    OR w.token_id <> (SELECT MAX(m2.token_id) FROM rw_mint_evt m2
	                      WHERE m2.contract_aid=w.contract_aid
	                        AND m2.evtlog_id < w.evtlog_id))
) v;

-- B5b: withdrawal amount replay: the contract balance is the sum of mint
-- prices minus previous withdrawals, and each withdrawal pays half of it.
-- WARN because force-sent ETH (selfdestruct) would legitimately break this.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. RandomWalkNFT','withdrawal amount = trunc(replayed balance / 2)','WARN',COUNT(*),
       LEFT(STRING_AGG('evtlog='||COALESCE(v.evtlog_id::TEXT,'?')||' '||v.amount||'/'||v.exp_amount,
                       ',' ORDER BY v.evtlog_id),300)
FROM (
	SELECT w.evtlog_id, w.amount,
	       TRUNC((
	         (SELECT COALESCE(SUM(m.price),0) FROM rw_mint_evt m
	          WHERE m.contract_aid=w.contract_aid AND m.evtlog_id<w.evtlog_id)
	       - (SELECT COALESCE(SUM(w2.amount),0) FROM rw_withdrawal w2
	          WHERE w2.contract_aid=w.contract_aid AND w2.evtlog_id<w.evtlog_id)
	       )/2) AS exp_amount
	FROM rw_withdrawal w
	WHERE w.contract_aid=(SELECT rwalk_aid FROM rwc)
) v
WHERE v.amount IS DISTINCT FROM v.exp_amount;

-- B6: token names respect the contract's 32-byte limit.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. RandomWalkNFT','token names <= 32 bytes','ERROR',COUNT(*),
       LEFT(STRING_AGG('token '||v.token_id,',' ORDER BY v.token_id),300)
FROM (
	SELECT token_id FROM rw_token_name
	WHERE OCTET_LENGTH(COALESCE(new_name,'')) > 32
) v;

-- B7: seed integrity: 64 hex chars, unique per contract, and seed_num (when
-- present) equals the numeric value of the hex seed.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. RandomWalkNFT','seeds: 64-hex, unique, seed_num matches hex','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.what,',' ORDER BY v.what),300)
FROM (
	SELECT 'malformed token '||token_id AS what FROM rw_mint_evt
	WHERE contract_aid=(SELECT rwalk_aid FROM rwc)
	  AND LOWER(REPLACE(seed,'0x','')) !~ '^[0-9a-f]{64}$'
	UNION ALL
	SELECT 'duplicate seed x'||COUNT(*) FROM rw_mint_evt
	WHERE contract_aid=(SELECT rwalk_aid FROM rwc)
	GROUP BY LOWER(REPLACE(seed,'0x','')) HAVING COUNT(*)>1
	UNION ALL
	SELECT 'seed_num token '||m.token_id FROM (
		SELECT token_id, seed_num,
		       (SELECT SUM( ('x'||SUBSTR(LOWER(REPLACE(seed,'0x','')), i*8+1, 8))::BIT(32)::BIGINT::NUMERIC
		                    * (2::NUMERIC ^ (32*(7-i))) )
		        FROM GENERATE_SERIES(0,7) i) AS exp_num
		FROM rw_mint_evt
		WHERE contract_aid=(SELECT rwalk_aid FROM rwc)
		  AND seed_num IS NOT NULL
		  AND LOWER(REPLACE(seed,'0x','')) ~ '^[0-9a-f]{64}$'
	) m WHERE m.seed_num <> m.exp_num
) v;

-- ===========================================================================
-- SECTION C: NFTMarket mechanics
-- ===========================================================================

-- C1: offer ids are dense: 0..N-1 per marketplace (numOffers+=1).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. NFTMarket','offer ids dense 0..N-1','ERROR',COUNT(*),
       LEFT(STRING_AGG('count='||v.c||' distinct='||v.d||' max='||v.m,','),300)
FROM (
	SELECT COUNT(*) c, COUNT(DISTINCT offer_id) d, MAX(offer_id) m
	FROM rw_new_offer WHERE contract_aid=(SELECT market_aid FROM rwc)
	HAVING COUNT(*)<>COUNT(DISTINCT offer_id) OR COUNT(*)<>MAX(offer_id)+1
) v;

-- C2: every purchase and cancellation references a recorded offer.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. NFTMarket','purchases/cancellations reference existing offers','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.what,',' ORDER BY v.what),300)
FROM (
	SELECT 'bought offer '||b.offer_id AS what
	FROM rw_item_bought b
	LEFT JOIN rw_new_offer o ON o.contract_aid=b.contract_aid AND o.offer_id=b.offer_id
	WHERE o.id IS NULL
	UNION ALL
	SELECT 'canceled offer '||c.offer_id
	FROM rw_offer_canceled c
	LEFT JOIN rw_new_offer o ON o.contract_aid=c.contract_aid AND o.offer_id=c.offer_id
	WHERE o.id IS NULL
) v;

-- C3: an offer terminates at most once (one purchase xor one cancellation).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. NFTMarket','offers terminated at most once','ERROR',COUNT(*),
       LEFT(STRING_AGG('offer '||v.offer_id||' x'||v.c,',' ORDER BY v.offer_id),300)
FROM (
	SELECT offer_id, COUNT(*) c FROM (
		SELECT contract_aid, offer_id FROM rw_item_bought
		UNION ALL
		SELECT contract_aid, offer_id FROM rw_offer_canceled
	) t GROUP BY contract_aid, offer_id HAVING COUNT(*)>1
) v;

-- C4: purchase parties match the offer: a sell offer (otype=1) is sold by
-- its maker, a buy offer (otype=0) is bought by its maker.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. NFTMarket','purchase parties consistent with offer type','ERROR',COUNT(*),
       LEFT(STRING_AGG('offer '||v.offer_id,',' ORDER BY v.offer_id),300)
FROM (
	SELECT b.offer_id
	FROM rw_item_bought b
	JOIN rw_new_offer o ON o.contract_aid=b.contract_aid AND o.offer_id=b.offer_id
	WHERE (o.otype=1 AND b.seller_aid<>o.seller_aid)
	   OR (o.otype<>1 AND b.buyer_aid<>o.buyer_aid)
) v;

-- C5: every purchase of a RandomWalk token moves it to the buyer in the same
-- tx. Scoped to RandomWalk offers: the marketplace accepts any ERC-721, and
-- transfers of foreign NFT contracts are not indexed (e.g. production offers
-- 459-461).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. NFTMarket','each RW purchase has the token transfer to the buyer in-tx','ERROR',COUNT(*),
       LEFT(STRING_AGG('offer '||v.offer_id,',' ORDER BY v.offer_id),300)
FROM (
	SELECT b.offer_id
	FROM rw_item_bought b
	JOIN rw_new_offer o ON o.contract_aid=b.contract_aid AND o.offer_id=b.offer_id
	WHERE o.rwalk_aid=(SELECT rwalk_aid FROM rwc)
	  AND NOT EXISTS (
		SELECT 1 FROM rw_transfer t
		WHERE t.contract_aid=o.rwalk_aid AND t.token_id=o.token_id
		  AND t.tx_id=b.tx_id AND t.to_aid=b.buyer_aid)
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
\warn 'RW DB VALIDATION FAILED: at least one ERROR-severity check found violations.'
DO $$ BEGIN RAISE EXCEPTION 'randomwalk db validation failed'; END $$;
\else
\echo 'RW DB VALIDATION OK: no ERROR-severity violations found.'
\endif
