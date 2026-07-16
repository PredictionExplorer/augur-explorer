package randomwalk

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/PredictionExplorer/augur-explorer/internal/beautyrank"
)

const rankingTestPrivateKey = "0000000000000000000000000000000000000000000000000000000000000001"

func signBeautyVoteForTest(
	t *testing.T,
	chainID int64,
	nonce string,
	nft1, nft2 int64,
	nft1Won bool,
) (signature, signer string) {
	t.Helper()

	key, err := crypto.HexToECDSA(rankingTestPrivateKey)
	if err != nil {
		t.Fatalf("parse deterministic private key: %v", err)
	}
	winner := nft2
	if nft1Won {
		winner = nft1
	}
	message := beautyrank.VoteMessage(chainID, nonce, nft1, nft2, winner)
	sig, err := crypto.Sign(accounts.TextHash([]byte(message)), key)
	if err != nil {
		t.Fatalf("sign beauty vote: %v", err)
	}
	return hex.EncodeToString(sig), crypto.PubkeyToAddress(key.PublicKey).Hex()
}

// TestExploreMaxTokenIDOption pins the module-side defaulting: malformed
// values are rejected at startup by the typed configuration (config tests),
// so the module only needs to guard the zero value.
func TestExploreMaxTokenIDOption(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		opt  int64
		want int64
	}{
		{name: "zero uses legacy default", opt: 0, want: 3766},
		{name: "negative uses legacy default", opt: -1, want: 3766},
		{name: "positive value", opt: 42, want: 42},
		{name: "maximum int64", opt: math.MaxInt64, want: math.MaxInt64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := New(nil, Options{ExploreMaxTokenID: tt.opt})
			if a.exploreMaxTokenID != tt.want {
				t.Fatalf("exploreMaxTokenID = %d, want %d for option %d", a.exploreMaxTokenID, tt.want, tt.opt)
			}
		})
	}
}

func TestExploreRandomLimit(t *testing.T) {
	tests := []struct {
		raw  string
		want int
	}{
		{raw: "", want: 2},
		{raw: " \t", want: 2},
		{raw: "1", want: 1},
		{raw: " 7 ", want: 7},
		{raw: "+9", want: 9},
		{raw: "100", want: 100},
		{raw: "0", want: 2},
		{raw: "-1", want: 2},
		{raw: "101", want: 2},
		{raw: "1.5", want: 2},
		{raw: "12tokens", want: 2},
		{raw: "999999999999999999999999", want: 2},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q", tt.raw), func(t *testing.T) {
			if got := exploreRandomLimit(tt.raw); got != tt.want {
				t.Fatalf("exploreRandomLimit(%q) = %d, want %d", tt.raw, got, tt.want)
			}
		})
	}
}

// TestRankingVoteChainAllowed pins the allowlist semantics over the typed
// configuration values (RANKING_VOTE_CHAIN_IDS parsing — trimming, blank
// entries, malformed-entry rejection — is the config loader's job now and
// is tested there).
func TestRankingVoteChainAllowed(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		allowed []int64
		chainID int64
		want    bool
	}{
		{name: "positive allowed when unset", allowed: nil, chainID: 1, want: true},
		{name: "positive allowed when empty", allowed: []int64{}, chainID: 42161, want: true},
		{name: "zero always rejected", allowed: nil, chainID: 0, want: false},
		{name: "negative always rejected", allowed: nil, chainID: -1, want: false},
		{name: "first allowlist entry", allowed: []int64{1, 42161, 8453}, chainID: 1, want: true},
		{name: "middle allowlist entry", allowed: []int64{1, 42161, 8453}, chainID: 42161, want: true},
		{name: "last allowlist entry", allowed: []int64{1, 42161, 8453}, chainID: 8453, want: true},
		{name: "not in allowlist", allowed: []int64{1, 42161, 8453}, chainID: 10, want: false},
		{name: "exact match not substring", allowed: []int64{11, 421610}, chainID: 1, want: false},
		{name: "maximum int64", allowed: []int64{math.MaxInt64}, chainID: math.MaxInt64, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := New(nil, Options{VoteChainIDs: tt.allowed})
			if got := a.rankingVoteChainAllowed(tt.chainID); got != tt.want {
				t.Fatalf("rankingVoteChainAllowed(%d) = %v, want %v for %v", tt.chainID, got, tt.want, tt.allowed)
			}
		})
	}
}

