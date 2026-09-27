# Consolidation decision log

## 2026-09-27

Decision: stop feature expansion temporarily and consolidate repository work before further development.

Reason: the repository accumulated many overlapping security/recovery/documentation branches and pull requests. Continuing to add independent branches would increase uncertainty over which tree represents the intended product.

Action: establish `integration/canonical` from the exact current `main` head; absorb work by objective; preserve failed recovery evidence; do not merge to `main` or delete historical branches during the evidence-gathering phase.

Exit condition: all unique open work has a documented disposition and the canonical integration PR has reviewable validation evidence.
