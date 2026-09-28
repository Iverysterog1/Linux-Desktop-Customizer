# Linux Desktop Customizer — project status

## Canonical state

The repository is undergoing source recovery and consolidation. The verified public baseline is 0.9.0. Preserved 0.11 Effects Composer material is development/recovery evidence and must not be represented as a verified final release until reconstruction and integrity validation pass.

`main` remains the stable default branch. During consolidation, `integration/canonical` is the single integration line. The branch/PR inventory and no-loss classification are complete in PR #37, but feature development remains paused until that convergence is reviewed and authoritative application source is recovered.

## Release discipline

Do not publish an installer, package, tag, or release from incomplete recovered material. A release candidate must have a complete source tree, reproducible provenance, passing tests/builds, packaging validation, privacy/security review, and explicit human publication approval.

## Current recovery blockers

Two preserved recovery routes have independent integrity blockers. The ten-part 0.11 stream decodes to a 120,000-byte XZ archive with SHA-256 `321faa7210e54178d33da1ea9f194df67e64d1ea9ff1e10a0e066010b779b1fd`, but strict XZ validation reports unexpected end of input. The 56-object beta archive cannot be reconstructed because authoritative blob `cfd027fdf53fcd403a020f7a934b74d0febdc98c` is unavailable. Neither failed input is canonical application source; no stream may be spliced or guessed.

## Safety

Direct or force writes to `main` are prohibited by project governance. PR #37 proposes the consolidated read-only retirement of the remaining legacy writers. Until it is reviewed and integrated, the old canonical base must still be treated as write-capable.
