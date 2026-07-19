//go:build integration

package randomwalk

// Integration coverage for the v2 ranking store surface: the rating
// directory keyset against the frozen v1 RatingOrder, the one-snapshot
// statistics, and the two transactional write methods with exact fixture
// restoration.

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/store"
)

// fixtureRatings is the seeded rating directory in ascending
// (rating, token_id) order: every minted token appears in exactly one
// ranking match.
func fixtureRatings() []RankingRatingRecord {
	return []RankingRatingRecord{
		{TokenID: 11, Rating: 1189.5, MatchCount: 1},
		{TokenID: 12, Rating: 1195, MatchCount: 1},
		{TokenID: 13, Rating: 1205, MatchCount: 1},
		{TokenID: 10, Rating: 1210.5, MatchCount: 1},
	}
}

func TestRankingRatingsPage(t *testing.T) {
	r := repo(t)
	ctx := context.Background()
	want := fixtureRatings()

	t.Run("single page holds the whole directory", func(t *testing.T) {
		records, hasMore, err := r.RankingRatingsPage(ctx, aidRandomWalk, nil, 10)
		if err != nil {
			t.Fatalf("RankingRatingsPage: %v", err)
		}
		if hasMore {
			t.Error("hasMore = true for the complete directory")
		}
		if len(records) != len(want) {
			t.Fatalf("records = %+v, want %+v", records, want)
		}
		for i := range want {
			if records[i] != want[i] {
				t.Fatalf("record[%d] = %+v, want %+v", i, records[i], want[i])
			}
		}
	})

	t.Run("page walks agree with the frozen v1 order at every size", func(t *testing.T) {
		v1Order, err := r.RatingOrder(ctx, aidRandomWalk)
		if err != nil {
			t.Fatalf("RatingOrder: %v", err)
		}
		for pageSize := 1; pageSize <= len(v1Order)+1; pageSize++ {
			var walked []int64
			var cursor *RankingRatingPageCursor
			for {
				records, hasMore, err := r.RankingRatingsPage(ctx, aidRandomWalk, cursor, pageSize)
				if err != nil {
					t.Fatalf("page size %d: %v", pageSize, err)
				}
				if len(records) > pageSize {
					t.Fatalf("page size %d returned %d records", pageSize, len(records))
				}
				for _, record := range records {
					walked = append(walked, record.TokenID)
				}
				if !hasMore {
					break
				}
				last := records[len(records)-1]
				cursor = &RankingRatingPageCursor{Rating: last.Rating, TokenID: last.TokenID}
			}
			if len(walked) != len(v1Order) {
				t.Fatalf("page size %d walked %v, want %v", pageSize, walked, v1Order)
			}
			for i := range v1Order {
				if walked[i] != v1Order[i] {
					t.Fatalf("page size %d walked %v, want %v", pageSize, walked, v1Order)
				}
			}
		}
	})

	t.Run("equal ratings tie-break on token id across page boundaries", func(t *testing.T) {
		p := pool(t)
		// Flatten tokens 12 and 13 onto the 1200 default so the directory
		// holds a genuine tie; restore the fixture rows afterwards.
		if _, err := p.Exec(ctx, "UPDATE rw_token_ranking SET rating=1200 WHERE token_id IN (12, 13)"); err != nil {
			t.Fatalf("flattening ratings: %v", err)
		}
		t.Cleanup(func() {
			if _, err := p.Exec(ctx, "UPDATE rw_token_ranking SET rating=1195, updated_at=TO_TIMESTAMP(1767232100) WHERE token_id=12"); err != nil {
				t.Errorf("restoring rating of token 12: %v", err)
			}
			if _, err := p.Exec(ctx, "UPDATE rw_token_ranking SET rating=1205, updated_at=TO_TIMESTAMP(1767232100) WHERE token_id=13"); err != nil {
				t.Errorf("restoring rating of token 13: %v", err)
			}
		})

		wantOrder := []int64{11, 12, 13, 10}
		var walked []int64
		var cursor *RankingRatingPageCursor
		for {
			records, hasMore, err := r.RankingRatingsPage(ctx, aidRandomWalk, cursor, 1)
			if err != nil {
				t.Fatalf("tie walk: %v", err)
			}
			for _, record := range records {
				walked = append(walked, record.TokenID)
			}
			if !hasMore {
				break
			}
			last := records[len(records)-1]
			cursor = &RankingRatingPageCursor{Rating: last.Rating, TokenID: last.TokenID}
		}
		if len(walked) != len(wantOrder) {
			t.Fatalf("tie walk = %v, want %v", walked, wantOrder)
		}
		for i := range wantOrder {
			if walked[i] != wantOrder[i] {
				t.Fatalf("tie walk = %v, want %v", walked, wantOrder)
			}
		}
	})

	t.Run("foreign contract scope is empty", func(t *testing.T) {
		records, hasMore, err := r.RankingRatingsPage(ctx, aidMarketplace, nil, 10)
		if err != nil {
			t.Fatalf("RankingRatingsPage(foreign): %v", err)
		}
		if len(records) != 0 || hasMore {
			t.Fatalf("foreign scope = %+v hasMore=%v, want empty", records, hasMore)
		}
	})

	t.Run("invalid arguments", func(t *testing.T) {
		if _, _, err := r.RankingRatingsPage(ctx, aidRandomWalk, nil, 0); err == nil {
			t.Error("limit 0 accepted")
		}
		if _, _, err := r.RankingRatingsPage(ctx, aidRandomWalk, &RankingRatingPageCursor{TokenID: -1}, 5); err == nil {
			t.Error("negative cursor token accepted")
		}
	})
}

