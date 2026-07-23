# Contract bindings

Generated Go bindings for the on-chain contracts the backend indexes and
serves:

| Package | Contracts | Generated file |
|---|---|---|
| `contracts/cosmicgame` | CosmicSignatureGame proxy (V1 + V2 + V3 mechanics), CosmicSignature NFT, CosmicSignatureToken (ERC-20), DAO, Charity/Prizes/Marketing wallets, both staking wallets, Arbitrum precompile helpers, OpenZeppelin bases and the SampToken test ERC-20 (113 types) | `bindings.gen.go` |
| `contracts/randomwalk` | RandomWalk NFT, RandomWalk marketplace (2 types) | `bindings.gen.go` |

Each package is one abigen invocation over one committed build artifact,
`buildjson/combined.json`, so shared helper structs (e.g. the
`ICosmicSignatureTokenMintSpec` tuple that both the token and the marketing
wallet reference) are emitted exactly once. Historically the bindings were
one file per contract, which forced a manual edit to a generated file to
resolve that duplicate — the single-file layout removed the hack.

## Regenerating (routine)

The abigen version is pinned through the `tool` directive in `go.mod` (it
tracks the `github.com/ethereum/go-ethereum` module version, currently
1.17.4):

```bash
make generate        # go generate ./internal/api/v2 + ./contracts/...
make generate-check  # regenerate and fail on any diff (runs in CI)
```

Reproducibility is enforced three ways:

- `make generate-check` proves the committed `bindings.gen.go` files are
  exactly what the pinned abigen produces from the committed artifacts.
- `contracts/internal/buildjson`'s `TestArtifactsMatchBindings` pins the
  artifacts to the bindings' embedded `bind.MetaData` in both directions.
- The ETL topic constants, decode fixtures and ABI-unpack fuzz targets pin
  the *behavior* of the ABIs across any regeneration.

## Build artifacts (`buildjson/combined.json`)

The deployed contracts' Solidity sources live in the Cosmic-Signature
repository (outside this monorepo), so the committed artifacts are the
in-repo source of truth for generation. They hold, per contract type, the
ABI document and creation bytecode in solc `--combined-json` shape (the
subset abigen consumes). They were reconstructed from the previously
committed bindings' embedded metadata with:

```bash
go run ./contracts/internal/buildjson/extract contracts/cosmicgame contracts/randomwalk
```

One reconstruction subtlety, handled by that tool: abigen strips every
whitespace rune from the ABI before embedding it — including the spaces in
solc's `"struct Foo.Bar"` / `"enum Foo.Bar"` / `"contract Foo"`
internalType values that abigen itself needs to derive named tuple structs.
The extractor restores those spaces; the drift test compares modulo the
same stripping.

## Full rebuild from Solidity (contract upgrades)

When the contracts change on-chain, the bindings must be rebuilt from real
compiler output — the artifacts here cannot invent new ABIs:

1. On the machine with the Solidity sources, compile with the pinned solc
   versions (`build-wrappers.sh` in `contracts/cosmicgame/` documents the
   exact invocations: solc 0.8.30 for V1 contracts, 0.8.34 for
   CosmicSignatureGameV2/V3, `--via-ir`, OpenZeppelin 5.1/5.02 remappings) and
   produce the per-contract `combined.json` outputs.
2. Merge the needed contract types into the package's
   `buildjson/combined.json` (keep the `<File>.sol:<Type>` naming; drop
   types the backend does not consume, mirroring the old `--exc` lists).
3. Run `make generate`, commit artifacts + bindings together, and let the
   test suite arbitrate: the topic-constant tests
   (`internal/indexer/*/topics.go`) fail loudly if an event signature
   changed, and the decode fixtures replay every event type through the
   real pipeline.

Mechanics upgrades are additive binding types: V2 added
`CosmicSignatureGameV2`, and V3 adds `CosmicSignatureGameV3` reconstructed
from the authoritative `a1eb87d6` generated wrapper. Older types remain
unchanged so all historical event generations stay decodable.
