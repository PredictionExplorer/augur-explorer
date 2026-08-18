-- ============================================================================
-- validate-db.sql -- Cosmic Signature database GAME MECHANICS validator
--
-- Validates the game-mechanics side of the indexer database against the
-- on-chain rules of the Cosmic-Signature contracts (V2, production `main`
-- branch): layer-1 integrity, event decoding/linkage, per-round prize
-- structure, withdrawals, bidding, NFT/token state and donation facts.
--
-- The trigger-maintained STATISTICAL counters (cg_round_stats, cg_glob_stats,
-- per-account stats, staking accruals, RandomWalk marketplace stats) are
-- lower priority and validated separately by scripts/validate-stats.sql.
--
-- Usage (reads the same env vars as cg-etl: DATABASE_URL, or PGSQL_HOST /
-- PGSQL_USERNAME / PGSQL_PASSWORD / PGSQL_DATABASE):
--   scripts/validate-db.sh
-- or invoke directly:
--   psql "$DATABASE_URL" -f scripts/validate-db.sql
--
-- The script is read-only with respect to your data: it only creates
-- TEMPORARY tables and everything runs inside a transaction that is rolled
-- back at the end. It prints a report of every check and exits with a
-- non-zero status (via a raised exception) if any ERROR-severity check found
-- violations, so it can be used from cron/CI.
--
-- Severity levels:
--   ERROR - a hard invariant of the contracts / schema is broken.
--   WARN  - order-dependent or heuristic bookkeeping (may need manual review).
--
-- Contract constants assumed (deploy-time defaults, CosmicSignatureConstants.sol;
-- any later admin change is picked up automatically from the cg_adm_* tables):
--   numRaffleEthPrizesForBidders                            = 3
--   numRaffleCosmicSignatureNftsForBidders                  = 10
--   numRaffleCosmicSignatureNftsForRandomWalkNftStakers     = 10
--   mainEthPrizeAmountPercentage                            = 25
--   chronoWarriorEthPrizeAmountPercentage                   = 8
--   raffleTotalEthPrizeAmountForBiddersPercentage           = 4
--   cosmicSignatureNftStakingTotalEthRewardAmountPercentage = 6
--   charityEthDonationAmountPercentage                      = 7
--   timeoutDurationToWithdrawPrizes                         = 3024000 (5 weeks)
-- ============================================================================

\set ON_ERROR_STOP on
\pset pager off
\echo ''
\echo '=== Cosmic Signature DB validator (game mechanics) ==='
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

-- ---------------------------------------------------------------------------
-- Per-round effective configuration (deploy defaults + cg_adm_* overrides
-- effective at the claim block). One row per completed round.
-- ---------------------------------------------------------------------------
CREATE TEMP TABLE cfg ON COMMIT DROP AS
SELECT
	pc.round_num,
	pc.evtlog_id  AS claim_evtlog_id,
	pc.block_num  AS claim_block,
	pc.tx_id      AS claim_tx,
	pc.time_stamp AS claim_ts,
	pc.winner_aid AS main_winner_aid,
	pc.amount     AS main_eth,
	pc.cst_amount AS main_cst,
	pc.token_id   AS main_token_id,
	pc.timeout    AS claim_timeout,
	COALESCE((SELECT a.num_winners::BIGINT FROM cg_adm_raf_eth_bidding a
	          WHERE a.block_num <= pc.block_num ORDER BY a.evtlog_id DESC LIMIT 1), 3)  AS n_raffle_eth,
	COALESCE((SELECT a.num_winners::BIGINT FROM cg_adm_raf_nft_bidding a
	          WHERE a.block_num <= pc.block_num ORDER BY a.evtlog_id DESC LIMIT 1), 10) AS n_nft_bidders,
	COALESCE((SELECT a.num_winners::BIGINT FROM cg_adm_raf_nft_staking_rwalk a
	          WHERE a.block_num <= pc.block_num ORDER BY a.evtlog_id DESC LIMIT 1), 10) AS n_nft_rwalk,
	COALESCE((SELECT a.percentage::NUMERIC FROM cg_adm_main_prize_pcent a
	          WHERE a.block_num <= pc.block_num ORDER BY a.evtlog_id DESC LIMIT 1), 25) AS pct_main,
	COALESCE((SELECT a.percentage::NUMERIC FROM cg_adm_chrono_pcent a
	          WHERE a.block_num <= pc.block_num ORDER BY a.evtlog_id DESC LIMIT 1), 8)  AS pct_chrono,
	COALESCE((SELECT a.percentage::NUMERIC FROM cg_adm_raffle_pcent a
	          WHERE a.block_num <= pc.block_num ORDER BY a.evtlog_id DESC LIMIT 1), 4)  AS pct_raffle,
	COALESCE((SELECT a.percentage::NUMERIC FROM cg_adm_stake_pcent a
	          WHERE a.block_num <= pc.block_num ORDER BY a.evtlog_id DESC LIMIT 1), 6)  AS pct_stake,
	COALESCE((SELECT a.percentage::NUMERIC FROM cg_adm_charity_pcent a
	          WHERE a.block_num <= pc.block_num ORDER BY a.evtlog_id DESC LIMIT 1), 7)  AS pct_charity,
	COALESCE((SELECT a.new_timeout FROM cg_adm_timeout_withdraw a
	          WHERE a.block_num <= pc.block_num ORDER BY a.evtlog_id DESC LIMIT 1), 3024000) AS withdraw_timeout
FROM cg_prize_claim pc;

-- Well-known address ids.
CREATE TEMP TABLE known_aid ON COMMIT DROP AS
SELECT
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.cosmic_game_addr))      AS game_aid,
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.cosmic_signature_addr)) AS nft_aid,
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.cosmic_token_addr))     AS cst_aid,
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.charity_wallet_addr))   AS charity_aid,
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.marketing_wallet_addr)) AS marketing_aid,
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.prizes_wallet_addr))    AS prizes_aid,
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.staking_wallet_cst_addr))   AS staking_cst_aid,
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.staking_wallet_rwalk_addr)) AS staking_rwalk_aid,
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.cosmic_dao_addr))       AS dao_aid,
	(SELECT address_id FROM address a JOIN cg_contracts c ON LOWER(a.addr)=LOWER(c.implementation_addr))   AS impl_aid,
	(SELECT address_id FROM address WHERE addr='0x0000000000000000000000000000000000000000')               AS zero_aid;

-- ===========================================================================
-- SECTION A: Layer-1 integrity (block / transaction / evt_log / address)
-- ===========================================================================

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. layer1','evt_log references missing block','ERROR',COUNT(*),
       LEFT(STRING_AGG('evt_log.id='||v.id,',' ORDER BY v.id),300)
FROM (SELECT e.id FROM evt_log e LEFT JOIN block b ON b.block_num=e.block_num
      WHERE b.block_num IS NULL LIMIT 100) v;

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. layer1','evt_log references missing/mismatched transaction','ERROR',COUNT(*),
       LEFT(STRING_AGG('evt_log.id='||v.id,',' ORDER BY v.id),300)
FROM (SELECT e.id FROM evt_log e LEFT JOIN transaction t ON t.id=e.tx_id
      WHERE t.id IS NULL OR t.block_num <> e.block_num LIMIT 100) v;

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. layer1','transaction references missing block','ERROR',COUNT(*),
       LEFT(STRING_AGG('transaction.id='||v.id,',' ORDER BY v.id),300)
FROM (SELECT t.id FROM transaction t LEFT JOIN block b ON b.block_num=t.block_num
      WHERE b.block_num IS NULL LIMIT 100) v;

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. layer1','block timestamps monotonic (non-decreasing by block_num)','ERROR',COUNT(*),
       LEFT(STRING_AGG('block '||v.block_num,',' ORDER BY v.block_num),300)
FROM (SELECT block_num FROM
        (SELECT block_num, ts, LAG(ts) OVER (ORDER BY block_num) prev_ts FROM block) x
      WHERE prev_ts > ts LIMIT 100) v;

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. layer1','parent_hash chain of consecutively stored blocks','ERROR',COUNT(*),
       LEFT(STRING_AGG('block '||v.block_num,',' ORDER BY v.block_num),300)
FROM (SELECT b.block_num FROM block b JOIN block p ON p.block_num = b.block_num-1
      WHERE b.parent_hash <> p.block_hash LIMIT 100) v;

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. layer1','evt_log.topic0_sig is 8 hex chars','ERROR',COUNT(*),
       LEFT(STRING_AGG('evt_log.id='||v.id,',' ORDER BY v.id),300)
FROM (SELECT id FROM evt_log WHERE topic0_sig !~* '^[0-9a-f]{8}$' LIMIT 100) v;

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'A. layer1','address.addr format 0x + 40 hex chars','ERROR',COUNT(*),
       LEFT(STRING_AGG('address_id='||v.address_id,',' ORDER BY v.address_id),300)
FROM (SELECT address_id FROM address WHERE addr !~* '^0x[0-9a-f]{40}$' LIMIT 100) v;

