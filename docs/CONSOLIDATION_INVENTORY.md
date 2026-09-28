# Consolidation inventory

This inventory prevents branch count from being mistaken for canonical state.

Snapshot: 2026-09-28, canonical base `9cc15cd76a6ce7c50c5f3572d92c6e2dc8963328`.

## Measured repository state

- 39 pre-existing named remote branches and 34 pre-existing pull-request refs were reviewed. Every reviewed pull-request ref was an exact alias of a named head. The dedicated convergence branch and PR #37 bring the known current totals to 40 named remote branches and 35 pull requests.
- 35 pull requests exist: 34 open, one closed, none merged. Eighteen target `main`; the remaining PRs, including #37, target `integration/canonical`.
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

- `docs/github-page-refresh` / PR #4: no-loss review completed. Its recovery-state README plus applicable changelog, metadata, release-checklist and media-policy detail is integrated; its older status rewrite is superseded by the more current canonical status and inventory.
- `quality/validate-preserved-source-package` / PR #5: complete strict validation logic is integrated under the canonical hyphenated names; the workflow remains fail-closed and is expected to `FAIL` while the XZ stream is truncated.
- `quality/validate-historical-beta-blob-archive` / PR #9: the 56-blob manifest and reconstruction/test procedure is preserved in `scripts/recovery/validate-historical-beta-blobs.sh`. It is deliberately not wired to a token-bearing workflow; the canonical status workflow remains fail-closed on missing blob `cfd027fdf53fcd403a020f7a934b74d0febdc98c`.

The original branches remain audit evidence until cleanup is separately reviewed, but they no longer contain unclassified documentation or validator logic.

## Security branches

The older `security/*` branches based on `main` are superseded by canonical-based equivalents. PRs #22, #26, #28, #31, #32, #34, #35 and #36 supplied CI-validated final workflow versions; PRs #23 and #24 supplied statically validated workflow versions whose runs failed only in the inherited broken guard. The convergence branch combines those ten workflow remediations with the final proven guard from #36 and the separately preserved #4/#5/#9 no-loss material.

## Cleanup rule

A branch is eligible for deletion only when its unique work is confirmed integrated or intentionally superseded and its PR/history remains sufficient for audit. Consolidation itself does not authorize deletion or merge to `main`.
