# Consolidation inventory

This inventory prevents branch count from being mistaken for canonical state.

Snapshot: 2026-09-28, canonical base `9cc15cd76a6ce7c50c5f3572d92c6e2dc8963328`.

## Measured repository state

- 39 named remote branches and 34 pull-request refs were reviewed. Every pull-request ref is an exact alias of a named head; there is no additional hidden content in those refs.
- 34 pull requests exist: 33 open, one closed, none merged. Eighteen target `main`; the active consolidation changes target `integration/canonical`.
- `main` and `effects-composer-beta` are ancestors of the canonical line and have no commits ahead of it.
- The pre-convergence canonical tree still had ten write-capable workflows. The dedicated convergence branch replaces all ten with read-only tombstones and carries one final regression guard; combined CI remains required.

## Canonical

- `main`: stable default branch; no direct writes.
- `integration/canonical`: sole integration line during consolidation.

## Historical product/recovery branches

- `effects-composer-beta`: ancestor only; preserve as historical provenance.
- `effects-import`: one exclusive placeholder/evidence commit; no verified product source.
- `publish-0.11-effects-composer`: ten exclusive encoded fragments plus an inspection workflow; reconstruction is truncated and remains `BLOCKED`.

Preserve until source provenance/recovery decisions are complete. They are not automatically canonical merely because they contain newer-looking material.

## Documentation / quality branches

- `docs/github-page-refresh` / PR #4: six documentation commits require a no-loss review before cleanup.
- `quality/validate-preserved-source-package` / PR #5: the complete validator remains exclusive; its recorded run is `FAIL` because the XZ stream is truncated.
- `quality/validate-historical-beta-blob-archive` / PR #9: the complete 56-blob validator remains exclusive; its recorded run is `FAIL` on missing blob `cfd027fdf53fcd403a020f7a934b74d0febdc98c`.

The canonical line represents the blocker outcomes, but not every byte of the full validators or the six-document branch. Original branches therefore remain evidence until an explicit preservation decision is reviewed.

## Security branches

The older `security/*` branches based on `main` are superseded by canonical-based equivalents. PRs #22, #26, #28, #31, #32, #34, #35 and #36 supplied CI-validated final workflow versions; PRs #23 and #24 supplied statically validated workflow versions whose runs failed only in the inherited broken guard. The convergence branch combines those ten workflow remediations with the final proven guard from #36 so they can be tested once as a coherent set.

## Cleanup rule

A branch is eligible for deletion only when its unique work is confirmed integrated or intentionally superseded and its PR/history remains sufficient for audit. Consolidation itself does not authorize deletion or merge to `main`.
