# Canonical source policy

## Current rule

`integration/canonical` is the sole integrated development truth. `main` is the stable/release line.

A topic branch, old recovery branch, archive, closed PR or generated artifact is never an alternate source of truth.

Every change follows:

`current canonical -> short-lived branch -> validation/security review -> PR -> canonical`

Promotion to `main` is a separate protected PR and requires explicit final authorization.

## Historical-source provenance

The complete original historical application source was not recoverable from surviving GitHub bytes. That conclusion is preserved as evidence and must not be reversed by combining incomplete archives or guessing missing data.

The present application is a clean rebuild based on the preserved product/security contract. It must never be mislabeled as byte-for-byte recovered historical source.

## Evidence

Keep the minimum evidence needed to reproduce important historical/security decisions:
- commit IDs and PR/issue history;
- integrity hashes and failed recovery evidence;
- current validation results and blockers.

Do not keep stale working branches or contradictory current-state documents merely as “history”; Git and closed review records already preserve that history.

## Protection

`integration/canonical` must be repository-protected before FINAL READY. Direct writable canonical refs are not an acceptable substitute for PR/check enforcement.
