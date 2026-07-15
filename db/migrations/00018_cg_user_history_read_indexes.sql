-- +goose NO TRANSACTION
-- +goose Up
-- API v2 user histories page one wallet's winnings and donations newest
-- first. The prize union filters each per-type source table by its winner
-- column; the event ledgers key on (wallet, evtlog_id DESC); the donated
-- ERC-20 summaries group claims by (round_num, token_aid).
CREATE INDEX CONCURRENTLY cg_prize_deposit_winner_evt_idx
	ON cg_prize_deposit (winner_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_raffle_nft_prize_winner_evt_idx
	ON cg_raffle_nft_prize (winner_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_raffle_eth_prize_winner_idx
	ON cg_raffle_eth_prize (winner_aid);

CREATE INDEX CONCURRENTLY cg_lastcst_prize_winner_idx
	ON cg_lastcst_prize (winner_aid);

CREATE INDEX CONCURRENTLY cg_endurance_prize_winner_idx
	ON cg_endurance_prize (winner_aid);

CREATE INDEX CONCURRENTLY cg_chrono_warrior_prize_winner_idx
	ON cg_chrono_warrior_prize (winner_aid);

CREATE INDEX CONCURRENTLY cg_eth_donated_donor_evt_idx
	ON cg_eth_donated (donor_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_eth_donated_wi_donor_evt_idx
	ON cg_eth_donated_wi (donor_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_erc20_donation_donor_evt_idx
	ON cg_erc20_donation (donor_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_nft_donation_donor_evt_idx
	ON cg_nft_donation (donor_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_donated_nft_claimed_winner_evt_idx
	ON cg_donated_nft_claimed (winner_aid, evtlog_id DESC);

CREATE INDEX CONCURRENTLY cg_donated_nft_claimed_donation_idx
	ON cg_donated_nft_claimed (idx);

CREATE INDEX CONCURRENTLY cg_donated_tok_claimed_winner_idx
	ON cg_donated_tok_claimed (winner_aid);

CREATE INDEX CONCURRENTLY cg_donated_tok_claimed_round_token_idx
	ON cg_donated_tok_claimed (round_num, token_aid);

-- +goose Down
DROP INDEX CONCURRENTLY IF EXISTS cg_donated_tok_claimed_round_token_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_donated_tok_claimed_winner_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_donated_nft_claimed_donation_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_donated_nft_claimed_winner_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_nft_donation_donor_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_erc20_donation_donor_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_eth_donated_wi_donor_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_eth_donated_donor_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_chrono_warrior_prize_winner_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_endurance_prize_winner_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_lastcst_prize_winner_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_raffle_eth_prize_winner_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_raffle_nft_prize_winner_evt_idx;
DROP INDEX CONCURRENTLY IF EXISTS cg_prize_deposit_winner_evt_idx;
