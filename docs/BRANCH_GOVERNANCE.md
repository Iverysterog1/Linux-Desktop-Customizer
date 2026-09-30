# Branch governance

## Canonical flow

- `main` is the stable/release branch. Never push directly or force-push to it.
- `integration/canonical` is the sole integration/development truth.
- All changes use: latest canonical -> short-lived topic branch -> validation/review -> PR -> canonical.
- A merged/superseded topic branch is disposable after no-loss verification; its audit trail remains in Git and the closed PR.

## Required protection

Repository rules are currently active and verified.

For `integration/canonical`:
- pull-request-based integration;
- strict/up-to-date required checks: `Canonical CI`, `Race verification`, `Workflow write guard`;
- required conversation resolution;
- squash-only merging;
- no force pushes;
- no branch deletion;
- no bypass actors.

For `main`:
- pull-request-based integration;
- strict/up-to-date required checks: `Canonical CI`, `Race verification`, `Workflow write guard`, `Publication verification`;
- required conversation resolution;
- squash-only merging;
- no force pushes;
- no branch deletion;
- no bypass actors.

Review requirements match the current single-owner reviewer model with zero mandatory independent approvals, avoiding an impossible self-approval gate.

Security B is **PASS** for the current stabilization boundary. Re-verify these controls on security-sensitive changes and again on the final release candidate.

## Coordination

Issue #47 and the Supervisor own cross-team coordination. Specialist branches are narrow, non-overlapping and never alternate sources of truth.

No direct canonical push is an accepted substitute for PR review, even when the branch is technically writable.

## Cleanup

Old topic branches are not historical archives. Once unique content is verified integrated/superseded and the PR/history preserves evidence, delete the branch when permissions allow. The repository currently retains only `main` and `integration/canonical`; automatic deletion of merged head branches is enabled.

Recovery evidence that documents genuinely unrecoverable historical source remains preserved under explicit evidence paths and Git history.

## Safety invariants

- no automatic merge to `main`;
- no automatic release/tag/package publication;
- least-privilege GitHub Actions;
- immutable third-party Action pins;
- no hidden failed checks;
- no protection weakening to unblock development;
- no arbitrary profile/downloaded execution.
