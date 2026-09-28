# Project status

Date: 2026-09-28

## Repository

The repository has one intended working truth: `integration/canonical`. `main` remains stable and is not used for direct development during reconstruction.

Repository security and consolidation work is substantially complete: legacy direct/force publishers are removed from the current tree, active validation workflows are read-only, and the workflow regression guard remains in place.

## Application

The complete original application source is not present in the surviving GitHub bytes. Multiple preserved recovery routes are incomplete and have been validated as such; no external PC/source copy exists.

The project therefore moves from **source recovery** to **clean source rebuild**. Historical fragments, patch material, product documentation and requirements are evidence/reference inputs, not a buildable source tree.

## Release state

No release is authorized. Before any release or promotion to `main`, the rebuilt application must pass:

- unit/integration tests;
- `go vet` and race detector;
- CLI/UI builds;
- installer and uninstaller syntax/behavior checks;
- clean-install and rollback validation;
- functional desktop-adapter tests with failures visible.
