// Package assets contains the context-aware operational engines used by the
// opsctl asset commands.
package assets

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5"
)

// SeedHexLength is the full 32-byte seed width as lowercase hexadecimal.
const SeedHexLength = 64

// Token identifies one minted token and its normalized generator seed.
type Token struct {
	TokenID int64
	Seed    string
}

// TokenSource loads minted token seeds for an asset operation.
type TokenSource interface {
	TokenSeeds(ctx context.Context, schema string) ([]Token, error)
}

// Querier is the narrow pgx query surface used by SQLTokenSource.
// *pgxpool.Pool, *pgx.Conn and pgx.Tx satisfy it.
type Querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

// SQLTokenSource loads token seeds from cg_mint_event.
type SQLTokenSource struct {
	DB Querier
}

var (
	schemaNameRE = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_$]*$`)
	seedValueRE  = regexp.MustCompile(`^[0-9a-f]{1,64}$`)
)

// ValidateSchema rejects identifiers that cannot safely be interpolated into
// the fixed cg_mint_event query.
func ValidateSchema(schema string) error {
	if !schemaNameRE.MatchString(schema) {
		return fmt.Errorf("invalid database schema %q", schema)
	}
	return nil
}

// TokenSeeds returns one row per distinct normalized seed, ordered by token id.
func (s SQLTokenSource) TokenSeeds(ctx context.Context, schema string) ([]Token, error) {
	if err := ValidateSchema(schema); err != nil {
		return nil, err
	}
	if s.DB == nil {
		return nil, fmt.Errorf("token database is nil")
	}

	// #nosec G201 -- schema is accepted only after the identifier whitelist above.
	query := fmt.Sprintf(
		"SELECT token_id, seed FROM %s.cg_mint_event WHERE seed IS NOT NULL AND seed <> '' ORDER BY token_id",
		schema,
	)
	rows, err := s.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tokens := make([]Token, 0)
	for rows.Next() {
		var token Token
		if err := rows.Scan(&token.TokenID, &token.Seed); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return NormalizeTokens(tokens), nil
}

// NormalizeSeed lowercases a seed, trims surrounding whitespace, and strips
// one optional 0x prefix.
func NormalizeSeed(seed string) string {
	seed = strings.ToLower(strings.TrimSpace(seed))
	return strings.TrimPrefix(seed, "0x")
}

// LeftPadSeed zero-pads a short seed to the full 64-hex-character width.
func LeftPadSeed(seed string) string {
	seed = NormalizeSeed(seed)
	if len(seed) >= SeedHexLength {
		return seed
	}
	return strings.Repeat("0", SeedHexLength-len(seed)) + seed
}

// SeedName is one candidate on-disk 0x-prefixed seed basename.
type SeedName struct {
	Name   string
	Padded bool
}

// SeedNameCandidates returns the normalized raw basename followed, when
// needed, by the zero-padded 64-character basename.
func SeedNameCandidates(seed string) []SeedName {
	seed = NormalizeSeed(seed)
	if !seedValueRE.MatchString(seed) {
		return nil
	}
	names := []SeedName{{Name: "0x" + seed}}
	if padded := LeftPadSeed(seed); padded != seed {
		names = append(names, SeedName{Name: "0x" + padded, Padded: true})
	}
	return names
}

// NormalizeTokens normalizes, deterministically orders, and de-duplicates
// tokens by seed. Empty seeds are omitted.
func NormalizeTokens(tokens []Token) []Token {
	normalized := make([]Token, 0, len(tokens))
	for _, token := range tokens {
		token.Seed = NormalizeSeed(token.Seed)
		if token.Seed != "" {
			normalized = append(normalized, token)
		}
	}
	sort.SliceStable(normalized, func(i, j int) bool {
		if normalized[i].TokenID != normalized[j].TokenID {
			return normalized[i].TokenID < normalized[j].TokenID
		}
		return normalized[i].Seed < normalized[j].Seed
	})

	seen := make(map[string]struct{}, len(normalized))
	out := normalized[:0]
	for _, token := range normalized {
		if _, ok := seen[token.Seed]; ok {
			continue
		}
		seen[token.Seed] = struct{}{}
		out = append(out, token)
	}
	return out
}
