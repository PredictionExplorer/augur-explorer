# Contributing

Before picking up work, check [docs/MODERNIZATION.md](docs/MODERNIZATION.md) —
the tracked roadmap for the ongoing Go-native rewrite. It defines the current
phase, the per-file checklists, and the golden rule that every rewrite must
keep the characterization test suite green.

## Development workflow

```bash
make generate     # regenerate committed OpenAPI v2 code
make build       # compile all binaries into ./bin
make test        # unit tests (race detector, shuffled)
make coverage-check # full integration + global/staged coverage policy
make fuzz-smoke  # every fuzz target briefly (FUZZTIME=30s to change)
make lint        # golangci-lint (config in .golangci.yml)
make fix         # apply the Go modernize fixers (CI fails on pending fixes)
make fmt         # gofmt the tree
make hooks-install # install the repository's pre-commit coverage gate
```

CI runs build, tidy-check, unit + integration tests, lint (including a
`go fix -diff` modern-idiom gate), and govulncheck on every PR — all must
pass.

## Code standards

- **Formatting:** gofmt, enforced by CI. No exceptions.
- **Naming:** idiomatic Go — `CamelCase` exported, `camelCase` unexported.
  Legacy `snake_case` identifiers are being renamed as files are touched;
  never introduce new ones.
- **Imports:** no dot-imports. Alias imports when the package name differs
  from the directory or collides (`cgstore`, `rwcontracts`, ...).
- **Errors:** return them. `os.Exit`/`panic` are allowed only in `main`
  startup paths. Wrap with `fmt.Errorf("context: %w", err)`.
- **SQL:** schema changes are goose migrations in `db/migrations`; queries
  are hand-written pgx methods on the `Store`/repos in `internal/store`.
  Always parameterized — never concatenate values.
- **HTTP handlers:** validate inputs at the boundary, keep handlers thin,
  reuse `internal/api/common` helpers. Mutating routes must be authenticated
  and rate limited (see ADR-0004).
- **Comments:** every exported symbol has a doc comment; comments explain
  *why*, not *what*.
- **Generated code** (`contracts/`) is never edited by hand; regenerate
  instead.

## Tests

- Unit tests live next to the code; use table-driven style and `t.Parallel()`
  where the code allows it.
- Integration tests that need PostgreSQL use `internal/testdb` and the
  `integration` build tag: `make test-integration`.
- **API parity suite** (`internal/api/apitest`): boots the real router against
  a seeded testcontainers database and pins every v1 GET route as a golden
  file under `testdata/golden/`. It is the contract for the v1 freeze — any
  rewrite must keep it green. If a response change is *intentional*,
  regenerate the goldens in the same PR and call it out in the description:

  ```bash
  go test -tags=integration ./internal/api/apitest/ -run TestAPIParity -update
  ```

  A route-drift test also asserts `docs/openapi.yaml` and the registered
  routes match exactly, in both directions; new routes must be added to the
  spec and given a parity case in the same PR.
- **API v2** (`internal/api/v2`): edit `docs/openapi-v2.yaml` first, run
  `make generate`, then implement the generated strict interface. V2 route
  drift and kin-openapi response validation run in the unit/integration
  suites; generated-code drift is a CI failure.
- Coverage policy is defined by
  [ADR-0006](docs/adr/0006-coverage-policy.md) and
  [`coverage/policy.json`](coverage/policy.json). CI and the local hook use the
  same tested `covergate` implementation. Current policy enforces legacy and
  handwritten-internal ratchets, a truthful all-production ratchet, and at
  least 95% coverage of changed executable Go statements. Floors only move up.
- Run `make hooks-install` once per clone. Commit enforcement is active and
  fail-closed now that handwritten internal coverage is above 90%; both the
  hook and CI enforce every current ratchet plus 95% changed-code coverage.
  The hook requires Docker, caches successful profiles under
  `.git/coverage-gate/`, and refuses ambiguous partially staged Go changes.
  Local hooks can be bypassed by Git, so the GitHub **Coverage Gate** check
  must remain required by branch protection.
- Fuzz targets (`func Fuzz*`) guard everything that parses untrusted bytes
  (chain data, HTTP input, signatures) — see the inventory in
  [docs/MODERNIZATION.md §4.4](docs/MODERNIZATION.md). Seeds are committed
  inline via `f.Add`; a nightly CI job fuzzes every target and files an issue
  on failure. When a fuzz run finds a crasher, commit the minimized input
  under `testdata/fuzz/` with the fix.
- Bug fixes come with a regression test.

## Commits and PRs

- Small, focused PRs with a clear "why" in the description.
- Update documentation (`docs/`, READMEs, `.env.example`) in the same PR as
  the behavior change.
- Architectural decisions get an ADR in `docs/adr/`.