-- ===========================================================================
-- SECTION B: every event table row must link to a consistent evt_log + block
-- (auto-discovers all cg_*/rw_* tables carrying evtlog_id/block_num/tx_id/time_stamp)
-- ===========================================================================

DO $do$
DECLARE
	t   TEXT;
	bad BIGINT;
	det TEXT;
BEGIN
	FOR t IN
		SELECT c.table_name FROM information_schema.columns c
		WHERE c.table_schema='public' AND c.column_name='evtlog_id'
		  AND (c.table_name LIKE 'cg\_%' OR c.table_name LIKE 'rw\_%')
		  AND EXISTS (SELECT 1 FROM information_schema.columns c2 WHERE c2.table_schema='public'
		              AND c2.table_name=c.table_name AND c2.column_name='block_num')
		  AND EXISTS (SELECT 1 FROM information_schema.columns c2 WHERE c2.table_schema='public'
		              AND c2.table_name=c.table_name AND c2.column_name='tx_id')
		  AND EXISTS (SELECT 1 FROM information_schema.columns c2 WHERE c2.table_schema='public'
		              AND c2.table_name=c.table_name AND c2.column_name='time_stamp')
		  AND EXISTS (SELECT 1 FROM information_schema.columns c2 WHERE c2.table_schema='public'
		              AND c2.table_name=c.table_name AND c2.column_name='id')
		ORDER BY c.table_name
	LOOP
		EXECUTE format(
			'SELECT COUNT(*), LEFT(STRING_AGG(''id=''||x.id::text, '','' ORDER BY x.id),300)
			 FROM (SELECT t.id FROM %I t
			       LEFT JOIN evt_log e ON e.id = t.evtlog_id
			       LEFT JOIN block b ON b.block_num = t.block_num
			       WHERE t.evtlog_id IS NULL OR e.id IS NULL
			          OR e.block_num <> t.block_num OR e.tx_id <> t.tx_id
			          OR b.block_num IS NULL OR b.ts <> t.time_stamp
			       LIMIT 100) x', t)
		INTO bad, det;
		INSERT INTO vr(section,check_name,severity,violations,details)
		VALUES ('B. evt_log linkage', t||': evtlog_id/block_num/tx_id/time_stamp consistent', 'ERROR', bad, det);
	END LOOP;
END $do$;

-- ===========================================================================
-- SECTION C: evt_log <-> decoded-table reconciliation by topic0 signature.
-- Every raw log the indexer stored for a dispatched topic must have been
-- decoded into its cg_* table (and vice versa: no orphan decoded rows).
-- ===========================================================================

CREATE TEMP TABLE tmap (sigs TEXT[], tbls TEXT[], label TEXT, sev TEXT) ON COMMIT DROP;
INSERT INTO tmap VALUES
 (ARRAY['8c551ec2','9314e785'], ARRAY['cg_prize_claim'],          'MainPrizeClaimed (V2+V3)',          'ERROR'),
 (ARRAY['bcb004d6','1d1f406c'], ARRAY['cg_bid'],                  'BidPlaced (V1+V2)',                 'ERROR'),
 (ARRAY['028a5264'],            ARRAY['cg_first_bid'],            'FirstBidPlacedInRound',             'ERROR'),
 (ARRAY['e32cacf2'],            ARRAY['cg_eth_donated'],          'EthDonated',                        'ERROR'),
 (ARRAY['a0804956'],            ARRAY['cg_eth_donated_wi'],       'EthDonatedWithInfo',                'ERROR'),
 (ARRAY['264f630d'],            ARRAY['cg_donation_received'],    'CharityWallet.DonationReceived',    'ERROR'),
 -- topic 1222634b (FundsTransferredToCharity / DonationSent) is handled
 -- contract-scoped below: the game's own emission at claim time is
 -- intentionally NOT decoded by the indexer (registry only maps the charity
 -- and marketing wallet emitters).
 (ARRAY['154fb6c6'],            ARRAY['cg_fund_transf_err'],      'FundTransferFailed',                'ERROR'),
 (ARRAY['a14cfb0f'],            ARRAY['cg_token_name'],           'NftNameChanged',                    'ERROR'),
 (ARRAY['c2115f21'],            ARRAY['cg_mint_event'],           'NftMinted',                         'ERROR'),
 (ARRAY['b12e72ba'],            ARRAY['cg_nft_donation'],         'NftDonated',                        'ERROR'),
 (ARRAY['3f94f617'],            ARRAY['cg_erc20_donation'],       'TokenDonated',                      'ERROR'),
 (ARRAY['af1adae2'],            ARRAY['cg_donated_tok_claimed'],  'DonatedTokenClaimed',               'ERROR'),
 (ARRAY['03c2b6e0'],            ARRAY['cg_donated_nft_claimed'],  'DonatedNftClaimed',                 'ERROR'),
 (ARRAY['8e369548'],            ARRAY['cg_prize_deposit'],        'PrizesWallet.EthReceived',          'ERROR'),
 (ARRAY['172b54ba'],            ARRAY['cg_prize_withdrawal'],     'PrizesWallet.EthWithdrawn',         'ERROR'),
 (ARRAY['9c62e2cb'],            ARRAY['cg_raffle_eth_prize'],     'RaffleWinnerBidderEthPrizeAllocated','ERROR'),
 (ARRAY['27c21fe4'],            ARRAY['cg_raffle_nft_prize'],     'RaffleWinnerPrizePaid',             'ERROR'),
 (ARRAY['838ec9dd'],            ARRAY['cg_endurance_prize'],      'EnduranceChampionPrizePaid',        'ERROR'),
 (ARRAY['3901b643'],            ARRAY['cg_lastcst_prize'],        'LastCstBidderPrizePaid',            'ERROR'),
 (ARRAY['aa858ae2'],            ARRAY['cg_chrono_warrior_prize'], 'ChronoWarriorPrizePaid',            'ERROR'),
 (ARRAY['26726e1a'],            ARRAY['cg_staking_eth_deposit'],  'StakingWallet.EthDepositReceived',  'ERROR'),
 (ARRAY['e2403640'],            ARRAY['cg_mkt_reward'],           'MarketingWallet.RewardPaid',        'ERROR'),
 (ARRAY['e09cd972'],            ARRAY['cg_nft_staked_cst'],       'NftStaked (CS NFT)',                'ERROR'),
 (ARRAY['62773741'],            ARRAY['cg_nft_staked_rwalk'],     'NftStaked (RandomWalk)',            'ERROR'),
 (ARRAY['ec478a78'],            ARRAY['cg_nft_unstaked_cst'],     'NftUnstaked (CS NFT)',              'ERROR'),
 (ARRAY['08e7047c'],            ARRAY['cg_nft_unstaked_rwalk'],   'NftUnstaked (RandomWalk)',          'ERROR'),
 -- cg_st_reward is NOT event-backed: it is an internal accrual table the
 -- staking-deposit trigger populates per (deposit, staked token); validated
 -- in scripts/validate-stats.sql instead.
 (ARRAY['1c7efd98'],            ARRAY['cg_charity_receiver_changed','cg_adm_charity_wallet'],
                                'CharityAddressChanged (shared topic)',                                'WARN'),
 (ARRAY['47870287'],            ARRAY['cg_adm_raf_eth_bidding'],  'NumRaffleEthPrizesForBiddersChanged','ERROR'),
 (ARRAY['85d8bf21'],            ARRAY['cg_adm_raf_nft_bidding'],  'NumRaffleCSNftsForBiddersChanged',  'ERROR'),
 (ARRAY['3312247f'],            ARRAY['cg_adm_raf_nft_staking_rwalk'],'NumRaffleCSNftsForRWStakersChanged','ERROR'),
 (ARRAY['fe65b6d5'],            ARRAY['cg_adm_charity_pcent'],    'CharityEthDonationPctChanged',      'ERROR'),
 (ARRAY['b5a05ec7'],            ARRAY['cg_adm_main_prize_pcent'], 'MainEthPrizePctChanged',            'ERROR'),
 (ARRAY['bfcd8fb9'],            ARRAY['cg_adm_raffle_pcent'],     'RaffleTotalEthPctChanged',          'ERROR'),
 (ARRAY['9e44c04f'],            ARRAY['cg_adm_stake_pcent'],      'StakingTotalEthRewardPctChanged',   'ERROR'),
 (ARRAY['5581e31f'],            ARRAY['cg_adm_chrono_pcent'],     'ChronoWarriorEthPctChanged',        'ERROR'),
 (ARRAY['8717bb19'],            ARRAY['cg_adm_timeout_withdraw'], 'TimeoutDurationToWithdrawPrizesChanged','ERROR'),
 (ARRAY['37a33291'],            ARRAY['cg_adm_timeout_claimprize'],'TimeoutDurationToClaimMainPrizeChanged','ERROR'),
 (ARRAY['f7fce645'],            ARRAY['cg_erc20_transf_err'],     'ERC20TransferFailed',               'ERROR'),
 (ARRAY['dab38e33'],            ARRAY['cg_adm_rwalk_addr'],       'RandomWalkNftAddressChanged',       'ERROR'),
 (ARRAY['b4cecfe1'],            ARRAY['cg_adm_prizes_wallet_addr'],'PrizesWalletAddressChanged',       'ERROR'),
 (ARRAY['4da1815c'],            ARRAY['cg_adm_staking_cst_addr'], 'StakingWalletCsNftAddressChanged',  'ERROR'),
 (ARRAY['bf6e296f'],            ARRAY['cg_adm_staking_rwalk_addr'],'StakingWalletRWalkNftAddressChanged','ERROR'),
 (ARRAY['4d03942c'],            ARRAY['cg_adm_marketing_addr'],   'MarketingWalletAddressChanged',     'ERROR'),
 (ARRAY['df73fc12'],            ARRAY['cg_adm_treasurer_addr'],   'TreasurerAddressChanged',           'ERROR'),
 (ARRAY['9b3eda10'],            ARRAY['cg_adm_costok_addr'],      'CosmicSignatureTokenAddressChanged','ERROR'),
 (ARRAY['5bde6238'],            ARRAY['cg_adm_cossig_addr'],      'CosmicSignatureNftAddressChanged',  'ERROR'),
 (ARRAY['ed46e73b'],            ARRAY['cg_adm_time_inc'],         'MainPrizeTimeIncrementIncreaseDivisorChanged','ERROR'),
 (ARRAY['deb71e1d'],            ARRAY['cg_adm_price_inc'],        'EthBidPriceIncreaseDivisorChanged', 'ERROR'),
 (ARRAY['07417920'],            ARRAY['cg_adm_prize_microsec'],   'MainPrizeTimeIncrementInMicroSecondsChanged','ERROR'),
 (ARRAY['b5edd1f3'],            ARRAY['cg_adm_inisecprize'],      'InitialDurationUntilMainPrizeDivisorChanged','ERROR'),
 (ARRAY['9a2159c1'],            ARRAY['cg_adm_acttime'],          'RoundActivationTimeChanged',        'ERROR'),
 -- Both the divisor-changed and duration-changed variants store into
 -- cg_adm_cst_auclen (storeCstAuctionLengthChange in handlers_admin.go).
 (ARRAY['c95d03f6','4abea08c'], ARRAY['cg_adm_cst_auclen'],       'CstDutchAuctionDuration(Divisor)Changed','ERROR'),
 (ARRAY['acbc6b69'],            ARRAY['cg_adm_cst_auclen_chg_div'],'CstDutchAuctionDurationChangeDivisorChanged','ERROR'),
 (ARRAY['7acba37d'],            ARRAY['cg_adm_late_bid_dur_divisor'],'RoundLateBidDurationDivisorChanged','ERROR'),
 (ARRAY['169f25ec'],            ARRAY['cg_adm_late_bid_premium_base_mul'],'RoundLateBidPremiumBaseMultiplierChanged','ERROR'),
 (ARRAY['cb78cca7'],            ARRAY['cg_adm_late_bid_premium_exponent'],'RoundLateBidPremiumExponentChanged','ERROR'),
 (ARRAY['5a775510'],            ARRAY['cg_adm_cst_price_decline_mul'],'CstBidPriceDeclineMultiplierChanged','ERROR'),
 (ARRAY['dca564ea'],            ARRAY['cg_adm_cst_price_decline_mul_div'],'CstBidPriceDeclineMultiplierChangeDivisorChanged','ERROR'),
 (ARRAY['616bfcaa'],            ARRAY['cg_adm_main_prize_num_nfts'],'MainPrizeNumCosmicSignatureNftsChanged','ERROR'),
 (ARRAY['fdf6043c'],            ARRAY['cg_adm_eth_auclen'],       'EthDutchAuctionDurationDivisorChanged','ERROR'),
 (ARRAY['b6f6af60'],            ARRAY['cg_adm_eth_auc_endprice'], 'EthDutchAuctionEndingBidPriceDivisorChanged','ERROR'),
 -- All three CST-bid-reward event variants store into cg_adm_erc20_reward
 -- (storeCstRewardForBiddingChange in handlers_admin.go).
 (ARRAY['70ad04ce','96978b83','40b9c59a'], ARRAY['cg_adm_erc20_reward'],'BidCstReward(Amount/Multiplier)Changed','ERROR'),
 (ARRAY['d95e7f96'],            ARRAY['cg_adm_erc_rwd_mul'],      'CstPrizeAmountChanged',             'ERROR'),
 (ARRAY['2652e665'],            ARRAY['cg_adm_mkt_reward'],       'MarketingWalletCstContributionAmountChanged','ERROR'),
 (ARRAY['157c413b'],            ARRAY['cg_adm_msg_len'],          'BidMessageLengthMaxLimitChanged',   'ERROR'),
 (ARRAY['27e2bd70'],            ARRAY['cg_adm_script_url'],       'NftGenerationScriptUrlChanged',     'ERROR'),
 (ARRAY['bdfd8152'],            ARRAY['cg_adm_base_uri_cs'],      'BaseUriChanged',                    'ERROR'),
 (ARRAY['4e8c80fe'],            ARRAY['cg_adm_cst_min_limit'],    'CstDutchAuctionBeginningBidPriceMinLimitChanged','ERROR'),
 (ARRAY['b0868a72'],            ARRAY['cg_delay_duration'],       'DelayDurationBeforeRoundActivationChanged','ERROR');

DO $do$
DECLARE
	r        RECORD;
	raw_cnt  BIGINT;
	dec_cnt  BIGINT;
	tot      BIGINT;
	tb       TEXT;
BEGIN
	FOR r IN SELECT * FROM tmap LOOP
		SELECT COUNT(*) INTO raw_cnt FROM evt_log WHERE topic0_sig = ANY (r.sigs);
		dec_cnt := 0;
		FOREACH tb IN ARRAY r.tbls LOOP
			EXECUTE format('SELECT COUNT(*) FROM %I', tb) INTO tot;
			dec_cnt := dec_cnt + tot;
		END LOOP;
		INSERT INTO vr(section,check_name,severity,violations,details)
		VALUES ('C. topic reconciliation',
		        r.label||' ['||ARRAY_TO_STRING(r.sigs,'/')||'] -> '||ARRAY_TO_STRING(r.tbls,'+'),
		        r.sev,
		        CASE WHEN raw_cnt = dec_cnt THEN 0 ELSE 1 END,
		        CASE WHEN raw_cnt = dec_cnt THEN NULL
		             ELSE 'evt_log has '||raw_cnt||' logs but tables hold '||dec_cnt||' rows' END);
	END LOOP;

	-- Topic 1222634b (FundsTransferredToCharity, historically "DonationSent")
	-- is emitted by three contracts but the indexer only decodes two of them:
	--   charity wallet  -> cg_donation_sent
	--   marketing wallet-> cg_funds_to_charity
	--   game (at claim) -> intentionally undecoded (the paired
	--                      CharityWallet.DonationReceived covers it).
	SELECT COUNT(*) INTO raw_cnt FROM evt_log e, known_aid k
	WHERE e.topic0_sig='1222634b' AND e.contract_aid=k.charity_aid;
	SELECT COUNT(*) INTO dec_cnt FROM cg_donation_sent;
	INSERT INTO vr(section,check_name,severity,violations,details)
	VALUES ('C. topic reconciliation','FundsTransferredToCharity [1222634b] from CharityWallet -> cg_donation_sent','ERROR',
	        CASE WHEN raw_cnt=dec_cnt THEN 0 ELSE 1 END,
	        CASE WHEN raw_cnt=dec_cnt THEN NULL ELSE 'evt_log='||raw_cnt||' table='||dec_cnt END);

	SELECT COUNT(*) INTO raw_cnt FROM evt_log e, known_aid k
	WHERE e.topic0_sig='1222634b' AND e.contract_aid=k.marketing_aid;
	SELECT COUNT(*) INTO dec_cnt FROM cg_funds_to_charity;
	INSERT INTO vr(section,check_name,severity,violations,details)
	VALUES ('C. topic reconciliation','FundsTransferredToCharity [1222634b] from MarketingWallet -> cg_funds_to_charity','ERROR',
	        CASE WHEN raw_cnt=dec_cnt THEN 0 ELSE 1 END,
	        CASE WHEN raw_cnt=dec_cnt THEN NULL ELSE 'evt_log='||raw_cnt||' table='||dec_cnt END);

	-- Any other emitter of this topic besides game/charity/marketing is
	-- unexpected and would go undecoded silently.
	SELECT COUNT(*) INTO raw_cnt FROM evt_log e, known_aid k
	WHERE e.topic0_sig='1222634b'
	  AND e.contract_aid NOT IN (k.game_aid, k.charity_aid, k.marketing_aid);
	INSERT INTO vr(section,check_name,severity,violations,details)
	VALUES ('C. topic reconciliation','FundsTransferredToCharity [1222634b] from unexpected emitter','WARN',
	        CASE WHEN raw_cnt=0 THEN 0 ELSE 1 END,
	        CASE WHEN raw_cnt=0 THEN NULL ELSE raw_cnt||' logs from contracts other than game/charity/marketing' END);

	-- ERC1967 Upgraded / AdminChanged: the registry only decodes the game
	-- proxy's emissions; other proxies (if any) are out of scope.
	SELECT COUNT(*) INTO raw_cnt FROM evt_log e, known_aid k
	WHERE e.topic0_sig='bc7cd75a' AND e.contract_aid=k.game_aid;
	SELECT COUNT(*) INTO dec_cnt FROM cg_adm_upgraded;
	INSERT INTO vr(section,check_name,severity,violations,details)
	VALUES ('C. topic reconciliation','Upgraded [bc7cd75a] from game proxy -> cg_adm_upgraded','ERROR',
	        CASE WHEN raw_cnt=dec_cnt THEN 0 ELSE 1 END,
	        CASE WHEN raw_cnt=dec_cnt THEN NULL ELSE 'evt_log='||raw_cnt||' table='||dec_cnt END);

	SELECT COUNT(*) INTO raw_cnt FROM evt_log e, known_aid k
	WHERE e.topic0_sig='7e644d79' AND e.contract_aid=k.game_aid;
	SELECT COUNT(*) INTO dec_cnt FROM cg_adm_admin_changed;
	INSERT INTO vr(section,check_name,severity,violations,details)
	VALUES ('C. topic reconciliation','AdminChanged [7e644d79] from game proxy -> cg_adm_admin_changed','ERROR',
	        CASE WHEN raw_cnt=dec_cnt THEN 0 ELSE 1 END,
	        CASE WHEN raw_cnt=dec_cnt THEN NULL ELSE 'evt_log='||raw_cnt||' table='||dec_cnt END);

	-- OwnershipTransferred: decoded for the nine platform contracts listed in
	-- ownershipSources() (handlers_admin.go); RandomWalk contracts and any
	-- other Ownable emitters are out of scope.
	SELECT COUNT(*) INTO raw_cnt FROM evt_log e, known_aid k
	WHERE e.topic0_sig='8be0079c'
	  AND e.contract_aid IN (k.game_aid, k.nft_aid, k.cst_aid, k.charity_aid,
	                         k.prizes_aid, k.staking_cst_aid, k.staking_rwalk_aid,
	                         k.marketing_aid, k.dao_aid);
	SELECT COUNT(*) INTO dec_cnt FROM cg_adm_ownership;
	INSERT INTO vr(section,check_name,severity,violations,details)
	VALUES ('C. topic reconciliation','OwnershipTransferred [8be0079c] from platform contracts -> cg_adm_ownership','ERROR',
	        CASE WHEN raw_cnt=dec_cnt THEN 0 ELSE 1 END,
	        CASE WHEN raw_cnt=dec_cnt THEN NULL ELSE 'evt_log='||raw_cnt||' table='||dec_cnt END);

	-- Initialized: same source set as OwnershipTransferred plus the
	-- implementation contract (initializedSources() in handlers_admin.go).
	SELECT COUNT(*) INTO raw_cnt FROM evt_log e, known_aid k
	WHERE e.topic0_sig='c7f505b2'
	  AND e.contract_aid IN (k.game_aid, k.nft_aid, k.cst_aid, k.charity_aid,
	                         k.prizes_aid, k.staking_cst_aid, k.staking_rwalk_aid,
	                         k.marketing_aid, k.dao_aid, k.impl_aid);
	SELECT COUNT(*) INTO dec_cnt FROM cg_adm_initialized;
	INSERT INTO vr(section,check_name,severity,violations,details)
	VALUES ('C. topic reconciliation','Initialized [c7f505b2] from platform contracts -> cg_adm_initialized','ERROR',
	        CASE WHEN raw_cnt=dec_cnt THEN 0 ELSE 1 END,
	        CASE WHEN raw_cnt=dec_cnt THEN NULL ELSE 'evt_log='||raw_cnt||' table='||dec_cnt END);

	-- ERC721/ERC20 Transfer share topic0; split by emitting contract.
	SELECT COUNT(*) INTO raw_cnt FROM evt_log e, known_aid k
	WHERE e.topic0_sig='ddf252ad' AND e.contract_aid=k.nft_aid;
	SELECT COUNT(*) INTO dec_cnt FROM cg_erc721_transfer;
	INSERT INTO vr(section,check_name,severity,violations,details)
	VALUES ('C. topic reconciliation','ERC721 Transfer (CosmicSignature NFT) -> cg_erc721_transfer','ERROR',
	        CASE WHEN raw_cnt=dec_cnt THEN 0 ELSE 1 END,
	        CASE WHEN raw_cnt=dec_cnt THEN NULL ELSE 'evt_log='||raw_cnt||' table='||dec_cnt END);

	SELECT COUNT(*) INTO raw_cnt FROM evt_log e, known_aid k
	WHERE e.topic0_sig='ddf252ad' AND e.contract_aid=k.cst_aid;
	SELECT COUNT(*) INTO dec_cnt FROM cg_erc20_transfer;
	INSERT INTO vr(section,check_name,severity,violations,details)
	VALUES ('C. topic reconciliation','ERC20 Transfer (CosmicSignature Token) -> cg_erc20_transfer','ERROR',
	        CASE WHEN raw_cnt=dec_cnt THEN 0 ELSE 1 END,
	        CASE WHEN raw_cnt=dec_cnt THEN NULL ELSE 'evt_log='||raw_cnt||' table='||dec_cnt END);
END $do$;

-- ===========================================================================
-- SECTION D: per-round prize structure (MainPrize._distributePrizes, V2)
-- ===========================================================================

-- D1: rounds are contiguous 0..max with exactly one claim each
-- (claimMainPrize increments roundNum by exactly 1).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','one MainPrizeClaimed per round, rounds contiguous from 0','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.msg,'; '),300)
FROM (
	SELECT 'round '||round_num||' has '||COUNT(*)||' claims' AS msg
	FROM cg_prize_claim GROUP BY round_num HAVING COUNT(*) <> 1
	UNION ALL
	SELECT 'round '||g||' missing'
	FROM generate_series(0,(SELECT MAX(round_num) FROM cg_prize_claim)) g
	WHERE NOT EXISTS (SELECT 1 FROM cg_prize_claim WHERE round_num=g)
) v;

