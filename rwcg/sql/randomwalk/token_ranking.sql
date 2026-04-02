-- RandomWalk token ranking (pairwise "cool or not" style), migrated from Python games/tokens.rating.
-- Apply after randomwalk base tables (needs rw_contracts for default rwalk_aid in app logic).

CREATE TABLE IF NOT EXISTS rw_ranking_match (
	id				BIGSERIAL PRIMARY KEY,
	nft1			BIGINT NOT NULL,
	nft2			BIGINT NOT NULL,
	nft1_won		BOOLEAN NOT NULL,
	voter_aid		BIGINT REFERENCES address(address_id),
	created_at		TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_rw_ranking_match_nft1 ON rw_ranking_match(nft1);
CREATE INDEX IF NOT EXISTS idx_rw_ranking_match_nft2 ON rw_ranking_match(nft2);

CREATE TABLE IF NOT EXISTS rw_ranking_vote_nonce (
	nonce		TEXT PRIMARY KEY,
	expires_at	TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_rw_ranking_vote_nonce_expires ON rw_ranking_vote_nonce(expires_at);

CREATE UNIQUE INDEX IF NOT EXISTS idx_rw_ranking_match_voter_pair
	ON rw_ranking_match (voter_aid, LEAST(nft1, nft2), GREATEST(nft1, nft2))
	WHERE voter_aid IS NOT NULL;

-- Elo-style rating per token (single RandomWalk collection).
CREATE TABLE IF NOT EXISTS rw_token_ranking (
	token_id		BIGINT PRIMARY KEY,
	rating			DOUBLE PRECISION NOT NULL DEFAULT 1200,
	updated_at		TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
