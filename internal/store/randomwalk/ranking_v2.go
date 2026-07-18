// V2 beauty-contest ranking reads and transactional writes: the bounded
// rating directory, the statistics snapshot, and the two Elo-recording
// mutations the v2 API performs as single Repo calls (the v1 handlers keep
// their own transaction seams over the shared ApplyRankingMatch /
// ConsumeRankingVoteNonce primitives).

package randomwalk

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// ErrRankingNonceInvalid reports that a wallet-signed vote presented a
// nonce that does not exist, was already consumed, or has expired.
var ErrRankingNonceInvalid = errors.New("ranking vote nonce is invalid or expired")

// RankingRatingRecord is one token's beauty-contest standing: its Elo
// rating (the 1200 default when never voted on) and how many recorded
// matches it appeared in.
type RankingRatingRecord struct {
	TokenID    int64
	Rating     float64
	MatchCount int64
}

// RankingRatingPageCursor is the exclusive keyset position of a rating
// directory page: the last returned row's (rating, tokenID) under the
// ascending (rating, token_id) order.
type RankingRatingPageCursor struct {
	Rating  float64
	TokenID int64
}

func (c *RankingRatingPageCursor) valid() bool {
	if c == nil {
		return true
	}
	return c.TokenID >= 0 && !math.IsNaN(c.Rating) && !math.IsInf(c.Rating, 0)
}

// RankingRatingsPage returns at most limit collection tokens ordered by
// Elo rating ascending then token id, resuming strictly after the cursor.
// Every minted token appears: tokens without a recorded rating carry the
// 1200 default and a zero match count.
func (r *Repo) RankingRatingsPage(
	ctx context.Context,
	rwalkAid int64,
	after *RankingRatingPageCursor,
	limit int,
) (records []RankingRatingRecord, hasMore bool, err error) {
	const op = "ranking ratings page"
	if limit <= 0 {
		return nil, false, fmt.Errorf("%s: invalid limit", op)
	}
	if !after.valid() {
		return nil, false, fmt.Errorf("%s: invalid cursor", op)
	}
	query := `
WITH counts AS (
	SELECT token_id, COUNT(*)::bigint AS cnt FROM (
		SELECT nft1 AS token_id FROM rw_ranking_match
		UNION ALL
		SELECT nft2 AS token_id FROM rw_ranking_match
	) u
	GROUP BY token_id
)
SELECT
	t.token_id,
	COALESCE(r.rating, 1200)::double precision,
	COALESCE(c.cnt, 0)
FROM rw_token t
LEFT JOIN rw_token_ranking r ON r.token_id = t.token_id
LEFT JOIN counts c ON c.token_id = t.token_id
WHERE t.rwalk_aid = $1`
	args := []any{rwalkAid}
	if after != nil {
		args = append(args, after.Rating, after.TokenID)
		query += fmt.Sprintf(
			" AND (COALESCE(r.rating, 1200)::double precision, t.token_id) > ($%d::double precision, $%d::bigint)",
			len(args)-1, len(args))
	}
	args = append(args, limit+1)
	query += fmt.Sprintf(`
ORDER BY COALESCE(r.rating, 1200) ASC, t.token_id ASC
LIMIT $%d`, len(args))
	scan := func(rows pgx.Rows, record *RankingRatingRecord) error {
		return rows.Scan(&record.TokenID, &record.Rating, &record.MatchCount)
	}
	records, err = queryList(ctx, r, op, limit+1, query, scan, args...)
	if err != nil {
		return nil, false, err
	}
	records, hasMore = truncatePage(records, limit)
	return records, hasMore, nil
}

// RankingStatisticsRecord is one consistent snapshot of the beauty
// contest's counters.
type RankingStatisticsRecord struct {
	TotalVotes     int64
	WalletVotes    int64
	DistinctVoters int64
	RatedTokens    int64
}

