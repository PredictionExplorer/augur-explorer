package cosmicgame

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/PredictionExplorer/augur-explorer/internal/api/httpx"
	"github.com/PredictionExplorer/augur-explorer/internal/nftraits"
	cgstore "github.com/PredictionExplorer/augur-explorer/internal/store/cosmicgame"
)

// metadataVersion is the version of the *served* metadata layout, bumped
// when this handler changes the shape it emits. It is independent of the
// trait file's schema_version, which versions the generator's contract.
//
// 2.0.0 is the first layout without properties.owner and without the
// per-token seed attribute, and the first that can carry art traits.
const metadataVersion = "2.0.0"

// Package-relative paths of the two assets the served metadata describes in
// detail. They are the canonical still and the canonical animation; the
// remaining package files are advertised through properties.media.
const (
	masterImagePath = "images/source/master.png"
	mainVideoPath   = "videos/web/main.mp4"
)

// tokenMetadataInput is everything the assembly needs: chain facts on the
// left, the optional art contract on the right.
type tokenMetadataInput struct {
	TokenID   int64
	RoundNum  int64
	TokenName string
	// MintTimestamp is unix seconds, zero when unavailable.
	MintTimestamp int64
	// Seed is canonical ("0x" + lowercase hex) whenever the stored value
	// was hexadecimal; otherwise it is the raw stored value with the prefix
	// added, matching the long-standing asset URL behaviour.
	Seed string
	// AssetBase is the public /images prefix, without a trailing slash.
	AssetBase string
	// SourceBase is the asset host root that publishes the trait contract
	// files, empty when trait ingestion is not configured.
	SourceBase string
	// Allocation is the display name of the prize path that minted the
	// token, empty when provenance could not be resolved.
	Allocation string
	// Traits is the stored art contract, nil on the fallback path.
	Traits *tokenTraits
}

// tokenTraits is the decoded art contract for one token.
type tokenTraits struct {
	Attributes     []nftraits.Attribute
	DescriptionArt string
	Simulation     json.RawMessage
	Generation     json.RawMessage
	Manifest       *nftraits.Manifest
}

// decodeTokenTraits turns a stored row into the assembly's view of it. A row
// whose art blocks cannot be decoded is treated as absent: serving the
// fallback is always better than serving a half-built token.
func decodeTokenTraits(row cgstore.TokenTraitsRow) (*tokenTraits, error) {
	var attributes []nftraits.Attribute
	if err := json.Unmarshal(row.Attributes, &attributes); err != nil {
		return nil, err
	}
	traits := &tokenTraits{
		Attributes:     attributes,
		DescriptionArt: row.DescriptionArt,
		Simulation:     json.RawMessage(row.Simulation),
		Generation:     json.RawMessage(row.Generation),
	}
	if len(row.Assets) > 0 {
		var manifest nftraits.Manifest
		if err := json.Unmarshal(row.Assets, &manifest); err != nil {
			return nil, err
		}
		traits.Manifest = &manifest
	}
	return traits, nil
}

// buildTokenMetadata assembles the ERC-721 metadata document.
//
// The merge rule is fixed: art facts are copied verbatim from the trait
// contract and chain facts are appended after them. Nothing mutable by a
// transfer is included — an owner baked into a document that marketplaces
// cache for days is permanently stale, and ownerOf() is the only truthful
// answer.
func buildTokenMetadata(in tokenMetadataInput) httpx.H {
	meta := httpx.H{
		"metadata_version": metadataVersion,
		"name":             tokenDisplayName(in.TokenName, in.TokenID),
		"description":      buildDescription(in),
		"image":            in.AssetBase + "/new/cosmicsignature/" + in.Seed + ".png",
		// animation_url is the marketplace-standard field (OpenSea et al.)
		// for the animated representation of the same artwork.
		"animation_url": in.AssetBase + "/new/cosmicsignature/" + in.Seed + ".mp4",
		// The art is rendered on pure black; this matches the frame the
		// marketplace draws around it.
		"background_color": "000000",
		"external_url":     "https://www.cosmicsignature.com/detail/" + strconv.FormatInt(in.TokenID, 10),
		"attributes":       buildAttributes(in),
		"properties":       buildProperties(in),
	}
	if in.Traits != nil {
		if details := imageDetails(in.Traits.Manifest); details != nil {
			meta["image_details"] = details
		}
		if details := animationDetails(in.Traits.Manifest); details != nil {
			meta["animation_details"] = details
		}
	}
	return meta
}