func TestRankingStatisticsSnapshot(t *testing.T) {
	r := repo(t)
	ctx := context.Background()

	record, err := r.RankingStatistics(ctx)
	if err != nil {
		t.Fatalf("RankingStatistics: %v", err)
	}
	// Fixture truth: alice's wallet vote on (10,11) plus one admin match
	// on (12,13); all four tokens hold rating rows.
	want := RankingStatisticsRecord{
		TotalVotes:     2,
		WalletVotes:    1,
		DistinctVoters: 1,
		RatedTokens:    4,
	}
	if record != want {
		t.Fatalf("statistics = %+v, want %+v", record, want)
	}

	// The snapshot agrees with the standalone count the v1 route serves.
	total, err := r.CountRankingMatches(ctx)
	if err != nil {
		t.Fatalf("CountRankingMatches: %v", err)
	}
	if record.TotalVotes != total {
		t.Fatalf("TotalVotes = %d, want %d", record.TotalVotes, total)
	}
}

func TestRecordRankingMatchTransaction(t *testing.T) {
	r := repo(t)
	p := pool(t)
	ctx := context.Background()

	t.Cleanup(func() {
		if _, err := p.Exec(ctx, "DELETE FROM rw_ranking_match WHERE voter_aid IS NULL AND LEAST(nft1,nft2)=10 AND GREATEST(nft1,nft2)=13"); err != nil {
			t.Errorf("removing recorded match: %v", err)
		}
		if _, err := p.Exec(ctx, "UPDATE rw_token_ranking SET rating=1210.5, updated_at=TO_TIMESTAMP(1767232000) WHERE token_id=10"); err != nil {
			t.Errorf("restoring rating of token 10: %v", err)
		}
		if _, err := p.Exec(ctx, "UPDATE rw_token_ranking SET rating=1205, updated_at=TO_TIMESTAMP(1767232100) WHERE token_id=13"); err != nil {
			t.Errorf("restoring rating of token 13: %v", err)
		}
	})

	countBefore, err := r.CountRankingMatches(ctx)
	if err != nil {
		t.Fatalf("CountRankingMatches: %v", err)
	}
	if err := r.RecordRankingMatch(ctx, 10, 13, false, 1201.25, 1214.25); err != nil {
		t.Fatalf("RecordRankingMatch: %v", err)
	}

	countAfter, err := r.CountRankingMatches(ctx)
	if err != nil {
		t.Fatalf("CountRankingMatches after: %v", err)
	}
	if countAfter != countBefore+1 {
		t.Errorf("match count = %d, want %d", countAfter, countBefore+1)
	}
	first, second, err := r.RatingPair(ctx, 10, 13)
	if err != nil {
		t.Fatalf("RatingPair: %v", err)
	}
	if first != 1201.25 || second != 1214.25 {
		t.Errorf("ratings = (%v, %v), want (1201.25, 1214.25)", first, second)
	}
	// Admin matches carry no voter of record, so the voter-pair lookup
	// stays empty for every wallet.
	var voterRows int
	if err := p.QueryRow(ctx,
		"SELECT COUNT(*) FROM rw_ranking_match WHERE voter_aid IS NOT NULL AND LEAST(nft1,nft2)=10 AND GREATEST(nft1,nft2)=13",
	).Scan(&voterRows); err != nil {
		t.Fatalf("counting voter rows: %v", err)
	}
	if voterRows != 0 {
		t.Errorf("admin match stored a voter of record")
	}
}

