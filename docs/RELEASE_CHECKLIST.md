# Release readiness checklist

A green build alone is not FINAL READY. Every applicable gate below needs evidence on the intended release candidate.

## Source / governance

- [x] Clean rebuilt Go foundation exists on `integration/canonical`.
- [x] CLI/UI and safe transaction/profile boundaries exist.
- [x] `integration/canonical` is protected: PR required, required checks enforced, force-push/delete blocked.
- [x] Legacy write/force-push/tag workflows have been removed from `main` through PR #75.
- [x] A temporary canonical-to-main validation PR proved `Canonical CI`, `Race verification`, `Workflow write guard`, and `Publication verification` all PASS without merging product code.
- [x] `main` has final protected-PR rules requiring those four checks.
- [x] Superseded topic branches are cleaned up; only `main` and `integration/canonical` remain, with automatic merged-head deletion enabled.
- [ ] Final promotion to `main` occurs only through a protected PR.
- [ ] Target version/changelog/release notes are internally consistent.

## Tests and security

- [ ] Unit/integration tests PASS on the exact release candidate.
- [ ] `go vet ./...` PASS.
- [ ] `go test -race ./...` PASS.
- [ ] CLI and UI builds PASS.
- [ ] Publication verification gate PASS on the release candidate.
- [x] Security A has no unresolved stabilization blocker on the current canonical boundary.
- [x] Security B repository-governance/supply-chain stabilization review PASS; re-run on the final release candidate.
- [ ] No credentials, tokens, private paths or user data in source/artifacts.

## KDE product validation

- [ ] Real Plasma session: detect -> preview -> snapshot -> apply -> effective/visible validation -> rollback/unapply PASS.
- [ ] Wayland representative validation PASS.
- [ ] X11 representative validation PASS where supported.
- [ ] Fractional-scaling compatibility state is explicit.
- [ ] Unsupported/unknown states never silently become PASS.
- [ ] KDE-008 is implemented before any monitor-specific layout/panel/widget/wallpaper restore ships.

## Graphical UX / localization

- [ ] Real graphical preview/review/apply/history/undo flow is usable.
- [ ] English shipped flow is complete.
- [ ] Portuguese (pt-PT) shipped flow is complete.
- [ ] Keyboard, reduced-motion, readability/high-contrast requirements are verified.
- [ ] Capability reporting matches actual runtime support.

## Distribution

- [x] User-local installer/uninstaller baseline tests PASS.
- [ ] DEB package built and validated.
- [ ] RPM package built and validated.
- [ ] Trusted/versioned terminal installer uses the same release artifacts and verifies integrity.
- [ ] Clean Linux install -> launch -> use -> uninstall PASS.
- [ ] Ordinary uninstall preserves user data/recovery state.
- [ ] Package/release artifact checksums (and signing where adopted) are verified.
- [ ] Reproducibility/equivalence expectations are documented.

## Licensing / provenance

- [x] The owner selected 0BSD and a canonical `LICENSE` file is present.
- [ ] Project-owned/generated assets have documented licensing.
- [ ] Third-party assets/components have provenance and compatible licensing.

## Documentation / presentation

- [ ] README reflects the actual shipped state.
- [ ] User-facing documentation is available in EN + pt-PT for shipped flows.
- [ ] Supported distributions/desktops and limitations are explicit.
- [ ] Real screenshots from the validated application are committed and shown on GitHub.
- [ ] Screenshot evidence contains no private/sensitive user data.

## Publication

A release/tag/package publication happens only after every mandatory gate is PASS and the owner gives explicit final publication authorization.
