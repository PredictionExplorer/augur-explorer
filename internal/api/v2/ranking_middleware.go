// Strict-middleware enforcement for the v2 write conventions (ADR-0008):
// per-operation write rate limits answering RFC 9457 429 problems, and the
// spec-declared apiKey admin authentication with v1's fail-closed contract.

package v2

import (
	"context"
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/PredictionExplorer/augur-explorer/internal/api/common"
	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
)

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

// rankingWriteRateLimitMiddleware enforces the per-IP write buckets. The
// limiters are keyed by generated operation id, so the policy lives next
// to the operations it protects; unlisted operations pass through
// untouched. Over-limit requests answer the spec-declared 429 problem with
// Retry-After.
func (s *Server) rankingWriteRateLimitMiddleware() StrictMiddlewareFunc {
	limits := s.rankingConfig.writeLimits()
	limiters := map[string]*common.IPRateLimiter{
		"CreateRandomWalkRankingChallenge": common.NewIPRateLimiter(limits.Challenges.PerSecond, limits.Challenges.Burst),
		"CreateRandomWalkRankingVote":      common.NewIPRateLimiter(limits.Votes.PerSecond, limits.Votes.Burst),
		"RecordRandomWalkRankingMatch":     common.NewIPRateLimiter(limits.Matches.PerSecond, limits.Matches.Burst),
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

// adminKeyMiddleware enforces the RankingAdminKey security scheme on every
// operation the spec declares it for (the generated wrapper marks those
// requests through the scopes context key, so a future secured operation
// is enforced automatically). Enforcement fails closed: with no configured
// key the operation answers 503, and a wrong or missing header answers 401
// after a constant-time compare.
func (s *Server) adminKeyMiddleware() StrictMiddlewareFunc {
	var secret string
	names := make([]string, 0, len(s.rankingConfig.AdminKeys))
	for _, key := range s.rankingConfig.AdminKeys {
		names = append(names, key.Name)
		if secret == "" {
			secret = strings.TrimSpace(key.Value)
		}
	}
	disabledDetail := "Admin authentication is not configured (" + strings.Join(names, " or ") + ")."
	if len(names) == 0 {
		disabledDetail = "Admin authentication is not configured."
	}
	return func(f StrictHandlerFunc, _ string) StrictHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request any) (any, error) {
			if ctx.Value(RankingAdminKeyScopes) == nil {
				return f(ctx, w, r, request)
			}
			if secret == "" {
				s.writeProblem(w, newProblem(
					http.StatusServiceUnavailable,
					"admin-disabled",
					"Admin operation disabled",
					disabledDetail,
					r.URL.Path,
				))
				return nil, nil
			}
			provided := r.Header.Get("X-Ranking-Admin-Key")
			if subtle.ConstantTimeCompare([]byte(provided), []byte(secret)) != 1 {
				s.writeProblem(w, newProblem(
					http.StatusUnauthorized,
					"unauthorized",
					"Unauthorized",
					"The X-Ranking-Admin-Key header is missing or wrong.",
					r.URL.Path,
				))
				return nil, nil
			}
			return f(ctx, w, r, request)
		}
	}
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