-- D2: singleton prizes per completed round.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','exactly 1 endurance + 1 chrono per round; <=1 lastCst; <=1 staking deposit','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.msg,'; '),300)
FROM (
	SELECT 'round '||c.round_num||': endurance='||
	       (SELECT COUNT(*) FROM cg_endurance_prize e WHERE e.round_num=c.round_num)||
	       ' chrono='||(SELECT COUNT(*) FROM cg_chrono_warrior_prize w WHERE w.round_num=c.round_num)||
	       ' lastcst='||(SELECT COUNT(*) FROM cg_lastcst_prize l WHERE l.round_num=c.round_num)||
	       ' staking='||(SELECT COUNT(*) FROM cg_staking_eth_deposit s WHERE s.round_num=c.round_num) AS msg
	FROM cfg c
	WHERE (SELECT COUNT(*) FROM cg_endurance_prize e WHERE e.round_num=c.round_num) <> 1
	   OR (SELECT COUNT(*) FROM cg_chrono_warrior_prize w WHERE w.round_num=c.round_num) <> 1
	   OR (SELECT COUNT(*) FROM cg_lastcst_prize l WHERE l.round_num=c.round_num) > 1
	   OR (SELECT COUNT(*) FROM cg_staking_eth_deposit s WHERE s.round_num=c.round_num) > 1
) v;