func TestRecordSignedRankingVoteTransaction(t *testing.T) {
	r := repo(t)
	p := pool(t)
	ctx := context.Background()

	cleanupVote := func() {
		if _, err := p.Exec(ctx, "DELETE FROM rw_ranking_match WHERE voter_aid=$1 AND LEAST(nft1,nft2)=12 AND GREATEST(nft1,nft2)=13", int64(aidBob)); err != nil {
			t.Errorf("removing recorded vote: %v", err)
		}
		if _, err := p.Exec(ctx, "UPDATE rw_token_ranking SET rating=1195, updated_at=TO_TIMESTAMP(1767232100) WHERE token_id=12"); err != nil {
			t.Errorf("restoring rating of token 12: %v", err)
		}
		if _, err := p.Exec(ctx, "UPDATE rw_token_ranking SET rating=1205, updated_at=TO_TIMESTAMP(1767232100) WHERE token_id=13"); err != nil {
			t.Errorf("restoring rating of token 13: %v", err)
		}
		if _, err := p.Exec(ctx, "DELETE FROM rw_ranking_vote_nonce WHERE nonce LIKE 'v2-suite-%'"); err != nil {
			t.Errorf("cleaning up nonces: %v", err)
		}
	}
	t.Cleanup(cleanupVote)

	t.Run("missing nonce leaves no trace", func(t *testing.T) {
		countBefore, err := r.CountRankingMatches(ctx)
		if err != nil {
			t.Fatalf("CountRankingMatches: %v", err)
		}
		err = r.RecordSignedRankingVote(ctx, "v2-suite-unknown", 12, 13, true, 1211, 1189, aidBob)
		if !errors.Is(err, ErrRankingNonceInvalid) {
			t.Fatalf("error = %v, want ErrRankingNonceInvalid", err)
		}
		countAfter, err := r.CountRankingMatches(ctx)
		if err != nil {
			t.Fatalf("CountRankingMatches after: %v", err)
		}
		if countAfter != countBefore {
			t.Errorf("match count changed on rejected nonce: %d -> %d", countBefore, countAfter)
		}
		first, second, err := r.RatingPair(ctx, 12, 13)
		if err != nil || first != 1195 || second != 1205 {
			t.Errorf("ratings = (%v, %v) err=%v, want fixture (1195, 1205)", first, second, err)
		}
	})

	t.Run("challenge creation returns the database expiry", func(t *testing.T) {
		before := time.Now().UTC()
		expiresAt, err := r.CreateRankingVoteNonce(ctx, "v2-suite-expiry", 15*time.Minute)
		if err != nil {
			t.Fatalf("CreateRankingVoteNonce: %v", err)
		}
		// The expiry is NOW() + ttl on the database clock; allow generous
		// skew between the container and process clocks.
		want := before.Add(15 * time.Minute)
		if drift := expiresAt.Sub(want); drift < -time.Minute || drift > time.Minute {
			t.Fatalf("expiresAt = %v, want about %v", expiresAt, want)
		}
		if _, err := r.CreateRankingVoteNonce(ctx, "v2-suite-bad-ttl", 0); err == nil {
			t.Fatal("CreateRankingVoteNonce accepted a zero ttl")
		}
		// The nonce is a primary key: re-issuing the same value is a
		// conflict, not a silent overwrite of the previous expiry.
		if _, err := r.CreateRankingVoteNonce(ctx, "v2-suite-expiry", 15*time.Minute); !errors.Is(err, store.ErrConflict) {
			t.Fatalf("duplicate nonce error = %v, want store.ErrConflict", err)
		}
	})

	t.Run("expired nonce is rejected", func(t *testing.T) {
		if _, err := p.Exec(ctx,
			"INSERT INTO rw_ranking_vote_nonce (nonce, expires_at) VALUES ($1, NOW() - INTERVAL '1 minute')",
			"v2-suite-expired",
		); err != nil {
			t.Fatalf("inserting expired nonce: %v", err)
		}
		err := r.RecordSignedRankingVote(ctx, "v2-suite-expired", 12, 13, true, 1211, 1189, aidBob)
		if !errors.Is(err, ErrRankingNonceInvalid) {
			t.Fatalf("error = %v, want ErrRankingNonceInvalid", err)
		}
	})

	t.Run("vote consumes the nonce and applies the update", func(t *testing.T) {
		if _, err := r.CreateRankingVoteNonce(ctx, "v2-suite-vote", time.Minute); err != nil {
			t.Fatalf("CreateRankingVoteNonce: %v", err)
		}
		if err := r.RecordSignedRankingVote(ctx, "v2-suite-vote", 12, 13, true, 1211, 1189, aidBob); err != nil {
			t.Fatalf("RecordSignedRankingVote: %v", err)
		}

		has, err := r.HasRankingVoteForVoterPair(ctx, aidBob, 13, 12)
		if err != nil || !has {
			t.Errorf("recorded vote not visible: has=%v err=%v", has, err)
		}
		first, second, err := r.RatingPair(ctx, 12, 13)
		if err != nil || first != 1211 || second != 1189 {
			t.Errorf("ratings = (%v, %v) err=%v, want (1211, 1189)", first, second, err)
		}
		var nonceLeft int
		if err := p.QueryRow(ctx, "SELECT COUNT(*) FROM rw_ranking_vote_nonce WHERE nonce='v2-suite-vote'").Scan(&nonceLeft); err != nil {
			t.Fatalf("counting nonce: %v", err)
		}
		if nonceLeft != 0 {
			t.Error("vote did not consume its nonce")
		}
	})

	t.Run("hidden nonce table surfaces the consume failure", func(t *testing.T) {
		// A real database failure during nonce consumption (not a missing
		// nonce) must abort the vote with the underlying error, never the
		// invalid-nonce sentinel a client could retry around.
		if _, err := p.Exec(ctx, "ALTER TABLE rw_ranking_vote_nonce RENAME TO rw_ranking_vote_nonce_hidden"); err != nil {
			t.Fatalf("hiding nonce table: %v", err)
		}
		restored := false
		restore := func() {
			if restored {
				return
			}
			restored = true
			if _, err := p.Exec(ctx, "ALTER TABLE rw_ranking_vote_nonce_hidden RENAME TO rw_ranking_vote_nonce"); err != nil {
				t.Fatalf("restoring nonce table: %v", err)
			}
		}
		defer restore()

		err := r.RecordSignedRankingVote(ctx, "v2-suite-hidden", 12, 13, true, 1211, 1189, aidBob)
		if err == nil {
			t.Fatal("vote succeeded with the nonce table hidden")
		}
		if errors.Is(err, ErrRankingNonceInvalid) || errors.Is(err, store.ErrConflict) {
			t.Fatalf("database failure misclassified: %v", err)
		}
		restore()
	})

	t.Run("duplicate vote returns conflict and the nonce survives", func(t *testing.T) {
		// The previous subtest recorded bob's vote on (12, 13); a second
		// vote on the reversed pair must fail the unordered uniqueness
		// constraint and roll the nonce consumption back with it.
		if _, err := r.CreateRankingVoteNonce(ctx, "v2-suite-dup", time.Minute); err != nil {
			t.Fatalf("CreateRankingVoteNonce: %v", err)
		}
		err := r.RecordSignedRankingVote(ctx, "v2-suite-dup", 13, 12, true, 1300, 1100, aidBob)
		if !errors.Is(err, store.ErrConflict) {
			t.Fatalf("error = %v, want store.ErrConflict", err)
		}
		var nonceLeft int
		if err := p.QueryRow(ctx, "SELECT COUNT(*) FROM rw_ranking_vote_nonce WHERE nonce='v2-suite-dup'").Scan(&nonceLeft); err != nil {
			t.Fatalf("counting nonce: %v", err)
		}
		if nonceLeft != 1 {
			t.Error("failed vote consumed its nonce; the wallet cannot retry")
		}
		first, second, err := r.RatingPair(ctx, 12, 13)
		if err != nil || first != 1211 || second != 1189 {
			t.Errorf("ratings = (%v, %v) err=%v, want unchanged (1211, 1189)", first, second, err)
		}
	})
}