func TestPerformRankingMatchRejectsBadPairsBeforeDependencies(t *testing.T) {
	a := NewBare()
	tests := []struct {
		name       string
		nft1, nft2 int64
	}{
		{name: "negative first token", nft1: -1, nft2: 2},
		{name: "negative second token", nft1: 1, nft2: -2},
		{name: "both negative", nft1: -1, nft2: -2},
		{name: "same zero token", nft1: 0, nft2: 0},
		{name: "same positive token", nft1: 7, nft2: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ra, rb, err := a.performRankingMatch(context.Background(), tt.nft1, tt.nft2, true)
			if !errors.Is(err, errRankingBadPair) {
				t.Fatalf("error = %v, want errRankingBadPair", err)
			}
			if ra != 0 || rb != 0 {
				t.Fatalf("ratings = (%v, %v), want (0, 0)", ra, rb)
			}
		})
	}
}

func TestPerformRankingMatchWithDependencies(t *testing.T) {
	errCount := errors.New("count failed")
	errRating := errors.New("rating failed")
	errBegin := errors.New("begin failed")
	errApply := errors.New("apply failed")
	errCommit := errors.New("commit failed")

	tests := []struct {
		name                 string
		nft1Won              bool
		countErr             error
		ratingErr            error
		transactionStartErr  error
		applyErr             error
		transactionFinishErr error
		wantErr              error
		wantSteps            string
	}{
		{name: "count failure", countErr: errCount, wantErr: errCount, wantSteps: "count"},
		{name: "rating failure", ratingErr: errRating, wantErr: errRating, wantSteps: "count,rating"},
		{name: "begin failure", transactionStartErr: errBegin, wantErr: errBegin, wantSteps: "count,rating,transaction"},
		{name: "apply failure", applyErr: errApply, wantErr: errApply, wantSteps: "count,rating,transaction,apply"},
		{name: "commit failure", transactionFinishErr: errCommit, wantErr: errCommit, wantSteps: "count,rating,transaction,apply"},
		{name: "first token wins", nft1Won: true, wantSteps: "count,rating,transaction,apply"},
		{name: "second token wins", nft1Won: false, wantSteps: "count,rating,transaction,apply"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var steps []string
			deps := rankingMatchDependencies{
				countRankingMatches: func(context.Context) (int64, error) {
					steps = append(steps, "count")
					return 17, tt.countErr
				},
				ratingPair: func(_ context.Context, nft1, nft2 int64) (float64, float64, error) {
					steps = append(steps, "rating")
					if nft1 != 7 || nft2 != 9 {
						t.Errorf("ratingPair tokens = (%d, %d), want (7, 9)", nft1, nft2)
					}
					return 1300, 1100, tt.ratingErr
				},
				withTransaction: func(ctx context.Context, fn func(pgx.Tx) error) error {
					steps = append(steps, "transaction")
					if tt.transactionStartErr != nil {
						return tt.transactionStartErr
					}
					if err := fn(nil); err != nil {
						return err
					}
					return tt.transactionFinishErr
				},
				applyRankingMatch: func(
					_ context.Context,
					_ pgx.Tx,
					nft1, nft2 int64,
					nft1Won bool,
					raNew, rbNew float64,
					voterAid *int64,
				) error {
					steps = append(steps, "apply")
					if nft1 != 7 || nft2 != 9 || nft1Won != tt.nft1Won {
						t.Errorf("apply args = (%d, %d, %v), want (7, 9, %v)", nft1, nft2, nft1Won, tt.nft1Won)
					}
					if voterAid != nil {
						t.Errorf("admin ranking match voterAid = %v, want nil", *voterAid)
					}
					wantRA, wantRB := beautyrank.EloUpdate(1300, 1100, map[bool]float64{false: 0, true: 1}[tt.nft1Won], 17)
					if raNew != wantRA || rbNew != wantRB {
						t.Errorf("applied ratings = (%v, %v), want (%v, %v)", raNew, rbNew, wantRA, wantRB)
					}
					return tt.applyErr
				},
			}

			ra, rb, err := performRankingMatchWithDependencies(context.Background(), 7, 9, tt.nft1Won, deps)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("error = %v, want %v", err, tt.wantErr)
				}
				if ra != 0 || rb != 0 {
					t.Fatalf("ratings on failure = (%v, %v), want (0, 0)", ra, rb)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				score := 0.0
				if tt.nft1Won {
					score = 1
				}
				wantRA, wantRB := beautyrank.EloUpdate(1300, 1100, score, 17)
				if ra != wantRA || rb != wantRB {
					t.Fatalf("ratings = (%v, %v), want (%v, %v)", ra, rb, wantRA, wantRB)
				}
				if math.Abs((ra+rb)-2400) > 1e-9 {
					t.Fatalf("rating sum = %v, want 2400", ra+rb)
				}
			}
			if gotSteps := strings.Join(steps, ","); gotSteps != tt.wantSteps {
				t.Fatalf("steps = %q, want %q", gotSteps, tt.wantSteps)
			}
		})
	}
}

