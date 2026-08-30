// Package nftraits models the trait contract published by the
// CS-Image-Generation pipeline: the per-seed nft_traits.json file and the
// assets.json manifest that accompanies it.
//
// The contract's one rule is that every fact derivable from the seed is
// computed by the Rust generator and copied verbatim here — this package
// never recomputes, renames, re-buckets or reorders an art fact. Attribute
// values and the simulation/generation blocks are therefore carried as
// json.RawMessage so they survive a decode/encode round trip byte-exact.
//
// Facts that come from chain events (round, mint time, prize track, token
// name, ownership) are not modeled here at all; they belong to the indexer.
package nftraits

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// SupportedSchemaMajor is the nft_traits.json layout this package
// understands. Minor and patch bumps only add fields or attribute kinds and
// are safe to ingest unchanged; a major bump means the layout changed and
// the file must be refused so the last known-good row keeps serving.
const SupportedSchemaMajor = 1

// SupportedManifestSchema is the assets.json schema_version that carries a
// sha256 on every single-file entry. Older manifests lack the hashes that
// image_details and animation_details are built from.
const SupportedManifestSchema = 2

// ErrUnsupportedSchema reports an nft_traits.json whose major schema version
// this build cannot interpret. Callers must keep serving the previous row
// (or the fallback) and alert rather than ingesting the file.
var ErrUnsupportedSchema = errors.New("unsupported nft_traits.json schema version")

// ErrSeedMismatch reports an nft_traits.json served for one seed whose body
// describes another. It means the asset host is misconfigured and the file
// must not be attributed to the requested token.
var ErrSeedMismatch = errors.New("nft_traits.json seed does not match the requested seed")

// Attribute is one ready-to-serve marketplace attribute. The formal schema
// declares additionalProperties: false over exactly these four fields, so
// the struct is a lossless representation of any conforming entry.
//
// Value and MaxValue stay raw: a JSON number decoded into float64 and
// re-encoded can change spelling, and the attribute values are part of the
// frozen art contract.
type Attribute struct {
	TraitType   string          `json:"trait_type"`
	DisplayType string          `json:"display_type,omitempty"`
	Value       json.RawMessage `json:"value"`
	MaxValue    json.RawMessage `json:"max_value,omitempty"`
}

// File is the parsed nft_traits.json contract for one seed.
type File struct {
	SchemaVersion   string          `json:"schema_version"`
	Seed            string          `json:"seed"`
	PipelineVersion string          `json:"pipeline_version"`
	GeneratedAt     string          `json:"generated_at"`
	Attributes      []Attribute     `json:"attributes"`
	DescriptionArt  string          `json:"description_art"`
	Simulation      json.RawMessage `json:"simulation"`
	Generation      json.RawMessage `json:"generation"`
}

// ManifestEntry describes one generated file (or, for spectral bins, one
// directory of files) in assets.json.
type ManifestEntry struct {
	Path            string   `json:"path"`
	Kind            string   `json:"kind"`
	Role            string   `json:"role"`
	Format          string   `json:"format"`
	Width           *int     `json:"width,omitempty"`
	Height          *int     `json:"height,omitempty"`
	DurationSeconds *float64 `json:"duration_seconds,omitempty"`
	FrameRate       *int     `json:"frame_rate,omitempty"`
	Codec           string   `json:"codec,omitempty"`
	PixelFormat     string   `json:"pixel_format,omitempty"`
	FileCount       *int     `json:"file_count,omitempty"`
	Bytes           *int64   `json:"bytes,omitempty"`
	SHA256          string   `json:"sha256,omitempty"`
}

// Manifest is the parsed assets.json for one seed.
type Manifest struct {
	SchemaVersion int             `json:"schema_version"`
	GeneratedAt   string          `json:"generated_at"`
	Assets        []ManifestEntry `json:"assets"`
}

