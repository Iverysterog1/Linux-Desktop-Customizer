# Changelog

All notable project changes should be recorded here. Development branch names are not releases by themselves.

## Unreleased

### Rebuilt foundation
- Re-established a dependency-free Go application foundation with CLI and graphical UI entry points.
- Added capability/adapters architecture, durable transactions, preview, post-apply read-back validation and rollback/unapply.
- Added bounded declarative portable profiles and routed the real CLI profile path through the reviewed parser.
- Added KDE/Plasma detection/preflight and conditional runtime registration of the reviewed KDE color-scheme adapter.
- Added English + Portuguese (pt-PT) UI foundation and accessibility behavior.
- Added user-local install/uninstall permission and data-preservation validation.
- Restored least-privilege CI with unit/integration, vet, race, build and installer checks.

### Stabilization
- Adopted the BSD Zero Clause License (0BSD) for unrestricted use, modification and distribution subject to the license disclaimer.
- Foundation milestone #40 completed; final product readiness continues under #47/#6.
- The temporary stabilization freeze completed successfully on 2026-09-30; controlled product development may resume.
- Activated repository rules protecting `integration/canonical` with required PR/check enforcement and no bypass actors.
- Removed the legacy write/force-push/tag workflow surface from `main` via security-only PR #75.
- Proved all four release-line checks PASS on the full canonical candidate using temporary PR #76, then closed it without merge.
- Activated the final `main` ruleset with strict required release-line checks, force-push/deletion blocking and no bypass actors.
- Removed all stale topic branches; only `main` and `integration/canonical` remain, with automatic merged-head branch deletion enabled.
- Security B moved to PASS for the current stabilization boundary.

### Historical recovery
- The complete original application source was not recoverable from surviving GitHub bytes.
- Historical fragments and failed recovery evidence remain preserved explicitly and are not presented as reconstructed source.

## 0.9.0

Verified historical public baseline retained for release/history reference.
