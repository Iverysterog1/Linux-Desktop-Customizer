# Project status

Date: 2026-09-30

## Current phase

The clean rebuild foundation milestone is complete; Issue #40 is closed. The project remains under a temporary **stabilization freeze** coordinated in Issue #47 before feature expansion resumes.

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
- unit/integration, vet, race, build and installer checks in least-privilege CI;
- 0BSD software license;
- active repository ruleset protecting `integration/canonical`.

## Not final yet

The project is not release-ready. Outstanding gates include:
- equivalent required-check protection for `main`;
- cleanup/revalidation of superseded topic branches;
- real Plasma end-to-end apply/visible validation/rollback evidence;
- complete graphical preview/apply/history/undo flow;
- DEB/RPM and trusted terminal release installer;
- clean real-Linux install/run/uninstall validation;
- real application screenshots;
- final release-candidate security/supply-chain validation.

## Security

Security A is currently PASS on the integrated code boundary.

The previous `integration/canonical` protection finding is resolved: an active ruleset now requires pull requests, up-to-date required checks (`Canonical CI`, `Race verification`, `Workflow write guard`), conversation resolution, blocks force pushes and deletion, and has no bypass actors.

Repository-wide Security B remains **FAIL — MEDIUM** until `main` receives the equivalent release-line rules with `Publication verification` additionally required and is reverified.

Software licensing is settled as **0BSD**. No release or promotion to `main` is authorized.
