# Release readiness checklist

A release is ready only when every required item below has evidence. A branch name, green documentation check, or successfully created archive is not sufficient by itself.

## Source and version

- [ ] Canonical source tree is reconstructed from preserved material.
- [ ] Reconstruction provenance and SHA-256 are recorded.
- [ ] Internal `VERSION` agrees with package metadata, application metadata and release notes.
- [ ] Changelog contains the target version and only verified changes.
- [ ] Working tree used for packaging is clean and tied to one reviewed commit SHA.

## Tests and CI

- [ ] Unit tests pass from a clean checkout.
- [ ] `go vet ./...` passes.
- [ ] Supported race tests pass.
- [ ] `git diff --check` passes.
- [ ] CLI and graphical UI targets build from a clean checkout.
- [ ] CI does not hide failures with unconditional `|| true` on release-critical validation.

## Packaging and reproducibility

- [ ] Installer and uninstaller pass syntax and functional smoke tests.
- [ ] Package contains the expected binaries, license, documentation and required theme assets.
- [ ] Package file list is recorded.
- [ ] SHA-256 is generated for every distributable artifact.
- [ ] A second clean build produces equivalent release contents; any unavoidable nondeterminism is documented.
- [ ] No release job force-pushes or writes directly to `main`.
- [ ] Publishing/tagging remains a separate reviewed/manual action.

## Desktop integration

- [ ] `.desktop` entry validates and launches the installed graphical binary.
- [ ] Application name, executable, icon and categories are consistent.
- [ ] Icon is installed at appropriate freedesktop sizes/paths or as a valid scalable icon.
- [ ] AppStream/metainfo is validated if shipped.
- [ ] Application appears correctly in a representative Linux application menu after installation.
- [ ] Uninstall removes only files owned by the application.

## Licensing and privacy

- [ ] Software license is present in the source and release archive.
- [ ] Generated/project-owned theme asset licensing is documented.
- [ ] Third-party assets have provenance and compatible licensing.
- [ ] No credentials, tokens, private paths or user data are present in source/package/artifacts.
- [ ] Network/update behavior matches the documented privacy model.

## User experience

- [ ] Primary path is `download → install → application menu → launch`.
- [ ] Terminal fallback is documented.
- [ ] Supported distributions/desktops and known limitations are stated.
- [ ] Real screenshots come from the validated build; no fabricated UI is presented as the application.
- [ ] Upgrade and rollback expectations are documented.

## Release gate

Do not create a public release or version tag until all mandatory checks above have evidence attached to the release PR or linked validation records.
