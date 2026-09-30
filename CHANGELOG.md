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
- Foundation milestone #40 completed; final product readiness continues under #47/#6.
- Current work is temporarily frozen to correctness/security/CI/docs/cleanup before feature expansion resumes.
- Branch protection for `integration/canonical` remains an administrative release blocker until enabled and reverified.

### Historical recovery
- The complete original application source was not recoverable from surviving GitHub bytes.
- Historical fragments and failed recovery evidence remain preserved explicitly and are not presented as reconstructed source.

## 0.9.0

Verified historical public baseline retained for release/history reference.
