package cosmicgame

// The bindings live in one generated file (bindings.gen.go) produced by a
// single abigen invocation, so shared tuple structs (e.g. the MintSpec
// helper both the token and marketing-wallet contracts reference) are
// emitted exactly once. The abigen version is pinned in go.mod (`tool`
// directive; the go-ethereum module version). `make generate-check` proves
// the committed file matches; see contracts/README.md for provenance and
// the full-from-Solidity workflow.

//go:generate go tool abigen --combined-json buildjson/combined.json --pkg cosmicgame --out bindings.gen.go
