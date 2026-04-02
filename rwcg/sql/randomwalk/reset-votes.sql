BEGIN;
TRUNCATE rw_ranking_match RESTART IDENTITY;
TRUNCATE rw_token_ranking;  -- or update ratings as above
-- optional:
TRUNCATE rw_ranking_vote_nonce;
COMMIT;
