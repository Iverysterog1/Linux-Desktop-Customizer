# Canonical development workflow

During consolidation:

`main` -> `integration/canonical` -> reviewed consolidation PR -> `main`

After consolidation:

`integration/canonical` -> short-lived topic branch -> validation/review -> `integration/canonical`; periodically a reviewed release/integration PR targets `main`.

Never use direct default-branch publication as a substitute for review. Never use force-push as a source-recovery mechanism.
