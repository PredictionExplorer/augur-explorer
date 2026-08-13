-- ============================================================================
-- validate-stats.sql -- Cosmic Signature STATISTICAL COUNTERS validator
--
-- Companion to scripts/validate-db.sql (game mechanics). This script checks
-- every trigger-maintained statistics table against a recomputation from the
-- event fact tables:
--   * cg_round_stats / cg_glob_stats
--   * per-account stats (cg_winner, cg_bidder, cg_raffle_winner_stats,
--     cg_raffle_nft_winner_stats, cg_donor, cg_transfer_stats, cg_costok_owner)
--   * donation stats (cg_erc20_donation_stats, cg_winner.unclaimed_nfts,
--     cg_nft_stats)
--   * the CST staking reward pipeline (cg_st_reward, cg_staked_token_cst,
--     cg_staked_token_cst_rewards, cg_staker_cst, cg_staker_deposit,
--     cg_stake_stats_cst) via a full replay of stake/unstake/deposit events
--   * RandomWalk staking (cg_staker_rwalk, cg_staked_token_rwalk,
--     cg_stake_stats_rwalk)
--   * the RandomWalk marketplace (rw_stats, rw_mkt_stats, rw_token,
--     rw_user_stats, rw_user_rwtok, rw_new_offer.active)
--
-- Columns that no trigger or handler ever writes are intentionally NOT
-- checked (dead schema): rw_stats.total_withdrawals,
-- rw_user_stats.total_withdrawals, cg_stake_stats_cst.num_charity_deposits,
-- cg_stake_stats_cst.total_charity_amount,
-- cg_staked_token_cst_rewards.claimed_reward, cg_bidder.tokens_minted.
--
-- Usage (reads the same env vars as cg-etl: DATABASE_URL, or PGSQL_HOST /
-- PGSQL_USERNAME / PGSQL_PASSWORD / PGSQL_DATABASE):
--   scripts/validate-stats.sh
-- or invoke directly:
--   psql "$DATABASE_URL" -f scripts/validate-stats.sql
--
-- Read-only: only TEMPORARY tables are created and the whole run is rolled
-- back. Exits non-zero (raised exception) if any ERROR-severity check fails.
-- ============================================================================

\set ON_ERROR_STOP on
\pset pager off
\echo ''
\echo '=== Cosmic Signature DB validator (statistical counters) ==='
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

-- One row per completed round (claim facts used by the timing checks).
CREATE TEMP TABLE claims ON COMMIT DROP AS
SELECT round_num, evtlog_id AS claim_evtlog_id, tx_id AS claim_tx,
       time_stamp AS claim_ts, winner_aid AS main_winner_aid
FROM cg_prize_claim;

-- Well-known address ids.
CREATE TEMP TABLE known_aid ON COMMIT DROP AS
SELECT
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.cosmic_game_addr)) AS game_aid,
	(SELECT address_id FROM address WHERE addr='0x0000000000000000000000000000000000000000')          AS zero_aid;

-- ===========================================================================
-- SECTION A: cg_round_stats vs recomputation
-- ===========================================================================

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. round stats','cg_round_stats columns match recomputation from fact tables','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num||'.'||v.col||': stored '||v.got||' expected '||v.want,'; '),500)
FROM (
	WITH bids AS (
		SELECT round_num, COUNT(*) cnt, SUM(eth_price) eth,
		       COALESCE(SUM(cst_price) FILTER (WHERE bid_type=2),0) cst
		FROM cg_bid GROUP BY round_num),
	nftdon AS (SELECT round_num, COUNT(*) cnt FROM cg_nft_donation GROUP BY round_num),
	e20don AS (SELECT round_num, COUNT(*) cnt, COUNT(DISTINCT token_aid) ctr FROM cg_erc20_donation GROUP BY round_num),
	reth   AS (SELECT round_num, SUM(amount) s FROM cg_raffle_eth_prize GROUP BY round_num),
	rnft   AS (SELECT round_num, COUNT(*) cnt, SUM(cst_amount) cst FROM cg_raffle_nft_prize GROUP BY round_num),
	chrono AS (SELECT round_num, SUM(eth_amount) eth, SUM(cst_amount) cst, COUNT(*) cnt FROM cg_chrono_warrior_prize GROUP BY round_num),
	endur  AS (SELECT round_num, SUM(erc20_amount) cst, COUNT(*) cnt FROM cg_endurance_prize GROUP BY round_num),
	lcst   AS (SELECT round_num, SUM(erc20_amount) cst, COUNT(*) cnt FROM cg_lastcst_prize GROUP BY round_num),
	claimx AS (SELECT round_num, SUM(cst_amount) cst, COUNT(*) cnt FROM cg_prize_claim GROUP BY round_num),
	dons   AS (SELECT round_num, SUM(amount) total, COUNT(*) cnt FROM
	           (SELECT round_num, amount FROM cg_eth_donated
	            UNION ALL SELECT round_num, amount FROM cg_eth_donated_wi) x GROUP BY round_num)
	SELECT r.round_num, v.col, v.got::TEXT AS got, v.want::TEXT AS want
	FROM cg_round_stats r
	LEFT JOIN bids   b  USING (round_num)
	LEFT JOIN nftdon nd USING (round_num)
	LEFT JOIN e20don ed USING (round_num)
	LEFT JOIN reth   re USING (round_num)
	LEFT JOIN rnft   rn USING (round_num)
	LEFT JOIN chrono ch USING (round_num)
	LEFT JOIN endur  en USING (round_num)
	LEFT JOIN lcst   lc USING (round_num)
	LEFT JOIN claimx cl USING (round_num)
	LEFT JOIN dons   dn USING (round_num),
	LATERAL (VALUES
		('total_bids',                r.total_bids::NUMERIC,                COALESCE(b.cnt,0)::NUMERIC),
		('total_eth_in_bids',         r.total_eth_in_bids::NUMERIC,         COALESCE(b.eth,0)::NUMERIC),
		('total_cst_in_bids',         r.total_cst_in_bids::NUMERIC,         COALESCE(b.cst,0)::NUMERIC),
		('total_nft_donated',         r.total_nft_donated::NUMERIC,         COALESCE(nd.cnt,0)::NUMERIC),
		('num_erc20_donations',       r.num_erc20_donations::NUMERIC,       COALESCE(ed.cnt,0)::NUMERIC),
		('num_contracts_donated_erc20',r.num_contracts_donated_erc20::NUMERIC,COALESCE(ed.ctr,0)::NUMERIC),
		('total_raffle_eth_deposits', r.total_raffle_eth_deposits::NUMERIC, COALESCE(re.s,0)::NUMERIC),
		('total_raffle_nfts',         r.total_raffle_nfts::NUMERIC,         COALESCE(rn.cnt,0)::NUMERIC),
		('chrono_warrior_prize_eth',  r.chrono_warrior_prize_eth::NUMERIC,  COALESCE(ch.eth,0)::NUMERIC),
		('total_cst_paid_in_prizes',  r.total_cst_paid_in_prizes::NUMERIC,
		     COALESCE(cl.cst,0)+COALESCE(en.cst,0)+COALESCE(lc.cst,0)+COALESCE(ch.cst,0)+COALESCE(rn.cst,0)),
		('total_nfts_minted',         r.total_nfts_minted::NUMERIC,
		     COALESCE(cl.cnt,0)+COALESCE(en.cnt,0)+COALESCE(lc.cnt,0)+COALESCE(ch.cnt,0)+COALESCE(rn.cnt,0)),
		('donations_round_total',     r.donations_round_total::NUMERIC,     COALESCE(dn.total,0)::NUMERIC),
		('donations_round_count',     r.donations_round_count::NUMERIC,     COALESCE(dn.cnt,0)::NUMERIC)
	) v(col, got, want)
	WHERE v.got IS DISTINCT FROM v.want
) v;

