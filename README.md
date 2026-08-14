# Linux Desktop Customizer

> **Created by one. Improved by many. Available to all.**

[![CI](https://github.com/Iverysterog1/Linux-Desktop-Customizer/actions/workflows/ci.yml/badge.svg)](https://github.com/Iverysterog1/Linux-Desktop-Customizer/actions/workflows/ci.yml)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE)
[![Version](https://img.shields.io/badge/version-0.9.0--visual--composer-informational)](VERSION)

**Linux Desktop Customizer** is a privacy-first open-source application for changing the whole Linux desktop from one place: themes, wallpapers, icons, cursors, windows, panels, effects, sounds, login/boot artwork and more.

The project focuses on a simple workflow: **preview → review → use → undo**. Changes are validated, journalled and reversible, downloads are quarantined, and arbitrary theme scripts are never executed automatically.

## Signature themes

The repository includes the source previews for six original JP99 signature themes. Their full passive theme payloads are generated reproducibly from code-owned design tokens.

| Fire | Water | Wind |
|---|---|---|
| ![Fire](assets/theme-previews/fire.png) | ![Water](assets/theme-previews/water.png) | ![Wind](assets/theme-previews/wind.png) |

| Earth | Mac Inspired | Windows Inspired |
|---|---|---|
| ![Earth](assets/theme-previews/earth.png) | ![Mac Inspired](assets/theme-previews/mac-inspired.png) | ![Windows Inspired](assets/theme-previews/windows-inspired.png) |

The “Mac Inspired” and “Windows Inspired” themes are original, unaffiliated designs inspired by broad desktop design language; they do not include vendor-owned assets.

## What it includes

- One local UI for desktop personalization.
- Six JP99 signature themes: **Fire, Water, Wind, Earth, Mac Inspired and Windows Inspired**.
- **Creator Forge** for building complete themes from a visual direction instead of hand-editing configuration files.
- **Visual Composer** for shaping windows, panels, docks, widgets, typography, wallpaper and effect intent from a visual preview.
- **Harmony Guard** to keep palettes and visual choices coherent.
- **Adapter Forge** to preview how a composition maps to supported compositor/desktop backends.
- Safe theme acquisition, inspection, quarantine and provenance handling.
- Reviewed plans before apply.
- Transaction journal and rollback/undo support.
- Diagnostics designed to avoid leaking secrets.
- Local-first operation with no telemetry, analytics, device ID or automatic crash reporting.

## Safety model

Linux Desktop Customizer deliberately separates **passive theme assets** from active code.

- Theme packages do not get to run arbitrary install scripts automatically.
- Downloaded content is inspected before use.
- Unsafe archive entries and path escapes are rejected.
- Changes are reviewed before apply and recorded for rollback.
- Official theme payloads are generated from repository-owned source definitions.
- The application does not require telemetry to function.

See [SECURITY.md](SECURITY.md), [docs/PRIVACY.md](docs/PRIVACY.md), [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) and [docs/VALIDATION.md](docs/VALIDATION.md) for the detailed model and current validation status.

## Current validation status

Version `0.9.0-visual-composer` has automated Go tests covering the application core, update path, transactions, manifests, planning, UI, remote content handling, diagnostics and signature themes.

The public repository CI rebuilds the six official themes before executing:

```text
go test ./...
go vet ./...
go test -race ./...
git diff --check
go build ./cmd/ltc
go build ./cmd/ltc-ui
```

Real graphical validation across the full Linux desktop matrix is still an explicit work item. A green CI build must not be interpreted as proof that every compositor/desktop combination has been tested on real hardware.

See [docs/VALIDATION.md](docs/VALIDATION.md) for the exact evidence and remaining gaps.

## Build from source

### Requirements

- Go 1.23+
- Python 3.10+
- Python packages listed in `requirements-assets.txt`

### Prepare generated theme assets

```bash
python3 -m pip install -r requirements-assets.txt
python3 scripts/generate_signature_themes.py
```

Or:

```bash
make assets
```

### Test

```bash
make test
make vet
```

### Build

```bash
make build
```

The binaries are written to:

```text
bin/ltc
bin/ltc-ui
```

## Why generated theme payloads are not committed

The complete `themes/official/` tree contains hundreds of generated cursor, sound, wallpaper and theme files. Keeping those outputs in every Git commit would make the repository substantially larger without adding source information.

Instead, the repository stores:

- the generator: `scripts/generate_signature_themes.py`;
- six approved source previews under `assets/theme-previews/`;
- the code-owned design tokens and generation logic.

A clean regeneration was checked against the supplied `0.9.0-visual-composer` source package: all **379 generated files** were recreated with the same paths and SHA-256 hashes.

GitHub Actions regenerates the payload before tests and before packaging Linux releases, so published release archives still contain the complete `themes/official/` directory.

## Linux releases

The `Linux builds` GitHub Actions workflow builds Linux `amd64` and `arm64` archives. Tags matching `v*` trigger testing, cross-compilation, SHA-256 generation and publication of GitHub Release assets.

A release archive contains:

- `ltc`
- `ltc-ui`
- per-user installer/uninstaller
- README and GPL license
- all six regenerated official theme payloads

## Documentation

Useful project documents:

- [Architecture](docs/ARCHITECTURE.md)
- [Visual Composer](docs/VISUAL_COMPOSER.md)
- [Creator Forge](docs/CREATOR_FORGE.md)
- [Adapter Forge](docs/ADAPTER_FORGE.md)
- [Theme format](docs/THEME_FORMAT.md)
- [Personalization model](docs/PERSONALIZATION.md)
- [Updates](docs/UPDATES.md)
- [Validation](docs/VALIDATION.md)
- [Roadmap](docs/ROADMAP.md)
- [Community](docs/COMMUNITY.md)
- [Publishing](docs/PUBLISHING.md)

## Contributing

Contributions are welcome. Please read [CONTRIBUTING.md](CONTRIBUTING.md) before submitting changes, and use the repository issue tracker for reproducible bugs or design discussions.

Security-sensitive reports should follow [SECURITY.md](SECURITY.md) instead of being posted publicly.

## Licensing

- **Software code:** GNU General Public License v3.0 — see [LICENSE](LICENSE).
- **Official JP99 theme assets:** CC BY-SA 4.0 as declared by their generated manifests.

Third-party/community themes retain their own upstream licensing and provenance and are not silently relicensed by this project.

## Credits

**Created by JP99** — concept, product direction and testing.

**GPT-5.6 Sol by OpenAI** — AI engineering collaboration.

See [CREDITS.md](CREDITS.md) for additional project credits and attribution details.