// RankingStatistics computes every beauty-contest counter in one query so
// the values are mutually consistent.
func (r *Repo) RankingStatistics(ctx context.Context) (RankingStatisticsRecord, error) {
	const op = "ranking statistics"
	var rec RankingStatisticsRecord
	err := r.q(ctx).QueryRow(ctx, `
SELECT
	(SELECT COUNT(*) FROM rw_ranking_match),
	(SELECT COUNT(*) FROM rw_ranking_match WHERE voter_aid IS NOT NULL),
	(SELECT COUNT(DISTINCT voter_aid) FROM rw_ranking_match WHERE voter_aid IS NOT NULL),
	(SELECT COUNT(*) FROM rw_token_ranking)`).
		Scan(&rec.TotalVotes, &rec.WalletVotes, &rec.DistinctVoters, &rec.RatedTokens)
	if err != nil {
		return RankingStatisticsRecord{}, store.WrapError(op, err)
	}
	return rec, nil
}

// EnsureVoterAddress resolves (or creates) the address row for a recovered
// vote signer, returning its address_id. Voter rows created here carry no
// chain provenance (block and transaction zero): the wallet may never have
// transacted on-chain with the collection.
func (r *Repo) EnsureVoterAddress(ctx context.Context, addr string) (int64, error) {
	return r.addrID(ctx, addr, 0, 0)
}

// CreateRankingVoteNonce stores a one-time nonce valid for ttl and returns
// the expiry instant, purging already-expired nonces on the way. The expiry
// is computed by the database clock — the same clock ConsumeRankingVoteNonce
// compares against — so validity can never drift with the process clock.
func (r *Repo) CreateRankingVoteNonce(ctx context.Context, nonce string, ttl time.Duration) (time.Time, error) {
	if ttl <= 0 {
		return time.Time{}, errors.New("create ranking vote nonce: ttl must be positive")
	}
	if _, err := r.q(ctx).Exec(ctx, `DELETE FROM rw_ranking_vote_nonce WHERE expires_at < NOW()`); err != nil {
		return time.Time{}, store.WrapError("purge expired ranking vote nonces", err)
	}
	var expiresAt time.Time
	err := r.q(ctx).QueryRow(ctx,
		`INSERT INTO rw_ranking_vote_nonce (nonce, expires_at) VALUES ($1, NOW() + $2) RETURNING expires_at`,
		nonce, ttl).Scan(&expiresAt)
	if err != nil {
		return time.Time{}, store.WrapError("insert ranking vote nonce", err)
	}
	return expiresAt.UTC(), nil
}

// RecordRankingMatch inserts one match row without a voter of record and
// applies both new ratings in a single transaction (the v2 admin
// operation).
func (r *Repo) RecordRankingMatch(
	ctx context.Context,
	first, second int64,
	firstWon bool,
	newFirst, newSecond float64,
) error {
	const op = "record ranking match"
	return r.inRankingTransaction(ctx, op, func(tx pgx.Tx) error {
		return ApplyRankingMatch(ctx, tx, first, second, firstWon, newFirst, newSecond, nil)
	})
}

// RecordSignedRankingVote consumes the one-time nonce, inserts the voter's
// match row and applies both new ratings in a single transaction, so a
// failed insert also returns the nonce unconsumed. A missing or expired
// nonce fails with ErrRankingNonceInvalid; a repeated (voter, pair) vote
// fails with store.ErrConflict.
func (r *Repo) RecordSignedRankingVote(
	ctx context.Context,
	nonce string,
	first, second int64,
	firstWon bool,
	newFirst, newSecond float64,
	voterAid int64,
) error {
	const op = "record signed ranking vote"
	return r.inRankingTransaction(ctx, op, func(tx pgx.Tx) error {
		ok, err := ConsumeRankingVoteNonce(ctx, tx, nonce)
		if err != nil {
			return err
		}
		if !ok {
			return ErrRankingNonceInvalid
		}
		return ApplyRankingMatch(ctx, tx, first, second, firstWon, newFirst, newSecond, &voterAid)
	})
}

// inRankingTransaction runs fn inside one transaction and wraps failures
// with op context and the store sentinels (unique violations surface as
// store.ErrConflict; the ranking sentinels pass through unwrapped).
func (r *Repo) inRankingTransaction(ctx context.Context, op string, fn func(pgx.Tx) error) error {
	tx, err := r.pool().Begin(ctx)
	if err != nil {
		return store.WrapError(op, err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if err := fn(tx); err != nil {
		if errors.Is(err, ErrRankingNonceInvalid) {
			return err
		}
		return store.WrapError(op, err)
	}
	return store.WrapError(op, tx.Commit(ctx))
}