func TestPerformSignedBeautyVoteValidatesBeforeDatabase(t *testing.T) {
	t.Parallel()
	a := New(nil, Options{VoteChainIDs: []int64{1}})

	tests := []struct {
		name                  string
		nft1, nft2, chainID   int64
		nonce, signature      string
		wantErr               error
		wantErrorTextContains string
	}{
		{
			name: "negative first token", nft1: -1, nft2: 2, chainID: 0,
			wantErr: errRankingBadPair,
		},
		{
			name: "negative second token", nft1: 1, nft2: -2, chainID: 0,
			wantErr: errRankingBadPair,
		},
		{
			name: "same token", nft1: 3, nft2: 3, chainID: 0,
			wantErr: errRankingBadPair,
		},
		{
			name: "missing nonce", nft1: 1, nft2: 2, chainID: 1, signature: "00",
			wantErr: errRankingVoteCredentialsRequired,
		},
		{
			name: "blank nonce", nft1: 1, nft2: 2, chainID: 1, nonce: " \t", signature: "00",
			wantErr: errRankingVoteCredentialsRequired,
		},
		{
			name: "missing signature", nft1: 1, nft2: 2, chainID: 1, nonce: "nonce",
			wantErr: errRankingVoteCredentialsRequired,
		},
		{
			name: "zero chain", nft1: 1, nft2: 2, chainID: 0, nonce: "nonce", signature: "00",
			wantErr: errRankingVoteChainNotAllowed,
		},
		{
			name: "chain outside allowlist", nft1: 1, nft2: 2, chainID: 42161, nonce: "nonce", signature: "00",
			wantErr: errRankingVoteChainNotAllowed,
		},
		{
			name: "malformed signature", nft1: 1, nft2: 2, chainID: 1, nonce: "nonce", signature: "not-hex",
			wantErr: errRankingVoteInvalidSignature, wantErrorTextContains: "decode signature",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := a.performSignedBeautyVote(
				context.Background(),
				tt.nft1,
				tt.nft2,
				true,
				tt.chainID,
				tt.nonce,
				tt.signature,
			)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
			if tt.wantErrorTextContains != "" && !strings.Contains(err.Error(), tt.wantErrorTextContains) {
				t.Fatalf("error %q does not contain %q", err, tt.wantErrorTextContains)
			}
		})
	}
}

