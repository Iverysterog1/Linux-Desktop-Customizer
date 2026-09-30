# Project status

Date: 2026-09-30

## Current phase

The clean rebuild foundation milestone is complete; Issue #40 is closed. The project is now under a temporary **stabilization freeze** coordinated in Issue #47 before feature expansion resumes.

`integration/canonical` is the sole integrated development truth. `main` remains the stable/release line.

## Implemented foundation

Current canonical includes:
- dependency-free Go module;
- CLI and graphical UI entry points;
- capability/adapter registry;
- bounded local-file and KDE color-scheme adapters;
- declarative bounded portable profiles wired into the real CLI path;
- preview, snapshot/journal, apply, read-back validation, rollback/unapply;
- KDE/Plasma session preflight and conditional KDE runtime registration;
- EN + pt-PT UI foundation and accessibility support;
- user-local install/uninstall baseline with permission and data-preservation tests;
- unit/integration, vet, race, build and installer checks in least-privilege CI.

## Not final yet

The project is not release-ready. Outstanding gates include:
- real Plasma end-to-end apply/visible validation/rollback evidence;
- complete graphical preview/apply/history/undo flow;
- DEB/RPM and trusted terminal release installer;
- clean real-Linux install/run/uninstall validation;
- real application screenshots;
- final security/supply-chain validation;
- branch protection on `integration/canonical`.

## Security

Security A is currently PASS on the integrated code boundary.

Security B remains **FAIL — MEDIUM** until repository administration protects `integration/canonical` with PR-required integration, required checks, and force-push/delete prevention.

Software licensing is now settled as **0BSD**. No release or promotion to `main` is authorized.
