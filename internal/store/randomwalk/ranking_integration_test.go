//go:build integration

package randomwalk

import (
	"context"
	"testing"
	"time"
)

func TestExploreRandomTokenIDs(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// All four tokens have one match each, so rating ascending decides:
	// 11 (1189.5), 12 (1195), 13 (1205), 10 (1210.5).
	ids, err := r.ExploreRandomTokenIDs(ctx, aidRandomWalk, 1_000_000, 2)
	if err != nil {
		t.Fatalf("ExploreRandomTokenIDs: %v", err)
	}
	if len(ids) != 2 || ids[0] != 11 || ids[1] != 12 {
		t.Errorf("explore ids: got %v, want [11 12]", ids)
	}
	// limit <= 0 falls back to 2.
	ids, err = r.ExploreRandomTokenIDs(ctx, aidRandomWalk, 1_000_000, 0)
	if err != nil {
		t.Fatalf("ExploreRandomTokenIDs(limit=0): %v", err)
	}
	if len(ids) != 2 {
		t.Errorf("explore ids with limit=0: got %d ids, want 2", len(ids))
	}
	// max_id scopes the candidate set.
	ids, err = r.ExploreRandomTokenIDs(ctx, aidRandomWalk, 11, 10)
	if err != nil {
		t.Fatalf("ExploreRandomTokenIDs(max_id=11): %v", err)
	}
	if len(ids) != 2 || ids[0] != 11 || ids[1] != 10 {
		t.Errorf("explore ids with max_id=11: got %v, want [11 10]", ids)
	}
}

func TestFallbackRandomTokenIDs(t *testing.T) {
	r := repo(t)
	// ORDER BY RANDOM(): assert shape, not order.
	ids, err := r.FallbackRandomTokenIDs(context.Background(), aidRandomWalk, 1_000_000, 3)
	if err != nil {
		t.Fatalf("FallbackRandomTokenIDs: %v", err)
	}
	if len(ids) != 3 {
		t.Fatalf("fallback ids: got %d, want 3", len(ids))
	}
	seen := map[int64]bool{}
	for _, id := range ids {
		if id < 10 || id > 13 {
			t.Errorf("fallback id %d outside minted range [10,13]", id)
		}
		if seen[id] {
			t.Errorf("fallback ids contain duplicate %d", id)
		}
		seen[id] = true
	}
}

func TestHasRankingVoteForVoterPair(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	// alice voted on (10, 11); the pair is unordered.
	for _, pair := range [][2]int64{{10, 11}, {11, 10}} {
		has, err := r.HasRankingVoteForVoterPair(ctx, aidAlice, pair[0], pair[1])
		if err != nil {
			t.Fatalf("HasRankingVoteForVoterPair(%v): %v", pair, err)
		}
		if !has {
			t.Errorf("expected alice's vote on pair %v to be found", pair)
		}
	}
	has, err := r.HasRankingVoteForVoterPair(ctx, aidBob, 10, 11)
	if err != nil {
		t.Fatalf("HasRankingVoteForVoterPair(bob): %v", err)
	}
	if has {
		t.Error("bob never voted; expected no vote found")
	}
	// Invalid inputs short-circuit without touching the database.
	for _, c := range []struct{ voter, nft1, nft2 int64 }{
		{0, 10, 11}, {aidAlice, -1, 11}, {aidAlice, 10, 10},
	} {
		has, err := r.HasRankingVoteForVoterPair(ctx, c.voter, c.nft1, c.nft2)
		if err != nil || has {
			t.Errorf("invalid input %+v: got (%v,%v), want (false,nil)", c, has, err)
		}
	}
}

func TestCountRankingMatches(t *testing.T) {
	r := repo(t)
	n, err := r.CountRankingMatches(context.Background())
	if err != nil {
		t.Fatalf("CountRankingMatches: %v", err)
	}
	if n != 2 {
		t.Errorf("ranking matches: got %d, want 2", n)
	}
}

func TestRatingOrder(t *testing.T) {
	r := repo(t)
	ids, err := r.RatingOrder(context.Background(), aidRandomWalk)
	if err != nil {
		t.Fatalf("RatingOrder: %v", err)
	}
	want := []int64{11, 12, 13, 10}
	if len(ids) != len(want) {
		t.Fatalf("rating order: got %v, want %v", ids, want)
	}
	for i := range want {
		if ids[i] != want[i] {
			t.Fatalf("rating order: got %v, want %v", ids, want)
		}
	}
}

func TestRatingPair(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	r1, r2, err := r.RatingPair(ctx, 10, 11)
	if err != nil {
		t.Fatalf("RatingPair: %v", err)
	}
	if r1 != 1210.5 || r2 != 1189.5 {
		t.Errorf("rating pair (10,11): got (%v,%v), want (1210.5,1189.5)", r1, r2)
	}
	// Unrated tokens default to 1200.
	r1, r2, err = r.RatingPair(ctx, 999, 10)
	if err != nil {
		t.Fatalf("RatingPair(999,10): %v", err)
	}
	if r1 != 1200 || r2 != 1210.5 {
		t.Errorf("rating pair (999,10): got (%v,%v), want (1200,1210.5)", r1, r2)
	}
}

