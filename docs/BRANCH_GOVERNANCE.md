# Branch governance

## Canonical flow

- `main` is the stable/release branch. Never push directly or force-push to it.
- `integration/canonical` is the sole integration/development truth.
- All changes use: latest canonical -> short-lived topic branch -> validation/review -> PR -> canonical.
- A merged/superseded topic branch is disposable after no-loss verification; its audit trail remains in Git and the closed PR.

## Required protection

Before FINAL READY, repository rules for `integration/canonical` must enforce:
- pull-request-based integration;
- required validation/status checks appropriate to the changed scope;
- no force pushes;
- no branch deletion.

Review requirements must match the repository's actual reviewer model; do not configure an impossible self-approval gate for a single-owner workflow.

Until these rules are enabled and reverified, Security B remains FAIL — MEDIUM.

## Coordination

Issue #47 and the Supervisor own cross-team coordination. Specialist branches are narrow, non-overlapping and never alternate sources of truth.

No direct canonical push is an accepted substitute for PR review, even when the branch is technically writable.

## Cleanup

Old topic branches are not historical archives. Once unique content is verified integrated/superseded and the PR/history preserves evidence, delete the branch when permissions allow.

Recovery evidence that documents genuinely unrecoverable historical source remains preserved under explicit evidence paths and Git history.

## Safety invariants

- no automatic merge to `main`;
- no automatic release/tag/package publication;
- least-privilege GitHub Actions;
- immutable third-party Action pins;
- no hidden failed checks;
- no protection weakening to unblock development;
- no arbitrary profile/downloaded execution.
