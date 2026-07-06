package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// assetsCmd groups the NFT asset utilities: on-disk inventory, thumbnail
// generation and token image URL verification.
var assetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "NFT asset utilities (inventory, thumbnails, image checks)",
}

func init() { register(assetsCmd) }

// seedHexLen is the full 32-byte seed width as lowercase hex (no 0x prefix).
const seedHexLen = 64

// tokenRow is one minted Cosmic Signature token with its normalized seed.
type tokenRow struct {
	tokenID int64
	seed    string // lowercase hex, no 0x prefix
}

// seedName is one candidate "0x<seed>" directory/file basename.
type seedName struct {
	name   string
	padded bool // zero-padded 64-char variant of a shorter DB seed
}

// fetchTokenSeeds returns one row per minted token (deduped by seed) from
// cg_mint_event in the given schema, with normalized seeds.
func fetchTokenSeeds(db *sql.DB, schema string) ([]tokenRow, error) {
	query := fmt.Sprintf(
		"SELECT token_id, seed FROM %s.cg_mint_event WHERE seed IS NOT NULL AND seed <> '' ORDER BY token_id",
		schema,
	)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	seen := make(map[string]bool)
	var out []tokenRow
	for rows.Next() {
		var tokenID int64
		var seed string
		if err := rows.Scan(&tokenID, &seed); err != nil {
			return nil, err
		}
		seed = normalizeSeed(seed)
		if seed == "" || seen[seed] {
			continue
		}
		seen[seed] = true
		out = append(out, tokenRow{tokenID: tokenID, seed: seed})
	}
	return out, rows.Err()
}

// normalizeSeed lowercases a seed and strips an optional 0x prefix.
func normalizeSeed(seed string) string {
	seed = strings.ToLower(strings.TrimSpace(seed))
	return strings.TrimPrefix(seed, "0x")
}

// leftPadSeed returns the seed zero-padded to the full 64-hex-char width.
func leftPadSeed(seed string) string {
	if len(seed) >= seedHexLen {
		return seed
	}
	return strings.Repeat("0", seedHexLen-len(seed)) + seed
}

// seedNameCandidates returns candidate "0x<seed>" basenames: the raw DB seed
// and, when shorter than 64 hex chars, the zero-padded variant. Some DB seeds
// drop leading hex zeros while the generator names its package directory with
// the full 64-char seed.
func seedNameCandidates(seed string) []seedName {
	out := []seedName{{name: "0x" + seed}}
	if padded := leftPadSeed(seed); padded != seed {
		out = append(out, seedName{name: "0x" + padded, padded: true})
	}
	return out
}

// isRegularFile reports whether path exists and is not a directory.
func isRegularFile(path string) bool {
	st, err := os.Stat(path)
	return err == nil && !st.IsDir()
}
