# Consolidation checklist

- [x] Establish `integration/canonical` as the single consolidation line.
- [x] Inventory all known branches and pull requests.
- [x] Classify every pre-existing PR as integrated, superseded, recovery evidence, or active.
- [x] Centralize every unique validated security/workflow change.
- [x] Preserve recovery evidence without promoting incomplete bytes to application source.
- [x] Re-test the older three-part source archive using only GitHub-hosted bytes.
- [x] Confirm there are no additional `source.part.03+` files or readable import patch parts in accessible history.
- [x] Incorporate immutable Action pinning into active validation workflows.
- [x] Remove obsolete publisher tombstones and READY trigger scaffolding from the current canonical tree.
- [x] Collapse repetitive consolidation documentation.
- [x] Preserve historical fragments under explicit evidence paths.
- [x] Close superseded pull requests while retaining their audit history.
- [x] Delete superseded branch refs after no-loss classification.
- [x] Verify only `main` and `integration/canonical` remain.
- [x] Close the unrecoverable historical-source P0 and create rebuild P0 #40.

## Result

**Centralization COMPLETE.**

Next phase: rebuild a complete, testable application source tree from `integration/canonical`.