-- D3: raffle ETH allocations: exactly N per round, winner_idx = 0..N-1,
-- all amounts identical (perWinner = floor(raffleTotal/N)) and > 0.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','raffle ETH: N winners, idx 0..N-1, equal positive amounts','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT c.round_num FROM cfg c
	LEFT JOIN (SELECT round_num, COUNT(*) cnt, COUNT(DISTINCT winner_idx) dcnt,
	                  MIN(winner_idx) mn, MAX(winner_idx) mx, MIN(amount) mna, MAX(amount) mxa
	           FROM cg_raffle_eth_prize GROUP BY round_num) r USING (round_num)
	WHERE COALESCE(r.cnt,0) <> c.n_raffle_eth
	   OR r.dcnt <> r.cnt OR r.mn <> 0 OR r.mx <> c.n_raffle_eth-1
	   OR r.mna <> r.mxa OR r.mna <= 0
) v;

-- D4: raffle NFTs: bidders exactly M_b (idx 0..M_b-1); RandomWalk stakers 0 or
-- M_rw (0 when nothing was staked); is_staker must equal is_rwalk in V2.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','raffle NFTs: bidder count = M_b, RW-staker count in {0, M_rw}, idx contiguous, flags coherent','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT c.round_num FROM cfg c
	LEFT JOIN (SELECT round_num, COUNT(*) cnt, COUNT(DISTINCT winner_idx) dcnt, MIN(winner_idx) mn, MAX(winner_idx) mx
	           FROM cg_raffle_nft_prize WHERE NOT is_rwalk GROUP BY round_num) b USING (round_num)
	LEFT JOIN (SELECT round_num, COUNT(*) cnt, COUNT(DISTINCT winner_idx) dcnt, MIN(winner_idx) mn, MAX(winner_idx) mx
	           FROM cg_raffle_nft_prize WHERE is_rwalk GROUP BY round_num) s USING (round_num)
	WHERE COALESCE(b.cnt,0) <> c.n_nft_bidders
	   OR b.dcnt <> b.cnt OR b.mn <> 0 OR b.mx <> c.n_nft_bidders-1
	   OR COALESCE(s.cnt,0) NOT IN (0, c.n_nft_rwalk)
	   OR (s.cnt IS NOT NULL AND (s.dcnt <> s.cnt OR s.mn <> 0 OR s.mx <> c.n_nft_rwalk-1))
	   OR EXISTS (SELECT 1 FROM cg_raffle_nft_prize p
	              WHERE p.round_num=c.round_num AND p.is_staker <> p.is_rwalk)
) v;