// tokenDisplayName prefers the on-chain custom name and falls back to the
// numbered default.
func tokenDisplayName(tokenName string, tokenID int64) string {
	if tokenName != "" {
		return tokenName
	}
	return "Cosmic Signature #" + strconv.FormatInt(tokenID, 10)
}

// buildAttributes concatenates the art attributes, verbatim and first, with
// the chain-derived ones.
//
// There is deliberately no seed attribute: every value is unique, so it can
// never filter anything, and it crowds out the traits that can. The
// canonical seed is in properties.seed.
func buildAttributes(in tokenMetadataInput) []nftraits.Attribute {
	var attributes []nftraits.Attribute
	if in.Traits != nil {
		// Copied as-is: never renamed, re-bucketed or reordered. The
		// generator owns everything derivable from the seed.
		attributes = append(attributes, in.Traits.Attributes...)
	}
	attributes = append(attributes, numberAttribute("Round", in.RoundNum))
	if in.MintTimestamp > 0 {
		attributes = append(attributes, dateAttribute("Imprinted", in.MintTimestamp))
	}
	if in.Allocation != "" {
		attributes = append(attributes, stringAttribute("Allocation", in.Allocation))
	}
	return attributes
}

// buildDescription puts the art first and the provenance second. On the
// fallback path the art sentence is unavailable, so the long-standing
// generic description stands in until the package is ingested.
func buildDescription(in tokenMetadataInput) string {
	if in.Traits == nil {
		return "Discover the unique attributes and ownership history of Cosmic Signature Token #" +
			strconv.FormatInt(in.TokenID, 10) +
			", an exclusive digital collectible from the Cosmic Signature game."
	}
	description := in.Traits.DescriptionArt + " Imprinted in Round " + strconv.FormatInt(in.RoundNum, 10)
	if in.MintTimestamp > 0 {
		description += " on " + time.Unix(in.MintTimestamp, 0).UTC().Format("02 Jan 2006")
	}
	return description + ". Same seed, same pixels — re-render it to verify."
}

// buildProperties carries the machine-readable half of the document: the
// canonical seed, the reproducibility certificate copied verbatim from the
// trait file, and every hosted asset of the package.
func buildProperties(in tokenMetadataInput) httpx.H {
	properties := httpx.H{
		"seed":      in.Seed,
		"token_id":  in.TokenID,
		"round_num": in.RoundNum,
	}
	if in.Traits == nil {
		return properties
	}
	properties["simulation"] = in.Traits.Simulation
	properties["generation"] = in.Traits.Generation
	properties["media"] = buildMedia(in)
	return properties
}

// buildMedia advertises the package files that the image and animation_url
// fields cannot reach: the HQ master video, both spectral sweeps and the
// 64-bin spectral image set. Without this they are hosted but undiscoverable.
//
// The manifest and trait-source URLs are only emitted when an asset host is
// configured, because those two paths are published by the asset host rather
// than by the /images mount.
func buildMedia(in tokenMetadataInput) httpx.H {
	pkg := in.AssetBase + "/new/cosmicsignature/" + in.Seed
	media := httpx.H{
		"hq_video":           pkg + "/videos/hq/main.mp4",
		"spectral_sweep":     pkg + "/videos/web/spectral_sweep.mp4",
		"spectral_sweep_hq":  pkg + "/videos/hq/spectral_sweep.mp4",
		"spectral_bins":      pkg + "/spectral/",
		"web_image":          pkg + "/images/web/full.webp",
		"source_image":       pkg + "/" + masterImagePath,
		"preview_image":      pkg + "/images/web/preview.webp",
		"generation_records": pkg + "/metadata/",
	}
	if in.SourceBase != "" {
		media["asset_manifest"] = in.SourceBase + "/asset-manifests/" + in.Seed + ".json"
		media["trait_source"] = in.SourceBase + "/traits/" + in.Seed + ".json"
	}
	return media
}

