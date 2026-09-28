# Pull-request consolidation map

Snapshot: 2026-09-28.

- **PR #37**: canonical convergence. It centralized the selected security/workflow remediations, #4 documentation, #5/#9 validators and historical recovery evidence.
- **PR #38**: immutable Action pinning; incorporated by the final cleanup line.
- **PRs #1, #3-#5, #7-#36**: their unique useful content is represented by PR #37, the final cleanup, or preserved recovery evidence. They are superseded as active development lines.
- Historical recovery PRs remain readable as audit history after closure; closing them does not erase commit/PR history.

No old PR should be treated as an alternate source-of-truth branch. New work starts only from `integration/canonical`.