-- A2: round timing columns.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. round stats','round timing: start=FirstBid, end=claim, duration, next param window','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.msg,'; '),400)
FROM (
	SELECT 'round '||c.round_num||' end/duration' AS msg
	FROM claims c JOIN cg_round_stats r USING (round_num)
	WHERE r.round_end_time IS DISTINCT FROM c.claim_ts
	   OR (r.round_start_time IS NOT NULL AND r.round_duration_seconds IS DISTINCT FROM
	       EXTRACT(EPOCH FROM (c.claim_ts - r.round_start_time))::BIGINT)
	UNION ALL
	SELECT 'round '||f.round_num||' start_time'
	FROM cg_first_bid f JOIN cg_round_stats r USING (round_num)
	WHERE r.round_start_time IS DISTINCT FROM TO_TIMESTAMP(f.start_ts)
	UNION ALL
	SELECT 'round '||(c.round_num+1)||' param_window_start'
	FROM claims c JOIN cg_round_stats r ON r.round_num = c.round_num+1
	WHERE r.param_window_start_time IS DISTINCT FROM c.claim_ts
) v;

-- ===========================================================================
-- SECTION B: cg_glob_stats vs recomputation (trigger-exact semantics)
-- ===========================================================================

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'B. global stats','cg_glob_stats matches recomputation from fact tables','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.col||': stored '||v.got||' expected '||v.want,'; '),500)
FROM (
	SELECT v.col, v.got::TEXT AS got, v.want::TEXT AS want
	FROM cg_glob_stats g, known_aid k,
	LATERAL (VALUES
		('num_bids',       g.num_bids::NUMERIC,       (SELECT COUNT(*) FROM cg_bid)::NUMERIC),
		('num_rwalk_used', g.num_rwalk_used::NUMERIC, (SELECT COUNT(*) FROM cg_bid WHERE rwalk_nft_id > -1)::NUMERIC),
		('num_bids_cst',   g.num_bids_cst::NUMERIC,   (SELECT COUNT(*) FROM cg_bid WHERE bid_type=2 AND rwalk_nft_id=-1)::NUMERIC),
		('total_cst_consumed', g.total_cst_consumed::NUMERIC,
		    (SELECT COALESCE(SUM(cst_price),0) FROM cg_bid WHERE bid_type=2 AND rwalk_nft_id=-1)),
		('cur_num_bids',   g.cur_num_bids::NUMERIC,
		    (SELECT COUNT(*) FROM cg_bid WHERE round_num > (SELECT COALESCE(MAX(round_num),-1) FROM cg_prize_claim))::NUMERIC),
		('num_wins',       g.num_wins::NUMERIC,       (SELECT COUNT(*) FROM cg_prize_claim)::NUMERIC),
		('num_mints',      g.num_mints::NUMERIC,      (SELECT COUNT(*) FROM cg_mint_event)::NUMERIC),
		('total_raffle_eth_deposits', g.total_raffle_eth_deposits::NUMERIC,
		    (SELECT COALESCE(SUM(amount),0) FROM cg_raffle_eth_prize)),
		('total_raffle_eth_withdrawn', g.total_raffle_eth_withdrawn::NUMERIC,
		    (SELECT COALESCE(SUM(amount),0) FROM cg_prize_withdrawal)),
		('total_chrono_warrior_eth_deposits', g.total_chrono_warrior_eth_deposits::NUMERIC,
		    (SELECT COALESCE(SUM(eth_amount),0) FROM cg_chrono_warrior_prize)),
		('total_cst_given_in_prizes', g.total_cst_given_in_prizes::NUMERIC,
		    (SELECT COALESCE(SUM(cst_amount),0) FROM cg_prize_claim)
		  + (SELECT COALESCE(SUM(erc20_amount),0) FROM cg_endurance_prize)
		  + (SELECT COALESCE(SUM(erc20_amount),0) FROM cg_lastcst_prize)
		  + (SELECT COALESCE(SUM(cst_amount),0) FROM cg_chrono_warrior_prize)
		  + (SELECT COALESCE(SUM(cst_amount),0) FROM cg_raffle_nft_prize)),
		('num_vol_donations', g.num_vol_donations::NUMERIC,
		    (SELECT COUNT(*) FROM cg_donation_received WHERE donor_aid <> k.game_aid)::NUMERIC),
		('vol_donations_total', g.vol_donations_total::NUMERIC,
		    (SELECT COALESCE(SUM(amount),0) FROM cg_donation_received WHERE donor_aid <> k.game_aid)),
		('num_cg_donations', g.num_cg_donations::NUMERIC,
		    (SELECT COUNT(*) FROM cg_donation_received WHERE donor_aid = k.game_aid)::NUMERIC),
		('cg_donations_total', g.cg_donations_total::NUMERIC,
		    (SELECT COALESCE(SUM(amount),0) FROM cg_donation_received WHERE donor_aid = k.game_aid)),
		('num_direct_donations', g.num_direct_donations::NUMERIC,
		    ((SELECT COUNT(*) FROM cg_eth_donated) + (SELECT COUNT(*) FROM cg_eth_donated_wi))::NUMERIC),
		('direct_donations', g.direct_donations::NUMERIC,
		    (SELECT COALESCE(SUM(amount),0) FROM cg_eth_donated)
		  + (SELECT COALESCE(SUM(amount),0) FROM cg_eth_donated_wi)),
		('num_withdrawals', g.num_withdrawals::NUMERIC, (SELECT COUNT(*) FROM cg_donation_sent)::NUMERIC),
		('sum_withdrawals', g.sum_withdrawals::NUMERIC, (SELECT COALESCE(SUM(amount),0) FROM cg_donation_sent)),
		('total_nft_donated', g.total_nft_donated::NUMERIC, (SELECT COUNT(*) FROM cg_nft_donation)::NUMERIC),
		('total_erc20_donations', g.total_erc20_donations::NUMERIC, (SELECT COUNT(*) FROM cg_erc20_donation)::NUMERIC),
		('num_mkt_rewards', g.num_mkt_rewards::NUMERIC, (SELECT COUNT(*) FROM cg_mkt_reward)::NUMERIC),
		('total_mkt_rewards', g.total_mkt_rewards::NUMERIC, (SELECT COALESCE(SUM(amount),0) FROM cg_mkt_reward))
	) v(col, got, want)
	WHERE v.got IS DISTINCT FROM v.want
) v;