func TestEnsureVoterAddress(t *testing.T) {
	r := repo(t)
	p := pool(t)
	ctx := context.Background()

	// An existing wallet resolves to its fixture address id.
	aid, err := r.EnsureVoterAddress(ctx, "0x2100000000000000000000000000000000000021")
	if err != nil {
		t.Fatalf("EnsureVoterAddress(alice): %v", err)
	}
	if aid != aidAlice {
		t.Errorf("alice aid = %d, want %d", aid, aidAlice)
	}

	// A brand-new wallet gets a fresh row; the second call reuses it.
	const fresh = "0x00000000000000000000000000000000000B0075"
	t.Cleanup(func() {
		if _, err := p.Exec(ctx, "DELETE FROM address WHERE addr=$1", fresh); err != nil {
			t.Errorf("removing voter address: %v", err)
		}
		sharedStore.ResetAddressCache()
	})
	created, err := r.EnsureVoterAddress(ctx, fresh)
	if err != nil {
		t.Fatalf("EnsureVoterAddress(new): %v", err)
	}
	if created <= 0 {
		t.Fatalf("created aid = %d", created)
	}
	again, err := r.EnsureVoterAddress(ctx, fresh)
	if err != nil {
		t.Fatalf("EnsureVoterAddress(again): %v", err)
	}
	if again != created {
		t.Errorf("repeated resolution = %d, want %d", again, created)
	}
}

