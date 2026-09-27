# Release checklist

A release is allowed only after all applicable gates pass.

- canonical source provenance and version are consistent;
- source recovery/integrity checks pass where recovery material is involved;
- tests, vet/static checks, and builds pass;
- installer/uninstaller and desktop integration are validated;
- icons and AppStream/metainfo are present and correct where applicable;
- dependency and supply-chain review is complete;
- privacy behavior and network access are documented and reviewed;
- packaging is reproducible enough to identify exactly what source produced the artifact;
- no workflow requires direct or force writes to `main`;
- release/tag/package publication requires explicit human authorization;
- public documentation, screenshots, version labels, and changelog match the actual shipped build.

No failed integrity gate may be bypassed by combining unrelated historical source streams without byte-level provenance.