-- D5: chrono warrior prizeWinnerIndex equals numRaffleEthPrizesForBidders.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','chrono warrior winner_index = N (PrizesWallet slot after raffle winners)','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (SELECT w.round_num FROM cg_chrono_warrior_prize w JOIN cfg c USING (round_num)
      WHERE w.winner_index <> c.n_raffle_eth) v;

-- D6a: PrizesWallet deposits: N+1 per round, winner_index exactly {0..N}.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','prize deposits: N+1 per round with winner_index 0..N','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT c.round_num FROM cfg c
	LEFT JOIN (SELECT round_num, COUNT(*) cnt, COUNT(DISTINCT winner_index) dcnt,
	                  MIN(winner_index) mn, MAX(winner_index) mx
	           FROM cg_prize_deposit GROUP BY round_num) d USING (round_num)
	WHERE COALESCE(d.cnt,0) <> c.n_raffle_eth+1
	   OR d.dcnt <> d.cnt OR d.mn <> 0 OR d.mx <> c.n_raffle_eth
) v;

-- D6b: each raffle ETH allocation has a matching deposit (same idx/winner/amount).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','raffle ETH allocation matches PrizesWallet deposit (winner+amount)','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num||' idx '||v.winner_idx,',' ),300)
FROM (
	SELECT r.round_num, r.winner_idx FROM cg_raffle_eth_prize r
	LEFT JOIN cg_prize_deposit d ON d.round_num=r.round_num AND d.winner_index=r.winner_idx
	WHERE d.id IS NULL OR d.winner_aid <> r.winner_aid OR d.amount <> r.amount
) v;

-- D6c: chrono ETH matches its deposit at winner_index = N.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','chrono ETH matches PrizesWallet deposit at index N','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT w.round_num FROM cg_chrono_warrior_prize w
	JOIN cfg c USING (round_num)
	LEFT JOIN cg_prize_deposit d ON d.round_num=w.round_num AND d.winner_index=c.n_raffle_eth
	WHERE d.id IS NULL OR d.winner_aid <> w.winner_aid OR d.amount <> w.eth_amount
) v;

-- D7: ETH split cross-ratios. All slices are floor(balance*pct/100) of the SAME
-- balance snapshot, so amountA*pctB and amountB*pctA differ by < pctA+pctB wei.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','ETH split ratio: main prize vs chrono warrior (25:8 by default)','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT c.round_num FROM cfg c JOIN cg_chrono_warrior_prize w USING (round_num)
	WHERE ABS(c.main_eth*c.pct_chrono - w.eth_amount*c.pct_main) > (c.pct_main + c.pct_chrono)
) v;

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','ETH split ratio: main prize vs staking deposit (25:6 by default)','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT c.round_num FROM cfg c JOIN cg_staking_eth_deposit s USING (round_num)
	WHERE ABS(c.main_eth*c.pct_stake - s.deposit_amount*c.pct_main) > (c.pct_main + c.pct_stake)
) v;

INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','ETH split ratio: main prize vs charity donation in claim tx (25:7 by default)','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT c.round_num FROM cfg c
	JOIN cg_donation_received d ON d.tx_id = c.claim_tx
	JOIN known_aid k ON d.donor_aid = k.game_aid
	WHERE ABS(c.main_eth*c.pct_charity - d.amount*c.pct_main) > (c.pct_main + c.pct_charity)
) v;

-- Chrono vs sum(raffle ETH): chrono = floor(bal*8/100), raffleTotal = floor(bal*4/100),
-- deposited raffle sum = raffleTotal - (raffleTotal % N). Range check with dust bounds.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','ETH split ratio: chrono vs sum of raffle ETH (2x by default, +division dust)','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT c.round_num FROM cfg c
	JOIN cg_chrono_warrior_prize w USING (round_num)
	JOIN (SELECT round_num, SUM(amount) s FROM cg_raffle_eth_prize GROUP BY round_num) r USING (round_num)
	WHERE (w.eth_amount*c.pct_raffle - r.s*c.pct_chrono) NOT BETWEEN
	      -(c.pct_raffle + c.pct_chrono)
	      AND (c.n_raffle_eth-1)*c.pct_chrono + c.pct_raffle + c.pct_chrono
) v;

-- D8: CST prize amounts: cstPrizeAmount is one game parameter, so within one
-- round every CST prize (main/endurance/lastCst/chrono/raffle NFT) is identical.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','all CST prize amounts equal within a round and > 0 (cstPrizeAmount)','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT x.round_num FROM (
		SELECT round_num, cst_amount AS amt FROM cg_prize_claim
		UNION ALL SELECT round_num, erc20_amount FROM cg_endurance_prize
		UNION ALL SELECT round_num, erc20_amount FROM cg_lastcst_prize
		UNION ALL SELECT round_num, cst_amount FROM cg_chrono_warrior_prize
		UNION ALL SELECT round_num, cst_amount FROM cg_raffle_nft_prize
	) x GROUP BY x.round_num
	HAVING MIN(x.amt) <> MAX(x.amt) OR MIN(x.amt) <= 0
) v;

-- D9a: 1:1 correspondence between prize-NFT awards and NftMinted rows,
-- and the minted NFT's owner must be the prize winner.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','every prize NFT has a matching mint (same round, owner = winner) and vice versa','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.msg,'; '),300)
FROM (
	SELECT COALESCE('token '||p.token_id||' ('||p.src||')','mint token '||m.token_id)||
	       ' round '||COALESCE(p.round_num,m.round_num) AS msg
	FROM (
		SELECT round_num, token_id, winner_aid, 'main' AS src FROM cg_prize_claim
		UNION ALL SELECT round_num, erc721_token_id, winner_aid, 'endurance' FROM cg_endurance_prize
		UNION ALL SELECT round_num, erc721_token_id, winner_aid, 'lastcst' FROM cg_lastcst_prize
		UNION ALL SELECT round_num, nft_id, winner_aid, 'chrono' FROM cg_chrono_warrior_prize
		UNION ALL SELECT round_num, token_id, winner_aid, 'raffle' FROM cg_raffle_nft_prize
	) p
	FULL JOIN cg_mint_event m ON m.round_num=p.round_num AND m.token_id=p.token_id
	WHERE p.token_id IS NULL OR m.token_id IS NULL OR m.owner_aid <> p.winner_aid
	LIMIT 100
) v;

-- D9b: token ids minted in a round are contiguous (nft.mintMany) and the main
-- prize NFT (MainPrizeClaimed.prizeCosmicSignatureNftId) is the FIRST minted id.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','round mints contiguous; claim token_id = first minted id of round','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT c.round_num FROM cfg c
	JOIN (SELECT round_num, COUNT(*) cnt, MIN(token_id) mn, MAX(token_id) mx
	      FROM cg_mint_event GROUP BY round_num) m USING (round_num)
	WHERE m.mx - m.mn + 1 <> m.cnt OR c.main_token_id <> m.mn
) v;