// EntryByPath returns the manifest entry for a package-relative path such as
// "images/source/master.png", or nil when the manifest does not describe it
// (image-only generator runs omit every video entry).
func (m *Manifest) EntryByPath(path string) *ManifestEntry {
	if m == nil {
		return nil
	}
	for i := range m.Assets {
		if m.Assets[i].Path == path {
			return &m.Assets[i]
		}
	}
	return nil
}

// Parse decodes and validates an nft_traits.json body. It enforces the
// required fields of the formal schema but deliberately tolerates unknown
// ones: a minor bump may add fields, and refusing them would strand every
// token behind a Go deploy.
//
// Parse does not check the schema major or the seed — those are policy
// decisions the caller makes through Gate, which produces actionable errors.
func Parse(body []byte) (*File, error) {
	var f File
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	if err := dec.Decode(&f); err != nil {
		return nil, fmt.Errorf("decoding nft_traits.json: %w", err)
	}
	if dec.More() {
		return nil, errors.New("decoding nft_traits.json: trailing content after the top-level object")
	}
	if err := f.validate(); err != nil {
		return nil, err
	}
	return &f, nil
}

// validate enforces the schema's required fields. The blocks are checked for
// JSON object-ness rather than shape: their contents are copied verbatim and
// a minor bump may extend them.
func (f *File) validate() error {
	switch {
	case f.SchemaVersion == "":
		return errors.New("nft_traits.json: missing schema_version")
	case f.Seed == "":
		return errors.New("nft_traits.json: missing seed")
	case f.PipelineVersion == "":
		return errors.New("nft_traits.json: missing pipeline_version")
	case f.DescriptionArt == "":
		return errors.New("nft_traits.json: missing description_art")
	case f.Attributes == nil:
		return errors.New("nft_traits.json: missing attributes")
	}
	if !isJSONObject(f.Simulation) {
		return errors.New("nft_traits.json: simulation must be an object")
	}
	if !isJSONObject(f.Generation) {
		return errors.New("nft_traits.json: generation must be an object")
	}
	for i, attr := range f.Attributes {
		if attr.TraitType == "" {
			return fmt.Errorf("nft_traits.json: attribute %d has no trait_type", i)
		}
		if len(attr.Value) == 0 {
			return fmt.Errorf("nft_traits.json: attribute %q has no value", attr.TraitType)
		}
	}
	return nil
}

// Gate applies the ingest policy to a parsed file fetched for wantSeed: the
// major schema version must be the supported one and the body must describe
// the seed that was requested. A gated file is safe to store and serve.
func (f *File) Gate(wantSeed string) error {
	major, err := SemverMajor(f.SchemaVersion)
	if err != nil {
		return err
	}
	if major != SupportedSchemaMajor {
		return fmt.Errorf("%w: got %s, want major %d", ErrUnsupportedSchema, f.SchemaVersion, SupportedSchemaMajor)
	}
	if !SeedsEquivalent(f.Seed, wantSeed) {
		return fmt.Errorf("%w: file says %q, requested %q", ErrSeedMismatch, f.Seed, wantSeed)
	}
	return nil
}

// ParseManifest decodes an assets.json body. Unlike the trait file this is
// advisory data — it only enriches image_details and animation_details — so
// an unexpected schema_version is reported but never fatal to the caller.
func ParseManifest(body []byte) (*Manifest, error) {
	var m Manifest
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	if err := dec.Decode(&m); err != nil {
		return nil, fmt.Errorf("decoding assets.json: %w", err)
	}
	if m.SchemaVersion != SupportedManifestSchema {
		return nil, fmt.Errorf("assets.json: unsupported schema_version %d, want %d", m.SchemaVersion, SupportedManifestSchema)
	}
	return &m, nil
}

// SemverMajor extracts the major component of a "MAJOR.MINOR.PATCH" string.
func SemverMajor(version string) (int, error) {
	major, _, ok := strings.Cut(strings.TrimSpace(version), ".")
	if !ok || major == "" {
		return 0, fmt.Errorf("malformed semver %q", version)
	}
	n, err := strconv.Atoi(major)
	if err != nil {
		return 0, fmt.Errorf("malformed semver %q: %w", version, err)
	}
	if n < 0 {
		return 0, fmt.Errorf("malformed semver %q: negative major", version)
	}
	return n, nil
}

