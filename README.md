# Linux Desktop Customizer

> **Created by one. Improved by many. Available to all.**

![Status](https://img.shields.io/badge/status-stabilization-blue)
![Next](https://img.shields.io/badge/next-real%20KDE%20validation-orange)
![Safety](https://img.shields.io/badge/repository-read--only%20CI-blue)

**Linux Desktop Customizer** is a privacy-first open-source project for personalizing the Linux desktop from one place: themes, wallpapers, icons, cursors, windows, panels, effects, sounds and other desktop appearance settings.

The product direction remains simple: **preview → review → apply → undo**.

> [!IMPORTANT]
> The clean rebuild foundation is complete. The project is currently under a stabilization freeze before feature expansion resumes. The original complete historical source was not recoverable; preserved fragments remain evidence only and are not presented as reconstructed source.

## Repository state

- **`integration/canonical`** — sole active integration/development truth.
- **`main`** — stable/release line; no direct development or force-push.
- Legacy publisher workflows and trigger scaffolding have been removed from the current canonical tree.
- Historical recovery fragments are preserved only as explicit evidence under `recovery-evidence/`, `source-package/`, and Git history.

See [Project status](docs/PROJECT_STATUS.md), [branch governance](docs/BRANCH_GOVERNANCE.md) and the historical [consolidation snapshot](docs/CONSOLIDATION_STATUS.md).

## What happened to the original source?

Every surviving GitHub-only recovery route was checked:

- the historical 56-object archive is missing one authoritative Git blob;
- the preserved ten-part XZ stream is truncated;
- the older three-part `source.part.*` archive fails its pinned SHA-256;
- accessible Git history contains no later `source.part.03+` and only one readable import patch part.

There is no separate PC or external authoritative copy. The failures are documented in [Recovery evidence](docs/RECOVERY_EVIDENCE.md).

## Current phase

Issue #40 (rebuild foundation) is complete and closed. Current coordination and final-readiness work is tracked in **Issue #47**, while **Issue #6** remains the product/UX contract.

Before new features resume, the project is solidifying the current canonical state:

1. align capability/UI reporting with the integrated runtime;
2. keep profiles, transactions, filesystem and KDE boundaries security-reviewed;
3. keep CI/publication verification aligned with the reconstructed repository;
4. remove verified-obsolete active-tree material and stale topic refs;
5. prove the KDE flow on a real Plasma session;
6. complete graphical apply/history/undo and distribution only after stabilization is green.

## Resumo em português

O LDC já tem a fundação reconstruída: CLI, interface gráfica local, perfis declarativos seguros, transações com rollback, integração KDE limitada ao esquema de cores, EN + pt-PT, instalação por utilizador e CI com testes/vet/race.

Ainda **não é uma versão final**. Falta provar o percurso completo num Plasma real, completar o fluxo gráfico de aplicar/desfazer, criar e validar pacotes DEB/RPM, testar instalação limpa em Linux, obter screenshots reais e fechar a proteção administrativa de `integration/canonical`.

## Product direction

The preserved product direction includes:

- one local interface for Linux desktop personalization;
- original signature themes;
- Creator Forge for complete theme creation;
- Visual Composer for windows, panels, docks, widgets, typography and wallpapers;
- Effects Composer for advanced visual effects;
- Harmony Guard for coherent palette/visual choices;
- Adapter Forge for supported desktop/compositor backends;
- reviewed changes before apply;
- transaction history with rollback/undo;
- local-first operation with no telemetry requirement;
- defensive handling of downloaded content.

These are requirements, not claims that the current repository already ships a working build.

## Installation and graphical preview

There is currently **no supported release build**. The graphical runtime is still read-only for mutation, while the CLI/backend already contains a reviewed, scoped KDE color-scheme path with preview/apply/rollback primitives. That backend path still requires real Plasma end-to-end validation before the graphical mutation flow is enabled.

For bilingual English/Portuguese build and safe loopback launch instructions, see [Graphical UI onboarding](docs/GETTING_STARTED_UI.md). Real project screenshots must follow the [real Linux screenshot evidence procedure](docs/REAL_LINUX_SCREENSHOTS.md); mockups or generated substitutes do not satisfy the release gate.

Packaging/release remains gated on real Plasma validation, complete graphical flow, DEB/RPM and trusted terminal distribution, clean Linux install/run/uninstall, real screenshots, final security/supply-chain review, branch protection and explicit publication authorization.

## Safety and privacy

Linux Desktop Customizer is intended to keep personalization reversible and inspectable.

- No arbitrary theme install scripts should execute automatically.
- Downloaded content should be inspected before use.
- Unsafe archive paths and path traversal must be rejected.
- Changes should be previewed/reviewed before apply and recorded for rollback.
- Diagnostics should avoid leaking secrets.
- Telemetry is not required for the application to function.
- CI should use least privilege and immutable Action pins.

See [Security testing policy](SECURITY_TESTING_POLICY.md) and [Security baseline](docs/SECURITY_BASELINE.md).

## Contributing

New implementation work starts from **`integration/canonical`** and uses a short-lived topic branch with visible validation. During the stabilization freeze, only correctness/security/CI/docs/cleanup work should land. Do not write or force-push directly to `main` or bypass review on canonical.

Security-sensitive reports must not include credentials, private files or secrets.

## Credits

**Created by JP99** — concept, product direction and testing.

AI engineering assistance has been used during development, repository consolidation and reconstruction.
