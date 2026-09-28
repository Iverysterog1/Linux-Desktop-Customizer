# Consolidation inventory

Final snapshot: 2026-09-28.

## Live branches

Exactly two branches remain:

- `main` — preserved stable historical branch.
- `integration/canonical` — sole active source-of-truth and rebuild line.

All other reviewed topic, security, recovery, documentation and historical working branches were deleted after their unique useful content was integrated, preserved as evidence, or intentionally superseded. Branch cleanup completed successfully in Actions run `36414812932`.

## Pull requests

There are **0 open pull requests**.

PR #37 supplied the major classified convergence. The final cleanup and immutable Action pinning are represented in canonical. Older PR discussions remain available as audit history after closure.

## Current canonical shape

The current tree retains only active project governance/security material plus explicit recovery evidence:

- root project documentation and policies;
- five active validation/security workflows;
- `recovery-evidence/` for legacy fragments/failures;
- `source-package/` for the preserved truncated 0.11 stream;
- recovery/validation scripts.

Obsolete publisher workflows, temporary READY files and duplicated consolidation notes are removed from the current tree.

## Product source

A complete original buildable application source tree is not recoverable from surviving GitHub bytes. The project has moved to issue #40: a clean application rebuild on the canonical foundation.
