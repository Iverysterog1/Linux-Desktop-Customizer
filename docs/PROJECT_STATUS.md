# Project status

Last reviewed: 2026-09-26

## Current state

The repository is in a controlled **source-recovery and hardening phase**.

### Verified facts

- The initial public release line is **0.9.0**.
- A newer preserved branch is named **`publish-0.11-effects-composer`**.
- That branch contains ten staged source-package fragments (`part00.b64` through `part09.b64`).
- The combined staged package decodes to XZ data, so it is materially different from the earlier truncated three-part bootstrap archive.
- The current `main` branch is not a normal complete source tree.
- Legacy workflows capable of direct/force writes to `main` are being neutralized through separate pull requests.

## Version policy

Until the recovered package is fully unpacked and validated:

- **0.9.0** is the public baseline.
- **0.11 Effects Composer** is a development/recovery line, **not a published final release**.
- The README must not present an unverified development version as downloadable or production-ready.

## Release blockers

1. Reconstruct and verify the ten-part preserved XZ source package.
2. Confirm the internal version and source provenance.
3. Restore the complete source tree on a dedicated branch.
4. Run tests, vet/static checks and supported race tests.
5. Build CLI and UI targets.
6. Validate installer/uninstaller scripts.
7. Restore verified screenshots/assets.
8. Produce reproducible Linux packages.
9. Test on representative Linux distributions/desktops.
10. Review before merge/release.

## Installation goal

The target user experience is:

**download → install → application menu → launch**

A one-command terminal fallback should also exist, but the primary path must not require cloning the repository or manually assembling dependencies.
