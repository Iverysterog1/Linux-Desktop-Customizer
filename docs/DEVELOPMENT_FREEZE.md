# Temporary development freeze

Feature development is temporarily paused while repository branches and overlapping pull requests are consolidated into `integration/canonical`.

Allowed during the freeze: branch/PR inventory, provenance checks, read-only validation, security neutralization, consolidation documentation, and non-destructive CI guards.

Not allowed as part of consolidation: direct writes to `main`, force-push to `main`, merging PRs, release/tag/package publication, credential rotation, weakening protections, or deleting evidence branches before their unique work is classified.

The freeze ends when the canonical integration line has a disposition for all open work and remaining branches are clearly historical, blocked recovery evidence, or active short-lived topics.
