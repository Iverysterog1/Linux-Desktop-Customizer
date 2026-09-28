# Linux Desktop Customizer

> **Created by one. Improved by many. Available to all.**

![Status](https://img.shields.io/badge/status-canonical%20consolidated-brightgreen)
![Next](https://img.shields.io/badge/P0-application%20rebuild-orange)
![Safety](https://img.shields.io/badge/repository-read--only%20CI-blue)

**Linux Desktop Customizer** is a privacy-first open-source project for personalizing the Linux desktop from one place: themes, wallpapers, icons, cursors, windows, panels, effects, sounds and other desktop appearance settings.

The product direction remains simple: **preview → review → apply → undo**.

> [!IMPORTANT]
> Repository consolidation is complete on `integration/canonical`, but the complete original application source did not survive in the GitHub data that remains. The project is therefore rebuilding a clean, testable application foundation rather than pretending partial archives are valid source.

## Repository state

- **`integration/canonical`** — sole active source-of-truth and rebuild line.
- **`main`** — preserved stable historical branch; no direct development or force-push.
- Legacy publisher workflows and trigger scaffolding have been removed from the current canonical tree.
- Historical recovery fragments are preserved only as explicit evidence under `recovery-evidence/`, `source-package/`, and Git history.

See [Project status](docs/PROJECT_STATUS.md) and [Consolidation status](docs/CONSOLIDATION_STATUS.md).

## What happened to the original source?

Every surviving GitHub-only recovery route was checked:

- the historical 56-object archive is missing one authoritative Git blob;
- the preserved ten-part XZ stream is truncated;
- the older three-part `source.part.*` archive fails its pinned SHA-256;
- accessible Git history contains no later `source.part.03+` and only one readable import patch part.

There is no separate PC or external authoritative copy. The failures are documented in [Recovery evidence](docs/RECOVERY_EVIDENCE.md).

## Current P0

The current P0 is [Issue #40: rebuild complete application source on canonical foundation](../../issues/40).

The first rebuild milestone is deliberately small and verifiable:

1. restore a coherent Go module;
2. restore CLI and UI entry points;
3. establish capability/adapters boundaries;
4. implement safe snapshot, transaction journal and rollback primitives;
5. restore installer/uninstaller paths;
6. run unit/integration tests, `go vet`, race detector and Linux builds;
7. only then resume broader desktop customization features.

Issue #6 remains the product/UX contract for the first KDE full-desktop-transformation vertical slice after the foundation is buildable.

## Product direction

The preserved product direction includes:

- one local interface for Linux desktop personalization;
- original signature themes;
- Creator Forge for complete theme creation;
- Visual Composer for windows, panels, docks, widgets, typography and wallpapers;
- Effects Composer for advanced visual effects;
- Harmony Guard for coherent palette/visual choices;
- Adapter Forge for supported desktop/compositor backends;
- reviewed changes before apply;
- transaction history with rollback/undo;
- local-first operation with no telemetry requirement;
- defensive handling of downloaded content.

These are requirements, not claims that the current repository already ships a working build.

## Installation

There is currently **no supported installable build** from the reconstructed canonical line.

Packaging resumes only after the new source foundation passes build, test, installer/uninstaller and rollback validation. No release should be published from historical fragments or incomplete recovery archives.

## Safety and privacy

Linux Desktop Customizer is intended to keep personalization reversible and inspectable.

- No arbitrary theme install scripts should execute automatically.
- Downloaded content should be inspected before use.
- Unsafe archive paths and path traversal must be rejected.
- Changes should be previewed/reviewed before apply and recorded for rollback.
- Diagnostics should avoid leaking secrets.
- Telemetry is not required for the application to function.
- CI should use least privilege and immutable Action pins.

See [Security testing policy](SECURITY_TESTING_POLICY.md) and [Security baseline](docs/SECURITY_BASELINE.md).

## Contributing

New implementation work starts from **`integration/canonical`** and should use a short-lived topic branch with visible validation. Do not write or force-push directly to `main`.

Security-sensitive reports must not include credentials, private files or secrets.

## Credits

**Created by JP99** — concept, product direction and testing.

AI engineering assistance has been used during development, repository consolidation and reconstruction.
