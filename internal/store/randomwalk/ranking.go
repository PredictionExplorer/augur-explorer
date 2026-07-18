// Beauty-contest ranking storage: Elo ratings, pairwise match rows and the
// one-time nonces protecting wallet-signed votes.

package randomwalk

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// ExploreRandomTokenIDs returns up to limit token_ids with fewest ranking
// matches, then lowest rating (GET /api/randomwalk/random), scoped to one
// RandomWalk contract and token_id <= maxID.
func (r *Repo) ExploreRandomTokenIDs(ctx context.Context, rwalkAid, maxID int64, limit int) ([]int64, error) {
	if limit <= 0 {
		limit = 2
	}
	q := `
WITH counts AS (
	SELECT token_id, COUNT(*)::bigint AS cnt FROM (
		SELECT nft1 AS token_id FROM rw_ranking_match
		UNION ALL
		SELECT nft2 AS token_id FROM rw_ranking_match
	) u
	GROUP BY token_id
)
SELECT t.token_id
FROM rw_token t
LEFT JOIN counts c ON c.token_id = t.token_id
LEFT JOIN rw_token_ranking r ON r.token_id = t.token_id
WHERE t.rwalk_aid = $1 AND t.token_id <= $2
ORDER BY COALESCE(c.cnt, 0) ASC, COALESCE(r.rating, 1200) ASC, t.token_id ASC
LIMIT $3`
	return r.scanTokenIDs(ctx, "explore random token ids", q, rwalkAid, maxID, limit)
}

// FallbackRandomTokenIDs picks random minted tokens (used when ranking
// tables are missing or the ranked query fails).
func (r *Repo) FallbackRandomTokenIDs(ctx context.Context, rwalkAid, maxID int64, limit int) ([]int64, error) {
	if limit <= 0 {
		limit = 2
	}
	q := `SELECT token_id FROM rw_token WHERE rwalk_aid = $1 AND token_id <= $2 ORDER BY RANDOM() LIMIT $3`
	return r.scanTokenIDs(ctx, "fallback random token ids", q, rwalkAid, maxID, limit)
}

// scanTokenIDs collects a single-column token_id result. The nil zero slice
// is deliberate: these feed handlers that marshal the value directly, and
// the legacy layer returned nil for an empty result.
func (r *Repo) scanTokenIDs(ctx context.Context, op, query string, args ...any) ([]int64, error) {
	rows, err := r.q(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, store.WrapError(op, err)
	}
	defer rows.Close()
	var out []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, store.WrapError(op, err)
		}
		out = append(out, id)
	}
	if err := rows.Err(); err != nil {
		return nil, store.WrapError(op, err)
	}
	return out, nil
}

// HasRankingVoteForVoterPair reports whether this voter already has a row
// for the unordered pair (nft1, nft2).
func (r *Repo) HasRankingVoteForVoterPair(ctx context.Context, voterAid, nft1, nft2 int64) (bool, error) {
	if voterAid <= 0 || nft1 < 0 || nft2 < 0 || nft1 == nft2 {
		return false, nil
	}
	var one int
	err := r.q(ctx).QueryRow(ctx, `
SELECT 1 FROM rw_ranking_match
WHERE voter_aid = $1
  AND LEAST(nft1, nft2) = LEAST($2::bigint, $3::bigint)
  AND GREATEST(nft1, nft2) = GREATEST($2::bigint, $3::bigint)
LIMIT 1`,
		voterAid, nft1, nft2).Scan(&one)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, store.WrapError("ranking vote lookup for voter pair", err)
	}
	return true, nil
}

// CountRankingMatches returns total pairwise rows (for the Elo K factor).
func (r *Repo) CountRankingMatches(ctx context.Context) (int64, error) {
	var n int64
	err := r.q(ctx).QueryRow(ctx, `SELECT COUNT(*) FROM rw_ranking_match`).Scan(&n)
	if err != nil {
		return 0, store.WrapError("count ranking matches", err)
	}
	return n, nil
}