-- ===========================================================================
-- SECTION C: per-account statistics tables
-- ===========================================================================

-- C1: cg_winner counters (increments per trigger: main +3/+1/+1, raffleETH +1,
-- raffleNFT/endurance/lastCst +2/+1/+1, chrono +3/+1/+1).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. account stats','cg_winner prizes_count/erc20_count/erc721_count/max_win_amount','ERROR',COUNT(*),
       LEFT(STRING_AGG('winner_aid='||v.winner_aid,',' ORDER BY v.winner_aid),300)
FROM (
	WITH m  AS (SELECT winner_aid, COUNT(*) c, MAX(amount) mx FROM cg_prize_claim GROUP BY 1),
	re AS (SELECT winner_aid, COUNT(*) c FROM cg_raffle_eth_prize GROUP BY 1),
	rn AS (SELECT winner_aid, COUNT(*) c FROM cg_raffle_nft_prize GROUP BY 1),
	en AS (SELECT winner_aid, COUNT(*) c FROM cg_endurance_prize GROUP BY 1),
	lc AS (SELECT winner_aid, COUNT(*) c FROM cg_lastcst_prize GROUP BY 1),
	ch AS (SELECT winner_aid, COUNT(*) c FROM cg_chrono_warrior_prize GROUP BY 1),
	exp AS (
		SELECT winner_aid,
		       3*COALESCE(m.c,0)+COALESCE(re.c,0)+2*COALESCE(rn.c,0)
		      +2*COALESCE(en.c,0)+2*COALESCE(lc.c,0)+3*COALESCE(ch.c,0) AS prizes_count,
		       COALESCE(m.c,0)+COALESCE(rn.c,0)+COALESCE(en.c,0)+COALESCE(lc.c,0)+COALESCE(ch.c,0) AS erc_count,
		       COALESCE(m.mx,0) AS max_win
		FROM m FULL JOIN re USING (winner_aid) FULL JOIN rn USING (winner_aid)
		       FULL JOIN en USING (winner_aid) FULL JOIN lc USING (winner_aid)
		       FULL JOIN ch USING (winner_aid))
	SELECT COALESCE(w.winner_aid, e.winner_aid) AS winner_aid
	FROM cg_winner w FULL JOIN exp e USING (winner_aid)
	WHERE COALESCE(w.prizes_count,0) <> COALESCE(e.prizes_count,0)
	   OR COALESCE(w.erc20_count,0)  <> COALESCE(e.erc_count,0)
	   OR COALESCE(w.erc721_count,0) <> COALESCE(e.erc_count,0)
	   OR COALESCE(w.max_win_amount,0) <> COALESCE(e.max_win,0)
) v;

-- C2 (WARN): cg_winner.prizes_sum is order-dependent (the claim trigger
-- overwrites it with SUM(main) while raffle/chrono triggers increment it), so
-- this models the log order: main sum + raffle/chrono amounts recorded after
-- the winner's last claim.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. account stats','cg_winner.prizes_sum (order-dependent model)','WARN',COUNT(*),
       LEFT(STRING_AGG('winner_aid='||v.winner_aid,',' ORDER BY v.winner_aid),300)
