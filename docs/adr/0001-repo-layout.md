# ADR-0001: Single module with cmd/ + internal/ layout

Date: 2026-07-06. Status: accepted.

## Context

The repository historically carried two generations of code: the frozen
`previous-code/` tree (an Augur/DeFi explorer, ~450k lines) and the active
`rwcg/` tree, each with its own ad-hoc layout, per-directory `make.sh`
scripts, and ~90 single-file `package main` scripts that could not compile
together (`go build ./...` failed).

## Decision

- Delete `previous-code/` (recoverable from git history). Actively maintained
  daemons that still lived there (`srvmonitor`, `loganomaly`) were migrated
  into the active tree first.
- Adopt the standard Go layout at the repository root:
  - `cmd/<name>` — one directory per binary; self-contained tool internals
    may nest under `cmd/<name>/internal`.
  - `internal/` — shared packages, unimportable from outside the module.
  - `contracts/` — generated abigen bindings, kept importable and excluded
    from lint.
- Consolidate the one-off scripts into three cobra CLIs (`cgctl`, `rwctl`,
  `opsctl`); delete scripts that were unreferenced by any build or ops script.
- Replace the per-directory build scripts with a single root `Makefile` that
  builds every binary into `bin/`.

## Consequences

`go build ./...`, `go vet ./...` and `go test ./...` operate on the whole
repository. CI can enforce them. New binaries follow one obvious pattern, and
operator functionality is discoverable via `--help` instead of directory
spelunking.
