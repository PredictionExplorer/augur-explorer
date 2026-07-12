# ADR-0006: Coverage quality policy and gates

- Status: Accepted
- Date: 2026-07-11

## Context

The previous CI gate reported 75.5% statement coverage from
`-coverpkg=./internal/...`. That number was useful as a historical ratchet, but
it was not whole-codebase production coverage: it included generated API code
and test-only harnesses while excluding all `cmd/` services and tools.

Coverage is a gap detector, not a proof of correctness. A high percentage can
still hide missing assertions, branches, concurrency failures and incorrect
requirements. Conversely, generated routing code and thin startup wiring can
distort an otherwise useful metric.

## Decision

The versioned policy in [`coverage/policy.json`](../../coverage/policy.json) is
the single source of truth for local hooks and CI.

We report and ratchet four complementary metrics:

1. **Legacy internal:** the original unfiltered `internal/...` denominator, so
   the historical metric cannot silently regress while the policy changes.
2. **Handwritten internal:** production `internal/...`, excluding generated
   OpenAPI output and packages that exist only to support tests.
3. **All production:** handwritten `cmd/...` plus handwritten production
   `internal/...`. This exposes rather than hides untested operational tools.
4. **Changed code:** executable statement blocks intersecting staged or PR Go
   changes, excluding the same generated/test-only paths.

The staged targets are:

- handwritten internal: **90% minimum**;
- all changed code: **95% minimum immediately**;
- critical packages: **95% final target**;
- all-production coverage: ratchet from its truthful baseline toward 90% as
  command logic moves behind testable package seams.

Global floors never decrease. Raising a floor requires a green full integration
profile proving the new value. Lowering or changing scope requires a new ADR.

Go's coverage profile is emitted once per test binary when `-coverpkg` spans
the repository. `internal/covergate` therefore deduplicates blocks by source
location and unions execution counts before calculating any percentage.

## Enforcement

- `scripts/coverage-gate.sh` generates an atomic integration profile over
  `cmd/...` and `internal/...`, caches successful staged-source runs under
  `.git/coverage-gate`, and invokes the tested `cmd/covergate` policy engine.
- `.githooks/pre-commit` reads the policy's commit-gate status. **Activated
  2026-07-12:** handwritten internal coverage first reached 90% (91.82%
  measured under the race-enabled profile), so `commitGateEnabled` is `true`
  with a 90.0 commit floor and a 91.5 internal floor; the hook now fails
  closed on test, Docker, profile, policy, global or changed-code failures.
  `make hooks-install` installs the tracked hook without changing Git
  configuration or overwriting unrelated hooks.
- `.cursor/hooks.json` denies Agent attempts to use `git commit --no-verify`
  and denies commits when the native hook is missing or stale.
- GitHub Actions runs the same policy engine with race-enabled integration
  coverage and the PR diff. The stable **Coverage Gate** check must be required
  by branch protection and is the authoritative merge barrier.

The activation change set both the handwritten internal and commit floors to
at least 90%, proved a sub-floor profile is rejected and the measured ≥90
profile passes; neither floor is ever lowered afterward.

Native Git hooks are local and can be bypassed deliberately. Repository code
cannot make that impossible; required CI branch protection closes that gap.

## Test-quality requirements

Coverage-only tests are not acceptable. New cases must assert externally
observable behavior, invariants, error secrecy, state transitions or exact
wire/storage contracts. Race tests, fuzzing, integration fixtures, golden
contracts and vulnerability scanning remain independent mandatory gates.

The program prioritizes live API error paths, signed mutations, store
transactions, indexer replay/sync behavior and external boundary adapters.
Generated response visitors are not tested merely to inflate the percentage.

## Consequences

With the commit gate active, commits that change Go code can take tens of
seconds with a warm Docker cache and longer on first use. Successful profiles
are cached by staged-source hash; ambiguous partial staging, missing Docker
and malformed evidence fail closed.

The first policy sprint raised handwritten internal coverage to 80.43%, the
legacy metric to 79.07%, and truthful all-production coverage to 52.80% under
the authoritative race-enabled CI command.
The API-boundary sprint raised those metrics to 86.07%, 83.73%, and 56.26%.
The store/indexer/notification sprint raised them to 91.82%, 88.48%, and
59.65% — past the 90% handwritten-internal milestone — and activated the
commit gate; the enforced floors are now 91.5%, 88.2%, and 59.4%, while
changed executable code remains gated at 95%. The next milestones are 95%
handwritten internal coverage and the operational-command extraction that
moves all-production coverage toward 90%.
