# Consolidation checklist

- [x] Create one canonical integration branch from current `main`.
- [x] Define branch/source governance and feature freeze.
- [x] Centralize dependency-update configuration.
- [x] Centralize changed-workflow write/direct-push regression guard.
- [x] Absorb the #20 legacy main/tag writer retirement family.
- [x] Preserve #5/#9 recovery blockers without promoting corrupt/incomplete input.
- [x] Assemble every unique legacy workflow remediation into one dedicated branch based on `integration/canonical`.
- [x] Obtain truthful combined CI for the ten read-only workflow replacements, final guard and fail-closed recovery checks; source-recovery failures remain `BLOCKED`, not hidden.
- [x] Preserve the complete validator logic from #5/#9 and complete the no-loss review of #4.
- [x] Compare the canonical diff against every open PR and classify each as absorbed, superseded, blocked, or still active.
- [x] Resolve or formally preserve the three historical recovery inputs without substituting unauthoritative bytes.
- [ ] Recover the authoritative historical blob and validate the 56-object archive, XZ stream, source tree, build, tests, and installer.
- [ ] Only then perform separate, explicit PR/branch cleanup.
- [ ] Resume feature development from the canonical line.

No unchecked item may be silently treated as complete.