FROM (
	WITH lastclaim AS (SELECT winner_aid, MAX(evtlog_id) ev FROM cg_prize_claim GROUP BY 1),
	exp AS (
		SELECT a.winner_aid,
		       COALESCE((SELECT SUM(amount) FROM cg_prize_claim p WHERE p.winner_aid=a.winner_aid),0)
		     + COALESCE((SELECT SUM(r.amount) FROM cg_raffle_eth_prize r
		                 LEFT JOIN lastclaim l ON l.winner_aid=r.winner_aid
		                 WHERE r.winner_aid=a.winner_aid AND (l.ev IS NULL OR r.evtlog_id > l.ev)),0)
		     + COALESCE((SELECT SUM(c.eth_amount) FROM cg_chrono_warrior_prize c
		                 LEFT JOIN lastclaim l ON l.winner_aid=c.winner_aid
		                 WHERE c.winner_aid=a.winner_aid AND (l.ev IS NULL OR c.evtlog_id > l.ev)),0) AS s
		FROM (SELECT winner_aid FROM cg_winner
		      UNION SELECT winner_aid FROM cg_prize_claim
		      UNION SELECT winner_aid FROM cg_raffle_eth_prize
		      UNION SELECT winner_aid FROM cg_chrono_warrior_prize) a)
	SELECT e.winner_aid FROM exp e LEFT JOIN cg_winner w USING (winner_aid)
	WHERE COALESCE(w.prizes_sum,0) <> e.s
) v;

-- C3: cg_bidder counters. Details name each divergent column with
-- stored/expected values so a failure pinpoints the drift.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. account stats','cg_bidder num_bids/max_bid/total_eth_spent/total_cst_spent','ERROR',COUNT(*),
       LEFT(STRING_AGG(
	'aid='||v.bidder_aid
	|| CASE WHEN v.s_cnt <> v.e_cnt THEN ' num_bids '||v.s_cnt||'/'||v.e_cnt ELSE '' END
	|| CASE WHEN v.s_mx  <> v.e_mx  THEN ' max_bid '||v.s_mx||'/'||v.e_mx ELSE '' END
	|| CASE WHEN v.s_eth <> v.e_eth THEN ' eth '||v.s_eth||'/'||v.e_eth ELSE '' END
	|| CASE WHEN v.s_cst <> v.e_cst THEN ' cst '||v.s_cst||'/'||v.e_cst ELSE '' END
	,'; ' ORDER BY v.bidder_aid),1500)
FROM (
	SELECT COALESCE(bd.bidder_aid, e.bidder_aid) AS bidder_aid,
	       COALESCE(bd.num_bids,0) s_cnt,        COALESCE(e.cnt,0) e_cnt,
	       COALESCE(bd.max_bid,0)  s_mx,         COALESCE(e.mx,0)  e_mx,
	       COALESCE(bd.total_eth_spent,0) s_eth, COALESCE(e.eth,0) e_eth,
	       COALESCE(bd.total_cst_spent,0) s_cst, COALESCE(e.cst,0) e_cst
	FROM cg_bidder bd
	FULL JOIN (SELECT bidder_aid, COUNT(*) cnt, MAX(eth_price) mx,
	                  COALESCE(SUM(eth_price) FILTER (WHERE eth_price > 0),0) eth,
	                  COALESCE(SUM(cst_price) FILTER (WHERE cst_price > 0),0) cst
	           FROM cg_bid GROUP BY bidder_aid) e USING (bidder_aid)
	WHERE COALESCE(bd.num_bids,0) <> COALESCE(e.cnt,0)
	   OR COALESCE(bd.max_bid,0)  <> COALESCE(e.mx,0)
	   OR COALESCE(bd.total_eth_spent,0) <> COALESCE(e.eth,0)
	   OR COALESCE(bd.total_cst_spent,0) <> COALESCE(e.cst,0)
) v;

-- C4: cg_raffle_winner_stats (amount_sum accumulates raffle ETH + chrono ETH
-- and is reduced by withdrawals; withdrawal_sum keeps history).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. account stats','cg_raffle_winner_stats amount_sum/withdrawal_sum/raffles_count','ERROR',COUNT(*),
       LEFT(STRING_AGG('winner_aid='||v.winner_aid,',' ORDER BY v.winner_aid),300)
FROM (
	WITH re AS (SELECT winner_aid, SUM(amount) s, COUNT(*) c FROM cg_raffle_eth_prize GROUP BY 1),
	ch AS (SELECT winner_aid, SUM(eth_amount) s FROM cg_chrono_warrior_prize GROUP BY 1),
	wd AS (SELECT winner_aid, SUM(amount) s FROM cg_prize_withdrawal GROUP BY 1),
	exp AS (SELECT winner_aid,
	               COALESCE(re.s,0)+COALESCE(ch.s,0)-COALESCE(wd.s,0) AS amount_sum,
	               COALESCE(wd.s,0) AS withdrawal_sum,
	               COALESCE(re.c,0) AS raffles_count
	        FROM re FULL JOIN ch USING (winner_aid) FULL JOIN wd USING (winner_aid))
	SELECT COALESCE(s.winner_aid, e.winner_aid) AS winner_aid
	FROM cg_raffle_winner_stats s FULL JOIN exp e USING (winner_aid)
	WHERE COALESCE(s.amount_sum,0)     <> COALESCE(e.amount_sum,0)
	   OR COALESCE(s.withdrawal_sum,0) <> COALESCE(e.withdrawal_sum,0)
	   OR COALESCE(s.raffles_count,0)  <> COALESCE(e.raffles_count,0)
) v;

-- C5: cg_raffle_nft_winner_stats.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. account stats','cg_raffle_nft_winner_stats num_won','ERROR',COUNT(*),
       LEFT(STRING_AGG('winner_aid='||v.winner_aid,',' ORDER BY v.winner_aid),300)
FROM (
	SELECT COALESCE(s.winner_aid, e.winner_aid) AS winner_aid
	FROM cg_raffle_nft_winner_stats s
	FULL JOIN (SELECT winner_aid, COUNT(*) c FROM cg_raffle_nft_prize GROUP BY 1) e USING (winner_aid)
	WHERE COALESCE(s.num_won,0) <> COALESCE(e.c,0)
) v;

-- C6: cg_donor (direct donations to the game via donate()/donateWithInfo()).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. account stats','cg_donor count_donations/total_eth_donated','ERROR',COUNT(*),
       LEFT(STRING_AGG('donor_aid='||v.donor_aid,',' ORDER BY v.donor_aid),300)