-- D10: cg_prize registry (populated by triggers) matches the fact tables in
-- both directions. Ptype codes: 0-2 main, 3-4 lastCst, 5-6 endurance,
-- 7-9 chrono, 10 raffle ETH, 11-12 raffle NFT bidder, 13-14 raffle NFT RW
-- staker, 15 staking deposit.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','cg_prize registry matches fact tables (both directions)','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.msg,'; '),300)
FROM (
	SELECT COALESCE('missing (r'||e.round_num||',i'||e.wi||',pt'||e.pt||')',
	                'orphan (r'||p.round_num||',i'||p.winner_index||',pt'||p.ptype||')') AS msg
	FROM (
		SELECT round_num, 0::BIGINT wi, x.pt FROM cg_prize_claim, LATERAL (VALUES (0),(1),(2)) x(pt)
		UNION ALL SELECT round_num, 0, x.pt FROM cg_lastcst_prize, LATERAL (VALUES (3),(4)) x(pt)
		UNION ALL SELECT round_num, 0, x.pt FROM cg_endurance_prize, LATERAL (VALUES (5),(6)) x(pt)
		UNION ALL SELECT round_num, winner_index, x.pt FROM cg_chrono_warrior_prize, LATERAL (VALUES (7),(8),(9)) x(pt)
		UNION ALL SELECT round_num, winner_idx, 10 FROM cg_raffle_eth_prize
		UNION ALL SELECT r.round_num, r.winner_idx, x.pt FROM cg_raffle_nft_prize r, LATERAL (VALUES (11),(12)) x(pt) WHERE NOT r.is_rwalk
		UNION ALL SELECT r.round_num, r.winner_idx, x.pt FROM cg_raffle_nft_prize r, LATERAL (VALUES (13),(14)) x(pt) WHERE r.is_rwalk
		UNION ALL SELECT round_num, 0, 15 FROM cg_staking_eth_deposit
	) e
	FULL JOIN cg_prize p ON p.round_num=e.round_num AND p.winner_index=e.wi AND p.ptype=e.pt
	WHERE e.pt IS NULL OR p.ptype IS NULL
	LIMIT 100
) v;

-- D11: MainPrizeClaimed.timeout = claim block timestamp + timeoutDurationToWithdrawPrizes.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','claim timeout = claim ts + timeoutDurationToWithdrawPrizes (5 weeks default)','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num||' (delta '||v.delta||'s)',',' ),300)
FROM (
	SELECT round_num,
	       claim_timeout - (EXTRACT(EPOCH FROM claim_ts)::BIGINT + withdraw_timeout) AS delta
	FROM cfg
	WHERE claim_timeout <> EXTRACT(EPOCH FROM claim_ts)::BIGINT + withdraw_timeout
) v;

-- D12: charity DonationReceived inside a claim tx carries the claim's round
-- (regression check for the round_num=-1 bug fixed in the indexer).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','charity donation in claim tx has round_num of the claim (round_num=-1 bug)','ERROR',COUNT(*),
       LEFT(STRING_AGG('donation id='||v.id||' round='||v.round_num||' expected '||v.exp,',' ),300)
FROM (
	SELECT d.id, d.round_num, c.round_num AS exp
	FROM cg_donation_received d JOIN cfg c ON c.claim_tx = d.tx_id
	WHERE d.round_num <> c.round_num
) v;

-- D13: every claim tx emitted a charity outcome: DonationReceived or
-- FundsTransferredToCharity or FundTransferFailed.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'D. round structure','claim tx contains a charity transfer event (received/sent/failed)','WARN',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT c.round_num FROM cfg c
	WHERE NOT EXISTS (SELECT 1 FROM cg_donation_received d WHERE d.tx_id=c.claim_tx)
	  AND NOT EXISTS (SELECT 1 FROM cg_funds_to_charity f WHERE f.tx_id=c.claim_tx)
	  AND NOT EXISTS (SELECT 1 FROM cg_fund_transf_err e WHERE e.tx_id=c.claim_tx)
) v;

-- ===========================================================================
-- SECTION E: PrizesWallet withdrawals
-- ===========================================================================

-- E1: PrizesWallet._prepareWithdrawEth deletes the (round, winner) balance
-- and emits EthWithdrawn even when that balance is zero ("It's OK if this is
-- zero"), so repeat withdrawals and withdrawals of rounds one never won
-- legitimately produce amount=0 events. Per (round, winner) group:
--   * at most one withdrawal may carry a non-zero amount,
--   * that non-zero amount must equal the deposited sum,
--   * hence SUM(withdrawn) is either 0 or exactly the deposited sum.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. withdrawals','withdrawn total per (round, winner) is 0 or the exact deposit sum','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num||' winner_aid '||v.winner_aid
                       ||' (withdrawn '||v.wsum||', deposited '||v.dsum||')',',' ORDER BY v.round_num, v.winner_aid),300)
FROM (
	SELECT w.round_num, w.winner_aid, SUM(w.amount) wsum,
	       COALESCE(MAX(d.s),0) dsum
	FROM cg_prize_withdrawal w
	LEFT JOIN (SELECT round_num, winner_aid, SUM(amount) s
	           FROM cg_prize_deposit GROUP BY round_num, winner_aid) d
	       ON d.round_num=w.round_num AND d.winner_aid=w.winner_aid
	GROUP BY w.round_num, w.winner_aid
	HAVING SUM(w.amount) NOT IN (0, COALESCE(MAX(d.s),0))
	    OR COUNT(*) FILTER (WHERE w.amount > 0) > 1
) v;

-- E2: a non-winner beneficiary may withdraw only at/after the round timeout.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. withdrawals','third-party withdrawal only after timeoutTimeToWithdrawSecondaryPrizes','ERROR',COUNT(*),
       LEFT(STRING_AGG('withdrawal id='||v.id,',' ORDER BY v.id),300)
FROM (
	SELECT w.id FROM cg_prize_withdrawal w JOIN cfg c USING (round_num)
	WHERE w.beneficiary_aid <> w.winner_aid
	  AND EXTRACT(EPOCH FROM w.time_stamp)::BIGINT < c.claim_timeout
) v;

-- E3: deposit.claimed flag / withdrawal_id linkage is coherent.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'E. withdrawals','deposit claimed flag consistent with withdrawal records','ERROR',COUNT(*),
       LEFT(STRING_AGG('deposit id='||v.id,',' ORDER BY v.id),300)
FROM (
	SELECT d.id FROM cg_prize_deposit d
	WHERE (d.claimed AND (d.withdrawal_id = 0
	        OR NOT EXISTS (SELECT 1 FROM cg_prize_withdrawal w
	                       WHERE w.evtlog_id=d.withdrawal_id
	                         AND w.round_num=d.round_num AND w.winner_aid=d.winner_aid)))
	   OR (NOT d.claimed AND d.withdrawal_id <> 0)
	   OR (NOT d.claimed AND EXISTS (SELECT 1 FROM cg_prize_withdrawal w
	                                 WHERE w.round_num=d.round_num AND w.winner_aid=d.winner_aid))
) v;

-- ===========================================================================
-- SECTION F: bidding
-- ===========================================================================

-- F1: every completed round has at least one bid (claim reverts otherwise).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'F. bidding','every completed round has >= 1 bid','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (SELECT c.round_num FROM cfg c
      WHERE NOT EXISTS (SELECT 1 FROM cg_bid b WHERE b.round_num=c.round_num)) v;

-- F2: the first bid of a round must be paid in ETH (bid_type 0 or 1, never CST).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'F. bidding','first bid of each round is an ETH (or RandomWalk) bid, not CST','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT DISTINCT b.round_num
	FROM cg_bid b
	JOIN (SELECT round_num, MIN(id) first_id FROM cg_bid GROUP BY round_num) f
	  ON f.round_num=b.round_num AND f.first_id=b.id
	WHERE b.bid_type = 2
) v;

-- F3: exactly one FirstBidPlacedInRound per round that has bids, in the same
-- tx as the actual first bid.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'F. bidding','FirstBidPlacedInRound: one per round with bids, same tx as first bid','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT r.round_num FROM (SELECT DISTINCT round_num FROM cg_bid) r
	LEFT JOIN (SELECT round_num, COUNT(*) cnt, MIN(tx_id) tx FROM cg_first_bid GROUP BY round_num) f USING (round_num)
	LEFT JOIN (SELECT round_num, (ARRAY_AGG(tx_id ORDER BY id))[1] first_tx FROM cg_bid GROUP BY round_num) b USING (round_num)
	WHERE COALESCE(f.cnt,0) <> 1 OR f.tx <> b.first_tx
) v;

-- F4: bid_position values are contiguous within each round.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'F. bidding','bid_position contiguous and unique per round','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num,',' ORDER BY v.round_num),300)
FROM (
	SELECT round_num FROM cg_bid GROUP BY round_num
	HAVING COUNT(DISTINCT bid_position) <> COUNT(*)
	    OR MAX(bid_position) - MIN(bid_position) + 1 <> COUNT(*)
) v;

-- F5: bids of a completed round happened at/before the claim.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'F. bidding','no bid recorded after its round was claimed','ERROR',COUNT(*),
       LEFT(STRING_AGG('bid id='||v.id,',' ORDER BY v.id),300)
FROM (SELECT b.id FROM cg_bid b JOIN cfg c USING (round_num)
      WHERE b.time_stamp > c.claim_ts LIMIT 100) v;