// RatingOrder returns all token_ids ordered by rating ascending
// (GET /api/randomwalk/rating_order).
func (r *Repo) RatingOrder(ctx context.Context, rwalkAid int64) ([]int64, error) {
	q := `
SELECT t.token_id
FROM rw_token t
LEFT JOIN rw_token_ranking r ON r.token_id = t.token_id
WHERE t.rwalk_aid = $1
ORDER BY COALESCE(r.rating, 1200) ASC, t.token_id ASC`
	return r.scanTokenIDs(ctx, "rating order", q, rwalkAid)
}

// RatingPair loads current ratings for two tokens (defaults 1200 if missing).
func (r *Repo) RatingPair(ctx context.Context, nft1, nft2 int64) (r1, r2 float64, err error) {
	r1, r2 = 1200, 1200
	for _, pair := range []struct {
		id  int64
		dst *float64
	}{{nft1, &r1}, {nft2, &r2}} {
		var rating *float64
		err := r.q(ctx).QueryRow(ctx, `SELECT rating FROM rw_token_ranking WHERE token_id = $1`, pair.id).Scan(&rating)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return 0, 0, store.WrapError("rating pair lookup", err)
		}
		if err == nil && rating != nil {
			*pair.dst = *rating
		}
	}
	return r1, r2, nil
}

// ApplyRankingMatch inserts a match and updates both ratings inside tx.
// voterAid nil => legacy/admin row without wallet voter (voter_aid NULL).
func ApplyRankingMatch(ctx context.Context, tx pgx.Tx, nft1, nft2 int64, nft1Won bool, raNew, rbNew float64, voterAid *int64) error {
	var err error
	if voterAid != nil {
		_, err = tx.Exec(ctx, `INSERT INTO rw_ranking_match (nft1, nft2, nft1_won, voter_aid) VALUES ($1,$2,$3,$4)`, nft1, nft2, nft1Won, *voterAid)
	} else {
		_, err = tx.Exec(ctx, `INSERT INTO rw_ranking_match (nft1, nft2, nft1_won) VALUES ($1,$2,$3)`, nft1, nft2, nft1Won)
	}
	if err != nil {
		return fmt.Errorf("insert match: %w", err)
	}
	_, err = tx.Exec(ctx, `
INSERT INTO rw_token_ranking (token_id, rating, updated_at) VALUES ($1, $2, NOW())
ON CONFLICT (token_id) DO UPDATE SET rating = EXCLUDED.rating, updated_at = NOW()`, nft1, raNew)
	if err != nil {
		return fmt.Errorf("upsert rating nft1: %w", err)
	}
	_, err = tx.Exec(ctx, `
INSERT INTO rw_token_ranking (token_id, rating, updated_at) VALUES ($1, $2, NOW())
ON CONFLICT (token_id) DO UPDATE SET rating = EXCLUDED.rating, updated_at = NOW()`, nft2, rbNew)
	if err != nil {
		return fmt.Errorf("upsert rating nft2: %w", err)
	}
	return nil
}

// InsertRankingVoteNonce stores a one-time nonce for wallet-signed beauty
// votes, purging expired nonces on the way.
func (r *Repo) InsertRankingVoteNonce(ctx context.Context, nonce string, ttl time.Duration) error {
	if ttl <= 0 {
		ttl = 15 * time.Minute
	}
	_, err := r.CreateRankingVoteNonce(ctx, nonce, ttl)
	return err
}

// ConsumeRankingVoteNonce removes nonce inside tx if present and unexpired
// (single row). Returns false if missing or expired.
func ConsumeRankingVoteNonce(ctx context.Context, tx pgx.Tx, nonce string) (bool, error) {
	var got string
	err := tx.QueryRow(ctx,
		`DELETE FROM rw_ranking_vote_nonce WHERE nonce = $1 AND expires_at > NOW() RETURNING nonce`,
		nonce,
	).Scan(&got)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, store.WrapError("consume ranking vote nonce", err)
	}
	return true, nil
}