FROM (
	SELECT COALESCE(d.donor_aid, e.donor_aid) AS donor_aid
	FROM cg_donor d
	FULL JOIN (SELECT donor_aid, COUNT(*) c, SUM(amount) s FROM
	           (SELECT donor_aid, amount FROM cg_eth_donated
	            UNION ALL SELECT donor_aid, amount FROM cg_eth_donated_wi) x
	           GROUP BY donor_aid) e USING (donor_aid)
	WHERE COALESCE(d.count_donations,0) <> COALESCE(e.c,0)
	   OR COALESCE(d.total_eth_donated,0) <> COALESCE(e.s,0)
) v;

-- C7: cg_transfer_stats (each transfer: +1 for from and to; self-transfer +1 once).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. account stats','cg_transfer_stats erc721/erc20 transfer counters','ERROR',COUNT(*),
       LEFT(STRING_AGG('user_aid='||v.user_aid,',' ORDER BY v.user_aid),300)
FROM (
	WITH e721 AS (
		SELECT user_aid, SUM(c) c FROM (
			SELECT from_aid user_aid, COUNT(*) c FROM cg_erc721_transfer GROUP BY 1
			UNION ALL
			SELECT to_aid, COUNT(*) FROM cg_erc721_transfer WHERE from_aid <> to_aid GROUP BY 1) x
		GROUP BY 1),
	e20 AS (
		SELECT user_aid, SUM(c) c FROM (
			SELECT from_aid user_aid, COUNT(*) c FROM cg_erc20_transfer GROUP BY 1
			UNION ALL
			SELECT to_aid, COUNT(*) FROM cg_erc20_transfer WHERE from_aid <> to_aid GROUP BY 1) x
		GROUP BY 1)
	SELECT COALESCE(s.user_aid, COALESCE(a.user_aid, b.user_aid)) AS user_aid
	FROM cg_transfer_stats s
	FULL JOIN e721 a ON a.user_aid = s.user_aid
	FULL JOIN e20  b ON b.user_aid = COALESCE(s.user_aid, a.user_aid)
	WHERE COALESCE(s.erc721_num_transfers,0) <> COALESCE(a.c,0)
	   OR COALESCE(s.erc20_num_transfers,0)  <> COALESCE(b.c,0)
) v;

-- C8: cg_costok_owner balances (CST): received minus sent, zero address excluded.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'C. account stats','cg_costok_owner cur_balance = CST received - CST sent','ERROR',COUNT(*),
       LEFT(STRING_AGG('owner_aid='||v.owner_aid,',' ORDER BY v.owner_aid),300)
FROM (
	WITH flows AS (
		SELECT owner_aid, SUM(v) bal FROM (
			SELECT to_aid owner_aid, SUM(value) v FROM cg_erc20_transfer GROUP BY 1
			UNION ALL
			SELECT from_aid, -SUM(value) FROM cg_erc20_transfer GROUP BY 1) x
		GROUP BY 1)
	SELECT COALESCE(o.owner_aid, f.owner_aid) AS owner_aid
	FROM cg_costok_owner o
	FULL JOIN flows f USING (owner_aid), known_aid k
	WHERE COALESCE(o.owner_aid, f.owner_aid) IS DISTINCT FROM k.zero_aid
	  AND COALESCE(o.cur_balance,0) <> COALESCE(f.bal,0)
) v;

-- ===========================================================================
-- SECTION D: donation statistics
-- ===========================================================================

-- D1: cg_erc20_donation_stats totals match the donation fact table.
-- on_donated_tok_claimed_insert subtracts each claimed amount from
-- total_amount (and flips claimed once it reaches 0), so the maintained
-- figure is the UNCLAIMED remainder: donations - claims.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. donation stats','cg_erc20_donation_stats total_amount = donations - claims per (round, token)','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num||' token_aid '||v.token_aid
                       ||' (stored '||v.stored||', expected '||v.expected||')',',' ),300)
FROM (
	SELECT COALESCE(s.round_num,d.round_num) round_num,
	       COALESCE(s.token_aid,d.token_aid) token_aid,
	       COALESCE(s.total_amount,0) stored,
	       COALESCE(d.s,0) - COALESCE(c.s,0) expected
	FROM cg_erc20_donation_stats s
	FULL JOIN (SELECT round_num, token_aid, SUM(amount) s FROM cg_erc20_donation GROUP BY 1,2) d
	       USING (round_num, token_aid)
	LEFT JOIN (SELECT round_num, token_aid, SUM(amount) s FROM cg_donated_tok_claimed GROUP BY 1,2) c
	       ON c.round_num=COALESCE(s.round_num,d.round_num) AND c.token_aid=COALESCE(s.token_aid,d.token_aid)
	WHERE COALESCE(s.total_amount,0) <> COALESCE(d.s,0) - COALESCE(c.s,0)
) v;

-- D2: cg_winner.unclaimed_nfts. The claim trigger adds the round's donated-NFT
-- count to the round winner; every DonatedNftClaimed (by anyone) decrements
-- the ORIGINAL round winner, clamped at 0.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. donation stats','cg_winner.unclaimed_nfts = donated NFTs of won rounds - claims','ERROR',COUNT(*),
       LEFT(STRING_AGG('winner_aid='||v.winner_aid||' (stored '||v.stored||', expected '||v.expected||')',
                       ',' ORDER BY v.winner_aid),300)
FROM (
	WITH per_round AS (
		SELECT c.main_winner_aid AS winner_aid,
		       (SELECT COUNT(*) FROM cg_nft_donation d WHERE d.round_num=c.round_num)        AS donated,
		       (SELECT COUNT(*) FROM cg_donated_nft_claimed x WHERE x.round_num=c.round_num) AS claimed
		FROM claims c),
	exp AS (
		SELECT winner_aid, GREATEST(0, SUM(donated) - SUM(claimed)) AS unclaimed
		FROM per_round GROUP BY winner_aid)
	SELECT COALESCE(w.winner_aid, e.winner_aid) AS winner_aid,
	       COALESCE(w.unclaimed_nfts,0) stored, COALESCE(e.unclaimed,0) expected
	FROM cg_winner w FULL JOIN exp e USING (winner_aid)
	WHERE COALESCE(w.unclaimed_nfts,0) <> COALESCE(e.unclaimed,0)
) v;

