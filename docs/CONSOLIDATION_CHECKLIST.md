# Consolidation checklist

- [x] Create one canonical integration branch from current `main`.
- [x] Define branch/source governance and feature freeze.
- [x] Centralize dependency-update configuration.
- [x] Centralize changed-workflow write/direct-push regression guard.
- [x] Absorb the #20 legacy main/tag writer retirement family.
- [x] Preserve #5/#9 recovery blockers without promoting corrupt/incomplete input.
- [ ] Absorb every unique remediation from #1/#3/#7/#8/#10-#18 not already represented.
- [ ] Compare final canonical diff against every open PR and classify each as absorbed, superseded, blocked, or still active.
- [ ] Obtain CI evidence for the canonical integration PR.
- [ ] Only then perform separate, explicit PR/branch cleanup.
- [ ] Resume feature development from the canonical line.

No unchecked item may be silently treated as complete.
