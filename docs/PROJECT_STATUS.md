# Linux Desktop Customizer — project status

## Canonical state

The repository is undergoing source recovery and consolidation. The verified public baseline is 0.9.0. Preserved 0.11 Effects Composer material is development/recovery evidence and must not be represented as a verified final release until reconstruction and integrity validation pass.

`main` remains the stable default branch. During consolidation, `integration/canonical` is the single integration line. Feature development is paused until overlapping branches and pull requests have been classified and safely consolidated.

## Release discipline

Do not publish an installer, package, tag, or release from incomplete recovered material. A release candidate must have a complete source tree, reproducible provenance, passing tests/builds, packaging validation, privacy/security review, and explicit human publication approval.

## Current recovery blockers

Two preserved recovery routes have independent integrity blockers documented by their validation pull requests. Neither failed input is canonical application source. Historical material is preserved as evidence rather than silently spliced or reconstructed by guesswork.

## Safety

Direct or force writes to `main` are prohibited by project governance. GitHub Actions should use least privilege, avoid persisted credentials when unnecessary, and route repository changes through reviewed pull requests.
