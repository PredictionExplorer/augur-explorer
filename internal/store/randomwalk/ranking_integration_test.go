//go:build integration

package randomwalk

import (
	"testing"
	"time"
)

func TestGetExploreRandomTokenIds(t *testing.T) {
	sw := store(t)
	// All four tokens have one match each, so rating ascending decides:
	// 11 (1189.5), 12 (1195), 13 (1205), 10 (1210.5).
	ids, err := sw.Get_explore_random_token_ids(aidRandomWalk, 1_000_000, 2)
	if err != nil {
		t.Fatalf("Get_explore_random_token_ids: %v", err)
	}
	if len(ids) != 2 || ids[0] != 11 || ids[1] != 12 {
		t.Errorf("explore ids: got %v, want [11 12]", ids)
	}
	// limit <= 0 falls back to 2.
	ids, err = sw.Get_explore_random_token_ids(aidRandomWalk, 1_000_000, 0)
	if err != nil {
		t.Fatalf("Get_explore_random_token_ids(limit=0): %v", err)
	}
	if len(ids) != 2 {
		t.Errorf("explore ids with limit=0: got %d ids, want 2", len(ids))
	}
	// max_id scopes the candidate set.
	ids, err = sw.Get_explore_random_token_ids(aidRandomWalk, 11, 10)
	if err != nil {
		t.Fatalf("Get_explore_random_token_ids(max_id=11): %v", err)
	}
	if len(ids) != 2 || ids[0] != 11 || ids[1] != 10 {
		t.Errorf("explore ids with max_id=11: got %v, want [11 10]", ids)
	}
}

