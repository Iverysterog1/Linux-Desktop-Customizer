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

## Repository rule blocker

`integration/canonical` must be protected before FINAL READY:
- pull-request-based integration;
- required validation/status checks appropriate to changed scope;
- force pushes disabled;
- branch deletion disabled.

The current GitHub connection used by the agents can verify this state but does not expose an administrative branch-protection write operation. Until the owner/admin enables and the Security B role reverifies these rules, **Security B remains FAIL — MEDIUM**.

No code or workflow workaround is considered equivalent to repository-enforced branch protection.