// TestApplyRankingMatchRollback proves the Elo write is transactional:
// a rolled-back match leaves the match count, the ratings and the vote
// lookup exactly as they were.
func TestApplyRankingMatchRollback(t *testing.T) {
	r := repo(t)
	p := pool(t)
	ctx := context.Background()

	countBefore, err := r.CountRankingMatches(ctx)
	if err != nil {
		t.Fatalf("CountRankingMatches: %v", err)
	}

	tx, err := p.Begin(ctx)
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	voter := int64(aidBob)
	if err := ApplyRankingMatch(ctx, tx, 12, 13, true, 1211.0, 1189.0, &voter); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("ApplyRankingMatch: %v", err)
	}
	if err := tx.Rollback(ctx); err != nil {
		t.Fatalf("Rollback: %v", err)
	}

	countAfter, err := r.CountRankingMatches(ctx)
	if err != nil {
		t.Fatalf("CountRankingMatches after rollback: %v", err)
	}
	if countAfter != countBefore {
		t.Errorf("match count changed across rollback: %d -> %d", countBefore, countAfter)
	}
	r1, r2, err := r.RatingPair(ctx, 12, 13)
	if err != nil {
		t.Fatalf("RatingPair after rollback: %v", err)
	}
	if r1 != 1195 || r2 != 1205 {
		t.Errorf("ratings changed across rollback: got (%v,%v), want (1195,1205)", r1, r2)
	}
}

// TestApplyRankingMatchCommit covers the commit path (match insert + both
// rating upserts) and restores the fixture state afterwards.
func TestApplyRankingMatchCommit(t *testing.T) {
	r := repo(t)
	p := pool(t)
	ctx := context.Background()

	t.Cleanup(func() {
		if _, err := p.Exec(ctx, "DELETE FROM rw_ranking_match WHERE voter_aid=$1 AND LEAST(nft1,nft2)=12 AND GREATEST(nft1,nft2)=13", int64(aidBob)); err != nil {
			t.Errorf("removing committed match: %v", err)
		}
		if _, err := p.Exec(ctx, "UPDATE rw_token_ranking SET rating=1195, updated_at=TO_TIMESTAMP(1767232100) WHERE token_id=12"); err != nil {
			t.Errorf("restoring rating of token 12: %v", err)
		}
		if _, err := p.Exec(ctx, "UPDATE rw_token_ranking SET rating=1205, updated_at=TO_TIMESTAMP(1767232100) WHERE token_id=13"); err != nil {
			t.Errorf("restoring rating of token 13: %v", err)
		}
	})

	tx, err := p.Begin(ctx)
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	voter := int64(aidBob)
	if err := ApplyRankingMatch(ctx, tx, 12, 13, true, 1211.0, 1189.0, &voter); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("ApplyRankingMatch: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("Commit: %v", err)
	}

	has, err := r.HasRankingVoteForVoterPair(ctx, aidBob, 13, 12)
	if err != nil || !has {
		t.Errorf("committed vote not visible: has=%v err=%v", has, err)
	}
	r1, r2, err := r.RatingPair(ctx, 12, 13)
	if err != nil {
		t.Fatalf("RatingPair after commit: %v", err)
	}
	if r1 != 1211.0 || r2 != 1189.0 {
		t.Errorf("ratings after commit: got (%v,%v), want (1211,1189)", r1, r2)
	}
}

func TestRankingVoteNonceLifecycle(t *testing.T) {
	r := repo(t)
	p := pool(t)
	ctx := context.Background()

	const nonce = "store-suite-nonce-1"
	if err := r.InsertRankingVoteNonce(ctx, nonce, time.Minute); err != nil {
		t.Fatalf("InsertRankingVoteNonce: %v", err)
	}

	tx, err := p.Begin(ctx)
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	ok, err := ConsumeRankingVoteNonce(ctx, tx, nonce)
	if err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("ConsumeRankingVoteNonce: %v", err)
	}
	if !ok {
		_ = tx.Rollback(ctx)
		t.Fatal("expected fresh nonce to validate")
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("Commit: %v", err)
	}

	// Consumed: second use must fail (replay protection).
	tx, err = p.Begin(ctx)
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	ok, err = ConsumeRankingVoteNonce(ctx, tx, nonce)
	if err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("ConsumeRankingVoteNonce(replay): %v", err)
	}
	_ = tx.Rollback(ctx)
	if ok {
		t.Error("expected consumed nonce to be rejected")
	}

	// Expired nonces are rejected and purged by the next insert.
	if _, err := p.Exec(ctx,
		"INSERT INTO rw_ranking_vote_nonce (nonce, expires_at) VALUES ($1, NOW() - INTERVAL '1 minute')",
		"store-suite-expired",
	); err != nil {
		t.Fatalf("inserting expired nonce: %v", err)
	}
	tx, err = p.Begin(ctx)
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	ok, err = ConsumeRankingVoteNonce(ctx, tx, "store-suite-expired")
	if err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("ConsumeRankingVoteNonce(expired): %v", err)
	}
	_ = tx.Rollback(ctx)
	if ok {
		t.Error("expected expired nonce to be rejected")
	}
	if err := r.InsertRankingVoteNonce(ctx, "store-suite-nonce-2", time.Minute); err != nil {
		t.Fatalf("InsertRankingVoteNonce(purge): %v", err)
	}
	t.Cleanup(func() {
		if _, err := p.Exec(ctx, "DELETE FROM rw_ranking_vote_nonce WHERE nonce LIKE 'store-suite-%'"); err != nil {
			t.Errorf("cleaning up nonces: %v", err)
		}
	})
	var expiredLeft int
	if err := p.QueryRow(ctx, "SELECT COUNT(*) FROM rw_ranking_vote_nonce WHERE nonce='store-suite-expired'").Scan(&expiredLeft); err != nil {
		t.Fatalf("counting expired nonces: %v", err)
	}
	if expiredLeft != 0 {
		t.Error("expected the insert to purge expired nonces")
	}
}