func TestGetFallbackRandomTokenIds(t *testing.T) {
	sw := store(t)
	// ORDER BY RANDOM(): assert shape, not order.
	ids, err := sw.Get_fallback_random_token_ids(aidRandomWalk, 1_000_000, 3)
	if err != nil {
		t.Fatalf("Get_fallback_random_token_ids: %v", err)
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
	sw := store(t)
	// alice voted on (10, 11); the pair is unordered.
	for _, pair := range [][2]int64{{10, 11}, {11, 10}} {
		has, err := sw.Has_ranking_vote_for_voter_pair(aidAlice, pair[0], pair[1])
		if err != nil {
			t.Fatalf("Has_ranking_vote_for_voter_pair(%v): %v", pair, err)
		}
		if !has {
			t.Errorf("expected alice's vote on pair %v to be found", pair)
		}
	}
	has, err := sw.Has_ranking_vote_for_voter_pair(aidBob, 10, 11)
	if err != nil {
		t.Fatalf("Has_ranking_vote_for_voter_pair(bob): %v", err)
	}
	if has {
		t.Error("bob never voted; expected no vote found")
	}
	// Invalid inputs short-circuit without touching the database.
	for _, c := range []struct{ voter, nft1, nft2 int64 }{
		{0, 10, 11}, {aidAlice, -1, 11}, {aidAlice, 10, 10},
	} {
		has, err := sw.Has_ranking_vote_for_voter_pair(c.voter, c.nft1, c.nft2)
		if err != nil || has {
			t.Errorf("invalid input %+v: got (%v,%v), want (false,nil)", c, has, err)
		}
	}
}

func TestCountRankingMatches(t *testing.T) {
	sw := store(t)
	n, err := sw.Count_ranking_matches()
	if err != nil {
		t.Fatalf("Count_ranking_matches: %v", err)
	}
	if n != 2 {
		t.Errorf("ranking matches: got %d, want 2", n)
	}
}

func TestGetRatingOrder(t *testing.T) {
	sw := store(t)
	ids, err := sw.Get_rating_order(aidRandomWalk)
	if err != nil {
		t.Fatalf("Get_rating_order: %v", err)
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

func TestGetRatingPair(t *testing.T) {
	sw := store(t)
	r1, r2, err := sw.Get_rating_pair(10, 11)
	if err != nil {
		t.Fatalf("Get_rating_pair: %v", err)
	}
	if r1 != 1210.5 || r2 != 1189.5 {
		t.Errorf("rating pair (10,11): got (%v,%v), want (1210.5,1189.5)", r1, r2)
	}
	// Unrated tokens default to 1200.
	r1, r2, err = sw.Get_rating_pair(999, 10)
	if err != nil {
		t.Fatalf("Get_rating_pair(999,10): %v", err)
	}
	if r1 != 1200 || r2 != 1210.5 {
		t.Errorf("rating pair (999,10): got (%v,%v), want (1200,1210.5)", r1, r2)
	}
}

// TestApplyRankingMatchTxRollback proves the Elo write is transactional:
// a rolled-back match leaves the match count, the ratings and the vote
// lookup exactly as they were.
func TestApplyRankingMatchTxRollback(t *testing.T) {
	sw := store(t)
	db := sw.S.Db()

	countBefore, err := sw.Count_ranking_matches()
	if err != nil {
		t.Fatalf("Count_ranking_matches: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	voter := int64(aidBob)
	if err := Apply_ranking_match_tx(tx, 12, 13, true, 1211.0, 1189.0, &voter); err != nil {
		_ = tx.Rollback()
		t.Fatalf("Apply_ranking_match_tx: %v", err)
	}
	if err := tx.Rollback(); err != nil {
		t.Fatalf("Rollback: %v", err)
	}

	countAfter, err := sw.Count_ranking_matches()
	if err != nil {
		t.Fatalf("Count_ranking_matches after rollback: %v", err)
	}
	if countAfter != countBefore {
		t.Errorf("match count changed across rollback: %d -> %d", countBefore, countAfter)
	}
	r1, r2, err := sw.Get_rating_pair(12, 13)
	if err != nil {
		t.Fatalf("Get_rating_pair after rollback: %v", err)
	}
	if r1 != 1195 || r2 != 1205 {
		t.Errorf("ratings changed across rollback: got (%v,%v), want (1195,1205)", r1, r2)
	}
}

// TestApplyRankingMatchTxCommit covers the commit path (match insert + both
// rating upserts) and restores the fixture state afterwards.
func TestApplyRankingMatchTxCommit(t *testing.T) {
	sw := store(t)
	db := sw.S.Db()

	t.Cleanup(func() {
		if _, err := db.Exec("DELETE FROM rw_ranking_match WHERE voter_aid=$1 AND LEAST(nft1,nft2)=12 AND GREATEST(nft1,nft2)=13", int64(aidBob)); err != nil {
			t.Errorf("removing committed match: %v", err)
		}
		if _, err := db.Exec("UPDATE rw_token_ranking SET rating=1195, updated_at=TO_TIMESTAMP(1767232100) WHERE token_id=12"); err != nil {
			t.Errorf("restoring rating of token 12: %v", err)
		}
		if _, err := db.Exec("UPDATE rw_token_ranking SET rating=1205, updated_at=TO_TIMESTAMP(1767232100) WHERE token_id=13"); err != nil {
			t.Errorf("restoring rating of token 13: %v", err)
		}
	})

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	voter := int64(aidBob)
	if err := Apply_ranking_match_tx(tx, 12, 13, true, 1211.0, 1189.0, &voter); err != nil {
		_ = tx.Rollback()
		t.Fatalf("Apply_ranking_match_tx: %v", err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatalf("Commit: %v", err)
	}

	has, err := sw.Has_ranking_vote_for_voter_pair(aidBob, 13, 12)
	if err != nil || !has {
		t.Errorf("committed vote not visible: has=%v err=%v", has, err)
	}
	r1, r2, err := sw.Get_rating_pair(12, 13)
	if err != nil {
		t.Fatalf("Get_rating_pair after commit: %v", err)
	}
	if r1 != 1211.0 || r2 != 1189.0 {
		t.Errorf("ratings after commit: got (%v,%v), want (1211,1189)", r1, r2)
	}
}

func TestRankingVoteNonceLifecycle(t *testing.T) {
	sw := store(t)
	db := sw.S.Db()

	const nonce = "store-suite-nonce-1"
	if err := sw.Insert_ranking_vote_nonce(nonce, time.Minute); err != nil {
		t.Fatalf("Insert_ranking_vote_nonce: %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	ok, err := sw.Delete_ranking_vote_nonce_if_valid(tx, nonce)
	if err != nil {
		_ = tx.Rollback()
		t.Fatalf("Delete_ranking_vote_nonce_if_valid: %v", err)
	}
	if !ok {
		_ = tx.Rollback()
		t.Fatal("expected fresh nonce to validate")
	}
	if err := tx.Commit(); err != nil {
		t.Fatalf("Commit: %v", err)
	}

	// Consumed: second use must fail (replay protection).
	tx, err = db.Begin()
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	ok, err = sw.Delete_ranking_vote_nonce_if_valid(tx, nonce)
	if err != nil {
		_ = tx.Rollback()
		t.Fatalf("Delete_ranking_vote_nonce_if_valid(replay): %v", err)
	}
	_ = tx.Rollback()
	if ok {
		t.Error("expected consumed nonce to be rejected")
	}

	// Expired nonces are rejected and purged by the next insert.
	if _, err := db.Exec(
		"INSERT INTO rw_ranking_vote_nonce (nonce, expires_at) VALUES ($1, NOW() - INTERVAL '1 minute')",
		"store-suite-expired",
	); err != nil {
		t.Fatalf("inserting expired nonce: %v", err)
	}
	tx, err = db.Begin()
	if err != nil {
		t.Fatalf("Begin: %v", err)
	}
	ok, err = sw.Delete_ranking_vote_nonce_if_valid(tx, "store-suite-expired")
	if err != nil {
		_ = tx.Rollback()
		t.Fatalf("Delete_ranking_vote_nonce_if_valid(expired): %v", err)
	}
	_ = tx.Rollback()
	if ok {
		t.Error("expected expired nonce to be rejected")
	}
	if err := sw.Insert_ranking_vote_nonce("store-suite-nonce-2", time.Minute); err != nil {
		t.Fatalf("Insert_ranking_vote_nonce(purge): %v", err)
	}
	t.Cleanup(func() {
		if _, err := db.Exec("DELETE FROM rw_ranking_vote_nonce WHERE nonce LIKE 'store-suite-%'"); err != nil {
			t.Errorf("cleaning up nonces: %v", err)
		}
	})
	var expiredLeft int
	if err := db.QueryRow("SELECT COUNT(*) FROM rw_ranking_vote_nonce WHERE nonce='store-suite-expired'").Scan(&expiredLeft); err != nil {
		t.Fatalf("counting expired nonces: %v", err)
	}
	if expiredLeft != 0 {
		t.Error("expected the insert to purge expired nonces")
	}
}
