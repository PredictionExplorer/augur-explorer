package randomwalk

// The bindings live in one generated file (bindings.gen.go) produced by a
// single abigen invocation. The abigen version is pinned in go.mod (`tool`
// directive; the go-ethereum module version). `make generate-check` proves
// the committed file matches; see contracts/README.md for provenance and
// the full-from-Solidity workflow.

//go:generate go tool abigen --combined-json buildjson/combined.json --pkg randomwalk --out bindings.gen.go