-- D3: cg_nft_stats per donated-NFT contract.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. donation stats','cg_nft_stats.num_donated = NftDonated count per contract','ERROR',COUNT(*),
       LEFT(STRING_AGG('contract_aid='||v.contract_aid,',' ORDER BY v.contract_aid),300)
FROM (
	SELECT COALESCE(s.contract_aid, e.token_aid) AS contract_aid
	FROM cg_nft_stats s
	FULL JOIN (SELECT token_aid, COUNT(*) c FROM cg_nft_donation GROUP BY 1) e
	       ON e.token_aid = s.contract_aid
	WHERE COALESCE(s.num_donated,0) <> COALESCE(e.c,0)
) v;

-- ===========================================================================
-- SECTION E: CST staking reward pipeline.
-- Replays NftStaked / NftUnstaked / EthDepositReceived in log order:
-- a token is "staked at deposit time" when its stake event precedes the
-- deposit and no unstake of the same action precedes it. Each such token
-- accrues floor((depositAmount - depositAmount % n) / n); an unstake marks
-- every accrued reward of that action collected (V2 pays rewards on unstake).
-- ===========================================================================

CREATE TEMP TABLE st_exp ON COMMIT DROP AS
SELECT d.deposit_id, d.round_num, d.deposit_amount, d.evtlog_id AS dep_evtlog_id,
       s.action_id, s.token_id, s.staker_aid,
       EXISTS (SELECT 1 FROM cg_nft_unstaked_cst u WHERE u.action_id = s.action_id) AS collected
FROM cg_staking_eth_deposit d
JOIN cg_nft_staked_cst s
  ON s.evtlog_id < d.evtlog_id
 AND NOT EXISTS (SELECT 1 FROM cg_nft_unstaked_cst u
                 WHERE u.action_id = s.action_id AND u.evtlog_id < d.evtlog_id);

CREATE TEMP TABLE st_dep ON COMMIT DROP AS
SELECT deposit_id, round_num, deposit_amount,
       COUNT(*)                                     AS n,
       MOD(deposit_amount, COUNT(*)::NUMERIC)       AS dep_mod,
       (deposit_amount - MOD(deposit_amount, COUNT(*)::NUMERIC)) / COUNT(*) AS per_token
FROM st_exp GROUP BY deposit_id, round_num, deposit_amount;

-- E1 (WARN): a deposit event with no staked tokens should not exist (the
-- contract forwards the funds to charity instead of emitting a deposit).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. CST staking','staking deposit has >= 1 token staked at deposit time','WARN',COUNT(*),
       LEFT(STRING_AGG('deposit_id='||v.deposit_id,',' ORDER BY v.deposit_id),300)
FROM (
	SELECT d.deposit_id FROM cg_staking_eth_deposit d
	WHERE NOT EXISTS (SELECT 1 FROM st_dep x WHERE x.deposit_id=d.deposit_id)
) v;

-- E2: cg_st_reward = replayed accrual set, one row per (deposit, staked
-- token), with reward = per-token slice and collected/is_unstake flags set
-- exactly when the token's action was unstaked.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. CST staking','cg_st_reward rows/reward/collected match stake+deposit replay','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.msg,'; '),400)
FROM (
	SELECT COALESCE('deposit '||e.deposit_id||' action '||e.action_id,
	                'orphan row deposit '||r.deposit_id||' action '||r.action_id) AS msg
	FROM (SELECT x.*, d.per_token FROM st_exp x JOIN st_dep d USING (deposit_id)) e
	FULL JOIN cg_st_reward r ON r.deposit_id=e.deposit_id AND r.action_id=e.action_id
	WHERE e.action_id IS NULL OR r.action_id IS NULL
	   OR r.token_id  <> e.token_id
	   OR r.staker_aid <> e.staker_aid
	   OR r.round_num <> e.round_num
	   OR r.reward    <> e.per_token
	   OR r.collected <> e.collected
	   OR r.is_unstake <> e.collected
	LIMIT 100
) v;

-- E3: cg_staking_eth_deposit accumulator columns rewritten by the trigger:
-- accumulated_nfts = tokens staked at deposit time, num_staked_nfts = tokens
-- added since the previous deposit, accumulated_amount / accumulated_per_token
-- = running sums in deposit_id order, amount_per_token = this deposit's slice.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. CST staking','cg_staking_eth_deposit accumulated_nfts/num_staked_nfts/amounts match replay','ERROR',COUNT(*),
       LEFT(STRING_AGG('deposit_id='||v.deposit_id,',' ORDER BY v.deposit_id),300)
FROM (
	WITH e AS (
		SELECT deposit_id, n, per_token, dep_mod,
		       n - COALESCE(LAG(n) OVER w, 0)          AS added,
		       SUM(deposit_amount) OVER w              AS acc_amount,
		       SUM(per_token) OVER w                   AS acc_per_token
		FROM st_dep WINDOW w AS (ORDER BY deposit_id) )
	SELECT d.deposit_id
	FROM cg_staking_eth_deposit d JOIN e ON e.deposit_id=d.deposit_id
	WHERE d.accumulated_nfts      <> e.n
	   OR d.num_staked_nfts       <> e.added
	   OR d.amount_per_token      <> e.per_token
	   OR d.accumulated_amount    <> e.acc_amount
	   OR d.accumulated_per_token <> e.acc_per_token
) v;

-- E4: cg_staked_token_cst = the currently staked set (stake without unstake).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. CST staking','cg_staked_token_cst = stakes without a matching unstake','ERROR',COUNT(*),
       LEFT(STRING_AGG('token '||v.token_id,',' ORDER BY v.token_id),300)
