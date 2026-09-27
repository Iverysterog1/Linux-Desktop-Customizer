# Canonical source policy

Canonical does not mean newest branch or largest archive. A source tree becomes canonical only when its provenance and integrity are established and its required validation gates pass.

During the current recovery period:

- `main` is the stable repository baseline;
- `integration/canonical` is the sole integration line for reviewed consolidation work;
- incomplete/truncated archives and manifests with unavailable objects are evidence, not source candidates for silent promotion;
- never combine unrelated recovery streams merely to make an archive extract successfully;
- retain commit IDs, hashes, validation output, and blockers needed to reproduce recovery decisions.

Once a complete source tree is recovered and validated, its promotion must occur through a reviewed PR rather than a workflow that replaces `main` history.
