-- Wallet-signed beauty votes: voter address id + one vote per voter per unordered token pair.
-- Apply after token_ranking.sql and address table exist.

ALTER TABLE rw_ranking_match ADD COLUMN IF NOT EXISTS voter_aid BIGINT REFERENCES address(address_id);

CREATE TABLE IF NOT EXISTS rw_ranking_vote_nonce (
	nonce		TEXT PRIMARY KEY,
	expires_at	TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_rw_ranking_vote_nonce_expires ON rw_ranking_vote_nonce(expires_at);

CREATE UNIQUE INDEX IF NOT EXISTS idx_rw_ranking_match_voter_pair
	ON rw_ranking_match (voter_aid, LEAST(nft1, nft2), GREATEST(nft1, nft2))
	WHERE voter_aid IS NOT NULL;
