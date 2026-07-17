// Strict-middleware enforcement for the v2 write conventions (ADR-0008/0009):
// per-operation write rate limits answering RFC 9457 429 problems, and the
// spec-declared apiKey admin authentication with v1's fail-closed contract.

package v2

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

// AdminConfig carries CosmicGame admin authentication and bid-moderation
// write limits. The zero value fails closed and applies production limits.
type AdminConfig struct {
	// AdminKeys guard the bid-ban create and delete operations. The first
	// non-empty value is the expected secret.
	AdminKeys []common.AdminKey

	// WriteLimits overrides the two moderation write buckets. Nil applies
	// the production defaults.
	WriteLimits *AdminWriteLimits
}

// AdminWriteLimits are the per-IP token buckets for bid-ban mutations.
type AdminWriteLimits struct {
	CreateBidBan RateLimitSpec
	DeleteBidBan RateLimitSpec
}

// RankingConfig carries the ranking slice's deployment configuration. The
// zero value is a working fail-closed default: the admin operation answers
// 503, any positive vote chain id is allowed, the exploration bound is the
// 3766 collection default and the production write rate limits apply.
type RankingConfig struct {
	// AdminKeys guard the admin match operation. The first key with a
	// non-empty value is the expected secret; when every value is empty
	// the operation fails closed with 503 (never anonymous access). Key
	// names appear in the disabled problem so operators know which
	// variables to configure.
	AdminKeys []common.AdminKey

	// VoteChainIDs allowlists the wallet-signed vote chain ids
	// (RANKING_VOTE_CHAIN_IDS); empty allows any positive chain id.
	VoteChainIDs []int64

	// ExploreMaxTokenID bounds exploration token ids
	// (RANDOMWALK_EXPLORE_MAX_TOKEN_ID); zero or negative applies the
	// 3766 collection default.
	ExploreMaxTokenID int64

	// WriteLimits overrides the per-operation write rate limits; nil
	// applies the production defaults. Test harnesses that drive many
	// mutations per second raise them.
	WriteLimits *RankingWriteLimits
}

// RankingWriteLimits are the per-IP token buckets of the three ranking
// write operations.
type RankingWriteLimits struct {
	Challenges RateLimitSpec
	Votes      RateLimitSpec
	Matches    RateLimitSpec
}

// RateLimitSpec is one token bucket: sustained requests per second and
// burst capacity.
type RateLimitSpec struct {
	PerSecond float64
	Burst     int
}

// defaultRankingWriteLimits mirror the v1 registration-time limits (votes
// 1/10, admin matches 2/5) plus a challenge budget that always covers the
// vote budget, since every vote consumes one challenge.
func defaultRankingWriteLimits() RankingWriteLimits {
	return RankingWriteLimits{
		Challenges: RateLimitSpec{PerSecond: 2, Burst: 20},
		Votes:      RateLimitSpec{PerSecond: 1, Burst: 10},
		Matches:    RateLimitSpec{PerSecond: 2, Burst: 5},
	}
}

func defaultAdminWriteLimits() AdminWriteLimits {
	return AdminWriteLimits{
		CreateBidBan: RateLimitSpec{PerSecond: 2, Burst: 5},
		DeleteBidBan: RateLimitSpec{PerSecond: 2, Burst: 5},
	}
}

// defaultExploreMaxTokenID bounds exploration token ids when the
// configuration does not override it: the frozen RandomWalk collection's
// final token id (same default as v1 and RANDOMWALK_EXPLORE_MAX_TOKEN_ID).
const defaultExploreMaxTokenID = 3766

func (c RankingConfig) writeLimits() RankingWriteLimits {
	if c.WriteLimits != nil {
		return *c.WriteLimits
	}
	return defaultRankingWriteLimits()
}

func (c AdminConfig) writeLimits() AdminWriteLimits {
	if c.WriteLimits != nil {
		return *c.WriteLimits
	}
	return defaultAdminWriteLimits()
}

func (c RankingConfig) exploreMaxTokenID() int64 {
	if c.ExploreMaxTokenID > 0 {
		return c.ExploreMaxTokenID
	}
	return defaultExploreMaxTokenID
}

// voteChainAllowed reports whether a wallet-signed vote's chain id passes
// the configured allowlist (empty allows any positive chain id).
func (c RankingConfig) voteChainAllowed(chainID int64) bool {
	if chainID <= 0 {
		return false
	}
	if len(c.VoteChainIDs) == 0 {
		return true
	}
	for _, allowed := range c.VoteChainIDs {
		if allowed == chainID {
			return true
		}
	}
	return false
}