// CanonicalSeed normalizes a seed to the "0x" + lowercase hex form used as
// the trait table's key and in the served properties.seed. Input may or may
// not carry the prefix and may be in any case; anything that is not
// hexadecimal is rejected so a malformed value can never become a URL.
func CanonicalSeed(seed string) (string, error) {
	s := strings.ToLower(strings.TrimSpace(seed))
	s = strings.TrimPrefix(s, "0x")
	if s == "" {
		return "", errors.New("empty seed")
	}
	for _, r := range s {
		if (r < '0' || r > '9') && (r < 'a' || r > 'f') {
			return "", fmt.Errorf("seed %q is not hexadecimal", seed)
		}
	}
	return "0x" + s, nil
}

// SeedsEquivalent reports whether two seeds denote the same value, ignoring
// the "0x" prefix, case and leading zeros.
//
// The indexer stores seeds zero-padded to 64 hex characters because
// NftMinted carries a uint256, while the generator writes whatever width the
// seed was given to it in. "0x100033" and the 64-character padded form are
// the same seed and must compare equal.
func SeedsEquivalent(a, b string) bool {
	na, errA := CanonicalSeed(a)
	nb, errB := CanonicalSeed(b)
	if errA != nil || errB != nil {
		return false
	}
	return trimSeedZeros(na) == trimSeedZeros(nb)
}

// trimSeedZeros strips the prefix and leading zeros, keeping at least one
// digit so a zero seed does not become the empty string.
func trimSeedZeros(canonical string) string {
	digits := strings.TrimLeft(strings.TrimPrefix(canonical, "0x"), "0")
	if digits == "" {
		return "0"
	}
	return digits
}

// ContentHash is a stable digest of the art facts a served token freezes at
// imprint: the attributes, the art description and the simulation and
// generation blocks. The generator is deterministic, so re-running it for
// the same seed must reproduce the same hash; a change is a determinism
// regression to investigate, not an update to apply.
//
// generated_at, pipeline_version and the asset manifest are excluded: the
// first is explicitly non-deterministic and the others can legitimately move
// without the art changing.
func ContentHash(f *File) (string, error) {
	attrs, err := json.Marshal(f.Attributes)
	if err != nil {
		return "", fmt.Errorf("hashing attributes: %w", err)
	}
	simulation, err := CanonicalJSON(f.Simulation)
	if err != nil {
		return "", fmt.Errorf("hashing simulation: %w", err)
	}
	generation, err := CanonicalJSON(f.Generation)
	if err != nil {
		return "", fmt.Errorf("hashing generation: %w", err)
	}
	h := sha256.New()
	for _, part := range [][]byte{attrs, []byte(f.DescriptionArt), simulation, generation} {
		// Length-prefixed so no concatenation of two parts can collide with
		// a different split of the same bytes.
		_, _ = h.Write([]byte(strconv.Itoa(len(part)) + ":"))
		_, _ = h.Write(part)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// CanonicalJSON re-encodes a JSON value with object keys sorted and
// insignificant whitespace removed, preserving number literals exactly. Two
// encodings of the same value always produce the same bytes, which is what
// makes the content hash and the stored blocks comparable across runs.
func CanonicalJSON(raw json.RawMessage) ([]byte, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	var v any
	if err := dec.Decode(&v); err != nil {
		return nil, err
	}
	// Go sorts map keys when marshaling, and json.Number marshals as its
	// original literal, so this pass is both canonical and lossless.
	return json.Marshal(v)
}

// isJSONObject reports whether raw holds a JSON object.
func isJSONObject(raw json.RawMessage) bool {
	trimmed := bytes.TrimSpace(raw)
	return len(trimmed) > 0 && trimmed[0] == '{'
}
