# Security baseline for canonical development

- GitHub Actions default to read-only repository access.
- Checkout credentials are not persisted in validation workflows.
- Third-party Actions used by security/release-critical validation are pinned to immutable commit SHAs.
- Validation/recovery jobs do not publish source, refs, tags or packages.
- No workflow force-pushes or directly pushes `main` or `integration/canonical`.
- Release/tag/package publication remains separate from ordinary validation and requires explicit final authorization.
- Profiles are declarative/bounded and may not carry arbitrary executable commands.
- Filesystem/profile/transaction boundaries reject unsafe traversal/symlink/ambiguous input where applicable.
- User-local install/uninstall must not silently broaden permissions or delete recovery/profile data.
- External/mutable inputs require provenance and integrity controls before any trusted use.
- Secrets must not be exported into arbitrary build steps when narrower authentication is possible.

## Repository rule state

### integration/canonical — PASS

An active repository ruleset now enforces:
- exact target `integration/canonical`;
- pull-request-based integration;
- required up-to-date checks: `Canonical CI`, `Race verification`, `Workflow write guard`;
- required conversation resolution;
- squash merge method;
- force pushes blocked;
- branch deletion blocked;
- zero bypass actors.

### main — PASS

PR #75 removed the legacy write-capable bootstrap/finalize/beta/recovery/sync/tag workflows from `main` and replaced them with the reviewed read-only validation set. A temporary draft PR from `integration/canonical` to `main` (#76) then proved all four release-line checks PASS on the full canonical candidate and was closed without merge.

`main` now has an active repository ruleset enforcing:
- pull requests;
- strict/up-to-date `Canonical CI`, `Race verification`, `Workflow write guard`, and `Publication verification`;
- required conversation resolution;
- squash-only merging;
- force-push protection;
- deletion protection;
- zero bypass actors.

The active branch inventory contains only `main` and `integration/canonical`, and automatic deletion of merged head branches is enabled.

Repository-wide **Security B: PASS** for the current stabilization boundary.

No code or workflow workaround is considered equivalent to repository-enforced protection.
