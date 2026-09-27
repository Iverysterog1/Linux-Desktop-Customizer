# Branch governance

## Canonical flow

- `main` is the stable default branch. Never push directly to it.
- `integration/canonical` is the single integration line while repository recovery and consolidation are in progress.
- New implementation work should branch from the current canonical integration head and return through a reviewed pull request.
- Short-lived topic branches are preferred. Do not create one branch per tiny remediation when changes belong to the same reviewable objective.

## Consolidation rule

Historical/recovery/security branches are evidence until their unique commits are classified as integrated, superseded, blocked, or intentionally archived. They must not be deleted merely to reduce the branch count.

Overlapping legacy-workflow security changes should be consolidated into the canonical integration line before further feature development. Recovery validators remain evidence and must not be treated as canonical application source until their integrity gates pass.

## Safety invariants

- no force-push or direct push to `main`;
- no automatic merge;
- no automatic release/tag/package publication;
- least-privilege GitHub Actions permissions;
- no credential rotation or protection weakening as part of consolidation;
- preserve provenance for recovered source and historical artifacts.

## Development restart gate

Feature development resumes only after the open branch/PR inventory has been classified and the canonical integration PR contains the selected non-conflicting work with validation evidence.
