# Consolidation checklist

- [x] Establish `integration/canonical` as the single consolidation line.
- [x] Inventory all known branches and pull requests.
- [x] Classify every pre-existing PR as integrated, superseded, recovery evidence, or active.
- [x] Centralize all unique validated workflow/security changes through PR #37.
- [x] Preserve recovery evidence without pretending incomplete bytes are application source.
- [x] Re-test the older three-part source archive using only GitHub-hosted bytes; exact SHA-256 validation failed.
- [x] Confirm there are no additional `source.part.03+` files or additional readable import patch parts in accessible Git history.
- [x] Pin active third-party Actions used by publication/race validation.
- [x] Remove obsolete write-publisher tombstone workflows from the current canonical tree.
- [x] Remove obsolete trigger/READY scaffolding from the current canonical tree.
- [x] Collapse repetitive `CONSOLIDATION_*` documentation into a small canonical set.
- [x] Preserve historical source fragments under explicit recovery evidence paths or existing validated source-package paths.
- [ ] Close superseded open pull requests after this final cleanup is incorporated.
- [ ] Delete superseded branch refs after their content is proven represented and an authorized branch-deletion mechanism is available.
- [ ] Begin clean application-source rebuild from `integration/canonical`.

Unchecked items must remain visible; they are not implicit PASS.