-- F6: bid type / price field coherence
-- (ETH bid: eth_price>0, cst_price=-1; CST bid: cst_price>=0, eth_price=-1;
--  RandomWalk flag <-> rwalk_nft_id).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'F. bidding','bid_type coherent with eth_price/cst_price/rwalk_nft_id','ERROR',COUNT(*),
       LEFT(STRING_AGG('bid id='||v.id,',' ORDER BY v.id),300)
FROM (
	SELECT id FROM cg_bid
	WHERE bid_type NOT IN (0,1,2)
	   OR (bid_type IN (0,1) AND (eth_price <= 0 OR cst_price <> -1))
	   OR (bid_type = 2 AND (cst_price < 0 OR eth_price <> -1))
	   OR (bid_type = 1) <> (rwalk_nft_id > -1)
	LIMIT 100
) v;

-- ===========================================================================
-- SECTION G: NFT / token consistency
-- ===========================================================================

-- G1: minted token ids are unique.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. NFT/token','cg_mint_event token_id unique','ERROR',COUNT(*),
       LEFT(STRING_AGG('token '||v.token_id,',' ORDER BY v.token_id),300)
FROM (SELECT token_id FROM cg_mint_event GROUP BY token_id HAVING COUNT(*) > 1) v;

-- G2: every mint has its ERC721 mint Transfer (otype=1) in the same tx.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. NFT/token','every NftMinted has a mint Transfer (otype=1) in the same tx','ERROR',COUNT(*),
       LEFT(STRING_AGG('token '||v.token_id,',' ORDER BY v.token_id),300)
FROM (
	SELECT m.token_id FROM cg_mint_event m
	WHERE NOT EXISTS (SELECT 1 FROM cg_erc721_transfer t
	                  WHERE t.token_id=m.token_id AND t.tx_id=m.tx_id
	                    AND t.otype=1 AND t.to_aid=m.owner_aid)
) v;

-- G3: cur_owner_aid equals the recipient of the latest transfer of the token.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. NFT/token','cg_mint_event.cur_owner_aid = recipient of latest Transfer','ERROR',COUNT(*),
       LEFT(STRING_AGG('token '||v.token_id,',' ORDER BY v.token_id),300)
FROM (
	SELECT m.token_id FROM cg_mint_event m
	JOIN LATERAL (SELECT t.to_aid FROM cg_erc721_transfer t
	              WHERE t.token_id=m.token_id ORDER BY t.id DESC LIMIT 1) l ON TRUE
	WHERE l.to_aid <> m.cur_owner_aid
) v;

-- G4: denormalized token_name matches the latest NftNameChanged (or '' if none).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. NFT/token','cg_mint_event.token_name = latest cg_token_name entry','ERROR',COUNT(*),
       LEFT(STRING_AGG('token '||v.token_id,',' ORDER BY v.token_id),300)
FROM (
	SELECT m.token_id FROM cg_mint_event m
	LEFT JOIN LATERAL (SELECT n.token_name FROM cg_token_name n
	                   WHERE n.token_id=m.token_id ORDER BY n.id DESC LIMIT 1) l ON TRUE
	WHERE COALESCE(m.token_name,'') <> COALESCE(l.token_name,'')
) v;

-- G5: staking sanity: a token cannot be unstaked more times than staked.
-- (Deeper staking-accrual validation lives in validate-stats.sql.)
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'G. NFT/token','staking: unstake count <= stake count per token (CST + RandomWalk wallets)','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.msg,',' ),300)
FROM (
	SELECT 'cst token '||token_id AS msg FROM
	  (SELECT token_id, COUNT(*) c FROM cg_nft_staked_cst GROUP BY 1) s
	  FULL JOIN (SELECT token_id, COUNT(*) c FROM cg_nft_unstaked_cst GROUP BY 1) u USING (token_id)
	WHERE COALESCE(u.c,0) > COALESCE(s.c,0)
	UNION ALL
	SELECT 'rwalk token '||token_id FROM
	  (SELECT token_id, COUNT(*) c FROM cg_nft_staked_rwalk GROUP BY 1) s
	  FULL JOIN (SELECT token_id, COUNT(*) c FROM cg_nft_unstaked_rwalk GROUP BY 1) u USING (token_id)
	WHERE COALESCE(u.c,0) > COALESCE(s.c,0)
) v;

-- (cg_st_reward accrual validation moved to validate-stats.sql.)

-- ===========================================================================
-- SECTION H: donations (ERC20 / NFT / claims)
-- ===========================================================================

-- H1: donation round numbers must reference an existing (or currently open) round.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'H. donations','donation round_num within [0, current round]','ERROR',COUNT(*),
       LEFT(STRING_AGG(v.msg,',' ),300)
FROM (
	SELECT t||' id='||id AS msg FROM (
		SELECT 'cg_eth_donated' t, id, round_num FROM cg_eth_donated
		UNION ALL SELECT 'cg_eth_donated_wi', id, round_num FROM cg_eth_donated_wi
		UNION ALL SELECT 'cg_erc20_donation', id, round_num FROM cg_erc20_donation
		UNION ALL SELECT 'cg_nft_donation', id, round_num FROM cg_nft_donation
	) x
	WHERE round_num < 0
	   OR round_num > (SELECT COALESCE(MAX(round_num),-1)+1 FROM cg_prize_claim)
) v;

-- H2: DonatedNftClaimed must reference an existing NftDonated (round, idx).
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'H. donations','DonatedNftClaimed references an existing NFT donation (round, idx, token)','ERROR',COUNT(*),
       LEFT(STRING_AGG('claim id='||v.id,',' ORDER BY v.id),300)
FROM (
	SELECT c.id FROM cg_donated_nft_claimed c
	LEFT JOIN cg_nft_donation d ON d.round_num=c.round_num AND d.idx=c.idx
	WHERE d.id IS NULL OR d.token_aid <> c.token_aid OR d.token_id <> c.token_id
) v;

-- H3: DonatedTokenClaimed must reference ERC20 donations of that round/token,
-- claimed amount cannot exceed what was donated.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'H. donations','DonatedTokenClaimed <= donated ERC20 amount per (round, token)','ERROR',COUNT(*),
       LEFT(STRING_AGG('round '||v.round_num||' token_aid '||v.token_aid,',' ),300)
FROM (
	SELECT c.round_num, c.token_aid
	FROM (SELECT round_num, token_aid, SUM(amount) s FROM cg_donated_tok_claimed GROUP BY 1,2) c
	LEFT JOIN (SELECT round_num, token_aid, SUM(amount) s FROM cg_erc20_donation GROUP BY 1,2) d
	       USING (round_num, token_aid)
	WHERE d.s IS NULL OR c.s > d.s
) v;

-- (cg_erc20_donation_stats totals moved to validate-stats.sql.)

-- H4 (WARN): donations tied to a bid should reference an existing cg_bid row.
INSERT INTO vr(section,check_name,severity,violations,details)
SELECT 'H. donations','erc20/NFT donation bid_id references cg_bid','WARN',COUNT(*),
       LEFT(STRING_AGG(v.msg,',' ),300)
FROM (
	SELECT 'erc20 id='||d.id AS msg FROM cg_erc20_donation d
	WHERE d.bid_id > 0 AND NOT EXISTS (SELECT 1 FROM cg_bid b WHERE b.id=d.bid_id)
	UNION ALL
	SELECT 'nft id='||d.id FROM cg_nft_donation d
	WHERE d.bid_id > 0 AND NOT EXISTS (SELECT 1 FROM cg_bid b WHERE b.id=d.bid_id)
) v;

-- ===========================================================================
-- SECTION I (informational, never fails): evt_log decode coverage.
-- Every evt_log row the indexer decoded is referenced by some table's
-- evtlog_id column; whatever is left over was stored but not tracked in the
-- SQL model (some on purpose, e.g. the game's FundsTransferredToCharity and
-- most admin/config events without a cg_adm_* handler).
-- ===========================================================================

