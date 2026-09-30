# Project-wide agent operating contract

This file applies to every automated agent and contributor working in this repository.

## Source of truth

- `integration/canonical` is the sole integrated development truth.
- `main` is the stable/release line and must never receive direct or force pushes.
- Every implementation starts from the latest verified canonical SHA and returns through a reviewed pull request.
- Closed PRs and Git history preserve audit evidence; old topic branches are not alternate sources of truth.

## Current phase — post-stabilization KDE validation

The 2026-09-30 stabilization gate is **PASS** and the temporary feature freeze is lifted.

Controlled product development may resume under the Supervisor and Issue #47. The next priority is real Plasma/KDE evidence for the existing safe vertical slice before broader feature expansion:
- detect real Plasma capability;
- preview the actual current state;
- snapshot/journal;
- apply a safe reviewed change;
- validate the effective/visible result;
- rollback/unapply to the exact prior state.

A1/A2/A3/B1/B2/B3 retain bounded, non-overlapping write-sets. Research remains backlog input until the Supervisor explicitly adopts it. Security A/B re-audit exact heads whenever security-sensitive boundaries change.

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

`integration/canonical` and `main` are both protected by active repository rulesets with PR-required integration, strict/up-to-date required checks, no force pushes/deletion and no bypass actors. Legacy write-capable `main` workflows have been removed, the four release-line checks were proven on a temporary validation PR, stale topic branches were cleaned, and automatic merged-head deletion is enabled. Repository-wide Security B is PASS for the current stabilization boundary.

## Release

No merge/push to `main`, tag, package publication, release publication, or automatic updater activation without explicit final owner authorization and all Definition-of-Done gates passing.

## Evidence convention

Use:
- **PASS** — executed and successful on the stated commit/input;
- **FAIL** — executed and failed;
- **BLOCKED** — authoritative prerequisite or permission is missing;
- **NOT RUN** — not executed.

Assumptions and historical results are never converted into current PASS evidence.
