# Security baseline for canonical development

- GitHub Actions default to read-only repository access.
- Checkout credentials are not persisted unless a reviewed workflow has a documented need.
- Recovery/validation jobs do not publish source or refs.
- No workflow force-pushes or directly pushes `main`.
- Release/tag/package publication is separated from ordinary validation and requires explicit authorization.
- New Actions should prefer immutable commit references for security-sensitive controls.
- Dependency updates remain reviewable; no automatic merge is enabled by consolidation.
- External downloads used in build/recovery paths require provenance and integrity controls and must not flow directly into privileged publication.
- Secrets must not be exported into arbitrary build steps when narrower authentication is possible.
