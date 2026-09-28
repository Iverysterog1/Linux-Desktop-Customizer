# Pull-request consolidation map

Snapshot: 2026-09-28. Classification is based on exact commit/file comparison, not branch names.

## Represented or preserved by the convergence branch

- #4: no-loss review complete. Recovery-state README plus applicable changelog, metadata, release checklist and media policy detail were retained; its older status rewrite is superseded by the current canonical status and inventory.
- #5: strict ten-part reconstruction, XZ/path/source/build validator absorbed into the canonical workflow/script names. It remains fail-closed on the known truncated stream.
- #9: full 56-blob manifest and reconstruction/test procedure preserved as a read-only manual recovery script; no token-bearing Actions job was introduced. The canonical status check remains fail-closed on the missing authoritative blob.
- #19: reviewed GitHub Actions dependency-update configuration.
- #20: three legacy main/tag writer retirements plus regression protection.
- #21: repository-wide changed-workflow write/direct-push regression guard.

## Selected for the single workflow convergence

- #22: `publish-clean-beta.yml`.
- #23: `publish-clean-beta-v2.yml`.
- #24: `publish-beta-stream.yml`.
- #26: `finalize-beta.yml`.
- #28: `build-beta-cleanup-branch.yml`.
- #31: `recover-current-source.yml`.
- #32: `race-verification.yml`.
- #34: `finalize-from-tunnel.yml`.
- #35: `bootstrap.yml`.
- #36: `publication-gate.yml` and the final regression guard.

PRs #22, #26, #28, #31, #32, #34, #35 and #36 have a passing guard run at their selected heads. PRs #23 and #24 failed because the then-current canonical guard had an `unexpected EOF`; their selected workflow files pass the final static guard and require combined CI here.

## Superseded security branches

- Old `main`-based PRs #1, #3, #7, #8, #10-#18 are superseded by the canonical-based selections above or were already represented by #20.
- #25 is superseded by #36; #27 by #34; #29 by #35; #30 contributes no workflow beyond #36; #33's guard is superseded by the final guard carried by #36.

Superseded does not mean deletable. The original PRs and branches remain audit evidence until the convergence PR is reviewed, recovery blockers are resolved or formally accepted, and cleanup is separately authorized.

## No destructive cleanup yet

PR closure and branch deletion are deliberately deferred until the canonical line contains every unique remediation and the resulting diff/CI has been reviewed. This prevents branch-count cleanup from causing loss of work.