func TestPerformSignedBeautyVoteWithDependencies(t *testing.T) {
	t.Parallel()

	const (
		nft1     = int64(7)
		nft2     = int64(9)
		chainID  = int64(1)
		nonce    = "nonce-123"
		voterAid = int64(77)
	)
	signature, signer := signBeautyVoteForTest(t, chainID, nonce, nft1, nft2, true)

	errCount := errors.New("count failed")
	errRating := errors.New("rating failed")
	errLookup := errors.New("lookup failed")
	errBegin := errors.New("begin failed")
	errConsume := errors.New("consume failed")
	errApply := errors.New("apply failed")
	errCommit := errors.New("commit failed")

	tests := []struct {
		name                 string
		countErr             error
		ratingErr            error
		lookupErr            error
		transactionStartErr  error
		consumeErr           error
		nonceMissing         bool
		applyErr             error
		transactionFinishErr error
		wantErr              error
		wantSteps            string
	}{
		{name: "count failure", countErr: errCount, wantErr: errCount, wantSteps: "count"},
		{name: "rating failure", ratingErr: errRating, wantErr: errRating, wantSteps: "count,rating"},
		{name: "address failure", lookupErr: errLookup, wantErr: errLookup, wantSteps: "count,rating,lookup"},
		{name: "begin failure", transactionStartErr: errBegin, wantErr: errBegin, wantSteps: "count,rating,lookup,transaction"},
		{name: "nonce store failure", consumeErr: errConsume, wantErr: errConsume, wantSteps: "count,rating,lookup,transaction,consume"},
		{name: "replayed or expired nonce", nonceMissing: true, wantErr: errRankingVoteInvalidNonce, wantSteps: "count,rating,lookup,transaction,consume"},
		{
			name: "duplicate voter pair",
			applyErr: &pgconn.PgError{
				Code:    "23505",
				Message: "duplicate key",
			},
			wantErr:   errRankingDuplicateVoterPair,
			wantSteps: "count,rating,lookup,transaction,consume,apply",
		},
		{
			name: "wrapped duplicate voter pair",
			applyErr: fmt.Errorf("insert ranking match: %w", &pgconn.PgError{
				Code:    "23505",
				Message: "duplicate key",
			}),
			wantErr:   errRankingDuplicateVoterPair,
			wantSteps: "count,rating,lookup,transaction,consume,apply",
		},
		{name: "apply failure", applyErr: errApply, wantErr: errApply, wantSteps: "count,rating,lookup,transaction,consume,apply"},
		{name: "commit failure", transactionFinishErr: errCommit, wantErr: errCommit, wantSteps: "count,rating,lookup,transaction,consume,apply"},
		{name: "success", wantSteps: "count,rating,lookup,transaction,consume,apply"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var steps []string
			deps := signedBeautyVoteDependencies{
				rankingMatchDependencies: rankingMatchDependencies{
					countRankingMatches: func(context.Context) (int64, error) {
						steps = append(steps, "count")
						return 0, tt.countErr
					},
					ratingPair: func(_ context.Context, gotNFT1, gotNFT2 int64) (float64, float64, error) {
						steps = append(steps, "rating")
						if gotNFT1 != nft1 || gotNFT2 != nft2 {
							t.Errorf("ratingPair tokens = (%d, %d), want (%d, %d)", gotNFT1, gotNFT2, nft1, nft2)
						}
						return 1000, 1000, tt.ratingErr
					},
					withTransaction: func(_ context.Context, fn func(pgx.Tx) error) error {
						steps = append(steps, "transaction")
						if tt.transactionStartErr != nil {
							return tt.transactionStartErr
						}
						if err := fn(nil); err != nil {
							return err
						}
						return tt.transactionFinishErr
					},
					applyRankingMatch: func(
						_ context.Context,
						_ pgx.Tx,
						gotNFT1, gotNFT2 int64,
						nft1Won bool,
						raNew, rbNew float64,
						gotVoterAid *int64,
					) error {
						steps = append(steps, "apply")
						if gotNFT1 != nft1 || gotNFT2 != nft2 || !nft1Won {
							t.Errorf("apply args = (%d, %d, %v), want (%d, %d, true)", gotNFT1, gotNFT2, nft1Won, nft1, nft2)
						}
						if gotVoterAid == nil || *gotVoterAid != voterAid {
							t.Errorf("apply voterAid = %v, want %d", gotVoterAid, voterAid)
						}
						if raNew != 1125 || rbNew != 875 {
							t.Errorf("applied ratings = (%v, %v), want (1125, 875)", raNew, rbNew)
						}
						return tt.applyErr
					},
				},
				lookupOrCreateAddress: func(_ context.Context, addr string, blockNum, txID int64) (int64, error) {
					steps = append(steps, "lookup")
					if addr != signer {
						t.Errorf("signer = %q, want %q", addr, signer)
					}
					if blockNum != 0 || txID != 0 {
						t.Errorf("address metadata = (%d, %d), want (0, 0)", blockNum, txID)
					}
					return voterAid, tt.lookupErr
				},
				consumeRankingVoteNonce: func(_ context.Context, _ pgx.Tx, gotNonce string) (bool, error) {
					steps = append(steps, "consume")
					if gotNonce != nonce {
						t.Errorf("consumed nonce = %q, want %q", gotNonce, nonce)
					}
					return !tt.nonceMissing, tt.consumeErr
				},
				chainAllowed: func(id int64) bool { return id == chainID },
			}

			err := performSignedBeautyVoteWithDependencies(
				context.Background(),
				nft1,
				nft2,
				true,
				chainID,
				nonce,
				signature,
				deps,
			)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("error = %v, want %v", err, tt.wantErr)
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotSteps := strings.Join(steps, ","); gotSteps != tt.wantSteps {
				t.Fatalf("steps = %q, want %q", gotSteps, tt.wantSteps)
			}
		})
	}
}

func TestClassifyRankingMatchApplyError(t *testing.T) {
	plainErr := errors.New("write failed")
	notUnique := &pgconn.PgError{Code: "23503", Message: "foreign key"}
	unique := &pgconn.PgError{Code: "23505", Message: "duplicate key"}

	tests := []struct {
		name    string
		err     error
		wantErr error
	}{
		{name: "nil", err: nil, wantErr: nil},
		{name: "plain error unchanged", err: plainErr, wantErr: plainErr},
		{name: "other postgres error unchanged", err: notUnique, wantErr: notUnique},
		{name: "unique violation classified", err: unique, wantErr: errRankingDuplicateVoterPair},
		{
			name:    "wrapped unique violation classified",
			err:     fmt.Errorf("store: %w", unique),
			wantErr: errRankingDuplicateVoterPair,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := classifyRankingMatchApplyError(tt.err)
			if tt.wantErr == nil {
				if got != nil {
					t.Fatalf("error = %v, want nil", got)
				}
				return
			}
			if !errors.Is(got, tt.wantErr) {
				t.Fatalf("error = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