-- Event names by topic0 signature (from internal/indexer/{cosmicgame,randomwalk}/topics.go).
CREATE TEMP TABLE topic_names (sig TEXT PRIMARY KEY, name TEXT) ON COMMIT DROP;
INSERT INTO topic_names VALUES
 ('8c551ec2','MainPrizeClaimed (V1/V2)'),
 ('9314e785','MainPrizeClaimed (V3)'),
 ('bcb004d6','BidPlaced (legacy)'),
 ('1d1f406c','BidPlaced (V2)'),
 ('e32cacf2','EthDonated'),
 ('a0804956','EthDonatedWithInfo'),
 ('264f630d','CharityWallet.DonationReceived'),
 ('1222634b','FundsTransferredToCharity / DonationSent'),
 ('1c7efd98','CharityAddressChanged'),
 ('a14cfb0f','NftNameChanged'),
 ('c2115f21','NftMinted'),
 ('b12e72ba','NftDonated'),
 ('3f94f617','TokenDonated'),
 ('af1adae2','DonatedTokenClaimed'),
 ('03c2b6e0','DonatedNftClaimed'),
 ('8e369548','PrizesWallet.EthReceived'),
 ('172b54ba','PrizesWallet.EthWithdrawn'),
 ('9c62e2cb','RaffleWinnerBidderEthPrizeAllocated'),
 ('27c21fe4','RaffleWinnerPrizePaid'),
 ('838ec9dd','EnduranceChampionPrizePaid'),
 ('3901b643','LastCstBidderPrizePaid'),
 ('aa858ae2','ChronoWarriorPrizePaid'),
 ('ddf252ad','Transfer (ERC721/ERC20)'),
 ('e09cd972','NftStaked (CS NFT)'),
 ('62773741','NftStaked (RandomWalk)'),
 ('08e7047c','NftUnstaked (RandomWalk)'),
 ('ec478a78','NftUnstaked (CS NFT)'),
 ('26726e1a','StakingWallet.EthDepositReceived'),
 ('dde81df5','StakingWallet.RewardClaimed'),
 ('154fb6c6','FundTransferFailed'),
 ('f7fce645','ERC20TransferFailed'),
 ('028a5264','FirstBidPlacedInRound'),
 ('bc7cd75a','Upgraded (ERC1967)'),
 ('7e644d79','AdminChanged (ERC1967)'),
 ('df73fc12','TreasurerAddressChanged'),
 ('c7f505b2','Initialized'),
 ('fe65b6d5','CharityEthDonationAmountPercentageChanged'),
 ('b5a05ec7','MainEthPrizeAmountPercentageChanged'),
 ('bfcd8fb9','RaffleTotalEthPrizeAmountForBiddersPercentageChanged'),
 ('9e44c04f','CosmicSignatureNftStakingTotalEthRewardAmountPercentageChanged'),
 ('5581e31f','ChronoWarriorEthPrizeAmountPercentageChanged'),
 ('47870287','NumRaffleEthPrizesForBiddersChanged'),
 ('85d8bf21','NumRaffleCosmicSignatureNftsForBiddersChanged'),
 ('3312247f','NumRaffleCosmicSignatureNftsForRandomWalkNftStakersChanged'),
 ('f24e774c','SystemModeChanged'),
 ('dab38e33','RandomWalkNftAddressChanged'),
 ('b4cecfe1','PrizesWalletAddressChanged'),
 ('4da1815c','StakingWalletCosmicSignatureNftAddressChanged'),
 ('bf6e296f','StakingWalletRandomWalkNftAddressChanged'),
 ('4d03942c','MarketingWalletAddressChanged'),
 ('9b3eda10','CosmicSignatureTokenAddressChanged'),
 ('5bde6238','CosmicSignatureNftAddressChanged'),
 ('77ddb5e9','BusinessLogicContractAddressChanged'),
 ('ed46e73b','MainPrizeTimeIncrementIncreaseDivisorChanged'),
 ('37a33291','TimeoutDurationToClaimMainPrizeChanged'),
 ('8717bb19','TimeoutDurationToWithdrawPrizesChanged'),
 ('deb71e1d','EthBidPriceIncreaseDivisorChanged'),
 ('07417920','MainPrizeTimeIncrementInMicroSecondsChanged'),
 ('b5edd1f3','InitialDurationUntilMainPrizeDivisorChanged'),
 ('9a2159c1','RoundActivationTimeChanged'),
 ('fdf6043c','EthDutchAuctionDurationDivisorChanged'),
 ('c95d03f6','CstDutchAuctionDurationDivisorChanged'),
 ('b6f6af60','EthDutchAuctionEndingBidPriceDivisorChanged'),
 ('e2403640','MarketingWallet.RewardPaid'),
 ('2652e665','MarketingWalletCstContributionAmountChanged'),
 ('70ad04ce','CstRewardAmountForBiddingChanged'),
 ('96978b83','BidCstRewardAmountChanged'),
 ('40b9c59a','BidCstRewardAmountMultiplierChanged'),
 ('4abea08c','CstDutchAuctionDurationChanged'),
 ('acbc6b69','CstDutchAuctionDurationChangeDivisorChanged'),
 ('7acba37d','RoundLateBidDurationDivisorChanged'),
 ('169f25ec','RoundLateBidPricePremiumAmountBaseMultiplierChanged'),
 ('cb78cca7','RoundLateBidPricePremiumAmountExponentChanged'),
 ('5a775510','CstBidPriceDeclineMultiplierChanged'),
 ('dca564ea','CstBidPriceDeclineMultiplierChangeDivisorChanged'),
 ('616bfcaa','MainPrizeNumCosmicSignatureNftsChanged'),
 ('d95e7f96','CstPrizeAmountChanged'),
 ('157c413b','BidMessageLengthMaxLimitChanged'),
 ('27e2bd70','NftGenerationScriptUrlChanged'),
 ('bdfd8152','BaseUriChanged'),
 ('8be0079c','OwnershipTransferred'),
 ('4e8c80fe','CstDutchAuctionBeginningBidPriceMinLimitChanged'),
 ('b0868a72','DelayDurationBeforeRoundActivationChanged'),
 ('8c5be1e5','Approval (ERC20/ERC721)'),
 ('17307eab','ApprovalForAll (ERC721)'),
 ('c565b045','VotingDelaySet (CosmicDAO Governor)'),
 ('7e3f7f07','VotingPeriodSet (CosmicDAO Governor)'),
 ('ccb45da8','ProposalThresholdSet (CosmicDAO Governor)'),
 ('0553476b','QuorumNumeratorUpdated (CosmicDAO Governor)'),
 ('3134e8a2','DelegateChanged (ERC20Votes)'),
 ('dec2bacd','DelegateVotesChanged (ERC20Votes)'),
 ('0a6387c9','EIP712DomainChanged'),
 ('55076e90','RandomWalk NewOffer'),
 ('caacc56f','RandomWalk ItemBought'),
 ('0ff09947','RandomWalk OfferCanceled'),
 ('a11b556a','RandomWalk Withdrawal'),
 ('8ad5e159','RandomWalk TokenNameEvent'),
 ('ad2bc79f','RandomWalk Mint');

-- Collect every evtlog_id referenced by any public table.
CREATE TEMP TABLE tracked_evt (id BIGINT PRIMARY KEY) ON COMMIT DROP;
DO $do$
DECLARE
	t TEXT;
BEGIN
	FOR t IN
		SELECT c.table_name FROM information_schema.columns c
		JOIN information_schema.tables tb
		  ON tb.table_schema=c.table_schema AND tb.table_name=c.table_name
		WHERE c.table_schema='public' AND c.column_name='evtlog_id'
		  AND tb.table_type='BASE TABLE'
	LOOP
		EXECUTE format(
			'INSERT INTO tracked_evt(id)
			 SELECT DISTINCT evtlog_id FROM %I WHERE evtlog_id IS NOT NULL
			 ON CONFLICT DO NOTHING', t);
	END LOOP;
END $do$;

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
\echo '--- Events not tracked by SQL DB (informational) ---------------------------'
\echo 'evt_log rows no table references via evtlog_id; some are intentional'
\echo '(e.g. the game''s FundsTransferredToCharity, admin/config events without'
\echo 'a dedicated handler).'
SELECT e.topic0_sig,
       COALESCE(n.name,'(unknown event)')                       AS event_name,
       COUNT(*) FILTER (WHERE t.id IS NULL)                     AS untracked_logs,
       COUNT(*)                                                 AS total_logs
FROM evt_log e
LEFT JOIN tracked_evt t ON t.id = e.id
LEFT JOIN topic_names n ON n.sig = e.topic0_sig
GROUP BY e.topic0_sig, n.name
HAVING COUNT(*) FILTER (WHERE t.id IS NULL) > 0
ORDER BY untracked_logs DESC, e.topic0_sig;

SELECT COUNT(*) FILTER (WHERE t.id IS NULL) AS untracked_logs_total,
       COUNT(*)                             AS evt_log_total
FROM evt_log e LEFT JOIN tracked_evt t ON t.id = e.id;

\echo ''
\echo '--- Zero-amount EthWithdrawn events (informational) ------------------------'
\echo 'PrizesWallet emits EthWithdrawn even when the (round, winner) balance is'
\echo 'already zero; these rows are legitimate no-op withdrawals.'
SELECT COUNT(*)                                        AS zero_amount_withdrawals,
       (SELECT COUNT(*) FROM cg_prize_withdrawal)      AS withdrawals_total
FROM cg_prize_withdrawal WHERE amount = 0;

SELECT w.id, w.round_num, w.winner_aid, w.beneficiary_aid, w.time_stamp
FROM cg_prize_withdrawal w WHERE w.amount = 0 ORDER BY w.id LIMIT 20;

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
\warn 'VALIDATION FAILED: at least one ERROR-severity check found violations.'
-- Force a non-zero psql exit status for cron/CI usage.
DO $$ BEGIN RAISE EXCEPTION 'database validation failed'; END $$;
\else
\echo 'VALIDATION OK: no ERROR-severity violations found.'
\endif
