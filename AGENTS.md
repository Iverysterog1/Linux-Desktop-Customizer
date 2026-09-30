# Project-wide agent operating contract

This file applies to every automated agent and contributor working in this repository.

## Source of truth

- `integration/canonical` is the sole integrated development truth.
- `main` is the stable/release line and must never receive direct or force pushes.
- Every implementation starts from the latest verified canonical SHA and returns through a reviewed pull request.
- Closed PRs and Git history preserve audit evidence; old topic branches are not alternate sources of truth.

## Current phase — stabilization freeze

As of 2026-09-30, feature expansion is temporarily frozen while the current product foundation is solidified.

During the freeze:
- A1/A2/A3/B1/B2/B3 work only on correctness, tests, security, CI, documentation, packaging baseline, stale capability reporting, validation, and verified-obsolete cleanup.
- Research may continue but produces backlog/architecture input only; it does not activate implementation.
- Security A/B re-audit the exact stabilization head.
- No new product feature is activated until the Supervisor records the freeze as PASS in Issue #47.

The freeze ends only after the stabilization PR is green, documentation/process state matches reality, obsolete active-tree material is classified, and every remaining blocker is explicit.

## Coordination

Issue #47 is the shared coordination ledger and the Supervisor is the single coordination authority.

Programmer roles do not independently create coordination issues or redefine ownership. Before coding, each active slice must have a Supervisor-recorded objective and bounded write-set. Scope expansion requires a Supervisor amendment before editing additional files.

Before creating a branch or PR:
1. verify current `integration/canonical` SHA and protection state;
2. read Issue #47 and open PRs;
3. confirm no overlapping active slice;
4. use a short-lived specialist branch;
5. preserve PASS / FAIL / BLOCKED / NOT RUN evidence.

Never hide or rewrite a failure to make the project appear cleaner.

## Cleanup

Merged/superseded topic branches should be deleted after no-loss verification when repository permissions allow. Historical evidence belongs in Git history, closed PRs/issues, and explicit recovery-evidence paths—not in stale active branches or contradictory current-state documentation.

Do not delete:
- authoritative recovery evidence;
- unresolved unique work;
- security evidence needed to explain a current or historical finding.

## Security and execution boundaries

- no arbitrary downloaded/profile scripts;
- least privilege;
- immutable third-party GitHub Action pins;
- active validation workflows remain read-only;
- no silent privilege escalation;
- no direct publication from development workflows;
- filesystem/profile/transaction boundaries remain bounded and regression-tested.

`integration/canonical` must be protected by repository rules before FINAL READY: PR-required integration, required validation checks, no force pushes, and no branch deletion. Until that is enforced and reverified, Security B remains FAIL.

## Release

No merge/push to `main`, tag, package publication, release publication, or automatic updater activation without explicit final owner authorization and all Definition-of-Done gates passing.

## Evidence convention

Use:
- **PASS** — executed and successful on the stated commit/input;
- **FAIL** — executed and failed;
- **BLOCKED** — authoritative prerequisite or permission is missing;
- **NOT RUN** — not executed.

Assumptions and historical results are never converted into current PASS evidence.
