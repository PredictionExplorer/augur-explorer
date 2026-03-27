package randomwalk

import (
	"database/sql"
	"fmt"
)

// Get_explore_random_token_ids returns up to limit token_ids with fewest ranking matches, then lowest rating
// (parity with legacy Python GET /random), scoped to one RandomWalk contract and token_id <= max_id.
func (sw *SQLStorageWrapper) Get_explore_random_token_ids(rwalk_aid, max_id int64, limit int) ([]int64, error) {
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
	rows, err := sw.S.Db().Query(q, rwalk_aid, max_id, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, rows.Err()
}

// Get_fallback_random_token_ids picks random minted tokens (used when ranking tables are missing or query fails).
func (sw *SQLStorageWrapper) Get_fallback_random_token_ids(rwalk_aid, max_id int64, limit int) ([]int64, error) {
	if limit <= 0 {
		limit = 2
	}
	q := `SELECT token_id FROM rw_token WHERE rwalk_aid = $1 AND token_id <= $2 ORDER BY RANDOM() LIMIT $3`
	rows, err := sw.S.Db().Query(q, rwalk_aid, max_id, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, rows.Err()
}

// Count_ranking_matches returns total pairwise rows (for Elo K factor, legacy parity).
func (sw *SQLStorageWrapper) Count_ranking_matches() (int64, error) {
	var n sql.NullInt64
	err := sw.S.Db().QueryRow(`SELECT COUNT(*) FROM rw_ranking_match`).Scan(&n)
	if err != nil {
		return 0, err
	}
	if n.Valid {
		return n.Int64, nil
	}
	return 0, nil
}

// Get_rating_order returns all token_ids ordered by rating ascending (legacy GET /rating_order).
func (sw *SQLStorageWrapper) Get_rating_order(rwalk_aid int64) ([]int64, error) {
	q := `
SELECT t.token_id
FROM rw_token t
LEFT JOIN rw_token_ranking r ON r.token_id = t.token_id
WHERE t.rwalk_aid = $1
ORDER BY COALESCE(r.rating, 1200) ASC, t.token_id ASC`
	rows, err := sw.S.Db().Query(q, rwalk_aid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out = append(out, id)
	}
	return out, rows.Err()
}

// Get_rating_pair loads current ratings for two tokens (defaults 1200 if missing).
func (sw *SQLStorageWrapper) Get_rating_pair(nft1, nft2 int64) (r1, r2 float64, err error) {
	r1, r2 = 1200, 1200
	var a, b sql.NullFloat64
	err = sw.S.Db().QueryRow(`SELECT rating FROM rw_token_ranking WHERE token_id = $1`, nft1).Scan(&a)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, err
	}
	if err == nil && a.Valid {
		r1 = a.Float64
	}
	err = sw.S.Db().QueryRow(`SELECT rating FROM rw_token_ranking WHERE token_id = $1`, nft2).Scan(&b)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, err
	}
	if err == nil && b.Valid {
		r2 = b.Float64
	}
	return r1, r2, nil
}

// Apply_ranking_match_tx inserts a match and updates ratings inside tx (call after BEGIN).
func Apply_ranking_match_tx(tx *sql.Tx, nft1, nft2 int64, nft1Won bool, raNew, rbNew float64) error {
	_, err := tx.Exec(`INSERT INTO rw_ranking_match (nft1, nft2, nft1_won) VALUES ($1,$2,$3)`, nft1, nft2, nft1Won)
	if err != nil {
		return fmt.Errorf("insert match: %w", err)
	}
	_, err = tx.Exec(`
INSERT INTO rw_token_ranking (token_id, rating, updated_at) VALUES ($1, $2, NOW())
ON CONFLICT (token_id) DO UPDATE SET rating = EXCLUDED.rating, updated_at = NOW()`, nft1, raNew)
	if err != nil {
		return fmt.Errorf("upsert rating nft1: %w", err)
	}
	_, err = tx.Exec(`
INSERT INTO rw_token_ranking (token_id, rating, updated_at) VALUES ($1, $2, NOW())
ON CONFLICT (token_id) DO UPDATE SET rating = EXCLUDED.rating, updated_at = NOW()`, nft2, rbNew)
	if err != nil {
		return fmt.Errorf("upsert rating nft2: %w", err)
	}
	return nil
}
