# Source recovery evidence

Two independently investigated recovery routes are currently blocked and must remain distinct.

## Preserved 0.11 stream

The ten preserved source-package fragments reconstruct to an input that fails strict XZ integrity checking as incomplete/truncated. This route is BLOCKED; downstream source tests/builds cannot establish canonicality from an incomplete archive.

## Historical blob-manifest route

A historical manifest references an authoritative Git blob that is not currently retrievable. Reconstruction therefore stops before complete-archive digest/XZ/source validation. This route is also BLOCKED.

## Rule

Do not combine these streams or invent replacement bytes. A future recovery attempt must preserve byte-level provenance and satisfy the expected complete-stream integrity checks before source promotion.
