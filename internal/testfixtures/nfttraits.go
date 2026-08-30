package testfixtures

import (
	_ "embed"
)

// NftTraitsExample is docs/fixtures/nft_traits.example.json vendored from
// PredictionExplorer/CS-Image-Generation. It is real generator output for
// seed 0x100033 at reduced simulation settings, and it is the contract
// fixture every trait test runs against.
//
// Regenerate it by re-copying the upstream file. For the same seed and
// parameters, any diff other than generated_at is a determinism regression
// in the generator: report it upstream rather than adapting the Go side.
//
//go:embed nfttraits/nft_traits.example.json
var NftTraitsExample []byte

// NftTraitsSchema is docs/nft_traits.schema.json vendored from the same
// upstream commit. It is the formal JSON Schema (draft 2020-12) the example
// conforms to, kept beside it so the contract this repository implements is
// reviewable in-tree.
//
//go:embed nfttraits/nft_traits.schema.json
var NftTraitsSchema []byte

// AssetManifestExample is a SYNTHETIC assets.json (schema_version 2) for the
// same seed. Unlike NftTraitsExample it is not generator output: the upstream
// repository publishes no manifest fixture, and the real file's sizes and
// hashes describe multi-megabyte binaries that cannot be vendored. Its shape,
// roles, paths and field set mirror write_asset_manifest in the generator's
// src/app.rs; the byte counts and digests are invented.
//
//go:embed nfttraits/assets.synthetic.json
var AssetManifestExample []byte