// imageDetails describes the still that the image field resolves to. The
// hashes come from the generator's manifest; the metadata server never
// hashes an asset itself.
func imageDetails(manifest *nftraits.Manifest) httpx.H {
	entry := manifest.EntryByPath(masterImagePath)
	if entry == nil {
		return nil
	}
	details := httpx.H{}
	putIfSet(details, "width", entry.Width)
	putIfSet(details, "height", entry.Height)
	putIfNotEmpty(details, "format", entry.Format)
	putIfSet(details, "bytes", entry.Bytes)
	putIfNotEmpty(details, "sha256", entry.SHA256)
	if len(details) == 0 {
		return nil
	}
	return details
}

// animationDetails describes the video that animation_url resolves to.
func animationDetails(manifest *nftraits.Manifest) httpx.H {
	entry := manifest.EntryByPath(mainVideoPath)
	if entry == nil {
		return nil
	}
	details := httpx.H{}
	putIfSet(details, "width", entry.Width)
	putIfSet(details, "height", entry.Height)
	putIfNotEmpty(details, "format", entry.Format)
	putIfNotEmpty(details, "codec", entry.Codec)
	putIfSet(details, "duration_seconds", entry.DurationSeconds)
	putIfSet(details, "bytes", entry.Bytes)
	putIfNotEmpty(details, "sha256", entry.SHA256)
	if len(details) == 0 {
		return nil
	}
	return details
}

// allocationLabels maps the indexer's prize families onto the display names
// the collection uses for provenance. Both staker raffle pools share
// "Anchored Selection": the distinction between a Cosmic Signature stake and
// a RandomWalk stake is a mechanic, not a provenance story.
var allocationLabels = map[cgstore.CosmicSignatureMintSource]string{
	cgstore.MintSourceMainPrize:          "Final Gesture",
	cgstore.MintSourceEnduranceChampion:  "Endurance Champion",
	cgstore.MintSourceChronoWarriorPrize: "Chrono-Warrior",
	cgstore.MintSourceBidderRaffle:       "Stellar Selection",
	cgstore.MintSourceCosmicSigStaker:    "Anchored Selection",
	cgstore.MintSourceRandomWalkStaker:   "Anchored Selection",
	cgstore.MintSourceLastCstBidder:      "Last CST Gesture",
}

// allocationLabel returns the display name for a mint source, or the empty
// string for an unresolved one. An unknown provenance omits the attribute
// rather than guessing at it.
func allocationLabel(source cgstore.CosmicSignatureMintSource) string {
	return allocationLabels[source]
}

// metadataSeed normalizes a stored seed for use in URLs and properties.
// Hexadecimal seeds become the canonical "0x" + lowercase form; anything
// else keeps the long-standing behaviour of prefixing the raw value, so a
// non-conforming legacy row still resolves to its assets.
func metadataSeed(stored string) string {
	if canonical, err := nftraits.CanonicalSeed(stored); err == nil {
		return canonical
	}
	return "0x" + stored
}

// numberAttribute builds a numeric marketplace attribute.
func numberAttribute(traitType string, value int64) nftraits.Attribute {
	return nftraits.Attribute{
		TraitType:   traitType,
		DisplayType: "number",
		Value:       json.RawMessage(strconv.FormatInt(value, 10)),
	}
}

// dateAttribute builds a date attribute from unix seconds, which is the
// representation marketplaces expect for display_type "date".
func dateAttribute(traitType string, unixSeconds int64) nftraits.Attribute {
	return nftraits.Attribute{
		TraitType:   traitType,
		DisplayType: "date",
		Value:       json.RawMessage(strconv.FormatInt(unixSeconds, 10)),
	}
}

// stringAttribute builds a plain string attribute.
func stringAttribute(traitType, value string) nftraits.Attribute {
	// Encoding a Go string is total: invalid UTF-8 is replaced rather than
	// rejected, and only channels, functions and non-finite floats can make
	// json.Marshal fail. There is no error to handle.
	encoded, _ := json.Marshal(value) //nolint:errchkjson // string encoding cannot fail
	return nftraits.Attribute{TraitType: traitType, Value: encoded}
}

// putIfSet stores a pointer's target when the manifest supplied one.
func putIfSet[T any](h httpx.H, key string, value *T) {
	if value != nil {
		h[key] = *value
	}
}

// putIfNotEmpty stores a string when the manifest supplied one.
func putIfNotEmpty(h httpx.H, key, value string) {
	if value != "" {
		h[key] = value
	}
}