// writeRateLimitMiddleware enforces all per-operation v2 write buckets.
// Unlisted operations pass through untouched. Over-limit requests answer the
// spec-declared 429 problem with Retry-After.
func (s *Server) writeRateLimitMiddleware() StrictMiddlewareFunc {
	rankingLimits := s.rankingConfig.writeLimits()
	adminLimits := s.adminConfig.writeLimits()
	limiters := map[string]*common.IPRateLimiter{
		"CreateCosmicGameBidBan":           common.NewIPRateLimiter(adminLimits.CreateBidBan.PerSecond, adminLimits.CreateBidBan.Burst),
		"DeleteCosmicGameBidBan":           common.NewIPRateLimiter(adminLimits.DeleteBidBan.PerSecond, adminLimits.DeleteBidBan.Burst),
		"CreateRandomWalkRankingChallenge": common.NewIPRateLimiter(rankingLimits.Challenges.PerSecond, rankingLimits.Challenges.Burst),
		"CreateRandomWalkRankingVote":      common.NewIPRateLimiter(rankingLimits.Votes.PerSecond, rankingLimits.Votes.Burst),
		"RecordRandomWalkRankingMatch":     common.NewIPRateLimiter(rankingLimits.Matches.PerSecond, rankingLimits.Matches.Burst),
	}
	return func(f StrictHandlerFunc, operationID string) StrictHandlerFunc {
		limiter, limited := limiters[operationID]
		if !limited {
			return f
		}
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request any) (any, error) {
			if !limiter.Allow(httpx.ClientIP(r)) {
				w.Header().Set("Retry-After", "1")
				s.writeProblem(w, rateLimitedProblem(r.URL.Path))
				return nil, nil
			}
			return f(ctx, w, r, request)
		}
	}
}

type adminKeyPolicy struct {
	active     func(context.Context) bool
	headerName string
	keys       []common.AdminKey
}

// adminKeyMiddleware enforces both generated admin security schemes. The
// generated wrapper marks secured requests through scheme-specific context
// keys, so future operations using either scheme inherit fail-closed
// enforcement automatically.
func (s *Server) adminKeyMiddleware() StrictMiddlewareFunc {
	policies := []adminKeyPolicy{
		{
			active: func(ctx context.Context) bool {
				return ctx.Value(AdminKeyScopes) != nil
			},
			headerName: "X-Admin-Key",
			keys:       s.adminConfig.AdminKeys,
		},
		{
			active: func(ctx context.Context) bool {
				return ctx.Value(RankingAdminKeyScopes) != nil
			},
			headerName: "X-Ranking-Admin-Key",
			keys:       s.rankingConfig.AdminKeys,
		},
	}
	return func(f StrictHandlerFunc, _ string) StrictHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request any) (any, error) {
			for _, policy := range policies {
				if !policy.active(ctx) {
					continue
				}
				secret, names := configuredAdminSecret(policy.keys)
				if secret == "" {
					detail := "Admin authentication is not configured."
					if len(names) > 0 {
						detail = "Admin authentication is not configured (" + strings.Join(names, " or ") + ")."
					}
					s.writeProblem(w, newProblem(
						http.StatusServiceUnavailable,
						"admin-disabled",
						"Admin operation disabled",
						detail,
						r.URL.Path,
					))
					return nil, nil
				}

				expectedHash := sha256.Sum256([]byte(secret))
				providedHash := sha256.Sum256([]byte(r.Header.Get(policy.headerName)))
				if subtle.ConstantTimeCompare(providedHash[:], expectedHash[:]) != 1 {
					s.writeProblem(w, newProblem(
						http.StatusUnauthorized,
						"unauthorized",
						"Unauthorized",
						"The "+policy.headerName+" header is missing or wrong.",
						r.URL.Path,
					))
					return nil, nil
				}
				return f(ctx, w, r, request)
			}
			return f(ctx, w, r, request)
		}
	}
}

func configuredAdminSecret(keys []common.AdminKey) (string, []string) {
	var secret string
	names := make([]string, 0, len(keys))
	for _, key := range keys {
		names = append(names, key.Name)
		if secret == "" {
			secret = strings.TrimSpace(key.Value)
		}
	}
	return secret, names
}

func rateLimitedProblem(instance string) Problem {
	return newProblem(
		http.StatusTooManyRequests,
		"rate-limited",
		"Rate limit exceeded",
		"Too many requests for this operation; retry after the indicated delay.",
		instance,
	)
}
