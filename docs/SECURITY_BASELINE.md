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

### main — pending

`main` is marked protected by GitHub, but the currently observable legacy protection reports required-status-check enforcement off and there is no repository ruleset for it yet.

Before FINAL READY, `main` must have repository-enforced:
- pull requests;
- `Canonical CI`, `Race verification`, `Workflow write guard`, and `Publication verification`;
- force-push protection;
- deletion protection;
- no broad bypass.

Until the `main` release-line rules are applied and reverified, repository-wide **Security B remains FAIL — MEDIUM**.

No code or workflow workaround is considered equivalent to repository-enforced protection.
