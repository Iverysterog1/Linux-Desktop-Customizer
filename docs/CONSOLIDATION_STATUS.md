# Consolidation status

Date: 2026-09-28

## State

**Canonical content consolidation: COMPLETE.**
**Legacy pull-request/branch cleanup: in progress.**

The only authoritative storage for this project is GitHub. There is no separate PC or external source tree to recover from.

- Stable branch: `main` at `56fa162d076cbe96c3ebdca828f480070293708a` (untouched by consolidation).
- Sole integration/source-of-truth line: `integration/canonical`.
- PR #37 centralized every pre-existing classified workflow remediation, validator, documentation contribution, and preserved recovery fragment.
- The immutable Action pinning from PR #38 is incorporated into the final cleanup line.

## Recovery truth

The original complete application source cannot be reconstructed from the GitHub bytes currently available:

1. The 56-object historical archive is missing authoritative blob `cfd027fdf53fcd403a020f7a934b74d0febdc98c`.
2. The ten-part preserved XZ stream is truncated and fails integrity validation.
3. The older three-part `bootstrap/source.part.*.b64` path was re-tested on GitHub Actions run `36413607567`; the reconstructed archive does **not** match its pinned SHA-256 `444f8803b764e2e7f709185428d1b519a44cd902291c40aef891cee8992413bd`.
4. Git history contains no `source.part.03+` and only one readable `import-parts/part-000.patch`.

These failures are preserved as evidence. No bytes are invented or silently substituted.

## Decision

Do not keep waiting for a non-existent off-GitHub copy. Historical fragments are retained as evidence, but the next product phase is a **clean rebuild of the missing application source from the surviving specifications, README/history, patches and validated project requirements**.

No release or promotion to `main` is authorized until that rebuilt source has tests, race checks, builds, installer/uninstaller checks and functional validation.
