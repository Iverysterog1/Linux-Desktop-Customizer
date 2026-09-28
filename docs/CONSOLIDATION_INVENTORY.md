# Consolidation inventory

Snapshot: 2026-09-28.

## Canonical

`integration/canonical` is the sole consolidation/source-of-truth line. `main` remains the untouched stable branch.

## Preserved historical evidence

- `source-package/`: ten exact fragments from the historical 0.11 effects-composer path; known truncated stream, retained for provenance.
- `recovery-evidence/legacy/bootstrap/`: historical payload/source fragments formerly stored at repository root under `bootstrap/`.
- `recovery-evidence/legacy/import-parts/part-000.patch`: the only readable import patch found in accessible history.
- `recovery-evidence/legacy/beta-archive-RESULT.txt`: preserved failure record from the old archive reconstruction attempt.
- `scripts/recovery/validate-historical-beta-blobs.sh` and validation scripts: retained as reproducible evidence.

## Removed from the current tree

Historical publisher workflows, one-shot READY markers, obsolete synchronization/release triggers, and repetitive consolidation notes are removed from the current tree after their meaning is captured by canonical documentation and Git history.

## Pull requests

PR #37 is the convergence commit that represented all earlier classified work. PR #38 contributes immutable Action pinning and is represented by the final cleanup. Older PRs remain useful only as audit history and are eligible to be closed once the final cleanup lands.

## Product source

A complete buildable application tree is **not** present. This is no longer described as an unresolved external recovery dependency: GitHub is the only source and the available recovery inputs are demonstrably incomplete. The next phase is a clean rebuild.