FROM (
	SELECT COALESCE(t.token_id, e.token_id) AS token_id
	FROM cg_staked_token_cst t
	FULL JOIN (SELECT s.staker_aid, s.token_id, s.action_id FROM cg_nft_staked_cst s
	           WHERE NOT EXISTS (SELECT 1 FROM cg_nft_unstaked_cst u WHERE u.action_id=s.action_id)) e
	       ON e.action_id = t.stake_action_id
	WHERE t.stake_action_id IS NULL OR e.action_id IS NULL
	   OR t.token_id <> e.token_id OR t.staker_aid <> e.staker_aid
) v;

-- E5: cg_staked_token_cst_rewards keeps one row per stake action forever;
-- accumulated_reward = sum of that action's replayed rewards.
-- (claimed_reward has no writer anywhere - not checked.)
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. CST staking','cg_staked_token_cst_rewards accumulated_reward per stake action','ERROR',COUNT(*),
       LEFT(STRING_AGG('action '||v.action_id,',' ORDER BY v.action_id),300)
FROM (
	SELECT COALESCE(t.stake_action_id, s.action_id) AS action_id
	FROM cg_staked_token_cst_rewards t
	FULL JOIN cg_nft_staked_cst s ON s.action_id = t.stake_action_id
	LEFT JOIN (SELECT e.action_id, SUM(d.per_token) rew
	           FROM st_exp e JOIN st_dep d USING (deposit_id) GROUP BY e.action_id) r
	       ON r.action_id = COALESCE(t.stake_action_id, s.action_id)
	WHERE t.stake_action_id IS NULL OR s.action_id IS NULL
	   OR t.token_id <> s.token_id OR t.staker_aid <> s.staker_aid
	   OR COALESCE(t.accumulated_reward,0) <> COALESCE(r.rew,0)
) v;

-- E6: cg_staker_cst per-staker counters and rewards.
-- num_tokens_minted counts raffle CS-NFT prizes won as a (CST) staker.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. CST staking','cg_staker_cst stake/unstake counts, staked tokens, rewards, mints','ERROR',COUNT(*),
       LEFT(STRING_AGG('staker_aid='||v.staker_aid,',' ORDER BY v.staker_aid),300)
FROM (
	WITH st AS (SELECT staker_aid, COUNT(*) c FROM cg_nft_staked_cst GROUP BY 1),
	un AS (SELECT staker_aid, COUNT(*) c FROM cg_nft_unstaked_cst GROUP BY 1),
	rw AS (SELECT e.staker_aid, SUM(d.per_token) total,
	              SUM(d.per_token) FILTER (WHERE NOT e.collected) unclaimed
	       FROM st_exp e JOIN st_dep d USING (deposit_id) GROUP BY 1),
	mn AS (SELECT winner_aid staker_aid, COUNT(*) c FROM cg_raffle_nft_prize
	       WHERE is_staker AND NOT is_rwalk GROUP BY 1),
	exp AS (
		SELECT staker_aid,
		       COALESCE(st.c,0)                  AS n_stake,
		       COALESCE(un.c,0)                  AS n_unstake,
		       COALESCE(st.c,0)-COALESCE(un.c,0) AS staked,
		       COALESCE(rw.total,0)              AS total_reward,
		       COALESCE(rw.unclaimed,0)          AS unclaimed_reward,
		       COALESCE(mn.c,0)                  AS minted
		FROM st FULL JOIN un USING (staker_aid) FULL JOIN rw USING (staker_aid)
		        FULL JOIN mn USING (staker_aid))
	SELECT COALESCE(s.staker_aid, e.staker_aid) AS staker_aid
	FROM cg_staker_cst s FULL JOIN exp e USING (staker_aid)
	WHERE COALESCE(s.num_stake_actions,0)   <> COALESCE(e.n_stake,0)
	   OR COALESCE(s.num_unstake_actions,0) <> COALESCE(e.n_unstake,0)
	   OR COALESCE(s.total_tokens_staked,0) <> COALESCE(e.staked,0)
	   OR COALESCE(s.total_reward,0)        <> COALESCE(e.total_reward,0)
	   OR COALESCE(s.unclaimed_reward,0)    <> COALESCE(e.unclaimed_reward,0)
	   OR COALESCE(s.num_tokens_minted,0)   <> COALESCE(e.minted,0)
) v;

-- E7: cg_staker_deposit per (staker, deposit): tokens staked at deposit time,
-- the deposited slice, and the still-unclaimed part.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. CST staking','cg_staker_deposit tokens_staked/amount_deposited/amount_to_claim','ERROR',COUNT(*),
       LEFT(STRING_AGG('staker '||v.staker_aid||' deposit '||v.deposit_id,',' ),300)
FROM (
	WITH exp AS (
		SELECT e.staker_aid, e.deposit_id,
		       COUNT(*)                                          AS toks,
		       SUM(d.per_token)                                  AS dep,
		       COALESCE(SUM(d.per_token) FILTER (WHERE NOT e.collected),0) AS to_claim
		FROM st_exp e JOIN st_dep d USING (deposit_id)
		GROUP BY e.staker_aid, e.deposit_id)
	SELECT COALESCE(s.staker_aid,e.staker_aid) staker_aid,
	       COALESCE(s.deposit_id,e.deposit_id) deposit_id
	FROM cg_staker_deposit s
	FULL JOIN exp e ON e.staker_aid=s.staker_aid AND e.deposit_id=s.deposit_id
	WHERE e.deposit_id IS NULL OR s.deposit_id IS NULL
	   OR s.tokens_staked    <> e.toks
	   OR s.amount_deposited <> e.dep
	   OR s.amount_to_claim  <> e.to_claim
) v;

