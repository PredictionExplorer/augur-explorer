// Package cosmicgame contains the abigen-generated Go bindings for the
// Cosmic Signature contract family (game proxy V1/V2/V3, NFT, token, DAO
// and the charity/prizes/marketing/staking wallets). The bindings are
// generated from the combined-JSON artifacts in buildjson/ — see
// contracts/README.md for provenance and the pinned regeneration workflow
// (go generate ./contracts/...).
//
// The SampToken test ERC-20 (used to exercise bids with ERC-20 donations in
// development) is not part of these bindings; its ABI and creation bytecode
// are embedded in cmd/cgctl/cmd_deploy_erc20.go (`cgctl deploy-erc20`).
package cosmicgame