// TestRankingV2ErrorPaths pins cancelled-context and closed-pool behavior
// for every new ranking method, mirroring TestErrorPathsConvertedFiles.
func TestRankingV2ErrorPaths(t *testing.T) {
	r := repo(t)

	cancelled, cancel := context.WithCancel(context.Background())
	cancel()

	t.Run("cancelled context", func(t *testing.T) {
		if _, _, err := r.RankingRatingsPage(cancelled, aidRandomWalk, nil, 5); err == nil {
			t.Error("RankingRatingsPage succeeded on a cancelled context")
		}
		if _, err := r.RankingStatistics(cancelled); err == nil {
			t.Error("RankingStatistics succeeded on a cancelled context")
		}
		if _, err := r.CreateRankingVoteNonce(cancelled, "v2-cancelled", time.Minute); err == nil {
			t.Error("CreateRankingVoteNonce succeeded on a cancelled context")
		}
		if err := r.RecordRankingMatch(cancelled, 10, 13, true, 1, 2); err == nil {
			t.Error("RecordRankingMatch succeeded on a cancelled context")
		}
		if err := r.RecordSignedRankingVote(cancelled, "v2-cancelled", 10, 13, true, 1, 2, aidBob); err == nil {
			t.Error("RecordSignedRankingVote succeeded on a cancelled context")
		}
		if _, err := r.EnsureVoterAddress(cancelled, "0x00000000000000000000000000000000000B0076"); err == nil {
			t.Error("EnsureVoterAddress succeeded on a cancelled context")
		}
	})

	t.Run("closed pool", func(t *testing.T) {
		ctx := context.Background()
		spare, err := spareStore(ctx)
		if err != nil {
			t.Fatalf("spareStore: %v", err)
		}
		spare.Close()
		closedRepo := NewRepo(spare)
		if _, _, err := closedRepo.RankingRatingsPage(ctx, aidRandomWalk, nil, 5); err == nil {
			t.Error("RankingRatingsPage succeeded on a closed pool")
		}
		if _, err := closedRepo.RankingStatistics(ctx); err == nil {
			t.Error("RankingStatistics succeeded on a closed pool")
		}
		if _, err := closedRepo.CreateRankingVoteNonce(ctx, "v2-closed", time.Minute); err == nil {
			t.Error("CreateRankingVoteNonce succeeded on a closed pool")
		}
		if err := closedRepo.RecordRankingMatch(ctx, 10, 13, true, 1, 2); err == nil {
			t.Error("RecordRankingMatch succeeded on a closed pool")
		}
		if err := closedRepo.RecordSignedRankingVote(ctx, "v2-closed", 10, 13, true, 1, 2, aidBob); err == nil {
			t.Error("RecordSignedRankingVote succeeded on a closed pool")
		}
	})
}

func BenchmarkRankingQueries(b *testing.B) {
	r := benchRepo(b)
	ctx := context.Background()

	b.Run("ratings_page", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			records, _, err := r.RankingRatingsPage(ctx, aidRandomWalk, nil, 50)
			if err != nil || len(records) == 0 {
				b.Fatalf("ratings page: %v (%d rows)", err, len(records))
			}
		}
	})

	b.Run("statistics", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			record, err := r.RankingStatistics(ctx)
			if err != nil || record.TotalVotes == 0 {
				b.Fatalf("statistics: %v (%+v)", err, record)
			}
		}
	})
}