-- E8: cg_stake_stats_cst global counters.
-- (num_charity_deposits / total_charity_amount have no writer - not checked.)
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. CST staking','cg_stake_stats_cst totals match replay','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.col||': stored '||v.got||' expected '||v.want,'; '),400)
FROM (
	SELECT v.col, v.got::TEXT AS got, v.want::TEXT AS want
	FROM cg_stake_stats_cst g,
	LATERAL (VALUES
		('total_tokens_staked', g.total_tokens_staked::NUMERIC,
		    (SELECT COUNT(*) FROM cg_nft_staked_cst)::NUMERIC
		  - (SELECT COUNT(*) FROM cg_nft_unstaked_cst)::NUMERIC),
		('total_num_stakers', g.total_num_stakers::NUMERIC,
		    (SELECT COUNT(*) FROM (
		        SELECT s.staker_aid FROM cg_nft_staked_cst s GROUP BY 1
		        HAVING COUNT(*) > (SELECT COUNT(*) FROM cg_nft_unstaked_cst u
		                           WHERE u.staker_aid=s.staker_aid)) x)::NUMERIC),
		('num_deposits',       g.num_deposits::NUMERIC, (SELECT COUNT(*) FROM st_dep)::NUMERIC),
		('total_reward_amount', g.total_reward_amount::NUMERIC,
		    (SELECT COALESCE(SUM(deposit_amount - dep_mod),0) FROM st_dep)),
		('total_unclaimed_reward', g.total_unclaimed_reward::NUMERIC,
		    (SELECT COALESCE(SUM(d.per_token),0) FROM st_exp e JOIN st_dep d USING (deposit_id)
		     WHERE NOT e.collected)),
		('total_modulo', g.total_modulo::NUMERIC, (SELECT COALESCE(SUM(dep_mod),0) FROM st_dep))
	) v(col, got, want)
	WHERE v.got IS DISTINCT FROM v.want
) v;

-- ===========================================================================
-- SECTION F: RandomWalk NFT staking (no ETH rewards, counters only)
-- ===========================================================================

-- F1: cg_staked_token_rwalk = stakes without a matching unstake.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'F. RandomWalk staking','cg_staked_token_rwalk = stakes without a matching unstake','ERROR',COUNT(*),
       LEFT(STRING_AGG('token '||v.token_id,',' ORDER BY v.token_id),300)
FROM (
	SELECT COALESCE(t.token_id, e.token_id) AS token_id
	FROM cg_staked_token_rwalk t
	FULL JOIN (SELECT s.staker_aid, s.token_id, s.action_id FROM cg_nft_staked_rwalk s
	           WHERE NOT EXISTS (SELECT 1 FROM cg_nft_unstaked_rwalk u WHERE u.action_id=s.action_id)) e
	       ON e.action_id = t.stake_action_id
	WHERE t.stake_action_id IS NULL OR e.action_id IS NULL
	   OR t.token_id <> e.token_id OR t.staker_aid <> e.staker_aid
) v;

-- F2: cg_staker_rwalk counters; num_tokens_minted counts raffle CS-NFT prizes
-- won as a RandomWalk staker.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'F. RandomWalk staking','cg_staker_rwalk stake/unstake/staked/minted counters','ERROR',COUNT(*),
       LEFT(STRING_AGG('staker_aid='||v.staker_aid,',' ORDER BY v.staker_aid),300)
FROM (
	WITH st AS (SELECT staker_aid, COUNT(*) c FROM cg_nft_staked_rwalk GROUP BY 1),
	un AS (SELECT staker_aid, COUNT(*) c FROM cg_nft_unstaked_rwalk GROUP BY 1),
	mn AS (SELECT winner_aid staker_aid, COUNT(*) c FROM cg_raffle_nft_prize
	       WHERE is_staker AND is_rwalk GROUP BY 1),
	exp AS (
		SELECT staker_aid, COALESCE(st.c,0) n_stake, COALESCE(un.c,0) n_unstake,
		       COALESCE(st.c,0)-COALESCE(un.c,0) staked, COALESCE(mn.c,0) minted
		FROM st FULL JOIN un USING (staker_aid) FULL JOIN mn USING (staker_aid))
	SELECT COALESCE(s.staker_aid, e.staker_aid) AS staker_aid
	FROM cg_staker_rwalk s FULL JOIN exp e USING (staker_aid)
	WHERE COALESCE(s.num_stake_actions,0)   <> COALESCE(e.n_stake,0)
	   OR COALESCE(s.num_unstake_actions,0) <> COALESCE(e.n_unstake,0)
	   OR COALESCE(s.total_tokens_staked,0) <> COALESCE(e.staked,0)
	   OR COALESCE(s.num_tokens_minted,0)   <> COALESCE(e.minted,0)
) v;

-- F3: cg_stake_stats_rwalk global counters.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'F. RandomWalk staking','cg_stake_stats_rwalk totals','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.col||': stored '||v.got||' expected '||v.want,'; '),300)
FROM (
	SELECT v.col, v.got::TEXT AS got, v.want::TEXT AS want
	FROM cg_stake_stats_rwalk g,
	LATERAL (VALUES
		('total_tokens_staked', g.total_tokens_staked::NUMERIC,
		    (SELECT COUNT(*) FROM cg_nft_staked_rwalk)::NUMERIC
		  - (SELECT COUNT(*) FROM cg_nft_unstaked_rwalk)::NUMERIC),
		('total_num_stakers', g.total_num_stakers::NUMERIC,
		    (SELECT COUNT(*) FROM (
		        SELECT s.staker_aid FROM cg_nft_staked_rwalk s GROUP BY 1
		        HAVING COUNT(*) > (SELECT COUNT(*) FROM cg_nft_unstaked_rwalk u
		                           WHERE u.staker_aid=s.staker_aid)) x)::NUMERIC),
		('total_nft_mints', g.total_nft_mints::NUMERIC,
		    (SELECT COUNT(*) FROM cg_raffle_nft_prize WHERE is_staker AND is_rwalk)::NUMERIC)
	) v(col, got, want)
	WHERE v.got IS DISTINCT FROM v.want
) v;

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
