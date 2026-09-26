# Linux Desktop Customizer

> **Created by one. Improved by many. Available to all.**

![Status](https://img.shields.io/badge/status-source%20recovery-orange)
![Public baseline](https://img.shields.io/badge/public%20baseline-0.9.0-blue)
![Development line](https://img.shields.io/badge/development-0.11%20Effects%20Composer-purple)

**Linux Desktop Customizer** is a privacy-first open-source project for personalizing the Linux desktop from one place: themes, wallpapers, icons, cursors, windows, panels, effects, sounds and other desktop appearance settings.

The product direction is deliberately simple: **preview → review → apply → undo**.

> [!IMPORTANT]
> The current `main` branch is a **source-recovery workspace**, not a complete release tree.  
> Do not treat the files currently visible on `main` as an installable build.

## Version status

| Line | Status | Meaning |
|---|---|---|
| **0.9.0** | Public baseline | Initial public release line preserved in Git history. |
| **0.11 Effects Composer** | Development / recovery | Newer preserved development line. Its source package is being reconstructed and validated before publication. |
| **main** | Recovery workspace | Contains recovery/security scaffolding while the normal source tree is restored. |

The project previously displayed `0.9.0-visual-composer` as if it were the current installable version. That was misleading while `main` did not contain the complete source tree. The repository now separates **released/public baseline** from **newer development work**.

See [Project status](docs/PROJECT_STATUS.md) for the current recovery and release gates.

## What the project is designed to include

- One local interface for Linux desktop personalization.
- Original signature themes.
- **Creator Forge** for building complete themes from a visual direction.
- **Visual Composer** for shaping windows, panels, docks, widgets, typography and wallpapers.
- **Effects Composer** for advanced visual-effect configuration.
- **Harmony Guard** for coherent palettes and visual choices.
- **Adapter Forge** for mapping compositions to supported desktop/compositor backends.
- Reviewed changes before apply.
- Transaction history with rollback/undo.
- Local-first operation with no telemetry requirement.
- Defensive handling of downloaded theme content.

These capabilities describe the preserved product direction. They will be marked individually as validated once the canonical source tree is restored and tested.

## Screenshots and images

Real application screenshots and theme previews are being restored from the preserved source package.

**We do not publish fabricated UI screenshots or broken placeholders.** Images will return to this README only after they are recovered from the verified source tree or captured from a reproducible build.

The image restoration rules are documented in [docs/media/README.md](docs/media/README.md).

## Installation

### Current status

There is **no supported one-click installer from the current `main` recovery workspace yet**.

The release target is:

1. Download the Linux package.
2. Open/install it normally.
3. Find **Linux Desktop Customizer** in the application menu.
4. Launch and use it without manual repository setup.

A terminal-based fallback will also be provided for users who prefer it.

Packaging work starts only after the canonical source tree passes recovery, build and installer validation. This avoids shipping an installer built from incomplete source.

## Recovery and validation gates

Before the next installable beta is published, the project must pass:

```text
1. Recover the complete canonical source tree
2. Verify source-package integrity and provenance
3. Restore normal project files and assets
4. Run unit tests and static checks
5. Build CLI and graphical UI targets
6. Validate installer/uninstaller behavior
7. Test representative Linux desktop environments
8. Restore verified screenshots and release metadata
9. Produce reproducible Linux packages
10. Publish a reviewed beta
```

## Safety and privacy

Linux Desktop Customizer is intended to keep personalization reversible and inspectable.

- No arbitrary theme install scripts should execute automatically.
- Downloaded content should be inspected before use.
- Unsafe archive paths and path traversal must be rejected.
- Changes should be reviewed before apply and recorded for rollback.
- Diagnostics should avoid leaking secrets.
- Telemetry is not required for the application to function.

Current repository security work is also removing legacy GitHub Actions paths that could write or force-push directly to `main`.

See [Security testing policy](SECURITY_TESTING_POLICY.md).

## Repository status

The active priorities are:

- **P0:** recover and validate the complete canonical source tree;
- neutralize legacy workflows that can mutate `main`;
- restore the normal source/assets layout;
- restore CI and build verification;
- make Linux installation extremely simple;
- restore verified screenshots, version metadata and release documentation.

For the exact state, see [docs/PROJECT_STATUS.md](docs/PROJECT_STATUS.md).

## GitHub presentation metadata

The canonical project description, suggested topics and social-preview requirements are tracked in [docs/GITHUB_METADATA.md](docs/GITHUB_METADATA.md) so the repository page stays consistent with the actual product state.

## Contributing

Until source recovery is complete, changes should use **branch → validation → pull request**. Do not force-push or write directly to `main`.

Security-sensitive issues should not be posted with secrets or credentials.

## Credits

**Created by JP99** — concept, product direction and testing.

AI engineering assistance has been used during development and repository recovery.
